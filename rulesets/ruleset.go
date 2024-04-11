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

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JsChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule] or [rulesets.SetCacheSettingsRule].
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
			Type:               reflect.TypeOf(ChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(JsChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
	)
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

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JsChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule] or [rulesets.SetCacheSettingsRule].
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
			Type:               reflect.TypeOf(ChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(JsChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
	)
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

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JsChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule] or [rulesets.SetCacheSettingsRule].
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
			Type:               reflect.TypeOf(ChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(CompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(JsChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(RouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetConfigRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
	)
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

// Satisfied by [rulesets.BlockRuleParam], [rulesets.ChallengeRuleParam],
// [rulesets.CompressResponseRuleParam], [rulesets.ExecuteRuleParam],
// [rulesets.JsChallengeRuleParam], [rulesets.LogRuleParam],
// [rulesets.ManagedChallengeRuleParam], [rulesets.RedirectRuleParam],
// [rulesets.RewriteRuleParam], [rulesets.RouteRuleParam],
// [rulesets.ScoreRuleParam], [rulesets.ServeErrorRuleParam],
// [rulesets.SetConfigRuleParam], [rulesets.SkipRuleParam],
// [rulesets.SetCacheSettingsRuleParam], [RulesetNewParamsRule].
type RulesetNewParamsRuleUnion interface {
	implementsRulesetsRulesetNewParamsRuleUnion()
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

// Satisfied by [rulesets.BlockRuleParam], [rulesets.ChallengeRuleParam],
// [rulesets.CompressResponseRuleParam], [rulesets.ExecuteRuleParam],
// [rulesets.JsChallengeRuleParam], [rulesets.LogRuleParam],
// [rulesets.ManagedChallengeRuleParam], [rulesets.RedirectRuleParam],
// [rulesets.RewriteRuleParam], [rulesets.RouteRuleParam],
// [rulesets.ScoreRuleParam], [rulesets.ServeErrorRuleParam],
// [rulesets.SetConfigRuleParam], [rulesets.SkipRuleParam],
// [rulesets.SetCacheSettingsRuleParam], [RulesetUpdateParamsRule].
type RulesetUpdateParamsRuleUnion interface {
	implementsRulesetsRulesetUpdateParamsRuleUnion()
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
