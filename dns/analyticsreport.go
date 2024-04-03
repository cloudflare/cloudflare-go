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
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AnalyticsReportService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAnalyticsReportService] method
// instead.
type AnalyticsReportService struct {
	Options []option.RequestOption
	Bytimes *AnalyticsReportBytimeService
}

// NewAnalyticsReportService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnalyticsReportService(opts ...option.RequestOption) (r *AnalyticsReportService) {
	r = &AnalyticsReportService{}
	r.Options = opts
	r.Bytimes = NewAnalyticsReportBytimeService(opts...)
	return
}

// Retrieves a list of summarised aggregate metrics over a given time period.
//
// See
// [Analytics API properties](https://developers.cloudflare.com/dns/reference/analytics-api-properties/)
// for detailed information about the available query parameters.
func (r *AnalyticsReportService) Get(ctx context.Context, params AnalyticsReportGetParams, opts ...option.RequestOption) (res *DNSAnalyticsReport, err error) {
	opts = append(r.Options[:], opts...)
	var env AnalyticsReportGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/dns_analytics/report", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DNSAnalyticsReport struct {
	// Array with one row per combination of dimension values.
	Data []DNSAnalyticsReportData `json:"data,required"`
	// Number of seconds between current time and last processed event, in another
	// words how many seconds of data could be missing.
	DataLag float64 `json:"data_lag,required"`
	// Maximum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Max interface{} `json:"max,required"`
	// Minimum results for each metric (object mapping metric names to values).
	// Currently always an empty object.
	Min   interface{}             `json:"min,required"`
	Query DNSAnalyticsReportQuery `json:"query,required"`
	// Total number of rows in the result.
	Rows float64 `json:"rows,required"`
	// Total results for metrics across all data (object mapping metric names to
	// values).
	Totals interface{}            `json:"totals,required"`
	JSON   dnsAnalyticsReportJSON `json:"-"`
}

// dnsAnalyticsReportJSON contains the JSON metadata for the struct
// [DNSAnalyticsReport]
type dnsAnalyticsReportJSON struct {
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

func (r *DNSAnalyticsReport) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsAnalyticsReportJSON) RawJSON() string {
	return r.raw
}

type DNSAnalyticsReportData struct {
	// Array of dimension values, representing the combination of dimension values
	// corresponding to this row.
	Dimensions []string `json:"dimensions,required"`
	// Array with one item per requested metric. Each item is a single value.
	Metrics []float64                  `json:"metrics,required"`
	JSON    dnsAnalyticsReportDataJSON `json:"-"`
}

// dnsAnalyticsReportDataJSON contains the JSON metadata for the struct
// [DNSAnalyticsReportData]
type dnsAnalyticsReportDataJSON struct {
	Dimensions  apijson.Field
	Metrics     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DNSAnalyticsReportData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsAnalyticsReportDataJSON) RawJSON() string {
	return r.raw
}

type DNSAnalyticsReportQuery struct {
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
	Sort []string                    `json:"sort"`
	JSON dnsAnalyticsReportQueryJSON `json:"-"`
}

// dnsAnalyticsReportQueryJSON contains the JSON metadata for the struct
// [DNSAnalyticsReportQuery]
type dnsAnalyticsReportQueryJSON struct {
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

func (r *DNSAnalyticsReportQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dnsAnalyticsReportQueryJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportGetParams struct {
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
	// End date and time of requesting data period in ISO 8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [AnalyticsReportGetParams]'s query parameters as
// `url.Values`.
func (r AnalyticsReportGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AnalyticsReportGetResponseEnvelope struct {
	Errors   []AnalyticsReportGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AnalyticsReportGetResponseEnvelopeMessages `json:"messages,required"`
	Result   DNSAnalyticsReport                           `json:"result,required"`
	// Whether the API call was successful
	Success AnalyticsReportGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    analyticsReportGetResponseEnvelopeJSON    `json:"-"`
}

// analyticsReportGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AnalyticsReportGetResponseEnvelope]
type analyticsReportGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsReportGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportGetResponseEnvelopeErrors struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    analyticsReportGetResponseEnvelopeErrorsJSON `json:"-"`
}

// analyticsReportGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AnalyticsReportGetResponseEnvelopeErrors]
type analyticsReportGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsReportGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AnalyticsReportGetResponseEnvelopeMessages struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    analyticsReportGetResponseEnvelopeMessagesJSON `json:"-"`
}

// analyticsReportGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AnalyticsReportGetResponseEnvelopeMessages]
type analyticsReportGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsReportGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r analyticsReportGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AnalyticsReportGetResponseEnvelopeSuccess bool

const (
	AnalyticsReportGetResponseEnvelopeSuccessTrue AnalyticsReportGetResponseEnvelopeSuccess = true
)

func (r AnalyticsReportGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AnalyticsReportGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
