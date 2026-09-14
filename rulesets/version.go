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
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/tidwall/gjson"
)

// VersionService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVersionService] method instead.
type VersionService struct {
	Options []option.RequestOption
}

// NewVersionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVersionService(opts ...option.RequestOption) (r *VersionService) {
	r = &VersionService{}
	r.Options = opts
	return
}

// Fetches the versions of an account or zone ruleset.
func (r *VersionService) List(ctx context.Context, rulesetID string, query VersionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[VersionListResponse], err error) {
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
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions", accountOrZone, accountOrZoneID, rulesetID)
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

// Fetches the versions of an account or zone ruleset.
func (r *VersionService) ListAutoPaging(ctx context.Context, rulesetID string, query VersionListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[VersionListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, rulesetID, query, opts...))
}

// Deletes an existing version of an account or zone ruleset.
func (r *VersionService) Delete(ctx context.Context, rulesetID string, rulesetVersion string, params VersionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
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
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return err
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return err
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions/%s", accountOrZone, accountOrZoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

// Fetches a specific version of an account or zone ruleset.
func (r *VersionService) Get(ctx context.Context, rulesetID string, rulesetVersion string, query VersionGetParams, opts ...option.RequestOption) (res *VersionGetResponse, err error) {
	var env VersionGetResponseEnvelope
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
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return nil, err
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return nil, err
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions/%s", accountOrZone, accountOrZoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// A ruleset object.
type VersionListResponse struct {
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
	Description string                  `json:"description"`
	JSON        versionListResponseJSON `json:"-"`
}

// versionListResponseJSON contains the JSON metadata for the struct
// [VersionListResponse]
type versionListResponseJSON struct {
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

func (r *VersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseJSON) RawJSON() string {
	return r.raw
}

// A ruleset object.
type VersionGetResponse struct {
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
	Rules []VersionGetResponseRule `json:"rules" api:"required"`
	// The version of the ruleset.
	Version string `json:"version" api:"required"`
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

type VersionGetResponseRule struct {
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
	// [VersionGetResponseRulesSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [VersionGetResponseRulesSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [VersionGetResponseRulesTransformResponseHTMLRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [VersionGetResponseRulesChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [VersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [VersionGetResponseRulesSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [VersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck],
	// [VersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [VersionGetResponseRulesChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [VersionGetResponseRulesJavaScriptChallengeRuleRatelimit], [LogRuleRatelimit],
	// [LogCustomFieldRuleRatelimit], [ManagedChallengeRuleRatelimit],
	// [RedirectRuleRatelimit], [RewriteRuleRatelimit], [RouteRuleRatelimit],
	// [ScoreRuleRatelimit], [ServeErrorRuleRatelimit],
	// [VersionGetResponseRulesSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [VersionGetResponseRulesSetCacheTagsRuleRatelimit], [SetConfigRuleRatelimit],
	// [SkipRuleRatelimit],
	// [VersionGetResponseRulesTransformResponseHTMLRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                     `json:"version"`
	JSON    versionGetResponseRuleJSON `json:"-"`
	union   VersionGetResponseRulesUnion
}

// versionGetResponseRuleJSON contains the JSON metadata for the struct
// [VersionGetResponseRule]
type versionGetResponseRuleJSON struct {
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

func (r versionGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [VersionGetResponseRulesUnion] interface which you can cast to
// the specific types for more type safety.
//
// Possible runtime types of the union are [VersionGetResponseRulesBlockRule],
// [VersionGetResponseRulesChallengeRule],
// [VersionGetResponseRulesResponseCompressionRule],
// [VersionGetResponseRulesDDoSDynamicRule], [VersionGetResponseRulesExecuteRule],
// [VersionGetResponseRulesForceConnectionCloseRule],
// [VersionGetResponseRulesJavaScriptChallengeRule],
// [VersionGetResponseRulesLogRule], [VersionGetResponseRulesLogCustomFieldRule],
// [VersionGetResponseRulesManagedChallengeRule],
// [VersionGetResponseRulesRedirectRule], [VersionGetResponseRulesRewriteRule],
// [VersionGetResponseRulesRouteRule], [VersionGetResponseRulesScoreRule],
// [VersionGetResponseRulesServeErrorRule],
// [VersionGetResponseRulesSetCacheControlRule],
// [VersionGetResponseRulesSetCacheSettingsRule],
// [VersionGetResponseRulesSetCacheTagsRule],
// [VersionGetResponseRulesSetConfigurationRule],
// [VersionGetResponseRulesSkipRule],
// [VersionGetResponseRulesTransformResponseHTMLRule].
func (r VersionGetResponseRule) AsUnion() VersionGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [VersionGetResponseRulesBlockRule],
// [VersionGetResponseRulesChallengeRule],
// [VersionGetResponseRulesResponseCompressionRule],
// [VersionGetResponseRulesDDoSDynamicRule], [VersionGetResponseRulesExecuteRule],
// [VersionGetResponseRulesForceConnectionCloseRule],
// [VersionGetResponseRulesJavaScriptChallengeRule],
// [VersionGetResponseRulesLogRule], [VersionGetResponseRulesLogCustomFieldRule],
// [VersionGetResponseRulesManagedChallengeRule],
// [VersionGetResponseRulesRedirectRule], [VersionGetResponseRulesRewriteRule],
// [VersionGetResponseRulesRouteRule], [VersionGetResponseRulesScoreRule],
// [VersionGetResponseRulesServeErrorRule],
// [VersionGetResponseRulesSetCacheControlRule],
// [VersionGetResponseRulesSetCacheSettingsRule],
// [VersionGetResponseRulesSetCacheTagsRule],
// [VersionGetResponseRulesSetConfigurationRule], [VersionGetResponseRulesSkipRule]
// or [VersionGetResponseRulesTransformResponseHTMLRule].
type VersionGetResponseRulesUnion interface {
	implementsVersionGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesResponseCompressionRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesDDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesJavaScriptChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesRouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesSetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesSetCacheTagsRule{}),
			DiscriminatorValue: "set_cache_tags",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesSetConfigurationRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesSkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseRulesTransformResponseHTMLRule{}),
			DiscriminatorValue: "transform_response_html",
		},
	)
}

type VersionGetResponseRulesBlockRule struct {
	ID         string                               `json:"id" api:"required"`
	Action     string                               `json:"action" api:"required"`
	Enabled    bool                                 `json:"enabled" api:"required"`
	Expression string                               `json:"expression" api:"required"`
	Ref        string                               `json:"ref" api:"required"`
	JSON       versionGetResponseRulesBlockRuleJSON `json:"-"`
	BlockRule
}

// versionGetResponseRulesBlockRuleJSON contains the JSON metadata for the struct
// [VersionGetResponseRulesBlockRule]
type versionGetResponseRulesBlockRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesBlockRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesChallengeRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesChallengeRuleAction `json:"action" api:"required"`
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
	ExposedCredentialCheck VersionGetResponseRulesChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit VersionGetResponseRulesChallengeRuleRatelimit `json:"ratelimit"`
	JSON      versionGetResponseRulesChallengeRuleJSON      `json:"-"`
}

// versionGetResponseRulesChallengeRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseRulesChallengeRule]
type versionGetResponseRulesChallengeRuleJSON struct {
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

func (r *VersionGetResponseRulesChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesChallengeRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesChallengeRuleAction string

const (
	VersionGetResponseRulesChallengeRuleActionChallenge VersionGetResponseRulesChallengeRuleAction = "challenge"
)

func (r VersionGetResponseRulesChallengeRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type VersionGetResponseRulesChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                         `json:"username_expression" api:"required"`
	JSON               versionGetResponseRulesChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// versionGetResponseRulesChallengeRuleExposedCredentialCheckJSON contains the JSON
// metadata for the struct
// [VersionGetResponseRulesChallengeRuleExposedCredentialCheck]
type versionGetResponseRulesChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionGetResponseRulesChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type VersionGetResponseRulesChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                            `json:"score_response_header_name"`
	JSON                    versionGetResponseRulesChallengeRuleRatelimitJSON `json:"-"`
}

// versionGetResponseRulesChallengeRuleRatelimitJSON contains the JSON metadata for
// the struct [VersionGetResponseRulesChallengeRuleRatelimit]
type versionGetResponseRulesChallengeRuleRatelimitJSON struct {
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

func (r *VersionGetResponseRulesChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseRulesResponseCompressionRule struct {
	ID         string                                             `json:"id" api:"required"`
	Action     string                                             `json:"action" api:"required"`
	Enabled    bool                                               `json:"enabled" api:"required"`
	Expression string                                             `json:"expression" api:"required"`
	Ref        string                                             `json:"ref" api:"required"`
	JSON       versionGetResponseRulesResponseCompressionRuleJSON `json:"-"`
	CompressResponseRule
}

// versionGetResponseRulesResponseCompressionRuleJSON contains the JSON metadata
// for the struct [VersionGetResponseRulesResponseCompressionRule]
type versionGetResponseRulesResponseCompressionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesResponseCompressionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesResponseCompressionRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesResponseCompressionRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesDDoSDynamicRule struct {
	ID         string                                     `json:"id" api:"required"`
	Action     string                                     `json:"action" api:"required"`
	Enabled    bool                                       `json:"enabled" api:"required"`
	Expression string                                     `json:"expression" api:"required"`
	Ref        string                                     `json:"ref" api:"required"`
	JSON       versionGetResponseRulesDDoSDynamicRuleJSON `json:"-"`
	DDoSDynamicRule
}

// versionGetResponseRulesDDoSDynamicRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseRulesDDoSDynamicRule]
type versionGetResponseRulesDDoSDynamicRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesDDoSDynamicRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesDDoSDynamicRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesDDoSDynamicRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesExecuteRule struct {
	ID         string                                 `json:"id" api:"required"`
	Action     string                                 `json:"action" api:"required"`
	Enabled    bool                                   `json:"enabled" api:"required"`
	Expression string                                 `json:"expression" api:"required"`
	Ref        string                                 `json:"ref" api:"required"`
	JSON       versionGetResponseRulesExecuteRuleJSON `json:"-"`
	ExecuteRule
}

// versionGetResponseRulesExecuteRuleJSON contains the JSON metadata for the struct
// [VersionGetResponseRulesExecuteRule]
type versionGetResponseRulesExecuteRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesExecuteRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesForceConnectionCloseRule struct {
	ID         string                                              `json:"id" api:"required"`
	Action     string                                              `json:"action" api:"required"`
	Enabled    bool                                                `json:"enabled" api:"required"`
	Expression string                                              `json:"expression" api:"required"`
	Ref        string                                              `json:"ref" api:"required"`
	JSON       versionGetResponseRulesForceConnectionCloseRuleJSON `json:"-"`
	ForceConnectionCloseRule
}

// versionGetResponseRulesForceConnectionCloseRuleJSON contains the JSON metadata
// for the struct [VersionGetResponseRulesForceConnectionCloseRule]
type versionGetResponseRulesForceConnectionCloseRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesForceConnectionCloseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesForceConnectionCloseRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesForceConnectionCloseRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesJavaScriptChallengeRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesJavaScriptChallengeRuleAction `json:"action" api:"required"`
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
	ExposedCredentialCheck VersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit VersionGetResponseRulesJavaScriptChallengeRuleRatelimit `json:"ratelimit"`
	JSON      versionGetResponseRulesJavaScriptChallengeRuleJSON      `json:"-"`
}

// versionGetResponseRulesJavaScriptChallengeRuleJSON contains the JSON metadata
// for the struct [VersionGetResponseRulesJavaScriptChallengeRule]
type versionGetResponseRulesJavaScriptChallengeRuleJSON struct {
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

func (r *VersionGetResponseRulesJavaScriptChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesJavaScriptChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesJavaScriptChallengeRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesJavaScriptChallengeRuleAction string

const (
	VersionGetResponseRulesJavaScriptChallengeRuleActionJSChallenge VersionGetResponseRulesJavaScriptChallengeRuleAction = "js_challenge"
)

func (r VersionGetResponseRulesJavaScriptChallengeRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesJavaScriptChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type VersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                   `json:"username_expression" api:"required"`
	JSON               versionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// versionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheck]
type versionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesJavaScriptChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type VersionGetResponseRulesJavaScriptChallengeRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                                      `json:"score_response_header_name"`
	JSON                    versionGetResponseRulesJavaScriptChallengeRuleRatelimitJSON `json:"-"`
}

// versionGetResponseRulesJavaScriptChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct
// [VersionGetResponseRulesJavaScriptChallengeRuleRatelimit]
type versionGetResponseRulesJavaScriptChallengeRuleRatelimitJSON struct {
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

func (r *VersionGetResponseRulesJavaScriptChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesJavaScriptChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseRulesLogRule struct {
	ID         string                             `json:"id" api:"required"`
	Action     string                             `json:"action" api:"required"`
	Enabled    bool                               `json:"enabled" api:"required"`
	Expression string                             `json:"expression" api:"required"`
	Ref        string                             `json:"ref" api:"required"`
	JSON       versionGetResponseRulesLogRuleJSON `json:"-"`
	LogRule
}

// versionGetResponseRulesLogRuleJSON contains the JSON metadata for the struct
// [VersionGetResponseRulesLogRule]
type versionGetResponseRulesLogRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesLogRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesLogCustomFieldRule struct {
	ID         string                                        `json:"id" api:"required"`
	Action     string                                        `json:"action" api:"required"`
	Enabled    bool                                          `json:"enabled" api:"required"`
	Expression string                                        `json:"expression" api:"required"`
	Ref        string                                        `json:"ref" api:"required"`
	JSON       versionGetResponseRulesLogCustomFieldRuleJSON `json:"-"`
	LogCustomFieldRule
}

// versionGetResponseRulesLogCustomFieldRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseRulesLogCustomFieldRule]
type versionGetResponseRulesLogCustomFieldRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesLogCustomFieldRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesManagedChallengeRule struct {
	ID         string                                          `json:"id" api:"required"`
	Action     string                                          `json:"action" api:"required"`
	Enabled    bool                                            `json:"enabled" api:"required"`
	Expression string                                          `json:"expression" api:"required"`
	Ref        string                                          `json:"ref" api:"required"`
	JSON       versionGetResponseRulesManagedChallengeRuleJSON `json:"-"`
	ManagedChallengeRule
}

// versionGetResponseRulesManagedChallengeRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseRulesManagedChallengeRule]
type versionGetResponseRulesManagedChallengeRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesManagedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesManagedChallengeRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesRedirectRule struct {
	ID         string                                  `json:"id" api:"required"`
	Action     string                                  `json:"action" api:"required"`
	Enabled    bool                                    `json:"enabled" api:"required"`
	Expression string                                  `json:"expression" api:"required"`
	Ref        string                                  `json:"ref" api:"required"`
	JSON       versionGetResponseRulesRedirectRuleJSON `json:"-"`
	RedirectRule
}

// versionGetResponseRulesRedirectRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseRulesRedirectRule]
type versionGetResponseRulesRedirectRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRedirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRedirectRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesRewriteRule struct {
	ID         string                                 `json:"id" api:"required"`
	Action     string                                 `json:"action" api:"required"`
	Enabled    bool                                   `json:"enabled" api:"required"`
	Expression string                                 `json:"expression" api:"required"`
	Ref        string                                 `json:"ref" api:"required"`
	JSON       versionGetResponseRulesRewriteRuleJSON `json:"-"`
	RewriteRule
}

// versionGetResponseRulesRewriteRuleJSON contains the JSON metadata for the struct
// [VersionGetResponseRulesRewriteRule]
type versionGetResponseRulesRewriteRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRewriteRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesRouteRule struct {
	ID         string                               `json:"id" api:"required"`
	Action     string                               `json:"action" api:"required"`
	Enabled    bool                                 `json:"enabled" api:"required"`
	Expression string                               `json:"expression" api:"required"`
	Ref        string                               `json:"ref" api:"required"`
	JSON       versionGetResponseRulesRouteRuleJSON `json:"-"`
	RouteRule
}

// versionGetResponseRulesRouteRuleJSON contains the JSON metadata for the struct
// [VersionGetResponseRulesRouteRule]
type versionGetResponseRulesRouteRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesRouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRouteRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRouteRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesScoreRule struct {
	ID         string                               `json:"id" api:"required"`
	Action     string                               `json:"action" api:"required"`
	Enabled    bool                                 `json:"enabled" api:"required"`
	Expression string                               `json:"expression" api:"required"`
	Ref        string                               `json:"ref" api:"required"`
	JSON       versionGetResponseRulesScoreRuleJSON `json:"-"`
	ScoreRule
}

// versionGetResponseRulesScoreRuleJSON contains the JSON metadata for the struct
// [VersionGetResponseRulesScoreRule]
type versionGetResponseRulesScoreRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesScoreRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesServeErrorRule struct {
	ID         string                                    `json:"id" api:"required"`
	Action     string                                    `json:"action" api:"required"`
	Enabled    bool                                      `json:"enabled" api:"required"`
	Expression string                                    `json:"expression" api:"required"`
	Ref        string                                    `json:"ref" api:"required"`
	JSON       versionGetResponseRulesServeErrorRuleJSON `json:"-"`
	ServeErrorRule
}

// versionGetResponseRulesServeErrorRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseRulesServeErrorRule]
type versionGetResponseRulesServeErrorRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesServeErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesServeErrorRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesSetCacheControlRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action  VersionGetResponseRulesSetCacheControlRuleAction `json:"action" api:"required"`
	Enabled bool                                             `json:"enabled" api:"required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression" api:"required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref" api:"required"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck VersionGetResponseRulesSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit VersionGetResponseRulesSetCacheControlRuleRatelimit `json:"ratelimit"`
	JSON      versionGetResponseRulesSetCacheControlRuleJSON      `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseRulesSetCacheControlRule]
type versionGetResponseRulesSetCacheControlRuleJSON struct {
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

func (r *VersionGetResponseRulesSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesSetCacheControlRuleAction string

const (
	VersionGetResponseRulesSetCacheControlRuleActionSetCacheControl VersionGetResponseRulesSetCacheControlRuleAction = "set_cache_control"
)

func (r VersionGetResponseRulesSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable VersionGetResponseRulesSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache VersionGetResponseRulesSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore VersionGetResponseRulesSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private VersionGetResponseRulesSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public VersionGetResponseRulesSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 versionGetResponseRulesSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersJSON contains the JSON
// metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParameters]
type versionGetResponseRulesSetCacheControlRuleActionParametersJSON struct {
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

func (r *VersionGetResponseRulesSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type VersionGetResponseRulesSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                    `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersImmutableJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersImmutable]
type versionGetResponseRulesSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersImmutable) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                   `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                `json:"value"`
	JSON  versionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge]
type versionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                             `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                         `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate]
type versionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                     `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                         `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand]
type versionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                     `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                           `json:"qualifiers"`
	JSON       versionGetResponseRulesSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersNoCacheJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoCache]
type versionGetResponseRulesSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoCache) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                          `json:"qualifiers"`
	JSON       versionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                 `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                  `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersNoStoreJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoStore]
type versionGetResponseRulesSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoStore) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                              `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                 `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                      `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform]
type versionGetResponseRulesSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                  `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                     `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type VersionGetResponseRulesSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                           `json:"qualifiers"`
	JSON       versionGetResponseRulesSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersPrivateJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPrivate]
type versionGetResponseRulesSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersPrivate) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                          `json:"qualifiers"`
	JSON       versionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                 `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                          `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate]
type versionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type VersionGetResponseRulesSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                 `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          VersionGetResponseRulesSetCacheControlRuleActionParametersPublicUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersPublicJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPublic]
type versionGetResponseRulesSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPublicUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersPublic) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersPublicUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersPublicUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersPublicUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                             `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                 `json:"value"`
	JSON  versionGetResponseRulesSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersSMaxageJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage]
type versionGetResponseRulesSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                              `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                 `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                      `json:"value"`
	JSON  versionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError]
type versionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                   `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                              `json:"value"`
	JSON  versionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion
}

// versionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate]
type versionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           versionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type versionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsVersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, VersionGetResponseRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type VersionGetResponseRulesSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                               `json:"username_expression" api:"required"`
	JSON               versionGetResponseRulesSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheControlRuleExposedCredentialCheck]
type versionGetResponseRulesSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type VersionGetResponseRulesSetCacheControlRuleRatelimit struct {
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
	JSON                    versionGetResponseRulesSetCacheControlRuleRatelimitJSON `json:"-"`
}

// versionGetResponseRulesSetCacheControlRuleRatelimitJSON contains the JSON
// metadata for the struct [VersionGetResponseRulesSetCacheControlRuleRatelimit]
type versionGetResponseRulesSetCacheControlRuleRatelimitJSON struct {
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

func (r *VersionGetResponseRulesSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseRulesSetCacheSettingsRule struct {
	ID         string                                          `json:"id" api:"required"`
	Action     string                                          `json:"action" api:"required"`
	Enabled    bool                                            `json:"enabled" api:"required"`
	Expression string                                          `json:"expression" api:"required"`
	Ref        string                                          `json:"ref" api:"required"`
	JSON       versionGetResponseRulesSetCacheSettingsRuleJSON `json:"-"`
	SetCacheSettingsRule
}

// versionGetResponseRulesSetCacheSettingsRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseRulesSetCacheSettingsRule]
type versionGetResponseRulesSetCacheSettingsRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheSettingsRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesSetCacheTagsRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action  VersionGetResponseRulesSetCacheTagsRuleAction `json:"action" api:"required"`
	Enabled bool                                          `json:"enabled" api:"required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression" api:"required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref" api:"required"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck VersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit VersionGetResponseRulesSetCacheTagsRuleRatelimit `json:"ratelimit"`
	JSON      versionGetResponseRulesSetCacheTagsRuleJSON      `json:"-"`
}

// versionGetResponseRulesSetCacheTagsRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseRulesSetCacheTagsRule]
type versionGetResponseRulesSetCacheTagsRuleJSON struct {
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

func (r *VersionGetResponseRulesSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheTagsRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesSetCacheTagsRuleAction string

const (
	VersionGetResponseRulesSetCacheTagsRuleActionSetCacheTags VersionGetResponseRulesSetCacheTagsRuleAction = "set_cache_tags"
)

func (r VersionGetResponseRulesSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation VersionGetResponseRulesSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                 `json:"values"`
	JSON   versionGetResponseRulesSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  VersionGetResponseRulesSetCacheTagsRuleActionParametersUnion
}

// versionGetResponseRulesSetCacheTagsRuleActionParametersJSON contains the JSON
// metadata for the struct
// [VersionGetResponseRulesSetCacheTagsRuleActionParameters]
type versionGetResponseRulesSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r versionGetResponseRulesSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseRulesSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseRulesSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [VersionGetResponseRulesSetCacheTagsRuleActionParametersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r VersionGetResponseRulesSetCacheTagsRuleActionParameters) AsUnion() VersionGetResponseRulesSetCacheTagsRuleActionParametersUnion {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues] or
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type VersionGetResponseRulesSetCacheTagsRuleActionParametersUnion interface {
	implementsVersionGetResponseRulesSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseRulesSetCacheTagsRuleActionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                      `json:"values" api:"required"`
	JSON   versionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// versionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues]
type versionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsVersionGetResponseRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      versionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// versionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type versionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsVersionGetResponseRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, VersionGetResponseRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                         `json:"values" api:"required"`
	JSON   versionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// versionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type versionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsVersionGetResponseRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      versionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// versionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type versionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsVersionGetResponseRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, VersionGetResponseRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                      `json:"values" api:"required"`
	JSON   versionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// versionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues]
type versionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsVersionGetResponseRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      versionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// versionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type versionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsVersionGetResponseRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, VersionGetResponseRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type VersionGetResponseRulesSetCacheTagsRuleActionParametersOperation string

const (
	VersionGetResponseRulesSetCacheTagsRuleActionParametersOperationAdd    VersionGetResponseRulesSetCacheTagsRuleActionParametersOperation = "add"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersOperationRemove VersionGetResponseRulesSetCacheTagsRuleActionParametersOperation = "remove"
	VersionGetResponseRulesSetCacheTagsRuleActionParametersOperationSet    VersionGetResponseRulesSetCacheTagsRuleActionParametersOperation = "set"
)

func (r VersionGetResponseRulesSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesSetCacheTagsRuleActionParametersOperationAdd, VersionGetResponseRulesSetCacheTagsRuleActionParametersOperationRemove, VersionGetResponseRulesSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type VersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                            `json:"username_expression" api:"required"`
	JSON               versionGetResponseRulesSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// versionGetResponseRulesSetCacheTagsRuleExposedCredentialCheckJSON contains the
// JSON metadata for the struct
// [VersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheck]
type versionGetResponseRulesSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type VersionGetResponseRulesSetCacheTagsRuleRatelimit struct {
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
	ScoreResponseHeaderName string                                               `json:"score_response_header_name"`
	JSON                    versionGetResponseRulesSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// versionGetResponseRulesSetCacheTagsRuleRatelimitJSON contains the JSON metadata
// for the struct [VersionGetResponseRulesSetCacheTagsRuleRatelimit]
type versionGetResponseRulesSetCacheTagsRuleRatelimitJSON struct {
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

func (r *VersionGetResponseRulesSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseRulesSetConfigurationRule struct {
	ID         string                                          `json:"id" api:"required"`
	Action     string                                          `json:"action" api:"required"`
	Enabled    bool                                            `json:"enabled" api:"required"`
	Expression string                                          `json:"expression" api:"required"`
	Ref        string                                          `json:"ref" api:"required"`
	JSON       versionGetResponseRulesSetConfigurationRuleJSON `json:"-"`
	SetConfigRule
}

// versionGetResponseRulesSetConfigurationRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseRulesSetConfigurationRule]
type versionGetResponseRulesSetConfigurationRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesSetConfigurationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSetConfigurationRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSetConfigurationRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesSkipRule struct {
	ID         string                              `json:"id" api:"required"`
	Action     string                              `json:"action" api:"required"`
	Enabled    bool                                `json:"enabled" api:"required"`
	Expression string                              `json:"expression" api:"required"`
	Ref        string                              `json:"ref" api:"required"`
	JSON       versionGetResponseRulesSkipRuleJSON `json:"-"`
	SkipRule
}

// versionGetResponseRulesSkipRuleJSON contains the JSON metadata for the struct
// [VersionGetResponseRulesSkipRule]
type versionGetResponseRulesSkipRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesSkipRule) implementsVersionGetResponseRule() {}

type VersionGetResponseRulesTransformResponseHTMLRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action  VersionGetResponseRulesTransformResponseHTMLRuleAction `json:"action" api:"required"`
	Enabled bool                                                   `json:"enabled" api:"required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression" api:"required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref" api:"required"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseRulesTransformResponseHTMLRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck VersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit VersionGetResponseRulesTransformResponseHTMLRuleRatelimit `json:"ratelimit"`
	JSON      versionGetResponseRulesTransformResponseHTMLRuleJSON      `json:"-"`
}

// versionGetResponseRulesTransformResponseHTMLRuleJSON contains the JSON metadata
// for the struct [VersionGetResponseRulesTransformResponseHTMLRule]
type versionGetResponseRulesTransformResponseHTMLRuleJSON struct {
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

func (r *VersionGetResponseRulesTransformResponseHTMLRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesTransformResponseHTMLRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesTransformResponseHTMLRule) implementsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesTransformResponseHTMLRuleAction string

const (
	VersionGetResponseRulesTransformResponseHTMLRuleActionTransformResponseHTML VersionGetResponseRulesTransformResponseHTMLRuleAction = "transform_response_html"
)

func (r VersionGetResponseRulesTransformResponseHTMLRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesTransformResponseHTMLRuleActionTransformResponseHTML:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseRulesTransformResponseHTMLRuleActionParameters struct {
	// Enables the link maze transformation on the response.
	LinkMaze interface{}                                                          `json:"link_maze" api:"required"`
	JSON     versionGetResponseRulesTransformResponseHTMLRuleActionParametersJSON `json:"-"`
}

// versionGetResponseRulesTransformResponseHTMLRuleActionParametersJSON contains
// the JSON metadata for the struct
// [VersionGetResponseRulesTransformResponseHTMLRuleActionParameters]
type versionGetResponseRulesTransformResponseHTMLRuleActionParametersJSON struct {
	LinkMaze    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseRulesTransformResponseHTMLRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesTransformResponseHTMLRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Configuration for exposed credential checking.
type VersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                     `json:"username_expression" api:"required"`
	JSON               versionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheckJSON `json:"-"`
}

// versionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [VersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheck]
type versionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesTransformResponseHTMLRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type VersionGetResponseRulesTransformResponseHTMLRuleRatelimit struct {
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
	JSON                    versionGetResponseRulesTransformResponseHTMLRuleRatelimitJSON `json:"-"`
}

// versionGetResponseRulesTransformResponseHTMLRuleRatelimitJSON contains the JSON
// metadata for the struct
// [VersionGetResponseRulesTransformResponseHTMLRuleRatelimit]
type versionGetResponseRulesTransformResponseHTMLRuleRatelimitJSON struct {
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

func (r *VersionGetResponseRulesTransformResponseHTMLRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesTransformResponseHTMLRuleRatelimitJSON) RawJSON() string {
	return r.raw
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
	// Validates the request without persisting changes when set to `true`. Responses
	// that normally return 200 return `result: null`; endpoints that normally return
	// 204 continue to return 204.
	DryRun param.Field[bool] `query:"dry_run"`
}

// URLQuery serializes [VersionDeleteParams]'s query parameters as `url.Values`.
func (r VersionDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
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
	Errors []VersionGetResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []VersionGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// A ruleset object.
	Result VersionGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success VersionGetResponseEnvelopeSuccess `json:"success" api:"required"`
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
	Message string `json:"message" api:"required"`
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
	Pointer string                                     `json:"pointer" api:"required"`
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
	Message string `json:"message" api:"required"`
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
	Pointer string                                       `json:"pointer" api:"required"`
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
