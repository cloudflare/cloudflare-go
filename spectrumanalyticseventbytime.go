// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// SpectrumAnalyticsEventBytimeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewSpectrumAnalyticsEventBytimeService] method instead.
type SpectrumAnalyticsEventBytimeService struct {
	Options []option.RequestOption
}

// NewSpectrumAnalyticsEventBytimeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSpectrumAnalyticsEventBytimeService(opts ...option.RequestOption) (r *SpectrumAnalyticsEventBytimeService) {
	r = &SpectrumAnalyticsEventBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
func (r *SpectrumAnalyticsEventBytimeService) SpectrumAnalyticsByTimeGetAnalyticsByTime(ctx context.Context, zone string, query SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParams, opts ...option.RequestOption) (res *SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelope
	path := fmt.Sprintf("zones/%s/spectrum/analytics/events/bytime", zone)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by
// [SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseUnknown]
// or [shared.UnionString].
type SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponse interface {
	ImplementsSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParams struct {
	// Can be used to break down the data by given attributes. Options are:
	//
	// | Dimension | Name                          | Example                                                    |
	// | --------- | ----------------------------- | ---------------------------------------------------------- |
	// | event     | Connection Event              | connect, progress, disconnect, originError, clientFiltered |
	// | appID     | Application ID                | 40d67c87c6cd4b889a4fd57805225e85                           |
	// | coloName  | Colo Name                     | SFO                                                        |
	// | ipVersion | IP version used by the client | 4, 6.                                                      |
	Dimensions param.Field[[]SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension] `query:"dimensions"`
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
	Metrics param.Field[[]SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric] `query:"metrics"`
	// Start of time interval to query, defaults to `until` - 6 hours. Timestamp must
	// be in RFC3339 format and uses UTC unless otherwise specified.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// The sort order for the result set; sort fields must be included in `metrics` or
	// `dimensions`.
	Sort param.Field[[]interface{}] `query:"sort"`
	// Used to select time series resolution.
	TimeDelta param.Field[SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta] `query:"time_delta"`
	// End of time interval to query, defaults to current time. Timestamp must be in
	// RFC3339 format and uses UTC unless otherwise specified.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes
// [SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParams]'s
// query parameters as `url.Values`.
func (r SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension string

const (
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimensionEvent     SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension = "event"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimensionAppID     SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension = "appID"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimensionColoName  SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension = "coloName"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimensionIPVersion SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension = "ipVersion"
)

type SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric string

const (
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricCount          SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "count"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricBytesIngress   SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "bytesIngress"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricBytesEgress    SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "bytesEgress"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricDurationAvg    SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "durationAvg"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricDurationMedian SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "durationMedian"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricDuration90th   SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "duration90th"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricDuration99th   SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "duration99th"
)

// Used to select time series resolution.
type SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta string

const (
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaYear       SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "year"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaQuarter    SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "quarter"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaMonth      SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "month"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaWeek       SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "week"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaDay        SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "day"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaHour       SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "hour"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaDekaminute SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "dekaminute"
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaMinute     SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "minute"
)

type SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelope struct {
	Errors   []SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeMessages `json:"messages,required"`
	Result   SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeSuccess `json:"success,required"`
	JSON    spectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeJSON    `json:"-"`
}

// spectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelope]
type spectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeErrors struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    spectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeErrorsJSON `json:"-"`
}

// spectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeErrors]
type spectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeMessages struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    spectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeMessagesJSON `json:"-"`
}

// spectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeMessages]
type spectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeSuccess bool

const (
	SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeSuccessTrue SpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeResponseEnvelopeSuccess = true
)
