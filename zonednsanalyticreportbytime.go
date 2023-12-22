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

// ZoneDNSAnalyticReportBytimeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneDNSAnalyticReportBytimeService] method instead.
type ZoneDNSAnalyticReportBytimeService struct {
	Options []option.RequestOption
}

// NewZoneDNSAnalyticReportBytimeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneDNSAnalyticReportBytimeService(opts ...option.RequestOption) (r *ZoneDNSAnalyticReportBytimeService) {
	r = &ZoneDNSAnalyticReportBytimeService{}
	r.Options = opts
	return
}

// Retrieves a list of aggregate metrics grouped by time interval.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *ZoneDNSAnalyticReportBytimeService) List(ctx context.Context, identifier string, query ZoneDNSAnalyticReportBytimeListParams, opts ...option.RequestOption) (res *ZoneDNSAnalyticReportBytimeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_analytics/report/bytime", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneDNSAnalyticReportBytimeListResponse struct {
	Errors   []ZoneDNSAnalyticReportBytimeListResponseError   `json:"errors"`
	Messages []ZoneDNSAnalyticReportBytimeListResponseMessage `json:"messages"`
	Result   ZoneDNSAnalyticReportBytimeListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneDNSAnalyticReportBytimeListResponseSuccess `json:"success"`
	JSON    zoneDNSAnalyticReportBytimeListResponseJSON    `json:"-"`
}

// zoneDNSAnalyticReportBytimeListResponseJSON contains the JSON metadata for the
// struct [ZoneDNSAnalyticReportBytimeListResponse]
type zoneDNSAnalyticReportBytimeListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportBytimeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSAnalyticReportBytimeListResponseError struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zoneDNSAnalyticReportBytimeListResponseErrorJSON `json:"-"`
}

// zoneDNSAnalyticReportBytimeListResponseErrorJSON contains the JSON metadata for
// the struct [ZoneDNSAnalyticReportBytimeListResponseError]
type zoneDNSAnalyticReportBytimeListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportBytimeListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSAnalyticReportBytimeListResponseMessage struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zoneDNSAnalyticReportBytimeListResponseMessageJSON `json:"-"`
}

// zoneDNSAnalyticReportBytimeListResponseMessageJSON contains the JSON metadata
// for the struct [ZoneDNSAnalyticReportBytimeListResponseMessage]
type zoneDNSAnalyticReportBytimeListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportBytimeListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSAnalyticReportBytimeListResponseResult struct {
	Data []ZoneDNSAnalyticReportBytimeListResponseResultData `json:"data"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}                                        `json:"min"`
	Query ZoneDNSAnalyticReportBytimeListResponseResultQuery `json:"query"`
	// Total number of rows in the result.
	Rows float64 `json:"rows"`
	// Array of time intervals in the response data. Each interval is represented as an
	// array containing two values: the start time, and the end time.
	TimeIntervals [][]time.Time `json:"time_intervals" format:"date-time"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                                       `json:"totals"`
	JSON   zoneDNSAnalyticReportBytimeListResponseResultJSON `json:"-"`
}

// zoneDNSAnalyticReportBytimeListResponseResultJSON contains the JSON metadata for
// the struct [ZoneDNSAnalyticReportBytimeListResponseResult]
type zoneDNSAnalyticReportBytimeListResponseResultJSON struct {
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

func (r *ZoneDNSAnalyticReportBytimeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSAnalyticReportBytimeListResponseResultData struct {
	// Array with one item per requested metric. Each item is an array of values,
	// broken down by time interval.
	Metrics [][]interface{}                                       `json:"metrics,required"`
	JSON    zoneDNSAnalyticReportBytimeListResponseResultDataJSON `json:"-"`
}

// zoneDNSAnalyticReportBytimeListResponseResultDataJSON contains the JSON metadata
// for the struct [ZoneDNSAnalyticReportBytimeListResponseResultData]
type zoneDNSAnalyticReportBytimeListResponseResultDataJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportBytimeListResponseResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSAnalyticReportBytimeListResponseResultQuery struct {
	// Unit of time to group data by.
	TimeDelta ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDelta `json:"time_delta,required"`
	JSON      zoneDNSAnalyticReportBytimeListResponseResultQueryJSON      `json:"-"`
}

// zoneDNSAnalyticReportBytimeListResponseResultQueryJSON contains the JSON
// metadata for the struct [ZoneDNSAnalyticReportBytimeListResponseResultQuery]
type zoneDNSAnalyticReportBytimeListResponseResultQueryJSON struct {
	TimeDelta   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportBytimeListResponseResultQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Unit of time to group data by.
type ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDelta string

const (
	ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaAll        ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "all"
	ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaAuto       ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "auto"
	ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaYear       ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "year"
	ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaQuarter    ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "quarter"
	ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaMonth      ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "month"
	ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaWeek       ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "week"
	ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaDay        ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "day"
	ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaHour       ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "hour"
	ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaDekaminute ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "dekaminute"
	ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDeltaMinute     ZoneDNSAnalyticReportBytimeListResponseResultQueryTimeDelta = "minute"
)

// Whether the API call was successful
type ZoneDNSAnalyticReportBytimeListResponseSuccess bool

const (
	ZoneDNSAnalyticReportBytimeListResponseSuccessTrue ZoneDNSAnalyticReportBytimeListResponseSuccess = true
)

type ZoneDNSAnalyticReportBytimeListParams struct {
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
	TimeDelta param.Field[ZoneDNSAnalyticReportBytimeListParamsTimeDelta] `query:"time_delta"`
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [ZoneDNSAnalyticReportBytimeListParams]'s query parameters
// as `url.Values`.
func (r ZoneDNSAnalyticReportBytimeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Unit of time to group data by.
type ZoneDNSAnalyticReportBytimeListParamsTimeDelta string

const (
	ZoneDNSAnalyticReportBytimeListParamsTimeDeltaAll        ZoneDNSAnalyticReportBytimeListParamsTimeDelta = "all"
	ZoneDNSAnalyticReportBytimeListParamsTimeDeltaAuto       ZoneDNSAnalyticReportBytimeListParamsTimeDelta = "auto"
	ZoneDNSAnalyticReportBytimeListParamsTimeDeltaYear       ZoneDNSAnalyticReportBytimeListParamsTimeDelta = "year"
	ZoneDNSAnalyticReportBytimeListParamsTimeDeltaQuarter    ZoneDNSAnalyticReportBytimeListParamsTimeDelta = "quarter"
	ZoneDNSAnalyticReportBytimeListParamsTimeDeltaMonth      ZoneDNSAnalyticReportBytimeListParamsTimeDelta = "month"
	ZoneDNSAnalyticReportBytimeListParamsTimeDeltaWeek       ZoneDNSAnalyticReportBytimeListParamsTimeDelta = "week"
	ZoneDNSAnalyticReportBytimeListParamsTimeDeltaDay        ZoneDNSAnalyticReportBytimeListParamsTimeDelta = "day"
	ZoneDNSAnalyticReportBytimeListParamsTimeDeltaHour       ZoneDNSAnalyticReportBytimeListParamsTimeDelta = "hour"
	ZoneDNSAnalyticReportBytimeListParamsTimeDeltaDekaminute ZoneDNSAnalyticReportBytimeListParamsTimeDelta = "dekaminute"
	ZoneDNSAnalyticReportBytimeListParamsTimeDeltaMinute     ZoneDNSAnalyticReportBytimeListParamsTimeDelta = "minute"
)
