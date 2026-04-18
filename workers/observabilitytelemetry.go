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

type ObservabilityTelemetryQueryResponse struct {
	// A Workers Observability Query Object
	Run ObservabilityTelemetryQueryResponseRun `json:"run" api:"required"`
	// The statistics object contains information about query performance from the
	// database, it does not include any network latency
	Statistics   ObservabilityTelemetryQueryResponseStatistics              `json:"statistics" api:"required"`
	Agents       []ObservabilityTelemetryQueryResponseAgent                 `json:"agents"`
	Calculations []ObservabilityTelemetryQueryResponseCalculation           `json:"calculations"`
	Compare      []ObservabilityTelemetryQueryResponseCompare               `json:"compare"`
	Events       ObservabilityTelemetryQueryResponseEvents                  `json:"events"`
	Invocations  map[string][]ObservabilityTelemetryQueryResponseInvocation `json:"invocations"`
	Traces       []ObservabilityTelemetryQueryResponseTrace                 `json:"traces"`
	JSON         observabilityTelemetryQueryResponseJSON                    `json:"-"`
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

// A Workers Observability Query Object
type ObservabilityTelemetryQueryResponseRun struct {
	ID          string                                       `json:"id" api:"required"`
	AccountID   string                                       `json:"accountId" api:"required"`
	Dry         bool                                         `json:"dry" api:"required"`
	Granularity float64                                      `json:"granularity" api:"required"`
	Query       ObservabilityTelemetryQueryResponseRunQuery  `json:"query" api:"required"`
	Status      ObservabilityTelemetryQueryResponseRunStatus `json:"status" api:"required"`
	// Time range for the query execution
	Timeframe  ObservabilityTelemetryQueryResponseRunTimeframe  `json:"timeframe" api:"required"`
	UserID     string                                           `json:"userId" api:"required"`
	Created    string                                           `json:"created"`
	Statistics ObservabilityTelemetryQueryResponseRunStatistics `json:"statistics"`
	Updated    string                                           `json:"updated"`
	JSON       observabilityTelemetryQueryResponseRunJSON       `json:"-"`
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
	// Filter field name. IMPORTANT: do not guess keys. Always use verified keys from
	// previous query results or the observability_keys response. Preferred keys:
	// $metadata.service, $metadata.origin, $metadata.trigger, $metadata.message,
	// $metadata.error.
	Key       string                                                                `json:"key"`
	Kind      ObservabilityTelemetryQueryResponseRunQueryParametersFiltersKind      `json:"kind"`
	Operation ObservabilityTelemetryQueryResponseRunQueryParametersFiltersOperation `json:"operation"`
	Type      ObservabilityTelemetryQueryResponseRunQueryParametersFiltersType      `json:"type"`
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

// Filtering best practices: use observability_keys and observability_values to
// confirm available fields and values. If searching for errors, filter for
// $metadata.error exists.
type ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. IMPORTANT: do not guess keys. Always use verified keys from
	// previous query results or the observability_keys response. Preferred keys:
	// $metadata.service, $metadata.origin, $metadata.trigger, $metadata.message,
	// $metadata.error.
	Key       string                                                                                              `json:"key" api:"required"`
	Operation ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafOperation `json:"operation" api:"required"`
	Type      ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafType      `json:"type" api:"required"`
	Kind      ObservabilityTelemetryQueryResponseRunQueryParametersFiltersWorkersObservabilityFilterLeafKind      `json:"kind"`
	// Filter comparison value. IMPORTANT: must match actual values in your logs.
	// Verify using previous query results or the /values endpoint. Ensure value type
	// matches the field type. String comparisons are case-sensitive unless using
	// specific operations. Regex uses ClickHouse RE2 syntax (no
	// lookaheads/lookbehinds); examples: ^5\d{2}$ for HTTP 5xx, \bERROR\b for word
	// boundary.
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

// Filter comparison value. IMPORTANT: must match actual values in your logs.
// Verify using previous query results or the /values endpoint. Ensure value type
// matches the field type. String comparisons are case-sensitive unless using
// specific operations. Regex uses ClickHouse RE2 syntax (no
// lookaheads/lookbehinds); examples: ^5\d{2}$ for HTTP 5xx, \bERROR\b for word
// boundary.
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

// The statistics object contains information about query performance from the
// database, it does not include any network latency
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
	AgentClass      string                                       `json:"agentClass" api:"required"`
	EventTypeCounts map[string]float64                           `json:"eventTypeCounts" api:"required"`
	FirstEventMs    float64                                      `json:"firstEventMs" api:"required"`
	HasErrors       bool                                         `json:"hasErrors" api:"required"`
	LastEventMs     float64                                      `json:"lastEventMs" api:"required"`
	Namespace       string                                       `json:"namespace" api:"required"`
	Service         string                                       `json:"service" api:"required"`
	TotalEvents     float64                                      `json:"totalEvents" api:"required"`
	JSON            observabilityTelemetryQueryResponseAgentJSON `json:"-"`
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

type ObservabilityTelemetryQueryResponseEvents struct {
	Count  float64                                           `json:"count"`
	Events []ObservabilityTelemetryQueryResponseEventsEvent  `json:"events"`
	Fields []ObservabilityTelemetryQueryResponseEventsField  `json:"fields"`
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

// The data structure of a telemetry event
type ObservabilityTelemetryQueryResponseEventsEvent struct {
	Metadata  ObservabilityTelemetryQueryResponseEventsEventsMetadata `json:"$metadata" api:"required"`
	Dataset   string                                                  `json:"dataset" api:"required"`
	Source    interface{}                                             `json:"source" api:"required"`
	Timestamp int64                                                   `json:"timestamp" api:"required"`
	// Cloudflare Containers event information enriches your logs so you can easily
	// identify and debug issues.
	Containers interface{} `json:"$containers"`
	// Cloudflare Workers event information enriches your logs so you can easily
	// identify and debug issues.
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

type ObservabilityTelemetryQueryResponseEventsEventsMetadata struct {
	// Unique event ID. Use as the cursor for offset-based pagination.
	ID              string                                                      `json:"id" api:"required"`
	Account         string                                                      `json:"account"`
	CloudService    string                                                      `json:"cloudService"`
	ColdStart       int64                                                       `json:"coldStart"`
	Cost            int64                                                       `json:"cost"`
	Duration        int64                                                       `json:"duration"`
	EndTime         int64                                                       `json:"endTime"`
	Error           string                                                      `json:"error"`
	ErrorTemplate   string                                                      `json:"errorTemplate"`
	Fingerprint     string                                                      `json:"fingerprint"`
	Level           string                                                      `json:"level"`
	Message         string                                                      `json:"message"`
	MessageTemplate string                                                      `json:"messageTemplate"`
	MetricName      string                                                      `json:"metricName"`
	Origin          string                                                      `json:"origin"`
	ParentSpanID    string                                                      `json:"parentSpanId"`
	Provider        string                                                      `json:"provider"`
	Region          string                                                      `json:"region"`
	RequestID       string                                                      `json:"requestId"`
	Service         string                                                      `json:"service"`
	SpanID          string                                                      `json:"spanId"`
	SpanName        string                                                      `json:"spanName"`
	StackID         string                                                      `json:"stackId"`
	StartTime       int64                                                       `json:"startTime"`
	StatusCode      int64                                                       `json:"statusCode"`
	TraceDuration   int64                                                       `json:"traceDuration"`
	TraceID         string                                                      `json:"traceId"`
	TransactionName string                                                      `json:"transactionName"`
	Trigger         string                                                      `json:"trigger"`
	Type            string                                                      `json:"type"`
	URL             string                                                      `json:"url"`
	JSON            observabilityTelemetryQueryResponseEventsEventsMetadataJSON `json:"-"`
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

// Cloudflare Workers event information enriches your logs so you can easily
// identify and debug issues.
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

// Cloudflare Workers event information enriches your logs so you can easily
// identify and debug issues.
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
	Key  string                                             `json:"key" api:"required"`
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

// The data structure of a telemetry event
type ObservabilityTelemetryQueryResponseInvocation struct {
	Metadata  ObservabilityTelemetryQueryResponseInvocationsMetadata `json:"$metadata" api:"required"`
	Dataset   string                                                 `json:"dataset" api:"required"`
	Source    interface{}                                            `json:"source" api:"required"`
	Timestamp int64                                                  `json:"timestamp" api:"required"`
	// Cloudflare Containers event information enriches your logs so you can easily
	// identify and debug issues.
	Containers interface{} `json:"$containers"`
	// Cloudflare Workers event information enriches your logs so you can easily
	// identify and debug issues.
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

type ObservabilityTelemetryQueryResponseInvocationsMetadata struct {
	// Unique event ID. Use as the cursor for offset-based pagination.
	ID              string                                                     `json:"id" api:"required"`
	Account         string                                                     `json:"account"`
	CloudService    string                                                     `json:"cloudService"`
	ColdStart       int64                                                      `json:"coldStart"`
	Cost            int64                                                      `json:"cost"`
	Duration        int64                                                      `json:"duration"`
	EndTime         int64                                                      `json:"endTime"`
	Error           string                                                     `json:"error"`
	ErrorTemplate   string                                                     `json:"errorTemplate"`
	Fingerprint     string                                                     `json:"fingerprint"`
	Level           string                                                     `json:"level"`
	Message         string                                                     `json:"message"`
	MessageTemplate string                                                     `json:"messageTemplate"`
	MetricName      string                                                     `json:"metricName"`
	Origin          string                                                     `json:"origin"`
	ParentSpanID    string                                                     `json:"parentSpanId"`
	Provider        string                                                     `json:"provider"`
	Region          string                                                     `json:"region"`
	RequestID       string                                                     `json:"requestId"`
	Service         string                                                     `json:"service"`
	SpanID          string                                                     `json:"spanId"`
	SpanName        string                                                     `json:"spanName"`
	StackID         string                                                     `json:"stackId"`
	StartTime       int64                                                      `json:"startTime"`
	StatusCode      int64                                                      `json:"statusCode"`
	TraceDuration   int64                                                      `json:"traceDuration"`
	TraceID         string                                                     `json:"traceId"`
	TransactionName string                                                     `json:"transactionName"`
	Trigger         string                                                     `json:"trigger"`
	Type            string                                                     `json:"type"`
	URL             string                                                     `json:"url"`
	JSON            observabilityTelemetryQueryResponseInvocationsMetadataJSON `json:"-"`
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

// Cloudflare Workers event information enriches your logs so you can easily
// identify and debug issues.
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

// Cloudflare Workers event information enriches your logs so you can easily
// identify and debug issues.
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
	RootSpanName        string                                       `json:"rootSpanName" api:"required"`
	RootTransactionName string                                       `json:"rootTransactionName" api:"required"`
	Service             []string                                     `json:"service" api:"required"`
	Spans               float64                                      `json:"spans" api:"required"`
	TraceDurationMs     float64                                      `json:"traceDurationMs" api:"required"`
	TraceEndMs          float64                                      `json:"traceEndMs" api:"required"`
	TraceID             string                                       `json:"traceId" api:"required"`
	TraceStartMs        float64                                      `json:"traceStartMs" api:"required"`
	Errors              []string                                     `json:"errors"`
	JSON                observabilityTelemetryQueryResponseTraceJSON `json:"-"`
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

// Supports nested groups via kind: 'group'. Maximum nesting depth is 4.
type ObservabilityTelemetryKeysParamsFilter struct {
	FilterCombination param.Field[ObservabilityTelemetryKeysParamsFiltersFilterCombination] `json:"filterCombination"`
	Filters           param.Field[interface{}]                                              `json:"filters"`
	// Filter field name. IMPORTANT: do not guess keys. Always use verified keys from
	// previous query results or the observability_keys response. Preferred keys:
	// $metadata.service, $metadata.origin, $metadata.trigger, $metadata.message,
	// $metadata.error.
	Key       param.Field[string]                                           `json:"key"`
	Kind      param.Field[ObservabilityTelemetryKeysParamsFiltersKind]      `json:"kind"`
	Operation param.Field[ObservabilityTelemetryKeysParamsFiltersOperation] `json:"operation"`
	Type      param.Field[ObservabilityTelemetryKeysParamsFiltersType]      `json:"type"`
	Value     param.Field[interface{}]                                      `json:"value"`
}

func (r ObservabilityTelemetryKeysParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryKeysParamsFilter) implementsObservabilityTelemetryKeysParamsFilterUnion() {
}

// Supports nested groups via kind: 'group'. Maximum nesting depth is 4.
//
// Satisfied by [workers.ObservabilityTelemetryKeysParamsFiltersObject],
// [workers.ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeaf],
// [ObservabilityTelemetryKeysParamsFilter].
type ObservabilityTelemetryKeysParamsFilterUnion interface {
	implementsObservabilityTelemetryKeysParamsFilterUnion()
}

type ObservabilityTelemetryKeysParamsFiltersObject struct {
	FilterCombination param.Field[ObservabilityTelemetryKeysParamsFiltersObjectFilterCombination] `json:"filterCombination" api:"required"`
	Filters           param.Field[[]interface{}]                                                  `json:"filters" api:"required"`
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

// Filtering best practices: use observability_keys and observability_values to
// confirm available fields and values. If searching for errors, filter for
// $metadata.error exists.
type ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. IMPORTANT: do not guess keys. Always use verified keys from
	// previous query results or the observability_keys response. Preferred keys:
	// $metadata.service, $metadata.origin, $metadata.trigger, $metadata.message,
	// $metadata.error.
	Key       param.Field[string]                                                                         `json:"key" api:"required"`
	Operation param.Field[ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafOperation] `json:"operation" api:"required"`
	Type      param.Field[ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafType]      `json:"type" api:"required"`
	Kind      param.Field[ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafKind]      `json:"kind"`
	// Filter comparison value. IMPORTANT: must match actual values in your logs.
	// Verify using previous query results or the /values endpoint. Ensure value type
	// matches the field type. String comparisons are case-sensitive unless using
	// specific operations. Regex uses ClickHouse RE2 syntax (no
	// lookaheads/lookbehinds); examples: ^5\d{2}$ for HTTP 5xx, \bERROR\b for word
	// boundary.
	Value param.Field[ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeafValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryKeysParamsFiltersWorkersObservabilityFilterLeaf) implementsObservabilityTelemetryKeysParamsFilterUnion() {
}

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

// Filter comparison value. IMPORTANT: must match actual values in your logs.
// Verify using previous query results or the /values endpoint. Ensure value type
// matches the field type. String comparisons are case-sensitive unless using
// specific operations. Regex uses ClickHouse RE2 syntax (no
// lookaheads/lookbehinds); examples: ^5\d{2}$ for HTTP 5xx, \bERROR\b for word
// boundary.
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
	Value     param.Field[ObservabilityTelemetryKeysParamsKeyNeedleValueUnion] `json:"value" api:"required"`
	IsRegex   param.Field[bool]                                                `json:"isRegex"`
	MatchCase param.Field[bool]                                                `json:"matchCase"`
}

func (r ObservabilityTelemetryKeysParamsKeyNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryKeysParamsKeyNeedleValueUnion interface {
	ImplementsObservabilityTelemetryKeysParamsKeyNeedleValueUnion()
}

// Search for a specific substring in any of the events
type ObservabilityTelemetryKeysParamsNeedle struct {
	Value     param.Field[ObservabilityTelemetryKeysParamsNeedleValueUnion] `json:"value" api:"required"`
	IsRegex   param.Field[bool]                                             `json:"isRegex"`
	MatchCase param.Field[bool]                                             `json:"matchCase"`
}

func (r ObservabilityTelemetryKeysParamsNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryKeysParamsNeedleValueUnion interface {
	ImplementsObservabilityTelemetryKeysParamsNeedleValueUnion()
}

type ObservabilityTelemetryQueryParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Unique identifier for the query to execute
	QueryID param.Field[string] `json:"queryId" api:"required"`
	// Timeframe for your query using Unix timestamps in milliseconds. Provide from/to
	// epoch ms; narrower timeframes provide faster responses and more specific
	// results.
	Timeframe param.Field[ObservabilityTelemetryQueryParamsTimeframe] `json:"timeframe" api:"required"`
	// Whether to include timeseties data in the response
	Chart param.Field[bool] `json:"chart"`
	// Whether to include comparison data with previous time periods
	Compare param.Field[bool] `json:"compare"`
	// Whether to perform a dry run without saving the results of the query. Useful for
	// validation
	Dry param.Field[bool] `json:"dry"`
	// This is only used when the view is calculations. Leaving it empty lets Workers
	// Observability detect the correct granularity.
	Granularity param.Field[float64] `json:"granularity"`
	// Whether to ignore time-series data in the results and return only aggregated
	// values
	IgnoreSeries param.Field[bool] `json:"ignoreSeries"`
	// Use this limit to cap the number of events returned when the view is events.
	Limit param.Field[float64] `json:"limit"`
	// Cursor pagination for event/trace/invocation views. Pass the last item's
	// $metadata.id as the next offset.
	Offset param.Field[string] `json:"offset"`
	// Numeric offset for pattern results (top-N list). Use with limit to page pattern
	// groups; not used by cursor pagination.
	OffsetBy param.Field[float64] `json:"offsetBy"`
	// Direction for offset-based pagination (e.g., 'next', 'prev')
	OffsetDirection param.Field[string] `json:"offsetDirection"`
	// Optional parameters to pass to the query execution
	Parameters param.Field[ObservabilityTelemetryQueryParamsParameters] `json:"parameters"`
	// Examples by view type. Events: show errors for a worker in the last 30 minutes.
	// Calculations: p99 of wall time or count by status code. Invocations: find a
	// specific request that resulted in a 500.
	View param.Field[ObservabilityTelemetryQueryParamsView] `json:"view"`
}

func (r ObservabilityTelemetryQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Timeframe for your query using Unix timestamps in milliseconds. Provide from/to
// epoch ms; narrower timeframes provide faster responses and more specific
// results.
type ObservabilityTelemetryQueryParamsTimeframe struct {
	// Start timestamp for the query timeframe (Unix timestamp in milliseconds)
	From param.Field[float64] `json:"from" api:"required"`
	// End timestamp for the query timeframe (Unix timestamp in milliseconds)
	To param.Field[float64] `json:"to" api:"required"`
}

func (r ObservabilityTelemetryQueryParamsTimeframe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Optional parameters to pass to the query execution
type ObservabilityTelemetryQueryParamsParameters struct {
	// Create Calculations to compute as part of the query.
	Calculations param.Field[[]ObservabilityTelemetryQueryParamsParametersCalculation] `json:"calculations"`
	// Set the Datasets to query. Leave it empty to query all the datasets.
	Datasets param.Field[[]string] `json:"datasets"`
	// Set a Flag to describe how to combine the filters on the query.
	FilterCombination param.Field[ObservabilityTelemetryQueryParamsParametersFilterCombination] `json:"filterCombination"`
	// Configure the Filters to apply to the query. Supports nested groups via kind:
	// 'group'. Maximum nesting depth is 4.
	Filters param.Field[[]ObservabilityTelemetryQueryParamsParametersFilterUnion] `json:"filters"`
	// Define how to group the results of the query.
	GroupBys param.Field[[]ObservabilityTelemetryQueryParamsParametersGroupBy] `json:"groupBys"`
	// Configure the Having clauses that filter on calculations in the query result.
	Havings param.Field[[]ObservabilityTelemetryQueryParamsParametersHaving] `json:"havings"`
	// Set a limit on the number of results / records returned by the query
	Limit param.Field[int64] `json:"limit"`
	// Define an expression to search using full-text search.
	Needle param.Field[ObservabilityTelemetryQueryParamsParametersNeedle] `json:"needle"`
	// Configure the order of the results returned by the query.
	OrderBy param.Field[ObservabilityTelemetryQueryParamsParametersOrderBy] `json:"orderBy"`
}

func (r ObservabilityTelemetryQueryParamsParameters) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryQueryParamsParametersCalculation struct {
	Operator param.Field[ObservabilityTelemetryQueryParamsParametersCalculationsOperator] `json:"operator" api:"required"`
	Alias    param.Field[string]                                                          `json:"alias"`
	// The key to use for the calculation. This key must exist in the logs. Use the
	// observability_keys response to confirm. Do not guess keys.
	Key     param.Field[string]                                                         `json:"key"`
	KeyType param.Field[ObservabilityTelemetryQueryParamsParametersCalculationsKeyType] `json:"keyType"`
}

func (r ObservabilityTelemetryQueryParamsParametersCalculation) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

// Set a Flag to describe how to combine the filters on the query.
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

// Supports nested groups via kind: 'group'. Maximum nesting depth is 4.
type ObservabilityTelemetryQueryParamsParametersFilter struct {
	FilterCombination param.Field[ObservabilityTelemetryQueryParamsParametersFiltersFilterCombination] `json:"filterCombination"`
	Filters           param.Field[interface{}]                                                         `json:"filters"`
	// Filter field name. IMPORTANT: do not guess keys. Always use verified keys from
	// previous query results or the observability_keys response. Preferred keys:
	// $metadata.service, $metadata.origin, $metadata.trigger, $metadata.message,
	// $metadata.error.
	Key       param.Field[string]                                                      `json:"key"`
	Kind      param.Field[ObservabilityTelemetryQueryParamsParametersFiltersKind]      `json:"kind"`
	Operation param.Field[ObservabilityTelemetryQueryParamsParametersFiltersOperation] `json:"operation"`
	Type      param.Field[ObservabilityTelemetryQueryParamsParametersFiltersType]      `json:"type"`
	Value     param.Field[interface{}]                                                 `json:"value"`
}

func (r ObservabilityTelemetryQueryParamsParametersFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryQueryParamsParametersFilter) implementsObservabilityTelemetryQueryParamsParametersFilterUnion() {
}

// Supports nested groups via kind: 'group'. Maximum nesting depth is 4.
//
// Satisfied by [workers.ObservabilityTelemetryQueryParamsParametersFiltersObject],
// [workers.ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeaf],
// [ObservabilityTelemetryQueryParamsParametersFilter].
type ObservabilityTelemetryQueryParamsParametersFilterUnion interface {
	implementsObservabilityTelemetryQueryParamsParametersFilterUnion()
}

type ObservabilityTelemetryQueryParamsParametersFiltersObject struct {
	FilterCombination param.Field[ObservabilityTelemetryQueryParamsParametersFiltersObjectFilterCombination] `json:"filterCombination" api:"required"`
	Filters           param.Field[[]interface{}]                                                             `json:"filters" api:"required"`
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

// Filtering best practices: use observability_keys and observability_values to
// confirm available fields and values. If searching for errors, filter for
// $metadata.error exists.
type ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. IMPORTANT: do not guess keys. Always use verified keys from
	// previous query results or the observability_keys response. Preferred keys:
	// $metadata.service, $metadata.origin, $metadata.trigger, $metadata.message,
	// $metadata.error.
	Key       param.Field[string]                                                                                    `json:"key" api:"required"`
	Operation param.Field[ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafOperation] `json:"operation" api:"required"`
	Type      param.Field[ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafType]      `json:"type" api:"required"`
	Kind      param.Field[ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafKind]      `json:"kind"`
	// Filter comparison value. IMPORTANT: must match actual values in your logs.
	// Verify using previous query results or the /values endpoint. Ensure value type
	// matches the field type. String comparisons are case-sensitive unless using
	// specific operations. Regex uses ClickHouse RE2 syntax (no
	// lookaheads/lookbehinds); examples: ^5\d{2}$ for HTTP 5xx, \bERROR\b for word
	// boundary.
	Value param.Field[ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeafValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryQueryParamsParametersFiltersWorkersObservabilityFilterLeaf) implementsObservabilityTelemetryQueryParamsParametersFilterUnion() {
}

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

// Filter comparison value. IMPORTANT: must match actual values in your logs.
// Verify using previous query results or the /values endpoint. Ensure value type
// matches the field type. String comparisons are case-sensitive unless using
// specific operations. Regex uses ClickHouse RE2 syntax (no
// lookaheads/lookbehinds); examples: ^5\d{2}$ for HTTP 5xx, \bERROR\b for word
// boundary.
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
	Type  param.Field[ObservabilityTelemetryQueryParamsParametersGroupBysType] `json:"type" api:"required"`
	Value param.Field[string]                                                  `json:"value" api:"required"`
}

func (r ObservabilityTelemetryQueryParamsParametersGroupBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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
	Key       param.Field[string]                                                      `json:"key" api:"required"`
	Operation param.Field[ObservabilityTelemetryQueryParamsParametersHavingsOperation] `json:"operation" api:"required"`
	Value     param.Field[float64]                                                     `json:"value" api:"required"`
}

func (r ObservabilityTelemetryQueryParamsParametersHaving) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

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

// Define an expression to search using full-text search.
type ObservabilityTelemetryQueryParamsParametersNeedle struct {
	Value     param.Field[ObservabilityTelemetryQueryParamsParametersNeedleValueUnion] `json:"value" api:"required"`
	IsRegex   param.Field[bool]                                                        `json:"isRegex"`
	MatchCase param.Field[bool]                                                        `json:"matchCase"`
}

func (r ObservabilityTelemetryQueryParamsParametersNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryQueryParamsParametersNeedleValueUnion interface {
	ImplementsObservabilityTelemetryQueryParamsParametersNeedleValueUnion()
}

// Configure the order of the results returned by the query.
type ObservabilityTelemetryQueryParamsParametersOrderBy struct {
	// Configure which Calculation to order the results by.
	Value param.Field[string] `json:"value" api:"required"`
	// Set the order of the results
	Order param.Field[ObservabilityTelemetryQueryParamsParametersOrderByOrder] `json:"order"`
}

func (r ObservabilityTelemetryQueryParamsParametersOrderBy) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Set the order of the results
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

// Examples by view type. Events: show errors for a worker in the last 30 minutes.
// Calculations: p99 of wall time or count by status code. Invocations: find a
// specific request that resulted in a 500.
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
	Result   ObservabilityTelemetryQueryResponse                   `json:"result" api:"required"`
	Success  ObservabilityTelemetryQueryResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityTelemetryQueryResponseEnvelopeJSON       `json:"-"`
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
	// Search for a specific substring in the event.
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

// Supports nested groups via kind: 'group'. Maximum nesting depth is 4.
type ObservabilityTelemetryValuesParamsFilter struct {
	FilterCombination param.Field[ObservabilityTelemetryValuesParamsFiltersFilterCombination] `json:"filterCombination"`
	Filters           param.Field[interface{}]                                                `json:"filters"`
	// Filter field name. IMPORTANT: do not guess keys. Always use verified keys from
	// previous query results or the observability_keys response. Preferred keys:
	// $metadata.service, $metadata.origin, $metadata.trigger, $metadata.message,
	// $metadata.error.
	Key       param.Field[string]                                             `json:"key"`
	Kind      param.Field[ObservabilityTelemetryValuesParamsFiltersKind]      `json:"kind"`
	Operation param.Field[ObservabilityTelemetryValuesParamsFiltersOperation] `json:"operation"`
	Type      param.Field[ObservabilityTelemetryValuesParamsFiltersType]      `json:"type"`
	Value     param.Field[interface{}]                                        `json:"value"`
}

func (r ObservabilityTelemetryValuesParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryValuesParamsFilter) implementsObservabilityTelemetryValuesParamsFilterUnion() {
}

// Supports nested groups via kind: 'group'. Maximum nesting depth is 4.
//
// Satisfied by [workers.ObservabilityTelemetryValuesParamsFiltersObject],
// [workers.ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeaf],
// [ObservabilityTelemetryValuesParamsFilter].
type ObservabilityTelemetryValuesParamsFilterUnion interface {
	implementsObservabilityTelemetryValuesParamsFilterUnion()
}

type ObservabilityTelemetryValuesParamsFiltersObject struct {
	FilterCombination param.Field[ObservabilityTelemetryValuesParamsFiltersObjectFilterCombination] `json:"filterCombination" api:"required"`
	Filters           param.Field[[]interface{}]                                                    `json:"filters" api:"required"`
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

// Filtering best practices: use observability_keys and observability_values to
// confirm available fields and values. If searching for errors, filter for
// $metadata.error exists.
type ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeaf struct {
	// Filter field name. IMPORTANT: do not guess keys. Always use verified keys from
	// previous query results or the observability_keys response. Preferred keys:
	// $metadata.service, $metadata.origin, $metadata.trigger, $metadata.message,
	// $metadata.error.
	Key       param.Field[string]                                                                           `json:"key" api:"required"`
	Operation param.Field[ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafOperation] `json:"operation" api:"required"`
	Type      param.Field[ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafType]      `json:"type" api:"required"`
	Kind      param.Field[ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafKind]      `json:"kind"`
	// Filter comparison value. IMPORTANT: must match actual values in your logs.
	// Verify using previous query results or the /values endpoint. Ensure value type
	// matches the field type. String comparisons are case-sensitive unless using
	// specific operations. Regex uses ClickHouse RE2 syntax (no
	// lookaheads/lookbehinds); examples: ^5\d{2}$ for HTTP 5xx, \bERROR\b for word
	// boundary.
	Value param.Field[ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeafValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeaf) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ObservabilityTelemetryValuesParamsFiltersWorkersObservabilityFilterLeaf) implementsObservabilityTelemetryValuesParamsFilterUnion() {
}

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

// Filter comparison value. IMPORTANT: must match actual values in your logs.
// Verify using previous query results or the /values endpoint. Ensure value type
// matches the field type. String comparisons are case-sensitive unless using
// specific operations. Regex uses ClickHouse RE2 syntax (no
// lookaheads/lookbehinds); examples: ^5\d{2}$ for HTTP 5xx, \bERROR\b for word
// boundary.
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

// Search for a specific substring in the event.
type ObservabilityTelemetryValuesParamsNeedle struct {
	Value     param.Field[ObservabilityTelemetryValuesParamsNeedleValueUnion] `json:"value" api:"required"`
	IsRegex   param.Field[bool]                                               `json:"isRegex"`
	MatchCase param.Field[bool]                                               `json:"matchCase"`
}

func (r ObservabilityTelemetryValuesParamsNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryValuesParamsNeedleValueUnion interface {
	ImplementsObservabilityTelemetryValuesParamsNeedleValueUnion()
}
