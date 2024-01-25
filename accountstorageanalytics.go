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

// AccountStorageAnalyticsService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountStorageAnalyticsService] method instead.
type AccountStorageAnalyticsService struct {
	Options []option.RequestOption
}

// NewAccountStorageAnalyticsService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStorageAnalyticsService(opts ...option.RequestOption) (r *AccountStorageAnalyticsService) {
	r = &AccountStorageAnalyticsService{}
	r.Options = opts
	return
}

// Retrieves Workers KV request metrics for the given account.
func (r *AccountStorageAnalyticsService) List(ctx context.Context, accountIdentifier string, query AccountStorageAnalyticsListParams, opts ...option.RequestOption) (res *AccountStorageAnalyticsListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/analytics", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieves Workers KV stored data metrics for the given account.
func (r *AccountStorageAnalyticsService) Stored(ctx context.Context, accountIdentifier string, query AccountStorageAnalyticsStoredParams, opts ...option.RequestOption) (res *AccountStorageAnalyticsStoredResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/storage/analytics/stored", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountStorageAnalyticsListResponse struct {
	Errors   []AccountStorageAnalyticsListResponseError   `json:"errors"`
	Messages []AccountStorageAnalyticsListResponseMessage `json:"messages"`
	Result   AccountStorageAnalyticsListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStorageAnalyticsListResponseSuccess `json:"success"`
	JSON    accountStorageAnalyticsListResponseJSON    `json:"-"`
}

// accountStorageAnalyticsListResponseJSON contains the JSON metadata for the
// struct [AccountStorageAnalyticsListResponse]
type accountStorageAnalyticsListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageAnalyticsListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageAnalyticsListResponseError struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    accountStorageAnalyticsListResponseErrorJSON `json:"-"`
}

// accountStorageAnalyticsListResponseErrorJSON contains the JSON metadata for the
// struct [AccountStorageAnalyticsListResponseError]
type accountStorageAnalyticsListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageAnalyticsListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageAnalyticsListResponseMessage struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountStorageAnalyticsListResponseMessageJSON `json:"-"`
}

// accountStorageAnalyticsListResponseMessageJSON contains the JSON metadata for
// the struct [AccountStorageAnalyticsListResponseMessage]
type accountStorageAnalyticsListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageAnalyticsListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageAnalyticsListResponseResult struct {
	Data interface{} `json:"data"`
	// Number of seconds between current time and last processed event, i.e. how many
	// seconds of data could be missing.
	DataLag float64     `json:"data_lag"`
	Max     interface{} `json:"max"`
	Min     interface{} `json:"min"`
	// For specifying result metrics.
	Query AccountStorageAnalyticsListResponseResultQuery `json:"query"`
	// Total number of rows in the result.
	Rows   float64                                       `json:"rows"`
	Totals interface{}                                   `json:"totals"`
	JSON   accountStorageAnalyticsListResponseResultJSON `json:"-"`
}

// accountStorageAnalyticsListResponseResultJSON contains the JSON metadata for the
// struct [AccountStorageAnalyticsListResponseResult]
type accountStorageAnalyticsListResponseResultJSON struct {
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

func (r *AccountStorageAnalyticsListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// For specifying result metrics.
type AccountStorageAnalyticsListResponseResultQuery struct {
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
	Until time.Time                                          `json:"until" format:"date-time"`
	JSON  accountStorageAnalyticsListResponseResultQueryJSON `json:"-"`
}

// accountStorageAnalyticsListResponseResultQueryJSON contains the JSON metadata
// for the struct [AccountStorageAnalyticsListResponseResultQuery]
type accountStorageAnalyticsListResponseResultQueryJSON struct {
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

func (r *AccountStorageAnalyticsListResponseResultQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStorageAnalyticsListResponseSuccess bool

const (
	AccountStorageAnalyticsListResponseSuccessTrue AccountStorageAnalyticsListResponseSuccess = true
)

type AccountStorageAnalyticsStoredResponse struct {
	Errors   []AccountStorageAnalyticsStoredResponseError   `json:"errors"`
	Messages []AccountStorageAnalyticsStoredResponseMessage `json:"messages"`
	Result   AccountStorageAnalyticsStoredResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountStorageAnalyticsStoredResponseSuccess `json:"success"`
	JSON    accountStorageAnalyticsStoredResponseJSON    `json:"-"`
}

// accountStorageAnalyticsStoredResponseJSON contains the JSON metadata for the
// struct [AccountStorageAnalyticsStoredResponse]
type accountStorageAnalyticsStoredResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageAnalyticsStoredResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageAnalyticsStoredResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountStorageAnalyticsStoredResponseErrorJSON `json:"-"`
}

// accountStorageAnalyticsStoredResponseErrorJSON contains the JSON metadata for
// the struct [AccountStorageAnalyticsStoredResponseError]
type accountStorageAnalyticsStoredResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageAnalyticsStoredResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageAnalyticsStoredResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountStorageAnalyticsStoredResponseMessageJSON `json:"-"`
}

// accountStorageAnalyticsStoredResponseMessageJSON contains the JSON metadata for
// the struct [AccountStorageAnalyticsStoredResponseMessage]
type accountStorageAnalyticsStoredResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountStorageAnalyticsStoredResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountStorageAnalyticsStoredResponseResult struct {
	Data interface{} `json:"data"`
	// Number of seconds between current time and last processed event, i.e. how many
	// seconds of data could be missing.
	DataLag float64     `json:"data_lag"`
	Max     interface{} `json:"max"`
	Min     interface{} `json:"min"`
	// For specifying result metrics.
	Query AccountStorageAnalyticsStoredResponseResultQuery `json:"query"`
	// Total number of rows in the result.
	Rows   float64                                         `json:"rows"`
	Totals interface{}                                     `json:"totals"`
	JSON   accountStorageAnalyticsStoredResponseResultJSON `json:"-"`
}

// accountStorageAnalyticsStoredResponseResultJSON contains the JSON metadata for
// the struct [AccountStorageAnalyticsStoredResponseResult]
type accountStorageAnalyticsStoredResponseResultJSON struct {
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

func (r *AccountStorageAnalyticsStoredResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// For specifying result metrics.
type AccountStorageAnalyticsStoredResponseResultQuery struct {
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
	Until time.Time                                            `json:"until" format:"date-time"`
	JSON  accountStorageAnalyticsStoredResponseResultQueryJSON `json:"-"`
}

// accountStorageAnalyticsStoredResponseResultQueryJSON contains the JSON metadata
// for the struct [AccountStorageAnalyticsStoredResponseResultQuery]
type accountStorageAnalyticsStoredResponseResultQueryJSON struct {
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

func (r *AccountStorageAnalyticsStoredResponseResultQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountStorageAnalyticsStoredResponseSuccess bool

const (
	AccountStorageAnalyticsStoredResponseSuccessTrue AccountStorageAnalyticsStoredResponseSuccess = true
)

type AccountStorageAnalyticsListParams struct {
	Query param.Field[AccountStorageAnalyticsListParamsQuery] `query:"query"`
}

// URLQuery serializes [AccountStorageAnalyticsListParams]'s query parameters as
// `url.Values`.
func (r AccountStorageAnalyticsListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountStorageAnalyticsListParamsQuery struct {
	Dimensions param.Field[[]AccountStorageAnalyticsListParamsQueryDimension] `query:"dimensions"`
	Filters    param.Field[interface{}]                                       `query:"filters"`
	// Limit number of returned metrics.
	Limit   param.Field[int64]                                          `query:"limit"`
	Metrics param.Field[[]AccountStorageAnalyticsListParamsQueryMetric] `query:"metrics"`
	// Start of time interval to query, defaults to 6 hours before request received.
	Since param.Field[time.Time]   `query:"since" format:"date-time"`
	Sort  param.Field[interface{}] `query:"sort"`
	// End of time interval to query, defaults to current time.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AccountStorageAnalyticsListParamsQuery]'s query parameters
// as `url.Values`.
func (r AccountStorageAnalyticsListParamsQuery) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountStorageAnalyticsListParamsQueryDimension string

const (
	AccountStorageAnalyticsListParamsQueryDimensionAccountID    AccountStorageAnalyticsListParamsQueryDimension = "accountId"
	AccountStorageAnalyticsListParamsQueryDimensionResponseCode AccountStorageAnalyticsListParamsQueryDimension = "responseCode"
	AccountStorageAnalyticsListParamsQueryDimensionRequestType  AccountStorageAnalyticsListParamsQueryDimension = "requestType"
)

type AccountStorageAnalyticsListParamsQueryMetric string

const (
	AccountStorageAnalyticsListParamsQueryMetricRequests AccountStorageAnalyticsListParamsQueryMetric = "requests"
	AccountStorageAnalyticsListParamsQueryMetricWriteKiB AccountStorageAnalyticsListParamsQueryMetric = "writeKiB"
	AccountStorageAnalyticsListParamsQueryMetricReadKiB  AccountStorageAnalyticsListParamsQueryMetric = "readKiB"
)

type AccountStorageAnalyticsStoredParams struct {
	Query param.Field[AccountStorageAnalyticsStoredParamsQuery] `query:"query"`
}

// URLQuery serializes [AccountStorageAnalyticsStoredParams]'s query parameters as
// `url.Values`.
func (r AccountStorageAnalyticsStoredParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountStorageAnalyticsStoredParamsQuery struct {
	Dimensions param.Field[[]AccountStorageAnalyticsStoredParamsQueryDimension] `query:"dimensions"`
	Filters    param.Field[interface{}]                                         `query:"filters"`
	// Limit number of returned metrics.
	Limit   param.Field[int64]                                            `query:"limit"`
	Metrics param.Field[[]AccountStorageAnalyticsStoredParamsQueryMetric] `query:"metrics"`
	// Start of time interval to query, defaults to 6 hours before request received.
	Since param.Field[time.Time]   `query:"since" format:"date-time"`
	Sort  param.Field[interface{}] `query:"sort"`
	// End of time interval to query, defaults to current time.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AccountStorageAnalyticsStoredParamsQuery]'s query
// parameters as `url.Values`.
func (r AccountStorageAnalyticsStoredParamsQuery) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AccountStorageAnalyticsStoredParamsQueryDimension string

const (
	AccountStorageAnalyticsStoredParamsQueryDimensionNamespaceID AccountStorageAnalyticsStoredParamsQueryDimension = "namespaceId"
)

type AccountStorageAnalyticsStoredParamsQueryMetric string

const (
	AccountStorageAnalyticsStoredParamsQueryMetricStoredBytes AccountStorageAnalyticsStoredParamsQueryMetric = "storedBytes"
	AccountStorageAnalyticsStoredParamsQueryMetricStoredKeys  AccountStorageAnalyticsStoredParamsQueryMetric = "storedKeys"
)
