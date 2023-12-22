// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSpectrumAnalyticsEventBytimeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneSpectrumAnalyticsEventBytimeService] method instead.
type ZoneSpectrumAnalyticsEventBytimeService struct {
	Options []option.RequestOption
}

// NewZoneSpectrumAnalyticsEventBytimeService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSpectrumAnalyticsEventBytimeService(opts ...option.RequestOption) (r *ZoneSpectrumAnalyticsEventBytimeService) {
	r = &ZoneSpectrumAnalyticsEventBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
func (r *ZoneSpectrumAnalyticsEventBytimeService) SpectrumAnalyticsByTimeGetAnalyticsByTime(ctx context.Context, zone string, query ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParams, opts ...option.RequestOption) (res *APIResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/spectrum/analytics/events/bytime", zone)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParams struct {
	// Can be used to break down the data by given attributes. Options are:
	//
	// | Dimension | Name                          | Example                                                    |
	// | --------- | ----------------------------- | ---------------------------------------------------------- |
	// | event     | Connection Event              | connect, progress, disconnect, originError, clientFiltered |
	// | appID     | Application ID                | 40d67c87c6cd4b889a4fd57805225e85                           |
	// | coloName  | Colo Name                     | SFO                                                        |
	// | ipVersion | IP version used by the client | 4, 6.                                                      |
	Dimensions param.Field[[]ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension] `query:"dimensions"`
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
	Metrics param.Field[[]ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric] `query:"metrics"`
	// Start of time interval to query, defaults to `until` - 6 hours. Timestamp must
	// be in RFC3339 format and uses UTC unless otherwise specified.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// The sort order for the result set; sort fields must be included in `metrics` or
	// `dimensions`.
	Sort param.Field[[]interface{}] `query:"sort"`
	// Used to select time series resolution.
	TimeDelta param.Field[ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta] `query:"time_delta"`
	// End of time interval to query, defaults to current time. Timestamp must be in
	// RFC3339 format and uses UTC unless otherwise specified.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes
// [ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParams]'s
// query parameters as `url.Values`.
func (r ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension string

const (
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimensionEvent     ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension = "event"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimensionAppID     ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension = "appID"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimensionColoName  ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension = "coloName"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimensionIPVersion ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsDimension = "ipVersion"
)

type ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric string

const (
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricCount          ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "count"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricBytesIngress   ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "bytesIngress"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricBytesEgress    ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "bytesEgress"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricDurationAvg    ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "durationAvg"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricDurationMedian ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "durationMedian"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricDuration90th   ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "duration90th"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetricDuration99th   ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsMetric = "duration99th"
)

// Used to select time series resolution.
type ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta string

const (
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaYear       ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "year"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaQuarter    ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "quarter"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaMonth      ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "month"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaWeek       ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "week"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaDay        ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "day"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaHour       ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "hour"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaDekaminute ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "dekaminute"
	ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDeltaMinute     ZoneSpectrumAnalyticsEventBytimeSpectrumAnalyticsByTimeGetAnalyticsByTimeParamsTimeDelta = "minute"
)
