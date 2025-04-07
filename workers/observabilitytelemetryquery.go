// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/tidwall/gjson"
)

// ObservabilityTelemetryQueryService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityTelemetryQueryService] method instead.
type ObservabilityTelemetryQueryService struct {
	Options []option.RequestOption
}

// NewObservabilityTelemetryQueryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservabilityTelemetryQueryService(opts ...option.RequestOption) (r *ObservabilityTelemetryQueryService) {
	r = &ObservabilityTelemetryQueryService{}
	r.Options = opts
	return
}

// Runs a temporary or saved query
func (r *ObservabilityTelemetryQueryService) New(ctx context.Context, params ObservabilityTelemetryQueryNewParams, opts ...option.RequestOption) (res *ObservabilityTelemetryQueryNewResponse, err error) {
	var env ObservabilityTelemetryQueryNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/telemetry/query", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ObservabilityTelemetryQueryNewResponse struct {
	// A Workers Observability Query Object
	Run ObservabilityTelemetryQueryNewResponseRun `json:"run,required"`
	// The statistics object contains information about query performance from the
	// database, it does not include any network latency
	Statistics   ObservabilityTelemetryQueryNewResponseStatistics              `json:"statistics,required"`
	Calculations []ObservabilityTelemetryQueryNewResponseCalculation           `json:"calculations"`
	Compare      []ObservabilityTelemetryQueryNewResponseCompare               `json:"compare"`
	Events       ObservabilityTelemetryQueryNewResponseEvents                  `json:"events"`
	Invocations  map[string][]ObservabilityTelemetryQueryNewResponseInvocation `json:"invocations"`
	Patterns     []ObservabilityTelemetryQueryNewResponsePattern               `json:"patterns"`
	JSON         observabilityTelemetryQueryNewResponseJSON                    `json:"-"`
}

// observabilityTelemetryQueryNewResponseJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryQueryNewResponse]
type observabilityTelemetryQueryNewResponseJSON struct {
	Run          apijson.Field
	Statistics   apijson.Field
	Calculations apijson.Field
	Compare      apijson.Field
	Events       apijson.Field
	Invocations  apijson.Field
	Patterns     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseJSON) RawJSON() string {
	return r.raw
}

// A Workers Observability Query Object
type ObservabilityTelemetryQueryNewResponseRun struct {
	ID        string `json:"id,required"`
	AccountID string `json:"accountId,required"`
	Dry       bool   `json:"dry,required"`
	// Deprecated: deprecated
	EnvironmentID string                                             `json:"environmentId,required"`
	Granularity   float64                                            `json:"granularity,required"`
	Query         ObservabilityTelemetryQueryNewResponseRunQuery     `json:"query,required"`
	Status        ObservabilityTelemetryQueryNewResponseRunStatus    `json:"status,required"`
	Timeframe     ObservabilityTelemetryQueryNewResponseRunTimeframe `json:"timeframe,required"`
	UserID        string                                             `json:"userId,required"`
	// Deprecated: deprecated
	WorkspaceID string                                              `json:"workspaceId,required"`
	Created     string                                              `json:"created"`
	Statistics  ObservabilityTelemetryQueryNewResponseRunStatistics `json:"statistics"`
	Updated     string                                              `json:"updated"`
	JSON        observabilityTelemetryQueryNewResponseRunJSON       `json:"-"`
}

// observabilityTelemetryQueryNewResponseRunJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryQueryNewResponseRun]
type observabilityTelemetryQueryNewResponseRunJSON struct {
	ID            apijson.Field
	AccountID     apijson.Field
	Dry           apijson.Field
	EnvironmentID apijson.Field
	Granularity   apijson.Field
	Query         apijson.Field
	Status        apijson.Field
	Timeframe     apijson.Field
	UserID        apijson.Field
	WorkspaceID   apijson.Field
	Created       apijson.Field
	Statistics    apijson.Field
	Updated       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseRun) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseRunJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseRunQuery struct {
	// ID of the query
	ID          string `json:"id,required"`
	Created     string `json:"created,required"`
	Description string `json:"description,required,nullable"`
	// ID of your environment
	EnvironmentID string `json:"environmentId,required"`
	// Flag for alerts automatically created
	Generated bool `json:"generated,required,nullable"`
	// Query name
	Name       string                                                   `json:"name,required,nullable"`
	Parameters ObservabilityTelemetryQueryNewResponseRunQueryParameters `json:"parameters,required"`
	Updated    string                                                   `json:"updated,required"`
	UserID     string                                                   `json:"userId,required"`
	// ID of your workspace
	WorkspaceID string                                             `json:"workspaceId,required"`
	JSON        observabilityTelemetryQueryNewResponseRunQueryJSON `json:"-"`
}

// observabilityTelemetryQueryNewResponseRunQueryJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryNewResponseRunQuery]
type observabilityTelemetryQueryNewResponseRunQueryJSON struct {
	ID            apijson.Field
	Created       apijson.Field
	Description   apijson.Field
	EnvironmentID apijson.Field
	Generated     apijson.Field
	Name          apijson.Field
	Parameters    apijson.Field
	Updated       apijson.Field
	UserID        apijson.Field
	WorkspaceID   apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseRunQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseRunQueryJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseRunQueryParameters struct {
	// Create Calculations to compute as part of the query.
	Calculations []ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculation `json:"calculations"`
	// Set the Datasets to query. Leave it empty to query all the datasets.
	Datasets []string `json:"datasets"`
	// Set a Flag to describe how to combine the filters on the query.
	FilterCombination ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombination `json:"filterCombination"`
	// Configure the Filters to apply to the query.
	Filters []ObservabilityTelemetryQueryNewResponseRunQueryParametersFilter `json:"filters"`
	// Define how to group the results of the query.
	GroupBys []ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBy `json:"groupBys"`
	// Configure the Having clauses that filter on calculations in the query result.
	Havings []ObservabilityTelemetryQueryNewResponseRunQueryParametersHaving `json:"havings"`
	// Set a limit on the number of results / records returned by the query
	Limit int64 `json:"limit"`
	// Define an expression to search using full-text search.
	Needle ObservabilityTelemetryQueryNewResponseRunQueryParametersNeedle `json:"needle"`
	// Configure the order of the results returned by the query.
	OrderBy ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderBy `json:"orderBy"`
	JSON    observabilityTelemetryQueryNewResponseRunQueryParametersJSON    `json:"-"`
}

// observabilityTelemetryQueryNewResponseRunQueryParametersJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryNewResponseRunQueryParameters]
type observabilityTelemetryQueryNewResponseRunQueryParametersJSON struct {
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

func (r *ObservabilityTelemetryQueryNewResponseRunQueryParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseRunQueryParametersJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculation struct {
	Operator ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator `json:"operator,required"`
	Alias    string                                                                       `json:"alias"`
	Key      string                                                                       `json:"key"`
	KeyType  ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsKeyType  `json:"keyType"`
	JSON     observabilityTelemetryQueryNewResponseRunQueryParametersCalculationJSON      `json:"-"`
}

// observabilityTelemetryQueryNewResponseRunQueryParametersCalculationJSON contains
// the JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculation]
type observabilityTelemetryQueryNewResponseRunQueryParametersCalculationJSON struct {
	Operator    apijson.Field
	Alias       apijson.Field
	Key         apijson.Field
	KeyType     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseRunQueryParametersCalculationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator string

const (
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorUniq              ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "uniq"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorCount             ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "count"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorMax               ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "max"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorMin               ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "min"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorSum               ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "sum"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorAvg               ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "avg"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorMedian            ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "median"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP001              ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "p001"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP01               ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "p01"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP05               ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "p05"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP10               ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "p10"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP25               ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "p25"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP75               ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "p75"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP90               ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "p90"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP95               ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "p95"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP99               ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "p99"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP999              ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "p999"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorStddev            ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "stddev"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorVariance          ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "variance"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorCountDistinct     ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "COUNT_DISTINCT"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorCountUppercase    ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "COUNT"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorMaxUppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "MAX"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorMinUppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "MIN"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorSumUppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "SUM"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorAvgUppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "AVG"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorMedianUppercase   ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "MEDIAN"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP001Uppercase     ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "P001"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP01Uppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "P01"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP05Uppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "P05"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP10Uppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "P10"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP25Uppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "P25"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP75Uppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "P75"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP90Uppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "P90"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP95Uppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "P95"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP99Uppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "P99"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP999Uppercase     ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "P999"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorStddevUppercase   ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "STDDEV"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorVarianceUppercase ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator = "VARIANCE"
)

func (r ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperator) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorUniq, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorCount, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorMax, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorMin, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorSum, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorAvg, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorMedian, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP001, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP01, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP05, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP10, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP25, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP75, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP90, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP95, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP99, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP999, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorStddev, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorVariance, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorCountDistinct, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorCountUppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorMaxUppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorMinUppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorSumUppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorAvgUppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorMedianUppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP001Uppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP01Uppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP05Uppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP10Uppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP25Uppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP75Uppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP90Uppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP95Uppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP99Uppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorP999Uppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorStddevUppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsOperatorVarianceUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsKeyType string

const (
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsKeyTypeString  ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsKeyType = "string"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsKeyTypeNumber  ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsKeyType = "number"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsKeyTypeBoolean ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsKeyType = "boolean"
)

func (r ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsKeyType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsKeyTypeString, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsKeyTypeNumber, ObservabilityTelemetryQueryNewResponseRunQueryParametersCalculationsKeyTypeBoolean:
		return true
	}
	return false
}

// Set a Flag to describe how to combine the filters on the query.
type ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombination string

const (
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombinationAnd          ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombination = "and"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombinationOr           ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombination = "or"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombinationAndUppercase ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombination = "AND"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombinationOrUppercase  ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombination = "OR"
)

func (r ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombinationAnd, ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombinationOr, ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombinationAndUppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseRunQueryParametersFilter struct {
	Key       string                                                                    `json:"key,required"`
	Operation ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation  `json:"operation,required"`
	Type      ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersType       `json:"type,required"`
	Value     ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersValueUnion `json:"value"`
	JSON      observabilityTelemetryQueryNewResponseRunQueryParametersFilterJSON        `json:"-"`
}

// observabilityTelemetryQueryNewResponseRunQueryParametersFilterJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseRunQueryParametersFilter]
type observabilityTelemetryQueryNewResponseRunQueryParametersFilterJSON struct {
	Key         apijson.Field
	Operation   apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseRunQueryParametersFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseRunQueryParametersFilterJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation string

const (
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationIncludes            ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "includes"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationNotIncludes         ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "not_includes"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationStartsWith          ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "starts_with"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationRegex               ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "regex"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationExists              ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "exists"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationIsNull              ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "is_null"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationIn                  ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "in"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationNotIn               ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "not_in"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationEq                  ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "eq"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationNeq                 ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "neq"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationGt                  ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "gt"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationGte                 ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "gte"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationLt                  ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "lt"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationLte                 ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "lte"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationUnknown3            ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "="
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationNotEquals           ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "!="
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationGreater             ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = ">"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationGreaterOrEquals     ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = ">="
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationLess                ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "<"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationLessOrEquals        ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "<="
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationIncludesUppercase   ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "INCLUDES"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationDoesNotInclude      ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationMatchRegex          ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationExistsUppercase     ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "EXISTS"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationDoesNotExist        ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationInUppercase         ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "IN"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationNotInUppercase      ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "NOT_IN"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationStartsWithUppercase ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationIncludes, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationNotIncludes, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationStartsWith, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationRegex, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationExists, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationIsNull, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationIn, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationNotIn, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationEq, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationNeq, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationGt, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationGte, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationLt, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationLte, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationUnknown3, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationNotEquals, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationGreater, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationGreaterOrEquals, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationLess, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationLessOrEquals, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationIncludesUppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationDoesNotInclude, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationMatchRegex, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationExistsUppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationDoesNotExist, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationInUppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationNotInUppercase, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersType string

const (
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersTypeString  ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersType = "string"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersTypeNumber  ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersType = "number"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersTypeBoolean ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersType = "boolean"
)

func (r ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersTypeString, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersTypeNumber, ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersTypeBoolean:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersValueUnion interface {
	ImplementsObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseRunQueryParametersFiltersValueUnion)(nil)).Elem(),
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

type ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBy struct {
	Type  ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBysType `json:"type,required"`
	Value string                                                               `json:"value,required"`
	JSON  observabilityTelemetryQueryNewResponseRunQueryParametersGroupByJSON  `json:"-"`
}

// observabilityTelemetryQueryNewResponseRunQueryParametersGroupByJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBy]
type observabilityTelemetryQueryNewResponseRunQueryParametersGroupByJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseRunQueryParametersGroupByJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBysType string

const (
	ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBysTypeString  ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBysType = "string"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBysTypeNumber  ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBysType = "number"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBysTypeBoolean ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBysType = "boolean"
)

func (r ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBysType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBysTypeString, ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBysTypeNumber, ObservabilityTelemetryQueryNewResponseRunQueryParametersGroupBysTypeBoolean:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseRunQueryParametersHaving struct {
	Key       string                                                                   `json:"key,required"`
	Operation ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperation `json:"operation,required"`
	Value     float64                                                                  `json:"value,required"`
	JSON      observabilityTelemetryQueryNewResponseRunQueryParametersHavingJSON       `json:"-"`
}

// observabilityTelemetryQueryNewResponseRunQueryParametersHavingJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseRunQueryParametersHaving]
type observabilityTelemetryQueryNewResponseRunQueryParametersHavingJSON struct {
	Key         apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseRunQueryParametersHaving) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseRunQueryParametersHavingJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperation string

const (
	ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperationEq  ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperation = "eq"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperationNeq ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperation = "neq"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperationGt  ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperation = "gt"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperationGte ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperation = "gte"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperationLt  ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperation = "lt"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperationLte ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperation = "lte"
)

func (r ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperationEq, ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperationNeq, ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperationGt, ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperationGte, ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperationLt, ObservabilityTelemetryQueryNewResponseRunQueryParametersHavingsOperationLte:
		return true
	}
	return false
}

// Define an expression to search using full-text search.
type ObservabilityTelemetryQueryNewResponseRunQueryParametersNeedle struct {
	Value     ObservabilityTelemetryQueryNewResponseRunQueryParametersNeedleValueUnion `json:"value,required"`
	IsRegex   bool                                                                     `json:"isRegex"`
	MatchCase bool                                                                     `json:"matchCase"`
	JSON      observabilityTelemetryQueryNewResponseRunQueryParametersNeedleJSON       `json:"-"`
}

// observabilityTelemetryQueryNewResponseRunQueryParametersNeedleJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseRunQueryParametersNeedle]
type observabilityTelemetryQueryNewResponseRunQueryParametersNeedleJSON struct {
	Value       apijson.Field
	IsRegex     apijson.Field
	MatchCase   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseRunQueryParametersNeedle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseRunQueryParametersNeedleJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryNewResponseRunQueryParametersNeedleValueUnion interface {
	ImplementsObservabilityTelemetryQueryNewResponseRunQueryParametersNeedleValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseRunQueryParametersNeedleValueUnion)(nil)).Elem(),
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
type ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderBy struct {
	// Configure which Calculation to order the results by.
	Value string `json:"value,required"`
	// Set the order of the results
	Order ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderByOrder `json:"order"`
	JSON  observabilityTelemetryQueryNewResponseRunQueryParametersOrderByJSON  `json:"-"`
}

// observabilityTelemetryQueryNewResponseRunQueryParametersOrderByJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderBy]
type observabilityTelemetryQueryNewResponseRunQueryParametersOrderByJSON struct {
	Value       apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseRunQueryParametersOrderByJSON) RawJSON() string {
	return r.raw
}

// Set the order of the results
type ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderByOrder string

const (
	ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderByOrderAsc  ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderByOrder = "asc"
	ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderByOrderDesc ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderByOrder = "desc"
)

func (r ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderByOrder) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderByOrderAsc, ObservabilityTelemetryQueryNewResponseRunQueryParametersOrderByOrderDesc:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseRunStatus string

const (
	ObservabilityTelemetryQueryNewResponseRunStatusStarted   ObservabilityTelemetryQueryNewResponseRunStatus = "STARTED"
	ObservabilityTelemetryQueryNewResponseRunStatusCompleted ObservabilityTelemetryQueryNewResponseRunStatus = "COMPLETED"
)

func (r ObservabilityTelemetryQueryNewResponseRunStatus) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseRunStatusStarted, ObservabilityTelemetryQueryNewResponseRunStatusCompleted:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseRunTimeframe struct {
	// Set the start time for your query using UNIX time in milliseconds.
	From float64 `json:"from,required"`
	// Set the end time for your query using UNIX time in milliseconds.
	To   float64                                                `json:"to,required"`
	JSON observabilityTelemetryQueryNewResponseRunTimeframeJSON `json:"-"`
}

// observabilityTelemetryQueryNewResponseRunTimeframeJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryNewResponseRunTimeframe]
type observabilityTelemetryQueryNewResponseRunTimeframeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseRunTimeframe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseRunTimeframeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseRunStatistics struct {
	// Number of uncompressed bytes read from the table.
	BytesRead float64 `json:"bytes_read,required"`
	// Time in seconds for the query to run.
	Elapsed float64 `json:"elapsed,required"`
	// Number of rows scanned from the table.
	RowsRead float64                                                 `json:"rows_read,required"`
	JSON     observabilityTelemetryQueryNewResponseRunStatisticsJSON `json:"-"`
}

// observabilityTelemetryQueryNewResponseRunStatisticsJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryNewResponseRunStatistics]
type observabilityTelemetryQueryNewResponseRunStatisticsJSON struct {
	BytesRead   apijson.Field
	Elapsed     apijson.Field
	RowsRead    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseRunStatistics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseRunStatisticsJSON) RawJSON() string {
	return r.raw
}

// The statistics object contains information about query performance from the
// database, it does not include any network latency
type ObservabilityTelemetryQueryNewResponseStatistics struct {
	// Number of uncompressed bytes read from the table.
	BytesRead float64 `json:"bytes_read,required"`
	// Time in seconds for the query to run.
	Elapsed float64 `json:"elapsed,required"`
	// Number of rows scanned from the table.
	RowsRead float64                                              `json:"rows_read,required"`
	JSON     observabilityTelemetryQueryNewResponseStatisticsJSON `json:"-"`
}

// observabilityTelemetryQueryNewResponseStatisticsJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryNewResponseStatistics]
type observabilityTelemetryQueryNewResponseStatisticsJSON struct {
	BytesRead   apijson.Field
	Elapsed     apijson.Field
	RowsRead    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseStatistics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseStatisticsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseCalculation struct {
	Aggregates  []ObservabilityTelemetryQueryNewResponseCalculationsAggregate `json:"aggregates,required"`
	Calculation string                                                        `json:"calculation,required"`
	Series      []ObservabilityTelemetryQueryNewResponseCalculationsSery      `json:"series,required"`
	Alias       string                                                        `json:"alias"`
	JSON        observabilityTelemetryQueryNewResponseCalculationJSON         `json:"-"`
}

// observabilityTelemetryQueryNewResponseCalculationJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryNewResponseCalculation]
type observabilityTelemetryQueryNewResponseCalculationJSON struct {
	Aggregates  apijson.Field
	Calculation apijson.Field
	Series      apijson.Field
	Alias       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseCalculation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseCalculationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseCalculationsAggregate struct {
	Count          float64                                                             `json:"count,required"`
	Interval       float64                                                             `json:"interval,required"`
	SampleInterval float64                                                             `json:"sampleInterval,required"`
	Value          float64                                                             `json:"value,required"`
	Groups         []ObservabilityTelemetryQueryNewResponseCalculationsAggregatesGroup `json:"groups"`
	JSON           observabilityTelemetryQueryNewResponseCalculationsAggregateJSON     `json:"-"`
}

// observabilityTelemetryQueryNewResponseCalculationsAggregateJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseCalculationsAggregate]
type observabilityTelemetryQueryNewResponseCalculationsAggregateJSON struct {
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseCalculationsAggregate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseCalculationsAggregateJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseCalculationsAggregatesGroup struct {
	Key   string                                                                       `json:"key,required"`
	Value ObservabilityTelemetryQueryNewResponseCalculationsAggregatesGroupsValueUnion `json:"value,required"`
	JSON  observabilityTelemetryQueryNewResponseCalculationsAggregatesGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryNewResponseCalculationsAggregatesGroupJSON contains
// the JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseCalculationsAggregatesGroup]
type observabilityTelemetryQueryNewResponseCalculationsAggregatesGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseCalculationsAggregatesGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseCalculationsAggregatesGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryNewResponseCalculationsAggregatesGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryNewResponseCalculationsAggregatesGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseCalculationsAggregatesGroupsValueUnion)(nil)).Elem(),
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

type ObservabilityTelemetryQueryNewResponseCalculationsSery struct {
	Data []ObservabilityTelemetryQueryNewResponseCalculationsSeriesData `json:"data,required"`
	Time string                                                         `json:"time,required"`
	JSON observabilityTelemetryQueryNewResponseCalculationsSeryJSON     `json:"-"`
}

// observabilityTelemetryQueryNewResponseCalculationsSeryJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryNewResponseCalculationsSery]
type observabilityTelemetryQueryNewResponseCalculationsSeryJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseCalculationsSery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseCalculationsSeryJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseCalculationsSeriesData struct {
	Count          float64                                                             `json:"count,required"`
	FirstSeen      string                                                              `json:"firstSeen,required"`
	Interval       float64                                                             `json:"interval,required"`
	LastSeen       string                                                              `json:"lastSeen,required"`
	SampleInterval float64                                                             `json:"sampleInterval,required"`
	Value          float64                                                             `json:"value,required"`
	Groups         []ObservabilityTelemetryQueryNewResponseCalculationsSeriesDataGroup `json:"groups"`
	JSON           observabilityTelemetryQueryNewResponseCalculationsSeriesDataJSON    `json:"-"`
}

// observabilityTelemetryQueryNewResponseCalculationsSeriesDataJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseCalculationsSeriesData]
type observabilityTelemetryQueryNewResponseCalculationsSeriesDataJSON struct {
	Count          apijson.Field
	FirstSeen      apijson.Field
	Interval       apijson.Field
	LastSeen       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseCalculationsSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseCalculationsSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseCalculationsSeriesDataGroup struct {
	Key   string                                                                       `json:"key,required"`
	Value ObservabilityTelemetryQueryNewResponseCalculationsSeriesDataGroupsValueUnion `json:"value,required"`
	JSON  observabilityTelemetryQueryNewResponseCalculationsSeriesDataGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryNewResponseCalculationsSeriesDataGroupJSON contains
// the JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseCalculationsSeriesDataGroup]
type observabilityTelemetryQueryNewResponseCalculationsSeriesDataGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseCalculationsSeriesDataGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseCalculationsSeriesDataGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryNewResponseCalculationsSeriesDataGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryNewResponseCalculationsSeriesDataGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseCalculationsSeriesDataGroupsValueUnion)(nil)).Elem(),
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

type ObservabilityTelemetryQueryNewResponseCompare struct {
	Aggregates  []ObservabilityTelemetryQueryNewResponseCompareAggregate `json:"aggregates,required"`
	Calculation string                                                   `json:"calculation,required"`
	Series      []ObservabilityTelemetryQueryNewResponseCompareSery      `json:"series,required"`
	Alias       string                                                   `json:"alias"`
	JSON        observabilityTelemetryQueryNewResponseCompareJSON        `json:"-"`
}

// observabilityTelemetryQueryNewResponseCompareJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryNewResponseCompare]
type observabilityTelemetryQueryNewResponseCompareJSON struct {
	Aggregates  apijson.Field
	Calculation apijson.Field
	Series      apijson.Field
	Alias       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseCompare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseCompareJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseCompareAggregate struct {
	Count          float64                                                        `json:"count,required"`
	Interval       float64                                                        `json:"interval,required"`
	SampleInterval float64                                                        `json:"sampleInterval,required"`
	Value          float64                                                        `json:"value,required"`
	Groups         []ObservabilityTelemetryQueryNewResponseCompareAggregatesGroup `json:"groups"`
	JSON           observabilityTelemetryQueryNewResponseCompareAggregateJSON     `json:"-"`
}

// observabilityTelemetryQueryNewResponseCompareAggregateJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryNewResponseCompareAggregate]
type observabilityTelemetryQueryNewResponseCompareAggregateJSON struct {
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseCompareAggregate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseCompareAggregateJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseCompareAggregatesGroup struct {
	Key   string                                                                  `json:"key,required"`
	Value ObservabilityTelemetryQueryNewResponseCompareAggregatesGroupsValueUnion `json:"value,required"`
	JSON  observabilityTelemetryQueryNewResponseCompareAggregatesGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryNewResponseCompareAggregatesGroupJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseCompareAggregatesGroup]
type observabilityTelemetryQueryNewResponseCompareAggregatesGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseCompareAggregatesGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseCompareAggregatesGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryNewResponseCompareAggregatesGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryNewResponseCompareAggregatesGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseCompareAggregatesGroupsValueUnion)(nil)).Elem(),
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

type ObservabilityTelemetryQueryNewResponseCompareSery struct {
	Data []ObservabilityTelemetryQueryNewResponseCompareSeriesData `json:"data,required"`
	Time string                                                    `json:"time,required"`
	JSON observabilityTelemetryQueryNewResponseCompareSeryJSON     `json:"-"`
}

// observabilityTelemetryQueryNewResponseCompareSeryJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryNewResponseCompareSery]
type observabilityTelemetryQueryNewResponseCompareSeryJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseCompareSery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseCompareSeryJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseCompareSeriesData struct {
	Count          float64                                                        `json:"count,required"`
	FirstSeen      string                                                         `json:"firstSeen,required"`
	Interval       float64                                                        `json:"interval,required"`
	LastSeen       string                                                         `json:"lastSeen,required"`
	SampleInterval float64                                                        `json:"sampleInterval,required"`
	Value          float64                                                        `json:"value,required"`
	Groups         []ObservabilityTelemetryQueryNewResponseCompareSeriesDataGroup `json:"groups"`
	JSON           observabilityTelemetryQueryNewResponseCompareSeriesDataJSON    `json:"-"`
}

// observabilityTelemetryQueryNewResponseCompareSeriesDataJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryNewResponseCompareSeriesData]
type observabilityTelemetryQueryNewResponseCompareSeriesDataJSON struct {
	Count          apijson.Field
	FirstSeen      apijson.Field
	Interval       apijson.Field
	LastSeen       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseCompareSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseCompareSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseCompareSeriesDataGroup struct {
	Key   string                                                                  `json:"key,required"`
	Value ObservabilityTelemetryQueryNewResponseCompareSeriesDataGroupsValueUnion `json:"value,required"`
	JSON  observabilityTelemetryQueryNewResponseCompareSeriesDataGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryNewResponseCompareSeriesDataGroupJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseCompareSeriesDataGroup]
type observabilityTelemetryQueryNewResponseCompareSeriesDataGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseCompareSeriesDataGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseCompareSeriesDataGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryNewResponseCompareSeriesDataGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryNewResponseCompareSeriesDataGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseCompareSeriesDataGroupsValueUnion)(nil)).Elem(),
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

type ObservabilityTelemetryQueryNewResponseEvents struct {
	Count  float64                                             `json:"count"`
	Events []ObservabilityTelemetryQueryNewResponseEventsEvent `json:"events"`
	Fields []ObservabilityTelemetryQueryNewResponseEventsField `json:"fields"`
	Series []ObservabilityTelemetryQueryNewResponseEventsSery  `json:"series"`
	JSON   observabilityTelemetryQueryNewResponseEventsJSON    `json:"-"`
}

// observabilityTelemetryQueryNewResponseEventsJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryNewResponseEvents]
type observabilityTelemetryQueryNewResponseEventsJSON struct {
	Count       apijson.Field
	Events      apijson.Field
	Fields      apijson.Field
	Series      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseEvents) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseEventsJSON) RawJSON() string {
	return r.raw
}

// The data structure of a telemetry event
type ObservabilityTelemetryQueryNewResponseEventsEvent struct {
	Metadata  ObservabilityTelemetryQueryNewResponseEventsEventsMetadata `json:"$metadata,required"`
	Dataset   string                                                     `json:"dataset,required"`
	Source    interface{}                                                `json:"source,required"`
	Timestamp int64                                                      `json:"timestamp,required"`
	// Cloudflare Workers event information enriches your logs so you can easily
	// identify and debug issues.
	Workers ObservabilityTelemetryQueryNewResponseEventsEventsWorkers `json:"$workers"`
	JSON    observabilityTelemetryQueryNewResponseEventsEventJSON     `json:"-"`
}

// observabilityTelemetryQueryNewResponseEventsEventJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryNewResponseEventsEvent]
type observabilityTelemetryQueryNewResponseEventsEventJSON struct {
	Metadata    apijson.Field
	Dataset     apijson.Field
	Source      apijson.Field
	Timestamp   apijson.Field
	Workers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseEventsEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseEventsEventJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseEventsEventsMetadata struct {
	ID              string                                                         `json:"id,required"`
	Account         string                                                         `json:"account"`
	CloudService    string                                                         `json:"cloudService"`
	ColdStart       int64                                                          `json:"coldStart"`
	Cost            int64                                                          `json:"cost"`
	Duration        int64                                                          `json:"duration"`
	EndTime         int64                                                          `json:"endTime"`
	Error           string                                                         `json:"error"`
	ErrorTemplate   string                                                         `json:"errorTemplate"`
	Fingerprint     string                                                         `json:"fingerprint"`
	Level           string                                                         `json:"level"`
	Message         string                                                         `json:"message"`
	MessageTemplate string                                                         `json:"messageTemplate"`
	MetricName      string                                                         `json:"metricName"`
	Origin          string                                                         `json:"origin"`
	ParentSpanID    string                                                         `json:"parentSpanId"`
	Provider        string                                                         `json:"provider"`
	Region          string                                                         `json:"region"`
	RequestID       string                                                         `json:"requestId"`
	Service         string                                                         `json:"service"`
	SpanID          string                                                         `json:"spanId"`
	SpanName        string                                                         `json:"spanName"`
	StackID         string                                                         `json:"stackId"`
	StartTime       int64                                                          `json:"startTime"`
	StatusCode      int64                                                          `json:"statusCode"`
	TraceDuration   int64                                                          `json:"traceDuration"`
	TraceID         string                                                         `json:"traceId"`
	Trigger         string                                                         `json:"trigger"`
	Type            string                                                         `json:"type"`
	URL             string                                                         `json:"url"`
	JSON            observabilityTelemetryQueryNewResponseEventsEventsMetadataJSON `json:"-"`
}

// observabilityTelemetryQueryNewResponseEventsEventsMetadataJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryNewResponseEventsEventsMetadata]
type observabilityTelemetryQueryNewResponseEventsEventsMetadataJSON struct {
	ID              apijson.Field
	Account         apijson.Field
	CloudService    apijson.Field
	ColdStart       apijson.Field
	Cost            apijson.Field
	Duration        apijson.Field
	EndTime         apijson.Field
	Error           apijson.Field
	ErrorTemplate   apijson.Field
	Fingerprint     apijson.Field
	Level           apijson.Field
	Message         apijson.Field
	MessageTemplate apijson.Field
	MetricName      apijson.Field
	Origin          apijson.Field
	ParentSpanID    apijson.Field
	Provider        apijson.Field
	Region          apijson.Field
	RequestID       apijson.Field
	Service         apijson.Field
	SpanID          apijson.Field
	SpanName        apijson.Field
	StackID         apijson.Field
	StartTime       apijson.Field
	StatusCode      apijson.Field
	TraceDuration   apijson.Field
	TraceID         apijson.Field
	Trigger         apijson.Field
	Type            apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseEventsEventsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseEventsEventsMetadataJSON) RawJSON() string {
	return r.raw
}

// Cloudflare Workers event information enriches your logs so you can easily
// identify and debug issues.
type ObservabilityTelemetryQueryNewResponseEventsEventsWorkers struct {
	EventType  ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType `json:"eventType,required"`
	Outcome    string                                                             `json:"outcome,required"`
	RequestID  string                                                             `json:"requestId,required"`
	ScriptName string                                                             `json:"scriptName,required"`
	CPUTimeMs  float64                                                            `json:"cpuTimeMs"`
	// This field can have the runtime type of
	// [[]ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectDiagnosticsChannelEvent].
	DiagnosticsChannelEvents interface{} `json:"diagnosticsChannelEvents"`
	DispatchNamespace        string      `json:"dispatchNamespace"`
	Entrypoint               string      `json:"entrypoint"`
	// This field can have the runtime type of
	// [map[string]ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventUnion].
	Event          interface{}                                                             `json:"event"`
	ExecutionModel ObservabilityTelemetryQueryNewResponseEventsEventsWorkersExecutionModel `json:"executionModel"`
	// This field can have the runtime type of
	// [ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectScriptVersion].
	ScriptVersion interface{}                                                   `json:"scriptVersion"`
	Truncated     bool                                                          `json:"truncated"`
	WallTimeMs    float64                                                       `json:"wallTimeMs"`
	JSON          observabilityTelemetryQueryNewResponseEventsEventsWorkersJSON `json:"-"`
	union         ObservabilityTelemetryQueryNewResponseEventsEventsWorkersUnion
}

// observabilityTelemetryQueryNewResponseEventsEventsWorkersJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryNewResponseEventsEventsWorkers]
type observabilityTelemetryQueryNewResponseEventsEventsWorkersJSON struct {
	EventType                apijson.Field
	Outcome                  apijson.Field
	RequestID                apijson.Field
	ScriptName               apijson.Field
	CPUTimeMs                apijson.Field
	DiagnosticsChannelEvents apijson.Field
	DispatchNamespace        apijson.Field
	Entrypoint               apijson.Field
	Event                    apijson.Field
	ExecutionModel           apijson.Field
	ScriptVersion            apijson.Field
	Truncated                apijson.Field
	WallTimeMs               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r observabilityTelemetryQueryNewResponseEventsEventsWorkersJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilityTelemetryQueryNewResponseEventsEventsWorkers) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilityTelemetryQueryNewResponseEventsEventsWorkers{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ObservabilityTelemetryQueryNewResponseEventsEventsWorkersUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [workers.ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObject],
// [workers.ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObject].
func (r ObservabilityTelemetryQueryNewResponseEventsEventsWorkers) AsUnion() ObservabilityTelemetryQueryNewResponseEventsEventsWorkersUnion {
	return r.union
}

// Cloudflare Workers event information enriches your logs so you can easily
// identify and debug issues.
//
// Union satisfied by
// [workers.ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObject] or
// [workers.ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObject].
type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersUnion interface {
	implementsObservabilityTelemetryQueryNewResponseEventsEventsWorkers()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseEventsEventsWorkersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObject{}),
		},
	)
}

type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObject struct {
	EventType      ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType             `json:"eventType,required"`
	Outcome        string                                                                               `json:"outcome,required"`
	RequestID      string                                                                               `json:"requestId,required"`
	ScriptName     string                                                                               `json:"scriptName,required"`
	Entrypoint     string                                                                               `json:"entrypoint"`
	Event          map[string]ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventUnion `json:"event"`
	ExecutionModel ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectExecutionModel        `json:"executionModel"`
	ScriptVersion  ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectScriptVersion         `json:"scriptVersion"`
	Truncated      bool                                                                                 `json:"truncated"`
	JSON           observabilityTelemetryQueryNewResponseEventsEventsWorkersObjectJSON                  `json:"-"`
}

// observabilityTelemetryQueryNewResponseEventsEventsWorkersObjectJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObject]
type observabilityTelemetryQueryNewResponseEventsEventsWorkersObjectJSON struct {
	EventType      apijson.Field
	Outcome        apijson.Field
	RequestID      apijson.Field
	ScriptName     apijson.Field
	Entrypoint     apijson.Field
	Event          apijson.Field
	ExecutionModel apijson.Field
	ScriptVersion  apijson.Field
	Truncated      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseEventsEventsWorkersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObject) implementsObservabilityTelemetryQueryNewResponseEventsEventsWorkers() {
}

type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType string

const (
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeFetch     ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType = "fetch"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeScheduled ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType = "scheduled"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeAlarm     ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType = "alarm"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeCron      ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType = "cron"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeQueue     ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType = "queue"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeEmail     ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType = "email"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeTail      ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType = "tail"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeRpc       ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType = "rpc"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeWebsocket ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType = "websocket"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeUnknown   ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType = "unknown"
)

func (r ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeFetch, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeScheduled, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeAlarm, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeCron, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeQueue, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeEmail, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeTail, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeRpc, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeWebsocket, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventTypeUnknown:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool]
// or
// [workers.ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMap].
type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventUnion interface {
	ImplementsObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventUnion)(nil)).Elem(),
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

type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMap map[string]ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapUnionItem

func (r ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMap) ImplementsObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventUnion() {
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool]
// or
// [workers.ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMap].
type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapUnionItem interface {
	ImplementsObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapUnionItem)(nil)).Elem(),
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

type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMap map[string]ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapUnionItem

func (r ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMap) ImplementsObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapUnionItem() {
}

// Union satisfied by
// [workers.ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapArray],
// [shared.UnionString], [shared.UnionFloat] or [shared.UnionBool].
type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapUnionItem interface {
	ImplementsObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapArray{}),
		},
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

type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapArray []ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapArrayUnionItem

func (r ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapArray) ImplementsObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapUnionItem() {
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapArrayUnionItem interface {
	ImplementsObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapArrayUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectEventMapMapArrayUnionItem)(nil)).Elem(),
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

type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectExecutionModel string

const (
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectExecutionModelDurableObject ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectExecutionModel = "durableObject"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectExecutionModelStateless     ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectExecutionModel = "stateless"
)

func (r ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectExecutionModelDurableObject, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectScriptVersion struct {
	ID      string                                                                           `json:"id"`
	Message string                                                                           `json:"message"`
	Tag     string                                                                           `json:"tag"`
	JSON    observabilityTelemetryQueryNewResponseEventsEventsWorkersObjectScriptVersionJSON `json:"-"`
}

// observabilityTelemetryQueryNewResponseEventsEventsWorkersObjectScriptVersionJSON
// contains the JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectScriptVersion]
type observabilityTelemetryQueryNewResponseEventsEventsWorkersObjectScriptVersionJSON struct {
	ID          apijson.Field
	Message     apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseEventsEventsWorkersObjectScriptVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseEventsEventsWorkersObjectScriptVersionJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType string

const (
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeFetch     ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType = "fetch"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeScheduled ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType = "scheduled"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeAlarm     ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType = "alarm"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeCron      ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType = "cron"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeQueue     ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType = "queue"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeEmail     ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType = "email"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeTail      ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType = "tail"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeRpc       ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType = "rpc"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeWebsocket ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType = "websocket"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeUnknown   ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType = "unknown"
)

func (r ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeFetch, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeScheduled, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeAlarm, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeCron, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeQueue, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeEmail, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeTail, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeRpc, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeWebsocket, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseEventsEventsWorkersExecutionModel string

const (
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersExecutionModelDurableObject ObservabilityTelemetryQueryNewResponseEventsEventsWorkersExecutionModel = "durableObject"
	ObservabilityTelemetryQueryNewResponseEventsEventsWorkersExecutionModelStateless     ObservabilityTelemetryQueryNewResponseEventsEventsWorkersExecutionModel = "stateless"
)

func (r ObservabilityTelemetryQueryNewResponseEventsEventsWorkersExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseEventsEventsWorkersExecutionModelDurableObject, ObservabilityTelemetryQueryNewResponseEventsEventsWorkersExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseEventsField struct {
	Key  string                                                `json:"key,required"`
	Type string                                                `json:"type,required"`
	JSON observabilityTelemetryQueryNewResponseEventsFieldJSON `json:"-"`
}

// observabilityTelemetryQueryNewResponseEventsFieldJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryNewResponseEventsField]
type observabilityTelemetryQueryNewResponseEventsFieldJSON struct {
	Key         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseEventsField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseEventsFieldJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseEventsSery struct {
	Data []ObservabilityTelemetryQueryNewResponseEventsSeriesData `json:"data,required"`
	Time string                                                   `json:"time,required"`
	JSON observabilityTelemetryQueryNewResponseEventsSeryJSON     `json:"-"`
}

// observabilityTelemetryQueryNewResponseEventsSeryJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryNewResponseEventsSery]
type observabilityTelemetryQueryNewResponseEventsSeryJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseEventsSery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseEventsSeryJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseEventsSeriesData struct {
	Aggregates     ObservabilityTelemetryQueryNewResponseEventsSeriesDataAggregates `json:"aggregates,required"`
	Count          float64                                                          `json:"count,required"`
	Interval       float64                                                          `json:"interval,required"`
	SampleInterval float64                                                          `json:"sampleInterval,required"`
	Errors         float64                                                          `json:"errors"`
	// Groups in the query results.
	Groups map[string]ObservabilityTelemetryQueryNewResponseEventsSeriesDataGroupsUnion `json:"groups"`
	JSON   observabilityTelemetryQueryNewResponseEventsSeriesDataJSON                   `json:"-"`
}

// observabilityTelemetryQueryNewResponseEventsSeriesDataJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryNewResponseEventsSeriesData]
type observabilityTelemetryQueryNewResponseEventsSeriesDataJSON struct {
	Aggregates     apijson.Field
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Errors         apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseEventsSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseEventsSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseEventsSeriesDataAggregates struct {
	// Deprecated: deprecated
	Count int64 `json:"_count,required"`
	// Deprecated: deprecated
	FirstSeen string `json:"_firstSeen,required"`
	// Deprecated: deprecated
	Interval int64 `json:"_interval,required"`
	// Deprecated: deprecated
	LastSeen string `json:"_lastSeen,required"`
	// Deprecated: deprecated
	Bin  interface{}                                                          `json:"bin"`
	JSON observabilityTelemetryQueryNewResponseEventsSeriesDataAggregatesJSON `json:"-"`
}

// observabilityTelemetryQueryNewResponseEventsSeriesDataAggregatesJSON contains
// the JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseEventsSeriesDataAggregates]
type observabilityTelemetryQueryNewResponseEventsSeriesDataAggregatesJSON struct {
	Count       apijson.Field
	FirstSeen   apijson.Field
	Interval    apijson.Field
	LastSeen    apijson.Field
	Bin         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseEventsSeriesDataAggregates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseEventsSeriesDataAggregatesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryNewResponseEventsSeriesDataGroupsUnion interface {
	ImplementsObservabilityTelemetryQueryNewResponseEventsSeriesDataGroupsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseEventsSeriesDataGroupsUnion)(nil)).Elem(),
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

// The data structure of a telemetry event
type ObservabilityTelemetryQueryNewResponseInvocation struct {
	Metadata  ObservabilityTelemetryQueryNewResponseInvocationsMetadata `json:"$metadata,required"`
	Dataset   string                                                    `json:"dataset,required"`
	Source    interface{}                                               `json:"source,required"`
	Timestamp int64                                                     `json:"timestamp,required"`
	// Cloudflare Workers event information enriches your logs so you can easily
	// identify and debug issues.
	Workers ObservabilityTelemetryQueryNewResponseInvocationsWorkers `json:"$workers"`
	JSON    observabilityTelemetryQueryNewResponseInvocationJSON     `json:"-"`
}

// observabilityTelemetryQueryNewResponseInvocationJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryNewResponseInvocation]
type observabilityTelemetryQueryNewResponseInvocationJSON struct {
	Metadata    apijson.Field
	Dataset     apijson.Field
	Source      apijson.Field
	Timestamp   apijson.Field
	Workers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseInvocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseInvocationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseInvocationsMetadata struct {
	ID              string                                                        `json:"id,required"`
	Account         string                                                        `json:"account"`
	CloudService    string                                                        `json:"cloudService"`
	ColdStart       int64                                                         `json:"coldStart"`
	Cost            int64                                                         `json:"cost"`
	Duration        int64                                                         `json:"duration"`
	EndTime         int64                                                         `json:"endTime"`
	Error           string                                                        `json:"error"`
	ErrorTemplate   string                                                        `json:"errorTemplate"`
	Fingerprint     string                                                        `json:"fingerprint"`
	Level           string                                                        `json:"level"`
	Message         string                                                        `json:"message"`
	MessageTemplate string                                                        `json:"messageTemplate"`
	MetricName      string                                                        `json:"metricName"`
	Origin          string                                                        `json:"origin"`
	ParentSpanID    string                                                        `json:"parentSpanId"`
	Provider        string                                                        `json:"provider"`
	Region          string                                                        `json:"region"`
	RequestID       string                                                        `json:"requestId"`
	Service         string                                                        `json:"service"`
	SpanID          string                                                        `json:"spanId"`
	SpanName        string                                                        `json:"spanName"`
	StackID         string                                                        `json:"stackId"`
	StartTime       int64                                                         `json:"startTime"`
	StatusCode      int64                                                         `json:"statusCode"`
	TraceDuration   int64                                                         `json:"traceDuration"`
	TraceID         string                                                        `json:"traceId"`
	Trigger         string                                                        `json:"trigger"`
	Type            string                                                        `json:"type"`
	URL             string                                                        `json:"url"`
	JSON            observabilityTelemetryQueryNewResponseInvocationsMetadataJSON `json:"-"`
}

// observabilityTelemetryQueryNewResponseInvocationsMetadataJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryNewResponseInvocationsMetadata]
type observabilityTelemetryQueryNewResponseInvocationsMetadataJSON struct {
	ID              apijson.Field
	Account         apijson.Field
	CloudService    apijson.Field
	ColdStart       apijson.Field
	Cost            apijson.Field
	Duration        apijson.Field
	EndTime         apijson.Field
	Error           apijson.Field
	ErrorTemplate   apijson.Field
	Fingerprint     apijson.Field
	Level           apijson.Field
	Message         apijson.Field
	MessageTemplate apijson.Field
	MetricName      apijson.Field
	Origin          apijson.Field
	ParentSpanID    apijson.Field
	Provider        apijson.Field
	Region          apijson.Field
	RequestID       apijson.Field
	Service         apijson.Field
	SpanID          apijson.Field
	SpanName        apijson.Field
	StackID         apijson.Field
	StartTime       apijson.Field
	StatusCode      apijson.Field
	TraceDuration   apijson.Field
	TraceID         apijson.Field
	Trigger         apijson.Field
	Type            apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseInvocationsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseInvocationsMetadataJSON) RawJSON() string {
	return r.raw
}

// Cloudflare Workers event information enriches your logs so you can easily
// identify and debug issues.
type ObservabilityTelemetryQueryNewResponseInvocationsWorkers struct {
	EventType  ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType `json:"eventType,required"`
	Outcome    string                                                            `json:"outcome,required"`
	RequestID  string                                                            `json:"requestId,required"`
	ScriptName string                                                            `json:"scriptName,required"`
	CPUTimeMs  float64                                                           `json:"cpuTimeMs"`
	// This field can have the runtime type of
	// [[]ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectDiagnosticsChannelEvent].
	DiagnosticsChannelEvents interface{} `json:"diagnosticsChannelEvents"`
	DispatchNamespace        string      `json:"dispatchNamespace"`
	Entrypoint               string      `json:"entrypoint"`
	// This field can have the runtime type of
	// [map[string]ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventUnion].
	Event          interface{}                                                            `json:"event"`
	ExecutionModel ObservabilityTelemetryQueryNewResponseInvocationsWorkersExecutionModel `json:"executionModel"`
	// This field can have the runtime type of
	// [ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectScriptVersion].
	ScriptVersion interface{}                                                  `json:"scriptVersion"`
	Truncated     bool                                                         `json:"truncated"`
	WallTimeMs    float64                                                      `json:"wallTimeMs"`
	JSON          observabilityTelemetryQueryNewResponseInvocationsWorkersJSON `json:"-"`
	union         ObservabilityTelemetryQueryNewResponseInvocationsWorkersUnion
}

// observabilityTelemetryQueryNewResponseInvocationsWorkersJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryNewResponseInvocationsWorkers]
type observabilityTelemetryQueryNewResponseInvocationsWorkersJSON struct {
	EventType                apijson.Field
	Outcome                  apijson.Field
	RequestID                apijson.Field
	ScriptName               apijson.Field
	CPUTimeMs                apijson.Field
	DiagnosticsChannelEvents apijson.Field
	DispatchNamespace        apijson.Field
	Entrypoint               apijson.Field
	Event                    apijson.Field
	ExecutionModel           apijson.Field
	ScriptVersion            apijson.Field
	Truncated                apijson.Field
	WallTimeMs               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r observabilityTelemetryQueryNewResponseInvocationsWorkersJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilityTelemetryQueryNewResponseInvocationsWorkers) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilityTelemetryQueryNewResponseInvocationsWorkers{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ObservabilityTelemetryQueryNewResponseInvocationsWorkersUnion] interface which
// you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [workers.ObservabilityTelemetryQueryNewResponseInvocationsWorkersObject],
// [workers.ObservabilityTelemetryQueryNewResponseInvocationsWorkersObject].
func (r ObservabilityTelemetryQueryNewResponseInvocationsWorkers) AsUnion() ObservabilityTelemetryQueryNewResponseInvocationsWorkersUnion {
	return r.union
}

// Cloudflare Workers event information enriches your logs so you can easily
// identify and debug issues.
//
// Union satisfied by
// [workers.ObservabilityTelemetryQueryNewResponseInvocationsWorkersObject] or
// [workers.ObservabilityTelemetryQueryNewResponseInvocationsWorkersObject].
type ObservabilityTelemetryQueryNewResponseInvocationsWorkersUnion interface {
	implementsObservabilityTelemetryQueryNewResponseInvocationsWorkers()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseInvocationsWorkersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryNewResponseInvocationsWorkersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryNewResponseInvocationsWorkersObject{}),
		},
	)
}

type ObservabilityTelemetryQueryNewResponseInvocationsWorkersObject struct {
	EventType      ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType             `json:"eventType,required"`
	Outcome        string                                                                              `json:"outcome,required"`
	RequestID      string                                                                              `json:"requestId,required"`
	ScriptName     string                                                                              `json:"scriptName,required"`
	Entrypoint     string                                                                              `json:"entrypoint"`
	Event          map[string]ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventUnion `json:"event"`
	ExecutionModel ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectExecutionModel        `json:"executionModel"`
	ScriptVersion  ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectScriptVersion         `json:"scriptVersion"`
	Truncated      bool                                                                                `json:"truncated"`
	JSON           observabilityTelemetryQueryNewResponseInvocationsWorkersObjectJSON                  `json:"-"`
}

// observabilityTelemetryQueryNewResponseInvocationsWorkersObjectJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseInvocationsWorkersObject]
type observabilityTelemetryQueryNewResponseInvocationsWorkersObjectJSON struct {
	EventType      apijson.Field
	Outcome        apijson.Field
	RequestID      apijson.Field
	ScriptName     apijson.Field
	Entrypoint     apijson.Field
	Event          apijson.Field
	ExecutionModel apijson.Field
	ScriptVersion  apijson.Field
	Truncated      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseInvocationsWorkersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseInvocationsWorkersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityTelemetryQueryNewResponseInvocationsWorkersObject) implementsObservabilityTelemetryQueryNewResponseInvocationsWorkers() {
}

type ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType string

const (
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeFetch     ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType = "fetch"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeScheduled ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType = "scheduled"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeAlarm     ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType = "alarm"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeCron      ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType = "cron"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeQueue     ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType = "queue"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeEmail     ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType = "email"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeTail      ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType = "tail"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeRpc       ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType = "rpc"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeWebsocket ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType = "websocket"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeUnknown   ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType = "unknown"
)

func (r ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeFetch, ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeScheduled, ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeAlarm, ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeCron, ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeQueue, ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeEmail, ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeTail, ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeRpc, ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeWebsocket, ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventTypeUnknown:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool]
// or
// [workers.ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMap].
type ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventUnion interface {
	ImplementsObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventUnion)(nil)).Elem(),
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

type ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMap map[string]ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapUnionItem

func (r ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMap) ImplementsObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventUnion() {
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool]
// or
// [workers.ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMap].
type ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapUnionItem interface {
	ImplementsObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapUnionItem)(nil)).Elem(),
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

type ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMap map[string]ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapUnionItem

func (r ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMap) ImplementsObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapUnionItem() {
}

// Union satisfied by
// [workers.ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapArray],
// [shared.UnionString], [shared.UnionFloat] or [shared.UnionBool].
type ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapUnionItem interface {
	ImplementsObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapUnionItem)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapArray{}),
		},
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

type ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapArray []ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapArrayUnionItem

func (r ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapArray) ImplementsObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapUnionItem() {
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapArrayUnionItem interface {
	ImplementsObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapArrayUnionItem()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectEventMapMapArrayUnionItem)(nil)).Elem(),
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

type ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectExecutionModel string

const (
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectExecutionModelDurableObject ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectExecutionModel = "durableObject"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectExecutionModelStateless     ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectExecutionModel = "stateless"
)

func (r ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectExecutionModelDurableObject, ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectScriptVersion struct {
	ID      string                                                                          `json:"id"`
	Message string                                                                          `json:"message"`
	Tag     string                                                                          `json:"tag"`
	JSON    observabilityTelemetryQueryNewResponseInvocationsWorkersObjectScriptVersionJSON `json:"-"`
}

// observabilityTelemetryQueryNewResponseInvocationsWorkersObjectScriptVersionJSON
// contains the JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectScriptVersion]
type observabilityTelemetryQueryNewResponseInvocationsWorkersObjectScriptVersionJSON struct {
	ID          apijson.Field
	Message     apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseInvocationsWorkersObjectScriptVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseInvocationsWorkersObjectScriptVersionJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType string

const (
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeFetch     ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType = "fetch"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeScheduled ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType = "scheduled"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeAlarm     ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType = "alarm"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeCron      ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType = "cron"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeQueue     ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType = "queue"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeEmail     ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType = "email"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeTail      ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType = "tail"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeRpc       ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType = "rpc"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeWebsocket ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType = "websocket"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeUnknown   ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType = "unknown"
)

func (r ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeFetch, ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeScheduled, ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeAlarm, ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeCron, ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeQueue, ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeEmail, ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeTail, ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeRpc, ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeWebsocket, ObservabilityTelemetryQueryNewResponseInvocationsWorkersEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseInvocationsWorkersExecutionModel string

const (
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersExecutionModelDurableObject ObservabilityTelemetryQueryNewResponseInvocationsWorkersExecutionModel = "durableObject"
	ObservabilityTelemetryQueryNewResponseInvocationsWorkersExecutionModelStateless     ObservabilityTelemetryQueryNewResponseInvocationsWorkersExecutionModel = "stateless"
)

func (r ObservabilityTelemetryQueryNewResponseInvocationsWorkersExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseInvocationsWorkersExecutionModelDurableObject, ObservabilityTelemetryQueryNewResponseInvocationsWorkersExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponsePattern struct {
	Count   float64                                              `json:"count,required"`
	Pattern string                                               `json:"pattern,required"`
	Series  []ObservabilityTelemetryQueryNewResponsePatternsSery `json:"series,required"`
	Service string                                               `json:"service,required"`
	JSON    observabilityTelemetryQueryNewResponsePatternJSON    `json:"-"`
}

// observabilityTelemetryQueryNewResponsePatternJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryNewResponsePattern]
type observabilityTelemetryQueryNewResponsePatternJSON struct {
	Count       apijson.Field
	Pattern     apijson.Field
	Series      apijson.Field
	Service     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponsePattern) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponsePatternJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponsePatternsSery struct {
	Data ObservabilityTelemetryQueryNewResponsePatternsSeriesData `json:"data,required"`
	Time string                                                   `json:"time,required"`
	JSON observabilityTelemetryQueryNewResponsePatternsSeryJSON   `json:"-"`
}

// observabilityTelemetryQueryNewResponsePatternsSeryJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryNewResponsePatternsSery]
type observabilityTelemetryQueryNewResponsePatternsSeryJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponsePatternsSery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponsePatternsSeryJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponsePatternsSeriesData struct {
	Count          float64                                                         `json:"count,required"`
	Interval       float64                                                         `json:"interval,required"`
	SampleInterval float64                                                         `json:"sampleInterval,required"`
	Value          float64                                                         `json:"value,required"`
	Groups         []ObservabilityTelemetryQueryNewResponsePatternsSeriesDataGroup `json:"groups"`
	JSON           observabilityTelemetryQueryNewResponsePatternsSeriesDataJSON    `json:"-"`
}

// observabilityTelemetryQueryNewResponsePatternsSeriesDataJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryNewResponsePatternsSeriesData]
type observabilityTelemetryQueryNewResponsePatternsSeriesDataJSON struct {
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponsePatternsSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponsePatternsSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponsePatternsSeriesDataGroup struct {
	Key   string                                                                   `json:"key,required"`
	Value ObservabilityTelemetryQueryNewResponsePatternsSeriesDataGroupsValueUnion `json:"value,required"`
	JSON  observabilityTelemetryQueryNewResponsePatternsSeriesDataGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryNewResponsePatternsSeriesDataGroupJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryNewResponsePatternsSeriesDataGroup]
type observabilityTelemetryQueryNewResponsePatternsSeriesDataGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponsePatternsSeriesDataGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponsePatternsSeriesDataGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryNewResponsePatternsSeriesDataGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryNewResponsePatternsSeriesDataGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryNewResponsePatternsSeriesDataGroupsValueUnion)(nil)).Elem(),
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

type ObservabilityTelemetryQueryNewParams struct {
	AccountID       param.Field[string]                                          `path:"account_id,required"`
	QueryID         param.Field[string]                                          `json:"queryId,required"`
	Timeframe       param.Field[ObservabilityTelemetryQueryNewParamsTimeframe]   `json:"timeframe,required"`
	Chart           param.Field[bool]                                            `json:"chart"`
	Compare         param.Field[bool]                                            `json:"compare"`
	Dry             param.Field[bool]                                            `json:"dry"`
	Granularity     param.Field[float64]                                         `json:"granularity"`
	IgnoreSeries    param.Field[bool]                                            `json:"ignoreSeries"`
	Limit           param.Field[float64]                                         `json:"limit"`
	Offset          param.Field[string]                                          `json:"offset"`
	OffsetBy        param.Field[float64]                                         `json:"offsetBy"`
	OffsetDirection param.Field[string]                                          `json:"offsetDirection"`
	Parameters      param.Field[ObservabilityTelemetryQueryNewParamsParameters]  `json:"parameters"`
	PatternType     param.Field[ObservabilityTelemetryQueryNewParamsPatternType] `json:"patternType"`
	View            param.Field[ObservabilityTelemetryQueryNewParamsView]        `json:"view"`
}

func (r ObservabilityTelemetryQueryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryNewParamsTimeframe struct {
	From param.Field[float64] `json:"from,required"`
	To   param.Field[float64] `json:"to,required"`
}

func (r ObservabilityTelemetryQueryNewParamsTimeframe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryNewParamsParameters struct {
	// Create Calculations to compute as part of the query.
	Calculations param.Field[[]ObservabilityTelemetryQueryNewParamsParametersCalculation] `json:"calculations"`
	// Set the Datasets to query. Leave it empty to query all the datasets.
	Datasets param.Field[[]string] `json:"datasets"`
	// Set a Flag to describe how to combine the filters on the query.
	FilterCombination param.Field[ObservabilityTelemetryQueryNewParamsParametersFilterCombination] `json:"filterCombination"`
	// Configure the Filters to apply to the query.
	Filters param.Field[[]ObservabilityTelemetryQueryNewParamsParametersFilter] `json:"filters"`
	// Define how to group the results of the query.
	GroupBys param.Field[[]ObservabilityTelemetryQueryNewParamsParametersGroupBy] `json:"groupBys"`
	// Configure the Having clauses that filter on calculations in the query result.
	Havings param.Field[[]ObservabilityTelemetryQueryNewParamsParametersHaving] `json:"havings"`
	// Set a limit on the number of results / records returned by the query
	Limit param.Field[int64] `json:"limit"`
	// Define an expression to search using full-text search.
	Needle param.Field[ObservabilityTelemetryQueryNewParamsParametersNeedle] `json:"needle"`
	// Configure the order of the results returned by the query.
	OrderBy param.Field[ObservabilityTelemetryQueryNewParamsParametersOrderBy] `json:"orderBy"`
}

func (r ObservabilityTelemetryQueryNewParamsParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryNewParamsParametersCalculation struct {
	Operator param.Field[ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator] `json:"operator,required"`
	Alias    param.Field[string]                                                             `json:"alias"`
	Key      param.Field[string]                                                             `json:"key"`
	KeyType  param.Field[ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyType]  `json:"keyType"`
}

func (r ObservabilityTelemetryQueryNewParamsParametersCalculation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator string

const (
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorUniq              ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "uniq"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorCount             ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "count"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorMax               ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "max"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorMin               ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "min"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorSum               ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "sum"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorAvg               ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "avg"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorMedian            ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "median"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP001              ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "p001"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP01               ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "p01"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP05               ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "p05"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP10               ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "p10"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP25               ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "p25"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP75               ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "p75"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP90               ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "p90"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP95               ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "p95"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP99               ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "p99"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP999              ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "p999"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorStddev            ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "stddev"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorVariance          ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "variance"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorCountDistinct     ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "COUNT_DISTINCT"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorCountUppercase    ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "COUNT"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorMaxUppercase      ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "MAX"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorMinUppercase      ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "MIN"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorSumUppercase      ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "SUM"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorAvgUppercase      ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "AVG"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorMedianUppercase   ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "MEDIAN"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP001Uppercase     ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "P001"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP01Uppercase      ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "P01"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP05Uppercase      ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "P05"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP10Uppercase      ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "P10"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP25Uppercase      ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "P25"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP75Uppercase      ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "P75"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP90Uppercase      ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "P90"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP95Uppercase      ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "P95"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP99Uppercase      ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "P99"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP999Uppercase     ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "P999"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorStddevUppercase   ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "STDDEV"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorVarianceUppercase ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator = "VARIANCE"
)

func (r ObservabilityTelemetryQueryNewParamsParametersCalculationsOperator) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorUniq, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorCount, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorMax, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorMin, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorSum, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorAvg, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorMedian, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP001, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP01, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP05, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP10, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP25, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP75, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP90, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP95, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP99, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP999, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorStddev, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorVariance, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorCountDistinct, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorCountUppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorMaxUppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorMinUppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorSumUppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorAvgUppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorMedianUppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP001Uppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP01Uppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP05Uppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP10Uppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP25Uppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP75Uppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP90Uppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP95Uppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP99Uppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorP999Uppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorStddevUppercase, ObservabilityTelemetryQueryNewParamsParametersCalculationsOperatorVarianceUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyType string

const (
	ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyTypeString  ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyType = "string"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyTypeNumber  ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyType = "number"
	ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyTypeBoolean ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyType = "boolean"
)

func (r ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyTypeString, ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyTypeNumber, ObservabilityTelemetryQueryNewParamsParametersCalculationsKeyTypeBoolean:
		return true
	}
	return false
}

// Set a Flag to describe how to combine the filters on the query.
type ObservabilityTelemetryQueryNewParamsParametersFilterCombination string

const (
	ObservabilityTelemetryQueryNewParamsParametersFilterCombinationAnd          ObservabilityTelemetryQueryNewParamsParametersFilterCombination = "and"
	ObservabilityTelemetryQueryNewParamsParametersFilterCombinationOr           ObservabilityTelemetryQueryNewParamsParametersFilterCombination = "or"
	ObservabilityTelemetryQueryNewParamsParametersFilterCombinationAndUppercase ObservabilityTelemetryQueryNewParamsParametersFilterCombination = "AND"
	ObservabilityTelemetryQueryNewParamsParametersFilterCombinationOrUppercase  ObservabilityTelemetryQueryNewParamsParametersFilterCombination = "OR"
)

func (r ObservabilityTelemetryQueryNewParamsParametersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewParamsParametersFilterCombinationAnd, ObservabilityTelemetryQueryNewParamsParametersFilterCombinationOr, ObservabilityTelemetryQueryNewParamsParametersFilterCombinationAndUppercase, ObservabilityTelemetryQueryNewParamsParametersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewParamsParametersFilter struct {
	Key       param.Field[string]                                                          `json:"key,required"`
	Operation param.Field[ObservabilityTelemetryQueryNewParamsParametersFiltersOperation]  `json:"operation,required"`
	Type      param.Field[ObservabilityTelemetryQueryNewParamsParametersFiltersType]       `json:"type,required"`
	Value     param.Field[ObservabilityTelemetryQueryNewParamsParametersFiltersValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryQueryNewParamsParametersFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryNewParamsParametersFiltersOperation string

const (
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationIncludes            ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "includes"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationNotIncludes         ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "not_includes"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationStartsWith          ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "starts_with"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationRegex               ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "regex"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationExists              ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "exists"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationIsNull              ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "is_null"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationIn                  ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "in"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationNotIn               ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "not_in"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationEq                  ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "eq"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationNeq                 ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "neq"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationGt                  ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "gt"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationGte                 ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "gte"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationLt                  ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "lt"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationLte                 ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "lte"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationUnknown5            ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "="
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationNotEquals           ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "!="
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationGreater             ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = ">"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationGreaterOrEquals     ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = ">="
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationLess                ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "<"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationLessOrEquals        ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "<="
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationIncludesUppercase   ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "INCLUDES"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationDoesNotInclude      ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationMatchRegex          ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationExistsUppercase     ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "EXISTS"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationDoesNotExist        ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationInUppercase         ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "IN"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationNotInUppercase      ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "NOT_IN"
	ObservabilityTelemetryQueryNewParamsParametersFiltersOperationStartsWithUppercase ObservabilityTelemetryQueryNewParamsParametersFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryQueryNewParamsParametersFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewParamsParametersFiltersOperationIncludes, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationNotIncludes, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationStartsWith, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationRegex, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationExists, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationIsNull, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationIn, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationNotIn, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationEq, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationNeq, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationGt, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationGte, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationLt, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationLte, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationUnknown5, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationNotEquals, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationGreater, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationGreaterOrEquals, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationLess, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationLessOrEquals, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationIncludesUppercase, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationDoesNotInclude, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationMatchRegex, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationExistsUppercase, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationDoesNotExist, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationInUppercase, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationNotInUppercase, ObservabilityTelemetryQueryNewParamsParametersFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewParamsParametersFiltersType string

const (
	ObservabilityTelemetryQueryNewParamsParametersFiltersTypeString  ObservabilityTelemetryQueryNewParamsParametersFiltersType = "string"
	ObservabilityTelemetryQueryNewParamsParametersFiltersTypeNumber  ObservabilityTelemetryQueryNewParamsParametersFiltersType = "number"
	ObservabilityTelemetryQueryNewParamsParametersFiltersTypeBoolean ObservabilityTelemetryQueryNewParamsParametersFiltersType = "boolean"
)

func (r ObservabilityTelemetryQueryNewParamsParametersFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewParamsParametersFiltersTypeString, ObservabilityTelemetryQueryNewParamsParametersFiltersTypeNumber, ObservabilityTelemetryQueryNewParamsParametersFiltersTypeBoolean:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryQueryNewParamsParametersFiltersValueUnion interface {
	ImplementsObservabilityTelemetryQueryNewParamsParametersFiltersValueUnion()
}

type ObservabilityTelemetryQueryNewParamsParametersGroupBy struct {
	Type  param.Field[ObservabilityTelemetryQueryNewParamsParametersGroupBysType] `json:"type,required"`
	Value param.Field[string]                                                     `json:"value,required"`
}

func (r ObservabilityTelemetryQueryNewParamsParametersGroupBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryNewParamsParametersGroupBysType string

const (
	ObservabilityTelemetryQueryNewParamsParametersGroupBysTypeString  ObservabilityTelemetryQueryNewParamsParametersGroupBysType = "string"
	ObservabilityTelemetryQueryNewParamsParametersGroupBysTypeNumber  ObservabilityTelemetryQueryNewParamsParametersGroupBysType = "number"
	ObservabilityTelemetryQueryNewParamsParametersGroupBysTypeBoolean ObservabilityTelemetryQueryNewParamsParametersGroupBysType = "boolean"
)

func (r ObservabilityTelemetryQueryNewParamsParametersGroupBysType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewParamsParametersGroupBysTypeString, ObservabilityTelemetryQueryNewParamsParametersGroupBysTypeNumber, ObservabilityTelemetryQueryNewParamsParametersGroupBysTypeBoolean:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewParamsParametersHaving struct {
	Key       param.Field[string]                                                         `json:"key,required"`
	Operation param.Field[ObservabilityTelemetryQueryNewParamsParametersHavingsOperation] `json:"operation,required"`
	Value     param.Field[float64]                                                        `json:"value,required"`
}

func (r ObservabilityTelemetryQueryNewParamsParametersHaving) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryNewParamsParametersHavingsOperation string

const (
	ObservabilityTelemetryQueryNewParamsParametersHavingsOperationEq  ObservabilityTelemetryQueryNewParamsParametersHavingsOperation = "eq"
	ObservabilityTelemetryQueryNewParamsParametersHavingsOperationNeq ObservabilityTelemetryQueryNewParamsParametersHavingsOperation = "neq"
	ObservabilityTelemetryQueryNewParamsParametersHavingsOperationGt  ObservabilityTelemetryQueryNewParamsParametersHavingsOperation = "gt"
	ObservabilityTelemetryQueryNewParamsParametersHavingsOperationGte ObservabilityTelemetryQueryNewParamsParametersHavingsOperation = "gte"
	ObservabilityTelemetryQueryNewParamsParametersHavingsOperationLt  ObservabilityTelemetryQueryNewParamsParametersHavingsOperation = "lt"
	ObservabilityTelemetryQueryNewParamsParametersHavingsOperationLte ObservabilityTelemetryQueryNewParamsParametersHavingsOperation = "lte"
)

func (r ObservabilityTelemetryQueryNewParamsParametersHavingsOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewParamsParametersHavingsOperationEq, ObservabilityTelemetryQueryNewParamsParametersHavingsOperationNeq, ObservabilityTelemetryQueryNewParamsParametersHavingsOperationGt, ObservabilityTelemetryQueryNewParamsParametersHavingsOperationGte, ObservabilityTelemetryQueryNewParamsParametersHavingsOperationLt, ObservabilityTelemetryQueryNewParamsParametersHavingsOperationLte:
		return true
	}
	return false
}

// Define an expression to search using full-text search.
type ObservabilityTelemetryQueryNewParamsParametersNeedle struct {
	Value     param.Field[ObservabilityTelemetryQueryNewParamsParametersNeedleValueUnion] `json:"value,required"`
	IsRegex   param.Field[bool]                                                           `json:"isRegex"`
	MatchCase param.Field[bool]                                                           `json:"matchCase"`
}

func (r ObservabilityTelemetryQueryNewParamsParametersNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryQueryNewParamsParametersNeedleValueUnion interface {
	ImplementsObservabilityTelemetryQueryNewParamsParametersNeedleValueUnion()
}

// Configure the order of the results returned by the query.
type ObservabilityTelemetryQueryNewParamsParametersOrderBy struct {
	// Configure which Calculation to order the results by.
	Value param.Field[string] `json:"value,required"`
	// Set the order of the results
	Order param.Field[ObservabilityTelemetryQueryNewParamsParametersOrderByOrder] `json:"order"`
}

func (r ObservabilityTelemetryQueryNewParamsParametersOrderBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set the order of the results
type ObservabilityTelemetryQueryNewParamsParametersOrderByOrder string

const (
	ObservabilityTelemetryQueryNewParamsParametersOrderByOrderAsc  ObservabilityTelemetryQueryNewParamsParametersOrderByOrder = "asc"
	ObservabilityTelemetryQueryNewParamsParametersOrderByOrderDesc ObservabilityTelemetryQueryNewParamsParametersOrderByOrder = "desc"
)

func (r ObservabilityTelemetryQueryNewParamsParametersOrderByOrder) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewParamsParametersOrderByOrderAsc, ObservabilityTelemetryQueryNewParamsParametersOrderByOrderDesc:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewParamsPatternType string

const (
	ObservabilityTelemetryQueryNewParamsPatternTypeMessage ObservabilityTelemetryQueryNewParamsPatternType = "message"
	ObservabilityTelemetryQueryNewParamsPatternTypeError   ObservabilityTelemetryQueryNewParamsPatternType = "error"
)

func (r ObservabilityTelemetryQueryNewParamsPatternType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewParamsPatternTypeMessage, ObservabilityTelemetryQueryNewParamsPatternTypeError:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewParamsView string

const (
	ObservabilityTelemetryQueryNewParamsViewTraces       ObservabilityTelemetryQueryNewParamsView = "traces"
	ObservabilityTelemetryQueryNewParamsViewEvents       ObservabilityTelemetryQueryNewParamsView = "events"
	ObservabilityTelemetryQueryNewParamsViewCalculations ObservabilityTelemetryQueryNewParamsView = "calculations"
	ObservabilityTelemetryQueryNewParamsViewInvocations  ObservabilityTelemetryQueryNewParamsView = "invocations"
	ObservabilityTelemetryQueryNewParamsViewRequests     ObservabilityTelemetryQueryNewParamsView = "requests"
	ObservabilityTelemetryQueryNewParamsViewPatterns     ObservabilityTelemetryQueryNewParamsView = "patterns"
)

func (r ObservabilityTelemetryQueryNewParamsView) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewParamsViewTraces, ObservabilityTelemetryQueryNewParamsViewEvents, ObservabilityTelemetryQueryNewParamsViewCalculations, ObservabilityTelemetryQueryNewParamsViewInvocations, ObservabilityTelemetryQueryNewParamsViewRequests, ObservabilityTelemetryQueryNewParamsViewPatterns:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseEnvelope struct {
	Errors   []ObservabilityTelemetryQueryNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ObservabilityTelemetryQueryNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ObservabilityTelemetryQueryNewResponse                   `json:"result,required"`
	Success  ObservabilityTelemetryQueryNewResponseEnvelopeSuccess    `json:"success,required"`
	JSON     observabilityTelemetryQueryNewResponseEnvelopeJSON       `json:"-"`
}

// observabilityTelemetryQueryNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryNewResponseEnvelope]
type observabilityTelemetryQueryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseEnvelopeErrors struct {
	Message string                                                   `json:"message,required"`
	JSON    observabilityTelemetryQueryNewResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityTelemetryQueryNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryNewResponseEnvelopeErrors]
type observabilityTelemetryQueryNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseEnvelopeMessages struct {
	Message ObservabilityTelemetryQueryNewResponseEnvelopeMessagesMessage `json:"message,required"`
	JSON    observabilityTelemetryQueryNewResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityTelemetryQueryNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryNewResponseEnvelopeMessages]
type observabilityTelemetryQueryNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryNewResponseEnvelopeMessagesMessage string

const (
	ObservabilityTelemetryQueryNewResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityTelemetryQueryNewResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityTelemetryQueryNewResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryNewResponseEnvelopeSuccess bool

const (
	ObservabilityTelemetryQueryNewResponseEnvelopeSuccessTrue ObservabilityTelemetryQueryNewResponseEnvelopeSuccess = true
)

func (r ObservabilityTelemetryQueryNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
