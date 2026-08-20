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
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// ObservabilitySharedQueryService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilitySharedQueryService] method instead.
type ObservabilitySharedQueryService struct {
	Options []option.RequestOption
}

// NewObservabilitySharedQueryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservabilitySharedQueryService(opts ...option.RequestOption) (r *ObservabilitySharedQueryService) {
	r = &ObservabilitySharedQueryService{}
	r.Options = opts
	return
}

// Shared queries store the results of a previously run query, allowing you to
// share the results with others.
func (r *ObservabilitySharedQueryService) New(ctx context.Context, params ObservabilitySharedQueryNewParams, opts ...option.RequestOption) (res *ObservabilitySharedQueryNewResponse, err error) {
	var env ObservabilitySharedQueryNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/shared/query", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Shared queries store the results of a previously run query, allowing you to
// share the results with others.
func (r *ObservabilitySharedQueryService) Get(ctx context.Context, id string, params ObservabilitySharedQueryGetParams, opts ...option.RequestOption) (res *ObservabilitySharedQueryGetResponse, err error) {
	var env ObservabilitySharedQueryGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/shared/query/%s", params.AccountID, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ObservabilitySharedQueryNewResponse struct {
	// Specify the ID of the shared query.
	ID   string                                  `json:"id" api:"required"`
	JSON observabilitySharedQueryNewResponseJSON `json:"-"`
}

// observabilitySharedQueryNewResponseJSON contains the JSON metadata for the
// struct [ObservabilitySharedQueryNewResponse]
type observabilitySharedQueryNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryNewResponseJSON) RawJSON() string {
	return r.raw
}

// Complete results of a query run. The populated fields depend on the requested
// view type (events, calculations, invocations, traces, or agents).
type ObservabilitySharedQueryGetResponse struct {
	// Represents a single execution of a query against Workers Observability data,
	// including the query definition, execution status, and performance statistics.
	Run ObservabilitySharedQueryGetResponseRun `json:"run" api:"required"`
	// Query performance statistics from the database. Includes execution time, rows
	// scanned, and bytes read. Does not include network latency.
	Statistics ObservabilitySharedQueryGetResponseStatistics `json:"statistics" api:"required"`
	// Durable Object agent summaries. Present when the query view is 'agents'. Each
	// entry represents an agent with its event counts and status.
	Agents []ObservabilitySharedQueryGetResponseAgent `json:"agents"`
	// Aggregated calculation results. Present when the query view is 'calculations'.
	// Contains computed metrics (count, avg, p99, etc.) with optional group-by
	// breakdowns and time-series data.
	Calculations []ObservabilitySharedQueryGetResponseCalculation `json:"calculations"`
	// Comparison calculation results from the previous time period. Present when the
	// compare option is enabled. Same structure as calculations.
	Compare []ObservabilitySharedQueryGetResponseCompare `json:"compare"`
	// Individual event results. Present when the query view is 'events'. Contains the
	// matching log lines and their metadata.
	Events ObservabilitySharedQueryGetResponseEvents `json:"events"`
	// Events grouped by invocation (request ID). Present when the query view is
	// 'invocations'. Each key is a request ID mapping to all events from that
	// invocation.
	Invocations map[string][]ObservabilitySharedQueryGetResponseInvocation `json:"invocations"`
	// Trace summaries matching the query. Present when the query view is 'traces'.
	// Each entry represents a distributed trace with its spans, duration, and services
	// involved.
	Traces []ObservabilitySharedQueryGetResponseTrace `json:"traces"`
	JSON   observabilitySharedQueryGetResponseJSON    `json:"-"`
}

// observabilitySharedQueryGetResponseJSON contains the JSON metadata for the
// struct [ObservabilitySharedQueryGetResponse]
type observabilitySharedQueryGetResponseJSON struct {
	Run          apijson.Field
	Statistics   apijson.Field
	Agents       apijson.Field
	Calculations apijson.Field
	Compare      apijson.Field
	Events       apijson.Field
	Invocations  apijson.Field
	Traces       apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseJSON) RawJSON() string {
	return r.raw
}

// Represents a single execution of a query against Workers Observability data,
// including the query definition, execution status, and performance statistics.
type ObservabilitySharedQueryGetResponseRun struct {
	// Unique identifier for this query run.
	ID string `json:"id" api:"required"`
	// Cloudflare account ID that owns this query run.
	AccountID string `json:"accountId" api:"required"`
	// Whether this was a dry run (results not persisted).
	Dry bool `json:"dry" api:"required"`
	// Number of time-series buckets used for the query. Higher values produce more
	// detailed series data.
	Granularity float64 `json:"granularity" api:"required"`
	// A saved query definition with its parameters, metadata, and ownership
	// information.
	Query ObservabilitySharedQueryGetResponseRunQuery `json:"query" api:"required"`
	// Current execution status of the query run.
	Status ObservabilitySharedQueryGetResponseRunStatus `json:"status" api:"required"`
	// Time range for the query execution
	Timeframe ObservabilitySharedQueryGetResponseRunTimeframe `json:"timeframe" api:"required"`
	// ID of the user who initiated the query run.
	UserID string `json:"userId" api:"required"`
	// ISO-8601 timestamp when the query run was created.
	Created string `json:"created"`
	// Query performance statistics from the database (does not include network
	// latency).
	Statistics ObservabilitySharedQueryGetResponseRunStatistics `json:"statistics"`
	// ISO-8601 timestamp when the query run was last updated.
	Updated string                                     `json:"updated"`
	JSON    observabilitySharedQueryGetResponseRunJSON `json:"-"`
}

// observabilitySharedQueryGetResponseRunJSON contains the JSON metadata for the
// struct [ObservabilitySharedQueryGetResponseRun]
type observabilitySharedQueryGetResponseRunJSON struct {
	ID          apijson.Field
	AccountID   apijson.Field
	Dry         apijson.Field
	Granularity apijson.Field
	Query       apijson.Field
	Status      apijson.Field
	Timeframe   apijson.Field
	UserID      apijson.Field
	Created     apijson.Field
	Statistics  apijson.Field
	Updated     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseRun) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunJSON) RawJSON() string {
	return r.raw
}

// A saved query definition with its parameters, metadata, and ownership
// information.
type ObservabilitySharedQueryGetResponseRunQuery struct {
	ID string `json:"id" api:"required"`
	// If the query wasn't explcitly saved
	Adhoc       bool   `json:"adhoc" api:"required"`
	Created     string `json:"created" api:"required"`
	CreatedBy   string `json:"createdBy" api:"required"`
	Description string `json:"description" api:"required,nullable"`
	// Query name
	Name       string                                                `json:"name" api:"required"`
	Parameters ObservabilitySharedQueryGetResponseRunQueryParameters `json:"parameters" api:"required"`
	Updated    string                                                `json:"updated" api:"required"`
	UpdatedBy  string                                                `json:"updatedBy" api:"required"`
	JSON       observabilitySharedQueryGetResponseRunQueryJSON       `json:"-"`
}

// observabilitySharedQueryGetResponseRunQueryJSON contains the JSON metadata for
// the struct [ObservabilitySharedQueryGetResponseRunQuery]
type observabilitySharedQueryGetResponseRunQueryJSON struct {
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

func (r *ObservabilitySharedQueryGetResponseRunQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunQueryJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseRunQueryParameters struct {
	// Create Calculations to compute as part of the query.
	Calculations []ObservabilitySharedQueryGetResponseRunQueryParametersCalculation `json:"calculations"`
	// Set the Datasets to query. Leave it empty to query all the datasets.
	Datasets []string `json:"datasets"`
	// Set a Flag to describe how to combine the filters on the query.
	FilterCombination ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombination `json:"filterCombination"`
	// Configure the Filters to apply to the query. Supports nested groups via kind:
	// 'group'.
	Filters []ObservabilitySharedQueryGetResponseRunQueryParametersFilter `json:"filters"`
	// Define how to group the results of the query.
	GroupBys []ObservabilitySharedQueryGetResponseRunQueryParametersGroupBy `json:"groupBys"`
	// Configure the Having clauses that filter on calculations in the query result.
	Havings []ObservabilitySharedQueryGetResponseRunQueryParametersHaving `json:"havings"`
	// Set a limit on the number of results / records returned by the query
	Limit int64 `json:"limit"`
	// Define an expression to search using full-text search.
	Needle ObservabilitySharedQueryGetResponseRunQueryParametersNeedle `json:"needle"`
	// Configure the order of the results returned by the query.
	OrderBy ObservabilitySharedQueryGetResponseRunQueryParametersOrderBy `json:"orderBy"`
	JSON    observabilitySharedQueryGetResponseRunQueryParametersJSON    `json:"-"`
}

// observabilitySharedQueryGetResponseRunQueryParametersJSON contains the JSON
// metadata for the struct [ObservabilitySharedQueryGetResponseRunQueryParameters]
type observabilitySharedQueryGetResponseRunQueryParametersJSON struct {
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

func (r *ObservabilitySharedQueryGetResponseRunQueryParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunQueryParametersJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseRunQueryParametersCalculation struct {
	Operator ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator `json:"operator" api:"required"`
	Alias    string                                                                    `json:"alias"`
	Key      string                                                                    `json:"key"`
	KeyType  ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsKeyType  `json:"keyType"`
	JSON     observabilitySharedQueryGetResponseRunQueryParametersCalculationJSON      `json:"-"`
}

// observabilitySharedQueryGetResponseRunQueryParametersCalculationJSON contains
// the JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseRunQueryParametersCalculation]
type observabilitySharedQueryGetResponseRunQueryParametersCalculationJSON struct {
	Operator    apijson.Field
	Alias       apijson.Field
	Key         apijson.Field
	KeyType     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseRunQueryParametersCalculation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunQueryParametersCalculationJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorUniq              ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "uniq"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorCount             ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "count"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorMax               ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "max"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorMin               ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "min"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorSum               ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "sum"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorAvg               ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "avg"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorMedian            ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "median"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP001              ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "p001"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP01               ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "p01"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP05               ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "p05"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP10               ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "p10"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP25               ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "p25"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP75               ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "p75"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP90               ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "p90"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP95               ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "p95"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP99               ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "p99"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP999              ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "p999"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorStddev            ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "stddev"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorVariance          ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "variance"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorCountDistinct     ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "COUNT_DISTINCT"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorCountUppercase    ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "COUNT"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorMaxUppercase      ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "MAX"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorMinUppercase      ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "MIN"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorSumUppercase      ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "SUM"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorAvgUppercase      ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "AVG"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorMedianUppercase   ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "MEDIAN"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP001Uppercase     ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "P001"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP01Uppercase      ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "P01"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP05Uppercase      ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "P05"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP10Uppercase      ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "P10"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP25Uppercase      ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "P25"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP75Uppercase      ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "P75"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP90Uppercase      ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "P90"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP95Uppercase      ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "P95"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP99Uppercase      ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "P99"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP999Uppercase     ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "P999"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorStddevUppercase   ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "STDDEV"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorVarianceUppercase ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator = "VARIANCE"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperator) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorUniq, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorCount, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorMax, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorMin, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorSum, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorAvg, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorMedian, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP001, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP01, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP05, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP10, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP25, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP75, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP90, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP95, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP99, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP999, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorStddev, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorVariance, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorCountDistinct, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorCountUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorMaxUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorMinUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorSumUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorAvgUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorMedianUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP001Uppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP01Uppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP05Uppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP10Uppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP25Uppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP75Uppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP90Uppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP95Uppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP99Uppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorP999Uppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorStddevUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsOperatorVarianceUppercase:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsKeyType string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsKeyTypeString  ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsKeyType = "string"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsKeyTypeNumber  ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsKeyType = "number"
	ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsKeyTypeBoolean ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsKeyType = "boolean"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsKeyType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsKeyTypeString, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsKeyTypeNumber, ObservabilitySharedQueryGetResponseRunQueryParametersCalculationsKeyTypeBoolean:
		return true
	}
	return false
}

// Set a Flag to describe how to combine the filters on the query.
type ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombination string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombinationAnd          ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombination = "and"
	ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombinationOr           ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombination = "or"
	ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombinationAndUppercase ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombination = "AND"
	ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombinationOrUppercase  ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombination = "OR"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombinationAnd, ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombinationOr, ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombinationAndUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFilterCombinationOrUppercase:
		return true
	}
	return false
}

// Supports nested groups via kind: 'group'.
type ObservabilitySharedQueryGetResponseRunQueryParametersFilter struct {
	FilterCombination ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombination `json:"filterCombination"`
	// This field can have the runtime type of [[]interface{}].
	Filters interface{} `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  string                                                           `json:"key"`
	Kind ObservabilitySharedQueryGetResponseRunQueryParametersFiltersKind `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
	// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersType `json:"type"`
	// This field can have the runtime type of
	// [ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion].
	Value interface{}                                                     `json:"value"`
	JSON  observabilitySharedQueryGetResponseRunQueryParametersFilterJSON `json:"-"`
	union ObservabilitySharedQueryGetResponseRunQueryParametersFiltersUnion
}

// observabilitySharedQueryGetResponseRunQueryParametersFilterJSON contains the
// JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseRunQueryParametersFilter]
type observabilitySharedQueryGetResponseRunQueryParametersFilterJSON struct {
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

func (r observabilitySharedQueryGetResponseRunQueryParametersFilterJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilitySharedQueryGetResponseRunQueryParametersFilter) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilitySharedQueryGetResponseRunQueryParametersFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ObservabilitySharedQueryGetResponseRunQueryParametersFiltersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObject],
// [ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf].
func (r ObservabilitySharedQueryGetResponseRunQueryParametersFilter) AsUnion() ObservabilitySharedQueryGetResponseRunQueryParametersFiltersUnion {
	return r.union
}

// Supports nested groups via kind: 'group'.
//
// Union satisfied by
// [ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObject] or
// [ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf].
type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersUnion interface {
	implementsObservabilitySharedQueryGetResponseRunQueryParametersFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilitySharedQueryGetResponseRunQueryParametersFiltersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf{}),
		},
	)
}

type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObject struct {
	FilterCombination ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombination `json:"filterCombination" api:"required"`
	Filters           []interface{}                                                                       `json:"filters" api:"required"`
	Kind              ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectKind              `json:"kind" api:"required"`
	JSON              observabilitySharedQueryGetResponseRunQueryParametersFiltersObjectJSON              `json:"-"`
}

// observabilitySharedQueryGetResponseRunQueryParametersFiltersObjectJSON contains
// the JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObject]
type observabilitySharedQueryGetResponseRunQueryParametersFiltersObjectJSON struct {
	FilterCombination apijson.Field
	Filters           apijson.Field
	Kind              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunQueryParametersFiltersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObject) implementsObservabilitySharedQueryGetResponseRunQueryParametersFilter() {
}

type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombination string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombinationAnd          ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombination = "and"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombinationOr           ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombination = "or"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombinationAndUppercase ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombination = "AND"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombinationOrUppercase  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombination = "OR"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombinationAnd, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombinationOr, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombinationAndUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectKind string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectKindGroup ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectKind = "group"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key string `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
	// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKind `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion `json:"value"`
	JSON  observabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafJSON       `json:"-"`
}

// observabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafJSON
// contains the JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf]
type observabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafJSON struct {
	Key         apijson.Field
	Operation   apijson.Field
	Type        apijson.Field
	Kind        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf) implementsObservabilitySharedQueryGetResponseRunQueryParametersFilter() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationEndsWith            ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "ends_with"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationEndsWithUppercase   ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "ENDS_WITH"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationEndsWith, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationEndsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafTypeString  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafTypeString, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKindFilter ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKindFilter:
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
type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilitySharedQueryGetResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion)(nil)).Elem(),
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

type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombination string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombinationAnd          ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombination = "and"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombinationOr           ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombination = "or"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombinationAndUppercase ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombination = "AND"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombinationOrUppercase  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombination = "OR"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombinationAnd, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombinationOr, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombinationAndUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersKind string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersKindGroup  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersKind = "group"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersKindFilter ObservabilitySharedQueryGetResponseRunQueryParametersFiltersKind = "filter"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersFiltersKindGroup, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationIncludes            ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "includes"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationNotIncludes         ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "not_includes"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationStartsWith          ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "starts_with"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationEndsWith            ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "ends_with"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationRegex               ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "regex"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationExists              ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "exists"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationIsNull              ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "is_null"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationIn                  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "in"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationNotIn               ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "not_in"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationEq                  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "eq"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationNeq                 ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "neq"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationGt                  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "gt"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationGte                 ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "gte"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationLt                  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "lt"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationLte                 ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "lte"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationEquals              ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "="
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationNotEquals           ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "!="
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationGreater             ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = ">"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationGreaterOrEquals     ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = ">="
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationLess                ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "<"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationLessOrEquals        ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "<="
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationIncludesUppercase   ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "INCLUDES"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationDoesNotInclude      ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationMatchRegex          ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "MATCH_REGEX"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationExistsUppercase     ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "EXISTS"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationDoesNotExist        ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "DOES_NOT_EXIST"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationInUppercase         ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "IN"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationNotInUppercase      ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "NOT_IN"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationStartsWithUppercase ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "STARTS_WITH"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationEndsWithUppercase   ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation = "ENDS_WITH"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationIncludes, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationNotIncludes, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationStartsWith, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationEndsWith, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationRegex, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationExists, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationIsNull, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationIn, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationNotIn, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationEq, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationNeq, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationGt, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationGte, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationLt, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationLte, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationEquals, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationNotEquals, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationGreater, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationGreaterOrEquals, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationLess, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationLessOrEquals, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationIncludesUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationDoesNotInclude, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationMatchRegex, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationExistsUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationDoesNotExist, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationInUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationNotInUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationStartsWithUppercase, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersOperationEndsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilitySharedQueryGetResponseRunQueryParametersFiltersType string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersTypeString  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersType = "string"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersTypeNumber  ObservabilitySharedQueryGetResponseRunQueryParametersFiltersType = "number"
	ObservabilitySharedQueryGetResponseRunQueryParametersFiltersTypeBoolean ObservabilitySharedQueryGetResponseRunQueryParametersFiltersType = "boolean"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersFiltersType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersFiltersTypeString, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersTypeNumber, ObservabilitySharedQueryGetResponseRunQueryParametersFiltersTypeBoolean:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseRunQueryParametersGroupBy struct {
	Type  ObservabilitySharedQueryGetResponseRunQueryParametersGroupBysType `json:"type" api:"required"`
	Value string                                                            `json:"value" api:"required"`
	JSON  observabilitySharedQueryGetResponseRunQueryParametersGroupByJSON  `json:"-"`
}

// observabilitySharedQueryGetResponseRunQueryParametersGroupByJSON contains the
// JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseRunQueryParametersGroupBy]
type observabilitySharedQueryGetResponseRunQueryParametersGroupByJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseRunQueryParametersGroupBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunQueryParametersGroupByJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseRunQueryParametersGroupBysType string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersGroupBysTypeString  ObservabilitySharedQueryGetResponseRunQueryParametersGroupBysType = "string"
	ObservabilitySharedQueryGetResponseRunQueryParametersGroupBysTypeNumber  ObservabilitySharedQueryGetResponseRunQueryParametersGroupBysType = "number"
	ObservabilitySharedQueryGetResponseRunQueryParametersGroupBysTypeBoolean ObservabilitySharedQueryGetResponseRunQueryParametersGroupBysType = "boolean"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersGroupBysType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersGroupBysTypeString, ObservabilitySharedQueryGetResponseRunQueryParametersGroupBysTypeNumber, ObservabilitySharedQueryGetResponseRunQueryParametersGroupBysTypeBoolean:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseRunQueryParametersHaving struct {
	Key       string                                                                `json:"key" api:"required"`
	Operation ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperation `json:"operation" api:"required"`
	Value     float64                                                               `json:"value" api:"required"`
	JSON      observabilitySharedQueryGetResponseRunQueryParametersHavingJSON       `json:"-"`
}

// observabilitySharedQueryGetResponseRunQueryParametersHavingJSON contains the
// JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseRunQueryParametersHaving]
type observabilitySharedQueryGetResponseRunQueryParametersHavingJSON struct {
	Key         apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseRunQueryParametersHaving) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunQueryParametersHavingJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperation string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperationEq  ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperation = "eq"
	ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperationNeq ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperation = "neq"
	ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperationGt  ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperation = "gt"
	ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperationGte ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperation = "gte"
	ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperationLt  ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperation = "lt"
	ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperationLte ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperation = "lte"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperation) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperationEq, ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperationNeq, ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperationGt, ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperationGte, ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperationLt, ObservabilitySharedQueryGetResponseRunQueryParametersHavingsOperationLte:
		return true
	}
	return false
}

// Define an expression to search using full-text search.
type ObservabilitySharedQueryGetResponseRunQueryParametersNeedle struct {
	Value     ObservabilitySharedQueryGetResponseRunQueryParametersNeedleValue `json:"value" api:"required"`
	IsRegex   bool                                                             `json:"isRegex"`
	MatchCase bool                                                             `json:"matchCase"`
	JSON      observabilitySharedQueryGetResponseRunQueryParametersNeedleJSON  `json:"-"`
}

// observabilitySharedQueryGetResponseRunQueryParametersNeedleJSON contains the
// JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseRunQueryParametersNeedle]
type observabilitySharedQueryGetResponseRunQueryParametersNeedleJSON struct {
	Value       apijson.Field
	IsRegex     apijson.Field
	MatchCase   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseRunQueryParametersNeedle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunQueryParametersNeedleJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseRunQueryParametersNeedleValue struct {
	JSON observabilitySharedQueryGetResponseRunQueryParametersNeedleValueJSON `json:"-"`
}

// observabilitySharedQueryGetResponseRunQueryParametersNeedleValueJSON contains
// the JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseRunQueryParametersNeedleValue]
type observabilitySharedQueryGetResponseRunQueryParametersNeedleValueJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseRunQueryParametersNeedleValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunQueryParametersNeedleValueJSON) RawJSON() string {
	return r.raw
}

// Configure the order of the results returned by the query.
type ObservabilitySharedQueryGetResponseRunQueryParametersOrderBy struct {
	// Configure which Calculation to order the results by.
	Value string `json:"value" api:"required"`
	// Set the order of the results
	Order ObservabilitySharedQueryGetResponseRunQueryParametersOrderByOrder `json:"order"`
	JSON  observabilitySharedQueryGetResponseRunQueryParametersOrderByJSON  `json:"-"`
}

// observabilitySharedQueryGetResponseRunQueryParametersOrderByJSON contains the
// JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseRunQueryParametersOrderBy]
type observabilitySharedQueryGetResponseRunQueryParametersOrderByJSON struct {
	Value       apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseRunQueryParametersOrderBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunQueryParametersOrderByJSON) RawJSON() string {
	return r.raw
}

// Set the order of the results
type ObservabilitySharedQueryGetResponseRunQueryParametersOrderByOrder string

const (
	ObservabilitySharedQueryGetResponseRunQueryParametersOrderByOrderAsc  ObservabilitySharedQueryGetResponseRunQueryParametersOrderByOrder = "asc"
	ObservabilitySharedQueryGetResponseRunQueryParametersOrderByOrderDesc ObservabilitySharedQueryGetResponseRunQueryParametersOrderByOrder = "desc"
)

func (r ObservabilitySharedQueryGetResponseRunQueryParametersOrderByOrder) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunQueryParametersOrderByOrderAsc, ObservabilitySharedQueryGetResponseRunQueryParametersOrderByOrderDesc:
		return true
	}
	return false
}

// Current execution status of the query run.
type ObservabilitySharedQueryGetResponseRunStatus string

const (
	ObservabilitySharedQueryGetResponseRunStatusStarted   ObservabilitySharedQueryGetResponseRunStatus = "STARTED"
	ObservabilitySharedQueryGetResponseRunStatusCompleted ObservabilitySharedQueryGetResponseRunStatus = "COMPLETED"
)

func (r ObservabilitySharedQueryGetResponseRunStatus) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseRunStatusStarted, ObservabilitySharedQueryGetResponseRunStatusCompleted:
		return true
	}
	return false
}

// Time range for the query execution
type ObservabilitySharedQueryGetResponseRunTimeframe struct {
	// Start timestamp for the query timeframe (Unix timestamp in milliseconds)
	From float64 `json:"from" api:"required"`
	// End timestamp for the query timeframe (Unix timestamp in milliseconds)
	To   float64                                             `json:"to" api:"required"`
	JSON observabilitySharedQueryGetResponseRunTimeframeJSON `json:"-"`
}

// observabilitySharedQueryGetResponseRunTimeframeJSON contains the JSON metadata
// for the struct [ObservabilitySharedQueryGetResponseRunTimeframe]
type observabilitySharedQueryGetResponseRunTimeframeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseRunTimeframe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunTimeframeJSON) RawJSON() string {
	return r.raw
}

// Query performance statistics from the database (does not include network
// latency).
type ObservabilitySharedQueryGetResponseRunStatistics struct {
	// Number of uncompressed bytes read from the table.
	BytesRead float64 `json:"bytes_read" api:"required"`
	// Time in seconds for the query to run.
	Elapsed float64 `json:"elapsed" api:"required"`
	// Number of rows scanned from the table.
	RowsRead float64 `json:"rows_read" api:"required"`
	// The level of Adaptive Bit Rate (ABR) sampling used for the query. If empty the
	// ABR level is 1
	AbrLevel float64                                              `json:"abr_level"`
	JSON     observabilitySharedQueryGetResponseRunStatisticsJSON `json:"-"`
}

// observabilitySharedQueryGetResponseRunStatisticsJSON contains the JSON metadata
// for the struct [ObservabilitySharedQueryGetResponseRunStatistics]
type observabilitySharedQueryGetResponseRunStatisticsJSON struct {
	BytesRead   apijson.Field
	Elapsed     apijson.Field
	RowsRead    apijson.Field
	AbrLevel    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseRunStatistics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseRunStatisticsJSON) RawJSON() string {
	return r.raw
}

// Query performance statistics from the database. Includes execution time, rows
// scanned, and bytes read. Does not include network latency.
type ObservabilitySharedQueryGetResponseStatistics struct {
	// Number of uncompressed bytes read from the table.
	BytesRead float64 `json:"bytes_read" api:"required"`
	// Time in seconds for the query to run.
	Elapsed float64 `json:"elapsed" api:"required"`
	// Number of rows scanned from the table.
	RowsRead float64 `json:"rows_read" api:"required"`
	// The level of Adaptive Bit Rate (ABR) sampling used for the query. If empty the
	// ABR level is 1
	AbrLevel float64                                           `json:"abr_level"`
	JSON     observabilitySharedQueryGetResponseStatisticsJSON `json:"-"`
}

// observabilitySharedQueryGetResponseStatisticsJSON contains the JSON metadata for
// the struct [ObservabilitySharedQueryGetResponseStatistics]
type observabilitySharedQueryGetResponseStatisticsJSON struct {
	BytesRead   apijson.Field
	Elapsed     apijson.Field
	RowsRead    apijson.Field
	AbrLevel    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseStatistics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseStatisticsJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseAgent struct {
	// Class name of the Durable Object agent.
	AgentClass string `json:"agentClass" api:"required"`
	// Breakdown of event counts by event type.
	EventTypeCounts map[string]float64 `json:"eventTypeCounts" api:"required"`
	// Timestamp of the earliest event from this agent in the queried window (Unix
	// epoch ms).
	FirstEventMs float64 `json:"firstEventMs" api:"required"`
	// Whether the agent emitted any error events in the queried window.
	HasErrors bool `json:"hasErrors" api:"required"`
	// Timestamp of the most recent event from this agent (Unix epoch ms).
	LastEventMs float64 `json:"lastEventMs" api:"required"`
	// Durable Object namespace the agent belongs to.
	Namespace string `json:"namespace" api:"required"`
	// Worker service name that hosts this agent.
	Service string `json:"service" api:"required"`
	// Total number of events emitted by this agent in the queried window.
	TotalEvents float64                                      `json:"totalEvents" api:"required"`
	JSON        observabilitySharedQueryGetResponseAgentJSON `json:"-"`
}

// observabilitySharedQueryGetResponseAgentJSON contains the JSON metadata for the
// struct [ObservabilitySharedQueryGetResponseAgent]
type observabilitySharedQueryGetResponseAgentJSON struct {
	AgentClass      apijson.Field
	EventTypeCounts apijson.Field
	FirstEventMs    apijson.Field
	HasErrors       apijson.Field
	LastEventMs     apijson.Field
	Namespace       apijson.Field
	Service         apijson.Field
	TotalEvents     apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseAgent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseAgentJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseCalculation struct {
	Aggregates  []ObservabilitySharedQueryGetResponseCalculationsAggregate `json:"aggregates" api:"required"`
	Calculation string                                                     `json:"calculation" api:"required"`
	Series      []ObservabilitySharedQueryGetResponseCalculationsSeries    `json:"series" api:"required"`
	Alias       string                                                     `json:"alias"`
	JSON        observabilitySharedQueryGetResponseCalculationJSON         `json:"-"`
}

// observabilitySharedQueryGetResponseCalculationJSON contains the JSON metadata
// for the struct [ObservabilitySharedQueryGetResponseCalculation]
type observabilitySharedQueryGetResponseCalculationJSON struct {
	Aggregates  apijson.Field
	Calculation apijson.Field
	Series      apijson.Field
	Alias       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseCalculation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseCalculationJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseCalculationsAggregate struct {
	Count          float64                                                          `json:"count" api:"required"`
	Interval       float64                                                          `json:"interval" api:"required"`
	SampleInterval float64                                                          `json:"sampleInterval" api:"required"`
	Value          float64                                                          `json:"value" api:"required"`
	Groups         []ObservabilitySharedQueryGetResponseCalculationsAggregatesGroup `json:"groups"`
	JSON           observabilitySharedQueryGetResponseCalculationsAggregateJSON     `json:"-"`
}

// observabilitySharedQueryGetResponseCalculationsAggregateJSON contains the JSON
// metadata for the struct
// [ObservabilitySharedQueryGetResponseCalculationsAggregate]
type observabilitySharedQueryGetResponseCalculationsAggregateJSON struct {
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseCalculationsAggregate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseCalculationsAggregateJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseCalculationsAggregatesGroup struct {
	Key   string                                                                    `json:"key" api:"required"`
	Value ObservabilitySharedQueryGetResponseCalculationsAggregatesGroupsValueUnion `json:"value" api:"required"`
	JSON  observabilitySharedQueryGetResponseCalculationsAggregatesGroupJSON        `json:"-"`
}

// observabilitySharedQueryGetResponseCalculationsAggregatesGroupJSON contains the
// JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseCalculationsAggregatesGroup]
type observabilitySharedQueryGetResponseCalculationsAggregatesGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseCalculationsAggregatesGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseCalculationsAggregatesGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilitySharedQueryGetResponseCalculationsAggregatesGroupsValueUnion interface {
	ImplementsObservabilitySharedQueryGetResponseCalculationsAggregatesGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilitySharedQueryGetResponseCalculationsAggregatesGroupsValueUnion)(nil)).Elem(),
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

type ObservabilitySharedQueryGetResponseCalculationsSeries struct {
	Data []ObservabilitySharedQueryGetResponseCalculationsSeriesData `json:"data" api:"required"`
	Time string                                                      `json:"time" api:"required"`
	JSON observabilitySharedQueryGetResponseCalculationsSeriesJSON   `json:"-"`
}

// observabilitySharedQueryGetResponseCalculationsSeriesJSON contains the JSON
// metadata for the struct [ObservabilitySharedQueryGetResponseCalculationsSeries]
type observabilitySharedQueryGetResponseCalculationsSeriesJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseCalculationsSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseCalculationsSeriesJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseCalculationsSeriesData struct {
	Count          float64                                                          `json:"count" api:"required"`
	Interval       float64                                                          `json:"interval" api:"required"`
	SampleInterval float64                                                          `json:"sampleInterval" api:"required"`
	Value          float64                                                          `json:"value" api:"required"`
	FirstSeen      string                                                           `json:"firstSeen"`
	Groups         []ObservabilitySharedQueryGetResponseCalculationsSeriesDataGroup `json:"groups"`
	LastSeen       string                                                           `json:"lastSeen"`
	JSON           observabilitySharedQueryGetResponseCalculationsSeriesDataJSON    `json:"-"`
}

// observabilitySharedQueryGetResponseCalculationsSeriesDataJSON contains the JSON
// metadata for the struct
// [ObservabilitySharedQueryGetResponseCalculationsSeriesData]
type observabilitySharedQueryGetResponseCalculationsSeriesDataJSON struct {
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	FirstSeen      apijson.Field
	Groups         apijson.Field
	LastSeen       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseCalculationsSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseCalculationsSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseCalculationsSeriesDataGroup struct {
	Key   string                                                                    `json:"key" api:"required"`
	Value ObservabilitySharedQueryGetResponseCalculationsSeriesDataGroupsValueUnion `json:"value" api:"required"`
	JSON  observabilitySharedQueryGetResponseCalculationsSeriesDataGroupJSON        `json:"-"`
}

// observabilitySharedQueryGetResponseCalculationsSeriesDataGroupJSON contains the
// JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseCalculationsSeriesDataGroup]
type observabilitySharedQueryGetResponseCalculationsSeriesDataGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseCalculationsSeriesDataGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseCalculationsSeriesDataGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilitySharedQueryGetResponseCalculationsSeriesDataGroupsValueUnion interface {
	ImplementsObservabilitySharedQueryGetResponseCalculationsSeriesDataGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilitySharedQueryGetResponseCalculationsSeriesDataGroupsValueUnion)(nil)).Elem(),
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

type ObservabilitySharedQueryGetResponseCompare struct {
	Aggregates  []ObservabilitySharedQueryGetResponseCompareAggregate `json:"aggregates" api:"required"`
	Calculation string                                                `json:"calculation" api:"required"`
	Series      []ObservabilitySharedQueryGetResponseCompareSeries    `json:"series" api:"required"`
	Alias       string                                                `json:"alias"`
	JSON        observabilitySharedQueryGetResponseCompareJSON        `json:"-"`
}

// observabilitySharedQueryGetResponseCompareJSON contains the JSON metadata for
// the struct [ObservabilitySharedQueryGetResponseCompare]
type observabilitySharedQueryGetResponseCompareJSON struct {
	Aggregates  apijson.Field
	Calculation apijson.Field
	Series      apijson.Field
	Alias       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseCompare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseCompareJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseCompareAggregate struct {
	Count          float64                                                     `json:"count" api:"required"`
	Interval       float64                                                     `json:"interval" api:"required"`
	SampleInterval float64                                                     `json:"sampleInterval" api:"required"`
	Value          float64                                                     `json:"value" api:"required"`
	Groups         []ObservabilitySharedQueryGetResponseCompareAggregatesGroup `json:"groups"`
	JSON           observabilitySharedQueryGetResponseCompareAggregateJSON     `json:"-"`
}

// observabilitySharedQueryGetResponseCompareAggregateJSON contains the JSON
// metadata for the struct [ObservabilitySharedQueryGetResponseCompareAggregate]
type observabilitySharedQueryGetResponseCompareAggregateJSON struct {
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseCompareAggregate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseCompareAggregateJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseCompareAggregatesGroup struct {
	Key   string                                                               `json:"key" api:"required"`
	Value ObservabilitySharedQueryGetResponseCompareAggregatesGroupsValueUnion `json:"value" api:"required"`
	JSON  observabilitySharedQueryGetResponseCompareAggregatesGroupJSON        `json:"-"`
}

// observabilitySharedQueryGetResponseCompareAggregatesGroupJSON contains the JSON
// metadata for the struct
// [ObservabilitySharedQueryGetResponseCompareAggregatesGroup]
type observabilitySharedQueryGetResponseCompareAggregatesGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseCompareAggregatesGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseCompareAggregatesGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilitySharedQueryGetResponseCompareAggregatesGroupsValueUnion interface {
	ImplementsObservabilitySharedQueryGetResponseCompareAggregatesGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilitySharedQueryGetResponseCompareAggregatesGroupsValueUnion)(nil)).Elem(),
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

type ObservabilitySharedQueryGetResponseCompareSeries struct {
	Data []ObservabilitySharedQueryGetResponseCompareSeriesData `json:"data" api:"required"`
	Time string                                                 `json:"time" api:"required"`
	JSON observabilitySharedQueryGetResponseCompareSeriesJSON   `json:"-"`
}

// observabilitySharedQueryGetResponseCompareSeriesJSON contains the JSON metadata
// for the struct [ObservabilitySharedQueryGetResponseCompareSeries]
type observabilitySharedQueryGetResponseCompareSeriesJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseCompareSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseCompareSeriesJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseCompareSeriesData struct {
	Count          float64                                                     `json:"count" api:"required"`
	Interval       float64                                                     `json:"interval" api:"required"`
	SampleInterval float64                                                     `json:"sampleInterval" api:"required"`
	Value          float64                                                     `json:"value" api:"required"`
	FirstSeen      string                                                      `json:"firstSeen"`
	Groups         []ObservabilitySharedQueryGetResponseCompareSeriesDataGroup `json:"groups"`
	LastSeen       string                                                      `json:"lastSeen"`
	JSON           observabilitySharedQueryGetResponseCompareSeriesDataJSON    `json:"-"`
}

// observabilitySharedQueryGetResponseCompareSeriesDataJSON contains the JSON
// metadata for the struct [ObservabilitySharedQueryGetResponseCompareSeriesData]
type observabilitySharedQueryGetResponseCompareSeriesDataJSON struct {
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	FirstSeen      apijson.Field
	Groups         apijson.Field
	LastSeen       apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseCompareSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseCompareSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseCompareSeriesDataGroup struct {
	Key   string                                                               `json:"key" api:"required"`
	Value ObservabilitySharedQueryGetResponseCompareSeriesDataGroupsValueUnion `json:"value" api:"required"`
	JSON  observabilitySharedQueryGetResponseCompareSeriesDataGroupJSON        `json:"-"`
}

// observabilitySharedQueryGetResponseCompareSeriesDataGroupJSON contains the JSON
// metadata for the struct
// [ObservabilitySharedQueryGetResponseCompareSeriesDataGroup]
type observabilitySharedQueryGetResponseCompareSeriesDataGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseCompareSeriesDataGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseCompareSeriesDataGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilitySharedQueryGetResponseCompareSeriesDataGroupsValueUnion interface {
	ImplementsObservabilitySharedQueryGetResponseCompareSeriesDataGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilitySharedQueryGetResponseCompareSeriesDataGroupsValueUnion)(nil)).Elem(),
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

// Individual event results. Present when the query view is 'events'. Contains the
// matching log lines and their metadata.
type ObservabilitySharedQueryGetResponseEvents struct {
	// Total number of events matching the query (may exceed the number returned due to
	// limits).
	Count float64 `json:"count"`
	// List of individual telemetry events matching the query.
	Events []ObservabilitySharedQueryGetResponseEventsEvent `json:"events"`
	// List of fields discovered in the matched events. Useful for building dynamic
	// UIs.
	Fields []ObservabilitySharedQueryGetResponseEventsField `json:"fields"`
	// Time-series data for the matched events, bucketed by the query granularity.
	Series []ObservabilitySharedQueryGetResponseEventsSeries `json:"series"`
	JSON   observabilitySharedQueryGetResponseEventsJSON     `json:"-"`
}

// observabilitySharedQueryGetResponseEventsJSON contains the JSON metadata for the
// struct [ObservabilitySharedQueryGetResponseEvents]
type observabilitySharedQueryGetResponseEventsJSON struct {
	Count       apijson.Field
	Events      apijson.Field
	Fields      apijson.Field
	Series      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEvents) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEventsJSON) RawJSON() string {
	return r.raw
}

// A single telemetry event representing a log line, span, or metric data point
// emitted by a Worker.
type ObservabilitySharedQueryGetResponseEventsEvent struct {
	// Structured metadata extracted from the event. These fields are indexed and
	// available for filtering and aggregation.
	Metadata ObservabilitySharedQueryGetResponseEventsEventsMetadata `json:"$metadata" api:"required"`
	// The dataset this event belongs to (e.g. cloudflare-workers).
	Dataset string `json:"dataset" api:"required"`
	// Raw log payload. May be a string or a structured object depending on how the log
	// was emitted.
	Source ObservabilitySharedQueryGetResponseEventsEventsSourceUnion `json:"source" api:"required"`
	// Event timestamp as a Unix epoch in milliseconds.
	Timestamp int64 `json:"timestamp" api:"required"`
	// Cloudflare Containers event information that enriches your logs for identifying
	// and debugging issues.
	Containers map[string]interface{} `json:"$containers"`
	// Cloudflare Workers event information that enriches your logs for identifying and
	// debugging issues.
	Workers ObservabilitySharedQueryGetResponseEventsEventsWorkers `json:"$workers"`
	JSON    observabilitySharedQueryGetResponseEventsEventJSON     `json:"-"`
}

// observabilitySharedQueryGetResponseEventsEventJSON contains the JSON metadata
// for the struct [ObservabilitySharedQueryGetResponseEventsEvent]
type observabilitySharedQueryGetResponseEventsEventJSON struct {
	Metadata    apijson.Field
	Dataset     apijson.Field
	Source      apijson.Field
	Timestamp   apijson.Field
	Containers  apijson.Field
	Workers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEventsEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEventsEventJSON) RawJSON() string {
	return r.raw
}

// Structured metadata extracted from the event. These fields are indexed and
// available for filtering and aggregation.
type ObservabilitySharedQueryGetResponseEventsEventsMetadata struct {
	// Unique event ID. Use as the cursor value for offset-based pagination.
	ID string `json:"id" api:"required"`
	// Cloudflare account identifier.
	Account string `json:"account"`
	// Cloudflare product that generated this event (e.g. workers, pages).
	CloudService string `json:"cloudService"`
	ColdStart    int64  `json:"coldStart"`
	// Estimated cost units for this invocation.
	Cost int64 `json:"cost"`
	// Span duration in milliseconds.
	Duration int64 `json:"duration"`
	// Span end time as a Unix epoch in milliseconds.
	EndTime int64 `json:"endTime"`
	// Error message, present when the log represents an error.
	Error string `json:"error"`
	// Templatized version of the error message used for grouping similar errors.
	ErrorTemplate string `json:"errorTemplate"`
	// Content-based fingerprint used to group similar events.
	Fingerprint string `json:"fingerprint"`
	// Log level (e.g. log, debug, info, warn, error).
	Level string `json:"level"`
	// Log message text.
	Message string `json:"message"`
	// Templatized version of the log message used for grouping similar messages.
	MessageTemplate string `json:"messageTemplate"`
	// Metric name when the event represents a metric data point.
	MetricName string `json:"metricName"`
	// Origin of the event (e.g. fetch, scheduled, queue).
	Origin string `json:"origin"`
	// Span ID of the parent span in the trace hierarchy.
	ParentSpanID string `json:"parentSpanId"`
	// Infrastructure provider identifier.
	Provider string `json:"provider"`
	// Cloudflare data center / region that handled the request.
	Region string `json:"region"`
	// Cloudflare request ID that ties all logs from a single invocation together.
	RequestID string `json:"requestId"`
	// Worker script name that produced this event.
	Service string `json:"service"`
	// Span ID for this individual unit of work within a trace.
	SpanID string `json:"spanId"`
	// Human-readable name for this span.
	SpanName string `json:"spanName"`
	// Stack / deployment identifier.
	StackID string `json:"stackId"`
	// Span start time as a Unix epoch in milliseconds.
	StartTime int64 `json:"startTime"`
	// HTTP response status code returned by the Worker.
	StatusCode int64 `json:"statusCode"`
	// Total duration of the entire trace in milliseconds.
	TraceDuration int64 `json:"traceDuration"`
	// Distributed trace ID linking spans across services.
	TraceID string `json:"traceId"`
	// Logical transaction name for this request.
	TransactionName string `json:"transactionName"`
	// What triggered the invocation (e.g. GET /users, POST /orders, queue message).
	Trigger string `json:"trigger"`
	// Event type classifier (e.g. cf-worker-event, cf-worker-log).
	Type string `json:"type"`
	// Request URL that triggered the Worker invocation.
	URL  string                                                      `json:"url"`
	JSON observabilitySharedQueryGetResponseEventsEventsMetadataJSON `json:"-"`
}

// observabilitySharedQueryGetResponseEventsEventsMetadataJSON contains the JSON
// metadata for the struct
// [ObservabilitySharedQueryGetResponseEventsEventsMetadata]
type observabilitySharedQueryGetResponseEventsEventsMetadataJSON struct {
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
	TransactionName apijson.Field
	Trigger         apijson.Field
	Type            apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEventsEventsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEventsEventsMetadataJSON) RawJSON() string {
	return r.raw
}

// Raw log payload. May be a string or a structured object depending on how the log
// was emitted.
//
// Union satisfied by [shared.UnionString] or
// [ObservabilitySharedQueryGetResponseEventsEventsSourceMap].
type ObservabilitySharedQueryGetResponseEventsEventsSourceUnion interface {
	ImplementsObservabilitySharedQueryGetResponseEventsEventsSourceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilitySharedQueryGetResponseEventsEventsSourceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilitySharedQueryGetResponseEventsEventsSourceMap{}),
		},
	)
}

type ObservabilitySharedQueryGetResponseEventsEventsSourceMap map[string]interface{}

func (r ObservabilitySharedQueryGetResponseEventsEventsSourceMap) ImplementsObservabilitySharedQueryGetResponseEventsEventsSourceUnion() {
}

// Cloudflare Workers event information that enriches your logs for identifying and
// debugging issues.
type ObservabilitySharedQueryGetResponseEventsEventsWorkers struct {
	EventType  ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType `json:"eventType" api:"required"`
	RequestID  string                                                          `json:"requestId" api:"required"`
	ScriptName string                                                          `json:"scriptName" api:"required"`
	CPUTimeMs  float64                                                         `json:"cpuTimeMs"`
	// This field can have the runtime type of
	// [[]ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectDiagnosticsChannelEvent].
	DiagnosticsChannelEvents interface{} `json:"diagnosticsChannelEvents"`
	DispatchNamespace        string      `json:"dispatchNamespace"`
	DurableObjectID          string      `json:"durableObjectId"`
	Entrypoint               string      `json:"entrypoint"`
	// This field can have the runtime type of [map[string]interface{}].
	Event          interface{}                                                          `json:"event"`
	ExecutionModel ObservabilitySharedQueryGetResponseEventsEventsWorkersExecutionModel `json:"executionModel"`
	Outcome        string                                                               `json:"outcome"`
	// This field can have the runtime type of
	// [ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectPreview].
	Preview interface{} `json:"preview"`
	// This field can have the runtime type of
	// [ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectScriptVersion].
	ScriptVersion interface{}                                                `json:"scriptVersion"`
	SpanID        string                                                     `json:"spanId"`
	TraceID       string                                                     `json:"traceId"`
	Truncated     bool                                                       `json:"truncated"`
	WallTimeMs    float64                                                    `json:"wallTimeMs"`
	JSON          observabilitySharedQueryGetResponseEventsEventsWorkersJSON `json:"-"`
	union         ObservabilitySharedQueryGetResponseEventsEventsWorkersUnion
}

// observabilitySharedQueryGetResponseEventsEventsWorkersJSON contains the JSON
// metadata for the struct [ObservabilitySharedQueryGetResponseEventsEventsWorkers]
type observabilitySharedQueryGetResponseEventsEventsWorkersJSON struct {
	EventType                apijson.Field
	RequestID                apijson.Field
	ScriptName               apijson.Field
	CPUTimeMs                apijson.Field
	DiagnosticsChannelEvents apijson.Field
	DispatchNamespace        apijson.Field
	DurableObjectID          apijson.Field
	Entrypoint               apijson.Field
	Event                    apijson.Field
	ExecutionModel           apijson.Field
	Outcome                  apijson.Field
	Preview                  apijson.Field
	ScriptVersion            apijson.Field
	SpanID                   apijson.Field
	TraceID                  apijson.Field
	Truncated                apijson.Field
	WallTimeMs               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r observabilitySharedQueryGetResponseEventsEventsWorkersJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilitySharedQueryGetResponseEventsEventsWorkers) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilitySharedQueryGetResponseEventsEventsWorkers{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObservabilitySharedQueryGetResponseEventsEventsWorkersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ObservabilitySharedQueryGetResponseEventsEventsWorkersObject],
// [ObservabilitySharedQueryGetResponseEventsEventsWorkersObject].
func (r ObservabilitySharedQueryGetResponseEventsEventsWorkers) AsUnion() ObservabilitySharedQueryGetResponseEventsEventsWorkersUnion {
	return r.union
}

// Cloudflare Workers event information that enriches your logs for identifying and
// debugging issues.
//
// Union satisfied by
// [ObservabilitySharedQueryGetResponseEventsEventsWorkersObject] or
// [ObservabilitySharedQueryGetResponseEventsEventsWorkersObject].
type ObservabilitySharedQueryGetResponseEventsEventsWorkersUnion interface {
	implementsObservabilitySharedQueryGetResponseEventsEventsWorkers()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilitySharedQueryGetResponseEventsEventsWorkersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilitySharedQueryGetResponseEventsEventsWorkersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilitySharedQueryGetResponseEventsEventsWorkersObject{}),
		},
	)
}

type ObservabilitySharedQueryGetResponseEventsEventsWorkersObject struct {
	EventType       ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType      `json:"eventType" api:"required"`
	RequestID       string                                                                     `json:"requestId" api:"required"`
	ScriptName      string                                                                     `json:"scriptName" api:"required"`
	DurableObjectID string                                                                     `json:"durableObjectId"`
	Entrypoint      string                                                                     `json:"entrypoint"`
	Event           map[string]interface{}                                                     `json:"event"`
	ExecutionModel  ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectExecutionModel `json:"executionModel"`
	Outcome         string                                                                     `json:"outcome"`
	Preview         ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectPreview        `json:"preview"`
	ScriptVersion   ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectScriptVersion  `json:"scriptVersion"`
	SpanID          string                                                                     `json:"spanId"`
	TraceID         string                                                                     `json:"traceId"`
	Truncated       bool                                                                       `json:"truncated"`
	JSON            observabilitySharedQueryGetResponseEventsEventsWorkersObjectJSON           `json:"-"`
}

// observabilitySharedQueryGetResponseEventsEventsWorkersObjectJSON contains the
// JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseEventsEventsWorkersObject]
type observabilitySharedQueryGetResponseEventsEventsWorkersObjectJSON struct {
	EventType       apijson.Field
	RequestID       apijson.Field
	ScriptName      apijson.Field
	DurableObjectID apijson.Field
	Entrypoint      apijson.Field
	Event           apijson.Field
	ExecutionModel  apijson.Field
	Outcome         apijson.Field
	Preview         apijson.Field
	ScriptVersion   apijson.Field
	SpanID          apijson.Field
	TraceID         apijson.Field
	Truncated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEventsEventsWorkersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEventsEventsWorkersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilitySharedQueryGetResponseEventsEventsWorkersObject) implementsObservabilitySharedQueryGetResponseEventsEventsWorkers() {
}

type ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType string

const (
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeFetch     ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType = "fetch"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeScheduled ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType = "scheduled"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeAlarm     ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType = "alarm"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeCron      ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType = "cron"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeQueue     ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType = "queue"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeEmail     ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType = "email"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeTail      ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType = "tail"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeRpc       ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType = "rpc"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeWebsocket ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType = "websocket"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeWorkflow  ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType = "workflow"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeUnknown   ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType = "unknown"
)

func (r ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeFetch, ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeScheduled, ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeAlarm, ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeCron, ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeQueue, ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeEmail, ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeTail, ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeRpc, ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeWebsocket, ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeWorkflow, ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectExecutionModel string

const (
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectExecutionModelDurableObject ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectExecutionModel = "durableObject"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectExecutionModelStateless     ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectExecutionModel = "stateless"
)

func (r ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectExecutionModelDurableObject, ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectPreview struct {
	ID   string                                                                  `json:"id"`
	Name string                                                                  `json:"name"`
	Slug string                                                                  `json:"slug"`
	JSON observabilitySharedQueryGetResponseEventsEventsWorkersObjectPreviewJSON `json:"-"`
}

// observabilitySharedQueryGetResponseEventsEventsWorkersObjectPreviewJSON contains
// the JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectPreview]
type observabilitySharedQueryGetResponseEventsEventsWorkersObjectPreviewJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEventsEventsWorkersObjectPreviewJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectScriptVersion struct {
	ID      string                                                                        `json:"id"`
	Message string                                                                        `json:"message"`
	Tag     string                                                                        `json:"tag"`
	JSON    observabilitySharedQueryGetResponseEventsEventsWorkersObjectScriptVersionJSON `json:"-"`
}

// observabilitySharedQueryGetResponseEventsEventsWorkersObjectScriptVersionJSON
// contains the JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectScriptVersion]
type observabilitySharedQueryGetResponseEventsEventsWorkersObjectScriptVersionJSON struct {
	ID          apijson.Field
	Message     apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEventsEventsWorkersObjectScriptVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEventsEventsWorkersObjectScriptVersionJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType string

const (
	ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeFetch     ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType = "fetch"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeScheduled ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType = "scheduled"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeAlarm     ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType = "alarm"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeCron      ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType = "cron"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeQueue     ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType = "queue"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeEmail     ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType = "email"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeTail      ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType = "tail"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeRpc       ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType = "rpc"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeWebsocket ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType = "websocket"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeWorkflow  ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType = "workflow"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeUnknown   ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType = "unknown"
)

func (r ObservabilitySharedQueryGetResponseEventsEventsWorkersEventType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeFetch, ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeScheduled, ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeAlarm, ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeCron, ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeQueue, ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeEmail, ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeTail, ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeRpc, ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeWebsocket, ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeWorkflow, ObservabilitySharedQueryGetResponseEventsEventsWorkersEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseEventsEventsWorkersExecutionModel string

const (
	ObservabilitySharedQueryGetResponseEventsEventsWorkersExecutionModelDurableObject ObservabilitySharedQueryGetResponseEventsEventsWorkersExecutionModel = "durableObject"
	ObservabilitySharedQueryGetResponseEventsEventsWorkersExecutionModelStateless     ObservabilitySharedQueryGetResponseEventsEventsWorkersExecutionModel = "stateless"
)

func (r ObservabilitySharedQueryGetResponseEventsEventsWorkersExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseEventsEventsWorkersExecutionModelDurableObject, ObservabilitySharedQueryGetResponseEventsEventsWorkersExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseEventsField struct {
	// Field name present in the matched events.
	Key string `json:"key" api:"required"`
	// Data type of the field (string, number, or boolean).
	Type string                                             `json:"type" api:"required"`
	JSON observabilitySharedQueryGetResponseEventsFieldJSON `json:"-"`
}

// observabilitySharedQueryGetResponseEventsFieldJSON contains the JSON metadata
// for the struct [ObservabilitySharedQueryGetResponseEventsField]
type observabilitySharedQueryGetResponseEventsFieldJSON struct {
	Key         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEventsField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEventsFieldJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseEventsSeries struct {
	Data []ObservabilitySharedQueryGetResponseEventsSeriesData `json:"data" api:"required"`
	Time string                                                `json:"time" api:"required"`
	JSON observabilitySharedQueryGetResponseEventsSeriesJSON   `json:"-"`
}

// observabilitySharedQueryGetResponseEventsSeriesJSON contains the JSON metadata
// for the struct [ObservabilitySharedQueryGetResponseEventsSeries]
type observabilitySharedQueryGetResponseEventsSeriesJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEventsSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEventsSeriesJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseEventsSeriesData struct {
	Aggregates     ObservabilitySharedQueryGetResponseEventsSeriesDataAggregates `json:"aggregates" api:"required"`
	Count          float64                                                       `json:"count" api:"required"`
	Interval       float64                                                       `json:"interval" api:"required"`
	SampleInterval float64                                                       `json:"sampleInterval" api:"required"`
	Errors         float64                                                       `json:"errors"`
	// Groups in the query results.
	Groups map[string]ObservabilitySharedQueryGetResponseEventsSeriesDataGroupsUnion `json:"groups"`
	JSON   observabilitySharedQueryGetResponseEventsSeriesDataJSON                   `json:"-"`
}

// observabilitySharedQueryGetResponseEventsSeriesDataJSON contains the JSON
// metadata for the struct [ObservabilitySharedQueryGetResponseEventsSeriesData]
type observabilitySharedQueryGetResponseEventsSeriesDataJSON struct {
	Aggregates     apijson.Field
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Errors         apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEventsSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEventsSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseEventsSeriesDataAggregates struct {
	// Deprecated: deprecated
	Count int64 `json:"_count" api:"required"`
	// Deprecated: deprecated
	Interval float64 `json:"_interval" api:"required"`
	// Deprecated: deprecated
	FirstSeen string `json:"_firstSeen"`
	// Deprecated: deprecated
	LastSeen string `json:"_lastSeen"`
	// Deprecated: deprecated
	Bin  interface{}                                                       `json:"bin"`
	JSON observabilitySharedQueryGetResponseEventsSeriesDataAggregatesJSON `json:"-"`
}

// observabilitySharedQueryGetResponseEventsSeriesDataAggregatesJSON contains the
// JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseEventsSeriesDataAggregates]
type observabilitySharedQueryGetResponseEventsSeriesDataAggregatesJSON struct {
	Count       apijson.Field
	Interval    apijson.Field
	FirstSeen   apijson.Field
	LastSeen    apijson.Field
	Bin         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEventsSeriesDataAggregates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEventsSeriesDataAggregatesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilitySharedQueryGetResponseEventsSeriesDataGroupsUnion interface {
	ImplementsObservabilitySharedQueryGetResponseEventsSeriesDataGroupsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilitySharedQueryGetResponseEventsSeriesDataGroupsUnion)(nil)).Elem(),
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

// A single telemetry event representing a log line, span, or metric data point
// emitted by a Worker.
type ObservabilitySharedQueryGetResponseInvocation struct {
	// Structured metadata extracted from the event. These fields are indexed and
	// available for filtering and aggregation.
	Metadata ObservabilitySharedQueryGetResponseInvocationsMetadata `json:"$metadata" api:"required"`
	// The dataset this event belongs to (e.g. cloudflare-workers).
	Dataset string `json:"dataset" api:"required"`
	// Raw log payload. May be a string or a structured object depending on how the log
	// was emitted.
	Source ObservabilitySharedQueryGetResponseInvocationsSourceUnion `json:"source" api:"required"`
	// Event timestamp as a Unix epoch in milliseconds.
	Timestamp int64 `json:"timestamp" api:"required"`
	// Cloudflare Containers event information that enriches your logs for identifying
	// and debugging issues.
	Containers map[string]interface{} `json:"$containers"`
	// Cloudflare Workers event information that enriches your logs for identifying and
	// debugging issues.
	Workers ObservabilitySharedQueryGetResponseInvocationsWorkers `json:"$workers"`
	JSON    observabilitySharedQueryGetResponseInvocationJSON     `json:"-"`
}

// observabilitySharedQueryGetResponseInvocationJSON contains the JSON metadata for
// the struct [ObservabilitySharedQueryGetResponseInvocation]
type observabilitySharedQueryGetResponseInvocationJSON struct {
	Metadata    apijson.Field
	Dataset     apijson.Field
	Source      apijson.Field
	Timestamp   apijson.Field
	Containers  apijson.Field
	Workers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseInvocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseInvocationJSON) RawJSON() string {
	return r.raw
}

// Structured metadata extracted from the event. These fields are indexed and
// available for filtering and aggregation.
type ObservabilitySharedQueryGetResponseInvocationsMetadata struct {
	// Unique event ID. Use as the cursor value for offset-based pagination.
	ID string `json:"id" api:"required"`
	// Cloudflare account identifier.
	Account string `json:"account"`
	// Cloudflare product that generated this event (e.g. workers, pages).
	CloudService string `json:"cloudService"`
	ColdStart    int64  `json:"coldStart"`
	// Estimated cost units for this invocation.
	Cost int64 `json:"cost"`
	// Span duration in milliseconds.
	Duration int64 `json:"duration"`
	// Span end time as a Unix epoch in milliseconds.
	EndTime int64 `json:"endTime"`
	// Error message, present when the log represents an error.
	Error string `json:"error"`
	// Templatized version of the error message used for grouping similar errors.
	ErrorTemplate string `json:"errorTemplate"`
	// Content-based fingerprint used to group similar events.
	Fingerprint string `json:"fingerprint"`
	// Log level (e.g. log, debug, info, warn, error).
	Level string `json:"level"`
	// Log message text.
	Message string `json:"message"`
	// Templatized version of the log message used for grouping similar messages.
	MessageTemplate string `json:"messageTemplate"`
	// Metric name when the event represents a metric data point.
	MetricName string `json:"metricName"`
	// Origin of the event (e.g. fetch, scheduled, queue).
	Origin string `json:"origin"`
	// Span ID of the parent span in the trace hierarchy.
	ParentSpanID string `json:"parentSpanId"`
	// Infrastructure provider identifier.
	Provider string `json:"provider"`
	// Cloudflare data center / region that handled the request.
	Region string `json:"region"`
	// Cloudflare request ID that ties all logs from a single invocation together.
	RequestID string `json:"requestId"`
	// Worker script name that produced this event.
	Service string `json:"service"`
	// Span ID for this individual unit of work within a trace.
	SpanID string `json:"spanId"`
	// Human-readable name for this span.
	SpanName string `json:"spanName"`
	// Stack / deployment identifier.
	StackID string `json:"stackId"`
	// Span start time as a Unix epoch in milliseconds.
	StartTime int64 `json:"startTime"`
	// HTTP response status code returned by the Worker.
	StatusCode int64 `json:"statusCode"`
	// Total duration of the entire trace in milliseconds.
	TraceDuration int64 `json:"traceDuration"`
	// Distributed trace ID linking spans across services.
	TraceID string `json:"traceId"`
	// Logical transaction name for this request.
	TransactionName string `json:"transactionName"`
	// What triggered the invocation (e.g. GET /users, POST /orders, queue message).
	Trigger string `json:"trigger"`
	// Event type classifier (e.g. cf-worker-event, cf-worker-log).
	Type string `json:"type"`
	// Request URL that triggered the Worker invocation.
	URL  string                                                     `json:"url"`
	JSON observabilitySharedQueryGetResponseInvocationsMetadataJSON `json:"-"`
}

// observabilitySharedQueryGetResponseInvocationsMetadataJSON contains the JSON
// metadata for the struct [ObservabilitySharedQueryGetResponseInvocationsMetadata]
type observabilitySharedQueryGetResponseInvocationsMetadataJSON struct {
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
	TransactionName apijson.Field
	Trigger         apijson.Field
	Type            apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseInvocationsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseInvocationsMetadataJSON) RawJSON() string {
	return r.raw
}

// Raw log payload. May be a string or a structured object depending on how the log
// was emitted.
//
// Union satisfied by [shared.UnionString] or
// [ObservabilitySharedQueryGetResponseInvocationsSourceMap].
type ObservabilitySharedQueryGetResponseInvocationsSourceUnion interface {
	ImplementsObservabilitySharedQueryGetResponseInvocationsSourceUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilitySharedQueryGetResponseInvocationsSourceUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilitySharedQueryGetResponseInvocationsSourceMap{}),
		},
	)
}

type ObservabilitySharedQueryGetResponseInvocationsSourceMap map[string]interface{}

func (r ObservabilitySharedQueryGetResponseInvocationsSourceMap) ImplementsObservabilitySharedQueryGetResponseInvocationsSourceUnion() {
}

// Cloudflare Workers event information that enriches your logs for identifying and
// debugging issues.
type ObservabilitySharedQueryGetResponseInvocationsWorkers struct {
	EventType  ObservabilitySharedQueryGetResponseInvocationsWorkersEventType `json:"eventType" api:"required"`
	RequestID  string                                                         `json:"requestId" api:"required"`
	ScriptName string                                                         `json:"scriptName" api:"required"`
	CPUTimeMs  float64                                                        `json:"cpuTimeMs"`
	// This field can have the runtime type of
	// [[]ObservabilitySharedQueryGetResponseInvocationsWorkersObjectDiagnosticsChannelEvent].
	DiagnosticsChannelEvents interface{} `json:"diagnosticsChannelEvents"`
	DispatchNamespace        string      `json:"dispatchNamespace"`
	DurableObjectID          string      `json:"durableObjectId"`
	Entrypoint               string      `json:"entrypoint"`
	// This field can have the runtime type of [map[string]interface{}].
	Event          interface{}                                                         `json:"event"`
	ExecutionModel ObservabilitySharedQueryGetResponseInvocationsWorkersExecutionModel `json:"executionModel"`
	Outcome        string                                                              `json:"outcome"`
	// This field can have the runtime type of
	// [ObservabilitySharedQueryGetResponseInvocationsWorkersObjectPreview].
	Preview interface{} `json:"preview"`
	// This field can have the runtime type of
	// [ObservabilitySharedQueryGetResponseInvocationsWorkersObjectScriptVersion].
	ScriptVersion interface{}                                               `json:"scriptVersion"`
	SpanID        string                                                    `json:"spanId"`
	TraceID       string                                                    `json:"traceId"`
	Truncated     bool                                                      `json:"truncated"`
	WallTimeMs    float64                                                   `json:"wallTimeMs"`
	JSON          observabilitySharedQueryGetResponseInvocationsWorkersJSON `json:"-"`
	union         ObservabilitySharedQueryGetResponseInvocationsWorkersUnion
}

// observabilitySharedQueryGetResponseInvocationsWorkersJSON contains the JSON
// metadata for the struct [ObservabilitySharedQueryGetResponseInvocationsWorkers]
type observabilitySharedQueryGetResponseInvocationsWorkersJSON struct {
	EventType                apijson.Field
	RequestID                apijson.Field
	ScriptName               apijson.Field
	CPUTimeMs                apijson.Field
	DiagnosticsChannelEvents apijson.Field
	DispatchNamespace        apijson.Field
	DurableObjectID          apijson.Field
	Entrypoint               apijson.Field
	Event                    apijson.Field
	ExecutionModel           apijson.Field
	Outcome                  apijson.Field
	Preview                  apijson.Field
	ScriptVersion            apijson.Field
	SpanID                   apijson.Field
	TraceID                  apijson.Field
	Truncated                apijson.Field
	WallTimeMs               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r observabilitySharedQueryGetResponseInvocationsWorkersJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilitySharedQueryGetResponseInvocationsWorkers) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilitySharedQueryGetResponseInvocationsWorkers{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObservabilitySharedQueryGetResponseInvocationsWorkersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ObservabilitySharedQueryGetResponseInvocationsWorkersObject],
// [ObservabilitySharedQueryGetResponseInvocationsWorkersObject].
func (r ObservabilitySharedQueryGetResponseInvocationsWorkers) AsUnion() ObservabilitySharedQueryGetResponseInvocationsWorkersUnion {
	return r.union
}

// Cloudflare Workers event information that enriches your logs for identifying and
// debugging issues.
//
// Union satisfied by [ObservabilitySharedQueryGetResponseInvocationsWorkersObject]
// or [ObservabilitySharedQueryGetResponseInvocationsWorkersObject].
type ObservabilitySharedQueryGetResponseInvocationsWorkersUnion interface {
	implementsObservabilitySharedQueryGetResponseInvocationsWorkers()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilitySharedQueryGetResponseInvocationsWorkersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilitySharedQueryGetResponseInvocationsWorkersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilitySharedQueryGetResponseInvocationsWorkersObject{}),
		},
	)
}

type ObservabilitySharedQueryGetResponseInvocationsWorkersObject struct {
	EventType       ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType      `json:"eventType" api:"required"`
	RequestID       string                                                                    `json:"requestId" api:"required"`
	ScriptName      string                                                                    `json:"scriptName" api:"required"`
	DurableObjectID string                                                                    `json:"durableObjectId"`
	Entrypoint      string                                                                    `json:"entrypoint"`
	Event           map[string]interface{}                                                    `json:"event"`
	ExecutionModel  ObservabilitySharedQueryGetResponseInvocationsWorkersObjectExecutionModel `json:"executionModel"`
	Outcome         string                                                                    `json:"outcome"`
	Preview         ObservabilitySharedQueryGetResponseInvocationsWorkersObjectPreview        `json:"preview"`
	ScriptVersion   ObservabilitySharedQueryGetResponseInvocationsWorkersObjectScriptVersion  `json:"scriptVersion"`
	SpanID          string                                                                    `json:"spanId"`
	TraceID         string                                                                    `json:"traceId"`
	Truncated       bool                                                                      `json:"truncated"`
	JSON            observabilitySharedQueryGetResponseInvocationsWorkersObjectJSON           `json:"-"`
}

// observabilitySharedQueryGetResponseInvocationsWorkersObjectJSON contains the
// JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseInvocationsWorkersObject]
type observabilitySharedQueryGetResponseInvocationsWorkersObjectJSON struct {
	EventType       apijson.Field
	RequestID       apijson.Field
	ScriptName      apijson.Field
	DurableObjectID apijson.Field
	Entrypoint      apijson.Field
	Event           apijson.Field
	ExecutionModel  apijson.Field
	Outcome         apijson.Field
	Preview         apijson.Field
	ScriptVersion   apijson.Field
	SpanID          apijson.Field
	TraceID         apijson.Field
	Truncated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseInvocationsWorkersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseInvocationsWorkersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilitySharedQueryGetResponseInvocationsWorkersObject) implementsObservabilitySharedQueryGetResponseInvocationsWorkers() {
}

type ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType string

const (
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeFetch     ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType = "fetch"
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeScheduled ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType = "scheduled"
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeAlarm     ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType = "alarm"
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeCron      ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType = "cron"
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeQueue     ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType = "queue"
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeEmail     ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType = "email"
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeTail      ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType = "tail"
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeRpc       ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType = "rpc"
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeWebsocket ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType = "websocket"
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeWorkflow  ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType = "workflow"
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeUnknown   ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType = "unknown"
)

func (r ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeFetch, ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeScheduled, ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeAlarm, ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeCron, ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeQueue, ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeEmail, ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeTail, ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeRpc, ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeWebsocket, ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeWorkflow, ObservabilitySharedQueryGetResponseInvocationsWorkersObjectEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseInvocationsWorkersObjectExecutionModel string

const (
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectExecutionModelDurableObject ObservabilitySharedQueryGetResponseInvocationsWorkersObjectExecutionModel = "durableObject"
	ObservabilitySharedQueryGetResponseInvocationsWorkersObjectExecutionModelStateless     ObservabilitySharedQueryGetResponseInvocationsWorkersObjectExecutionModel = "stateless"
)

func (r ObservabilitySharedQueryGetResponseInvocationsWorkersObjectExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseInvocationsWorkersObjectExecutionModelDurableObject, ObservabilitySharedQueryGetResponseInvocationsWorkersObjectExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseInvocationsWorkersObjectPreview struct {
	ID   string                                                                 `json:"id"`
	Name string                                                                 `json:"name"`
	Slug string                                                                 `json:"slug"`
	JSON observabilitySharedQueryGetResponseInvocationsWorkersObjectPreviewJSON `json:"-"`
}

// observabilitySharedQueryGetResponseInvocationsWorkersObjectPreviewJSON contains
// the JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseInvocationsWorkersObjectPreview]
type observabilitySharedQueryGetResponseInvocationsWorkersObjectPreviewJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseInvocationsWorkersObjectPreview) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseInvocationsWorkersObjectPreviewJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseInvocationsWorkersObjectScriptVersion struct {
	ID      string                                                                       `json:"id"`
	Message string                                                                       `json:"message"`
	Tag     string                                                                       `json:"tag"`
	JSON    observabilitySharedQueryGetResponseInvocationsWorkersObjectScriptVersionJSON `json:"-"`
}

// observabilitySharedQueryGetResponseInvocationsWorkersObjectScriptVersionJSON
// contains the JSON metadata for the struct
// [ObservabilitySharedQueryGetResponseInvocationsWorkersObjectScriptVersion]
type observabilitySharedQueryGetResponseInvocationsWorkersObjectScriptVersionJSON struct {
	ID          apijson.Field
	Message     apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseInvocationsWorkersObjectScriptVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseInvocationsWorkersObjectScriptVersionJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseInvocationsWorkersEventType string

const (
	ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeFetch     ObservabilitySharedQueryGetResponseInvocationsWorkersEventType = "fetch"
	ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeScheduled ObservabilitySharedQueryGetResponseInvocationsWorkersEventType = "scheduled"
	ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeAlarm     ObservabilitySharedQueryGetResponseInvocationsWorkersEventType = "alarm"
	ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeCron      ObservabilitySharedQueryGetResponseInvocationsWorkersEventType = "cron"
	ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeQueue     ObservabilitySharedQueryGetResponseInvocationsWorkersEventType = "queue"
	ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeEmail     ObservabilitySharedQueryGetResponseInvocationsWorkersEventType = "email"
	ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeTail      ObservabilitySharedQueryGetResponseInvocationsWorkersEventType = "tail"
	ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeRpc       ObservabilitySharedQueryGetResponseInvocationsWorkersEventType = "rpc"
	ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeWebsocket ObservabilitySharedQueryGetResponseInvocationsWorkersEventType = "websocket"
	ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeWorkflow  ObservabilitySharedQueryGetResponseInvocationsWorkersEventType = "workflow"
	ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeUnknown   ObservabilitySharedQueryGetResponseInvocationsWorkersEventType = "unknown"
)

func (r ObservabilitySharedQueryGetResponseInvocationsWorkersEventType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeFetch, ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeScheduled, ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeAlarm, ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeCron, ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeQueue, ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeEmail, ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeTail, ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeRpc, ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeWebsocket, ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeWorkflow, ObservabilitySharedQueryGetResponseInvocationsWorkersEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseInvocationsWorkersExecutionModel string

const (
	ObservabilitySharedQueryGetResponseInvocationsWorkersExecutionModelDurableObject ObservabilitySharedQueryGetResponseInvocationsWorkersExecutionModel = "durableObject"
	ObservabilitySharedQueryGetResponseInvocationsWorkersExecutionModelStateless     ObservabilitySharedQueryGetResponseInvocationsWorkersExecutionModel = "stateless"
)

func (r ObservabilitySharedQueryGetResponseInvocationsWorkersExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseInvocationsWorkersExecutionModelDurableObject, ObservabilitySharedQueryGetResponseInvocationsWorkersExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseTrace struct {
	// Name of the root span that initiated the trace.
	RootSpanName string `json:"rootSpanName" api:"required"`
	// Logical transaction name for the root span.
	RootTransactionName string `json:"rootTransactionName" api:"required"`
	// List of Worker services involved in the trace.
	Service []string `json:"service" api:"required"`
	// Total number of spans in the trace.
	Spans float64 `json:"spans" api:"required"`
	// Total duration of the trace in milliseconds.
	TraceDurationMs float64 `json:"traceDurationMs" api:"required"`
	// Trace end time as a Unix epoch in milliseconds.
	TraceEndMs float64 `json:"traceEndMs" api:"required"`
	// Unique identifier for the distributed trace.
	TraceID string `json:"traceId" api:"required"`
	// Trace start time as a Unix epoch in milliseconds.
	TraceStartMs float64 `json:"traceStartMs" api:"required"`
	// Error messages encountered during the trace, if any.
	Errors []string                                     `json:"errors"`
	JSON   observabilitySharedQueryGetResponseTraceJSON `json:"-"`
}

// observabilitySharedQueryGetResponseTraceJSON contains the JSON metadata for the
// struct [ObservabilitySharedQueryGetResponseTrace]
type observabilitySharedQueryGetResponseTraceJSON struct {
	RootSpanName        apijson.Field
	RootTransactionName apijson.Field
	Service             apijson.Field
	Spans               apijson.Field
	TraceDurationMs     apijson.Field
	TraceEndMs          apijson.Field
	TraceID             apijson.Field
	TraceStartMs        apijson.Field
	Errors              apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseTraceJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryNewParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Identifier for the query. When parameters are omitted, this ID is used to load a
	// previously saved query's parameters. When providing parameters inline, pass any
	// identifier (e.g. an ad-hoc ID).
	QueryID param.Field[string] `json:"queryId" api:"required"`
	// Timeframe for the query using Unix timestamps in milliseconds. Narrower
	// timeframes produce faster responses and more specific results.
	Timeframe param.Field[ObservabilitySharedQueryNewParamsTimeframe] `json:"timeframe" api:"required"`
	// When true, includes time-series data in the response.
	Chart param.Field[bool] `json:"chart"`
	// When true, includes a comparison dataset from the previous time period of equal
	// length.
	Compare param.Field[bool] `json:"compare"`
	// When true, executes the query without persisting the results. Useful for
	// validation or previewing.
	Dry param.Field[bool] `json:"dry"`
	// Number of time-series buckets. Only used when view is 'calculations'. Omit to
	// let the system auto-detect an appropriate granularity.
	Granularity param.Field[float64] `json:"granularity"`
	// When true, omits time-series data from the response and returns only aggregated
	// values. Reduces response size when series are not needed.
	IgnoreSeries param.Field[bool] `json:"ignoreSeries"`
	// Maximum number of events to return when view is 'events'. Also controls the
	// number of group-by rows when view is 'calculations'.
	Limit param.Field[float64] `json:"limit"`
	// Cursor for pagination in event, trace, and invocation views. Pass the
	// $metadata.id of the last returned item to fetch the next page.
	Offset param.Field[string] `json:"offset"`
	// Numeric offset for paginating grouped/pattern results (top-N lists). Use
	// together with limit. Not used by cursor-based pagination.
	OffsetBy param.Field[float64] `json:"offsetBy"`
	// Pagination direction: 'next' for forward, 'prev' for backward.
	OffsetDirection param.Field[string] `json:"offsetDirection"`
	// Query parameters defining what data to retrieve — filters, calculations,
	// group-bys, and ordering. In practice this should always be provided for ad-hoc
	// queries. Only omit when executing a previously saved query by queryId. Use the
	// keys and values endpoints to discover available fields before building filters.
	Parameters param.Field[ObservabilitySharedQueryNewParamsParameters] `json:"parameters"`
	// Controls the shape of the response. 'events': individual log lines matching the
	// query. 'calculations': aggregated metrics (count, avg, p99, etc.) with optional
	// group-by breakdowns and time-series. 'invocations': events grouped by request
	// ID. 'traces': distributed trace summaries. 'agents': Durable Object agent
	// summaries.
	View param.Field[ObservabilitySharedQueryNewParamsView] `json:"view"`
}

func (r ObservabilitySharedQueryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Timeframe for the query using Unix timestamps in milliseconds. Narrower
// timeframes produce faster responses and more specific results.
type ObservabilitySharedQueryNewParamsTimeframe struct {
	// Start timestamp for the query timeframe (Unix timestamp in milliseconds)
	From param.Field[float64] `json:"from" api:"required"`
	// End timestamp for the query timeframe (Unix timestamp in milliseconds)
	To param.Field[float64] `json:"to" api:"required"`
}

func (r ObservabilitySharedQueryNewParamsTimeframe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Query parameters defining what data to retrieve — filters, calculations,
// group-bys, and ordering. In practice this should always be provided for ad-hoc
// queries. Only omit when executing a previously saved query by queryId. Use the
// keys and values endpoints to discover available fields before building filters.
type ObservabilitySharedQueryNewParamsParameters struct {
	// Aggregation calculations to compute (e.g. count, avg, p99). Each calculation
	// produces aggregate values and optional time-series data.
	Calculations param.Field[[]ObservabilitySharedQueryNewParamsParametersCalculation] `json:"calculations"`
	// Datasets to query. Leave empty to query all available datasets.
	Datasets param.Field[[]string] `json:"datasets"`
	// Logical operator for combining top-level filters: 'and' (all must match) or 'or'
	// (any must match). Defaults to 'and'.
	FilterCombination param.Field[ObservabilitySharedQueryNewParamsParametersFilterCombination] `json:"filterCombination"`
	// Filters to narrow query results. Use the keys and values endpoints to discover
	// available fields before building filters. Supports nested groups via kind:
	// 'group'. Maximum nesting depth is 4.
	Filters param.Field[[]ObservabilitySharedQueryNewParamsParametersFilterUnion] `json:"filters"`
	// Fields to group calculation results by. Only applicable when the query view is
	// 'calculations'. Produces per-group aggregate values.
	GroupBys param.Field[[]ObservabilitySharedQueryNewParamsParametersGroupBy] `json:"groupBys"`
	// Post-aggregation filters applied to calculation results. Use to filter groups
	// after aggregation (e.g. only groups where count > 100).
	Havings param.Field[[]ObservabilitySharedQueryNewParamsParametersHaving] `json:"havings"`
	// Maximum number of group-by rows to return in calculation results. A value of 10
	// is a sensible default for most use cases.
	Limit param.Field[int64] `json:"limit"`
	// Full-text search expression applied across all event fields. Matches events
	// containing the specified text.
	Needle param.Field[ObservabilitySharedQueryNewParamsParametersNeedle] `json:"needle"`
	// Ordering for grouped calculation results. Only effective when a group-by is
	// present.
	OrderBy param.Field[ObservabilitySharedQueryNewParamsParametersOrderBy] `json:"orderBy"`
}

func (r ObservabilitySharedQueryNewParamsParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilitySharedQueryNewParamsParametersCalculation struct {
	// Aggregation operator to apply. Examples: count, avg, sum, min, max, median, p90,
	// p95, p99, uniq, stddev, variance.
	Operator param.Field[ObservabilitySharedQueryNewParamsParametersCalculationsOperator] `json:"operator" api:"required"`
	// Custom label for this calculation in the results. Useful for distinguishing
	// multiple calculations.
	Alias param.Field[string] `json:"alias"`
	// Field name to calculate over. Must exist in the data — verify with the keys
	// endpoint. Omit for operators that don't require a key (e.g. count).
	Key param.Field[string] `json:"key"`
	// Data type of the key. Required when key is provided to ensure correct
	// aggregation.
	KeyType param.Field[ObservabilitySharedQueryNewParamsParametersCalculationsKeyType] `json:"keyType"`
}

func (r ObservabilitySharedQueryNewParamsParametersCalculation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Aggregation operator to apply. Examples: count, avg, sum, min, max, median, p90,
// p95, p99, uniq, stddev, variance.
type ObservabilitySharedQueryNewParamsParametersCalculationsOperator string

const (
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorUniq              ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "uniq"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorCount             ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "count"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorMax               ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "max"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorMin               ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "min"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorSum               ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "sum"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorAvg               ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "avg"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorMedian            ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "median"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP001              ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "p001"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP01               ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "p01"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP05               ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "p05"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP10               ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "p10"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP25               ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "p25"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP75               ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "p75"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP90               ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "p90"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP95               ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "p95"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP99               ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "p99"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP999              ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "p999"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorStddev            ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "stddev"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorVariance          ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "variance"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorCountDistinct     ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "COUNT_DISTINCT"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorCountUppercase    ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "COUNT"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorMaxUppercase      ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "MAX"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorMinUppercase      ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "MIN"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorSumUppercase      ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "SUM"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorAvgUppercase      ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "AVG"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorMedianUppercase   ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "MEDIAN"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP001Uppercase     ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "P001"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP01Uppercase      ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "P01"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP05Uppercase      ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "P05"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP10Uppercase      ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "P10"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP25Uppercase      ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "P25"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP75Uppercase      ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "P75"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP90Uppercase      ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "P90"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP95Uppercase      ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "P95"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP99Uppercase      ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "P99"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP999Uppercase     ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "P999"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorStddevUppercase   ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "STDDEV"
	ObservabilitySharedQueryNewParamsParametersCalculationsOperatorVarianceUppercase ObservabilitySharedQueryNewParamsParametersCalculationsOperator = "VARIANCE"
)

func (r ObservabilitySharedQueryNewParamsParametersCalculationsOperator) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersCalculationsOperatorUniq, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorCount, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorMax, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorMin, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorSum, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorAvg, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorMedian, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP001, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP01, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP05, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP10, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP25, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP75, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP90, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP95, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP99, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP999, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorStddev, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorVariance, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorCountDistinct, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorCountUppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorMaxUppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorMinUppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorSumUppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorAvgUppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorMedianUppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP001Uppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP01Uppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP05Uppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP10Uppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP25Uppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP75Uppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP90Uppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP95Uppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP99Uppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorP999Uppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorStddevUppercase, ObservabilitySharedQueryNewParamsParametersCalculationsOperatorVarianceUppercase:
		return true
	}
	return false
}

// Data type of the key. Required when key is provided to ensure correct
// aggregation.
type ObservabilitySharedQueryNewParamsParametersCalculationsKeyType string

const (
	ObservabilitySharedQueryNewParamsParametersCalculationsKeyTypeString  ObservabilitySharedQueryNewParamsParametersCalculationsKeyType = "string"
	ObservabilitySharedQueryNewParamsParametersCalculationsKeyTypeNumber  ObservabilitySharedQueryNewParamsParametersCalculationsKeyType = "number"
	ObservabilitySharedQueryNewParamsParametersCalculationsKeyTypeBoolean ObservabilitySharedQueryNewParamsParametersCalculationsKeyType = "boolean"
)

func (r ObservabilitySharedQueryNewParamsParametersCalculationsKeyType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersCalculationsKeyTypeString, ObservabilitySharedQueryNewParamsParametersCalculationsKeyTypeNumber, ObservabilitySharedQueryNewParamsParametersCalculationsKeyTypeBoolean:
		return true
	}
	return false
}

// Logical operator for combining top-level filters: 'and' (all must match) or 'or'
// (any must match). Defaults to 'and'.
type ObservabilitySharedQueryNewParamsParametersFilterCombination string

const (
	ObservabilitySharedQueryNewParamsParametersFilterCombinationAnd          ObservabilitySharedQueryNewParamsParametersFilterCombination = "and"
	ObservabilitySharedQueryNewParamsParametersFilterCombinationOr           ObservabilitySharedQueryNewParamsParametersFilterCombination = "or"
	ObservabilitySharedQueryNewParamsParametersFilterCombinationAndUppercase ObservabilitySharedQueryNewParamsParametersFilterCombination = "AND"
	ObservabilitySharedQueryNewParamsParametersFilterCombinationOrUppercase  ObservabilitySharedQueryNewParamsParametersFilterCombination = "OR"
)

func (r ObservabilitySharedQueryNewParamsParametersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFilterCombinationAnd, ObservabilitySharedQueryNewParamsParametersFilterCombinationOr, ObservabilitySharedQueryNewParamsParametersFilterCombinationAndUppercase, ObservabilitySharedQueryNewParamsParametersFilterCombinationOrUppercase:
		return true
	}
	return false
}

// Supports nested groups via kind: 'group'.
type ObservabilitySharedQueryNewParamsParametersFilter struct {
	FilterCombination param.Field[ObservabilitySharedQueryNewParamsParametersFiltersFilterCombination] `json:"filterCombination"`
	Filters           param.Field[interface{}]                                                         `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  param.Field[string]                                                 `json:"key"`
	Kind param.Field[ObservabilitySharedQueryNewParamsParametersFiltersKind] `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
	// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilitySharedQueryNewParamsParametersFiltersOperation] `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type  param.Field[ObservabilitySharedQueryNewParamsParametersFiltersType] `json:"type"`
	Value param.Field[interface{}]                                            `json:"value"`
}

func (r ObservabilitySharedQueryNewParamsParametersFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilitySharedQueryNewParamsParametersFilter) implementsObservabilitySharedQueryNewParamsParametersFilterUnion() {
}

// Supports nested groups via kind: 'group'.
//
// Satisfied by [workers.ObservabilitySharedQueryNewParamsParametersFiltersObject],
// [workers.ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeaf],
// [ObservabilitySharedQueryNewParamsParametersFilter].
type ObservabilitySharedQueryNewParamsParametersFilterUnion interface {
	implementsObservabilitySharedQueryNewParamsParametersFilterUnion()
}

type ObservabilitySharedQueryNewParamsParametersFiltersObject struct {
	FilterCombination param.Field[ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombination] `json:"filterCombination" api:"required"`
	Filters           param.Field[[]ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterUnion]     `json:"filters" api:"required"`
	Kind              param.Field[ObservabilitySharedQueryNewParamsParametersFiltersObjectKind]              `json:"kind" api:"required"`
}

func (r ObservabilitySharedQueryNewParamsParametersFiltersObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilitySharedQueryNewParamsParametersFiltersObject) implementsObservabilitySharedQueryNewParamsParametersFilterUnion() {
}

type ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombination string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombinationAnd          ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombination = "and"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombinationOr           ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombination = "or"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombinationAndUppercase ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombination = "AND"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombinationOrUppercase  ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombination = "OR"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombinationAnd, ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombinationOr, ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombinationAndUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

// Supports nested groups via kind: 'group'.
type ObservabilitySharedQueryNewParamsParametersFiltersObjectFilter struct {
	FilterCombination param.Field[ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombination] `json:"filterCombination"`
	Filters           param.Field[interface{}]                                                                      `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  param.Field[string]                                                              `json:"key"`
	Kind param.Field[ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersKind] `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
	// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation] `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type  param.Field[ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersType] `json:"type"`
	Value param.Field[interface{}]                                                         `json:"value"`
}

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFilter) implementsObservabilitySharedQueryNewParamsParametersFiltersObjectFilterUnion() {
}

// Supports nested groups via kind: 'group'.
//
// Satisfied by
// [workers.ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObject],
// [workers.ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeaf],
// [ObservabilitySharedQueryNewParamsParametersFiltersObjectFilter].
type ObservabilitySharedQueryNewParamsParametersFiltersObjectFilterUnion interface {
	implementsObservabilitySharedQueryNewParamsParametersFiltersObjectFilterUnion()
}

type ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObject struct {
	FilterCombination param.Field[ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombination] `json:"filterCombination" api:"required"`
	Filters           param.Field[[]interface{}]                                                                          `json:"filters" api:"required"`
	Kind              param.Field[ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectKind]              `json:"kind" api:"required"`
}

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObject) implementsObservabilitySharedQueryNewParamsParametersFiltersObjectFilterUnion() {
}

type ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombination string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombinationAnd          ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombination = "and"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombinationOr           ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombination = "or"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombinationAndUppercase ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombination = "AND"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombinationOrUppercase  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombination = "OR"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombinationAnd, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombinationOr, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombinationAndUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectKind string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectKindGroup ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectKind = "group"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key param.Field[string] `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
	// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation] `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type param.Field[ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafType] `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind param.Field[ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafKind] `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value param.Field[ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion] `json:"value"`
}

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeaf) implementsObservabilitySharedQueryNewParamsParametersFiltersObjectFilterUnion() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEndsWith            ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "ends_with"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEndsWithUppercase   ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "ENDS_WITH"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEndsWith, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEndsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafTypeString  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafTypeString, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafKindFilter ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafKindFilter:
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
type ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion()
}

type ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombination string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombinationAnd          ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombination = "and"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombinationOr           ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombination = "or"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombinationAndUppercase ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombination = "AND"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombinationOrUppercase  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombination = "OR"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombinationAnd, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombinationOr, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombinationAndUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersKind string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersKindGroup  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersKind = "group"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersKindFilter ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersKind = "filter"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersKindGroup, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationIncludes            ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "includes"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationNotIncludes         ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "not_includes"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationStartsWith          ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "starts_with"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationEndsWith            ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "ends_with"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationRegex               ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "regex"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationExists              ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "exists"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationIsNull              ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "is_null"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationIn                  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "in"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationNotIn               ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "not_in"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationEq                  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "eq"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationNeq                 ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "neq"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationGt                  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "gt"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationGte                 ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "gte"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationLt                  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "lt"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationLte                 ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "lte"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationEquals              ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "="
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationNotEquals           ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "!="
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationGreater             ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = ">"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationGreaterOrEquals     ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = ">="
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationLess                ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "<"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationLessOrEquals        ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "<="
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationIncludesUppercase   ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "INCLUDES"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationDoesNotInclude      ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationMatchRegex          ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "MATCH_REGEX"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationExistsUppercase     ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "EXISTS"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationDoesNotExist        ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "DOES_NOT_EXIST"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationInUppercase         ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "IN"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationNotInUppercase      ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "NOT_IN"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationStartsWithUppercase ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "STARTS_WITH"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationEndsWithUppercase   ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation = "ENDS_WITH"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationIncludes, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationNotIncludes, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationStartsWith, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationEndsWith, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationRegex, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationExists, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationIsNull, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationIn, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationNotIn, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationEq, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationNeq, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationGt, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationGte, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationLt, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationLte, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationEquals, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationNotEquals, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationGreater, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationGreaterOrEquals, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationLess, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationLessOrEquals, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationIncludesUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationDoesNotInclude, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationMatchRegex, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationExistsUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationDoesNotExist, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationInUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationNotInUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationStartsWithUppercase, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersOperationEndsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersType string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersTypeString  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersType = "string"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersTypeNumber  ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersType = "number"
	ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersTypeBoolean ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersType = "boolean"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersTypeString, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersTypeNumber, ObservabilitySharedQueryNewParamsParametersFiltersObjectFiltersTypeBoolean:
		return true
	}
	return false
}

type ObservabilitySharedQueryNewParamsParametersFiltersObjectKind string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersObjectKindGroup ObservabilitySharedQueryNewParamsParametersFiltersObjectKind = "group"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key param.Field[string] `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
	// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation] `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type param.Field[ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafType] `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind param.Field[ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafKind] `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value param.Field[ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion] `json:"value"`
}

func (r ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeaf) implementsObservabilitySharedQueryNewParamsParametersFilterUnion() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEndsWith            ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "ends_with"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEndsWithUppercase   ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "ENDS_WITH"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEndsWith, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafOperationEndsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafTypeString  ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafTypeString, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafKindFilter ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafKindFilter:
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
type ObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilitySharedQueryNewParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion()
}

type ObservabilitySharedQueryNewParamsParametersFiltersFilterCombination string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersFilterCombinationAnd          ObservabilitySharedQueryNewParamsParametersFiltersFilterCombination = "and"
	ObservabilitySharedQueryNewParamsParametersFiltersFilterCombinationOr           ObservabilitySharedQueryNewParamsParametersFiltersFilterCombination = "or"
	ObservabilitySharedQueryNewParamsParametersFiltersFilterCombinationAndUppercase ObservabilitySharedQueryNewParamsParametersFiltersFilterCombination = "AND"
	ObservabilitySharedQueryNewParamsParametersFiltersFilterCombinationOrUppercase  ObservabilitySharedQueryNewParamsParametersFiltersFilterCombination = "OR"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersFilterCombinationAnd, ObservabilitySharedQueryNewParamsParametersFiltersFilterCombinationOr, ObservabilitySharedQueryNewParamsParametersFiltersFilterCombinationAndUppercase, ObservabilitySharedQueryNewParamsParametersFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilitySharedQueryNewParamsParametersFiltersKind string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersKindGroup  ObservabilitySharedQueryNewParamsParametersFiltersKind = "group"
	ObservabilitySharedQueryNewParamsParametersFiltersKindFilter ObservabilitySharedQueryNewParamsParametersFiltersKind = "filter"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersKindGroup, ObservabilitySharedQueryNewParamsParametersFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// ends_with, regex. Existence: exists, is_null. Set membership: in, not_in
// (comma-separated values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilitySharedQueryNewParamsParametersFiltersOperation string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersOperationIncludes            ObservabilitySharedQueryNewParamsParametersFiltersOperation = "includes"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationNotIncludes         ObservabilitySharedQueryNewParamsParametersFiltersOperation = "not_includes"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationStartsWith          ObservabilitySharedQueryNewParamsParametersFiltersOperation = "starts_with"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationEndsWith            ObservabilitySharedQueryNewParamsParametersFiltersOperation = "ends_with"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationRegex               ObservabilitySharedQueryNewParamsParametersFiltersOperation = "regex"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationExists              ObservabilitySharedQueryNewParamsParametersFiltersOperation = "exists"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationIsNull              ObservabilitySharedQueryNewParamsParametersFiltersOperation = "is_null"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationIn                  ObservabilitySharedQueryNewParamsParametersFiltersOperation = "in"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationNotIn               ObservabilitySharedQueryNewParamsParametersFiltersOperation = "not_in"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationEq                  ObservabilitySharedQueryNewParamsParametersFiltersOperation = "eq"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationNeq                 ObservabilitySharedQueryNewParamsParametersFiltersOperation = "neq"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationGt                  ObservabilitySharedQueryNewParamsParametersFiltersOperation = "gt"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationGte                 ObservabilitySharedQueryNewParamsParametersFiltersOperation = "gte"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationLt                  ObservabilitySharedQueryNewParamsParametersFiltersOperation = "lt"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationLte                 ObservabilitySharedQueryNewParamsParametersFiltersOperation = "lte"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationEquals              ObservabilitySharedQueryNewParamsParametersFiltersOperation = "="
	ObservabilitySharedQueryNewParamsParametersFiltersOperationNotEquals           ObservabilitySharedQueryNewParamsParametersFiltersOperation = "!="
	ObservabilitySharedQueryNewParamsParametersFiltersOperationGreater             ObservabilitySharedQueryNewParamsParametersFiltersOperation = ">"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationGreaterOrEquals     ObservabilitySharedQueryNewParamsParametersFiltersOperation = ">="
	ObservabilitySharedQueryNewParamsParametersFiltersOperationLess                ObservabilitySharedQueryNewParamsParametersFiltersOperation = "<"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationLessOrEquals        ObservabilitySharedQueryNewParamsParametersFiltersOperation = "<="
	ObservabilitySharedQueryNewParamsParametersFiltersOperationIncludesUppercase   ObservabilitySharedQueryNewParamsParametersFiltersOperation = "INCLUDES"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationDoesNotInclude      ObservabilitySharedQueryNewParamsParametersFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationMatchRegex          ObservabilitySharedQueryNewParamsParametersFiltersOperation = "MATCH_REGEX"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationExistsUppercase     ObservabilitySharedQueryNewParamsParametersFiltersOperation = "EXISTS"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationDoesNotExist        ObservabilitySharedQueryNewParamsParametersFiltersOperation = "DOES_NOT_EXIST"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationInUppercase         ObservabilitySharedQueryNewParamsParametersFiltersOperation = "IN"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationNotInUppercase      ObservabilitySharedQueryNewParamsParametersFiltersOperation = "NOT_IN"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationStartsWithUppercase ObservabilitySharedQueryNewParamsParametersFiltersOperation = "STARTS_WITH"
	ObservabilitySharedQueryNewParamsParametersFiltersOperationEndsWithUppercase   ObservabilitySharedQueryNewParamsParametersFiltersOperation = "ENDS_WITH"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersOperationIncludes, ObservabilitySharedQueryNewParamsParametersFiltersOperationNotIncludes, ObservabilitySharedQueryNewParamsParametersFiltersOperationStartsWith, ObservabilitySharedQueryNewParamsParametersFiltersOperationEndsWith, ObservabilitySharedQueryNewParamsParametersFiltersOperationRegex, ObservabilitySharedQueryNewParamsParametersFiltersOperationExists, ObservabilitySharedQueryNewParamsParametersFiltersOperationIsNull, ObservabilitySharedQueryNewParamsParametersFiltersOperationIn, ObservabilitySharedQueryNewParamsParametersFiltersOperationNotIn, ObservabilitySharedQueryNewParamsParametersFiltersOperationEq, ObservabilitySharedQueryNewParamsParametersFiltersOperationNeq, ObservabilitySharedQueryNewParamsParametersFiltersOperationGt, ObservabilitySharedQueryNewParamsParametersFiltersOperationGte, ObservabilitySharedQueryNewParamsParametersFiltersOperationLt, ObservabilitySharedQueryNewParamsParametersFiltersOperationLte, ObservabilitySharedQueryNewParamsParametersFiltersOperationEquals, ObservabilitySharedQueryNewParamsParametersFiltersOperationNotEquals, ObservabilitySharedQueryNewParamsParametersFiltersOperationGreater, ObservabilitySharedQueryNewParamsParametersFiltersOperationGreaterOrEquals, ObservabilitySharedQueryNewParamsParametersFiltersOperationLess, ObservabilitySharedQueryNewParamsParametersFiltersOperationLessOrEquals, ObservabilitySharedQueryNewParamsParametersFiltersOperationIncludesUppercase, ObservabilitySharedQueryNewParamsParametersFiltersOperationDoesNotInclude, ObservabilitySharedQueryNewParamsParametersFiltersOperationMatchRegex, ObservabilitySharedQueryNewParamsParametersFiltersOperationExistsUppercase, ObservabilitySharedQueryNewParamsParametersFiltersOperationDoesNotExist, ObservabilitySharedQueryNewParamsParametersFiltersOperationInUppercase, ObservabilitySharedQueryNewParamsParametersFiltersOperationNotInUppercase, ObservabilitySharedQueryNewParamsParametersFiltersOperationStartsWithUppercase, ObservabilitySharedQueryNewParamsParametersFiltersOperationEndsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilitySharedQueryNewParamsParametersFiltersType string

const (
	ObservabilitySharedQueryNewParamsParametersFiltersTypeString  ObservabilitySharedQueryNewParamsParametersFiltersType = "string"
	ObservabilitySharedQueryNewParamsParametersFiltersTypeNumber  ObservabilitySharedQueryNewParamsParametersFiltersType = "number"
	ObservabilitySharedQueryNewParamsParametersFiltersTypeBoolean ObservabilitySharedQueryNewParamsParametersFiltersType = "boolean"
)

func (r ObservabilitySharedQueryNewParamsParametersFiltersType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersFiltersTypeString, ObservabilitySharedQueryNewParamsParametersFiltersTypeNumber, ObservabilitySharedQueryNewParamsParametersFiltersTypeBoolean:
		return true
	}
	return false
}

type ObservabilitySharedQueryNewParamsParametersGroupBy struct {
	// Data type of the group-by field.
	Type param.Field[ObservabilitySharedQueryNewParamsParametersGroupBysType] `json:"type" api:"required"`
	// Field name to group results by (e.g. $metadata.service, $metadata.statusCode).
	Value param.Field[string] `json:"value" api:"required"`
}

func (r ObservabilitySharedQueryNewParamsParametersGroupBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data type of the group-by field.
type ObservabilitySharedQueryNewParamsParametersGroupBysType string

const (
	ObservabilitySharedQueryNewParamsParametersGroupBysTypeString  ObservabilitySharedQueryNewParamsParametersGroupBysType = "string"
	ObservabilitySharedQueryNewParamsParametersGroupBysTypeNumber  ObservabilitySharedQueryNewParamsParametersGroupBysType = "number"
	ObservabilitySharedQueryNewParamsParametersGroupBysTypeBoolean ObservabilitySharedQueryNewParamsParametersGroupBysType = "boolean"
)

func (r ObservabilitySharedQueryNewParamsParametersGroupBysType) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersGroupBysTypeString, ObservabilitySharedQueryNewParamsParametersGroupBysTypeNumber, ObservabilitySharedQueryNewParamsParametersGroupBysTypeBoolean:
		return true
	}
	return false
}

type ObservabilitySharedQueryNewParamsParametersHaving struct {
	// Calculation alias or operator to filter on after aggregation.
	Key param.Field[string] `json:"key" api:"required"`
	// Numeric comparison operator: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilitySharedQueryNewParamsParametersHavingsOperation] `json:"operation" api:"required"`
	// Threshold value to compare the calculation result against.
	Value param.Field[float64] `json:"value" api:"required"`
}

func (r ObservabilitySharedQueryNewParamsParametersHaving) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Numeric comparison operator: eq, neq, gt, gte, lt, lte.
type ObservabilitySharedQueryNewParamsParametersHavingsOperation string

const (
	ObservabilitySharedQueryNewParamsParametersHavingsOperationEq  ObservabilitySharedQueryNewParamsParametersHavingsOperation = "eq"
	ObservabilitySharedQueryNewParamsParametersHavingsOperationNeq ObservabilitySharedQueryNewParamsParametersHavingsOperation = "neq"
	ObservabilitySharedQueryNewParamsParametersHavingsOperationGt  ObservabilitySharedQueryNewParamsParametersHavingsOperation = "gt"
	ObservabilitySharedQueryNewParamsParametersHavingsOperationGte ObservabilitySharedQueryNewParamsParametersHavingsOperation = "gte"
	ObservabilitySharedQueryNewParamsParametersHavingsOperationLt  ObservabilitySharedQueryNewParamsParametersHavingsOperation = "lt"
	ObservabilitySharedQueryNewParamsParametersHavingsOperationLte ObservabilitySharedQueryNewParamsParametersHavingsOperation = "lte"
)

func (r ObservabilitySharedQueryNewParamsParametersHavingsOperation) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersHavingsOperationEq, ObservabilitySharedQueryNewParamsParametersHavingsOperationNeq, ObservabilitySharedQueryNewParamsParametersHavingsOperationGt, ObservabilitySharedQueryNewParamsParametersHavingsOperationGte, ObservabilitySharedQueryNewParamsParametersHavingsOperationLt, ObservabilitySharedQueryNewParamsParametersHavingsOperationLte:
		return true
	}
	return false
}

// Full-text search expression applied across all event fields. Matches events
// containing the specified text.
type ObservabilitySharedQueryNewParamsParametersNeedle struct {
	// The text or pattern to search for.
	Value param.Field[ObservabilitySharedQueryNewParamsParametersNeedleValueUnion] `json:"value" api:"required"`
	// When true, treats the value as a regular expression (RE2 syntax).
	IsRegex param.Field[bool] `json:"isRegex"`
	// When true, performs a case-sensitive search. Defaults to case-insensitive.
	MatchCase param.Field[bool] `json:"matchCase"`
}

func (r ObservabilitySharedQueryNewParamsParametersNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The text or pattern to search for.
//
// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilitySharedQueryNewParamsParametersNeedleValueUnion interface {
	ImplementsObservabilitySharedQueryNewParamsParametersNeedleValueUnion()
}

// Ordering for grouped calculation results. Only effective when a group-by is
// present.
type ObservabilitySharedQueryNewParamsParametersOrderBy struct {
	// Alias of the calculation to order results by. Must match the alias (or operator)
	// of a calculation in the query.
	Value param.Field[string] `json:"value" api:"required"`
	// Sort direction: 'asc' for ascending, 'desc' for descending.
	Order param.Field[ObservabilitySharedQueryNewParamsParametersOrderByOrder] `json:"order"`
}

func (r ObservabilitySharedQueryNewParamsParametersOrderBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort direction: 'asc' for ascending, 'desc' for descending.
type ObservabilitySharedQueryNewParamsParametersOrderByOrder string

const (
	ObservabilitySharedQueryNewParamsParametersOrderByOrderAsc  ObservabilitySharedQueryNewParamsParametersOrderByOrder = "asc"
	ObservabilitySharedQueryNewParamsParametersOrderByOrderDesc ObservabilitySharedQueryNewParamsParametersOrderByOrder = "desc"
)

func (r ObservabilitySharedQueryNewParamsParametersOrderByOrder) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsParametersOrderByOrderAsc, ObservabilitySharedQueryNewParamsParametersOrderByOrderDesc:
		return true
	}
	return false
}

// Controls the shape of the response. 'events': individual log lines matching the
// query. 'calculations': aggregated metrics (count, avg, p99, etc.) with optional
// group-by breakdowns and time-series. 'invocations': events grouped by request
// ID. 'traces': distributed trace summaries. 'agents': Durable Object agent
// summaries.
type ObservabilitySharedQueryNewParamsView string

const (
	ObservabilitySharedQueryNewParamsViewTraces       ObservabilitySharedQueryNewParamsView = "traces"
	ObservabilitySharedQueryNewParamsViewEvents       ObservabilitySharedQueryNewParamsView = "events"
	ObservabilitySharedQueryNewParamsViewCalculations ObservabilitySharedQueryNewParamsView = "calculations"
	ObservabilitySharedQueryNewParamsViewInvocations  ObservabilitySharedQueryNewParamsView = "invocations"
	ObservabilitySharedQueryNewParamsViewRequests     ObservabilitySharedQueryNewParamsView = "requests"
	ObservabilitySharedQueryNewParamsViewAgents       ObservabilitySharedQueryNewParamsView = "agents"
)

func (r ObservabilitySharedQueryNewParamsView) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewParamsViewTraces, ObservabilitySharedQueryNewParamsViewEvents, ObservabilitySharedQueryNewParamsViewCalculations, ObservabilitySharedQueryNewParamsViewInvocations, ObservabilitySharedQueryNewParamsViewRequests, ObservabilitySharedQueryNewParamsViewAgents:
		return true
	}
	return false
}

type ObservabilitySharedQueryNewResponseEnvelope struct {
	Errors   []ObservabilitySharedQueryNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilitySharedQueryNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilitySharedQueryNewResponse                   `json:"result" api:"required"`
	Success  ObservabilitySharedQueryNewResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilitySharedQueryNewResponseEnvelopeJSON       `json:"-"`
}

// observabilitySharedQueryNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ObservabilitySharedQueryNewResponseEnvelope]
type observabilitySharedQueryNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryNewResponseEnvelopeErrors struct {
	Message string                                                `json:"message" api:"required"`
	JSON    observabilitySharedQueryNewResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilitySharedQueryNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ObservabilitySharedQueryNewResponseEnvelopeErrors]
type observabilitySharedQueryNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryNewResponseEnvelopeMessages struct {
	Message ObservabilitySharedQueryNewResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilitySharedQueryNewResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilitySharedQueryNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilitySharedQueryNewResponseEnvelopeMessages]
type observabilitySharedQueryNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryNewResponseEnvelopeMessagesMessage string

const (
	ObservabilitySharedQueryNewResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilitySharedQueryNewResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilitySharedQueryNewResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilitySharedQueryNewResponseEnvelopeSuccess bool

const (
	ObservabilitySharedQueryNewResponseEnvelopeSuccessTrue ObservabilitySharedQueryNewResponseEnvelopeSuccess = true
)

func (r ObservabilitySharedQueryNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Select the view of the query result to return, defaults to events.
	View param.Field[ObservabilitySharedQueryGetParamsView] `query:"view"`
}

// URLQuery serializes [ObservabilitySharedQueryGetParams]'s query parameters as
// `url.Values`.
func (r ObservabilitySharedQueryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Select the view of the query result to return, defaults to events.
type ObservabilitySharedQueryGetParamsView string

const (
	ObservabilitySharedQueryGetParamsViewEvents       ObservabilitySharedQueryGetParamsView = "events"
	ObservabilitySharedQueryGetParamsViewInvocations  ObservabilitySharedQueryGetParamsView = "invocations"
	ObservabilitySharedQueryGetParamsViewCalculations ObservabilitySharedQueryGetParamsView = "calculations"
)

func (r ObservabilitySharedQueryGetParamsView) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetParamsViewEvents, ObservabilitySharedQueryGetParamsViewInvocations, ObservabilitySharedQueryGetParamsViewCalculations:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseEnvelope struct {
	Errors   []ObservabilitySharedQueryGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilitySharedQueryGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Complete results of a query run. The populated fields depend on the requested
	// view type (events, calculations, invocations, traces, or agents).
	Result  ObservabilitySharedQueryGetResponse                `json:"result" api:"required"`
	Success ObservabilitySharedQueryGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    observabilitySharedQueryGetResponseEnvelopeJSON    `json:"-"`
}

// observabilitySharedQueryGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ObservabilitySharedQueryGetResponseEnvelope]
type observabilitySharedQueryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseEnvelopeErrors struct {
	Message string                                                `json:"message" api:"required"`
	JSON    observabilitySharedQueryGetResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilitySharedQueryGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ObservabilitySharedQueryGetResponseEnvelopeErrors]
type observabilitySharedQueryGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseEnvelopeMessages struct {
	Message ObservabilitySharedQueryGetResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilitySharedQueryGetResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilitySharedQueryGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilitySharedQueryGetResponseEnvelopeMessages]
type observabilitySharedQueryGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilitySharedQueryGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilitySharedQueryGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilitySharedQueryGetResponseEnvelopeMessagesMessage string

const (
	ObservabilitySharedQueryGetResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilitySharedQueryGetResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilitySharedQueryGetResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilitySharedQueryGetResponseEnvelopeSuccess bool

const (
	ObservabilitySharedQueryGetResponseEnvelopeSuccessTrue ObservabilitySharedQueryGetResponseEnvelopeSuccess = true
)

func (r ObservabilitySharedQueryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilitySharedQueryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
