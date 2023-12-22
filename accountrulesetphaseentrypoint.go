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

// AccountRulesetPhaseEntrypointService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountRulesetPhaseEntrypointService] method instead.
type AccountRulesetPhaseEntrypointService struct {
	Options  []option.RequestOption
	Versions *AccountRulesetPhaseEntrypointVersionService
}

// NewAccountRulesetPhaseEntrypointService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountRulesetPhaseEntrypointService(opts ...option.RequestOption) (r *AccountRulesetPhaseEntrypointService) {
	r = &AccountRulesetPhaseEntrypointService{}
	r.Options = opts
	r.Versions = NewAccountRulesetPhaseEntrypointVersionService(opts...)
	return
}

// Fetches the latest version of the account entry point ruleset for a given phase.
func (r *AccountRulesetPhaseEntrypointService) AccountRulesetsGetAnAccountEntryPointRuleset(ctx context.Context, accountID string, rulesetPhase string, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/phases/%s/entrypoint", accountID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an account entry point ruleset, creating a new version.
func (r *AccountRulesetPhaseEntrypointService) AccountRulesetsUpdateAnAccountEntryPointRuleset(ctx context.Context, accountID string, rulesetPhase string, body AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParams, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/phases/%s/entrypoint", accountID, rulesetPhase)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParams struct {
	// The list of rules in the ruleset.
	Rules param.Field[[]AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object.
//
// Satisfied by
// [AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesCreateUpdateRule],
// [AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesObject].
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRule interface {
	implementsAccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRule()
}

// A rule object.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesCreateUpdateRule struct {
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
	Logging param.Field[AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesCreateUpdateRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesCreateUpdateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesCreateUpdateRule) implementsAccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRule() {
}

// An object configuring the rule's logging behavior.
type AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesCreateUpdateRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountRulesetPhaseEntrypointAccountRulesetsUpdateAnAccountEntryPointRulesetParamsRulesCreateUpdateRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
