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

// DNSFirewallAnalyticsReportBytimeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDNSFirewallAnalyticsReportBytimeService] method instead.
type DNSFirewallAnalyticsReportBytimeService struct {
	Options []option.RequestOption
}

// NewDNSFirewallAnalyticsReportBytimeService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDNSFirewallAnalyticsReportBytimeService(opts ...option.RequestOption) (r *DNSFirewallAnalyticsReportBytimeService) {
	r = &DNSFirewallAnalyticsReportBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *DNSFirewallAnalyticsReportBytimeService) List(ctx context.Context, accountIdentifier string, identifier string, query DNSFirewallAnalyticsReportBytimeListParams, opts ...option.RequestOption) (res *DNSFirewallAnalyticsReportBytimeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSFirewallAnalyticsReportBytimeListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report/bytime", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSFirewallAnalyticsReportBytimeListResponse struct {
	// Array with one row per combination of dimension values.
	Data []DNSFirewallAnalyticsReportBytimeListResponseData `json:"data,required"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}                                       `json:"min,required"`
	Query DNSFirewallAnalyticsReportBytimeListResponseQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Array of time intervals in the response data. Each interval is represented as an
	// array containing two values: the start time, and the end time.
	TimeIntervals [][]time.Time `json:"time_intervals,required" format:"date-time"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                                      `json:"totals,required"`
	JSON   dnsFirewallAnalyticsReportBytimeListResponseJSON `json:"-"`
}

// dnsFirewallAnalyticsReportBytimeListResponseJSON contains the JSON metadata for
// the struct [DNSFirewallAnalyticsReportBytimeListResponse]
type dnsFirewallAnalyticsReportBytimeListResponseJSON struct {
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

func (r *DNSFirewallAnalyticsReportBytimeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallAnalyticsReportBytimeListResponseData struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is an array of values,
	// broken down by time interval.
	Metrics [][]interface{}                                      `json:"metrics,required"`
	JSON    dnsFirewallAnalyticsReportBytimeListResponseDataJSON `json:"-"`
}

// dnsFirewallAnalyticsReportBytimeListResponseDataJSON contains the JSON metadata
// for the struct [DNSFirewallAnalyticsReportBytimeListResponseData]
type dnsFirewallAnalyticsReportBytimeListResponseDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallAnalyticsReportBytimeListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallAnalyticsReportBytimeListResponseQuery struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// Unit of time to group data by.
	TimeDelta DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDelta `json:"time_delta,required"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string                                              `json:"sort"`
	JSON dnsFirewallAnalyticsReportBytimeListResponseQueryJSON `json:"-"`
}

// dnsFirewallAnalyticsReportBytimeListResponseQueryJSON contains the JSON metadata
// for the struct [DNSFirewallAnalyticsReportBytimeListResponseQuery]
type dnsFirewallAnalyticsReportBytimeListResponseQueryJSON struct {
	Dimensions  apijson.Field
	Limit       apijson.Field
	Metrics     apijson.Field
	Since       apijson.Field
	TimeDelta   apijson.Field
	Until       apijson.Field
	Filters     apijson.Field
	Sort        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallAnalyticsReportBytimeListResponseQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Unit of time to group data by.
type DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDelta string

const (
	DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDeltaAll        DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDelta = "all"
	DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDeltaAuto       DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDelta = "auto"
	DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDeltaYear       DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDelta = "year"
	DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDeltaQuarter    DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDelta = "quarter"
	DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDeltaMonth      DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDelta = "month"
	DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDeltaWeek       DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDelta = "week"
	DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDeltaDay        DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDelta = "day"
	DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDeltaHour       DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDelta = "hour"
	DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDeltaDekaminute DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDelta = "dekaminute"
	DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDeltaMinute     DNSFirewallAnalyticsReportBytimeListResponseQueryTimeDelta = "minute"
)

type DNSFirewallAnalyticsReportBytimeListParams struct {
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
	TimeDelta param.Field[DNSFirewallAnalyticsReportBytimeListParamsTimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [DNSFirewallAnalyticsReportBytimeListParams]'s query
// parameters as `url.Values`.
func (r DNSFirewallAnalyticsReportBytimeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Unit of time to group data by.
type DNSFirewallAnalyticsReportBytimeListParamsTimeDelta string

const (
	DNSFirewallAnalyticsReportBytimeListParamsTimeDeltaAll        DNSFirewallAnalyticsReportBytimeListParamsTimeDelta = "all"
	DNSFirewallAnalyticsReportBytimeListParamsTimeDeltaAuto       DNSFirewallAnalyticsReportBytimeListParamsTimeDelta = "auto"
	DNSFirewallAnalyticsReportBytimeListParamsTimeDeltaYear       DNSFirewallAnalyticsReportBytimeListParamsTimeDelta = "year"
	DNSFirewallAnalyticsReportBytimeListParamsTimeDeltaQuarter    DNSFirewallAnalyticsReportBytimeListParamsTimeDelta = "quarter"
	DNSFirewallAnalyticsReportBytimeListParamsTimeDeltaMonth      DNSFirewallAnalyticsReportBytimeListParamsTimeDelta = "month"
	DNSFirewallAnalyticsReportBytimeListParamsTimeDeltaWeek       DNSFirewallAnalyticsReportBytimeListParamsTimeDelta = "week"
	DNSFirewallAnalyticsReportBytimeListParamsTimeDeltaDay        DNSFirewallAnalyticsReportBytimeListParamsTimeDelta = "day"
	DNSFirewallAnalyticsReportBytimeListParamsTimeDeltaHour       DNSFirewallAnalyticsReportBytimeListParamsTimeDelta = "hour"
	DNSFirewallAnalyticsReportBytimeListParamsTimeDeltaDekaminute DNSFirewallAnalyticsReportBytimeListParamsTimeDelta = "dekaminute"
	DNSFirewallAnalyticsReportBytimeListParamsTimeDeltaMinute     DNSFirewallAnalyticsReportBytimeListParamsTimeDelta = "minute"
)

type DNSFirewallAnalyticsReportBytimeListResponseEnvelope struct {
	Errors   []DNSFirewallAnalyticsReportBytimeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallAnalyticsReportBytimeListResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSFirewallAnalyticsReportBytimeListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSFirewallAnalyticsReportBytimeListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsFirewallAnalyticsReportBytimeListResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallAnalyticsReportBytimeListResponseEnvelopeJSON contains the JSON
// metadata for the struct [DNSFirewallAnalyticsReportBytimeListResponseEnvelope]
type dnsFirewallAnalyticsReportBytimeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallAnalyticsReportBytimeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallAnalyticsReportBytimeListResponseEnvelopeErrors struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    dnsFirewallAnalyticsReportBytimeListResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsFirewallAnalyticsReportBytimeListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DNSFirewallAnalyticsReportBytimeListResponseEnvelopeErrors]
type dnsFirewallAnalyticsReportBytimeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallAnalyticsReportBytimeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallAnalyticsReportBytimeListResponseEnvelopeMessages struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    dnsFirewallAnalyticsReportBytimeListResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsFirewallAnalyticsReportBytimeListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DNSFirewallAnalyticsReportBytimeListResponseEnvelopeMessages]
type dnsFirewallAnalyticsReportBytimeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallAnalyticsReportBytimeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSFirewallAnalyticsReportBytimeListResponseEnvelopeSuccess bool

const (
	DNSFirewallAnalyticsReportBytimeListResponseEnvelopeSuccessTrue DNSFirewallAnalyticsReportBytimeListResponseEnvelopeSuccess = true
)
