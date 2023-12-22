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

// AccountRulesetRuleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRulesetRuleService] method
// instead.
type AccountRulesetRuleService struct {
	Options []option.RequestOption
}

// NewAccountRulesetRuleService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountRulesetRuleService(opts ...option.RequestOption) (r *AccountRulesetRuleService) {
	r = &AccountRulesetRuleService{}
	r.Options = opts
	return
}

// Updates an existing rule in an account ruleset.
func (r *AccountRulesetRuleService) Update(ctx context.Context, accountID string, rulesetID string, ruleID string, body AccountRulesetRuleUpdateParams, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/rules/%s", accountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Deletes an existing rule from an account ruleset.
func (r *AccountRulesetRuleService) Delete(ctx context.Context, accountID string, rulesetID string, ruleID string, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/rules/%s", accountID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Adds a new rule to an account ruleset. The rule will be added to the end of the
// existing list of rules in the ruleset.
func (r *AccountRulesetRuleService) AccountRulesetsNewAnAccountRulesetRule(ctx context.Context, accountID string, rulesetID string, body AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s/rules", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AccountRulesetRuleUpdateParams struct {
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
	Logging param.Field[AccountRulesetRuleUpdateParamsLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetRuleUpdateParamsLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountRulesetRuleUpdateParamsLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams struct {
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
	Logging param.Field[AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring the rule's logging behavior.
type AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountRulesetRuleAccountRulesetsNewAnAccountRulesetRuleParamsLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
