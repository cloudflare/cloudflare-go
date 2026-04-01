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

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
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
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [LogCustomFieldRuleActionParameters],
	// [RedirectRuleActionParameters], [RewriteRuleActionParameters],
	// [RouteRuleActionParameters], [ScoreRuleActionParameters],
	// [ServeErrorRuleActionParameters],
	// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [PhaseVersionGetResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [PhaseVersionGetResponseRulesRulesetsJSChallengeRuleRatelimit],
	// [LogRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit],
	// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleRatelimit],
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref   string                          `json:"ref"`
	JSON  phaseVersionGetResponseRuleJSON `json:"-"`
	union PhaseVersionGetResponseRulesUnion
}

// phaseVersionGetResponseRuleJSON contains the JSON metadata for the struct
// [PhaseVersionGetResponseRule]
type phaseVersionGetResponseRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
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
// Possible runtime types of the union are [BlockRule],
// [PhaseVersionGetResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [DDoSDynamicRule], [ExecuteRule], [ForceConnectionCloseRule],
// [PhaseVersionGetResponseRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRule],
// [SetCacheSettingsRule], [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRule],
// [SetConfigRule], [SkipRule].
func (r PhaseVersionGetResponseRule) AsUnion() PhaseVersionGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule],
// [PhaseVersionGetResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [DDoSDynamicRule], [ExecuteRule], [ForceConnectionCloseRule],
// [PhaseVersionGetResponseRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRule],
// [SetCacheSettingsRule], [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRule],
// [SetConfigRule] or [SkipRule].
type PhaseVersionGetResponseRulesUnion interface {
	implementsPhaseVersionGetResponseRule()
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
			Type:               reflect.TypeOf(CompressResponseRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsJSChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
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
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheTagsRule{}),
			DiscriminatorValue: "set_cache_tags",
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
	)
}

type PhaseVersionGetResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
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
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseVersionGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseRulesRulesetsChallengeRule]
type phaseVersionGetResponseRulesRulesetsChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsChallengeRule) implementsPhaseVersionGetResponseRule() {}

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

// Configuration for exposed credential checking.
type PhaseVersionGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                      `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type phaseVersionGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseRulesRulesetsChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                         `json:"score_response_header_name"`
	JSON                    phaseVersionGetResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsChallengeRuleRatelimit]
type phaseVersionGetResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseVersionGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                  `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsJSChallengeRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseRulesRulesetsJSChallengeRule]
type phaseVersionGetResponseRulesRulesetsJSChallengeRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsJSChallengeRule) implementsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsJSChallengeRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsJSChallengeRuleActionJSChallenge PhaseVersionGetResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r PhaseVersionGetResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                        `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type phaseVersionGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseRulesRulesetsJSChallengeRuleRatelimit struct {
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
	JSON                    phaseVersionGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsJSChallengeRuleRatelimit]
type phaseVersionGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseRulesRulesetsSetCacheControlRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                      `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsSetCacheControlRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleJSON contains the JSON
// metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRule]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRule) implementsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionSetCacheControl PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleAction = "set_cache_control"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                 `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                             `json:"value"`
	JSON  phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                  `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                     `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                  `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                     `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                        `json:"qualifiers"`
	JSON       phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                       `json:"qualifiers"`
	JSON       phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                               `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                   `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                               `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                  `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                        `json:"qualifiers"`
	JSON       phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                       `json:"qualifiers"`
	JSON       phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                   `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                      `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                              `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                              `json:"value"`
	JSON  phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                   `json:"value"`
	JSON  phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                   `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                           `json:"value"`
	JSON  phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                        `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                           `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsPhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                            `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                               `json:"score_response_header_name"`
	JSON                    phaseVersionGetResponseRulesRulesetsSetCacheControlRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheControlRuleRatelimitJSON contains
// the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleRatelimit]
type phaseVersionGetResponseRulesRulesetsSetCacheControlRuleRatelimitJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                   `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRule]
type phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleJSON struct {
	LastUpdated            apijson.Field
	Version                apijson.Field
	ID                     apijson.Field
	Action                 apijson.Field
	ActionParameters       apijson.Field
	Categories             apijson.Field
	Description            apijson.Field
	Enabled                apijson.Field
	ExposedCredentialCheck apijson.Field
	Expression             apijson.Field
	Logging                apijson.Field
	Ratelimit              apijson.Field
	Ref                    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRule) implementsPhaseVersionGetResponseRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleAction string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionSetCacheTags PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleAction = "set_cache_tags"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                              `json:"values"`
	JSON   phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion
}

// phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters]
type phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters) AsUnion() PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
// or
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion interface {
	implementsPhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                   `json:"values" api:"required"`
	JSON   phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues]
type phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsPhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsPhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                      `json:"values" api:"required"`
	JSON   phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsPhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsPhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                   `json:"values" api:"required"`
	JSON   phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
type phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsPhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsPhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation string

const (
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "add"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "remove"
	PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationSet    PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "set"
)

func (r PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove, PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                         `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck]
type phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                            `json:"score_response_header_name"`
	JSON                    phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleRatelimit]
type phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesAction string

const (
	PhaseVersionGetResponseRulesActionBlock                PhaseVersionGetResponseRulesAction = "block"
	PhaseVersionGetResponseRulesActionChallenge            PhaseVersionGetResponseRulesAction = "challenge"
	PhaseVersionGetResponseRulesActionCompressResponse     PhaseVersionGetResponseRulesAction = "compress_response"
	PhaseVersionGetResponseRulesActionDDoSDynamic          PhaseVersionGetResponseRulesAction = "ddos_dynamic"
	PhaseVersionGetResponseRulesActionExecute              PhaseVersionGetResponseRulesAction = "execute"
	PhaseVersionGetResponseRulesActionForceConnectionClose PhaseVersionGetResponseRulesAction = "force_connection_close"
	PhaseVersionGetResponseRulesActionJSChallenge          PhaseVersionGetResponseRulesAction = "js_challenge"
	PhaseVersionGetResponseRulesActionLog                  PhaseVersionGetResponseRulesAction = "log"
	PhaseVersionGetResponseRulesActionLogCustomField       PhaseVersionGetResponseRulesAction = "log_custom_field"
	PhaseVersionGetResponseRulesActionManagedChallenge     PhaseVersionGetResponseRulesAction = "managed_challenge"
	PhaseVersionGetResponseRulesActionRedirect             PhaseVersionGetResponseRulesAction = "redirect"
	PhaseVersionGetResponseRulesActionRewrite              PhaseVersionGetResponseRulesAction = "rewrite"
	PhaseVersionGetResponseRulesActionRoute                PhaseVersionGetResponseRulesAction = "route"
	PhaseVersionGetResponseRulesActionScore                PhaseVersionGetResponseRulesAction = "score"
	PhaseVersionGetResponseRulesActionServeError           PhaseVersionGetResponseRulesAction = "serve_error"
	PhaseVersionGetResponseRulesActionSetCacheControl      PhaseVersionGetResponseRulesAction = "set_cache_control"
	PhaseVersionGetResponseRulesActionSetCacheSettings     PhaseVersionGetResponseRulesAction = "set_cache_settings"
	PhaseVersionGetResponseRulesActionSetCacheTags         PhaseVersionGetResponseRulesAction = "set_cache_tags"
	PhaseVersionGetResponseRulesActionSetConfig            PhaseVersionGetResponseRulesAction = "set_config"
	PhaseVersionGetResponseRulesActionSkip                 PhaseVersionGetResponseRulesAction = "skip"
)

func (r PhaseVersionGetResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesActionBlock, PhaseVersionGetResponseRulesActionChallenge, PhaseVersionGetResponseRulesActionCompressResponse, PhaseVersionGetResponseRulesActionDDoSDynamic, PhaseVersionGetResponseRulesActionExecute, PhaseVersionGetResponseRulesActionForceConnectionClose, PhaseVersionGetResponseRulesActionJSChallenge, PhaseVersionGetResponseRulesActionLog, PhaseVersionGetResponseRulesActionLogCustomField, PhaseVersionGetResponseRulesActionManagedChallenge, PhaseVersionGetResponseRulesActionRedirect, PhaseVersionGetResponseRulesActionRewrite, PhaseVersionGetResponseRulesActionRoute, PhaseVersionGetResponseRulesActionScore, PhaseVersionGetResponseRulesActionServeError, PhaseVersionGetResponseRulesActionSetCacheControl, PhaseVersionGetResponseRulesActionSetCacheSettings, PhaseVersionGetResponseRulesActionSetCacheTags, PhaseVersionGetResponseRulesActionSetConfig, PhaseVersionGetResponseRulesActionSkip:
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
