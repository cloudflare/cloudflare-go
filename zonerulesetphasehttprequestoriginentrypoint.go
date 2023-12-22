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

// ZoneRulesetPhaseHTTPRequestOriginEntrypointService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneRulesetPhaseHTTPRequestOriginEntrypointService] method instead.
type ZoneRulesetPhaseHTTPRequestOriginEntrypointService struct {
	Options []option.RequestOption
}

// NewZoneRulesetPhaseHTTPRequestOriginEntrypointService generates a new service
// that applies the given options to each request. These options are applied after
// the parent client's options (if there is one), and before any request-specific
// options.
func NewZoneRulesetPhaseHTTPRequestOriginEntrypointService(opts ...option.RequestOption) (r *ZoneRulesetPhaseHTTPRequestOriginEntrypointService) {
	r = &ZoneRulesetPhaseHTTPRequestOriginEntrypointService{}
	r.Options = opts
	return
}

// Fetches all Origin Rules in a zone.
func (r *ZoneRulesetPhaseHTTPRequestOriginEntrypointService) OriginRulesListOriginRules(ctx context.Context, zoneID string, opts ...option.RequestOption) (res *SchemasRuleset, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/http_request_origin/entrypoint", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates the Origin Rules of a zone.
func (r *ZoneRulesetPhaseHTTPRequestOriginEntrypointService) OriginRulesUpdateOriginRules(ctx context.Context, zoneID string, body ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/phases/http_request_origin/entrypoint", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type SchemasRuleset struct {
	ID          interface{} `json:"id"`
	Description interface{} `json:"description"`
	Kind        interface{} `json:"kind"`
	Name        interface{} `json:"name"`
	Phase       interface{} `json:"phase"`
	// The rules in the ruleset.
	Rules []SchemasRulesetRule `json:"rules"`
	JSON  schemasRulesetJSON   `json:"-"`
}

// schemasRulesetJSON contains the JSON metadata for the struct [SchemasRuleset]
type schemasRulesetJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasRuleset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasRulesetRule struct {
	ID     interface{} `json:"id"`
	Action interface{} `json:"action"`
	// The parameters configuring the action.
	ActionParameters SchemasRulesetRulesActionParameters `json:"action_parameters"`
	Description      interface{}                         `json:"description"`
	Expression       interface{}                         `json:"expression"`
	Version          interface{}                         `json:"version"`
	JSON             schemasRulesetRuleJSON              `json:"-"`
}

// schemasRulesetRuleJSON contains the JSON metadata for the struct
// [SchemasRulesetRule]
type schemasRulesetRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Description      apijson.Field
	Expression       apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *SchemasRulesetRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The parameters configuring the action.
type SchemasRulesetRulesActionParameters struct {
	// The value of the Host header.
	HostHeader string `json:"host_header"`
	// The parameters that control where the origin is.
	Origin SchemasRulesetRulesActionParametersOrigin `json:"origin"`
	// The parameters that control the SNI.
	Sni  SchemasRulesetRulesActionParametersSni  `json:"sni"`
	JSON schemasRulesetRulesActionParametersJSON `json:"-"`
}

// schemasRulesetRulesActionParametersJSON contains the JSON metadata for the
// struct [SchemasRulesetRulesActionParameters]
type schemasRulesetRulesActionParametersJSON struct {
	HostHeader  apijson.Field
	Origin      apijson.Field
	Sni         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasRulesetRulesActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The parameters that control where the origin is.
type SchemasRulesetRulesActionParametersOrigin struct {
	// The host to use for origin.
	Host string `json:"host"`
	// The port to use for origin.
	Port int64                                         `json:"port"`
	JSON schemasRulesetRulesActionParametersOriginJSON `json:"-"`
}

// schemasRulesetRulesActionParametersOriginJSON contains the JSON metadata for the
// struct [SchemasRulesetRulesActionParametersOrigin]
type schemasRulesetRulesActionParametersOriginJSON struct {
	Host        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasRulesetRulesActionParametersOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The parameters that control the SNI.
type SchemasRulesetRulesActionParametersSni struct {
	// The SNI used to connect to the origin.
	Value string                                     `json:"value"`
	JSON  schemasRulesetRulesActionParametersSniJSON `json:"-"`
}

// schemasRulesetRulesActionParametersSniJSON contains the JSON metadata for the
// struct [SchemasRulesetRulesActionParametersSni]
type schemasRulesetRulesActionParametersSniJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasRulesetRulesActionParametersSni) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParams struct {
	// The list of rules in the ruleset.
	Rules param.Field[[]ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object.
//
// Satisfied by
// [ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParamsRulesCreateUpdateRule],
// [ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParamsRulesObject].
type ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParamsRule interface {
	implementsZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParamsRule()
}

// A rule object.
type ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParamsRulesCreateUpdateRule struct {
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
	Logging param.Field[ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParamsRulesCreateUpdateRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParamsRulesCreateUpdateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParamsRulesCreateUpdateRule) implementsZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParamsRule() {
}

// An object configuring the rule's logging behavior.
type ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParamsRulesCreateUpdateRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneRulesetPhaseHTTPRequestOriginEntrypointOriginRulesUpdateOriginRulesParamsRulesCreateUpdateRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
