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

// AccountDNSFirewallDNSAnalyticReportBytimeService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccountDNSFirewallDNSAnalyticReportBytimeService] method instead.
type AccountDNSFirewallDNSAnalyticReportBytimeService struct {
	Options []option.RequestOption
}

// NewAccountDNSFirewallDNSAnalyticReportBytimeService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountDNSFirewallDNSAnalyticReportBytimeService(opts ...option.RequestOption) (r *AccountDNSFirewallDNSAnalyticReportBytimeService) {
	r = &AccountDNSFirewallDNSAnalyticReportBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *AccountDNSFirewallDNSAnalyticReportBytimeService) List(ctx context.Context, accountIdentifier string, identifier string, query AccountDNSFirewallDNSAnalyticReportBytimeListParams, opts ...option.RequestOption) (res *AccountDNSFirewallDNSAnalyticReportBytimeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report/bytime", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountDNSFirewallDNSAnalyticReportBytimeListResponse struct {
	Errors   []AccountDNSFirewallDNSAnalyticReportBytimeListResponseError   `json:"errors"`
	Messages []AccountDNSFirewallDNSAnalyticReportBytimeListResponseMessage `json:"messages"`
	Result   AccountDNSFirewallDNSAnalyticReportBytimeListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDNSFirewallDNSAnalyticReportBytimeListResponseSuccess `json:"success"`
	JSON    accountDNSFirewallDNSAnalyticReportBytimeListResponseJSON    `json:"-"`
}

// accountDNSFirewallDNSAnalyticReportBytimeListResponseJSON contains the JSON
// metadata for the struct [AccountDNSFirewallDNSAnalyticReportBytimeListResponse]
type accountDNSFirewallDNSAnalyticReportBytimeListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportBytimeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSAnalyticReportBytimeListResponseError struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    accountDNSFirewallDNSAnalyticReportBytimeListResponseErrorJSON `json:"-"`
}

// accountDNSFirewallDNSAnalyticReportBytimeListResponseErrorJSON contains the JSON
// metadata for the struct
// [AccountDNSFirewallDNSAnalyticReportBytimeListResponseError]
type accountDNSFirewallDNSAnalyticReportBytimeListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportBytimeListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSAnalyticReportBytimeListResponseMessage struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    accountDNSFirewallDNSAnalyticReportBytimeListResponseMessageJSON `json:"-"`
}

// accountDNSFirewallDNSAnalyticReportBytimeListResponseMessageJSON contains the
// JSON metadata for the struct
// [AccountDNSFirewallDNSAnalyticReportBytimeListResponseMessage]
type accountDNSFirewallDNSAnalyticReportBytimeListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportBytimeListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSAnalyticReportBytimeListResponseResult struct {
	Data []AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultData `json:"data"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}                                                      `json:"min"`
	Query AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQuery `json:"query"`
	// Total number of rows in the result.
	Rows float64 `json:"rows"`
	// Array of time intervals in the response data. Each interval is represented as an
	// array containing two values: the start time, and the end time.
	TimeIntervals [][]time.Time `json:"time_intervals" format:"date-time"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                                                     `json:"totals"`
	JSON   accountDNSFirewallDNSAnalyticReportBytimeListResponseResultJSON `json:"-"`
}

// accountDNSFirewallDNSAnalyticReportBytimeListResponseResultJSON contains the
// JSON metadata for the struct
// [AccountDNSFirewallDNSAnalyticReportBytimeListResponseResult]
type accountDNSFirewallDNSAnalyticReportBytimeListResponseResultJSON struct {
	Data          apijson.Field
	DataLag       apijson.Field
	Max           apijson.Field
	Min           apijson.Field
	Query         apijson.Field
	Rows          apijson.Field
	TimeIntervals apijson.Field
	Totals        apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportBytimeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultData struct {
	// Array with one item per requested metric. Each item is an array of values,
	// broken down by time interval.
	Metrics [][]interface{}                                                     `json:"metrics,required"`
	JSON    accountDNSFirewallDNSAnalyticReportBytimeListResponseResultDataJSON `json:"-"`
}

// accountDNSFirewallDNSAnalyticReportBytimeListResponseResultDataJSON contains the
// JSON metadata for the struct
// [AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultData]
type accountDNSFirewallDNSAnalyticReportBytimeListResponseResultDataJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQuery struct {
	// Unit of time to group data by.
	TimeDelta AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDelta `json:"time_delta,required"`
	JSON      accountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryJSON      `json:"-"`
}

// accountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryJSON contains
// the JSON metadata for the struct
// [AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQuery]
type accountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryJSON struct {
	TimeDelta   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Unit of time to group data by.
type AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDelta string

const (
	AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaAll        AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "all"
	AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaAuto       AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "auto"
	AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaYear       AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "year"
	AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaQuarter    AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "quarter"
	AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaMonth      AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "month"
	AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaWeek       AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "week"
	AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaDay        AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "day"
	AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaHour       AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "hour"
	AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaDekaminute AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "dekaminute"
	AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaMinute     AccountDNSFirewallDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "minute"
)

// Whether the API call was successful
type AccountDNSFirewallDNSAnalyticReportBytimeListResponseSuccess bool

const (
	AccountDNSFirewallDNSAnalyticReportBytimeListResponseSuccessTrue AccountDNSFirewallDNSAnalyticReportBytimeListResponseSuccess = true
)

type AccountDNSFirewallDNSAnalyticReportBytimeListParams struct {
	// A comma-separated list of dimensions to group results by.
	Dimensions param.Field[string] `query:"dimensions"`
	// Segmentation filter in 'attribute operator value' format.
	Filters param.Field[string] `query:"filters"`
	// Limit number of returned metrics.
	Limit param.Field[int64] `query:"limit"`
	// A comma-separated list of metrics to query.
	Metrics param.Field[string] `query:"metrics"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// A comma-separated list of dimensions to sort by, where each dimension may be
	// prefixed by - (descending) or + (ascending).
	Sort param.Field[string] `query:"sort"`
	// Unit of time to group data by.
	TimeDelta param.Field[AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AccountDNSFirewallDNSAnalyticReportBytimeListParams]'s
// query parameters as `url.Values`.
func (r AccountDNSFirewallDNSAnalyticReportBytimeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Unit of time to group data by.
type AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta string

const (
	AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaAll        AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "all"
	AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaAuto       AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "auto"
	AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaYear       AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "year"
	AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaQuarter    AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "quarter"
	AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaMonth      AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "month"
	AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaWeek       AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "week"
	AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaDay        AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "day"
	AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaHour       AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "hour"
	AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaDekaminute AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "dekaminute"
	AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaMinute     AccountDNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "minute"
)
