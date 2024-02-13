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

// DNSFirewallDNSAnalyticReportBytimeService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDNSFirewallDNSAnalyticReportBytimeService] method instead.
type DNSFirewallDNSAnalyticReportBytimeService struct {
	Options []option.RequestOption
}

// NewDNSFirewallDNSAnalyticReportBytimeService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewDNSFirewallDNSAnalyticReportBytimeService(opts ...option.RequestOption) (r *DNSFirewallDNSAnalyticReportBytimeService) {
	r = &DNSFirewallDNSAnalyticReportBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *DNSFirewallDNSAnalyticReportBytimeService) List(ctx context.Context, accountIdentifier string, identifier string, query DNSFirewallDNSAnalyticReportBytimeListParams, opts ...option.RequestOption) (res *DNSFirewallDNSAnalyticReportBytimeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSFirewallDNSAnalyticReportBytimeListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report/bytime", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSFirewallDNSAnalyticReportBytimeListResponse struct {
	// Array with one row per combination of dimension values.
	Data []DNSFirewallDNSAnalyticReportBytimeListResponseData `json:"data,required"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}                                         `json:"min,required"`
	Query DNSFirewallDNSAnalyticReportBytimeListResponseQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Array of time intervals in the response data. Each interval is represented as an
	// array containing two values: the start time, and the end time.
	TimeIntervals [][]time.Time `json:"time_intervals,required" format:"date-time"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                                        `json:"totals,required"`
	JSON   dnsFirewallDNSAnalyticReportBytimeListResponseJSON `json:"-"`
}

// dnsFirewallDNSAnalyticReportBytimeListResponseJSON contains the JSON metadata
// for the struct [DNSFirewallDNSAnalyticReportBytimeListResponse]
type dnsFirewallDNSAnalyticReportBytimeListResponseJSON struct {
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

func (r *DNSFirewallDNSAnalyticReportBytimeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallDNSAnalyticReportBytimeListResponseData struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is an array of values,
	// broken down by time interval.
	Metrics [][]interface{}                                        `json:"metrics,required"`
	JSON    dnsFirewallDNSAnalyticReportBytimeListResponseDataJSON `json:"-"`
}

// dnsFirewallDNSAnalyticReportBytimeListResponseDataJSON contains the JSON
// metadata for the struct [DNSFirewallDNSAnalyticReportBytimeListResponseData]
type dnsFirewallDNSAnalyticReportBytimeListResponseDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDNSAnalyticReportBytimeListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallDNSAnalyticReportBytimeListResponseQuery struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// Unit of time to group data by.
	TimeDelta DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDelta `json:"time_delta,required"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string                                                `json:"sort"`
	JSON dnsFirewallDNSAnalyticReportBytimeListResponseQueryJSON `json:"-"`
}

// dnsFirewallDNSAnalyticReportBytimeListResponseQueryJSON contains the JSON
// metadata for the struct [DNSFirewallDNSAnalyticReportBytimeListResponseQuery]
type dnsFirewallDNSAnalyticReportBytimeListResponseQueryJSON struct {
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

func (r *DNSFirewallDNSAnalyticReportBytimeListResponseQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Unit of time to group data by.
type DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDelta string

const (
	DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDeltaAll        DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDelta = "all"
	DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDeltaAuto       DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDelta = "auto"
	DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDeltaYear       DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDelta = "year"
	DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDeltaQuarter    DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDelta = "quarter"
	DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDeltaMonth      DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDelta = "month"
	DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDeltaWeek       DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDelta = "week"
	DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDeltaDay        DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDelta = "day"
	DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDeltaHour       DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDelta = "hour"
	DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDeltaDekaminute DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDelta = "dekaminute"
	DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDeltaMinute     DNSFirewallDNSAnalyticReportBytimeListResponseQueryTimeDelta = "minute"
)

type DNSFirewallDNSAnalyticReportBytimeListParams struct {
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
	TimeDelta param.Field[DNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [DNSFirewallDNSAnalyticReportBytimeListParams]'s query
// parameters as `url.Values`.
func (r DNSFirewallDNSAnalyticReportBytimeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Unit of time to group data by.
type DNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta string

const (
	DNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaAll        DNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "all"
	DNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaAuto       DNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "auto"
	DNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaYear       DNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "year"
	DNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaQuarter    DNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "quarter"
	DNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaMonth      DNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "month"
	DNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaWeek       DNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "week"
	DNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaDay        DNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "day"
	DNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaHour       DNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "hour"
	DNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaDekaminute DNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "dekaminute"
	DNSFirewallDNSAnalyticReportBytimeListParamsTimeDeltaMinute     DNSFirewallDNSAnalyticReportBytimeListParamsTimeDelta = "minute"
)

type DNSFirewallDNSAnalyticReportBytimeListResponseEnvelope struct {
	Errors   []DNSFirewallDNSAnalyticReportBytimeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSFirewallDNSAnalyticReportBytimeListResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSFirewallDNSAnalyticReportBytimeListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSFirewallDNSAnalyticReportBytimeListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsFirewallDNSAnalyticReportBytimeListResponseEnvelopeJSON    `json:"-"`
}

// dnsFirewallDNSAnalyticReportBytimeListResponseEnvelopeJSON contains the JSON
// metadata for the struct [DNSFirewallDNSAnalyticReportBytimeListResponseEnvelope]
type dnsFirewallDNSAnalyticReportBytimeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDNSAnalyticReportBytimeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallDNSAnalyticReportBytimeListResponseEnvelopeErrors struct {
	Code    int64                                                            `json:"code,required"`
	Message string                                                           `json:"message,required"`
	JSON    dnsFirewallDNSAnalyticReportBytimeListResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsFirewallDNSAnalyticReportBytimeListResponseEnvelopeErrorsJSON contains the
// JSON metadata for the struct
// [DNSFirewallDNSAnalyticReportBytimeListResponseEnvelopeErrors]
type dnsFirewallDNSAnalyticReportBytimeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDNSAnalyticReportBytimeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSFirewallDNSAnalyticReportBytimeListResponseEnvelopeMessages struct {
	Code    int64                                                              `json:"code,required"`
	Message string                                                             `json:"message,required"`
	JSON    dnsFirewallDNSAnalyticReportBytimeListResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsFirewallDNSAnalyticReportBytimeListResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DNSFirewallDNSAnalyticReportBytimeListResponseEnvelopeMessages]
type dnsFirewallDNSAnalyticReportBytimeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSFirewallDNSAnalyticReportBytimeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSFirewallDNSAnalyticReportBytimeListResponseEnvelopeSuccess bool

const (
	DNSFirewallDNSAnalyticReportBytimeListResponseEnvelopeSuccessTrue DNSFirewallDNSAnalyticReportBytimeListResponseEnvelopeSuccess = true
)
