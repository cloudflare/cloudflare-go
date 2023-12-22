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

// ZoneRulesetRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneRulesetRuleService] method
// instead.
type ZoneRulesetRuleService struct {
	Options []option.RequestOption
}

// NewZoneRulesetRuleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneRulesetRuleService(opts ...option.RequestOption) (r *ZoneRulesetRuleService) {
	r = &ZoneRulesetRuleService{}
	r.Options = opts
	return
}

// Updates an existing rule in a zone ruleset.
func (r *ZoneRulesetRuleService) Update(ctx context.Context, zoneID string, rulesetID string, ruleID string, body ZoneRulesetRuleUpdateParams, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s/rules/%s", zoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an existing rule from a zone ruleset.
func (r *ZoneRulesetRuleService) Delete(ctx context.Context, zoneID string, rulesetID string, ruleID string, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s/rules/%s", zoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new rule to a zone ruleset. The rule will be added to the end of the
// existing list of rules in the ruleset.
func (r *ZoneRulesetRuleService) ZoneRulesetsNewAZoneRulesetRule(ctx context.Context, zoneID string, rulesetID string, body ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParams, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/rulesets/%s/rules", zoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneRulesetRuleUpdateParams struct {
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
	Logging param.Field[ZoneRulesetRuleUpdateParamsLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetRuleUpdateParamsLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneRulesetRuleUpdateParamsLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParams struct {
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
	Logging param.Field[ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r ZoneRulesetRuleZoneRulesetsNewAZoneRulesetRuleParamsLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
