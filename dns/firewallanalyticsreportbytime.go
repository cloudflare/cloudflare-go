// File generated from our OpenAPI spec by Stainless.

package dns

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// FirewallAnalyticsReportBytimeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewFirewallAnalyticsReportBytimeService] method instead.
type FirewallAnalyticsReportBytimeService struct {
	Options []option.RequestOption
}

// NewFirewallAnalyticsReportBytimeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewFirewallAnalyticsReportBytimeService(opts ...option.RequestOption) (r *FirewallAnalyticsReportBytimeService) {
	r = &FirewallAnalyticsReportBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *FirewallAnalyticsReportBytimeService) Get(ctx context.Context, accountIdentifier string, identifier string, query FirewallAnalyticsReportBytimeGetParams, opts ...option.RequestOption) (res *FirewallAnalyticsReportBytimeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FirewallAnalyticsReportBytimeGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dns_firewall/%s/dns_analytics/report/bytime", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type FirewallAnalyticsReportBytimeGetResponse struct {
	// Array with one row per combination of dimension values.
	Data []FirewallAnalyticsReportBytimeGetResponseData `json:"data,required"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}                                   `json:"min,required"`
	Query FirewallAnalyticsReportBytimeGetResponseQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Array of time intervals in the response data. Each interval is represented as an
	// array containing two values: the start time, and the end time.
	TimeIntervals [][]time.Time `json:"time_intervals,required" format:"date-time"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                                  `json:"totals,required"`
	JSON   firewallAnalyticsReportBytimeGetResponseJSON `json:"-"`
}

// firewallAnalyticsReportBytimeGetResponseJSON contains the JSON metadata for the
// struct [FirewallAnalyticsReportBytimeGetResponse]
type firewallAnalyticsReportBytimeGetResponseJSON struct {
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

func (r *FirewallAnalyticsReportBytimeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAnalyticsReportBytimeGetResponseJSON) RawJSON() string {
	return r.raw
}

type FirewallAnalyticsReportBytimeGetResponseData struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is an array of values,
	// broken down by time interval.
	Metrics [][]interface{}                                  `json:"metrics,required"`
	JSON    firewallAnalyticsReportBytimeGetResponseDataJSON `json:"-"`
}

// firewallAnalyticsReportBytimeGetResponseDataJSON contains the JSON metadata for
// the struct [FirewallAnalyticsReportBytimeGetResponseData]
type firewallAnalyticsReportBytimeGetResponseDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAnalyticsReportBytimeGetResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAnalyticsReportBytimeGetResponseDataJSON) RawJSON() string {
	return r.raw
}

type FirewallAnalyticsReportBytimeGetResponseQuery struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// Unit of time to group data by.
	TimeDelta FirewallAnalyticsReportBytimeGetResponseQueryTimeDelta `json:"time_delta,required"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string                                          `json:"sort"`
	JSON firewallAnalyticsReportBytimeGetResponseQueryJSON `json:"-"`
}

// firewallAnalyticsReportBytimeGetResponseQueryJSON contains the JSON metadata for
// the struct [FirewallAnalyticsReportBytimeGetResponseQuery]
type firewallAnalyticsReportBytimeGetResponseQueryJSON struct {
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

func (r *FirewallAnalyticsReportBytimeGetResponseQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAnalyticsReportBytimeGetResponseQueryJSON) RawJSON() string {
	return r.raw
}

// Unit of time to group data by.
type FirewallAnalyticsReportBytimeGetResponseQueryTimeDelta string

const (
	FirewallAnalyticsReportBytimeGetResponseQueryTimeDeltaAll        FirewallAnalyticsReportBytimeGetResponseQueryTimeDelta = "all"
	FirewallAnalyticsReportBytimeGetResponseQueryTimeDeltaAuto       FirewallAnalyticsReportBytimeGetResponseQueryTimeDelta = "auto"
	FirewallAnalyticsReportBytimeGetResponseQueryTimeDeltaYear       FirewallAnalyticsReportBytimeGetResponseQueryTimeDelta = "year"
	FirewallAnalyticsReportBytimeGetResponseQueryTimeDeltaQuarter    FirewallAnalyticsReportBytimeGetResponseQueryTimeDelta = "quarter"
	FirewallAnalyticsReportBytimeGetResponseQueryTimeDeltaMonth      FirewallAnalyticsReportBytimeGetResponseQueryTimeDelta = "month"
	FirewallAnalyticsReportBytimeGetResponseQueryTimeDeltaWeek       FirewallAnalyticsReportBytimeGetResponseQueryTimeDelta = "week"
	FirewallAnalyticsReportBytimeGetResponseQueryTimeDeltaDay        FirewallAnalyticsReportBytimeGetResponseQueryTimeDelta = "day"
	FirewallAnalyticsReportBytimeGetResponseQueryTimeDeltaHour       FirewallAnalyticsReportBytimeGetResponseQueryTimeDelta = "hour"
	FirewallAnalyticsReportBytimeGetResponseQueryTimeDeltaDekaminute FirewallAnalyticsReportBytimeGetResponseQueryTimeDelta = "dekaminute"
	FirewallAnalyticsReportBytimeGetResponseQueryTimeDeltaMinute     FirewallAnalyticsReportBytimeGetResponseQueryTimeDelta = "minute"
)

type FirewallAnalyticsReportBytimeGetParams struct {
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
	TimeDelta param.Field[FirewallAnalyticsReportBytimeGetParamsTimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [FirewallAnalyticsReportBytimeGetParams]'s query parameters
// as `url.Values`.
func (r FirewallAnalyticsReportBytimeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Unit of time to group data by.
type FirewallAnalyticsReportBytimeGetParamsTimeDelta string

const (
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaAll        FirewallAnalyticsReportBytimeGetParamsTimeDelta = "all"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaAuto       FirewallAnalyticsReportBytimeGetParamsTimeDelta = "auto"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaYear       FirewallAnalyticsReportBytimeGetParamsTimeDelta = "year"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaQuarter    FirewallAnalyticsReportBytimeGetParamsTimeDelta = "quarter"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaMonth      FirewallAnalyticsReportBytimeGetParamsTimeDelta = "month"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaWeek       FirewallAnalyticsReportBytimeGetParamsTimeDelta = "week"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaDay        FirewallAnalyticsReportBytimeGetParamsTimeDelta = "day"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaHour       FirewallAnalyticsReportBytimeGetParamsTimeDelta = "hour"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaDekaminute FirewallAnalyticsReportBytimeGetParamsTimeDelta = "dekaminute"
	FirewallAnalyticsReportBytimeGetParamsTimeDeltaMinute     FirewallAnalyticsReportBytimeGetParamsTimeDelta = "minute"
)

type FirewallAnalyticsReportBytimeGetResponseEnvelope struct {
	Errors   []FirewallAnalyticsReportBytimeGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FirewallAnalyticsReportBytimeGetResponseEnvelopeMessages `json:"messages,required"`
	Result   FirewallAnalyticsReportBytimeGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success FirewallAnalyticsReportBytimeGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    firewallAnalyticsReportBytimeGetResponseEnvelopeJSON    `json:"-"`
}

// firewallAnalyticsReportBytimeGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [FirewallAnalyticsReportBytimeGetResponseEnvelope]
type firewallAnalyticsReportBytimeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAnalyticsReportBytimeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAnalyticsReportBytimeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FirewallAnalyticsReportBytimeGetResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    firewallAnalyticsReportBytimeGetResponseEnvelopeErrorsJSON `json:"-"`
}

// firewallAnalyticsReportBytimeGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [FirewallAnalyticsReportBytimeGetResponseEnvelopeErrors]
type firewallAnalyticsReportBytimeGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAnalyticsReportBytimeGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAnalyticsReportBytimeGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FirewallAnalyticsReportBytimeGetResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    firewallAnalyticsReportBytimeGetResponseEnvelopeMessagesJSON `json:"-"`
}

// firewallAnalyticsReportBytimeGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [FirewallAnalyticsReportBytimeGetResponseEnvelopeMessages]
type firewallAnalyticsReportBytimeGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FirewallAnalyticsReportBytimeGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r firewallAnalyticsReportBytimeGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FirewallAnalyticsReportBytimeGetResponseEnvelopeSuccess bool

const (
	FirewallAnalyticsReportBytimeGetResponseEnvelopeSuccessTrue FirewallAnalyticsReportBytimeGetResponseEnvelopeSuccess = true
)
