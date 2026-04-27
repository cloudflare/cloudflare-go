// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v6/shared"
	"github.com/tidwall/gjson"
)

// ObservabilityTelemetryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityTelemetryService] method instead.
type ObservabilityTelemetryService struct {
	Options []option.RequestOption
}

// NewObservabilityTelemetryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewObservabilityTelemetryService(opts ...option.RequestOption) (r *ObservabilityTelemetryService) {
	r = &ObservabilityTelemetryService{}
	r.Options = opts
	return
}

// List all the keys in your telemetry events.
func (r *ObservabilityTelemetryService) Keys(ctx context.Context, params ObservabilityTelemetryKeysParams, opts ...option.RequestOption) (res *pagination.SinglePage[ObservabilityTelemetryKeysResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/telemetry/keys", params.AccountID)
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

// List all the keys in your telemetry events.
func (r *ObservabilityTelemetryService) KeysAutoPaging(ctx context.Context, params ObservabilityTelemetryKeysParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ObservabilityTelemetryKeysResponse] {
	return pagination.NewSinglePageAutoPager(r.Keys(ctx, params, opts...))
}

// Run a temporary or saved query.
func (r *ObservabilityTelemetryService) Query(ctx context.Context, params ObservabilityTelemetryQueryParams, opts ...option.RequestOption) (res *ObservabilityTelemetryQueryResponse, err error) {
	var env ObservabilityTelemetryQueryResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/telemetry/query", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List unique values found in your events.
func (r *ObservabilityTelemetryService) Values(ctx context.Context, params ObservabilityTelemetryValuesParams, opts ...option.RequestOption) (res *pagination.SinglePage[ObservabilityTelemetryValuesResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/telemetry/values", params.AccountID)
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

// List unique values found in your events.
func (r *ObservabilityTelemetryService) ValuesAutoPaging(ctx context.Context, params ObservabilityTelemetryValuesParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ObservabilityTelemetryValuesResponse] {
	return pagination.NewSinglePageAutoPager(r.Values(ctx, params, opts...))
}

type ObservabilityTelemetryKeysResponse struct {
	Key        string                                 `json:"key" api:"required"`
	LastSeenAt float64                                `json:"lastSeenAt" api:"required"`
	Type       ObservabilityTelemetryKeysResponseType `json:"type" api:"required"`
	JSON       observabilityTelemetryKeysResponseJSON `json:"-"`
}

// observabilityTelemetryKeysResponseJSON contains the JSON metadata for the struct
// [ObservabilityTelemetryKeysResponse]
type observabilityTelemetryKeysResponseJSON struct {
	Key         apijson.Field
	LastSeenAt  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryKeysResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryKeysResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryKeysResponseType string

const (
	ObservabilityTelemetryKeysResponseTypeString  ObservabilityTelemetryKeysResponseType = "string"
	ObservabilityTelemetryKeysResponseTypeBoolean ObservabilityTelemetryKeysResponseType = "boolean"
	ObservabilityTelemetryKeysResponseTypeNumber  ObservabilityTelemetryKeysResponseType = "number"
)

func (r ObservabilityTelemetryKeysResponseType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysResponseTypeString, ObservabilityTelemetryKeysResponseTypeBoolean, ObservabilityTelemetryKeysResponseTypeNumber:
		return true
	}
	return false
}

// Complete results of a query run. The populated fields depend on the requested
// view type (events, calculations, invocations, traces, or agents).
type ObservabilityTelemetryQueryResponse struct {
	// The query run metadata including the query definition, execution status, and
	// timeframe.
	Run ObservabilityTelemetryQueryResponseRun `json:"run" api:"required"`
	// Query performance statistics from the database. Includes execution time, rows
	// scanned, and bytes read. Does not include network latency.
	Statistics ObservabilityTelemetryQueryResponseStatistics `json:"statistics" api:"required"`
	// Durable Object agent summaries. Present when the query view is 'agents'. Each
	// entry represents an agent with its event counts and status.
	Agents []ObservabilityTelemetryQueryResponseAgent `json:"agents"`
	// Aggregated calculation results. Present when the query view is 'calculations'.
	// Contains computed metrics (count, avg, p99, etc.) with optional group-by
	// breakdowns and time-series data.
	Calculations []ObservabilityTelemetryQueryResponseCalculation `json:"calculations"`
	// Comparison calculation results from the previous time period. Present when the
	// compare option is enabled. Same structure as calculations.
	Compare []ObservabilityTelemetryQueryResponseCompare `json:"compare"`
	// Individual event results. Present when the query view is 'events'. Contains the
	// matching log lines and their metadata.
	Events ObservabilityTelemetryQueryResponseEvents `json:"events"`
	// Events grouped by invocation (request ID). Present when the query view is
	// 'invocations'. Each key is a request ID mapping to all events from that
	// invocation.
	Invocations map[string][]ObservabilityTelemetryQueryResponseInvocation `json:"invocations"`
	// Trace summaries matching the query. Present when the query view is 'traces'.
	// Each entry represents a distributed trace with its spans, duration, and services
	// involved.
	Traces []ObservabilityTelemetryQueryResponseTrace `json:"traces"`
	JSON   observabilityTelemetryQueryResponseJSON    `json:"-"`
}

// observabilityTelemetryQueryResponseJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryQueryResponse]
type observabilityTelemetryQueryResponseJSON struct {
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

func (r *ObservabilityTelemetryQueryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseJSON) RawJSON() string {
	return r.raw
}

// The query run metadata including the query definition, execution status, and
// timeframe.
type ObservabilityTelemetryQueryResponseRun struct {
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
	Query ObservabilityTelemetryQueryResponseRunQuery `json:"query" api:"required"`
	// Current execution status of the query run.
	Status ObservabilityTelemetryQueryResponseRunStatus `json:"status" api:"required"`
	// Time range for the query execution
	Timeframe ObservabilityTelemetryQueryResponseRunTimeframe `json:"timeframe" api:"required"`
	// ID of the user who initiated the query run.
	UserID string `json:"userId" api:"required"`
	// ISO-8601 timestamp when the query run was created.
	Created string `json:"created"`
	// Query performance statistics from the database (does not include network
	// latency).
	Statistics ObservabilityTelemetryQueryResponseRunStatistics `json:"statistics"`
	// ISO-8601 timestamp when the query run was last updated.
	Updated string                                     `json:"updated"`
	JSON    observabilityTelemetryQueryResponseRunJSON `json:"-"`
}

// observabilityTelemetryQueryResponseRunJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryQueryResponseRun]
type observabilityTelemetryQueryResponseRunJSON struct {
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

func (r *ObservabilityTelemetryQueryResponseRun) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunJSON) RawJSON() string {
	return r.raw
}

// A saved query definition with its parameters, metadata, and ownership
// information.
type ObservabilityTelemetryQueryResponseRunQuery struct {
	ID string `json:"id" api:"required"`
	// If the query wasn't explcitly saved
	Adhoc       bool   `json:"adhoc" api:"required"`
	Created     string `json:"created" api:"required"`
	CreatedBy   string `json:"createdBy" api:"required"`
	Description string `json:"description" api:"required,nullable"`
	// Query name
	Name       string                                                `json:"name" api:"required"`
	Parameters ObservabilityTelemetryQueryResponseRunQueryParameters `json:"parameters" api:"required"`
	Updated    string                                                `json:"updated" api:"required"`
	UpdatedBy  string                                                `json:"updatedBy" api:"required"`
	JSON       observabilityTelemetryQueryResponseRunQueryJSON       `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryResponseRunQuery]
type observabilityTelemetryQueryResponseRunQueryJSON struct {
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

func (r *ObservabilityTelemetryQueryResponseRunQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQueryParameters struct {
	// Create Calculations to compute as part of the query.
	Calculations []ObservabilityTelemetryQueryResponseRunQueryParametersCalculation `json:"calculations"`
	// Set the Datasets to query. Leave it empty to query all the datasets.
	Datasets []string `json:"datasets"`
	// Set a Flag to describe how to combine the filters on the query.
	FilterCombination ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination `json:"filterCombination"`
	// Configure the Filters to apply to the query. Supports nested groups via kind:
	// 'group'.
	Filters []ObservabilityTelemetryQueryResponseRunQueryParametersFilter `json:"filters"`
	// Define how to group the results of the query.
	GroupBys []ObservabilityTelemetryQueryResponseRunQueryParametersGroupBy `json:"groupBys"`
	// Configure the Having clauses that filter on calculations in the query result.
	Havings []ObservabilityTelemetryQueryResponseRunQueryParametersHaving `json:"havings"`
	// Set a limit on the number of results / records returned by the query
	Limit int64 `json:"limit"`
	// Define an expression to search using full-text search.
	Needle ObservabilityTelemetryQueryResponseRunQueryParametersNeedle `json:"needle"`
	// Configure the order of the results returned by the query.
	OrderBy ObservabilityTelemetryQueryResponseRunQueryParametersOrderBy `json:"orderBy"`
	JSON    observabilityTelemetryQueryResponseRunQueryParametersJSON    `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseRunQueryParameters]
type observabilityTelemetryQueryResponseRunQueryParametersJSON struct {
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

func (r *ObservabilityTelemetryQueryResponseRunQueryParameters) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQueryParametersCalculation struct {
	Operator ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator `json:"operator" api:"required"`
	Alias    string                                                                    `json:"alias"`
	Key      string                                                                    `json:"key"`
	KeyType  ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyType  `json:"keyType"`
	JSON     observabilityTelemetryQueryResponseRunQueryParametersCalculationJSON      `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersCalculationJSON contains
// the JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersCalculation]
type observabilityTelemetryQueryResponseRunQueryParametersCalculationJSON struct {
	Operator    apijson.Field
	Alias       apijson.Field
	Key         apijson.Field
	KeyType     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersCalculation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersCalculationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorUniq              ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "uniq"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorCount             ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "count"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMax               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "max"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMin               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "min"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorSum               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "sum"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorAvg               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "avg"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMedian            ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "median"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP001              ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p001"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP01               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p01"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP05               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p05"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP10               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p10"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP25               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p25"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP75               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p75"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP90               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p90"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP95               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p95"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP99               ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p99"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP999              ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "p999"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorStddev            ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "stddev"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorVariance          ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "variance"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorCountDistinct     ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "COUNT_DISTINCT"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorCountUppercase    ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "COUNT"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMaxUppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "MAX"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMinUppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "MIN"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorSumUppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "SUM"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorAvgUppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "AVG"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMedianUppercase   ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "MEDIAN"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP001Uppercase     ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P001"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP01Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P01"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP05Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P05"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP10Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P10"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP25Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P25"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP75Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P75"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP90Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P90"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP95Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P95"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP99Uppercase      ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P99"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP999Uppercase     ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "P999"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorStddevUppercase   ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "STDDEV"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorVarianceUppercase ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator = "VARIANCE"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperator) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorUniq, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorCount, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMax, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMin, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorSum, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorAvg, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMedian, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP001, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP01, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP05, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP10, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP25, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP75, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP90, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP95, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP99, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP999, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorStddev, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorVariance, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorCountDistinct, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorCountUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMaxUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMinUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorSumUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorAvgUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorMedianUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP001Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP01Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP05Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP10Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP25Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP75Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP90Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP95Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP99Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorP999Uppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorStddevUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsOperatorVarianceUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyType string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyTypeString  ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyType = "string"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyTypeNumber  ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyType = "number"
	ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyTypeBoolean ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyType = "boolean"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyTypeString, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyTypeNumber, ObservabilityTelemetryQueryResponseRunQueryParametersCalculationsKeyTypeBoolean:
		return true
	}
	return false
}

// Set a Flag to describe how to combine the filters on the query.
type ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationAnd          ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination = "and"
	ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationOr           ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination = "or"
	ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationAndUppercase ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination = "AND"
	ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationOrUppercase  ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination = "OR"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationAnd, ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationOr, ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationAndUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFilterCombinationOrUppercase:
		return true
	}
	return false
}

// Supports nested groups via kind: 'group'.
type ObservabilityTelemetryQueryResponseRunQueryParametersFilter struct {
	FilterCombination ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombination `json:"filterCombination"`
	// This field can have the runtime type of [[]interface{}].
	Filters interface{} `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  string                                                           `json:"key"`
	Kind ObservabilityTelemetryQueryResponseRunQueryParametersFiltersKind `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType `json:"type"`
	// This field can have the runtime type of
	// [ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion].
	Value interface{}                                                     `json:"value"`
	JSON  observabilityTelemetryQueryResponseRunQueryParametersFilterJSON `json:"-"`
	union ObservabilityTelemetryQueryResponseRunQueryParametersFiltersUnion
}

// observabilityTelemetryQueryResponseRunQueryParametersFilterJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersFilter]
type observabilityTelemetryQueryResponseRunQueryParametersFilterJSON struct {
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

func (r observabilityTelemetryQueryResponseRunQueryParametersFilterJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersFilter) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilityTelemetryQueryResponseRunQueryParametersFilter{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a
// [ObservabilityTelemetryQueryResponseRunQueryParametersFiltersUnion] interface
// which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObject],
// [ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf].
func (r ObservabilityTelemetryQueryResponseRunQueryParametersFilter) AsUnion() ObservabilityTelemetryQueryResponseRunQueryParametersFiltersUnion {
	return r.union
}

// Supports nested groups via kind: 'group'.
//
// Union satisfied by
// [ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObject] or
// [ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf].
type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersUnion interface {
	implementsObservabilityTelemetryQueryResponseRunQueryParametersFilter()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseRunQueryParametersFiltersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf{}),
		},
	)
}

type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObject struct {
	FilterCombination ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombination `json:"filterCombination" api:"required"`
	Filters           []interface{}                                                                       `json:"filters" api:"required"`
	Kind              ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectKind              `json:"kind" api:"required"`
	JSON              observabilityTelemetryQueryResponseRunQueryParametersFiltersObjectJSON              `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersFiltersObjectJSON contains
// the JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObject]
type observabilityTelemetryQueryResponseRunQueryParametersFiltersObjectJSON struct {
	FilterCombination apijson.Field
	Filters           apijson.Field
	Kind              apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersFiltersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObject) implementsObservabilityTelemetryQueryResponseRunQueryParametersFilter() {
}

type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombination string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombinationAnd          ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombination = "and"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombinationOr           ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombination = "or"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombinationAndUppercase ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombination = "AND"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombinationOrUppercase  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombination = "OR"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombinationAnd, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombinationOr, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombinationAndUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectKind string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectKindGroup ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectKind = "group"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key string `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKind `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion `json:"value"`
	JSON  observabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafJSON       `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafJSON
// contains the JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf]
type observabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafJSON struct {
	Key         apijson.Field
	Operation   apijson.Field
	Type        apijson.Field
	Kind        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf) implementsObservabilityTelemetryQueryResponseRunQueryParametersFilter() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafTypeString  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafTypeString, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKindFilter ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKindFilter:
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
type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafValueUnion)(nil)).Elem(),
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

type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombination string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombinationAnd          ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombination = "and"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombinationOr           ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombination = "or"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombinationAndUppercase ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombination = "AND"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombinationOrUppercase  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombination = "OR"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombinationAnd, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombinationOr, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombinationAndUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersKind string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersKindGroup  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersKind = "group"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersKindFilter ObservabilityTelemetryQueryResponseRunQueryParametersFiltersKind = "filter"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFiltersKindGroup, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIncludes            ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "includes"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotIncludes         ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "not_includes"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationStartsWith          ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "starts_with"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationRegex               ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "regex"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationExists              ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "exists"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIsNull              ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "is_null"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIn                  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "in"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotIn               ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "not_in"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationEq                  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "eq"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNeq                 ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "neq"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGt                  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "gt"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGte                 ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "gte"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLt                  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "lt"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLte                 ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "lte"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationEquals              ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "="
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotEquals           ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "!="
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGreater             ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = ">"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGreaterOrEquals     ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = ">="
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLess                ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "<"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLessOrEquals        ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "<="
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIncludesUppercase   ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "INCLUDES"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationDoesNotInclude      ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationMatchRegex          ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationExistsUppercase     ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "EXISTS"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationDoesNotExist        ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationInUppercase         ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "IN"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotInUppercase      ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "NOT_IN"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationStartsWithUppercase ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIncludes, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotIncludes, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationStartsWith, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationRegex, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationExists, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIsNull, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIn, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotIn, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationEq, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNeq, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGt, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGte, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLt, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLte, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationEquals, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotEquals, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGreater, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationGreaterOrEquals, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLess, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationLessOrEquals, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationIncludesUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationDoesNotInclude, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationMatchRegex, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationExistsUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationDoesNotExist, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationInUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationNotInUppercase, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersTypeString  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType = "string"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersTypeNumber  ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType = "number"
	ObservabilityTelemetryQueryResponseRunQueryParametersFiltersTypeBoolean ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType = "boolean"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersFiltersTypeString, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersTypeNumber, ObservabilityTelemetryQueryResponseRunQueryParametersFiltersTypeBoolean:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseRunQueryParametersGroupBy struct {
	Type  ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysType `json:"type" api:"required"`
	Value string                                                            `json:"value" api:"required"`
	JSON  observabilityTelemetryQueryResponseRunQueryParametersGroupByJSON  `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersGroupByJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersGroupBy]
type observabilityTelemetryQueryResponseRunQueryParametersGroupByJSON struct {
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersGroupBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersGroupByJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysType string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysTypeString  ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysType = "string"
	ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysTypeNumber  ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysType = "number"
	ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysTypeBoolean ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysType = "boolean"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysTypeString, ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysTypeNumber, ObservabilityTelemetryQueryResponseRunQueryParametersGroupBysTypeBoolean:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseRunQueryParametersHaving struct {
	Key       string                                                                `json:"key" api:"required"`
	Operation ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation `json:"operation" api:"required"`
	Value     float64                                                               `json:"value" api:"required"`
	JSON      observabilityTelemetryQueryResponseRunQueryParametersHavingJSON       `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersHavingJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersHaving]
type observabilityTelemetryQueryResponseRunQueryParametersHavingJSON struct {
	Key         apijson.Field
	Operation   apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersHaving) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersHavingJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationEq  ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation = "eq"
	ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationNeq ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation = "neq"
	ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationGt  ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation = "gt"
	ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationGte ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation = "gte"
	ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationLt  ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation = "lt"
	ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationLte ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation = "lte"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationEq, ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationNeq, ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationGt, ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationGte, ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationLt, ObservabilityTelemetryQueryResponseRunQueryParametersHavingsOperationLte:
		return true
	}
	return false
}

// Define an expression to search using full-text search.
type ObservabilityTelemetryQueryResponseRunQueryParametersNeedle struct {
	Value     ObservabilityTelemetryQueryResponseRunQueryParametersNeedleValue `json:"value" api:"required"`
	IsRegex   bool                                                             `json:"isRegex"`
	MatchCase bool                                                             `json:"matchCase"`
	JSON      observabilityTelemetryQueryResponseRunQueryParametersNeedleJSON  `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersNeedleJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersNeedle]
type observabilityTelemetryQueryResponseRunQueryParametersNeedleJSON struct {
	Value       apijson.Field
	IsRegex     apijson.Field
	MatchCase   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersNeedle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersNeedleJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseRunQueryParametersNeedleValue struct {
	JSON observabilityTelemetryQueryResponseRunQueryParametersNeedleValueJSON `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersNeedleValueJSON contains
// the JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersNeedleValue]
type observabilityTelemetryQueryResponseRunQueryParametersNeedleValueJSON struct {
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersNeedleValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersNeedleValueJSON) RawJSON() string {
	return r.raw
}

// Configure the order of the results returned by the query.
type ObservabilityTelemetryQueryResponseRunQueryParametersOrderBy struct {
	// Configure which Calculation to order the results by.
	Value string `json:"value" api:"required"`
	// Set the order of the results
	Order ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrder `json:"order"`
	JSON  observabilityTelemetryQueryResponseRunQueryParametersOrderByJSON  `json:"-"`
}

// observabilityTelemetryQueryResponseRunQueryParametersOrderByJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseRunQueryParametersOrderBy]
type observabilityTelemetryQueryResponseRunQueryParametersOrderByJSON struct {
	Value       apijson.Field
	Order       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunQueryParametersOrderBy) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunQueryParametersOrderByJSON) RawJSON() string {
	return r.raw
}

// Set the order of the results
type ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrder string

const (
	ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrderAsc  ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrder = "asc"
	ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrderDesc ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrder = "desc"
)

func (r ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrder) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrderAsc, ObservabilityTelemetryQueryResponseRunQueryParametersOrderByOrderDesc:
		return true
	}
	return false
}

// Current execution status of the query run.
type ObservabilityTelemetryQueryResponseRunStatus string

const (
	ObservabilityTelemetryQueryResponseRunStatusStarted   ObservabilityTelemetryQueryResponseRunStatus = "STARTED"
	ObservabilityTelemetryQueryResponseRunStatusCompleted ObservabilityTelemetryQueryResponseRunStatus = "COMPLETED"
)

func (r ObservabilityTelemetryQueryResponseRunStatus) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseRunStatusStarted, ObservabilityTelemetryQueryResponseRunStatusCompleted:
		return true
	}
	return false
}

// Time range for the query execution
type ObservabilityTelemetryQueryResponseRunTimeframe struct {
	// Start timestamp for the query timeframe (Unix timestamp in milliseconds)
	From float64 `json:"from" api:"required"`
	// End timestamp for the query timeframe (Unix timestamp in milliseconds)
	To   float64                                             `json:"to" api:"required"`
	JSON observabilityTelemetryQueryResponseRunTimeframeJSON `json:"-"`
}

// observabilityTelemetryQueryResponseRunTimeframeJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseRunTimeframe]
type observabilityTelemetryQueryResponseRunTimeframeJSON struct {
	From        apijson.Field
	To          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunTimeframe) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunTimeframeJSON) RawJSON() string {
	return r.raw
}

// Query performance statistics from the database (does not include network
// latency).
type ObservabilityTelemetryQueryResponseRunStatistics struct {
	// Number of uncompressed bytes read from the table.
	BytesRead float64 `json:"bytes_read" api:"required"`
	// Time in seconds for the query to run.
	Elapsed float64 `json:"elapsed" api:"required"`
	// Number of rows scanned from the table.
	RowsRead float64 `json:"rows_read" api:"required"`
	// The level of Adaptive Bit Rate (ABR) sampling used for the query. If empty the
	// ABR level is 1
	AbrLevel float64                                              `json:"abr_level"`
	JSON     observabilityTelemetryQueryResponseRunStatisticsJSON `json:"-"`
}

// observabilityTelemetryQueryResponseRunStatisticsJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseRunStatistics]
type observabilityTelemetryQueryResponseRunStatisticsJSON struct {
	BytesRead   apijson.Field
	Elapsed     apijson.Field
	RowsRead    apijson.Field
	AbrLevel    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseRunStatistics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseRunStatisticsJSON) RawJSON() string {
	return r.raw
}

// Query performance statistics from the database. Includes execution time, rows
// scanned, and bytes read. Does not include network latency.
type ObservabilityTelemetryQueryResponseStatistics struct {
	// Number of uncompressed bytes read from the table.
	BytesRead float64 `json:"bytes_read" api:"required"`
	// Time in seconds for the query to run.
	Elapsed float64 `json:"elapsed" api:"required"`
	// Number of rows scanned from the table.
	RowsRead float64 `json:"rows_read" api:"required"`
	// The level of Adaptive Bit Rate (ABR) sampling used for the query. If empty the
	// ABR level is 1
	AbrLevel float64                                           `json:"abr_level"`
	JSON     observabilityTelemetryQueryResponseStatisticsJSON `json:"-"`
}

// observabilityTelemetryQueryResponseStatisticsJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryResponseStatistics]
type observabilityTelemetryQueryResponseStatisticsJSON struct {
	BytesRead   apijson.Field
	Elapsed     apijson.Field
	RowsRead    apijson.Field
	AbrLevel    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseStatistics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseStatisticsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseAgent struct {
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
	JSON        observabilityTelemetryQueryResponseAgentJSON `json:"-"`
}

// observabilityTelemetryQueryResponseAgentJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryQueryResponseAgent]
type observabilityTelemetryQueryResponseAgentJSON struct {
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

func (r *ObservabilityTelemetryQueryResponseAgent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseAgentJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCalculation struct {
	Aggregates  []ObservabilityTelemetryQueryResponseCalculationsAggregate `json:"aggregates" api:"required"`
	Calculation string                                                     `json:"calculation" api:"required"`
	Series      []ObservabilityTelemetryQueryResponseCalculationsSeries    `json:"series" api:"required"`
	Alias       string                                                     `json:"alias"`
	JSON        observabilityTelemetryQueryResponseCalculationJSON         `json:"-"`
}

// observabilityTelemetryQueryResponseCalculationJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseCalculation]
type observabilityTelemetryQueryResponseCalculationJSON struct {
	Aggregates  apijson.Field
	Calculation apijson.Field
	Series      apijson.Field
	Alias       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCalculation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCalculationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCalculationsAggregate struct {
	Count          float64                                                          `json:"count" api:"required"`
	Interval       float64                                                          `json:"interval" api:"required"`
	SampleInterval float64                                                          `json:"sampleInterval" api:"required"`
	Value          float64                                                          `json:"value" api:"required"`
	Groups         []ObservabilityTelemetryQueryResponseCalculationsAggregatesGroup `json:"groups"`
	JSON           observabilityTelemetryQueryResponseCalculationsAggregateJSON     `json:"-"`
}

// observabilityTelemetryQueryResponseCalculationsAggregateJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryResponseCalculationsAggregate]
type observabilityTelemetryQueryResponseCalculationsAggregateJSON struct {
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCalculationsAggregate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCalculationsAggregateJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCalculationsAggregatesGroup struct {
	Key   string                                                                    `json:"key" api:"required"`
	Value ObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion `json:"value" api:"required"`
	JSON  observabilityTelemetryQueryResponseCalculationsAggregatesGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryResponseCalculationsAggregatesGroupJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseCalculationsAggregatesGroup]
type observabilityTelemetryQueryResponseCalculationsAggregatesGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCalculationsAggregatesGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCalculationsAggregatesGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseCalculationsAggregatesGroupsValueUnion)(nil)).Elem(),
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

type ObservabilityTelemetryQueryResponseCalculationsSeries struct {
	Data []ObservabilityTelemetryQueryResponseCalculationsSeriesData `json:"data" api:"required"`
	Time string                                                      `json:"time" api:"required"`
	JSON observabilityTelemetryQueryResponseCalculationsSeriesJSON   `json:"-"`
}

// observabilityTelemetryQueryResponseCalculationsSeriesJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseCalculationsSeries]
type observabilityTelemetryQueryResponseCalculationsSeriesJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCalculationsSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCalculationsSeriesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCalculationsSeriesData struct {
	Count          float64                                                          `json:"count" api:"required"`
	Interval       float64                                                          `json:"interval" api:"required"`
	SampleInterval float64                                                          `json:"sampleInterval" api:"required"`
	Value          float64                                                          `json:"value" api:"required"`
	FirstSeen      string                                                           `json:"firstSeen"`
	Groups         []ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroup `json:"groups"`
	LastSeen       string                                                           `json:"lastSeen"`
	JSON           observabilityTelemetryQueryResponseCalculationsSeriesDataJSON    `json:"-"`
}

// observabilityTelemetryQueryResponseCalculationsSeriesDataJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryResponseCalculationsSeriesData]
type observabilityTelemetryQueryResponseCalculationsSeriesDataJSON struct {
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

func (r *ObservabilityTelemetryQueryResponseCalculationsSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCalculationsSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroup struct {
	Key   string                                                                    `json:"key" api:"required"`
	Value ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion `json:"value" api:"required"`
	JSON  observabilityTelemetryQueryResponseCalculationsSeriesDataGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryResponseCalculationsSeriesDataGroupJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroup]
type observabilityTelemetryQueryResponseCalculationsSeriesDataGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCalculationsSeriesDataGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseCalculationsSeriesDataGroupsValueUnion)(nil)).Elem(),
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

type ObservabilityTelemetryQueryResponseCompare struct {
	Aggregates  []ObservabilityTelemetryQueryResponseCompareAggregate `json:"aggregates" api:"required"`
	Calculation string                                                `json:"calculation" api:"required"`
	Series      []ObservabilityTelemetryQueryResponseCompareSeries    `json:"series" api:"required"`
	Alias       string                                                `json:"alias"`
	JSON        observabilityTelemetryQueryResponseCompareJSON        `json:"-"`
}

// observabilityTelemetryQueryResponseCompareJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryResponseCompare]
type observabilityTelemetryQueryResponseCompareJSON struct {
	Aggregates  apijson.Field
	Calculation apijson.Field
	Series      apijson.Field
	Alias       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCompare) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCompareJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCompareAggregate struct {
	Count          float64                                                     `json:"count" api:"required"`
	Interval       float64                                                     `json:"interval" api:"required"`
	SampleInterval float64                                                     `json:"sampleInterval" api:"required"`
	Value          float64                                                     `json:"value" api:"required"`
	Groups         []ObservabilityTelemetryQueryResponseCompareAggregatesGroup `json:"groups"`
	JSON           observabilityTelemetryQueryResponseCompareAggregateJSON     `json:"-"`
}

// observabilityTelemetryQueryResponseCompareAggregateJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseCompareAggregate]
type observabilityTelemetryQueryResponseCompareAggregateJSON struct {
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Value          apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCompareAggregate) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCompareAggregateJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCompareAggregatesGroup struct {
	Key   string                                                               `json:"key" api:"required"`
	Value ObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion `json:"value" api:"required"`
	JSON  observabilityTelemetryQueryResponseCompareAggregatesGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryResponseCompareAggregatesGroupJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryResponseCompareAggregatesGroup]
type observabilityTelemetryQueryResponseCompareAggregatesGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCompareAggregatesGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCompareAggregatesGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseCompareAggregatesGroupsValueUnion)(nil)).Elem(),
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

type ObservabilityTelemetryQueryResponseCompareSeries struct {
	Data []ObservabilityTelemetryQueryResponseCompareSeriesData `json:"data" api:"required"`
	Time string                                                 `json:"time" api:"required"`
	JSON observabilityTelemetryQueryResponseCompareSeriesJSON   `json:"-"`
}

// observabilityTelemetryQueryResponseCompareSeriesJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseCompareSeries]
type observabilityTelemetryQueryResponseCompareSeriesJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCompareSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCompareSeriesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCompareSeriesData struct {
	Count          float64                                                     `json:"count" api:"required"`
	Interval       float64                                                     `json:"interval" api:"required"`
	SampleInterval float64                                                     `json:"sampleInterval" api:"required"`
	Value          float64                                                     `json:"value" api:"required"`
	FirstSeen      string                                                      `json:"firstSeen"`
	Groups         []ObservabilityTelemetryQueryResponseCompareSeriesDataGroup `json:"groups"`
	LastSeen       string                                                      `json:"lastSeen"`
	JSON           observabilityTelemetryQueryResponseCompareSeriesDataJSON    `json:"-"`
}

// observabilityTelemetryQueryResponseCompareSeriesDataJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseCompareSeriesData]
type observabilityTelemetryQueryResponseCompareSeriesDataJSON struct {
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

func (r *ObservabilityTelemetryQueryResponseCompareSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCompareSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseCompareSeriesDataGroup struct {
	Key   string                                                               `json:"key" api:"required"`
	Value ObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion `json:"value" api:"required"`
	JSON  observabilityTelemetryQueryResponseCompareSeriesDataGroupJSON        `json:"-"`
}

// observabilityTelemetryQueryResponseCompareSeriesDataGroupJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryResponseCompareSeriesDataGroup]
type observabilityTelemetryQueryResponseCompareSeriesDataGroupJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseCompareSeriesDataGroup) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseCompareSeriesDataGroupJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion interface {
	ImplementsObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseCompareSeriesDataGroupsValueUnion)(nil)).Elem(),
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
type ObservabilityTelemetryQueryResponseEvents struct {
	// Total number of events matching the query (may exceed the number returned due to
	// limits).
	Count float64 `json:"count"`
	// List of individual telemetry events matching the query.
	Events []ObservabilityTelemetryQueryResponseEventsEvent `json:"events"`
	// List of fields discovered in the matched events. Useful for building dynamic
	// UIs.
	Fields []ObservabilityTelemetryQueryResponseEventsField `json:"fields"`
	// Time-series data for the matched events, bucketed by the query granularity.
	Series []ObservabilityTelemetryQueryResponseEventsSeries `json:"series"`
	JSON   observabilityTelemetryQueryResponseEventsJSON     `json:"-"`
}

// observabilityTelemetryQueryResponseEventsJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryQueryResponseEvents]
type observabilityTelemetryQueryResponseEventsJSON struct {
	Count       apijson.Field
	Events      apijson.Field
	Fields      apijson.Field
	Series      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEvents) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsJSON) RawJSON() string {
	return r.raw
}

// A single telemetry event representing a log line, span, or metric data point
// emitted by a Worker.
type ObservabilityTelemetryQueryResponseEventsEvent struct {
	// Structured metadata extracted from the event. These fields are indexed and
	// available for filtering and aggregation.
	Metadata ObservabilityTelemetryQueryResponseEventsEventsMetadata `json:"$metadata" api:"required"`
	// The dataset this event belongs to (e.g. cloudflare-workers).
	Dataset string `json:"dataset" api:"required"`
	// Raw log payload. May be a string or a structured object depending on how the log
	// was emitted.
	Source interface{} `json:"source" api:"required"`
	// Event timestamp as a Unix epoch in milliseconds.
	Timestamp int64 `json:"timestamp" api:"required"`
	// Cloudflare Containers event information that enriches your logs for identifying
	// and debugging issues.
	Containers interface{} `json:"$containers"`
	// Cloudflare Workers event information that enriches your logs for identifying and
	// debugging issues.
	Workers ObservabilityTelemetryQueryResponseEventsEventsWorkers `json:"$workers"`
	JSON    observabilityTelemetryQueryResponseEventsEventJSON     `json:"-"`
}

// observabilityTelemetryQueryResponseEventsEventJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseEventsEvent]
type observabilityTelemetryQueryResponseEventsEventJSON struct {
	Metadata    apijson.Field
	Dataset     apijson.Field
	Source      apijson.Field
	Timestamp   apijson.Field
	Containers  apijson.Field
	Workers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsEvent) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsEventJSON) RawJSON() string {
	return r.raw
}

// Structured metadata extracted from the event. These fields are indexed and
// available for filtering and aggregation.
type ObservabilityTelemetryQueryResponseEventsEventsMetadata struct {
	// Unique event ID. Use as the cursor value for offset-based pagination.
	ID string `json:"id" api:"required"`
	// Cloudflare account identifier.
	Account string `json:"account"`
	// Cloudflare product that generated this event (e.g. workers, pages).
	CloudService string `json:"cloudService"`
	// Whether this was a cold start (1) or warm invocation (0).
	ColdStart int64 `json:"coldStart"`
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
	JSON observabilityTelemetryQueryResponseEventsEventsMetadataJSON `json:"-"`
}

// observabilityTelemetryQueryResponseEventsEventsMetadataJSON contains the JSON
// metadata for the struct
// [ObservabilityTelemetryQueryResponseEventsEventsMetadata]
type observabilityTelemetryQueryResponseEventsEventsMetadataJSON struct {
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

func (r *ObservabilityTelemetryQueryResponseEventsEventsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsEventsMetadataJSON) RawJSON() string {
	return r.raw
}

// Cloudflare Workers event information that enriches your logs for identifying and
// debugging issues.
type ObservabilityTelemetryQueryResponseEventsEventsWorkers struct {
	EventType  ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType `json:"eventType" api:"required"`
	RequestID  string                                                          `json:"requestId" api:"required"`
	ScriptName string                                                          `json:"scriptName" api:"required"`
	CPUTimeMs  float64                                                         `json:"cpuTimeMs"`
	// This field can have the runtime type of
	// [[]ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectDiagnosticsChannelEvent].
	DiagnosticsChannelEvents interface{} `json:"diagnosticsChannelEvents"`
	DispatchNamespace        string      `json:"dispatchNamespace"`
	DurableObjectID          string      `json:"durableObjectId"`
	Entrypoint               string      `json:"entrypoint"`
	// This field can have the runtime type of [map[string]interface{}].
	Event          interface{}                                                          `json:"event"`
	ExecutionModel ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModel `json:"executionModel"`
	Outcome        string                                                               `json:"outcome"`
	// This field can have the runtime type of
	// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersion].
	ScriptVersion interface{}                                                `json:"scriptVersion"`
	SpanID        string                                                     `json:"spanId"`
	TraceID       string                                                     `json:"traceId"`
	Truncated     bool                                                       `json:"truncated"`
	WallTimeMs    float64                                                    `json:"wallTimeMs"`
	JSON          observabilityTelemetryQueryResponseEventsEventsWorkersJSON `json:"-"`
	union         ObservabilityTelemetryQueryResponseEventsEventsWorkersUnion
}

// observabilityTelemetryQueryResponseEventsEventsWorkersJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseEventsEventsWorkers]
type observabilityTelemetryQueryResponseEventsEventsWorkersJSON struct {
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
	ScriptVersion            apijson.Field
	SpanID                   apijson.Field
	TraceID                  apijson.Field
	Truncated                apijson.Field
	WallTimeMs               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r observabilityTelemetryQueryResponseEventsEventsWorkersJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilityTelemetryQueryResponseEventsEventsWorkers) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilityTelemetryQueryResponseEventsEventsWorkers{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObservabilityTelemetryQueryResponseEventsEventsWorkersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObject],
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObject].
func (r ObservabilityTelemetryQueryResponseEventsEventsWorkers) AsUnion() ObservabilityTelemetryQueryResponseEventsEventsWorkersUnion {
	return r.union
}

// Cloudflare Workers event information that enriches your logs for identifying and
// debugging issues.
//
// Union satisfied by
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObject] or
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObject].
type ObservabilityTelemetryQueryResponseEventsEventsWorkersUnion interface {
	implementsObservabilityTelemetryQueryResponseEventsEventsWorkers()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseEventsEventsWorkersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseEventsEventsWorkersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseEventsEventsWorkersObject{}),
		},
	)
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersObject struct {
	EventType       ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType      `json:"eventType" api:"required"`
	RequestID       string                                                                     `json:"requestId" api:"required"`
	ScriptName      string                                                                     `json:"scriptName" api:"required"`
	DurableObjectID string                                                                     `json:"durableObjectId"`
	Entrypoint      string                                                                     `json:"entrypoint"`
	Event           map[string]interface{}                                                     `json:"event"`
	ExecutionModel  ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModel `json:"executionModel"`
	Outcome         string                                                                     `json:"outcome"`
	ScriptVersion   ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersion  `json:"scriptVersion"`
	SpanID          string                                                                     `json:"spanId"`
	TraceID         string                                                                     `json:"traceId"`
	Truncated       bool                                                                       `json:"truncated"`
	JSON            observabilityTelemetryQueryResponseEventsEventsWorkersObjectJSON           `json:"-"`
}

// observabilityTelemetryQueryResponseEventsEventsWorkersObjectJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObject]
type observabilityTelemetryQueryResponseEventsEventsWorkersObjectJSON struct {
	EventType       apijson.Field
	RequestID       apijson.Field
	ScriptName      apijson.Field
	DurableObjectID apijson.Field
	Entrypoint      apijson.Field
	Event           apijson.Field
	ExecutionModel  apijson.Field
	Outcome         apijson.Field
	ScriptVersion   apijson.Field
	SpanID          apijson.Field
	TraceID         apijson.Field
	Truncated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsEventsWorkersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsEventsWorkersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersObject) implementsObservabilityTelemetryQueryResponseEventsEventsWorkers() {
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType string

const (
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeFetch     ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "fetch"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeScheduled ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "scheduled"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeAlarm     ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "alarm"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeCron      ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "cron"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeQueue     ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "queue"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeEmail     ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "email"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeTail      ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "tail"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeRpc       ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "rpc"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeWebsocket ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "websocket"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeWorkflow  ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "workflow"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeUnknown   ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType = "unknown"
)

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeFetch, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeScheduled, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeAlarm, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeCron, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeQueue, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeEmail, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeTail, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeRpc, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeWebsocket, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeWorkflow, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModel string

const (
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModelDurableObject ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModel = "durableObject"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModelStateless     ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModel = "stateless"
)

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModelDurableObject, ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersion struct {
	ID      string                                                                        `json:"id"`
	Message string                                                                        `json:"message"`
	Tag     string                                                                        `json:"tag"`
	JSON    observabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersionJSON `json:"-"`
}

// observabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersionJSON
// contains the JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersion]
type observabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersionJSON struct {
	ID          apijson.Field
	Message     apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsEventsWorkersObjectScriptVersionJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType string

const (
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeFetch     ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "fetch"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeScheduled ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "scheduled"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeAlarm     ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "alarm"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeCron      ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "cron"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeQueue     ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "queue"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeEmail     ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "email"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeTail      ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "tail"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeRpc       ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "rpc"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeWebsocket ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "websocket"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeWorkflow  ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "workflow"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeUnknown   ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType = "unknown"
)

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersEventType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeFetch, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeScheduled, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeAlarm, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeCron, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeQueue, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeEmail, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeTail, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeRpc, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeWebsocket, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeWorkflow, ObservabilityTelemetryQueryResponseEventsEventsWorkersEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModel string

const (
	ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModelDurableObject ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModel = "durableObject"
	ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModelStateless     ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModel = "stateless"
)

func (r ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModelDurableObject, ObservabilityTelemetryQueryResponseEventsEventsWorkersExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseEventsField struct {
	// Field name present in the matched events.
	Key string `json:"key" api:"required"`
	// Data type of the field (string, number, or boolean).
	Type string                                             `json:"type" api:"required"`
	JSON observabilityTelemetryQueryResponseEventsFieldJSON `json:"-"`
}

// observabilityTelemetryQueryResponseEventsFieldJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseEventsField]
type observabilityTelemetryQueryResponseEventsFieldJSON struct {
	Key         apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsField) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsFieldJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEventsSeries struct {
	Data []ObservabilityTelemetryQueryResponseEventsSeriesData `json:"data" api:"required"`
	Time string                                                `json:"time" api:"required"`
	JSON observabilityTelemetryQueryResponseEventsSeriesJSON   `json:"-"`
}

// observabilityTelemetryQueryResponseEventsSeriesJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseEventsSeries]
type observabilityTelemetryQueryResponseEventsSeriesJSON struct {
	Data        apijson.Field
	Time        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsSeries) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsSeriesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEventsSeriesData struct {
	Aggregates     ObservabilityTelemetryQueryResponseEventsSeriesDataAggregates `json:"aggregates" api:"required"`
	Count          float64                                                       `json:"count" api:"required"`
	Interval       float64                                                       `json:"interval" api:"required"`
	SampleInterval float64                                                       `json:"sampleInterval" api:"required"`
	Errors         float64                                                       `json:"errors"`
	// Groups in the query results.
	Groups map[string]ObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion `json:"groups"`
	JSON   observabilityTelemetryQueryResponseEventsSeriesDataJSON                   `json:"-"`
}

// observabilityTelemetryQueryResponseEventsSeriesDataJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseEventsSeriesData]
type observabilityTelemetryQueryResponseEventsSeriesDataJSON struct {
	Aggregates     apijson.Field
	Count          apijson.Field
	Interval       apijson.Field
	SampleInterval apijson.Field
	Errors         apijson.Field
	Groups         apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsSeriesData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsSeriesDataJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEventsSeriesDataAggregates struct {
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
	JSON observabilityTelemetryQueryResponseEventsSeriesDataAggregatesJSON `json:"-"`
}

// observabilityTelemetryQueryResponseEventsSeriesDataAggregatesJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseEventsSeriesDataAggregates]
type observabilityTelemetryQueryResponseEventsSeriesDataAggregatesJSON struct {
	Count       apijson.Field
	Interval    apijson.Field
	FirstSeen   apijson.Field
	LastSeen    apijson.Field
	Bin         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEventsSeriesDataAggregates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEventsSeriesDataAggregatesJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion interface {
	ImplementsObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseEventsSeriesDataGroupsUnion)(nil)).Elem(),
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
type ObservabilityTelemetryQueryResponseInvocation struct {
	// Structured metadata extracted from the event. These fields are indexed and
	// available for filtering and aggregation.
	Metadata ObservabilityTelemetryQueryResponseInvocationsMetadata `json:"$metadata" api:"required"`
	// The dataset this event belongs to (e.g. cloudflare-workers).
	Dataset string `json:"dataset" api:"required"`
	// Raw log payload. May be a string or a structured object depending on how the log
	// was emitted.
	Source interface{} `json:"source" api:"required"`
	// Event timestamp as a Unix epoch in milliseconds.
	Timestamp int64 `json:"timestamp" api:"required"`
	// Cloudflare Containers event information that enriches your logs for identifying
	// and debugging issues.
	Containers interface{} `json:"$containers"`
	// Cloudflare Workers event information that enriches your logs for identifying and
	// debugging issues.
	Workers ObservabilityTelemetryQueryResponseInvocationsWorkers `json:"$workers"`
	JSON    observabilityTelemetryQueryResponseInvocationJSON     `json:"-"`
}

// observabilityTelemetryQueryResponseInvocationJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryResponseInvocation]
type observabilityTelemetryQueryResponseInvocationJSON struct {
	Metadata    apijson.Field
	Dataset     apijson.Field
	Source      apijson.Field
	Timestamp   apijson.Field
	Containers  apijson.Field
	Workers     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseInvocation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseInvocationJSON) RawJSON() string {
	return r.raw
}

// Structured metadata extracted from the event. These fields are indexed and
// available for filtering and aggregation.
type ObservabilityTelemetryQueryResponseInvocationsMetadata struct {
	// Unique event ID. Use as the cursor value for offset-based pagination.
	ID string `json:"id" api:"required"`
	// Cloudflare account identifier.
	Account string `json:"account"`
	// Cloudflare product that generated this event (e.g. workers, pages).
	CloudService string `json:"cloudService"`
	// Whether this was a cold start (1) or warm invocation (0).
	ColdStart int64 `json:"coldStart"`
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
	JSON observabilityTelemetryQueryResponseInvocationsMetadataJSON `json:"-"`
}

// observabilityTelemetryQueryResponseInvocationsMetadataJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseInvocationsMetadata]
type observabilityTelemetryQueryResponseInvocationsMetadataJSON struct {
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

func (r *ObservabilityTelemetryQueryResponseInvocationsMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseInvocationsMetadataJSON) RawJSON() string {
	return r.raw
}

// Cloudflare Workers event information that enriches your logs for identifying and
// debugging issues.
type ObservabilityTelemetryQueryResponseInvocationsWorkers struct {
	EventType  ObservabilityTelemetryQueryResponseInvocationsWorkersEventType `json:"eventType" api:"required"`
	RequestID  string                                                         `json:"requestId" api:"required"`
	ScriptName string                                                         `json:"scriptName" api:"required"`
	CPUTimeMs  float64                                                        `json:"cpuTimeMs"`
	// This field can have the runtime type of
	// [[]ObservabilityTelemetryQueryResponseInvocationsWorkersObjectDiagnosticsChannelEvent].
	DiagnosticsChannelEvents interface{} `json:"diagnosticsChannelEvents"`
	DispatchNamespace        string      `json:"dispatchNamespace"`
	DurableObjectID          string      `json:"durableObjectId"`
	Entrypoint               string      `json:"entrypoint"`
	// This field can have the runtime type of [map[string]interface{}].
	Event          interface{}                                                         `json:"event"`
	ExecutionModel ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModel `json:"executionModel"`
	Outcome        string                                                              `json:"outcome"`
	// This field can have the runtime type of
	// [ObservabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersion].
	ScriptVersion interface{}                                               `json:"scriptVersion"`
	SpanID        string                                                    `json:"spanId"`
	TraceID       string                                                    `json:"traceId"`
	Truncated     bool                                                      `json:"truncated"`
	WallTimeMs    float64                                                   `json:"wallTimeMs"`
	JSON          observabilityTelemetryQueryResponseInvocationsWorkersJSON `json:"-"`
	union         ObservabilityTelemetryQueryResponseInvocationsWorkersUnion
}

// observabilityTelemetryQueryResponseInvocationsWorkersJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseInvocationsWorkers]
type observabilityTelemetryQueryResponseInvocationsWorkersJSON struct {
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
	ScriptVersion            apijson.Field
	SpanID                   apijson.Field
	TraceID                  apijson.Field
	Truncated                apijson.Field
	WallTimeMs               apijson.Field
	raw                      string
	ExtraFields              map[string]apijson.Field
}

func (r observabilityTelemetryQueryResponseInvocationsWorkersJSON) RawJSON() string {
	return r.raw
}

func (r *ObservabilityTelemetryQueryResponseInvocationsWorkers) UnmarshalJSON(data []byte) (err error) {
	*r = ObservabilityTelemetryQueryResponseInvocationsWorkers{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [ObservabilityTelemetryQueryResponseInvocationsWorkersUnion]
// interface which you can cast to the specific types for more type safety.
//
// Possible runtime types of the union are
// [ObservabilityTelemetryQueryResponseInvocationsWorkersObject],
// [ObservabilityTelemetryQueryResponseInvocationsWorkersObject].
func (r ObservabilityTelemetryQueryResponseInvocationsWorkers) AsUnion() ObservabilityTelemetryQueryResponseInvocationsWorkersUnion {
	return r.union
}

// Cloudflare Workers event information that enriches your logs for identifying and
// debugging issues.
//
// Union satisfied by [ObservabilityTelemetryQueryResponseInvocationsWorkersObject]
// or [ObservabilityTelemetryQueryResponseInvocationsWorkersObject].
type ObservabilityTelemetryQueryResponseInvocationsWorkersUnion interface {
	implementsObservabilityTelemetryQueryResponseInvocationsWorkers()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryQueryResponseInvocationsWorkersUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseInvocationsWorkersObject{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(ObservabilityTelemetryQueryResponseInvocationsWorkersObject{}),
		},
	)
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersObject struct {
	EventType       ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType      `json:"eventType" api:"required"`
	RequestID       string                                                                    `json:"requestId" api:"required"`
	ScriptName      string                                                                    `json:"scriptName" api:"required"`
	DurableObjectID string                                                                    `json:"durableObjectId"`
	Entrypoint      string                                                                    `json:"entrypoint"`
	Event           map[string]interface{}                                                    `json:"event"`
	ExecutionModel  ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModel `json:"executionModel"`
	Outcome         string                                                                    `json:"outcome"`
	ScriptVersion   ObservabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersion  `json:"scriptVersion"`
	SpanID          string                                                                    `json:"spanId"`
	TraceID         string                                                                    `json:"traceId"`
	Truncated       bool                                                                      `json:"truncated"`
	JSON            observabilityTelemetryQueryResponseInvocationsWorkersObjectJSON           `json:"-"`
}

// observabilityTelemetryQueryResponseInvocationsWorkersObjectJSON contains the
// JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseInvocationsWorkersObject]
type observabilityTelemetryQueryResponseInvocationsWorkersObjectJSON struct {
	EventType       apijson.Field
	RequestID       apijson.Field
	ScriptName      apijson.Field
	DurableObjectID apijson.Field
	Entrypoint      apijson.Field
	Event           apijson.Field
	ExecutionModel  apijson.Field
	Outcome         apijson.Field
	ScriptVersion   apijson.Field
	SpanID          apijson.Field
	TraceID         apijson.Field
	Truncated       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseInvocationsWorkersObject) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseInvocationsWorkersObjectJSON) RawJSON() string {
	return r.raw
}

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersObject) implementsObservabilityTelemetryQueryResponseInvocationsWorkers() {
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType string

const (
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeFetch     ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "fetch"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeScheduled ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "scheduled"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeAlarm     ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "alarm"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeCron      ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "cron"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeQueue     ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "queue"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeEmail     ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "email"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeTail      ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "tail"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeRpc       ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "rpc"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeWebsocket ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "websocket"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeWorkflow  ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "workflow"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeUnknown   ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType = "unknown"
)

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeFetch, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeScheduled, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeAlarm, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeCron, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeQueue, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeEmail, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeTail, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeRpc, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeWebsocket, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeWorkflow, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModel string

const (
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModelDurableObject ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModel = "durableObject"
	ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModelStateless     ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModel = "stateless"
)

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModelDurableObject, ObservabilityTelemetryQueryResponseInvocationsWorkersObjectExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersion struct {
	ID      string                                                                       `json:"id"`
	Message string                                                                       `json:"message"`
	Tag     string                                                                       `json:"tag"`
	JSON    observabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersionJSON `json:"-"`
}

// observabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersionJSON
// contains the JSON metadata for the struct
// [ObservabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersion]
type observabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersionJSON struct {
	ID          apijson.Field
	Message     apijson.Field
	Tag         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseInvocationsWorkersObjectScriptVersionJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersEventType string

const (
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeFetch     ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "fetch"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeScheduled ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "scheduled"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeAlarm     ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "alarm"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeCron      ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "cron"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeQueue     ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "queue"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeEmail     ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "email"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeTail      ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "tail"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeRpc       ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "rpc"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeWebsocket ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "websocket"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeWorkflow  ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "workflow"
	ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeUnknown   ObservabilityTelemetryQueryResponseInvocationsWorkersEventType = "unknown"
)

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersEventType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeFetch, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeScheduled, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeAlarm, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeCron, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeQueue, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeEmail, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeTail, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeRpc, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeWebsocket, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeWorkflow, ObservabilityTelemetryQueryResponseInvocationsWorkersEventTypeUnknown:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModel string

const (
	ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModelDurableObject ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModel = "durableObject"
	ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModelStateless     ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModel = "stateless"
)

func (r ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModel) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModelDurableObject, ObservabilityTelemetryQueryResponseInvocationsWorkersExecutionModelStateless:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseTrace struct {
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
	JSON   observabilityTelemetryQueryResponseTraceJSON `json:"-"`
}

// observabilityTelemetryQueryResponseTraceJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryQueryResponseTrace]
type observabilityTelemetryQueryResponseTraceJSON struct {
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

func (r *ObservabilityTelemetryQueryResponseTrace) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseTraceJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryValuesResponse struct {
	Dataset string                                         `json:"dataset" api:"required"`
	Key     string                                         `json:"key" api:"required"`
	Type    ObservabilityTelemetryValuesResponseType       `json:"type" api:"required"`
	Value   ObservabilityTelemetryValuesResponseValueUnion `json:"value" api:"required"`
	JSON    observabilityTelemetryValuesResponseJSON       `json:"-"`
}

// observabilityTelemetryValuesResponseJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryValuesResponse]
type observabilityTelemetryValuesResponseJSON struct {
	Dataset     apijson.Field
	Key         apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryValuesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryValuesResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryValuesResponseType string

const (
	ObservabilityTelemetryValuesResponseTypeString  ObservabilityTelemetryValuesResponseType = "string"
	ObservabilityTelemetryValuesResponseTypeBoolean ObservabilityTelemetryValuesResponseType = "boolean"
	ObservabilityTelemetryValuesResponseTypeNumber  ObservabilityTelemetryValuesResponseType = "number"
)

func (r ObservabilityTelemetryValuesResponseType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesResponseTypeString, ObservabilityTelemetryValuesResponseTypeBoolean, ObservabilityTelemetryValuesResponseTypeNumber:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryValuesResponseValueUnion interface {
	ImplementsObservabilityTelemetryValuesResponseValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryValuesResponseValueUnion)(nil)).Elem(),
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

type ObservabilityTelemetryKeysParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Leave this empty to use the default datasets
	Datasets param.Field[[]string] `json:"datasets"`
	// Apply filters to narrow key discovery. Supports nested groups via kind: 'group'.
	// Maximum nesting depth is 4.
	Filters param.Field[[]ObservabilityTelemetryKeysParamsFilterUnion] `json:"filters"`
	From    param.Field[float64]                                       `json:"from"`
	// If the user suggests a key, use this to narrow down the list of keys returned.
	// Make sure matchCase is false to avoid case sensitivity issues.
	KeyNeedle param.Field[ObservabilityTelemetryKeysParamsKeyNeedle] `json:"keyNeedle"`
	// Advanced usage: set limit=1000+ to retrieve comprehensive key options without
	// needing additional filtering.
	Limit param.Field[float64] `json:"limit"`
	// Search for a specific substring in any of the events
	Needle param.Field[ObservabilityTelemetryKeysParamsNeedle] `json:"needle"`
	To     param.Field[float64]                                `json:"to"`
}

func (r ObservabilityTelemetryKeysParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Supports nested groups via kind: 'group'.
type ObservabilityTelemetryKeysParamsFilter struct {
	FilterCombination param.Field[ObservabilityTelemetryKeysParamsFiltersFilterCombination] `json:"filterCombination"`
	Filters           param.Field[interface{}]                                              `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  param.Field[string]                                      `json:"key"`
	Kind param.Field[ObservabilityTelemetryKeysParamsFiltersKind] `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryKeysParamsFiltersOperation] `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type  param.Field[ObservabilityTelemetryKeysParamsFiltersType] `json:"type"`
	Value param.Field[interface{}]                                 `json:"value"`
}

func (r ObservabilityTelemetryKeysParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryKeysParamsFilter) implementsObservabilityTelemetryKeysParamsFilterUnion() {
}

// Supports nested groups via kind: 'group'.
//
// Satisfied by [workers.ObservabilityTelemetryKeysParamsFiltersObject],
// [workers.ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeaf],
// [ObservabilityTelemetryKeysParamsFilter].
type ObservabilityTelemetryKeysParamsFilterUnion interface {
	implementsObservabilityTelemetryKeysParamsFilterUnion()
}

type ObservabilityTelemetryKeysParamsFiltersObject struct {
	FilterCombination param.Field[ObservabilityTelemetryKeysParamsFiltersObjectFilterCombination] `json:"filterCombination" api:"required"`
	Filters           param.Field[[]ObservabilityTelemetryKeysParamsFiltersObjectFilterUnion]     `json:"filters" api:"required"`
	Kind              param.Field[ObservabilityTelemetryKeysParamsFiltersObjectKind]              `json:"kind" api:"required"`
}

func (r ObservabilityTelemetryKeysParamsFiltersObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryKeysParamsFiltersObject) implementsObservabilityTelemetryKeysParamsFilterUnion() {
}

type ObservabilityTelemetryKeysParamsFiltersObjectFilterCombination string

const (
	ObservabilityTelemetryKeysParamsFiltersObjectFilterCombinationAnd          ObservabilityTelemetryKeysParamsFiltersObjectFilterCombination = "and"
	ObservabilityTelemetryKeysParamsFiltersObjectFilterCombinationOr           ObservabilityTelemetryKeysParamsFiltersObjectFilterCombination = "or"
	ObservabilityTelemetryKeysParamsFiltersObjectFilterCombinationAndUppercase ObservabilityTelemetryKeysParamsFiltersObjectFilterCombination = "AND"
	ObservabilityTelemetryKeysParamsFiltersObjectFilterCombinationOrUppercase  ObservabilityTelemetryKeysParamsFiltersObjectFilterCombination = "OR"
)

func (r ObservabilityTelemetryKeysParamsFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersObjectFilterCombinationAnd, ObservabilityTelemetryKeysParamsFiltersObjectFilterCombinationOr, ObservabilityTelemetryKeysParamsFiltersObjectFilterCombinationAndUppercase, ObservabilityTelemetryKeysParamsFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

// Supports nested groups via kind: 'group'.
type ObservabilityTelemetryKeysParamsFiltersObjectFilter struct {
	FilterCombination param.Field[ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombination] `json:"filterCombination"`
	Filters           param.Field[interface{}]                                                           `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  param.Field[string]                                                   `json:"key"`
	Kind param.Field[ObservabilityTelemetryKeysParamsFiltersObjectFiltersKind] `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation] `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type  param.Field[ObservabilityTelemetryKeysParamsFiltersObjectFiltersType] `json:"type"`
	Value param.Field[interface{}]                                              `json:"value"`
}

func (r ObservabilityTelemetryKeysParamsFiltersObjectFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryKeysParamsFiltersObjectFilter) implementsObservabilityTelemetryKeysParamsFiltersObjectFilterUnion() {
}

// Supports nested groups via kind: 'group'.
//
// Satisfied by
// [workers.ObservabilityTelemetryKeysParamsFiltersObjectFiltersObject],
// [workers.ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeaf],
// [ObservabilityTelemetryKeysParamsFiltersObjectFilter].
type ObservabilityTelemetryKeysParamsFiltersObjectFilterUnion interface {
	implementsObservabilityTelemetryKeysParamsFiltersObjectFilterUnion()
}

type ObservabilityTelemetryKeysParamsFiltersObjectFiltersObject struct {
	FilterCombination param.Field[ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombination] `json:"filterCombination" api:"required"`
	Filters           param.Field[[]interface{}]                                                               `json:"filters" api:"required"`
	Kind              param.Field[ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectKind]              `json:"kind" api:"required"`
}

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersObject) implementsObservabilityTelemetryKeysParamsFiltersObjectFilterUnion() {
}

type ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombination string

const (
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombinationAnd          ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombination = "and"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombinationOr           ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombination = "or"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombinationAndUppercase ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombination = "AND"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombinationOrUppercase  ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombination = "OR"
)

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombinationAnd, ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombinationOr, ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombinationAndUppercase, ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectKind string

const (
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectKindGroup ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectKind = "group"
)

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersObjectFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key param.Field[string] `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation] `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type param.Field[ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafType] `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind param.Field[ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafKind] `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value param.Field[ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeaf) implementsObservabilityTelemetryKeysParamsFiltersObjectFilterUnion() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafTypeString  ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafTypeString, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafKindFilter ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafKindFilter:
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
type ObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilityTelemetryKeysParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion()
}

type ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombination string

const (
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombinationAnd          ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombination = "and"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombinationOr           ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombination = "or"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombinationAndUppercase ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombination = "AND"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombinationOrUppercase  ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombination = "OR"
)

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombinationAnd, ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombinationOr, ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombinationAndUppercase, ObservabilityTelemetryKeysParamsFiltersObjectFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryKeysParamsFiltersObjectFiltersKind string

const (
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersKindGroup  ObservabilityTelemetryKeysParamsFiltersObjectFiltersKind = "group"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersKindFilter ObservabilityTelemetryKeysParamsFiltersObjectFiltersKind = "filter"
)

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersObjectFiltersKindGroup, ObservabilityTelemetryKeysParamsFiltersObjectFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation string

const (
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationIncludes            ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "includes"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationNotIncludes         ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "not_includes"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationStartsWith          ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "starts_with"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationRegex               ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "regex"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationExists              ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "exists"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationIsNull              ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "is_null"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationIn                  ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "in"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationNotIn               ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "not_in"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationEq                  ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "eq"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationNeq                 ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "neq"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationGt                  ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "gt"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationGte                 ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "gte"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationLt                  ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "lt"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationLte                 ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "lte"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationEquals              ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "="
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationNotEquals           ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "!="
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationGreater             ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = ">"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationGreaterOrEquals     ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = ">="
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationLess                ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "<"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationLessOrEquals        ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "<="
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationIncludesUppercase   ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "INCLUDES"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationDoesNotInclude      ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationMatchRegex          ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationExistsUppercase     ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "EXISTS"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationDoesNotExist        ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationInUppercase         ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "IN"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationNotInUppercase      ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "NOT_IN"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationStartsWithUppercase ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationIncludes, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationNotIncludes, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationStartsWith, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationRegex, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationExists, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationIsNull, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationIn, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationNotIn, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationEq, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationNeq, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationGt, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationGte, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationLt, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationLte, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationEquals, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationNotEquals, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationGreater, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationGreaterOrEquals, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationLess, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationLessOrEquals, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationIncludesUppercase, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationDoesNotInclude, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationMatchRegex, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationExistsUppercase, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationDoesNotExist, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationInUppercase, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationNotInUppercase, ObservabilityTelemetryKeysParamsFiltersObjectFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryKeysParamsFiltersObjectFiltersType string

const (
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersTypeString  ObservabilityTelemetryKeysParamsFiltersObjectFiltersType = "string"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersTypeNumber  ObservabilityTelemetryKeysParamsFiltersObjectFiltersType = "number"
	ObservabilityTelemetryKeysParamsFiltersObjectFiltersTypeBoolean ObservabilityTelemetryKeysParamsFiltersObjectFiltersType = "boolean"
)

func (r ObservabilityTelemetryKeysParamsFiltersObjectFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersObjectFiltersTypeString, ObservabilityTelemetryKeysParamsFiltersObjectFiltersTypeNumber, ObservabilityTelemetryKeysParamsFiltersObjectFiltersTypeBoolean:
		return true
	}
	return false
}

type ObservabilityTelemetryKeysParamsFiltersObjectKind string

const (
	ObservabilityTelemetryKeysParamsFiltersObjectKindGroup ObservabilityTelemetryKeysParamsFiltersObjectKind = "group"
)

func (r ObservabilityTelemetryKeysParamsFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key param.Field[string] `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation] `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type param.Field[ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafType] `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind param.Field[ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafKind] `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value param.Field[ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeaf) implementsObservabilityTelemetryKeysParamsFilterUnion() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafTypeString  ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafTypeString, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafKindFilter ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafKindFilter:
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
type ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafValueUnion()
}

type ObservabilityTelemetryKeysParamsFiltersFilterCombination string

const (
	ObservabilityTelemetryKeysParamsFiltersFilterCombinationAnd          ObservabilityTelemetryKeysParamsFiltersFilterCombination = "and"
	ObservabilityTelemetryKeysParamsFiltersFilterCombinationOr           ObservabilityTelemetryKeysParamsFiltersFilterCombination = "or"
	ObservabilityTelemetryKeysParamsFiltersFilterCombinationAndUppercase ObservabilityTelemetryKeysParamsFiltersFilterCombination = "AND"
	ObservabilityTelemetryKeysParamsFiltersFilterCombinationOrUppercase  ObservabilityTelemetryKeysParamsFiltersFilterCombination = "OR"
)

func (r ObservabilityTelemetryKeysParamsFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersFilterCombinationAnd, ObservabilityTelemetryKeysParamsFiltersFilterCombinationOr, ObservabilityTelemetryKeysParamsFiltersFilterCombinationAndUppercase, ObservabilityTelemetryKeysParamsFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryKeysParamsFiltersKind string

const (
	ObservabilityTelemetryKeysParamsFiltersKindGroup  ObservabilityTelemetryKeysParamsFiltersKind = "group"
	ObservabilityTelemetryKeysParamsFiltersKindFilter ObservabilityTelemetryKeysParamsFiltersKind = "filter"
)

func (r ObservabilityTelemetryKeysParamsFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersKindGroup, ObservabilityTelemetryKeysParamsFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryKeysParamsFiltersOperation string

const (
	ObservabilityTelemetryKeysParamsFiltersOperationIncludes            ObservabilityTelemetryKeysParamsFiltersOperation = "includes"
	ObservabilityTelemetryKeysParamsFiltersOperationNotIncludes         ObservabilityTelemetryKeysParamsFiltersOperation = "not_includes"
	ObservabilityTelemetryKeysParamsFiltersOperationStartsWith          ObservabilityTelemetryKeysParamsFiltersOperation = "starts_with"
	ObservabilityTelemetryKeysParamsFiltersOperationRegex               ObservabilityTelemetryKeysParamsFiltersOperation = "regex"
	ObservabilityTelemetryKeysParamsFiltersOperationExists              ObservabilityTelemetryKeysParamsFiltersOperation = "exists"
	ObservabilityTelemetryKeysParamsFiltersOperationIsNull              ObservabilityTelemetryKeysParamsFiltersOperation = "is_null"
	ObservabilityTelemetryKeysParamsFiltersOperationIn                  ObservabilityTelemetryKeysParamsFiltersOperation = "in"
	ObservabilityTelemetryKeysParamsFiltersOperationNotIn               ObservabilityTelemetryKeysParamsFiltersOperation = "not_in"
	ObservabilityTelemetryKeysParamsFiltersOperationEq                  ObservabilityTelemetryKeysParamsFiltersOperation = "eq"
	ObservabilityTelemetryKeysParamsFiltersOperationNeq                 ObservabilityTelemetryKeysParamsFiltersOperation = "neq"
	ObservabilityTelemetryKeysParamsFiltersOperationGt                  ObservabilityTelemetryKeysParamsFiltersOperation = "gt"
	ObservabilityTelemetryKeysParamsFiltersOperationGte                 ObservabilityTelemetryKeysParamsFiltersOperation = "gte"
	ObservabilityTelemetryKeysParamsFiltersOperationLt                  ObservabilityTelemetryKeysParamsFiltersOperation = "lt"
	ObservabilityTelemetryKeysParamsFiltersOperationLte                 ObservabilityTelemetryKeysParamsFiltersOperation = "lte"
	ObservabilityTelemetryKeysParamsFiltersOperationEquals              ObservabilityTelemetryKeysParamsFiltersOperation = "="
	ObservabilityTelemetryKeysParamsFiltersOperationNotEquals           ObservabilityTelemetryKeysParamsFiltersOperation = "!="
	ObservabilityTelemetryKeysParamsFiltersOperationGreater             ObservabilityTelemetryKeysParamsFiltersOperation = ">"
	ObservabilityTelemetryKeysParamsFiltersOperationGreaterOrEquals     ObservabilityTelemetryKeysParamsFiltersOperation = ">="
	ObservabilityTelemetryKeysParamsFiltersOperationLess                ObservabilityTelemetryKeysParamsFiltersOperation = "<"
	ObservabilityTelemetryKeysParamsFiltersOperationLessOrEquals        ObservabilityTelemetryKeysParamsFiltersOperation = "<="
	ObservabilityTelemetryKeysParamsFiltersOperationIncludesUppercase   ObservabilityTelemetryKeysParamsFiltersOperation = "INCLUDES"
	ObservabilityTelemetryKeysParamsFiltersOperationDoesNotInclude      ObservabilityTelemetryKeysParamsFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryKeysParamsFiltersOperationMatchRegex          ObservabilityTelemetryKeysParamsFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryKeysParamsFiltersOperationExistsUppercase     ObservabilityTelemetryKeysParamsFiltersOperation = "EXISTS"
	ObservabilityTelemetryKeysParamsFiltersOperationDoesNotExist        ObservabilityTelemetryKeysParamsFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryKeysParamsFiltersOperationInUppercase         ObservabilityTelemetryKeysParamsFiltersOperation = "IN"
	ObservabilityTelemetryKeysParamsFiltersOperationNotInUppercase      ObservabilityTelemetryKeysParamsFiltersOperation = "NOT_IN"
	ObservabilityTelemetryKeysParamsFiltersOperationStartsWithUppercase ObservabilityTelemetryKeysParamsFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryKeysParamsFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersOperationIncludes, ObservabilityTelemetryKeysParamsFiltersOperationNotIncludes, ObservabilityTelemetryKeysParamsFiltersOperationStartsWith, ObservabilityTelemetryKeysParamsFiltersOperationRegex, ObservabilityTelemetryKeysParamsFiltersOperationExists, ObservabilityTelemetryKeysParamsFiltersOperationIsNull, ObservabilityTelemetryKeysParamsFiltersOperationIn, ObservabilityTelemetryKeysParamsFiltersOperationNotIn, ObservabilityTelemetryKeysParamsFiltersOperationEq, ObservabilityTelemetryKeysParamsFiltersOperationNeq, ObservabilityTelemetryKeysParamsFiltersOperationGt, ObservabilityTelemetryKeysParamsFiltersOperationGte, ObservabilityTelemetryKeysParamsFiltersOperationLt, ObservabilityTelemetryKeysParamsFiltersOperationLte, ObservabilityTelemetryKeysParamsFiltersOperationEquals, ObservabilityTelemetryKeysParamsFiltersOperationNotEquals, ObservabilityTelemetryKeysParamsFiltersOperationGreater, ObservabilityTelemetryKeysParamsFiltersOperationGreaterOrEquals, ObservabilityTelemetryKeysParamsFiltersOperationLess, ObservabilityTelemetryKeysParamsFiltersOperationLessOrEquals, ObservabilityTelemetryKeysParamsFiltersOperationIncludesUppercase, ObservabilityTelemetryKeysParamsFiltersOperationDoesNotInclude, ObservabilityTelemetryKeysParamsFiltersOperationMatchRegex, ObservabilityTelemetryKeysParamsFiltersOperationExistsUppercase, ObservabilityTelemetryKeysParamsFiltersOperationDoesNotExist, ObservabilityTelemetryKeysParamsFiltersOperationInUppercase, ObservabilityTelemetryKeysParamsFiltersOperationNotInUppercase, ObservabilityTelemetryKeysParamsFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryKeysParamsFiltersType string

const (
	ObservabilityTelemetryKeysParamsFiltersTypeString  ObservabilityTelemetryKeysParamsFiltersType = "string"
	ObservabilityTelemetryKeysParamsFiltersTypeNumber  ObservabilityTelemetryKeysParamsFiltersType = "number"
	ObservabilityTelemetryKeysParamsFiltersTypeBoolean ObservabilityTelemetryKeysParamsFiltersType = "boolean"
)

func (r ObservabilityTelemetryKeysParamsFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeysParamsFiltersTypeString, ObservabilityTelemetryKeysParamsFiltersTypeNumber, ObservabilityTelemetryKeysParamsFiltersTypeBoolean:
		return true
	}
	return false
}

// If the user suggests a key, use this to narrow down the list of keys returned.
// Make sure matchCase is false to avoid case sensitivity issues.
type ObservabilityTelemetryKeysParamsKeyNeedle struct {
	// The text or pattern to search for.
	Value param.Field[ObservabilityTelemetryKeysParamsKeyNeedleValueUnion] `json:"value" api:"required"`
	// When true, treats the value as a regular expression (RE2 syntax).
	IsRegex param.Field[bool] `json:"isRegex"`
	// When true, performs a case-sensitive search. Defaults to case-insensitive.
	MatchCase param.Field[bool] `json:"matchCase"`
}

func (r ObservabilityTelemetryKeysParamsKeyNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The text or pattern to search for.
//
// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryKeysParamsKeyNeedleValueUnion interface {
	ImplementsObservabilityTelemetryKeysParamsKeyNeedleValueUnion()
}

// Search for a specific substring in any of the events
type ObservabilityTelemetryKeysParamsNeedle struct {
	// The text or pattern to search for.
	Value param.Field[ObservabilityTelemetryKeysParamsNeedleValueUnion] `json:"value" api:"required"`
	// When true, treats the value as a regular expression (RE2 syntax).
	IsRegex param.Field[bool] `json:"isRegex"`
	// When true, performs a case-sensitive search. Defaults to case-insensitive.
	MatchCase param.Field[bool] `json:"matchCase"`
}

func (r ObservabilityTelemetryKeysParamsNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The text or pattern to search for.
//
// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryKeysParamsNeedleValueUnion interface {
	ImplementsObservabilityTelemetryKeysParamsNeedleValueUnion()
}

type ObservabilityTelemetryQueryParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Identifier for the query. When parameters are omitted, this ID is used to load a
	// previously saved query's parameters. When providing parameters inline, pass any
	// identifier (e.g. an ad-hoc ID).
	QueryID param.Field[string] `json:"queryId" api:"required"`
	// Timeframe for the query using Unix timestamps in milliseconds. Narrower
	// timeframes produce faster responses and more specific results.
	Timeframe param.Field[ObservabilityTelemetryQueryParamsTimeframe] `json:"timeframe" api:"required"`
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
	Parameters param.Field[ObservabilityTelemetryQueryParamsParameters] `json:"parameters"`
	// Controls the shape of the response. 'events': individual log lines matching the
	// query. 'calculations': aggregated metrics (count, avg, p99, etc.) with optional
	// group-by breakdowns and time-series. 'invocations': events grouped by request
	// ID. 'traces': distributed trace summaries. 'agents': Durable Object agent
	// summaries.
	View param.Field[ObservabilityTelemetryQueryParamsView] `json:"view"`
}

func (r ObservabilityTelemetryQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Timeframe for the query using Unix timestamps in milliseconds. Narrower
// timeframes produce faster responses and more specific results.
type ObservabilityTelemetryQueryParamsTimeframe struct {
	// Start timestamp for the query timeframe (Unix timestamp in milliseconds)
	From param.Field[float64] `json:"from" api:"required"`
	// End timestamp for the query timeframe (Unix timestamp in milliseconds)
	To param.Field[float64] `json:"to" api:"required"`
}

func (r ObservabilityTelemetryQueryParamsTimeframe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Query parameters defining what data to retrieve — filters, calculations,
// group-bys, and ordering. In practice this should always be provided for ad-hoc
// queries. Only omit when executing a previously saved query by queryId. Use the
// keys and values endpoints to discover available fields before building filters.
type ObservabilityTelemetryQueryParamsParameters struct {
	// Aggregation calculations to compute (e.g. count, avg, p99). Each calculation
	// produces aggregate values and optional time-series data.
	Calculations param.Field[[]ObservabilityTelemetryQueryParamsParametersCalculation] `json:"calculations"`
	// Datasets to query. Leave empty to query all available datasets.
	Datasets param.Field[[]string] `json:"datasets"`
	// Logical operator for combining top-level filters: 'and' (all must match) or 'or'
	// (any must match). Defaults to 'and'.
	FilterCombination param.Field[ObservabilityTelemetryQueryParamsParametersFilterCombination] `json:"filterCombination"`
	// Filters to narrow query results. Use the keys and values endpoints to discover
	// available fields before building filters. Supports nested groups via kind:
	// 'group'. Maximum nesting depth is 4.
	Filters param.Field[[]ObservabilityTelemetryQueryParamsParametersFilterUnion] `json:"filters"`
	// Fields to group calculation results by. Only applicable when the query view is
	// 'calculations'. Produces per-group aggregate values.
	GroupBys param.Field[[]ObservabilityTelemetryQueryParamsParametersGroupBy] `json:"groupBys"`
	// Post-aggregation filters applied to calculation results. Use to filter groups
	// after aggregation (e.g. only groups where count > 100).
	Havings param.Field[[]ObservabilityTelemetryQueryParamsParametersHaving] `json:"havings"`
	// Maximum number of group-by rows to return in calculation results. A value of 10
	// is a sensible default for most use cases.
	Limit param.Field[int64] `json:"limit"`
	// Full-text search expression applied across all event fields. Matches events
	// containing the specified text.
	Needle param.Field[ObservabilityTelemetryQueryParamsParametersNeedle] `json:"needle"`
	// Ordering for grouped calculation results. Only effective when a group-by is
	// present.
	OrderBy param.Field[ObservabilityTelemetryQueryParamsParametersOrderBy] `json:"orderBy"`
}

func (r ObservabilityTelemetryQueryParamsParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryParamsParametersCalculation struct {
	// Aggregation operator to apply. Examples: count, avg, sum, min, max, p50, p90,
	// p95, p99, uniq, stddev, variance.
	Operator param.Field[ObservabilityTelemetryQueryParamsParametersCalculationsOperator] `json:"operator" api:"required"`
	// Custom label for this calculation in the results. Useful for distinguishing
	// multiple calculations.
	Alias param.Field[string] `json:"alias"`
	// Field name to calculate over. Must exist in the data — verify with the keys
	// endpoint. Omit for operators that don't require a key (e.g. count).
	Key param.Field[string] `json:"key"`
	// Data type of the key. Required when key is provided to ensure correct
	// aggregation.
	KeyType param.Field[ObservabilityTelemetryQueryParamsParametersCalculationsKeyType] `json:"keyType"`
}

func (r ObservabilityTelemetryQueryParamsParametersCalculation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Aggregation operator to apply. Examples: count, avg, sum, min, max, p50, p90,
// p95, p99, uniq, stddev, variance.
type ObservabilityTelemetryQueryParamsParametersCalculationsOperator string

const (
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorUniq              ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "uniq"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorCount             ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "count"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMax               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "max"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMin               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "min"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorSum               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "sum"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorAvg               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "avg"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMedian            ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "median"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP001              ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p001"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP01               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p01"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP05               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p05"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP10               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p10"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP25               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p25"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP75               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p75"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP90               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p90"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP95               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p95"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP99               ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p99"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP999              ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "p999"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorStddev            ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "stddev"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorVariance          ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "variance"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorCountDistinct     ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "COUNT_DISTINCT"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorCountUppercase    ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "COUNT"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMaxUppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "MAX"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMinUppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "MIN"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorSumUppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "SUM"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorAvgUppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "AVG"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMedianUppercase   ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "MEDIAN"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP001Uppercase     ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P001"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP01Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P01"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP05Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P05"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP10Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P10"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP25Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P25"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP75Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P75"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP90Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P90"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP95Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P95"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP99Uppercase      ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P99"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP999Uppercase     ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "P999"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorStddevUppercase   ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "STDDEV"
	ObservabilityTelemetryQueryParamsParametersCalculationsOperatorVarianceUppercase ObservabilityTelemetryQueryParamsParametersCalculationsOperator = "VARIANCE"
)

func (r ObservabilityTelemetryQueryParamsParametersCalculationsOperator) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersCalculationsOperatorUniq, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorCount, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMax, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMin, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorSum, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorAvg, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMedian, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP001, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP01, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP05, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP10, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP25, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP75, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP90, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP95, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP99, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP999, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorStddev, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorVariance, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorCountDistinct, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorCountUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMaxUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMinUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorSumUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorAvgUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorMedianUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP001Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP01Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP05Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP10Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP25Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP75Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP90Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP95Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP99Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorP999Uppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorStddevUppercase, ObservabilityTelemetryQueryParamsParametersCalculationsOperatorVarianceUppercase:
		return true
	}
	return false
}

// Data type of the key. Required when key is provided to ensure correct
// aggregation.
type ObservabilityTelemetryQueryParamsParametersCalculationsKeyType string

const (
	ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeString  ObservabilityTelemetryQueryParamsParametersCalculationsKeyType = "string"
	ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeNumber  ObservabilityTelemetryQueryParamsParametersCalculationsKeyType = "number"
	ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeBoolean ObservabilityTelemetryQueryParamsParametersCalculationsKeyType = "boolean"
)

func (r ObservabilityTelemetryQueryParamsParametersCalculationsKeyType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeString, ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeNumber, ObservabilityTelemetryQueryParamsParametersCalculationsKeyTypeBoolean:
		return true
	}
	return false
}

// Logical operator for combining top-level filters: 'and' (all must match) or 'or'
// (any must match). Defaults to 'and'.
type ObservabilityTelemetryQueryParamsParametersFilterCombination string

const (
	ObservabilityTelemetryQueryParamsParametersFilterCombinationAnd          ObservabilityTelemetryQueryParamsParametersFilterCombination = "and"
	ObservabilityTelemetryQueryParamsParametersFilterCombinationOr           ObservabilityTelemetryQueryParamsParametersFilterCombination = "or"
	ObservabilityTelemetryQueryParamsParametersFilterCombinationAndUppercase ObservabilityTelemetryQueryParamsParametersFilterCombination = "AND"
	ObservabilityTelemetryQueryParamsParametersFilterCombinationOrUppercase  ObservabilityTelemetryQueryParamsParametersFilterCombination = "OR"
)

func (r ObservabilityTelemetryQueryParamsParametersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFilterCombinationAnd, ObservabilityTelemetryQueryParamsParametersFilterCombinationOr, ObservabilityTelemetryQueryParamsParametersFilterCombinationAndUppercase, ObservabilityTelemetryQueryParamsParametersFilterCombinationOrUppercase:
		return true
	}
	return false
}

// Supports nested groups via kind: 'group'.
type ObservabilityTelemetryQueryParamsParametersFilter struct {
	FilterCombination param.Field[ObservabilityTelemetryQueryParamsParametersFiltersFilterCombination] `json:"filterCombination"`
	Filters           param.Field[interface{}]                                                         `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  param.Field[string]                                                 `json:"key"`
	Kind param.Field[ObservabilityTelemetryQueryParamsParametersFiltersKind] `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryQueryParamsParametersFiltersOperation] `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type  param.Field[ObservabilityTelemetryQueryParamsParametersFiltersType] `json:"type"`
	Value param.Field[interface{}]                                            `json:"value"`
}

func (r ObservabilityTelemetryQueryParamsParametersFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryQueryParamsParametersFilter) implementsObservabilityTelemetryQueryParamsParametersFilterUnion() {
}

// Supports nested groups via kind: 'group'.
//
// Satisfied by [workers.ObservabilityTelemetryQueryParamsParametersFiltersObject],
// [workers.ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeaf],
// [ObservabilityTelemetryQueryParamsParametersFilter].
type ObservabilityTelemetryQueryParamsParametersFilterUnion interface {
	implementsObservabilityTelemetryQueryParamsParametersFilterUnion()
}

type ObservabilityTelemetryQueryParamsParametersFiltersObject struct {
	FilterCombination param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombination] `json:"filterCombination" api:"required"`
	Filters           param.Field[[]ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterUnion]     `json:"filters" api:"required"`
	Kind              param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectKind]              `json:"kind" api:"required"`
}

func (r ObservabilityTelemetryQueryParamsParametersFiltersObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryQueryParamsParametersFiltersObject) implementsObservabilityTelemetryQueryParamsParametersFilterUnion() {
}

type ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombination string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombinationAnd          ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombination = "and"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombinationOr           ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombination = "or"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombinationAndUppercase ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombination = "AND"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombinationOrUppercase  ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombination = "OR"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombinationAnd, ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombinationOr, ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombinationAndUppercase, ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

// Supports nested groups via kind: 'group'.
type ObservabilityTelemetryQueryParamsParametersFiltersObjectFilter struct {
	FilterCombination param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombination] `json:"filterCombination"`
	Filters           param.Field[interface{}]                                                                      `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  param.Field[string]                                                              `json:"key"`
	Kind param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersKind] `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation] `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type  param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersType] `json:"type"`
	Value param.Field[interface{}]                                                         `json:"value"`
}

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFilter) implementsObservabilityTelemetryQueryParamsParametersFiltersObjectFilterUnion() {
}

// Supports nested groups via kind: 'group'.
//
// Satisfied by
// [workers.ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObject],
// [workers.ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeaf],
// [ObservabilityTelemetryQueryParamsParametersFiltersObjectFilter].
type ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterUnion interface {
	implementsObservabilityTelemetryQueryParamsParametersFiltersObjectFilterUnion()
}

type ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObject struct {
	FilterCombination param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombination] `json:"filterCombination" api:"required"`
	Filters           param.Field[[]interface{}]                                                                          `json:"filters" api:"required"`
	Kind              param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectKind]              `json:"kind" api:"required"`
}

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObject) implementsObservabilityTelemetryQueryParamsParametersFiltersObjectFilterUnion() {
}

type ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombination string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombinationAnd          ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombination = "and"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombinationOr           ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombination = "or"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombinationAndUppercase ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombination = "AND"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombinationOrUppercase  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombination = "OR"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombinationAnd, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombinationOr, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombinationAndUppercase, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectKind string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectKindGroup ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectKind = "group"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key param.Field[string] `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation] `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafType] `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafKind] `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeaf) implementsObservabilityTelemetryQueryParamsParametersFiltersObjectFilterUnion() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafTypeString  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafTypeString, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafKindFilter ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafKindFilter:
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
type ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion()
}

type ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombination string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombinationAnd          ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombination = "and"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombinationOr           ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombination = "or"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombinationAndUppercase ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombination = "AND"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombinationOrUppercase  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombination = "OR"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombinationAnd, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombinationOr, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombinationAndUppercase, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersKind string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersKindGroup  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersKind = "group"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersKindFilter ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersKind = "filter"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersKindGroup, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationIncludes            ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "includes"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationNotIncludes         ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "not_includes"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationStartsWith          ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "starts_with"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationRegex               ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "regex"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationExists              ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "exists"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationIsNull              ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "is_null"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationIn                  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "in"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationNotIn               ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "not_in"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationEq                  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "eq"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationNeq                 ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "neq"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationGt                  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "gt"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationGte                 ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "gte"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationLt                  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "lt"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationLte                 ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "lte"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationEquals              ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "="
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationNotEquals           ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "!="
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationGreater             ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = ">"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationGreaterOrEquals     ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = ">="
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationLess                ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "<"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationLessOrEquals        ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "<="
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationIncludesUppercase   ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "INCLUDES"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationDoesNotInclude      ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationMatchRegex          ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationExistsUppercase     ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "EXISTS"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationDoesNotExist        ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationInUppercase         ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "IN"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationNotInUppercase      ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "NOT_IN"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationStartsWithUppercase ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationIncludes, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationNotIncludes, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationStartsWith, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationRegex, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationExists, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationIsNull, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationIn, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationNotIn, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationEq, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationNeq, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationGt, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationGte, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationLt, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationLte, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationEquals, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationNotEquals, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationGreater, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationGreaterOrEquals, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationLess, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationLessOrEquals, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationIncludesUppercase, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationDoesNotInclude, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationMatchRegex, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationExistsUppercase, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationDoesNotExist, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationInUppercase, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationNotInUppercase, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersType string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersTypeString  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersType = "string"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersTypeNumber  ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersType = "number"
	ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersTypeBoolean ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersType = "boolean"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersTypeString, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersTypeNumber, ObservabilityTelemetryQueryParamsParametersFiltersObjectFiltersTypeBoolean:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryParamsParametersFiltersObjectKind string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersObjectKindGroup ObservabilityTelemetryQueryParamsParametersFiltersObjectKind = "group"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key param.Field[string] `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation] `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type param.Field[ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafType] `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind param.Field[ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafKind] `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value param.Field[ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeaf) implementsObservabilityTelemetryQueryParamsParametersFilterUnion() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafTypeString  ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafTypeString, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafKindFilter ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafKindFilter:
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
type ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion()
}

type ObservabilityTelemetryQueryParamsParametersFiltersFilterCombination string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersFilterCombinationAnd          ObservabilityTelemetryQueryParamsParametersFiltersFilterCombination = "and"
	ObservabilityTelemetryQueryParamsParametersFiltersFilterCombinationOr           ObservabilityTelemetryQueryParamsParametersFiltersFilterCombination = "or"
	ObservabilityTelemetryQueryParamsParametersFiltersFilterCombinationAndUppercase ObservabilityTelemetryQueryParamsParametersFiltersFilterCombination = "AND"
	ObservabilityTelemetryQueryParamsParametersFiltersFilterCombinationOrUppercase  ObservabilityTelemetryQueryParamsParametersFiltersFilterCombination = "OR"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersFilterCombinationAnd, ObservabilityTelemetryQueryParamsParametersFiltersFilterCombinationOr, ObservabilityTelemetryQueryParamsParametersFiltersFilterCombinationAndUppercase, ObservabilityTelemetryQueryParamsParametersFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryParamsParametersFiltersKind string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersKindGroup  ObservabilityTelemetryQueryParamsParametersFiltersKind = "group"
	ObservabilityTelemetryQueryParamsParametersFiltersKindFilter ObservabilityTelemetryQueryParamsParametersFiltersKind = "filter"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersKindGroup, ObservabilityTelemetryQueryParamsParametersFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryQueryParamsParametersFiltersOperation string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersOperationIncludes            ObservabilityTelemetryQueryParamsParametersFiltersOperation = "includes"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationNotIncludes         ObservabilityTelemetryQueryParamsParametersFiltersOperation = "not_includes"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationStartsWith          ObservabilityTelemetryQueryParamsParametersFiltersOperation = "starts_with"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationRegex               ObservabilityTelemetryQueryParamsParametersFiltersOperation = "regex"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationExists              ObservabilityTelemetryQueryParamsParametersFiltersOperation = "exists"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationIsNull              ObservabilityTelemetryQueryParamsParametersFiltersOperation = "is_null"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationIn                  ObservabilityTelemetryQueryParamsParametersFiltersOperation = "in"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationNotIn               ObservabilityTelemetryQueryParamsParametersFiltersOperation = "not_in"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationEq                  ObservabilityTelemetryQueryParamsParametersFiltersOperation = "eq"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationNeq                 ObservabilityTelemetryQueryParamsParametersFiltersOperation = "neq"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationGt                  ObservabilityTelemetryQueryParamsParametersFiltersOperation = "gt"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationGte                 ObservabilityTelemetryQueryParamsParametersFiltersOperation = "gte"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationLt                  ObservabilityTelemetryQueryParamsParametersFiltersOperation = "lt"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationLte                 ObservabilityTelemetryQueryParamsParametersFiltersOperation = "lte"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationEquals              ObservabilityTelemetryQueryParamsParametersFiltersOperation = "="
	ObservabilityTelemetryQueryParamsParametersFiltersOperationNotEquals           ObservabilityTelemetryQueryParamsParametersFiltersOperation = "!="
	ObservabilityTelemetryQueryParamsParametersFiltersOperationGreater             ObservabilityTelemetryQueryParamsParametersFiltersOperation = ">"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationGreaterOrEquals     ObservabilityTelemetryQueryParamsParametersFiltersOperation = ">="
	ObservabilityTelemetryQueryParamsParametersFiltersOperationLess                ObservabilityTelemetryQueryParamsParametersFiltersOperation = "<"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationLessOrEquals        ObservabilityTelemetryQueryParamsParametersFiltersOperation = "<="
	ObservabilityTelemetryQueryParamsParametersFiltersOperationIncludesUppercase   ObservabilityTelemetryQueryParamsParametersFiltersOperation = "INCLUDES"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationDoesNotInclude      ObservabilityTelemetryQueryParamsParametersFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationMatchRegex          ObservabilityTelemetryQueryParamsParametersFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationExistsUppercase     ObservabilityTelemetryQueryParamsParametersFiltersOperation = "EXISTS"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationDoesNotExist        ObservabilityTelemetryQueryParamsParametersFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationInUppercase         ObservabilityTelemetryQueryParamsParametersFiltersOperation = "IN"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationNotInUppercase      ObservabilityTelemetryQueryParamsParametersFiltersOperation = "NOT_IN"
	ObservabilityTelemetryQueryParamsParametersFiltersOperationStartsWithUppercase ObservabilityTelemetryQueryParamsParametersFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersOperationIncludes, ObservabilityTelemetryQueryParamsParametersFiltersOperationNotIncludes, ObservabilityTelemetryQueryParamsParametersFiltersOperationStartsWith, ObservabilityTelemetryQueryParamsParametersFiltersOperationRegex, ObservabilityTelemetryQueryParamsParametersFiltersOperationExists, ObservabilityTelemetryQueryParamsParametersFiltersOperationIsNull, ObservabilityTelemetryQueryParamsParametersFiltersOperationIn, ObservabilityTelemetryQueryParamsParametersFiltersOperationNotIn, ObservabilityTelemetryQueryParamsParametersFiltersOperationEq, ObservabilityTelemetryQueryParamsParametersFiltersOperationNeq, ObservabilityTelemetryQueryParamsParametersFiltersOperationGt, ObservabilityTelemetryQueryParamsParametersFiltersOperationGte, ObservabilityTelemetryQueryParamsParametersFiltersOperationLt, ObservabilityTelemetryQueryParamsParametersFiltersOperationLte, ObservabilityTelemetryQueryParamsParametersFiltersOperationEquals, ObservabilityTelemetryQueryParamsParametersFiltersOperationNotEquals, ObservabilityTelemetryQueryParamsParametersFiltersOperationGreater, ObservabilityTelemetryQueryParamsParametersFiltersOperationGreaterOrEquals, ObservabilityTelemetryQueryParamsParametersFiltersOperationLess, ObservabilityTelemetryQueryParamsParametersFiltersOperationLessOrEquals, ObservabilityTelemetryQueryParamsParametersFiltersOperationIncludesUppercase, ObservabilityTelemetryQueryParamsParametersFiltersOperationDoesNotInclude, ObservabilityTelemetryQueryParamsParametersFiltersOperationMatchRegex, ObservabilityTelemetryQueryParamsParametersFiltersOperationExistsUppercase, ObservabilityTelemetryQueryParamsParametersFiltersOperationDoesNotExist, ObservabilityTelemetryQueryParamsParametersFiltersOperationInUppercase, ObservabilityTelemetryQueryParamsParametersFiltersOperationNotInUppercase, ObservabilityTelemetryQueryParamsParametersFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryQueryParamsParametersFiltersType string

const (
	ObservabilityTelemetryQueryParamsParametersFiltersTypeString  ObservabilityTelemetryQueryParamsParametersFiltersType = "string"
	ObservabilityTelemetryQueryParamsParametersFiltersTypeNumber  ObservabilityTelemetryQueryParamsParametersFiltersType = "number"
	ObservabilityTelemetryQueryParamsParametersFiltersTypeBoolean ObservabilityTelemetryQueryParamsParametersFiltersType = "boolean"
)

func (r ObservabilityTelemetryQueryParamsParametersFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersFiltersTypeString, ObservabilityTelemetryQueryParamsParametersFiltersTypeNumber, ObservabilityTelemetryQueryParamsParametersFiltersTypeBoolean:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryParamsParametersGroupBy struct {
	// Data type of the group-by field.
	Type param.Field[ObservabilityTelemetryQueryParamsParametersGroupBysType] `json:"type" api:"required"`
	// Field name to group results by (e.g. $metadata.service, $metadata.statusCode).
	Value param.Field[string] `json:"value" api:"required"`
}

func (r ObservabilityTelemetryQueryParamsParametersGroupBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Data type of the group-by field.
type ObservabilityTelemetryQueryParamsParametersGroupBysType string

const (
	ObservabilityTelemetryQueryParamsParametersGroupBysTypeString  ObservabilityTelemetryQueryParamsParametersGroupBysType = "string"
	ObservabilityTelemetryQueryParamsParametersGroupBysTypeNumber  ObservabilityTelemetryQueryParamsParametersGroupBysType = "number"
	ObservabilityTelemetryQueryParamsParametersGroupBysTypeBoolean ObservabilityTelemetryQueryParamsParametersGroupBysType = "boolean"
)

func (r ObservabilityTelemetryQueryParamsParametersGroupBysType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersGroupBysTypeString, ObservabilityTelemetryQueryParamsParametersGroupBysTypeNumber, ObservabilityTelemetryQueryParamsParametersGroupBysTypeBoolean:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryParamsParametersHaving struct {
	// Calculation alias or operator to filter on after aggregation.
	Key param.Field[string] `json:"key" api:"required"`
	// Numeric comparison operator: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryQueryParamsParametersHavingsOperation] `json:"operation" api:"required"`
	// Threshold value to compare the calculation result against.
	Value param.Field[float64] `json:"value" api:"required"`
}

func (r ObservabilityTelemetryQueryParamsParametersHaving) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Numeric comparison operator: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryQueryParamsParametersHavingsOperation string

const (
	ObservabilityTelemetryQueryParamsParametersHavingsOperationEq  ObservabilityTelemetryQueryParamsParametersHavingsOperation = "eq"
	ObservabilityTelemetryQueryParamsParametersHavingsOperationNeq ObservabilityTelemetryQueryParamsParametersHavingsOperation = "neq"
	ObservabilityTelemetryQueryParamsParametersHavingsOperationGt  ObservabilityTelemetryQueryParamsParametersHavingsOperation = "gt"
	ObservabilityTelemetryQueryParamsParametersHavingsOperationGte ObservabilityTelemetryQueryParamsParametersHavingsOperation = "gte"
	ObservabilityTelemetryQueryParamsParametersHavingsOperationLt  ObservabilityTelemetryQueryParamsParametersHavingsOperation = "lt"
	ObservabilityTelemetryQueryParamsParametersHavingsOperationLte ObservabilityTelemetryQueryParamsParametersHavingsOperation = "lte"
)

func (r ObservabilityTelemetryQueryParamsParametersHavingsOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersHavingsOperationEq, ObservabilityTelemetryQueryParamsParametersHavingsOperationNeq, ObservabilityTelemetryQueryParamsParametersHavingsOperationGt, ObservabilityTelemetryQueryParamsParametersHavingsOperationGte, ObservabilityTelemetryQueryParamsParametersHavingsOperationLt, ObservabilityTelemetryQueryParamsParametersHavingsOperationLte:
		return true
	}
	return false
}

// Full-text search expression applied across all event fields. Matches events
// containing the specified text.
type ObservabilityTelemetryQueryParamsParametersNeedle struct {
	// The text or pattern to search for.
	Value param.Field[ObservabilityTelemetryQueryParamsParametersNeedleValueUnion] `json:"value" api:"required"`
	// When true, treats the value as a regular expression (RE2 syntax).
	IsRegex param.Field[bool] `json:"isRegex"`
	// When true, performs a case-sensitive search. Defaults to case-insensitive.
	MatchCase param.Field[bool] `json:"matchCase"`
}

func (r ObservabilityTelemetryQueryParamsParametersNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The text or pattern to search for.
//
// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryQueryParamsParametersNeedleValueUnion interface {
	ImplementsObservabilityTelemetryQueryParamsParametersNeedleValueUnion()
}

// Ordering for grouped calculation results. Only effective when a group-by is
// present.
type ObservabilityTelemetryQueryParamsParametersOrderBy struct {
	// Alias of the calculation to order results by. Must match the alias (or operator)
	// of a calculation in the query.
	Value param.Field[string] `json:"value" api:"required"`
	// Sort direction: 'asc' for ascending, 'desc' for descending.
	Order param.Field[ObservabilityTelemetryQueryParamsParametersOrderByOrder] `json:"order"`
}

func (r ObservabilityTelemetryQueryParamsParametersOrderBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Sort direction: 'asc' for ascending, 'desc' for descending.
type ObservabilityTelemetryQueryParamsParametersOrderByOrder string

const (
	ObservabilityTelemetryQueryParamsParametersOrderByOrderAsc  ObservabilityTelemetryQueryParamsParametersOrderByOrder = "asc"
	ObservabilityTelemetryQueryParamsParametersOrderByOrderDesc ObservabilityTelemetryQueryParamsParametersOrderByOrder = "desc"
)

func (r ObservabilityTelemetryQueryParamsParametersOrderByOrder) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsParametersOrderByOrderAsc, ObservabilityTelemetryQueryParamsParametersOrderByOrderDesc:
		return true
	}
	return false
}

// Controls the shape of the response. 'events': individual log lines matching the
// query. 'calculations': aggregated metrics (count, avg, p99, etc.) with optional
// group-by breakdowns and time-series. 'invocations': events grouped by request
// ID. 'traces': distributed trace summaries. 'agents': Durable Object agent
// summaries.
type ObservabilityTelemetryQueryParamsView string

const (
	ObservabilityTelemetryQueryParamsViewTraces       ObservabilityTelemetryQueryParamsView = "traces"
	ObservabilityTelemetryQueryParamsViewEvents       ObservabilityTelemetryQueryParamsView = "events"
	ObservabilityTelemetryQueryParamsViewCalculations ObservabilityTelemetryQueryParamsView = "calculations"
	ObservabilityTelemetryQueryParamsViewInvocations  ObservabilityTelemetryQueryParamsView = "invocations"
	ObservabilityTelemetryQueryParamsViewRequests     ObservabilityTelemetryQueryParamsView = "requests"
	ObservabilityTelemetryQueryParamsViewAgents       ObservabilityTelemetryQueryParamsView = "agents"
)

func (r ObservabilityTelemetryQueryParamsView) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryParamsViewTraces, ObservabilityTelemetryQueryParamsViewEvents, ObservabilityTelemetryQueryParamsViewCalculations, ObservabilityTelemetryQueryParamsViewInvocations, ObservabilityTelemetryQueryParamsViewRequests, ObservabilityTelemetryQueryParamsViewAgents:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseEnvelope struct {
	Errors   []ObservabilityTelemetryQueryResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityTelemetryQueryResponseEnvelopeMessages `json:"messages" api:"required"`
	// Complete results of a query run. The populated fields depend on the requested
	// view type (events, calculations, invocations, traces, or agents).
	Result  ObservabilityTelemetryQueryResponse                `json:"result" api:"required"`
	Success ObservabilityTelemetryQueryResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    observabilityTelemetryQueryResponseEnvelopeJSON    `json:"-"`
}

// observabilityTelemetryQueryResponseEnvelopeJSON contains the JSON metadata for
// the struct [ObservabilityTelemetryQueryResponseEnvelope]
type observabilityTelemetryQueryResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEnvelopeErrors struct {
	Message string                                                `json:"message" api:"required"`
	JSON    observabilityTelemetryQueryResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityTelemetryQueryResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ObservabilityTelemetryQueryResponseEnvelopeErrors]
type observabilityTelemetryQueryResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEnvelopeMessages struct {
	Message ObservabilityTelemetryQueryResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityTelemetryQueryResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityTelemetryQueryResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityTelemetryQueryResponseEnvelopeMessages]
type observabilityTelemetryQueryResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryQueryResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryQueryResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryQueryResponseEnvelopeMessagesMessage string

const (
	ObservabilityTelemetryQueryResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityTelemetryQueryResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityTelemetryQueryResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityTelemetryQueryResponseEnvelopeSuccess bool

const (
	ObservabilityTelemetryQueryResponseEnvelopeSuccessTrue ObservabilityTelemetryQueryResponseEnvelopeSuccess = true
)

func (r ObservabilityTelemetryQueryResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryQueryResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityTelemetryValuesParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Leave this empty to use the default datasets
	Datasets  param.Field[[]string]                                    `json:"datasets" api:"required"`
	Key       param.Field[string]                                      `json:"key" api:"required"`
	Timeframe param.Field[ObservabilityTelemetryValuesParamsTimeframe] `json:"timeframe" api:"required"`
	Type      param.Field[ObservabilityTelemetryValuesParamsType]      `json:"type" api:"required"`
	// Apply filters before listing values. Supports nested groups via kind: 'group'.
	// Maximum nesting depth is 4.
	Filters param.Field[[]ObservabilityTelemetryValuesParamsFilterUnion] `json:"filters"`
	Limit   param.Field[float64]                                         `json:"limit"`
	// Full-text search expression to match events containing the specified text or
	// pattern.
	Needle param.Field[ObservabilityTelemetryValuesParamsNeedle] `json:"needle"`
}

func (r ObservabilityTelemetryValuesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryValuesParamsTimeframe struct {
	From param.Field[float64] `json:"from" api:"required"`
	To   param.Field[float64] `json:"to" api:"required"`
}

func (r ObservabilityTelemetryValuesParamsTimeframe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryValuesParamsType string

const (
	ObservabilityTelemetryValuesParamsTypeString  ObservabilityTelemetryValuesParamsType = "string"
	ObservabilityTelemetryValuesParamsTypeBoolean ObservabilityTelemetryValuesParamsType = "boolean"
	ObservabilityTelemetryValuesParamsTypeNumber  ObservabilityTelemetryValuesParamsType = "number"
)

func (r ObservabilityTelemetryValuesParamsType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsTypeString, ObservabilityTelemetryValuesParamsTypeBoolean, ObservabilityTelemetryValuesParamsTypeNumber:
		return true
	}
	return false
}

// Supports nested groups via kind: 'group'.
type ObservabilityTelemetryValuesParamsFilter struct {
	FilterCombination param.Field[ObservabilityTelemetryValuesParamsFiltersFilterCombination] `json:"filterCombination"`
	Filters           param.Field[interface{}]                                                `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  param.Field[string]                                        `json:"key"`
	Kind param.Field[ObservabilityTelemetryValuesParamsFiltersKind] `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryValuesParamsFiltersOperation] `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type  param.Field[ObservabilityTelemetryValuesParamsFiltersType] `json:"type"`
	Value param.Field[interface{}]                                   `json:"value"`
}

func (r ObservabilityTelemetryValuesParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryValuesParamsFilter) implementsObservabilityTelemetryValuesParamsFilterUnion() {
}

// Supports nested groups via kind: 'group'.
//
// Satisfied by [workers.ObservabilityTelemetryValuesParamsFiltersObject],
// [workers.ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeaf],
// [ObservabilityTelemetryValuesParamsFilter].
type ObservabilityTelemetryValuesParamsFilterUnion interface {
	implementsObservabilityTelemetryValuesParamsFilterUnion()
}

type ObservabilityTelemetryValuesParamsFiltersObject struct {
	FilterCombination param.Field[ObservabilityTelemetryValuesParamsFiltersObjectFilterCombination] `json:"filterCombination" api:"required"`
	Filters           param.Field[[]ObservabilityTelemetryValuesParamsFiltersObjectFilterUnion]     `json:"filters" api:"required"`
	Kind              param.Field[ObservabilityTelemetryValuesParamsFiltersObjectKind]              `json:"kind" api:"required"`
}

func (r ObservabilityTelemetryValuesParamsFiltersObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryValuesParamsFiltersObject) implementsObservabilityTelemetryValuesParamsFilterUnion() {
}

type ObservabilityTelemetryValuesParamsFiltersObjectFilterCombination string

const (
	ObservabilityTelemetryValuesParamsFiltersObjectFilterCombinationAnd          ObservabilityTelemetryValuesParamsFiltersObjectFilterCombination = "and"
	ObservabilityTelemetryValuesParamsFiltersObjectFilterCombinationOr           ObservabilityTelemetryValuesParamsFiltersObjectFilterCombination = "or"
	ObservabilityTelemetryValuesParamsFiltersObjectFilterCombinationAndUppercase ObservabilityTelemetryValuesParamsFiltersObjectFilterCombination = "AND"
	ObservabilityTelemetryValuesParamsFiltersObjectFilterCombinationOrUppercase  ObservabilityTelemetryValuesParamsFiltersObjectFilterCombination = "OR"
)

func (r ObservabilityTelemetryValuesParamsFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersObjectFilterCombinationAnd, ObservabilityTelemetryValuesParamsFiltersObjectFilterCombinationOr, ObservabilityTelemetryValuesParamsFiltersObjectFilterCombinationAndUppercase, ObservabilityTelemetryValuesParamsFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

// Supports nested groups via kind: 'group'.
type ObservabilityTelemetryValuesParamsFiltersObjectFilter struct {
	FilterCombination param.Field[ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombination] `json:"filterCombination"`
	Filters           param.Field[interface{}]                                                             `json:"filters"`
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key  param.Field[string]                                                     `json:"key"`
	Kind param.Field[ObservabilityTelemetryValuesParamsFiltersObjectFiltersKind] `json:"kind"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation] `json:"operation"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type  param.Field[ObservabilityTelemetryValuesParamsFiltersObjectFiltersType] `json:"type"`
	Value param.Field[interface{}]                                                `json:"value"`
}

func (r ObservabilityTelemetryValuesParamsFiltersObjectFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryValuesParamsFiltersObjectFilter) implementsObservabilityTelemetryValuesParamsFiltersObjectFilterUnion() {
}

// Supports nested groups via kind: 'group'.
//
// Satisfied by
// [workers.ObservabilityTelemetryValuesParamsFiltersObjectFiltersObject],
// [workers.ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeaf],
// [ObservabilityTelemetryValuesParamsFiltersObjectFilter].
type ObservabilityTelemetryValuesParamsFiltersObjectFilterUnion interface {
	implementsObservabilityTelemetryValuesParamsFiltersObjectFilterUnion()
}

type ObservabilityTelemetryValuesParamsFiltersObjectFiltersObject struct {
	FilterCombination param.Field[ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombination] `json:"filterCombination" api:"required"`
	Filters           param.Field[[]interface{}]                                                                 `json:"filters" api:"required"`
	Kind              param.Field[ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectKind]              `json:"kind" api:"required"`
}

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersObject) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersObject) implementsObservabilityTelemetryValuesParamsFiltersObjectFilterUnion() {
}

type ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombination string

const (
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombinationAnd          ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombination = "and"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombinationOr           ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombination = "or"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombinationAndUppercase ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombination = "AND"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombinationOrUppercase  ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombination = "OR"
)

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombinationAnd, ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombinationOr, ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombinationAndUppercase, ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectKind string

const (
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectKindGroup ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectKind = "group"
)

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersObjectFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key param.Field[string] `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation] `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type param.Field[ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafType] `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind param.Field[ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafKind] `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value param.Field[ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeaf) implementsObservabilityTelemetryValuesParamsFiltersObjectFilterUnion() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafTypeString  ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafTypeString, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafKindFilter ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafKindFilter:
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
type ObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilityTelemetryValuesParamsFiltersObjectFiltersWorkersObservabilityFilterLeafValueUnion()
}

type ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombination string

const (
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombinationAnd          ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombination = "and"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombinationOr           ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombination = "or"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombinationAndUppercase ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombination = "AND"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombinationOrUppercase  ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombination = "OR"
)

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombinationAnd, ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombinationOr, ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombinationAndUppercase, ObservabilityTelemetryValuesParamsFiltersObjectFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryValuesParamsFiltersObjectFiltersKind string

const (
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersKindGroup  ObservabilityTelemetryValuesParamsFiltersObjectFiltersKind = "group"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersKindFilter ObservabilityTelemetryValuesParamsFiltersObjectFiltersKind = "filter"
)

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersObjectFiltersKindGroup, ObservabilityTelemetryValuesParamsFiltersObjectFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation string

const (
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationIncludes            ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "includes"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationNotIncludes         ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "not_includes"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationStartsWith          ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "starts_with"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationRegex               ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "regex"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationExists              ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "exists"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationIsNull              ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "is_null"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationIn                  ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "in"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationNotIn               ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "not_in"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationEq                  ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "eq"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationNeq                 ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "neq"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationGt                  ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "gt"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationGte                 ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "gte"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationLt                  ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "lt"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationLte                 ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "lte"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationEquals              ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "="
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationNotEquals           ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "!="
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationGreater             ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = ">"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationGreaterOrEquals     ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = ">="
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationLess                ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "<"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationLessOrEquals        ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "<="
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationIncludesUppercase   ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "INCLUDES"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationDoesNotInclude      ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationMatchRegex          ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationExistsUppercase     ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "EXISTS"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationDoesNotExist        ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationInUppercase         ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "IN"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationNotInUppercase      ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "NOT_IN"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationStartsWithUppercase ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationIncludes, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationNotIncludes, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationStartsWith, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationRegex, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationExists, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationIsNull, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationIn, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationNotIn, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationEq, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationNeq, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationGt, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationGte, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationLt, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationLte, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationEquals, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationNotEquals, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationGreater, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationGreaterOrEquals, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationLess, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationLessOrEquals, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationIncludesUppercase, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationDoesNotInclude, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationMatchRegex, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationExistsUppercase, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationDoesNotExist, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationInUppercase, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationNotInUppercase, ObservabilityTelemetryValuesParamsFiltersObjectFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryValuesParamsFiltersObjectFiltersType string

const (
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersTypeString  ObservabilityTelemetryValuesParamsFiltersObjectFiltersType = "string"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersTypeNumber  ObservabilityTelemetryValuesParamsFiltersObjectFiltersType = "number"
	ObservabilityTelemetryValuesParamsFiltersObjectFiltersTypeBoolean ObservabilityTelemetryValuesParamsFiltersObjectFiltersType = "boolean"
)

func (r ObservabilityTelemetryValuesParamsFiltersObjectFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersObjectFiltersTypeString, ObservabilityTelemetryValuesParamsFiltersObjectFiltersTypeNumber, ObservabilityTelemetryValuesParamsFiltersObjectFiltersTypeBoolean:
		return true
	}
	return false
}

type ObservabilityTelemetryValuesParamsFiltersObjectKind string

const (
	ObservabilityTelemetryValuesParamsFiltersObjectKindGroup ObservabilityTelemetryValuesParamsFiltersObjectKind = "group"
)

func (r ObservabilityTelemetryValuesParamsFiltersObjectKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersObjectKindGroup:
		return true
	}
	return false
}

// A filter condition applied to query results. Use the keys and values endpoints
// to discover available fields and their values before constructing filters.
type ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. Use verified keys from previous query results or the keys
	// endpoint. Common keys include $metadata.service, $metadata.origin,
	// $metadata.trigger, $metadata.message, and $metadata.error.
	Key param.Field[string] `json:"key" api:"required"`
	// Comparison operator. String operators: includes, not_includes, starts_with,
	// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
	// values). Numeric: eq, neq, gt, gte, lt, lte.
	Operation param.Field[ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation] `json:"operation" api:"required"`
	// Data type of the filter field. Must match the actual type of the key being
	// filtered.
	Type param.Field[ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafType] `json:"type" api:"required"`
	// Discriminator for leaf filter nodes. Always 'filter' when present; may be
	// omitted.
	Kind param.Field[ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafKind] `json:"kind"`
	// Comparison value. Must match actual values in your data — verify with the values
	// endpoint. Ensure the value type (string/number/boolean) matches the field type.
	// String comparisons are case-sensitive. Regex uses RE2 syntax (no
	// lookaheads/lookbehinds).
	Value param.Field[ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeaf) implementsObservabilityTelemetryValuesParamsFilterUnion() {
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation string

const (
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationIncludes            ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "includes"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationNotIncludes         ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "not_includes"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationStartsWith          ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "starts_with"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationRegex               ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "regex"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationExists              ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "exists"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationIsNull              ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "is_null"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationIn                  ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "in"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationNotIn               ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "not_in"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationEq                  ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "eq"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationNeq                 ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "neq"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationGt                  ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "gt"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationGte                 ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "gte"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationLt                  ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "lt"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationLte                 ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "lte"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationEquals              ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "="
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationNotEquals           ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "!="
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationGreater             ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = ">"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals     ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = ">="
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationLess                ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "<"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationLessOrEquals        ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "<="
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase   ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "INCLUDES"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude      ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationMatchRegex          ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "MATCH_REGEX"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationExistsUppercase     ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "EXISTS"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationDoesNotExist        ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationInUppercase         ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "IN"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationNotInUppercase      ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "NOT_IN"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationIncludes, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationNotIncludes, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationStartsWith, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationRegex, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationExists, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationIsNull, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationIn, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationNotIn, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationEq, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationNeq, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationGt, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationGte, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationLt, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationLte, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationEquals, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationNotEquals, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationGreater, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationGreaterOrEquals, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationLess, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationLessOrEquals, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationIncludesUppercase, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationDoesNotInclude, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationMatchRegex, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationExistsUppercase, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationDoesNotExist, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationInUppercase, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationNotInUppercase, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafType string

const (
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafTypeString  ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafType = "string"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafTypeNumber  ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafType = "number"
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafTypeBoolean ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafType = "boolean"
)

func (r ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafTypeString, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafTypeNumber, ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafTypeBoolean:
		return true
	}
	return false
}

// Discriminator for leaf filter nodes. Always 'filter' when present; may be
// omitted.
type ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafKind string

const (
	ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafKindFilter ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafKind = "filter"
)

func (r ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafKindFilter:
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
type ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafValueUnion interface {
	ImplementsObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafValueUnion()
}

type ObservabilityTelemetryValuesParamsFiltersFilterCombination string

const (
	ObservabilityTelemetryValuesParamsFiltersFilterCombinationAnd          ObservabilityTelemetryValuesParamsFiltersFilterCombination = "and"
	ObservabilityTelemetryValuesParamsFiltersFilterCombinationOr           ObservabilityTelemetryValuesParamsFiltersFilterCombination = "or"
	ObservabilityTelemetryValuesParamsFiltersFilterCombinationAndUppercase ObservabilityTelemetryValuesParamsFiltersFilterCombination = "AND"
	ObservabilityTelemetryValuesParamsFiltersFilterCombinationOrUppercase  ObservabilityTelemetryValuesParamsFiltersFilterCombination = "OR"
)

func (r ObservabilityTelemetryValuesParamsFiltersFilterCombination) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersFilterCombinationAnd, ObservabilityTelemetryValuesParamsFiltersFilterCombinationOr, ObservabilityTelemetryValuesParamsFiltersFilterCombinationAndUppercase, ObservabilityTelemetryValuesParamsFiltersFilterCombinationOrUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryValuesParamsFiltersKind string

const (
	ObservabilityTelemetryValuesParamsFiltersKindGroup  ObservabilityTelemetryValuesParamsFiltersKind = "group"
	ObservabilityTelemetryValuesParamsFiltersKindFilter ObservabilityTelemetryValuesParamsFiltersKind = "filter"
)

func (r ObservabilityTelemetryValuesParamsFiltersKind) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersKindGroup, ObservabilityTelemetryValuesParamsFiltersKindFilter:
		return true
	}
	return false
}

// Comparison operator. String operators: includes, not_includes, starts_with,
// regex. Existence: exists, is_null. Set membership: in, not_in (comma-separated
// values). Numeric: eq, neq, gt, gte, lt, lte.
type ObservabilityTelemetryValuesParamsFiltersOperation string

const (
	ObservabilityTelemetryValuesParamsFiltersOperationIncludes            ObservabilityTelemetryValuesParamsFiltersOperation = "includes"
	ObservabilityTelemetryValuesParamsFiltersOperationNotIncludes         ObservabilityTelemetryValuesParamsFiltersOperation = "not_includes"
	ObservabilityTelemetryValuesParamsFiltersOperationStartsWith          ObservabilityTelemetryValuesParamsFiltersOperation = "starts_with"
	ObservabilityTelemetryValuesParamsFiltersOperationRegex               ObservabilityTelemetryValuesParamsFiltersOperation = "regex"
	ObservabilityTelemetryValuesParamsFiltersOperationExists              ObservabilityTelemetryValuesParamsFiltersOperation = "exists"
	ObservabilityTelemetryValuesParamsFiltersOperationIsNull              ObservabilityTelemetryValuesParamsFiltersOperation = "is_null"
	ObservabilityTelemetryValuesParamsFiltersOperationIn                  ObservabilityTelemetryValuesParamsFiltersOperation = "in"
	ObservabilityTelemetryValuesParamsFiltersOperationNotIn               ObservabilityTelemetryValuesParamsFiltersOperation = "not_in"
	ObservabilityTelemetryValuesParamsFiltersOperationEq                  ObservabilityTelemetryValuesParamsFiltersOperation = "eq"
	ObservabilityTelemetryValuesParamsFiltersOperationNeq                 ObservabilityTelemetryValuesParamsFiltersOperation = "neq"
	ObservabilityTelemetryValuesParamsFiltersOperationGt                  ObservabilityTelemetryValuesParamsFiltersOperation = "gt"
	ObservabilityTelemetryValuesParamsFiltersOperationGte                 ObservabilityTelemetryValuesParamsFiltersOperation = "gte"
	ObservabilityTelemetryValuesParamsFiltersOperationLt                  ObservabilityTelemetryValuesParamsFiltersOperation = "lt"
	ObservabilityTelemetryValuesParamsFiltersOperationLte                 ObservabilityTelemetryValuesParamsFiltersOperation = "lte"
	ObservabilityTelemetryValuesParamsFiltersOperationEquals              ObservabilityTelemetryValuesParamsFiltersOperation = "="
	ObservabilityTelemetryValuesParamsFiltersOperationNotEquals           ObservabilityTelemetryValuesParamsFiltersOperation = "!="
	ObservabilityTelemetryValuesParamsFiltersOperationGreater             ObservabilityTelemetryValuesParamsFiltersOperation = ">"
	ObservabilityTelemetryValuesParamsFiltersOperationGreaterOrEquals     ObservabilityTelemetryValuesParamsFiltersOperation = ">="
	ObservabilityTelemetryValuesParamsFiltersOperationLess                ObservabilityTelemetryValuesParamsFiltersOperation = "<"
	ObservabilityTelemetryValuesParamsFiltersOperationLessOrEquals        ObservabilityTelemetryValuesParamsFiltersOperation = "<="
	ObservabilityTelemetryValuesParamsFiltersOperationIncludesUppercase   ObservabilityTelemetryValuesParamsFiltersOperation = "INCLUDES"
	ObservabilityTelemetryValuesParamsFiltersOperationDoesNotInclude      ObservabilityTelemetryValuesParamsFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryValuesParamsFiltersOperationMatchRegex          ObservabilityTelemetryValuesParamsFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryValuesParamsFiltersOperationExistsUppercase     ObservabilityTelemetryValuesParamsFiltersOperation = "EXISTS"
	ObservabilityTelemetryValuesParamsFiltersOperationDoesNotExist        ObservabilityTelemetryValuesParamsFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryValuesParamsFiltersOperationInUppercase         ObservabilityTelemetryValuesParamsFiltersOperation = "IN"
	ObservabilityTelemetryValuesParamsFiltersOperationNotInUppercase      ObservabilityTelemetryValuesParamsFiltersOperation = "NOT_IN"
	ObservabilityTelemetryValuesParamsFiltersOperationStartsWithUppercase ObservabilityTelemetryValuesParamsFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryValuesParamsFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersOperationIncludes, ObservabilityTelemetryValuesParamsFiltersOperationNotIncludes, ObservabilityTelemetryValuesParamsFiltersOperationStartsWith, ObservabilityTelemetryValuesParamsFiltersOperationRegex, ObservabilityTelemetryValuesParamsFiltersOperationExists, ObservabilityTelemetryValuesParamsFiltersOperationIsNull, ObservabilityTelemetryValuesParamsFiltersOperationIn, ObservabilityTelemetryValuesParamsFiltersOperationNotIn, ObservabilityTelemetryValuesParamsFiltersOperationEq, ObservabilityTelemetryValuesParamsFiltersOperationNeq, ObservabilityTelemetryValuesParamsFiltersOperationGt, ObservabilityTelemetryValuesParamsFiltersOperationGte, ObservabilityTelemetryValuesParamsFiltersOperationLt, ObservabilityTelemetryValuesParamsFiltersOperationLte, ObservabilityTelemetryValuesParamsFiltersOperationEquals, ObservabilityTelemetryValuesParamsFiltersOperationNotEquals, ObservabilityTelemetryValuesParamsFiltersOperationGreater, ObservabilityTelemetryValuesParamsFiltersOperationGreaterOrEquals, ObservabilityTelemetryValuesParamsFiltersOperationLess, ObservabilityTelemetryValuesParamsFiltersOperationLessOrEquals, ObservabilityTelemetryValuesParamsFiltersOperationIncludesUppercase, ObservabilityTelemetryValuesParamsFiltersOperationDoesNotInclude, ObservabilityTelemetryValuesParamsFiltersOperationMatchRegex, ObservabilityTelemetryValuesParamsFiltersOperationExistsUppercase, ObservabilityTelemetryValuesParamsFiltersOperationDoesNotExist, ObservabilityTelemetryValuesParamsFiltersOperationInUppercase, ObservabilityTelemetryValuesParamsFiltersOperationNotInUppercase, ObservabilityTelemetryValuesParamsFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

// Data type of the filter field. Must match the actual type of the key being
// filtered.
type ObservabilityTelemetryValuesParamsFiltersType string

const (
	ObservabilityTelemetryValuesParamsFiltersTypeString  ObservabilityTelemetryValuesParamsFiltersType = "string"
	ObservabilityTelemetryValuesParamsFiltersTypeNumber  ObservabilityTelemetryValuesParamsFiltersType = "number"
	ObservabilityTelemetryValuesParamsFiltersTypeBoolean ObservabilityTelemetryValuesParamsFiltersType = "boolean"
)

func (r ObservabilityTelemetryValuesParamsFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValuesParamsFiltersTypeString, ObservabilityTelemetryValuesParamsFiltersTypeNumber, ObservabilityTelemetryValuesParamsFiltersTypeBoolean:
		return true
	}
	return false
}

// Full-text search expression to match events containing the specified text or
// pattern.
type ObservabilityTelemetryValuesParamsNeedle struct {
	// The text or pattern to search for.
	Value param.Field[ObservabilityTelemetryValuesParamsNeedleValueUnion] `json:"value" api:"required"`
	// When true, treats the value as a regular expression (RE2 syntax).
	IsRegex param.Field[bool] `json:"isRegex"`
	// When true, performs a case-sensitive search. Defaults to case-insensitive.
	MatchCase param.Field[bool] `json:"matchCase"`
}

func (r ObservabilityTelemetryValuesParamsNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// The text or pattern to search for.
//
// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryValuesParamsNeedleValueUnion interface {
	ImplementsObservabilityTelemetryValuesParamsNeedleValueUnion()
}
