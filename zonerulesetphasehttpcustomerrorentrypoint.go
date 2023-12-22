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

// ZoneRulesetPhaseHTTPCustomErrorEntrypointService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRulesetPhaseHTTPCustomErrorEntrypointService] method instead.
type ZoneRulesetPhaseHTTPCustomErrorEntrypointService struct {
	Options []option.RequestOption
}

// NewZoneRulesetPhaseHTTPCustomErrorEntrypointService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewZoneRulesetPhaseHTTPCustomErrorEntrypointService(opts ...option.RequestOption) (r *ZoneRulesetPhaseHTTPCustomErrorEntrypointService) {
	r = &ZoneRulesetPhaseHTTPCustomErrorEntrypointService{}
	r.Options = opts
	return
}

// Fetches all Custom Error Responses in a zone.
func (r *ZoneRulesetPhaseHTTPCustomErrorEntrypointService) CustomErrorResponsesGetCustomErrorResponses(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *CustomErrorResponsesRuleset, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/http_custom_errors/entrypoint", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Custom Error Responses of a zone.
func (r *ZoneRulesetPhaseHTTPCustomErrorEntrypointService) CustomErrorResponsesUpdateCustomErrorResponses(ctx context.Context, zoneID string, body ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/http_custom_errors/entrypoint", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type CustomErrorResponsesRuleset struct {
	ID          interface{} `json:"id"`
	Description interface{} `json:"description"`
	Kind        interface{} `json:"kind"`
	Name        interface{} `json:"name"`
	Phase       interface{} `json:"phase"`
	// The rules in the ruleset.
	Rules []CustomErrorResponsesRulesetRule `json:"rules"`
	JSON  customErrorResponsesRulesetJSON   `json:"-"`
}

// customErrorResponsesRulesetJSON contains the JSON metadata for the struct
// [CustomErrorResponsesRuleset]
type customErrorResponsesRulesetJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomErrorResponsesRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type CustomErrorResponsesRulesetRule struct {
	ID     interface{} `json:"id"`
	Action interface{} `json:"action"`
	// The action parameters for the serve_error action.
	ActionParameters CustomErrorResponsesRulesetRulesActionParameters `json:"action_parameters"`
	Description      interface{}                                      `json:"description"`
	Expression       interface{}                                      `json:"expression"`
	Version          interface{}                                      `json:"version"`
	JSON             customErrorResponsesRulesetRuleJSON              `json:"-"`
}

// customErrorResponsesRulesetRuleJSON contains the JSON metadata for the struct
// [CustomErrorResponsesRulesetRule]
type customErrorResponsesRulesetRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Description      apijson.Field
	Expression       apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *CustomErrorResponsesRulesetRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The action parameters for the serve_error action.
type CustomErrorResponsesRulesetRulesActionParameters struct {
	// The new content for the response error.
	Content string `json:"content"`
	// The content-type of the response error.
	ContentType string `json:"content_type"`
	// The HTTP status code of the response error.
	StatusCode float64                                              `json:"status_code"`
	JSON       customErrorResponsesRulesetRulesActionParametersJSON `json:"-"`
}

// customErrorResponsesRulesetRulesActionParametersJSON contains the JSON metadata
// for the struct [CustomErrorResponsesRulesetRulesActionParameters]
type customErrorResponsesRulesetRulesActionParametersJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CustomErrorResponsesRulesetRulesActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParams struct {
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object.
//
// Satisfied by
// [ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParamsRulesCreateUpdateRule],
// [ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParamsRulesObject].
type ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParamsRule interface {
	implementsZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParamsRule()
}

// A rule object.
type ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParamsRulesCreateUpdateRule struct {
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
	Logging param.Field[ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParamsRulesCreateUpdateRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParamsRulesCreateUpdateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParamsRulesCreateUpdateRule) implementsZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParamsRule() {
}

// An object configuring the rule's logging behavior.
type ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParamsRulesCreateUpdateRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneRulesetPhaseHTTPCustomErrorEntrypointCustomErrorResponsesUpdateCustomErrorResponsesParamsRulesCreateUpdateRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
