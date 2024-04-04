// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagerules

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
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
func (r *PageruleService) New(ctx context.Context, params PageruleNewParams, opts ...option.RequestOption) (res *PageruleNewResponse, err error) {
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
func (r *PageruleService) Update(ctx context.Context, pageruleID string, params PageruleUpdateParams, opts ...option.RequestOption) (res *PageruleUpdateResponse, err error) {
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
func (r *PageruleService) List(ctx context.Context, params PageruleListParams, opts ...option.RequestOption) (res *[]ZonesPagerule, err error) {
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
func (r *PageruleService) Edit(ctx context.Context, pageruleID string, params PageruleEditParams, opts ...option.RequestOption) (res *PageruleEditResponse, err error) {
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
func (r *PageruleService) Get(ctx context.Context, pageruleID string, query PageruleGetParams, opts ...option.RequestOption) (res *PageruleGetResponse, err error) {
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

type ZonesPagerule struct {
	// Identifier
	ID string `json:"id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions []ZonesPageruleAction `json:"actions,required"`
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
	Status ZonesPageruleStatus `json:"status,required"`
	// The rule targets to evaluate on each request.
	Targets []ZonesPageruleTarget `json:"targets,required"`
	JSON    zonesPageruleJSON     `json:"-"`
}

// zonesPageruleJSON contains the JSON metadata for the struct [ZonesPagerule]
type zonesPageruleJSON struct {
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

func (r *ZonesPagerule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesPageruleJSON) RawJSON() string {
	return r.raw
}

type ZonesPageruleAction struct {
	// The timestamp of when the override was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The type of route.
	Name  ZonesPageruleActionsName  `json:"name"`
	Value ZonesPageruleActionsValue `json:"value"`
	JSON  zonesPageruleActionJSON   `json:"-"`
}

// zonesPageruleActionJSON contains the JSON metadata for the struct
// [ZonesPageruleAction]
type zonesPageruleActionJSON struct {
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesPageruleAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesPageruleActionJSON) RawJSON() string {
	return r.raw
}

// The type of route.
type ZonesPageruleActionsName string

const (
	ZonesPageruleActionsNameForwardURL ZonesPageruleActionsName = "forward_url"
)

func (r ZonesPageruleActionsName) IsKnown() bool {
	switch r {
	case ZonesPageruleActionsNameForwardURL:
		return true
	}
	return false
}

type ZonesPageruleActionsValue struct {
	// The response type for the URL redirect.
	Type ZonesPageruleActionsValueType `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string                        `json:"url"`
	JSON zonesPageruleActionsValueJSON `json:"-"`
}

// zonesPageruleActionsValueJSON contains the JSON metadata for the struct
// [ZonesPageruleActionsValue]
type zonesPageruleActionsValueJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesPageruleActionsValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesPageruleActionsValueJSON) RawJSON() string {
	return r.raw
}

// The response type for the URL redirect.
type ZonesPageruleActionsValueType string

const (
	ZonesPageruleActionsValueTypeTemporary ZonesPageruleActionsValueType = "temporary"
	ZonesPageruleActionsValueTypePermanent ZonesPageruleActionsValueType = "permanent"
)

func (r ZonesPageruleActionsValueType) IsKnown() bool {
	switch r {
	case ZonesPageruleActionsValueTypeTemporary, ZonesPageruleActionsValueTypePermanent:
		return true
	}
	return false
}

// The status of the Page Rule.
type ZonesPageruleStatus string

const (
	ZonesPageruleStatusActive   ZonesPageruleStatus = "active"
	ZonesPageruleStatusDisabled ZonesPageruleStatus = "disabled"
)

func (r ZonesPageruleStatus) IsKnown() bool {
	switch r {
	case ZonesPageruleStatusActive, ZonesPageruleStatusDisabled:
		return true
	}
	return false
}

// A request condition target.
type ZonesPageruleTarget struct {
	// String constraint.
	Constraint ZonesPageruleTargetsConstraint `json:"constraint,required"`
	// A target based on the URL of the request.
	Target ZonesPageruleTargetsTarget `json:"target,required"`
	JSON   zonesPageruleTargetJSON    `json:"-"`
}

// zonesPageruleTargetJSON contains the JSON metadata for the struct
// [ZonesPageruleTarget]
type zonesPageruleTargetJSON struct {
	Constraint  apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesPageruleTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesPageruleTargetJSON) RawJSON() string {
	return r.raw
}

// String constraint.
type ZonesPageruleTargetsConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator ZonesPageruleTargetsConstraintOperator `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value string                             `json:"value,required"`
	JSON  zonesPageruleTargetsConstraintJSON `json:"-"`
}

// zonesPageruleTargetsConstraintJSON contains the JSON metadata for the struct
// [ZonesPageruleTargetsConstraint]
type zonesPageruleTargetsConstraintJSON struct {
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonesPageruleTargetsConstraint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zonesPageruleTargetsConstraintJSON) RawJSON() string {
	return r.raw
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type ZonesPageruleTargetsConstraintOperator string

const (
	ZonesPageruleTargetsConstraintOperatorMatches    ZonesPageruleTargetsConstraintOperator = "matches"
	ZonesPageruleTargetsConstraintOperatorContains   ZonesPageruleTargetsConstraintOperator = "contains"
	ZonesPageruleTargetsConstraintOperatorEquals     ZonesPageruleTargetsConstraintOperator = "equals"
	ZonesPageruleTargetsConstraintOperatorNotEqual   ZonesPageruleTargetsConstraintOperator = "not_equal"
	ZonesPageruleTargetsConstraintOperatorNotContain ZonesPageruleTargetsConstraintOperator = "not_contain"
)

func (r ZonesPageruleTargetsConstraintOperator) IsKnown() bool {
	switch r {
	case ZonesPageruleTargetsConstraintOperatorMatches, ZonesPageruleTargetsConstraintOperatorContains, ZonesPageruleTargetsConstraintOperatorEquals, ZonesPageruleTargetsConstraintOperatorNotEqual, ZonesPageruleTargetsConstraintOperatorNotContain:
		return true
	}
	return false
}

// A target based on the URL of the request.
type ZonesPageruleTargetsTarget string

const (
	ZonesPageruleTargetsTargetURL ZonesPageruleTargetsTarget = "url"
)

func (r ZonesPageruleTargetsTarget) IsKnown() bool {
	switch r {
	case ZonesPageruleTargetsTargetURL:
		return true
	}
	return false
}

// Union satisfied by [pagerules.PageruleNewResponseUnknown] or
// [shared.UnionString].
type PageruleNewResponse interface {
	ImplementsPagerulesPageruleNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageruleNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [pagerules.PageruleUpdateResponseUnknown] or
// [shared.UnionString].
type PageruleUpdateResponse interface {
	ImplementsPagerulesPageruleUpdateResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageruleUpdateResponse)(nil)).Elem(),
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
type PageruleEditResponse interface {
	ImplementsPagerulesPageruleEditResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageruleEditResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

// Union satisfied by [pagerules.PageruleGetResponseUnknown] or
// [shared.UnionString].
type PageruleGetResponse interface {
	ImplementsPagerulesPageruleGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*PageruleGetResponse)(nil)).Elem(),
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
	Actions param.Field[[]PageruleNewParamsAction] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]PageruleNewParamsTarget] `json:"targets,required"`
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

type PageruleNewParamsAction struct {
	// The type of route.
	Name  param.Field[PageruleNewParamsActionsName]  `json:"name"`
	Value param.Field[PageruleNewParamsActionsValue] `json:"value"`
}

func (r PageruleNewParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of route.
type PageruleNewParamsActionsName string

const (
	PageruleNewParamsActionsNameForwardURL PageruleNewParamsActionsName = "forward_url"
)

func (r PageruleNewParamsActionsName) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsNameForwardURL:
		return true
	}
	return false
}

type PageruleNewParamsActionsValue struct {
	// The response type for the URL redirect.
	Type param.Field[PageruleNewParamsActionsValueType] `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageruleNewParamsActionsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response type for the URL redirect.
type PageruleNewParamsActionsValueType string

const (
	PageruleNewParamsActionsValueTypeTemporary PageruleNewParamsActionsValueType = "temporary"
	PageruleNewParamsActionsValueTypePermanent PageruleNewParamsActionsValueType = "permanent"
)

func (r PageruleNewParamsActionsValueType) IsKnown() bool {
	switch r {
	case PageruleNewParamsActionsValueTypeTemporary, PageruleNewParamsActionsValueTypePermanent:
		return true
	}
	return false
}

// A request condition target.
type PageruleNewParamsTarget struct {
	// String constraint.
	Constraint param.Field[PageruleNewParamsTargetsConstraint] `json:"constraint,required"`
	// A target based on the URL of the request.
	Target param.Field[PageruleNewParamsTargetsTarget] `json:"target,required"`
}

func (r PageruleNewParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// String constraint.
type PageruleNewParamsTargetsConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator param.Field[PageruleNewParamsTargetsConstraintOperator] `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value param.Field[string] `json:"value,required"`
}

func (r PageruleNewParamsTargetsConstraint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type PageruleNewParamsTargetsConstraintOperator string

const (
	PageruleNewParamsTargetsConstraintOperatorMatches    PageruleNewParamsTargetsConstraintOperator = "matches"
	PageruleNewParamsTargetsConstraintOperatorContains   PageruleNewParamsTargetsConstraintOperator = "contains"
	PageruleNewParamsTargetsConstraintOperatorEquals     PageruleNewParamsTargetsConstraintOperator = "equals"
	PageruleNewParamsTargetsConstraintOperatorNotEqual   PageruleNewParamsTargetsConstraintOperator = "not_equal"
	PageruleNewParamsTargetsConstraintOperatorNotContain PageruleNewParamsTargetsConstraintOperator = "not_contain"
)

func (r PageruleNewParamsTargetsConstraintOperator) IsKnown() bool {
	switch r {
	case PageruleNewParamsTargetsConstraintOperatorMatches, PageruleNewParamsTargetsConstraintOperatorContains, PageruleNewParamsTargetsConstraintOperatorEquals, PageruleNewParamsTargetsConstraintOperatorNotEqual, PageruleNewParamsTargetsConstraintOperatorNotContain:
		return true
	}
	return false
}

// A target based on the URL of the request.
type PageruleNewParamsTargetsTarget string

const (
	PageruleNewParamsTargetsTargetURL PageruleNewParamsTargetsTarget = "url"
)

func (r PageruleNewParamsTargetsTarget) IsKnown() bool {
	switch r {
	case PageruleNewParamsTargetsTargetURL:
		return true
	}
	return false
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   PageruleNewResponse   `json:"result,required"`
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
	Actions param.Field[[]PageruleUpdateParamsAction] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]PageruleUpdateParamsTarget] `json:"targets,required"`
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

type PageruleUpdateParamsAction struct {
	// The type of route.
	Name  param.Field[PageruleUpdateParamsActionsName]  `json:"name"`
	Value param.Field[PageruleUpdateParamsActionsValue] `json:"value"`
}

func (r PageruleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of route.
type PageruleUpdateParamsActionsName string

const (
	PageruleUpdateParamsActionsNameForwardURL PageruleUpdateParamsActionsName = "forward_url"
)

func (r PageruleUpdateParamsActionsName) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsNameForwardURL:
		return true
	}
	return false
}

type PageruleUpdateParamsActionsValue struct {
	// The response type for the URL redirect.
	Type param.Field[PageruleUpdateParamsActionsValueType] `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageruleUpdateParamsActionsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response type for the URL redirect.
type PageruleUpdateParamsActionsValueType string

const (
	PageruleUpdateParamsActionsValueTypeTemporary PageruleUpdateParamsActionsValueType = "temporary"
	PageruleUpdateParamsActionsValueTypePermanent PageruleUpdateParamsActionsValueType = "permanent"
)

func (r PageruleUpdateParamsActionsValueType) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsActionsValueTypeTemporary, PageruleUpdateParamsActionsValueTypePermanent:
		return true
	}
	return false
}

// A request condition target.
type PageruleUpdateParamsTarget struct {
	// String constraint.
	Constraint param.Field[PageruleUpdateParamsTargetsConstraint] `json:"constraint,required"`
	// A target based on the URL of the request.
	Target param.Field[PageruleUpdateParamsTargetsTarget] `json:"target,required"`
}

func (r PageruleUpdateParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// String constraint.
type PageruleUpdateParamsTargetsConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator param.Field[PageruleUpdateParamsTargetsConstraintOperator] `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value param.Field[string] `json:"value,required"`
}

func (r PageruleUpdateParamsTargetsConstraint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type PageruleUpdateParamsTargetsConstraintOperator string

const (
	PageruleUpdateParamsTargetsConstraintOperatorMatches    PageruleUpdateParamsTargetsConstraintOperator = "matches"
	PageruleUpdateParamsTargetsConstraintOperatorContains   PageruleUpdateParamsTargetsConstraintOperator = "contains"
	PageruleUpdateParamsTargetsConstraintOperatorEquals     PageruleUpdateParamsTargetsConstraintOperator = "equals"
	PageruleUpdateParamsTargetsConstraintOperatorNotEqual   PageruleUpdateParamsTargetsConstraintOperator = "not_equal"
	PageruleUpdateParamsTargetsConstraintOperatorNotContain PageruleUpdateParamsTargetsConstraintOperator = "not_contain"
)

func (r PageruleUpdateParamsTargetsConstraintOperator) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsTargetsConstraintOperatorMatches, PageruleUpdateParamsTargetsConstraintOperatorContains, PageruleUpdateParamsTargetsConstraintOperatorEquals, PageruleUpdateParamsTargetsConstraintOperatorNotEqual, PageruleUpdateParamsTargetsConstraintOperatorNotContain:
		return true
	}
	return false
}

// A target based on the URL of the request.
type PageruleUpdateParamsTargetsTarget string

const (
	PageruleUpdateParamsTargetsTargetURL PageruleUpdateParamsTargetsTarget = "url"
)

func (r PageruleUpdateParamsTargetsTarget) IsKnown() bool {
	switch r {
	case PageruleUpdateParamsTargetsTargetURL:
		return true
	}
	return false
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
	Errors   []shared.ResponseInfo  `json:"errors,required"`
	Messages []shared.ResponseInfo  `json:"messages,required"`
	Result   PageruleUpdateResponse `json:"result,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []ZonesPagerule       `json:"result,required"`
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
	Actions param.Field[[]PageruleEditParamsAction] `json:"actions"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[PageruleEditParamsStatus] `json:"status"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]PageruleEditParamsTarget] `json:"targets"`
}

func (r PageruleEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type PageruleEditParamsAction struct {
	// The type of route.
	Name  param.Field[PageruleEditParamsActionsName]  `json:"name"`
	Value param.Field[PageruleEditParamsActionsValue] `json:"value"`
}

func (r PageruleEditParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of route.
type PageruleEditParamsActionsName string

const (
	PageruleEditParamsActionsNameForwardURL PageruleEditParamsActionsName = "forward_url"
)

func (r PageruleEditParamsActionsName) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsNameForwardURL:
		return true
	}
	return false
}

type PageruleEditParamsActionsValue struct {
	// The response type for the URL redirect.
	Type param.Field[PageruleEditParamsActionsValueType] `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r PageruleEditParamsActionsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response type for the URL redirect.
type PageruleEditParamsActionsValueType string

const (
	PageruleEditParamsActionsValueTypeTemporary PageruleEditParamsActionsValueType = "temporary"
	PageruleEditParamsActionsValueTypePermanent PageruleEditParamsActionsValueType = "permanent"
)

func (r PageruleEditParamsActionsValueType) IsKnown() bool {
	switch r {
	case PageruleEditParamsActionsValueTypeTemporary, PageruleEditParamsActionsValueTypePermanent:
		return true
	}
	return false
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

// A request condition target.
type PageruleEditParamsTarget struct {
	// String constraint.
	Constraint param.Field[PageruleEditParamsTargetsConstraint] `json:"constraint,required"`
	// A target based on the URL of the request.
	Target param.Field[PageruleEditParamsTargetsTarget] `json:"target,required"`
}

func (r PageruleEditParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// String constraint.
type PageruleEditParamsTargetsConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator param.Field[PageruleEditParamsTargetsConstraintOperator] `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value param.Field[string] `json:"value,required"`
}

func (r PageruleEditParamsTargetsConstraint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type PageruleEditParamsTargetsConstraintOperator string

const (
	PageruleEditParamsTargetsConstraintOperatorMatches    PageruleEditParamsTargetsConstraintOperator = "matches"
	PageruleEditParamsTargetsConstraintOperatorContains   PageruleEditParamsTargetsConstraintOperator = "contains"
	PageruleEditParamsTargetsConstraintOperatorEquals     PageruleEditParamsTargetsConstraintOperator = "equals"
	PageruleEditParamsTargetsConstraintOperatorNotEqual   PageruleEditParamsTargetsConstraintOperator = "not_equal"
	PageruleEditParamsTargetsConstraintOperatorNotContain PageruleEditParamsTargetsConstraintOperator = "not_contain"
)

func (r PageruleEditParamsTargetsConstraintOperator) IsKnown() bool {
	switch r {
	case PageruleEditParamsTargetsConstraintOperatorMatches, PageruleEditParamsTargetsConstraintOperatorContains, PageruleEditParamsTargetsConstraintOperatorEquals, PageruleEditParamsTargetsConstraintOperatorNotEqual, PageruleEditParamsTargetsConstraintOperatorNotContain:
		return true
	}
	return false
}

// A target based on the URL of the request.
type PageruleEditParamsTargetsTarget string

const (
	PageruleEditParamsTargetsTargetURL PageruleEditParamsTargetsTarget = "url"
)

func (r PageruleEditParamsTargetsTarget) IsKnown() bool {
	switch r {
	case PageruleEditParamsTargetsTargetURL:
		return true
	}
	return false
}

type PageruleEditResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   PageruleEditResponse  `json:"result,required"`
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
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   PageruleGetResponse   `json:"result,required"`
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
