// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// ObservabilityQueryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityQueryService] method instead.
type ObservabilityQueryService struct {
	Options []option.RequestOption
}

// NewObservabilityQueryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObservabilityQueryService(opts ...option.RequestOption) (r *ObservabilityQueryService) {
	r = &ObservabilityQueryService{}
	r.Options = opts
	return
}

// Persist query for later use.
func (r *ObservabilityQueryService) New(ctx context.Context, params ObservabilityQueryNewParams, opts ...option.RequestOption) (res *ObservabilityQueryNewResponse, err error) {
	var env ObservabilityQueryNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/queries", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List saved queries.
func (r *ObservabilityQueryService) List(ctx context.Context, params ObservabilityQueryListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ObservabilityQueryListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/queries", params.AccountID)
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

// List saved queries.
func (r *ObservabilityQueryService) ListAutoPaging(ctx context.Context, params ObservabilityQueryListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ObservabilityQueryListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

type ObservabilityQueryNewResponse struct {
	ID string `json:"id" api:"required"`
	// If the query wasn't explcitly saved
	Adhoc       bool   `json:"adhoc" api:"required"`
	Created     string `json:"created" api:"required"`
	CreatedBy   string `json:"createdBy" api:"required"`
	Description string `json:"description" api:"required,nullable"`
	// Query name
	Name       string                                  `json:"name" api:"required"`
	Parameters ObservabilityQueryNewResponseParameters `json:"parameters" api:"required"`
	Updated    string                                  `json:"updated" api:"required"`
	UpdatedBy  string                                  `json:"updatedBy" api:"required"`
	JSON       observabilityQueryNewResponseJSON       `json:"-"`
}

// observabilityQueryNewResponseJSON contains the JSON metadata for the struct
// [ObservabilityQueryNewResponse]
type observabilityQueryNewResponseJSON struct {
	ID          apijson.Field
	Adhoc       apijson.Field
	Created     apijson.Field
	CreatedBy   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Parameters  apijson.Field
	Updated     apijson.Field
	UpdatedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryNewResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryNewResponseParameters struct {
	// Create Calculations to compute as part of the query.
	Calculations []ObservabilityQueryNewResponseParametersCalculation `json:"calculations"`
	// Set the Datasets to query. Leave it empty to query all the datasets.
	Datasets []string `json:"datasets"`
	// Set a Flag to describe how to combine the filters on the query.
	FilterCombination ObservabilityQueryNewResponseParametersFilterCombination `json:"filterCombination"`
	// Configure the Filters to apply to the query. Supports nested groups via kind:
	// 'group'.
	Filters []ObservabilityQueryNewResponseParametersFilter `json:"filters"`
	// Define how to group the results of the query.
	GroupBys []ObservabilityQueryNewResponseParametersGroupBy `json:"groupBys"`
	// Configure the Having clauses that filter on calculations in the query result.
	Havings []ObservabilityQueryNewResponseParametersHaving `json:"havings"`
	// Set a limit on the number of results / records returned by the query
	Limit int64 `json:"limit"`
	// Define an expression to search using full-text search.
	Needle ObservabilityQueryNewResponseParametersNeedle `json:"needle"`
	// Configure the order of the results returned by the query.
	OrderBy ObservabilityQueryNewResponseParametersOrderBy `json:"orderBy"`
	JSON    observabilityQueryNewResponseParametersJSON    `json:"-"`
}

// observabilityQueryNewResponseParametersJSON contains the JSON metadata for the
// struct [ObservabilityQueryNewResponseParameters]
type observabilityQueryNewResponseParametersJSON struct {
	Calculations      apijson.Field
	Datasets          apijson.Field
	FilterCombination apijson.Field
	Filters           apijson.Field
	GroupBys          apijson.Field
	Havings           apijson.Field
	Limit             apijson.Field
	Needle            apijson.Field
	OrderBy           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityQueryNewResponseParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryNewResponseParametersJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryNewResponseParametersCalculation struct {
	Operator ObservabilityQueryNewResponseParametersCalculationsOperator `json:"operator" api:"required"`
	Alias    string                                                      `json:"alias"`
	Key      string                                                      `json:"key"`
	KeyType  ObservabilityQueryNewResponseParametersCalculationsKeyType  `json:"keyType"`
	JSON     observabilityQueryNewResponseParametersCalculationJSON      `json:"-"`
}

// observabilityQueryNewResponseParametersCalculationJSON contains the JSON
// metadata for the struct [ObservabilityQueryNewResponseParametersCalculation]
type observabilityQueryNewResponseParametersCalculationJSON struct {
	Operator    apijson.Field
	Alias       apijson.Field
	Key         apijson.Field
	KeyType     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryNewResponseParametersCalculation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryNewResponseParametersCalculationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryNewResponseParametersCalculationsOperator string

const (
	ObservabilityQueryNewResponseParametersCalculationsOperatorUniq              ObservabilityQueryNewResponseParametersCalculationsOperator = "uniq"
	ObservabilityQueryNewResponseParametersCalculationsOperatorCount             ObservabilityQueryNewResponseParametersCalculationsOperator = "count"
	ObservabilityQueryNewResponseParametersCalculationsOperatorMax               ObservabilityQueryNewResponseParametersCalculationsOperator = "max"
	ObservabilityQueryNewResponseParametersCalculationsOperatorMin               ObservabilityQueryNewResponseParametersCalculationsOperator = "min"
	ObservabilityQueryNewResponseParametersCalculationsOperatorSum               ObservabilityQueryNewResponseParametersCalculationsOperator = "sum"
	ObservabilityQueryNewResponseParametersCalculationsOperatorAvg               ObservabilityQueryNewResponseParametersCalculationsOperator = "avg"
	ObservabilityQueryNewResponseParametersCalculationsOperatorMedian            ObservabilityQueryNewResponseParametersCalculationsOperator = "median"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP001              ObservabilityQueryNewResponseParametersCalculationsOperator = "p001"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP01               ObservabilityQueryNewResponseParametersCalculationsOperator = "p01"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP05               ObservabilityQueryNewResponseParametersCalculationsOperator = "p05"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP10               ObservabilityQueryNewResponseParametersCalculationsOperator = "p10"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP25               ObservabilityQueryNewResponseParametersCalculationsOperator = "p25"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP75               ObservabilityQueryNewResponseParametersCalculationsOperator = "p75"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP90               ObservabilityQueryNewResponseParametersCalculationsOperator = "p90"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP95               ObservabilityQueryNewResponseParametersCalculationsOperator = "p95"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP99               ObservabilityQueryNewResponseParametersCalculationsOperator = "p99"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP999              ObservabilityQueryNewResponseParametersCalculationsOperator = "p999"
	ObservabilityQueryNewResponseParametersCalculationsOperatorStddev            ObservabilityQueryNewResponseParametersCalculationsOperator = "stddev"
	ObservabilityQueryNewResponseParametersCalculationsOperatorVariance          ObservabilityQueryNewResponseParametersCalculationsOperator = "variance"
	ObservabilityQueryNewResponseParametersCalculationsOperatorCountDistinct     ObservabilityQueryNewResponseParametersCalculationsOperator = "COUNT_DISTINCT"
	ObservabilityQueryNewResponseParametersCalculationsOperatorCountUppercase    ObservabilityQueryNewResponseParametersCalculationsOperator = "COUNT"
	ObservabilityQueryNewResponseParametersCalculationsOperatorMaxUppercase      ObservabilityQueryNewResponseParametersCalculationsOperator = "MAX"
	ObservabilityQueryNewResponseParametersCalculationsOperatorMinUppercase      ObservabilityQueryNewResponseParametersCalculationsOperator = "MIN"
	ObservabilityQueryNewResponseParametersCalculationsOperatorSumUppercase      ObservabilityQueryNewResponseParametersCalculationsOperator = "SUM"
	ObservabilityQueryNewResponseParametersCalculationsOperatorAvgUppercase      ObservabilityQueryNewResponseParametersCalculationsOperator = "AVG"
	ObservabilityQueryNewResponseParametersCalculationsOperatorMedianUppercase   ObservabilityQueryNewResponseParametersCalculationsOperator = "MEDIAN"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP001Uppercase     ObservabilityQueryNewResponseParametersCalculationsOperator = "P001"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP01Uppercase      ObservabilityQueryNewResponseParametersCalculationsOperator = "P01"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP05Uppercase      ObservabilityQueryNewResponseParametersCalculationsOperator = "P05"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP10Uppercase      ObservabilityQueryNewResponseParametersCalculationsOperator = "P10"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP25Uppercase      ObservabilityQueryNewResponseParametersCalculationsOperator = "P25"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP75Uppercase      ObservabilityQueryNewResponseParametersCalculationsOperator = "P75"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP90Uppercase      ObservabilityQueryNewResponseParametersCalculationsOperator = "P90"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP95Uppercase      ObservabilityQueryNewResponseParametersCalculationsOperator = "P95"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP99Uppercase      ObservabilityQueryNewResponseParametersCalculationsOperator = "P99"
	ObservabilityQueryNewResponseParametersCalculationsOperatorP999Uppercase     ObservabilityQueryNewResponseParametersCalculationsOperator = "P999"
	ObservabilityQueryNewResponseParametersCalculationsOperatorStddevUppercase   ObservabilityQueryNewResponseParametersCalculationsOperator = "STDDEV"
	ObservabilityQueryNewResponseParametersCalculationsOperatorVarianceUppercase ObservabilityQueryNewResponseParametersCalculationsOperator = "VARIANCE"
)

func (r ObservabilityQueryNewResponseParametersCalculationsOperator) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersCalculationsOperatorUniq, ObservabilityQueryNewResponseParametersCalculationsOperatorCount, ObservabilityQueryNewResponseParametersCalculationsOperatorMax, ObservabilityQueryNewResponseParametersCalculationsOperatorMin, ObservabilityQueryNewResponseParametersCalculationsOperatorSum, ObservabilityQueryNewResponseParametersCalculationsOperatorAvg, ObservabilityQueryNewResponseParametersCalculationsOperatorMedian, ObservabilityQueryNewResponseParametersCalculationsOperatorP001, ObservabilityQueryNewResponseParametersCalculationsOperatorP01, ObservabilityQueryNewResponseParametersCalculationsOperatorP05, ObservabilityQueryNewResponseParametersCalculationsOperatorP10, ObservabilityQueryNewResponseParametersCalculationsOperatorP25, ObservabilityQueryNewResponseParametersCalculationsOperatorP75, ObservabilityQueryNewResponseParametersCalculationsOperatorP90, ObservabilityQueryNewResponseParametersCalculationsOperatorP95, ObservabilityQueryNewResponseParametersCalculationsOperatorP99, ObservabilityQueryNewResponseParametersCalculationsOperatorP999, ObservabilityQueryNewResponseParametersCalculationsOperatorStddev, ObservabilityQueryNewResponseParametersCalculationsOperatorVariance, ObservabilityQueryNewResponseParametersCalculationsOperatorCountDistinct, ObservabilityQueryNewResponseParametersCalculationsOperatorCountUppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorMaxUppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorMinUppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorSumUppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorAvgUppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorMedianUppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorP001Uppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorP01Uppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorP05Uppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorP10Uppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorP25Uppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorP75Uppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorP90Uppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorP95Uppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorP99Uppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorP999Uppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorStddevUppercase, ObservabilityQueryNewResponseParametersCalculationsOperatorVarianceUppercase:
		return true
	}
	return false
}

type ObservabilityQueryNewResponseParametersCalculationsKeyType string

const (
	ObservabilityQueryNewResponseParametersCalculationsKeyTypeString  ObservabilityQueryNewResponseParametersCalculationsKeyType = "string"
	ObservabilityQueryNewResponseParametersCalculationsKeyTypeNumber  ObservabilityQueryNewResponseParametersCalculationsKeyType = "number"
	ObservabilityQueryNewResponseParametersCalculationsKeyTypeBoolean ObservabilityQueryNewResponseParametersCalculationsKeyType = "boolean"
)

func (r ObservabilityQueryNewResponseParametersCalculationsKeyType) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersCalculationsKeyTypeString, ObservabilityQueryNewResponseParametersCalculationsKeyTypeNumber, ObservabilityQueryNewResponseParametersCalculationsKeyTypeBoolean:
		return true
	}
	return false
}

// Set a Flag to describe how to combine the filters on the query.
type ObservabilityQueryNewResponseParametersFilterCombination string

const (
	ObservabilityQueryNewResponseParametersFilterCombinationAnd          ObservabilityQueryNewResponseParametersFilterCombination = "and"
	ObservabilityQueryNewResponseParametersFilterCombinationOr           ObservabilityQueryNewResponseParametersFilterCombination = "or"
	ObservabilityQueryNewResponseParametersFilterCombinationAndUppercase ObservabilityQueryNewResponseParametersFilterCombination = "AND"
	ObservabilityQueryNewResponseParametersFilterCombinationOrUppercase  ObservabilityQueryNewResponseParametersFilterCombination = "OR"
)

func (r ObservabilityQueryNewResponseParametersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersFilterCombinationAnd, ObservabilityQueryNewResponseParametersFilterCombinationOr, ObservabilityQueryNewResponseParametersFilterCombinationAndUppercase, ObservabilityQueryNewResponseParametersFilterCombinationOrUppercase:
		return true
	}
	return false
}

// Supports nested groups via kind: 'group'.
type ObservabilityQueryNewResponseParametersFilter struct {
	FilterCombination ObservabilityQueryNewResponseParametersFiltersFilterCombination `json:"filterCombination"`
	// This field can have the runtime type of [[]interface{}].
	Filters interface{} `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  string                                             `json:"key"`
	Kind ObservabilityQueryNewResponseParametersFiltersKind `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
	// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation ObservabilityQueryNewResponseParametersFiltersOperation `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type ObservabilityQueryNewResponseParametersFiltersType `json:"type"`
	// This field can have the runtime type of
	// [ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion].
	Value interface{}                                       `json:"value"`
	JSON  observabilityQueryNewResponseParametersFilterJSON `json:"-"`
	union ObservabilityQueryNewResponseParametersFiltersUnion
}

// observabilityQueryNewResponseParametersFilterJSON contains the JSON metadata for
// the struct [ObservabilityQueryNewResponseParametersFilter]
type observabilityQueryNewResponseParametersFilterJSON struct {
	FilterCombination apijson.Field
	Filters           apijson.Field
	Key               apijson.Field
	Kind              apijson.Field
	Operation         apijson.Field
	Type              apijson.Field
	Value             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r observabilityQueryNewResponseParametersFilterJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilityQueryNewResponseParametersFilter) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilityQueryNewResponseParametersFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObservabilityQueryNewResponseParametersFiltersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ObservabilityQueryNewResponseParametersFiltersObject],
// [ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeaf].
func (r ObservabilityQueryNewResponseParametersFilter) AsUnion() ObservabilityQueryNewResponseParametersFiltersUnion {
	return r.union
}

// Supports nested groups via kind: 'group'.
//
// Union satisfied by [ObservabilityQueryNewResponseParametersFiltersObject] or
// [ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeaf].
type ObservabilityQueryNewResponseParametersFiltersUnion interface {
	implementsObservabilityQueryNewResponseParametersFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityQueryNewResponseParametersFiltersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityQueryNewResponseParametersFiltersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeaf{}),
		},
	)
}

type ObservabilityQueryNewResponseParametersFiltersObject struct {
	FilterCombination ObservabilityQueryNewResponseParametersFiltersObjectFilterCombination `json:"filterCombination" api:"required"`
	Filters           []interface{}                                                         `json:"filters" api:"required"`
	Kind              ObservabilityQueryNewResponseParametersFiltersObjectKind              `json:"kind" api:"required"`
	JSON              observabilityQueryNewResponseParametersFiltersObjectJSON              `json:"-"`
}

// observabilityQueryNewResponseParametersFiltersObjectJSON contains the JSON
// metadata for the struct [ObservabilityQueryNewResponseParametersFiltersObject]
type observabilityQueryNewResponseParametersFiltersObjectJSON struct {
	FilterCombination apijson.Field
	Filters           apijson.Field
	Kind              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityQueryNewResponseParametersFiltersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryNewResponseParametersFiltersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityQueryNewResponseParametersFiltersObject) implementsObservabilityQueryNewResponseParametersFilter() {
}

type ObservabilityQueryNewResponseParametersFiltersObjectFilterCombination string

const (
	ObservabilityQueryNewResponseParametersFiltersObjectFilterCombinationAnd          ObservabilityQueryNewResponseParametersFiltersObjectFilterCombination = "and"
	ObservabilityQueryNewResponseParametersFiltersObjectFilterCombinationOr           ObservabilityQueryNewResponseParametersFiltersObjectFilterCombination = "or"
	ObservabilityQueryNewResponseParametersFiltersObjectFilterCombinationAndUppercase ObservabilityQueryNewResponseParametersFiltersObjectFilterCombination = "AND"
	ObservabilityQueryNewResponseParametersFiltersObjectFilterCombinationOrUppercase  ObservabilityQueryNewResponseParametersFiltersObjectFilterCombination = "OR"
)

func (r ObservabilityQueryNewResponseParametersFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersFiltersObjectFilterCombinationAnd, ObservabilityQueryNewResponseParametersFiltersObjectFilterCombinationOr, ObservabilityQueryNewResponseParametersFiltersObjectFilterCombinationAndUppercase, ObservabilityQueryNewResponseParametersFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityQueryNewResponseParametersFiltersObjectKind string

const (
	ObservabilityQueryNewResponseParametersFiltersObjectKindGroup ObservabilityQueryNewResponseParametersFiltersObjectKind = "group"
)

func (r ObservabilityQueryNewResponseParametersFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key string `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
	// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafType `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafKind `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion `json:"value"`
	JSON  observabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafJSON       `json:"-"`
}

// observabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafJSON
// contains the JSON metadata for the struct
// [ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeaf]
type observabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafJSON struct {
	Key         apijson.Field
	Operation   apijson.Field
	Type        apijson.Field
	Kind        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeaf) implementsObservabilityQueryNewResponseParametersFilter() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationEndsWith            ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "ends_with"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationEndsWithUppercase   ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "ENDS_WITH"
)

func (r ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationEndsWith, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafOperationEndsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafTypeString  ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafTypeString, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafKindFilter ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafKindFilter:
		return true
	}
	return false
}

// Comparison value. Must match actual values in your data — verify with the values
// endpoint. Ensure the value type (string/number/boolean) matches the field type.
// String comparisons are case-sensitive. Regex uses RE2 syntax (no
// lookaheads/lookbehinds).
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityQueryNewResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityQueryNewResponseParametersFiltersFilterCombination string

const (
	ObservabilityQueryNewResponseParametersFiltersFilterCombinationAnd          ObservabilityQueryNewResponseParametersFiltersFilterCombination = "and"
	ObservabilityQueryNewResponseParametersFiltersFilterCombinationOr           ObservabilityQueryNewResponseParametersFiltersFilterCombination = "or"
	ObservabilityQueryNewResponseParametersFiltersFilterCombinationAndUppercase ObservabilityQueryNewResponseParametersFiltersFilterCombination = "AND"
	ObservabilityQueryNewResponseParametersFiltersFilterCombinationOrUppercase  ObservabilityQueryNewResponseParametersFiltersFilterCombination = "OR"
)

func (r ObservabilityQueryNewResponseParametersFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersFiltersFilterCombinationAnd, ObservabilityQueryNewResponseParametersFiltersFilterCombinationOr, ObservabilityQueryNewResponseParametersFiltersFilterCombinationAndUppercase, ObservabilityQueryNewResponseParametersFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityQueryNewResponseParametersFiltersKind string

const (
	ObservabilityQueryNewResponseParametersFiltersKindGroup  ObservabilityQueryNewResponseParametersFiltersKind = "group"
	ObservabilityQueryNewResponseParametersFiltersKindFilter ObservabilityQueryNewResponseParametersFiltersKind = "filter"
)

func (r ObservabilityQueryNewResponseParametersFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersFiltersKindGroup, ObservabilityQueryNewResponseParametersFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityQueryNewResponseParametersFiltersOperation string

const (
	ObservabilityQueryNewResponseParametersFiltersOperationIncludes            ObservabilityQueryNewResponseParametersFiltersOperation = "includes"
	ObservabilityQueryNewResponseParametersFiltersOperationNotIncludes         ObservabilityQueryNewResponseParametersFiltersOperation = "not_includes"
	ObservabilityQueryNewResponseParametersFiltersOperationStartsWith          ObservabilityQueryNewResponseParametersFiltersOperation = "starts_with"
	ObservabilityQueryNewResponseParametersFiltersOperationEndsWith            ObservabilityQueryNewResponseParametersFiltersOperation = "ends_with"
	ObservabilityQueryNewResponseParametersFiltersOperationRegex               ObservabilityQueryNewResponseParametersFiltersOperation = "regex"
	ObservabilityQueryNewResponseParametersFiltersOperationExists              ObservabilityQueryNewResponseParametersFiltersOperation = "exists"
	ObservabilityQueryNewResponseParametersFiltersOperationIsNull              ObservabilityQueryNewResponseParametersFiltersOperation = "is_null"
	ObservabilityQueryNewResponseParametersFiltersOperationIn                  ObservabilityQueryNewResponseParametersFiltersOperation = "in"
	ObservabilityQueryNewResponseParametersFiltersOperationNotIn               ObservabilityQueryNewResponseParametersFiltersOperation = "not_in"
	ObservabilityQueryNewResponseParametersFiltersOperationEq                  ObservabilityQueryNewResponseParametersFiltersOperation = "eq"
	ObservabilityQueryNewResponseParametersFiltersOperationNeq                 ObservabilityQueryNewResponseParametersFiltersOperation = "neq"
	ObservabilityQueryNewResponseParametersFiltersOperationGt                  ObservabilityQueryNewResponseParametersFiltersOperation = "gt"
	ObservabilityQueryNewResponseParametersFiltersOperationGte                 ObservabilityQueryNewResponseParametersFiltersOperation = "gte"
	ObservabilityQueryNewResponseParametersFiltersOperationLt                  ObservabilityQueryNewResponseParametersFiltersOperation = "lt"
	ObservabilityQueryNewResponseParametersFiltersOperationLte                 ObservabilityQueryNewResponseParametersFiltersOperation = "lte"
	ObservabilityQueryNewResponseParametersFiltersOperationEquals              ObservabilityQueryNewResponseParametersFiltersOperation = "="
	ObservabilityQueryNewResponseParametersFiltersOperationNotEquals           ObservabilityQueryNewResponseParametersFiltersOperation = "!="
	ObservabilityQueryNewResponseParametersFiltersOperationGreater             ObservabilityQueryNewResponseParametersFiltersOperation = ">"
	ObservabilityQueryNewResponseParametersFiltersOperationGreaterOrEquals     ObservabilityQueryNewResponseParametersFiltersOperation = ">="
	ObservabilityQueryNewResponseParametersFiltersOperationLess                ObservabilityQueryNewResponseParametersFiltersOperation = "<"
	ObservabilityQueryNewResponseParametersFiltersOperationLessOrEquals        ObservabilityQueryNewResponseParametersFiltersOperation = "<="
	ObservabilityQueryNewResponseParametersFiltersOperationIncludesUppercase   ObservabilityQueryNewResponseParametersFiltersOperation = "INCLUDES"
	ObservabilityQueryNewResponseParametersFiltersOperationDoesNotInclude      ObservabilityQueryNewResponseParametersFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityQueryNewResponseParametersFiltersOperationMatchRegex          ObservabilityQueryNewResponseParametersFiltersOperation = "MATCH_REGEX"
	ObservabilityQueryNewResponseParametersFiltersOperationExistsUppercase     ObservabilityQueryNewResponseParametersFiltersOperation = "EXISTS"
	ObservabilityQueryNewResponseParametersFiltersOperationDoesNotExist        ObservabilityQueryNewResponseParametersFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityQueryNewResponseParametersFiltersOperationInUppercase         ObservabilityQueryNewResponseParametersFiltersOperation = "IN"
	ObservabilityQueryNewResponseParametersFiltersOperationNotInUppercase      ObservabilityQueryNewResponseParametersFiltersOperation = "NOT_IN"
	ObservabilityQueryNewResponseParametersFiltersOperationStartsWithUppercase ObservabilityQueryNewResponseParametersFiltersOperation = "STARTS_WITH"
	ObservabilityQueryNewResponseParametersFiltersOperationEndsWithUppercase   ObservabilityQueryNewResponseParametersFiltersOperation = "ENDS_WITH"
)

func (r ObservabilityQueryNewResponseParametersFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersFiltersOperationIncludes, ObservabilityQueryNewResponseParametersFiltersOperationNotIncludes, ObservabilityQueryNewResponseParametersFiltersOperationStartsWith, ObservabilityQueryNewResponseParametersFiltersOperationEndsWith, ObservabilityQueryNewResponseParametersFiltersOperationRegex, ObservabilityQueryNewResponseParametersFiltersOperationExists, ObservabilityQueryNewResponseParametersFiltersOperationIsNull, ObservabilityQueryNewResponseParametersFiltersOperationIn, ObservabilityQueryNewResponseParametersFiltersOperationNotIn, ObservabilityQueryNewResponseParametersFiltersOperationEq, ObservabilityQueryNewResponseParametersFiltersOperationNeq, ObservabilityQueryNewResponseParametersFiltersOperationGt, ObservabilityQueryNewResponseParametersFiltersOperationGte, ObservabilityQueryNewResponseParametersFiltersOperationLt, ObservabilityQueryNewResponseParametersFiltersOperationLte, ObservabilityQueryNewResponseParametersFiltersOperationEquals, ObservabilityQueryNewResponseParametersFiltersOperationNotEquals, ObservabilityQueryNewResponseParametersFiltersOperationGreater, ObservabilityQueryNewResponseParametersFiltersOperationGreaterOrEquals, ObservabilityQueryNewResponseParametersFiltersOperationLess, ObservabilityQueryNewResponseParametersFiltersOperationLessOrEquals, ObservabilityQueryNewResponseParametersFiltersOperationIncludesUppercase, ObservabilityQueryNewResponseParametersFiltersOperationDoesNotInclude, ObservabilityQueryNewResponseParametersFiltersOperationMatchRegex, ObservabilityQueryNewResponseParametersFiltersOperationExistsUppercase, ObservabilityQueryNewResponseParametersFiltersOperationDoesNotExist, ObservabilityQueryNewResponseParametersFiltersOperationInUppercase, ObservabilityQueryNewResponseParametersFiltersOperationNotInUppercase, ObservabilityQueryNewResponseParametersFiltersOperationStartsWithUppercase, ObservabilityQueryNewResponseParametersFiltersOperationEndsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityQueryNewResponseParametersFiltersType string

const (
	ObservabilityQueryNewResponseParametersFiltersTypeString  ObservabilityQueryNewResponseParametersFiltersType = "string"
	ObservabilityQueryNewResponseParametersFiltersTypeNumber  ObservabilityQueryNewResponseParametersFiltersType = "number"
	ObservabilityQueryNewResponseParametersFiltersTypeBoolean ObservabilityQueryNewResponseParametersFiltersType = "boolean"
)

func (r ObservabilityQueryNewResponseParametersFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersFiltersTypeString, ObservabilityQueryNewResponseParametersFiltersTypeNumber, ObservabilityQueryNewResponseParametersFiltersTypeBoolean:
		return true
	}
	return false
}

type ObservabilityQueryNewResponseParametersGroupBy struct {
	Type  ObservabilityQueryNewResponseParametersGroupBysType `json:"type" api:"required"`
	Value string                                              `json:"value" api:"required"`
	JSON  observabilityQueryNewResponseParametersGroupByJSON  `json:"-"`
}

// observabilityQueryNewResponseParametersGroupByJSON contains the JSON metadata
// for the struct [ObservabilityQueryNewResponseParametersGroupBy]
type observabilityQueryNewResponseParametersGroupByJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryNewResponseParametersGroupBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryNewResponseParametersGroupByJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryNewResponseParametersGroupBysType string

const (
	ObservabilityQueryNewResponseParametersGroupBysTypeString  ObservabilityQueryNewResponseParametersGroupBysType = "string"
	ObservabilityQueryNewResponseParametersGroupBysTypeNumber  ObservabilityQueryNewResponseParametersGroupBysType = "number"
	ObservabilityQueryNewResponseParametersGroupBysTypeBoolean ObservabilityQueryNewResponseParametersGroupBysType = "boolean"
)

func (r ObservabilityQueryNewResponseParametersGroupBysType) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersGroupBysTypeString, ObservabilityQueryNewResponseParametersGroupBysTypeNumber, ObservabilityQueryNewResponseParametersGroupBysTypeBoolean:
		return true
	}
	return false
}

type ObservabilityQueryNewResponseParametersHaving struct {
	Key       string                                                  `json:"key" api:"required"`
	Operation ObservabilityQueryNewResponseParametersHavingsOperation `json:"operation" api:"required"`
	Value     float64                                                 `json:"value" api:"required"`
	JSON      observabilityQueryNewResponseParametersHavingJSON       `json:"-"`
}

// observabilityQueryNewResponseParametersHavingJSON contains the JSON metadata for
// the struct [ObservabilityQueryNewResponseParametersHaving]
type observabilityQueryNewResponseParametersHavingJSON struct {
	Key         apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryNewResponseParametersHaving) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryNewResponseParametersHavingJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryNewResponseParametersHavingsOperation string

const (
	ObservabilityQueryNewResponseParametersHavingsOperationEq  ObservabilityQueryNewResponseParametersHavingsOperation = "eq"
	ObservabilityQueryNewResponseParametersHavingsOperationNeq ObservabilityQueryNewResponseParametersHavingsOperation = "neq"
	ObservabilityQueryNewResponseParametersHavingsOperationGt  ObservabilityQueryNewResponseParametersHavingsOperation = "gt"
	ObservabilityQueryNewResponseParametersHavingsOperationGte ObservabilityQueryNewResponseParametersHavingsOperation = "gte"
	ObservabilityQueryNewResponseParametersHavingsOperationLt  ObservabilityQueryNewResponseParametersHavingsOperation = "lt"
	ObservabilityQueryNewResponseParametersHavingsOperationLte ObservabilityQueryNewResponseParametersHavingsOperation = "lte"
)

func (r ObservabilityQueryNewResponseParametersHavingsOperation) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersHavingsOperationEq, ObservabilityQueryNewResponseParametersHavingsOperationNeq, ObservabilityQueryNewResponseParametersHavingsOperationGt, ObservabilityQueryNewResponseParametersHavingsOperationGte, ObservabilityQueryNewResponseParametersHavingsOperationLt, ObservabilityQueryNewResponseParametersHavingsOperationLte:
		return true
	}
	return false
}

// Define an expression to search using full-text search.
type ObservabilityQueryNewResponseParametersNeedle struct {
	Value     ObservabilityQueryNewResponseParametersNeedleValueUnion `json:"value" api:"required"`
	IsRegex   bool                                                    `json:"isRegex"`
	MatchCase bool                                                    `json:"matchCase"`
	JSON      observabilityQueryNewResponseParametersNeedleJSON       `json:"-"`
}

// observabilityQueryNewResponseParametersNeedleJSON contains the JSON metadata for
// the struct [ObservabilityQueryNewResponseParametersNeedle]
type observabilityQueryNewResponseParametersNeedleJSON struct {
	Value       apijson.Field
	IsRegex     apijson.Field
	MatchCase   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryNewResponseParametersNeedle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryNewResponseParametersNeedleJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityQueryNewResponseParametersNeedleValueUnion interface {
	ImplementsObservabilityQueryNewResponseParametersNeedleValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityQueryNewResponseParametersNeedleValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Configure the order of the results returned by the query.
type ObservabilityQueryNewResponseParametersOrderBy struct {
	// Configure which Calculation to order the results by.
	Value string `json:"value" api:"required"`
	// Set the order of the results
	Order ObservabilityQueryNewResponseParametersOrderByOrder `json:"order"`
	JSON  observabilityQueryNewResponseParametersOrderByJSON  `json:"-"`
}

// observabilityQueryNewResponseParametersOrderByJSON contains the JSON metadata
// for the struct [ObservabilityQueryNewResponseParametersOrderBy]
type observabilityQueryNewResponseParametersOrderByJSON struct {
	Value       apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryNewResponseParametersOrderBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryNewResponseParametersOrderByJSON) RawJSON() string {
	return r.raw
}

// Set the order of the results
type ObservabilityQueryNewResponseParametersOrderByOrder string

const (
	ObservabilityQueryNewResponseParametersOrderByOrderAsc  ObservabilityQueryNewResponseParametersOrderByOrder = "asc"
	ObservabilityQueryNewResponseParametersOrderByOrderDesc ObservabilityQueryNewResponseParametersOrderByOrder = "desc"
)

func (r ObservabilityQueryNewResponseParametersOrderByOrder) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseParametersOrderByOrderAsc, ObservabilityQueryNewResponseParametersOrderByOrderDesc:
		return true
	}
	return false
}

type ObservabilityQueryListResponse struct {
	ID string `json:"id" api:"required"`
	// If the query wasn't explcitly saved
	Adhoc       bool   `json:"adhoc" api:"required"`
	Created     string `json:"created" api:"required"`
	CreatedBy   string `json:"createdBy" api:"required"`
	Description string `json:"description" api:"required,nullable"`
	// Query name
	Name       string                                   `json:"name" api:"required"`
	Parameters ObservabilityQueryListResponseParameters `json:"parameters" api:"required"`
	Updated    string                                   `json:"updated" api:"required"`
	UpdatedBy  string                                   `json:"updatedBy" api:"required"`
	JSON       observabilityQueryListResponseJSON       `json:"-"`
}

// observabilityQueryListResponseJSON contains the JSON metadata for the struct
// [ObservabilityQueryListResponse]
type observabilityQueryListResponseJSON struct {
	ID          apijson.Field
	Adhoc       apijson.Field
	Created     apijson.Field
	CreatedBy   apijson.Field
	Description apijson.Field
	Name        apijson.Field
	Parameters  apijson.Field
	Updated     apijson.Field
	UpdatedBy   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryListResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryListResponseParameters struct {
	// Create Calculations to compute as part of the query.
	Calculations []ObservabilityQueryListResponseParametersCalculation `json:"calculations"`
	// Set the Datasets to query. Leave it empty to query all the datasets.
	Datasets []string `json:"datasets"`
	// Set a Flag to describe how to combine the filters on the query.
	FilterCombination ObservabilityQueryListResponseParametersFilterCombination `json:"filterCombination"`
	// Configure the Filters to apply to the query. Supports nested groups via kind:
	// 'group'.
	Filters []ObservabilityQueryListResponseParametersFilter `json:"filters"`
	// Define how to group the results of the query.
	GroupBys []ObservabilityQueryListResponseParametersGroupBy `json:"groupBys"`
	// Configure the Having clauses that filter on calculations in the query result.
	Havings []ObservabilityQueryListResponseParametersHaving `json:"havings"`
	// Set a limit on the number of results / records returned by the query
	Limit int64 `json:"limit"`
	// Define an expression to search using full-text search.
	Needle ObservabilityQueryListResponseParametersNeedle `json:"needle"`
	// Configure the order of the results returned by the query.
	OrderBy ObservabilityQueryListResponseParametersOrderBy `json:"orderBy"`
	JSON    observabilityQueryListResponseParametersJSON    `json:"-"`
}

// observabilityQueryListResponseParametersJSON contains the JSON metadata for the
// struct [ObservabilityQueryListResponseParameters]
type observabilityQueryListResponseParametersJSON struct {
	Calculations      apijson.Field
	Datasets          apijson.Field
	FilterCombination apijson.Field
	Filters           apijson.Field
	GroupBys          apijson.Field
	Havings           apijson.Field
	Limit             apijson.Field
	Needle            apijson.Field
	OrderBy           apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityQueryListResponseParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryListResponseParametersJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryListResponseParametersCalculation struct {
	Operator ObservabilityQueryListResponseParametersCalculationsOperator `json:"operator" api:"required"`
	Alias    string                                                       `json:"alias"`
	Key      string                                                       `json:"key"`
	KeyType  ObservabilityQueryListResponseParametersCalculationsKeyType  `json:"keyType"`
	JSON     observabilityQueryListResponseParametersCalculationJSON      `json:"-"`
}

// observabilityQueryListResponseParametersCalculationJSON contains the JSON
// metadata for the struct [ObservabilityQueryListResponseParametersCalculation]
type observabilityQueryListResponseParametersCalculationJSON struct {
	Operator    apijson.Field
	Alias       apijson.Field
	Key         apijson.Field
	KeyType     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryListResponseParametersCalculation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryListResponseParametersCalculationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryListResponseParametersCalculationsOperator string

const (
	ObservabilityQueryListResponseParametersCalculationsOperatorUniq              ObservabilityQueryListResponseParametersCalculationsOperator = "uniq"
	ObservabilityQueryListResponseParametersCalculationsOperatorCount             ObservabilityQueryListResponseParametersCalculationsOperator = "count"
	ObservabilityQueryListResponseParametersCalculationsOperatorMax               ObservabilityQueryListResponseParametersCalculationsOperator = "max"
	ObservabilityQueryListResponseParametersCalculationsOperatorMin               ObservabilityQueryListResponseParametersCalculationsOperator = "min"
	ObservabilityQueryListResponseParametersCalculationsOperatorSum               ObservabilityQueryListResponseParametersCalculationsOperator = "sum"
	ObservabilityQueryListResponseParametersCalculationsOperatorAvg               ObservabilityQueryListResponseParametersCalculationsOperator = "avg"
	ObservabilityQueryListResponseParametersCalculationsOperatorMedian            ObservabilityQueryListResponseParametersCalculationsOperator = "median"
	ObservabilityQueryListResponseParametersCalculationsOperatorP001              ObservabilityQueryListResponseParametersCalculationsOperator = "p001"
	ObservabilityQueryListResponseParametersCalculationsOperatorP01               ObservabilityQueryListResponseParametersCalculationsOperator = "p01"
	ObservabilityQueryListResponseParametersCalculationsOperatorP05               ObservabilityQueryListResponseParametersCalculationsOperator = "p05"
	ObservabilityQueryListResponseParametersCalculationsOperatorP10               ObservabilityQueryListResponseParametersCalculationsOperator = "p10"
	ObservabilityQueryListResponseParametersCalculationsOperatorP25               ObservabilityQueryListResponseParametersCalculationsOperator = "p25"
	ObservabilityQueryListResponseParametersCalculationsOperatorP75               ObservabilityQueryListResponseParametersCalculationsOperator = "p75"
	ObservabilityQueryListResponseParametersCalculationsOperatorP90               ObservabilityQueryListResponseParametersCalculationsOperator = "p90"
	ObservabilityQueryListResponseParametersCalculationsOperatorP95               ObservabilityQueryListResponseParametersCalculationsOperator = "p95"
	ObservabilityQueryListResponseParametersCalculationsOperatorP99               ObservabilityQueryListResponseParametersCalculationsOperator = "p99"
	ObservabilityQueryListResponseParametersCalculationsOperatorP999              ObservabilityQueryListResponseParametersCalculationsOperator = "p999"
	ObservabilityQueryListResponseParametersCalculationsOperatorStddev            ObservabilityQueryListResponseParametersCalculationsOperator = "stddev"
	ObservabilityQueryListResponseParametersCalculationsOperatorVariance          ObservabilityQueryListResponseParametersCalculationsOperator = "variance"
	ObservabilityQueryListResponseParametersCalculationsOperatorCountDistinct     ObservabilityQueryListResponseParametersCalculationsOperator = "COUNT_DISTINCT"
	ObservabilityQueryListResponseParametersCalculationsOperatorCountUppercase    ObservabilityQueryListResponseParametersCalculationsOperator = "COUNT"
	ObservabilityQueryListResponseParametersCalculationsOperatorMaxUppercase      ObservabilityQueryListResponseParametersCalculationsOperator = "MAX"
	ObservabilityQueryListResponseParametersCalculationsOperatorMinUppercase      ObservabilityQueryListResponseParametersCalculationsOperator = "MIN"
	ObservabilityQueryListResponseParametersCalculationsOperatorSumUppercase      ObservabilityQueryListResponseParametersCalculationsOperator = "SUM"
	ObservabilityQueryListResponseParametersCalculationsOperatorAvgUppercase      ObservabilityQueryListResponseParametersCalculationsOperator = "AVG"
	ObservabilityQueryListResponseParametersCalculationsOperatorMedianUppercase   ObservabilityQueryListResponseParametersCalculationsOperator = "MEDIAN"
	ObservabilityQueryListResponseParametersCalculationsOperatorP001Uppercase     ObservabilityQueryListResponseParametersCalculationsOperator = "P001"
	ObservabilityQueryListResponseParametersCalculationsOperatorP01Uppercase      ObservabilityQueryListResponseParametersCalculationsOperator = "P01"
	ObservabilityQueryListResponseParametersCalculationsOperatorP05Uppercase      ObservabilityQueryListResponseParametersCalculationsOperator = "P05"
	ObservabilityQueryListResponseParametersCalculationsOperatorP10Uppercase      ObservabilityQueryListResponseParametersCalculationsOperator = "P10"
	ObservabilityQueryListResponseParametersCalculationsOperatorP25Uppercase      ObservabilityQueryListResponseParametersCalculationsOperator = "P25"
	ObservabilityQueryListResponseParametersCalculationsOperatorP75Uppercase      ObservabilityQueryListResponseParametersCalculationsOperator = "P75"
	ObservabilityQueryListResponseParametersCalculationsOperatorP90Uppercase      ObservabilityQueryListResponseParametersCalculationsOperator = "P90"
	ObservabilityQueryListResponseParametersCalculationsOperatorP95Uppercase      ObservabilityQueryListResponseParametersCalculationsOperator = "P95"
	ObservabilityQueryListResponseParametersCalculationsOperatorP99Uppercase      ObservabilityQueryListResponseParametersCalculationsOperator = "P99"
	ObservabilityQueryListResponseParametersCalculationsOperatorP999Uppercase     ObservabilityQueryListResponseParametersCalculationsOperator = "P999"
	ObservabilityQueryListResponseParametersCalculationsOperatorStddevUppercase   ObservabilityQueryListResponseParametersCalculationsOperator = "STDDEV"
	ObservabilityQueryListResponseParametersCalculationsOperatorVarianceUppercase ObservabilityQueryListResponseParametersCalculationsOperator = "VARIANCE"
)

func (r ObservabilityQueryListResponseParametersCalculationsOperator) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersCalculationsOperatorUniq, ObservabilityQueryListResponseParametersCalculationsOperatorCount, ObservabilityQueryListResponseParametersCalculationsOperatorMax, ObservabilityQueryListResponseParametersCalculationsOperatorMin, ObservabilityQueryListResponseParametersCalculationsOperatorSum, ObservabilityQueryListResponseParametersCalculationsOperatorAvg, ObservabilityQueryListResponseParametersCalculationsOperatorMedian, ObservabilityQueryListResponseParametersCalculationsOperatorP001, ObservabilityQueryListResponseParametersCalculationsOperatorP01, ObservabilityQueryListResponseParametersCalculationsOperatorP05, ObservabilityQueryListResponseParametersCalculationsOperatorP10, ObservabilityQueryListResponseParametersCalculationsOperatorP25, ObservabilityQueryListResponseParametersCalculationsOperatorP75, ObservabilityQueryListResponseParametersCalculationsOperatorP90, ObservabilityQueryListResponseParametersCalculationsOperatorP95, ObservabilityQueryListResponseParametersCalculationsOperatorP99, ObservabilityQueryListResponseParametersCalculationsOperatorP999, ObservabilityQueryListResponseParametersCalculationsOperatorStddev, ObservabilityQueryListResponseParametersCalculationsOperatorVariance, ObservabilityQueryListResponseParametersCalculationsOperatorCountDistinct, ObservabilityQueryListResponseParametersCalculationsOperatorCountUppercase, ObservabilityQueryListResponseParametersCalculationsOperatorMaxUppercase, ObservabilityQueryListResponseParametersCalculationsOperatorMinUppercase, ObservabilityQueryListResponseParametersCalculationsOperatorSumUppercase, ObservabilityQueryListResponseParametersCalculationsOperatorAvgUppercase, ObservabilityQueryListResponseParametersCalculationsOperatorMedianUppercase, ObservabilityQueryListResponseParametersCalculationsOperatorP001Uppercase, ObservabilityQueryListResponseParametersCalculationsOperatorP01Uppercase, ObservabilityQueryListResponseParametersCalculationsOperatorP05Uppercase, ObservabilityQueryListResponseParametersCalculationsOperatorP10Uppercase, ObservabilityQueryListResponseParametersCalculationsOperatorP25Uppercase, ObservabilityQueryListResponseParametersCalculationsOperatorP75Uppercase, ObservabilityQueryListResponseParametersCalculationsOperatorP90Uppercase, ObservabilityQueryListResponseParametersCalculationsOperatorP95Uppercase, ObservabilityQueryListResponseParametersCalculationsOperatorP99Uppercase, ObservabilityQueryListResponseParametersCalculationsOperatorP999Uppercase, ObservabilityQueryListResponseParametersCalculationsOperatorStddevUppercase, ObservabilityQueryListResponseParametersCalculationsOperatorVarianceUppercase:
		return true
	}
	return false
}

type ObservabilityQueryListResponseParametersCalculationsKeyType string

const (
	ObservabilityQueryListResponseParametersCalculationsKeyTypeString  ObservabilityQueryListResponseParametersCalculationsKeyType = "string"
	ObservabilityQueryListResponseParametersCalculationsKeyTypeNumber  ObservabilityQueryListResponseParametersCalculationsKeyType = "number"
	ObservabilityQueryListResponseParametersCalculationsKeyTypeBoolean ObservabilityQueryListResponseParametersCalculationsKeyType = "boolean"
)

func (r ObservabilityQueryListResponseParametersCalculationsKeyType) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersCalculationsKeyTypeString, ObservabilityQueryListResponseParametersCalculationsKeyTypeNumber, ObservabilityQueryListResponseParametersCalculationsKeyTypeBoolean:
		return true
	}
	return false
}

// Set a Flag to describe how to combine the filters on the query.
type ObservabilityQueryListResponseParametersFilterCombination string

const (
	ObservabilityQueryListResponseParametersFilterCombinationAnd          ObservabilityQueryListResponseParametersFilterCombination = "and"
	ObservabilityQueryListResponseParametersFilterCombinationOr           ObservabilityQueryListResponseParametersFilterCombination = "or"
	ObservabilityQueryListResponseParametersFilterCombinationAndUppercase ObservabilityQueryListResponseParametersFilterCombination = "AND"
	ObservabilityQueryListResponseParametersFilterCombinationOrUppercase  ObservabilityQueryListResponseParametersFilterCombination = "OR"
)

func (r ObservabilityQueryListResponseParametersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersFilterCombinationAnd, ObservabilityQueryListResponseParametersFilterCombinationOr, ObservabilityQueryListResponseParametersFilterCombinationAndUppercase, ObservabilityQueryListResponseParametersFilterCombinationOrUppercase:
		return true
	}
	return false
}

// Supports nested groups via kind: 'group'.
type ObservabilityQueryListResponseParametersFilter struct {
	FilterCombination ObservabilityQueryListResponseParametersFiltersFilterCombination `json:"filterCombination"`
	// This field can have the runtime type of [[]interface{}].
	Filters interface{} `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  string                                              `json:"key"`
	Kind ObservabilityQueryListResponseParametersFiltersKind `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
	// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation ObservabilityQueryListResponseParametersFiltersOperation `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type ObservabilityQueryListResponseParametersFiltersType `json:"type"`
	// This field can have the runtime type of
	// [ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion].
	Value interface{}                                        `json:"value"`
	JSON  observabilityQueryListResponseParametersFilterJSON `json:"-"`
	union ObservabilityQueryListResponseParametersFiltersUnion
}

// observabilityQueryListResponseParametersFilterJSON contains the JSON metadata
// for the struct [ObservabilityQueryListResponseParametersFilter]
type observabilityQueryListResponseParametersFilterJSON struct {
	FilterCombination apijson.Field
	Filters           apijson.Field
	Key               apijson.Field
	Kind              apijson.Field
	Operation         apijson.Field
	Type              apijson.Field
	Value             apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r observabilityQueryListResponseParametersFilterJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilityQueryListResponseParametersFilter) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilityQueryListResponseParametersFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObservabilityQueryListResponseParametersFiltersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ObservabilityQueryListResponseParametersFiltersObject],
// [ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeaf].
func (r ObservabilityQueryListResponseParametersFilter) AsUnion() ObservabilityQueryListResponseParametersFiltersUnion {
	return r.union
}

// Supports nested groups via kind: 'group'.
//
// Union satisfied by [ObservabilityQueryListResponseParametersFiltersObject] or
// [ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeaf].
type ObservabilityQueryListResponseParametersFiltersUnion interface {
	implementsObservabilityQueryListResponseParametersFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityQueryListResponseParametersFiltersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityQueryListResponseParametersFiltersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeaf{}),
		},
	)
}

type ObservabilityQueryListResponseParametersFiltersObject struct {
	FilterCombination ObservabilityQueryListResponseParametersFiltersObjectFilterCombination `json:"filterCombination" api:"required"`
	Filters           []interface{}                                                          `json:"filters" api:"required"`
	Kind              ObservabilityQueryListResponseParametersFiltersObjectKind              `json:"kind" api:"required"`
	JSON              observabilityQueryListResponseParametersFiltersObjectJSON              `json:"-"`
}

// observabilityQueryListResponseParametersFiltersObjectJSON contains the JSON
// metadata for the struct [ObservabilityQueryListResponseParametersFiltersObject]
type observabilityQueryListResponseParametersFiltersObjectJSON struct {
	FilterCombination apijson.Field
	Filters           apijson.Field
	Kind              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityQueryListResponseParametersFiltersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryListResponseParametersFiltersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityQueryListResponseParametersFiltersObject) implementsObservabilityQueryListResponseParametersFilter() {
}

type ObservabilityQueryListResponseParametersFiltersObjectFilterCombination string

const (
	ObservabilityQueryListResponseParametersFiltersObjectFilterCombinationAnd          ObservabilityQueryListResponseParametersFiltersObjectFilterCombination = "and"
	ObservabilityQueryListResponseParametersFiltersObjectFilterCombinationOr           ObservabilityQueryListResponseParametersFiltersObjectFilterCombination = "or"
	ObservabilityQueryListResponseParametersFiltersObjectFilterCombinationAndUppercase ObservabilityQueryListResponseParametersFiltersObjectFilterCombination = "AND"
	ObservabilityQueryListResponseParametersFiltersObjectFilterCombinationOrUppercase  ObservabilityQueryListResponseParametersFiltersObjectFilterCombination = "OR"
)

func (r ObservabilityQueryListResponseParametersFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersFiltersObjectFilterCombinationAnd, ObservabilityQueryListResponseParametersFiltersObjectFilterCombinationOr, ObservabilityQueryListResponseParametersFiltersObjectFilterCombinationAndUppercase, ObservabilityQueryListResponseParametersFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityQueryListResponseParametersFiltersObjectKind string

const (
	ObservabilityQueryListResponseParametersFiltersObjectKindGroup ObservabilityQueryListResponseParametersFiltersObjectKind = "group"
)

func (r ObservabilityQueryListResponseParametersFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key string `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
	// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafType `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafKind `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion `json:"value"`
	JSON  observabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafJSON       `json:"-"`
}

// observabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafJSON
// contains the JSON metadata for the struct
// [ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeaf]
type observabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafJSON struct {
	Key         apijson.Field
	Operation   apijson.Field
	Type        apijson.Field
	Kind        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeaf) implementsObservabilityQueryListResponseParametersFilter() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationEndsWith            ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "ends_with"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationEndsWithUppercase   ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation = "ENDS_WITH"
)

func (r ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationEndsWith, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafOperationEndsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafTypeString  ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafTypeString, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafKindFilter ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafKindFilter:
		return true
	}
	return false
}

// Comparison value. Must match actual values in your data — verify with the values
// endpoint. Ensure the value type (string/number/boolean) matches the field type.
// String comparisons are case-sensitive. Regex uses RE2 syntax (no
// lookaheads/lookbehinds).
//
// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityQueryListResponseParametersFiltersWorkersObservabilityFilterLeafValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

type ObservabilityQueryListResponseParametersFiltersFilterCombination string

const (
	ObservabilityQueryListResponseParametersFiltersFilterCombinationAnd          ObservabilityQueryListResponseParametersFiltersFilterCombination = "and"
	ObservabilityQueryListResponseParametersFiltersFilterCombinationOr           ObservabilityQueryListResponseParametersFiltersFilterCombination = "or"
	ObservabilityQueryListResponseParametersFiltersFilterCombinationAndUppercase ObservabilityQueryListResponseParametersFiltersFilterCombination = "AND"
	ObservabilityQueryListResponseParametersFiltersFilterCombinationOrUppercase  ObservabilityQueryListResponseParametersFiltersFilterCombination = "OR"
)

func (r ObservabilityQueryListResponseParametersFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersFiltersFilterCombinationAnd, ObservabilityQueryListResponseParametersFiltersFilterCombinationOr, ObservabilityQueryListResponseParametersFiltersFilterCombinationAndUppercase, ObservabilityQueryListResponseParametersFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityQueryListResponseParametersFiltersKind string

const (
	ObservabilityQueryListResponseParametersFiltersKindGroup  ObservabilityQueryListResponseParametersFiltersKind = "group"
	ObservabilityQueryListResponseParametersFiltersKindFilter ObservabilityQueryListResponseParametersFiltersKind = "filter"
)

func (r ObservabilityQueryListResponseParametersFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersFiltersKindGroup, ObservabilityQueryListResponseParametersFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityQueryListResponseParametersFiltersOperation string

const (
	ObservabilityQueryListResponseParametersFiltersOperationIncludes            ObservabilityQueryListResponseParametersFiltersOperation = "includes"
	ObservabilityQueryListResponseParametersFiltersOperationNotIncludes         ObservabilityQueryListResponseParametersFiltersOperation = "not_includes"
	ObservabilityQueryListResponseParametersFiltersOperationStartsWith          ObservabilityQueryListResponseParametersFiltersOperation = "starts_with"
	ObservabilityQueryListResponseParametersFiltersOperationEndsWith            ObservabilityQueryListResponseParametersFiltersOperation = "ends_with"
	ObservabilityQueryListResponseParametersFiltersOperationRegex               ObservabilityQueryListResponseParametersFiltersOperation = "regex"
	ObservabilityQueryListResponseParametersFiltersOperationExists              ObservabilityQueryListResponseParametersFiltersOperation = "exists"
	ObservabilityQueryListResponseParametersFiltersOperationIsNull              ObservabilityQueryListResponseParametersFiltersOperation = "is_null"
	ObservabilityQueryListResponseParametersFiltersOperationIn                  ObservabilityQueryListResponseParametersFiltersOperation = "in"
	ObservabilityQueryListResponseParametersFiltersOperationNotIn               ObservabilityQueryListResponseParametersFiltersOperation = "not_in"
	ObservabilityQueryListResponseParametersFiltersOperationEq                  ObservabilityQueryListResponseParametersFiltersOperation = "eq"
	ObservabilityQueryListResponseParametersFiltersOperationNeq                 ObservabilityQueryListResponseParametersFiltersOperation = "neq"
	ObservabilityQueryListResponseParametersFiltersOperationGt                  ObservabilityQueryListResponseParametersFiltersOperation = "gt"
	ObservabilityQueryListResponseParametersFiltersOperationGte                 ObservabilityQueryListResponseParametersFiltersOperation = "gte"
	ObservabilityQueryListResponseParametersFiltersOperationLt                  ObservabilityQueryListResponseParametersFiltersOperation = "lt"
	ObservabilityQueryListResponseParametersFiltersOperationLte                 ObservabilityQueryListResponseParametersFiltersOperation = "lte"
	ObservabilityQueryListResponseParametersFiltersOperationEquals              ObservabilityQueryListResponseParametersFiltersOperation = "="
	ObservabilityQueryListResponseParametersFiltersOperationNotEquals           ObservabilityQueryListResponseParametersFiltersOperation = "!="
	ObservabilityQueryListResponseParametersFiltersOperationGreater             ObservabilityQueryListResponseParametersFiltersOperation = ">"
	ObservabilityQueryListResponseParametersFiltersOperationGreaterOrEquals     ObservabilityQueryListResponseParametersFiltersOperation = ">="
	ObservabilityQueryListResponseParametersFiltersOperationLess                ObservabilityQueryListResponseParametersFiltersOperation = "<"
	ObservabilityQueryListResponseParametersFiltersOperationLessOrEquals        ObservabilityQueryListResponseParametersFiltersOperation = "<="
	ObservabilityQueryListResponseParametersFiltersOperationIncludesUppercase   ObservabilityQueryListResponseParametersFiltersOperation = "INCLUDES"
	ObservabilityQueryListResponseParametersFiltersOperationDoesNotInclude      ObservabilityQueryListResponseParametersFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityQueryListResponseParametersFiltersOperationMatchRegex          ObservabilityQueryListResponseParametersFiltersOperation = "MATCH_REGEX"
	ObservabilityQueryListResponseParametersFiltersOperationExistsUppercase     ObservabilityQueryListResponseParametersFiltersOperation = "EXISTS"
	ObservabilityQueryListResponseParametersFiltersOperationDoesNotExist        ObservabilityQueryListResponseParametersFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityQueryListResponseParametersFiltersOperationInUppercase         ObservabilityQueryListResponseParametersFiltersOperation = "IN"
	ObservabilityQueryListResponseParametersFiltersOperationNotInUppercase      ObservabilityQueryListResponseParametersFiltersOperation = "NOT_IN"
	ObservabilityQueryListResponseParametersFiltersOperationStartsWithUppercase ObservabilityQueryListResponseParametersFiltersOperation = "STARTS_WITH"
	ObservabilityQueryListResponseParametersFiltersOperationEndsWithUppercase   ObservabilityQueryListResponseParametersFiltersOperation = "ENDS_WITH"
)

func (r ObservabilityQueryListResponseParametersFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersFiltersOperationIncludes, ObservabilityQueryListResponseParametersFiltersOperationNotIncludes, ObservabilityQueryListResponseParametersFiltersOperationStartsWith, ObservabilityQueryListResponseParametersFiltersOperationEndsWith, ObservabilityQueryListResponseParametersFiltersOperationRegex, ObservabilityQueryListResponseParametersFiltersOperationExists, ObservabilityQueryListResponseParametersFiltersOperationIsNull, ObservabilityQueryListResponseParametersFiltersOperationIn, ObservabilityQueryListResponseParametersFiltersOperationNotIn, ObservabilityQueryListResponseParametersFiltersOperationEq, ObservabilityQueryListResponseParametersFiltersOperationNeq, ObservabilityQueryListResponseParametersFiltersOperationGt, ObservabilityQueryListResponseParametersFiltersOperationGte, ObservabilityQueryListResponseParametersFiltersOperationLt, ObservabilityQueryListResponseParametersFiltersOperationLte, ObservabilityQueryListResponseParametersFiltersOperationEquals, ObservabilityQueryListResponseParametersFiltersOperationNotEquals, ObservabilityQueryListResponseParametersFiltersOperationGreater, ObservabilityQueryListResponseParametersFiltersOperationGreaterOrEquals, ObservabilityQueryListResponseParametersFiltersOperationLess, ObservabilityQueryListResponseParametersFiltersOperationLessOrEquals, ObservabilityQueryListResponseParametersFiltersOperationIncludesUppercase, ObservabilityQueryListResponseParametersFiltersOperationDoesNotInclude, ObservabilityQueryListResponseParametersFiltersOperationMatchRegex, ObservabilityQueryListResponseParametersFiltersOperationExistsUppercase, ObservabilityQueryListResponseParametersFiltersOperationDoesNotExist, ObservabilityQueryListResponseParametersFiltersOperationInUppercase, ObservabilityQueryListResponseParametersFiltersOperationNotInUppercase, ObservabilityQueryListResponseParametersFiltersOperationStartsWithUppercase, ObservabilityQueryListResponseParametersFiltersOperationEndsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityQueryListResponseParametersFiltersType string

const (
	ObservabilityQueryListResponseParametersFiltersTypeString  ObservabilityQueryListResponseParametersFiltersType = "string"
	ObservabilityQueryListResponseParametersFiltersTypeNumber  ObservabilityQueryListResponseParametersFiltersType = "number"
	ObservabilityQueryListResponseParametersFiltersTypeBoolean ObservabilityQueryListResponseParametersFiltersType = "boolean"
)

func (r ObservabilityQueryListResponseParametersFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersFiltersTypeString, ObservabilityQueryListResponseParametersFiltersTypeNumber, ObservabilityQueryListResponseParametersFiltersTypeBoolean:
		return true
	}
	return false
}

type ObservabilityQueryListResponseParametersGroupBy struct {
	Type  ObservabilityQueryListResponseParametersGroupBysType `json:"type" api:"required"`
	Value string                                               `json:"value" api:"required"`
	JSON  observabilityQueryListResponseParametersGroupByJSON  `json:"-"`
}

// observabilityQueryListResponseParametersGroupByJSON contains the JSON metadata
// for the struct [ObservabilityQueryListResponseParametersGroupBy]
type observabilityQueryListResponseParametersGroupByJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryListResponseParametersGroupBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryListResponseParametersGroupByJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryListResponseParametersGroupBysType string

const (
	ObservabilityQueryListResponseParametersGroupBysTypeString  ObservabilityQueryListResponseParametersGroupBysType = "string"
	ObservabilityQueryListResponseParametersGroupBysTypeNumber  ObservabilityQueryListResponseParametersGroupBysType = "number"
	ObservabilityQueryListResponseParametersGroupBysTypeBoolean ObservabilityQueryListResponseParametersGroupBysType = "boolean"
)

func (r ObservabilityQueryListResponseParametersGroupBysType) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersGroupBysTypeString, ObservabilityQueryListResponseParametersGroupBysTypeNumber, ObservabilityQueryListResponseParametersGroupBysTypeBoolean:
		return true
	}
	return false
}

type ObservabilityQueryListResponseParametersHaving struct {
	Key       string                                                   `json:"key" api:"required"`
	Operation ObservabilityQueryListResponseParametersHavingsOperation `json:"operation" api:"required"`
	Value     float64                                                  `json:"value" api:"required"`
	JSON      observabilityQueryListResponseParametersHavingJSON       `json:"-"`
}

// observabilityQueryListResponseParametersHavingJSON contains the JSON metadata
// for the struct [ObservabilityQueryListResponseParametersHaving]
type observabilityQueryListResponseParametersHavingJSON struct {
	Key         apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryListResponseParametersHaving) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryListResponseParametersHavingJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryListResponseParametersHavingsOperation string

const (
	ObservabilityQueryListResponseParametersHavingsOperationEq  ObservabilityQueryListResponseParametersHavingsOperation = "eq"
	ObservabilityQueryListResponseParametersHavingsOperationNeq ObservabilityQueryListResponseParametersHavingsOperation = "neq"
	ObservabilityQueryListResponseParametersHavingsOperationGt  ObservabilityQueryListResponseParametersHavingsOperation = "gt"
	ObservabilityQueryListResponseParametersHavingsOperationGte ObservabilityQueryListResponseParametersHavingsOperation = "gte"
	ObservabilityQueryListResponseParametersHavingsOperationLt  ObservabilityQueryListResponseParametersHavingsOperation = "lt"
	ObservabilityQueryListResponseParametersHavingsOperationLte ObservabilityQueryListResponseParametersHavingsOperation = "lte"
)

func (r ObservabilityQueryListResponseParametersHavingsOperation) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersHavingsOperationEq, ObservabilityQueryListResponseParametersHavingsOperationNeq, ObservabilityQueryListResponseParametersHavingsOperationGt, ObservabilityQueryListResponseParametersHavingsOperationGte, ObservabilityQueryListResponseParametersHavingsOperationLt, ObservabilityQueryListResponseParametersHavingsOperationLte:
		return true
	}
	return false
}

// Define an expression to search using full-text search.
type ObservabilityQueryListResponseParametersNeedle struct {
	Value     ObservabilityQueryListResponseParametersNeedleValueUnion `json:"value" api:"required"`
	IsRegex   bool                                                     `json:"isRegex"`
	MatchCase bool                                                     `json:"matchCase"`
	JSON      observabilityQueryListResponseParametersNeedleJSON       `json:"-"`
}

// observabilityQueryListResponseParametersNeedleJSON contains the JSON metadata
// for the struct [ObservabilityQueryListResponseParametersNeedle]
type observabilityQueryListResponseParametersNeedleJSON struct {
	Value       apijson.Field
	IsRegex     apijson.Field
	MatchCase   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryListResponseParametersNeedle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryListResponseParametersNeedleJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityQueryListResponseParametersNeedleValueUnion interface {
	ImplementsObservabilityQueryListResponseParametersNeedleValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityQueryListResponseParametersNeedleValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
	)
}

// Configure the order of the results returned by the query.
type ObservabilityQueryListResponseParametersOrderBy struct {
	// Configure which Calculation to order the results by.
	Value string `json:"value" api:"required"`
	// Set the order of the results
	Order ObservabilityQueryListResponseParametersOrderByOrder `json:"order"`
	JSON  observabilityQueryListResponseParametersOrderByJSON  `json:"-"`
}

// observabilityQueryListResponseParametersOrderByJSON contains the JSON metadata
// for the struct [ObservabilityQueryListResponseParametersOrderBy]
type observabilityQueryListResponseParametersOrderByJSON struct {
	Value       apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryListResponseParametersOrderBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryListResponseParametersOrderByJSON) RawJSON() string {
	return r.raw
}

// Set the order of the results
type ObservabilityQueryListResponseParametersOrderByOrder string

const (
	ObservabilityQueryListResponseParametersOrderByOrderAsc  ObservabilityQueryListResponseParametersOrderByOrder = "asc"
	ObservabilityQueryListResponseParametersOrderByOrderDesc ObservabilityQueryListResponseParametersOrderByOrder = "desc"
)

func (r ObservabilityQueryListResponseParametersOrderByOrder) IsKnown() bool {
	switch r {
	case ObservabilityQueryListResponseParametersOrderByOrderAsc, ObservabilityQueryListResponseParametersOrderByOrderDesc:
		return true
	}
	return false
}

type ObservabilityQueryNewParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Description param.Field[string] `json:"description" api:"required"`
	// Query name
	Name       param.Field[string]                                `json:"name" api:"required"`
	Parameters param.Field[ObservabilityQueryNewParamsParameters] `json:"parameters" api:"required"`
}

func (r ObservabilityQueryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityQueryNewParamsParameters struct {
	// Create Calculations to compute as part of the query.
	Calculations param.Field[[]ObservabilityQueryNewParamsParametersCalculation] `json:"calculations"`
	// Set the Datasets to query. Leave it empty to query all the datasets.
	Datasets param.Field[[]string] `json:"datasets"`
	// Set a Flag to describe how to combine the filters on the query.
	FilterCombination param.Field[ObservabilityQueryNewParamsParametersFilterCombination] `json:"filterCombination"`
	// Configure the Filters to apply to the query. Supports nested groups via kind:
	// 'group'.
	Filters param.Field[[]ObservabilityQueryNewParamsParametersFilterUnion] `json:"filters"`
	// Define how to group the results of the query.
	GroupBys param.Field[[]ObservabilityQueryNewParamsParametersGroupBy] `json:"groupBys"`
	// Configure the Having clauses that filter on calculations in the query result.
	Havings param.Field[[]ObservabilityQueryNewParamsParametersHaving] `json:"havings"`
	// Set a limit on the number of results / records returned by the query
	Limit param.Field[int64] `json:"limit"`
	// Define an expression to search using full-text search.
	Needle param.Field[ObservabilityQueryNewParamsParametersNeedle] `json:"needle"`
	// Configure the order of the results returned by the query.
	OrderBy param.Field[ObservabilityQueryNewParamsParametersOrderBy] `json:"orderBy"`
}

func (r ObservabilityQueryNewParamsParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityQueryNewParamsParametersCalculation struct {
	Operator param.Field[ObservabilityQueryNewParamsParametersCalculationsOperator] `json:"operator" api:"required"`
	Alias    param.Field[string]                                                    `json:"alias"`
	Key      param.Field[string]                                                    `json:"key"`
	KeyType  param.Field[ObservabilityQueryNewParamsParametersCalculationsKeyType]  `json:"keyType"`
}

func (r ObservabilityQueryNewParamsParametersCalculation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityQueryNewParamsParametersCalculationsOperator string

const (
	ObservabilityQueryNewParamsParametersCalculationsOperatorUniq              ObservabilityQueryNewParamsParametersCalculationsOperator = "uniq"
	ObservabilityQueryNewParamsParametersCalculationsOperatorCount             ObservabilityQueryNewParamsParametersCalculationsOperator = "count"
	ObservabilityQueryNewParamsParametersCalculationsOperatorMax               ObservabilityQueryNewParamsParametersCalculationsOperator = "max"
	ObservabilityQueryNewParamsParametersCalculationsOperatorMin               ObservabilityQueryNewParamsParametersCalculationsOperator = "min"
	ObservabilityQueryNewParamsParametersCalculationsOperatorSum               ObservabilityQueryNewParamsParametersCalculationsOperator = "sum"
	ObservabilityQueryNewParamsParametersCalculationsOperatorAvg               ObservabilityQueryNewParamsParametersCalculationsOperator = "avg"
	ObservabilityQueryNewParamsParametersCalculationsOperatorMedian            ObservabilityQueryNewParamsParametersCalculationsOperator = "median"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP001              ObservabilityQueryNewParamsParametersCalculationsOperator = "p001"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP01               ObservabilityQueryNewParamsParametersCalculationsOperator = "p01"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP05               ObservabilityQueryNewParamsParametersCalculationsOperator = "p05"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP10               ObservabilityQueryNewParamsParametersCalculationsOperator = "p10"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP25               ObservabilityQueryNewParamsParametersCalculationsOperator = "p25"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP75               ObservabilityQueryNewParamsParametersCalculationsOperator = "p75"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP90               ObservabilityQueryNewParamsParametersCalculationsOperator = "p90"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP95               ObservabilityQueryNewParamsParametersCalculationsOperator = "p95"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP99               ObservabilityQueryNewParamsParametersCalculationsOperator = "p99"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP999              ObservabilityQueryNewParamsParametersCalculationsOperator = "p999"
	ObservabilityQueryNewParamsParametersCalculationsOperatorStddev            ObservabilityQueryNewParamsParametersCalculationsOperator = "stddev"
	ObservabilityQueryNewParamsParametersCalculationsOperatorVariance          ObservabilityQueryNewParamsParametersCalculationsOperator = "variance"
	ObservabilityQueryNewParamsParametersCalculationsOperatorCountDistinct     ObservabilityQueryNewParamsParametersCalculationsOperator = "COUNT_DISTINCT"
	ObservabilityQueryNewParamsParametersCalculationsOperatorCountUppercase    ObservabilityQueryNewParamsParametersCalculationsOperator = "COUNT"
	ObservabilityQueryNewParamsParametersCalculationsOperatorMaxUppercase      ObservabilityQueryNewParamsParametersCalculationsOperator = "MAX"
	ObservabilityQueryNewParamsParametersCalculationsOperatorMinUppercase      ObservabilityQueryNewParamsParametersCalculationsOperator = "MIN"
	ObservabilityQueryNewParamsParametersCalculationsOperatorSumUppercase      ObservabilityQueryNewParamsParametersCalculationsOperator = "SUM"
	ObservabilityQueryNewParamsParametersCalculationsOperatorAvgUppercase      ObservabilityQueryNewParamsParametersCalculationsOperator = "AVG"
	ObservabilityQueryNewParamsParametersCalculationsOperatorMedianUppercase   ObservabilityQueryNewParamsParametersCalculationsOperator = "MEDIAN"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP001Uppercase     ObservabilityQueryNewParamsParametersCalculationsOperator = "P001"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP01Uppercase      ObservabilityQueryNewParamsParametersCalculationsOperator = "P01"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP05Uppercase      ObservabilityQueryNewParamsParametersCalculationsOperator = "P05"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP10Uppercase      ObservabilityQueryNewParamsParametersCalculationsOperator = "P10"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP25Uppercase      ObservabilityQueryNewParamsParametersCalculationsOperator = "P25"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP75Uppercase      ObservabilityQueryNewParamsParametersCalculationsOperator = "P75"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP90Uppercase      ObservabilityQueryNewParamsParametersCalculationsOperator = "P90"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP95Uppercase      ObservabilityQueryNewParamsParametersCalculationsOperator = "P95"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP99Uppercase      ObservabilityQueryNewParamsParametersCalculationsOperator = "P99"
	ObservabilityQueryNewParamsParametersCalculationsOperatorP999Uppercase     ObservabilityQueryNewParamsParametersCalculationsOperator = "P999"
	ObservabilityQueryNewParamsParametersCalculationsOperatorStddevUppercase   ObservabilityQueryNewParamsParametersCalculationsOperator = "STDDEV"
	ObservabilityQueryNewParamsParametersCalculationsOperatorVarianceUppercase ObservabilityQueryNewParamsParametersCalculationsOperator = "VARIANCE"
)

func (r ObservabilityQueryNewParamsParametersCalculationsOperator) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersCalculationsOperatorUniq, ObservabilityQueryNewParamsParametersCalculationsOperatorCount, ObservabilityQueryNewParamsParametersCalculationsOperatorMax, ObservabilityQueryNewParamsParametersCalculationsOperatorMin, ObservabilityQueryNewParamsParametersCalculationsOperatorSum, ObservabilityQueryNewParamsParametersCalculationsOperatorAvg, ObservabilityQueryNewParamsParametersCalculationsOperatorMedian, ObservabilityQueryNewParamsParametersCalculationsOperatorP001, ObservabilityQueryNewParamsParametersCalculationsOperatorP01, ObservabilityQueryNewParamsParametersCalculationsOperatorP05, ObservabilityQueryNewParamsParametersCalculationsOperatorP10, ObservabilityQueryNewParamsParametersCalculationsOperatorP25, ObservabilityQueryNewParamsParametersCalculationsOperatorP75, ObservabilityQueryNewParamsParametersCalculationsOperatorP90, ObservabilityQueryNewParamsParametersCalculationsOperatorP95, ObservabilityQueryNewParamsParametersCalculationsOperatorP99, ObservabilityQueryNewParamsParametersCalculationsOperatorP999, ObservabilityQueryNewParamsParametersCalculationsOperatorStddev, ObservabilityQueryNewParamsParametersCalculationsOperatorVariance, ObservabilityQueryNewParamsParametersCalculationsOperatorCountDistinct, ObservabilityQueryNewParamsParametersCalculationsOperatorCountUppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorMaxUppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorMinUppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorSumUppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorAvgUppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorMedianUppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorP001Uppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorP01Uppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorP05Uppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorP10Uppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorP25Uppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorP75Uppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorP90Uppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorP95Uppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorP99Uppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorP999Uppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorStddevUppercase, ObservabilityQueryNewParamsParametersCalculationsOperatorVarianceUppercase:
		return true
	}
	return false
}

type ObservabilityQueryNewParamsParametersCalculationsKeyType string

const (
	ObservabilityQueryNewParamsParametersCalculationsKeyTypeString  ObservabilityQueryNewParamsParametersCalculationsKeyType = "string"
	ObservabilityQueryNewParamsParametersCalculationsKeyTypeNumber  ObservabilityQueryNewParamsParametersCalculationsKeyType = "number"
	ObservabilityQueryNewParamsParametersCalculationsKeyTypeBoolean ObservabilityQueryNewParamsParametersCalculationsKeyType = "boolean"
)

func (r ObservabilityQueryNewParamsParametersCalculationsKeyType) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersCalculationsKeyTypeString, ObservabilityQueryNewParamsParametersCalculationsKeyTypeNumber, ObservabilityQueryNewParamsParametersCalculationsKeyTypeBoolean:
		return true
	}
	return false
}

// Set a Flag to describe how to combine the filters on the query.
type ObservabilityQueryNewParamsParametersFilterCombination string

const (
	ObservabilityQueryNewParamsParametersFilterCombinationAnd          ObservabilityQueryNewParamsParametersFilterCombination = "and"
	ObservabilityQueryNewParamsParametersFilterCombinationOr           ObservabilityQueryNewParamsParametersFilterCombination = "or"
	ObservabilityQueryNewParamsParametersFilterCombinationAndUppercase ObservabilityQueryNewParamsParametersFilterCombination = "AND"
	ObservabilityQueryNewParamsParametersFilterCombinationOrUppercase  ObservabilityQueryNewParamsParametersFilterCombination = "OR"
)

func (r ObservabilityQueryNewParamsParametersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersFilterCombinationAnd, ObservabilityQueryNewParamsParametersFilterCombinationOr, ObservabilityQueryNewParamsParametersFilterCombinationAndUppercase, ObservabilityQueryNewParamsParametersFilterCombinationOrUppercase:
		return true
	}
	return false
}

// Supports nested groups via kind: 'group'.
type ObservabilityQueryNewParamsParametersFilter struct {
	FilterCombination param.Field[ObservabilityQueryNewParamsParametersFiltersFilterCombination] `json:"filterCombination"`
	Filters           param.Field[interface{}]                                                   `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  param.Field[string]                                           `json:"key"`
	Kind param.Field[ObservabilityQueryNewParamsParametersFiltersKind] `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
	// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityQueryNewParamsParametersFiltersOperation] `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type  param.Field[ObservabilityQueryNewParamsParametersFiltersType] `json:"type"`
	Value param.Field[interface{}]                                      `json:"value"`
}

func (r ObservabilityQueryNewParamsParametersFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityQueryNewParamsParametersFilter) implementsObservabilityQueryNewParamsParametersFilterUnion() {
}

// Supports nested groups via kind: 'group'.
//
// Satisfied by [workers.ObservabilityQueryNewParamsParametersFiltersObject],
// [workers.ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeaf],
// [ObservabilityQueryNewParamsParametersFilter].
type ObservabilityQueryNewParamsParametersFilterUnion interface {
	implementsObservabilityQueryNewParamsParametersFilterUnion()
}

type ObservabilityQueryNewParamsParametersFiltersObject struct {
	FilterCombination param.Field[ObservabilityQueryNewParamsParametersFiltersObjectFilterCombination] `json:"filterCombination" api:"required"`
	Filters           param.Field[[]interface{}]                                                       `json:"filters" api:"required"`
	Kind              param.Field[ObservabilityQueryNewParamsParametersFiltersObjectKind]              `json:"kind" api:"required"`
}

func (r ObservabilityQueryNewParamsParametersFiltersObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityQueryNewParamsParametersFiltersObject) implementsObservabilityQueryNewParamsParametersFilterUnion() {
}

type ObservabilityQueryNewParamsParametersFiltersObjectFilterCombination string

const (
	ObservabilityQueryNewParamsParametersFiltersObjectFilterCombinationAnd          ObservabilityQueryNewParamsParametersFiltersObjectFilterCombination = "and"
	ObservabilityQueryNewParamsParametersFiltersObjectFilterCombinationOr           ObservabilityQueryNewParamsParametersFiltersObjectFilterCombination = "or"
	ObservabilityQueryNewParamsParametersFiltersObjectFilterCombinationAndUppercase ObservabilityQueryNewParamsParametersFiltersObjectFilterCombination = "AND"
	ObservabilityQueryNewParamsParametersFiltersObjectFilterCombinationOrUppercase  ObservabilityQueryNewParamsParametersFiltersObjectFilterCombination = "OR"
)

func (r ObservabilityQueryNewParamsParametersFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersFiltersObjectFilterCombinationAnd, ObservabilityQueryNewParamsParametersFiltersObjectFilterCombinationOr, ObservabilityQueryNewParamsParametersFiltersObjectFilterCombinationAndUppercase, ObservabilityQueryNewParamsParametersFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityQueryNewParamsParametersFiltersObjectKind string

const (
	ObservabilityQueryNewParamsParametersFiltersObjectKindGroup ObservabilityQueryNewParamsParametersFiltersObjectKind = "group"
)

func (r ObservabilityQueryNewParamsParametersFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key param.Field[string] `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
	// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation] `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type param.Field[ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafType] `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind param.Field[ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafKind] `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value param.Field[ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion] `json:"value"`
}

func (r ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeaf) implementsObservabilityQueryNewParamsParametersFilterUnion() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEndsWith            ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "ends_with"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEndsWithUppercase   ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "ENDS_WITH"
)

func (r ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEndsWith, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEndsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafTypeString  ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafTypeString, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafKindFilter ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafKindFilter:
		return true
	}
	return false
}

// Comparison value. Must match actual values in your data — verify with the values
// endpoint. Ensure the value type (string/number/boolean) matches the field type.
// String comparisons are case-sensitive. Regex uses RE2 syntax (no
// lookaheads/lookbehinds).
//
// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilityQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion()
}

type ObservabilityQueryNewParamsParametersFiltersFilterCombination string

const (
	ObservabilityQueryNewParamsParametersFiltersFilterCombinationAnd          ObservabilityQueryNewParamsParametersFiltersFilterCombination = "and"
	ObservabilityQueryNewParamsParametersFiltersFilterCombinationOr           ObservabilityQueryNewParamsParametersFiltersFilterCombination = "or"
	ObservabilityQueryNewParamsParametersFiltersFilterCombinationAndUppercase ObservabilityQueryNewParamsParametersFiltersFilterCombination = "AND"
	ObservabilityQueryNewParamsParametersFiltersFilterCombinationOrUppercase  ObservabilityQueryNewParamsParametersFiltersFilterCombination = "OR"
)

func (r ObservabilityQueryNewParamsParametersFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersFiltersFilterCombinationAnd, ObservabilityQueryNewParamsParametersFiltersFilterCombinationOr, ObservabilityQueryNewParamsParametersFiltersFilterCombinationAndUppercase, ObservabilityQueryNewParamsParametersFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityQueryNewParamsParametersFiltersKind string

const (
	ObservabilityQueryNewParamsParametersFiltersKindGroup  ObservabilityQueryNewParamsParametersFiltersKind = "group"
	ObservabilityQueryNewParamsParametersFiltersKindFilter ObservabilityQueryNewParamsParametersFiltersKind = "filter"
)

func (r ObservabilityQueryNewParamsParametersFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersFiltersKindGroup, ObservabilityQueryNewParamsParametersFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityQueryNewParamsParametersFiltersOperation string

const (
	ObservabilityQueryNewParamsParametersFiltersOperationIncludes            ObservabilityQueryNewParamsParametersFiltersOperation = "includes"
	ObservabilityQueryNewParamsParametersFiltersOperationNotIncludes         ObservabilityQueryNewParamsParametersFiltersOperation = "not_includes"
	ObservabilityQueryNewParamsParametersFiltersOperationStartsWith          ObservabilityQueryNewParamsParametersFiltersOperation = "starts_with"
	ObservabilityQueryNewParamsParametersFiltersOperationEndsWith            ObservabilityQueryNewParamsParametersFiltersOperation = "ends_with"
	ObservabilityQueryNewParamsParametersFiltersOperationRegex               ObservabilityQueryNewParamsParametersFiltersOperation = "regex"
	ObservabilityQueryNewParamsParametersFiltersOperationExists              ObservabilityQueryNewParamsParametersFiltersOperation = "exists"
	ObservabilityQueryNewParamsParametersFiltersOperationIsNull              ObservabilityQueryNewParamsParametersFiltersOperation = "is_null"
	ObservabilityQueryNewParamsParametersFiltersOperationIn                  ObservabilityQueryNewParamsParametersFiltersOperation = "in"
	ObservabilityQueryNewParamsParametersFiltersOperationNotIn               ObservabilityQueryNewParamsParametersFiltersOperation = "not_in"
	ObservabilityQueryNewParamsParametersFiltersOperationEq                  ObservabilityQueryNewParamsParametersFiltersOperation = "eq"
	ObservabilityQueryNewParamsParametersFiltersOperationNeq                 ObservabilityQueryNewParamsParametersFiltersOperation = "neq"
	ObservabilityQueryNewParamsParametersFiltersOperationGt                  ObservabilityQueryNewParamsParametersFiltersOperation = "gt"
	ObservabilityQueryNewParamsParametersFiltersOperationGte                 ObservabilityQueryNewParamsParametersFiltersOperation = "gte"
	ObservabilityQueryNewParamsParametersFiltersOperationLt                  ObservabilityQueryNewParamsParametersFiltersOperation = "lt"
	ObservabilityQueryNewParamsParametersFiltersOperationLte                 ObservabilityQueryNewParamsParametersFiltersOperation = "lte"
	ObservabilityQueryNewParamsParametersFiltersOperationEquals              ObservabilityQueryNewParamsParametersFiltersOperation = "="
	ObservabilityQueryNewParamsParametersFiltersOperationNotEquals           ObservabilityQueryNewParamsParametersFiltersOperation = "!="
	ObservabilityQueryNewParamsParametersFiltersOperationGreater             ObservabilityQueryNewParamsParametersFiltersOperation = ">"
	ObservabilityQueryNewParamsParametersFiltersOperationGreaterOrEquals     ObservabilityQueryNewParamsParametersFiltersOperation = ">="
	ObservabilityQueryNewParamsParametersFiltersOperationLess                ObservabilityQueryNewParamsParametersFiltersOperation = "<"
	ObservabilityQueryNewParamsParametersFiltersOperationLessOrEquals        ObservabilityQueryNewParamsParametersFiltersOperation = "<="
	ObservabilityQueryNewParamsParametersFiltersOperationIncludesUppercase   ObservabilityQueryNewParamsParametersFiltersOperation = "INCLUDES"
	ObservabilityQueryNewParamsParametersFiltersOperationDoesNotInclude      ObservabilityQueryNewParamsParametersFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityQueryNewParamsParametersFiltersOperationMatchRegex          ObservabilityQueryNewParamsParametersFiltersOperation = "MATCH_REGEX"
	ObservabilityQueryNewParamsParametersFiltersOperationExistsUppercase     ObservabilityQueryNewParamsParametersFiltersOperation = "EXISTS"
	ObservabilityQueryNewParamsParametersFiltersOperationDoesNotExist        ObservabilityQueryNewParamsParametersFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityQueryNewParamsParametersFiltersOperationInUppercase         ObservabilityQueryNewParamsParametersFiltersOperation = "IN"
	ObservabilityQueryNewParamsParametersFiltersOperationNotInUppercase      ObservabilityQueryNewParamsParametersFiltersOperation = "NOT_IN"
	ObservabilityQueryNewParamsParametersFiltersOperationStartsWithUppercase ObservabilityQueryNewParamsParametersFiltersOperation = "STARTS_WITH"
	ObservabilityQueryNewParamsParametersFiltersOperationEndsWithUppercase   ObservabilityQueryNewParamsParametersFiltersOperation = "ENDS_WITH"
)

func (r ObservabilityQueryNewParamsParametersFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersFiltersOperationIncludes, ObservabilityQueryNewParamsParametersFiltersOperationNotIncludes, ObservabilityQueryNewParamsParametersFiltersOperationStartsWith, ObservabilityQueryNewParamsParametersFiltersOperationEndsWith, ObservabilityQueryNewParamsParametersFiltersOperationRegex, ObservabilityQueryNewParamsParametersFiltersOperationExists, ObservabilityQueryNewParamsParametersFiltersOperationIsNull, ObservabilityQueryNewParamsParametersFiltersOperationIn, ObservabilityQueryNewParamsParametersFiltersOperationNotIn, ObservabilityQueryNewParamsParametersFiltersOperationEq, ObservabilityQueryNewParamsParametersFiltersOperationNeq, ObservabilityQueryNewParamsParametersFiltersOperationGt, ObservabilityQueryNewParamsParametersFiltersOperationGte, ObservabilityQueryNewParamsParametersFiltersOperationLt, ObservabilityQueryNewParamsParametersFiltersOperationLte, ObservabilityQueryNewParamsParametersFiltersOperationEquals, ObservabilityQueryNewParamsParametersFiltersOperationNotEquals, ObservabilityQueryNewParamsParametersFiltersOperationGreater, ObservabilityQueryNewParamsParametersFiltersOperationGreaterOrEquals, ObservabilityQueryNewParamsParametersFiltersOperationLess, ObservabilityQueryNewParamsParametersFiltersOperationLessOrEquals, ObservabilityQueryNewParamsParametersFiltersOperationIncludesUppercase, ObservabilityQueryNewParamsParametersFiltersOperationDoesNotInclude, ObservabilityQueryNewParamsParametersFiltersOperationMatchRegex, ObservabilityQueryNewParamsParametersFiltersOperationExistsUppercase, ObservabilityQueryNewParamsParametersFiltersOperationDoesNotExist, ObservabilityQueryNewParamsParametersFiltersOperationInUppercase, ObservabilityQueryNewParamsParametersFiltersOperationNotInUppercase, ObservabilityQueryNewParamsParametersFiltersOperationStartsWithUppercase, ObservabilityQueryNewParamsParametersFiltersOperationEndsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityQueryNewParamsParametersFiltersType string

const (
	ObservabilityQueryNewParamsParametersFiltersTypeString  ObservabilityQueryNewParamsParametersFiltersType = "string"
	ObservabilityQueryNewParamsParametersFiltersTypeNumber  ObservabilityQueryNewParamsParametersFiltersType = "number"
	ObservabilityQueryNewParamsParametersFiltersTypeBoolean ObservabilityQueryNewParamsParametersFiltersType = "boolean"
)

func (r ObservabilityQueryNewParamsParametersFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersFiltersTypeString, ObservabilityQueryNewParamsParametersFiltersTypeNumber, ObservabilityQueryNewParamsParametersFiltersTypeBoolean:
		return true
	}
	return false
}

type ObservabilityQueryNewParamsParametersGroupBy struct {
	Type  param.Field[ObservabilityQueryNewParamsParametersGroupBysType] `json:"type" api:"required"`
	Value param.Field[string]                                            `json:"value" api:"required"`
}

func (r ObservabilityQueryNewParamsParametersGroupBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityQueryNewParamsParametersGroupBysType string

const (
	ObservabilityQueryNewParamsParametersGroupBysTypeString  ObservabilityQueryNewParamsParametersGroupBysType = "string"
	ObservabilityQueryNewParamsParametersGroupBysTypeNumber  ObservabilityQueryNewParamsParametersGroupBysType = "number"
	ObservabilityQueryNewParamsParametersGroupBysTypeBoolean ObservabilityQueryNewParamsParametersGroupBysType = "boolean"
)

func (r ObservabilityQueryNewParamsParametersGroupBysType) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersGroupBysTypeString, ObservabilityQueryNewParamsParametersGroupBysTypeNumber, ObservabilityQueryNewParamsParametersGroupBysTypeBoolean:
		return true
	}
	return false
}

type ObservabilityQueryNewParamsParametersHaving struct {
	Key       param.Field[string]                                                `json:"key" api:"required"`
	Operation param.Field[ObservabilityQueryNewParamsParametersHavingsOperation] `json:"operation" api:"required"`
	Value     param.Field[float64]                                               `json:"value" api:"required"`
}

func (r ObservabilityQueryNewParamsParametersHaving) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityQueryNewParamsParametersHavingsOperation string

const (
	ObservabilityQueryNewParamsParametersHavingsOperationEq  ObservabilityQueryNewParamsParametersHavingsOperation = "eq"
	ObservabilityQueryNewParamsParametersHavingsOperationNeq ObservabilityQueryNewParamsParametersHavingsOperation = "neq"
	ObservabilityQueryNewParamsParametersHavingsOperationGt  ObservabilityQueryNewParamsParametersHavingsOperation = "gt"
	ObservabilityQueryNewParamsParametersHavingsOperationGte ObservabilityQueryNewParamsParametersHavingsOperation = "gte"
	ObservabilityQueryNewParamsParametersHavingsOperationLt  ObservabilityQueryNewParamsParametersHavingsOperation = "lt"
	ObservabilityQueryNewParamsParametersHavingsOperationLte ObservabilityQueryNewParamsParametersHavingsOperation = "lte"
)

func (r ObservabilityQueryNewParamsParametersHavingsOperation) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersHavingsOperationEq, ObservabilityQueryNewParamsParametersHavingsOperationNeq, ObservabilityQueryNewParamsParametersHavingsOperationGt, ObservabilityQueryNewParamsParametersHavingsOperationGte, ObservabilityQueryNewParamsParametersHavingsOperationLt, ObservabilityQueryNewParamsParametersHavingsOperationLte:
		return true
	}
	return false
}

// Define an expression to search using full-text search.
type ObservabilityQueryNewParamsParametersNeedle struct {
	Value     param.Field[ObservabilityQueryNewParamsParametersNeedleValueUnion] `json:"value" api:"required"`
	IsRegex   param.Field[bool]                                                  `json:"isRegex"`
	MatchCase param.Field[bool]                                                  `json:"matchCase"`
}

func (r ObservabilityQueryNewParamsParametersNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityQueryNewParamsParametersNeedleValueUnion interface {
	ImplementsObservabilityQueryNewParamsParametersNeedleValueUnion()
}

// Configure the order of the results returned by the query.
type ObservabilityQueryNewParamsParametersOrderBy struct {
	// Configure which Calculation to order the results by.
	Value param.Field[string] `json:"value" api:"required"`
	// Set the order of the results
	Order param.Field[ObservabilityQueryNewParamsParametersOrderByOrder] `json:"order"`
}

func (r ObservabilityQueryNewParamsParametersOrderBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set the order of the results
type ObservabilityQueryNewParamsParametersOrderByOrder string

const (
	ObservabilityQueryNewParamsParametersOrderByOrderAsc  ObservabilityQueryNewParamsParametersOrderByOrder = "asc"
	ObservabilityQueryNewParamsParametersOrderByOrderDesc ObservabilityQueryNewParamsParametersOrderByOrder = "desc"
)

func (r ObservabilityQueryNewParamsParametersOrderByOrder) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewParamsParametersOrderByOrderAsc, ObservabilityQueryNewParamsParametersOrderByOrderDesc:
		return true
	}
	return false
}

type ObservabilityQueryNewResponseEnvelope struct {
	Errors   []ObservabilityQueryNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityQueryNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityQueryNewResponse                   `json:"result" api:"required"`
	Success  ObservabilityQueryNewResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityQueryNewResponseEnvelopeJSON       `json:"-"`
}

// observabilityQueryNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ObservabilityQueryNewResponseEnvelope]
type observabilityQueryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryNewResponseEnvelopeErrors struct {
	Message string                                          `json:"message" api:"required"`
	JSON    observabilityQueryNewResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityQueryNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ObservabilityQueryNewResponseEnvelopeErrors]
type observabilityQueryNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryNewResponseEnvelopeMessages struct {
	Message ObservabilityQueryNewResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityQueryNewResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityQueryNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ObservabilityQueryNewResponseEnvelopeMessages]
type observabilityQueryNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityQueryNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityQueryNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityQueryNewResponseEnvelopeMessagesMessage string

const (
	ObservabilityQueryNewResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityQueryNewResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityQueryNewResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityQueryNewResponseEnvelopeSuccess bool

const (
	ObservabilityQueryNewResponseEnvelopeSuccessTrue ObservabilityQueryNewResponseEnvelopeSuccess = true
)

func (r ObservabilityQueryNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityQueryNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityQueryListParams struct {
	AccountID param.Field[string]                              `path:"account_id" api:"required"`
	Order     param.Field[ObservabilityQueryListParamsOrder]   `query:"order"`
	OrderBy   param.Field[ObservabilityQueryListParamsOrderBy] `query:"orderBy"`
	Page      param.Field[float64]                             `query:"page"`
	PerPage   param.Field[float64]                             `query:"perPage"`
}

// URLQuery serializes [ObservabilityQueryListParams]'s query parameters as
// `url.Values`.
func (r ObservabilityQueryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ObservabilityQueryListParamsOrder string

const (
	ObservabilityQueryListParamsOrderAsc  ObservabilityQueryListParamsOrder = "asc"
	ObservabilityQueryListParamsOrderDesc ObservabilityQueryListParamsOrder = "desc"
)

func (r ObservabilityQueryListParamsOrder) IsKnown() bool {
	switch r {
	case ObservabilityQueryListParamsOrderAsc, ObservabilityQueryListParamsOrderDesc:
		return true
	}
	return false
}

type ObservabilityQueryListParamsOrderBy string

const (
	ObservabilityQueryListParamsOrderByCreated ObservabilityQueryListParamsOrderBy = "created"
	ObservabilityQueryListParamsOrderByUpdated ObservabilityQueryListParamsOrderBy = "updated"
)

func (r ObservabilityQueryListParamsOrderBy) IsKnown() bool {
	switch r {
	case ObservabilityQueryListParamsOrderByCreated, ObservabilityQueryListParamsOrderByUpdated:
		return true
	}
	return false
}
