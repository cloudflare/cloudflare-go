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

// AccountRulesetService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountRulesetService] method
// instead.
type AccountRulesetService struct {
	Options  []option.RequestOption
	Phases   *AccountRulesetPhaseService
	Rules    *AccountRulesetRuleService
	Versions *AccountRulesetVersionService
}

// NewAccountRulesetService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountRulesetService(opts ...option.RequestOption) (r *AccountRulesetService) {
	r = &AccountRulesetService{}
	r.Options = opts
	r.Phases = NewAccountRulesetPhaseService(opts...)
	r.Rules = NewAccountRulesetRuleService(opts...)
	r.Versions = NewAccountRulesetVersionService(opts...)
	return
}

// Fetches the latest version of an account ruleset.
func (r *AccountRulesetService) Get(ctx context.Context, accountID string, rulesetID string, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an account ruleset, creating a new version.
func (r *AccountRulesetService) Update(ctx context.Context, accountID string, rulesetID string, body AccountRulesetUpdateParams, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes all versions of an existing account ruleset.
func (r *AccountRulesetService) Delete(ctx context.Context, accountID string, rulesetID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := fmt.Sprintf("accounts/%s/rulesets/%s", accountID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Creates a ruleset at the account level.
func (r *AccountRulesetService) AccountRulesetsNewAnAccountRuleset(ctx context.Context, accountID string, body AccountRulesetAccountRulesetsNewAnAccountRulesetParams, opts ...option.RequestOption) (res *RulesetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches all rulesets at the account level.
func (r *AccountRulesetService) AccountRulesetsListAccountRulesets(ctx context.Context, accountID string, opts ...option.RequestOption) (res *RulesetsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/rulesets", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type RulesetResponse struct {
	Errors   []RulesetResponseError   `json:"errors"`
	Messages []RulesetResponseMessage `json:"messages"`
	// A ruleset object.
	Result RulesetResponseResult `json:"result"`
	// Whether the API call was successful
	Success RulesetResponseSuccess `json:"success"`
	JSON    rulesetResponseJSON    `json:"-"`
}

// rulesetResponseJSON contains the JSON metadata for the struct [RulesetResponse]
type rulesetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetResponseError struct {
	Code    int64                    `json:"code,required"`
	Message string                   `json:"message,required"`
	JSON    rulesetResponseErrorJSON `json:"-"`
}

// rulesetResponseErrorJSON contains the JSON metadata for the struct
// [RulesetResponseError]
type rulesetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetResponseMessage struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    rulesetResponseMessageJSON `json:"-"`
}

// rulesetResponseMessageJSON contains the JSON metadata for the struct
// [RulesetResponseMessage]
type rulesetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ruleset object.
type RulesetResponseResult struct {
	// The unique ID of the ruleset.
	ID string `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind RulesetResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated string `json:"last_updated"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase string `json:"phase"`
	// The list of rules in the ruleset.
	Rules []RulesetResponseResultRule `json:"rules"`
	// The version of the ruleset.
	Version string                    `json:"version"`
	JSON    rulesetResponseResultJSON `json:"-"`
}

// rulesetResponseResultJSON contains the JSON metadata for the struct
// [RulesetResponseResult]
type rulesetResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Rules       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetResponseResultKind string

const (
	RulesetResponseResultKindCustom RulesetResponseResultKind = "custom"
	RulesetResponseResultKindRoot   RulesetResponseResultKind = "root"
	RulesetResponseResultKindZone   RulesetResponseResultKind = "zone"
)

// A rule object.
type RulesetResponseResultRule struct {
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action string `json:"action"`
	// The parameters configuring the rule action.
	ActionParameters interface{} `json:"action_parameters"`
	// The categories of the rule.
	Categories []string `json:"categories"`
	// An informative description of the rule.
	Description string `json:"description"`
	// Whether the rule should be executed.
	Enabled bool `json:"enabled"`
	// The expression defining which traffic will match the rule.
	Expression string `json:"expression"`
	// The timestamp of when the rule was last modified.
	LastUpdated string `json:"last_updated"`
	// An object configuring the rule's logging behavior.
	Logging RulesetResponseResultRulesLogging `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref string `json:"ref"`
	// The version of the rule.
	Version string                        `json:"version"`
	JSON    rulesetResponseResultRuleJSON `json:"-"`
}

// rulesetResponseResultRuleJSON contains the JSON metadata for the struct
// [RulesetResponseResultRule]
type rulesetResponseResultRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	ActionParameters apijson.Field
	Categories       apijson.Field
	Description      apijson.Field
	Enabled          apijson.Field
	Expression       apijson.Field
	LastUpdated      apijson.Field
	Logging          apijson.Field
	Ref              apijson.Field
	Version          apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RulesetResponseResultRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// An object configuring the rule's logging behavior.
type RulesetResponseResultRulesLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool                                  `json:"enabled"`
	JSON    rulesetResponseResultRulesLoggingJSON `json:"-"`
}

// rulesetResponseResultRulesLoggingJSON contains the JSON metadata for the struct
// [RulesetResponseResultRulesLogging]
type rulesetResponseResultRulesLoggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetResponseResultRulesLogging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type RulesetResponseSuccess bool

const (
	RulesetResponseSuccessTrue RulesetResponseSuccess = true
)

type RulesetsResponse struct {
	Errors   []RulesetsResponseError   `json:"errors"`
	Messages []RulesetsResponseMessage `json:"messages"`
	// A list of rulesets. The returned information will not include the rules in each
	// ruleset.
	Result []RulesetsResponseResult `json:"result"`
	// Whether the API call was successful
	Success RulesetsResponseSuccess `json:"success"`
	JSON    rulesetsResponseJSON    `json:"-"`
}

// rulesetsResponseJSON contains the JSON metadata for the struct
// [RulesetsResponse]
type rulesetsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetsResponseError struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    rulesetsResponseErrorJSON `json:"-"`
}

// rulesetsResponseErrorJSON contains the JSON metadata for the struct
// [RulesetsResponseError]
type rulesetsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RulesetsResponseMessage struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    rulesetsResponseMessageJSON `json:"-"`
}

// rulesetsResponseMessageJSON contains the JSON metadata for the struct
// [RulesetsResponseMessage]
type rulesetsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A ruleset object.
type RulesetsResponseResult struct {
	// The unique ID of the ruleset.
	ID string `json:"id"`
	// An informative description of the ruleset.
	Description string `json:"description"`
	// The kind of the ruleset.
	Kind RulesetsResponseResultKind `json:"kind"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated string `json:"last_updated"`
	// The human-readable name of the ruleset.
	Name string `json:"name"`
	// The phase of the ruleset.
	Phase string `json:"phase"`
	// The version of the ruleset.
	Version string                     `json:"version"`
	JSON    rulesetsResponseResultJSON `json:"-"`
}

// rulesetsResponseResultJSON contains the JSON metadata for the struct
// [RulesetsResponseResult]
type rulesetsResponseResultJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Kind        apijson.Field
	LastUpdated apijson.Field
	Name        apijson.Field
	Phase       apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RulesetsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The kind of the ruleset.
type RulesetsResponseResultKind string

const (
	RulesetsResponseResultKindCustom RulesetsResponseResultKind = "custom"
	RulesetsResponseResultKindRoot   RulesetsResponseResultKind = "root"
	RulesetsResponseResultKindZone   RulesetsResponseResultKind = "zone"
)

// Whether the API call was successful
type RulesetsResponseSuccess bool

const (
	RulesetsResponseSuccessTrue RulesetsResponseSuccess = true
)

type AccountRulesetUpdateParams struct {
	// The list of rules in the ruleset.
	Rules param.Field[[]AccountRulesetUpdateParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r AccountRulesetUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule object.
//
// Satisfied by [AccountRulesetUpdateParamsRulesCreateUpdateRule],
// [AccountRulesetUpdateParamsRulesObject].
type AccountRulesetUpdateParamsRule interface {
	implementsAccountRulesetUpdateParamsRule()
}

// A rule object.
type AccountRulesetUpdateParamsRulesCreateUpdateRule struct {
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
	Logging param.Field[AccountRulesetUpdateParamsRulesCreateUpdateRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetUpdateParamsRulesCreateUpdateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetUpdateParamsRulesCreateUpdateRule) implementsAccountRulesetUpdateParamsRule() {}

// An object configuring the rule's logging behavior.
type AccountRulesetUpdateParamsRulesCreateUpdateRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountRulesetUpdateParamsRulesCreateUpdateRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountRulesetAccountRulesetsNewAnAccountRulesetParams struct {
	// The kind of the ruleset.
	Kind param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKind] `json:"kind,required"`
	// The human-readable name of the ruleset.
	Name param.Field[string] `json:"name,required"`
	// The phase of the ruleset.
	Phase param.Field[string] `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules param.Field[[]AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule] `json:"rules,required"`
	// An informative description of the ruleset.
	Description param.Field[string] `json:"description"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The kind of the ruleset.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKind string

const (
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKindCustom AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKind = "custom"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKindRoot   AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKind = "root"
	AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKindZone   AccountRulesetAccountRulesetsNewAnAccountRulesetParamsKind = "zone"
)

// A rule object.
//
// Satisfied by
// [AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRule],
// [AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesObject].
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule interface {
	implementsAccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule()
}

// A rule object.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRule struct {
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
	Logging param.Field[AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRuleLogging] `json:"logging"`
	// The reference of the rule (the rule ID by default).
	Ref param.Field[string] `json:"ref"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRule) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRule) implementsAccountRulesetAccountRulesetsNewAnAccountRulesetParamsRule() {
}

// An object configuring the rule's logging behavior.
type AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRuleLogging struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled"`
}

func (r AccountRulesetAccountRulesetsNewAnAccountRulesetParamsRulesCreateUpdateRuleLogging) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
