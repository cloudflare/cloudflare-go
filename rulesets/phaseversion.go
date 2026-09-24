// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/tidwall/gjson"
)

// PhaseVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhaseVersionService] method instead.
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
func (r *PhaseVersionService) List(ctx context.Context, rulesetPhase Phase, query PhaseVersionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[PhaseVersionListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint/versions", accountOrZone, accountOrZoneID, rulesetPhase)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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
func (r *PhaseVersionService) ListAutoPaging(ctx context.Context, rulesetPhase Phase, query PhaseVersionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[PhaseVersionListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, rulesetPhase, query, opts...))
}

// Fetches a specific version of an account or zone entry point ruleset.
func (r *PhaseVersionService) Get(ctx context.Context, rulesetPhase Phase, rulesetVersion string, query PhaseVersionGetParams, opts ...option.RequestOption) (res *PhaseVersionGetResponse, err error) {
	var env PhaseVersionGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Value != "" && query.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if query.AccountID.Value == "" && query.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if query.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	}
	if query.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint/versions/%s", accountOrZone, accountOrZoneID, rulesetPhase, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A ruleset object.
type PhaseVersionListResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id" api:"required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind" api:"required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name" api:"required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase" api:"required"`
	// The version of the ruleset.
	Version string `json:"version" api:"required"`
	// An informative description of the ruleset.
	Description string                       `json:"description"`
	JSON        phaseVersionListResponseJSON `json:"-"`
}

// phaseVersionListResponseJSON contains the JSON metadata for the struct
// [PhaseVersionListResponse]
type phaseVersionListResponseJSON struct {
	ID          apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Version     apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionListResponseJSON) RawJSON() string {
	return r.raw
}

// A ruleset object.
type PhaseVersionGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id" api:"required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind" api:"required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name" api:"required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase" api:"required"`
	// The list of rules in the ruleset.
	Rules []PhaseVersionGetResponseRule `json:"rules" api:"required"`
	// The version of the ruleset.
	Version string `json:"version" api:"required"`
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

type PhaseVersionGetResponseRule struct {
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action string `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [LogCustomFieldRuleActionParameters],
	// [RedirectRuleActionParameters], [RewriteRuleActionParameters],
	// [RouteRuleActionParameters], [ScoreRuleActionParameters],
	// [ServeErrorRuleActionParameters],
	// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [PhaseVersionGetResponseRulesTransformResponseHTMLRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseRulesChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseRulesSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [PhaseVersionGetResponseRulesChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [PhaseVersionGetResponseRulesJavaScriptChallengeRuleRatelimit],
	// [LogRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit],
	// [PhaseVersionGetResponseRulesSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [PhaseVersionGetResponseRulesSetCacheTagsRuleRatelimit],
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [PhaseVersionGetResponseRulesTransformResponseHTMLRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                          `json:"version"`
	JSON    phaseVersionGetResponseRuleJSON `json:"-"`
	union   PhaseVersionGetResponseRulesUnion
}

// phaseVersionGetResponseRuleJSON contains the JSON metadata for the struct
// [PhaseVersionGetResponseRule]
type phaseVersionGetResponseRuleJSON struct {
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	LastUpdated            apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	Version                apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r phaseVersionGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PhaseVersionGetResponseRulesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [PhaseVersionGetResponseRulesBlockRule],
// [PhaseVersionGetResponseRulesChallengeRule],
// [PhaseVersionGetResponseRulesResponseCompressionRule],
// [PhaseVersionGetResponseRulesDDoSDynamicRule],
// [PhaseVersionGetResponseRulesExecuteRule],
// [PhaseVersionGetResponseRulesForceConnectionCloseRule],
// [PhaseVersionGetResponseRulesJavaScriptChallengeRule],
// [PhaseVersionGetResponseRulesLogRule],
// [PhaseVersionGetResponseRulesLogCustomFieldRule],
// [PhaseVersionGetResponseRulesManagedChallengeRule],
// [PhaseVersionGetResponseRulesRedirectRule],
// [PhaseVersionGetResponseRulesRewriteRule],
// [PhaseVersionGetResponseRulesRouteRule],
// [PhaseVersionGetResponseRulesScoreRule],
// [PhaseVersionGetResponseRulesServeErrorRule],
// [PhaseVersionGetResponseRulesSetCacheControlRule],
// [PhaseVersionGetResponseRulesSetCacheSettingsRule],
// [PhaseVersionGetResponseRulesSetCacheTagsRule],
// [PhaseVersionGetResponseRulesSetConfigurationRule],
// [PhaseVersionGetResponseRulesSkipRule],
// [PhaseVersionGetResponseRulesTransformResponseHTMLRule].
func (r PhaseVersionGetResponseRule) AsUnion() PhaseVersionGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [PhaseVersionGetResponseRulesBlockRule],
// [PhaseVersionGetResponseRulesChallengeRule],
// [PhaseVersionGetResponseRulesResponseCompressionRule],
// [PhaseVersionGetResponseRulesDDoSDynamicRule],
// [PhaseVersionGetResponseRulesExecuteRule],
// [PhaseVersionGetResponseRulesForceConnectionCloseRule],
// [PhaseVersionGetResponseRulesJavaScriptChallengeRule],
// [PhaseVersionGetResponseRulesLogRule],
// [PhaseVersionGetResponseRulesLogCustomFieldRule],
// [PhaseVersionGetResponseRulesManagedChallengeRule],
// [PhaseVersionGetResponseRulesRedirectRule],
// [PhaseVersionGetResponseRulesRewriteRule],
// [PhaseVersionGetResponseRulesRouteRule],
// [PhaseVersionGetResponseRulesScoreRule],
// [PhaseVersionGetResponseRulesServeErrorRule],
// [PhaseVersionGetResponseRulesSetCacheControlRule],
// [PhaseVersionGetResponseRulesSetCacheSettingsRule],
// [PhaseVersionGetResponseRulesSetCacheTagsRule],
// [PhaseVersionGetResponseRulesSetConfigurationRule],
// [PhaseVersionGetResponseRulesSkipRule] or
// [PhaseVersionGetResponseRulesTransformResponseHTMLRule].
type PhaseVersionGetResponseRulesUnion interface {
	implementsPhaseVersionGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesResponseCompressionRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesDDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesJavaScriptChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheTagsRule{}),
			DiscriminatorValue: "set_cache_tags",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesSetConfigurationRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesSkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesTransformResponseHTMLRule{}),
			DiscriminatorValue: "transform_response_html",
		},
	)
}

type PhaseVersionGetResponseRulesBlockRule struct {
	ID         string                                    `json:"id" api:"required"`
	Action     string                                    `json:"action" api:"required"`
	Enabled    bool                                      `json:"enabled" api:"required"`
	Expression string                                    `json:"expression" api:"required"`
	Ref        string                                    `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesBlockRuleJSON `json:"-"`
	BlockRule
}

// phaseVersionGetResponseRulesBlockRuleJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseRulesBlockRule]
type phaseVersionGetResponseRulesBlockRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesBlockRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesChallengeRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesChallengeRuleAction `json:"action" api:"required"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled" api:"required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression" api:"required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref" api:"required"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseVersionGetResponseRulesChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseRulesChallengeRuleRatelimit `json:"ratelimit"`
	JSON      phaseVersionGetResponseRulesChallengeRuleJSON      `json:"-"`
}

// phaseVersionGetResponseRulesChallengeRuleJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseRulesChallengeRule]
type phaseVersionGetResponseRulesChallengeRuleJSON struct {
	ID                     apijson.Field
	Action                 apijson.Field
	Enabled                apijson.Field
	Expression             apijson.Field
	LastUpdated            apijson.Field
	Ref                    apijson.Field
	Version                apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	ExposedCredentialCheck apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesChallengeRule) implementsPhaseVersionGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesChallengeRuleAction string

const (
	PhaseVersionGetResponseRulesChallengeRuleActionChallenge PhaseVersionGetResponseRulesChallengeRuleAction = "challenge"
)

func (r PhaseVersionGetResponseRulesChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseRulesChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                              `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseRulesChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseRulesChallengeRuleExposedCredentialCheckJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesChallengeRuleExposedCredentialCheck]
type phaseVersionGetResponseRulesChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseRulesChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                 `json:"score_response_header_name"`
	JSON                    phaseVersionGetResponseRulesChallengeRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseRulesChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseRulesChallengeRuleRatelimit]
type phaseVersionGetResponseRulesChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseRulesResponseCompressionRule struct {
	ID         string                                                  `json:"id" api:"required"`
	Action     string                                                  `json:"action" api:"required"`
	Enabled    bool                                                    `json:"enabled" api:"required"`
	Expression string                                                  `json:"expression" api:"required"`
	Ref        string                                                  `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesResponseCompressionRuleJSON `json:"-"`
	CompressResponseRule
}

// phaseVersionGetResponseRulesResponseCompressionRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseRulesResponseCompressionRule]
type phaseVersionGetResponseRulesResponseCompressionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesResponseCompressionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesResponseCompressionRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesResponseCompressionRule) implementsPhaseVersionGetResponseRule() {
}

type PhaseVersionGetResponseRulesDDoSDynamicRule struct {
	ID         string                                          `json:"id" api:"required"`
	Action     string                                          `json:"action" api:"required"`
	Enabled    bool                                            `json:"enabled" api:"required"`
	Expression string                                          `json:"expression" api:"required"`
	Ref        string                                          `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesDDoSDynamicRuleJSON `json:"-"`
	DDoSDynamicRule
}

// phaseVersionGetResponseRulesDDoSDynamicRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseRulesDDoSDynamicRule]
type phaseVersionGetResponseRulesDDoSDynamicRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesDDoSDynamicRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesDDoSDynamicRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesDDoSDynamicRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesExecuteRule struct {
	ID         string                                      `json:"id" api:"required"`
	Action     string                                      `json:"action" api:"required"`
	Enabled    bool                                        `json:"enabled" api:"required"`
	Expression string                                      `json:"expression" api:"required"`
	Ref        string                                      `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesExecuteRuleJSON `json:"-"`
	ExecuteRule
}

// phaseVersionGetResponseRulesExecuteRuleJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseRulesExecuteRule]
type phaseVersionGetResponseRulesExecuteRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesExecuteRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesForceConnectionCloseRule struct {
	ID         string                                                   `json:"id" api:"required"`
	Action     string                                                   `json:"action" api:"required"`
	Enabled    bool                                                     `json:"enabled" api:"required"`
	Expression string                                                   `json:"expression" api:"required"`
	Ref        string                                                   `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesForceConnectionCloseRuleJSON `json:"-"`
	ForceConnectionCloseRule
}

// phaseVersionGetResponseRulesForceConnectionCloseRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseRulesForceConnectionCloseRule]
type phaseVersionGetResponseRulesForceConnectionCloseRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesForceConnectionCloseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesForceConnectionCloseRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesForceConnectionCloseRule) implementsPhaseVersionGetResponseRule() {
}

type PhaseVersionGetResponseRulesJavaScriptChallengeRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesJavaScriptChallengeRuleAction `json:"action" api:"required"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled" api:"required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression" api:"required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref" api:"required"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseVersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseRulesJavaScriptChallengeRuleRatelimit `json:"ratelimit"`
	JSON      phaseVersionGetResponseRulesJavaScriptChallengeRuleJSON      `json:"-"`
}

// phaseVersionGetResponseRulesJavaScriptChallengeRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseRulesJavaScriptChallengeRule]
type phaseVersionGetResponseRulesJavaScriptChallengeRuleJSON struct {
	ID                     apijson.Field
	Action                 apijson.Field
	Enabled                apijson.Field
	Expression             apijson.Field
	LastUpdated            apijson.Field
	Ref                    apijson.Field
	Version                apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	ExposedCredentialCheck apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesJavaScriptChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesJavaScriptChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesJavaScriptChallengeRule) implementsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesJavaScriptChallengeRuleAction string

const (
	PhaseVersionGetResponseRulesJavaScriptChallengeRuleActionJSChallenge PhaseVersionGetResponseRulesJavaScriptChallengeRuleAction = "js_challenge"
)

func (r PhaseVersionGetResponseRulesJavaScriptChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesJavaScriptChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                        `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheck]
type phaseVersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseRulesJavaScriptChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                           `json:"score_response_header_name"`
	JSON                    phaseVersionGetResponseRulesJavaScriptChallengeRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseRulesJavaScriptChallengeRuleRatelimitJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesJavaScriptChallengeRuleRatelimit]
type phaseVersionGetResponseRulesJavaScriptChallengeRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesJavaScriptChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesJavaScriptChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseRulesLogRule struct {
	ID         string                                  `json:"id" api:"required"`
	Action     string                                  `json:"action" api:"required"`
	Enabled    bool                                    `json:"enabled" api:"required"`
	Expression string                                  `json:"expression" api:"required"`
	Ref        string                                  `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesLogRuleJSON `json:"-"`
	LogRule
}

// phaseVersionGetResponseRulesLogRuleJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseRulesLogRule]
type phaseVersionGetResponseRulesLogRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesLogRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesLogCustomFieldRule struct {
	ID         string                                             `json:"id" api:"required"`
	Action     string                                             `json:"action" api:"required"`
	Enabled    bool                                               `json:"enabled" api:"required"`
	Expression string                                             `json:"expression" api:"required"`
	Ref        string                                             `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesLogCustomFieldRuleJSON `json:"-"`
	LogCustomFieldRule
}

// phaseVersionGetResponseRulesLogCustomFieldRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseRulesLogCustomFieldRule]
type phaseVersionGetResponseRulesLogCustomFieldRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesLogCustomFieldRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesManagedChallengeRule struct {
	ID         string                                               `json:"id" api:"required"`
	Action     string                                               `json:"action" api:"required"`
	Enabled    bool                                                 `json:"enabled" api:"required"`
	Expression string                                               `json:"expression" api:"required"`
	Ref        string                                               `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesManagedChallengeRuleJSON `json:"-"`
	ManagedChallengeRule
}

// phaseVersionGetResponseRulesManagedChallengeRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseRulesManagedChallengeRule]
type phaseVersionGetResponseRulesManagedChallengeRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesManagedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesManagedChallengeRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesRedirectRule struct {
	ID         string                                       `json:"id" api:"required"`
	Action     string                                       `json:"action" api:"required"`
	Enabled    bool                                         `json:"enabled" api:"required"`
	Expression string                                       `json:"expression" api:"required"`
	Ref        string                                       `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesRedirectRuleJSON `json:"-"`
	RedirectRule
}

// phaseVersionGetResponseRulesRedirectRuleJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseRulesRedirectRule]
type phaseVersionGetResponseRulesRedirectRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRedirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRedirectRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesRewriteRule struct {
	ID         string                                      `json:"id" api:"required"`
	Action     string                                      `json:"action" api:"required"`
	Enabled    bool                                        `json:"enabled" api:"required"`
	Expression string                                      `json:"expression" api:"required"`
	Ref        string                                      `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesRewriteRuleJSON `json:"-"`
	RewriteRule
}

// phaseVersionGetResponseRulesRewriteRuleJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseRulesRewriteRule]
type phaseVersionGetResponseRulesRewriteRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRewriteRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesRouteRule struct {
	ID         string                                    `json:"id" api:"required"`
	Action     string                                    `json:"action" api:"required"`
	Enabled    bool                                      `json:"enabled" api:"required"`
	Expression string                                    `json:"expression" api:"required"`
	Ref        string                                    `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesRouteRuleJSON `json:"-"`
	RouteRule
}

// phaseVersionGetResponseRulesRouteRuleJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseRulesRouteRule]
type phaseVersionGetResponseRulesRouteRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRouteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRouteRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesScoreRule struct {
	ID         string                                    `json:"id" api:"required"`
	Action     string                                    `json:"action" api:"required"`
	Enabled    bool                                      `json:"enabled" api:"required"`
	Expression string                                    `json:"expression" api:"required"`
	Ref        string                                    `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesScoreRuleJSON `json:"-"`
	ScoreRule
}

// phaseVersionGetResponseRulesScoreRuleJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseRulesScoreRule]
type phaseVersionGetResponseRulesScoreRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesScoreRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesServeErrorRule struct {
	ID         string                                         `json:"id" api:"required"`
	Action     string                                         `json:"action" api:"required"`
	Enabled    bool                                           `json:"enabled" api:"required"`
	Expression string                                         `json:"expression" api:"required"`
	Ref        string                                         `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesServeErrorRuleJSON `json:"-"`
	ServeErrorRule
}

// phaseVersionGetResponseRulesServeErrorRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseRulesServeErrorRule]
type phaseVersionGetResponseRulesServeErrorRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesServeErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesServeErrorRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesSetCacheControlRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action  PhaseVersionGetResponseRulesSetCacheControlRuleAction `json:"action" api:"required"`
	Enabled bool                                                  `json:"enabled" api:"required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression" api:"required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref" api:"required"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseVersionGetResponseRulesSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseRulesSetCacheControlRuleRatelimit `json:"ratelimit"`
	JSON      phaseVersionGetResponseRulesSetCacheControlRuleJSON      `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseRulesSetCacheControlRule]
type phaseVersionGetResponseRulesSetCacheControlRuleJSON struct {
	ID                     apijson.Field
	Action                 apijson.Field
	Enabled                apijson.Field
	Expression             apijson.Field
	LastUpdated            apijson.Field
	Ref                    apijson.Field
	Version                apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	ExposedCredentialCheck apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRule) implementsPhaseVersionGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesSetCacheControlRuleAction string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionSetCacheControl PhaseVersionGetResponseRulesSetCacheControlRuleAction = "set_cache_control"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 phaseVersionGetResponseRulesSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParameters]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersJSON struct {
	Immutable            apijson.Field
	MaxAge               apijson.Field
	MustRevalidate       apijson.Field
	MustUnderstand       apijson.Field
	NoCache              apijson.Field
	NoStore              apijson.Field
	NoTransform          apijson.Field
	Private              apijson.Field
	ProxyRevalidate      apijson.Field
	Public               apijson.Field
	SMaxage              apijson.Field
	StaleIfError         apijson.Field
	StaleWhileRevalidate apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                         `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutable]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutable) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                     `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                     `json:"value"`
	JSON  phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                  `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                     `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                              `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                              `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                `json:"qualifiers"`
	JSON       phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCache]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCache) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                               `json:"qualifiers"`
	JSON       phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                       `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStore]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStore) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                   `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                           `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                `json:"qualifiers"`
	JSON       phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivate]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivate) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                               `json:"qualifiers"`
	JSON       phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                               `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                      `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublic]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublic) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                  `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                     `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                      `json:"value"`
	JSON  phaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                   `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                           `json:"value"`
	JSON  phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                   `json:"value"`
	JSON  phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                   `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsPhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, PhaseVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseRulesSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                    `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseRulesSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleExposedCredentialCheck]
type phaseVersionGetResponseRulesSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseRulesSetCacheControlRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                       `json:"score_response_header_name"`
	JSON                    phaseVersionGetResponseRulesSetCacheControlRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheControlRuleRatelimitJSON contains the JSON
// metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheControlRuleRatelimit]
type phaseVersionGetResponseRulesSetCacheControlRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseRulesSetCacheSettingsRule struct {
	ID         string                                               `json:"id" api:"required"`
	Action     string                                               `json:"action" api:"required"`
	Enabled    bool                                                 `json:"enabled" api:"required"`
	Expression string                                               `json:"expression" api:"required"`
	Ref        string                                               `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesSetCacheSettingsRuleJSON `json:"-"`
	SetCacheSettingsRule
}

// phaseVersionGetResponseRulesSetCacheSettingsRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseRulesSetCacheSettingsRule]
type phaseVersionGetResponseRulesSetCacheSettingsRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheSettingsRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesSetCacheTagsRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action  PhaseVersionGetResponseRulesSetCacheTagsRuleAction `json:"action" api:"required"`
	Enabled bool                                               `json:"enabled" api:"required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression" api:"required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref" api:"required"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseVersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseRulesSetCacheTagsRuleRatelimit `json:"ratelimit"`
	JSON      phaseVersionGetResponseRulesSetCacheTagsRuleJSON      `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheTagsRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseRulesSetCacheTagsRule]
type phaseVersionGetResponseRulesSetCacheTagsRuleJSON struct {
	ID                     apijson.Field
	Action                 apijson.Field
	Enabled                apijson.Field
	Expression             apijson.Field
	LastUpdated            apijson.Field
	Ref                    apijson.Field
	Version                apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	ExposedCredentialCheck apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheTagsRule) implementsPhaseVersionGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesSetCacheTagsRuleAction string

const (
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionSetCacheTags PhaseVersionGetResponseRulesSetCacheTagsRuleAction = "set_cache_tags"
)

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                      `json:"values"`
	JSON   phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersUnion
}

// phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters]
type phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters) AsUnion() PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersUnion {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues]
// or
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersUnion interface {
	implementsPhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                           `json:"values" api:"required"`
	JSON   phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues]
type phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsPhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsPhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                              `json:"values" api:"required"`
	JSON   phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsPhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsPhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                           `json:"values" api:"required"`
	JSON   phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues]
type phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsPhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsPhaseVersionGetResponseRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersOperation string

const (
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersOperationAdd    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersOperation = "add"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersOperationRemove PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersOperation = "remove"
	PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersOperationSet    PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersOperation = "set"
)

func (r PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersOperationAdd, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersOperationRemove, PhaseVersionGetResponseRulesSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                 `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [PhaseVersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheck]
type phaseVersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseRulesSetCacheTagsRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                    `json:"score_response_header_name"`
	JSON                    phaseVersionGetResponseRulesSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseRulesSetCacheTagsRuleRatelimitJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseRulesSetCacheTagsRuleRatelimit]
type phaseVersionGetResponseRulesSetCacheTagsRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseRulesSetConfigurationRule struct {
	ID         string                                               `json:"id" api:"required"`
	Action     string                                               `json:"action" api:"required"`
	Enabled    bool                                                 `json:"enabled" api:"required"`
	Expression string                                               `json:"expression" api:"required"`
	Ref        string                                               `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesSetConfigurationRuleJSON `json:"-"`
	SetConfigRule
}

// phaseVersionGetResponseRulesSetConfigurationRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseRulesSetConfigurationRule]
type phaseVersionGetResponseRulesSetConfigurationRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSetConfigurationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSetConfigurationRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSetConfigurationRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesSkipRule struct {
	ID         string                                   `json:"id" api:"required"`
	Action     string                                   `json:"action" api:"required"`
	Enabled    bool                                     `json:"enabled" api:"required"`
	Expression string                                   `json:"expression" api:"required"`
	Ref        string                                   `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseRulesSkipRuleJSON `json:"-"`
	SkipRule
}

// phaseVersionGetResponseRulesSkipRuleJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseRulesSkipRule]
type phaseVersionGetResponseRulesSkipRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesSkipRule) implementsPhaseVersionGetResponseRule() {}

type PhaseVersionGetResponseRulesTransformResponseHTMLRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action  PhaseVersionGetResponseRulesTransformResponseHTMLRuleAction `json:"action" api:"required"`
	Enabled bool                                                        `json:"enabled" api:"required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression" api:"required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref" api:"required"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesTransformResponseHTMLRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseVersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseRulesTransformResponseHTMLRuleRatelimit `json:"ratelimit"`
	JSON      phaseVersionGetResponseRulesTransformResponseHTMLRuleJSON      `json:"-"`
}

// phaseVersionGetResponseRulesTransformResponseHTMLRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseRulesTransformResponseHTMLRule]
type phaseVersionGetResponseRulesTransformResponseHTMLRuleJSON struct {
	ID                     apijson.Field
	Action                 apijson.Field
	Enabled                apijson.Field
	Expression             apijson.Field
	LastUpdated            apijson.Field
	Ref                    apijson.Field
	Version                apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	ExposedCredentialCheck apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesTransformResponseHTMLRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesTransformResponseHTMLRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesTransformResponseHTMLRule) implementsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesTransformResponseHTMLRuleAction string

const (
	PhaseVersionGetResponseRulesTransformResponseHTMLRuleActionTransformResponseHTML PhaseVersionGetResponseRulesTransformResponseHTMLRuleAction = "transform_response_html"
)

func (r PhaseVersionGetResponseRulesTransformResponseHTMLRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesTransformResponseHTMLRuleActionTransformResponseHTML:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesTransformResponseHTMLRuleActionParameters struct {
	// Enables the link maze transformation on the response.
	LinkMaze interface{}                                                               `json:"link_maze" api:"required"`
	JSON     phaseVersionGetResponseRulesTransformResponseHTMLRuleActionParametersJSON `json:"-"`
}

// phaseVersionGetResponseRulesTransformResponseHTMLRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesTransformResponseHTMLRuleActionParameters]
type phaseVersionGetResponseRulesTransformResponseHTMLRuleActionParametersJSON struct {
	LinkMaze    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesTransformResponseHTMLRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesTransformResponseHTMLRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                          `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheck]
type phaseVersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseRulesTransformResponseHTMLRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics []string `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period int64 `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression string `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout int64 `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod int64 `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin bool `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod int64 `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName string                                                             `json:"score_response_header_name"`
	JSON                    phaseVersionGetResponseRulesTransformResponseHTMLRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseRulesTransformResponseHTMLRuleRatelimitJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesTransformResponseHTMLRuleRatelimit]
type phaseVersionGetResponseRulesTransformResponseHTMLRuleRatelimitJSON struct {
	Characteristics         apijson.Field
	Period                  apijson.Field
	CountingExpression      apijson.Field
	MitigationTimeout       apijson.Field
	RequestsPerPeriod       apijson.Field
	RequestsToOrigin        apijson.Field
	ScorePerPeriod          apijson.Field
	ScoreResponseHeaderName apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesTransformResponseHTMLRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesTransformResponseHTMLRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type PhaseVersionGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type PhaseVersionGetResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseVersionGetResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []PhaseVersionGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// A ruleset object.
	Result PhaseVersionGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success PhaseVersionGetResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Message string `json:"message" api:"required"`
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
	Pointer string                                          `json:"pointer" api:"required"`
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
	Message string `json:"message" api:"required"`
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
	Pointer string                                            `json:"pointer" api:"required"`
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
