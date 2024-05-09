// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// PhaseService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewPhaseService] method instead.
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
	opts = append(r.Options[:], opts...)
	var env PhaseUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the latest version of the account or zone entry point ruleset for a
// given phase.
func (r *PhaseService) Get(ctx context.Context, rulesetPhase Phase, query PhaseGetParams, opts ...option.RequestOption) (res *PhaseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PhaseGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/phases/%v/entrypoint", accountOrZone, accountOrZoneID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type PhaseUpdateResponse struct {
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
	Rules []PhaseUpdateResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
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
	// The action to perform when the rule matches.
	Action           PhaseUpdateResponseRulesAction `json:"action"`
	ActionParameters interface{}                    `json:"action_parameters,required"`
	Categories       interface{}                    `json:"categories,required"`
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
	Version string                      `json:"version,required"`
	JSON    phaseUpdateResponseRuleJSON `json:"-"`
	union   PhaseUpdateResponseRulesUnion
}

// phaseUpdateResponseRuleJSON contains the JSON metadata for the struct
// [PhaseUpdateResponseRule]
type phaseUpdateResponseRuleJSON struct {
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

func (r phaseUpdateResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseUpdateResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseUpdateResponseRule) AsUnion() PhaseUpdateResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JSChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule], [rulesets.SetCacheSettingsRule] or
// [rulesets.PhaseUpdateResponseRulesRulesetsLogCustomFieldRule].
type PhaseUpdateResponseRulesUnion interface {
	implementsRulesetsPhaseUpdateResponseRule()
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
			Type:               reflect.TypeOf(JSChallengeRule{}),
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
			Type:               reflect.TypeOf(PhaseUpdateResponseRulesRulesetsLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
	)
}

type PhaseUpdateResponseRulesRulesetsLogCustomFieldRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParameters `json:"action_parameters"`
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
	Ref  string                                                 `json:"ref"`
	JSON phaseUpdateResponseRulesRulesetsLogCustomFieldRuleJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsLogCustomFieldRuleJSON contains the JSON
// metadata for the struct [PhaseUpdateResponseRulesRulesetsLogCustomFieldRule]
type phaseUpdateResponseRulesRulesetsLogCustomFieldRuleJSON struct {
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

func (r *PhaseUpdateResponseRulesRulesetsLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseUpdateResponseRulesRulesetsLogCustomFieldRule) implementsRulesetsPhaseUpdateResponseRule() {
}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleAction string

const (
	PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleAction = "log_custom_field"
)

func (r PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields []PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField `json:"cookie_fields"`
	// The request fields to log.
	RequestFields []PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField `json:"request_fields"`
	// The response fields to log.
	ResponseFields []PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField `json:"response_fields"`
	JSON           phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON            `json:"-"`
}

// phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON contains
// the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParameters]
type phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON struct {
	CookieFields   apijson.Field
	RequestFields  apijson.Field
	ResponseFields apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The cookie field to log.
type PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name string                                                                            `json:"name,required"`
	JSON phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField]
type phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON) RawJSON() string {
	return r.raw
}

// The request field to log.
type PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name string                                                                             `json:"name,required"`
	JSON phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField]
type phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON) RawJSON() string {
	return r.raw
}

// The response field to log.
type PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name string                                                                              `json:"name,required"`
	JSON phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON `json:"-"`
}

// phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON
// contains the JSON metadata for the struct
// [PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField]
type phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseUpdateResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type PhaseUpdateResponseRulesAction string

const (
	PhaseUpdateResponseRulesActionBlock            PhaseUpdateResponseRulesAction = "block"
	PhaseUpdateResponseRulesActionChallenge        PhaseUpdateResponseRulesAction = "challenge"
	PhaseUpdateResponseRulesActionCompressResponse PhaseUpdateResponseRulesAction = "compress_response"
	PhaseUpdateResponseRulesActionExecute          PhaseUpdateResponseRulesAction = "execute"
	PhaseUpdateResponseRulesActionJSChallenge      PhaseUpdateResponseRulesAction = "js_challenge"
	PhaseUpdateResponseRulesActionLog              PhaseUpdateResponseRulesAction = "log"
	PhaseUpdateResponseRulesActionManagedChallenge PhaseUpdateResponseRulesAction = "managed_challenge"
	PhaseUpdateResponseRulesActionRedirect         PhaseUpdateResponseRulesAction = "redirect"
	PhaseUpdateResponseRulesActionRewrite          PhaseUpdateResponseRulesAction = "rewrite"
	PhaseUpdateResponseRulesActionRoute            PhaseUpdateResponseRulesAction = "route"
	PhaseUpdateResponseRulesActionScore            PhaseUpdateResponseRulesAction = "score"
	PhaseUpdateResponseRulesActionServeError       PhaseUpdateResponseRulesAction = "serve_error"
	PhaseUpdateResponseRulesActionSetConfig        PhaseUpdateResponseRulesAction = "set_config"
	PhaseUpdateResponseRulesActionSkip             PhaseUpdateResponseRulesAction = "skip"
	PhaseUpdateResponseRulesActionSetCacheSettings PhaseUpdateResponseRulesAction = "set_cache_settings"
	PhaseUpdateResponseRulesActionLogCustomField   PhaseUpdateResponseRulesAction = "log_custom_field"
)

func (r PhaseUpdateResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseUpdateResponseRulesActionBlock, PhaseUpdateResponseRulesActionChallenge, PhaseUpdateResponseRulesActionCompressResponse, PhaseUpdateResponseRulesActionExecute, PhaseUpdateResponseRulesActionJSChallenge, PhaseUpdateResponseRulesActionLog, PhaseUpdateResponseRulesActionManagedChallenge, PhaseUpdateResponseRulesActionRedirect, PhaseUpdateResponseRulesActionRewrite, PhaseUpdateResponseRulesActionRoute, PhaseUpdateResponseRulesActionScore, PhaseUpdateResponseRulesActionServeError, PhaseUpdateResponseRulesActionSetConfig, PhaseUpdateResponseRulesActionSkip, PhaseUpdateResponseRulesActionSetCacheSettings, PhaseUpdateResponseRulesActionLogCustomField:
		return true
	}
	return false
}

// A ruleset object.
type PhaseGetResponse struct {
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
	Rules []PhaseGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
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
	// The action to perform when the rule matches.
	Action           PhaseGetResponseRulesAction `json:"action"`
	ActionParameters interface{}                 `json:"action_parameters,required"`
	Categories       interface{}                 `json:"categories,required"`
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
	Version string                   `json:"version,required"`
	JSON    phaseGetResponseRuleJSON `json:"-"`
	union   PhaseGetResponseRulesUnion
}

// phaseGetResponseRuleJSON contains the JSON metadata for the struct
// [PhaseGetResponseRule]
type phaseGetResponseRuleJSON struct {
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

func (r phaseGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *PhaseGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r PhaseGetResponseRule) AsUnion() PhaseGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JSChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule], [rulesets.SetCacheSettingsRule] or
// [rulesets.PhaseGetResponseRulesRulesetsLogCustomFieldRule].
type PhaseGetResponseRulesUnion interface {
	implementsRulesetsPhaseGetResponseRule()
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
			Type:               reflect.TypeOf(JSChallengeRule{}),
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
			Type:               reflect.TypeOf(PhaseGetResponseRulesRulesetsLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
	)
}

type PhaseGetResponseRulesRulesetsLogCustomFieldRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action PhaseGetResponseRulesRulesetsLogCustomFieldRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParameters `json:"action_parameters"`
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
	Ref  string                                              `json:"ref"`
	JSON phaseGetResponseRulesRulesetsLogCustomFieldRuleJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsLogCustomFieldRuleJSON contains the JSON metadata
// for the struct [PhaseGetResponseRulesRulesetsLogCustomFieldRule]
type phaseGetResponseRulesRulesetsLogCustomFieldRuleJSON struct {
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

func (r *PhaseGetResponseRulesRulesetsLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r PhaseGetResponseRulesRulesetsLogCustomFieldRule) implementsRulesetsPhaseGetResponseRule() {}

// The action to perform when the rule matches.
type PhaseGetResponseRulesRulesetsLogCustomFieldRuleAction string

const (
	PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField PhaseGetResponseRulesRulesetsLogCustomFieldRuleAction = "log_custom_field"
)

func (r PhaseGetResponseRulesRulesetsLogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields []PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField `json:"cookie_fields"`
	// The request fields to log.
	RequestFields []PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField `json:"request_fields"`
	// The response fields to log.
	ResponseFields []PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField `json:"response_fields"`
	JSON           phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON            `json:"-"`
}

// phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON contains the
// JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParameters]
type phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON struct {
	CookieFields   apijson.Field
	RequestFields  apijson.Field
	ResponseFields apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The cookie field to log.
type PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name string                                                                         `json:"name,required"`
	JSON phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField]
type phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON) RawJSON() string {
	return r.raw
}

// The request field to log.
type PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name string                                                                          `json:"name,required"`
	JSON phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField]
type phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON) RawJSON() string {
	return r.raw
}

// The response field to log.
type PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name string                                                                           `json:"name,required"`
	JSON phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON `json:"-"`
}

// phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON
// contains the JSON metadata for the struct
// [PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField]
type phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PhaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r phaseGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type PhaseGetResponseRulesAction string

const (
	PhaseGetResponseRulesActionBlock            PhaseGetResponseRulesAction = "block"
	PhaseGetResponseRulesActionChallenge        PhaseGetResponseRulesAction = "challenge"
	PhaseGetResponseRulesActionCompressResponse PhaseGetResponseRulesAction = "compress_response"
	PhaseGetResponseRulesActionExecute          PhaseGetResponseRulesAction = "execute"
	PhaseGetResponseRulesActionJSChallenge      PhaseGetResponseRulesAction = "js_challenge"
	PhaseGetResponseRulesActionLog              PhaseGetResponseRulesAction = "log"
	PhaseGetResponseRulesActionManagedChallenge PhaseGetResponseRulesAction = "managed_challenge"
	PhaseGetResponseRulesActionRedirect         PhaseGetResponseRulesAction = "redirect"
	PhaseGetResponseRulesActionRewrite          PhaseGetResponseRulesAction = "rewrite"
	PhaseGetResponseRulesActionRoute            PhaseGetResponseRulesAction = "route"
	PhaseGetResponseRulesActionScore            PhaseGetResponseRulesAction = "score"
	PhaseGetResponseRulesActionServeError       PhaseGetResponseRulesAction = "serve_error"
	PhaseGetResponseRulesActionSetConfig        PhaseGetResponseRulesAction = "set_config"
	PhaseGetResponseRulesActionSkip             PhaseGetResponseRulesAction = "skip"
	PhaseGetResponseRulesActionSetCacheSettings PhaseGetResponseRulesAction = "set_cache_settings"
	PhaseGetResponseRulesActionLogCustomField   PhaseGetResponseRulesAction = "log_custom_field"
)

func (r PhaseGetResponseRulesAction) IsKnown() bool {
	switch r {
	case PhaseGetResponseRulesActionBlock, PhaseGetResponseRulesActionChallenge, PhaseGetResponseRulesActionCompressResponse, PhaseGetResponseRulesActionExecute, PhaseGetResponseRulesActionJSChallenge, PhaseGetResponseRulesActionLog, PhaseGetResponseRulesActionManagedChallenge, PhaseGetResponseRulesActionRedirect, PhaseGetResponseRulesActionRewrite, PhaseGetResponseRulesActionRoute, PhaseGetResponseRulesActionScore, PhaseGetResponseRulesActionServeError, PhaseGetResponseRulesActionSetConfig, PhaseGetResponseRulesActionSkip, PhaseGetResponseRulesActionSetCacheSettings, PhaseGetResponseRulesActionLogCustomField:
		return true
	}
	return false
}

type PhaseUpdateParams struct {
	// The list of rules in the ruleset.
	Rules param.Field[[]PhaseUpdateParamsRuleUnion] `json:"rules,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
	// The kind of the ruleset.
	Kind param.Field[Kind] `json:"kind"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name"`
	// The phase of the ruleset.
	Phase param.Field[Phase] `json:"phase"`
}

func (r PhaseUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PhaseUpdateParamsRule struct {
	// The action to perform when the rule matches.
	Action           param.Field[PhaseUpdateParamsRulesAction] `json:"action"`
	ActionParameters param.Field[interface{}]                  `json:"action_parameters,required"`
	Categories       param.Field[interface{}]                  `json:"categories,required"`
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

func (r PhaseUpdateParamsRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// Satisfied by [rulesets.BlockRuleParam], [rulesets.ChallengeRuleParam],
// [rulesets.CompressResponseRuleParam], [rulesets.ExecuteRuleParam],
// [rulesets.JSChallengeRuleParam], [rulesets.LogRuleParam],
// [rulesets.ManagedChallengeRuleParam], [rulesets.RedirectRuleParam],
// [rulesets.RewriteRuleParam], [rulesets.RouteRuleParam],
// [rulesets.ScoreRuleParam], [rulesets.ServeErrorRuleParam],
// [rulesets.SetConfigRuleParam], [rulesets.SkipRuleParam],
// [rulesets.SetCacheSettingsRuleParam],
// [rulesets.PhaseUpdateParamsRulesRulesetsLogCustomFieldRule],
// [PhaseUpdateParamsRule].
type PhaseUpdateParamsRuleUnion interface {
	implementsRulesetsPhaseUpdateParamsRuleUnion()
}

type PhaseUpdateParamsRulesRulesetsLogCustomFieldRule struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionParameters] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[LoggingParam] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r PhaseUpdateParamsRulesRulesetsLogCustomFieldRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PhaseUpdateParamsRulesRulesetsLogCustomFieldRule) implementsRulesetsPhaseUpdateParamsRuleUnion() {
}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleAction string

const (
	PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionLogCustomField PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleAction = "log_custom_field"
)

func (r PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields param.Field[[]PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersCookieField] `json:"cookie_fields"`
	// The request fields to log.
	RequestFields param.Field[[]PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersRequestField] `json:"request_fields"`
	// The response fields to log.
	ResponseFields param.Field[[]PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersResponseField] `json:"response_fields"`
}

func (r PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cookie field to log.
type PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
}

func (r PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersCookieField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The request field to log.
type PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
}

func (r PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersRequestField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response field to log.
type PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name param.Field[string] `json:"name,required"`
}

func (r PhaseUpdateParamsRulesRulesetsLogCustomFieldRuleActionParametersResponseField) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The action to perform when the rule matches.
type PhaseUpdateParamsRulesAction string

const (
	PhaseUpdateParamsRulesActionBlock            PhaseUpdateParamsRulesAction = "block"
	PhaseUpdateParamsRulesActionChallenge        PhaseUpdateParamsRulesAction = "challenge"
	PhaseUpdateParamsRulesActionCompressResponse PhaseUpdateParamsRulesAction = "compress_response"
	PhaseUpdateParamsRulesActionExecute          PhaseUpdateParamsRulesAction = "execute"
	PhaseUpdateParamsRulesActionJSChallenge      PhaseUpdateParamsRulesAction = "js_challenge"
	PhaseUpdateParamsRulesActionLog              PhaseUpdateParamsRulesAction = "log"
	PhaseUpdateParamsRulesActionManagedChallenge PhaseUpdateParamsRulesAction = "managed_challenge"
	PhaseUpdateParamsRulesActionRedirect         PhaseUpdateParamsRulesAction = "redirect"
	PhaseUpdateParamsRulesActionRewrite          PhaseUpdateParamsRulesAction = "rewrite"
	PhaseUpdateParamsRulesActionRoute            PhaseUpdateParamsRulesAction = "route"
	PhaseUpdateParamsRulesActionScore            PhaseUpdateParamsRulesAction = "score"
	PhaseUpdateParamsRulesActionServeError       PhaseUpdateParamsRulesAction = "serve_error"
	PhaseUpdateParamsRulesActionSetConfig        PhaseUpdateParamsRulesAction = "set_config"
	PhaseUpdateParamsRulesActionSkip             PhaseUpdateParamsRulesAction = "skip"
	PhaseUpdateParamsRulesActionSetCacheSettings PhaseUpdateParamsRulesAction = "set_cache_settings"
	PhaseUpdateParamsRulesActionLogCustomField   PhaseUpdateParamsRulesAction = "log_custom_field"
)

func (r PhaseUpdateParamsRulesAction) IsKnown() bool {
	switch r {
	case PhaseUpdateParamsRulesActionBlock, PhaseUpdateParamsRulesActionChallenge, PhaseUpdateParamsRulesActionCompressResponse, PhaseUpdateParamsRulesActionExecute, PhaseUpdateParamsRulesActionJSChallenge, PhaseUpdateParamsRulesActionLog, PhaseUpdateParamsRulesActionManagedChallenge, PhaseUpdateParamsRulesActionRedirect, PhaseUpdateParamsRulesActionRewrite, PhaseUpdateParamsRulesActionRoute, PhaseUpdateParamsRulesActionScore, PhaseUpdateParamsRulesActionServeError, PhaseUpdateParamsRulesActionSetConfig, PhaseUpdateParamsRulesActionSkip, PhaseUpdateParamsRulesActionSetCacheSettings, PhaseUpdateParamsRulesActionLogCustomField:
		return true
	}
	return false
}

// A response object.
type PhaseUpdateResponseEnvelope struct {
	// A list of error messages.
	Errors []PhaseUpdateResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseUpdateResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result PhaseUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success PhaseUpdateResponseEnvelopeSuccess `json:"success,required"`
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
	Message string `json:"message,required"`
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
	Pointer string                                      `json:"pointer,required"`
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
	Message string `json:"message,required"`
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
	Pointer string                                        `json:"pointer,required"`
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
	Errors []PhaseGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []PhaseGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result PhaseGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success PhaseGetResponseEnvelopeSuccess `json:"success,required"`
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
	Message string `json:"message,required"`
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
	Pointer string                                   `json:"pointer,required"`
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
	Message string `json:"message,required"`
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
	Pointer string                                     `json:"pointer,required"`
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
