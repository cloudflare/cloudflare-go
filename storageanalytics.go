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

// StorageAnalyticsService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewStorageAnalyticsService] method
// instead.
type StorageAnalyticsService struct {
	Options []option.RequestOption
}

// NewStorageAnalyticsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewStorageAnalyticsService(opts ...option.RequestOption) (r *StorageAnalyticsService) {
	r = &StorageAnalyticsService{}
	r.Options = opts
	return
}

// Retrieves Workers KV request metrics for the given account.
func (r *StorageAnalyticsService) List(ctx context.Context, params StorageAnalyticsListParams, opts ...option.RequestOption) (res *WorkersKVSchemasResult, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageAnalyticsListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/storage/analytics", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves Workers KV stored data metrics for the given account.
func (r *StorageAnalyticsService) Stored(ctx context.Context, params StorageAnalyticsStoredParams, opts ...option.RequestOption) (res *WorkersKVComponentsSchemasResult, err error) {
	opts = append(r.Options[:], opts...)
	var env StorageAnalyticsStoredResponseEnvelope
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

type StorageAnalyticsListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// For specifying result metrics.
	Query param.Field[StorageAnalyticsListParamsQuery] `query:"query"`
}

// URLQuery serializes [StorageAnalyticsListParams]'s query parameters as
// `url.Values`.
func (r StorageAnalyticsListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// For specifying result metrics.
type StorageAnalyticsListParamsQuery struct {
	// Can be used to break down the data by given attributes.
	Dimensions param.Field[[]StorageAnalyticsListParamsQueryDimension] `query:"dimensions"`
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
	Metrics param.Field[[]StorageAnalyticsListParamsQueryMetric] `query:"metrics"`
	// Start of time interval to query, defaults to 6 hours before request received.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// Array of dimensions or metrics to sort by, each dimension/metric may be prefixed
	// by - (descending) or + (ascending).
	Sort param.Field[[]interface{}] `query:"sort"`
	// End of time interval to query, defaults to current time.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [StorageAnalyticsListParamsQuery]'s query parameters as
// `url.Values`.
func (r StorageAnalyticsListParamsQuery) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// For drilling down on metrics.
type StorageAnalyticsListParamsQueryDimension string

const (
	StorageAnalyticsListParamsQueryDimensionAccountID    StorageAnalyticsListParamsQueryDimension = "accountId"
	StorageAnalyticsListParamsQueryDimensionResponseCode StorageAnalyticsListParamsQueryDimension = "responseCode"
	StorageAnalyticsListParamsQueryDimensionRequestType  StorageAnalyticsListParamsQueryDimension = "requestType"
)

// A quantitative measurement of KV usage.
type StorageAnalyticsListParamsQueryMetric string

const (
	StorageAnalyticsListParamsQueryMetricRequests StorageAnalyticsListParamsQueryMetric = "requests"
	StorageAnalyticsListParamsQueryMetricWriteKiB StorageAnalyticsListParamsQueryMetric = "writeKiB"
	StorageAnalyticsListParamsQueryMetricReadKiB  StorageAnalyticsListParamsQueryMetric = "readKiB"
)

type StorageAnalyticsListResponseEnvelope struct {
	Errors   []StorageAnalyticsListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageAnalyticsListResponseEnvelopeMessages `json:"messages,required"`
	// Metrics on Workers KV requests.
	Result WorkersKVSchemasResult `json:"result,required"`
	// Whether the API call was successful
	Success StorageAnalyticsListResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageAnalyticsListResponseEnvelopeJSON    `json:"-"`
}

// storageAnalyticsListResponseEnvelopeJSON contains the JSON metadata for the
// struct [StorageAnalyticsListResponseEnvelope]
type storageAnalyticsListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageAnalyticsListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageAnalyticsListResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    storageAnalyticsListResponseEnvelopeErrorsJSON `json:"-"`
}

// storageAnalyticsListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StorageAnalyticsListResponseEnvelopeErrors]
type storageAnalyticsListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageAnalyticsListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageAnalyticsListResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    storageAnalyticsListResponseEnvelopeMessagesJSON `json:"-"`
}

// storageAnalyticsListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [StorageAnalyticsListResponseEnvelopeMessages]
type storageAnalyticsListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageAnalyticsListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageAnalyticsListResponseEnvelopeSuccess bool

const (
	StorageAnalyticsListResponseEnvelopeSuccessTrue StorageAnalyticsListResponseEnvelopeSuccess = true
)

type StorageAnalyticsStoredParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// For specifying result metrics.
	Query param.Field[StorageAnalyticsStoredParamsQuery] `query:"query"`
}

// URLQuery serializes [StorageAnalyticsStoredParams]'s query parameters as
// `url.Values`.
func (r StorageAnalyticsStoredParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// For specifying result metrics.
type StorageAnalyticsStoredParamsQuery struct {
	// Can be used to break down the data by given attributes.
	Dimensions param.Field[[]StorageAnalyticsStoredParamsQueryDimension] `query:"dimensions"`
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
	Metrics param.Field[[]StorageAnalyticsStoredParamsQueryMetric] `query:"metrics"`
	// Start of time interval to query, defaults to 6 hours before request received.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// Array of dimensions or metrics to sort by, each dimension/metric may be prefixed
	// by - (descending) or + (ascending).
	Sort param.Field[[]interface{}] `query:"sort"`
	// End of time interval to query, defaults to current time.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [StorageAnalyticsStoredParamsQuery]'s query parameters as
// `url.Values`.
func (r StorageAnalyticsStoredParamsQuery) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// For drilling down on metrics.
type StorageAnalyticsStoredParamsQueryDimension string

const (
	StorageAnalyticsStoredParamsQueryDimensionNamespaceID StorageAnalyticsStoredParamsQueryDimension = "namespaceId"
)

// A quantitative measurement of KV usage.
type StorageAnalyticsStoredParamsQueryMetric string

const (
	StorageAnalyticsStoredParamsQueryMetricStoredBytes StorageAnalyticsStoredParamsQueryMetric = "storedBytes"
	StorageAnalyticsStoredParamsQueryMetricStoredKeys  StorageAnalyticsStoredParamsQueryMetric = "storedKeys"
)

type StorageAnalyticsStoredResponseEnvelope struct {
	Errors   []StorageAnalyticsStoredResponseEnvelopeErrors   `json:"errors,required"`
	Messages []StorageAnalyticsStoredResponseEnvelopeMessages `json:"messages,required"`
	// Metrics on Workers KV requests.
	Result WorkersKVComponentsSchemasResult `json:"result,required"`
	// Whether the API call was successful
	Success StorageAnalyticsStoredResponseEnvelopeSuccess `json:"success,required"`
	JSON    storageAnalyticsStoredResponseEnvelopeJSON    `json:"-"`
}

// storageAnalyticsStoredResponseEnvelopeJSON contains the JSON metadata for the
// struct [StorageAnalyticsStoredResponseEnvelope]
type storageAnalyticsStoredResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageAnalyticsStoredResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageAnalyticsStoredResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    storageAnalyticsStoredResponseEnvelopeErrorsJSON `json:"-"`
}

// storageAnalyticsStoredResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [StorageAnalyticsStoredResponseEnvelopeErrors]
type storageAnalyticsStoredResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageAnalyticsStoredResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type StorageAnalyticsStoredResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    storageAnalyticsStoredResponseEnvelopeMessagesJSON `json:"-"`
}

// storageAnalyticsStoredResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [StorageAnalyticsStoredResponseEnvelopeMessages]
type storageAnalyticsStoredResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *StorageAnalyticsStoredResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type StorageAnalyticsStoredResponseEnvelopeSuccess bool

const (
	StorageAnalyticsStoredResponseEnvelopeSuccessTrue StorageAnalyticsStoredResponseEnvelopeSuccess = true
)
