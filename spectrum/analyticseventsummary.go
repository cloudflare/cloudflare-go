// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package spectrum

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// AnalyticsEventSummaryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAnalyticsEventSummaryService]
// method instead.
type AnalyticsEventSummaryService struct {
	Options []option.RequestOption
}

// NewAnalyticsEventSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAnalyticsEventSummaryService(opts ...option.RequestOption) (r *AnalyticsEventSummaryService) {
	r = &AnalyticsEventSummaryService{}
	r.Options = opts
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
func (r *AnalyticsEventSummaryService) Get(ctx context.Context, zone string, query AnalyticsEventSummaryGetParams, opts ...option.RequestOption) (res *AnalyticsEventSummaryGetResponseUnion, err error) {
	opts = append(r.Options[:], opts...)
	var env AnalyticsEventSummaryGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/analytics/events/summary", zone)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [spectrum.AnalyticsEventSummaryGetResponseUnknown] or
// [shared.UnionString].
type AnalyticsEventSummaryGetResponseUnion interface {
	ImplementsSpectrumAnalyticsEventSummaryGetResponseUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AnalyticsEventSummaryGetResponseUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AnalyticsEventSummaryGetParams struct {
	// Can be used to break down the data by given attributes. Options are:
	//
	// | Dimension | Name                          | Example                                                    |
	// | --------- | ----------------------------- | ---------------------------------------------------------- |
	// | event     | Connection Event              | connect, progress, disconnect, originError, clientFiltered |
	// | appID     | Application ID                | 40d67c87c6cd4b889a4fd57805225e85                           |
	// | coloName  | Colo Name                     | SFO                                                        |
	// | ipVersion | IP version used by the client | 4, 6.                                                      |
	Dimensions param.Field[[]Dimension] `query:"dimensions"`
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
	// One or more metrics to compute. Options are:
	//
	// | Metric         | Name                                | Example | Unit                  |
	// | -------------- | ----------------------------------- | ------- | --------------------- |
	// | count          | Count of total events               | 1000    | Count                 |
	// | bytesIngress   | Sum of ingress bytes                | 1000    | Sum                   |
	// | bytesEgress    | Sum of egress bytes                 | 1000    | Sum                   |
	// | durationAvg    | Average connection duration         | 1.0     | Time in milliseconds  |
	// | durationMedian | Median connection duration          | 1.0     | Time in milliseconds  |
	// | duration90th   | 90th percentile connection duration | 1.0     | Time in milliseconds  |
	// | duration99th   | 99th percentile connection duration | 1.0     | Time in milliseconds. |
	Metrics param.Field[[]AnalyticsEventSummaryGetParamsMetric] `query:"metrics"`
	// Start of time interval to query, defaults to `until` - 6 hours. Timestamp must
	// be in RFC3339 format and uses UTC unless otherwise specified.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// The sort order for the result set; sort fields must be included in `metrics` or
	// `dimensions`.
	Sort param.Field[[]interface{}] `query:"sort"`
	// End of time interval to query, defaults to current time. Timestamp must be in
	// RFC3339 format and uses UTC unless otherwise specified.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AnalyticsEventSummaryGetParams]'s query parameters as
// `url.Values`.
func (r AnalyticsEventSummaryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticsEventSummaryGetParamsMetric string

const (
	AnalyticsEventSummaryGetParamsMetricCount          AnalyticsEventSummaryGetParamsMetric = "count"
	AnalyticsEventSummaryGetParamsMetricBytesIngress   AnalyticsEventSummaryGetParamsMetric = "bytesIngress"
	AnalyticsEventSummaryGetParamsMetricBytesEgress    AnalyticsEventSummaryGetParamsMetric = "bytesEgress"
	AnalyticsEventSummaryGetParamsMetricDurationAvg    AnalyticsEventSummaryGetParamsMetric = "durationAvg"
	AnalyticsEventSummaryGetParamsMetricDurationMedian AnalyticsEventSummaryGetParamsMetric = "durationMedian"
	AnalyticsEventSummaryGetParamsMetricDuration90th   AnalyticsEventSummaryGetParamsMetric = "duration90th"
	AnalyticsEventSummaryGetParamsMetricDuration99th   AnalyticsEventSummaryGetParamsMetric = "duration99th"
)

func (r AnalyticsEventSummaryGetParamsMetric) IsKnown() bool {
	switch r {
	case AnalyticsEventSummaryGetParamsMetricCount, AnalyticsEventSummaryGetParamsMetricBytesIngress, AnalyticsEventSummaryGetParamsMetricBytesEgress, AnalyticsEventSummaryGetParamsMetricDurationAvg, AnalyticsEventSummaryGetParamsMetricDurationMedian, AnalyticsEventSummaryGetParamsMetricDuration90th, AnalyticsEventSummaryGetParamsMetricDuration99th:
		return true
	}
	return false
}

type AnalyticsEventSummaryGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo                 `json:"errors,required"`
	Messages []shared.ResponseInfo                 `json:"messages,required"`
	Result   AnalyticsEventSummaryGetResponseUnion `json:"result,required,nullable"`
	// Whether the API call was successful
	Success AnalyticsEventSummaryGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    analyticsEventSummaryGetResponseEnvelopeJSON    `json:"-"`
}

// analyticsEventSummaryGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AnalyticsEventSummaryGetResponseEnvelope]
type analyticsEventSummaryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsEventSummaryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsEventSummaryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AnalyticsEventSummaryGetResponseEnvelopeSuccess bool

const (
	AnalyticsEventSummaryGetResponseEnvelopeSuccessTrue AnalyticsEventSummaryGetResponseEnvelopeSuccess = true
)

func (r AnalyticsEventSummaryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AnalyticsEventSummaryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
