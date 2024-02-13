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

// DNSAnalyticReportService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDNSAnalyticReportService] method
// instead.
type DNSAnalyticReportService struct {
	Options []option.RequestOption
	Bytimes *DNSAnalyticReportBytimeService
}

// NewDNSAnalyticReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDNSAnalyticReportService(opts ...option.RequestOption) (r *DNSAnalyticReportService) {
	r = &DNSAnalyticReportService{}
	r.Options = opts
	r.Bytimes = NewDNSAnalyticReportBytimeService(opts...)
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *DNSAnalyticReportService) List(ctx context.Context, identifier string, query DNSAnalyticReportListParams, opts ...option.RequestOption) (res *DNSAnalyticReportListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSAnalyticReportListResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_analytics/report", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSAnalyticReportListResponse struct {
	// Array with one row per combination of dimension values.
	Data []DNSAnalyticReportListResponseData `json:"data,required"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}                        `json:"min,required"`
	Query DNSAnalyticReportListResponseQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                       `json:"totals,required"`
	JSON   dnsAnalyticReportListResponseJSON `json:"-"`
}

// dnsAnalyticReportListResponseJSON contains the JSON metadata for the struct
// [DNSAnalyticReportListResponse]
type dnsAnalyticReportListResponseJSON struct {
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

func (r *DNSAnalyticReportListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSAnalyticReportListResponseData struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is a single value.
	Metrics []float64                             `json:"metrics,required"`
	JSON    dnsAnalyticReportListResponseDataJSON `json:"-"`
}

// dnsAnalyticReportListResponseDataJSON contains the JSON metadata for the struct
// [DNSAnalyticReportListResponseData]
type dnsAnalyticReportListResponseDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticReportListResponseData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSAnalyticReportListResponseQuery struct {
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
	Sort []string                               `json:"sort"`
	JSON dnsAnalyticReportListResponseQueryJSON `json:"-"`
}

// dnsAnalyticReportListResponseQueryJSON contains the JSON metadata for the struct
// [DNSAnalyticReportListResponseQuery]
type dnsAnalyticReportListResponseQueryJSON struct {
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

func (r *DNSAnalyticReportListResponseQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSAnalyticReportListParams struct {
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

// URLQuery serializes [DNSAnalyticReportListParams]'s query parameters as
// `url.Values`.
func (r DNSAnalyticReportListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DNSAnalyticReportListResponseEnvelope struct {
	Errors   []DNSAnalyticReportListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSAnalyticReportListResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSAnalyticReportListResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success DNSAnalyticReportListResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsAnalyticReportListResponseEnvelopeJSON    `json:"-"`
}

// dnsAnalyticReportListResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSAnalyticReportListResponseEnvelope]
type dnsAnalyticReportListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticReportListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSAnalyticReportListResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    dnsAnalyticReportListResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsAnalyticReportListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DNSAnalyticReportListResponseEnvelopeErrors]
type dnsAnalyticReportListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticReportListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSAnalyticReportListResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    dnsAnalyticReportListResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsAnalyticReportListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DNSAnalyticReportListResponseEnvelopeMessages]
type dnsAnalyticReportListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticReportListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSAnalyticReportListResponseEnvelopeSuccess bool

const (
	DNSAnalyticReportListResponseEnvelopeSuccessTrue DNSAnalyticReportListResponseEnvelopeSuccess = true
)
