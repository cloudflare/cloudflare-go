// File generated from our OpenAPI spec by Stainless.

package pagerules

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
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
func (r *PageruleService) List(ctx context.Context, params PageruleListParams, opts ...option.RequestOption) (res *[]PageruleListResponse, err error) {
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
func (r *PageruleService) Delete(ctx context.Context, pageruleID string, body PageruleDeleteParams, opts ...option.RequestOption) (res *PageruleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env PageruleDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/pagerules/%s", body.ZoneID, pageruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type PageruleListResponse struct {
	// Identifier
	ID string `json:"id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions []PageruleListResponseAction `json:"actions,required"`
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
	Status PageruleListResponseStatus `json:"status,required"`
	// The rule targets to evaluate on each request.
	Targets []PageruleListResponseTarget `json:"targets,required"`
	JSON    pageruleListResponseJSON     `json:"-"`
}

// pageruleListResponseJSON contains the JSON metadata for the struct
// [PageruleListResponse]
type pageruleListResponseJSON struct {
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

func (r *PageruleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseJSON) RawJSON() string {
	return r.raw
}

type PageruleListResponseAction struct {
	// The timestamp of when the override was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The type of route.
	Name  PageruleListResponseActionsName  `json:"name"`
	Value PageruleListResponseActionsValue `json:"value"`
	JSON  pageruleListResponseActionJSON   `json:"-"`
}

// pageruleListResponseActionJSON contains the JSON metadata for the struct
// [PageruleListResponseAction]
type pageruleListResponseActionJSON struct {
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleListResponseAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseActionJSON) RawJSON() string {
	return r.raw
}

// The type of route.
type PageruleListResponseActionsName string

const (
	PageruleListResponseActionsNameForwardURL PageruleListResponseActionsName = "forward_url"
)

type PageruleListResponseActionsValue struct {
	// The response type for the URL redirect.
	Type PageruleListResponseActionsValueType `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string                               `json:"url"`
	JSON pageruleListResponseActionsValueJSON `json:"-"`
}

// pageruleListResponseActionsValueJSON contains the JSON metadata for the struct
// [PageruleListResponseActionsValue]
type pageruleListResponseActionsValueJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleListResponseActionsValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseActionsValueJSON) RawJSON() string {
	return r.raw
}

// The response type for the URL redirect.
type PageruleListResponseActionsValueType string

const (
	PageruleListResponseActionsValueTypeTemporary PageruleListResponseActionsValueType = "temporary"
	PageruleListResponseActionsValueTypePermanent PageruleListResponseActionsValueType = "permanent"
)

// The status of the Page Rule.
type PageruleListResponseStatus string

const (
	PageruleListResponseStatusActive   PageruleListResponseStatus = "active"
	PageruleListResponseStatusDisabled PageruleListResponseStatus = "disabled"
)

// A request condition target.
type PageruleListResponseTarget struct {
	// String constraint.
	Constraint PageruleListResponseTargetsConstraint `json:"constraint,required"`
	// A target based on the URL of the request.
	Target PageruleListResponseTargetsTarget `json:"target,required"`
	JSON   pageruleListResponseTargetJSON    `json:"-"`
}

// pageruleListResponseTargetJSON contains the JSON metadata for the struct
// [PageruleListResponseTarget]
type pageruleListResponseTargetJSON struct {
	Constraint  apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleListResponseTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseTargetJSON) RawJSON() string {
	return r.raw
}

// String constraint.
type PageruleListResponseTargetsConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator PageruleListResponseTargetsConstraintOperator `json:"operator,required"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value string                                    `json:"value,required"`
	JSON  pageruleListResponseTargetsConstraintJSON `json:"-"`
}

// pageruleListResponseTargetsConstraintJSON contains the JSON metadata for the
// struct [PageruleListResponseTargetsConstraint]
type pageruleListResponseTargetsConstraintJSON struct {
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleListResponseTargetsConstraint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseTargetsConstraintJSON) RawJSON() string {
	return r.raw
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type PageruleListResponseTargetsConstraintOperator string

const (
	PageruleListResponseTargetsConstraintOperatorMatches    PageruleListResponseTargetsConstraintOperator = "matches"
	PageruleListResponseTargetsConstraintOperatorContains   PageruleListResponseTargetsConstraintOperator = "contains"
	PageruleListResponseTargetsConstraintOperatorEquals     PageruleListResponseTargetsConstraintOperator = "equals"
	PageruleListResponseTargetsConstraintOperatorNotEqual   PageruleListResponseTargetsConstraintOperator = "not_equal"
	PageruleListResponseTargetsConstraintOperatorNotContain PageruleListResponseTargetsConstraintOperator = "not_contain"
)

// A target based on the URL of the request.
type PageruleListResponseTargetsTarget string

const (
	PageruleListResponseTargetsTargetURL PageruleListResponseTargetsTarget = "url"
)

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

// A target based on the URL of the request.
type PageruleNewParamsTargetsTarget string

const (
	PageruleNewParamsTargetsTargetURL PageruleNewParamsTargetsTarget = "url"
)

// The status of the Page Rule.
type PageruleNewParamsStatus string

const (
	PageruleNewParamsStatusActive   PageruleNewParamsStatus = "active"
	PageruleNewParamsStatusDisabled PageruleNewParamsStatus = "disabled"
)

type PageruleNewResponseEnvelope struct {
	Errors   []PageruleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageruleNewResponseEnvelopeMessages `json:"messages,required"`
	Result   PageruleNewResponse                   `json:"result,required"`
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

type PageruleNewResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    pageruleNewResponseEnvelopeErrorsJSON `json:"-"`
}

// pageruleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PageruleNewResponseEnvelopeErrors]
type pageruleNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageruleNewResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    pageruleNewResponseEnvelopeMessagesJSON `json:"-"`
}

// pageruleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageruleNewResponseEnvelopeMessages]
type pageruleNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleNewResponseEnvelopeSuccess bool

const (
	PageruleNewResponseEnvelopeSuccessTrue PageruleNewResponseEnvelopeSuccess = true
)

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

// A target based on the URL of the request.
type PageruleUpdateParamsTargetsTarget string

const (
	PageruleUpdateParamsTargetsTargetURL PageruleUpdateParamsTargetsTarget = "url"
)

// The status of the Page Rule.
type PageruleUpdateParamsStatus string

const (
	PageruleUpdateParamsStatusActive   PageruleUpdateParamsStatus = "active"
	PageruleUpdateParamsStatusDisabled PageruleUpdateParamsStatus = "disabled"
)

type PageruleUpdateResponseEnvelope struct {
	Errors   []PageruleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageruleUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   PageruleUpdateResponse                   `json:"result,required"`
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

type PageruleUpdateResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    pageruleUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// pageruleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PageruleUpdateResponseEnvelopeErrors]
type pageruleUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageruleUpdateResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    pageruleUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// pageruleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageruleUpdateResponseEnvelopeMessages]
type pageruleUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleUpdateResponseEnvelopeSuccess bool

const (
	PageruleUpdateResponseEnvelopeSuccessTrue PageruleUpdateResponseEnvelopeSuccess = true
)

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
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned Page Rules.
type PageruleListParamsDirection string

const (
	PageruleListParamsDirectionAsc  PageruleListParamsDirection = "asc"
	PageruleListParamsDirectionDesc PageruleListParamsDirection = "desc"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type PageruleListParamsMatch string

const (
	PageruleListParamsMatchAny PageruleListParamsMatch = "any"
	PageruleListParamsMatchAll PageruleListParamsMatch = "all"
)

// The field used to sort returned Page Rules.
type PageruleListParamsOrder string

const (
	PageruleListParamsOrderStatus   PageruleListParamsOrder = "status"
	PageruleListParamsOrderPriority PageruleListParamsOrder = "priority"
)

// The status of the Page Rule.
type PageruleListParamsStatus string

const (
	PageruleListParamsStatusActive   PageruleListParamsStatus = "active"
	PageruleListParamsStatusDisabled PageruleListParamsStatus = "disabled"
)

type PageruleListResponseEnvelope struct {
	Errors   []PageruleListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageruleListResponseEnvelopeMessages `json:"messages,required"`
	Result   []PageruleListResponse                 `json:"result,required"`
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

type PageruleListResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    pageruleListResponseEnvelopeErrorsJSON `json:"-"`
}

// pageruleListResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PageruleListResponseEnvelopeErrors]
type pageruleListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageruleListResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    pageruleListResponseEnvelopeMessagesJSON `json:"-"`
}

// pageruleListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageruleListResponseEnvelopeMessages]
type pageruleListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleListResponseEnvelopeSuccess bool

const (
	PageruleListResponseEnvelopeSuccessTrue PageruleListResponseEnvelopeSuccess = true
)

type PageruleDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageruleDeleteResponseEnvelope struct {
	Errors   []PageruleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageruleDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   PageruleDeleteResponse                   `json:"result,required,nullable"`
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

type PageruleDeleteResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    pageruleDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// pageruleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [PageruleDeleteResponseEnvelopeErrors]
type pageruleDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageruleDeleteResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    pageruleDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// pageruleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageruleDeleteResponseEnvelopeMessages]
type pageruleDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleDeleteResponseEnvelopeSuccess bool

const (
	PageruleDeleteResponseEnvelopeSuccessTrue PageruleDeleteResponseEnvelopeSuccess = true
)

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

// The status of the Page Rule.
type PageruleEditParamsStatus string

const (
	PageruleEditParamsStatusActive   PageruleEditParamsStatus = "active"
	PageruleEditParamsStatusDisabled PageruleEditParamsStatus = "disabled"
)

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

// A target based on the URL of the request.
type PageruleEditParamsTargetsTarget string

const (
	PageruleEditParamsTargetsTargetURL PageruleEditParamsTargetsTarget = "url"
)

type PageruleEditResponseEnvelope struct {
	Errors   []PageruleEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageruleEditResponseEnvelopeMessages `json:"messages,required"`
	Result   PageruleEditResponse                   `json:"result,required"`
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

type PageruleEditResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    pageruleEditResponseEnvelopeErrorsJSON `json:"-"`
}

// pageruleEditResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PageruleEditResponseEnvelopeErrors]
type pageruleEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageruleEditResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    pageruleEditResponseEnvelopeMessagesJSON `json:"-"`
}

// pageruleEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageruleEditResponseEnvelopeMessages]
type pageruleEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleEditResponseEnvelopeSuccess bool

const (
	PageruleEditResponseEnvelopeSuccessTrue PageruleEditResponseEnvelopeSuccess = true
)

type PageruleGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type PageruleGetResponseEnvelope struct {
	Errors   []PageruleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []PageruleGetResponseEnvelopeMessages `json:"messages,required"`
	Result   PageruleGetResponse                   `json:"result,required"`
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

type PageruleGetResponseEnvelopeErrors struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    pageruleGetResponseEnvelopeErrorsJSON `json:"-"`
}

// pageruleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [PageruleGetResponseEnvelopeErrors]
type pageruleGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type PageruleGetResponseEnvelopeMessages struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    pageruleGetResponseEnvelopeMessagesJSON `json:"-"`
}

// pageruleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [PageruleGetResponseEnvelopeMessages]
type pageruleGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pageruleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PageruleGetResponseEnvelopeSuccess bool

const (
	PageruleGetResponseEnvelopeSuccessTrue PageruleGetResponseEnvelopeSuccess = true
)
