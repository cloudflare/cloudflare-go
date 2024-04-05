// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagerules

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// PageruleService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewPageruleService] method instead.
type PageruleService struct {
	Options  []option.RequestOption
	Settings *SettingService
}

// NewPageruleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPageruleService(opts ...option.RequestOption) (r *PageruleService) {
	r = &PageruleService{}
	r.Options = opts
	r.Settings = NewSettingService(opts...)
	return
}

// Creates a new Page Rule.
func (r *PageruleService) New(ctx context.Context, params PageruleNewParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env PageruleNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/pagerules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Replaces the configuration of an existing Page Rule. The configuration of the
// updated Page Rule will exactly match the data passed in the API request.
func (r *PageruleService) Update(ctx context.Context, pageruleID string, params PageruleUpdateParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env PageruleUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/pagerules/%s", params.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Page Rules in a zone.
func (r *PageruleService) List(ctx context.Context, params PageruleListParams, opts ...option.RequestOption) (res *[]PageRule, err error) {
	opts = append(r.Options[:], opts...)
	var env PageruleListResponseEnvelope
	path := fmt.Sprintf("zones/%s/pagerules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing Page Rule.
func (r *PageruleService) Delete(ctx context.Context, pageruleID string, params PageruleDeleteParams, opts ...option.RequestOption) (res *PageruleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageruleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/pagerules/%s", params.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates one or more fields of an existing Page Rule.
func (r *PageruleService) Edit(ctx context.Context, pageruleID string, params PageruleEditParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env PageruleEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/pagerules/%s", params.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a Page Rule.
func (r *PageruleService) Get(ctx context.Context, pageruleID string, query PageruleGetParams, opts ...option.RequestOption) (res *shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env PageruleGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/pagerules/%s", query.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ActionItem struct {
	// The timestamp of when the override was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The type of route.
	Name  ActionItemName  `json:"name"`
	Value ActionItemValue `json:"value"`
	JSON  actionItemJSON  `json:"-"`
}

// actionItemJSON contains the JSON metadata for the struct [ActionItem]
type actionItemJSON struct {
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActionItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r actionItemJSON) RawJSON() string {
	return r.raw
}

// The type of route.
type ActionItemName string

const (
	ActionItemNameForwardURL ActionItemName = "forward_url"
)

func (r ActionItemName) IsKnown() bool {
	switch r {
	case ActionItemNameForwardURL:
		return true
	}
	return false
}

type ActionItemValue struct {
	// The response type for the URL redirect.
	Type ActionItemValueType `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string              `json:"url"`
	JSON actionItemValueJSON `json:"-"`
}

// actionItemValueJSON contains the JSON metadata for the struct [ActionItemValue]
type actionItemValueJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ActionItemValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r actionItemValueJSON) RawJSON() string {
	return r.raw
}

// The response type for the URL redirect.
type ActionItemValueType string

const (
	ActionItemValueTypeTemporary ActionItemValueType = "temporary"
	ActionItemValueTypePermanent ActionItemValueType = "permanent"
)

func (r ActionItemValueType) IsKnown() bool {
	switch r {
	case ActionItemValueTypeTemporary, ActionItemValueTypePermanent:
		return true
	}
	return false
}

type ActionItemParam struct {
	// The type of route.
	Name  param.Field[ActionItemName]       `json:"name"`
	Value param.Field[ActionItemValueParam] `json:"value"`
}

func (r ActionItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ActionItemValueParam struct {
	// The response type for the URL redirect.
	Type param.Field[ActionItemValueType] `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r ActionItemValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageRule struct {
	// Identifier
	ID string `json:"id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions []Route `json:"actions,required"`
	// The timestamp of when the Page Rule was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// The timestamp of when the Page Rule was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority int64 `json:"priority,required"`
	// The status of the Page Rule.
	Status PageRuleStatus `json:"status,required"`
	// The rule targets to evaluate on each request.
	Targets []URLTarget  `json:"targets,required"`
	JSON    pageRuleJSON `json:"-"`
}

// pageRuleJSON contains the JSON metadata for the struct [PageRule]
type pageRuleJSON struct {
	ID          apijson.Field
	Actions     apijson.Field
	CreatedOn   apijson.Field
	ModifiedOn  apijson.Field
	Priority    apijson.Field
	Status      apijson.Field
	Targets     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageRuleJSON) RawJSON() string {
	return r.raw
}

// The status of the Page Rule.
type PageRuleStatus string

const (
	PageRuleStatusActive   PageRuleStatus = "active"
	PageRuleStatusDisabled PageRuleStatus = "disabled"
)

func (r PageRuleStatus) IsKnown() bool {
	switch r {
	case PageRuleStatusActive, PageRuleStatusDisabled:
		return true
	}
	return false
}

// A request condition target.
type TargesItemParam struct {
	// String constraint.
	Constraint param.Field[TargesItemConstraintParam] `json:"constraint,required"`
	// A target based on the URL of the request.
	Target param.Field[TargesItemTarget] `json:"target,required"`
}

func (r TargesItemParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// String constraint.
type TargesItemConstraintParam struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator param.Field[TargesItemConstraintOperator] `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value param.Field[string] `json:"value,required"`
}

func (r TargesItemConstraintParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type TargesItemConstraintOperator string

const (
	TargesItemConstraintOperatorMatches    TargesItemConstraintOperator = "matches"
	TargesItemConstraintOperatorContains   TargesItemConstraintOperator = "contains"
	TargesItemConstraintOperatorEquals     TargesItemConstraintOperator = "equals"
	TargesItemConstraintOperatorNotEqual   TargesItemConstraintOperator = "not_equal"
	TargesItemConstraintOperatorNotContain TargesItemConstraintOperator = "not_contain"
)

func (r TargesItemConstraintOperator) IsKnown() bool {
	switch r {
	case TargesItemConstraintOperatorMatches, TargesItemConstraintOperatorContains, TargesItemConstraintOperatorEquals, TargesItemConstraintOperatorNotEqual, TargesItemConstraintOperatorNotContain:
		return true
	}
	return false
}

// A target based on the URL of the request.
type TargesItemTarget string

const (
	TargesItemTargetURL TargesItemTarget = "url"
)

func (r TargesItemTarget) IsKnown() bool {
	switch r {
	case TargesItemTargetURL:
		return true
	}
	return false
}

// URL target.
type URLTarget struct {
	// String constraint.
	Constraint URLTargetConstraint `json:"constraint"`
	// A target based on the URL of the request.
	Target URLTargetTarget `json:"target"`
	JSON   urlTargetJSON   `json:"-"`
}

// urlTargetJSON contains the JSON metadata for the struct [URLTarget]
type urlTargetJSON struct {
	Constraint  apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlTargetJSON) RawJSON() string {
	return r.raw
}

// String constraint.
type URLTargetConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator URLTargetConstraintOperator `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value string                  `json:"value,required"`
	JSON  urlTargetConstraintJSON `json:"-"`
}

// urlTargetConstraintJSON contains the JSON metadata for the struct
// [URLTargetConstraint]
type urlTargetConstraintJSON struct {
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *URLTargetConstraint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r urlTargetConstraintJSON) RawJSON() string {
	return r.raw
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type URLTargetConstraintOperator string

const (
	URLTargetConstraintOperatorMatches    URLTargetConstraintOperator = "matches"
	URLTargetConstraintOperatorContains   URLTargetConstraintOperator = "contains"
	URLTargetConstraintOperatorEquals     URLTargetConstraintOperator = "equals"
	URLTargetConstraintOperatorNotEqual   URLTargetConstraintOperator = "not_equal"
	URLTargetConstraintOperatorNotContain URLTargetConstraintOperator = "not_contain"
)

func (r URLTargetConstraintOperator) IsKnown() bool {
	switch r {
	case URLTargetConstraintOperatorMatches, URLTargetConstraintOperatorContains, URLTargetConstraintOperatorEquals, URLTargetConstraintOperatorNotEqual, URLTargetConstraintOperatorNotContain:
		return true
	}
	return false
}

// A target based on the URL of the request.
type URLTargetTarget string

const (
	URLTargetTargetURL URLTargetTarget = "url"
)

func (r URLTargetTarget) IsKnown() bool {
	switch r {
	case URLTargetTargetURL:
		return true
	}
	return false
}

type PageruleDeleteResponse struct {
	// Identifier
	ID   string                     `json:"id,required"`
	JSON pageruleDeleteResponseJSON `json:"-"`
}

// pageruleDeleteResponseJSON contains the JSON metadata for the struct
// [PageruleDeleteResponse]
type pageruleDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type PageruleNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]ActionItemParam] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargesItemParam] `json:"targets,required"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageruleNewParamsStatus] `json:"status"`
}

func (r PageruleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the Page Rule.
type PageruleNewParamsStatus string

const (
	PageruleNewParamsStatusActive   PageruleNewParamsStatus = "active"
	PageruleNewParamsStatusDisabled PageruleNewParamsStatus = "disabled"
)

func (r PageruleNewParamsStatus) IsKnown() bool {
	switch r {
	case PageruleNewParamsStatusActive, PageruleNewParamsStatusDisabled:
		return true
	}
	return false
}

type PageruleNewResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion `json:"result,required"`
	// Whether the API call was successful
	Success PageruleNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageruleNewResponseEnvelopeJSON    `json:"-"`
}

// pageruleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleNewResponseEnvelope]
type pageruleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleNewResponseEnvelopeSuccess bool

const (
	PageruleNewResponseEnvelopeSuccessTrue PageruleNewResponseEnvelopeSuccess = true
)

func (r PageruleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleUpdateParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]ActionItemParam] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargesItemParam] `json:"targets,required"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageruleUpdateParamsStatus] `json:"status"`
}

func (r PageruleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the Page Rule.
type PageruleUpdateParamsStatus string

const (
	PageruleUpdateParamsStatusActive   PageruleUpdateParamsStatus = "active"
	PageruleUpdateParamsStatusDisabled PageruleUpdateParamsStatus = "disabled"
)

func (r PageruleUpdateParamsStatus) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsStatusActive, PageruleUpdateParamsStatusDisabled:
		return true
	}
	return false
}

type PageruleUpdateResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion `json:"result,required"`
	// Whether the API call was successful
	Success PageruleUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageruleUpdateResponseEnvelopeJSON    `json:"-"`
}

// pageruleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleUpdateResponseEnvelope]
type pageruleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleUpdateResponseEnvelopeSuccess bool

const (
	PageruleUpdateResponseEnvelopeSuccessTrue PageruleUpdateResponseEnvelopeSuccess = true
)

func (r PageruleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The direction used to sort returned Page Rules.
	Direction param.Field[PageruleListParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[PageruleListParamsMatch] `query:"match"`
	// The field used to sort returned Page Rules.
	Order param.Field[PageruleListParamsOrder] `query:"order"`
	// The status of the Page Rule.
	Status param.Field[PageruleListParamsStatus] `query:"status"`
}

// URLQuery serializes [PageruleListParams]'s query parameters as `url.Values`.
func (r PageruleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned Page Rules.
type PageruleListParamsDirection string

const (
	PageruleListParamsDirectionAsc  PageruleListParamsDirection = "asc"
	PageruleListParamsDirectionDesc PageruleListParamsDirection = "desc"
)

func (r PageruleListParamsDirection) IsKnown() bool {
	switch r {
	case PageruleListParamsDirectionAsc, PageruleListParamsDirectionDesc:
		return true
	}
	return false
}

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type PageruleListParamsMatch string

const (
	PageruleListParamsMatchAny PageruleListParamsMatch = "any"
	PageruleListParamsMatchAll PageruleListParamsMatch = "all"
)

func (r PageruleListParamsMatch) IsKnown() bool {
	switch r {
	case PageruleListParamsMatchAny, PageruleListParamsMatchAll:
		return true
	}
	return false
}

// The field used to sort returned Page Rules.
type PageruleListParamsOrder string

const (
	PageruleListParamsOrderStatus   PageruleListParamsOrder = "status"
	PageruleListParamsOrderPriority PageruleListParamsOrder = "priority"
)

func (r PageruleListParamsOrder) IsKnown() bool {
	switch r {
	case PageruleListParamsOrderStatus, PageruleListParamsOrderPriority:
		return true
	}
	return false
}

// The status of the Page Rule.
type PageruleListParamsStatus string

const (
	PageruleListParamsStatusActive   PageruleListParamsStatus = "active"
	PageruleListParamsStatusDisabled PageruleListParamsStatus = "disabled"
)

func (r PageruleListParamsStatus) IsKnown() bool {
	switch r {
	case PageruleListParamsStatusActive, PageruleListParamsStatusDisabled:
		return true
	}
	return false
}

type PageruleListResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   []PageRule                                                `json:"result,required"`
	// Whether the API call was successful
	Success PageruleListResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageruleListResponseEnvelopeJSON    `json:"-"`
}

// pageruleListResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleListResponseEnvelope]
type pageruleListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleListResponseEnvelopeSuccess bool

const (
	PageruleListResponseEnvelopeSuccessTrue PageruleListResponseEnvelopeSuccess = true
)

func (r PageruleListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleDeleteParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r PageruleDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type PageruleDeleteResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72 `json:"messages,required"`
	Result   PageruleDeleteResponse                                    `json:"result,required,nullable"`
	// Whether the API call was successful
	Success PageruleDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageruleDeleteResponseEnvelopeJSON    `json:"-"`
}

// pageruleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleDeleteResponseEnvelope]
type pageruleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleDeleteResponseEnvelopeSuccess bool

const (
	PageruleDeleteResponseEnvelopeSuccessTrue PageruleDeleteResponseEnvelopeSuccess = true
)

func (r PageruleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]ActionItemParam] `json:"actions"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageruleEditParamsStatus] `json:"status"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargesItemParam] `json:"targets"`
}

func (r PageruleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The status of the Page Rule.
type PageruleEditParamsStatus string

const (
	PageruleEditParamsStatusActive   PageruleEditParamsStatus = "active"
	PageruleEditParamsStatusDisabled PageruleEditParamsStatus = "disabled"
)

func (r PageruleEditParamsStatus) IsKnown() bool {
	switch r {
	case PageruleEditParamsStatusActive, PageruleEditParamsStatusDisabled:
		return true
	}
	return false
}

type PageruleEditResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion `json:"result,required"`
	// Whether the API call was successful
	Success PageruleEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageruleEditResponseEnvelopeJSON    `json:"-"`
}

// pageruleEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleEditResponseEnvelope]
type pageruleEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleEditResponseEnvelopeSuccess bool

const (
	PageruleEditResponseEnvelopeSuccessTrue PageruleEditResponseEnvelopeSuccess = true
)

func (r PageruleEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type PageruleGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageruleGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef3248f24329456e19dfa042fff9986f72    `json:"messages,required"`
	Result   shared.UnnamedSchemaRef9444735ca60712dbcf8afd832eb5716aUnion `json:"result,required"`
	// Whether the API call was successful
	Success PageruleGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    pageruleGetResponseEnvelopeJSON    `json:"-"`
}

// pageruleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PageruleGetResponseEnvelope]
type pageruleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleGetResponseEnvelopeSuccess bool

const (
	PageruleGetResponseEnvelopeSuccessTrue PageruleGetResponseEnvelopeSuccess = true
)

func (r PageruleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PageruleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
