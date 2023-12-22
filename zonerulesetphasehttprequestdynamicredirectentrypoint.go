// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointService contains methods and
// other services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointService] method
// instead.
type ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointService struct {
	Options []option.RequestOption
}

// NewZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointService generates a new
// service that applies the given options to each request. These options are
// applied after the parent client's options (if there is one), and before any
// request-specific options.
func NewZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointService(opts ...option.RequestOption) (r *ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointService) {
	r = &ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointService{}
	r.Options = opts
	return
}

// Fetches all Single Redirect Rules in a zone.
func (r *ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointService) SingleRedirectRulesListSingleRedirectRules(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *Ruleset, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/http_request_dynamic_redirect/entrypoint", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Single Redirect Rules of a zone.
func (r *ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointService) SingleRedirectRulesUpdateSingleRedirectRules(ctx context.Context, zoneID string, body ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/http_request_dynamic_redirect/entrypoint", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type Ruleset struct {
	ID          interface{} `json:"id"`
	Description interface{} `json:"description"`
	Kind        interface{} `json:"kind"`
	Name        interface{} `json:"name"`
	Phase       interface{} `json:"phase"`
	// The rules in the ruleset.
	Rules []RulesetRule `json:"rules"`
	JSON  rulesetJSON   `json:"-"`
}

// rulesetJSON contains the JSON metadata for the struct [Ruleset]
type rulesetJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Ruleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetRule struct {
	ID     interface{} `json:"id"`
	Action interface{} `json:"action"`
	// The parameters configuring the action.
	ActionParameters RulesetRulesActionParameters `json:"action_parameters"`
	Description      interface{}                  `json:"description"`
	Expression       interface{}                  `json:"expression"`
	Version          interface{}                  `json:"version"`
	JSON             rulesetRuleJSON              `json:"-"`
}

// rulesetRuleJSON contains the JSON metadata for the struct [RulesetRule]
type rulesetRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Description      apijson.Field
	Expression       apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The parameters configuring the action.
type RulesetRulesActionParameters struct {
	// The parameters that control the redirect.
	FromValue RulesetRulesActionParametersFromValue `json:"from_value"`
	JSON      rulesetRulesActionParametersJSON      `json:"-"`
}

// rulesetRulesActionParametersJSON contains the JSON metadata for the struct
// [RulesetRulesActionParameters]
type rulesetRulesActionParametersJSON struct {
	FromValue   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRulesActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The parameters that control the redirect.
type RulesetRulesActionParametersFromValue struct {
	// Whether the query string for the request should be carried to the redirect's
	// target url.
	PreserveQueryString bool `json:"preserve_query_string"`
	// The status code to use for the redirect.
	StatusCode int64                                          `json:"status_code"`
	TargetURL  RulesetRulesActionParametersFromValueTargetURL `json:"target_url"`
	JSON       rulesetRulesActionParametersFromValueJSON      `json:"-"`
}

// rulesetRulesActionParametersFromValueJSON contains the JSON metadata for the
// struct [RulesetRulesActionParametersFromValue]
type rulesetRulesActionParametersFromValueJSON struct {
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	TargetURL           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RulesetRulesActionParametersFromValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [RulesetRulesActionParametersFromValueTargetURLObject] or
// [RulesetRulesActionParametersFromValueTargetURLObject].
type RulesetRulesActionParametersFromValueTargetURL interface {
	implementsRulesetRulesActionParametersFromValueTargetURL()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*RulesetRulesActionParametersFromValueTargetURL)(nil)).Elem(), "")
}

type RulesetRulesActionParametersFromValueTargetURLObject struct {
	// An expression defining a dynamic value for the target url of the redirect.
	Expression string                                                   `json:"expression"`
	JSON       rulesetRulesActionParametersFromValueTargetURLObjectJSON `json:"-"`
}

// rulesetRulesActionParametersFromValueTargetURLObjectJSON contains the JSON
// metadata for the struct [RulesetRulesActionParametersFromValueTargetURLObject]
type rulesetRulesActionParametersFromValueTargetURLObjectJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetRulesActionParametersFromValueTargetURLObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r RulesetRulesActionParametersFromValueTargetURLObject) implementsRulesetRulesActionParametersFromValueTargetURL() {
}

type ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParams struct {
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object.
//
// Satisfied by
// [ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParamsRulesCreateUpdateRule],
// [ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParamsRulesObject].
type ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParamsRule interface {
	implementsZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParamsRule()
}

// A rule object.
type ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParamsRulesCreateUpdateRule struct {
	// The action to perform when the rule matches.
	Action param.Field[string] `json:"action,required"`
	// The expression defining which traffic will match the rule.
	Expression param.Field[string] `json:"expression,required"`
	// The parameters configuring the rule action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
	// An informative description of the rule.
	Description param.Field[string] `json:"description"`
	// Whether the rule should be executed.
	Enabled param.Field[bool] `json:"enabled"`
	// An object configuring the rule's logging behavior.
	Logging param.Field[ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParamsRulesCreateUpdateRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParamsRulesCreateUpdateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParamsRulesCreateUpdateRule) implementsZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParamsRule() {
}

// An object configuring the rule's logging behavior.
type ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParamsRulesCreateUpdateRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneRulesetPhaseHTTPRequestDynamicRedirectEntrypointSingleRedirectRulesUpdateSingleRedirectRulesParamsRulesCreateUpdateRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
