// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/tidwall/gjson"
)

// PhaseService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhaseService] method instead.
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
func (r *PhaseService) Update(ctx context.Context, rulesetPhase Phase, params PhaseUpdateParams, opts ...option.RequestOption) (res *PhaseUpdateResponseEnvelopeResult, err error) {
	var env PhaseUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetches the latest version of the account or zone entry point ruleset for a
// given phase.
func (r *PhaseService) Get(ctx context.Context, rulesetPhase Phase, query PhaseGetParams, opts ...option.RequestOption) (res *PhaseGetResponse, err error) {
	var env PhaseGetResponseEnvelope
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
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A ruleset object.
type PhaseGetResponse struct {
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
	Rules []PhaseGetResponseRule `json:"rules" api:"required"`
	// The version of the ruleset.
	Version string `json:"version" api:"required"`
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

type PhaseGetResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [LogCustomFieldRuleActionParameters],
	// [RedirectRuleActionParameters], [RewriteRuleActionParameters],
	// [RouteRuleActionParameters], [ScoreRuleActionParameters],
	// [ServeErrorRuleActionParameters],
	// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [PhaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [PhaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [PhaseGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck],
	// [PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [PhaseGetResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [PhaseGetResponseRulesRulesetsJSChallengeRuleRatelimit], [LogRuleRatelimit],
	// [LogCustomFieldRuleRatelimit], [ManagedChallengeRuleRatelimit],
	// [RedirectRuleRatelimit], [RewriteRuleRatelimit], [RouteRuleRatelimit],
	// [ScoreRuleRatelimit], [ServeErrorRuleRatelimit],
	// [PhaseGetResponseRulesRulesetsSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleRatelimit],
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref   string                   `json:"ref"`
	JSON  phaseGetResponseRuleJSON `json:"-"`
	union PhaseGetResponseRulesUnion
}

// phaseGetResponseRuleJSON contains the JSON metadata for the struct
// [PhaseGetResponseRule]
type phaseGetResponseRuleJSON struct {
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

func (r phaseGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PhaseGetResponseRulesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [PhaseGetResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [DDoSDynamicRule], [ExecuteRule], [ForceConnectionCloseRule],
// [PhaseGetResponseRulesRulesetsJSChallengeRule], [LogRule], [LogCustomFieldRule],
// [ManagedChallengeRule], [RedirectRule], [RewriteRule], [RouteRule], [ScoreRule],
// [ServeErrorRule], [PhaseGetResponseRulesRulesetsSetCacheControlRule],
// [SetCacheSettingsRule], [PhaseGetResponseRulesRulesetsSetCacheTagsRule],
// [SetConfigRule], [SkipRule],
// [PhaseGetResponseRulesRulesetsTransformResponseHTMLRule].
func (r PhaseGetResponseRule) AsUnion() PhaseGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [PhaseGetResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [DDoSDynamicRule], [ExecuteRule],
// [ForceConnectionCloseRule], [PhaseGetResponseRulesRulesetsJSChallengeRule],
// [LogRule], [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule],
// [RewriteRule], [RouteRule], [ScoreRule], [ServeErrorRule],
// [PhaseGetResponseRulesRulesetsSetCacheControlRule], [SetCacheSettingsRule],
// [PhaseGetResponseRulesRulesetsSetCacheTagsRule], [SetConfigRule], [SkipRule] or
// [PhaseGetResponseRulesRulesetsTransformResponseHTMLRule].
type PhaseGetResponseRulesUnion interface {
	implementsPhaseGetResponseRule()
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
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsJSChallengeRule{}),
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
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheTagsRule{}),
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
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsTransformResponseHTMLRule{}),
			DiscriminatorValue: "transform_response_html",
		},
	)
}

type PhaseGetResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
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
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseGetResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                         `json:"ref"`
	JSON phaseGetResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [PhaseGetResponseRulesRulesetsChallengeRule]
type phaseGetResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsChallengeRule) implementsPhaseGetResponseRule() {}

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

// Configuration for exposed credential checking.
type PhaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                               `json:"username_expression" api:"required"`
	JSON               phaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type phaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseGetResponseRulesRulesetsChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                  `json:"score_response_header_name"`
	JSON                    phaseGetResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [PhaseGetResponseRulesRulesetsChallengeRuleRatelimit]
type phaseGetResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseGetResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseGetResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                           `json:"ref"`
	JSON phaseGetResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata for
// the struct [PhaseGetResponseRulesRulesetsJSChallengeRule]
type phaseGetResponseRulesRulesetsJSChallengeRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsJSChallengeRule) implementsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsJSChallengeRuleAction string

const (
	PhaseGetResponseRulesRulesetsJSChallengeRuleActionJSChallenge PhaseGetResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r PhaseGetResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                 `json:"username_expression" api:"required"`
	JSON               phaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type phaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseGetResponseRulesRulesetsJSChallengeRuleRatelimit struct {
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
	JSON                    phaseGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [PhaseGetResponseRulesRulesetsJSChallengeRuleRatelimit]
type phaseGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseGetResponseRulesRulesetsSetCacheControlRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsSetCacheControlRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseGetResponseRulesRulesetsSetCacheControlRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                               `json:"ref"`
	JSON phaseGetResponseRulesRulesetsSetCacheControlRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleJSON contains the JSON metadata
// for the struct [PhaseGetResponseRulesRulesetsSetCacheControlRule]
type phaseGetResponseRulesRulesetsSetCacheControlRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRule) implementsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleAction string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionSetCacheControl PhaseGetResponseRulesRulesetsSetCacheControlRuleAction = "set_cache_control"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParameters]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                          `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                      `json:"value"`
	JSON  phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                   `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                               `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                               `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                 `json:"qualifiers"`
	JSON       phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                `json:"qualifiers"`
	JSON       phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                        `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                    `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                            `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                 `json:"qualifiers"`
	JSON       phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                `json:"qualifiers"`
	JSON       phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                               `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                       `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                   `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                       `json:"value"`
	JSON  phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                    `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                            `json:"value"`
	JSON  phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                    `json:"value"`
	JSON  phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                 `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                    `json:"cloudflare_only"`
	JSON           phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsPhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, PhaseGetResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                     `json:"username_expression" api:"required"`
	JSON               phaseGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck]
type phaseGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseGetResponseRulesRulesetsSetCacheControlRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                        `json:"score_response_header_name"`
	JSON                    phaseGetResponseRulesRulesetsSetCacheControlRuleRatelimitJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheControlRuleRatelimitJSON contains the JSON
// metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheControlRuleRatelimit]
type phaseGetResponseRulesRulesetsSetCacheControlRuleRatelimitJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseGetResponseRulesRulesetsSetCacheTagsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsSetCacheTagsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseGetResponseRulesRulesetsSetCacheTagsRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                            `json:"ref"`
	JSON phaseGetResponseRulesRulesetsSetCacheTagsRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheTagsRuleJSON contains the JSON metadata for
// the struct [PhaseGetResponseRulesRulesetsSetCacheTagsRule]
type phaseGetResponseRulesRulesetsSetCacheTagsRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRule) implementsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleAction string

const (
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionSetCacheTags PhaseGetResponseRulesRulesetsSetCacheTagsRuleAction = "set_cache_tags"
)

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                       `json:"values"`
	JSON   phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion
}

// phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters]
type phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters) AsUnion() PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
// or
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion interface {
	implementsPhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                            `json:"values" api:"required"`
	JSON   phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues]
type phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsPhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsPhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                               `json:"values" api:"required"`
	JSON   phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsPhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsPhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                            `json:"values" api:"required"`
	JSON   phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
type phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsPhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsPhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation string

const (
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "add"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "remove"
	PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationSet    PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "set"
)

func (r PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove, PhaseGetResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                  `json:"username_expression" api:"required"`
	JSON               phaseGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck]
type phaseGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseGetResponseRulesRulesetsSetCacheTagsRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                     `json:"score_response_header_name"`
	JSON                    phaseGetResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON contains the JSON
// metadata for the struct [PhaseGetResponseRulesRulesetsSetCacheTagsRuleRatelimit]
type phaseGetResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseGetResponseRulesRulesetsTransformResponseHTMLRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                     `json:"ref"`
	JSON phaseGetResponseRulesRulesetsTransformResponseHTMLRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsTransformResponseHTMLRuleJSON contains the JSON
// metadata for the struct [PhaseGetResponseRulesRulesetsTransformResponseHTMLRule]
type phaseGetResponseRulesRulesetsTransformResponseHTMLRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsTransformResponseHTMLRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsTransformResponseHTMLRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsTransformResponseHTMLRule) implementsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleAction string

const (
	PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleAction = "transform_response_html"
)

func (r PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleActionParameters struct {
	// Enables the link maze transformation on the response.
	LinkMaze interface{}                                                                `json:"link_maze" api:"required"`
	JSON     phaseGetResponseRulesRulesetsTransformResponseHTMLRuleActionParametersJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsTransformResponseHTMLRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleActionParameters]
type phaseGetResponseRulesRulesetsTransformResponseHTMLRuleActionParametersJSON struct {
	LinkMaze    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsTransformResponseHTMLRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Configuration for exposed credential checking.
type PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                           `json:"username_expression" api:"required"`
	JSON               phaseGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck]
type phaseGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                              `json:"score_response_header_name"`
	JSON                    phaseGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimitJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimitJSON contains the
// JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimit]
type phaseGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimitJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsTransformResponseHTMLRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type PhaseGetResponseRulesAction string

const (
	PhaseGetResponseRulesActionBlock                 PhaseGetResponseRulesAction = "block"
	PhaseGetResponseRulesActionChallenge             PhaseGetResponseRulesAction = "challenge"
	PhaseGetResponseRulesActionCompressResponse      PhaseGetResponseRulesAction = "compress_response"
	PhaseGetResponseRulesActionDDoSDynamic           PhaseGetResponseRulesAction = "ddos_dynamic"
	PhaseGetResponseRulesActionExecute               PhaseGetResponseRulesAction = "execute"
	PhaseGetResponseRulesActionForceConnectionClose  PhaseGetResponseRulesAction = "force_connection_close"
	PhaseGetResponseRulesActionJSChallenge           PhaseGetResponseRulesAction = "js_challenge"
	PhaseGetResponseRulesActionLog                   PhaseGetResponseRulesAction = "log"
	PhaseGetResponseRulesActionLogCustomField        PhaseGetResponseRulesAction = "log_custom_field"
	PhaseGetResponseRulesActionManagedChallenge      PhaseGetResponseRulesAction = "managed_challenge"
	PhaseGetResponseRulesActionRedirect              PhaseGetResponseRulesAction = "redirect"
	PhaseGetResponseRulesActionRewrite               PhaseGetResponseRulesAction = "rewrite"
	PhaseGetResponseRulesActionRoute                 PhaseGetResponseRulesAction = "route"
	PhaseGetResponseRulesActionScore                 PhaseGetResponseRulesAction = "score"
	PhaseGetResponseRulesActionServeError            PhaseGetResponseRulesAction = "serve_error"
	PhaseGetResponseRulesActionSetCacheControl       PhaseGetResponseRulesAction = "set_cache_control"
	PhaseGetResponseRulesActionSetCacheSettings      PhaseGetResponseRulesAction = "set_cache_settings"
	PhaseGetResponseRulesActionSetCacheTags          PhaseGetResponseRulesAction = "set_cache_tags"
	PhaseGetResponseRulesActionSetConfig             PhaseGetResponseRulesAction = "set_config"
	PhaseGetResponseRulesActionSkip                  PhaseGetResponseRulesAction = "skip"
	PhaseGetResponseRulesActionTransformResponseHTML PhaseGetResponseRulesAction = "transform_response_html"
)

func (r PhaseGetResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesActionBlock, PhaseGetResponseRulesActionChallenge, PhaseGetResponseRulesActionCompressResponse, PhaseGetResponseRulesActionDDoSDynamic, PhaseGetResponseRulesActionExecute, PhaseGetResponseRulesActionForceConnectionClose, PhaseGetResponseRulesActionJSChallenge, PhaseGetResponseRulesActionLog, PhaseGetResponseRulesActionLogCustomField, PhaseGetResponseRulesActionManagedChallenge, PhaseGetResponseRulesActionRedirect, PhaseGetResponseRulesActionRewrite, PhaseGetResponseRulesActionRoute, PhaseGetResponseRulesActionScore, PhaseGetResponseRulesActionServeError, PhaseGetResponseRulesActionSetCacheControl, PhaseGetResponseRulesActionSetCacheSettings, PhaseGetResponseRulesActionSetCacheTags, PhaseGetResponseRulesActionSetConfig, PhaseGetResponseRulesActionSkip, PhaseGetResponseRulesActionTransformResponseHTML:
		return true
	}
	return false
}

type PhaseUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// Validates the request without persisting changes when set to `true`. Responses
	// that normally return 200 return `result: null`; endpoints that normally return
	// 204 continue to return 204.
	DryRun param.Field[bool] `query:"dry_run"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The list of rules in the ruleset.
	Rules param.Field[[]PhaseUpdateParamsRuleUnion] `json:"rules"`
}

func (r PhaseUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [PhaseUpdateParams]'s query parameters as `url.Values`.
func (r PhaseUpdateParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type PhaseUpdateParamsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action           param.Field[PhaseUpdateParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                  `json:"action_parameters"`
	Categories       param.Field[interface{}]                  `json:"categories"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled                param.Field[bool]        `json:"enabled"`
	ExposedCredentialCheck param.Field[interface{}] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging   param.Field[LoggingParam] `json:"logging"`
	Ratelimit param.Field[interface{}]  `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRule) implementsPhaseUpdateParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam],
// [rulesets.PhaseUpdateParamsRulesRulesetsChallengeRule],
// [rulesets.CompressResponseRuleParam], [rulesets.DDoSDynamicRuleParam],
// [rulesets.ExecuteRuleParam], [rulesets.ForceConnectionCloseRuleParam],
// [rulesets.PhaseUpdateParamsRulesRulesetsJSChallengeRule],
// [rulesets.LogRuleParam], [rulesets.LogCustomFieldRuleParam],
// [rulesets.ManagedChallengeRuleParam], [rulesets.RedirectRuleParam],
// [rulesets.RewriteRuleParam], [rulesets.RouteRuleParam],
// [rulesets.ScoreRuleParam], [rulesets.ServeErrorRuleParam],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRule],
// [rulesets.SetCacheSettingsRuleParam],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheTagsRule],
// [rulesets.SetConfigRuleParam], [rulesets.SkipRuleParam],
// [rulesets.PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRule],
// [PhaseUpdateParamsRule].
type PhaseUpdateParamsRuleUnion interface {
	implementsPhaseUpdateParamsRuleUnion()
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
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[PhaseUpdateParamsRulesRulesetsChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[PhaseUpdateParamsRulesRulesetsChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsChallengeRule) implementsPhaseUpdateParamsRuleUnion() {}

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

// Configuration for exposed credential checking.
type PhaseUpdateParamsRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r PhaseUpdateParamsRulesRulesetsChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateParamsRulesRulesetsChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r PhaseUpdateParamsRulesRulesetsChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhaseUpdateParamsRulesRulesetsJSChallengeRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsJSChallengeRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[PhaseUpdateParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[PhaseUpdateParamsRulesRulesetsJSChallengeRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsJSChallengeRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsJSChallengeRule) implementsPhaseUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsJSChallengeRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsJSChallengeRuleActionJSChallenge PhaseUpdateParamsRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r PhaseUpdateParamsRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseUpdateParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r PhaseUpdateParamsRulesRulesetsJSChallengeRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateParamsRulesRulesetsJSChallengeRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r PhaseUpdateParamsRulesRulesetsJSChallengeRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhaseUpdateParamsRulesRulesetsSetCacheControlRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRule) implementsPhaseUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionSetCacheControl PhaseUpdateParamsRulesRulesetsSetCacheControlRuleAction = "set_cache_control"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion] `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion] `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion] `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion] `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion] `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion] `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion] `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion] `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion] `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion] `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion] `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion] `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion] `json:"stale-while-revalidate"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A cache-control directive configuration.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutable) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutable) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutable].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion()
}

// Set the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAge) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAge) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion() {
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAge].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion()
}

// Set the directive with a duration value in seconds.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation] `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion()
}

// Set the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion()
}

// Set the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool]        `json:"cloudflare_only"`
	Qualifiers     param.Field[interface{}] `json:"qualifiers"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCache) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCache) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion() {
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCache].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion()
}

// Set the directive with optional qualifiers.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers param.Field[[]string] `json:"qualifiers"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStore) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStore].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion()
}

// Set the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransform) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransform) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransform].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion()
}

// Set the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool]        `json:"cloudflare_only"`
	Qualifiers     param.Field[interface{}] `json:"qualifiers"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivate) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion() {
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivate].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion()
}

// Set the directive with optional qualifiers.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers param.Field[[]string] `json:"qualifiers"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion()
}

// Set the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublic) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublic) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion() {
}

// A cache-control directive configuration.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublic].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion()
}

// Set the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxage) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxage) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion() {
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxage].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion()
}

// Set the directive with a duration value in seconds.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation] `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion() {
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfError].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion()
}

// Set the directive with a duration value in seconds.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation] `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion() {
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective],
// [PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate].
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion()
}

// Set the directive with a duration value in seconds.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation] `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value param.Field[int64] `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation] `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly param.Field[bool] `json:"cloudflare_only"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsPhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, PhaseUpdateParamsRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateParamsRulesRulesetsSetCacheControlRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheControlRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhaseUpdateParamsRulesRulesetsSetCacheTagsRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRule) implementsPhaseUpdateParamsRuleUnion() {}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionSetCacheTags PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleAction = "set_cache_tags"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation] `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression param.Field[string]      `json:"expression"`
	Values     param.Field[interface{}] `json:"values"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParameters) implementsPhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The parameters configuring the rule's action.
//
// Satisfied by
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [rulesets.PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression],
// [PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParameters].
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion interface {
	implementsPhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion()
}

// Add cache tags using a list of values.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation] `json:"operation" api:"required"`
	// A list of cache tag values.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsPhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression param.Field[string] `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation] `json:"operation" api:"required"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsPhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation] `json:"operation" api:"required"`
	// A list of cache tag values.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsPhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression param.Field[string] `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation] `json:"operation" api:"required"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsPhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation] `json:"operation" api:"required"`
	// A list of cache tag values.
	Values param.Field[[]string] `json:"values" api:"required"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsPhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression param.Field[string] `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation param.Field[PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation] `json:"operation" api:"required"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsPhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersUnion() {
}

// The operation to perform on the cache tags.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation string

const (
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation = "add"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation = "remove"
	PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationSet    PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation = "set"
)

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove, PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r PhaseUpdateParamsRulesRulesetsSetCacheTagsRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck param.Field[PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck] `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit param.Field[PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleRatelimit] `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRule) implementsPhaseUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleAction = "transform_response_html"
)

func (r PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleActionParameters struct {
	// Enables the link maze transformation on the response.
	LinkMaze param.Field[interface{}] `json:"link_maze" api:"required"`
}

func (r PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for exposed credential checking.
type PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression param.Field[string] `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression param.Field[string] `json:"username_expression" api:"required"`
}

func (r PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleRatelimit struct {
	// Characteristics of the request on which the rate limit counter will be
	// incremented.
	Characteristics param.Field[[]string] `json:"characteristics" api:"required"`
	// Period in seconds over which the counter is being incremented.
	Period param.Field[int64] `json:"period" api:"required"`
	// An expression that defines when the rate limit counter should be incremented. It
	// defaults to the same as the rule's expression.
	CountingExpression param.Field[string] `json:"counting_expression"`
	// Period of time in seconds after which the action will be disabled following its
	// first execution.
	MitigationTimeout param.Field[int64] `json:"mitigation_timeout"`
	// The threshold of requests per period after which the action will be executed for
	// the first time.
	RequestsPerPeriod param.Field[int64] `json:"requests_per_period"`
	// Whether counting is only performed when an origin is reached.
	RequestsToOrigin param.Field[bool] `json:"requests_to_origin"`
	// The score threshold per period for which the action will be executed the first
	// time.
	ScorePerPeriod param.Field[int64] `json:"score_per_period"`
	// A response header name provided by the origin, which contains the score to
	// increment rate limit counter with.
	ScoreResponseHeaderName param.Field[string] `json:"score_response_header_name"`
}

func (r PhaseUpdateParamsRulesRulesetsTransformResponseHTMLRuleRatelimit) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesAction string

const (
	PhaseUpdateParamsRulesActionBlock                 PhaseUpdateParamsRulesAction = "block"
	PhaseUpdateParamsRulesActionChallenge             PhaseUpdateParamsRulesAction = "challenge"
	PhaseUpdateParamsRulesActionCompressResponse      PhaseUpdateParamsRulesAction = "compress_response"
	PhaseUpdateParamsRulesActionDDoSDynamic           PhaseUpdateParamsRulesAction = "ddos_dynamic"
	PhaseUpdateParamsRulesActionExecute               PhaseUpdateParamsRulesAction = "execute"
	PhaseUpdateParamsRulesActionForceConnectionClose  PhaseUpdateParamsRulesAction = "force_connection_close"
	PhaseUpdateParamsRulesActionJSChallenge           PhaseUpdateParamsRulesAction = "js_challenge"
	PhaseUpdateParamsRulesActionLog                   PhaseUpdateParamsRulesAction = "log"
	PhaseUpdateParamsRulesActionLogCustomField        PhaseUpdateParamsRulesAction = "log_custom_field"
	PhaseUpdateParamsRulesActionManagedChallenge      PhaseUpdateParamsRulesAction = "managed_challenge"
	PhaseUpdateParamsRulesActionRedirect              PhaseUpdateParamsRulesAction = "redirect"
	PhaseUpdateParamsRulesActionRewrite               PhaseUpdateParamsRulesAction = "rewrite"
	PhaseUpdateParamsRulesActionRoute                 PhaseUpdateParamsRulesAction = "route"
	PhaseUpdateParamsRulesActionScore                 PhaseUpdateParamsRulesAction = "score"
	PhaseUpdateParamsRulesActionServeError            PhaseUpdateParamsRulesAction = "serve_error"
	PhaseUpdateParamsRulesActionSetCacheControl       PhaseUpdateParamsRulesAction = "set_cache_control"
	PhaseUpdateParamsRulesActionSetCacheSettings      PhaseUpdateParamsRulesAction = "set_cache_settings"
	PhaseUpdateParamsRulesActionSetCacheTags          PhaseUpdateParamsRulesAction = "set_cache_tags"
	PhaseUpdateParamsRulesActionSetConfig             PhaseUpdateParamsRulesAction = "set_config"
	PhaseUpdateParamsRulesActionSkip                  PhaseUpdateParamsRulesAction = "skip"
	PhaseUpdateParamsRulesActionTransformResponseHTML PhaseUpdateParamsRulesAction = "transform_response_html"
)

func (r PhaseUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesActionBlock, PhaseUpdateParamsRulesActionChallenge, PhaseUpdateParamsRulesActionCompressResponse, PhaseUpdateParamsRulesActionDDoSDynamic, PhaseUpdateParamsRulesActionExecute, PhaseUpdateParamsRulesActionForceConnectionClose, PhaseUpdateParamsRulesActionJSChallenge, PhaseUpdateParamsRulesActionLog, PhaseUpdateParamsRulesActionLogCustomField, PhaseUpdateParamsRulesActionManagedChallenge, PhaseUpdateParamsRulesActionRedirect, PhaseUpdateParamsRulesActionRewrite, PhaseUpdateParamsRulesActionRoute, PhaseUpdateParamsRulesActionScore, PhaseUpdateParamsRulesActionServeError, PhaseUpdateParamsRulesActionSetCacheControl, PhaseUpdateParamsRulesActionSetCacheSettings, PhaseUpdateParamsRulesActionSetCacheTags, PhaseUpdateParamsRulesActionSetConfig, PhaseUpdateParamsRulesActionSkip, PhaseUpdateParamsRulesActionTransformResponseHTML:
		return true
	}
	return false
}

// A response object.
type PhaseUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseUpdateResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []PhaseUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// A result.
	Result PhaseUpdateResponseEnvelopeResult `json:"result" api:"required"`
	// Whether the API call was successful.
	Success PhaseUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Message string `json:"message" api:"required"`
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
	Pointer string                                      `json:"pointer" api:"required"`
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
	Message string `json:"message" api:"required"`
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
	Pointer string                                        `json:"pointer" api:"required"`
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

// A ruleset object.
type PhaseUpdateResponseEnvelopeResult struct {
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
	Rules []PhaseUpdateResponseEnvelopeResultRules `json:"rules" api:"required"`
	// The version of the ruleset.
	Version string `json:"version" api:"required"`
	// An informative description of the ruleset.
	Description string                                `json:"description"`
	JSON        phaseUpdateResponseEnvelopeResultJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultJSON contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResult]
type phaseUpdateResponseEnvelopeResultJSON struct {
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

func (r *PhaseUpdateResponseEnvelopeResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseEnvelopeResultRules struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseEnvelopeResultRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [LogCustomFieldRuleActionParameters],
	// [RedirectRuleActionParameters], [RewriteRuleActionParameters],
	// [RouteRuleActionParameters], [ScoreRuleActionParameters],
	// [ServeErrorRuleActionParameters],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit],
	// [LogRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit],
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref   string                                     `json:"ref"`
	JSON  phaseUpdateResponseEnvelopeResultRulesJSON `json:"-"`
	union PhaseUpdateResponseEnvelopeResultRules
}

// phaseUpdateResponseEnvelopeResultRulesJSON contains the JSON metadata for the
// struct [PhaseUpdateResponseEnvelopeResultRules]
type phaseUpdateResponseEnvelopeResultRulesJSON struct {
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

func (r phaseUpdateResponseEnvelopeResultRulesJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRules) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRules{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PhaseUpdateResponseEnvelopeResultRules] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRule],
// [CompressResponseRule], [DDoSDynamicRule], [ExecuteRule],
// [ForceConnectionCloseRule],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule],
// [SetCacheSettingsRule],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule],
// [SetConfigRule], [SkipRule],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule].
func (r PhaseUpdateResponseEnvelopeResultRules) AsUnion() PhaseUpdateResponseEnvelopeResultRules {
	return r.union
}

// Union satisfied by [BlockRule],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRule],
// [CompressResponseRule], [DDoSDynamicRule], [ExecuteRule],
// [ForceConnectionCloseRule],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule],
// [SetCacheSettingsRule],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule],
// [SetConfigRule], [SkipRule] or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule].
type PhaseUpdateResponseEnvelopeResultRules interface {
	implementsPhaseUpdateResponseEnvelopeResultRules()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRules)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(BlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRule{}),
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
			Type:               reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule{}),
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
			Type:               reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule{}),
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
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule{}),
			DiscriminatorValue: "transform_response_html",
		},
	)
}

type PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                          `json:"ref"`
	JSON phaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleJSON contains the
// JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRule]
type phaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleJSON struct {
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

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRule) implementsPhaseUpdateResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleAction string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleActionChallenge PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleAction = "challenge"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                `json:"username_expression" api:"required"`
	JSON               phaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck]
type phaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                   `json:"score_response_header_name"`
	JSON                    phaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimitJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit]
type phaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimitJSON struct {
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

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                            `json:"ref"`
	JSON phaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleJSON contains the
// JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule]
type phaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleJSON struct {
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

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRule) implementsPhaseUpdateResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleAction string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleActionJSChallenge PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                  `json:"username_expression" api:"required"`
	JSON               phaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type phaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                     `json:"score_response_header_name"`
	JSON                    phaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimitJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit]
type phaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimitJSON struct {
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

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                                `json:"ref"`
	JSON phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleJSON contains
// the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleJSON struct {
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

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRule) implementsPhaseUpdateResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleAction string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionSetCacheControl PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleAction = "set_cache_control"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersJSON struct {
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

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                       `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                          `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                       `json:"value"`
	JSON  phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                    `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                       `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                            `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                               `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                            `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                               `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                                  `json:"qualifiers"`
	JSON       phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                                 `json:"qualifiers"`
	JSON       phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                        `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                     `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                        `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                         `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                            `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                                  `json:"qualifiers"`
	JSON       phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                                 `json:"qualifiers"`
	JSON       phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                        `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                 `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                             `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                                `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                    `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                       `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                        `json:"value"`
	JSON  phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                     `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                        `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                             `json:"value"`
	JSON  phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                          `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                             `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                                     `json:"value"`
	JSON  phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                                  `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                                     `json:"cloudflare_only"`
	JSON           phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                      `json:"username_expression" api:"required"`
	JSON               phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                         `json:"score_response_header_name"`
	JSON                    phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimitJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimitJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimitJSON struct {
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

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                             `json:"ref"`
	JSON phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleJSON contains the
// JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleJSON struct {
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

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRule) implementsPhaseUpdateResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleAction string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionSetCacheTags PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleAction = "set_cache_tags"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                                        `json:"values"`
	JSON   phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters) AsUnion() PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
// or
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters interface {
	implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                             `json:"values" api:"required"`
	JSON   phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                                `json:"values" api:"required"`
	JSON   phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                             `json:"values" api:"required"`
	JSON   phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsPhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation = "add"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation = "remove"
	PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationSet    PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation = "set"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove, PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                   `json:"username_expression" api:"required"`
	JSON               phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                      `json:"score_response_header_name"`
	JSON                    phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimitJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit]
type phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimitJSON struct {
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

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                                      `json:"ref"`
	JSON phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule]
type phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleJSON struct {
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

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRule) implementsPhaseUpdateResponseEnvelopeResultRules() {
}

// The action to perform when the rule matches.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleAction string

const (
	PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleAction = "transform_response_html"
)

func (r PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionTransformResponseHTML:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters struct {
	// Enables the link maze transformation on the response.
	LinkMaze interface{}                                                                                 `json:"link_maze" api:"required"`
	JSON     phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParametersJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters]
type phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParametersJSON struct {
	LinkMaze    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Configuration for exposed credential checking.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                            `json:"username_expression" api:"required"`
	JSON               phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck]
type phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                               `json:"score_response_header_name"`
	JSON                    phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimitJSON `json:"-"`
}

// phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimitJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit]
type phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimitJSON struct {
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

func (r *PhaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseEnvelopeResultRulesRulesetsTransformResponseHTMLRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type PhaseUpdateResponseEnvelopeResultRulesAction string

const (
	PhaseUpdateResponseEnvelopeResultRulesActionBlock                 PhaseUpdateResponseEnvelopeResultRulesAction = "block"
	PhaseUpdateResponseEnvelopeResultRulesActionChallenge             PhaseUpdateResponseEnvelopeResultRulesAction = "challenge"
	PhaseUpdateResponseEnvelopeResultRulesActionCompressResponse      PhaseUpdateResponseEnvelopeResultRulesAction = "compress_response"
	PhaseUpdateResponseEnvelopeResultRulesActionDDoSDynamic           PhaseUpdateResponseEnvelopeResultRulesAction = "ddos_dynamic"
	PhaseUpdateResponseEnvelopeResultRulesActionExecute               PhaseUpdateResponseEnvelopeResultRulesAction = "execute"
	PhaseUpdateResponseEnvelopeResultRulesActionForceConnectionClose  PhaseUpdateResponseEnvelopeResultRulesAction = "force_connection_close"
	PhaseUpdateResponseEnvelopeResultRulesActionJSChallenge           PhaseUpdateResponseEnvelopeResultRulesAction = "js_challenge"
	PhaseUpdateResponseEnvelopeResultRulesActionLog                   PhaseUpdateResponseEnvelopeResultRulesAction = "log"
	PhaseUpdateResponseEnvelopeResultRulesActionLogCustomField        PhaseUpdateResponseEnvelopeResultRulesAction = "log_custom_field"
	PhaseUpdateResponseEnvelopeResultRulesActionManagedChallenge      PhaseUpdateResponseEnvelopeResultRulesAction = "managed_challenge"
	PhaseUpdateResponseEnvelopeResultRulesActionRedirect              PhaseUpdateResponseEnvelopeResultRulesAction = "redirect"
	PhaseUpdateResponseEnvelopeResultRulesActionRewrite               PhaseUpdateResponseEnvelopeResultRulesAction = "rewrite"
	PhaseUpdateResponseEnvelopeResultRulesActionRoute                 PhaseUpdateResponseEnvelopeResultRulesAction = "route"
	PhaseUpdateResponseEnvelopeResultRulesActionScore                 PhaseUpdateResponseEnvelopeResultRulesAction = "score"
	PhaseUpdateResponseEnvelopeResultRulesActionServeError            PhaseUpdateResponseEnvelopeResultRulesAction = "serve_error"
	PhaseUpdateResponseEnvelopeResultRulesActionSetCacheControl       PhaseUpdateResponseEnvelopeResultRulesAction = "set_cache_control"
	PhaseUpdateResponseEnvelopeResultRulesActionSetCacheSettings      PhaseUpdateResponseEnvelopeResultRulesAction = "set_cache_settings"
	PhaseUpdateResponseEnvelopeResultRulesActionSetCacheTags          PhaseUpdateResponseEnvelopeResultRulesAction = "set_cache_tags"
	PhaseUpdateResponseEnvelopeResultRulesActionSetConfig             PhaseUpdateResponseEnvelopeResultRulesAction = "set_config"
	PhaseUpdateResponseEnvelopeResultRulesActionSkip                  PhaseUpdateResponseEnvelopeResultRulesAction = "skip"
	PhaseUpdateResponseEnvelopeResultRulesActionTransformResponseHTML PhaseUpdateResponseEnvelopeResultRulesAction = "transform_response_html"
)

func (r PhaseUpdateResponseEnvelopeResultRulesAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseEnvelopeResultRulesActionBlock, PhaseUpdateResponseEnvelopeResultRulesActionChallenge, PhaseUpdateResponseEnvelopeResultRulesActionCompressResponse, PhaseUpdateResponseEnvelopeResultRulesActionDDoSDynamic, PhaseUpdateResponseEnvelopeResultRulesActionExecute, PhaseUpdateResponseEnvelopeResultRulesActionForceConnectionClose, PhaseUpdateResponseEnvelopeResultRulesActionJSChallenge, PhaseUpdateResponseEnvelopeResultRulesActionLog, PhaseUpdateResponseEnvelopeResultRulesActionLogCustomField, PhaseUpdateResponseEnvelopeResultRulesActionManagedChallenge, PhaseUpdateResponseEnvelopeResultRulesActionRedirect, PhaseUpdateResponseEnvelopeResultRulesActionRewrite, PhaseUpdateResponseEnvelopeResultRulesActionRoute, PhaseUpdateResponseEnvelopeResultRulesActionScore, PhaseUpdateResponseEnvelopeResultRulesActionServeError, PhaseUpdateResponseEnvelopeResultRulesActionSetCacheControl, PhaseUpdateResponseEnvelopeResultRulesActionSetCacheSettings, PhaseUpdateResponseEnvelopeResultRulesActionSetCacheTags, PhaseUpdateResponseEnvelopeResultRulesActionSetConfig, PhaseUpdateResponseEnvelopeResultRulesActionSkip, PhaseUpdateResponseEnvelopeResultRulesActionTransformResponseHTML:
		return true
	}
	return false
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

// A response object.
type PhaseGetResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseGetResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []PhaseGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// A ruleset object.
	Result PhaseGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success PhaseGetResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Message string `json:"message" api:"required"`
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
	Pointer string                                   `json:"pointer" api:"required"`
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
	Message string `json:"message" api:"required"`
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
	Pointer string                                     `json:"pointer" api:"required"`
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
