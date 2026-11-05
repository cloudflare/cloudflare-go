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
	"github.com/cloudflare/cloudflare-go/v7/shared"
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
func (r *PhaseVersionService) List(ctx context.Context, rulesetPhase Phase, query PhaseVersionListParams, opts ...option.RequestOption) (res *PhaseVersionListResponseUnion, err error) {
	var env PhaseVersionListResponseEnvelope
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
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint/versions", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
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

// A result.
//
// Union satisfied by [PhaseVersionListResponseArray], [shared.UnionString],
// [shared.UnionFloat] or [shared.UnionBool].
type PhaseVersionListResponseUnion interface {
	ImplementsPhaseVersionListResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionListResponseArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type PhaseVersionListResponseArray []interface{}

func (r PhaseVersionListResponseArray) ImplementsPhaseVersionListResponseUnion() {}

// A result.
type PhaseVersionGetResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind Kind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase Phase `json:"phase"`
	// This field can have the runtime type of [[]PhaseVersionGetResponseObjectRule].
	Rules interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                      `json:"version"`
	JSON    phaseVersionGetResponseJSON `json:"-"`
	union   PhaseVersionGetResponseUnion
}

// phaseVersionGetResponseJSON contains the JSON metadata for the struct
// [PhaseVersionGetResponse]
type phaseVersionGetResponseJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PhaseVersionGetResponseUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [PhaseVersionGetResponseObject],
// [PhaseVersionGetResponseObject], [PhaseVersionGetResponseObject],
// [PhaseVersionGetResponseObject], [PhaseVersionGetResponseObject].
func (r PhaseVersionGetResponse) AsUnion() PhaseVersionGetResponseUnion {
	return r.union
}

// A result.
//
// Union satisfied by [PhaseVersionGetResponseObject],
// [PhaseVersionGetResponseObject], [PhaseVersionGetResponseObject],
// [PhaseVersionGetResponseObject] or [PhaseVersionGetResponseObject].
type PhaseVersionGetResponseUnion interface {
	implementsPhaseVersionGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObject{}),
		},
	)
}

type PhaseVersionGetResponseObject struct {
	// The unique ID of the ruleset.
	ID string `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind Kind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase Phase `json:"phase"`
	// The list of rules in the ruleset.
	Rules []PhaseVersionGetResponseObjectRule `json:"rules"`
	// The version of the ruleset.
	Version string                            `json:"version"`
	JSON    phaseVersionGetResponseObjectJSON `json:"-"`
}

// phaseVersionGetResponseObjectJSON contains the JSON metadata for the struct
// [PhaseVersionGetResponseObject]
type phaseVersionGetResponseObjectJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObject) implementsPhaseVersionGetResponse() {}

type PhaseVersionGetResponseObjectRule struct {
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
	// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseObjectRulesChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck],
	// [PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [PhaseVersionGetResponseObjectRulesChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimit],
	// [LogRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit],
	// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleRatelimit],
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                                `json:"version"`
	JSON    phaseVersionGetResponseObjectRuleJSON `json:"-"`
	union   PhaseVersionGetResponseObjectRulesUnion
}

// phaseVersionGetResponseObjectRuleJSON contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRule]
type phaseVersionGetResponseObjectRuleJSON struct {
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

func (r phaseVersionGetResponseObjectRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRule) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [PhaseVersionGetResponseObjectRulesUnion] interface which you
// can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesBlockRule],
// [PhaseVersionGetResponseObjectRulesChallengeRule],
// [PhaseVersionGetResponseObjectRulesResponseCompressionRule],
// [PhaseVersionGetResponseObjectRulesDDoSDynamicRule],
// [PhaseVersionGetResponseObjectRulesExecuteRule],
// [PhaseVersionGetResponseObjectRulesForceConnectionCloseRule],
// [PhaseVersionGetResponseObjectRulesJavaScriptChallengeRule],
// [PhaseVersionGetResponseObjectRulesLogRule],
// [PhaseVersionGetResponseObjectRulesLogCustomFieldRule],
// [PhaseVersionGetResponseObjectRulesManagedChallengeRule],
// [PhaseVersionGetResponseObjectRulesRedirectRule],
// [PhaseVersionGetResponseObjectRulesRewriteRule],
// [PhaseVersionGetResponseObjectRulesRouteRule],
// [PhaseVersionGetResponseObjectRulesScoreRule],
// [PhaseVersionGetResponseObjectRulesServeErrorRule],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRule],
// [PhaseVersionGetResponseObjectRulesSetCacheSettingsRule],
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRule],
// [PhaseVersionGetResponseObjectRulesSetConfigurationRule],
// [PhaseVersionGetResponseObjectRulesSkipRule],
// [PhaseVersionGetResponseObjectRulesTransformResponseHTMLRule].
func (r PhaseVersionGetResponseObjectRule) AsUnion() PhaseVersionGetResponseObjectRulesUnion {
	return r.union
}

// Union satisfied by [PhaseVersionGetResponseObjectRulesBlockRule],
// [PhaseVersionGetResponseObjectRulesChallengeRule],
// [PhaseVersionGetResponseObjectRulesResponseCompressionRule],
// [PhaseVersionGetResponseObjectRulesDDoSDynamicRule],
// [PhaseVersionGetResponseObjectRulesExecuteRule],
// [PhaseVersionGetResponseObjectRulesForceConnectionCloseRule],
// [PhaseVersionGetResponseObjectRulesJavaScriptChallengeRule],
// [PhaseVersionGetResponseObjectRulesLogRule],
// [PhaseVersionGetResponseObjectRulesLogCustomFieldRule],
// [PhaseVersionGetResponseObjectRulesManagedChallengeRule],
// [PhaseVersionGetResponseObjectRulesRedirectRule],
// [PhaseVersionGetResponseObjectRulesRewriteRule],
// [PhaseVersionGetResponseObjectRulesRouteRule],
// [PhaseVersionGetResponseObjectRulesScoreRule],
// [PhaseVersionGetResponseObjectRulesServeErrorRule],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRule],
// [PhaseVersionGetResponseObjectRulesSetCacheSettingsRule],
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRule],
// [PhaseVersionGetResponseObjectRulesSetConfigurationRule],
// [PhaseVersionGetResponseObjectRulesSkipRule] or
// [PhaseVersionGetResponseObjectRulesTransformResponseHTMLRule].
type PhaseVersionGetResponseObjectRulesUnion interface {
	implementsPhaseVersionGetResponseObjectRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesResponseCompressionRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesDDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesJavaScriptChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesRedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesRewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesRouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheTagsRule{}),
			DiscriminatorValue: "set_cache_tags",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetConfigurationRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesSkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PhaseVersionGetResponseObjectRulesTransformResponseHTMLRule{}),
			DiscriminatorValue: "transform_response_html",
		},
	)
}

type PhaseVersionGetResponseObjectRulesBlockRule struct {
	ID         string                                          `json:"id" api:"required"`
	Action     string                                          `json:"action" api:"required"`
	Enabled    bool                                            `json:"enabled" api:"required"`
	Expression string                                          `json:"expression" api:"required"`
	Ref        string                                          `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesBlockRuleJSON `json:"-"`
	BlockRule
}

// phaseVersionGetResponseObjectRulesBlockRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseObjectRulesBlockRule]
type phaseVersionGetResponseObjectRulesBlockRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesBlockRule) implementsPhaseVersionGetResponseObjectRule() {}

type PhaseVersionGetResponseObjectRulesChallengeRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseObjectRulesChallengeRuleAction `json:"action" api:"required"`
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
	ExposedCredentialCheck PhaseVersionGetResponseObjectRulesChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseObjectRulesChallengeRuleRatelimit `json:"ratelimit"`
	JSON      phaseVersionGetResponseObjectRulesChallengeRuleJSON      `json:"-"`
}

// phaseVersionGetResponseObjectRulesChallengeRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseObjectRulesChallengeRule]
type phaseVersionGetResponseObjectRulesChallengeRuleJSON struct {
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

func (r *PhaseVersionGetResponseObjectRulesChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesChallengeRule) implementsPhaseVersionGetResponseObjectRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseObjectRulesChallengeRuleAction string

const (
	PhaseVersionGetResponseObjectRulesChallengeRuleActionChallenge PhaseVersionGetResponseObjectRulesChallengeRuleAction = "challenge"
)

func (r PhaseVersionGetResponseObjectRulesChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseObjectRulesChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                    `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseObjectRulesChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesChallengeRuleExposedCredentialCheck]
type phaseVersionGetResponseObjectRulesChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseObjectRulesChallengeRuleRatelimit struct {
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
	JSON                    phaseVersionGetResponseObjectRulesChallengeRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [PhaseVersionGetResponseObjectRulesChallengeRuleRatelimit]
type phaseVersionGetResponseObjectRulesChallengeRuleRatelimitJSON struct {
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

func (r *PhaseVersionGetResponseObjectRulesChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseObjectRulesResponseCompressionRule struct {
	ID         string                                                        `json:"id" api:"required"`
	Action     string                                                        `json:"action" api:"required"`
	Enabled    bool                                                          `json:"enabled" api:"required"`
	Expression string                                                        `json:"expression" api:"required"`
	Ref        string                                                        `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesResponseCompressionRuleJSON `json:"-"`
	CompressResponseRule
}

// phaseVersionGetResponseObjectRulesResponseCompressionRuleJSON contains the JSON
// metadata for the struct
// [PhaseVersionGetResponseObjectRulesResponseCompressionRule]
type phaseVersionGetResponseObjectRulesResponseCompressionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesResponseCompressionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesResponseCompressionRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesResponseCompressionRule) implementsPhaseVersionGetResponseObjectRule() {
}

type PhaseVersionGetResponseObjectRulesDDoSDynamicRule struct {
	ID         string                                                `json:"id" api:"required"`
	Action     string                                                `json:"action" api:"required"`
	Enabled    bool                                                  `json:"enabled" api:"required"`
	Expression string                                                `json:"expression" api:"required"`
	Ref        string                                                `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesDDoSDynamicRuleJSON `json:"-"`
	DDoSDynamicRule
}

// phaseVersionGetResponseObjectRulesDDoSDynamicRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseObjectRulesDDoSDynamicRule]
type phaseVersionGetResponseObjectRulesDDoSDynamicRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesDDoSDynamicRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesDDoSDynamicRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesDDoSDynamicRule) implementsPhaseVersionGetResponseObjectRule() {
}

type PhaseVersionGetResponseObjectRulesExecuteRule struct {
	ID         string                                            `json:"id" api:"required"`
	Action     string                                            `json:"action" api:"required"`
	Enabled    bool                                              `json:"enabled" api:"required"`
	Expression string                                            `json:"expression" api:"required"`
	Ref        string                                            `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesExecuteRuleJSON `json:"-"`
	ExecuteRule
}

// phaseVersionGetResponseObjectRulesExecuteRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseObjectRulesExecuteRule]
type phaseVersionGetResponseObjectRulesExecuteRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesExecuteRule) implementsPhaseVersionGetResponseObjectRule() {
}

type PhaseVersionGetResponseObjectRulesForceConnectionCloseRule struct {
	ID         string                                                         `json:"id" api:"required"`
	Action     string                                                         `json:"action" api:"required"`
	Enabled    bool                                                           `json:"enabled" api:"required"`
	Expression string                                                         `json:"expression" api:"required"`
	Ref        string                                                         `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesForceConnectionCloseRuleJSON `json:"-"`
	ForceConnectionCloseRule
}

// phaseVersionGetResponseObjectRulesForceConnectionCloseRuleJSON contains the JSON
// metadata for the struct
// [PhaseVersionGetResponseObjectRulesForceConnectionCloseRule]
type phaseVersionGetResponseObjectRulesForceConnectionCloseRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesForceConnectionCloseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesForceConnectionCloseRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesForceConnectionCloseRule) implementsPhaseVersionGetResponseObjectRule() {
}

type PhaseVersionGetResponseObjectRulesJavaScriptChallengeRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleAction `json:"action" api:"required"`
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
	ExposedCredentialCheck PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimit `json:"ratelimit"`
	JSON      phaseVersionGetResponseObjectRulesJavaScriptChallengeRuleJSON      `json:"-"`
}

// phaseVersionGetResponseObjectRulesJavaScriptChallengeRuleJSON contains the JSON
// metadata for the struct
// [PhaseVersionGetResponseObjectRulesJavaScriptChallengeRule]
type phaseVersionGetResponseObjectRulesJavaScriptChallengeRuleJSON struct {
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

func (r *PhaseVersionGetResponseObjectRulesJavaScriptChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesJavaScriptChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesJavaScriptChallengeRule) implementsPhaseVersionGetResponseObjectRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleAction string

const (
	PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleActionJSChallenge PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleAction = "js_challenge"
)

func (r PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                              `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheck]
type phaseVersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                                 `json:"score_response_header_name"`
	JSON                    phaseVersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimitJSON contains
// the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimit]
type phaseVersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimitJSON struct {
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

func (r *PhaseVersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseObjectRulesLogRule struct {
	ID         string                                        `json:"id" api:"required"`
	Action     string                                        `json:"action" api:"required"`
	Enabled    bool                                          `json:"enabled" api:"required"`
	Expression string                                        `json:"expression" api:"required"`
	Ref        string                                        `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesLogRuleJSON `json:"-"`
	LogRule
}

// phaseVersionGetResponseObjectRulesLogRuleJSON contains the JSON metadata for the
// struct [PhaseVersionGetResponseObjectRulesLogRule]
type phaseVersionGetResponseObjectRulesLogRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesLogRule) implementsPhaseVersionGetResponseObjectRule() {}

type PhaseVersionGetResponseObjectRulesLogCustomFieldRule struct {
	ID         string                                                   `json:"id" api:"required"`
	Action     string                                                   `json:"action" api:"required"`
	Enabled    bool                                                     `json:"enabled" api:"required"`
	Expression string                                                   `json:"expression" api:"required"`
	Ref        string                                                   `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesLogCustomFieldRuleJSON `json:"-"`
	LogCustomFieldRule
}

// phaseVersionGetResponseObjectRulesLogCustomFieldRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseObjectRulesLogCustomFieldRule]
type phaseVersionGetResponseObjectRulesLogCustomFieldRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesLogCustomFieldRule) implementsPhaseVersionGetResponseObjectRule() {
}

type PhaseVersionGetResponseObjectRulesManagedChallengeRule struct {
	ID         string                                                     `json:"id" api:"required"`
	Action     string                                                     `json:"action" api:"required"`
	Enabled    bool                                                       `json:"enabled" api:"required"`
	Expression string                                                     `json:"expression" api:"required"`
	Ref        string                                                     `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesManagedChallengeRuleJSON `json:"-"`
	ManagedChallengeRule
}

// phaseVersionGetResponseObjectRulesManagedChallengeRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseObjectRulesManagedChallengeRule]
type phaseVersionGetResponseObjectRulesManagedChallengeRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesManagedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesManagedChallengeRule) implementsPhaseVersionGetResponseObjectRule() {
}

type PhaseVersionGetResponseObjectRulesRedirectRule struct {
	ID         string                                             `json:"id" api:"required"`
	Action     string                                             `json:"action" api:"required"`
	Enabled    bool                                               `json:"enabled" api:"required"`
	Expression string                                             `json:"expression" api:"required"`
	Ref        string                                             `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesRedirectRuleJSON `json:"-"`
	RedirectRule
}

// phaseVersionGetResponseObjectRulesRedirectRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseObjectRulesRedirectRule]
type phaseVersionGetResponseObjectRulesRedirectRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesRedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesRedirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesRedirectRule) implementsPhaseVersionGetResponseObjectRule() {
}

type PhaseVersionGetResponseObjectRulesRewriteRule struct {
	ID         string                                            `json:"id" api:"required"`
	Action     string                                            `json:"action" api:"required"`
	Enabled    bool                                              `json:"enabled" api:"required"`
	Expression string                                            `json:"expression" api:"required"`
	Ref        string                                            `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesRewriteRuleJSON `json:"-"`
	RewriteRule
}

// phaseVersionGetResponseObjectRulesRewriteRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseObjectRulesRewriteRule]
type phaseVersionGetResponseObjectRulesRewriteRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesRewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesRewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesRewriteRule) implementsPhaseVersionGetResponseObjectRule() {
}

type PhaseVersionGetResponseObjectRulesRouteRule struct {
	ID         string                                          `json:"id" api:"required"`
	Action     string                                          `json:"action" api:"required"`
	Enabled    bool                                            `json:"enabled" api:"required"`
	Expression string                                          `json:"expression" api:"required"`
	Ref        string                                          `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesRouteRuleJSON `json:"-"`
	RouteRule
}

// phaseVersionGetResponseObjectRulesRouteRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseObjectRulesRouteRule]
type phaseVersionGetResponseObjectRulesRouteRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesRouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesRouteRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesRouteRule) implementsPhaseVersionGetResponseObjectRule() {}

type PhaseVersionGetResponseObjectRulesScoreRule struct {
	ID         string                                          `json:"id" api:"required"`
	Action     string                                          `json:"action" api:"required"`
	Enabled    bool                                            `json:"enabled" api:"required"`
	Expression string                                          `json:"expression" api:"required"`
	Ref        string                                          `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesScoreRuleJSON `json:"-"`
	ScoreRule
}

// phaseVersionGetResponseObjectRulesScoreRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseObjectRulesScoreRule]
type phaseVersionGetResponseObjectRulesScoreRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesScoreRule) implementsPhaseVersionGetResponseObjectRule() {}

type PhaseVersionGetResponseObjectRulesServeErrorRule struct {
	ID         string                                               `json:"id" api:"required"`
	Action     string                                               `json:"action" api:"required"`
	Enabled    bool                                                 `json:"enabled" api:"required"`
	Expression string                                               `json:"expression" api:"required"`
	Ref        string                                               `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesServeErrorRuleJSON `json:"-"`
	ServeErrorRule
}

// phaseVersionGetResponseObjectRulesServeErrorRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseObjectRulesServeErrorRule]
type phaseVersionGetResponseObjectRulesServeErrorRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesServeErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesServeErrorRule) implementsPhaseVersionGetResponseObjectRule() {
}

type PhaseVersionGetResponseObjectRulesSetCacheControlRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action  PhaseVersionGetResponseObjectRulesSetCacheControlRuleAction `json:"action" api:"required"`
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
	ActionParameters PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseVersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseObjectRulesSetCacheControlRuleRatelimit `json:"ratelimit"`
	JSON      phaseVersionGetResponseObjectRulesSetCacheControlRuleJSON      `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseObjectRulesSetCacheControlRule]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleJSON struct {
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

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRule) implementsPhaseVersionGetResponseObjectRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleAction string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionSetCacheControl PhaseVersionGetResponseObjectRulesSetCacheControlRuleAction = "set_cache_control"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParameters]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersJSON struct {
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

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                               `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                           `json:"value"`
	JSON  phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                    `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                   `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                    `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                   `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                      `json:"qualifiers"`
	JSON       phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                     `json:"qualifiers"`
	JSON       phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                             `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                 `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                             `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                      `json:"qualifiers"`
	JSON       phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                     `json:"qualifiers"`
	JSON       phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                     `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                 `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                    `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                            `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                            `json:"value"`
	JSON  phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                 `json:"value"`
	JSON  phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                 `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                         `json:"value"`
	JSON  phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                      `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                         `json:"cloudflare_only"`
	JSON           phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsPhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, PhaseVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                          `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheck]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseObjectRulesSetCacheControlRuleRatelimit struct {
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
	JSON                    phaseVersionGetResponseObjectRulesSetCacheControlRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheControlRuleRatelimitJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheControlRuleRatelimit]
type phaseVersionGetResponseObjectRulesSetCacheControlRuleRatelimitJSON struct {
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

func (r *PhaseVersionGetResponseObjectRulesSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseObjectRulesSetCacheSettingsRule struct {
	ID         string                                                     `json:"id" api:"required"`
	Action     string                                                     `json:"action" api:"required"`
	Enabled    bool                                                       `json:"enabled" api:"required"`
	Expression string                                                     `json:"expression" api:"required"`
	Ref        string                                                     `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesSetCacheSettingsRuleJSON `json:"-"`
	SetCacheSettingsRule
}

// phaseVersionGetResponseObjectRulesSetCacheSettingsRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseObjectRulesSetCacheSettingsRule]
type phaseVersionGetResponseObjectRulesSetCacheSettingsRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheSettingsRule) implementsPhaseVersionGetResponseObjectRule() {
}

type PhaseVersionGetResponseObjectRulesSetCacheTagsRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action  PhaseVersionGetResponseObjectRulesSetCacheTagsRuleAction `json:"action" api:"required"`
	Enabled bool                                                     `json:"enabled" api:"required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression" api:"required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref" api:"required"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseVersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseObjectRulesSetCacheTagsRuleRatelimit `json:"ratelimit"`
	JSON      phaseVersionGetResponseObjectRulesSetCacheTagsRuleJSON      `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheTagsRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseObjectRulesSetCacheTagsRule]
type phaseVersionGetResponseObjectRulesSetCacheTagsRuleJSON struct {
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

func (r *PhaseVersionGetResponseObjectRulesSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRule) implementsPhaseVersionGetResponseObjectRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleAction string

const (
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionSetCacheTags PhaseVersionGetResponseObjectRulesSetCacheTagsRuleAction = "set_cache_tags"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                            `json:"values"`
	JSON   phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersUnion
}

// phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersJSON contains
// the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters]
type phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters) AsUnion() PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersUnion {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues]
// or
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersUnion interface {
	implementsPhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                 `json:"values" api:"required"`
	JSON   phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues]
type phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsPhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsPhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                    `json:"values" api:"required"`
	JSON   phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsPhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsPhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                                 `json:"values" api:"required"`
	JSON   phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues]
type phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsPhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsPhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperation string

const (
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperationAdd    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperation = "add"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperationRemove PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperation = "remove"
	PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperationSet    PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperation = "set"
)

func (r PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperationAdd, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperationRemove, PhaseVersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                       `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheck]
type phaseVersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseObjectRulesSetCacheTagsRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                          `json:"score_response_header_name"`
	JSON                    phaseVersionGetResponseObjectRulesSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesSetCacheTagsRuleRatelimitJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesSetCacheTagsRuleRatelimit]
type phaseVersionGetResponseObjectRulesSetCacheTagsRuleRatelimitJSON struct {
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

func (r *PhaseVersionGetResponseObjectRulesSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionGetResponseObjectRulesSetConfigurationRule struct {
	ID         string                                                     `json:"id" api:"required"`
	Action     string                                                     `json:"action" api:"required"`
	Enabled    bool                                                       `json:"enabled" api:"required"`
	Expression string                                                     `json:"expression" api:"required"`
	Ref        string                                                     `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesSetConfigurationRuleJSON `json:"-"`
	SetConfigRule
}

// phaseVersionGetResponseObjectRulesSetConfigurationRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseObjectRulesSetConfigurationRule]
type phaseVersionGetResponseObjectRulesSetConfigurationRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSetConfigurationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSetConfigurationRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSetConfigurationRule) implementsPhaseVersionGetResponseObjectRule() {
}

type PhaseVersionGetResponseObjectRulesSkipRule struct {
	ID         string                                         `json:"id" api:"required"`
	Action     string                                         `json:"action" api:"required"`
	Enabled    bool                                           `json:"enabled" api:"required"`
	Expression string                                         `json:"expression" api:"required"`
	Ref        string                                         `json:"ref" api:"required"`
	JSON       phaseVersionGetResponseObjectRulesSkipRuleJSON `json:"-"`
	SkipRule
}

// phaseVersionGetResponseObjectRulesSkipRuleJSON contains the JSON metadata for
// the struct [PhaseVersionGetResponseObjectRulesSkipRule]
type phaseVersionGetResponseObjectRulesSkipRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesSkipRule) implementsPhaseVersionGetResponseObjectRule() {}

type PhaseVersionGetResponseObjectRulesTransformResponseHTMLRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action  PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleAction `json:"action" api:"required"`
	Enabled bool                                                              `json:"enabled" api:"required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression" api:"required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref" api:"required"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimit `json:"ratelimit"`
	JSON      phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleJSON      `json:"-"`
}

// phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleJSON contains the
// JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesTransformResponseHTMLRule]
type phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleJSON struct {
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

func (r *PhaseVersionGetResponseObjectRulesTransformResponseHTMLRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseObjectRulesTransformResponseHTMLRule) implementsPhaseVersionGetResponseObjectRule() {
}

// The action to perform when the rule matches.
type PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleAction string

const (
	PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleActionTransformResponseHTML PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleAction = "transform_response_html"
)

func (r PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleActionTransformResponseHTML:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleActionParameters struct {
	// Enables the link maze transformation on the response.
	LinkMaze interface{}                                                                     `json:"link_maze" api:"required"`
	JSON     phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleActionParametersJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleActionParametersJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleActionParameters]
type phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleActionParametersJSON struct {
	LinkMaze    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Configuration for exposed credential checking.
type PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                                `json:"username_expression" api:"required"`
	JSON               phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheckJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheck]
type phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimit struct {
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
	JSON                    phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimitJSON `json:"-"`
}

// phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimitJSON
// contains the JSON metadata for the struct
// [PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimit]
type phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimitJSON struct {
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

func (r *PhaseVersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type PhaseVersionListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type PhaseVersionListResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseVersionListResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []PhaseVersionListResponseEnvelopeMessages `json:"messages" api:"required"`
	// A result.
	Result PhaseVersionListResponseUnion `json:"result" api:"required"`
	// Whether the API call was successful.
	Success PhaseVersionListResponseEnvelopeSuccess `json:"success" api:"required"`
	// Information to navigate the results.
	ResultInfo PhaseVersionListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       phaseVersionListResponseEnvelopeJSON       `json:"-"`
}

// phaseVersionListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PhaseVersionListResponseEnvelope]
type phaseVersionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseVersionListResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseVersionListResponseEnvelopeErrorsSource `json:"source"`
	JSON   phaseVersionListResponseEnvelopeErrorsJSON   `json:"-"`
}

// phaseVersionListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PhaseVersionListResponseEnvelopeErrors]
type phaseVersionListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseVersionListResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                           `json:"pointer" api:"required"`
	JSON    phaseVersionListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// phaseVersionListResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [PhaseVersionListResponseEnvelopeErrorsSource]
type phaseVersionListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type PhaseVersionListResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source PhaseVersionListResponseEnvelopeMessagesSource `json:"source"`
	JSON   phaseVersionListResponseEnvelopeMessagesJSON   `json:"-"`
}

// phaseVersionListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PhaseVersionListResponseEnvelopeMessages]
type phaseVersionListResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type PhaseVersionListResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                             `json:"pointer" api:"required"`
	JSON    phaseVersionListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// phaseVersionListResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [PhaseVersionListResponseEnvelopeMessagesSource]
type phaseVersionListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type PhaseVersionListResponseEnvelopeSuccess bool

const (
	PhaseVersionListResponseEnvelopeSuccessTrue PhaseVersionListResponseEnvelopeSuccess = true
)

func (r PhaseVersionListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PhaseVersionListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

// Information to navigate the results.
type PhaseVersionListResponseEnvelopeResultInfo struct {
	// The set of cursors.
	Cursors PhaseVersionListResponseEnvelopeResultInfoCursors `json:"cursors"`
	JSON    phaseVersionListResponseEnvelopeResultInfoJSON    `json:"-"`
}

// phaseVersionListResponseEnvelopeResultInfoJSON contains the JSON metadata for
// the struct [PhaseVersionListResponseEnvelopeResultInfo]
type phaseVersionListResponseEnvelopeResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

// The set of cursors.
type PhaseVersionListResponseEnvelopeResultInfoCursors struct {
	// The cursor to use for the next page.
	After string                                                `json:"after" api:"required"`
	JSON  phaseVersionListResponseEnvelopeResultInfoCursorsJSON `json:"-"`
}

// phaseVersionListResponseEnvelopeResultInfoCursorsJSON contains the JSON metadata
// for the struct [PhaseVersionListResponseEnvelopeResultInfoCursors]
type phaseVersionListResponseEnvelopeResultInfoCursorsJSON struct {
	After       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseVersionListResponseEnvelopeResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionListResponseEnvelopeResultInfoCursorsJSON) RawJSON() string {
	return r.raw
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
	// A result.
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
