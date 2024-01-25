// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZonePageruleService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZonePageruleService] method
// instead.
type ZonePageruleService struct {
	Options  []option.RequestOption
	Settings *ZonePageruleSettingService
}

// NewZonePageruleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZonePageruleService(opts ...option.RequestOption) (r *ZonePageruleService) {
	r = &ZonePageruleService{}
	r.Options = opts
	r.Settings = NewZonePageruleSettingService(opts...)
	return
}

// Fetches the details of a Page Rule.
func (r *ZonePageruleService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *PageruleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/pagerules/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Replaces the configuration of an existing Page Rule. The configuration of the
// updated Page Rule will exactly match the data passed in the API request.
func (r *ZonePageruleService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZonePageruleUpdateParams, opts ...option.RequestOption) (res *PageruleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/pagerules/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes an existing Page Rule.
func (r *ZonePageruleService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZonePageruleDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/pagerules/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new Page Rule.
func (r *ZonePageruleService) PageRulesNewAPageRule(ctx context.Context, zoneIdentifier string, body ZonePagerulePageRulesNewAPageRuleParams, opts ...option.RequestOption) (res *PageruleResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/pagerules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Fetches Page Rules in a zone.
func (r *ZonePageruleService) PageRulesListPageRules(ctx context.Context, zoneIdentifier string, query ZonePagerulePageRulesListPageRulesParams, opts ...option.RequestOption) (res *ZonePagerulePageRulesListPageRulesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/pagerules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PageruleResponseSingle struct {
	Errors   []PageruleResponseSingleError   `json:"errors"`
	Messages []PageruleResponseSingleMessage `json:"messages"`
	Result   interface{}                     `json:"result"`
	// Whether the API call was successful
	Success PageruleResponseSingleSuccess `json:"success"`
	JSON    pageruleResponseSingleJSON    `json:"-"`
}

// pageruleResponseSingleJSON contains the JSON metadata for the struct
// [PageruleResponseSingle]
type pageruleResponseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageruleResponseSingleError struct {
	Code    int64                           `json:"code,required"`
	Message string                          `json:"message,required"`
	JSON    pageruleResponseSingleErrorJSON `json:"-"`
}

// pageruleResponseSingleErrorJSON contains the JSON metadata for the struct
// [PageruleResponseSingleError]
type pageruleResponseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageruleResponseSingleMessage struct {
	Code    int64                             `json:"code,required"`
	Message string                            `json:"message,required"`
	JSON    pageruleResponseSingleMessageJSON `json:"-"`
}

// pageruleResponseSingleMessageJSON contains the JSON metadata for the struct
// [PageruleResponseSingleMessage]
type pageruleResponseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PageruleResponseSingleSuccess bool

const (
	PageruleResponseSingleSuccessTrue PageruleResponseSingleSuccess = true
)

type ZonePageruleDeleteResponse struct {
	Errors   []ZonePageruleDeleteResponseError   `json:"errors"`
	Messages []ZonePageruleDeleteResponseMessage `json:"messages"`
	Result   ZonePageruleDeleteResponseResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success ZonePageruleDeleteResponseSuccess `json:"success"`
	JSON    zonePageruleDeleteResponseJSON    `json:"-"`
}

// zonePageruleDeleteResponseJSON contains the JSON metadata for the struct
// [ZonePageruleDeleteResponse]
type zonePageruleDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageruleDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageruleDeleteResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    zonePageruleDeleteResponseErrorJSON `json:"-"`
}

// zonePageruleDeleteResponseErrorJSON contains the JSON metadata for the struct
// [ZonePageruleDeleteResponseError]
type zonePageruleDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageruleDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageruleDeleteResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    zonePageruleDeleteResponseMessageJSON `json:"-"`
}

// zonePageruleDeleteResponseMessageJSON contains the JSON metadata for the struct
// [ZonePageruleDeleteResponseMessage]
type zonePageruleDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageruleDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePageruleDeleteResponseResult struct {
	// Identifier
	ID   string                               `json:"id,required"`
	JSON zonePageruleDeleteResponseResultJSON `json:"-"`
}

// zonePageruleDeleteResponseResultJSON contains the JSON metadata for the struct
// [ZonePageruleDeleteResponseResult]
type zonePageruleDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePageruleDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZonePageruleDeleteResponseSuccess bool

const (
	ZonePageruleDeleteResponseSuccessTrue ZonePageruleDeleteResponseSuccess = true
)

type ZonePagerulePageRulesListPageRulesResponse struct {
	Errors   []ZonePagerulePageRulesListPageRulesResponseError   `json:"errors"`
	Messages []ZonePagerulePageRulesListPageRulesResponseMessage `json:"messages"`
	Result   []ZonePagerulePageRulesListPageRulesResponseResult  `json:"result"`
	// Whether the API call was successful
	Success ZonePagerulePageRulesListPageRulesResponseSuccess `json:"success"`
	JSON    zonePagerulePageRulesListPageRulesResponseJSON    `json:"-"`
}

// zonePagerulePageRulesListPageRulesResponseJSON contains the JSON metadata for
// the struct [ZonePagerulePageRulesListPageRulesResponse]
type zonePagerulePageRulesListPageRulesResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePagerulePageRulesListPageRulesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePagerulePageRulesListPageRulesResponseError struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zonePagerulePageRulesListPageRulesResponseErrorJSON `json:"-"`
}

// zonePagerulePageRulesListPageRulesResponseErrorJSON contains the JSON metadata
// for the struct [ZonePagerulePageRulesListPageRulesResponseError]
type zonePagerulePageRulesListPageRulesResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePagerulePageRulesListPageRulesResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePagerulePageRulesListPageRulesResponseMessage struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    zonePagerulePageRulesListPageRulesResponseMessageJSON `json:"-"`
}

// zonePagerulePageRulesListPageRulesResponseMessageJSON contains the JSON metadata
// for the struct [ZonePagerulePageRulesListPageRulesResponseMessage]
type zonePagerulePageRulesListPageRulesResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePagerulePageRulesListPageRulesResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePagerulePageRulesListPageRulesResponseResult struct {
	// Identifier
	ID string `json:"id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions []ZonePagerulePageRulesListPageRulesResponseResultAction `json:"actions,required"`
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
	Status ZonePagerulePageRulesListPageRulesResponseResultStatus `json:"status,required"`
	// The rule targets to evaluate on each request.
	Targets []ZonePagerulePageRulesListPageRulesResponseResultTarget `json:"targets,required"`
	JSON    zonePagerulePageRulesListPageRulesResponseResultJSON     `json:"-"`
}

// zonePagerulePageRulesListPageRulesResponseResultJSON contains the JSON metadata
// for the struct [ZonePagerulePageRulesListPageRulesResponseResult]
type zonePagerulePageRulesListPageRulesResponseResultJSON struct {
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

func (r *ZonePagerulePageRulesListPageRulesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZonePagerulePageRulesListPageRulesResponseResultAction struct {
	// The timestamp of when the override was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The type of route.
	Name  ZonePagerulePageRulesListPageRulesResponseResultActionsName  `json:"name"`
	Value ZonePagerulePageRulesListPageRulesResponseResultActionsValue `json:"value"`
	JSON  zonePagerulePageRulesListPageRulesResponseResultActionJSON   `json:"-"`
}

// zonePagerulePageRulesListPageRulesResponseResultActionJSON contains the JSON
// metadata for the struct [ZonePagerulePageRulesListPageRulesResponseResultAction]
type zonePagerulePageRulesListPageRulesResponseResultActionJSON struct {
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePagerulePageRulesListPageRulesResponseResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of route.
type ZonePagerulePageRulesListPageRulesResponseResultActionsName string

const (
	ZonePagerulePageRulesListPageRulesResponseResultActionsNameForwardURL ZonePagerulePageRulesListPageRulesResponseResultActionsName = "forward_url"
)

type ZonePagerulePageRulesListPageRulesResponseResultActionsValue struct {
	// The response type for the URL redirect.
	Type ZonePagerulePageRulesListPageRulesResponseResultActionsValueType `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string                                                           `json:"url"`
	JSON zonePagerulePageRulesListPageRulesResponseResultActionsValueJSON `json:"-"`
}

// zonePagerulePageRulesListPageRulesResponseResultActionsValueJSON contains the
// JSON metadata for the struct
// [ZonePagerulePageRulesListPageRulesResponseResultActionsValue]
type zonePagerulePageRulesListPageRulesResponseResultActionsValueJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePagerulePageRulesListPageRulesResponseResultActionsValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response type for the URL redirect.
type ZonePagerulePageRulesListPageRulesResponseResultActionsValueType string

const (
	ZonePagerulePageRulesListPageRulesResponseResultActionsValueTypeTemporary ZonePagerulePageRulesListPageRulesResponseResultActionsValueType = "temporary"
	ZonePagerulePageRulesListPageRulesResponseResultActionsValueTypePermanent ZonePagerulePageRulesListPageRulesResponseResultActionsValueType = "permanent"
)

// The status of the Page Rule.
type ZonePagerulePageRulesListPageRulesResponseResultStatus string

const (
	ZonePagerulePageRulesListPageRulesResponseResultStatusActive   ZonePagerulePageRulesListPageRulesResponseResultStatus = "active"
	ZonePagerulePageRulesListPageRulesResponseResultStatusDisabled ZonePagerulePageRulesListPageRulesResponseResultStatus = "disabled"
)

// A request condition target.
type ZonePagerulePageRulesListPageRulesResponseResultTarget struct {
	// The constraint of a target.
	Constraint ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraint `json:"constraint,required"`
	// A target based on the URL of the request.
	Target ZonePagerulePageRulesListPageRulesResponseResultTargetsTarget `json:"target,required"`
	JSON   zonePagerulePageRulesListPageRulesResponseResultTargetJSON    `json:"-"`
}

// zonePagerulePageRulesListPageRulesResponseResultTargetJSON contains the JSON
// metadata for the struct [ZonePagerulePageRulesListPageRulesResponseResultTarget]
type zonePagerulePageRulesListPageRulesResponseResultTargetJSON struct {
	Constraint  apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePagerulePageRulesListPageRulesResponseResultTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The constraint of a target.
type ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraintOperator `json:"operator"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value string                                                                `json:"value"`
	JSON  zonePagerulePageRulesListPageRulesResponseResultTargetsConstraintJSON `json:"-"`
}

// zonePagerulePageRulesListPageRulesResponseResultTargetsConstraintJSON contains
// the JSON metadata for the struct
// [ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraint]
type zonePagerulePageRulesListPageRulesResponseResultTargetsConstraintJSON struct {
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraintOperator string

const (
	ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraintOperatorMatches    ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraintOperator = "matches"
	ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraintOperatorContains   ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraintOperator = "contains"
	ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraintOperatorEquals     ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraintOperator = "equals"
	ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraintOperatorNotEqual   ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraintOperator = "not_equal"
	ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraintOperatorNotContain ZonePagerulePageRulesListPageRulesResponseResultTargetsConstraintOperator = "not_contain"
)

// A target based on the URL of the request.
type ZonePagerulePageRulesListPageRulesResponseResultTargetsTarget string

const (
	ZonePagerulePageRulesListPageRulesResponseResultTargetsTargetURL ZonePagerulePageRulesListPageRulesResponseResultTargetsTarget = "url"
)

// Whether the API call was successful
type ZonePagerulePageRulesListPageRulesResponseSuccess bool

const (
	ZonePagerulePageRulesListPageRulesResponseSuccessTrue ZonePagerulePageRulesListPageRulesResponseSuccess = true
)

type ZonePageruleUpdateParams struct {
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]ZonePageruleUpdateParamsAction] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]ZonePageruleUpdateParamsTarget] `json:"targets,required"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[ZonePageruleUpdateParamsStatus] `json:"status"`
}

func (r ZonePageruleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZonePageruleUpdateParamsAction struct {
	// The type of route.
	Name  param.Field[ZonePageruleUpdateParamsActionsName]  `json:"name"`
	Value param.Field[ZonePageruleUpdateParamsActionsValue] `json:"value"`
}

func (r ZonePageruleUpdateParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of route.
type ZonePageruleUpdateParamsActionsName string

const (
	ZonePageruleUpdateParamsActionsNameForwardURL ZonePageruleUpdateParamsActionsName = "forward_url"
)

type ZonePageruleUpdateParamsActionsValue struct {
	// The response type for the URL redirect.
	Type param.Field[ZonePageruleUpdateParamsActionsValueType] `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r ZonePageruleUpdateParamsActionsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response type for the URL redirect.
type ZonePageruleUpdateParamsActionsValueType string

const (
	ZonePageruleUpdateParamsActionsValueTypeTemporary ZonePageruleUpdateParamsActionsValueType = "temporary"
	ZonePageruleUpdateParamsActionsValueTypePermanent ZonePageruleUpdateParamsActionsValueType = "permanent"
)

// A request condition target.
type ZonePageruleUpdateParamsTarget struct {
	// The constraint of a target.
	Constraint param.Field[ZonePageruleUpdateParamsTargetsConstraint] `json:"constraint,required"`
	// A target based on the URL of the request.
	Target param.Field[ZonePageruleUpdateParamsTargetsTarget] `json:"target,required"`
}

func (r ZonePageruleUpdateParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The constraint of a target.
type ZonePageruleUpdateParamsTargetsConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator param.Field[ZonePageruleUpdateParamsTargetsConstraintOperator] `json:"operator"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value param.Field[string] `json:"value"`
}

func (r ZonePageruleUpdateParamsTargetsConstraint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type ZonePageruleUpdateParamsTargetsConstraintOperator string

const (
	ZonePageruleUpdateParamsTargetsConstraintOperatorMatches    ZonePageruleUpdateParamsTargetsConstraintOperator = "matches"
	ZonePageruleUpdateParamsTargetsConstraintOperatorContains   ZonePageruleUpdateParamsTargetsConstraintOperator = "contains"
	ZonePageruleUpdateParamsTargetsConstraintOperatorEquals     ZonePageruleUpdateParamsTargetsConstraintOperator = "equals"
	ZonePageruleUpdateParamsTargetsConstraintOperatorNotEqual   ZonePageruleUpdateParamsTargetsConstraintOperator = "not_equal"
	ZonePageruleUpdateParamsTargetsConstraintOperatorNotContain ZonePageruleUpdateParamsTargetsConstraintOperator = "not_contain"
)

// A target based on the URL of the request.
type ZonePageruleUpdateParamsTargetsTarget string

const (
	ZonePageruleUpdateParamsTargetsTargetURL ZonePageruleUpdateParamsTargetsTarget = "url"
)

// The status of the Page Rule.
type ZonePageruleUpdateParamsStatus string

const (
	ZonePageruleUpdateParamsStatusActive   ZonePageruleUpdateParamsStatus = "active"
	ZonePageruleUpdateParamsStatusDisabled ZonePageruleUpdateParamsStatus = "disabled"
)

type ZonePagerulePageRulesNewAPageRuleParams struct {
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions param.Field[[]ZonePagerulePageRulesNewAPageRuleParamsAction] `json:"actions,required"`
	// The rule targets to evaluate on each request.
	Targets param.Field[[]ZonePagerulePageRulesNewAPageRuleParamsTarget] `json:"targets,required"`
	// The priority of the rule, used to define which Page Rule is processed over
	// another. A higher number indicates a higher priority. For example, if you have a
	// catch-all Page Rule (rule A: `/images/*`) but want a more specific Page Rule to
	// take precedence (rule B: `/images/special/*`), specify a higher priority for
	// rule B so it overrides rule A.
	Priority param.Field[int64] `json:"priority"`
	// The status of the Page Rule.
	Status param.Field[ZonePagerulePageRulesNewAPageRuleParamsStatus] `json:"status"`
}

func (r ZonePagerulePageRulesNewAPageRuleParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZonePagerulePageRulesNewAPageRuleParamsAction struct {
	// The type of route.
	Name  param.Field[ZonePagerulePageRulesNewAPageRuleParamsActionsName]  `json:"name"`
	Value param.Field[ZonePagerulePageRulesNewAPageRuleParamsActionsValue] `json:"value"`
}

func (r ZonePagerulePageRulesNewAPageRuleParamsAction) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The type of route.
type ZonePagerulePageRulesNewAPageRuleParamsActionsName string

const (
	ZonePagerulePageRulesNewAPageRuleParamsActionsNameForwardURL ZonePagerulePageRulesNewAPageRuleParamsActionsName = "forward_url"
)

type ZonePagerulePageRulesNewAPageRuleParamsActionsValue struct {
	// The response type for the URL redirect.
	Type param.Field[ZonePagerulePageRulesNewAPageRuleParamsActionsValueType] `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL param.Field[string] `json:"url"`
}

func (r ZonePagerulePageRulesNewAPageRuleParamsActionsValue) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The response type for the URL redirect.
type ZonePagerulePageRulesNewAPageRuleParamsActionsValueType string

const (
	ZonePagerulePageRulesNewAPageRuleParamsActionsValueTypeTemporary ZonePagerulePageRulesNewAPageRuleParamsActionsValueType = "temporary"
	ZonePagerulePageRulesNewAPageRuleParamsActionsValueTypePermanent ZonePagerulePageRulesNewAPageRuleParamsActionsValueType = "permanent"
)

// A request condition target.
type ZonePagerulePageRulesNewAPageRuleParamsTarget struct {
	// The constraint of a target.
	Constraint param.Field[ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraint] `json:"constraint,required"`
	// A target based on the URL of the request.
	Target param.Field[ZonePagerulePageRulesNewAPageRuleParamsTargetsTarget] `json:"target,required"`
}

func (r ZonePagerulePageRulesNewAPageRuleParamsTarget) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The constraint of a target.
type ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator param.Field[ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperator] `json:"operator"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value param.Field[string] `json:"value"`
}

func (r ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperator string

const (
	ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperatorMatches    ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperator = "matches"
	ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperatorContains   ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperator = "contains"
	ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperatorEquals     ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperator = "equals"
	ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperatorNotEqual   ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperator = "not_equal"
	ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperatorNotContain ZonePagerulePageRulesNewAPageRuleParamsTargetsConstraintOperator = "not_contain"
)

// A target based on the URL of the request.
type ZonePagerulePageRulesNewAPageRuleParamsTargetsTarget string

const (
	ZonePagerulePageRulesNewAPageRuleParamsTargetsTargetURL ZonePagerulePageRulesNewAPageRuleParamsTargetsTarget = "url"
)

// The status of the Page Rule.
type ZonePagerulePageRulesNewAPageRuleParamsStatus string

const (
	ZonePagerulePageRulesNewAPageRuleParamsStatusActive   ZonePagerulePageRulesNewAPageRuleParamsStatus = "active"
	ZonePagerulePageRulesNewAPageRuleParamsStatusDisabled ZonePagerulePageRulesNewAPageRuleParamsStatus = "disabled"
)

type ZonePagerulePageRulesListPageRulesParams struct {
	// The direction used to sort returned Page Rules.
	Direction param.Field[ZonePagerulePageRulesListPageRulesParamsDirection] `query:"direction"`
	// When set to `all`, all the search requirements must match. When set to `any`,
	// only one of the search requirements has to match.
	Match param.Field[ZonePagerulePageRulesListPageRulesParamsMatch] `query:"match"`
	// The field used to sort returned Page Rules.
	Order param.Field[ZonePagerulePageRulesListPageRulesParamsOrder] `query:"order"`
	// The status of the Page Rule.
	Status param.Field[ZonePagerulePageRulesListPageRulesParamsStatus] `query:"status"`
}

// URLQuery serializes [ZonePagerulePageRulesListPageRulesParams]'s query
// parameters as `url.Values`.
func (r ZonePagerulePageRulesListPageRulesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The direction used to sort returned Page Rules.
type ZonePagerulePageRulesListPageRulesParamsDirection string

const (
	ZonePagerulePageRulesListPageRulesParamsDirectionAsc  ZonePagerulePageRulesListPageRulesParamsDirection = "asc"
	ZonePagerulePageRulesListPageRulesParamsDirectionDesc ZonePagerulePageRulesListPageRulesParamsDirection = "desc"
)

// When set to `all`, all the search requirements must match. When set to `any`,
// only one of the search requirements has to match.
type ZonePagerulePageRulesListPageRulesParamsMatch string

const (
	ZonePagerulePageRulesListPageRulesParamsMatchAny ZonePagerulePageRulesListPageRulesParamsMatch = "any"
	ZonePagerulePageRulesListPageRulesParamsMatchAll ZonePagerulePageRulesListPageRulesParamsMatch = "all"
)

// The field used to sort returned Page Rules.
type ZonePagerulePageRulesListPageRulesParamsOrder string

const (
	ZonePagerulePageRulesListPageRulesParamsOrderStatus   ZonePagerulePageRulesListPageRulesParamsOrder = "status"
	ZonePagerulePageRulesListPageRulesParamsOrderPriority ZonePagerulePageRulesListPageRulesParamsOrder = "priority"
)

// The status of the Page Rule.
type ZonePagerulePageRulesListPageRulesParamsStatus string

const (
	ZonePagerulePageRulesListPageRulesParamsStatusActive   ZonePagerulePageRulesListPageRulesParamsStatus = "active"
	ZonePagerulePageRulesListPageRulesParamsStatusDisabled ZonePagerulePageRulesListPageRulesParamsStatus = "disabled"
)
