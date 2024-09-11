// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/pagination"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *VersionService) List(ctx context.Context, rulesetID string, query VersionListParams, opts ...option.RequestOption) (res *pagination.SinglePage[VersionListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
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
		return
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
func (r *VersionService) Delete(ctx context.Context, rulesetID string, rulesetVersion string, body VersionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/versions/%s", accountOrZone, accountOrZoneID, rulesetID, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Fetches a specific version of an account or zone ruleset.
func (r *VersionService) Get(ctx context.Context, rulesetID string, rulesetVersion string, query VersionGetParams, opts ...option.RequestOption) (res *VersionGetResponse, err error) {
	var env VersionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
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
		return
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return
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
type VersionListResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
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
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind Kind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase,required"`
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

type VersionGetResponseRule struct {
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [RedirectRuleActionParameters],
	// [RewriteRuleActionParameters], [RouteRuleActionParameters],
	// [ScoreRuleActionParameters], [ServeErrorRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters], [LogCustomFieldRuleActionParameters].
	ActionParameters interface{} `json:"action_parameters,required"`
	// This field can have the runtime type of [[]string].
	Categories interface{} `json:"categories,required"`
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
// Possible runtime types of the union are [rulesets.BlockRule],
// [rulesets.VersionGetResponseRulesRulesetsChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.VersionGetResponseRulesRulesetsJSChallengeRule], [rulesets.LogRule],
// [rulesets.ManagedChallengeRule], [rulesets.RedirectRule],
// [rulesets.RewriteRule], [rulesets.RouteRule], [rulesets.ScoreRule],
// [rulesets.ServeErrorRule], [rulesets.SetConfigRule], [rulesets.SkipRule],
// [rulesets.SetCacheSettingsRule], [rulesets.LogCustomFieldRule],
// [rulesets.DDoSDynamicRule], [rulesets.ForceConnectionCloseRule].
func (r VersionGetResponseRule) AsUnion() VersionGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule],
// [rulesets.VersionGetResponseRulesRulesetsChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.VersionGetResponseRulesRulesetsJSChallengeRule], [rulesets.LogRule],
// [rulesets.ManagedChallengeRule], [rulesets.RedirectRule],
// [rulesets.RewriteRule], [rulesets.RouteRule], [rulesets.ScoreRule],
// [rulesets.ServeErrorRule], [rulesets.SetConfigRule], [rulesets.SkipRule],
// [rulesets.SetCacheSettingsRule], [rulesets.LogCustomFieldRule],
// [rulesets.DDoSDynamicRule] or [rulesets.ForceConnectionCloseRule].
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
			Type:               reflect.TypeOf(VersionGetResponseRulesRulesetsJSChallengeRule{}),
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
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(DDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(ForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
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

type VersionGetResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionGetResponseRulesRulesetsJSChallengeRuleAction `json:"action"`
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
	JSON versionGetResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// versionGetResponseRulesRulesetsJSChallengeRuleJSON contains the JSON metadata
// for the struct [VersionGetResponseRulesRulesetsJSChallengeRule]
type versionGetResponseRulesRulesetsJSChallengeRuleJSON struct {
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

func (r *VersionGetResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionGetResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionGetResponseRulesRulesetsJSChallengeRule) implementsRulesetsVersionGetResponseRule() {}

// The action to perform when the rule matches.
type VersionGetResponseRulesRulesetsJSChallengeRuleAction string

const (
	VersionGetResponseRulesRulesetsJSChallengeRuleActionJSChallenge VersionGetResponseRulesRulesetsJSChallengeRuleAction = "js_challenge"
)

func (r VersionGetResponseRulesRulesetsJSChallengeRuleAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesRulesetsJSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type VersionGetResponseRulesAction string

const (
	VersionGetResponseRulesActionBlock                VersionGetResponseRulesAction = "block"
	VersionGetResponseRulesActionChallenge            VersionGetResponseRulesAction = "challenge"
	VersionGetResponseRulesActionCompressResponse     VersionGetResponseRulesAction = "compress_response"
	VersionGetResponseRulesActionExecute              VersionGetResponseRulesAction = "execute"
	VersionGetResponseRulesActionJSChallenge          VersionGetResponseRulesAction = "js_challenge"
	VersionGetResponseRulesActionLog                  VersionGetResponseRulesAction = "log"
	VersionGetResponseRulesActionManagedChallenge     VersionGetResponseRulesAction = "managed_challenge"
	VersionGetResponseRulesActionRedirect             VersionGetResponseRulesAction = "redirect"
	VersionGetResponseRulesActionRewrite              VersionGetResponseRulesAction = "rewrite"
	VersionGetResponseRulesActionRoute                VersionGetResponseRulesAction = "route"
	VersionGetResponseRulesActionScore                VersionGetResponseRulesAction = "score"
	VersionGetResponseRulesActionServeError           VersionGetResponseRulesAction = "serve_error"
	VersionGetResponseRulesActionSetConfig            VersionGetResponseRulesAction = "set_config"
	VersionGetResponseRulesActionSkip                 VersionGetResponseRulesAction = "skip"
	VersionGetResponseRulesActionSetCacheSettings     VersionGetResponseRulesAction = "set_cache_settings"
	VersionGetResponseRulesActionLogCustomField       VersionGetResponseRulesAction = "log_custom_field"
	VersionGetResponseRulesActionDDoSDynamic          VersionGetResponseRulesAction = "ddos_dynamic"
	VersionGetResponseRulesActionForceConnectionClose VersionGetResponseRulesAction = "force_connection_close"
)

func (r VersionGetResponseRulesAction) IsKnown() bool {
	switch r {
	case VersionGetResponseRulesActionBlock, VersionGetResponseRulesActionChallenge, VersionGetResponseRulesActionCompressResponse, VersionGetResponseRulesActionExecute, VersionGetResponseRulesActionJSChallenge, VersionGetResponseRulesActionLog, VersionGetResponseRulesActionManagedChallenge, VersionGetResponseRulesActionRedirect, VersionGetResponseRulesActionRewrite, VersionGetResponseRulesActionRoute, VersionGetResponseRulesActionScore, VersionGetResponseRulesActionServeError, VersionGetResponseRulesActionSetConfig, VersionGetResponseRulesActionSkip, VersionGetResponseRulesActionSetCacheSettings, VersionGetResponseRulesActionLogCustomField, VersionGetResponseRulesActionDDoSDynamic, VersionGetResponseRulesActionForceConnectionClose:
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