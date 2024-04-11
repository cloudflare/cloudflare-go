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
func (r *RulesetService) List(ctx context.Context, query RulesetListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Ruleset], err error) {
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
	path := fmt.Sprintf("%s/%s/rulesets", accountOrZone, accountOrZoneID)
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

// Fetches all rulesets.
func (r *RulesetService) ListAutoPaging(ctx context.Context, query RulesetListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Ruleset] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type Ruleset struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind RulesetKind `json:"kind"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase RulesetPhase `json:"phase"`
	JSON  rulesetJSON  `json:"-"`
}

// rulesetJSON contains the JSON metadata for the struct [Ruleset]
type rulesetJSON struct {
	ID          apijson.Field
	LastUpdated apijson.Field
	Version     apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Ruleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RulesetKind string

const (
	RulesetKindManaged RulesetKind = "managed"
	RulesetKindCustom  RulesetKind = "custom"
	RulesetKindRoot    RulesetKind = "root"
	RulesetKindZone    RulesetKind = "zone"
)

func (r RulesetKind) IsKnown() bool {
	switch r {
	case RulesetKindManaged, RulesetKindCustom, RulesetKindRoot, RulesetKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type RulesetPhase string

const (
	RulesetPhaseDDoSL4                         RulesetPhase = "ddos_l4"
	RulesetPhaseDDoSL7                         RulesetPhase = "ddos_l7"
	RulesetPhaseHTTPConfigSettings             RulesetPhase = "http_config_settings"
	RulesetPhaseHTTPCustomErrors               RulesetPhase = "http_custom_errors"
	RulesetPhaseHTTPLogCustomFields            RulesetPhase = "http_log_custom_fields"
	RulesetPhaseHTTPRatelimit                  RulesetPhase = "http_ratelimit"
	RulesetPhaseHTTPRequestCacheSettings       RulesetPhase = "http_request_cache_settings"
	RulesetPhaseHTTPRequestDynamicRedirect     RulesetPhase = "http_request_dynamic_redirect"
	RulesetPhaseHTTPRequestFirewallCustom      RulesetPhase = "http_request_firewall_custom"
	RulesetPhaseHTTPRequestFirewallManaged     RulesetPhase = "http_request_firewall_managed"
	RulesetPhaseHTTPRequestLateTransform       RulesetPhase = "http_request_late_transform"
	RulesetPhaseHTTPRequestOrigin              RulesetPhase = "http_request_origin"
	RulesetPhaseHTTPRequestRedirect            RulesetPhase = "http_request_redirect"
	RulesetPhaseHTTPRequestSanitize            RulesetPhase = "http_request_sanitize"
	RulesetPhaseHTTPRequestSbfm                RulesetPhase = "http_request_sbfm"
	RulesetPhaseHTTPRequestSelectConfiguration RulesetPhase = "http_request_select_configuration"
	RulesetPhaseHTTPRequestTransform           RulesetPhase = "http_request_transform"
	RulesetPhaseHTTPResponseCompression        RulesetPhase = "http_response_compression"
	RulesetPhaseHTTPResponseFirewallManaged    RulesetPhase = "http_response_firewall_managed"
	RulesetPhaseHTTPResponseHeadersTransform   RulesetPhase = "http_response_headers_transform"
	RulesetPhaseMagicTransit                   RulesetPhase = "magic_transit"
	RulesetPhaseMagicTransitIDsManaged         RulesetPhase = "magic_transit_ids_managed"
	RulesetPhaseMagicTransitManaged            RulesetPhase = "magic_transit_managed"
)

func (r RulesetPhase) IsKnown() bool {
	switch r {
	case RulesetPhaseDDoSL4, RulesetPhaseDDoSL7, RulesetPhaseHTTPConfigSettings, RulesetPhaseHTTPCustomErrors, RulesetPhaseHTTPLogCustomFields, RulesetPhaseHTTPRatelimit, RulesetPhaseHTTPRequestCacheSettings, RulesetPhaseHTTPRequestDynamicRedirect, RulesetPhaseHTTPRequestFirewallCustom, RulesetPhaseHTTPRequestFirewallManaged, RulesetPhaseHTTPRequestLateTransform, RulesetPhaseHTTPRequestOrigin, RulesetPhaseHTTPRequestRedirect, RulesetPhaseHTTPRequestSanitize, RulesetPhaseHTTPRequestSbfm, RulesetPhaseHTTPRequestSelectConfiguration, RulesetPhaseHTTPRequestTransform, RulesetPhaseHTTPResponseCompression, RulesetPhaseHTTPResponseFirewallManaged, RulesetPhaseHTTPResponseHeadersTransform, RulesetPhaseMagicTransit, RulesetPhaseMagicTransitIDsManaged, RulesetPhaseMagicTransitManaged:
		return true
	}
	return false
}

// A ruleset object.
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

func (r RulesetNewResponseKind) IsKnown() bool {
	switch r {
	case RulesetNewResponseKindManaged, RulesetNewResponseKindCustom, RulesetNewResponseKindRoot, RulesetNewResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type RulesetNewResponsePhase string

const (
	RulesetNewResponsePhaseDDoSL4                         RulesetNewResponsePhase = "ddos_l4"
	RulesetNewResponsePhaseDDoSL7                         RulesetNewResponsePhase = "ddos_l7"
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

func (r RulesetNewResponsePhase) IsKnown() bool {
	switch r {
	case RulesetNewResponsePhaseDDoSL4, RulesetNewResponsePhaseDDoSL7, RulesetNewResponsePhaseHTTPConfigSettings, RulesetNewResponsePhaseHTTPCustomErrors, RulesetNewResponsePhaseHTTPLogCustomFields, RulesetNewResponsePhaseHTTPRatelimit, RulesetNewResponsePhaseHTTPRequestCacheSettings, RulesetNewResponsePhaseHTTPRequestDynamicRedirect, RulesetNewResponsePhaseHTTPRequestFirewallCustom, RulesetNewResponsePhaseHTTPRequestFirewallManaged, RulesetNewResponsePhaseHTTPRequestLateTransform, RulesetNewResponsePhaseHTTPRequestOrigin, RulesetNewResponsePhaseHTTPRequestRedirect, RulesetNewResponsePhaseHTTPRequestSanitize, RulesetNewResponsePhaseHTTPRequestSbfm, RulesetNewResponsePhaseHTTPRequestSelectConfiguration, RulesetNewResponsePhaseHTTPRequestTransform, RulesetNewResponsePhaseHTTPResponseCompression, RulesetNewResponsePhaseHTTPResponseFirewallManaged, RulesetNewResponsePhaseHTTPResponseHeadersTransform, RulesetNewResponsePhaseMagicTransit, RulesetNewResponsePhaseMagicTransitIDsManaged, RulesetNewResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type RulesetNewResponseRule struct {
	// The action to perform when the rule matches.
	Action           RulesetNewResponseRulesAction `json:"action"`
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
	JSON    rulesetNewResponseRuleJSON `json:"-"`
	union   RulesetNewResponseRulesUnion
}

// rulesetNewResponseRuleJSON contains the JSON metadata for the struct
// [RulesetNewResponseRule]
type rulesetNewResponseRuleJSON struct {
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

func (r rulesetNewResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetNewResponseRule) AsUnion() RulesetNewResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule],
// [rulesets.RulesetNewResponseRulesRulesetsChallengeRule],
// [rulesets.RulesetNewResponseRulesRulesetsCompressResponseRule],
// [rulesets.ExecuteRule],
// [rulesets.RulesetNewResponseRulesRulesetsJsChallengeRule], [rulesets.LogRule],
// [rulesets.RulesetNewResponseRulesRulesetsManagedChallengeRule],
// [rulesets.RulesetNewResponseRulesRulesetsRedirectRule],
// [rulesets.RulesetNewResponseRulesRulesetsRewriteRule],
// [rulesets.RulesetNewResponseRulesRulesetsRouteRule],
// [rulesets.RulesetNewResponseRulesRulesetsScoreRule],
// [rulesets.RulesetNewResponseRulesRulesetsServeErrorRule],
// [rulesets.RulesetNewResponseRulesRulesetsSetConfigRule], [rulesets.SkipRule] or
// [rulesets.RulesetNewResponseRulesRulesetsSetCacheSettingsRule].
type RulesetNewResponseRulesUnion interface {
	implementsRulesetsRulesetNewResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsCompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsJsChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsRedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsRewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsRouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsSetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetNewResponseRulesRulesetsSetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
	)
}

type RulesetNewResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsChallengeRuleAction `json:"action"`
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
	JSON rulesetNewResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [RulesetNewResponseRulesRulesetsChallengeRule]
type rulesetNewResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsChallengeRule) implementsRulesetsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsChallengeRuleAction string

const (
	RulesetNewResponseRulesRulesetsChallengeRuleActionChallenge RulesetNewResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetNewResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

type RulesetNewResponseRulesRulesetsCompressResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsCompressResponseRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsCompressResponseRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetNewResponseRulesRulesetsCompressResponseRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsCompressResponseRuleJSON contains the JSON
// metadata for the struct [RulesetNewResponseRulesRulesetsCompressResponseRule]
type rulesetNewResponseRulesRulesetsCompressResponseRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsCompressResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsCompressResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsCompressResponseRule) implementsRulesetsRulesetNewResponseRule() {
}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsCompressResponseRuleAction string

const (
	RulesetNewResponseRulesRulesetsCompressResponseRuleActionCompressResponse RulesetNewResponseRulesRulesetsCompressResponseRuleAction = "compress_response"
)

func (r RulesetNewResponseRulesRulesetsCompressResponseRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsCompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsCompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms []RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm `json:"algorithms"`
	JSON       rulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersJSON        `json:"-"`
}

// rulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsCompressResponseRuleActionParameters]
type rulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersJSON struct {
	Algorithms  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsCompressResponseRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Compression algorithm to enable.
type RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName `json:"name"`
	JSON rulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON  `json:"-"`
}

// rulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm]
type rulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON) RawJSON() string {
	return r.raw
}

// Name of compression algorithm to enable.
type RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName string

const (
	RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone    RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "none"
	RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto    RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "auto"
	RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "default"
	RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip    RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "gzip"
	RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli  RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "brotli"
)

func (r RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone, RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto, RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault, RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip, RulesetNewResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli:
		return true
	}
	return false
}

type RulesetNewResponseRulesRulesetsJsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsJsChallengeRuleAction `json:"action"`
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
	JSON rulesetNewResponseRulesRulesetsJsChallengeRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsJsChallengeRuleJSON contains the JSON metadata
// for the struct [RulesetNewResponseRulesRulesetsJsChallengeRule]
type rulesetNewResponseRulesRulesetsJsChallengeRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsJsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsJsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsJsChallengeRule) implementsRulesetsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsJsChallengeRuleAction string

const (
	RulesetNewResponseRulesRulesetsJsChallengeRuleActionJsChallenge RulesetNewResponseRulesRulesetsJsChallengeRuleAction = "js_challenge"
)

func (r RulesetNewResponseRulesRulesetsJsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsJsChallengeRuleActionJsChallenge:
		return true
	}
	return false
}

type RulesetNewResponseRulesRulesetsManagedChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsManagedChallengeRuleAction `json:"action"`
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
	JSON rulesetNewResponseRulesRulesetsManagedChallengeRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsManagedChallengeRuleJSON contains the JSON
// metadata for the struct [RulesetNewResponseRulesRulesetsManagedChallengeRule]
type rulesetNewResponseRulesRulesetsManagedChallengeRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsManagedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsManagedChallengeRule) implementsRulesetsRulesetNewResponseRule() {
}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsManagedChallengeRuleAction string

const (
	RulesetNewResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge RulesetNewResponseRulesRulesetsManagedChallengeRuleAction = "managed_challenge"
)

func (r RulesetNewResponseRulesRulesetsManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

type RulesetNewResponseRulesRulesetsRedirectRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsRedirectRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsRedirectRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetNewResponseRulesRulesetsRedirectRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRedirectRuleJSON contains the JSON metadata for
// the struct [RulesetNewResponseRulesRulesetsRedirectRule]
type rulesetNewResponseRulesRulesetsRedirectRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsRedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRedirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsRedirectRule) implementsRulesetsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsRedirectRuleAction string

const (
	RulesetNewResponseRulesRulesetsRedirectRuleActionRedirect RulesetNewResponseRulesRulesetsRedirectRuleAction = "redirect"
)

func (r RulesetNewResponseRulesRulesetsRedirectRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsRedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsRedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromList `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValue `json:"from_value"`
	JSON      rulesetNewResponseRulesRulesetsRedirectRuleActionParametersJSON      `json:"-"`
}

// rulesetNewResponseRulesRulesetsRedirectRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRedirectRuleActionParameters]
type rulesetNewResponseRulesRulesetsRedirectRuleActionParametersJSON struct {
	FromList    apijson.Field
	FromValue   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRedirectRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRedirectRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on a bulk list lookup.
type RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key string `json:"key"`
	// The name of the list to match against.
	Name string                                                                  `json:"name"`
	JSON rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromListJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromListJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromList]
type rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromListJSON struct {
	Key         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromListJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on the request properties.
type RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString bool `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL `json:"target_url"`
	JSON      rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON      `json:"-"`
}

// rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValue]
type rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON struct {
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	TargetURL           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON) RawJSON() string {
	return r.raw
}

// The status code to be used for the redirect.
type RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode float64

const (
	RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301 RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 301
	RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302 RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 302
	RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303 RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 303
	RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307 RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 307
	RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308 RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 308
)

func (r RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301, RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302, RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303, RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307, RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL struct {
	// The URL to redirect the request to.
	Value string `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                            `json:"expression"`
	JSON       rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON `json:"-"`
	union      RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion
}

// rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL]
type rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) AsUnion() RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion {
	return r.union
}

// The URL to redirect the request to.
//
// Union satisfied by
// [rulesets.RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
// or
// [rulesets.RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect].
type RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRulesetsRulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect{}),
		},
	)
}

type RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value string                                                                                             `json:"value"`
	JSON  rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
type rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRulesetsRulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                                              `json:"expression"`
	JSON       rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect]
type rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRulesetsRulesetNewResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type RulesetNewResponseRulesRulesetsRewriteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsRewriteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsRewriteRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetNewResponseRulesRulesetsRewriteRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRewriteRuleJSON contains the JSON metadata for
// the struct [RulesetNewResponseRulesRulesetsRewriteRule]
type rulesetNewResponseRulesRulesetsRewriteRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsRewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsRewriteRule) implementsRulesetsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsRewriteRuleAction string

const (
	RulesetNewResponseRulesRulesetsRewriteRuleActionRewrite RulesetNewResponseRulesRulesetsRewriteRuleAction = "rewrite"
)

func (r RulesetNewResponseRulesRulesetsRewriteRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsRewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsRewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers map[string]RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeader `json:"headers"`
	// URI to rewrite the request to.
	URI  RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURI  `json:"uri"`
	JSON rulesetNewResponseRulesRulesetsRewriteRuleActionParametersJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRewriteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetNewResponseRulesRulesetsRewriteRuleActionParameters]
type rulesetNewResponseRulesRulesetsRewriteRuleActionParametersJSON struct {
	Headers     apijson.Field
	URI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRewriteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Remove the header from the request.
type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeader struct {
	Operation RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation `json:"operation,required"`
	// Static value for the header.
	Value string `json:"value"`
	// Expression for the header value.
	Expression string                                                               `json:"expression"`
	JSON       rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON `json:"-"`
	union      RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion
}

// rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeader]
type rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeader) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeader) AsUnion() RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion {
	return r.union
}

// Remove the header from the request.
//
// Union satisfied by
// [rulesets.RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader],
// [rulesets.RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
// or
// [rulesets.RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader].
type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion interface {
	implementsRulesetsRulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeader()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader{}),
		},
	)
}

// Remove the header from the request.
type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation `json:"operation,required"`
	JSON      rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON      `json:"-"`
}

// rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader]
type rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON struct {
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) implementsRulesetsRulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Set a request header with a static value.
type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader struct {
	Operation RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation `json:"operation,required"`
	// Static value for the header.
	Value string                                                                            `json:"value,required"`
	JSON  rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
type rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) implementsRulesetsRulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation string

const (
	RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation = "set"
)

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader struct {
	// Expression for the header value.
	Expression string                                                                                  `json:"expression,required"`
	Operation  RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation `json:"operation,required"`
	JSON       rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON      `json:"-"`
}

// rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader]
type rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) implementsRulesetsRulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation string

const (
	RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation = "set"
)

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet:
		return true
	}
	return false
}

type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation string

const (
	RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "remove"
	RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet    RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "set"
)

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove, RulesetNewResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPath `json:"path"`
	// Query portion rewrite.
	Query RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQuery `json:"query"`
	JSON  rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIJSON  `json:"-"`
}

// rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIJSON contains the
// JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURI]
type rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIJSON struct {
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIJSON) RawJSON() string {
	return r.raw
}

// Path portion rewrite.
type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPath struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                `json:"expression"`
	JSON       rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON `json:"-"`
	union      RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion
}

// rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPath]
type rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPath) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPath) AsUnion() RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion {
	return r.union
}

// Path portion rewrite.
//
// Union satisfied by
// [rulesets.RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
// or
// [rulesets.RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue].
type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion interface {
	implementsRulesetsRulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPath()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue{}),
		},
	)
}

type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue struct {
	// Predefined replacement value.
	Value string                                                                           `json:"value,required"`
	JSON  rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
type rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) implementsRulesetsRulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                            `json:"expression,required"`
	JSON       rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue]
type rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) implementsRulesetsRulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

// Query portion rewrite.
type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQuery struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                 `json:"expression"`
	JSON       rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON `json:"-"`
	union      RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion
}

// rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQuery]
type rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQuery) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQuery) AsUnion() RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion {
	return r.union
}

// Query portion rewrite.
//
// Union satisfied by
// [rulesets.RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
// or
// [rulesets.RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue].
type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion interface {
	implementsRulesetsRulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQuery()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue{}),
		},
	)
}

type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue struct {
	// Predefined replacement value.
	Value string                                                                            `json:"value,required"`
	JSON  rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
type rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) implementsRulesetsRulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                             `json:"expression,required"`
	JSON       rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue]
type rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) implementsRulesetsRulesetNewResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type RulesetNewResponseRulesRulesetsRouteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsRouteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsRouteRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetNewResponseRulesRulesetsRouteRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRouteRuleJSON contains the JSON metadata for the
// struct [RulesetNewResponseRulesRulesetsRouteRule]
type rulesetNewResponseRulesRulesetsRouteRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsRouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRouteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsRouteRule) implementsRulesetsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsRouteRuleAction string

const (
	RulesetNewResponseRulesRulesetsRouteRuleActionRoute RulesetNewResponseRulesRulesetsRouteRuleAction = "route"
)

func (r RulesetNewResponseRulesRulesetsRouteRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsRouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsRouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader string `json:"host_header"`
	// Override the IP/TCP destination.
	Origin RulesetNewResponseRulesRulesetsRouteRuleActionParametersOrigin `json:"origin"`
	// Override the Server Name Indication (SNI).
	Sni  RulesetNewResponseRulesRulesetsRouteRuleActionParametersSni  `json:"sni"`
	JSON rulesetNewResponseRulesRulesetsRouteRuleActionParametersJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRouteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetNewResponseRulesRulesetsRouteRuleActionParameters]
type rulesetNewResponseRulesRulesetsRouteRuleActionParametersJSON struct {
	HostHeader  apijson.Field
	Origin      apijson.Field
	Sni         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRouteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRouteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Override the IP/TCP destination.
type RulesetNewResponseRulesRulesetsRouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host string `json:"host"`
	// Override the destination port.
	Port float64                                                            `json:"port"`
	JSON rulesetNewResponseRulesRulesetsRouteRuleActionParametersOriginJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRouteRuleActionParametersOriginJSON contains the
// JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRouteRuleActionParametersOrigin]
type rulesetNewResponseRulesRulesetsRouteRuleActionParametersOriginJSON struct {
	Host        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRouteRuleActionParametersOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRouteRuleActionParametersOriginJSON) RawJSON() string {
	return r.raw
}

// Override the Server Name Indication (SNI).
type RulesetNewResponseRulesRulesetsRouteRuleActionParametersSni struct {
	// The SNI override.
	Value string                                                          `json:"value,required"`
	JSON  rulesetNewResponseRulesRulesetsRouteRuleActionParametersSniJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsRouteRuleActionParametersSniJSON contains the
// JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsRouteRuleActionParametersSni]
type rulesetNewResponseRulesRulesetsRouteRuleActionParametersSniJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsRouteRuleActionParametersSni) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsRouteRuleActionParametersSniJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseRulesRulesetsScoreRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsScoreRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsScoreRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetNewResponseRulesRulesetsScoreRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsScoreRuleJSON contains the JSON metadata for the
// struct [RulesetNewResponseRulesRulesetsScoreRule]
type rulesetNewResponseRulesRulesetsScoreRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsScoreRule) implementsRulesetsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsScoreRuleAction string

const (
	RulesetNewResponseRulesRulesetsScoreRuleActionScore RulesetNewResponseRulesRulesetsScoreRuleAction = "score"
)

func (r RulesetNewResponseRulesRulesetsScoreRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment int64                                                        `json:"increment"`
	JSON      rulesetNewResponseRulesRulesetsScoreRuleActionParametersJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsScoreRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetNewResponseRulesRulesetsScoreRuleActionParameters]
type rulesetNewResponseRulesRulesetsScoreRuleActionParametersJSON struct {
	Increment   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsScoreRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsScoreRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

type RulesetNewResponseRulesRulesetsServeErrorRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsServeErrorRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsServeErrorRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetNewResponseRulesRulesetsServeErrorRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsServeErrorRuleJSON contains the JSON metadata for
// the struct [RulesetNewResponseRulesRulesetsServeErrorRule]
type rulesetNewResponseRulesRulesetsServeErrorRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsServeErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsServeErrorRule) implementsRulesetsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsServeErrorRuleAction string

const (
	RulesetNewResponseRulesRulesetsServeErrorRuleActionServeError RulesetNewResponseRulesRulesetsServeErrorRuleAction = "serve_error"
)

func (r RulesetNewResponseRulesRulesetsServeErrorRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsServeErrorRuleActionParameters struct {
	// Error response content.
	Content string `json:"content"`
	// Content-type header to set with the response.
	ContentType RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentType `json:"content_type"`
	// The status code to use for the error.
	StatusCode float64                                                           `json:"status_code"`
	JSON       rulesetNewResponseRulesRulesetsServeErrorRuleActionParametersJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsServeErrorRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsServeErrorRuleActionParameters]
type rulesetNewResponseRulesRulesetsServeErrorRuleActionParametersJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsServeErrorRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsServeErrorRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Content-type header to set with the response.
type RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentType string

const (
	RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentType = "application/json"
	RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml         RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/xml"
	RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain       RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/plain"
	RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML        RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/html"
)

func (r RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson, RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml, RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain, RulesetNewResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

type RulesetNewResponseRulesRulesetsSetConfigRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsSetConfigRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsSetConfigRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetNewResponseRulesRulesetsSetConfigRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetConfigRuleJSON contains the JSON metadata for
// the struct [RulesetNewResponseRulesRulesetsSetConfigRule]
type rulesetNewResponseRulesRulesetsSetConfigRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsSetConfigRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetConfigRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetConfigRule) implementsRulesetsRulesetNewResponseRule() {}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsSetConfigRuleAction string

const (
	RulesetNewResponseRulesRulesetsSetConfigRuleActionSetConfig RulesetNewResponseRulesRulesetsSetConfigRuleAction = "set_config"
)

func (r RulesetNewResponseRulesRulesetsSetConfigRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsSetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites bool `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersAutominify `json:"autominify"`
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
	Polish RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersPolish `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader bool `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes bool `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSL `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg  bool                                                             `json:"sxg"`
	JSON rulesetNewResponseRulesRulesetsSetConfigRuleActionParametersJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetConfigRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetConfigRuleActionParameters]
type rulesetNewResponseRulesRulesetsSetConfigRuleActionParametersJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsSetConfigRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetConfigRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Select which file extensions to minify automatically.
type RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	Css bool `json:"css"`
	// Minify HTML files.
	HTML bool `json:"html"`
	// Minify JS files.
	Js   bool                                                                       `json:"js"`
	JSON rulesetNewResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersAutominify]
type rulesetNewResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersAutominify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON) RawJSON() string {
	return r.raw
}

// Configure the Polish level.
type RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersPolish string

const (
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersPolishOff      RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersPolish = "off"
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossless"
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy    RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossy"
)

func (r RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersPolishOff, RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless, RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy:
		return true
	}
	return false
}

// Configure the Security Level.
type RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel string

const (
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff            RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "off"
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "essentially_off"
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow            RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "low"
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium         RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "medium"
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh           RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "high"
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack    RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff, RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff, RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow, RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium, RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh, RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSL string

const (
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSLOff        RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSL = "off"
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible   RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSL = "flexible"
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSLFull       RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSL = "full"
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict     RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSL = "strict"
	RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSL = "origin_pull"
)

func (r RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSLOff, RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible, RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSLFull, RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict, RulesetNewResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

type RulesetNewResponseRulesRulesetsSetCacheSettingsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetNewResponseRulesRulesetsSetCacheSettingsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetNewResponseRulesRulesetsSetCacheSettingsRuleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleJSON contains the JSON
// metadata for the struct [RulesetNewResponseRulesRulesetsSetCacheSettingsRule]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetNewResponseRulesRulesetsSetCacheSettingsRule) implementsRulesetsRulesetNewResponseRule() {
}

// The action to perform when the rule matches.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleAction string

const (
	RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings RulesetNewResponseRulesRulesetsSetCacheSettingsRuleAction = "set_cache_settings"
)

func (r RulesetNewResponseRulesRulesetsSetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts []int64 `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache bool `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL `json:"edge_ttl"`
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
	ServeStale RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale `json:"serve_stale"`
	JSON       rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON       `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParameters]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON struct {
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

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64                                                                             `json:"default"`
	JSON    rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON struct {
	Mode        apijson.Field
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON) RawJSON() string {
	return r.raw
}

// Determines which browser ttl mode to use.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
)

func (r RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType bool `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor bool `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder bool                                                                            `json:"ignore_query_strings_order"`
	JSON                    rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON struct {
	CacheByDeviceType       apijson.Field
	CacheDeceptionArmor     apijson.Field
	CustomKey               apijson.Field
	IgnoreQueryStringsOrder apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON) RawJSON() string {
	return r.raw
}

// Customize which components of the request are included or excluded from the
// cache key.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser `json:"user"`
	JSON rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON) RawJSON() string {
	return r.raw
}

// The cookies to include in building the cache key.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Include these cookies' names and their values.
	Include []string                                                                                       `json:"include"`
	JSON    rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON) RawJSON() string {
	return r.raw
}

// The header names and values to include in building the cache key.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin bool `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include []string                                                                                       `json:"include"`
	JSON    rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON struct {
	CheckPresence apijson.Field
	ExcludeOrigin apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON) RawJSON() string {
	return r.raw
}

// Whether to use the original host or the resolved host in the cache key.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved bool                                                                                         `json:"resolved"`
	JSON     rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON) RawJSON() string {
	return r.raw
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude `json:"include"`
	JSON    rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON    `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON) RawJSON() string {
	return r.raw
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Exclude all query string parameters from use in building the cache key.
	All bool `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List []string                                                                                                   `json:"list"`
	JSON rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON) RawJSON() string {
	return r.raw
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Use all query string parameters in the cache key.
	All bool `json:"all"`
	// A list of query string parameters used to build the cache key.
	List []string                                                                                                   `json:"list"`
	JSON rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON) RawJSON() string {
	return r.raw
}

// Characteristics of the request user agent used in building the cache key.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType bool `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo bool `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang bool                                                                                         `json:"lang"`
	JSON rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON) RawJSON() string {
	return r.raw
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible bool `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize int64                                                                               `json:"min_file_size,required"`
	JSON        rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON struct {
	Eligible    apijson.Field
	MinFileSize apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON) RawJSON() string {
	return r.raw
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64 `json:"default,required"`
	// edge ttl options
	Mode RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL []RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL `json:"status_code_ttl,required"`
	JSON          rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON            `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON struct {
	Default       apijson.Field
	Mode          apijson.Field
	StatusCodeTTL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON) RawJSON() string {
	return r.raw
}

// edge ttl options
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value int64 `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue int64                                                                                       `json:"status_code_value"`
	JSON            rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON struct {
	Value           apijson.Field
	StatusCodeRange apijson.Field
	StatusCodeValue apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON) RawJSON() string {
	return r.raw
}

// The range of status codes used to apply the selected mode.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// response status code lower bound
	From int64 `json:"from,required"`
	// response status code upper bound
	To   int64                                                                                                      `json:"to,required"`
	JSON rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON) RawJSON() string {
	return r.raw
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating bool                                                                              `json:"disable_stale_while_updating,required"`
	JSON                      rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON `json:"-"`
}

// rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON
// contains the JSON metadata for the struct
// [RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale]
type rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON struct {
	DisableStaleWhileUpdating apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *RulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetNewResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetNewResponseRulesAction string

const (
	RulesetNewResponseRulesActionBlock            RulesetNewResponseRulesAction = "block"
	RulesetNewResponseRulesActionChallenge        RulesetNewResponseRulesAction = "challenge"
	RulesetNewResponseRulesActionCompressResponse RulesetNewResponseRulesAction = "compress_response"
	RulesetNewResponseRulesActionExecute          RulesetNewResponseRulesAction = "execute"
	RulesetNewResponseRulesActionJsChallenge      RulesetNewResponseRulesAction = "js_challenge"
	RulesetNewResponseRulesActionLog              RulesetNewResponseRulesAction = "log"
	RulesetNewResponseRulesActionManagedChallenge RulesetNewResponseRulesAction = "managed_challenge"
	RulesetNewResponseRulesActionRedirect         RulesetNewResponseRulesAction = "redirect"
	RulesetNewResponseRulesActionRewrite          RulesetNewResponseRulesAction = "rewrite"
	RulesetNewResponseRulesActionRoute            RulesetNewResponseRulesAction = "route"
	RulesetNewResponseRulesActionScore            RulesetNewResponseRulesAction = "score"
	RulesetNewResponseRulesActionServeError       RulesetNewResponseRulesAction = "serve_error"
	RulesetNewResponseRulesActionSetConfig        RulesetNewResponseRulesAction = "set_config"
	RulesetNewResponseRulesActionSkip             RulesetNewResponseRulesAction = "skip"
	RulesetNewResponseRulesActionSetCacheSettings RulesetNewResponseRulesAction = "set_cache_settings"
)

func (r RulesetNewResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetNewResponseRulesActionBlock, RulesetNewResponseRulesActionChallenge, RulesetNewResponseRulesActionCompressResponse, RulesetNewResponseRulesActionExecute, RulesetNewResponseRulesActionJsChallenge, RulesetNewResponseRulesActionLog, RulesetNewResponseRulesActionManagedChallenge, RulesetNewResponseRulesActionRedirect, RulesetNewResponseRulesActionRewrite, RulesetNewResponseRulesActionRoute, RulesetNewResponseRulesActionScore, RulesetNewResponseRulesActionServeError, RulesetNewResponseRulesActionSetConfig, RulesetNewResponseRulesActionSkip, RulesetNewResponseRulesActionSetCacheSettings:
		return true
	}
	return false
}

// A ruleset object.
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

func (r RulesetUpdateResponseKind) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseKindManaged, RulesetUpdateResponseKindCustom, RulesetUpdateResponseKindRoot, RulesetUpdateResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type RulesetUpdateResponsePhase string

const (
	RulesetUpdateResponsePhaseDDoSL4                         RulesetUpdateResponsePhase = "ddos_l4"
	RulesetUpdateResponsePhaseDDoSL7                         RulesetUpdateResponsePhase = "ddos_l7"
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

func (r RulesetUpdateResponsePhase) IsKnown() bool {
	switch r {
	case RulesetUpdateResponsePhaseDDoSL4, RulesetUpdateResponsePhaseDDoSL7, RulesetUpdateResponsePhaseHTTPConfigSettings, RulesetUpdateResponsePhaseHTTPCustomErrors, RulesetUpdateResponsePhaseHTTPLogCustomFields, RulesetUpdateResponsePhaseHTTPRatelimit, RulesetUpdateResponsePhaseHTTPRequestCacheSettings, RulesetUpdateResponsePhaseHTTPRequestDynamicRedirect, RulesetUpdateResponsePhaseHTTPRequestFirewallCustom, RulesetUpdateResponsePhaseHTTPRequestFirewallManaged, RulesetUpdateResponsePhaseHTTPRequestLateTransform, RulesetUpdateResponsePhaseHTTPRequestOrigin, RulesetUpdateResponsePhaseHTTPRequestRedirect, RulesetUpdateResponsePhaseHTTPRequestSanitize, RulesetUpdateResponsePhaseHTTPRequestSbfm, RulesetUpdateResponsePhaseHTTPRequestSelectConfiguration, RulesetUpdateResponsePhaseHTTPRequestTransform, RulesetUpdateResponsePhaseHTTPResponseCompression, RulesetUpdateResponsePhaseHTTPResponseFirewallManaged, RulesetUpdateResponsePhaseHTTPResponseHeadersTransform, RulesetUpdateResponsePhaseMagicTransit, RulesetUpdateResponsePhaseMagicTransitIDsManaged, RulesetUpdateResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type RulesetUpdateResponseRule struct {
	// The action to perform when the rule matches.
	Action           RulesetUpdateResponseRulesAction `json:"action"`
	ActionParameters interface{}                      `json:"action_parameters,required"`
	Categories       interface{}                      `json:"categories,required"`
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
	Version string                        `json:"version,required"`
	JSON    rulesetUpdateResponseRuleJSON `json:"-"`
	union   RulesetUpdateResponseRulesUnion
}

// rulesetUpdateResponseRuleJSON contains the JSON metadata for the struct
// [RulesetUpdateResponseRule]
type rulesetUpdateResponseRuleJSON struct {
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

func (r rulesetUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetUpdateResponseRule) AsUnion() RulesetUpdateResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule],
// [rulesets.RulesetUpdateResponseRulesRulesetsChallengeRule],
// [rulesets.RulesetUpdateResponseRulesRulesetsCompressResponseRule],
// [rulesets.ExecuteRule],
// [rulesets.RulesetUpdateResponseRulesRulesetsJsChallengeRule],
// [rulesets.LogRule],
// [rulesets.RulesetUpdateResponseRulesRulesetsManagedChallengeRule],
// [rulesets.RulesetUpdateResponseRulesRulesetsRedirectRule],
// [rulesets.RulesetUpdateResponseRulesRulesetsRewriteRule],
// [rulesets.RulesetUpdateResponseRulesRulesetsRouteRule],
// [rulesets.RulesetUpdateResponseRulesRulesetsScoreRule],
// [rulesets.RulesetUpdateResponseRulesRulesetsServeErrorRule],
// [rulesets.RulesetUpdateResponseRulesRulesetsSetConfigRule], [rulesets.SkipRule]
// or [rulesets.RulesetUpdateResponseRulesRulesetsSetCacheSettingsRule].
type RulesetUpdateResponseRulesUnion interface {
	implementsRulesetsRulesetUpdateResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsCompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsJsChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsRedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsRewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsRouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetUpdateResponseRulesRulesetsSetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
	)
}

type RulesetUpdateResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsChallengeRuleAction `json:"action"`
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
	JSON rulesetUpdateResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata
// for the struct [RulesetUpdateResponseRulesRulesetsChallengeRule]
type rulesetUpdateResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsChallengeRule) implementsRulesetsRulesetUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsChallengeRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsChallengeRuleActionChallenge RulesetUpdateResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetUpdateResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

type RulesetUpdateResponseRulesRulesetsCompressResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsCompressResponseRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParameters `json:"action_parameters"`
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
	Ref  string                                                     `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsCompressResponseRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsCompressResponseRuleJSON contains the JSON
// metadata for the struct [RulesetUpdateResponseRulesRulesetsCompressResponseRule]
type rulesetUpdateResponseRulesRulesetsCompressResponseRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsCompressResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsCompressResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsCompressResponseRule) implementsRulesetsRulesetUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsCompressResponseRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionCompressResponse RulesetUpdateResponseRulesRulesetsCompressResponseRuleAction = "compress_response"
)

func (r RulesetUpdateResponseRulesRulesetsCompressResponseRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms []RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm `json:"algorithms"`
	JSON       rulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersJSON        `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersJSON struct {
	Algorithms  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Compression algorithm to enable.
type RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName `json:"name"`
	JSON rulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON  `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm]
type rulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON) RawJSON() string {
	return r.raw
}

// Name of compression algorithm to enable.
type RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName string

const (
	RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone    RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "none"
	RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto    RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "auto"
	RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "default"
	RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip    RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "gzip"
	RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli  RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "brotli"
)

func (r RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone, RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto, RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault, RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip, RulesetUpdateResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli:
		return true
	}
	return false
}

type RulesetUpdateResponseRulesRulesetsJsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsJsChallengeRuleAction `json:"action"`
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
	JSON rulesetUpdateResponseRulesRulesetsJsChallengeRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsJsChallengeRuleJSON contains the JSON metadata
// for the struct [RulesetUpdateResponseRulesRulesetsJsChallengeRule]
type rulesetUpdateResponseRulesRulesetsJsChallengeRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsJsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsJsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsJsChallengeRule) implementsRulesetsRulesetUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsJsChallengeRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsJsChallengeRuleActionJsChallenge RulesetUpdateResponseRulesRulesetsJsChallengeRuleAction = "js_challenge"
)

func (r RulesetUpdateResponseRulesRulesetsJsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsJsChallengeRuleActionJsChallenge:
		return true
	}
	return false
}

type RulesetUpdateResponseRulesRulesetsManagedChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsManagedChallengeRuleAction `json:"action"`
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
	Ref  string                                                     `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsManagedChallengeRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsManagedChallengeRuleJSON contains the JSON
// metadata for the struct [RulesetUpdateResponseRulesRulesetsManagedChallengeRule]
type rulesetUpdateResponseRulesRulesetsManagedChallengeRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsManagedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsManagedChallengeRule) implementsRulesetsRulesetUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsManagedChallengeRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge RulesetUpdateResponseRulesRulesetsManagedChallengeRuleAction = "managed_challenge"
)

func (r RulesetUpdateResponseRulesRulesetsManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

type RulesetUpdateResponseRulesRulesetsRedirectRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsRedirectRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsRedirectRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetUpdateResponseRulesRulesetsRedirectRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRedirectRuleJSON contains the JSON metadata
// for the struct [RulesetUpdateResponseRulesRulesetsRedirectRule]
type rulesetUpdateResponseRulesRulesetsRedirectRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsRedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRedirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsRedirectRule) implementsRulesetsRulesetUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsRedirectRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsRedirectRuleActionRedirect RulesetUpdateResponseRulesRulesetsRedirectRuleAction = "redirect"
)

func (r RulesetUpdateResponseRulesRulesetsRedirectRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsRedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsRedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromList `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValue `json:"from_value"`
	JSON      rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersJSON      `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRedirectRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersJSON struct {
	FromList    apijson.Field
	FromValue   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRedirectRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on a bulk list lookup.
type RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key string `json:"key"`
	// The name of the list to match against.
	Name string                                                                     `json:"name"`
	JSON rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromListJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromListJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromList]
type rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromListJSON struct {
	Key         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromListJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on the request properties.
type RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString bool `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL `json:"target_url"`
	JSON      rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON      `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValue]
type rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON struct {
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	TargetURL           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON) RawJSON() string {
	return r.raw
}

// The status code to be used for the redirect.
type RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode float64

const (
	RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301 RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 301
	RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302 RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 302
	RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303 RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 303
	RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307 RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 307
	RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308 RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 308
)

func (r RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301, RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302, RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303, RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307, RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL struct {
	// The URL to redirect the request to.
	Value string `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                               `json:"expression"`
	JSON       rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON `json:"-"`
	union      RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion
}

// rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL]
type rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) AsUnion() RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion {
	return r.union
}

// The URL to redirect the request to.
//
// Union satisfied by
// [rulesets.RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
// or
// [rulesets.RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect].
type RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRulesetsRulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect{}),
		},
	)
}

type RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value string                                                                                                `json:"value"`
	JSON  rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
type rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRulesetsRulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                                                 `json:"expression"`
	JSON       rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect]
type rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRulesetsRulesetUpdateResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type RulesetUpdateResponseRulesRulesetsRewriteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsRewriteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsRewriteRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetUpdateResponseRulesRulesetsRewriteRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleJSON contains the JSON metadata for
// the struct [RulesetUpdateResponseRulesRulesetsRewriteRule]
type rulesetUpdateResponseRulesRulesetsRewriteRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsRewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsRewriteRule) implementsRulesetsRulesetUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsRewriteRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsRewriteRuleActionRewrite RulesetUpdateResponseRulesRulesetsRewriteRuleAction = "rewrite"
)

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsRewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers map[string]RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader `json:"headers"`
	// URI to rewrite the request to.
	URI  RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURI  `json:"uri"`
	JSON rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRewriteRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersJSON struct {
	Headers     apijson.Field
	URI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRewriteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Remove the header from the request.
type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader struct {
	Operation RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation `json:"operation,required"`
	// Static value for the header.
	Value string `json:"value"`
	// Expression for the header value.
	Expression string                                                                  `json:"expression"`
	JSON       rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON `json:"-"`
	union      RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON contains
// the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader]
type rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader) AsUnion() RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion {
	return r.union
}

// Remove the header from the request.
//
// Union satisfied by
// [rulesets.RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader],
// [rulesets.RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
// or
// [rulesets.RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader].
type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion interface {
	implementsRulesetsRulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader{}),
		},
	)
}

// Remove the header from the request.
type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation `json:"operation,required"`
	JSON      rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON      `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader]
type rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON struct {
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) implementsRulesetsRulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Set a request header with a static value.
type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader struct {
	Operation RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation `json:"operation,required"`
	// Static value for the header.
	Value string                                                                               `json:"value,required"`
	JSON  rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
type rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) implementsRulesetsRulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation string

const (
	RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation = "set"
)

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader struct {
	// Expression for the header value.
	Expression string                                                                                     `json:"expression,required"`
	Operation  RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation `json:"operation,required"`
	JSON       rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON      `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader]
type rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) implementsRulesetsRulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation string

const (
	RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation = "set"
)

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet:
		return true
	}
	return false
}

type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation string

const (
	RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "remove"
	RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet    RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "set"
)

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove, RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath `json:"path"`
	// Query portion rewrite.
	Query RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery `json:"query"`
	JSON  rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIJSON  `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIJSON contains
// the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURI]
type rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIJSON struct {
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIJSON) RawJSON() string {
	return r.raw
}

// Path portion rewrite.
type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                   `json:"expression"`
	JSON       rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON `json:"-"`
	union      RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath]
type rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath) AsUnion() RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion {
	return r.union
}

// Path portion rewrite.
//
// Union satisfied by
// [rulesets.RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
// or
// [rulesets.RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue].
type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion interface {
	implementsRulesetsRulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue{}),
		},
	)
}

type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue struct {
	// Predefined replacement value.
	Value string                                                                              `json:"value,required"`
	JSON  rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
type rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) implementsRulesetsRulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                               `json:"expression,required"`
	JSON       rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue]
type rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) implementsRulesetsRulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

// Query portion rewrite.
type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                    `json:"expression"`
	JSON       rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON `json:"-"`
	union      RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery]
type rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery) AsUnion() RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion {
	return r.union
}

// Query portion rewrite.
//
// Union satisfied by
// [rulesets.RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
// or
// [rulesets.RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue].
type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion interface {
	implementsRulesetsRulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue{}),
		},
	)
}

type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue struct {
	// Predefined replacement value.
	Value string                                                                               `json:"value,required"`
	JSON  rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
type rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) implementsRulesetsRulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                                `json:"expression,required"`
	JSON       rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue]
type rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) implementsRulesetsRulesetUpdateResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type RulesetUpdateResponseRulesRulesetsRouteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsRouteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsRouteRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetUpdateResponseRulesRulesetsRouteRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRouteRuleJSON contains the JSON metadata for
// the struct [RulesetUpdateResponseRulesRulesetsRouteRule]
type rulesetUpdateResponseRulesRulesetsRouteRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsRouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRouteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsRouteRule) implementsRulesetsRulesetUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsRouteRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsRouteRuleActionRoute RulesetUpdateResponseRulesRulesetsRouteRuleAction = "route"
)

func (r RulesetUpdateResponseRulesRulesetsRouteRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsRouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsRouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader string `json:"host_header"`
	// Override the IP/TCP destination.
	Origin RulesetUpdateResponseRulesRulesetsRouteRuleActionParametersOrigin `json:"origin"`
	// Override the Server Name Indication (SNI).
	Sni  RulesetUpdateResponseRulesRulesetsRouteRuleActionParametersSni  `json:"sni"`
	JSON rulesetUpdateResponseRulesRulesetsRouteRuleActionParametersJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRouteRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRouteRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsRouteRuleActionParametersJSON struct {
	HostHeader  apijson.Field
	Origin      apijson.Field
	Sni         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRouteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRouteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Override the IP/TCP destination.
type RulesetUpdateResponseRulesRulesetsRouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host string `json:"host"`
	// Override the destination port.
	Port float64                                                               `json:"port"`
	JSON rulesetUpdateResponseRulesRulesetsRouteRuleActionParametersOriginJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRouteRuleActionParametersOriginJSON contains
// the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRouteRuleActionParametersOrigin]
type rulesetUpdateResponseRulesRulesetsRouteRuleActionParametersOriginJSON struct {
	Host        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRouteRuleActionParametersOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRouteRuleActionParametersOriginJSON) RawJSON() string {
	return r.raw
}

// Override the Server Name Indication (SNI).
type RulesetUpdateResponseRulesRulesetsRouteRuleActionParametersSni struct {
	// The SNI override.
	Value string                                                             `json:"value,required"`
	JSON  rulesetUpdateResponseRulesRulesetsRouteRuleActionParametersSniJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsRouteRuleActionParametersSniJSON contains the
// JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsRouteRuleActionParametersSni]
type rulesetUpdateResponseRulesRulesetsRouteRuleActionParametersSniJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsRouteRuleActionParametersSni) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsRouteRuleActionParametersSniJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseRulesRulesetsScoreRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsScoreRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsScoreRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetUpdateResponseRulesRulesetsScoreRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsScoreRuleJSON contains the JSON metadata for
// the struct [RulesetUpdateResponseRulesRulesetsScoreRule]
type rulesetUpdateResponseRulesRulesetsScoreRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsScoreRule) implementsRulesetsRulesetUpdateResponseRule() {}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsScoreRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsScoreRuleActionScore RulesetUpdateResponseRulesRulesetsScoreRuleAction = "score"
)

func (r RulesetUpdateResponseRulesRulesetsScoreRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment int64                                                           `json:"increment"`
	JSON      rulesetUpdateResponseRulesRulesetsScoreRuleActionParametersJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsScoreRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsScoreRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsScoreRuleActionParametersJSON struct {
	Increment   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsScoreRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsScoreRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

type RulesetUpdateResponseRulesRulesetsServeErrorRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsServeErrorRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetUpdateResponseRulesRulesetsServeErrorRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsServeErrorRuleJSON contains the JSON metadata
// for the struct [RulesetUpdateResponseRulesRulesetsServeErrorRule]
type rulesetUpdateResponseRulesRulesetsServeErrorRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsServeErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsServeErrorRule) implementsRulesetsRulesetUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsServeErrorRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsServeErrorRuleActionServeError RulesetUpdateResponseRulesRulesetsServeErrorRuleAction = "serve_error"
)

func (r RulesetUpdateResponseRulesRulesetsServeErrorRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParameters struct {
	// Error response content.
	Content string `json:"content"`
	// Content-type header to set with the response.
	ContentType RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType `json:"content_type"`
	// The status code to use for the error.
	StatusCode float64                                                              `json:"status_code"`
	JSON       rulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Content-type header to set with the response.
type RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType string

const (
	RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType = "application/json"
	RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml         RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/xml"
	RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain       RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/plain"
	RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML        RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/html"
)

func (r RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson, RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml, RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain, RulesetUpdateResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

type RulesetUpdateResponseRulesRulesetsSetConfigRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsSetConfigRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetUpdateResponseRulesRulesetsSetConfigRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetConfigRuleJSON contains the JSON metadata
// for the struct [RulesetUpdateResponseRulesRulesetsSetConfigRule]
type rulesetUpdateResponseRulesRulesetsSetConfigRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsSetConfigRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetConfigRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetConfigRule) implementsRulesetsRulesetUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsSetConfigRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionSetConfig RulesetUpdateResponseRulesRulesetsSetConfigRuleAction = "set_config"
)

func (r RulesetUpdateResponseRulesRulesetsSetConfigRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites bool `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominify `json:"autominify"`
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
	Polish RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolish `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader bool `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes bool `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg  bool                                                                `json:"sxg"`
	JSON rulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Select which file extensions to minify automatically.
type RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	Css bool `json:"css"`
	// Minify HTML files.
	HTML bool `json:"html"`
	// Minify JS files.
	Js   bool                                                                          `json:"js"`
	JSON rulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominify]
type rulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON) RawJSON() string {
	return r.raw
}

// Configure the Polish level.
type RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolish string

const (
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolishOff      RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolish = "off"
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossless"
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy    RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossy"
)

func (r RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolishOff, RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless, RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy:
		return true
	}
	return false
}

// Configure the Security Level.
type RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel string

const (
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff            RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "off"
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "essentially_off"
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow            RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "low"
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium         RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "medium"
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh           RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "high"
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack    RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff, RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff, RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow, RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium, RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh, RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL string

const (
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLOff        RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL = "off"
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible   RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL = "flexible"
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLFull       RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL = "full"
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict     RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL = "strict"
	RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL = "origin_pull"
)

func (r RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLOff, RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible, RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLFull, RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict, RulesetUpdateResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParameters `json:"action_parameters"`
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
	Ref  string                                                     `json:"ref"`
	JSON rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleJSON contains the JSON
// metadata for the struct [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRule]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetUpdateResponseRulesRulesetsSetCacheSettingsRule) implementsRulesetsRulesetUpdateResponseRule() {
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleAction string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleAction = "set_cache_settings"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts []int64 `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache bool `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL `json:"edge_ttl"`
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
	ServeStale RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale `json:"serve_stale"`
	JSON       rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON       `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParameters]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON struct {
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

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64                                                                                `json:"default"`
	JSON    rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON struct {
	Mode        apijson.Field
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON) RawJSON() string {
	return r.raw
}

// Determines which browser ttl mode to use.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType bool `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor bool `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder bool                                                                               `json:"ignore_query_strings_order"`
	JSON                    rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON struct {
	CacheByDeviceType       apijson.Field
	CacheDeceptionArmor     apijson.Field
	CustomKey               apijson.Field
	IgnoreQueryStringsOrder apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON) RawJSON() string {
	return r.raw
}

// Customize which components of the request are included or excluded from the
// cache key.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser `json:"user"`
	JSON rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON) RawJSON() string {
	return r.raw
}

// The cookies to include in building the cache key.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Include these cookies' names and their values.
	Include []string                                                                                          `json:"include"`
	JSON    rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON) RawJSON() string {
	return r.raw
}

// The header names and values to include in building the cache key.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin bool `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include []string                                                                                          `json:"include"`
	JSON    rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON struct {
	CheckPresence apijson.Field
	ExcludeOrigin apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON) RawJSON() string {
	return r.raw
}

// Whether to use the original host or the resolved host in the cache key.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved bool                                                                                            `json:"resolved"`
	JSON     rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON) RawJSON() string {
	return r.raw
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude `json:"include"`
	JSON    rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON    `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON) RawJSON() string {
	return r.raw
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Exclude all query string parameters from use in building the cache key.
	All bool `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List []string                                                                                                      `json:"list"`
	JSON rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON) RawJSON() string {
	return r.raw
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Use all query string parameters in the cache key.
	All bool `json:"all"`
	// A list of query string parameters used to build the cache key.
	List []string                                                                                                      `json:"list"`
	JSON rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON) RawJSON() string {
	return r.raw
}

// Characteristics of the request user agent used in building the cache key.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType bool `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo bool `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang bool                                                                                            `json:"lang"`
	JSON rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON) RawJSON() string {
	return r.raw
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible bool `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize int64                                                                                  `json:"min_file_size,required"`
	JSON        rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON struct {
	Eligible    apijson.Field
	MinFileSize apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON) RawJSON() string {
	return r.raw
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64 `json:"default,required"`
	// edge ttl options
	Mode RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL []RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL `json:"status_code_ttl,required"`
	JSON          rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON            `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON struct {
	Default       apijson.Field
	Mode          apijson.Field
	StatusCodeTTL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON) RawJSON() string {
	return r.raw
}

// edge ttl options
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value int64 `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue int64                                                                                          `json:"status_code_value"`
	JSON            rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON struct {
	Value           apijson.Field
	StatusCodeRange apijson.Field
	StatusCodeValue apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON) RawJSON() string {
	return r.raw
}

// The range of status codes used to apply the selected mode.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// response status code lower bound
	From int64 `json:"from,required"`
	// response status code upper bound
	To   int64                                                                                                         `json:"to,required"`
	JSON rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON) RawJSON() string {
	return r.raw
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating bool                                                                                 `json:"disable_stale_while_updating,required"`
	JSON                      rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON `json:"-"`
}

// rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON
// contains the JSON metadata for the struct
// [RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale]
type rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON struct {
	DisableStaleWhileUpdating apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *RulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetUpdateResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetUpdateResponseRulesAction string

const (
	RulesetUpdateResponseRulesActionBlock            RulesetUpdateResponseRulesAction = "block"
	RulesetUpdateResponseRulesActionChallenge        RulesetUpdateResponseRulesAction = "challenge"
	RulesetUpdateResponseRulesActionCompressResponse RulesetUpdateResponseRulesAction = "compress_response"
	RulesetUpdateResponseRulesActionExecute          RulesetUpdateResponseRulesAction = "execute"
	RulesetUpdateResponseRulesActionJsChallenge      RulesetUpdateResponseRulesAction = "js_challenge"
	RulesetUpdateResponseRulesActionLog              RulesetUpdateResponseRulesAction = "log"
	RulesetUpdateResponseRulesActionManagedChallenge RulesetUpdateResponseRulesAction = "managed_challenge"
	RulesetUpdateResponseRulesActionRedirect         RulesetUpdateResponseRulesAction = "redirect"
	RulesetUpdateResponseRulesActionRewrite          RulesetUpdateResponseRulesAction = "rewrite"
	RulesetUpdateResponseRulesActionRoute            RulesetUpdateResponseRulesAction = "route"
	RulesetUpdateResponseRulesActionScore            RulesetUpdateResponseRulesAction = "score"
	RulesetUpdateResponseRulesActionServeError       RulesetUpdateResponseRulesAction = "serve_error"
	RulesetUpdateResponseRulesActionSetConfig        RulesetUpdateResponseRulesAction = "set_config"
	RulesetUpdateResponseRulesActionSkip             RulesetUpdateResponseRulesAction = "skip"
	RulesetUpdateResponseRulesActionSetCacheSettings RulesetUpdateResponseRulesAction = "set_cache_settings"
)

func (r RulesetUpdateResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetUpdateResponseRulesActionBlock, RulesetUpdateResponseRulesActionChallenge, RulesetUpdateResponseRulesActionCompressResponse, RulesetUpdateResponseRulesActionExecute, RulesetUpdateResponseRulesActionJsChallenge, RulesetUpdateResponseRulesActionLog, RulesetUpdateResponseRulesActionManagedChallenge, RulesetUpdateResponseRulesActionRedirect, RulesetUpdateResponseRulesActionRewrite, RulesetUpdateResponseRulesActionRoute, RulesetUpdateResponseRulesActionScore, RulesetUpdateResponseRulesActionServeError, RulesetUpdateResponseRulesActionSetConfig, RulesetUpdateResponseRulesActionSkip, RulesetUpdateResponseRulesActionSetCacheSettings:
		return true
	}
	return false
}

// A ruleset object.
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

func (r RulesetGetResponseKind) IsKnown() bool {
	switch r {
	case RulesetGetResponseKindManaged, RulesetGetResponseKindCustom, RulesetGetResponseKindRoot, RulesetGetResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type RulesetGetResponsePhase string

const (
	RulesetGetResponsePhaseDDoSL4                         RulesetGetResponsePhase = "ddos_l4"
	RulesetGetResponsePhaseDDoSL7                         RulesetGetResponsePhase = "ddos_l7"
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

func (r RulesetGetResponsePhase) IsKnown() bool {
	switch r {
	case RulesetGetResponsePhaseDDoSL4, RulesetGetResponsePhaseDDoSL7, RulesetGetResponsePhaseHTTPConfigSettings, RulesetGetResponsePhaseHTTPCustomErrors, RulesetGetResponsePhaseHTTPLogCustomFields, RulesetGetResponsePhaseHTTPRatelimit, RulesetGetResponsePhaseHTTPRequestCacheSettings, RulesetGetResponsePhaseHTTPRequestDynamicRedirect, RulesetGetResponsePhaseHTTPRequestFirewallCustom, RulesetGetResponsePhaseHTTPRequestFirewallManaged, RulesetGetResponsePhaseHTTPRequestLateTransform, RulesetGetResponsePhaseHTTPRequestOrigin, RulesetGetResponsePhaseHTTPRequestRedirect, RulesetGetResponsePhaseHTTPRequestSanitize, RulesetGetResponsePhaseHTTPRequestSbfm, RulesetGetResponsePhaseHTTPRequestSelectConfiguration, RulesetGetResponsePhaseHTTPRequestTransform, RulesetGetResponsePhaseHTTPResponseCompression, RulesetGetResponsePhaseHTTPResponseFirewallManaged, RulesetGetResponsePhaseHTTPResponseHeadersTransform, RulesetGetResponsePhaseMagicTransit, RulesetGetResponsePhaseMagicTransitIDsManaged, RulesetGetResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type RulesetGetResponseRule struct {
	// The action to perform when the rule matches.
	Action           RulesetGetResponseRulesAction `json:"action"`
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
	JSON    rulesetGetResponseRuleJSON `json:"-"`
	union   RulesetGetResponseRulesUnion
}

// rulesetGetResponseRuleJSON contains the JSON metadata for the struct
// [RulesetGetResponseRule]
type rulesetGetResponseRuleJSON struct {
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

func (r rulesetGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetGetResponseRule) AsUnion() RulesetGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule],
// [rulesets.RulesetGetResponseRulesRulesetsChallengeRule],
// [rulesets.RulesetGetResponseRulesRulesetsCompressResponseRule],
// [rulesets.ExecuteRule],
// [rulesets.RulesetGetResponseRulesRulesetsJsChallengeRule], [rulesets.LogRule],
// [rulesets.RulesetGetResponseRulesRulesetsManagedChallengeRule],
// [rulesets.RulesetGetResponseRulesRulesetsRedirectRule],
// [rulesets.RulesetGetResponseRulesRulesetsRewriteRule],
// [rulesets.RulesetGetResponseRulesRulesetsRouteRule],
// [rulesets.RulesetGetResponseRulesRulesetsScoreRule],
// [rulesets.RulesetGetResponseRulesRulesetsServeErrorRule],
// [rulesets.RulesetGetResponseRulesRulesetsSetConfigRule], [rulesets.SkipRule] or
// [rulesets.RulesetGetResponseRulesRulesetsSetCacheSettingsRule].
type RulesetGetResponseRulesUnion interface {
	implementsRulesetsRulesetGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsCompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsJsChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsRedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsRewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsRouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsSetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RulesetGetResponseRulesRulesetsSetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
	)
}

type RulesetGetResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsChallengeRuleAction `json:"action"`
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
	JSON rulesetGetResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [RulesetGetResponseRulesRulesetsChallengeRule]
type rulesetGetResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsChallengeRule) implementsRulesetsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsChallengeRuleAction string

const (
	RulesetGetResponseRulesRulesetsChallengeRuleActionChallenge RulesetGetResponseRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetGetResponseRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

type RulesetGetResponseRulesRulesetsCompressResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsCompressResponseRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsCompressResponseRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetGetResponseRulesRulesetsCompressResponseRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsCompressResponseRuleJSON contains the JSON
// metadata for the struct [RulesetGetResponseRulesRulesetsCompressResponseRule]
type rulesetGetResponseRulesRulesetsCompressResponseRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsCompressResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsCompressResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsCompressResponseRule) implementsRulesetsRulesetGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsCompressResponseRuleAction string

const (
	RulesetGetResponseRulesRulesetsCompressResponseRuleActionCompressResponse RulesetGetResponseRulesRulesetsCompressResponseRuleAction = "compress_response"
)

func (r RulesetGetResponseRulesRulesetsCompressResponseRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsCompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsCompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms []RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm `json:"algorithms"`
	JSON       rulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON        `json:"-"`
}

// rulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsCompressResponseRuleActionParameters]
type rulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON struct {
	Algorithms  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsCompressResponseRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Compression algorithm to enable.
type RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName `json:"name"`
	JSON rulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON  `json:"-"`
}

// rulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm]
type rulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmJSON) RawJSON() string {
	return r.raw
}

// Name of compression algorithm to enable.
type RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName string

const (
	RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone    RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "none"
	RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto    RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "auto"
	RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "default"
	RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip    RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "gzip"
	RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli  RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "brotli"
)

func (r RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone, RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto, RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault, RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip, RulesetGetResponseRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli:
		return true
	}
	return false
}

type RulesetGetResponseRulesRulesetsJsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsJsChallengeRuleAction `json:"action"`
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
	JSON rulesetGetResponseRulesRulesetsJsChallengeRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsJsChallengeRuleJSON contains the JSON metadata
// for the struct [RulesetGetResponseRulesRulesetsJsChallengeRule]
type rulesetGetResponseRulesRulesetsJsChallengeRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsJsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsJsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsJsChallengeRule) implementsRulesetsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsJsChallengeRuleAction string

const (
	RulesetGetResponseRulesRulesetsJsChallengeRuleActionJsChallenge RulesetGetResponseRulesRulesetsJsChallengeRuleAction = "js_challenge"
)

func (r RulesetGetResponseRulesRulesetsJsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsJsChallengeRuleActionJsChallenge:
		return true
	}
	return false
}

type RulesetGetResponseRulesRulesetsManagedChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsManagedChallengeRuleAction `json:"action"`
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
	JSON rulesetGetResponseRulesRulesetsManagedChallengeRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsManagedChallengeRuleJSON contains the JSON
// metadata for the struct [RulesetGetResponseRulesRulesetsManagedChallengeRule]
type rulesetGetResponseRulesRulesetsManagedChallengeRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsManagedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsManagedChallengeRule) implementsRulesetsRulesetGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsManagedChallengeRuleAction string

const (
	RulesetGetResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge RulesetGetResponseRulesRulesetsManagedChallengeRuleAction = "managed_challenge"
)

func (r RulesetGetResponseRulesRulesetsManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

type RulesetGetResponseRulesRulesetsRedirectRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsRedirectRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsRedirectRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetGetResponseRulesRulesetsRedirectRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRedirectRuleJSON contains the JSON metadata for
// the struct [RulesetGetResponseRulesRulesetsRedirectRule]
type rulesetGetResponseRulesRulesetsRedirectRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsRedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRedirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsRedirectRule) implementsRulesetsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsRedirectRuleAction string

const (
	RulesetGetResponseRulesRulesetsRedirectRuleActionRedirect RulesetGetResponseRulesRulesetsRedirectRuleAction = "redirect"
)

func (r RulesetGetResponseRulesRulesetsRedirectRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsRedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsRedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromList `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValue `json:"from_value"`
	JSON      rulesetGetResponseRulesRulesetsRedirectRuleActionParametersJSON      `json:"-"`
}

// rulesetGetResponseRulesRulesetsRedirectRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRedirectRuleActionParameters]
type rulesetGetResponseRulesRulesetsRedirectRuleActionParametersJSON struct {
	FromList    apijson.Field
	FromValue   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRedirectRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRedirectRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on a bulk list lookup.
type RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key string `json:"key"`
	// The name of the list to match against.
	Name string                                                                  `json:"name"`
	JSON rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromList]
type rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON struct {
	Key         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromListJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on the request properties.
type RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString bool `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL `json:"target_url"`
	JSON      rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON      `json:"-"`
}

// rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValue]
type rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON struct {
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	TargetURL           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueJSON) RawJSON() string {
	return r.raw
}

// The status code to be used for the redirect.
type RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode float64

const (
	RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301 RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 301
	RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302 RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 302
	RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303 RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 303
	RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307 RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 307
	RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308 RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 308
)

func (r RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301, RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302, RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303, RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307, RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL struct {
	// The URL to redirect the request to.
	Value string `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                            `json:"expression"`
	JSON       rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON `json:"-"`
	union      RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion
}

// rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL]
type rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) AsUnion() RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion {
	return r.union
}

// The URL to redirect the request to.
//
// Union satisfied by
// [rulesets.RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
// or
// [rulesets.RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect].
type RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRulesetsRulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect{}),
		},
	)
}

type RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value string                                                                                             `json:"value"`
	JSON  rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
type rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRulesetsRulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                                                              `json:"expression"`
	JSON       rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect]
type rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRulesetsRulesetGetResponseRulesRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type RulesetGetResponseRulesRulesetsRewriteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsRewriteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsRewriteRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetGetResponseRulesRulesetsRewriteRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRewriteRuleJSON contains the JSON metadata for
// the struct [RulesetGetResponseRulesRulesetsRewriteRule]
type rulesetGetResponseRulesRulesetsRewriteRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsRewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsRewriteRule) implementsRulesetsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsRewriteRuleAction string

const (
	RulesetGetResponseRulesRulesetsRewriteRuleActionRewrite RulesetGetResponseRulesRulesetsRewriteRuleAction = "rewrite"
)

func (r RulesetGetResponseRulesRulesetsRewriteRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsRewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsRewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers map[string]RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeader `json:"headers"`
	// URI to rewrite the request to.
	URI  RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURI  `json:"uri"`
	JSON rulesetGetResponseRulesRulesetsRewriteRuleActionParametersJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRewriteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetGetResponseRulesRulesetsRewriteRuleActionParameters]
type rulesetGetResponseRulesRulesetsRewriteRuleActionParametersJSON struct {
	Headers     apijson.Field
	URI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRewriteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Remove the header from the request.
type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeader struct {
	Operation RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation `json:"operation,required"`
	// Static value for the header.
	Value string `json:"value"`
	// Expression for the header value.
	Expression string                                                               `json:"expression"`
	JSON       rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON `json:"-"`
	union      RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion
}

// rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeader]
type rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeaderJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeader) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeader) AsUnion() RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion {
	return r.union
}

// Remove the header from the request.
//
// Union satisfied by
// [rulesets.RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader],
// [rulesets.RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
// or
// [rulesets.RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader].
type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion interface {
	implementsRulesetsRulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeader()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader{}),
		},
	)
}

// Remove the header from the request.
type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation `json:"operation,required"`
	JSON      rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON      `json:"-"`
}

// rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader]
type rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON struct {
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) implementsRulesetsRulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Set a request header with a static value.
type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader struct {
	Operation RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation `json:"operation,required"`
	// Static value for the header.
	Value string                                                                            `json:"value,required"`
	JSON  rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader]
type rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) implementsRulesetsRulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation string

const (
	RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation = "set"
)

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader struct {
	// Expression for the header value.
	Expression string                                                                                  `json:"expression,required"`
	Operation  RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation `json:"operation,required"`
	JSON       rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON      `json:"-"`
}

// rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader]
type rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) implementsRulesetsRulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeader() {
}

type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation string

const (
	RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation = "set"
)

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet:
		return true
	}
	return false
}

type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation string

const (
	RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "remove"
	RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet    RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation = "set"
)

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove, RulesetGetResponseRulesRulesetsRewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPath `json:"path"`
	// Query portion rewrite.
	Query RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery `json:"query"`
	JSON  rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON  `json:"-"`
}

// rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON contains the
// JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURI]
type rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON struct {
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIJSON) RawJSON() string {
	return r.raw
}

// Path portion rewrite.
type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPath struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                `json:"expression"`
	JSON       rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON `json:"-"`
	union      RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion
}

// rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPath]
type rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPath) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPath) AsUnion() RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion {
	return r.union
}

// Path portion rewrite.
//
// Union satisfied by
// [rulesets.RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
// or
// [rulesets.RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue].
type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion interface {
	implementsRulesetsRulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPath()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue{}),
		},
	)
}

type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue struct {
	// Predefined replacement value.
	Value string                                                                           `json:"value,required"`
	JSON  rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue]
type rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) implementsRulesetsRulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                            `json:"expression,required"`
	JSON       rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue]
type rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) implementsRulesetsRulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIPath() {
}

// Query portion rewrite.
type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string                                                                 `json:"expression"`
	JSON       rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON `json:"-"`
	union      RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion
}

// rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery]
type rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryJSON) RawJSON() string {
	return r.raw
}

func (r *RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery) AsUnion() RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion {
	return r.union
}

// Query portion rewrite.
//
// Union satisfied by
// [rulesets.RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
// or
// [rulesets.RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue].
type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion interface {
	implementsRulesetsRulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue{}),
		},
	)
}

type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue struct {
	// Predefined replacement value.
	Value string                                                                            `json:"value,required"`
	JSON  rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue]
type rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) implementsRulesetsRulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                                                                             `json:"expression,required"`
	JSON       rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue]
type rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) implementsRulesetsRulesetGetResponseRulesRulesetsRewriteRuleActionParametersURIQuery() {
}

type RulesetGetResponseRulesRulesetsRouteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsRouteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsRouteRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetGetResponseRulesRulesetsRouteRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRouteRuleJSON contains the JSON metadata for the
// struct [RulesetGetResponseRulesRulesetsRouteRule]
type rulesetGetResponseRulesRulesetsRouteRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsRouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRouteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsRouteRule) implementsRulesetsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsRouteRuleAction string

const (
	RulesetGetResponseRulesRulesetsRouteRuleActionRoute RulesetGetResponseRulesRulesetsRouteRuleAction = "route"
)

func (r RulesetGetResponseRulesRulesetsRouteRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsRouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsRouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader string `json:"host_header"`
	// Override the IP/TCP destination.
	Origin RulesetGetResponseRulesRulesetsRouteRuleActionParametersOrigin `json:"origin"`
	// Override the Server Name Indication (SNI).
	Sni  RulesetGetResponseRulesRulesetsRouteRuleActionParametersSni  `json:"sni"`
	JSON rulesetGetResponseRulesRulesetsRouteRuleActionParametersJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRouteRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetGetResponseRulesRulesetsRouteRuleActionParameters]
type rulesetGetResponseRulesRulesetsRouteRuleActionParametersJSON struct {
	HostHeader  apijson.Field
	Origin      apijson.Field
	Sni         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRouteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRouteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Override the IP/TCP destination.
type RulesetGetResponseRulesRulesetsRouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host string `json:"host"`
	// Override the destination port.
	Port float64                                                            `json:"port"`
	JSON rulesetGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON contains the
// JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRouteRuleActionParametersOrigin]
type rulesetGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON struct {
	Host        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRouteRuleActionParametersOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRouteRuleActionParametersOriginJSON) RawJSON() string {
	return r.raw
}

// Override the Server Name Indication (SNI).
type RulesetGetResponseRulesRulesetsRouteRuleActionParametersSni struct {
	// The SNI override.
	Value string                                                          `json:"value,required"`
	JSON  rulesetGetResponseRulesRulesetsRouteRuleActionParametersSniJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsRouteRuleActionParametersSniJSON contains the
// JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsRouteRuleActionParametersSni]
type rulesetGetResponseRulesRulesetsRouteRuleActionParametersSniJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsRouteRuleActionParametersSni) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsRouteRuleActionParametersSniJSON) RawJSON() string {
	return r.raw
}

type RulesetGetResponseRulesRulesetsScoreRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsScoreRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsScoreRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetGetResponseRulesRulesetsScoreRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsScoreRuleJSON contains the JSON metadata for the
// struct [RulesetGetResponseRulesRulesetsScoreRule]
type rulesetGetResponseRulesRulesetsScoreRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsScoreRule) implementsRulesetsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsScoreRuleAction string

const (
	RulesetGetResponseRulesRulesetsScoreRuleActionScore RulesetGetResponseRulesRulesetsScoreRuleAction = "score"
)

func (r RulesetGetResponseRulesRulesetsScoreRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment int64                                                        `json:"increment"`
	JSON      rulesetGetResponseRulesRulesetsScoreRuleActionParametersJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsScoreRuleActionParametersJSON contains the JSON
// metadata for the struct
// [RulesetGetResponseRulesRulesetsScoreRuleActionParameters]
type rulesetGetResponseRulesRulesetsScoreRuleActionParametersJSON struct {
	Increment   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsScoreRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsScoreRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

type RulesetGetResponseRulesRulesetsServeErrorRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsServeErrorRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsServeErrorRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetGetResponseRulesRulesetsServeErrorRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsServeErrorRuleJSON contains the JSON metadata for
// the struct [RulesetGetResponseRulesRulesetsServeErrorRule]
type rulesetGetResponseRulesRulesetsServeErrorRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsServeErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsServeErrorRule) implementsRulesetsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsServeErrorRuleAction string

const (
	RulesetGetResponseRulesRulesetsServeErrorRuleActionServeError RulesetGetResponseRulesRulesetsServeErrorRuleAction = "serve_error"
)

func (r RulesetGetResponseRulesRulesetsServeErrorRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsServeErrorRuleActionParameters struct {
	// Error response content.
	Content string `json:"content"`
	// Content-type header to set with the response.
	ContentType RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentType `json:"content_type"`
	// The status code to use for the error.
	StatusCode float64                                                           `json:"status_code"`
	JSON       rulesetGetResponseRulesRulesetsServeErrorRuleActionParametersJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsServeErrorRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsServeErrorRuleActionParameters]
type rulesetGetResponseRulesRulesetsServeErrorRuleActionParametersJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsServeErrorRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsServeErrorRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Content-type header to set with the response.
type RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentType string

const (
	RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "application/json"
	RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml         RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/xml"
	RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain       RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/plain"
	RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML        RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentType = "text/html"
)

func (r RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson, RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml, RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain, RulesetGetResponseRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

type RulesetGetResponseRulesRulesetsSetConfigRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsSetConfigRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsSetConfigRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetGetResponseRulesRulesetsSetConfigRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetConfigRuleJSON contains the JSON metadata for
// the struct [RulesetGetResponseRulesRulesetsSetConfigRule]
type rulesetGetResponseRulesRulesetsSetConfigRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsSetConfigRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetConfigRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetConfigRule) implementsRulesetsRulesetGetResponseRule() {}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsSetConfigRuleAction string

const (
	RulesetGetResponseRulesRulesetsSetConfigRuleActionSetConfig RulesetGetResponseRulesRulesetsSetConfigRuleAction = "set_config"
)

func (r RulesetGetResponseRulesRulesetsSetConfigRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsSetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites bool `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify `json:"autominify"`
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
	Polish RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersPolish `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader bool `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes bool `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSL `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg  bool                                                             `json:"sxg"`
	JSON rulesetGetResponseRulesRulesetsSetConfigRuleActionParametersJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetConfigRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetConfigRuleActionParameters]
type rulesetGetResponseRulesRulesetsSetConfigRuleActionParametersJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsSetConfigRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetConfigRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Select which file extensions to minify automatically.
type RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	Css bool `json:"css"`
	// Minify HTML files.
	HTML bool `json:"html"`
	// Minify JS files.
	Js   bool                                                                       `json:"js"`
	JSON rulesetGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify]
type rulesetGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	Js          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersAutominify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetConfigRuleActionParametersAutominifyJSON) RawJSON() string {
	return r.raw
}

// Configure the Polish level.
type RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersPolish string

const (
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersPolishOff      RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "off"
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossless"
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy    RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersPolish = "lossy"
)

func (r RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersPolishOff, RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossless, RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersPolishLossy:
		return true
	}
	return false
}

// Configure the Security Level.
type RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel string

const (
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff            RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "off"
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "essentially_off"
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow            RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "low"
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium         RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "medium"
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh           RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "high"
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack    RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff, RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff, RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow, RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium, RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh, RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSL string

const (
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOff        RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "off"
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible   RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "flexible"
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFull       RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "full"
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict     RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "strict"
	RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSL = "origin_pull"
)

func (r RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOff, RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFlexible, RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSLFull, RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSLStrict, RulesetGetResponseRulesRulesetsSetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

type RulesetGetResponseRulesRulesetsSetCacheSettingsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RulesetGetResponseRulesRulesetsSetCacheSettingsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters `json:"action_parameters"`
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
	JSON rulesetGetResponseRulesRulesetsSetCacheSettingsRuleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleJSON contains the JSON
// metadata for the struct [RulesetGetResponseRulesRulesetsSetCacheSettingsRule]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r RulesetGetResponseRulesRulesetsSetCacheSettingsRule) implementsRulesetsRulesetGetResponseRule() {
}

// The action to perform when the rule matches.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleAction string

const (
	RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings RulesetGetResponseRulesRulesetsSetCacheSettingsRuleAction = "set_cache_settings"
)

func (r RulesetGetResponseRulesRulesetsSetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts []int64 `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache bool `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL `json:"edge_ttl"`
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
	ServeStale RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale `json:"serve_stale"`
	JSON       rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON       `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON struct {
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

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64                                                                             `json:"default"`
	JSON    rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON struct {
	Mode        apijson.Field
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLJSON) RawJSON() string {
	return r.raw
}

// Determines which browser ttl mode to use.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
)

func (r RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType bool `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor bool `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder bool                                                                            `json:"ignore_query_strings_order"`
	JSON                    rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON struct {
	CacheByDeviceType       apijson.Field
	CacheDeceptionArmor     apijson.Field
	CustomKey               apijson.Field
	IgnoreQueryStringsOrder apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyJSON) RawJSON() string {
	return r.raw
}

// Customize which components of the request are included or excluded from the
// cache key.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser `json:"user"`
	JSON rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON) RawJSON() string {
	return r.raw
}

// The cookies to include in building the cache key.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Include these cookies' names and their values.
	Include []string                                                                                       `json:"include"`
	JSON    rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON) RawJSON() string {
	return r.raw
}

// The header names and values to include in building the cache key.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin bool `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include []string                                                                                       `json:"include"`
	JSON    rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON struct {
	CheckPresence apijson.Field
	ExcludeOrigin apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON) RawJSON() string {
	return r.raw
}

// Whether to use the original host or the resolved host in the cache key.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved bool                                                                                         `json:"resolved"`
	JSON     rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON) RawJSON() string {
	return r.raw
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude `json:"include"`
	JSON    rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON    `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON) RawJSON() string {
	return r.raw
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Exclude all query string parameters from use in building the cache key.
	All bool `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List []string                                                                                                   `json:"list"`
	JSON rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON) RawJSON() string {
	return r.raw
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Use all query string parameters in the cache key.
	All bool `json:"all"`
	// A list of query string parameters used to build the cache key.
	List []string                                                                                                   `json:"list"`
	JSON rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON) RawJSON() string {
	return r.raw
}

// Characteristics of the request user agent used in building the cache key.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType bool `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo bool `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang bool                                                                                         `json:"lang"`
	JSON rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON) RawJSON() string {
	return r.raw
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible bool `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize int64                                                                               `json:"min_file_size,required"`
	JSON        rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON struct {
	Eligible    apijson.Field
	MinFileSize apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserveJSON) RawJSON() string {
	return r.raw
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64 `json:"default,required"`
	// edge ttl options
	Mode RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL []RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL `json:"status_code_ttl,required"`
	JSON          rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON            `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON struct {
	Default       apijson.Field
	Mode          apijson.Field
	StatusCodeTTL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLJSON) RawJSON() string {
	return r.raw
}

// edge ttl options
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value int64 `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue int64                                                                                       `json:"status_code_value"`
	JSON            rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON struct {
	Value           apijson.Field
	StatusCodeRange apijson.Field
	StatusCodeValue apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON) RawJSON() string {
	return r.raw
}

// The range of status codes used to apply the selected mode.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// response status code lower bound
	From int64 `json:"from,required"`
	// response status code upper bound
	To   int64                                                                                                      `json:"to,required"`
	JSON rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON) RawJSON() string {
	return r.raw
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating bool                                                                              `json:"disable_stale_while_updating,required"`
	JSON                      rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON `json:"-"`
}

// rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON
// contains the JSON metadata for the struct
// [RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale]
type rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON struct {
	DisableStaleWhileUpdating apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *RulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStale) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rulesetGetResponseRulesRulesetsSetCacheSettingsRuleActionParametersServeStaleJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RulesetGetResponseRulesAction string

const (
	RulesetGetResponseRulesActionBlock            RulesetGetResponseRulesAction = "block"
	RulesetGetResponseRulesActionChallenge        RulesetGetResponseRulesAction = "challenge"
	RulesetGetResponseRulesActionCompressResponse RulesetGetResponseRulesAction = "compress_response"
	RulesetGetResponseRulesActionExecute          RulesetGetResponseRulesAction = "execute"
	RulesetGetResponseRulesActionJsChallenge      RulesetGetResponseRulesAction = "js_challenge"
	RulesetGetResponseRulesActionLog              RulesetGetResponseRulesAction = "log"
	RulesetGetResponseRulesActionManagedChallenge RulesetGetResponseRulesAction = "managed_challenge"
	RulesetGetResponseRulesActionRedirect         RulesetGetResponseRulesAction = "redirect"
	RulesetGetResponseRulesActionRewrite          RulesetGetResponseRulesAction = "rewrite"
	RulesetGetResponseRulesActionRoute            RulesetGetResponseRulesAction = "route"
	RulesetGetResponseRulesActionScore            RulesetGetResponseRulesAction = "score"
	RulesetGetResponseRulesActionServeError       RulesetGetResponseRulesAction = "serve_error"
	RulesetGetResponseRulesActionSetConfig        RulesetGetResponseRulesAction = "set_config"
	RulesetGetResponseRulesActionSkip             RulesetGetResponseRulesAction = "skip"
	RulesetGetResponseRulesActionSetCacheSettings RulesetGetResponseRulesAction = "set_cache_settings"
)

func (r RulesetGetResponseRulesAction) IsKnown() bool {
	switch r {
	case RulesetGetResponseRulesActionBlock, RulesetGetResponseRulesActionChallenge, RulesetGetResponseRulesActionCompressResponse, RulesetGetResponseRulesActionExecute, RulesetGetResponseRulesActionJsChallenge, RulesetGetResponseRulesActionLog, RulesetGetResponseRulesActionManagedChallenge, RulesetGetResponseRulesActionRedirect, RulesetGetResponseRulesActionRewrite, RulesetGetResponseRulesActionRoute, RulesetGetResponseRulesActionScore, RulesetGetResponseRulesActionServeError, RulesetGetResponseRulesActionSetConfig, RulesetGetResponseRulesActionSkip, RulesetGetResponseRulesActionSetCacheSettings:
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
	Rules param.Field[[]RulesetNewParamsRuleUnion] `json:"rules,required"`
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
	RulesetNewParamsPhaseDDoSL4                         RulesetNewParamsPhase = "ddos_l4"
	RulesetNewParamsPhaseDDoSL7                         RulesetNewParamsPhase = "ddos_l7"
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
	case RulesetNewParamsPhaseDDoSL4, RulesetNewParamsPhaseDDoSL7, RulesetNewParamsPhaseHTTPConfigSettings, RulesetNewParamsPhaseHTTPCustomErrors, RulesetNewParamsPhaseHTTPLogCustomFields, RulesetNewParamsPhaseHTTPRatelimit, RulesetNewParamsPhaseHTTPRequestCacheSettings, RulesetNewParamsPhaseHTTPRequestDynamicRedirect, RulesetNewParamsPhaseHTTPRequestFirewallCustom, RulesetNewParamsPhaseHTTPRequestFirewallManaged, RulesetNewParamsPhaseHTTPRequestLateTransform, RulesetNewParamsPhaseHTTPRequestOrigin, RulesetNewParamsPhaseHTTPRequestRedirect, RulesetNewParamsPhaseHTTPRequestSanitize, RulesetNewParamsPhaseHTTPRequestSbfm, RulesetNewParamsPhaseHTTPRequestSelectConfiguration, RulesetNewParamsPhaseHTTPRequestTransform, RulesetNewParamsPhaseHTTPResponseCompression, RulesetNewParamsPhaseHTTPResponseFirewallManaged, RulesetNewParamsPhaseHTTPResponseHeadersTransform, RulesetNewParamsPhaseMagicTransit, RulesetNewParamsPhaseMagicTransitIDsManaged, RulesetNewParamsPhaseMagicTransitManaged:
		return true
	}
	return false
}

type RulesetNewParamsRule struct {
	// The action to perform when the rule matches.
	Action           param.Field[RulesetNewParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                 `json:"action_parameters,required"`
	Categories       param.Field[interface{}]                 `json:"categories,required"`
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

func (r RulesetNewParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRule) implementsRulesetsRulesetNewParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam],
// [rulesets.RulesetNewParamsRulesRulesetsChallengeRule],
// [rulesets.RulesetNewParamsRulesRulesetsCompressResponseRule],
// [rulesets.ExecuteRuleParam],
// [rulesets.RulesetNewParamsRulesRulesetsJsChallengeRule],
// [rulesets.LogRuleParam],
// [rulesets.RulesetNewParamsRulesRulesetsManagedChallengeRule],
// [rulesets.RulesetNewParamsRulesRulesetsRedirectRule],
// [rulesets.RulesetNewParamsRulesRulesetsRewriteRule],
// [rulesets.RulesetNewParamsRulesRulesetsRouteRule],
// [rulesets.RulesetNewParamsRulesRulesetsScoreRule],
// [rulesets.RulesetNewParamsRulesRulesetsServeErrorRule],
// [rulesets.RulesetNewParamsRulesRulesetsSetConfigRule], [rulesets.SkipRuleParam],
// [rulesets.RulesetNewParamsRulesRulesetsSetCacheSettingsRule],
// [RulesetNewParamsRule].
type RulesetNewParamsRuleUnion interface {
	implementsRulesetsRulesetNewParamsRuleUnion()
}

type RulesetNewParamsRulesRulesetsChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsChallengeRuleAction] `json:"action"`
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

func (r RulesetNewParamsRulesRulesetsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsChallengeRule) implementsRulesetsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsChallengeRuleAction string

const (
	RulesetNewParamsRulesRulesetsChallengeRuleActionChallenge RulesetNewParamsRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetNewParamsRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

type RulesetNewParamsRulesRulesetsCompressResponseRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsCompressResponseRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsCompressResponseRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetNewParamsRulesRulesetsCompressResponseRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsCompressResponseRule) implementsRulesetsRulesetNewParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsCompressResponseRuleAction string

const (
	RulesetNewParamsRulesRulesetsCompressResponseRuleActionCompressResponse RulesetNewParamsRulesRulesetsCompressResponseRuleAction = "compress_response"
)

func (r RulesetNewParamsRulesRulesetsCompressResponseRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsCompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsCompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms param.Field[[]RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithm] `json:"algorithms"`
}

func (r RulesetNewParamsRulesRulesetsCompressResponseRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Compression algorithm to enable.
type RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name param.Field[RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName] `json:"name"`
}

func (r RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithm) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Name of compression algorithm to enable.
type RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName string

const (
	RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone    RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "none"
	RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto    RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "auto"
	RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "default"
	RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip    RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "gzip"
	RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli  RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "brotli"
)

func (r RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone, RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto, RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault, RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip, RulesetNewParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli:
		return true
	}
	return false
}

type RulesetNewParamsRulesRulesetsJsChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsJsChallengeRuleAction] `json:"action"`
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

func (r RulesetNewParamsRulesRulesetsJsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsJsChallengeRule) implementsRulesetsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsJsChallengeRuleAction string

const (
	RulesetNewParamsRulesRulesetsJsChallengeRuleActionJsChallenge RulesetNewParamsRulesRulesetsJsChallengeRuleAction = "js_challenge"
)

func (r RulesetNewParamsRulesRulesetsJsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsJsChallengeRuleActionJsChallenge:
		return true
	}
	return false
}

type RulesetNewParamsRulesRulesetsManagedChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsManagedChallengeRuleAction] `json:"action"`
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

func (r RulesetNewParamsRulesRulesetsManagedChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsManagedChallengeRule) implementsRulesetsRulesetNewParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsManagedChallengeRuleAction string

const (
	RulesetNewParamsRulesRulesetsManagedChallengeRuleActionManagedChallenge RulesetNewParamsRulesRulesetsManagedChallengeRuleAction = "managed_challenge"
)

func (r RulesetNewParamsRulesRulesetsManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

type RulesetNewParamsRulesRulesetsRedirectRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsRedirectRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsRedirectRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetNewParamsRulesRulesetsRedirectRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRedirectRule) implementsRulesetsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsRedirectRuleAction string

const (
	RulesetNewParamsRulesRulesetsRedirectRuleActionRedirect RulesetNewParamsRulesRulesetsRedirectRuleAction = "redirect"
)

func (r RulesetNewParamsRulesRulesetsRedirectRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsRedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsRedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList param.Field[RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromList] `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue param.Field[RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValue] `json:"from_value"`
}

func (r RulesetNewParamsRulesRulesetsRedirectRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Serve a redirect based on a bulk list lookup.
type RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key param.Field[string] `json:"key"`
	// The name of the list to match against.
	Name param.Field[string] `json:"name"`
}

func (r RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Serve a redirect based on the request properties.
type RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString param.Field[bool] `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode param.Field[RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL param.Field[RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion] `json:"target_url"`
}

func (r RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status code to be used for the redirect.
type RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode float64

const (
	RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301 RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 301
	RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302 RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 302
	RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303 RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 303
	RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307 RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 307
	RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308 RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 308
)

func (r RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301, RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302, RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303, RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307, RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURL struct {
	// The URL to redirect the request to.
	Value param.Field[string] `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression param.Field[string] `json:"expression"`
}

func (r RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) implementsRulesetsRulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion() {
}

// The URL to redirect the request to.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect],
// [rulesets.RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect],
// [RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURL].
type RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRulesetsRulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion()
}

type RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value param.Field[string] `json:"value"`
}

func (r RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRulesetsRulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion() {
}

type RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression param.Field[string] `json:"expression"`
}

func (r RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRulesetsRulesetNewParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion() {
}

type RulesetNewParamsRulesRulesetsRewriteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsRewriteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsRewriteRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetNewParamsRulesRulesetsRewriteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRewriteRule) implementsRulesetsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsRewriteRuleAction string

const (
	RulesetNewParamsRulesRulesetsRewriteRuleActionRewrite RulesetNewParamsRulesRulesetsRewriteRuleAction = "rewrite"
)

func (r RulesetNewParamsRulesRulesetsRewriteRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsRewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsRewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers param.Field[map[string]RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion] `json:"headers"`
	// URI to rewrite the request to.
	URI param.Field[RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURI] `json:"uri"`
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Remove the header from the request.
type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeaders struct {
	Operation param.Field[RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation] `json:"operation,required"`
	// Static value for the header.
	Value param.Field[string] `json:"value"`
	// Expression for the header value.
	Expression param.Field[string] `json:"expression"`
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeaders) implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion() {
}

// Remove the header from the request.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader],
// [rulesets.RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader],
// [rulesets.RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader],
// [RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeaders].
type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion interface {
	implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion()
}

// Remove the header from the request.
type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation param.Field[RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation] `json:"operation,required"`
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion() {
}

type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Set a request header with a static value.
type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader struct {
	Operation param.Field[RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation] `json:"operation,required"`
	// Static value for the header.
	Value param.Field[string] `json:"value,required"`
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion() {
}

type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation string

const (
	RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation = "set"
)

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader struct {
	// Expression for the header value.
	Expression param.Field[string]                                                                                `json:"expression,required"`
	Operation  param.Field[RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation] `json:"operation,required"`
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion() {
}

type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation string

const (
	RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation = "set"
)

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet:
		return true
	}
	return false
}

type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation string

const (
	RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation = "remove"
	RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersOperationSet    RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation = "set"
)

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove, RulesetNewParamsRulesRulesetsRewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path param.Field[RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion] `json:"path"`
	// Query portion rewrite.
	Query param.Field[RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion] `json:"query"`
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Path portion rewrite.
type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPath struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression"`
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPath) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPath) implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion() {
}

// Path portion rewrite.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathStaticValue],
// [rulesets.RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue],
// [RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPath].
type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion interface {
	implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion()
}

type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathStaticValue struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value,required"`
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion() {
}

type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression,required"`
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion() {
}

// Query portion rewrite.
type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQuery struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression"`
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQuery) implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion() {
}

// Query portion rewrite.
//
// Satisfied by
// [rulesets.RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue],
// [rulesets.RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue],
// [RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQuery].
type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion interface {
	implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion()
}

type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value,required"`
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion() {
}

type RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression,required"`
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) implementsRulesetsRulesetNewParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion() {
}

type RulesetNewParamsRulesRulesetsRouteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsRouteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsRouteRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetNewParamsRulesRulesetsRouteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsRouteRule) implementsRulesetsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsRouteRuleAction string

const (
	RulesetNewParamsRulesRulesetsRouteRuleActionRoute RulesetNewParamsRulesRulesetsRouteRuleAction = "route"
)

func (r RulesetNewParamsRulesRulesetsRouteRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsRouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsRouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader param.Field[string] `json:"host_header"`
	// Override the IP/TCP destination.
	Origin param.Field[RulesetNewParamsRulesRulesetsRouteRuleActionParametersOrigin] `json:"origin"`
	// Override the Server Name Indication (SNI).
	Sni param.Field[RulesetNewParamsRulesRulesetsRouteRuleActionParametersSni] `json:"sni"`
}

func (r RulesetNewParamsRulesRulesetsRouteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Override the IP/TCP destination.
type RulesetNewParamsRulesRulesetsRouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host param.Field[string] `json:"host"`
	// Override the destination port.
	Port param.Field[float64] `json:"port"`
}

func (r RulesetNewParamsRulesRulesetsRouteRuleActionParametersOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Override the Server Name Indication (SNI).
type RulesetNewParamsRulesRulesetsRouteRuleActionParametersSni struct {
	// The SNI override.
	Value param.Field[string] `json:"value,required"`
}

func (r RulesetNewParamsRulesRulesetsRouteRuleActionParametersSni) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetNewParamsRulesRulesetsScoreRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsScoreRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsScoreRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetNewParamsRulesRulesetsScoreRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsScoreRule) implementsRulesetsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsScoreRuleAction string

const (
	RulesetNewParamsRulesRulesetsScoreRuleActionScore RulesetNewParamsRulesRulesetsScoreRuleAction = "score"
)

func (r RulesetNewParamsRulesRulesetsScoreRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment param.Field[int64] `json:"increment"`
}

func (r RulesetNewParamsRulesRulesetsScoreRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetNewParamsRulesRulesetsServeErrorRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsServeErrorRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsServeErrorRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetNewParamsRulesRulesetsServeErrorRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsServeErrorRule) implementsRulesetsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsServeErrorRuleAction string

const (
	RulesetNewParamsRulesRulesetsServeErrorRuleActionServeError RulesetNewParamsRulesRulesetsServeErrorRuleAction = "serve_error"
)

func (r RulesetNewParamsRulesRulesetsServeErrorRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsServeErrorRuleActionParameters struct {
	// Error response content.
	Content param.Field[string] `json:"content"`
	// Content-type header to set with the response.
	ContentType param.Field[RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentType] `json:"content_type"`
	// The status code to use for the error.
	StatusCode param.Field[float64] `json:"status_code"`
}

func (r RulesetNewParamsRulesRulesetsServeErrorRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Content-type header to set with the response.
type RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentType string

const (
	RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentType = "application/json"
	RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml         RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentType = "text/xml"
	RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain       RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentType = "text/plain"
	RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML        RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentType = "text/html"
)

func (r RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson, RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml, RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain, RulesetNewParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

type RulesetNewParamsRulesRulesetsSetConfigRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsSetConfigRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsSetConfigRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetNewParamsRulesRulesetsSetConfigRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetConfigRule) implementsRulesetsRulesetNewParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsSetConfigRuleAction string

const (
	RulesetNewParamsRulesRulesetsSetConfigRuleActionSetConfig RulesetNewParamsRulesRulesetsSetConfigRuleAction = "set_config"
)

func (r RulesetNewParamsRulesRulesetsSetConfigRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsSetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites param.Field[bool] `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify param.Field[RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersAutominify] `json:"autominify"`
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
	Polish param.Field[RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersPolish] `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader param.Field[bool] `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel param.Field[RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel] `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes param.Field[bool] `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL param.Field[RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSL] `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg param.Field[bool] `json:"sxg"`
}

func (r RulesetNewParamsRulesRulesetsSetConfigRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Select which file extensions to minify automatically.
type RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	Css param.Field[bool] `json:"css"`
	// Minify HTML files.
	HTML param.Field[bool] `json:"html"`
	// Minify JS files.
	Js param.Field[bool] `json:"js"`
}

func (r RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersAutominify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure the Polish level.
type RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersPolish string

const (
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersPolishOff      RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersPolish = "off"
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersPolishLossless RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersPolish = "lossless"
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersPolishLossy    RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersPolish = "lossy"
)

func (r RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersPolishOff, RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersPolishLossless, RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersPolishLossy:
		return true
	}
	return false
}

// Configure the Security Level.
type RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel string

const (
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff            RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "off"
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "essentially_off"
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow            RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "low"
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium         RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "medium"
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh           RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "high"
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack    RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff, RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff, RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow, RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium, RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh, RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSL string

const (
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSLOff        RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSL = "off"
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSLFlexible   RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSL = "flexible"
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSLFull       RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSL = "full"
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSLStrict     RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSL = "strict"
	RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSLOriginPull RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSL = "origin_pull"
)

func (r RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSLOff, RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSLFlexible, RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSLFull, RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSLStrict, RulesetNewParamsRulesRulesetsSetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

type RulesetNewParamsRulesRulesetsSetCacheSettingsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRule) implementsRulesetsRulesetNewParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleAction string

const (
	RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings RulesetNewParamsRulesRulesetsSetCacheSettingsRuleAction = "set_cache_settings"
)

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts param.Field[[]int64] `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL] `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache param.Field[bool] `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey] `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve] `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL] `json:"edge_ttl"`
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
	ServeStale param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersServeStale] `json:"serve_stale"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode] `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default param.Field[int64] `json:"default"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which browser ttl mode to use.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
)

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType param.Field[bool] `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor param.Field[bool] `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey] `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder param.Field[bool] `json:"ignore_query_strings_order"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Customize which components of the request are included or excluded from the
// cache key.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie] `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader] `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost] `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString] `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser] `json:"user"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cookies to include in building the cache key.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// Include these cookies' names and their values.
	Include param.Field[[]string] `json:"include"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The header names and values to include in building the cache key.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin param.Field[bool] `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include param.Field[[]string] `json:"include"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to use the original host or the resolved host in the cache key.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved param.Field[bool] `json:"resolved"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude] `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude] `json:"include"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Exclude all query string parameters from use in building the cache key.
	All param.Field[bool] `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List param.Field[[]string] `json:"list"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Use all query string parameters in the cache key.
	All param.Field[bool] `json:"all"`
	// A list of query string parameters used to build the cache key.
	List param.Field[[]string] `json:"list"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Characteristics of the request user agent used in building the cache key.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType param.Field[bool] `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo param.Field[bool] `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang param.Field[bool] `json:"lang"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible param.Field[bool] `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize param.Field[int64] `json:"min_file_size,required"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default param.Field[int64] `json:"default,required"`
	// edge ttl options
	Mode param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode] `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL param.Field[[]RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL] `json:"status_code_ttl,required"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// edge ttl options
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value param.Field[int64] `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange param.Field[RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange] `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue param.Field[int64] `json:"status_code_value"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The range of status codes used to apply the selected mode.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// response status code lower bound
	From param.Field[int64] `json:"from,required"`
	// response status code upper bound
	To param.Field[int64] `json:"to,required"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating param.Field[bool] `json:"disable_stale_while_updating,required"`
}

func (r RulesetNewParamsRulesRulesetsSetCacheSettingsRuleActionParametersServeStale) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type RulesetNewParamsRulesAction string

const (
	RulesetNewParamsRulesActionBlock            RulesetNewParamsRulesAction = "block"
	RulesetNewParamsRulesActionChallenge        RulesetNewParamsRulesAction = "challenge"
	RulesetNewParamsRulesActionCompressResponse RulesetNewParamsRulesAction = "compress_response"
	RulesetNewParamsRulesActionExecute          RulesetNewParamsRulesAction = "execute"
	RulesetNewParamsRulesActionJsChallenge      RulesetNewParamsRulesAction = "js_challenge"
	RulesetNewParamsRulesActionLog              RulesetNewParamsRulesAction = "log"
	RulesetNewParamsRulesActionManagedChallenge RulesetNewParamsRulesAction = "managed_challenge"
	RulesetNewParamsRulesActionRedirect         RulesetNewParamsRulesAction = "redirect"
	RulesetNewParamsRulesActionRewrite          RulesetNewParamsRulesAction = "rewrite"
	RulesetNewParamsRulesActionRoute            RulesetNewParamsRulesAction = "route"
	RulesetNewParamsRulesActionScore            RulesetNewParamsRulesAction = "score"
	RulesetNewParamsRulesActionServeError       RulesetNewParamsRulesAction = "serve_error"
	RulesetNewParamsRulesActionSetConfig        RulesetNewParamsRulesAction = "set_config"
	RulesetNewParamsRulesActionSkip             RulesetNewParamsRulesAction = "skip"
	RulesetNewParamsRulesActionSetCacheSettings RulesetNewParamsRulesAction = "set_cache_settings"
)

func (r RulesetNewParamsRulesAction) IsKnown() bool {
	switch r {
	case RulesetNewParamsRulesActionBlock, RulesetNewParamsRulesActionChallenge, RulesetNewParamsRulesActionCompressResponse, RulesetNewParamsRulesActionExecute, RulesetNewParamsRulesActionJsChallenge, RulesetNewParamsRulesActionLog, RulesetNewParamsRulesActionManagedChallenge, RulesetNewParamsRulesActionRedirect, RulesetNewParamsRulesActionRewrite, RulesetNewParamsRulesActionRoute, RulesetNewParamsRulesActionScore, RulesetNewParamsRulesActionServeError, RulesetNewParamsRulesActionSetConfig, RulesetNewParamsRulesActionSkip, RulesetNewParamsRulesActionSetCacheSettings:
		return true
	}
	return false
}

// A response object.
type RulesetNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RulesetNewResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RulesetNewResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
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

func (r RulesetNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RulesetNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RulesetUpdateParams struct {
	// The list of rules in the ruleset.
	Rules param.Field[[]RulesetUpdateParamsRuleUnion] `json:"rules,required"`
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

type RulesetUpdateParamsRule struct {
	// The action to perform when the rule matches.
	Action           param.Field[RulesetUpdateParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                    `json:"action_parameters,required"`
	Categories       param.Field[interface{}]                    `json:"categories,required"`
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

func (r RulesetUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam],
// [rulesets.RulesetUpdateParamsRulesRulesetsChallengeRule],
// [rulesets.RulesetUpdateParamsRulesRulesetsCompressResponseRule],
// [rulesets.ExecuteRuleParam],
// [rulesets.RulesetUpdateParamsRulesRulesetsJsChallengeRule],
// [rulesets.LogRuleParam],
// [rulesets.RulesetUpdateParamsRulesRulesetsManagedChallengeRule],
// [rulesets.RulesetUpdateParamsRulesRulesetsRedirectRule],
// [rulesets.RulesetUpdateParamsRulesRulesetsRewriteRule],
// [rulesets.RulesetUpdateParamsRulesRulesetsRouteRule],
// [rulesets.RulesetUpdateParamsRulesRulesetsScoreRule],
// [rulesets.RulesetUpdateParamsRulesRulesetsServeErrorRule],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetConfigRule],
// [rulesets.SkipRuleParam],
// [rulesets.RulesetUpdateParamsRulesRulesetsSetCacheSettingsRule],
// [RulesetUpdateParamsRule].
type RulesetUpdateParamsRuleUnion interface {
	implementsRulesetsRulesetUpdateParamsRuleUnion()
}

type RulesetUpdateParamsRulesRulesetsChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsChallengeRuleAction] `json:"action"`
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

func (r RulesetUpdateParamsRulesRulesetsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsChallengeRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsChallengeRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsChallengeRuleActionChallenge RulesetUpdateParamsRulesRulesetsChallengeRuleAction = "challenge"
)

func (r RulesetUpdateParamsRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

type RulesetUpdateParamsRulesRulesetsCompressResponseRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsCompressResponseRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetUpdateParamsRulesRulesetsCompressResponseRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsCompressResponseRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsCompressResponseRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionCompressResponse RulesetUpdateParamsRulesRulesetsCompressResponseRuleAction = "compress_response"
)

func (r RulesetUpdateParamsRulesRulesetsCompressResponseRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms param.Field[[]RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithm] `json:"algorithms"`
}

func (r RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Compression algorithm to enable.
type RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name param.Field[RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName] `json:"name"`
}

func (r RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithm) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Name of compression algorithm to enable.
type RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName string

const (
	RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone    RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "none"
	RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto    RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "auto"
	RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "default"
	RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip    RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "gzip"
	RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli  RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName = "brotli"
)

func (r RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameNone, RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameAuto, RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameDefault, RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameGzip, RulesetUpdateParamsRulesRulesetsCompressResponseRuleActionParametersAlgorithmsNameBrotli:
		return true
	}
	return false
}

type RulesetUpdateParamsRulesRulesetsJsChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsJsChallengeRuleAction] `json:"action"`
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

func (r RulesetUpdateParamsRulesRulesetsJsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsJsChallengeRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsJsChallengeRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsJsChallengeRuleActionJsChallenge RulesetUpdateParamsRulesRulesetsJsChallengeRuleAction = "js_challenge"
)

func (r RulesetUpdateParamsRulesRulesetsJsChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsJsChallengeRuleActionJsChallenge:
		return true
	}
	return false
}

type RulesetUpdateParamsRulesRulesetsManagedChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsManagedChallengeRuleAction] `json:"action"`
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

func (r RulesetUpdateParamsRulesRulesetsManagedChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsManagedChallengeRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsManagedChallengeRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsManagedChallengeRuleActionManagedChallenge RulesetUpdateParamsRulesRulesetsManagedChallengeRuleAction = "managed_challenge"
)

func (r RulesetUpdateParamsRulesRulesetsManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

type RulesetUpdateParamsRulesRulesetsRedirectRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsRedirectRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsRedirectRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetUpdateParamsRulesRulesetsRedirectRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRedirectRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsRedirectRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsRedirectRuleActionRedirect RulesetUpdateParamsRulesRulesetsRedirectRuleAction = "redirect"
)

func (r RulesetUpdateParamsRulesRulesetsRedirectRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsRedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsRedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList param.Field[RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromList] `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue param.Field[RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValue] `json:"from_value"`
}

func (r RulesetUpdateParamsRulesRulesetsRedirectRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Serve a redirect based on a bulk list lookup.
type RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key param.Field[string] `json:"key"`
	// The name of the list to match against.
	Name param.Field[string] `json:"name"`
}

func (r RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromList) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Serve a redirect based on the request properties.
type RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString param.Field[bool] `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode param.Field[RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL param.Field[RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion] `json:"target_url"`
}

func (r RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status code to be used for the redirect.
type RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode float64

const (
	RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301 RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 301
	RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302 RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 302
	RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303 RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 303
	RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307 RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 307
	RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308 RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode = 308
)

func (r RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode301, RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode302, RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode303, RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode307, RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURL struct {
	// The URL to redirect the request to.
	Value param.Field[string] `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression param.Field[string] `json:"expression"`
}

func (r RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURL) implementsRulesetsRulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion() {
}

// The URL to redirect the request to.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect],
// [rulesets.RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect],
// [RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURL].
type RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRulesetsRulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion()
}

type RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value param.Field[string] `json:"value"`
}

func (r RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRulesetsRulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion() {
}

type RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression param.Field[string] `json:"expression"`
}

func (r RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRulesetsRulesetUpdateParamsRulesRulesetsRedirectRuleActionParametersFromValueTargetURLUnion() {
}

type RulesetUpdateParamsRulesRulesetsRewriteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsRewriteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsRewriteRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetUpdateParamsRulesRulesetsRewriteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsRewriteRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsRewriteRuleActionRewrite RulesetUpdateParamsRulesRulesetsRewriteRuleAction = "rewrite"
)

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsRewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers param.Field[map[string]RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion] `json:"headers"`
	// URI to rewrite the request to.
	URI param.Field[RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURI] `json:"uri"`
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Remove the header from the request.
type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeaders struct {
	Operation param.Field[RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation] `json:"operation,required"`
	// Static value for the header.
	Value param.Field[string] `json:"value"`
	// Expression for the header value.
	Expression param.Field[string] `json:"expression"`
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeaders) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeaders) implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion() {
}

// Remove the header from the request.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader],
// [rulesets.RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader],
// [rulesets.RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader],
// [RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeaders].
type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion interface {
	implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion()
}

// Remove the header from the request.
type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation param.Field[RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation] `json:"operation,required"`
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeader) implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion() {
}

type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Set a request header with a static value.
type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader struct {
	Operation param.Field[RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation] `json:"operation,required"`
	// Static value for the header.
	Value param.Field[string] `json:"value,required"`
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeader) implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion() {
}

type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation string

const (
	RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation = "set"
)

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersStaticHeaderOperationSet:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader struct {
	// Expression for the header value.
	Expression param.Field[string]                                                                                   `json:"expression,required"`
	Operation  param.Field[RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation] `json:"operation,required"`
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeader) implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersUnion() {
}

type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation string

const (
	RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation = "set"
)

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersDynamicHeaderOperationSet:
		return true
	}
	return false
}

type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation string

const (
	RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation = "remove"
	RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperationSet    RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation = "set"
)

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperationRemove, RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path param.Field[RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion] `json:"path"`
	// Query portion rewrite.
	Query param.Field[RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion] `json:"query"`
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURI) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Path portion rewrite.
type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPath struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression"`
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPath) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPath) implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion() {
}

// Path portion rewrite.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathStaticValue],
// [rulesets.RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue],
// [RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPath].
type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion interface {
	implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion()
}

type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathStaticValue struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value,required"`
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathStaticValue) implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion() {
}

type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression,required"`
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathDynamicValue) implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIPathUnion() {
}

// Query portion rewrite.
type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQuery struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression"`
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQuery) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQuery) implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion() {
}

// Query portion rewrite.
//
// Satisfied by
// [rulesets.RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue],
// [rulesets.RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue],
// [RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQuery].
type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion interface {
	implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion()
}

type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value,required"`
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryStaticValue) implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion() {
}

type RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression,required"`
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryDynamicValue) implementsRulesetsRulesetUpdateParamsRulesRulesetsRewriteRuleActionParametersURIQueryUnion() {
}

type RulesetUpdateParamsRulesRulesetsRouteRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsRouteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsRouteRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetUpdateParamsRulesRulesetsRouteRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsRouteRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsRouteRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsRouteRuleActionRoute RulesetUpdateParamsRulesRulesetsRouteRuleAction = "route"
)

func (r RulesetUpdateParamsRulesRulesetsRouteRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsRouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsRouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader param.Field[string] `json:"host_header"`
	// Override the IP/TCP destination.
	Origin param.Field[RulesetUpdateParamsRulesRulesetsRouteRuleActionParametersOrigin] `json:"origin"`
	// Override the Server Name Indication (SNI).
	Sni param.Field[RulesetUpdateParamsRulesRulesetsRouteRuleActionParametersSni] `json:"sni"`
}

func (r RulesetUpdateParamsRulesRulesetsRouteRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Override the IP/TCP destination.
type RulesetUpdateParamsRulesRulesetsRouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host param.Field[string] `json:"host"`
	// Override the destination port.
	Port param.Field[float64] `json:"port"`
}

func (r RulesetUpdateParamsRulesRulesetsRouteRuleActionParametersOrigin) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Override the Server Name Indication (SNI).
type RulesetUpdateParamsRulesRulesetsRouteRuleActionParametersSni struct {
	// The SNI override.
	Value param.Field[string] `json:"value,required"`
}

func (r RulesetUpdateParamsRulesRulesetsRouteRuleActionParametersSni) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetUpdateParamsRulesRulesetsScoreRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsScoreRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsScoreRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetUpdateParamsRulesRulesetsScoreRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsScoreRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsScoreRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsScoreRuleActionScore RulesetUpdateParamsRulesRulesetsScoreRuleAction = "score"
)

func (r RulesetUpdateParamsRulesRulesetsScoreRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment param.Field[int64] `json:"increment"`
}

func (r RulesetUpdateParamsRulesRulesetsScoreRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RulesetUpdateParamsRulesRulesetsServeErrorRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsServeErrorRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetUpdateParamsRulesRulesetsServeErrorRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsServeErrorRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsServeErrorRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsServeErrorRuleActionServeError RulesetUpdateParamsRulesRulesetsServeErrorRuleAction = "serve_error"
)

func (r RulesetUpdateParamsRulesRulesetsServeErrorRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParameters struct {
	// Error response content.
	Content param.Field[string] `json:"content"`
	// Content-type header to set with the response.
	ContentType param.Field[RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType] `json:"content_type"`
	// The status code to use for the error.
	StatusCode param.Field[float64] `json:"status_code"`
}

func (r RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Content-type header to set with the response.
type RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType string

const (
	RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType = "application/json"
	RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml         RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType = "text/xml"
	RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain       RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType = "text/plain"
	RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML        RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType = "text/html"
)

func (r RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeApplicationJson, RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextXml, RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextPlain, RulesetUpdateParamsRulesRulesetsServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

type RulesetUpdateParamsRulesRulesetsSetConfigRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsSetConfigRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetUpdateParamsRulesRulesetsSetConfigRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetConfigRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsSetConfigRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionSetConfig RulesetUpdateParamsRulesRulesetsSetConfigRuleAction = "set_config"
)

func (r RulesetUpdateParamsRulesRulesetsSetConfigRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites param.Field[bool] `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify param.Field[RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersAutominify] `json:"autominify"`
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
	Polish param.Field[RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolish] `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader param.Field[bool] `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel param.Field[RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel] `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes param.Field[bool] `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL param.Field[RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL] `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg param.Field[bool] `json:"sxg"`
}

func (r RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Select which file extensions to minify automatically.
type RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	Css param.Field[bool] `json:"css"`
	// Minify HTML files.
	HTML param.Field[bool] `json:"html"`
	// Minify JS files.
	Js param.Field[bool] `json:"js"`
}

func (r RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersAutominify) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configure the Polish level.
type RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolish string

const (
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolishOff      RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolish = "off"
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolishLossless RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolish = "lossless"
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolishLossy    RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolish = "lossy"
)

func (r RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolishOff, RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolishLossless, RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersPolishLossy:
		return true
	}
	return false
}

// Configure the Security Level.
type RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel string

const (
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff            RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "off"
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "essentially_off"
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow            RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "low"
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium         RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "medium"
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh           RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "high"
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack    RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelOff, RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelEssentiallyOff, RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelLow, RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelMedium, RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelHigh, RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL string

const (
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLOff        RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL = "off"
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLFlexible   RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL = "flexible"
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLFull       RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL = "full"
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLStrict     RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL = "strict"
	RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLOriginPull RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL = "origin_pull"
)

func (r RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLOff, RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLFlexible, RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLFull, RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLStrict, RulesetUpdateParamsRulesRulesetsSetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParameters] `json:"action_parameters"`
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

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRule) implementsRulesetsRulesetUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleAction string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleAction = "set_cache_settings"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts param.Field[[]int64] `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL] `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache param.Field[bool] `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey] `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve] `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL] `json:"edge_ttl"`
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
	ServeStale param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersServeStale] `json:"serve_stale"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode] `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default param.Field[int64] `json:"default"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Determines which browser ttl mode to use.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType param.Field[bool] `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor param.Field[bool] `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey] `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder param.Field[bool] `json:"ignore_query_strings_order"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Customize which components of the request are included or excluded from the
// cache key.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie] `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader] `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost] `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString] `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser] `json:"user"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKey) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cookies to include in building the cache key.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// Include these cookies' names and their values.
	Include param.Field[[]string] `json:"include"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The header names and values to include in building the cache key.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin param.Field[bool] `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include param.Field[[]string] `json:"include"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to use the original host or the resolved host in the cache key.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved param.Field[bool] `json:"resolved"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude] `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude] `json:"include"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Exclude all query string parameters from use in building the cache key.
	All param.Field[bool] `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List param.Field[[]string] `json:"list"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Use all query string parameters in the cache key.
	All param.Field[bool] `json:"all"`
	// A list of query string parameters used to build the cache key.
	List param.Field[[]string] `json:"list"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Characteristics of the request user agent used in building the cache key.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType param.Field[bool] `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo param.Field[bool] `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang param.Field[bool] `json:"lang"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible param.Field[bool] `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize param.Field[int64] `json:"min_file_size,required"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersCacheReserve) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default param.Field[int64] `json:"default,required"`
	// edge ttl options
	Mode param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode] `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL param.Field[[]RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL] `json:"status_code_ttl,required"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// edge ttl options
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value param.Field[int64] `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange param.Field[RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange] `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue param.Field[int64] `json:"status_code_value"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The range of status codes used to apply the selected mode.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// response status code lower bound
	From param.Field[int64] `json:"from,required"`
	// response status code upper bound
	To param.Field[int64] `json:"to,required"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating param.Field[bool] `json:"disable_stale_while_updating,required"`
}

func (r RulesetUpdateParamsRulesRulesetsSetCacheSettingsRuleActionParametersServeStale) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type RulesetUpdateParamsRulesAction string

const (
	RulesetUpdateParamsRulesActionBlock            RulesetUpdateParamsRulesAction = "block"
	RulesetUpdateParamsRulesActionChallenge        RulesetUpdateParamsRulesAction = "challenge"
	RulesetUpdateParamsRulesActionCompressResponse RulesetUpdateParamsRulesAction = "compress_response"
	RulesetUpdateParamsRulesActionExecute          RulesetUpdateParamsRulesAction = "execute"
	RulesetUpdateParamsRulesActionJsChallenge      RulesetUpdateParamsRulesAction = "js_challenge"
	RulesetUpdateParamsRulesActionLog              RulesetUpdateParamsRulesAction = "log"
	RulesetUpdateParamsRulesActionManagedChallenge RulesetUpdateParamsRulesAction = "managed_challenge"
	RulesetUpdateParamsRulesActionRedirect         RulesetUpdateParamsRulesAction = "redirect"
	RulesetUpdateParamsRulesActionRewrite          RulesetUpdateParamsRulesAction = "rewrite"
	RulesetUpdateParamsRulesActionRoute            RulesetUpdateParamsRulesAction = "route"
	RulesetUpdateParamsRulesActionScore            RulesetUpdateParamsRulesAction = "score"
	RulesetUpdateParamsRulesActionServeError       RulesetUpdateParamsRulesAction = "serve_error"
	RulesetUpdateParamsRulesActionSetConfig        RulesetUpdateParamsRulesAction = "set_config"
	RulesetUpdateParamsRulesActionSkip             RulesetUpdateParamsRulesAction = "skip"
	RulesetUpdateParamsRulesActionSetCacheSettings RulesetUpdateParamsRulesAction = "set_cache_settings"
)

func (r RulesetUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case RulesetUpdateParamsRulesActionBlock, RulesetUpdateParamsRulesActionChallenge, RulesetUpdateParamsRulesActionCompressResponse, RulesetUpdateParamsRulesActionExecute, RulesetUpdateParamsRulesActionJsChallenge, RulesetUpdateParamsRulesActionLog, RulesetUpdateParamsRulesActionManagedChallenge, RulesetUpdateParamsRulesActionRedirect, RulesetUpdateParamsRulesActionRewrite, RulesetUpdateParamsRulesActionRoute, RulesetUpdateParamsRulesActionScore, RulesetUpdateParamsRulesActionServeError, RulesetUpdateParamsRulesActionSetConfig, RulesetUpdateParamsRulesActionSkip, RulesetUpdateParamsRulesActionSetCacheSettings:
		return true
	}
	return false
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
	RulesetUpdateParamsPhaseDDoSL4                         RulesetUpdateParamsPhase = "ddos_l4"
	RulesetUpdateParamsPhaseDDoSL7                         RulesetUpdateParamsPhase = "ddos_l7"
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
	case RulesetUpdateParamsPhaseDDoSL4, RulesetUpdateParamsPhaseDDoSL7, RulesetUpdateParamsPhaseHTTPConfigSettings, RulesetUpdateParamsPhaseHTTPCustomErrors, RulesetUpdateParamsPhaseHTTPLogCustomFields, RulesetUpdateParamsPhaseHTTPRatelimit, RulesetUpdateParamsPhaseHTTPRequestCacheSettings, RulesetUpdateParamsPhaseHTTPRequestDynamicRedirect, RulesetUpdateParamsPhaseHTTPRequestFirewallCustom, RulesetUpdateParamsPhaseHTTPRequestFirewallManaged, RulesetUpdateParamsPhaseHTTPRequestLateTransform, RulesetUpdateParamsPhaseHTTPRequestOrigin, RulesetUpdateParamsPhaseHTTPRequestRedirect, RulesetUpdateParamsPhaseHTTPRequestSanitize, RulesetUpdateParamsPhaseHTTPRequestSbfm, RulesetUpdateParamsPhaseHTTPRequestSelectConfiguration, RulesetUpdateParamsPhaseHTTPRequestTransform, RulesetUpdateParamsPhaseHTTPResponseCompression, RulesetUpdateParamsPhaseHTTPResponseFirewallManaged, RulesetUpdateParamsPhaseHTTPResponseHeadersTransform, RulesetUpdateParamsPhaseMagicTransit, RulesetUpdateParamsPhaseMagicTransitIDsManaged, RulesetUpdateParamsPhaseMagicTransitManaged:
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
	// A ruleset object.
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
	// A ruleset object.
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

func (r RulesetGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RulesetGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
