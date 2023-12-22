// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneRulesetPhaseEntrypointService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneRulesetPhaseEntrypointService] method instead.
type ZoneRulesetPhaseEntrypointService struct {
	Options  []option.RequestOption
	Versions *ZoneRulesetPhaseEntrypointVersionService
}

// NewZoneRulesetPhaseEntrypointService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneRulesetPhaseEntrypointService(opts ...option.RequestOption) (r *ZoneRulesetPhaseEntrypointService) {
	r = &ZoneRulesetPhaseEntrypointService{}
	r.Options = opts
	r.Versions = NewZoneRulesetPhaseEntrypointVersionService(opts...)
	return
}

// Fetches all Transform Rules in a zone.
func (r *ZoneRulesetPhaseEntrypointService) TransformRulesListTransformRules(ctx context.Context, zoneID string, query ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParams, opts ...option.RequestOption) (res *TransformRulesRuleset, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/rulesets/phases/%s/entrypoint", query.Phase, zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Updates the Transform Rules of a zone.
func (r *ZoneRulesetPhaseEntrypointService) TransformRulesUpdateTransformRules(ctx context.Context, zoneID string, params ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%v/rulesets/phases/%s/entrypoint", params.Phase, zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type TransformRulesRuleset struct {
	ID          interface{} `json:"id"`
	Description interface{} `json:"description"`
	Kind        interface{} `json:"kind"`
	Name        interface{} `json:"name"`
	Phase       interface{} `json:"phase"`
	// The rules in the ruleset.
	Rules []TransformRulesRulesetRule `json:"rules"`
	JSON  transformRulesRulesetJSON   `json:"-"`
}

// transformRulesRulesetJSON contains the JSON metadata for the struct
// [TransformRulesRuleset]
type transformRulesRulesetJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformRulesRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type TransformRulesRulesetRule struct {
	ID     interface{} `json:"id"`
	Action interface{} `json:"action"`
	// The parameters configuring the action.
	ActionParameters TransformRulesRulesetRulesActionParameters `json:"action_parameters"`
	Description      interface{}                                `json:"description"`
	Expression       interface{}                                `json:"expression"`
	Version          interface{}                                `json:"version"`
	JSON             transformRulesRulesetRuleJSON              `json:"-"`
}

// transformRulesRulesetRuleJSON contains the JSON metadata for the struct
// [TransformRulesRulesetRule]
type transformRulesRulesetRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Description      apijson.Field
	Expression       apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *TransformRulesRulesetRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The parameters configuring the action.
type TransformRulesRulesetRulesActionParameters struct {
	// The URI rewrite configuration to rewrite the URI path, the query string, or
	// both.
	Uri  TransformRulesRulesetRulesActionParametersUri  `json:"uri"`
	JSON transformRulesRulesetRulesActionParametersJSON `json:"-"`
}

// transformRulesRulesetRulesActionParametersJSON contains the JSON metadata for
// the struct [TransformRulesRulesetRulesActionParameters]
type transformRulesRulesetRulesActionParametersJSON struct {
	Uri         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformRulesRulesetRulesActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The URI rewrite configuration to rewrite the URI path, the query string, or
// both.
type TransformRulesRulesetRulesActionParametersUri struct {
	// The new URI path sent to the origin.
	Path interface{} `json:"path"`
	// The new query string sent to the origin.
	Query interface{}                                       `json:"query"`
	JSON  transformRulesRulesetRulesActionParametersUriJSON `json:"-"`
}

// transformRulesRulesetRulesActionParametersUriJSON contains the JSON metadata for
// the struct [TransformRulesRulesetRulesActionParametersUri]
type transformRulesRulesetRulesActionParametersUriJSON struct {
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TransformRulesRulesetRulesActionParametersUri) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParams struct {
	// The phase where the ruleset is executed.
	Phase param.Field[ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsPhase] `path:"phase,required"`
}

// The phase where the ruleset is executed.
type ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsPhase string

const (
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsPhaseHTTPRequestTransform         ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsPhase = "http_request_transform"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsPhaseHTTPRequestLateTransform     ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsPhase = "http_request_late_transform"
	ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsPhaseHTTPResponseHeadersTransform ZoneRulesetPhaseEntrypointTransformRulesListTransformRulesParamsPhase = "http_response_headers_transform"
)

type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParams struct {
	// The phase where the ruleset is executed.
	Phase param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase] `path:"phase,required"`
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The phase where the ruleset is executed.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase string

const (
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestTransform         ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_transform"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPRequestLateTransform     ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_request_late_transform"
	ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhaseHTTPResponseHeadersTransform ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsPhase = "http_response_headers_transform"
)

// A rule object.
//
// Satisfied by
// [ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesCreateUpdateRule],
// [ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesObject].
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRule interface {
	implementsZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRule()
}

// A rule object.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesCreateUpdateRule struct {
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
	Logging param.Field[ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesCreateUpdateRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesCreateUpdateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesCreateUpdateRule) implementsZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRule() {
}

// An object configuring the rule's logging behavior.
type ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesCreateUpdateRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneRulesetPhaseEntrypointTransformRulesUpdateTransformRulesParamsRulesCreateUpdateRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
