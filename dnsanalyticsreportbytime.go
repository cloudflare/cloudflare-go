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

// DNSAnalyticsReportBytimeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewDNSAnalyticsReportBytimeService] method instead.
type DNSAnalyticsReportBytimeService struct {
	Options []option.RequestOption
}

// NewDNSAnalyticsReportBytimeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDNSAnalyticsReportBytimeService(opts ...option.RequestOption) (r *DNSAnalyticsReportBytimeService) {
	r = &DNSAnalyticsReportBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *DNSAnalyticsReportBytimeService) Get(ctx context.Context, identifier string, query DNSAnalyticsReportBytimeGetParams, opts ...option.RequestOption) (res *DNSAnalyticsReportBytimeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSAnalyticsReportBytimeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_analytics/report/bytime", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSAnalyticsReportBytimeGetResponse struct {
	// Array with one row per combination of dimension values.
	Data []DNSAnalyticsReportBytimeGetResponseData `json:"data,required"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}                              `json:"min,required"`
	Query DNSAnalyticsReportBytimeGetResponseQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Array of time intervals in the response data. Each interval is represented as an
	// array containing two values: the start time, and the end time.
	TimeIntervals [][]time.Time `json:"time_intervals,required" format:"date-time"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                             `json:"totals,required"`
	JSON   dnsAnalyticsReportBytimeGetResponseJSON `json:"-"`
}

// dnsAnalyticsReportBytimeGetResponseJSON contains the JSON metadata for the
// struct [DNSAnalyticsReportBytimeGetResponse]
type dnsAnalyticsReportBytimeGetResponseJSON struct {
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

func (r *DNSAnalyticsReportBytimeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsAnalyticsReportBytimeGetResponseJSON) RawJSON() string {
	return r.raw
}

type DNSAnalyticsReportBytimeGetResponseData struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is an array of values,
	// broken down by time interval.
	Metrics [][]interface{}                             `json:"metrics,required"`
	JSON    dnsAnalyticsReportBytimeGetResponseDataJSON `json:"-"`
}

// dnsAnalyticsReportBytimeGetResponseDataJSON contains the JSON metadata for the
// struct [DNSAnalyticsReportBytimeGetResponseData]
type dnsAnalyticsReportBytimeGetResponseDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticsReportBytimeGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsAnalyticsReportBytimeGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type DNSAnalyticsReportBytimeGetResponseQuery struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// Unit of time to group data by.
	TimeDelta DNSAnalyticsReportBytimeGetResponseQueryTimeDelta `json:"time_delta,required"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string                                     `json:"sort"`
	JSON dnsAnalyticsReportBytimeGetResponseQueryJSON `json:"-"`
}

// dnsAnalyticsReportBytimeGetResponseQueryJSON contains the JSON metadata for the
// struct [DNSAnalyticsReportBytimeGetResponseQuery]
type dnsAnalyticsReportBytimeGetResponseQueryJSON struct {
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

func (r *DNSAnalyticsReportBytimeGetResponseQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsAnalyticsReportBytimeGetResponseQueryJSON) RawJSON() string {
	return r.raw
}

// Unit of time to group data by.
type DNSAnalyticsReportBytimeGetResponseQueryTimeDelta string

const (
	DNSAnalyticsReportBytimeGetResponseQueryTimeDeltaAll        DNSAnalyticsReportBytimeGetResponseQueryTimeDelta = "all"
	DNSAnalyticsReportBytimeGetResponseQueryTimeDeltaAuto       DNSAnalyticsReportBytimeGetResponseQueryTimeDelta = "auto"
	DNSAnalyticsReportBytimeGetResponseQueryTimeDeltaYear       DNSAnalyticsReportBytimeGetResponseQueryTimeDelta = "year"
	DNSAnalyticsReportBytimeGetResponseQueryTimeDeltaQuarter    DNSAnalyticsReportBytimeGetResponseQueryTimeDelta = "quarter"
	DNSAnalyticsReportBytimeGetResponseQueryTimeDeltaMonth      DNSAnalyticsReportBytimeGetResponseQueryTimeDelta = "month"
	DNSAnalyticsReportBytimeGetResponseQueryTimeDeltaWeek       DNSAnalyticsReportBytimeGetResponseQueryTimeDelta = "week"
	DNSAnalyticsReportBytimeGetResponseQueryTimeDeltaDay        DNSAnalyticsReportBytimeGetResponseQueryTimeDelta = "day"
	DNSAnalyticsReportBytimeGetResponseQueryTimeDeltaHour       DNSAnalyticsReportBytimeGetResponseQueryTimeDelta = "hour"
	DNSAnalyticsReportBytimeGetResponseQueryTimeDeltaDekaminute DNSAnalyticsReportBytimeGetResponseQueryTimeDelta = "dekaminute"
	DNSAnalyticsReportBytimeGetResponseQueryTimeDeltaMinute     DNSAnalyticsReportBytimeGetResponseQueryTimeDelta = "minute"
)

type DNSAnalyticsReportBytimeGetParams struct {
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
	TimeDelta param.Field[DNSAnalyticsReportBytimeGetParamsTimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [DNSAnalyticsReportBytimeGetParams]'s query parameters as
// `url.Values`.
func (r DNSAnalyticsReportBytimeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Unit of time to group data by.
type DNSAnalyticsReportBytimeGetParamsTimeDelta string

const (
	DNSAnalyticsReportBytimeGetParamsTimeDeltaAll        DNSAnalyticsReportBytimeGetParamsTimeDelta = "all"
	DNSAnalyticsReportBytimeGetParamsTimeDeltaAuto       DNSAnalyticsReportBytimeGetParamsTimeDelta = "auto"
	DNSAnalyticsReportBytimeGetParamsTimeDeltaYear       DNSAnalyticsReportBytimeGetParamsTimeDelta = "year"
	DNSAnalyticsReportBytimeGetParamsTimeDeltaQuarter    DNSAnalyticsReportBytimeGetParamsTimeDelta = "quarter"
	DNSAnalyticsReportBytimeGetParamsTimeDeltaMonth      DNSAnalyticsReportBytimeGetParamsTimeDelta = "month"
	DNSAnalyticsReportBytimeGetParamsTimeDeltaWeek       DNSAnalyticsReportBytimeGetParamsTimeDelta = "week"
	DNSAnalyticsReportBytimeGetParamsTimeDeltaDay        DNSAnalyticsReportBytimeGetParamsTimeDelta = "day"
	DNSAnalyticsReportBytimeGetParamsTimeDeltaHour       DNSAnalyticsReportBytimeGetParamsTimeDelta = "hour"
	DNSAnalyticsReportBytimeGetParamsTimeDeltaDekaminute DNSAnalyticsReportBytimeGetParamsTimeDelta = "dekaminute"
	DNSAnalyticsReportBytimeGetParamsTimeDeltaMinute     DNSAnalyticsReportBytimeGetParamsTimeDelta = "minute"
)

type DNSAnalyticsReportBytimeGetResponseEnvelope struct {
	Errors   []DNSAnalyticsReportBytimeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSAnalyticsReportBytimeGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSAnalyticsReportBytimeGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSAnalyticsReportBytimeGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsAnalyticsReportBytimeGetResponseEnvelopeJSON    `json:"-"`
}

// dnsAnalyticsReportBytimeGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DNSAnalyticsReportBytimeGetResponseEnvelope]
type dnsAnalyticsReportBytimeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticsReportBytimeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsAnalyticsReportBytimeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DNSAnalyticsReportBytimeGetResponseEnvelopeErrors struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    dnsAnalyticsReportBytimeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsAnalyticsReportBytimeGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DNSAnalyticsReportBytimeGetResponseEnvelopeErrors]
type dnsAnalyticsReportBytimeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticsReportBytimeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsAnalyticsReportBytimeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DNSAnalyticsReportBytimeGetResponseEnvelopeMessages struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    dnsAnalyticsReportBytimeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsAnalyticsReportBytimeGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DNSAnalyticsReportBytimeGetResponseEnvelopeMessages]
type dnsAnalyticsReportBytimeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticsReportBytimeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsAnalyticsReportBytimeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DNSAnalyticsReportBytimeGetResponseEnvelopeSuccess bool

const (
	DNSAnalyticsReportBytimeGetResponseEnvelopeSuccessTrue DNSAnalyticsReportBytimeGetResponseEnvelopeSuccess = true
)
