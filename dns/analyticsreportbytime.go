// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

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
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AnalyticsReportBytimeService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAnalyticsReportBytimeService]
// method instead.
type AnalyticsReportBytimeService struct {
	Options []option.RequestOption
}

// NewAnalyticsReportBytimeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsReportBytimeService(opts ...option.RequestOption) (r *AnalyticsReportBytimeService) {
	r = &AnalyticsReportBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *AnalyticsReportBytimeService) Get(ctx context.Context, params AnalyticsReportBytimeGetParams, opts ...option.RequestOption) (res *DNSAnalyticsReportByTime, err error) {
	opts = append(r.Options[:], opts...)
	var env AnalyticsReportBytimeGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_analytics/report/bytime", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSAnalyticsReportByTime struct {
	// Array with one row per combination of dimension values.
	Data []DNSAnalyticsReportByTimeData `json:"data,required"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}               `json:"min,required"`
	Query shared.UnnamedSchemaRef24 `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Array of time intervals in the response data. Each interval is represented as an
	// array containing two values: the start time, and the end time.
	TimeIntervals [][]time.Time `json:"time_intervals,required" format:"date-time"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                  `json:"totals,required"`
	JSON   dnsAnalyticsReportByTimeJSON `json:"-"`
}

// dnsAnalyticsReportByTimeJSON contains the JSON metadata for the struct
// [DNSAnalyticsReportByTime]
type dnsAnalyticsReportByTimeJSON struct {
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

func (r *DNSAnalyticsReportByTime) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsAnalyticsReportByTimeJSON) RawJSON() string {
	return r.raw
}

type DNSAnalyticsReportByTimeData struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is an array of values,
	// broken down by time interval.
	Metrics []shared.UnnamedSchemaRef23      `json:"metrics,required"`
	JSON    dnsAnalyticsReportByTimeDataJSON `json:"-"`
}

// dnsAnalyticsReportByTimeDataJSON contains the JSON metadata for the struct
// [DNSAnalyticsReportByTimeData]
type dnsAnalyticsReportByTimeDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticsReportByTimeData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsAnalyticsReportByTimeDataJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportBytimeGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
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
	TimeDelta param.Field[AnalyticsReportBytimeGetParamsTimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AnalyticsReportBytimeGetParams]'s query parameters as
// `url.Values`.
func (r AnalyticsReportBytimeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Unit of time to group data by.
type AnalyticsReportBytimeGetParamsTimeDelta string

const (
	AnalyticsReportBytimeGetParamsTimeDeltaAll        AnalyticsReportBytimeGetParamsTimeDelta = "all"
	AnalyticsReportBytimeGetParamsTimeDeltaAuto       AnalyticsReportBytimeGetParamsTimeDelta = "auto"
	AnalyticsReportBytimeGetParamsTimeDeltaYear       AnalyticsReportBytimeGetParamsTimeDelta = "year"
	AnalyticsReportBytimeGetParamsTimeDeltaQuarter    AnalyticsReportBytimeGetParamsTimeDelta = "quarter"
	AnalyticsReportBytimeGetParamsTimeDeltaMonth      AnalyticsReportBytimeGetParamsTimeDelta = "month"
	AnalyticsReportBytimeGetParamsTimeDeltaWeek       AnalyticsReportBytimeGetParamsTimeDelta = "week"
	AnalyticsReportBytimeGetParamsTimeDeltaDay        AnalyticsReportBytimeGetParamsTimeDelta = "day"
	AnalyticsReportBytimeGetParamsTimeDeltaHour       AnalyticsReportBytimeGetParamsTimeDelta = "hour"
	AnalyticsReportBytimeGetParamsTimeDeltaDekaminute AnalyticsReportBytimeGetParamsTimeDelta = "dekaminute"
	AnalyticsReportBytimeGetParamsTimeDeltaMinute     AnalyticsReportBytimeGetParamsTimeDelta = "minute"
)

func (r AnalyticsReportBytimeGetParamsTimeDelta) IsKnown() bool {
	switch r {
	case AnalyticsReportBytimeGetParamsTimeDeltaAll, AnalyticsReportBytimeGetParamsTimeDeltaAuto, AnalyticsReportBytimeGetParamsTimeDeltaYear, AnalyticsReportBytimeGetParamsTimeDeltaQuarter, AnalyticsReportBytimeGetParamsTimeDeltaMonth, AnalyticsReportBytimeGetParamsTimeDeltaWeek, AnalyticsReportBytimeGetParamsTimeDeltaDay, AnalyticsReportBytimeGetParamsTimeDeltaHour, AnalyticsReportBytimeGetParamsTimeDeltaDekaminute, AnalyticsReportBytimeGetParamsTimeDeltaMinute:
		return true
	}
	return false
}

type AnalyticsReportBytimeGetResponseEnvelope struct {
	Errors   []shared.UnnamedSchemaRef172 `json:"errors,required"`
	Messages []shared.UnnamedSchemaRef172 `json:"messages,required"`
	Result   DNSAnalyticsReportByTime     `json:"result,required"`
	// Whether the API call was successful
	Success AnalyticsReportBytimeGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    analyticsReportBytimeGetResponseEnvelopeJSON    `json:"-"`
}

// analyticsReportBytimeGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnalyticsReportBytimeGetResponseEnvelope]
type analyticsReportBytimeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsReportBytimeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportBytimeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AnalyticsReportBytimeGetResponseEnvelopeSuccess bool

const (
	AnalyticsReportBytimeGetResponseEnvelopeSuccessTrue AnalyticsReportBytimeGetResponseEnvelopeSuccess = true
)

func (r AnalyticsReportBytimeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AnalyticsReportBytimeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
