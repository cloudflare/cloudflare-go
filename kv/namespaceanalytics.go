// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package kv

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// NamespaceAnalyticsService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNamespaceAnalyticsService] method instead.
type NamespaceAnalyticsService struct {
	Options []option.RequestOption
}

// NewNamespaceAnalyticsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewNamespaceAnalyticsService(opts ...option.RequestOption) (r *NamespaceAnalyticsService) {
	r = &NamespaceAnalyticsService{}
	r.Options = opts
	return
}

// Retrieves Workers KV request metrics for the given account.
func (r *NamespaceAnalyticsService) List(ctx context.Context, params NamespaceAnalyticsListParams, opts ...option.RequestOption) (res *Schema, err error) {
	var env NamespaceAnalyticsListResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/analytics", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves Workers KV stored data metrics for the given account.
func (r *NamespaceAnalyticsService) Stored(ctx context.Context, params NamespaceAnalyticsStoredParams, opts ...option.RequestOption) (res *Components, err error) {
	var env NamespaceAnalyticsStoredResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/storage/analytics/stored", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Metrics on Workers KV requests.
type Components struct {
	Data []ComponentsData `json:"data,required,nullable"`
	// Number of seconds between current time and last processed event, i.e. how many
	// seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric.
	Max map[string]float64 `json:"max,required"`
	// Minimum results for each metric.
	Min map[string]float64 `json:"min,required"`
	// For specifying result metrics.
	Query ComponentsQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total results for metrics across all data.
	Totals map[string]float64 `json:"totals,required"`
	// Time interval buckets by beginning and ending
	TimeIntervals [][]time.Time  `json:"time_intervals" format:"date-time"`
	JSON          componentsJSON `json:"-"`
}

// componentsJSON contains the JSON metadata for the struct [Components]
type componentsJSON struct {
	Data          apijson.Field
	DataLag       apijson.Field
	Max           apijson.Field
	Min           apijson.Field
	Query         apijson.Field
	Rows          apijson.Field
	Totals        apijson.Field
	TimeIntervals apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Components) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r componentsJSON) RawJSON() string {
	return r.raw
}

type ComponentsData struct {
	// List of metrics returned by the query.
	Metrics    [][]float64        `json:"metrics,required"`
	Dimensions []string           `json:"dimensions"`
	JSON       componentsDataJSON `json:"-"`
}

// componentsDataJSON contains the JSON metadata for the struct [ComponentsData]
type componentsDataJSON struct {
	Metrics     apijson.Field
	Dimensions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComponentsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r componentsDataJSON) RawJSON() string {
	return r.raw
}

// For specifying result metrics.
type ComponentsQuery struct {
	// Can be used to break down the data by given attributes.
	Dimensions []string `json:"dimensions"`
	// Used to filter rows by one or more dimensions. Filters can be combined using OR
	// and AND boolean logic. AND takes precedence over OR in all the expressions. The
	// OR operator is defined using a comma (,) or OR keyword surrounded by whitespace.
	// The AND operator is defined using a semicolon (;) or AND keyword surrounded by
	// whitespace. Note that the semicolon is a reserved character in URLs (rfc1738)
	// and needs to be percent-encoded as %3B. Comparison options are:
	//
	// | Operator | Name                     | URL Encoded |
	// | -------- | ------------------------ | ----------- |
	// | ==       | Equals                   | %3D%3D      |
	// | !=       | Does not equals          | !%3D        |
	// | >        | Greater Than             | %3E         |
	// | <        | Less Than                | %3C         |
	// | >=       | Greater than or equal to | %3E%3D      |
	// | <=       | Less than or equal to    | %3C%3D .    |
	Filters string `json:"filters"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit"`
	// One or more metrics to compute.
	Metrics []string `json:"metrics"`
	// Start of time interval to query, defaults to 6 hours before request received.
	Since time.Time `json:"since" format:"date-time"`
	// Array of dimensions or metrics to sort by, each dimension/metric may be prefixed
	// by - (descending) or + (ascending).
	Sort []string `json:"sort"`
	// End of time interval to query, defaults to current time.
	Until time.Time           `json:"until" format:"date-time"`
	JSON  componentsQueryJSON `json:"-"`
}

// componentsQueryJSON contains the JSON metadata for the struct [ComponentsQuery]
type componentsQueryJSON struct {
	Dimensions  apijson.Field
	Filters     apijson.Field
	Limit       apijson.Field
	Metrics     apijson.Field
	Since       apijson.Field
	Sort        apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ComponentsQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r componentsQueryJSON) RawJSON() string {
	return r.raw
}

// Metrics on Workers KV requests.
type Schema struct {
	Data []SchemaData `json:"data,required,nullable"`
	// Number of seconds between current time and last processed event, i.e. how many
	// seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric.
	Max map[string]float64 `json:"max,required"`
	// Minimum results for each metric.
	Min map[string]float64 `json:"min,required"`
	// For specifying result metrics.
	Query SchemaQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total results for metrics across all data.
	Totals map[string]float64 `json:"totals,required"`
	// Time interval buckets by beginning and ending
	TimeIntervals [][]time.Time `json:"time_intervals" format:"date-time"`
	JSON          schemaJSON    `json:"-"`
}

// schemaJSON contains the JSON metadata for the struct [Schema]
type schemaJSON struct {
	Data          apijson.Field
	DataLag       apijson.Field
	Max           apijson.Field
	Min           apijson.Field
	Query         apijson.Field
	Rows          apijson.Field
	Totals        apijson.Field
	TimeIntervals apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Schema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaJSON) RawJSON() string {
	return r.raw
}

type SchemaData struct {
	// List of metrics returned by the query.
	Metrics    [][]float64    `json:"metrics,required"`
	Dimensions []string       `json:"dimensions"`
	JSON       schemaDataJSON `json:"-"`
}

// schemaDataJSON contains the JSON metadata for the struct [SchemaData]
type schemaDataJSON struct {
	Metrics     apijson.Field
	Dimensions  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaDataJSON) RawJSON() string {
	return r.raw
}

// For specifying result metrics.
type SchemaQuery struct {
	// Can be used to break down the data by given attributes.
	Dimensions []string `json:"dimensions"`
	// Used to filter rows by one or more dimensions. Filters can be combined using OR
	// and AND boolean logic. AND takes precedence over OR in all the expressions. The
	// OR operator is defined using a comma (,) or OR keyword surrounded by whitespace.
	// The AND operator is defined using a semicolon (;) or AND keyword surrounded by
	// whitespace. Note that the semicolon is a reserved character in URLs (rfc1738)
	// and needs to be percent-encoded as %3B. Comparison options are:
	//
	// | Operator | Name                     | URL Encoded |
	// | -------- | ------------------------ | ----------- |
	// | ==       | Equals                   | %3D%3D      |
	// | !=       | Does not equals          | !%3D        |
	// | >        | Greater Than             | %3E         |
	// | <        | Less Than                | %3C         |
	// | >=       | Greater than or equal to | %3E%3D      |
	// | <=       | Less than or equal to    | %3C%3D .    |
	Filters string `json:"filters"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit"`
	// One or more metrics to compute.
	Metrics []string `json:"metrics"`
	// Start of time interval to query, defaults to 6 hours before request received.
	Since time.Time `json:"since" format:"date-time"`
	// Array of dimensions or metrics to sort by, each dimension/metric may be prefixed
	// by - (descending) or + (ascending).
	Sort []string `json:"sort"`
	// End of time interval to query, defaults to current time.
	Until time.Time       `json:"until" format:"date-time"`
	JSON  schemaQueryJSON `json:"-"`
}

// schemaQueryJSON contains the JSON metadata for the struct [SchemaQuery]
type schemaQueryJSON struct {
	Dimensions  apijson.Field
	Filters     apijson.Field
	Limit       apijson.Field
	Metrics     apijson.Field
	Since       apijson.Field
	Sort        apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SchemaQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r schemaQueryJSON) RawJSON() string {
	return r.raw
}

type NamespaceAnalyticsListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// For specifying result metrics.
	Query param.Field[NamespaceAnalyticsListParamsQuery] `query:"query"`
}

// URLQuery serializes [NamespaceAnalyticsListParams]'s query parameters as
// `url.Values`.
func (r NamespaceAnalyticsListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// For specifying result metrics.
type NamespaceAnalyticsListParamsQuery struct {
	// Can be used to break down the data by given attributes.
	Dimensions param.Field[[]NamespaceAnalyticsListParamsQueryDimension] `query:"dimensions"`
	// Used to filter rows by one or more dimensions. Filters can be combined using OR
	// and AND boolean logic. AND takes precedence over OR in all the expressions. The
	// OR operator is defined using a comma (,) or OR keyword surrounded by whitespace.
	// The AND operator is defined using a semicolon (;) or AND keyword surrounded by
	// whitespace. Note that the semicolon is a reserved character in URLs (rfc1738)
	// and needs to be percent-encoded as %3B. Comparison options are:
	//
	// | Operator | Name                     | URL Encoded |
	// | -------- | ------------------------ | ----------- |
	// | ==       | Equals                   | %3D%3D      |
	// | !=       | Does not equals          | !%3D        |
	// | >        | Greater Than             | %3E         |
	// | <        | Less Than                | %3C         |
	// | >=       | Greater than or equal to | %3E%3D      |
	// | <=       | Less than or equal to    | %3C%3D .    |
	Filters param.Field[string] `query:"filters"`
	// Limit number of returned metrics.
	Limit param.Field[int64] `query:"limit"`
	// One or more metrics to compute.
	Metrics param.Field[[]NamespaceAnalyticsListParamsQueryMetric] `query:"metrics"`
	// Start of time interval to query, defaults to 6 hours before request received.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// Array of dimensions or metrics to sort by, each dimension/metric may be prefixed
	// by - (descending) or + (ascending).
	Sort param.Field[[]string] `query:"sort"`
	// End of time interval to query, defaults to current time.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [NamespaceAnalyticsListParamsQuery]'s query parameters as
// `url.Values`.
func (r NamespaceAnalyticsListParamsQuery) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// For drilling down on metrics.
type NamespaceAnalyticsListParamsQueryDimension string

const (
	NamespaceAnalyticsListParamsQueryDimensionAccountID    NamespaceAnalyticsListParamsQueryDimension = "accountId"
	NamespaceAnalyticsListParamsQueryDimensionResponseCode NamespaceAnalyticsListParamsQueryDimension = "responseCode"
	NamespaceAnalyticsListParamsQueryDimensionRequestType  NamespaceAnalyticsListParamsQueryDimension = "requestType"
)

func (r NamespaceAnalyticsListParamsQueryDimension) IsKnown() bool {
	switch r {
	case NamespaceAnalyticsListParamsQueryDimensionAccountID, NamespaceAnalyticsListParamsQueryDimensionResponseCode, NamespaceAnalyticsListParamsQueryDimensionRequestType:
		return true
	}
	return false
}

// A quantitative measurement of KV usage.
type NamespaceAnalyticsListParamsQueryMetric string

const (
	NamespaceAnalyticsListParamsQueryMetricRequests NamespaceAnalyticsListParamsQueryMetric = "requests"
	NamespaceAnalyticsListParamsQueryMetricWriteKiB NamespaceAnalyticsListParamsQueryMetric = "writeKiB"
	NamespaceAnalyticsListParamsQueryMetricReadKiB  NamespaceAnalyticsListParamsQueryMetric = "readKiB"
)

func (r NamespaceAnalyticsListParamsQueryMetric) IsKnown() bool {
	switch r {
	case NamespaceAnalyticsListParamsQueryMetricRequests, NamespaceAnalyticsListParamsQueryMetricWriteKiB, NamespaceAnalyticsListParamsQueryMetricReadKiB:
		return true
	}
	return false
}

type NamespaceAnalyticsListResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success NamespaceAnalyticsListResponseEnvelopeSuccess `json:"success,required"`
	// Metrics on Workers KV requests.
	Result Schema                                     `json:"result"`
	JSON   namespaceAnalyticsListResponseEnvelopeJSON `json:"-"`
}

// namespaceAnalyticsListResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceAnalyticsListResponseEnvelope]
type namespaceAnalyticsListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceAnalyticsListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceAnalyticsListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceAnalyticsListResponseEnvelopeSuccess bool

const (
	NamespaceAnalyticsListResponseEnvelopeSuccessTrue NamespaceAnalyticsListResponseEnvelopeSuccess = true
)

func (r NamespaceAnalyticsListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceAnalyticsListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type NamespaceAnalyticsStoredParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// For specifying result metrics.
	Query param.Field[NamespaceAnalyticsStoredParamsQuery] `query:"query"`
}

// URLQuery serializes [NamespaceAnalyticsStoredParams]'s query parameters as
// `url.Values`.
func (r NamespaceAnalyticsStoredParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// For specifying result metrics.
type NamespaceAnalyticsStoredParamsQuery struct {
	// Can be used to break down the data by given attributes.
	Dimensions param.Field[[]NamespaceAnalyticsStoredParamsQueryDimension] `query:"dimensions"`
	// Used to filter rows by one or more dimensions. Filters can be combined using OR
	// and AND boolean logic. AND takes precedence over OR in all the expressions. The
	// OR operator is defined using a comma (,) or OR keyword surrounded by whitespace.
	// The AND operator is defined using a semicolon (;) or AND keyword surrounded by
	// whitespace. Note that the semicolon is a reserved character in URLs (rfc1738)
	// and needs to be percent-encoded as %3B. Comparison options are:
	//
	// | Operator | Name                     | URL Encoded |
	// | -------- | ------------------------ | ----------- |
	// | ==       | Equals                   | %3D%3D      |
	// | !=       | Does not equals          | !%3D        |
	// | >        | Greater Than             | %3E         |
	// | <        | Less Than                | %3C         |
	// | >=       | Greater than or equal to | %3E%3D      |
	// | <=       | Less than or equal to    | %3C%3D .    |
	Filters param.Field[string] `query:"filters"`
	// Limit number of returned metrics.
	Limit param.Field[int64] `query:"limit"`
	// One or more metrics to compute.
	Metrics param.Field[[]NamespaceAnalyticsStoredParamsQueryMetric] `query:"metrics"`
	// Start of time interval to query, defaults to 6 hours before request received.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// Array of dimensions or metrics to sort by, each dimension/metric may be prefixed
	// by - (descending) or + (ascending).
	Sort param.Field[[]string] `query:"sort"`
	// End of time interval to query, defaults to current time.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [NamespaceAnalyticsStoredParamsQuery]'s query parameters as
// `url.Values`.
func (r NamespaceAnalyticsStoredParamsQuery) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// For drilling down on metrics.
type NamespaceAnalyticsStoredParamsQueryDimension string

const (
	NamespaceAnalyticsStoredParamsQueryDimensionNamespaceID NamespaceAnalyticsStoredParamsQueryDimension = "namespaceId"
)

func (r NamespaceAnalyticsStoredParamsQueryDimension) IsKnown() bool {
	switch r {
	case NamespaceAnalyticsStoredParamsQueryDimensionNamespaceID:
		return true
	}
	return false
}

// A quantitative measurement of KV usage.
type NamespaceAnalyticsStoredParamsQueryMetric string

const (
	NamespaceAnalyticsStoredParamsQueryMetricStoredBytes NamespaceAnalyticsStoredParamsQueryMetric = "storedBytes"
	NamespaceAnalyticsStoredParamsQueryMetricStoredKeys  NamespaceAnalyticsStoredParamsQueryMetric = "storedKeys"
)

func (r NamespaceAnalyticsStoredParamsQueryMetric) IsKnown() bool {
	switch r {
	case NamespaceAnalyticsStoredParamsQueryMetricStoredBytes, NamespaceAnalyticsStoredParamsQueryMetricStoredKeys:
		return true
	}
	return false
}

type NamespaceAnalyticsStoredResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success NamespaceAnalyticsStoredResponseEnvelopeSuccess `json:"success,required"`
	// Metrics on Workers KV requests.
	Result Components                                   `json:"result"`
	JSON   namespaceAnalyticsStoredResponseEnvelopeJSON `json:"-"`
}

// namespaceAnalyticsStoredResponseEnvelopeJSON contains the JSON metadata for the
// struct [NamespaceAnalyticsStoredResponseEnvelope]
type namespaceAnalyticsStoredResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NamespaceAnalyticsStoredResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r namespaceAnalyticsStoredResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type NamespaceAnalyticsStoredResponseEnvelopeSuccess bool

const (
	NamespaceAnalyticsStoredResponseEnvelopeSuccessTrue NamespaceAnalyticsStoredResponseEnvelopeSuccess = true
)

func (r NamespaceAnalyticsStoredResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case NamespaceAnalyticsStoredResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
