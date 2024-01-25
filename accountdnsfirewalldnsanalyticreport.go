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

// AccountDNSFirewallDNSAnalyticReportService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountDNSFirewallDNSAnalyticReportService] method instead.
type AccountDNSFirewallDNSAnalyticReportService struct {
	Options []option.RequestOption
	Bytimes *AccountDNSFirewallDNSAnalyticReportBytimeService
}

// NewAccountDNSFirewallDNSAnalyticReportService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewAccountDNSFirewallDNSAnalyticReportService(opts ...option.RequestOption) (r *AccountDNSFirewallDNSAnalyticReportService) {
	r = &AccountDNSFirewallDNSAnalyticReportService{}
	r.Options = opts
	r.Bytimes = NewAccountDNSFirewallDNSAnalyticReportBytimeService(opts...)
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *AccountDNSFirewallDNSAnalyticReportService) List(ctx context.Context, accountIdentifier string, identifier string, query AccountDNSFirewallDNSAnalyticReportListParams, opts ...option.RequestOption) (res *AccountDNSFirewallDNSAnalyticReportListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AccountDNSFirewallDNSAnalyticReportListResponse struct {
	Errors   []AccountDNSFirewallDNSAnalyticReportListResponseError   `json:"errors"`
	Messages []AccountDNSFirewallDNSAnalyticReportListResponseMessage `json:"messages"`
	Result   AccountDNSFirewallDNSAnalyticReportListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success AccountDNSFirewallDNSAnalyticReportListResponseSuccess `json:"success"`
	JSON    accountDNSFirewallDNSAnalyticReportListResponseJSON    `json:"-"`
}

// accountDNSFirewallDNSAnalyticReportListResponseJSON contains the JSON metadata
// for the struct [AccountDNSFirewallDNSAnalyticReportListResponse]
type accountDNSFirewallDNSAnalyticReportListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSAnalyticReportListResponseError struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    accountDNSFirewallDNSAnalyticReportListResponseErrorJSON `json:"-"`
}

// accountDNSFirewallDNSAnalyticReportListResponseErrorJSON contains the JSON
// metadata for the struct [AccountDNSFirewallDNSAnalyticReportListResponseError]
type accountDNSFirewallDNSAnalyticReportListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSAnalyticReportListResponseMessage struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    accountDNSFirewallDNSAnalyticReportListResponseMessageJSON `json:"-"`
}

// accountDNSFirewallDNSAnalyticReportListResponseMessageJSON contains the JSON
// metadata for the struct [AccountDNSFirewallDNSAnalyticReportListResponseMessage]
type accountDNSFirewallDNSAnalyticReportListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSAnalyticReportListResponseResult struct {
	Data []AccountDNSFirewallDNSAnalyticReportListResponseResultData `json:"data"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}                                                `json:"min"`
	Query AccountDNSFirewallDNSAnalyticReportListResponseResultQuery `json:"query"`
	// Total number of rows in the result.
	Rows float64 `json:"rows"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                                               `json:"totals"`
	JSON   accountDNSFirewallDNSAnalyticReportListResponseResultJSON `json:"-"`
}

// accountDNSFirewallDNSAnalyticReportListResponseResultJSON contains the JSON
// metadata for the struct [AccountDNSFirewallDNSAnalyticReportListResponseResult]
type accountDNSFirewallDNSAnalyticReportListResponseResultJSON struct {
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

func (r *AccountDNSFirewallDNSAnalyticReportListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSAnalyticReportListResponseResultData struct {
	// Array with one item per requested metric. Each item is a single value.
	Metrics []float64                                                     `json:"metrics,required"`
	JSON    accountDNSFirewallDNSAnalyticReportListResponseResultDataJSON `json:"-"`
}

// accountDNSFirewallDNSAnalyticReportListResponseResultDataJSON contains the JSON
// metadata for the struct
// [AccountDNSFirewallDNSAnalyticReportListResponseResultData]
type accountDNSFirewallDNSAnalyticReportListResponseResultDataJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportListResponseResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountDNSFirewallDNSAnalyticReportListResponseResultQuery struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string                                                       `json:"sort"`
	JSON accountDNSFirewallDNSAnalyticReportListResponseResultQueryJSON `json:"-"`
}

// accountDNSFirewallDNSAnalyticReportListResponseResultQueryJSON contains the JSON
// metadata for the struct
// [AccountDNSFirewallDNSAnalyticReportListResponseResultQuery]
type accountDNSFirewallDNSAnalyticReportListResponseResultQueryJSON struct {
	Dimensions  apijson.Field
	Limit       apijson.Field
	Metrics     apijson.Field
	Since       apijson.Field
	Until       apijson.Field
	Filters     apijson.Field
	Sort        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountDNSFirewallDNSAnalyticReportListResponseResultQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountDNSFirewallDNSAnalyticReportListResponseSuccess bool

const (
	AccountDNSFirewallDNSAnalyticReportListResponseSuccessTrue AccountDNSFirewallDNSAnalyticReportListResponseSuccess = true
)

type AccountDNSFirewallDNSAnalyticReportListParams struct {
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
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AccountDNSFirewallDNSAnalyticReportListParams]'s query
// parameters as `url.Values`.
func (r AccountDNSFirewallDNSAnalyticReportListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
