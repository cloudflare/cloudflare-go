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
	"github.com/cloudflare/cloudflare-go/v7/shared"
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
func (r *VersionService) List(ctx context.Context, rulesetID string, query VersionListParams, opts ...option.RequestOption) (res *VersionListResponseUnion, err error) {
	var env VersionListResponseEnvelope
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
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
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

// A result.
//
// Union satisfied by [VersionListResponseArray], [shared.UnionString],
// [shared.UnionFloat] or [shared.UnionBool].
type VersionListResponseUnion interface {
	ImplementsVersionListResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionListResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionListResponseArray{}),
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

type VersionListResponseArray []interface{}

func (r VersionListResponseArray) ImplementsVersionListResponseUnion() {}

// A result.
type VersionGetResponse struct {
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
	// This field can have the runtime type of [[]VersionGetResponseObjectRule].
	Rules interface{} `json:"rules"`
	// The version of the ruleset.
	Version string                 `json:"version"`
	JSON    versionGetResponseJSON `json:"-"`
	union   VersionGetResponseUnion
}

// versionGetResponseJSON contains the JSON metadata for the struct
// [VersionGetResponse]
type versionGetResponseJSON struct {
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

func (r versionGetResponseJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponse{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [VersionGetResponseUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [VersionGetResponseObject],
// [VersionGetResponseObject], [VersionGetResponseObject],
// [VersionGetResponseObject], [VersionGetResponseObject].
func (r VersionGetResponse) AsUnion() VersionGetResponseUnion {
	return r.union
}

// A result.
//
// Union satisfied by [VersionGetResponseObject], [VersionGetResponseObject],
// [VersionGetResponseObject], [VersionGetResponseObject] or
// [VersionGetResponseObject].
type VersionGetResponseUnion interface {
	implementsVersionGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObject{}),
		},
	)
}

type VersionGetResponseObject struct {
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
	Rules []VersionGetResponseObjectRule `json:"rules"`
	// The version of the ruleset.
	Version string                       `json:"version"`
	JSON    versionGetResponseObjectJSON `json:"-"`
}

// versionGetResponseObjectJSON contains the JSON metadata for the struct
// [VersionGetResponseObject]
type versionGetResponseObjectJSON struct {
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

func (r *VersionGetResponseObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObject) implementsVersionGetResponse() {}

type VersionGetResponseObjectRule struct {
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
	// [VersionGetResponseObjectRulesSetCacheControlRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [VersionGetResponseObjectRulesTransformResponseHTMLRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// This field can have the runtime type of [BlockRuleExposedCredentialCheck],
	// [VersionGetResponseObjectRulesChallengeRuleExposedCredentialCheck],
	// [CompressResponseRuleExposedCredentialCheck],
	// [DDoSDynamicRuleExposedCredentialCheck], [ExecuteRuleExposedCredentialCheck],
	// [ForceConnectionCloseRuleExposedCredentialCheck],
	// [VersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheck],
	// [LogRuleExposedCredentialCheck], [LogCustomFieldRuleExposedCredentialCheck],
	// [ManagedChallengeRuleExposedCredentialCheck],
	// [RedirectRuleExposedCredentialCheck], [RewriteRuleExposedCredentialCheck],
	// [RouteRuleExposedCredentialCheck], [ScoreRuleExposedCredentialCheck],
	// [ServeErrorRuleExposedCredentialCheck],
	// [VersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheck],
	// [SetCacheSettingsRuleExposedCredentialCheck],
	// [VersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheck],
	// [SetConfigRuleExposedCredentialCheck], [SkipRuleExposedCredentialCheck],
	// [VersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheck].
	ExposedCredentialCheck interface{} `json:"exposed_credential_check"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" format:"date-time"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// This field can have the runtime type of [BlockRuleRatelimit],
	// [VersionGetResponseObjectRulesChallengeRuleRatelimit],
	// [CompressResponseRuleRatelimit], [DDoSDynamicRuleRatelimit],
	// [ExecuteRuleRatelimit], [ForceConnectionCloseRuleRatelimit],
	// [VersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimit],
	// [LogRuleRatelimit], [LogCustomFieldRuleRatelimit],
	// [ManagedChallengeRuleRatelimit], [RedirectRuleRatelimit],
	// [RewriteRuleRatelimit], [RouteRuleRatelimit], [ScoreRuleRatelimit],
	// [ServeErrorRuleRatelimit],
	// [VersionGetResponseObjectRulesSetCacheControlRuleRatelimit],
	// [SetCacheSettingsRuleRatelimit],
	// [VersionGetResponseObjectRulesSetCacheTagsRuleRatelimit],
	// [SetConfigRuleRatelimit], [SkipRuleRatelimit],
	// [VersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimit].
	Ratelimit interface{} `json:"ratelimit"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                           `json:"version"`
	JSON    versionGetResponseObjectRuleJSON `json:"-"`
	union   VersionGetResponseObjectRulesUnion
}

// versionGetResponseObjectRuleJSON contains the JSON metadata for the struct
// [VersionGetResponseObjectRule]
type versionGetResponseObjectRuleJSON struct {
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

func (r versionGetResponseObjectRuleJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRule) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRule{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [VersionGetResponseObjectRulesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesBlockRule],
// [VersionGetResponseObjectRulesChallengeRule],
// [VersionGetResponseObjectRulesResponseCompressionRule],
// [VersionGetResponseObjectRulesDDoSDynamicRule],
// [VersionGetResponseObjectRulesExecuteRule],
// [VersionGetResponseObjectRulesForceConnectionCloseRule],
// [VersionGetResponseObjectRulesJavaScriptChallengeRule],
// [VersionGetResponseObjectRulesLogRule],
// [VersionGetResponseObjectRulesLogCustomFieldRule],
// [VersionGetResponseObjectRulesManagedChallengeRule],
// [VersionGetResponseObjectRulesRedirectRule],
// [VersionGetResponseObjectRulesRewriteRule],
// [VersionGetResponseObjectRulesRouteRule],
// [VersionGetResponseObjectRulesScoreRule],
// [VersionGetResponseObjectRulesServeErrorRule],
// [VersionGetResponseObjectRulesSetCacheControlRule],
// [VersionGetResponseObjectRulesSetCacheSettingsRule],
// [VersionGetResponseObjectRulesSetCacheTagsRule],
// [VersionGetResponseObjectRulesSetConfigurationRule],
// [VersionGetResponseObjectRulesSkipRule],
// [VersionGetResponseObjectRulesTransformResponseHTMLRule].
func (r VersionGetResponseObjectRule) AsUnion() VersionGetResponseObjectRulesUnion {
	return r.union
}

// Union satisfied by [VersionGetResponseObjectRulesBlockRule],
// [VersionGetResponseObjectRulesChallengeRule],
// [VersionGetResponseObjectRulesResponseCompressionRule],
// [VersionGetResponseObjectRulesDDoSDynamicRule],
// [VersionGetResponseObjectRulesExecuteRule],
// [VersionGetResponseObjectRulesForceConnectionCloseRule],
// [VersionGetResponseObjectRulesJavaScriptChallengeRule],
// [VersionGetResponseObjectRulesLogRule],
// [VersionGetResponseObjectRulesLogCustomFieldRule],
// [VersionGetResponseObjectRulesManagedChallengeRule],
// [VersionGetResponseObjectRulesRedirectRule],
// [VersionGetResponseObjectRulesRewriteRule],
// [VersionGetResponseObjectRulesRouteRule],
// [VersionGetResponseObjectRulesScoreRule],
// [VersionGetResponseObjectRulesServeErrorRule],
// [VersionGetResponseObjectRulesSetCacheControlRule],
// [VersionGetResponseObjectRulesSetCacheSettingsRule],
// [VersionGetResponseObjectRulesSetCacheTagsRule],
// [VersionGetResponseObjectRulesSetConfigurationRule],
// [VersionGetResponseObjectRulesSkipRule] or
// [VersionGetResponseObjectRulesTransformResponseHTMLRule].
type VersionGetResponseObjectRulesUnion interface {
	implementsVersionGetResponseObjectRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesUnion)(nil)).Elem(),
		"action",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesBlockRule{}),
			DiscriminatorValue: "block",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesChallengeRule{}),
			DiscriminatorValue: "challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesResponseCompressionRule{}),
			DiscriminatorValue: "compress_response",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesDDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesExecuteRule{}),
			DiscriminatorValue: "execute",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesJavaScriptChallengeRule{}),
			DiscriminatorValue: "js_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesLogRule{}),
			DiscriminatorValue: "log",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesManagedChallengeRule{}),
			DiscriminatorValue: "managed_challenge",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesRedirectRule{}),
			DiscriminatorValue: "redirect",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesRewriteRule{}),
			DiscriminatorValue: "rewrite",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesRouteRule{}),
			DiscriminatorValue: "route",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesScoreRule{}),
			DiscriminatorValue: "score",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesServeErrorRule{}),
			DiscriminatorValue: "serve_error",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRule{}),
			DiscriminatorValue: "set_cache_control",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesSetCacheSettingsRule{}),
			DiscriminatorValue: "set_cache_settings",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesSetCacheTagsRule{}),
			DiscriminatorValue: "set_cache_tags",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesSetConfigurationRule{}),
			DiscriminatorValue: "set_config",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesSkipRule{}),
			DiscriminatorValue: "skip",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionGetResponseObjectRulesTransformResponseHTMLRule{}),
			DiscriminatorValue: "transform_response_html",
		},
	)
}

type VersionGetResponseObjectRulesBlockRule struct {
	ID         string                                     `json:"id" api:"required"`
	Action     string                                     `json:"action" api:"required"`
	Enabled    bool                                       `json:"enabled" api:"required"`
	Expression string                                     `json:"expression" api:"required"`
	Ref        string                                     `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesBlockRuleJSON `json:"-"`
	BlockRule
}

// versionGetResponseObjectRulesBlockRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseObjectRulesBlockRule]
type versionGetResponseObjectRulesBlockRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesBlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesBlockRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesBlockRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesChallengeRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action VersionGetResponseObjectRulesChallengeRuleAction `json:"action" api:"required"`
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
	ExposedCredentialCheck VersionGetResponseObjectRulesChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit VersionGetResponseObjectRulesChallengeRuleRatelimit `json:"ratelimit"`
	JSON      versionGetResponseObjectRulesChallengeRuleJSON      `json:"-"`
}

// versionGetResponseObjectRulesChallengeRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseObjectRulesChallengeRule]
type versionGetResponseObjectRulesChallengeRuleJSON struct {
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

func (r *VersionGetResponseObjectRulesChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesChallengeRule) implementsVersionGetResponseObjectRule() {}

// The action to perform when the rule matches.
type VersionGetResponseObjectRulesChallengeRuleAction string

const (
	VersionGetResponseObjectRulesChallengeRuleActionChallenge VersionGetResponseObjectRulesChallengeRuleAction = "challenge"
)

func (r VersionGetResponseObjectRulesChallengeRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesChallengeRuleActionChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type VersionGetResponseObjectRulesChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                               `json:"username_expression" api:"required"`
	JSON               versionGetResponseObjectRulesChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// versionGetResponseObjectRulesChallengeRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [VersionGetResponseObjectRulesChallengeRuleExposedCredentialCheck]
type versionGetResponseObjectRulesChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type VersionGetResponseObjectRulesChallengeRuleRatelimit struct {
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
	JSON                    versionGetResponseObjectRulesChallengeRuleRatelimitJSON `json:"-"`
}

// versionGetResponseObjectRulesChallengeRuleRatelimitJSON contains the JSON
// metadata for the struct [VersionGetResponseObjectRulesChallengeRuleRatelimit]
type versionGetResponseObjectRulesChallengeRuleRatelimitJSON struct {
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

func (r *VersionGetResponseObjectRulesChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseObjectRulesResponseCompressionRule struct {
	ID         string                                                   `json:"id" api:"required"`
	Action     string                                                   `json:"action" api:"required"`
	Enabled    bool                                                     `json:"enabled" api:"required"`
	Expression string                                                   `json:"expression" api:"required"`
	Ref        string                                                   `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesResponseCompressionRuleJSON `json:"-"`
	CompressResponseRule
}

// versionGetResponseObjectRulesResponseCompressionRuleJSON contains the JSON
// metadata for the struct [VersionGetResponseObjectRulesResponseCompressionRule]
type versionGetResponseObjectRulesResponseCompressionRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesResponseCompressionRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesResponseCompressionRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesResponseCompressionRule) implementsVersionGetResponseObjectRule() {
}

type VersionGetResponseObjectRulesDDoSDynamicRule struct {
	ID         string                                           `json:"id" api:"required"`
	Action     string                                           `json:"action" api:"required"`
	Enabled    bool                                             `json:"enabled" api:"required"`
	Expression string                                           `json:"expression" api:"required"`
	Ref        string                                           `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesDDoSDynamicRuleJSON `json:"-"`
	DDoSDynamicRule
}

// versionGetResponseObjectRulesDDoSDynamicRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseObjectRulesDDoSDynamicRule]
type versionGetResponseObjectRulesDDoSDynamicRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesDDoSDynamicRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesDDoSDynamicRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesDDoSDynamicRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesExecuteRule struct {
	ID         string                                       `json:"id" api:"required"`
	Action     string                                       `json:"action" api:"required"`
	Enabled    bool                                         `json:"enabled" api:"required"`
	Expression string                                       `json:"expression" api:"required"`
	Ref        string                                       `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesExecuteRuleJSON `json:"-"`
	ExecuteRule
}

// versionGetResponseObjectRulesExecuteRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseObjectRulesExecuteRule]
type versionGetResponseObjectRulesExecuteRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesExecuteRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesExecuteRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesForceConnectionCloseRule struct {
	ID         string                                                    `json:"id" api:"required"`
	Action     string                                                    `json:"action" api:"required"`
	Enabled    bool                                                      `json:"enabled" api:"required"`
	Expression string                                                    `json:"expression" api:"required"`
	Ref        string                                                    `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesForceConnectionCloseRuleJSON `json:"-"`
	ForceConnectionCloseRule
}

// versionGetResponseObjectRulesForceConnectionCloseRuleJSON contains the JSON
// metadata for the struct [VersionGetResponseObjectRulesForceConnectionCloseRule]
type versionGetResponseObjectRulesForceConnectionCloseRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesForceConnectionCloseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesForceConnectionCloseRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesForceConnectionCloseRule) implementsVersionGetResponseObjectRule() {
}

type VersionGetResponseObjectRulesJavaScriptChallengeRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action VersionGetResponseObjectRulesJavaScriptChallengeRuleAction `json:"action" api:"required"`
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
	ExposedCredentialCheck VersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit VersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimit `json:"ratelimit"`
	JSON      versionGetResponseObjectRulesJavaScriptChallengeRuleJSON      `json:"-"`
}

// versionGetResponseObjectRulesJavaScriptChallengeRuleJSON contains the JSON
// metadata for the struct [VersionGetResponseObjectRulesJavaScriptChallengeRule]
type versionGetResponseObjectRulesJavaScriptChallengeRuleJSON struct {
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

func (r *VersionGetResponseObjectRulesJavaScriptChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesJavaScriptChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesJavaScriptChallengeRule) implementsVersionGetResponseObjectRule() {
}

// The action to perform when the rule matches.
type VersionGetResponseObjectRulesJavaScriptChallengeRuleAction string

const (
	VersionGetResponseObjectRulesJavaScriptChallengeRuleActionJSChallenge VersionGetResponseObjectRulesJavaScriptChallengeRuleAction = "js_challenge"
)

func (r VersionGetResponseObjectRulesJavaScriptChallengeRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesJavaScriptChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type VersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                         `json:"username_expression" api:"required"`
	JSON               versionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheckJSON `json:"-"`
}

// versionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheck]
type versionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesJavaScriptChallengeRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type VersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimit struct {
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
	JSON                    versionGetResponseObjectRulesJavaScriptChallengeRuleRatelimitJSON `json:"-"`
}

// versionGetResponseObjectRulesJavaScriptChallengeRuleRatelimitJSON contains the
// JSON metadata for the struct
// [VersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimit]
type versionGetResponseObjectRulesJavaScriptChallengeRuleRatelimitJSON struct {
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

func (r *VersionGetResponseObjectRulesJavaScriptChallengeRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesJavaScriptChallengeRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseObjectRulesLogRule struct {
	ID         string                                   `json:"id" api:"required"`
	Action     string                                   `json:"action" api:"required"`
	Enabled    bool                                     `json:"enabled" api:"required"`
	Expression string                                   `json:"expression" api:"required"`
	Ref        string                                   `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesLogRuleJSON `json:"-"`
	LogRule
}

// versionGetResponseObjectRulesLogRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseObjectRulesLogRule]
type versionGetResponseObjectRulesLogRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesLogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesLogRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesLogRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesLogCustomFieldRule struct {
	ID         string                                              `json:"id" api:"required"`
	Action     string                                              `json:"action" api:"required"`
	Enabled    bool                                                `json:"enabled" api:"required"`
	Expression string                                              `json:"expression" api:"required"`
	Ref        string                                              `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesLogCustomFieldRuleJSON `json:"-"`
	LogCustomFieldRule
}

// versionGetResponseObjectRulesLogCustomFieldRuleJSON contains the JSON metadata
// for the struct [VersionGetResponseObjectRulesLogCustomFieldRule]
type versionGetResponseObjectRulesLogCustomFieldRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesLogCustomFieldRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesManagedChallengeRule struct {
	ID         string                                                `json:"id" api:"required"`
	Action     string                                                `json:"action" api:"required"`
	Enabled    bool                                                  `json:"enabled" api:"required"`
	Expression string                                                `json:"expression" api:"required"`
	Ref        string                                                `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesManagedChallengeRuleJSON `json:"-"`
	ManagedChallengeRule
}

// versionGetResponseObjectRulesManagedChallengeRuleJSON contains the JSON metadata
// for the struct [VersionGetResponseObjectRulesManagedChallengeRule]
type versionGetResponseObjectRulesManagedChallengeRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesManagedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesManagedChallengeRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesRedirectRule struct {
	ID         string                                        `json:"id" api:"required"`
	Action     string                                        `json:"action" api:"required"`
	Enabled    bool                                          `json:"enabled" api:"required"`
	Expression string                                        `json:"expression" api:"required"`
	Ref        string                                        `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesRedirectRuleJSON `json:"-"`
	RedirectRule
}

// versionGetResponseObjectRulesRedirectRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseObjectRulesRedirectRule]
type versionGetResponseObjectRulesRedirectRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesRedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesRedirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesRedirectRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesRewriteRule struct {
	ID         string                                       `json:"id" api:"required"`
	Action     string                                       `json:"action" api:"required"`
	Enabled    bool                                         `json:"enabled" api:"required"`
	Expression string                                       `json:"expression" api:"required"`
	Ref        string                                       `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesRewriteRuleJSON `json:"-"`
	RewriteRule
}

// versionGetResponseObjectRulesRewriteRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseObjectRulesRewriteRule]
type versionGetResponseObjectRulesRewriteRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesRewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesRewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesRewriteRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesRouteRule struct {
	ID         string                                     `json:"id" api:"required"`
	Action     string                                     `json:"action" api:"required"`
	Enabled    bool                                       `json:"enabled" api:"required"`
	Expression string                                     `json:"expression" api:"required"`
	Ref        string                                     `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesRouteRuleJSON `json:"-"`
	RouteRule
}

// versionGetResponseObjectRulesRouteRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseObjectRulesRouteRule]
type versionGetResponseObjectRulesRouteRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesRouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesRouteRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesRouteRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesScoreRule struct {
	ID         string                                     `json:"id" api:"required"`
	Action     string                                     `json:"action" api:"required"`
	Enabled    bool                                       `json:"enabled" api:"required"`
	Expression string                                     `json:"expression" api:"required"`
	Ref        string                                     `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesScoreRuleJSON `json:"-"`
	ScoreRule
}

// versionGetResponseObjectRulesScoreRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseObjectRulesScoreRule]
type versionGetResponseObjectRulesScoreRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesScoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesScoreRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesServeErrorRule struct {
	ID         string                                          `json:"id" api:"required"`
	Action     string                                          `json:"action" api:"required"`
	Enabled    bool                                            `json:"enabled" api:"required"`
	Expression string                                          `json:"expression" api:"required"`
	Ref        string                                          `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesServeErrorRuleJSON `json:"-"`
	ServeErrorRule
}

// versionGetResponseObjectRulesServeErrorRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseObjectRulesServeErrorRule]
type versionGetResponseObjectRulesServeErrorRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesServeErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesServeErrorRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesSetCacheControlRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action  VersionGetResponseObjectRulesSetCacheControlRuleAction `json:"action" api:"required"`
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
	ActionParameters VersionGetResponseObjectRulesSetCacheControlRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck VersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit VersionGetResponseObjectRulesSetCacheControlRuleRatelimit `json:"ratelimit"`
	JSON      versionGetResponseObjectRulesSetCacheControlRuleJSON      `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleJSON contains the JSON metadata
// for the struct [VersionGetResponseObjectRulesSetCacheControlRule]
type versionGetResponseObjectRulesSetCacheControlRuleJSON struct {
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

func (r *VersionGetResponseObjectRulesSetCacheControlRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRule) implementsVersionGetResponseObjectRule() {}

// The action to perform when the rule matches.
type VersionGetResponseObjectRulesSetCacheControlRuleAction string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionSetCacheControl VersionGetResponseObjectRulesSetCacheControlRuleAction = "set_cache_control"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionSetCacheControl:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParameters struct {
	// A cache-control directive configuration.
	Immutable VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable `json:"immutable"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	MaxAge VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge `json:"max-age"`
	// A cache-control directive configuration.
	MustRevalidate VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate `json:"must-revalidate"`
	// A cache-control directive configuration.
	MustUnderstand VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand `json:"must-understand"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	NoCache VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache `json:"no-cache"`
	// A cache-control directive configuration.
	NoStore VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore `json:"no-store"`
	// A cache-control directive configuration.
	NoTransform VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform `json:"no-transform"`
	// A cache-control directive configuration that accepts optional qualifiers (header
	// names).
	Private VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate `json:"private"`
	// A cache-control directive configuration.
	ProxyRevalidate VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate `json:"proxy-revalidate"`
	// A cache-control directive configuration.
	Public VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic `json:"public"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	SMaxage VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage `json:"s-maxage"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleIfError VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError `json:"stale-if-error"`
	// A cache-control directive configuration that accepts a duration value in
	// seconds.
	StaleWhileRevalidate VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate `json:"stale-while-revalidate"`
	JSON                 versionGetResponseObjectRulesSetCacheControlRuleActionParametersJSON                 `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersJSON contains
// the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParameters]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersJSON struct {
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

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A cache-control directive configuration.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                          `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableJSON `json:"-"`
	union          VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutable() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersImmutableOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                      `json:"value"`
	JSON  versionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeJSON `json:"-"`
	union VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                   `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAge() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMaxAgeOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                               `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateJSON `json:"-"`
	union          VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                               `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandJSON `json:"-"`
	union          VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                              `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstand() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersMustUnderstandOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                 `json:"qualifiers"`
	JSON       versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheJSON `json:"-"`
	union      VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                `json:"qualifiers"`
	JSON       versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCache() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoCacheOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                        `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreJSON `json:"-"`
	union          VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                    `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStore() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoStoreOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                            `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformJSON `json:"-"`
	union          VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                        `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                           `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransform() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersNoTransformOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// This field can have the runtime type of [[]string].
	Qualifiers interface{}                                                                 `json:"qualifiers"`
	JSON       versionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateJSON `json:"-"`
	union      VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateUnion {
	return r.union
}

// A cache-control directive configuration that accepts optional qualifiers (header
// names).
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective{}),
		},
	)
}

// Set the directive with optional qualifiers.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// Optional list of header names to qualify the directive (e.g., for "private" or
	// "no-cache" directives).
	Qualifiers []string                                                                                `json:"qualifiers"`
	JSON       versionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Qualifiers     apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPrivateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateJSON `json:"-"`
	union          VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                               `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersProxyRevalidateOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                       `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicJSON `json:"-"`
	union          VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicUnion {
	return r.union
}

// A cache-control directive configuration.
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective{}),
		},
	)
}

// Set the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                   `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                      `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublic() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersPublicOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                       `json:"value"`
	JSON  versionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageJSON `json:"-"`
	union VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                    `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                       `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxage() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersSMaxageOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                            `json:"value"`
	JSON  versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorJSON `json:"-"`
	union VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                         `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                            `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfError() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleIfErrorOperationRemove:
		return true
	}
	return false
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool `json:"cloudflare_only"`
	// The duration value in seconds for the directive.
	Value int64                                                                                    `json:"value"`
	JSON  versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON `json:"-"`
	union VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	Value          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective],
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate) AsUnion() VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion {
	return r.union
}

// A cache-control directive configuration that accepts a duration value in
// seconds.
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
// or
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective].
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective{}),
		},
	)
}

// Set the directive with a duration value in seconds.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation `json:"operation" api:"required"`
	// The duration value in seconds for the directive.
	Value int64 `json:"value" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                 `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON struct {
	Operation      apijson.Field
	Value          apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateSetDirectiveOperationRemove:
		return true
	}
	return false
}

// Remove the directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective struct {
	// The operation to perform on the cache-control directive.
	Operation VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation `json:"operation" api:"required"`
	// Whether the directive should only be applied to the Cloudflare CDN cache.
	CloudflareOnly bool                                                                                                    `json:"cloudflare_only"`
	JSON           versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective]
type versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON struct {
	Operation      apijson.Field
	CloudflareOnly apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirective) implementsVersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidate() {
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateRemoveDirectiveOperationRemove:
		return true
	}
	return false
}

// The operation to perform on the cache-control directive.
type VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation string

const (
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet    VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "set"
	VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation = "remove"
)

func (r VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationSet, VersionGetResponseObjectRulesSetCacheControlRuleActionParametersStaleWhileRevalidateOperationRemove:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type VersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                     `json:"username_expression" api:"required"`
	JSON               versionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheckJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheck]
type versionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type VersionGetResponseObjectRulesSetCacheControlRuleRatelimit struct {
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
	JSON                    versionGetResponseObjectRulesSetCacheControlRuleRatelimitJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheControlRuleRatelimitJSON contains the JSON
// metadata for the struct
// [VersionGetResponseObjectRulesSetCacheControlRuleRatelimit]
type versionGetResponseObjectRulesSetCacheControlRuleRatelimitJSON struct {
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

func (r *VersionGetResponseObjectRulesSetCacheControlRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheControlRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseObjectRulesSetCacheSettingsRule struct {
	ID         string                                                `json:"id" api:"required"`
	Action     string                                                `json:"action" api:"required"`
	Enabled    bool                                                  `json:"enabled" api:"required"`
	Expression string                                                `json:"expression" api:"required"`
	Ref        string                                                `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesSetCacheSettingsRuleJSON `json:"-"`
	SetCacheSettingsRule
}

// versionGetResponseObjectRulesSetCacheSettingsRuleJSON contains the JSON metadata
// for the struct [VersionGetResponseObjectRulesSetCacheSettingsRule]
type versionGetResponseObjectRulesSetCacheSettingsRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheSettingsRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesSetCacheTagsRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action  VersionGetResponseObjectRulesSetCacheTagsRuleAction `json:"action" api:"required"`
	Enabled bool                                                `json:"enabled" api:"required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression" api:"required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref" api:"required"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseObjectRulesSetCacheTagsRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck VersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit VersionGetResponseObjectRulesSetCacheTagsRuleRatelimit `json:"ratelimit"`
	JSON      versionGetResponseObjectRulesSetCacheTagsRuleJSON      `json:"-"`
}

// versionGetResponseObjectRulesSetCacheTagsRuleJSON contains the JSON metadata for
// the struct [VersionGetResponseObjectRulesSetCacheTagsRule]
type versionGetResponseObjectRulesSetCacheTagsRuleJSON struct {
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

func (r *VersionGetResponseObjectRulesSetCacheTagsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheTagsRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheTagsRule) implementsVersionGetResponseObjectRule() {}

// The action to perform when the rule matches.
type VersionGetResponseObjectRulesSetCacheTagsRuleAction string

const (
	VersionGetResponseObjectRulesSetCacheTagsRuleActionSetCacheTags VersionGetResponseObjectRulesSetCacheTagsRuleAction = "set_cache_tags"
)

func (r VersionGetResponseObjectRulesSetCacheTagsRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheTagsRuleActionSetCacheTags:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParameters struct {
	// The operation to perform on the cache tags.
	Operation VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperation `json:"operation" api:"required"`
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression"`
	// This field can have the runtime type of [[]string].
	Values interface{}                                                       `json:"values"`
	JSON   versionGetResponseObjectRulesSetCacheTagsRuleActionParametersJSON `json:"-"`
	union  VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersUnion
}

// versionGetResponseObjectRulesSetCacheTagsRuleActionParametersJSON contains the
// JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParameters]
type versionGetResponseObjectRulesSetCacheTagsRuleActionParametersJSON struct {
	Operation   apijson.Field
	Expression  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r versionGetResponseObjectRulesSetCacheTagsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

func (r *VersionGetResponseObjectRulesSetCacheTagsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	*r = VersionGetResponseObjectRulesSetCacheTagsRuleActionParameters{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues],
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression].
func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParameters) AsUnion() VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersUnion {
	return r.union
}

// The parameters configuring the rule's action.
//
// Union satisfied by
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues],
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression],
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues],
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression],
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues]
// or
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression].
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersUnion interface {
	implementsVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression{}),
		},
	)
}

// Add cache tags using a list of values.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                            `json:"values" api:"required"`
	JSON   versionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues]
type versionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValues) implementsVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation string

const (
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "add"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "remove"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation = "set"
)

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationAdd, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationRemove, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Add cache tags using an expression.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      versionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON      `json:"-"`
}

// versionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression]
type versionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpression) implementsVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation string

const (
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "add"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "remove"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation = "set"
)

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationAdd, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationRemove, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersAddCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Remove cache tags using a list of values.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                               `json:"values" api:"required"`
	JSON   versionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues]
type versionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValues) implementsVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation string

const (
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "add"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "remove"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation = "set"
)

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationAdd, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationRemove, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Remove cache tags using an expression.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      versionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON      `json:"-"`
}

// versionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression]
type versionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpression) implementsVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation string

const (
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "add"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "remove"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation = "set"
)

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationAdd, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationRemove, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersRemoveCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// Set cache tags using a list of values.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues struct {
	// The operation to perform on the cache tags.
	Operation VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation `json:"operation" api:"required"`
	// A list of cache tag values.
	Values []string                                                                            `json:"values" api:"required"`
	JSON   versionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues]
type versionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON struct {
	Operation   apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValues) implementsVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation string

const (
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "add"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "remove"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation = "set"
)

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationAdd, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationRemove, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsValuesOperationSet:
		return true
	}
	return false
}

// Set cache tags using an expression.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression struct {
	// An expression that evaluates to an array of cache tag values.
	Expression string `json:"expression" api:"required"`
	// The operation to perform on the cache tags.
	Operation VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation `json:"operation" api:"required"`
	JSON      versionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON      `json:"-"`
}

// versionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression]
type versionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpression) implementsVersionGetResponseObjectRulesSetCacheTagsRuleActionParameters() {
}

// The operation to perform on the cache tags.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation string

const (
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "add"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "remove"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation = "set"
)

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationAdd, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationRemove, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersSetCacheTagsExpressionOperationSet:
		return true
	}
	return false
}

// The operation to perform on the cache tags.
type VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperation string

const (
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperationAdd    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperation = "add"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperationRemove VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperation = "remove"
	VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperationSet    VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperation = "set"
)

func (r VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperation) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperationAdd, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperationRemove, VersionGetResponseObjectRulesSetCacheTagsRuleActionParametersOperationSet:
		return true
	}
	return false
}

// Configuration for exposed credential checking.
type VersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                  `json:"username_expression" api:"required"`
	JSON               versionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheckJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheckJSON contains
// the JSON metadata for the struct
// [VersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheck]
type versionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheTagsRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type VersionGetResponseObjectRulesSetCacheTagsRuleRatelimit struct {
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
	JSON                    versionGetResponseObjectRulesSetCacheTagsRuleRatelimitJSON `json:"-"`
}

// versionGetResponseObjectRulesSetCacheTagsRuleRatelimitJSON contains the JSON
// metadata for the struct [VersionGetResponseObjectRulesSetCacheTagsRuleRatelimit]
type versionGetResponseObjectRulesSetCacheTagsRuleRatelimitJSON struct {
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

func (r *VersionGetResponseObjectRulesSetCacheTagsRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetCacheTagsRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type VersionGetResponseObjectRulesSetConfigurationRule struct {
	ID         string                                                `json:"id" api:"required"`
	Action     string                                                `json:"action" api:"required"`
	Enabled    bool                                                  `json:"enabled" api:"required"`
	Expression string                                                `json:"expression" api:"required"`
	Ref        string                                                `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesSetConfigurationRuleJSON `json:"-"`
	SetConfigRule
}

// versionGetResponseObjectRulesSetConfigurationRuleJSON contains the JSON metadata
// for the struct [VersionGetResponseObjectRulesSetConfigurationRule]
type versionGetResponseObjectRulesSetConfigurationRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSetConfigurationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSetConfigurationRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSetConfigurationRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesSkipRule struct {
	ID         string                                    `json:"id" api:"required"`
	Action     string                                    `json:"action" api:"required"`
	Enabled    bool                                      `json:"enabled" api:"required"`
	Expression string                                    `json:"expression" api:"required"`
	Ref        string                                    `json:"ref" api:"required"`
	JSON       versionGetResponseObjectRulesSkipRuleJSON `json:"-"`
	SkipRule
}

// versionGetResponseObjectRulesSkipRuleJSON contains the JSON metadata for the
// struct [VersionGetResponseObjectRulesSkipRule]
type versionGetResponseObjectRulesSkipRuleJSON struct {
	ID          apijson.Field
	Action      apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Ref         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesSkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesSkipRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesSkipRule) implementsVersionGetResponseObjectRule() {}

type VersionGetResponseObjectRulesTransformResponseHTMLRule struct {
	// The unique ID of the rule.
	ID string `json:"id" api:"required"`
	// The action to perform when the rule matches.
	Action  VersionGetResponseObjectRulesTransformResponseHTMLRuleAction `json:"action" api:"required"`
	Enabled bool                                                         `json:"enabled" api:"required"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression" api:"required"`
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// The reference of the rule (the rule's ID by default).
	Ref string `json:"ref" api:"required"`
	// The version of the rule.
	Version string `json:"version" api:"required"`
	// The parameters configuring the rule's action.
	ActionParameters VersionGetResponseObjectRulesTransformResponseHTMLRuleActionParameters `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Configuration for exposed credential checking.
	ExposedCredentialCheck VersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheck `json:"exposed_credential_check"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// An object configuring the rule's rate limit behavior.
	Ratelimit VersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimit `json:"ratelimit"`
	JSON      versionGetResponseObjectRulesTransformResponseHTMLRuleJSON      `json:"-"`
}

// versionGetResponseObjectRulesTransformResponseHTMLRuleJSON contains the JSON
// metadata for the struct [VersionGetResponseObjectRulesTransformResponseHTMLRule]
type versionGetResponseObjectRulesTransformResponseHTMLRuleJSON struct {
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

func (r *VersionGetResponseObjectRulesTransformResponseHTMLRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesTransformResponseHTMLRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseObjectRulesTransformResponseHTMLRule) implementsVersionGetResponseObjectRule() {
}

// The action to perform when the rule matches.
type VersionGetResponseObjectRulesTransformResponseHTMLRuleAction string

const (
	VersionGetResponseObjectRulesTransformResponseHTMLRuleActionTransformResponseHTML VersionGetResponseObjectRulesTransformResponseHTMLRuleAction = "transform_response_html"
)

func (r VersionGetResponseObjectRulesTransformResponseHTMLRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseObjectRulesTransformResponseHTMLRuleActionTransformResponseHTML:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionGetResponseObjectRulesTransformResponseHTMLRuleActionParameters struct {
	// Enables the link maze transformation on the response.
	LinkMaze interface{}                                                                `json:"link_maze" api:"required"`
	JSON     versionGetResponseObjectRulesTransformResponseHTMLRuleActionParametersJSON `json:"-"`
}

// versionGetResponseObjectRulesTransformResponseHTMLRuleActionParametersJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesTransformResponseHTMLRuleActionParameters]
type versionGetResponseObjectRulesTransformResponseHTMLRuleActionParametersJSON struct {
	LinkMaze    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesTransformResponseHTMLRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesTransformResponseHTMLRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Configuration for exposed credential checking.
type VersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheck struct {
	// An expression that selects the password used in the credentials check.
	PasswordExpression string `json:"password_expression" api:"required"`
	// An expression that selects the user ID used in the credentials check.
	UsernameExpression string                                                                           `json:"username_expression" api:"required"`
	JSON               versionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheckJSON `json:"-"`
}

// versionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheckJSON
// contains the JSON metadata for the struct
// [VersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheck]
type versionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheckJSON struct {
	PasswordExpression apijson.Field
	UsernameExpression apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *VersionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheck) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesTransformResponseHTMLRuleExposedCredentialCheckJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's rate limit behavior.
type VersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimit struct {
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
	JSON                    versionGetResponseObjectRulesTransformResponseHTMLRuleRatelimitJSON `json:"-"`
}

// versionGetResponseObjectRulesTransformResponseHTMLRuleRatelimitJSON contains the
// JSON metadata for the struct
// [VersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimit]
type versionGetResponseObjectRulesTransformResponseHTMLRuleRatelimitJSON struct {
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

func (r *VersionGetResponseObjectRulesTransformResponseHTMLRuleRatelimit) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseObjectRulesTransformResponseHTMLRuleRatelimitJSON) RawJSON() string {
	return r.raw
}

type VersionListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type VersionListResponseEnvelope struct {
	// A list of error messages.
	Errors []VersionListResponseEnvelopeErrors `json:"errors" api:"required"`
	// A list of warning messages.
	Messages []VersionListResponseEnvelopeMessages `json:"messages" api:"required"`
	// A result.
	Result VersionListResponseUnion `json:"result" api:"required"`
	// Whether the API call was successful.
	Success VersionListResponseEnvelopeSuccess `json:"success" api:"required"`
	// Information to navigate the results.
	ResultInfo VersionListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       versionListResponseEnvelopeJSON       `json:"-"`
}

// versionListResponseEnvelopeJSON contains the JSON metadata for the struct
// [VersionListResponseEnvelope]
type versionListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionListResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionListResponseEnvelopeErrorsSource `json:"source"`
	JSON   versionListResponseEnvelopeErrorsJSON   `json:"-"`
}

// versionListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [VersionListResponseEnvelopeErrors]
type versionListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionListResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                      `json:"pointer" api:"required"`
	JSON    versionListResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// versionListResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [VersionListResponseEnvelopeErrorsSource]
type versionListResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionListResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message" api:"required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionListResponseEnvelopeMessagesSource `json:"source"`
	JSON   versionListResponseEnvelopeMessagesJSON   `json:"-"`
}

// versionListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VersionListResponseEnvelopeMessages]
type versionListResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionListResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                        `json:"pointer" api:"required"`
	JSON    versionListResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// versionListResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [VersionListResponseEnvelopeMessagesSource]
type versionListResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type VersionListResponseEnvelopeSuccess bool

const (
	VersionListResponseEnvelopeSuccessTrue VersionListResponseEnvelopeSuccess = true
)

func (r VersionListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VersionListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

// Information to navigate the results.
type VersionListResponseEnvelopeResultInfo struct {
	// The set of cursors.
	Cursors VersionListResponseEnvelopeResultInfoCursors `json:"cursors"`
	JSON    versionListResponseEnvelopeResultInfoJSON    `json:"-"`
}

// versionListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [VersionListResponseEnvelopeResultInfo]
type versionListResponseEnvelopeResultInfoJSON struct {
	Cursors     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

// The set of cursors.
type VersionListResponseEnvelopeResultInfoCursors struct {
	// The cursor to use for the next page.
	After string                                           `json:"after" api:"required"`
	JSON  versionListResponseEnvelopeResultInfoCursorsJSON `json:"-"`
}

// versionListResponseEnvelopeResultInfoCursorsJSON contains the JSON metadata for
// the struct [VersionListResponseEnvelopeResultInfoCursors]
type versionListResponseEnvelopeResultInfoCursorsJSON struct {
	After       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionListResponseEnvelopeResultInfoCursors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionListResponseEnvelopeResultInfoCursorsJSON) RawJSON() string {
	return r.raw
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
	// A result.
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
