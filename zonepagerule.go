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
func (r *ZonePageruleService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *SchemasAPIResponseSingleID, err error) {
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
func (r *ZonePageruleService) PageRulesListPageRules(ctx context.Context, zoneIdentifier string, query ZonePagerulePageRulesListPageRulesParams, opts ...option.RequestOption) (res *PageruleResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/pagerules", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PageruleResponseCollection struct {
	Errors   []PageruleResponseCollectionError   `json:"errors"`
	Messages []PageruleResponseCollectionMessage `json:"messages"`
	Result   []PageruleResponseCollectionResult  `json:"result"`
	// Whether the API call was successful
	Success PageruleResponseCollectionSuccess `json:"success"`
	JSON    pageruleResponseCollectionJSON    `json:"-"`
}

// pageruleResponseCollectionJSON contains the JSON metadata for the struct
// [PageruleResponseCollection]
type pageruleResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageruleResponseCollectionError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    pageruleResponseCollectionErrorJSON `json:"-"`
}

// pageruleResponseCollectionErrorJSON contains the JSON metadata for the struct
// [PageruleResponseCollectionError]
type pageruleResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageruleResponseCollectionMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    pageruleResponseCollectionMessageJSON `json:"-"`
}

// pageruleResponseCollectionMessageJSON contains the JSON metadata for the struct
// [PageruleResponseCollectionMessage]
type pageruleResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageruleResponseCollectionResult struct {
	// Identifier
	ID string `json:"id,required"`
	// The set of actions to perform if the targets of this rule match the request.
	// Actions can redirect to another URL or override settings, but not both.
	Actions []PageruleResponseCollectionResultAction `json:"actions,required"`
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
	Status PageruleResponseCollectionResultStatus `json:"status,required"`
	// The rule targets to evaluate on each request.
	Targets []PageruleResponseCollectionResultTarget `json:"targets,required"`
	JSON    pageruleResponseCollectionResultJSON     `json:"-"`
}

// pageruleResponseCollectionResultJSON contains the JSON metadata for the struct
// [PageruleResponseCollectionResult]
type pageruleResponseCollectionResultJSON struct {
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

func (r *PageruleResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PageruleResponseCollectionResultAction struct {
	// The timestamp of when the override was last modified.
	ModifiedOn time.Time `json:"modified_on" format:"date-time"`
	// The type of route.
	Name  PageruleResponseCollectionResultActionsName  `json:"name"`
	Value PageruleResponseCollectionResultActionsValue `json:"value"`
	JSON  pageruleResponseCollectionResultActionJSON   `json:"-"`
}

// pageruleResponseCollectionResultActionJSON contains the JSON metadata for the
// struct [PageruleResponseCollectionResultAction]
type pageruleResponseCollectionResultActionJSON struct {
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleResponseCollectionResultAction) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of route.
type PageruleResponseCollectionResultActionsName string

const (
	PageruleResponseCollectionResultActionsNameForwardURL PageruleResponseCollectionResultActionsName = "forward_url"
)

type PageruleResponseCollectionResultActionsValue struct {
	// The response type for the URL redirect.
	Type PageruleResponseCollectionResultActionsValueType `json:"type"`
	// The URL to redirect the request to. Notes: ${num} refers to the position of '\*'
	// in the constraint value.
	URL  string                                           `json:"url"`
	JSON pageruleResponseCollectionResultActionsValueJSON `json:"-"`
}

// pageruleResponseCollectionResultActionsValueJSON contains the JSON metadata for
// the struct [PageruleResponseCollectionResultActionsValue]
type pageruleResponseCollectionResultActionsValueJSON struct {
	Type        apijson.Field
	URL         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleResponseCollectionResultActionsValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The response type for the URL redirect.
type PageruleResponseCollectionResultActionsValueType string

const (
	PageruleResponseCollectionResultActionsValueTypeTemporary PageruleResponseCollectionResultActionsValueType = "temporary"
	PageruleResponseCollectionResultActionsValueTypePermanent PageruleResponseCollectionResultActionsValueType = "permanent"
)

// The status of the Page Rule.
type PageruleResponseCollectionResultStatus string

const (
	PageruleResponseCollectionResultStatusActive   PageruleResponseCollectionResultStatus = "active"
	PageruleResponseCollectionResultStatusDisabled PageruleResponseCollectionResultStatus = "disabled"
)

// A request condition target.
type PageruleResponseCollectionResultTarget struct {
	// The constraint of a target.
	Constraint PageruleResponseCollectionResultTargetsConstraint `json:"constraint,required"`
	// A target based on the URL of the request.
	Target PageruleResponseCollectionResultTargetsTarget `json:"target,required"`
	JSON   pageruleResponseCollectionResultTargetJSON    `json:"-"`
}

// pageruleResponseCollectionResultTargetJSON contains the JSON metadata for the
// struct [PageruleResponseCollectionResultTarget]
type pageruleResponseCollectionResultTargetJSON struct {
	Constraint  apijson.Field
	Target      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleResponseCollectionResultTarget) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The constraint of a target.
type PageruleResponseCollectionResultTargetsConstraint struct {
	// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
	Operator PageruleResponseCollectionResultTargetsConstraintOperator `json:"operator"`
	// The URL pattern to match against the current request. The pattern may contain up
	// to four asterisks ('\*') as placeholders.
	Value string                                                `json:"value"`
	JSON  pageruleResponseCollectionResultTargetsConstraintJSON `json:"-"`
}

// pageruleResponseCollectionResultTargetsConstraintJSON contains the JSON metadata
// for the struct [PageruleResponseCollectionResultTargetsConstraint]
type pageruleResponseCollectionResultTargetsConstraintJSON struct {
	Operator    apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PageruleResponseCollectionResultTargetsConstraint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The matches operator can use asterisks and pipes as wildcard and 'or' operators.
type PageruleResponseCollectionResultTargetsConstraintOperator string

const (
	PageruleResponseCollectionResultTargetsConstraintOperatorMatches    PageruleResponseCollectionResultTargetsConstraintOperator = "matches"
	PageruleResponseCollectionResultTargetsConstraintOperatorContains   PageruleResponseCollectionResultTargetsConstraintOperator = "contains"
	PageruleResponseCollectionResultTargetsConstraintOperatorEquals     PageruleResponseCollectionResultTargetsConstraintOperator = "equals"
	PageruleResponseCollectionResultTargetsConstraintOperatorNotEqual   PageruleResponseCollectionResultTargetsConstraintOperator = "not_equal"
	PageruleResponseCollectionResultTargetsConstraintOperatorNotContain PageruleResponseCollectionResultTargetsConstraintOperator = "not_contain"
)

// A target based on the URL of the request.
type PageruleResponseCollectionResultTargetsTarget string

const (
	PageruleResponseCollectionResultTargetsTargetURL PageruleResponseCollectionResultTargetsTarget = "url"
)

// Whether the API call was successful
type PageruleResponseCollectionSuccess bool

const (
	PageruleResponseCollectionSuccessTrue PageruleResponseCollectionSuccess = true
)

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

type SchemasAPIResponseSingleID struct {
	Errors   []SchemasAPIResponseSingleIDError   `json:"errors"`
	Messages []SchemasAPIResponseSingleIDMessage `json:"messages"`
	Result   SchemasAPIResponseSingleIDResult    `json:"result,nullable"`
	// Whether the API call was successful
	Success SchemasAPIResponseSingleIDSuccess `json:"success"`
	JSON    schemasAPIResponseSingleIDJSON    `json:"-"`
}

// schemasAPIResponseSingleIDJSON contains the JSON metadata for the struct
// [SchemasAPIResponseSingleID]
type schemasAPIResponseSingleIDJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasAPIResponseSingleID) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasAPIResponseSingleIDError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    schemasAPIResponseSingleIDErrorJSON `json:"-"`
}

// schemasAPIResponseSingleIDErrorJSON contains the JSON metadata for the struct
// [SchemasAPIResponseSingleIDError]
type schemasAPIResponseSingleIDErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasAPIResponseSingleIDError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasAPIResponseSingleIDMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    schemasAPIResponseSingleIDMessageJSON `json:"-"`
}

// schemasAPIResponseSingleIDMessageJSON contains the JSON metadata for the struct
// [SchemasAPIResponseSingleIDMessage]
type schemasAPIResponseSingleIDMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasAPIResponseSingleIDMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SchemasAPIResponseSingleIDResult struct {
	// Identifier
	ID   string                               `json:"id,required"`
	JSON schemasAPIResponseSingleIDResultJSON `json:"-"`
}

// schemasAPIResponseSingleIDResultJSON contains the JSON metadata for the struct
// [SchemasAPIResponseSingleIDResult]
type schemasAPIResponseSingleIDResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemasAPIResponseSingleIDResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SchemasAPIResponseSingleIDSuccess bool

const (
	SchemasAPIResponseSingleIDSuccessTrue SchemasAPIResponseSingleIDSuccess = true
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
