// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagerules

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
	"github.com/tidwall/gjson"
)

// PageruleService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPageruleService] method instead.
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
//
// Deprecated: The Page Rules API is deprecated in favour of the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#page-rules
// for full details.
func (r *PageruleService) New(ctx context.Context, params PageruleNewParams, opts ...option.RequestOption) (res *PageruleNewResponseUnion, err error) {
	var env PageruleNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
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
//
// Deprecated: The Page Rules API is deprecated in favour of the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#page-rules
// for full details.
func (r *PageruleService) Update(ctx context.Context, pageruleID string, params PageruleUpdateParams, opts ...option.RequestOption) (res *PageruleUpdateResponseUnion, err error) {
	var env PageruleUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", params.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches Page Rules in a zone.
//
// Deprecated: The Page Rules API is deprecated in favour of the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#page-rules
// for full details.
func (r *PageruleService) List(ctx context.Context, params PageruleListParams, opts ...option.RequestOption) (res *[]PageRule, err error) {
	var env PageruleListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an existing Page Rule.
//
// Deprecated: The Page Rules API is deprecated in favour of the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#page-rules
// for full details.
func (r *PageruleService) Delete(ctx context.Context, pageruleID string, body PageruleDeleteParams, opts ...option.RequestOption) (res *PageruleDeleteResponse, err error) {
	var env PageruleDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", body.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates one or more fields of an existing Page Rule.
//
// Deprecated: The Page Rules API is deprecated in favour of the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#page-rules
// for full details.
func (r *PageruleService) Edit(ctx context.Context, pageruleID string, params PageruleEditParams, opts ...option.RequestOption) (res *PageruleEditResponseUnion, err error) {
	var env PageruleEditResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", params.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches the details of a Page Rule.
//
// Deprecated: The Page Rules API is deprecated in favour of the Ruleset Engine.
// See
// https://developers.cloudflare.com/fundamentals/api/reference/deprecations/#page-rules
// for full details.
func (r *PageruleService) Get(ctx context.Context, pageruleID string, query PageruleGetParams, opts ...option.RequestOption) (res *PageruleGetResponseUnion, err error) {
	var env PageruleGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if pageruleID == "" {
		err = errors.New("missing required pagerule_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/pagerules/%s", query.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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
	Targets []Target     `json:"targets,required"`
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

type Route struct {
	// The timestamp of when the override was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The type of route.
	Name  RouteName  `json:"name"`
	Value RouteValue `json:"value"`
	JSON  routeJSON  `json:"-"`
}

// routeJSON contains the JSON metadata for the struct [Route]
type routeJSON struct {
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Route) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeJSON) RawJSON() string {
	return r.raw
}

// The type of route.
type RouteName string

const (
	RouteNameForwardURL RouteName = "forward_url"
)

func (r RouteName) IsKnown() bool {
	switch r {
	case RouteNameForwardURL:
		return true
	}
	return false
}

type RouteValue struct {
	// The response type for the URL redirect.
	Type RouteValueType `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string         `json:"url"`
	JSON routeValueJSON `json:"-"`
}

// routeValueJSON contains the JSON metadata for the struct [RouteValue]
type routeValueJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeValueJSON) RawJSON() string {
	return r.raw
}

// The response type for the URL redirect.
type RouteValueType string

const (
	RouteValueTypeTemporary RouteValueType = "temporary"
	RouteValueTypePermanent RouteValueType = "permanent"
)

func (r RouteValueType) IsKnown() bool {
	switch r {
	case RouteValueTypeTemporary, RouteValueTypePermanent:
		return true
	}
	return false
}

type RouteParam struct {
	// The type of route.
	Name  param.Field[RouteName]       `json:"name"`
	Value param.Field[RouteValueParam] `json:"value"`
}

func (r RouteParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RouteValueParam struct {
	// The response type for the URL redirect.
	Type param.Field[RouteValueType] `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r RouteValueParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URL target.
type Target struct {
	// String constraint.
	Constraint TargetConstraint `json:"constraint"`
	// A target based on the URL of the request.
	Target TargetTarget `json:"target"`
	JSON   targetJSON   `json:"-"`
}

// targetJSON contains the JSON metadata for the struct [Target]
type targetJSON struct {
	Constraint  apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Target) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r targetJSON) RawJSON() string {
	return r.raw
}

// String constraint.
type TargetConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator TargetConstraintOperator `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value string               `json:"value,required"`
	JSON  targetConstraintJSON `json:"-"`
}

// targetConstraintJSON contains the JSON metadata for the struct
// [TargetConstraint]
type targetConstraintJSON struct {
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TargetConstraint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r targetConstraintJSON) RawJSON() string {
	return r.raw
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type TargetConstraintOperator string

const (
	TargetConstraintOperatorMatches    TargetConstraintOperator = "matches"
	TargetConstraintOperatorContains   TargetConstraintOperator = "contains"
	TargetConstraintOperatorEquals     TargetConstraintOperator = "equals"
	TargetConstraintOperatorNotEqual   TargetConstraintOperator = "not_equal"
	TargetConstraintOperatorNotContain TargetConstraintOperator = "not_contain"
)

func (r TargetConstraintOperator) IsKnown() bool {
	switch r {
	case TargetConstraintOperatorMatches, TargetConstraintOperatorContains, TargetConstraintOperatorEquals, TargetConstraintOperatorNotEqual, TargetConstraintOperatorNotContain:
		return true
	}
	return false
}

// A target based on the URL of the request.
type TargetTarget string

const (
	TargetTargetURL TargetTarget = "url"
)

func (r TargetTarget) IsKnown() bool {
	switch r {
	case TargetTargetURL:
		return true
	}
	return false
}

// URL target.
type TargetParam struct {
	// String constraint.
	Constraint param.Field[TargetConstraintParam] `json:"constraint"`
	// A target based on the URL of the request.
	Target param.Field[TargetTarget] `json:"target"`
}

func (r TargetParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// String constraint.
type TargetConstraintParam struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator param.Field[TargetConstraintOperator] `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value param.Field[string] `json:"value,required"`
}

func (r TargetConstraintParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Union satisfied by [pagerules.PageruleNewResponseUnknown] or
// [shared.UnionString].
type PageruleNewResponseUnion interface {
	ImplementsPagerulesPageruleNewResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageruleNewResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [pagerules.PageruleUpdateResponseUnknown] or
// [shared.UnionString].
type PageruleUpdateResponseUnion interface {
	ImplementsPagerulesPageruleUpdateResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageruleUpdateResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
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

// Union satisfied by [pagerules.PageruleEditResponseUnknown] or
// [shared.UnionString].
type PageruleEditResponseUnion interface {
	ImplementsPagerulesPageruleEditResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageruleEditResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [pagerules.PageruleGetResponseUnknown] or
// [shared.UnionString].
type PageruleGetResponseUnion interface {
	ImplementsPagerulesPageruleGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageruleGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type PageruleNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]RouteParam] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargetParam] `json:"targets,required"`
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
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   PageruleNewResponseUnion `json:"result,required"`
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
	Actions param.Field[[]RouteParam] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargetParam] `json:"targets,required"`
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
	Errors   []shared.ResponseInfo       `json:"errors,required"`
	Messages []shared.ResponseInfo       `json:"messages,required"`
	Result   PageruleUpdateResponseUnion `json:"result,required"`
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
		NestedFormat: apiquery.NestedQueryFormatDots,
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []PageRule            `json:"result,required"`
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
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageruleDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   PageruleDeleteResponse `json:"result,required,nullable"`
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
	Actions param.Field[[]RouteParam] `json:"actions"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageruleEditParamsStatus] `json:"status"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]TargetParam] `json:"targets"`
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
	Errors   []shared.ResponseInfo     `json:"errors,required"`
	Messages []shared.ResponseInfo     `json:"messages,required"`
	Result   PageruleEditResponseUnion `json:"result,required"`
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
	Errors   []shared.ResponseInfo    `json:"errors,required"`
	Messages []shared.ResponseInfo    `json:"messages,required"`
	Result   PageruleGetResponseUnion `json:"result,required"`
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
