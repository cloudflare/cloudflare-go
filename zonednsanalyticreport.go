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

// ZoneDNSAnalyticReportService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneDNSAnalyticReportService]
// method instead.
type ZoneDNSAnalyticReportService struct {
	Options []option.RequestOption
	Bytimes *ZoneDNSAnalyticReportBytimeService
}

// NewZoneDNSAnalyticReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneDNSAnalyticReportService(opts ...option.RequestOption) (r *ZoneDNSAnalyticReportService) {
	r = &ZoneDNSAnalyticReportService{}
	r.Options = opts
	r.Bytimes = NewZoneDNSAnalyticReportBytimeService(opts...)
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *ZoneDNSAnalyticReportService) List(ctx context.Context, identifier string, query ZoneDNSAnalyticReportListParams, opts ...option.RequestOption) (res *ZoneDNSAnalyticReportListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/dns_analytics/report", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneDNSAnalyticReportListResponse struct {
	Errors   []ZoneDNSAnalyticReportListResponseError   `json:"errors"`
	Messages []ZoneDNSAnalyticReportListResponseMessage `json:"messages"`
	Result   ZoneDNSAnalyticReportListResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneDNSAnalyticReportListResponseSuccess `json:"success"`
	JSON    zoneDNSAnalyticReportListResponseJSON    `json:"-"`
}

// zoneDNSAnalyticReportListResponseJSON contains the JSON metadata for the struct
// [ZoneDNSAnalyticReportListResponse]
type zoneDNSAnalyticReportListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSAnalyticReportListResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneDNSAnalyticReportListResponseErrorJSON `json:"-"`
}

// zoneDNSAnalyticReportListResponseErrorJSON contains the JSON metadata for the
// struct [ZoneDNSAnalyticReportListResponseError]
type zoneDNSAnalyticReportListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSAnalyticReportListResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneDNSAnalyticReportListResponseMessageJSON `json:"-"`
}

// zoneDNSAnalyticReportListResponseMessageJSON contains the JSON metadata for the
// struct [ZoneDNSAnalyticReportListResponseMessage]
type zoneDNSAnalyticReportListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSAnalyticReportListResponseResult struct {
	Data []ZoneDNSAnalyticReportListResponseResultData `json:"data"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}                                  `json:"min"`
	Query ZoneDNSAnalyticReportListResponseResultQuery `json:"query"`
	// Total number of rows in the result.
	Rows float64 `json:"rows"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                                 `json:"totals"`
	JSON   zoneDNSAnalyticReportListResponseResultJSON `json:"-"`
}

// zoneDNSAnalyticReportListResponseResultJSON contains the JSON metadata for the
// struct [ZoneDNSAnalyticReportListResponseResult]
type zoneDNSAnalyticReportListResponseResultJSON struct {
	Data        apijson.Field
	DataLag     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	Query       apijson.Field
	Rows        apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSAnalyticReportListResponseResultData struct {
	// Array with one item per requested metric. Each item is a single value.
	Metrics []float64                                       `json:"metrics,required"`
	JSON    zoneDNSAnalyticReportListResponseResultDataJSON `json:"-"`
}

// zoneDNSAnalyticReportListResponseResultDataJSON contains the JSON metadata for
// the struct [ZoneDNSAnalyticReportListResponseResultData]
type zoneDNSAnalyticReportListResponseResultDataJSON struct {
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportListResponseResultData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneDNSAnalyticReportListResponseResultQuery struct {
	// Array of dimension names.
	Dimensions []string `json:"dimensions,required"`
	// Limit number of returned metrics.
	Limit int64 `json:"limit,required"`
	// Array of metric names.
	Metrics []string `json:"metrics,required"`
	// Start date and time of requesting data period in ISO 8601 format.
	Since time.Time `json:"since,required" format:"date-time"`
	// End date and time of requesting data period in ISO 8601 format.
	Until time.Time `json:"until,required" format:"date-time"`
	// Segmentation filter in 'attribute operator value' format.
	Filters string `json:"filters"`
	// Array of dimensions to sort by, where each dimension may be prefixed by -
	// (descending) or + (ascending).
	Sort []string                                         `json:"sort"`
	JSON zoneDNSAnalyticReportListResponseResultQueryJSON `json:"-"`
}

// zoneDNSAnalyticReportListResponseResultQueryJSON contains the JSON metadata for
// the struct [ZoneDNSAnalyticReportListResponseResultQuery]
type zoneDNSAnalyticReportListResponseResultQueryJSON struct {
	Dimensions  apijson.Field
	Limit       apijson.Field
	Metrics     apijson.Field
	Since       apijson.Field
	Until       apijson.Field
	Filters     apijson.Field
	Sort        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneDNSAnalyticReportListResponseResultQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneDNSAnalyticReportListResponseSuccess bool

const (
	ZoneDNSAnalyticReportListResponseSuccessTrue ZoneDNSAnalyticReportListResponseSuccess = true
)

type ZoneDNSAnalyticReportListParams struct {
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
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [ZoneDNSAnalyticReportListParams]'s query parameters as
// `url.Values`.
func (r ZoneDNSAnalyticReportListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
