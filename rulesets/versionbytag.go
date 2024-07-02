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
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// VersionByTagService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVersionByTagService] method instead.
type VersionByTagService struct {
	Options []option.RequestOption
}

// NewVersionByTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVersionByTagService(opts ...option.RequestOption) (r *VersionByTagService) {
	r = &VersionByTagService{}
	r.Options = opts
	return
}

// Fetches the rules of a managed account ruleset version for a given tag.
func (r *VersionByTagService) Get(ctx context.Context, rulesetID string, rulesetVersion string, ruleTag string, query VersionByTagGetParams, opts ...option.RequestOption) (res *VersionByTagGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env VersionByTagGetResponseEnvelope
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if rulesetVersion == "" {
		err = errors.New("missing required ruleset_version parameter")
		return
	}
	if ruleTag == "" {
		err = errors.New("missing required rule_tag parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/rulesets/%s/versions/%s/by_tag/%s", query.AccountID, rulesetID, rulesetVersion, ruleTag)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A ruleset object.
type VersionByTagGetResponse struct {
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
	Rules []VersionByTagGetResponseRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                      `json:"description"`
	JSON        versionByTagGetResponseJSON `json:"-"`
}

// versionByTagGetResponseJSON contains the JSON metadata for the struct
// [VersionByTagGetResponse]
type versionByTagGetResponseJSON struct {
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

func (r *VersionByTagGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseJSON) RawJSON() string {
	return r.raw
}

type VersionByTagGetResponseRule struct {
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesAction `json:"action"`
	// This field can have the runtime type of [BlockRuleActionParameters],
	// [interface{}], [CompressResponseRuleActionParameters],
	// [ExecuteRuleActionParameters], [RedirectRuleActionParameters],
	// [RewriteRuleActionParameters], [RouteRuleActionParameters],
	// [ScoreRuleActionParameters], [ServeErrorRuleActionParameters],
	// [SetConfigRuleActionParameters], [SkipRuleActionParameters],
	// [SetCacheSettingsRuleActionParameters],
	// [VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParameters].
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
	JSON    versionByTagGetResponseRuleJSON `json:"-"`
	union   VersionByTagGetResponseRulesUnion
}

// versionByTagGetResponseRuleJSON contains the JSON metadata for the struct
// [VersionByTagGetResponseRule]
type versionByTagGetResponseRuleJSON struct {
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

func (r versionByTagGetResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *VersionByTagGetResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [VersionByTagGetResponseRulesUnion] interface which you can
// cast to the specific types for more type safety.
//
// Possible runtime types of the union are [rulesets.BlockRule],
// [rulesets.ChallengeRule], [rulesets.CompressResponseRule],
// [rulesets.ExecuteRule], [rulesets.JSChallengeRule], [rulesets.LogRule],
// [rulesets.ManagedChallengeRule], [rulesets.RedirectRule],
// [rulesets.RewriteRule], [rulesets.RouteRule], [rulesets.ScoreRule],
// [rulesets.ServeErrorRule], [rulesets.SetConfigRule], [rulesets.SkipRule],
// [rulesets.SetCacheSettingsRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsLogCustomFieldRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsDDoSDynamicRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsForceConnectionCloseRule].
func (r VersionByTagGetResponseRule) AsUnion() VersionByTagGetResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JSChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule], [rulesets.SetCacheSettingsRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsLogCustomFieldRule],
// [rulesets.VersionByTagGetResponseRulesRulesetsDDoSDynamicRule] or
// [rulesets.VersionByTagGetResponseRulesRulesetsForceConnectionCloseRule].
type VersionByTagGetResponseRulesUnion interface {
	implementsRulesetsVersionByTagGetResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*VersionByTagGetResponseRulesUnion)(nil)).Elem(),
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
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsDDoSDynamicRule{}),
			DiscriminatorValue: "ddos_dynamic",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(VersionByTagGetResponseRulesRulesetsForceConnectionCloseRule{}),
			DiscriminatorValue: "force_connection_close",
		},
	)
}

type VersionByTagGetResponseRulesRulesetsLogCustomFieldRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParameters `json:"action_parameters"`
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
	Ref  string                                                     `json:"ref"`
	JSON versionByTagGetResponseRulesRulesetsLogCustomFieldRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsLogCustomFieldRuleJSON contains the JSON
// metadata for the struct [VersionByTagGetResponseRulesRulesetsLogCustomFieldRule]
type versionByTagGetResponseRulesRulesetsLogCustomFieldRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsLogCustomFieldRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleAction = "log_custom_field"
)

func (r VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields []VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField `json:"cookie_fields"`
	// The request fields to log.
	RequestFields []VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField `json:"request_fields"`
	// The response fields to log.
	ResponseFields []VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField `json:"response_fields"`
	JSON           versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON            `json:"-"`
}

// versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParameters]
type versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON struct {
	CookieFields   apijson.Field
	RequestFields  apijson.Field
	ResponseFields apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The cookie field to log.
type VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name string                                                                                `json:"name,required"`
	JSON versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField]
type versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON) RawJSON() string {
	return r.raw
}

// The request field to log.
type VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name string                                                                                 `json:"name,required"`
	JSON versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField]
type versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON) RawJSON() string {
	return r.raw
}

// The response field to log.
type VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name string                                                                                  `json:"name,required"`
	JSON versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON
// contains the JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField]
type versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON) RawJSON() string {
	return r.raw
}

type VersionByTagGetResponseRulesRulesetsDDoSDynamicRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsDDoSDynamicRuleAction `json:"action"`
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
	JSON versionByTagGetResponseRulesRulesetsDDoSDynamicRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsDDoSDynamicRuleJSON contains the JSON
// metadata for the struct [VersionByTagGetResponseRulesRulesetsDDoSDynamicRule]
type versionByTagGetResponseRulesRulesetsDDoSDynamicRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsDDoSDynamicRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsDDoSDynamicRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsDDoSDynamicRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsDDoSDynamicRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsDDoSDynamicRuleActionDDoSDynamic VersionByTagGetResponseRulesRulesetsDDoSDynamicRuleAction = "ddos_dynamic"
)

func (r VersionByTagGetResponseRulesRulesetsDDoSDynamicRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsDDoSDynamicRuleActionDDoSDynamic:
		return true
	}
	return false
}

type VersionByTagGetResponseRulesRulesetsForceConnectionCloseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action VersionByTagGetResponseRulesRulesetsForceConnectionCloseRuleAction `json:"action"`
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
	Ref  string                                                           `json:"ref"`
	JSON versionByTagGetResponseRulesRulesetsForceConnectionCloseRuleJSON `json:"-"`
}

// versionByTagGetResponseRulesRulesetsForceConnectionCloseRuleJSON contains the
// JSON metadata for the struct
// [VersionByTagGetResponseRulesRulesetsForceConnectionCloseRule]
type versionByTagGetResponseRulesRulesetsForceConnectionCloseRuleJSON struct {
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

func (r *VersionByTagGetResponseRulesRulesetsForceConnectionCloseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseRulesRulesetsForceConnectionCloseRuleJSON) RawJSON() string {
	return r.raw
}

func (r VersionByTagGetResponseRulesRulesetsForceConnectionCloseRule) implementsRulesetsVersionByTagGetResponseRule() {
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesRulesetsForceConnectionCloseRuleAction string

const (
	VersionByTagGetResponseRulesRulesetsForceConnectionCloseRuleActionForceConnectionClose VersionByTagGetResponseRulesRulesetsForceConnectionCloseRuleAction = "force_connection_close"
)

func (r VersionByTagGetResponseRulesRulesetsForceConnectionCloseRuleAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesRulesetsForceConnectionCloseRuleActionForceConnectionClose:
		return true
	}
	return false
}

// The action to perform when the rule matches.
type VersionByTagGetResponseRulesAction string

const (
	VersionByTagGetResponseRulesActionBlock                VersionByTagGetResponseRulesAction = "block"
	VersionByTagGetResponseRulesActionChallenge            VersionByTagGetResponseRulesAction = "challenge"
	VersionByTagGetResponseRulesActionCompressResponse     VersionByTagGetResponseRulesAction = "compress_response"
	VersionByTagGetResponseRulesActionExecute              VersionByTagGetResponseRulesAction = "execute"
	VersionByTagGetResponseRulesActionJSChallenge          VersionByTagGetResponseRulesAction = "js_challenge"
	VersionByTagGetResponseRulesActionLog                  VersionByTagGetResponseRulesAction = "log"
	VersionByTagGetResponseRulesActionManagedChallenge     VersionByTagGetResponseRulesAction = "managed_challenge"
	VersionByTagGetResponseRulesActionRedirect             VersionByTagGetResponseRulesAction = "redirect"
	VersionByTagGetResponseRulesActionRewrite              VersionByTagGetResponseRulesAction = "rewrite"
	VersionByTagGetResponseRulesActionRoute                VersionByTagGetResponseRulesAction = "route"
	VersionByTagGetResponseRulesActionScore                VersionByTagGetResponseRulesAction = "score"
	VersionByTagGetResponseRulesActionServeError           VersionByTagGetResponseRulesAction = "serve_error"
	VersionByTagGetResponseRulesActionSetConfig            VersionByTagGetResponseRulesAction = "set_config"
	VersionByTagGetResponseRulesActionSkip                 VersionByTagGetResponseRulesAction = "skip"
	VersionByTagGetResponseRulesActionSetCacheSettings     VersionByTagGetResponseRulesAction = "set_cache_settings"
	VersionByTagGetResponseRulesActionLogCustomField       VersionByTagGetResponseRulesAction = "log_custom_field"
	VersionByTagGetResponseRulesActionDDoSDynamic          VersionByTagGetResponseRulesAction = "ddos_dynamic"
	VersionByTagGetResponseRulesActionForceConnectionClose VersionByTagGetResponseRulesAction = "force_connection_close"
)

func (r VersionByTagGetResponseRulesAction) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseRulesActionBlock, VersionByTagGetResponseRulesActionChallenge, VersionByTagGetResponseRulesActionCompressResponse, VersionByTagGetResponseRulesActionExecute, VersionByTagGetResponseRulesActionJSChallenge, VersionByTagGetResponseRulesActionLog, VersionByTagGetResponseRulesActionManagedChallenge, VersionByTagGetResponseRulesActionRedirect, VersionByTagGetResponseRulesActionRewrite, VersionByTagGetResponseRulesActionRoute, VersionByTagGetResponseRulesActionScore, VersionByTagGetResponseRulesActionServeError, VersionByTagGetResponseRulesActionSetConfig, VersionByTagGetResponseRulesActionSkip, VersionByTagGetResponseRulesActionSetCacheSettings, VersionByTagGetResponseRulesActionLogCustomField, VersionByTagGetResponseRulesActionDDoSDynamic, VersionByTagGetResponseRulesActionForceConnectionClose:
		return true
	}
	return false
}

type VersionByTagGetParams struct {
	// The unique ID of the account.
	AccountID param.Field[string] `path:"account_id,required"`
}

// A response object.
type VersionByTagGetResponseEnvelope struct {
	// A list of error messages.
	Errors []VersionByTagGetResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []VersionByTagGetResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result VersionByTagGetResponse `json:"result,required"`
	// Whether the API call was successful.
	Success VersionByTagGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    versionByTagGetResponseEnvelopeJSON    `json:"-"`
}

// versionByTagGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [VersionByTagGetResponseEnvelope]
type versionByTagGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionByTagGetResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionByTagGetResponseEnvelopeErrorsSource `json:"source"`
	JSON   versionByTagGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// versionByTagGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [VersionByTagGetResponseEnvelopeErrors]
type versionByTagGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionByTagGetResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                          `json:"pointer,required"`
	JSON    versionByTagGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// versionByTagGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseEnvelopeErrorsSource]
type versionByTagGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type VersionByTagGetResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source VersionByTagGetResponseEnvelopeMessagesSource `json:"source"`
	JSON   versionByTagGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// versionByTagGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [VersionByTagGetResponseEnvelopeMessages]
type versionByTagGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type VersionByTagGetResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                            `json:"pointer,required"`
	JSON    versionByTagGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// versionByTagGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [VersionByTagGetResponseEnvelopeMessagesSource]
type versionByTagGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *VersionByTagGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r versionByTagGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type VersionByTagGetResponseEnvelopeSuccess bool

const (
	VersionByTagGetResponseEnvelopeSuccessTrue VersionByTagGetResponseEnvelopeSuccess = true
)

func (r VersionByTagGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case VersionByTagGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
