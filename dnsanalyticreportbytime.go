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

// DNSAnalyticReportBytimeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewDNSAnalyticReportBytimeService] method instead.
type DNSAnalyticReportBytimeService struct {
	Options []option.RequestOption
}

// NewDNSAnalyticReportBytimeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDNSAnalyticReportBytimeService(opts ...option.RequestOption) (r *DNSAnalyticReportBytimeService) {
	r = &DNSAnalyticReportBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *DNSAnalyticReportBytimeService) List(ctx context.Context, identifier string, query DNSAnalyticReportBytimeListParams, opts ...option.RequestOption) (res *DNSAnalyticReportBytimeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSAnalyticReportBytimeListResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_analytics/report/bytime", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSAnalyticReportBytimeListResponse struct {
	// Array with one row per combination of dimension values.
	Data []DNSAnalyticReportBytimeListResponseData `json:"data,required"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}                              `json:"min,required"`
	Query DNSAnalyticReportBytimeListResponseQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Array of time intervals in the response data. Each interval is represented as an
	// array containing two values: the start time, and the end time.
	TimeIntervals [][]time.Time `json:"time_intervals,required" format:"date-time"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                             `json:"totals,required"`
	JSON   dnsAnalyticReportBytimeListResponseJSON `json:"-"`
}

// dnsAnalyticReportBytimeListResponseJSON contains the JSON metadata for the
// struct [DNSAnalyticReportBytimeListResponse]
type dnsAnalyticReportBytimeListResponseJSON struct {
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

func (r *DNSAnalyticReportBytimeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSAnalyticReportBytimeListResponseData struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is an array of values,
	// broken down by time interval.
	Metrics [][]interface{}                             `json:"metrics,required"`
	JSON    dnsAnalyticReportBytimeListResponseDataJSON `json:"-"`
}

// dnsAnalyticReportBytimeListResponseDataJSON contains the JSON metadata for the
// struct [DNSAnalyticReportBytimeListResponseData]
type dnsAnalyticReportBytimeListResponseDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticReportBytimeListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSAnalyticReportBytimeListResponseQuery struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// Unit of time to group data by.
	TimeDelta DNSAnalyticReportBytimeListResponseQueryTimeDelta `json:"time_delta,required"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string                                     `json:"sort"`
	JSON dnsAnalyticReportBytimeListResponseQueryJSON `json:"-"`
}

// dnsAnalyticReportBytimeListResponseQueryJSON contains the JSON metadata for the
// struct [DNSAnalyticReportBytimeListResponseQuery]
type dnsAnalyticReportBytimeListResponseQueryJSON struct {
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

func (r *DNSAnalyticReportBytimeListResponseQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Unit of time to group data by.
type DNSAnalyticReportBytimeListResponseQueryTimeDelta string

const (
	DNSAnalyticReportBytimeListResponseQueryTimeDeltaAll        DNSAnalyticReportBytimeListResponseQueryTimeDelta = "all"
	DNSAnalyticReportBytimeListResponseQueryTimeDeltaAuto       DNSAnalyticReportBytimeListResponseQueryTimeDelta = "auto"
	DNSAnalyticReportBytimeListResponseQueryTimeDeltaYear       DNSAnalyticReportBytimeListResponseQueryTimeDelta = "year"
	DNSAnalyticReportBytimeListResponseQueryTimeDeltaQuarter    DNSAnalyticReportBytimeListResponseQueryTimeDelta = "quarter"
	DNSAnalyticReportBytimeListResponseQueryTimeDeltaMonth      DNSAnalyticReportBytimeListResponseQueryTimeDelta = "month"
	DNSAnalyticReportBytimeListResponseQueryTimeDeltaWeek       DNSAnalyticReportBytimeListResponseQueryTimeDelta = "week"
	DNSAnalyticReportBytimeListResponseQueryTimeDeltaDay        DNSAnalyticReportBytimeListResponseQueryTimeDelta = "day"
	DNSAnalyticReportBytimeListResponseQueryTimeDeltaHour       DNSAnalyticReportBytimeListResponseQueryTimeDelta = "hour"
	DNSAnalyticReportBytimeListResponseQueryTimeDeltaDekaminute DNSAnalyticReportBytimeListResponseQueryTimeDelta = "dekaminute"
	DNSAnalyticReportBytimeListResponseQueryTimeDeltaMinute     DNSAnalyticReportBytimeListResponseQueryTimeDelta = "minute"
)

type DNSAnalyticReportBytimeListParams struct {
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
	TimeDelta param.Field[DNSAnalyticReportBytimeListParamsTimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [DNSAnalyticReportBytimeListParams]'s query parameters as
// `url.Values`.
func (r DNSAnalyticReportBytimeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Unit of time to group data by.
type DNSAnalyticReportBytimeListParamsTimeDelta string

const (
	DNSAnalyticReportBytimeListParamsTimeDeltaAll        DNSAnalyticReportBytimeListParamsTimeDelta = "all"
	DNSAnalyticReportBytimeListParamsTimeDeltaAuto       DNSAnalyticReportBytimeListParamsTimeDelta = "auto"
	DNSAnalyticReportBytimeListParamsTimeDeltaYear       DNSAnalyticReportBytimeListParamsTimeDelta = "year"
	DNSAnalyticReportBytimeListParamsTimeDeltaQuarter    DNSAnalyticReportBytimeListParamsTimeDelta = "quarter"
	DNSAnalyticReportBytimeListParamsTimeDeltaMonth      DNSAnalyticReportBytimeListParamsTimeDelta = "month"
	DNSAnalyticReportBytimeListParamsTimeDeltaWeek       DNSAnalyticReportBytimeListParamsTimeDelta = "week"
	DNSAnalyticReportBytimeListParamsTimeDeltaDay        DNSAnalyticReportBytimeListParamsTimeDelta = "day"
	DNSAnalyticReportBytimeListParamsTimeDeltaHour       DNSAnalyticReportBytimeListParamsTimeDelta = "hour"
	DNSAnalyticReportBytimeListParamsTimeDeltaDekaminute DNSAnalyticReportBytimeListParamsTimeDelta = "dekaminute"
	DNSAnalyticReportBytimeListParamsTimeDeltaMinute     DNSAnalyticReportBytimeListParamsTimeDelta = "minute"
)

type DNSAnalyticReportBytimeListResponseEnvelope struct {
	Errors   []DNSAnalyticReportBytimeListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSAnalyticReportBytimeListResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSAnalyticReportBytimeListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSAnalyticReportBytimeListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsAnalyticReportBytimeListResponseEnvelopeJSON    `json:"-"`
}

// dnsAnalyticReportBytimeListResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSAnalyticReportBytimeListResponseEnvelope]
type dnsAnalyticReportBytimeListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticReportBytimeListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSAnalyticReportBytimeListResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    dnsAnalyticReportBytimeListResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsAnalyticReportBytimeListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DNSAnalyticReportBytimeListResponseEnvelopeErrors]
type dnsAnalyticReportBytimeListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticReportBytimeListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSAnalyticReportBytimeListResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    dnsAnalyticReportBytimeListResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsAnalyticReportBytimeListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DNSAnalyticReportBytimeListResponseEnvelopeMessages]
type dnsAnalyticReportBytimeListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticReportBytimeListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSAnalyticReportBytimeListResponseEnvelopeSuccess bool

const (
	DNSAnalyticReportBytimeListResponseEnvelopeSuccessTrue DNSAnalyticReportBytimeListResponseEnvelopeSuccess = true
)
