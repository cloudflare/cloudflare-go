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

// DNSAnalyticsReportService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewDNSAnalyticsReportService] method
// instead.
type DNSAnalyticsReportService struct {
	Options []option.RequestOption
	Bytimes *DNSAnalyticsReportBytimeService
}

// NewDNSAnalyticsReportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDNSAnalyticsReportService(opts ...option.RequestOption) (r *DNSAnalyticsReportService) {
	r = &DNSAnalyticsReportService{}
	r.Options = opts
	r.Bytimes = NewDNSAnalyticsReportBytimeService(opts...)
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *DNSAnalyticsReportService) Get(ctx context.Context, identifier string, query DNSAnalyticsReportGetParams, opts ...option.RequestOption) (res *DNSDNSAnalyticsAPIReport, err error) {
	opts = append(r.Options[:], opts...)
	var env DNSAnalyticsReportGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_analytics/report", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSDNSAnalyticsAPIReport struct {
	// Array with one row per combination of dimension values.
	Data []DNSDNSAnalyticsAPIReportData `json:"data,required"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}                   `json:"min,required"`
	Query DNSDNSAnalyticsAPIReportQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}                  `json:"totals,required"`
	JSON   dnsdnsAnalyticsAPIReportJSON `json:"-"`
}

// dnsdnsAnalyticsAPIReportJSON contains the JSON metadata for the struct
// [DNSDNSAnalyticsAPIReport]
type dnsdnsAnalyticsAPIReportJSON struct {
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

func (r *DNSDNSAnalyticsAPIReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSDNSAnalyticsAPIReportData struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is a single value.
	Metrics []float64                        `json:"metrics,required"`
	JSON    dnsdnsAnalyticsAPIReportDataJSON `json:"-"`
}

// dnsdnsAnalyticsAPIReportDataJSON contains the JSON metadata for the struct
// [DNSDNSAnalyticsAPIReportData]
type dnsdnsAnalyticsAPIReportDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSDNSAnalyticsAPIReportData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSDNSAnalyticsAPIReportQuery struct {
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
	Sort []string                          `json:"sort"`
	JSON dnsdnsAnalyticsAPIReportQueryJSON `json:"-"`
}

// dnsdnsAnalyticsAPIReportQueryJSON contains the JSON metadata for the struct
// [DNSDNSAnalyticsAPIReportQuery]
type dnsdnsAnalyticsAPIReportQueryJSON struct {
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

func (r *DNSDNSAnalyticsAPIReportQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSAnalyticsReportGetParams struct {
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

// URLQuery serializes [DNSAnalyticsReportGetParams]'s query parameters as
// `url.Values`.
func (r DNSAnalyticsReportGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DNSAnalyticsReportGetResponseEnvelope struct {
	Errors   []DNSAnalyticsReportGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DNSAnalyticsReportGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSDNSAnalyticsAPIReport                        `json:"result,required"`
	// Whether the API call was successful
	Success DNSAnalyticsReportGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    dnsAnalyticsReportGetResponseEnvelopeJSON    `json:"-"`
}

// dnsAnalyticsReportGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DNSAnalyticsReportGetResponseEnvelope]
type dnsAnalyticsReportGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticsReportGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSAnalyticsReportGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    dnsAnalyticsReportGetResponseEnvelopeErrorsJSON `json:"-"`
}

// dnsAnalyticsReportGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DNSAnalyticsReportGetResponseEnvelopeErrors]
type dnsAnalyticsReportGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticsReportGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DNSAnalyticsReportGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    dnsAnalyticsReportGetResponseEnvelopeMessagesJSON `json:"-"`
}

// dnsAnalyticsReportGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DNSAnalyticsReportGetResponseEnvelopeMessages]
type dnsAnalyticsReportGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticsReportGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type DNSAnalyticsReportGetResponseEnvelopeSuccess bool

const (
	DNSAnalyticsReportGetResponseEnvelopeSuccessTrue DNSAnalyticsReportGetResponseEnvelopeSuccess = true
)
