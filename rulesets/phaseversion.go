// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"errors"
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
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint/versions/%s", accountOrZone, accountOrZoneID, rulesetPhase, rulesetVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type PhaseVersionListResponse struct {
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
	Rules []PhaseVersionGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
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
	// The action to perform when the rule matches.
	Action PhaseVersionGetResponseRulesAction `json:"action"`
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
	Version string                          `json:"version,required"`
	JSON    phaseVersionGetResponseRuleJSON `json:"-"`
	union   PhaseVersionGetResponseRulesUnion
}

// phaseVersionGetResponseRuleJSON contains the JSON metadata for the struct
// [PhaseVersionGetResponseRule]
type phaseVersionGetResponseRuleJSON struct {
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
// Possible runtime types of the union are [rulesets.BlockRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsJSChallengeRule],
// [rulesets.LogRule], [rulesets.ManagedChallengeRule], [rulesets.RedirectRule],
// [rulesets.RewriteRule], [rulesets.RouteRule], [rulesets.ScoreRule],
// [rulesets.ServeErrorRule], [rulesets.SetConfigRule], [rulesets.SkipRule],
// [rulesets.SetCacheSettingsRule], [rulesets.LogCustomFieldRule],
// [rulesets.DDoSDynamicRule], [rulesets.ForceConnectionCloseRule].
func (r PhaseVersionGetResponseRule) AsUnion() PhaseVersionGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.PhaseVersionGetResponseRulesRulesetsJSChallengeRule],
// [rulesets.LogRule], [rulesets.ManagedChallengeRule], [rulesets.RedirectRule],
// [rulesets.RewriteRule], [rulesets.RouteRule], [rulesets.ScoreRule],
// [rulesets.ServeErrorRule], [rulesets.SetConfigRule], [rulesets.SkipRule],
// [rulesets.SetCacheSettingsRule], [rulesets.LogCustomFieldRule],
// [rulesets.DDoSDynamicRule] or [rulesets.ForceConnectionCloseRule].
type PhaseVersionGetResponseRulesUnion interface {
	implementsRulesetsPhaseVersionGetResponseRule()
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
			Type:               reflect.TypeOf(ExecuteRule{}),
			DiscriminatorValue: "execute",
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

type PhaseVersionGetResponseRulesRulesetsChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
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
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsChallengeRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsChallengeRuleJSON contains the JSON metadata
// for the struct [PhaseVersionGetResponseRulesRulesetsChallengeRule]
type phaseVersionGetResponseRulesRulesetsChallengeRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsChallengeRule) implementsRulesetsPhaseVersionGetResponseRule() {
}

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

type PhaseVersionGetResponseRulesRulesetsJSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
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
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging Logging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref  string                                                  `json:"ref"`
	JSON phaseVersionGetResponseRulesRulesetsJSChallengeRuleJSON `json:"-"`
}

// phaseVersionGetResponseRulesRulesetsJSChallengeRuleJSON contains the JSON
// metadata for the struct [PhaseVersionGetResponseRulesRulesetsJSChallengeRule]
type phaseVersionGetResponseRulesRulesetsJSChallengeRuleJSON struct {
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

func (r *PhaseVersionGetResponseRulesRulesetsJSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseVersionGetResponseRulesRulesetsJSChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseVersionGetResponseRulesRulesetsJSChallengeRule) implementsRulesetsPhaseVersionGetResponseRule() {
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

// The action to perform when the rule matches.
type PhaseVersionGetResponseRulesAction string

const (
	PhaseVersionGetResponseRulesActionBlock                PhaseVersionGetResponseRulesAction = "block"
	PhaseVersionGetResponseRulesActionChallenge            PhaseVersionGetResponseRulesAction = "challenge"
	PhaseVersionGetResponseRulesActionCompressResponse     PhaseVersionGetResponseRulesAction = "compress_response"
	PhaseVersionGetResponseRulesActionExecute              PhaseVersionGetResponseRulesAction = "execute"
	PhaseVersionGetResponseRulesActionJSChallenge          PhaseVersionGetResponseRulesAction = "js_challenge"
	PhaseVersionGetResponseRulesActionLog                  PhaseVersionGetResponseRulesAction = "log"
	PhaseVersionGetResponseRulesActionManagedChallenge     PhaseVersionGetResponseRulesAction = "managed_challenge"
	PhaseVersionGetResponseRulesActionRedirect             PhaseVersionGetResponseRulesAction = "redirect"
	PhaseVersionGetResponseRulesActionRewrite              PhaseVersionGetResponseRulesAction = "rewrite"
	PhaseVersionGetResponseRulesActionRoute                PhaseVersionGetResponseRulesAction = "route"
	PhaseVersionGetResponseRulesActionScore                PhaseVersionGetResponseRulesAction = "score"
	PhaseVersionGetResponseRulesActionServeError           PhaseVersionGetResponseRulesAction = "serve_error"
	PhaseVersionGetResponseRulesActionSetConfig            PhaseVersionGetResponseRulesAction = "set_config"
	PhaseVersionGetResponseRulesActionSkip                 PhaseVersionGetResponseRulesAction = "skip"
	PhaseVersionGetResponseRulesActionSetCacheSettings     PhaseVersionGetResponseRulesAction = "set_cache_settings"
	PhaseVersionGetResponseRulesActionLogCustomField       PhaseVersionGetResponseRulesAction = "log_custom_field"
	PhaseVersionGetResponseRulesActionDDoSDynamic          PhaseVersionGetResponseRulesAction = "ddos_dynamic"
	PhaseVersionGetResponseRulesActionForceConnectionClose PhaseVersionGetResponseRulesAction = "force_connection_close"
)

func (r PhaseVersionGetResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseVersionGetResponseRulesActionBlock, PhaseVersionGetResponseRulesActionChallenge, PhaseVersionGetResponseRulesActionCompressResponse, PhaseVersionGetResponseRulesActionExecute, PhaseVersionGetResponseRulesActionJSChallenge, PhaseVersionGetResponseRulesActionLog, PhaseVersionGetResponseRulesActionManagedChallenge, PhaseVersionGetResponseRulesActionRedirect, PhaseVersionGetResponseRulesActionRewrite, PhaseVersionGetResponseRulesActionRoute, PhaseVersionGetResponseRulesActionScore, PhaseVersionGetResponseRulesActionServeError, PhaseVersionGetResponseRulesActionSetConfig, PhaseVersionGetResponseRulesActionSkip, PhaseVersionGetResponseRulesActionSetCacheSettings, PhaseVersionGetResponseRulesActionLogCustomField, PhaseVersionGetResponseRulesActionDDoSDynamic, PhaseVersionGetResponseRulesActionForceConnectionClose:
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
	Errors []PhaseVersionGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseVersionGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result PhaseVersionGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success PhaseVersionGetResponseEnvelopeSuccess `json:"success,required"`
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
	Message string `json:"message,required"`
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
	Pointer string                                          `json:"pointer,required"`
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
	Message string `json:"message,required"`
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
	Pointer string                                            `json:"pointer,required"`
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
