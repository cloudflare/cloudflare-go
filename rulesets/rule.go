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

// RuleService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRuleService] method instead.
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
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
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
	if body.AccountID.Value != "" && body.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if body.AccountID.Value == "" && body.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if body.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	}
	if body.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
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
	if params.AccountID.Value != "" && params.ZoneID.Value != "" {
		err = errors.New("account ID and zone ID are mutually exclusive")
		return
	}
	if params.AccountID.Value == "" && params.ZoneID.Value == "" {
		err = errors.New("either account ID or zone ID must be provided")
		return
	}
	if params.AccountID.Value != "" {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	}
	if params.ZoneID.Value != "" {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	if rulesetID == "" {
		err = errors.New("missing required ruleset_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
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

func (r BlockRule) implementsRulesetsRulesetNewResponseRule() {}

func (r BlockRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r BlockRule) implementsRulesetsRulesetGetResponseRule() {}

func (r BlockRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r BlockRule) implementsRulesetsPhaseGetResponseRule() {}

func (r BlockRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r BlockRule) implementsRulesetsRuleNewResponseRule() {}

func (r BlockRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r BlockRule) implementsRulesetsRuleEditResponseRule() {}

func (r BlockRule) implementsRulesetsVersionGetResponseRule() {}

func (r BlockRule) implementsRulesetsVersionByTagGetResponseRule() {}

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

func (r BlockRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r BlockRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r BlockRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

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

type ChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action ChallengeRuleAction `json:"action"`
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
	Ref  string            `json:"ref"`
	JSON challengeRuleJSON `json:"-"`
}

// challengeRuleJSON contains the JSON metadata for the struct [ChallengeRule]
type challengeRuleJSON struct {
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

func (r *ChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r challengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r ChallengeRule) implementsRulesetsRulesetNewResponseRule() {}

func (r ChallengeRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r ChallengeRule) implementsRulesetsRulesetGetResponseRule() {}

func (r ChallengeRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r ChallengeRule) implementsRulesetsPhaseGetResponseRule() {}

func (r ChallengeRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r ChallengeRule) implementsRulesetsRuleNewResponseRule() {}

func (r ChallengeRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r ChallengeRule) implementsRulesetsRuleEditResponseRule() {}

func (r ChallengeRule) implementsRulesetsVersionGetResponseRule() {}

func (r ChallengeRule) implementsRulesetsVersionByTagGetResponseRule() {}

// The action to perform when the rule matches.
type ChallengeRuleAction string

const (
	ChallengeRuleActionChallenge ChallengeRuleAction = "challenge"
)

func (r ChallengeRuleAction) IsKnown() bool {
	switch r {
	case ChallengeRuleActionChallenge:
		return true
	}
	return false
}

type ChallengeRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[ChallengeRuleAction] `json:"action"`
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

func (r ChallengeRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ChallengeRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r ChallengeRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r ChallengeRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

type CompressResponseRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action CompressResponseRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters CompressResponseRuleActionParameters `json:"action_parameters"`
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
	Ref  string                   `json:"ref"`
	JSON compressResponseRuleJSON `json:"-"`
}

// compressResponseRuleJSON contains the JSON metadata for the struct
// [CompressResponseRule]
type compressResponseRuleJSON struct {
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

func (r *CompressResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compressResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r CompressResponseRule) implementsRulesetsRulesetNewResponseRule() {}

func (r CompressResponseRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r CompressResponseRule) implementsRulesetsRulesetGetResponseRule() {}

func (r CompressResponseRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r CompressResponseRule) implementsRulesetsPhaseGetResponseRule() {}

func (r CompressResponseRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r CompressResponseRule) implementsRulesetsRuleNewResponseRule() {}

func (r CompressResponseRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r CompressResponseRule) implementsRulesetsRuleEditResponseRule() {}

func (r CompressResponseRule) implementsRulesetsVersionGetResponseRule() {}

func (r CompressResponseRule) implementsRulesetsVersionByTagGetResponseRule() {}

// The action to perform when the rule matches.
type CompressResponseRuleAction string

const (
	CompressResponseRuleActionCompressResponse CompressResponseRuleAction = "compress_response"
)

func (r CompressResponseRuleAction) IsKnown() bool {
	switch r {
	case CompressResponseRuleActionCompressResponse:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type CompressResponseRuleActionParameters struct {
	// Custom order for compression algorithms.
	Algorithms []CompressResponseRuleActionParametersAlgorithm `json:"algorithms"`
	JSON       compressResponseRuleActionParametersJSON        `json:"-"`
}

// compressResponseRuleActionParametersJSON contains the JSON metadata for the
// struct [CompressResponseRuleActionParameters]
type compressResponseRuleActionParametersJSON struct {
	Algorithms  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompressResponseRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compressResponseRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Compression algorithm to enable.
type CompressResponseRuleActionParametersAlgorithm struct {
	// Name of compression algorithm to enable.
	Name CompressResponseRuleActionParametersAlgorithmsName `json:"name"`
	JSON compressResponseRuleActionParametersAlgorithmJSON  `json:"-"`
}

// compressResponseRuleActionParametersAlgorithmJSON contains the JSON metadata for
// the struct [CompressResponseRuleActionParametersAlgorithm]
type compressResponseRuleActionParametersAlgorithmJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *CompressResponseRuleActionParametersAlgorithm) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r compressResponseRuleActionParametersAlgorithmJSON) RawJSON() string {
	return r.raw
}

// Name of compression algorithm to enable.
type CompressResponseRuleActionParametersAlgorithmsName string

const (
	CompressResponseRuleActionParametersAlgorithmsNameNone    CompressResponseRuleActionParametersAlgorithmsName = "none"
	CompressResponseRuleActionParametersAlgorithmsNameAuto    CompressResponseRuleActionParametersAlgorithmsName = "auto"
	CompressResponseRuleActionParametersAlgorithmsNameDefault CompressResponseRuleActionParametersAlgorithmsName = "default"
	CompressResponseRuleActionParametersAlgorithmsNameGzip    CompressResponseRuleActionParametersAlgorithmsName = "gzip"
	CompressResponseRuleActionParametersAlgorithmsNameBrotli  CompressResponseRuleActionParametersAlgorithmsName = "brotli"
)

func (r CompressResponseRuleActionParametersAlgorithmsName) IsKnown() bool {
	switch r {
	case CompressResponseRuleActionParametersAlgorithmsNameNone, CompressResponseRuleActionParametersAlgorithmsNameAuto, CompressResponseRuleActionParametersAlgorithmsNameDefault, CompressResponseRuleActionParametersAlgorithmsNameGzip, CompressResponseRuleActionParametersAlgorithmsNameBrotli:
		return true
	}
	return false
}

type CompressResponseRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[CompressResponseRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[CompressResponseRuleActionParametersParam] `json:"action_parameters"`
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

func (r CompressResponseRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r CompressResponseRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r CompressResponseRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r CompressResponseRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type CompressResponseRuleActionParametersParam struct {
	// Custom order for compression algorithms.
	Algorithms param.Field[[]CompressResponseRuleActionParametersAlgorithmParam] `json:"algorithms"`
}

func (r CompressResponseRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Compression algorithm to enable.
type CompressResponseRuleActionParametersAlgorithmParam struct {
	// Name of compression algorithm to enable.
	Name param.Field[CompressResponseRuleActionParametersAlgorithmsName] `json:"name"`
}

func (r CompressResponseRuleActionParametersAlgorithmParam) MarshalJSON() (data []byte, err error) {
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

func (r ExecuteRule) implementsRulesetsRulesetNewResponseRule() {}

func (r ExecuteRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r ExecuteRule) implementsRulesetsRulesetGetResponseRule() {}

func (r ExecuteRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r ExecuteRule) implementsRulesetsPhaseGetResponseRule() {}

func (r ExecuteRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r ExecuteRule) implementsRulesetsRuleNewResponseRule() {}

func (r ExecuteRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r ExecuteRule) implementsRulesetsRuleEditResponseRule() {}

func (r ExecuteRule) implementsRulesetsVersionGetResponseRule() {}

func (r ExecuteRule) implementsRulesetsVersionByTagGetResponseRule() {}

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

func (r ExecuteRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r ExecuteRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r ExecuteRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

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

type JSChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action JSChallengeRuleAction `json:"action"`
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
	Ref  string              `json:"ref"`
	JSON jsChallengeRuleJSON `json:"-"`
}

// jsChallengeRuleJSON contains the JSON metadata for the struct [JSChallengeRule]
type jsChallengeRuleJSON struct {
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

func (r *JSChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r jsChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r JSChallengeRule) implementsRulesetsRulesetNewResponseRule() {}

func (r JSChallengeRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r JSChallengeRule) implementsRulesetsRulesetGetResponseRule() {}

func (r JSChallengeRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r JSChallengeRule) implementsRulesetsPhaseGetResponseRule() {}

func (r JSChallengeRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r JSChallengeRule) implementsRulesetsRuleNewResponseRule() {}

func (r JSChallengeRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r JSChallengeRule) implementsRulesetsRuleEditResponseRule() {}

func (r JSChallengeRule) implementsRulesetsVersionGetResponseRule() {}

func (r JSChallengeRule) implementsRulesetsVersionByTagGetResponseRule() {}

// The action to perform when the rule matches.
type JSChallengeRuleAction string

const (
	JSChallengeRuleActionJSChallenge JSChallengeRuleAction = "js_challenge"
)

func (r JSChallengeRuleAction) IsKnown() bool {
	switch r {
	case JSChallengeRuleActionJSChallenge:
		return true
	}
	return false
}

type JSChallengeRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[JSChallengeRuleAction] `json:"action"`
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

func (r JSChallengeRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r JSChallengeRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r JSChallengeRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r JSChallengeRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

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

func (r LogRule) implementsRulesetsRulesetNewResponseRule() {}

func (r LogRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r LogRule) implementsRulesetsRulesetGetResponseRule() {}

func (r LogRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r LogRule) implementsRulesetsPhaseGetResponseRule() {}

func (r LogRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r LogRule) implementsRulesetsRuleNewResponseRule() {}

func (r LogRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r LogRule) implementsRulesetsRuleEditResponseRule() {}

func (r LogRule) implementsRulesetsVersionGetResponseRule() {}

func (r LogRule) implementsRulesetsVersionByTagGetResponseRule() {}

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

func (r LogRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r LogRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r LogRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

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

type ManagedChallengeRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action ManagedChallengeRuleAction `json:"action"`
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
	Ref  string                   `json:"ref"`
	JSON managedChallengeRuleJSON `json:"-"`
}

// managedChallengeRuleJSON contains the JSON metadata for the struct
// [ManagedChallengeRule]
type managedChallengeRuleJSON struct {
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

func (r *ManagedChallengeRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r managedChallengeRuleJSON) RawJSON() string {
	return r.raw
}

func (r ManagedChallengeRule) implementsRulesetsRulesetNewResponseRule() {}

func (r ManagedChallengeRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r ManagedChallengeRule) implementsRulesetsRulesetGetResponseRule() {}

func (r ManagedChallengeRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r ManagedChallengeRule) implementsRulesetsPhaseGetResponseRule() {}

func (r ManagedChallengeRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r ManagedChallengeRule) implementsRulesetsRuleNewResponseRule() {}

func (r ManagedChallengeRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r ManagedChallengeRule) implementsRulesetsRuleEditResponseRule() {}

func (r ManagedChallengeRule) implementsRulesetsVersionGetResponseRule() {}

func (r ManagedChallengeRule) implementsRulesetsVersionByTagGetResponseRule() {}

// The action to perform when the rule matches.
type ManagedChallengeRuleAction string

const (
	ManagedChallengeRuleActionManagedChallenge ManagedChallengeRuleAction = "managed_challenge"
)

func (r ManagedChallengeRuleAction) IsKnown() bool {
	switch r {
	case ManagedChallengeRuleActionManagedChallenge:
		return true
	}
	return false
}

type ManagedChallengeRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[ManagedChallengeRuleAction] `json:"action"`
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

func (r ManagedChallengeRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ManagedChallengeRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r ManagedChallengeRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r ManagedChallengeRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

type RedirectRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RedirectRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RedirectRuleActionParameters `json:"action_parameters"`
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
	Ref  string           `json:"ref"`
	JSON redirectRuleJSON `json:"-"`
}

// redirectRuleJSON contains the JSON metadata for the struct [RedirectRule]
type redirectRuleJSON struct {
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

func (r *RedirectRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleJSON) RawJSON() string {
	return r.raw
}

func (r RedirectRule) implementsRulesetsRulesetNewResponseRule() {}

func (r RedirectRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r RedirectRule) implementsRulesetsRulesetGetResponseRule() {}

func (r RedirectRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r RedirectRule) implementsRulesetsPhaseGetResponseRule() {}

func (r RedirectRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r RedirectRule) implementsRulesetsRuleNewResponseRule() {}

func (r RedirectRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r RedirectRule) implementsRulesetsRuleEditResponseRule() {}

func (r RedirectRule) implementsRulesetsVersionGetResponseRule() {}

func (r RedirectRule) implementsRulesetsVersionByTagGetResponseRule() {}

// The action to perform when the rule matches.
type RedirectRuleAction string

const (
	RedirectRuleActionRedirect RedirectRuleAction = "redirect"
)

func (r RedirectRuleAction) IsKnown() bool {
	switch r {
	case RedirectRuleActionRedirect:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RedirectRuleActionParameters struct {
	// Serve a redirect based on a bulk list lookup.
	FromList RedirectRuleActionParametersFromList `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue RedirectRuleActionParametersFromValue `json:"from_value"`
	JSON      redirectRuleActionParametersJSON      `json:"-"`
}

// redirectRuleActionParametersJSON contains the JSON metadata for the struct
// [RedirectRuleActionParameters]
type redirectRuleActionParametersJSON struct {
	FromList    apijson.Field
	FromValue   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RedirectRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on a bulk list lookup.
type RedirectRuleActionParametersFromList struct {
	// Expression that evaluates to the list lookup key.
	Key string `json:"key"`
	// The name of the list to match against.
	Name string                                   `json:"name"`
	JSON redirectRuleActionParametersFromListJSON `json:"-"`
}

// redirectRuleActionParametersFromListJSON contains the JSON metadata for the
// struct [RedirectRuleActionParametersFromList]
type redirectRuleActionParametersFromListJSON struct {
	Key         apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RedirectRuleActionParametersFromList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleActionParametersFromListJSON) RawJSON() string {
	return r.raw
}

// Serve a redirect based on the request properties.
type RedirectRuleActionParametersFromValue struct {
	// Keep the query string of the original request.
	PreserveQueryString bool `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode RedirectRuleActionParametersFromValueStatusCode `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL RedirectRuleActionParametersFromValueTargetURL `json:"target_url"`
	JSON      redirectRuleActionParametersFromValueJSON      `json:"-"`
}

// redirectRuleActionParametersFromValueJSON contains the JSON metadata for the
// struct [RedirectRuleActionParametersFromValue]
type redirectRuleActionParametersFromValueJSON struct {
	PreserveQueryString apijson.Field
	StatusCode          apijson.Field
	TargetURL           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RedirectRuleActionParametersFromValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleActionParametersFromValueJSON) RawJSON() string {
	return r.raw
}

// The status code to be used for the redirect.
type RedirectRuleActionParametersFromValueStatusCode float64

const (
	RedirectRuleActionParametersFromValueStatusCode301 RedirectRuleActionParametersFromValueStatusCode = 301
	RedirectRuleActionParametersFromValueStatusCode302 RedirectRuleActionParametersFromValueStatusCode = 302
	RedirectRuleActionParametersFromValueStatusCode303 RedirectRuleActionParametersFromValueStatusCode = 303
	RedirectRuleActionParametersFromValueStatusCode307 RedirectRuleActionParametersFromValueStatusCode = 307
	RedirectRuleActionParametersFromValueStatusCode308 RedirectRuleActionParametersFromValueStatusCode = 308
)

func (r RedirectRuleActionParametersFromValueStatusCode) IsKnown() bool {
	switch r {
	case RedirectRuleActionParametersFromValueStatusCode301, RedirectRuleActionParametersFromValueStatusCode302, RedirectRuleActionParametersFromValueStatusCode303, RedirectRuleActionParametersFromValueStatusCode307, RedirectRuleActionParametersFromValueStatusCode308:
		return true
	}
	return false
}

// The URL to redirect the request to.
type RedirectRuleActionParametersFromValueTargetURL struct {
	// The URL to redirect the request to.
	Value string `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                             `json:"expression"`
	JSON       redirectRuleActionParametersFromValueTargetURLJSON `json:"-"`
	union      RedirectRuleActionParametersFromValueTargetURLUnion
}

// redirectRuleActionParametersFromValueTargetURLJSON contains the JSON metadata
// for the struct [RedirectRuleActionParametersFromValueTargetURL]
type redirectRuleActionParametersFromValueTargetURLJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r redirectRuleActionParametersFromValueTargetURLJSON) RawJSON() string {
	return r.raw
}

func (r *RedirectRuleActionParametersFromValueTargetURL) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RedirectRuleActionParametersFromValueTargetURL) AsUnion() RedirectRuleActionParametersFromValueTargetURLUnion {
	return r.union
}

// The URL to redirect the request to.
//
// Union satisfied by
// [rulesets.RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect] or
// [rulesets.RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect].
type RedirectRuleActionParametersFromValueTargetURLUnion interface {
	implementsRulesetsRedirectRuleActionParametersFromValueTargetURL()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RedirectRuleActionParametersFromValueTargetURLUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect{}),
		},
	)
}

type RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect struct {
	// The URL to redirect the request to.
	Value string                                                              `json:"value"`
	JSON  redirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON `json:"-"`
}

// redirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON contains the
// JSON metadata for the struct
// [RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect]
type redirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleActionParametersFromValueTargetURLStaticURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r RedirectRuleActionParametersFromValueTargetURLStaticURLRedirect) implementsRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression string                                                               `json:"expression"`
	JSON       redirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON `json:"-"`
}

// redirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON contains
// the JSON metadata for the struct
// [RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect]
type redirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r redirectRuleActionParametersFromValueTargetURLDynamicURLRedirectJSON) RawJSON() string {
	return r.raw
}

func (r RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirect) implementsRulesetsRedirectRuleActionParametersFromValueTargetURL() {
}

type RedirectRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RedirectRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RedirectRuleActionParametersParam] `json:"action_parameters"`
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

func (r RedirectRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RedirectRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r RedirectRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r RedirectRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type RedirectRuleActionParametersParam struct {
	// Serve a redirect based on a bulk list lookup.
	FromList param.Field[RedirectRuleActionParametersFromListParam] `json:"from_list"`
	// Serve a redirect based on the request properties.
	FromValue param.Field[RedirectRuleActionParametersFromValueParam] `json:"from_value"`
}

func (r RedirectRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Serve a redirect based on a bulk list lookup.
type RedirectRuleActionParametersFromListParam struct {
	// Expression that evaluates to the list lookup key.
	Key param.Field[string] `json:"key"`
	// The name of the list to match against.
	Name param.Field[string] `json:"name"`
}

func (r RedirectRuleActionParametersFromListParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Serve a redirect based on the request properties.
type RedirectRuleActionParametersFromValueParam struct {
	// Keep the query string of the original request.
	PreserveQueryString param.Field[bool] `json:"preserve_query_string"`
	// The status code to be used for the redirect.
	StatusCode param.Field[RedirectRuleActionParametersFromValueStatusCode] `json:"status_code"`
	// The URL to redirect the request to.
	TargetURL param.Field[RedirectRuleActionParametersFromValueTargetURLUnionParam] `json:"target_url"`
}

func (r RedirectRuleActionParametersFromValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The URL to redirect the request to.
type RedirectRuleActionParametersFromValueTargetURLParam struct {
	// The URL to redirect the request to.
	Value param.Field[string] `json:"value"`
	// An expression to evaluate to get the URL to redirect the request to.
	Expression param.Field[string] `json:"expression"`
}

func (r RedirectRuleActionParametersFromValueTargetURLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RedirectRuleActionParametersFromValueTargetURLParam) implementsRulesetsRedirectRuleActionParametersFromValueTargetURLUnionParam() {
}

// The URL to redirect the request to.
//
// Satisfied by
// [rulesets.RedirectRuleActionParametersFromValueTargetURLStaticURLRedirectParam],
// [rulesets.RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectParam],
// [RedirectRuleActionParametersFromValueTargetURLParam].
type RedirectRuleActionParametersFromValueTargetURLUnionParam interface {
	implementsRulesetsRedirectRuleActionParametersFromValueTargetURLUnionParam()
}

type RedirectRuleActionParametersFromValueTargetURLStaticURLRedirectParam struct {
	// The URL to redirect the request to.
	Value param.Field[string] `json:"value"`
}

func (r RedirectRuleActionParametersFromValueTargetURLStaticURLRedirectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RedirectRuleActionParametersFromValueTargetURLStaticURLRedirectParam) implementsRulesetsRedirectRuleActionParametersFromValueTargetURLUnionParam() {
}

type RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectParam struct {
	// An expression to evaluate to get the URL to redirect the request to.
	Expression param.Field[string] `json:"expression"`
}

func (r RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RedirectRuleActionParametersFromValueTargetURLDynamicURLRedirectParam) implementsRulesetsRedirectRuleActionParametersFromValueTargetURLUnionParam() {
}

type RewriteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RewriteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RewriteRuleActionParameters `json:"action_parameters"`
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
	JSON rewriteRuleJSON `json:"-"`
}

// rewriteRuleJSON contains the JSON metadata for the struct [RewriteRule]
type rewriteRuleJSON struct {
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

func (r *RewriteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleJSON) RawJSON() string {
	return r.raw
}

func (r RewriteRule) implementsRulesetsRulesetNewResponseRule() {}

func (r RewriteRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r RewriteRule) implementsRulesetsRulesetGetResponseRule() {}

func (r RewriteRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r RewriteRule) implementsRulesetsPhaseGetResponseRule() {}

func (r RewriteRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r RewriteRule) implementsRulesetsRuleNewResponseRule() {}

func (r RewriteRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r RewriteRule) implementsRulesetsRuleEditResponseRule() {}

func (r RewriteRule) implementsRulesetsVersionGetResponseRule() {}

func (r RewriteRule) implementsRulesetsVersionByTagGetResponseRule() {}

// The action to perform when the rule matches.
type RewriteRuleAction string

const (
	RewriteRuleActionRewrite RewriteRuleAction = "rewrite"
)

func (r RewriteRuleAction) IsKnown() bool {
	switch r {
	case RewriteRuleActionRewrite:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RewriteRuleActionParameters struct {
	// Map of request headers to modify.
	Headers map[string]RewriteRuleActionParametersHeader `json:"headers"`
	// URI to rewrite the request to.
	URI  RewriteRuleActionParametersURI  `json:"uri"`
	JSON rewriteRuleActionParametersJSON `json:"-"`
}

// rewriteRuleActionParametersJSON contains the JSON metadata for the struct
// [RewriteRuleActionParameters]
type rewriteRuleActionParametersJSON struct {
	Headers     apijson.Field
	URI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Remove the header from the request.
type RewriteRuleActionParametersHeader struct {
	Operation RewriteRuleActionParametersHeadersOperation `json:"operation,required"`
	// Static value for the header.
	Value string `json:"value"`
	// Expression for the header value.
	Expression string                                `json:"expression"`
	JSON       rewriteRuleActionParametersHeaderJSON `json:"-"`
	union      RewriteRuleActionParametersHeadersUnion
}

// rewriteRuleActionParametersHeaderJSON contains the JSON metadata for the struct
// [RewriteRuleActionParametersHeader]
type rewriteRuleActionParametersHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rewriteRuleActionParametersHeaderJSON) RawJSON() string {
	return r.raw
}

func (r *RewriteRuleActionParametersHeader) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RewriteRuleActionParametersHeader) AsUnion() RewriteRuleActionParametersHeadersUnion {
	return r.union
}

// Remove the header from the request.
//
// Union satisfied by [rulesets.RewriteRuleActionParametersHeadersRemoveHeader],
// [rulesets.RewriteRuleActionParametersHeadersStaticHeader] or
// [rulesets.RewriteRuleActionParametersHeadersDynamicHeader].
type RewriteRuleActionParametersHeadersUnion interface {
	implementsRulesetsRewriteRuleActionParametersHeader()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RewriteRuleActionParametersHeadersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RewriteRuleActionParametersHeadersRemoveHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RewriteRuleActionParametersHeadersStaticHeader{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RewriteRuleActionParametersHeadersDynamicHeader{}),
		},
	)
}

// Remove the header from the request.
type RewriteRuleActionParametersHeadersRemoveHeader struct {
	Operation RewriteRuleActionParametersHeadersRemoveHeaderOperation `json:"operation,required"`
	JSON      rewriteRuleActionParametersHeadersRemoveHeaderJSON      `json:"-"`
}

// rewriteRuleActionParametersHeadersRemoveHeaderJSON contains the JSON metadata
// for the struct [RewriteRuleActionParametersHeadersRemoveHeader]
type rewriteRuleActionParametersHeadersRemoveHeaderJSON struct {
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteRuleActionParametersHeadersRemoveHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleActionParametersHeadersRemoveHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RewriteRuleActionParametersHeadersRemoveHeader) implementsRulesetsRewriteRuleActionParametersHeader() {
}

type RewriteRuleActionParametersHeadersRemoveHeaderOperation string

const (
	RewriteRuleActionParametersHeadersRemoveHeaderOperationRemove RewriteRuleActionParametersHeadersRemoveHeaderOperation = "remove"
)

func (r RewriteRuleActionParametersHeadersRemoveHeaderOperation) IsKnown() bool {
	switch r {
	case RewriteRuleActionParametersHeadersRemoveHeaderOperationRemove:
		return true
	}
	return false
}

// Set a request header with a static value.
type RewriteRuleActionParametersHeadersStaticHeader struct {
	Operation RewriteRuleActionParametersHeadersStaticHeaderOperation `json:"operation,required"`
	// Static value for the header.
	Value string                                             `json:"value,required"`
	JSON  rewriteRuleActionParametersHeadersStaticHeaderJSON `json:"-"`
}

// rewriteRuleActionParametersHeadersStaticHeaderJSON contains the JSON metadata
// for the struct [RewriteRuleActionParametersHeadersStaticHeader]
type rewriteRuleActionParametersHeadersStaticHeaderJSON struct {
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteRuleActionParametersHeadersStaticHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleActionParametersHeadersStaticHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RewriteRuleActionParametersHeadersStaticHeader) implementsRulesetsRewriteRuleActionParametersHeader() {
}

type RewriteRuleActionParametersHeadersStaticHeaderOperation string

const (
	RewriteRuleActionParametersHeadersStaticHeaderOperationSet RewriteRuleActionParametersHeadersStaticHeaderOperation = "set"
)

func (r RewriteRuleActionParametersHeadersStaticHeaderOperation) IsKnown() bool {
	switch r {
	case RewriteRuleActionParametersHeadersStaticHeaderOperationSet:
		return true
	}
	return false
}

// Set a request header with a dynamic value.
type RewriteRuleActionParametersHeadersDynamicHeader struct {
	// Expression for the header value.
	Expression string                                                   `json:"expression,required"`
	Operation  RewriteRuleActionParametersHeadersDynamicHeaderOperation `json:"operation,required"`
	JSON       rewriteRuleActionParametersHeadersDynamicHeaderJSON      `json:"-"`
}

// rewriteRuleActionParametersHeadersDynamicHeaderJSON contains the JSON metadata
// for the struct [RewriteRuleActionParametersHeadersDynamicHeader]
type rewriteRuleActionParametersHeadersDynamicHeaderJSON struct {
	Expression  apijson.Field
	Operation   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteRuleActionParametersHeadersDynamicHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleActionParametersHeadersDynamicHeaderJSON) RawJSON() string {
	return r.raw
}

func (r RewriteRuleActionParametersHeadersDynamicHeader) implementsRulesetsRewriteRuleActionParametersHeader() {
}

type RewriteRuleActionParametersHeadersDynamicHeaderOperation string

const (
	RewriteRuleActionParametersHeadersDynamicHeaderOperationSet RewriteRuleActionParametersHeadersDynamicHeaderOperation = "set"
)

func (r RewriteRuleActionParametersHeadersDynamicHeaderOperation) IsKnown() bool {
	switch r {
	case RewriteRuleActionParametersHeadersDynamicHeaderOperationSet:
		return true
	}
	return false
}

type RewriteRuleActionParametersHeadersOperation string

const (
	RewriteRuleActionParametersHeadersOperationRemove RewriteRuleActionParametersHeadersOperation = "remove"
	RewriteRuleActionParametersHeadersOperationSet    RewriteRuleActionParametersHeadersOperation = "set"
)

func (r RewriteRuleActionParametersHeadersOperation) IsKnown() bool {
	switch r {
	case RewriteRuleActionParametersHeadersOperationRemove, RewriteRuleActionParametersHeadersOperationSet:
		return true
	}
	return false
}

// URI to rewrite the request to.
type RewriteRuleActionParametersURI struct {
	// Path portion rewrite.
	Path RewriteURIPart `json:"path"`
	// Query portion rewrite.
	Query RewriteURIPart                     `json:"query"`
	JSON  rewriteRuleActionParametersURIJSON `json:"-"`
}

// rewriteRuleActionParametersURIJSON contains the JSON metadata for the struct
// [RewriteRuleActionParametersURI]
type rewriteRuleActionParametersURIJSON struct {
	Path        apijson.Field
	Query       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteRuleActionParametersURI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteRuleActionParametersURIJSON) RawJSON() string {
	return r.raw
}

type RewriteRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RewriteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RewriteRuleActionParametersParam] `json:"action_parameters"`
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

func (r RewriteRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r RewriteRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r RewriteRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type RewriteRuleActionParametersParam struct {
	// Map of request headers to modify.
	Headers param.Field[map[string]RewriteRuleActionParametersHeadersUnionParam] `json:"headers"`
	// URI to rewrite the request to.
	URI param.Field[RewriteRuleActionParametersURIParam] `json:"uri"`
}

func (r RewriteRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Remove the header from the request.
type RewriteRuleActionParametersHeaderParam struct {
	Operation param.Field[RewriteRuleActionParametersHeadersOperation] `json:"operation,required"`
	// Static value for the header.
	Value param.Field[string] `json:"value"`
	// Expression for the header value.
	Expression param.Field[string] `json:"expression"`
}

func (r RewriteRuleActionParametersHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteRuleActionParametersHeaderParam) implementsRulesetsRewriteRuleActionParametersHeadersUnionParam() {
}

// Remove the header from the request.
//
// Satisfied by [rulesets.RewriteRuleActionParametersHeadersRemoveHeaderParam],
// [rulesets.RewriteRuleActionParametersHeadersStaticHeaderParam],
// [rulesets.RewriteRuleActionParametersHeadersDynamicHeaderParam],
// [RewriteRuleActionParametersHeaderParam].
type RewriteRuleActionParametersHeadersUnionParam interface {
	implementsRulesetsRewriteRuleActionParametersHeadersUnionParam()
}

// Remove the header from the request.
type RewriteRuleActionParametersHeadersRemoveHeaderParam struct {
	Operation param.Field[RewriteRuleActionParametersHeadersRemoveHeaderOperation] `json:"operation,required"`
}

func (r RewriteRuleActionParametersHeadersRemoveHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteRuleActionParametersHeadersRemoveHeaderParam) implementsRulesetsRewriteRuleActionParametersHeadersUnionParam() {
}

// Set a request header with a static value.
type RewriteRuleActionParametersHeadersStaticHeaderParam struct {
	Operation param.Field[RewriteRuleActionParametersHeadersStaticHeaderOperation] `json:"operation,required"`
	// Static value for the header.
	Value param.Field[string] `json:"value,required"`
}

func (r RewriteRuleActionParametersHeadersStaticHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteRuleActionParametersHeadersStaticHeaderParam) implementsRulesetsRewriteRuleActionParametersHeadersUnionParam() {
}

// Set a request header with a dynamic value.
type RewriteRuleActionParametersHeadersDynamicHeaderParam struct {
	// Expression for the header value.
	Expression param.Field[string]                                                   `json:"expression,required"`
	Operation  param.Field[RewriteRuleActionParametersHeadersDynamicHeaderOperation] `json:"operation,required"`
}

func (r RewriteRuleActionParametersHeadersDynamicHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteRuleActionParametersHeadersDynamicHeaderParam) implementsRulesetsRewriteRuleActionParametersHeadersUnionParam() {
}

// URI to rewrite the request to.
type RewriteRuleActionParametersURIParam struct {
	// Path portion rewrite.
	Path param.Field[RewriteURIPartUnionParam] `json:"path"`
	// Query portion rewrite.
	Query param.Field[RewriteURIPartUnionParam] `json:"query"`
}

func (r RewriteRuleActionParametersURIParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RewriteURIPart struct {
	// Predefined replacement value.
	Value string `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression string             `json:"expression"`
	JSON       rewriteURIPartJSON `json:"-"`
	union      RewriteURIPartUnion
}

// rewriteURIPartJSON contains the JSON metadata for the struct [RewriteURIPart]
type rewriteURIPartJSON struct {
	Value       apijson.Field
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r rewriteURIPartJSON) RawJSON() string {
	return r.raw
}

func (r *RewriteURIPart) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RewriteURIPart) AsUnion() RewriteURIPartUnion {
	return r.union
}

// Union satisfied by [rulesets.RewriteURIPartStaticValue] or
// [rulesets.RewriteURIPartDynamicValue].
type RewriteURIPartUnion interface {
	implementsRulesetsRewriteURIPart()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RewriteURIPartUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RewriteURIPartStaticValue{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(RewriteURIPartDynamicValue{}),
		},
	)
}

type RewriteURIPartStaticValue struct {
	// Predefined replacement value.
	Value string                        `json:"value,required"`
	JSON  rewriteURIPartStaticValueJSON `json:"-"`
}

// rewriteURIPartStaticValueJSON contains the JSON metadata for the struct
// [RewriteURIPartStaticValue]
type rewriteURIPartStaticValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteURIPartStaticValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteURIPartStaticValueJSON) RawJSON() string {
	return r.raw
}

func (r RewriteURIPartStaticValue) implementsRulesetsRewriteURIPart() {}

type RewriteURIPartDynamicValue struct {
	// Expression to evaluate for the replacement value.
	Expression string                         `json:"expression,required"`
	JSON       rewriteURIPartDynamicValueJSON `json:"-"`
}

// rewriteURIPartDynamicValueJSON contains the JSON metadata for the struct
// [RewriteURIPartDynamicValue]
type rewriteURIPartDynamicValueJSON struct {
	Expression  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RewriteURIPartDynamicValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r rewriteURIPartDynamicValueJSON) RawJSON() string {
	return r.raw
}

func (r RewriteURIPartDynamicValue) implementsRulesetsRewriteURIPart() {}

type RewriteURIPartParam struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value"`
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression"`
}

func (r RewriteURIPartParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteURIPartParam) implementsRulesetsRewriteURIPartUnionParam() {}

// Satisfied by [rulesets.RewriteURIPartStaticValueParam],
// [rulesets.RewriteURIPartDynamicValueParam], [RewriteURIPartParam].
type RewriteURIPartUnionParam interface {
	implementsRulesetsRewriteURIPartUnionParam()
}

type RewriteURIPartStaticValueParam struct {
	// Predefined replacement value.
	Value param.Field[string] `json:"value,required"`
}

func (r RewriteURIPartStaticValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteURIPartStaticValueParam) implementsRulesetsRewriteURIPartUnionParam() {}

type RewriteURIPartDynamicValueParam struct {
	// Expression to evaluate for the replacement value.
	Expression param.Field[string] `json:"expression,required"`
}

func (r RewriteURIPartDynamicValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RewriteURIPartDynamicValueParam) implementsRulesetsRewriteURIPartUnionParam() {}

type RouteRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RouteRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RouteRuleActionParameters `json:"action_parameters"`
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
	JSON routeRuleJSON `json:"-"`
}

// routeRuleJSON contains the JSON metadata for the struct [RouteRule]
type routeRuleJSON struct {
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

func (r *RouteRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeRuleJSON) RawJSON() string {
	return r.raw
}

func (r RouteRule) implementsRulesetsRulesetNewResponseRule() {}

func (r RouteRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r RouteRule) implementsRulesetsRulesetGetResponseRule() {}

func (r RouteRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r RouteRule) implementsRulesetsPhaseGetResponseRule() {}

func (r RouteRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r RouteRule) implementsRulesetsRuleNewResponseRule() {}

func (r RouteRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r RouteRule) implementsRulesetsRuleEditResponseRule() {}

func (r RouteRule) implementsRulesetsVersionGetResponseRule() {}

func (r RouteRule) implementsRulesetsVersionByTagGetResponseRule() {}

// The action to perform when the rule matches.
type RouteRuleAction string

const (
	RouteRuleActionRoute RouteRuleAction = "route"
)

func (r RouteRuleAction) IsKnown() bool {
	switch r {
	case RouteRuleActionRoute:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RouteRuleActionParameters struct {
	// Rewrite the HTTP Host header.
	HostHeader string `json:"host_header"`
	// Override the IP/TCP destination.
	Origin RouteRuleActionParametersOrigin `json:"origin"`
	// Override the Server Name Indication (SNI).
	SNI  RouteRuleActionParametersSNI  `json:"sni"`
	JSON routeRuleActionParametersJSON `json:"-"`
}

// routeRuleActionParametersJSON contains the JSON metadata for the struct
// [RouteRuleActionParameters]
type routeRuleActionParametersJSON struct {
	HostHeader  apijson.Field
	Origin      apijson.Field
	SNI         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Override the IP/TCP destination.
type RouteRuleActionParametersOrigin struct {
	// Override the resolved hostname.
	Host string `json:"host"`
	// Override the destination port.
	Port float64                             `json:"port"`
	JSON routeRuleActionParametersOriginJSON `json:"-"`
}

// routeRuleActionParametersOriginJSON contains the JSON metadata for the struct
// [RouteRuleActionParametersOrigin]
type routeRuleActionParametersOriginJSON struct {
	Host        apijson.Field
	Port        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteRuleActionParametersOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeRuleActionParametersOriginJSON) RawJSON() string {
	return r.raw
}

// Override the Server Name Indication (SNI).
type RouteRuleActionParametersSNI struct {
	// The SNI override.
	Value string                           `json:"value,required"`
	JSON  routeRuleActionParametersSNIJSON `json:"-"`
}

// routeRuleActionParametersSNIJSON contains the JSON metadata for the struct
// [RouteRuleActionParametersSNI]
type routeRuleActionParametersSNIJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteRuleActionParametersSNI) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeRuleActionParametersSNIJSON) RawJSON() string {
	return r.raw
}

type RouteRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[RouteRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[RouteRuleActionParametersParam] `json:"action_parameters"`
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

func (r RouteRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RouteRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r RouteRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r RouteRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type RouteRuleActionParametersParam struct {
	// Rewrite the HTTP Host header.
	HostHeader param.Field[string] `json:"host_header"`
	// Override the IP/TCP destination.
	Origin param.Field[RouteRuleActionParametersOriginParam] `json:"origin"`
	// Override the Server Name Indication (SNI).
	SNI param.Field[RouteRuleActionParametersSNIParam] `json:"sni"`
}

func (r RouteRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Override the IP/TCP destination.
type RouteRuleActionParametersOriginParam struct {
	// Override the resolved hostname.
	Host param.Field[string] `json:"host"`
	// Override the destination port.
	Port param.Field[float64] `json:"port"`
}

func (r RouteRuleActionParametersOriginParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Override the Server Name Indication (SNI).
type RouteRuleActionParametersSNIParam struct {
	// The SNI override.
	Value param.Field[string] `json:"value,required"`
}

func (r RouteRuleActionParametersSNIParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ScoreRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action ScoreRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters ScoreRuleActionParameters `json:"action_parameters"`
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
	JSON scoreRuleJSON `json:"-"`
}

// scoreRuleJSON contains the JSON metadata for the struct [ScoreRule]
type scoreRuleJSON struct {
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

func (r *ScoreRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scoreRuleJSON) RawJSON() string {
	return r.raw
}

func (r ScoreRule) implementsRulesetsRulesetNewResponseRule() {}

func (r ScoreRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r ScoreRule) implementsRulesetsRulesetGetResponseRule() {}

func (r ScoreRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r ScoreRule) implementsRulesetsPhaseGetResponseRule() {}

func (r ScoreRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r ScoreRule) implementsRulesetsRuleNewResponseRule() {}

func (r ScoreRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r ScoreRule) implementsRulesetsRuleEditResponseRule() {}

func (r ScoreRule) implementsRulesetsVersionGetResponseRule() {}

func (r ScoreRule) implementsRulesetsVersionByTagGetResponseRule() {}

// The action to perform when the rule matches.
type ScoreRuleAction string

const (
	ScoreRuleActionScore ScoreRuleAction = "score"
)

func (r ScoreRuleAction) IsKnown() bool {
	switch r {
	case ScoreRuleActionScore:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type ScoreRuleActionParameters struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment int64                         `json:"increment"`
	JSON      scoreRuleActionParametersJSON `json:"-"`
}

// scoreRuleActionParametersJSON contains the JSON metadata for the struct
// [ScoreRuleActionParameters]
type scoreRuleActionParametersJSON struct {
	Increment   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ScoreRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r scoreRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

type ScoreRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[ScoreRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[ScoreRuleActionParametersParam] `json:"action_parameters"`
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

func (r ScoreRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ScoreRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r ScoreRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r ScoreRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type ScoreRuleActionParametersParam struct {
	// Increment contains the delta to change the score and can be either positive or
	// negative.
	Increment param.Field[int64] `json:"increment"`
}

func (r ScoreRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ServeErrorRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action ServeErrorRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters ServeErrorRuleActionParameters `json:"action_parameters"`
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
	Ref  string             `json:"ref"`
	JSON serveErrorRuleJSON `json:"-"`
}

// serveErrorRuleJSON contains the JSON metadata for the struct [ServeErrorRule]
type serveErrorRuleJSON struct {
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

func (r *ServeErrorRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serveErrorRuleJSON) RawJSON() string {
	return r.raw
}

func (r ServeErrorRule) implementsRulesetsRulesetNewResponseRule() {}

func (r ServeErrorRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r ServeErrorRule) implementsRulesetsRulesetGetResponseRule() {}

func (r ServeErrorRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r ServeErrorRule) implementsRulesetsPhaseGetResponseRule() {}

func (r ServeErrorRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r ServeErrorRule) implementsRulesetsRuleNewResponseRule() {}

func (r ServeErrorRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r ServeErrorRule) implementsRulesetsRuleEditResponseRule() {}

func (r ServeErrorRule) implementsRulesetsVersionGetResponseRule() {}

func (r ServeErrorRule) implementsRulesetsVersionByTagGetResponseRule() {}

// The action to perform when the rule matches.
type ServeErrorRuleAction string

const (
	ServeErrorRuleActionServeError ServeErrorRuleAction = "serve_error"
)

func (r ServeErrorRuleAction) IsKnown() bool {
	switch r {
	case ServeErrorRuleActionServeError:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type ServeErrorRuleActionParameters struct {
	// Error response content.
	Content string `json:"content"`
	// Content-type header to set with the response.
	ContentType ServeErrorRuleActionParametersContentType `json:"content_type"`
	// The status code to use for the error.
	StatusCode float64                            `json:"status_code"`
	JSON       serveErrorRuleActionParametersJSON `json:"-"`
}

// serveErrorRuleActionParametersJSON contains the JSON metadata for the struct
// [ServeErrorRuleActionParameters]
type serveErrorRuleActionParametersJSON struct {
	Content     apijson.Field
	ContentType apijson.Field
	StatusCode  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ServeErrorRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r serveErrorRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Content-type header to set with the response.
type ServeErrorRuleActionParametersContentType string

const (
	ServeErrorRuleActionParametersContentTypeApplicationJson ServeErrorRuleActionParametersContentType = "application/json"
	ServeErrorRuleActionParametersContentTypeTextXml         ServeErrorRuleActionParametersContentType = "text/xml"
	ServeErrorRuleActionParametersContentTypeTextPlain       ServeErrorRuleActionParametersContentType = "text/plain"
	ServeErrorRuleActionParametersContentTypeTextHTML        ServeErrorRuleActionParametersContentType = "text/html"
)

func (r ServeErrorRuleActionParametersContentType) IsKnown() bool {
	switch r {
	case ServeErrorRuleActionParametersContentTypeApplicationJson, ServeErrorRuleActionParametersContentTypeTextXml, ServeErrorRuleActionParametersContentTypeTextPlain, ServeErrorRuleActionParametersContentTypeTextHTML:
		return true
	}
	return false
}

type ServeErrorRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[ServeErrorRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[ServeErrorRuleActionParametersParam] `json:"action_parameters"`
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

func (r ServeErrorRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ServeErrorRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r ServeErrorRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r ServeErrorRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type ServeErrorRuleActionParametersParam struct {
	// Error response content.
	Content param.Field[string] `json:"content"`
	// Content-type header to set with the response.
	ContentType param.Field[ServeErrorRuleActionParametersContentType] `json:"content_type"`
	// The status code to use for the error.
	StatusCode param.Field[float64] `json:"status_code"`
}

func (r ServeErrorRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SetCacheSettingsRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action SetCacheSettingsRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters SetCacheSettingsRuleActionParameters `json:"action_parameters"`
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
	Ref  string                   `json:"ref"`
	JSON setCacheSettingsRuleJSON `json:"-"`
}

// setCacheSettingsRuleJSON contains the JSON metadata for the struct
// [SetCacheSettingsRule]
type setCacheSettingsRuleJSON struct {
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

func (r *SetCacheSettingsRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleJSON) RawJSON() string {
	return r.raw
}

func (r SetCacheSettingsRule) implementsRulesetsRulesetNewResponseRule() {}

func (r SetCacheSettingsRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r SetCacheSettingsRule) implementsRulesetsRulesetGetResponseRule() {}

func (r SetCacheSettingsRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r SetCacheSettingsRule) implementsRulesetsPhaseGetResponseRule() {}

func (r SetCacheSettingsRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r SetCacheSettingsRule) implementsRulesetsRuleNewResponseRule() {}

func (r SetCacheSettingsRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r SetCacheSettingsRule) implementsRulesetsRuleEditResponseRule() {}

func (r SetCacheSettingsRule) implementsRulesetsVersionGetResponseRule() {}

func (r SetCacheSettingsRule) implementsRulesetsVersionByTagGetResponseRule() {}

// The action to perform when the rule matches.
type SetCacheSettingsRuleAction string

const (
	SetCacheSettingsRuleActionSetCacheSettings SetCacheSettingsRuleAction = "set_cache_settings"
)

func (r SetCacheSettingsRuleAction) IsKnown() bool {
	switch r {
	case SetCacheSettingsRuleActionSetCacheSettings:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type SetCacheSettingsRuleActionParameters struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts []int64 `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL SetCacheSettingsRuleActionParametersBrowserTTL `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache bool `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey SetCacheSettingsRuleActionParametersCacheKey `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve SetCacheSettingsRuleActionParametersCacheReserve `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL SetCacheSettingsRuleActionParametersEdgeTTL `json:"edge_ttl"`
	// When enabled, Cloudflare will aim to strictly adhere to RFC 7234.
	OriginCacheControl bool `json:"origin_cache_control"`
	// Generate Cloudflare error pages from issues sent from the origin server. When
	// on, error pages will trigger for issues from the origin
	OriginErrorPagePassthru bool `json:"origin_error_page_passthru"`
	// Define a timeout value between two successive read operations to your origin
	// server. Historically, the timeout value between two read options from Cloudflare
	// to an origin server is 100 seconds. If you are attempting to reduce HTTP 524
	// errors because of timeouts from an origin server, try increasing this timeout
	// value.
	ReadTimeout int64 `json:"read_timeout"`
	// Specify whether or not Cloudflare should respect strong ETag (entity tag)
	// headers. When off, Cloudflare converts strong ETag headers to weak ETag headers.
	RespectStrongEtags bool `json:"respect_strong_etags"`
	// Define if Cloudflare should serve stale content while getting the latest content
	// from the origin. If on, Cloudflare will not serve stale content while getting
	// the latest content from the origin.
	ServeStale SetCacheSettingsRuleActionParametersServeStale `json:"serve_stale"`
	JSON       setCacheSettingsRuleActionParametersJSON       `json:"-"`
}

// setCacheSettingsRuleActionParametersJSON contains the JSON metadata for the
// struct [SetCacheSettingsRuleActionParameters]
type setCacheSettingsRuleActionParametersJSON struct {
	AdditionalCacheablePorts apijson.Field
	BrowserTTL               apijson.Field
	Cache                    apijson.Field
	CacheKey                 apijson.Field
	CacheReserve             apijson.Field
	EdgeTTL                  apijson.Field
	OriginCacheControl       apijson.Field
	OriginErrorPagePassthru  apijson.Field
	ReadTimeout              apijson.Field
	RespectStrongEtags       apijson.Field
	ServeStale               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type SetCacheSettingsRuleActionParametersBrowserTTL struct {
	// Determines which browser ttl mode to use.
	Mode SetCacheSettingsRuleActionParametersBrowserTTLMode `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64                                              `json:"default"`
	JSON    setCacheSettingsRuleActionParametersBrowserTTLJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersBrowserTTLJSON contains the JSON metadata
// for the struct [SetCacheSettingsRuleActionParametersBrowserTTL]
type setCacheSettingsRuleActionParametersBrowserTTLJSON struct {
	Mode        apijson.Field
	Default     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersBrowserTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersBrowserTTLJSON) RawJSON() string {
	return r.raw
}

// Determines which browser ttl mode to use.
type SetCacheSettingsRuleActionParametersBrowserTTLMode string

const (
	SetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin   SetCacheSettingsRuleActionParametersBrowserTTLMode = "respect_origin"
	SetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault SetCacheSettingsRuleActionParametersBrowserTTLMode = "bypass_by_default"
	SetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin  SetCacheSettingsRuleActionParametersBrowserTTLMode = "override_origin"
)

func (r SetCacheSettingsRuleActionParametersBrowserTTLMode) IsKnown() bool {
	switch r {
	case SetCacheSettingsRuleActionParametersBrowserTTLModeRespectOrigin, SetCacheSettingsRuleActionParametersBrowserTTLModeBypassByDefault, SetCacheSettingsRuleActionParametersBrowserTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type SetCacheSettingsRuleActionParametersCacheKey struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType bool `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor bool `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey SetCacheSettingsRuleActionParametersCacheKeyCustomKey `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder bool                                             `json:"ignore_query_strings_order"`
	JSON                    setCacheSettingsRuleActionParametersCacheKeyJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyJSON contains the JSON metadata for
// the struct [SetCacheSettingsRuleActionParametersCacheKey]
type setCacheSettingsRuleActionParametersCacheKeyJSON struct {
	CacheByDeviceType       apijson.Field
	CacheDeceptionArmor     apijson.Field
	CustomKey               apijson.Field
	IgnoreQueryStringsOrder apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyJSON) RawJSON() string {
	return r.raw
}

// Customize which components of the request are included or excluded from the
// cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKey struct {
	// The cookies to include in building the cache key.
	Cookie SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser `json:"user"`
	JSON setCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON contains the JSON
// metadata for the struct [SetCacheSettingsRuleActionParametersCacheKeyCustomKey]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON struct {
	Cookie      apijson.Field
	Header      apijson.Field
	Host        apijson.Field
	QueryString apijson.Field
	User        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyJSON) RawJSON() string {
	return r.raw
}

// The cookies to include in building the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Include these cookies' names and their values.
	Include []string                                                        `json:"include"`
	JSON    setCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON contains the
// JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON struct {
	CheckPresence apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookie) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieJSON) RawJSON() string {
	return r.raw
}

// The header names and values to include in building the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence []string `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin bool `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include []string                                                        `json:"include"`
	JSON    setCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON contains the
// JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON struct {
	CheckPresence apijson.Field
	ExcludeOrigin apijson.Field
	Include       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeader) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderJSON) RawJSON() string {
	return r.raw
}

// Whether to use the original host or the resolved host in the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved bool                                                          `json:"resolved"`
	JSON     setCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON contains the JSON
// metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON struct {
	Resolved    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHost) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyHostJSON) RawJSON() string {
	return r.raw
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude `json:"include"`
	JSON    setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON    `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON contains
// the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryString) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringJSON) RawJSON() string {
	return r.raw
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude struct {
	// Exclude all query string parameters from use in building the cache key.
	All bool `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List []string                                                                    `json:"list"`
	JSON setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON
// contains the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeJSON) RawJSON() string {
	return r.raw
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude struct {
	// Use all query string parameters in the cache key.
	All bool `json:"all"`
	// A list of query string parameters used to build the cache key.
	List []string                                                                    `json:"list"`
	JSON setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON
// contains the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON struct {
	All         apijson.Field
	List        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeJSON) RawJSON() string {
	return r.raw
}

// Characteristics of the request user agent used in building the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser struct {
	// Use the user agent's device type in the cache key.
	DeviceType bool `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo bool `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang bool                                                          `json:"lang"`
	JSON setCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON contains the JSON
// metadata for the struct
// [SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser]
type setCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON struct {
	DeviceType  apijson.Field
	Geo         apijson.Field
	Lang        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheKeyCustomKeyUserJSON) RawJSON() string {
	return r.raw
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type SetCacheSettingsRuleActionParametersCacheReserve struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible bool `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize int64                                                `json:"min_file_size,required"`
	JSON        setCacheSettingsRuleActionParametersCacheReserveJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersCacheReserveJSON contains the JSON metadata
// for the struct [SetCacheSettingsRuleActionParametersCacheReserve]
type setCacheSettingsRuleActionParametersCacheReserveJSON struct {
	Eligible    apijson.Field
	MinFileSize apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersCacheReserve) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersCacheReserveJSON) RawJSON() string {
	return r.raw
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type SetCacheSettingsRuleActionParametersEdgeTTL struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default int64 `json:"default,required"`
	// edge ttl options
	Mode SetCacheSettingsRuleActionParametersEdgeTTLMode `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL []SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL `json:"status_code_ttl,required"`
	JSON          setCacheSettingsRuleActionParametersEdgeTTLJSON            `json:"-"`
}

// setCacheSettingsRuleActionParametersEdgeTTLJSON contains the JSON metadata for
// the struct [SetCacheSettingsRuleActionParametersEdgeTTL]
type setCacheSettingsRuleActionParametersEdgeTTLJSON struct {
	Default       apijson.Field
	Mode          apijson.Field
	StatusCodeTTL apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersEdgeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersEdgeTTLJSON) RawJSON() string {
	return r.raw
}

// edge ttl options
type SetCacheSettingsRuleActionParametersEdgeTTLMode string

const (
	SetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin   SetCacheSettingsRuleActionParametersEdgeTTLMode = "respect_origin"
	SetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault SetCacheSettingsRuleActionParametersEdgeTTLMode = "bypass_by_default"
	SetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin  SetCacheSettingsRuleActionParametersEdgeTTLMode = "override_origin"
)

func (r SetCacheSettingsRuleActionParametersEdgeTTLMode) IsKnown() bool {
	switch r {
	case SetCacheSettingsRuleActionParametersEdgeTTLModeRespectOrigin, SetCacheSettingsRuleActionParametersEdgeTTLModeBypassByDefault, SetCacheSettingsRuleActionParametersEdgeTTLModeOverrideOrigin:
		return true
	}
	return false
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value int64 `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue int64                                                        `json:"status_code_value"`
	JSON            setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON contains the JSON
// metadata for the struct
// [SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL]
type setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON struct {
	Value           apijson.Field
	StatusCodeRange apijson.Field
	StatusCodeValue apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTL) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLJSON) RawJSON() string {
	return r.raw
}

// The range of status codes used to apply the selected mode.
type SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange struct {
	// response status code lower bound
	From int64 `json:"from,required"`
	// response status code upper bound
	To   int64                                                                       `json:"to,required"`
	JSON setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON
// contains the JSON metadata for the struct
// [SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange]
type setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeJSON) RawJSON() string {
	return r.raw
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type SetCacheSettingsRuleActionParametersServeStale struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating bool                                               `json:"disable_stale_while_updating,required"`
	JSON                      setCacheSettingsRuleActionParametersServeStaleJSON `json:"-"`
}

// setCacheSettingsRuleActionParametersServeStaleJSON contains the JSON metadata
// for the struct [SetCacheSettingsRuleActionParametersServeStale]
type setCacheSettingsRuleActionParametersServeStaleJSON struct {
	DisableStaleWhileUpdating apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *SetCacheSettingsRuleActionParametersServeStale) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setCacheSettingsRuleActionParametersServeStaleJSON) RawJSON() string {
	return r.raw
}

type SetCacheSettingsRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[SetCacheSettingsRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[SetCacheSettingsRuleActionParametersParam] `json:"action_parameters"`
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

func (r SetCacheSettingsRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetCacheSettingsRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r SetCacheSettingsRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r SetCacheSettingsRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type SetCacheSettingsRuleActionParametersParam struct {
	// List of additional ports that caching can be enabled on.
	AdditionalCacheablePorts param.Field[[]int64] `json:"additional_cacheable_ports"`
	// Specify how long client browsers should cache the response. Cloudflare cache
	// purge will not purge content cached on client browsers, so high browser TTLs may
	// lead to stale content.
	BrowserTTL param.Field[SetCacheSettingsRuleActionParametersBrowserTTLParam] `json:"browser_ttl"`
	// Mark whether the request’s response from origin is eligible for caching. Caching
	// itself will still depend on the cache-control header and your other caching
	// configurations.
	Cache param.Field[bool] `json:"cache"`
	// Define which components of the request are included or excluded from the cache
	// key Cloudflare uses to store the response in cache.
	CacheKey param.Field[SetCacheSettingsRuleActionParametersCacheKeyParam] `json:"cache_key"`
	// Mark whether the request's response from origin is eligible for Cache Reserve
	// (requires a Cache Reserve add-on plan).
	CacheReserve param.Field[SetCacheSettingsRuleActionParametersCacheReserveParam] `json:"cache_reserve"`
	// TTL (Time to Live) specifies the maximum time to cache a resource in the
	// Cloudflare edge network.
	EdgeTTL param.Field[SetCacheSettingsRuleActionParametersEdgeTTLParam] `json:"edge_ttl"`
	// When enabled, Cloudflare will aim to strictly adhere to RFC 7234.
	OriginCacheControl param.Field[bool] `json:"origin_cache_control"`
	// Generate Cloudflare error pages from issues sent from the origin server. When
	// on, error pages will trigger for issues from the origin
	OriginErrorPagePassthru param.Field[bool] `json:"origin_error_page_passthru"`
	// Define a timeout value between two successive read operations to your origin
	// server. Historically, the timeout value between two read options from Cloudflare
	// to an origin server is 100 seconds. If you are attempting to reduce HTTP 524
	// errors because of timeouts from an origin server, try increasing this timeout
	// value.
	ReadTimeout param.Field[int64] `json:"read_timeout"`
	// Specify whether or not Cloudflare should respect strong ETag (entity tag)
	// headers. When off, Cloudflare converts strong ETag headers to weak ETag headers.
	RespectStrongEtags param.Field[bool] `json:"respect_strong_etags"`
	// Define if Cloudflare should serve stale content while getting the latest content
	// from the origin. If on, Cloudflare will not serve stale content while getting
	// the latest content from the origin.
	ServeStale param.Field[SetCacheSettingsRuleActionParametersServeStaleParam] `json:"serve_stale"`
}

func (r SetCacheSettingsRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify how long client browsers should cache the response. Cloudflare cache
// purge will not purge content cached on client browsers, so high browser TTLs may
// lead to stale content.
type SetCacheSettingsRuleActionParametersBrowserTTLParam struct {
	// Determines which browser ttl mode to use.
	Mode param.Field[SetCacheSettingsRuleActionParametersBrowserTTLMode] `json:"mode,required"`
	// The TTL (in seconds) if you choose override_origin mode.
	Default param.Field[int64] `json:"default"`
}

func (r SetCacheSettingsRuleActionParametersBrowserTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define which components of the request are included or excluded from the cache
// key Cloudflare uses to store the response in cache.
type SetCacheSettingsRuleActionParametersCacheKeyParam struct {
	// Separate cached content based on the visitor’s device type
	CacheByDeviceType param.Field[bool] `json:"cache_by_device_type"`
	// Protect from web cache deception attacks while allowing static assets to be
	// cached
	CacheDeceptionArmor param.Field[bool] `json:"cache_deception_armor"`
	// Customize which components of the request are included or excluded from the
	// cache key.
	CustomKey param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyParam] `json:"custom_key"`
	// Treat requests with the same query parameters the same, regardless of the order
	// those query parameters are in. A value of true ignores the query strings' order.
	IgnoreQueryStringsOrder param.Field[bool] `json:"ignore_query_strings_order"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Customize which components of the request are included or excluded from the
// cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyParam struct {
	// The cookies to include in building the cache key.
	Cookie param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieParam] `json:"cookie"`
	// The header names and values to include in building the cache key.
	Header param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderParam] `json:"header"`
	// Whether to use the original host or the resolved host in the cache key.
	Host param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostParam] `json:"host"`
	// Use the presence or absence of parameters in the query string to build the cache
	// key.
	QueryString param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringParam] `json:"query_string"`
	// Characteristics of the request user agent used in building the cache key.
	User param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserParam] `json:"user"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The cookies to include in building the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieParam struct {
	// Checks for the presence of these cookie names. The presence of these cookies is
	// used in building the cache key.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// Include these cookies' names and their values.
	Include param.Field[[]string] `json:"include"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyCookieParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The header names and values to include in building the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderParam struct {
	// Checks for the presence of these header names. The presence of these headers is
	// used in building the cache key.
	CheckPresence param.Field[[]string] `json:"check_presence"`
	// Whether or not to include the origin header. A value of true will exclude the
	// origin header in the cache key.
	ExcludeOrigin param.Field[bool] `json:"exclude_origin"`
	// Include these headers' names and their values.
	Include param.Field[[]string] `json:"include"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHeaderParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Whether to use the original host or the resolved host in the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostParam struct {
	// Use the resolved host in the cache key. A value of true will use the resolved
	// host, while a value or false will use the original host.
	Resolved param.Field[bool] `json:"resolved"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyHostParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Use the presence or absence of parameters in the query string to build the cache
// key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringParam struct {
	// build the cache key using all query string parameters EXCECPT these excluded
	// parameters
	Exclude param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeParam] `json:"exclude"`
	// build the cache key using a list of query string parameters that ARE in the
	// request.
	Include param.Field[SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeParam] `json:"include"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// build the cache key using all query string parameters EXCECPT these excluded
// parameters
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeParam struct {
	// Exclude all query string parameters from use in building the cache key.
	All param.Field[bool] `json:"all"`
	// A list of query string parameters NOT used to build the cache key. All
	// parameters present in the request but missing in this list will be used to build
	// the cache key.
	List param.Field[[]string] `json:"list"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringExcludeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// build the cache key using a list of query string parameters that ARE in the
// request.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeParam struct {
	// Use all query string parameters in the cache key.
	All param.Field[bool] `json:"all"`
	// A list of query string parameters used to build the cache key.
	List param.Field[[]string] `json:"list"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyQueryStringIncludeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Characteristics of the request user agent used in building the cache key.
type SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserParam struct {
	// Use the user agent's device type in the cache key.
	DeviceType param.Field[bool] `json:"device_type"`
	// Use the user agents's country in the cache key.
	Geo param.Field[bool] `json:"geo"`
	// Use the user agent's language in the cache key.
	Lang param.Field[bool] `json:"lang"`
}

func (r SetCacheSettingsRuleActionParametersCacheKeyCustomKeyUserParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Mark whether the request's response from origin is eligible for Cache Reserve
// (requires a Cache Reserve add-on plan).
type SetCacheSettingsRuleActionParametersCacheReserveParam struct {
	// Determines whether cache reserve is enabled. If this is true and a request meets
	// eligibility criteria, Cloudflare will write the resource to cache reserve.
	Eligible param.Field[bool] `json:"eligible,required"`
	// The minimum file size eligible for store in cache reserve.
	MinFileSize param.Field[int64] `json:"min_file_size,required"`
}

func (r SetCacheSettingsRuleActionParametersCacheReserveParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// TTL (Time to Live) specifies the maximum time to cache a resource in the
// Cloudflare edge network.
type SetCacheSettingsRuleActionParametersEdgeTTLParam struct {
	// The TTL (in seconds) if you choose override_origin mode.
	Default param.Field[int64] `json:"default,required"`
	// edge ttl options
	Mode param.Field[SetCacheSettingsRuleActionParametersEdgeTTLMode] `json:"mode,required"`
	// List of single status codes, or status code ranges to apply the selected mode
	StatusCodeTTL param.Field[[]SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLParam] `json:"status_code_ttl,required"`
}

func (r SetCacheSettingsRuleActionParametersEdgeTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Specify how long Cloudflare should cache the response based on the status code
// from the origin. Can be a single status code or a range or status codes
type SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLParam struct {
	// Time to cache a response (in seconds). A value of 0 is equivalent to setting the
	// Cache-Control header with the value "no-cache". A value of -1 is equivalent to
	// setting Cache-Control header with the value of "no-store".
	Value param.Field[int64] `json:"value,required"`
	// The range of status codes used to apply the selected mode.
	StatusCodeRange param.Field[SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeParam] `json:"status_code_range"`
	// Set the ttl for responses with this specific status code
	StatusCodeValue param.Field[int64] `json:"status_code_value"`
}

func (r SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The range of status codes used to apply the selected mode.
type SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeParam struct {
	// response status code lower bound
	From param.Field[int64] `json:"from,required"`
	// response status code upper bound
	To param.Field[int64] `json:"to,required"`
}

func (r SetCacheSettingsRuleActionParametersEdgeTTLStatusCodeTTLStatusCodeRangeParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Define if Cloudflare should serve stale content while getting the latest content
// from the origin. If on, Cloudflare will not serve stale content while getting
// the latest content from the origin.
type SetCacheSettingsRuleActionParametersServeStaleParam struct {
	// Defines whether Cloudflare should serve stale content while updating. If true,
	// Cloudflare will not serve stale content while getting the latest content from
	// the origin.
	DisableStaleWhileUpdating param.Field[bool] `json:"disable_stale_while_updating,required"`
}

func (r SetCacheSettingsRuleActionParametersServeStaleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SetConfigRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action SetConfigRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters SetConfigRuleActionParameters `json:"action_parameters"`
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
	Ref  string            `json:"ref"`
	JSON setConfigRuleJSON `json:"-"`
}

// setConfigRuleJSON contains the JSON metadata for the struct [SetConfigRule]
type setConfigRuleJSON struct {
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

func (r *SetConfigRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setConfigRuleJSON) RawJSON() string {
	return r.raw
}

func (r SetConfigRule) implementsRulesetsRulesetNewResponseRule() {}

func (r SetConfigRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r SetConfigRule) implementsRulesetsRulesetGetResponseRule() {}

func (r SetConfigRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r SetConfigRule) implementsRulesetsPhaseGetResponseRule() {}

func (r SetConfigRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r SetConfigRule) implementsRulesetsRuleNewResponseRule() {}

func (r SetConfigRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r SetConfigRule) implementsRulesetsRuleEditResponseRule() {}

func (r SetConfigRule) implementsRulesetsVersionGetResponseRule() {}

func (r SetConfigRule) implementsRulesetsVersionByTagGetResponseRule() {}

// The action to perform when the rule matches.
type SetConfigRuleAction string

const (
	SetConfigRuleActionSetConfig SetConfigRuleAction = "set_config"
)

func (r SetConfigRuleAction) IsKnown() bool {
	switch r {
	case SetConfigRuleActionSetConfig:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type SetConfigRuleActionParameters struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites bool `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify SetConfigRuleActionParametersAutominify `json:"autominify"`
	// Turn on or off Browser Integrity Check.
	Bic bool `json:"bic"`
	// Turn off all active Cloudflare Apps.
	DisableApps SetConfigRuleActionParametersDisableApps `json:"disable_apps"`
	// Turn off Real User Monitoring (RUM).
	DisableRUM SetConfigRuleActionParametersDisableRUM `json:"disable_rum"`
	// Turn off Zaraz.
	DisableZaraz SetConfigRuleActionParametersDisableZaraz `json:"disable_zaraz"`
	// Turn on or off Email Obfuscation.
	EmailObfuscation bool `json:"email_obfuscation"`
	// Turn on or off Cloudflare Fonts.
	Fonts bool `json:"fonts"`
	// Turn on or off the Hotlink Protection.
	HotlinkProtection bool `json:"hotlink_protection"`
	// Turn on or off Mirage.
	Mirage bool `json:"mirage"`
	// Turn on or off Opportunistic Encryption.
	OpportunisticEncryption bool `json:"opportunistic_encryption"`
	// Configure the Polish level.
	Polish SetConfigRuleActionParametersPolish `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader bool `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel SetConfigRuleActionParametersSecurityLevel `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes bool `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL SetConfigRuleActionParametersSSL `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg  bool                              `json:"sxg"`
	JSON setConfigRuleActionParametersJSON `json:"-"`
}

// setConfigRuleActionParametersJSON contains the JSON metadata for the struct
// [SetConfigRuleActionParameters]
type setConfigRuleActionParametersJSON struct {
	AutomaticHTTPSRewrites  apijson.Field
	Autominify              apijson.Field
	Bic                     apijson.Field
	DisableApps             apijson.Field
	DisableRUM              apijson.Field
	DisableZaraz            apijson.Field
	EmailObfuscation        apijson.Field
	Fonts                   apijson.Field
	HotlinkProtection       apijson.Field
	Mirage                  apijson.Field
	OpportunisticEncryption apijson.Field
	Polish                  apijson.Field
	RocketLoader            apijson.Field
	SecurityLevel           apijson.Field
	ServerSideExcludes      apijson.Field
	SSL                     apijson.Field
	Sxg                     apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *SetConfigRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setConfigRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// Select which file extensions to minify automatically.
type SetConfigRuleActionParametersAutominify struct {
	// Minify CSS files.
	Css bool `json:"css"`
	// Minify HTML files.
	HTML bool `json:"html"`
	// Minify JS files.
	JS   bool                                        `json:"js"`
	JSON setConfigRuleActionParametersAutominifyJSON `json:"-"`
}

// setConfigRuleActionParametersAutominifyJSON contains the JSON metadata for the
// struct [SetConfigRuleActionParametersAutominify]
type setConfigRuleActionParametersAutominifyJSON struct {
	Css         apijson.Field
	HTML        apijson.Field
	JS          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SetConfigRuleActionParametersAutominify) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r setConfigRuleActionParametersAutominifyJSON) RawJSON() string {
	return r.raw
}

// Turn off all active Cloudflare Apps.
type SetConfigRuleActionParametersDisableApps bool

const (
	SetConfigRuleActionParametersDisableAppsTrue SetConfigRuleActionParametersDisableApps = true
)

func (r SetConfigRuleActionParametersDisableApps) IsKnown() bool {
	switch r {
	case SetConfigRuleActionParametersDisableAppsTrue:
		return true
	}
	return false
}

// Turn off Real User Monitoring (RUM).
type SetConfigRuleActionParametersDisableRUM bool

const (
	SetConfigRuleActionParametersDisableRUMTrue SetConfigRuleActionParametersDisableRUM = true
)

func (r SetConfigRuleActionParametersDisableRUM) IsKnown() bool {
	switch r {
	case SetConfigRuleActionParametersDisableRUMTrue:
		return true
	}
	return false
}

// Turn off Zaraz.
type SetConfigRuleActionParametersDisableZaraz bool

const (
	SetConfigRuleActionParametersDisableZarazTrue SetConfigRuleActionParametersDisableZaraz = true
)

func (r SetConfigRuleActionParametersDisableZaraz) IsKnown() bool {
	switch r {
	case SetConfigRuleActionParametersDisableZarazTrue:
		return true
	}
	return false
}

// Configure the Polish level.
type SetConfigRuleActionParametersPolish string

const (
	SetConfigRuleActionParametersPolishOff      SetConfigRuleActionParametersPolish = "off"
	SetConfigRuleActionParametersPolishLossless SetConfigRuleActionParametersPolish = "lossless"
	SetConfigRuleActionParametersPolishLossy    SetConfigRuleActionParametersPolish = "lossy"
)

func (r SetConfigRuleActionParametersPolish) IsKnown() bool {
	switch r {
	case SetConfigRuleActionParametersPolishOff, SetConfigRuleActionParametersPolishLossless, SetConfigRuleActionParametersPolishLossy:
		return true
	}
	return false
}

// Configure the Security Level.
type SetConfigRuleActionParametersSecurityLevel string

const (
	SetConfigRuleActionParametersSecurityLevelOff            SetConfigRuleActionParametersSecurityLevel = "off"
	SetConfigRuleActionParametersSecurityLevelEssentiallyOff SetConfigRuleActionParametersSecurityLevel = "essentially_off"
	SetConfigRuleActionParametersSecurityLevelLow            SetConfigRuleActionParametersSecurityLevel = "low"
	SetConfigRuleActionParametersSecurityLevelMedium         SetConfigRuleActionParametersSecurityLevel = "medium"
	SetConfigRuleActionParametersSecurityLevelHigh           SetConfigRuleActionParametersSecurityLevel = "high"
	SetConfigRuleActionParametersSecurityLevelUnderAttack    SetConfigRuleActionParametersSecurityLevel = "under_attack"
)

func (r SetConfigRuleActionParametersSecurityLevel) IsKnown() bool {
	switch r {
	case SetConfigRuleActionParametersSecurityLevelOff, SetConfigRuleActionParametersSecurityLevelEssentiallyOff, SetConfigRuleActionParametersSecurityLevelLow, SetConfigRuleActionParametersSecurityLevelMedium, SetConfigRuleActionParametersSecurityLevelHigh, SetConfigRuleActionParametersSecurityLevelUnderAttack:
		return true
	}
	return false
}

// Configure the SSL level.
type SetConfigRuleActionParametersSSL string

const (
	SetConfigRuleActionParametersSSLOff        SetConfigRuleActionParametersSSL = "off"
	SetConfigRuleActionParametersSSLFlexible   SetConfigRuleActionParametersSSL = "flexible"
	SetConfigRuleActionParametersSSLFull       SetConfigRuleActionParametersSSL = "full"
	SetConfigRuleActionParametersSSLStrict     SetConfigRuleActionParametersSSL = "strict"
	SetConfigRuleActionParametersSSLOriginPull SetConfigRuleActionParametersSSL = "origin_pull"
)

func (r SetConfigRuleActionParametersSSL) IsKnown() bool {
	switch r {
	case SetConfigRuleActionParametersSSLOff, SetConfigRuleActionParametersSSLFlexible, SetConfigRuleActionParametersSSLFull, SetConfigRuleActionParametersSSLStrict, SetConfigRuleActionParametersSSLOriginPull:
		return true
	}
	return false
}

type SetConfigRuleParam struct {
	// The unique ID of the rule.
	ID param.Field[string] `json:"id"`
	// The action to perform when the rule matches.
	Action param.Field[SetConfigRuleAction] `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters param.Field[SetConfigRuleActionParametersParam] `json:"action_parameters"`
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

func (r SetConfigRuleParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r SetConfigRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r SetConfigRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r SetConfigRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type SetConfigRuleActionParametersParam struct {
	// Turn on or off Automatic HTTPS Rewrites.
	AutomaticHTTPSRewrites param.Field[bool] `json:"automatic_https_rewrites"`
	// Select which file extensions to minify automatically.
	Autominify param.Field[SetConfigRuleActionParametersAutominifyParam] `json:"autominify"`
	// Turn on or off Browser Integrity Check.
	Bic param.Field[bool] `json:"bic"`
	// Turn off all active Cloudflare Apps.
	DisableApps param.Field[SetConfigRuleActionParametersDisableApps] `json:"disable_apps"`
	// Turn off Real User Monitoring (RUM).
	DisableRUM param.Field[SetConfigRuleActionParametersDisableRUM] `json:"disable_rum"`
	// Turn off Zaraz.
	DisableZaraz param.Field[SetConfigRuleActionParametersDisableZaraz] `json:"disable_zaraz"`
	// Turn on or off Email Obfuscation.
	EmailObfuscation param.Field[bool] `json:"email_obfuscation"`
	// Turn on or off Cloudflare Fonts.
	Fonts param.Field[bool] `json:"fonts"`
	// Turn on or off the Hotlink Protection.
	HotlinkProtection param.Field[bool] `json:"hotlink_protection"`
	// Turn on or off Mirage.
	Mirage param.Field[bool] `json:"mirage"`
	// Turn on or off Opportunistic Encryption.
	OpportunisticEncryption param.Field[bool] `json:"opportunistic_encryption"`
	// Configure the Polish level.
	Polish param.Field[SetConfigRuleActionParametersPolish] `json:"polish"`
	// Turn on or off Rocket Loader
	RocketLoader param.Field[bool] `json:"rocket_loader"`
	// Configure the Security Level.
	SecurityLevel param.Field[SetConfigRuleActionParametersSecurityLevel] `json:"security_level"`
	// Turn on or off Server Side Excludes.
	ServerSideExcludes param.Field[bool] `json:"server_side_excludes"`
	// Configure the SSL level.
	SSL param.Field[SetConfigRuleActionParametersSSL] `json:"ssl"`
	// Turn on or off Signed Exchanges (SXG).
	Sxg param.Field[bool] `json:"sxg"`
}

func (r SetConfigRuleActionParametersParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Select which file extensions to minify automatically.
type SetConfigRuleActionParametersAutominifyParam struct {
	// Minify CSS files.
	Css param.Field[bool] `json:"css"`
	// Minify HTML files.
	HTML param.Field[bool] `json:"html"`
	// Minify JS files.
	JS param.Field[bool] `json:"js"`
}

func (r SetConfigRuleActionParametersAutominifyParam) MarshalJSON() (data []byte, err error) {
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

func (r SkipRule) implementsRulesetsRulesetNewResponseRule() {}

func (r SkipRule) implementsRulesetsRulesetUpdateResponseRule() {}

func (r SkipRule) implementsRulesetsRulesetGetResponseRule() {}

func (r SkipRule) implementsRulesetsPhaseUpdateResponseRule() {}

func (r SkipRule) implementsRulesetsPhaseGetResponseRule() {}

func (r SkipRule) implementsRulesetsPhaseVersionGetResponseRule() {}

func (r SkipRule) implementsRulesetsRuleNewResponseRule() {}

func (r SkipRule) implementsRulesetsRuleDeleteResponseRule() {}

func (r SkipRule) implementsRulesetsRuleEditResponseRule() {}

func (r SkipRule) implementsRulesetsVersionGetResponseRule() {}

func (r SkipRule) implementsRulesetsVersionByTagGetResponseRule() {}

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
	Phases []Phase `json:"phases"`
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

func (r SkipRuleParam) implementsRulesetsRulesetNewParamsRuleUnion() {}

func (r SkipRuleParam) implementsRulesetsRulesetUpdateParamsRuleUnion() {}

func (r SkipRuleParam) implementsRulesetsPhaseUpdateParamsRuleUnion() {}

// The parameters configuring the rule's action.
type SkipRuleActionParametersParam struct {
	// A list of phases to skip the execution of. This option is incompatible with the
	// ruleset and rulesets options.
	Phases param.Field[[]Phase] `json:"phases"`
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
	Kind Kind `json:"kind,required"`
	// The timestamp of when the ruleset was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The human-readable name of the ruleset.
	Name string `json:"name,required"`
	// The phase of the ruleset.
	Phase Phase `json:"phase,required"`
	// The list of rules in the ruleset.
	Rules []RuleNewResponseRule `json:"rules,required"`
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

type RuleNewResponseRule struct {
	// The action to perform when the rule matches.
	Action           RuleNewResponseRulesAction `json:"action"`
	ActionParameters interface{}                `json:"action_parameters,required"`
	Categories       interface{}                `json:"categories,required"`
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
	Version string                  `json:"version,required"`
	JSON    ruleNewResponseRuleJSON `json:"-"`
	union   RuleNewResponseRulesUnion
}

// ruleNewResponseRuleJSON contains the JSON metadata for the struct
// [RuleNewResponseRule]
type ruleNewResponseRuleJSON struct {
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

func (r ruleNewResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RuleNewResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleNewResponseRule) AsUnion() RuleNewResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JSChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule], [rulesets.SetCacheSettingsRule] or
// [rulesets.RuleNewResponseRulesRulesetsLogCustomFieldRule].
type RuleNewResponseRulesUnion interface {
	implementsRulesetsRuleNewResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleNewResponseRulesUnion)(nil)).Elem(),
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
			Type:               reflect.TypeOf(RuleNewResponseRulesRulesetsLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
	)
}

type RuleNewResponseRulesRulesetsLogCustomFieldRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleNewResponseRulesRulesetsLogCustomFieldRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParameters `json:"action_parameters"`
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
	Ref  string                                             `json:"ref"`
	JSON ruleNewResponseRulesRulesetsLogCustomFieldRuleJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsLogCustomFieldRuleJSON contains the JSON metadata
// for the struct [RuleNewResponseRulesRulesetsLogCustomFieldRule]
type ruleNewResponseRulesRulesetsLogCustomFieldRuleJSON struct {
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

func (r *RuleNewResponseRulesRulesetsLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleNewResponseRulesRulesetsLogCustomFieldRule) implementsRulesetsRuleNewResponseRule() {}

// The action to perform when the rule matches.
type RuleNewResponseRulesRulesetsLogCustomFieldRuleAction string

const (
	RuleNewResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField RuleNewResponseRulesRulesetsLogCustomFieldRuleAction = "log_custom_field"
)

func (r RuleNewResponseRulesRulesetsLogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case RuleNewResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields []RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField `json:"cookie_fields"`
	// The request fields to log.
	RequestFields []RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField `json:"request_fields"`
	// The response fields to log.
	ResponseFields []RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField `json:"response_fields"`
	JSON           ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON            `json:"-"`
}

// ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParameters]
type ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON struct {
	CookieFields   apijson.Field
	RequestFields  apijson.Field
	ResponseFields apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The cookie field to log.
type RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name string                                                                        `json:"name,required"`
	JSON ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON
// contains the JSON metadata for the struct
// [RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField]
type ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON) RawJSON() string {
	return r.raw
}

// The request field to log.
type RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name string                                                                         `json:"name,required"`
	JSON ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON
// contains the JSON metadata for the struct
// [RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField]
type ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON) RawJSON() string {
	return r.raw
}

// The response field to log.
type RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name string                                                                          `json:"name,required"`
	JSON ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON `json:"-"`
}

// ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON
// contains the JSON metadata for the struct
// [RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField]
type ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleNewResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RuleNewResponseRulesAction string

const (
	RuleNewResponseRulesActionBlock            RuleNewResponseRulesAction = "block"
	RuleNewResponseRulesActionChallenge        RuleNewResponseRulesAction = "challenge"
	RuleNewResponseRulesActionCompressResponse RuleNewResponseRulesAction = "compress_response"
	RuleNewResponseRulesActionExecute          RuleNewResponseRulesAction = "execute"
	RuleNewResponseRulesActionJSChallenge      RuleNewResponseRulesAction = "js_challenge"
	RuleNewResponseRulesActionLog              RuleNewResponseRulesAction = "log"
	RuleNewResponseRulesActionManagedChallenge RuleNewResponseRulesAction = "managed_challenge"
	RuleNewResponseRulesActionRedirect         RuleNewResponseRulesAction = "redirect"
	RuleNewResponseRulesActionRewrite          RuleNewResponseRulesAction = "rewrite"
	RuleNewResponseRulesActionRoute            RuleNewResponseRulesAction = "route"
	RuleNewResponseRulesActionScore            RuleNewResponseRulesAction = "score"
	RuleNewResponseRulesActionServeError       RuleNewResponseRulesAction = "serve_error"
	RuleNewResponseRulesActionSetConfig        RuleNewResponseRulesAction = "set_config"
	RuleNewResponseRulesActionSkip             RuleNewResponseRulesAction = "skip"
	RuleNewResponseRulesActionSetCacheSettings RuleNewResponseRulesAction = "set_cache_settings"
	RuleNewResponseRulesActionLogCustomField   RuleNewResponseRulesAction = "log_custom_field"
)

func (r RuleNewResponseRulesAction) IsKnown() bool {
	switch r {
	case RuleNewResponseRulesActionBlock, RuleNewResponseRulesActionChallenge, RuleNewResponseRulesActionCompressResponse, RuleNewResponseRulesActionExecute, RuleNewResponseRulesActionJSChallenge, RuleNewResponseRulesActionLog, RuleNewResponseRulesActionManagedChallenge, RuleNewResponseRulesActionRedirect, RuleNewResponseRulesActionRewrite, RuleNewResponseRulesActionRoute, RuleNewResponseRulesActionScore, RuleNewResponseRulesActionServeError, RuleNewResponseRulesActionSetConfig, RuleNewResponseRulesActionSkip, RuleNewResponseRulesActionSetCacheSettings, RuleNewResponseRulesActionLogCustomField:
		return true
	}
	return false
}

// A ruleset object.
type RuleDeleteResponse struct {
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
	Rules []RuleDeleteResponseRule `json:"rules,required"`
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

type RuleDeleteResponseRule struct {
	// The action to perform when the rule matches.
	Action           RuleDeleteResponseRulesAction `json:"action"`
	ActionParameters interface{}                   `json:"action_parameters,required"`
	Categories       interface{}                   `json:"categories,required"`
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
	Version string                     `json:"version,required"`
	JSON    ruleDeleteResponseRuleJSON `json:"-"`
	union   RuleDeleteResponseRulesUnion
}

// ruleDeleteResponseRuleJSON contains the JSON metadata for the struct
// [RuleDeleteResponseRule]
type ruleDeleteResponseRuleJSON struct {
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

func (r ruleDeleteResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RuleDeleteResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleDeleteResponseRule) AsUnion() RuleDeleteResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JSChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule], [rulesets.SetCacheSettingsRule] or
// [rulesets.RuleDeleteResponseRulesRulesetsLogCustomFieldRule].
type RuleDeleteResponseRulesUnion interface {
	implementsRulesetsRuleDeleteResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleDeleteResponseRulesUnion)(nil)).Elem(),
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
			Type:               reflect.TypeOf(RuleDeleteResponseRulesRulesetsLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
	)
}

type RuleDeleteResponseRulesRulesetsLogCustomFieldRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleDeleteResponseRulesRulesetsLogCustomFieldRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParameters `json:"action_parameters"`
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
	Ref  string                                                `json:"ref"`
	JSON ruleDeleteResponseRulesRulesetsLogCustomFieldRuleJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsLogCustomFieldRuleJSON contains the JSON metadata
// for the struct [RuleDeleteResponseRulesRulesetsLogCustomFieldRule]
type ruleDeleteResponseRulesRulesetsLogCustomFieldRuleJSON struct {
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

func (r *RuleDeleteResponseRulesRulesetsLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleDeleteResponseRulesRulesetsLogCustomFieldRule) implementsRulesetsRuleDeleteResponseRule() {
}

// The action to perform when the rule matches.
type RuleDeleteResponseRulesRulesetsLogCustomFieldRuleAction string

const (
	RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField RuleDeleteResponseRulesRulesetsLogCustomFieldRuleAction = "log_custom_field"
)

func (r RuleDeleteResponseRulesRulesetsLogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields []RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField `json:"cookie_fields"`
	// The request fields to log.
	RequestFields []RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField `json:"request_fields"`
	// The response fields to log.
	ResponseFields []RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField `json:"response_fields"`
	JSON           ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON            `json:"-"`
}

// ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON contains
// the JSON metadata for the struct
// [RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParameters]
type ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON struct {
	CookieFields   apijson.Field
	RequestFields  apijson.Field
	ResponseFields apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The cookie field to log.
type RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name string                                                                           `json:"name,required"`
	JSON ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON
// contains the JSON metadata for the struct
// [RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField]
type ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON) RawJSON() string {
	return r.raw
}

// The request field to log.
type RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name string                                                                            `json:"name,required"`
	JSON ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON
// contains the JSON metadata for the struct
// [RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField]
type ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON) RawJSON() string {
	return r.raw
}

// The response field to log.
type RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name string                                                                             `json:"name,required"`
	JSON ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON `json:"-"`
}

// ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON
// contains the JSON metadata for the struct
// [RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField]
type ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RuleDeleteResponseRulesAction string

const (
	RuleDeleteResponseRulesActionBlock            RuleDeleteResponseRulesAction = "block"
	RuleDeleteResponseRulesActionChallenge        RuleDeleteResponseRulesAction = "challenge"
	RuleDeleteResponseRulesActionCompressResponse RuleDeleteResponseRulesAction = "compress_response"
	RuleDeleteResponseRulesActionExecute          RuleDeleteResponseRulesAction = "execute"
	RuleDeleteResponseRulesActionJSChallenge      RuleDeleteResponseRulesAction = "js_challenge"
	RuleDeleteResponseRulesActionLog              RuleDeleteResponseRulesAction = "log"
	RuleDeleteResponseRulesActionManagedChallenge RuleDeleteResponseRulesAction = "managed_challenge"
	RuleDeleteResponseRulesActionRedirect         RuleDeleteResponseRulesAction = "redirect"
	RuleDeleteResponseRulesActionRewrite          RuleDeleteResponseRulesAction = "rewrite"
	RuleDeleteResponseRulesActionRoute            RuleDeleteResponseRulesAction = "route"
	RuleDeleteResponseRulesActionScore            RuleDeleteResponseRulesAction = "score"
	RuleDeleteResponseRulesActionServeError       RuleDeleteResponseRulesAction = "serve_error"
	RuleDeleteResponseRulesActionSetConfig        RuleDeleteResponseRulesAction = "set_config"
	RuleDeleteResponseRulesActionSkip             RuleDeleteResponseRulesAction = "skip"
	RuleDeleteResponseRulesActionSetCacheSettings RuleDeleteResponseRulesAction = "set_cache_settings"
	RuleDeleteResponseRulesActionLogCustomField   RuleDeleteResponseRulesAction = "log_custom_field"
)

func (r RuleDeleteResponseRulesAction) IsKnown() bool {
	switch r {
	case RuleDeleteResponseRulesActionBlock, RuleDeleteResponseRulesActionChallenge, RuleDeleteResponseRulesActionCompressResponse, RuleDeleteResponseRulesActionExecute, RuleDeleteResponseRulesActionJSChallenge, RuleDeleteResponseRulesActionLog, RuleDeleteResponseRulesActionManagedChallenge, RuleDeleteResponseRulesActionRedirect, RuleDeleteResponseRulesActionRewrite, RuleDeleteResponseRulesActionRoute, RuleDeleteResponseRulesActionScore, RuleDeleteResponseRulesActionServeError, RuleDeleteResponseRulesActionSetConfig, RuleDeleteResponseRulesActionSkip, RuleDeleteResponseRulesActionSetCacheSettings, RuleDeleteResponseRulesActionLogCustomField:
		return true
	}
	return false
}

// A ruleset object.
type RuleEditResponse struct {
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
	Rules []RuleEditResponseRule `json:"rules,required"`
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

type RuleEditResponseRule struct {
	// The action to perform when the rule matches.
	Action           RuleEditResponseRulesAction `json:"action"`
	ActionParameters interface{}                 `json:"action_parameters,required"`
	Categories       interface{}                 `json:"categories,required"`
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
	Version string                   `json:"version,required"`
	JSON    ruleEditResponseRuleJSON `json:"-"`
	union   RuleEditResponseRulesUnion
}

// ruleEditResponseRuleJSON contains the JSON metadata for the struct
// [RuleEditResponseRule]
type ruleEditResponseRuleJSON struct {
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

func (r ruleEditResponseRuleJSON) RawJSON() string {
	return r.raw
}

func (r *RuleEditResponseRule) UnmarshalJSON(data []byte) (err error) {
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

func (r RuleEditResponseRule) AsUnion() RuleEditResponseRulesUnion {
	return r.union
}

// Union satisfied by [rulesets.BlockRule], [rulesets.ChallengeRule],
// [rulesets.CompressResponseRule], [rulesets.ExecuteRule],
// [rulesets.JSChallengeRule], [rulesets.LogRule], [rulesets.ManagedChallengeRule],
// [rulesets.RedirectRule], [rulesets.RewriteRule], [rulesets.RouteRule],
// [rulesets.ScoreRule], [rulesets.ServeErrorRule], [rulesets.SetConfigRule],
// [rulesets.SkipRule], [rulesets.SetCacheSettingsRule] or
// [rulesets.RuleEditResponseRulesRulesetsLogCustomFieldRule].
type RuleEditResponseRulesUnion interface {
	implementsRulesetsRuleEditResponseRule()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*RuleEditResponseRulesUnion)(nil)).Elem(),
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
			Type:               reflect.TypeOf(RuleEditResponseRulesRulesetsLogCustomFieldRule{}),
			DiscriminatorValue: "log_custom_field",
		},
	)
}

type RuleEditResponseRulesRulesetsLogCustomFieldRule struct {
	// The timestamp of when the rule was last modified.
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// The version of the rule.
	Version string `json:"version,required"`
	// The unique ID of the rule.
	ID string `json:"id"`
	// The action to perform when the rule matches.
	Action RuleEditResponseRulesRulesetsLogCustomFieldRuleAction `json:"action"`
	// The parameters configuring the rule's action.
	ActionParameters RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParameters `json:"action_parameters"`
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
	Ref  string                                              `json:"ref"`
	JSON ruleEditResponseRulesRulesetsLogCustomFieldRuleJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsLogCustomFieldRuleJSON contains the JSON metadata
// for the struct [RuleEditResponseRulesRulesetsLogCustomFieldRule]
type ruleEditResponseRulesRulesetsLogCustomFieldRuleJSON struct {
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

func (r *RuleEditResponseRulesRulesetsLogCustomFieldRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsLogCustomFieldRuleJSON) RawJSON() string {
	return r.raw
}

func (r RuleEditResponseRulesRulesetsLogCustomFieldRule) implementsRulesetsRuleEditResponseRule() {}

// The action to perform when the rule matches.
type RuleEditResponseRulesRulesetsLogCustomFieldRuleAction string

const (
	RuleEditResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField RuleEditResponseRulesRulesetsLogCustomFieldRuleAction = "log_custom_field"
)

func (r RuleEditResponseRulesRulesetsLogCustomFieldRuleAction) IsKnown() bool {
	switch r {
	case RuleEditResponseRulesRulesetsLogCustomFieldRuleActionLogCustomField:
		return true
	}
	return false
}

// The parameters configuring the rule's action.
type RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParameters struct {
	// The cookie fields to log.
	CookieFields []RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField `json:"cookie_fields"`
	// The request fields to log.
	RequestFields []RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField `json:"request_fields"`
	// The response fields to log.
	ResponseFields []RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField `json:"response_fields"`
	JSON           ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON            `json:"-"`
}

// ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON contains the
// JSON metadata for the struct
// [RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParameters]
type ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON struct {
	CookieFields   apijson.Field
	RequestFields  apijson.Field
	ResponseFields apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersJSON) RawJSON() string {
	return r.raw
}

// The cookie field to log.
type RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField struct {
	// The name of the field.
	Name string                                                                         `json:"name,required"`
	JSON ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON
// contains the JSON metadata for the struct
// [RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField]
type ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersCookieFieldJSON) RawJSON() string {
	return r.raw
}

// The request field to log.
type RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField struct {
	// The name of the field.
	Name string                                                                          `json:"name,required"`
	JSON ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON
// contains the JSON metadata for the struct
// [RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField]
type ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersRequestFieldJSON) RawJSON() string {
	return r.raw
}

// The response field to log.
type RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField struct {
	// The name of the field.
	Name string                                                                           `json:"name,required"`
	JSON ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON `json:"-"`
}

// ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON
// contains the JSON metadata for the struct
// [RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField]
type ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON struct {
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleEditResponseRulesRulesetsLogCustomFieldRuleActionParametersResponseFieldJSON) RawJSON() string {
	return r.raw
}

// The action to perform when the rule matches.
type RuleEditResponseRulesAction string

const (
	RuleEditResponseRulesActionBlock            RuleEditResponseRulesAction = "block"
	RuleEditResponseRulesActionChallenge        RuleEditResponseRulesAction = "challenge"
	RuleEditResponseRulesActionCompressResponse RuleEditResponseRulesAction = "compress_response"
	RuleEditResponseRulesActionExecute          RuleEditResponseRulesAction = "execute"
	RuleEditResponseRulesActionJSChallenge      RuleEditResponseRulesAction = "js_challenge"
	RuleEditResponseRulesActionLog              RuleEditResponseRulesAction = "log"
	RuleEditResponseRulesActionManagedChallenge RuleEditResponseRulesAction = "managed_challenge"
	RuleEditResponseRulesActionRedirect         RuleEditResponseRulesAction = "redirect"
	RuleEditResponseRulesActionRewrite          RuleEditResponseRulesAction = "rewrite"
	RuleEditResponseRulesActionRoute            RuleEditResponseRulesAction = "route"
	RuleEditResponseRulesActionScore            RuleEditResponseRulesAction = "score"
	RuleEditResponseRulesActionServeError       RuleEditResponseRulesAction = "serve_error"
	RuleEditResponseRulesActionSetConfig        RuleEditResponseRulesAction = "set_config"
	RuleEditResponseRulesActionSkip             RuleEditResponseRulesAction = "skip"
	RuleEditResponseRulesActionSetCacheSettings RuleEditResponseRulesAction = "set_cache_settings"
	RuleEditResponseRulesActionLogCustomField   RuleEditResponseRulesAction = "log_custom_field"
)

func (r RuleEditResponseRulesAction) IsKnown() bool {
	switch r {
	case RuleEditResponseRulesActionBlock, RuleEditResponseRulesActionChallenge, RuleEditResponseRulesActionCompressResponse, RuleEditResponseRulesActionExecute, RuleEditResponseRulesActionJSChallenge, RuleEditResponseRulesActionLog, RuleEditResponseRulesActionManagedChallenge, RuleEditResponseRulesActionRedirect, RuleEditResponseRulesActionRewrite, RuleEditResponseRulesActionRoute, RuleEditResponseRulesActionScore, RuleEditResponseRulesActionServeError, RuleEditResponseRulesActionSetConfig, RuleEditResponseRulesActionSkip, RuleEditResponseRulesActionSetCacheSettings, RuleEditResponseRulesActionLogCustomField:
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
