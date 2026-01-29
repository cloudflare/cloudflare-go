// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package token_validation

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/api_gateway"
	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
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

// Create a token validation rule.
func (r *RuleService) New(ctx context.Context, params RuleNewParams, opts ...option.RequestOption) (res *TokenValidationRule, err error) {
	var env RuleNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/rules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List token validation rules
func (r *RuleService) List(ctx context.Context, params RuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[TokenValidationRule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/rules", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List token validation rules
func (r *RuleService) ListAutoPaging(ctx context.Context, params RuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[TokenValidationRule] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Delete a zone token validation rule.
func (r *RuleService) Delete(ctx context.Context, ruleID string, body RuleDeleteParams, opts ...option.RequestOption) (res *RuleDeleteResponse, err error) {
	var env RuleDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/rules/%s", body.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create zone token validation rules.
//
// A request can create multiple Token Validation Rules.
func (r *RuleService) BulkNew(ctx context.Context, params RuleBulkNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[TokenValidationRule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/rules/bulk", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Create zone token validation rules.
//
// A request can create multiple Token Validation Rules.
func (r *RuleService) BulkNewAutoPaging(ctx context.Context, params RuleBulkNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TokenValidationRule] {
	return pagination.NewSinglePageAutoPager(r.BulkNew(ctx, params, opts...))
}

// Edit token validation rules.
//
// A request can update multiple Token Validation Rules.
//
// Rules can be re-ordered using the `position` field.
//
// Returns all updated rules.
func (r *RuleService) BulkEdit(ctx context.Context, params RuleBulkEditParams, opts ...option.RequestOption) (res *pagination.SinglePage[TokenValidationRule], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/rules/bulk", params.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPatch, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Edit token validation rules.
//
// A request can update multiple Token Validation Rules.
//
// Rules can be re-ordered using the `position` field.
//
// Returns all updated rules.
func (r *RuleService) BulkEditAutoPaging(ctx context.Context, params RuleBulkEditParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[TokenValidationRule] {
	return pagination.NewSinglePageAutoPager(r.BulkEdit(ctx, params, opts...))
}

// Edit a zone token validation rule.
func (r *RuleService) Edit(ctx context.Context, ruleID string, params RuleEditParams, opts ...option.RequestOption) (res *TokenValidationRule, err error) {
	var env RuleEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/rules/%s", params.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a zone token validation rule.
func (r *RuleService) Get(ctx context.Context, ruleID string, query RuleGetParams, opts ...option.RequestOption) (res *TokenValidationRule, err error) {
	var env RuleGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/token_validation/rules/%s", query.ZoneID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A Token Validation rule that can enforce security policies using JWT Tokens.
type TokenValidationRule struct {
	// Action to take on requests that match operations included in `selector` and fail
	// `expression`.
	Action TokenValidationRuleAction `json:"action,required"`
	// A human-readable description that gives more details than `title`.
	Description string `json:"description,required"`
	// Toggle rule on or off.
	Enabled bool `json:"enabled,required"`
	// Rule expression. Requests that fail to match this expression will be subject to
	// `action`.
	//
	// For details on expressions, see the
	// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
	Expression string `json:"expression,required"`
	// Select operations covered by this rule.
	//
	// For details on selectors, see the
	// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
	Selector TokenValidationRuleSelector `json:"selector,required"`
	// A human-readable name for the rule.
	Title string `json:"title,required"`
	// UUID.
	ID          string                  `json:"id"`
	CreatedAt   time.Time               `json:"created_at" format:"date-time"`
	LastUpdated time.Time               `json:"last_updated" format:"date-time"`
	JSON        tokenValidationRuleJSON `json:"-"`
}

// tokenValidationRuleJSON contains the JSON metadata for the struct
// [TokenValidationRule]
type tokenValidationRuleJSON struct {
	Action      apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Expression  apijson.Field
	Selector    apijson.Field
	Title       apijson.Field
	ID          apijson.Field
	CreatedAt   apijson.Field
	LastUpdated apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenValidationRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValidationRuleJSON) RawJSON() string {
	return r.raw
}

// Action to take on requests that match operations included in `selector` and fail
// `expression`.
type TokenValidationRuleAction string

const (
	TokenValidationRuleActionLog   TokenValidationRuleAction = "log"
	TokenValidationRuleActionBlock TokenValidationRuleAction = "block"
)

func (r TokenValidationRuleAction) IsKnown() bool {
	switch r {
	case TokenValidationRuleActionLog, TokenValidationRuleActionBlock:
		return true
	}
	return false
}

// Select operations covered by this rule.
//
// For details on selectors, see the
// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
type TokenValidationRuleSelector struct {
	// Ignore operations that were otherwise included by `include`.
	Exclude []TokenValidationRuleSelectorExclude `json:"exclude,nullable"`
	// Select all matching operations.
	Include []TokenValidationRuleSelectorInclude `json:"include,nullable"`
	JSON    tokenValidationRuleSelectorJSON      `json:"-"`
}

// tokenValidationRuleSelectorJSON contains the JSON metadata for the struct
// [TokenValidationRuleSelector]
type tokenValidationRuleSelectorJSON struct {
	Exclude     apijson.Field
	Include     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenValidationRuleSelector) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValidationRuleSelectorJSON) RawJSON() string {
	return r.raw
}

type TokenValidationRuleSelectorExclude struct {
	// Excluded operation IDs.
	OperationIDs []string                               `json:"operation_ids"`
	JSON         tokenValidationRuleSelectorExcludeJSON `json:"-"`
}

// tokenValidationRuleSelectorExcludeJSON contains the JSON metadata for the struct
// [TokenValidationRuleSelectorExclude]
type tokenValidationRuleSelectorExcludeJSON struct {
	OperationIDs apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TokenValidationRuleSelectorExclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValidationRuleSelectorExcludeJSON) RawJSON() string {
	return r.raw
}

type TokenValidationRuleSelectorInclude struct {
	// Included hostnames.
	Host []string                               `json:"host" format:"hostname"`
	JSON tokenValidationRuleSelectorIncludeJSON `json:"-"`
}

// tokenValidationRuleSelectorIncludeJSON contains the JSON metadata for the struct
// [TokenValidationRuleSelectorInclude]
type tokenValidationRuleSelectorIncludeJSON struct {
	Host        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TokenValidationRuleSelectorInclude) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tokenValidationRuleSelectorIncludeJSON) RawJSON() string {
	return r.raw
}

type RuleDeleteResponse = interface{}

type RuleNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Action to take on requests that match operations included in `selector` and fail
	// `expression`.
	Action param.Field[RuleNewParamsAction] `json:"action,required"`
	// A human-readable description that gives more details than `title`.
	Description param.Field[string] `json:"description,required"`
	// Toggle rule on or off.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Rule expression. Requests that fail to match this expression will be subject to
	// `action`.
	//
	// For details on expressions, see the
	// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
	Expression param.Field[string] `json:"expression,required"`
	// Select operations covered by this rule.
	//
	// For details on selectors, see the
	// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
	Selector param.Field[RuleNewParamsSelector] `json:"selector,required"`
	// A human-readable name for the rule.
	Title param.Field[string] `json:"title,required"`
}

func (r RuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action to take on requests that match operations included in `selector` and fail
// `expression`.
type RuleNewParamsAction string

const (
	RuleNewParamsActionLog   RuleNewParamsAction = "log"
	RuleNewParamsActionBlock RuleNewParamsAction = "block"
)

func (r RuleNewParamsAction) IsKnown() bool {
	switch r {
	case RuleNewParamsActionLog, RuleNewParamsActionBlock:
		return true
	}
	return false
}

// Select operations covered by this rule.
//
// For details on selectors, see the
// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
type RuleNewParamsSelector struct {
	// Ignore operations that were otherwise included by `include`.
	Exclude param.Field[[]RuleNewParamsSelectorExclude] `json:"exclude"`
	// Select all matching operations.
	Include param.Field[[]RuleNewParamsSelectorInclude] `json:"include"`
}

func (r RuleNewParamsSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewParamsSelectorExclude struct {
	// Excluded operation IDs.
	OperationIDs param.Field[[]string] `json:"operation_ids"`
}

func (r RuleNewParamsSelectorExclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewParamsSelectorInclude struct {
	// Included hostnames.
	Host param.Field[[]string] `json:"host" format:"hostname"`
}

func (r RuleNewParamsSelectorInclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleNewResponseEnvelope struct {
	Errors   api_gateway.Message `json:"errors,required"`
	Messages api_gateway.Message `json:"messages,required"`
	// A Token Validation rule that can enforce security policies using JWT Tokens.
	Result TokenValidationRule `json:"result,required"`
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

type RuleListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Select rules with these IDs.
	ID param.Field[string] `query:"id"`
	// Action to take on requests that match operations included in `selector` and fail
	// `expression`.
	Action param.Field[RuleListParamsAction] `query:"action"`
	// Toggle rule on or off.
	Enabled param.Field[bool] `query:"enabled"`
	// Select rules with this host in `include`.
	Host param.Field[string] `query:"host" format:"hostname"`
	// Select rules with this host in `include`.
	Hostname param.Field[string] `query:"hostname" format:"hostname"`
	// Page number of paginated results.
	Page param.Field[int64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[int64] `query:"per_page"`
	// Select rules with these IDs.
	RuleID param.Field[string] `query:"rule_id"`
	// Select rules using any of these token configurations.
	TokenConfiguration param.Field[[]string] `query:"token_configuration"`
}

// URLQuery serializes [RuleListParams]'s query parameters as `url.Values`.
func (r RuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Action to take on requests that match operations included in `selector` and fail
// `expression`.
type RuleListParamsAction string

const (
	RuleListParamsActionLog   RuleListParamsAction = "log"
	RuleListParamsActionBlock RuleListParamsAction = "block"
)

func (r RuleListParamsAction) IsKnown() bool {
	switch r {
	case RuleListParamsActionLog, RuleListParamsActionBlock:
		return true
	}
	return false
}

type RuleDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RuleDeleteResponseEnvelope struct {
	Errors   api_gateway.Message `json:"errors,required"`
	Messages api_gateway.Message `json:"messages,required"`
	// Whether the API call was successful.
	Success RuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  RuleDeleteResponse                `json:"result"`
	JSON    ruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// ruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleDeleteResponseEnvelope]
type ruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleDeleteResponseEnvelopeJSON) RawJSON() string {
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

type RuleBulkNewParams struct {
	// Identifier.
	ZoneID param.Field[string]     `path:"zone_id,required"`
	Body   []RuleBulkNewParamsBody `json:"body,required"`
}

func (r RuleBulkNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

// A Token Validation rule that can enforce security policies using JWT Tokens.
type RuleBulkNewParamsBody struct {
	// Action to take on requests that match operations included in `selector` and fail
	// `expression`.
	Action param.Field[RuleBulkNewParamsBodyAction] `json:"action,required"`
	// A human-readable description that gives more details than `title`.
	Description param.Field[string] `json:"description,required"`
	// Toggle rule on or off.
	Enabled param.Field[bool] `json:"enabled,required"`
	// Rule expression. Requests that fail to match this expression will be subject to
	// `action`.
	//
	// For details on expressions, see the
	// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
	Expression param.Field[string] `json:"expression,required"`
	// Select operations covered by this rule.
	//
	// For details on selectors, see the
	// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
	Selector param.Field[RuleBulkNewParamsBodySelector] `json:"selector,required"`
	// A human-readable name for the rule.
	Title param.Field[string] `json:"title,required"`
}

func (r RuleBulkNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action to take on requests that match operations included in `selector` and fail
// `expression`.
type RuleBulkNewParamsBodyAction string

const (
	RuleBulkNewParamsBodyActionLog   RuleBulkNewParamsBodyAction = "log"
	RuleBulkNewParamsBodyActionBlock RuleBulkNewParamsBodyAction = "block"
)

func (r RuleBulkNewParamsBodyAction) IsKnown() bool {
	switch r {
	case RuleBulkNewParamsBodyActionLog, RuleBulkNewParamsBodyActionBlock:
		return true
	}
	return false
}

// Select operations covered by this rule.
//
// For details on selectors, see the
// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
type RuleBulkNewParamsBodySelector struct {
	// Ignore operations that were otherwise included by `include`.
	Exclude param.Field[[]RuleBulkNewParamsBodySelectorExclude] `json:"exclude"`
	// Select all matching operations.
	Include param.Field[[]RuleBulkNewParamsBodySelectorInclude] `json:"include"`
}

func (r RuleBulkNewParamsBodySelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleBulkNewParamsBodySelectorExclude struct {
	// Excluded operation IDs.
	OperationIDs param.Field[[]string] `json:"operation_ids"`
}

func (r RuleBulkNewParamsBodySelectorExclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleBulkNewParamsBodySelectorInclude struct {
	// Included hostnames.
	Host param.Field[[]string] `json:"host" format:"hostname"`
}

func (r RuleBulkNewParamsBodySelectorInclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleBulkEditParams struct {
	// Identifier.
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   []RuleBulkEditParamsBody `json:"body,required"`
}

func (r RuleBulkEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type RuleBulkEditParamsBody struct {
	// Rule ID this patch applies to
	ID param.Field[string] `json:"id,required" format:"uuid"`
	// Action to take on requests that match operations included in `selector` and fail
	// `expression`.
	Action param.Field[RuleBulkEditParamsBodyAction] `json:"action"`
	// A human-readable description that gives more details than `title`.
	Description param.Field[string] `json:"description"`
	// Toggle rule on or off.
	Enabled param.Field[bool] `json:"enabled"`
	// Rule expression. Requests that fail to match this expression will be subject to
	// `action`.
	//
	// For details on expressions, see the
	// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
	Expression param.Field[string] `json:"expression"`
	// Update rule order among zone rules.
	Position param.Field[RuleBulkEditParamsBodyPositionUnion] `json:"position"`
	// Select operations covered by this rule.
	//
	// For details on selectors, see the
	// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
	Selector param.Field[RuleBulkEditParamsBodySelector] `json:"selector"`
	// A human-readable name for the rule.
	Title param.Field[string] `json:"title"`
}

func (r RuleBulkEditParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action to take on requests that match operations included in `selector` and fail
// `expression`.
type RuleBulkEditParamsBodyAction string

const (
	RuleBulkEditParamsBodyActionLog   RuleBulkEditParamsBodyAction = "log"
	RuleBulkEditParamsBodyActionBlock RuleBulkEditParamsBodyAction = "block"
)

func (r RuleBulkEditParamsBodyAction) IsKnown() bool {
	switch r {
	case RuleBulkEditParamsBodyActionLog, RuleBulkEditParamsBodyActionBlock:
		return true
	}
	return false
}

// Update rule order among zone rules.
type RuleBulkEditParamsBodyPosition struct {
	// Move rule to after rule with this ID.
	After param.Field[string] `json:"after" format:"uuid"`
	// Move rule to before rule with this ID.
	Before param.Field[string] `json:"before" format:"uuid"`
	// Move rule to this position
	Index param.Field[int64] `json:"index"`
}

func (r RuleBulkEditParamsBodyPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleBulkEditParamsBodyPosition) implementsRuleBulkEditParamsBodyPositionUnion() {}

// Update rule order among zone rules.
//
// Satisfied by [token_validation.RuleBulkEditParamsBodyPositionAPIShieldIndex],
// [token_validation.RuleBulkEditParamsBodyPositionAPIShieldBefore],
// [token_validation.RuleBulkEditParamsBodyPositionAPIShieldAfter],
// [RuleBulkEditParamsBodyPosition].
type RuleBulkEditParamsBodyPositionUnion interface {
	implementsRuleBulkEditParamsBodyPositionUnion()
}

type RuleBulkEditParamsBodyPositionAPIShieldIndex struct {
	// Move rule to this position
	Index param.Field[int64] `json:"index,required"`
}

func (r RuleBulkEditParamsBodyPositionAPIShieldIndex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleBulkEditParamsBodyPositionAPIShieldIndex) implementsRuleBulkEditParamsBodyPositionUnion() {
}

// Move rule to after rule with ID.
type RuleBulkEditParamsBodyPositionAPIShieldBefore struct {
	// Move rule to before rule with this ID.
	Before param.Field[string] `json:"before" format:"uuid"`
}

func (r RuleBulkEditParamsBodyPositionAPIShieldBefore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleBulkEditParamsBodyPositionAPIShieldBefore) implementsRuleBulkEditParamsBodyPositionUnion() {
}

// Move rule to before rule with ID.
type RuleBulkEditParamsBodyPositionAPIShieldAfter struct {
	// Move rule to after rule with this ID.
	After param.Field[string] `json:"after" format:"uuid"`
}

func (r RuleBulkEditParamsBodyPositionAPIShieldAfter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleBulkEditParamsBodyPositionAPIShieldAfter) implementsRuleBulkEditParamsBodyPositionUnion() {
}

// Select operations covered by this rule.
//
// For details on selectors, see the
// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
type RuleBulkEditParamsBodySelector struct {
	// Ignore operations that were otherwise included by `include`.
	Exclude param.Field[[]RuleBulkEditParamsBodySelectorExclude] `json:"exclude"`
	// Select all matching operations.
	Include param.Field[[]RuleBulkEditParamsBodySelectorInclude] `json:"include"`
}

func (r RuleBulkEditParamsBodySelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleBulkEditParamsBodySelectorExclude struct {
	// Excluded operation IDs.
	OperationIDs param.Field[[]string] `json:"operation_ids"`
}

func (r RuleBulkEditParamsBodySelectorExclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleBulkEditParamsBodySelectorInclude struct {
	// Included hostnames.
	Host param.Field[[]string] `json:"host" format:"hostname"`
}

func (r RuleBulkEditParamsBodySelectorInclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleEditParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Action to take on requests that match operations included in `selector` and fail
	// `expression`.
	Action param.Field[RuleEditParamsAction] `json:"action"`
	// A human-readable description that gives more details than `title`.
	Description param.Field[string] `json:"description"`
	// Toggle rule on or off.
	Enabled param.Field[bool] `json:"enabled"`
	// Rule expression. Requests that fail to match this expression will be subject to
	// `action`.
	//
	// For details on expressions, see the
	// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
	Expression param.Field[string] `json:"expression"`
	// Update rule order among zone rules.
	Position param.Field[RuleEditParamsPositionUnion] `json:"position"`
	// Select operations covered by this rule.
	//
	// For details on selectors, see the
	// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
	Selector param.Field[RuleEditParamsSelector] `json:"selector"`
	// A human-readable name for the rule.
	Title param.Field[string] `json:"title"`
}

func (r RuleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Action to take on requests that match operations included in `selector` and fail
// `expression`.
type RuleEditParamsAction string

const (
	RuleEditParamsActionLog   RuleEditParamsAction = "log"
	RuleEditParamsActionBlock RuleEditParamsAction = "block"
)

func (r RuleEditParamsAction) IsKnown() bool {
	switch r {
	case RuleEditParamsActionLog, RuleEditParamsActionBlock:
		return true
	}
	return false
}

// Update rule order among zone rules.
type RuleEditParamsPosition struct {
	// Move rule to after rule with this ID.
	After param.Field[string] `json:"after" format:"uuid"`
	// Move rule to before rule with this ID.
	Before param.Field[string] `json:"before" format:"uuid"`
	// Move rule to this position
	Index param.Field[int64] `json:"index"`
}

func (r RuleEditParamsPosition) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPosition) implementsRuleEditParamsPositionUnion() {}

// Update rule order among zone rules.
//
// Satisfied by [token_validation.RuleEditParamsPositionAPIShieldIndex],
// [token_validation.RuleEditParamsPositionAPIShieldBefore],
// [token_validation.RuleEditParamsPositionAPIShieldAfter],
// [RuleEditParamsPosition].
type RuleEditParamsPositionUnion interface {
	implementsRuleEditParamsPositionUnion()
}

type RuleEditParamsPositionAPIShieldIndex struct {
	// Move rule to this position
	Index param.Field[int64] `json:"index,required"`
}

func (r RuleEditParamsPositionAPIShieldIndex) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionAPIShieldIndex) implementsRuleEditParamsPositionUnion() {}

// Move rule to after rule with ID.
type RuleEditParamsPositionAPIShieldBefore struct {
	// Move rule to before rule with this ID.
	Before param.Field[string] `json:"before" format:"uuid"`
}

func (r RuleEditParamsPositionAPIShieldBefore) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionAPIShieldBefore) implementsRuleEditParamsPositionUnion() {}

// Move rule to before rule with ID.
type RuleEditParamsPositionAPIShieldAfter struct {
	// Move rule to after rule with this ID.
	After param.Field[string] `json:"after" format:"uuid"`
}

func (r RuleEditParamsPositionAPIShieldAfter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r RuleEditParamsPositionAPIShieldAfter) implementsRuleEditParamsPositionUnion() {}

// Select operations covered by this rule.
//
// For details on selectors, see the
// [Cloudflare Docs](https://developers.cloudflare.com/api-shield/security/jwt-validation/).
type RuleEditParamsSelector struct {
	// Ignore operations that were otherwise included by `include`.
	Exclude param.Field[[]RuleEditParamsSelectorExclude] `json:"exclude"`
	// Select all matching operations.
	Include param.Field[[]RuleEditParamsSelectorInclude] `json:"include"`
}

func (r RuleEditParamsSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleEditParamsSelectorExclude struct {
	// Excluded operation IDs.
	OperationIDs param.Field[[]string] `json:"operation_ids"`
}

func (r RuleEditParamsSelectorExclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleEditParamsSelectorInclude struct {
	// Included hostnames.
	Host param.Field[[]string] `json:"host" format:"hostname"`
}

func (r RuleEditParamsSelectorInclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RuleEditResponseEnvelope struct {
	Errors   api_gateway.Message `json:"errors,required"`
	Messages api_gateway.Message `json:"messages,required"`
	// A Token Validation rule that can enforce security policies using JWT Tokens.
	Result TokenValidationRule `json:"result,required"`
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

type RuleGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type RuleGetResponseEnvelope struct {
	Errors   api_gateway.Message `json:"errors,required"`
	Messages api_gateway.Message `json:"messages,required"`
	// A Token Validation rule that can enforce security policies using JWT Tokens.
	Result TokenValidationRule `json:"result,required"`
	// Whether the API call was successful.
	Success RuleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    ruleGetResponseEnvelopeJSON    `json:"-"`
}

// ruleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RuleGetResponseEnvelope]
type ruleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ruleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type RuleGetResponseEnvelopeSuccess bool

const (
	RuleGetResponseEnvelopeSuccessTrue RuleGetResponseEnvelopeSuccess = true
)

func (r RuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case RuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
