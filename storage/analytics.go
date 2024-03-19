// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package storage

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AnalyticsService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAnalyticsService] method instead.
type AnalyticsService struct {
	Options []option.RequestOption
}

// NewAnalyticsService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsService(opts ...option.RequestOption) (r *AnalyticsService) {
	r = &AnalyticsService{}
	r.Options = opts
	return
}

// Retrieves Workers KV request metrics for the given account.
func (r *AnalyticsService) List(ctx context.Context, params AnalyticsListParams, opts ...option.RequestOption) (res *WorkersKVSchemasResult, err error) {
	opts = append(r.Options[:], opts...)
	var env AnalyticsListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/analytics", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves Workers KV stored data metrics for the given account.
func (r *AnalyticsService) Stored(ctx context.Context, params AnalyticsStoredParams, opts ...option.RequestOption) (res *WorkersKVComponentsSchemasResult, err error) {
	opts = append(r.Options[:], opts...)
	var env AnalyticsStoredResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/analytics/stored", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Metrics on Workers KV requests.
type WorkersKVComponentsSchemasResult struct {
	Data []WorkersKVComponentsSchemasResultData `json:"data,required,nullable"`
	// Number of seconds between current time and last processed event, i.e. how many
	// seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric.
	Min interface{} `json:"min,required"`
	// For specifying result metrics.
	Query WorkersKVComponentsSchemasResultQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total results for metrics across all data.
	Totals interface{}                          `json:"totals,required"`
	JSON   workersKVComponentsSchemasResultJSON `json:"-"`
}

// workersKVComponentsSchemasResultJSON contains the JSON metadata for the struct
// [WorkersKVComponentsSchemasResult]
type workersKVComponentsSchemasResultJSON struct {
	Data        apijson.Field
	DataLag     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	Query       apijson.Field
	Rows        apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersKVComponentsSchemasResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersKVComponentsSchemasResultJSON) RawJSON() string {
	return r.raw
}

type WorkersKVComponentsSchemasResultData struct {
	// List of metrics returned by the query.
	Metrics []interface{}                            `json:"metrics,required"`
	JSON    workersKVComponentsSchemasResultDataJSON `json:"-"`
}

// workersKVComponentsSchemasResultDataJSON contains the JSON metadata for the
// struct [WorkersKVComponentsSchemasResultData]
type workersKVComponentsSchemasResultDataJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersKVComponentsSchemasResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersKVComponentsSchemasResultDataJSON) RawJSON() string {
	return r.raw
}

// For specifying result metrics.
type WorkersKVComponentsSchemasResultQuery struct {
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
	Sort []interface{} `json:"sort"`
	// End of time interval to query, defaults to current time.
	Until time.Time                                 `json:"until" format:"date-time"`
	JSON  workersKVComponentsSchemasResultQueryJSON `json:"-"`
}

// workersKVComponentsSchemasResultQueryJSON contains the JSON metadata for the
// struct [WorkersKVComponentsSchemasResultQuery]
type workersKVComponentsSchemasResultQueryJSON struct {
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

func (r *WorkersKVComponentsSchemasResultQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersKVComponentsSchemasResultQueryJSON) RawJSON() string {
	return r.raw
}

// Metrics on Workers KV requests.
type WorkersKVSchemasResult struct {
	Data []WorkersKVSchemasResultData `json:"data,required,nullable"`
	// Number of seconds between current time and last processed event, i.e. how many
	// seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric.
	Min interface{} `json:"min,required"`
	// For specifying result metrics.
	Query WorkersKVSchemasResultQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total results for metrics across all data.
	Totals interface{}                `json:"totals,required"`
	JSON   workersKVSchemasResultJSON `json:"-"`
}

// workersKVSchemasResultJSON contains the JSON metadata for the struct
// [WorkersKVSchemasResult]
type workersKVSchemasResultJSON struct {
	Data        apijson.Field
	DataLag     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	Query       apijson.Field
	Rows        apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersKVSchemasResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersKVSchemasResultJSON) RawJSON() string {
	return r.raw
}

type WorkersKVSchemasResultData struct {
	// List of metrics returned by the query.
	Metrics []interface{}                  `json:"metrics,required"`
	JSON    workersKVSchemasResultDataJSON `json:"-"`
}

// workersKVSchemasResultDataJSON contains the JSON metadata for the struct
// [WorkersKVSchemasResultData]
type workersKVSchemasResultDataJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersKVSchemasResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersKVSchemasResultDataJSON) RawJSON() string {
	return r.raw
}

// For specifying result metrics.
type WorkersKVSchemasResultQuery struct {
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
	Sort []interface{} `json:"sort"`
	// End of time interval to query, defaults to current time.
	Until time.Time                       `json:"until" format:"date-time"`
	JSON  workersKVSchemasResultQueryJSON `json:"-"`
}

// workersKVSchemasResultQueryJSON contains the JSON metadata for the struct
// [WorkersKVSchemasResultQuery]
type workersKVSchemasResultQueryJSON struct {
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

func (r *WorkersKVSchemasResultQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersKVSchemasResultQueryJSON) RawJSON() string {
	return r.raw
}

type AnalyticsListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// For specifying result metrics.
	Query param.Field[AnalyticsListParamsQuery] `query:"query"`
}

// URLQuery serializes [AnalyticsListParams]'s query parameters as `url.Values`.
func (r AnalyticsListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// For specifying result metrics.
type AnalyticsListParamsQuery struct {
	// Can be used to break down the data by given attributes.
	Dimensions param.Field[[]AnalyticsListParamsQueryDimension] `query:"dimensions"`
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
	Metrics param.Field[[]AnalyticsListParamsQueryMetric] `query:"metrics"`
	// Start of time interval to query, defaults to 6 hours before request received.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// Array of dimensions or metrics to sort by, each dimension/metric may be prefixed
	// by - (descending) or + (ascending).
	Sort param.Field[[]interface{}] `query:"sort"`
	// End of time interval to query, defaults to current time.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AnalyticsListParamsQuery]'s query parameters as
// `url.Values`.
func (r AnalyticsListParamsQuery) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// For drilling down on metrics.
type AnalyticsListParamsQueryDimension string

const (
	AnalyticsListParamsQueryDimensionAccountID    AnalyticsListParamsQueryDimension = "accountId"
	AnalyticsListParamsQueryDimensionResponseCode AnalyticsListParamsQueryDimension = "responseCode"
	AnalyticsListParamsQueryDimensionRequestType  AnalyticsListParamsQueryDimension = "requestType"
)

// A quantitative measurement of KV usage.
type AnalyticsListParamsQueryMetric string

const (
	AnalyticsListParamsQueryMetricRequests AnalyticsListParamsQueryMetric = "requests"
	AnalyticsListParamsQueryMetricWriteKiB AnalyticsListParamsQueryMetric = "writeKiB"
	AnalyticsListParamsQueryMetricReadKiB  AnalyticsListParamsQueryMetric = "readKiB"
)

type AnalyticsListResponseEnvelope struct {
	Errors   []AnalyticsListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsListResponseEnvelopeMessages `json:"messages,required"`
	// Metrics on Workers KV requests.
	Result WorkersKVSchemasResult `json:"result,required"`
	// Whether the API call was successful
	Success AnalyticsListResponseEnvelopeSuccess `json:"success,required"`
	JSON    analyticsListResponseEnvelopeJSON    `json:"-"`
}

// analyticsListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AnalyticsListResponseEnvelope]
type analyticsListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsListResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    analyticsListResponseEnvelopeErrorsJSON `json:"-"`
}

// analyticsListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AnalyticsListResponseEnvelopeErrors]
type analyticsListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsListResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    analyticsListResponseEnvelopeMessagesJSON `json:"-"`
}

// analyticsListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AnalyticsListResponseEnvelopeMessages]
type analyticsListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AnalyticsListResponseEnvelopeSuccess bool

const (
	AnalyticsListResponseEnvelopeSuccessTrue AnalyticsListResponseEnvelopeSuccess = true
)

type AnalyticsStoredParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// For specifying result metrics.
	Query param.Field[AnalyticsStoredParamsQuery] `query:"query"`
}

// URLQuery serializes [AnalyticsStoredParams]'s query parameters as `url.Values`.
func (r AnalyticsStoredParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// For specifying result metrics.
type AnalyticsStoredParamsQuery struct {
	// Can be used to break down the data by given attributes.
	Dimensions param.Field[[]AnalyticsStoredParamsQueryDimension] `query:"dimensions"`
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
	Metrics param.Field[[]AnalyticsStoredParamsQueryMetric] `query:"metrics"`
	// Start of time interval to query, defaults to 6 hours before request received.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// Array of dimensions or metrics to sort by, each dimension/metric may be prefixed
	// by - (descending) or + (ascending).
	Sort param.Field[[]interface{}] `query:"sort"`
	// End of time interval to query, defaults to current time.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AnalyticsStoredParamsQuery]'s query parameters as
// `url.Values`.
func (r AnalyticsStoredParamsQuery) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// For drilling down on metrics.
type AnalyticsStoredParamsQueryDimension string

const (
	AnalyticsStoredParamsQueryDimensionNamespaceID AnalyticsStoredParamsQueryDimension = "namespaceId"
)

// A quantitative measurement of KV usage.
type AnalyticsStoredParamsQueryMetric string

const (
	AnalyticsStoredParamsQueryMetricStoredBytes AnalyticsStoredParamsQueryMetric = "storedBytes"
	AnalyticsStoredParamsQueryMetricStoredKeys  AnalyticsStoredParamsQueryMetric = "storedKeys"
)

type AnalyticsStoredResponseEnvelope struct {
	Errors   []AnalyticsStoredResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsStoredResponseEnvelopeMessages `json:"messages,required"`
	// Metrics on Workers KV requests.
	Result WorkersKVComponentsSchemasResult `json:"result,required"`
	// Whether the API call was successful
	Success AnalyticsStoredResponseEnvelopeSuccess `json:"success,required"`
	JSON    analyticsStoredResponseEnvelopeJSON    `json:"-"`
}

// analyticsStoredResponseEnvelopeJSON contains the JSON metadata for the struct
// [AnalyticsStoredResponseEnvelope]
type analyticsStoredResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsStoredResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsStoredResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsStoredResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    analyticsStoredResponseEnvelopeErrorsJSON `json:"-"`
}

// analyticsStoredResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AnalyticsStoredResponseEnvelopeErrors]
type analyticsStoredResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsStoredResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsStoredResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsStoredResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    analyticsStoredResponseEnvelopeMessagesJSON `json:"-"`
}

// analyticsStoredResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AnalyticsStoredResponseEnvelopeMessages]
type analyticsStoredResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsStoredResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsStoredResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AnalyticsStoredResponseEnvelopeSuccess bool

const (
	AnalyticsStoredResponseEnvelopeSuccessTrue AnalyticsStoredResponseEnvelopeSuccess = true
)
