// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package rulesets

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RuleService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewRuleService] method instead.
type RuleService struct {
	Options []option.RequestOption
}

// NewRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRuleService(opts ...option.RequestOption) (r *RuleService) {
	r = &RuleService{}
	r.Options = opts
	return
}

// Adds a new rule to an account or zone ruleset. The rule will be added to the end
// of the existing list of rules in the ruleset by default.
func (r *RuleService) New(ctx context.Context, rulesetID string, params RuleNewParams, opts ...option.RequestOption) (res *RuleNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules", accountOrZone, accountOrZoneID, rulesetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing rule from an account or zone ruleset.
func (r *RuleService) Delete(ctx context.Context, rulesetID string, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) (res *RuleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/%s", accountOrZone, accountOrZoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates an existing rule in an account or zone ruleset.
func (r *RuleService) Edit(ctx context.Context, rulesetID string, ruleID string, params RuleEditParams, opts ...option.RequestOption) (res *RuleEditResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RuleEditResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/rulesets/%s/rules/%s", accountOrZone, accountOrZoneID, rulesetID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type BlockRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action BlockRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters BlockRuleActionParameters `json:"action_parameters"`
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
	Ref  string        `json:"ref"`
	JSON blockRuleJSON `json:"-"`
}

// blockRuleJSON contains the JSON metadata for the struct [BlockRule]
type blockRuleJSON struct {
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

func (r *BlockRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockRuleJSON) RawJSON() string {
	return r.raw
}

func (r BlockRule) implementsRulesetsResponeRuleItem() {}

// The action to perform when the rule matches.
type BlockRuleAction string

const (
	BlockRuleActionBlock BlockRuleAction = "block"
)

func (r BlockRuleAction) IsKnown() bool {
	switch r {
	case BlockRuleActionBlock:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type BlockRuleActionParameters struct {
	// The response to show when the block is applied.
	Response BlockRuleActionParametersResponse `json:"response"`
	JSON     blockRuleActionParametersJSON     `json:"-"`
}

// blockRuleActionParametersJSON contains the JSON metadata for the struct
// [BlockRuleActionParameters]
type blockRuleActionParametersJSON struct {
	Response    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The response to show when the block is applied.
type BlockRuleActionParametersResponse struct {
	// The content to return.
	Content string `json:"content,required"`
	// The type of the content to return.
	ContentType string `json:"content_type,required"`
	// The status code to return.
	StatusCode int64                                 `json:"status_code,required"`
	JSON       blockRuleActionParametersResponseJSON `json:"-"`
}

// blockRuleActionParametersResponseJSON contains the JSON metadata for the struct
// [BlockRuleActionParametersResponse]
type blockRuleActionParametersResponseJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *BlockRuleActionParametersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r blockRuleActionParametersResponseJSON) RawJSON() string {
	return r.raw
}

type BlockRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[BlockRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[BlockRuleActionParametersParam] `json:"action_parameters"`
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

func (r BlockRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r BlockRuleParam) implementsRulesetsRequestRuleUnionItemParam() {}

// The parameters configuring the rule's action.
type BlockRuleActionParametersParam struct {
	// The response to show when the block is applied.
	Response param.Field[BlockRuleActionParametersResponseParam] `json:"response"`
}

func (r BlockRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response to show when the block is applied.
type BlockRuleActionParametersResponseParam struct {
	// The content to return.
	Content param.Field[string] `json:"content,required"`
	// The type of the content to return.
	ContentType param.Field[string] `json:"content_type,required"`
	// The status code to return.
	StatusCode param.Field[int64] `json:"status_code,required"`
}

func (r BlockRuleActionParametersResponseParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ExecuteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action ExecuteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters ExecuteRuleActionParameters `json:"action_parameters"`
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
	Ref  string          `json:"ref"`
	JSON executeRuleJSON `json:"-"`
}

// executeRuleJSON contains the JSON metadata for the struct [ExecuteRule]
type executeRuleJSON struct {
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

func (r *ExecuteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleJSON) RawJSON() string {
	return r.raw
}

func (r ExecuteRule) implementsRulesetsResponeRuleItem() {}

// The action to perform when the rule matches.
type ExecuteRuleAction string

const (
	ExecuteRuleActionExecute ExecuteRuleAction = "execute"
)

func (r ExecuteRuleAction) IsKnown() bool {
	switch r {
	case ExecuteRuleActionExecute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type ExecuteRuleActionParameters struct {
	// The ID of the ruleset to execute.
	ID string `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData ExecuteRuleActionParametersMatchedData `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides ExecuteRuleActionParametersOverrides `json:"overrides"`
	JSON      executeRuleActionParametersJSON      `json:"-"`
}

// executeRuleActionParametersJSON contains the JSON metadata for the struct
// [ExecuteRuleActionParameters]
type executeRuleActionParametersJSON struct {
	ID          apijson.Field
	MatchedData apijson.Field
	Overrides   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecuteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The configuration to use for matched data logging.
type ExecuteRuleActionParametersMatchedData struct {
	// The public key to encrypt matched data logs with.
	PublicKey string                                     `json:"public_key,required"`
	JSON      executeRuleActionParametersMatchedDataJSON `json:"-"`
}

// executeRuleActionParametersMatchedDataJSON contains the JSON metadata for the
// struct [ExecuteRuleActionParametersMatchedData]
type executeRuleActionParametersMatchedDataJSON struct {
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ExecuteRuleActionParametersMatchedData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleActionParametersMatchedDataJSON) RawJSON() string {
	return r.raw
}

// A set of overrides to apply to the target ruleset.
type ExecuteRuleActionParametersOverrides struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action string `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories []ExecuteRuleActionParametersOverridesCategory `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled bool `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules []ExecuteRuleActionParametersOverridesRule `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel ExecuteRuleActionParametersOverridesSensitivityLevel `json:"sensitivity_level"`
	JSON             executeRuleActionParametersOverridesJSON             `json:"-"`
}

// executeRuleActionParametersOverridesJSON contains the JSON metadata for the
// struct [ExecuteRuleActionParametersOverrides]
type executeRuleActionParametersOverridesJSON struct {
	Action           apijson.Field
	Categories       apijson.Field
	Enabled          apijson.Field
	Rules            apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ExecuteRuleActionParametersOverrides) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleActionParametersOverridesJSON) RawJSON() string {
	return r.raw
}

// A category-level override
type ExecuteRuleActionParametersOverridesCategory struct {
	// The name of the category to override.
	Category string `json:"category,required"`
	// The action to override rules in the category with.
	Action string `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled bool `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel `json:"sensitivity_level"`
	JSON             executeRuleActionParametersOverridesCategoryJSON               `json:"-"`
}

// executeRuleActionParametersOverridesCategoryJSON contains the JSON metadata for
// the struct [ExecuteRuleActionParametersOverridesCategory]
type executeRuleActionParametersOverridesCategoryJSON struct {
	Category         apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ExecuteRuleActionParametersOverridesCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleActionParametersOverridesCategoryJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for rules in the category.
type ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel string

const (
	ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "default"
	ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium  ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "medium"
	ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow     ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "low"
	ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff    ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel = "eoff"
)

func (r ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel) IsKnown() bool {
	switch r {
	case ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelDefault, ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelMedium, ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelLow, ExecuteRuleActionParametersOverridesCategoriesSensitivityLevelEoff:
		return true
	}
	return false
}

// A rule-level override
type ExecuteRuleActionParametersOverridesRule struct {
	// The ID of the rule to override.
	ID string `json:"id,required"`
	// The action to override the rule with.
	Action string `json:"action"`
	// Whether to enable execution of the rule.
	Enabled bool `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold int64 `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel ExecuteRuleActionParametersOverridesRulesSensitivityLevel `json:"sensitivity_level"`
	JSON             executeRuleActionParametersOverridesRuleJSON              `json:"-"`
}

// executeRuleActionParametersOverridesRuleJSON contains the JSON metadata for the
// struct [ExecuteRuleActionParametersOverridesRule]
type executeRuleActionParametersOverridesRuleJSON struct {
	ID               apijson.Field
	Action           apijson.Field
	Enabled          apijson.Field
	ScoreThreshold   apijson.Field
	SensitivityLevel apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ExecuteRuleActionParametersOverridesRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r executeRuleActionParametersOverridesRuleJSON) RawJSON() string {
	return r.raw
}

// The sensitivity level to use for the rule.
type ExecuteRuleActionParametersOverridesRulesSensitivityLevel string

const (
	ExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault ExecuteRuleActionParametersOverridesRulesSensitivityLevel = "default"
	ExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium  ExecuteRuleActionParametersOverridesRulesSensitivityLevel = "medium"
	ExecuteRuleActionParametersOverridesRulesSensitivityLevelLow     ExecuteRuleActionParametersOverridesRulesSensitivityLevel = "low"
	ExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff    ExecuteRuleActionParametersOverridesRulesSensitivityLevel = "eoff"
)

func (r ExecuteRuleActionParametersOverridesRulesSensitivityLevel) IsKnown() bool {
	switch r {
	case ExecuteRuleActionParametersOverridesRulesSensitivityLevelDefault, ExecuteRuleActionParametersOverridesRulesSensitivityLevelMedium, ExecuteRuleActionParametersOverridesRulesSensitivityLevelLow, ExecuteRuleActionParametersOverridesRulesSensitivityLevelEoff:
		return true
	}
	return false
}

// A sensitivity level to set for all rules. This option has lower precedence than
// rule and category overrides and is only applicable for DDoS phases.
type ExecuteRuleActionParametersOverridesSensitivityLevel string

const (
	ExecuteRuleActionParametersOverridesSensitivityLevelDefault ExecuteRuleActionParametersOverridesSensitivityLevel = "default"
	ExecuteRuleActionParametersOverridesSensitivityLevelMedium  ExecuteRuleActionParametersOverridesSensitivityLevel = "medium"
	ExecuteRuleActionParametersOverridesSensitivityLevelLow     ExecuteRuleActionParametersOverridesSensitivityLevel = "low"
	ExecuteRuleActionParametersOverridesSensitivityLevelEoff    ExecuteRuleActionParametersOverridesSensitivityLevel = "eoff"
)

func (r ExecuteRuleActionParametersOverridesSensitivityLevel) IsKnown() bool {
	switch r {
	case ExecuteRuleActionParametersOverridesSensitivityLevelDefault, ExecuteRuleActionParametersOverridesSensitivityLevelMedium, ExecuteRuleActionParametersOverridesSensitivityLevelLow, ExecuteRuleActionParametersOverridesSensitivityLevelEoff:
		return true
	}
	return false
}

type ExecuteRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[ExecuteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[ExecuteRuleActionParametersParam] `json:"action_parameters"`
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

func (r ExecuteRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ExecuteRuleParam) implementsRulesetsRequestRuleUnionItemParam() {}

// The parameters configuring the rule's action.
type ExecuteRuleActionParametersParam struct {
	// The ID of the ruleset to execute.
	ID param.Field[string] `json:"id,required"`
	// The configuration to use for matched data logging.
	MatchedData param.Field[ExecuteRuleActionParametersMatchedDataParam] `json:"matched_data"`
	// A set of overrides to apply to the target ruleset.
	Overrides param.Field[ExecuteRuleActionParametersOverridesParam] `json:"overrides"`
}

func (r ExecuteRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The configuration to use for matched data logging.
type ExecuteRuleActionParametersMatchedDataParam struct {
	// The public key to encrypt matched data logs with.
	PublicKey param.Field[string] `json:"public_key,required"`
}

func (r ExecuteRuleActionParametersMatchedDataParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A set of overrides to apply to the target ruleset.
type ExecuteRuleActionParametersOverridesParam struct {
	// An action to override all rules with. This option has lower precedence than rule
	// and category overrides.
	Action param.Field[string] `json:"action"`
	// A list of category-level overrides. This option has the second-highest
	// precedence after rule-level overrides.
	Categories param.Field[[]ExecuteRuleActionParametersOverridesCategoryParam] `json:"categories"`
	// Whether to enable execution of all rules. This option has lower precedence than
	// rule and category overrides.
	Enabled param.Field[bool] `json:"enabled"`
	// A list of rule-level overrides. This option has the highest precedence.
	Rules param.Field[[]ExecuteRuleActionParametersOverridesRuleParam] `json:"rules"`
	// A sensitivity level to set for all rules. This option has lower precedence than
	// rule and category overrides and is only applicable for DDoS phases.
	SensitivityLevel param.Field[ExecuteRuleActionParametersOverridesSensitivityLevel] `json:"sensitivity_level"`
}

func (r ExecuteRuleActionParametersOverridesParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A category-level override
type ExecuteRuleActionParametersOverridesCategoryParam struct {
	// The name of the category to override.
	Category param.Field[string] `json:"category,required"`
	// The action to override rules in the category with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of rules in the category.
	Enabled param.Field[bool] `json:"enabled"`
	// The sensitivity level to use for rules in the category.
	SensitivityLevel param.Field[ExecuteRuleActionParametersOverridesCategoriesSensitivityLevel] `json:"sensitivity_level"`
}

func (r ExecuteRuleActionParametersOverridesCategoryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A rule-level override
type ExecuteRuleActionParametersOverridesRuleParam struct {
	// The ID of the rule to override.
	ID param.Field[string] `json:"id,required"`
	// The action to override the rule with.
	Action param.Field[string] `json:"action"`
	// Whether to enable execution of the rule.
	Enabled param.Field[bool] `json:"enabled"`
	// The score threshold to use for the rule.
	ScoreThreshold param.Field[int64] `json:"score_threshold"`
	// The sensitivity level to use for the rule.
	SensitivityLevel param.Field[ExecuteRuleActionParametersOverridesRulesSensitivityLevel] `json:"sensitivity_level"`
}

func (r ExecuteRuleActionParametersOverridesRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LogRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action LogRuleAction `json:"action"`
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
	Ref  string      `json:"ref"`
	JSON logRuleJSON `json:"-"`
}

// logRuleJSON contains the JSON metadata for the struct [LogRule]
type logRuleJSON struct {
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

func (r *LogRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r logRuleJSON) RawJSON() string {
	return r.raw
}

func (r LogRule) implementsRulesetsResponeRuleItem() {}

// The action to perform when the rule matches.
type LogRuleAction string

const (
	LogRuleActionLog LogRuleAction = "log"
)

func (r LogRuleAction) IsKnown() bool {
	switch r {
	case LogRuleActionLog:
		return true
	}
	return false
}

type LogRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[LogRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[interface{}] `json:"action_parameters"`
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

func (r LogRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LogRuleParam) implementsRulesetsRequestRuleUnionItemParam() {}

// An object configuring the rule's logging behavior.
type Logging struct {
	// Whether to generate a log when the rule matches.
	Enabled bool        `json:"enabled,required"`
	JSON    loggingJSON `json:"-"`
}

// loggingJSON contains the JSON metadata for the struct [Logging]
type loggingJSON struct {
	Enabled     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Logging) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loggingJSON) RawJSON() string {
	return r.raw
}

// An object configuring the rule's logging behavior.
type LoggingParam struct {
	// Whether to generate a log when the rule matches.
	Enabled param.Field[bool] `json:"enabled,required"`
}

func (r LoggingParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SkipRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action SkipRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters SkipRuleActionParameters `json:"action_parameters"`
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
	Ref  string       `json:"ref"`
	JSON skipRuleJSON `json:"-"`
}

// skipRuleJSON contains the JSON metadata for the struct [SkipRule]
type skipRuleJSON struct {
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

func (r *SkipRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skipRuleJSON) RawJSON() string {
	return r.raw
}

func (r SkipRule) implementsRulesetsResponeRuleItem() {}

// The action to perform when the rule matches.
type SkipRuleAction string

const (
	SkipRuleActionSkip SkipRuleAction = "skip"
)

func (r SkipRuleAction) IsKnown() bool {
	switch r {
	case SkipRuleActionSkip:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type SkipRuleActionParameters struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases []SkipRuleActionParametersPhase `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products []SkipRuleActionParametersProduct `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules map[string][]string `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset SkipRuleActionParametersRuleset `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets []string                     `json:"rulesets"`
	JSON     skipRuleActionParametersJSON `json:"-"`
}

// skipRuleActionParametersJSON contains the JSON metadata for the struct
// [SkipRuleActionParameters]
type skipRuleActionParametersJSON struct {
	Phases      apijson.Field
	Products    apijson.Field
	Rules       apijson.Field
	Ruleset     apijson.Field
	Rulesets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SkipRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r skipRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// A phase to skip the execution of.
type SkipRuleActionParametersPhase string

const (
	SkipRuleActionParametersPhaseDDoSL4                         SkipRuleActionParametersPhase = "ddos_l4"
	SkipRuleActionParametersPhaseDDoSL7                         SkipRuleActionParametersPhase = "ddos_l7"
	SkipRuleActionParametersPhaseHTTPConfigSettings             SkipRuleActionParametersPhase = "http_config_settings"
	SkipRuleActionParametersPhaseHTTPCustomErrors               SkipRuleActionParametersPhase = "http_custom_errors"
	SkipRuleActionParametersPhaseHTTPLogCustomFields            SkipRuleActionParametersPhase = "http_log_custom_fields"
	SkipRuleActionParametersPhaseHTTPRatelimit                  SkipRuleActionParametersPhase = "http_ratelimit"
	SkipRuleActionParametersPhaseHTTPRequestCacheSettings       SkipRuleActionParametersPhase = "http_request_cache_settings"
	SkipRuleActionParametersPhaseHTTPRequestDynamicRedirect     SkipRuleActionParametersPhase = "http_request_dynamic_redirect"
	SkipRuleActionParametersPhaseHTTPRequestFirewallCustom      SkipRuleActionParametersPhase = "http_request_firewall_custom"
	SkipRuleActionParametersPhaseHTTPRequestFirewallManaged     SkipRuleActionParametersPhase = "http_request_firewall_managed"
	SkipRuleActionParametersPhaseHTTPRequestLateTransform       SkipRuleActionParametersPhase = "http_request_late_transform"
	SkipRuleActionParametersPhaseHTTPRequestOrigin              SkipRuleActionParametersPhase = "http_request_origin"
	SkipRuleActionParametersPhaseHTTPRequestRedirect            SkipRuleActionParametersPhase = "http_request_redirect"
	SkipRuleActionParametersPhaseHTTPRequestSanitize            SkipRuleActionParametersPhase = "http_request_sanitize"
	SkipRuleActionParametersPhaseHTTPRequestSbfm                SkipRuleActionParametersPhase = "http_request_sbfm"
	SkipRuleActionParametersPhaseHTTPRequestSelectConfiguration SkipRuleActionParametersPhase = "http_request_select_configuration"
	SkipRuleActionParametersPhaseHTTPRequestTransform           SkipRuleActionParametersPhase = "http_request_transform"
	SkipRuleActionParametersPhaseHTTPResponseCompression        SkipRuleActionParametersPhase = "http_response_compression"
	SkipRuleActionParametersPhaseHTTPResponseFirewallManaged    SkipRuleActionParametersPhase = "http_response_firewall_managed"
	SkipRuleActionParametersPhaseHTTPResponseHeadersTransform   SkipRuleActionParametersPhase = "http_response_headers_transform"
	SkipRuleActionParametersPhaseMagicTransit                   SkipRuleActionParametersPhase = "magic_transit"
	SkipRuleActionParametersPhaseMagicTransitIDsManaged         SkipRuleActionParametersPhase = "magic_transit_ids_managed"
	SkipRuleActionParametersPhaseMagicTransitManaged            SkipRuleActionParametersPhase = "magic_transit_managed"
)

func (r SkipRuleActionParametersPhase) IsKnown() bool {
	switch r {
	case SkipRuleActionParametersPhaseDDoSL4, SkipRuleActionParametersPhaseDDoSL7, SkipRuleActionParametersPhaseHTTPConfigSettings, SkipRuleActionParametersPhaseHTTPCustomErrors, SkipRuleActionParametersPhaseHTTPLogCustomFields, SkipRuleActionParametersPhaseHTTPRatelimit, SkipRuleActionParametersPhaseHTTPRequestCacheSettings, SkipRuleActionParametersPhaseHTTPRequestDynamicRedirect, SkipRuleActionParametersPhaseHTTPRequestFirewallCustom, SkipRuleActionParametersPhaseHTTPRequestFirewallManaged, SkipRuleActionParametersPhaseHTTPRequestLateTransform, SkipRuleActionParametersPhaseHTTPRequestOrigin, SkipRuleActionParametersPhaseHTTPRequestRedirect, SkipRuleActionParametersPhaseHTTPRequestSanitize, SkipRuleActionParametersPhaseHTTPRequestSbfm, SkipRuleActionParametersPhaseHTTPRequestSelectConfiguration, SkipRuleActionParametersPhaseHTTPRequestTransform, SkipRuleActionParametersPhaseHTTPResponseCompression, SkipRuleActionParametersPhaseHTTPResponseFirewallManaged, SkipRuleActionParametersPhaseHTTPResponseHeadersTransform, SkipRuleActionParametersPhaseMagicTransit, SkipRuleActionParametersPhaseMagicTransitIDsManaged, SkipRuleActionParametersPhaseMagicTransitManaged:
		return true
	}
	return false
}

// The name of a legacy security product to skip the execution of.
type SkipRuleActionParametersProduct string

const (
	SkipRuleActionParametersProductBic           SkipRuleActionParametersProduct = "bic"
	SkipRuleActionParametersProductHot           SkipRuleActionParametersProduct = "hot"
	SkipRuleActionParametersProductRateLimit     SkipRuleActionParametersProduct = "rateLimit"
	SkipRuleActionParametersProductSecurityLevel SkipRuleActionParametersProduct = "securityLevel"
	SkipRuleActionParametersProductUABlock       SkipRuleActionParametersProduct = "uaBlock"
	SkipRuleActionParametersProductWAF           SkipRuleActionParametersProduct = "waf"
	SkipRuleActionParametersProductZoneLockdown  SkipRuleActionParametersProduct = "zoneLockdown"
)

func (r SkipRuleActionParametersProduct) IsKnown() bool {
	switch r {
	case SkipRuleActionParametersProductBic, SkipRuleActionParametersProductHot, SkipRuleActionParametersProductRateLimit, SkipRuleActionParametersProductSecurityLevel, SkipRuleActionParametersProductUABlock, SkipRuleActionParametersProductWAF, SkipRuleActionParametersProductZoneLockdown:
		return true
	}
	return false
}

// A ruleset to skip the execution of. This option is incompatible with the
// rulesets, rules and phases options.
type SkipRuleActionParametersRuleset string

const (
	SkipRuleActionParametersRulesetCurrent SkipRuleActionParametersRuleset = "current"
)

func (r SkipRuleActionParametersRuleset) IsKnown() bool {
	switch r {
	case SkipRuleActionParametersRulesetCurrent:
		return true
	}
	return false
}

type SkipRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[SkipRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[SkipRuleActionParametersParam] `json:"action_parameters"`
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

func (r SkipRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SkipRuleParam) implementsRulesetsRequestRuleUnionItemParam() {}

// The parameters configuring the rule's action.
type SkipRuleActionParametersParam struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]SkipRuleActionParametersPhase] `json:"phases"`
	// A list of legacy security products to skip the execution of.
	Products param.Field[[]SkipRuleActionParametersProduct] `json:"products"`
	// A mapping of ruleset IDs to a list of rule IDs in that ruleset to skip the
	// execution of. This option is incompatible with the ruleset option.
	Rules param.Field[map[string][]string] `json:"rules"`
	// A ruleset to skip the execution of. This option is incompatible with the
	// rulesets, rules and phases options.
	Ruleset param.Field[SkipRuleActionParametersRuleset] `json:"ruleset"`
	// A list of ruleset IDs to skip the execution of. This option is incompatible with
	// the ruleset and phases options.
	Rulesets param.Field[[]string] `json:"rulesets"`
}

func (r SkipRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ruleset object.
type RuleNewResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RuleNewResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RuleNewResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules ResponeRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string              `json:"description"`
	JSON        ruleNewResponseJSON `json:"-"`
}

// ruleNewResponseJSON contains the JSON metadata for the struct [RuleNewResponse]
type ruleNewResponseJSON struct {
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

func (r *RuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RuleNewResponseKind string

const (
	RuleNewResponseKindManaged RuleNewResponseKind = "managed"
	RuleNewResponseKindCustom  RuleNewResponseKind = "custom"
	RuleNewResponseKindRoot    RuleNewResponseKind = "root"
	RuleNewResponseKindZone    RuleNewResponseKind = "zone"
)

func (r RuleNewResponseKind) IsKnown() bool {
	switch r {
	case RuleNewResponseKindManaged, RuleNewResponseKindCustom, RuleNewResponseKindRoot, RuleNewResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type RuleNewResponsePhase string

const (
	RuleNewResponsePhaseDDoSL4                         RuleNewResponsePhase = "ddos_l4"
	RuleNewResponsePhaseDDoSL7                         RuleNewResponsePhase = "ddos_l7"
	RuleNewResponsePhaseHTTPConfigSettings             RuleNewResponsePhase = "http_config_settings"
	RuleNewResponsePhaseHTTPCustomErrors               RuleNewResponsePhase = "http_custom_errors"
	RuleNewResponsePhaseHTTPLogCustomFields            RuleNewResponsePhase = "http_log_custom_fields"
	RuleNewResponsePhaseHTTPRatelimit                  RuleNewResponsePhase = "http_ratelimit"
	RuleNewResponsePhaseHTTPRequestCacheSettings       RuleNewResponsePhase = "http_request_cache_settings"
	RuleNewResponsePhaseHTTPRequestDynamicRedirect     RuleNewResponsePhase = "http_request_dynamic_redirect"
	RuleNewResponsePhaseHTTPRequestFirewallCustom      RuleNewResponsePhase = "http_request_firewall_custom"
	RuleNewResponsePhaseHTTPRequestFirewallManaged     RuleNewResponsePhase = "http_request_firewall_managed"
	RuleNewResponsePhaseHTTPRequestLateTransform       RuleNewResponsePhase = "http_request_late_transform"
	RuleNewResponsePhaseHTTPRequestOrigin              RuleNewResponsePhase = "http_request_origin"
	RuleNewResponsePhaseHTTPRequestRedirect            RuleNewResponsePhase = "http_request_redirect"
	RuleNewResponsePhaseHTTPRequestSanitize            RuleNewResponsePhase = "http_request_sanitize"
	RuleNewResponsePhaseHTTPRequestSbfm                RuleNewResponsePhase = "http_request_sbfm"
	RuleNewResponsePhaseHTTPRequestSelectConfiguration RuleNewResponsePhase = "http_request_select_configuration"
	RuleNewResponsePhaseHTTPRequestTransform           RuleNewResponsePhase = "http_request_transform"
	RuleNewResponsePhaseHTTPResponseCompression        RuleNewResponsePhase = "http_response_compression"
	RuleNewResponsePhaseHTTPResponseFirewallManaged    RuleNewResponsePhase = "http_response_firewall_managed"
	RuleNewResponsePhaseHTTPResponseHeadersTransform   RuleNewResponsePhase = "http_response_headers_transform"
	RuleNewResponsePhaseMagicTransit                   RuleNewResponsePhase = "magic_transit"
	RuleNewResponsePhaseMagicTransitIDsManaged         RuleNewResponsePhase = "magic_transit_ids_managed"
	RuleNewResponsePhaseMagicTransitManaged            RuleNewResponsePhase = "magic_transit_managed"
)

func (r RuleNewResponsePhase) IsKnown() bool {
	switch r {
	case RuleNewResponsePhaseDDoSL4, RuleNewResponsePhaseDDoSL7, RuleNewResponsePhaseHTTPConfigSettings, RuleNewResponsePhaseHTTPCustomErrors, RuleNewResponsePhaseHTTPLogCustomFields, RuleNewResponsePhaseHTTPRatelimit, RuleNewResponsePhaseHTTPRequestCacheSettings, RuleNewResponsePhaseHTTPRequestDynamicRedirect, RuleNewResponsePhaseHTTPRequestFirewallCustom, RuleNewResponsePhaseHTTPRequestFirewallManaged, RuleNewResponsePhaseHTTPRequestLateTransform, RuleNewResponsePhaseHTTPRequestOrigin, RuleNewResponsePhaseHTTPRequestRedirect, RuleNewResponsePhaseHTTPRequestSanitize, RuleNewResponsePhaseHTTPRequestSbfm, RuleNewResponsePhaseHTTPRequestSelectConfiguration, RuleNewResponsePhaseHTTPRequestTransform, RuleNewResponsePhaseHTTPResponseCompression, RuleNewResponsePhaseHTTPResponseFirewallManaged, RuleNewResponsePhaseHTTPResponseHeadersTransform, RuleNewResponsePhaseMagicTransit, RuleNewResponsePhaseMagicTransitIDsManaged, RuleNewResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

// A ruleset object.
type RuleDeleteResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RuleDeleteResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RuleDeleteResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules ResponeRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string                 `json:"description"`
	JSON        ruleDeleteResponseJSON `json:"-"`
}

// ruleDeleteResponseJSON contains the JSON metadata for the struct
// [RuleDeleteResponse]
type ruleDeleteResponseJSON struct {
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

func (r *RuleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RuleDeleteResponseKind string

const (
	RuleDeleteResponseKindManaged RuleDeleteResponseKind = "managed"
	RuleDeleteResponseKindCustom  RuleDeleteResponseKind = "custom"
	RuleDeleteResponseKindRoot    RuleDeleteResponseKind = "root"
	RuleDeleteResponseKindZone    RuleDeleteResponseKind = "zone"
)

func (r RuleDeleteResponseKind) IsKnown() bool {
	switch r {
	case RuleDeleteResponseKindManaged, RuleDeleteResponseKindCustom, RuleDeleteResponseKindRoot, RuleDeleteResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type RuleDeleteResponsePhase string

const (
	RuleDeleteResponsePhaseDDoSL4                         RuleDeleteResponsePhase = "ddos_l4"
	RuleDeleteResponsePhaseDDoSL7                         RuleDeleteResponsePhase = "ddos_l7"
	RuleDeleteResponsePhaseHTTPConfigSettings             RuleDeleteResponsePhase = "http_config_settings"
	RuleDeleteResponsePhaseHTTPCustomErrors               RuleDeleteResponsePhase = "http_custom_errors"
	RuleDeleteResponsePhaseHTTPLogCustomFields            RuleDeleteResponsePhase = "http_log_custom_fields"
	RuleDeleteResponsePhaseHTTPRatelimit                  RuleDeleteResponsePhase = "http_ratelimit"
	RuleDeleteResponsePhaseHTTPRequestCacheSettings       RuleDeleteResponsePhase = "http_request_cache_settings"
	RuleDeleteResponsePhaseHTTPRequestDynamicRedirect     RuleDeleteResponsePhase = "http_request_dynamic_redirect"
	RuleDeleteResponsePhaseHTTPRequestFirewallCustom      RuleDeleteResponsePhase = "http_request_firewall_custom"
	RuleDeleteResponsePhaseHTTPRequestFirewallManaged     RuleDeleteResponsePhase = "http_request_firewall_managed"
	RuleDeleteResponsePhaseHTTPRequestLateTransform       RuleDeleteResponsePhase = "http_request_late_transform"
	RuleDeleteResponsePhaseHTTPRequestOrigin              RuleDeleteResponsePhase = "http_request_origin"
	RuleDeleteResponsePhaseHTTPRequestRedirect            RuleDeleteResponsePhase = "http_request_redirect"
	RuleDeleteResponsePhaseHTTPRequestSanitize            RuleDeleteResponsePhase = "http_request_sanitize"
	RuleDeleteResponsePhaseHTTPRequestSbfm                RuleDeleteResponsePhase = "http_request_sbfm"
	RuleDeleteResponsePhaseHTTPRequestSelectConfiguration RuleDeleteResponsePhase = "http_request_select_configuration"
	RuleDeleteResponsePhaseHTTPRequestTransform           RuleDeleteResponsePhase = "http_request_transform"
	RuleDeleteResponsePhaseHTTPResponseCompression        RuleDeleteResponsePhase = "http_response_compression"
	RuleDeleteResponsePhaseHTTPResponseFirewallManaged    RuleDeleteResponsePhase = "http_response_firewall_managed"
	RuleDeleteResponsePhaseHTTPResponseHeadersTransform   RuleDeleteResponsePhase = "http_response_headers_transform"
	RuleDeleteResponsePhaseMagicTransit                   RuleDeleteResponsePhase = "magic_transit"
	RuleDeleteResponsePhaseMagicTransitIDsManaged         RuleDeleteResponsePhase = "magic_transit_ids_managed"
	RuleDeleteResponsePhaseMagicTransitManaged            RuleDeleteResponsePhase = "magic_transit_managed"
)

func (r RuleDeleteResponsePhase) IsKnown() bool {
	switch r {
	case RuleDeleteResponsePhaseDDoSL4, RuleDeleteResponsePhaseDDoSL7, RuleDeleteResponsePhaseHTTPConfigSettings, RuleDeleteResponsePhaseHTTPCustomErrors, RuleDeleteResponsePhaseHTTPLogCustomFields, RuleDeleteResponsePhaseHTTPRatelimit, RuleDeleteResponsePhaseHTTPRequestCacheSettings, RuleDeleteResponsePhaseHTTPRequestDynamicRedirect, RuleDeleteResponsePhaseHTTPRequestFirewallCustom, RuleDeleteResponsePhaseHTTPRequestFirewallManaged, RuleDeleteResponsePhaseHTTPRequestLateTransform, RuleDeleteResponsePhaseHTTPRequestOrigin, RuleDeleteResponsePhaseHTTPRequestRedirect, RuleDeleteResponsePhaseHTTPRequestSanitize, RuleDeleteResponsePhaseHTTPRequestSbfm, RuleDeleteResponsePhaseHTTPRequestSelectConfiguration, RuleDeleteResponsePhaseHTTPRequestTransform, RuleDeleteResponsePhaseHTTPResponseCompression, RuleDeleteResponsePhaseHTTPResponseFirewallManaged, RuleDeleteResponsePhaseHTTPResponseHeadersTransform, RuleDeleteResponsePhaseMagicTransit, RuleDeleteResponsePhaseMagicTransitIDsManaged, RuleDeleteResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

// A ruleset object.
type RuleEditResponse struct {
	// The unique ID of the ruleset.
	ID string `json:"id,required"`
	// The kind of the ruleset.
	Kind RuleEditResponseKind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase RuleEditResponsePhase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules ResponeRule `json:"rules,required"`
	// The version of the ruleset.
	Version string `json:"version,required"`
	// An informative description of the ruleset.
	Description string               `json:"description"`
	JSON        ruleEditResponseJSON `json:"-"`
}

// ruleEditResponseJSON contains the JSON metadata for the struct
// [RuleEditResponse]
type ruleEditResponseJSON struct {
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

func (r *RuleEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseJSON) RawJSON() string {
	return r.raw
}

// The kind of the ruleset.
type RuleEditResponseKind string

const (
	RuleEditResponseKindManaged RuleEditResponseKind = "managed"
	RuleEditResponseKindCustom  RuleEditResponseKind = "custom"
	RuleEditResponseKindRoot    RuleEditResponseKind = "root"
	RuleEditResponseKindZone    RuleEditResponseKind = "zone"
)

func (r RuleEditResponseKind) IsKnown() bool {
	switch r {
	case RuleEditResponseKindManaged, RuleEditResponseKindCustom, RuleEditResponseKindRoot, RuleEditResponseKindZone:
		return true
	}
	return false
}

// The phase of the ruleset.
type RuleEditResponsePhase string

const (
	RuleEditResponsePhaseDDoSL4                         RuleEditResponsePhase = "ddos_l4"
	RuleEditResponsePhaseDDoSL7                         RuleEditResponsePhase = "ddos_l7"
	RuleEditResponsePhaseHTTPConfigSettings             RuleEditResponsePhase = "http_config_settings"
	RuleEditResponsePhaseHTTPCustomErrors               RuleEditResponsePhase = "http_custom_errors"
	RuleEditResponsePhaseHTTPLogCustomFields            RuleEditResponsePhase = "http_log_custom_fields"
	RuleEditResponsePhaseHTTPRatelimit                  RuleEditResponsePhase = "http_ratelimit"
	RuleEditResponsePhaseHTTPRequestCacheSettings       RuleEditResponsePhase = "http_request_cache_settings"
	RuleEditResponsePhaseHTTPRequestDynamicRedirect     RuleEditResponsePhase = "http_request_dynamic_redirect"
	RuleEditResponsePhaseHTTPRequestFirewallCustom      RuleEditResponsePhase = "http_request_firewall_custom"
	RuleEditResponsePhaseHTTPRequestFirewallManaged     RuleEditResponsePhase = "http_request_firewall_managed"
	RuleEditResponsePhaseHTTPRequestLateTransform       RuleEditResponsePhase = "http_request_late_transform"
	RuleEditResponsePhaseHTTPRequestOrigin              RuleEditResponsePhase = "http_request_origin"
	RuleEditResponsePhaseHTTPRequestRedirect            RuleEditResponsePhase = "http_request_redirect"
	RuleEditResponsePhaseHTTPRequestSanitize            RuleEditResponsePhase = "http_request_sanitize"
	RuleEditResponsePhaseHTTPRequestSbfm                RuleEditResponsePhase = "http_request_sbfm"
	RuleEditResponsePhaseHTTPRequestSelectConfiguration RuleEditResponsePhase = "http_request_select_configuration"
	RuleEditResponsePhaseHTTPRequestTransform           RuleEditResponsePhase = "http_request_transform"
	RuleEditResponsePhaseHTTPResponseCompression        RuleEditResponsePhase = "http_response_compression"
	RuleEditResponsePhaseHTTPResponseFirewallManaged    RuleEditResponsePhase = "http_response_firewall_managed"
	RuleEditResponsePhaseHTTPResponseHeadersTransform   RuleEditResponsePhase = "http_response_headers_transform"
	RuleEditResponsePhaseMagicTransit                   RuleEditResponsePhase = "magic_transit"
	RuleEditResponsePhaseMagicTransitIDsManaged         RuleEditResponsePhase = "magic_transit_ids_managed"
	RuleEditResponsePhaseMagicTransitManaged            RuleEditResponsePhase = "magic_transit_managed"
)

func (r RuleEditResponsePhase) IsKnown() bool {
	switch r {
	case RuleEditResponsePhaseDDoSL4, RuleEditResponsePhaseDDoSL7, RuleEditResponsePhaseHTTPConfigSettings, RuleEditResponsePhaseHTTPCustomErrors, RuleEditResponsePhaseHTTPLogCustomFields, RuleEditResponsePhaseHTTPRatelimit, RuleEditResponsePhaseHTTPRequestCacheSettings, RuleEditResponsePhaseHTTPRequestDynamicRedirect, RuleEditResponsePhaseHTTPRequestFirewallCustom, RuleEditResponsePhaseHTTPRequestFirewallManaged, RuleEditResponsePhaseHTTPRequestLateTransform, RuleEditResponsePhaseHTTPRequestOrigin, RuleEditResponsePhaseHTTPRequestRedirect, RuleEditResponsePhaseHTTPRequestSanitize, RuleEditResponsePhaseHTTPRequestSbfm, RuleEditResponsePhaseHTTPRequestSelectConfiguration, RuleEditResponsePhaseHTTPRequestTransform, RuleEditResponsePhaseHTTPResponseCompression, RuleEditResponsePhaseHTTPResponseFirewallManaged, RuleEditResponsePhaseHTTPResponseHeadersTransform, RuleEditResponsePhaseMagicTransit, RuleEditResponsePhaseMagicTransitIDsManaged, RuleEditResponsePhaseMagicTransitManaged:
		return true
	}
	return false
}

type RuleNewParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An object configuring where the rule will be placed.
	Position param.Field[RuleNewParamsPositionUnion] `json:"position"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
type RuleNewParamsPosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsPosition) implementsRulesetsRuleNewParamsPositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleNewParamsPositionBeforePosition],
// [rulesets.RuleNewParamsPositionAfterPosition],
// [rulesets.RuleNewParamsPositionIndexPosition], [RuleNewParamsPosition].
type RuleNewParamsPositionUnion interface {
	implementsRulesetsRuleNewParamsPositionUnion()
}

// An object configuring where the rule will be placed.
type RuleNewParamsPositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleNewParamsPositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsPositionBeforePosition) implementsRulesetsRuleNewParamsPositionUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsPositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleNewParamsPositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsPositionAfterPosition) implementsRulesetsRuleNewParamsPositionUnion() {}

// An object configuring where the rule will be placed.
type RuleNewParamsPositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleNewParamsPositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleNewParamsPositionIndexPosition) implementsRulesetsRuleNewParamsPositionUnion() {}

// A response object.
type RuleNewResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleNewResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleNewResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RuleNewResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RuleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleNewResponseEnvelopeJSON    `json:"-"`
}

// ruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelope]
type ruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleNewResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleNewResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeErrors]
type ruleNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleNewResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                  `json:"pointer,required"`
	JSON    ruleNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleNewResponseEnvelopeErrorsSource]
type ruleNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleNewResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleNewResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleNewResponseEnvelopeMessages]
type ruleNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleNewResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                    `json:"pointer,required"`
	JSON    ruleNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleNewResponseEnvelopeMessagesSource]
type ruleNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleNewResponseEnvelopeSuccess bool

const (
	RuleNewResponseEnvelopeSuccessTrue RuleNewResponseEnvelopeSuccess = true
)

func (r RuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

// A response object.
type RuleDeleteResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleDeleteResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RuleDeleteResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelope]
type ruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleDeleteResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeErrors]
type ruleDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleDeleteResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    ruleDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleDeleteResponseEnvelopeErrorsSource]
type ruleDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleDeleteResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelopeMessages]
type ruleDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleDeleteResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                       `json:"pointer,required"`
	JSON    ruleDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleDeleteResponseEnvelopeMessagesSource]
type ruleDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleDeleteResponseEnvelopeSuccess bool

const (
	RuleDeleteResponseEnvelopeSuccessTrue RuleDeleteResponseEnvelopeSuccess = true
)

func (r RuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type RuleEditParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// An object configuring where the rule will be placed.
	Position param.Field[RuleEditParamsPositionUnion] `json:"position"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// An object configuring where the rule will be placed.
type RuleEditParamsPosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPosition) implementsRulesetsRuleEditParamsPositionUnion() {}

// An object configuring where the rule will be placed.
//
// Satisfied by [rulesets.RuleEditParamsPositionBeforePosition],
// [rulesets.RuleEditParamsPositionAfterPosition],
// [rulesets.RuleEditParamsPositionIndexPosition], [RuleEditParamsPosition].
type RuleEditParamsPositionUnion interface {
	implementsRulesetsRuleEditParamsPositionUnion()
}

// An object configuring where the rule will be placed.
type RuleEditParamsPositionBeforePosition struct {
	// The ID of another rule to place the rule before. An empty value causes the rule
	// to be placed at the top.
	Before param.Field[string] `json:"before"`
}

func (r RuleEditParamsPositionBeforePosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionBeforePosition) implementsRulesetsRuleEditParamsPositionUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsPositionAfterPosition struct {
	// The ID of another rule to place the rule after. An empty value causes the rule
	// to be placed at the bottom.
	After param.Field[string] `json:"after"`
}

func (r RuleEditParamsPositionAfterPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionAfterPosition) implementsRulesetsRuleEditParamsPositionUnion() {}

// An object configuring where the rule will be placed.
type RuleEditParamsPositionIndexPosition struct {
	// An index at which to place the rule, where index 1 is the first rule.
	Index param.Field[float64] `json:"index"`
}

func (r RuleEditParamsPositionIndexPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionIndexPosition) implementsRulesetsRuleEditParamsPositionUnion() {}

// A response object.
type RuleEditResponseEnvelope struct {
	// A list of error messages.
	Errors []RuleEditResponseEnvelopeErrors `json:"errors,required"`
	// A list of warning messages.
	Messages []RuleEditResponseEnvelopeMessages `json:"messages,required"`
	// A ruleset object.
	Result RuleEditResponse `json:"result,required"`
	// Whether the API call was successful.
	Success RuleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleEditResponseEnvelopeJSON    `json:"-"`
}

// ruleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelope]
type ruleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleEditResponseEnvelopeErrors struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleEditResponseEnvelopeErrorsSource `json:"source"`
	JSON   ruleEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// ruleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeErrors]
type ruleEditResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleEditResponseEnvelopeErrorsSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                   `json:"pointer,required"`
	JSON    ruleEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// ruleEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [RuleEditResponseEnvelopeErrorsSource]
type ruleEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

// A message.
type RuleEditResponseEnvelopeMessages struct {
	// A text description of this message.
	Message string `json:"message,required"`
	// A unique code for this message.
	Code int64 `json:"code"`
	// The source of this message.
	Source RuleEditResponseEnvelopeMessagesSource `json:"source"`
	JSON   ruleEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// ruleEditResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [RuleEditResponseEnvelopeMessages]
type ruleEditResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// The source of this message.
type RuleEditResponseEnvelopeMessagesSource struct {
	// A JSON pointer to the field that is the source of the message.
	Pointer string                                     `json:"pointer,required"`
	JSON    ruleEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// ruleEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [RuleEditResponseEnvelopeMessagesSource]
type ruleEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleEditResponseEnvelopeSuccess bool

const (
	RuleEditResponseEnvelopeSuccessTrue RuleEditResponseEnvelopeSuccess = true
)

func (r RuleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
