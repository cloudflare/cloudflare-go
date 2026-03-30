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
func (r *PhaseService) Update(ctx context.Context, rulesetPhase Phase, params PhaseUpdateParams, opts ...option.RequestOption) (res *PhaseUpdateResponse, err error) {
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
type PhaseUpdateResponse struct {
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
	Rules []PhaseUpdateResponseRule `json:"rules" api:"required"`
	// The version of the ruleset.
	Version string `json:"version" api:"required"`
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

type PhaseUpdateResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [LogCustomFieldRuleActionParameters],
	// [RedirectRuleActionParameters], [RewriteRuleActionParameters],
	// [RouteRuleActionParameters], [ScoreRuleActionParameters],
	// [ServeErrorRuleActionParameters],
	// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [PhaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [PhaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [PhaseUpdateResponseRulesRulesetsChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [PhaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimit], [LogRuleRatelimit],
	// [LogCustomFieldRuleRatelimit], [ManagedChallengeRuleRatelimit],
	// [RedirectRuleRatelimit], [RewriteRuleRatelimit], [RouteRuleRatelimit],
	// [ScoreRuleRatelimit], [ServeErrorRuleRatelimit],
	// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimit],
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref   string                      `json:"ref"`
	JSON  phaseUpdateResponseRuleJSON `json:"-"`
	union PhaseUpdateResponseRulesUnion
}

// phaseUpdateResponseRuleJSON contains the JSON metadata for the struct
// [PhaseUpdateResponseRule]
type phaseUpdateResponseRuleJSON struct {
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

func (r phaseUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PhaseUpdateResponseRulesUnion] interface which you can cast
// to the specific types for more type safety.
//
// Possible runtime types of the union are [BlockRule],
// [PhaseUpdateResponseRulesRulesetsChallengeRule], [CompressResponseRule],
// [DDoSDynamicRule], [ExecuteRule], [ForceConnectionCloseRule],
// [PhaseUpdateResponseRulesRulesetsJSChallengeRule], [LogRule],
// [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule], [RewriteRule],
// [RouteRule], [ScoreRule], [ServeErrorRule],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRule], [SetCacheSettingsRule],
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRule], [SetConfigRule], [SkipRule].
func (r PhaseUpdateResponseRule) AsUnion() PhaseUpdateResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [PhaseUpdateResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [DDoSDynamicRule], [ExecuteRule],
// [ForceConnectionCloseRule], [PhaseUpdateResponseRulesRulesetsJSChallengeRule],
// [LogRule], [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule],
// [RewriteRule], [RouteRule], [ScoreRule], [ServeErrorRule],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRule], [SetCacheSettingsRule],
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRule], [SetConfigRule] or
// [SkipRule].
type PhaseUpdateResponseRulesUnion interface {
	implementsPhaseUpdateResponseRule()
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
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsJSChallengeRule{}),
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
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(SetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheTagsRule{}),
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

type PhaseUpdateResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
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
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseUpdateResponseRulesRulesetsChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                            `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata for
// the struct [PhaseUpdateResponseRulesRulesetsChallengeRule]
type phaseUpdateResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsChallengeRule) implementsPhaseUpdateResponseRule() {}

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

// Configuration for exposed credential checking.
type PhaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                  `json:"username_expression" api:"required"`
	JSON               phaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck]
type phaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateResponseRulesRulesetsChallengeRuleRatelimit struct {
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
	JSON                    phaseUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [PhaseUpdateResponseRulesRulesetsChallengeRuleRatelimit]
type phaseUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                              `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata
// for the struct [PhaseUpdateResponseRulesRulesetsJSChallengeRule]
type phaseUpdateResponseRulesRulesetsJSChallengeRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsJSChallengeRule) implementsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsJSChallengeRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsJSChallengeRuleActionJSChallenge PhaseUpdateResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r PhaseUpdateResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                    `json:"username_expression" api:"required"`
	JSON               phaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck]
type phaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsJSChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimit struct {
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
	JSON                    phaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [PhaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimit]
type phaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsJSChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseRulesRulesetsSetCacheControlRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsSetCacheControlRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseUpdateResponseRulesRulesetsSetCacheControlRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                                  `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsSetCacheControlRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleJSON contains the JSON
// metadata for the struct [PhaseUpdateResponseRulesRulesetsSetCacheControlRule]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRule) implementsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionSetCacheControl PhaseUpdateResponseRulesRulesetsSetCacheControlRuleAction = "set_cache_control"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersJSON contains
// the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParameters]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                             `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                         `json:"value"`
	JSON  phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                  `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                 `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                  `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                 `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                    `json:"qualifiers"`
	JSON       phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                   `json:"qualifiers"`
	JSON       phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                           `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                               `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                    `json:"qualifiers"`
	JSON       phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                   `json:"qualifiers"`
	JSON       phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                   `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                               `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                  `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                          `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                          `json:"value"`
	JSON  phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                          `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                               `json:"value"`
	JSON  phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                               `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                       `json:"value"`
	JSON  phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                    `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                       `json:"cloudflare_only"`
	JSON           phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsPhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, PhaseUpdateResponseRulesRulesetsSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                        `json:"username_expression" api:"required"`
	JSON               phaseUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateResponseRulesRulesetsSetCacheControlRuleRatelimit struct {
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
	JSON                    phaseUpdateResponseRulesRulesetsSetCacheControlRuleRatelimitJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheControlRuleRatelimitJSON contains the
// JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheControlRuleRatelimit]
type phaseUpdateResponseRulesRulesetsSetCacheControlRuleRatelimitJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseUpdateResponseRulesRulesetsSetCacheTagsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimit `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref  string                                               `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsSetCacheTagsRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheTagsRuleJSON contains the JSON metadata
// for the struct [PhaseUpdateResponseRulesRulesetsSetCacheTagsRule]
type phaseUpdateResponseRulesRulesetsSetCacheTagsRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRule) implementsPhaseUpdateResponseRule() {}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionSetCacheTags PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleAction = "set_cache_tags"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                          `json:"values"`
	JSON   phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion
}

// phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON contains
// the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters]
type phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters) AsUnion() PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
// or
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion interface {
	implementsPhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                               `json:"values" api:"required"`
	JSON   phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues]
type phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsPhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsPhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                  `json:"values" api:"required"`
	JSON   phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsPhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsPhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                               `json:"values" api:"required"`
	JSON   phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues]
type phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsPhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsPhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation string

const (
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "add"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "remove"
	PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationSet    PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation = "set"
)

func (r PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationAdd, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationRemove, PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                     `json:"username_expression" api:"required"`
	JSON               phaseUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck]
type phaseUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimit struct {
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
	JSON                    phaseUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON contains the JSON
// metadata for the struct
// [PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimit]
type phaseUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesAction string

const (
	PhaseUpdateResponseRulesActionBlock                PhaseUpdateResponseRulesAction = "block"
	PhaseUpdateResponseRulesActionChallenge            PhaseUpdateResponseRulesAction = "challenge"
	PhaseUpdateResponseRulesActionCompressResponse     PhaseUpdateResponseRulesAction = "compress_response"
	PhaseUpdateResponseRulesActionDDoSDynamic          PhaseUpdateResponseRulesAction = "ddos_dynamic"
	PhaseUpdateResponseRulesActionExecute              PhaseUpdateResponseRulesAction = "execute"
	PhaseUpdateResponseRulesActionForceConnectionClose PhaseUpdateResponseRulesAction = "force_connection_close"
	PhaseUpdateResponseRulesActionJSChallenge          PhaseUpdateResponseRulesAction = "js_challenge"
	PhaseUpdateResponseRulesActionLog                  PhaseUpdateResponseRulesAction = "log"
	PhaseUpdateResponseRulesActionLogCustomField       PhaseUpdateResponseRulesAction = "log_custom_field"
	PhaseUpdateResponseRulesActionManagedChallenge     PhaseUpdateResponseRulesAction = "managed_challenge"
	PhaseUpdateResponseRulesActionRedirect             PhaseUpdateResponseRulesAction = "redirect"
	PhaseUpdateResponseRulesActionRewrite              PhaseUpdateResponseRulesAction = "rewrite"
	PhaseUpdateResponseRulesActionRoute                PhaseUpdateResponseRulesAction = "route"
	PhaseUpdateResponseRulesActionScore                PhaseUpdateResponseRulesAction = "score"
	PhaseUpdateResponseRulesActionServeError           PhaseUpdateResponseRulesAction = "serve_error"
	PhaseUpdateResponseRulesActionSetCacheControl      PhaseUpdateResponseRulesAction = "set_cache_control"
	PhaseUpdateResponseRulesActionSetCacheSettings     PhaseUpdateResponseRulesAction = "set_cache_settings"
	PhaseUpdateResponseRulesActionSetCacheTags         PhaseUpdateResponseRulesAction = "set_cache_tags"
	PhaseUpdateResponseRulesActionSetConfig            PhaseUpdateResponseRulesAction = "set_config"
	PhaseUpdateResponseRulesActionSkip                 PhaseUpdateResponseRulesAction = "skip"
)

func (r PhaseUpdateResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesActionBlock, PhaseUpdateResponseRulesActionChallenge, PhaseUpdateResponseRulesActionCompressResponse, PhaseUpdateResponseRulesActionDDoSDynamic, PhaseUpdateResponseRulesActionExecute, PhaseUpdateResponseRulesActionForceConnectionClose, PhaseUpdateResponseRulesActionJSChallenge, PhaseUpdateResponseRulesActionLog, PhaseUpdateResponseRulesActionLogCustomField, PhaseUpdateResponseRulesActionManagedChallenge, PhaseUpdateResponseRulesActionRedirect, PhaseUpdateResponseRulesActionRewrite, PhaseUpdateResponseRulesActionRoute, PhaseUpdateResponseRulesActionScore, PhaseUpdateResponseRulesActionServeError, PhaseUpdateResponseRulesActionSetCacheControl, PhaseUpdateResponseRulesActionSetCacheSettings, PhaseUpdateResponseRulesActionSetCacheTags, PhaseUpdateResponseRulesActionSetConfig, PhaseUpdateResponseRulesActionSkip:
		return true
	}
	return false
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
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters].
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
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck].
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
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit].
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
// [SetConfigRule], [SkipRule].
func (r PhaseGetResponseRule) AsUnion() PhaseGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [BlockRule], [PhaseGetResponseRulesRulesetsChallengeRule],
// [CompressResponseRule], [DDoSDynamicRule], [ExecuteRule],
// [ForceConnectionCloseRule], [PhaseGetResponseRulesRulesetsJSChallengeRule],
// [LogRule], [LogCustomFieldRule], [ManagedChallengeRule], [RedirectRule],
// [RewriteRule], [RouteRule], [ScoreRule], [ServeErrorRule],
// [PhaseGetResponseRulesRulesetsSetCacheControlRule], [SetCacheSettingsRule],
// [PhaseGetResponseRulesRulesetsSetCacheTagsRule], [SetConfigRule] or [SkipRule].
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

// The action to perform when the rule matches.
type PhaseGetResponseRulesAction string

const (
	PhaseGetResponseRulesActionBlock                PhaseGetResponseRulesAction = "block"
	PhaseGetResponseRulesActionChallenge            PhaseGetResponseRulesAction = "challenge"
	PhaseGetResponseRulesActionCompressResponse     PhaseGetResponseRulesAction = "compress_response"
	PhaseGetResponseRulesActionDDoSDynamic          PhaseGetResponseRulesAction = "ddos_dynamic"
	PhaseGetResponseRulesActionExecute              PhaseGetResponseRulesAction = "execute"
	PhaseGetResponseRulesActionForceConnectionClose PhaseGetResponseRulesAction = "force_connection_close"
	PhaseGetResponseRulesActionJSChallenge          PhaseGetResponseRulesAction = "js_challenge"
	PhaseGetResponseRulesActionLog                  PhaseGetResponseRulesAction = "log"
	PhaseGetResponseRulesActionLogCustomField       PhaseGetResponseRulesAction = "log_custom_field"
	PhaseGetResponseRulesActionManagedChallenge     PhaseGetResponseRulesAction = "managed_challenge"
	PhaseGetResponseRulesActionRedirect             PhaseGetResponseRulesAction = "redirect"
	PhaseGetResponseRulesActionRewrite              PhaseGetResponseRulesAction = "rewrite"
	PhaseGetResponseRulesActionRoute                PhaseGetResponseRulesAction = "route"
	PhaseGetResponseRulesActionScore                PhaseGetResponseRulesAction = "score"
	PhaseGetResponseRulesActionServeError           PhaseGetResponseRulesAction = "serve_error"
	PhaseGetResponseRulesActionSetCacheControl      PhaseGetResponseRulesAction = "set_cache_control"
	PhaseGetResponseRulesActionSetCacheSettings     PhaseGetResponseRulesAction = "set_cache_settings"
	PhaseGetResponseRulesActionSetCacheTags         PhaseGetResponseRulesAction = "set_cache_tags"
	PhaseGetResponseRulesActionSetConfig            PhaseGetResponseRulesAction = "set_config"
	PhaseGetResponseRulesActionSkip                 PhaseGetResponseRulesAction = "skip"
)

func (r PhaseGetResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesActionBlock, PhaseGetResponseRulesActionChallenge, PhaseGetResponseRulesActionCompressResponse, PhaseGetResponseRulesActionDDoSDynamic, PhaseGetResponseRulesActionExecute, PhaseGetResponseRulesActionForceConnectionClose, PhaseGetResponseRulesActionJSChallenge, PhaseGetResponseRulesActionLog, PhaseGetResponseRulesActionLogCustomField, PhaseGetResponseRulesActionManagedChallenge, PhaseGetResponseRulesActionRedirect, PhaseGetResponseRulesActionRewrite, PhaseGetResponseRulesActionRoute, PhaseGetResponseRulesActionScore, PhaseGetResponseRulesActionServeError, PhaseGetResponseRulesActionSetCacheControl, PhaseGetResponseRulesActionSetCacheSettings, PhaseGetResponseRulesActionSetCacheTags, PhaseGetResponseRulesActionSetConfig, PhaseGetResponseRulesActionSkip:
		return true
	}
	return false
}

type PhaseUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
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

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesAction string

const (
	PhaseUpdateParamsRulesActionBlock                PhaseUpdateParamsRulesAction = "block"
	PhaseUpdateParamsRulesActionChallenge            PhaseUpdateParamsRulesAction = "challenge"
	PhaseUpdateParamsRulesActionCompressResponse     PhaseUpdateParamsRulesAction = "compress_response"
	PhaseUpdateParamsRulesActionDDoSDynamic          PhaseUpdateParamsRulesAction = "ddos_dynamic"
	PhaseUpdateParamsRulesActionExecute              PhaseUpdateParamsRulesAction = "execute"
	PhaseUpdateParamsRulesActionForceConnectionClose PhaseUpdateParamsRulesAction = "force_connection_close"
	PhaseUpdateParamsRulesActionJSChallenge          PhaseUpdateParamsRulesAction = "js_challenge"
	PhaseUpdateParamsRulesActionLog                  PhaseUpdateParamsRulesAction = "log"
	PhaseUpdateParamsRulesActionLogCustomField       PhaseUpdateParamsRulesAction = "log_custom_field"
	PhaseUpdateParamsRulesActionManagedChallenge     PhaseUpdateParamsRulesAction = "managed_challenge"
	PhaseUpdateParamsRulesActionRedirect             PhaseUpdateParamsRulesAction = "redirect"
	PhaseUpdateParamsRulesActionRewrite              PhaseUpdateParamsRulesAction = "rewrite"
	PhaseUpdateParamsRulesActionRoute                PhaseUpdateParamsRulesAction = "route"
	PhaseUpdateParamsRulesActionScore                PhaseUpdateParamsRulesAction = "score"
	PhaseUpdateParamsRulesActionServeError           PhaseUpdateParamsRulesAction = "serve_error"
	PhaseUpdateParamsRulesActionSetCacheControl      PhaseUpdateParamsRulesAction = "set_cache_control"
	PhaseUpdateParamsRulesActionSetCacheSettings     PhaseUpdateParamsRulesAction = "set_cache_settings"
	PhaseUpdateParamsRulesActionSetCacheTags         PhaseUpdateParamsRulesAction = "set_cache_tags"
	PhaseUpdateParamsRulesActionSetConfig            PhaseUpdateParamsRulesAction = "set_config"
	PhaseUpdateParamsRulesActionSkip                 PhaseUpdateParamsRulesAction = "skip"
)

func (r PhaseUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesActionBlock, PhaseUpdateParamsRulesActionChallenge, PhaseUpdateParamsRulesActionCompressResponse, PhaseUpdateParamsRulesActionDDoSDynamic, PhaseUpdateParamsRulesActionExecute, PhaseUpdateParamsRulesActionForceConnectionClose, PhaseUpdateParamsRulesActionJSChallenge, PhaseUpdateParamsRulesActionLog, PhaseUpdateParamsRulesActionLogCustomField, PhaseUpdateParamsRulesActionManagedChallenge, PhaseUpdateParamsRulesActionRedirect, PhaseUpdateParamsRulesActionRewrite, PhaseUpdateParamsRulesActionRoute, PhaseUpdateParamsRulesActionScore, PhaseUpdateParamsRulesActionServeError, PhaseUpdateParamsRulesActionSetCacheControl, PhaseUpdateParamsRulesActionSetCacheSettings, PhaseUpdateParamsRulesActionSetCacheTags, PhaseUpdateParamsRulesActionSetConfig, PhaseUpdateParamsRulesActionSkip:
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
	// A ruleset object.
	Result PhaseUpdateResponse `json:"result" api:"required"`
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
