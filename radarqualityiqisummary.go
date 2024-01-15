// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarQualityIqiSummaryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarQualityIqiSummaryService]
// method instead.
type RadarQualityIqiSummaryService struct {
	Options []option.RequestOption
}

// NewRadarQualityIqiSummaryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarQualityIqiSummaryService(opts ...option.RequestOption) (r *RadarQualityIqiSummaryService) {
	r = &RadarQualityIqiSummaryService{}
	r.Options = opts
	return
}

// Get a summary (percentiles) of bandwidth, latency or DNS response time from the
// Radar Internet Quality Index (IQI).
func (r *RadarQualityIqiSummaryService) Get(ctx context.Context, query RadarQualityIqiSummaryGetParams, opts ...option.RequestOption) (res *RadarQualityIqiSummaryGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/quality/iqi/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarQualityIqiSummaryGetResponse struct {
	Result  RadarQualityIqiSummaryGetResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarQualityIqiSummaryGetResponseJSON   `json:"-"`
}

// radarQualityIqiSummaryGetResponseJSON contains the JSON metadata for the struct
// [RadarQualityIqiSummaryGetResponse]
type radarQualityIqiSummaryGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryGetResponseResult struct {
	Meta     RadarQualityIqiSummaryGetResponseResultMeta     `json:"meta,required"`
	Summary0 RadarQualityIqiSummaryGetResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarQualityIqiSummaryGetResponseResultJSON     `json:"-"`
}

// radarQualityIqiSummaryGetResponseResultJSON contains the JSON metadata for the
// struct [RadarQualityIqiSummaryGetResponseResult]
type radarQualityIqiSummaryGetResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryGetResponseResultMeta struct {
	DateRange      []RadarQualityIqiSummaryGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                    `json:"lastUpdated,required"`
	Normalization  string                                                    `json:"normalization,required"`
	ConfidenceInfo RadarQualityIqiSummaryGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualityIqiSummaryGetResponseResultMetaJSON           `json:"-"`
}

// radarQualityIqiSummaryGetResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarQualityIqiSummaryGetResponseResultMeta]
type radarQualityIqiSummaryGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarQualityIqiSummaryGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarQualityIqiSummaryGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarQualityIqiSummaryGetResponseResultMetaDateRange]
type radarQualityIqiSummaryGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarQualityIqiSummaryGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                 `json:"level"`
	JSON        radarQualityIqiSummaryGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualityIqiSummaryGetResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarQualityIqiSummaryGetResponseResultMetaConfidenceInfo]
type radarQualityIqiSummaryGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                  `json:"dataSource,required"`
	Description     string                                                                  `json:"description,required"`
	EventType       string                                                                  `json:"eventType,required"`
	IsInstantaneous interface{}                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                  `json:"linkedUrl"`
	StartTime       time.Time                                                               `json:"startTime" format:"date-time"`
	JSON            radarQualityIqiSummaryGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualityIqiSummaryGetResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarQualityIqiSummaryGetResponseResultMetaConfidenceInfoAnnotation]
type radarQualityIqiSummaryGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryGetResponseResultSummary0 struct {
	P25  string                                              `json:"p25,required"`
	P50  string                                              `json:"p50,required"`
	P75  string                                              `json:"p75,required"`
	JSON radarQualityIqiSummaryGetResponseResultSummary0JSON `json:"-"`
}

// radarQualityIqiSummaryGetResponseResultSummary0JSON contains the JSON metadata
// for the struct [RadarQualityIqiSummaryGetResponseResultSummary0]
type radarQualityIqiSummaryGetResponseResultSummary0JSON struct {
	P25         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryGetResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryGetParams struct {
	// Which metric to return: bandwidth, latency or DNS response time.
	Metric param.Field[RadarQualityIqiSummaryGetParamsMetric] `query:"metric,required"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarQualityIqiSummaryGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualityIqiSummaryGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualityIqiSummaryGetParams]'s query parameters as
// `url.Values`.
func (r RadarQualityIqiSummaryGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which metric to return: bandwidth, latency or DNS response time.
type RadarQualityIqiSummaryGetParamsMetric string

const (
	RadarQualityIqiSummaryGetParamsMetricBandwidth RadarQualityIqiSummaryGetParamsMetric = "BANDWIDTH"
	RadarQualityIqiSummaryGetParamsMetricDNS       RadarQualityIqiSummaryGetParamsMetric = "DNS"
	RadarQualityIqiSummaryGetParamsMetricLatency   RadarQualityIqiSummaryGetParamsMetric = "LATENCY"
)

type RadarQualityIqiSummaryGetParamsDateRange string

const (
	RadarQualityIqiSummaryGetParamsDateRange1d         RadarQualityIqiSummaryGetParamsDateRange = "1d"
	RadarQualityIqiSummaryGetParamsDateRange2d         RadarQualityIqiSummaryGetParamsDateRange = "2d"
	RadarQualityIqiSummaryGetParamsDateRange7d         RadarQualityIqiSummaryGetParamsDateRange = "7d"
	RadarQualityIqiSummaryGetParamsDateRange14d        RadarQualityIqiSummaryGetParamsDateRange = "14d"
	RadarQualityIqiSummaryGetParamsDateRange28d        RadarQualityIqiSummaryGetParamsDateRange = "28d"
	RadarQualityIqiSummaryGetParamsDateRange12w        RadarQualityIqiSummaryGetParamsDateRange = "12w"
	RadarQualityIqiSummaryGetParamsDateRange24w        RadarQualityIqiSummaryGetParamsDateRange = "24w"
	RadarQualityIqiSummaryGetParamsDateRange52w        RadarQualityIqiSummaryGetParamsDateRange = "52w"
	RadarQualityIqiSummaryGetParamsDateRange1dControl  RadarQualityIqiSummaryGetParamsDateRange = "1dControl"
	RadarQualityIqiSummaryGetParamsDateRange2dControl  RadarQualityIqiSummaryGetParamsDateRange = "2dControl"
	RadarQualityIqiSummaryGetParamsDateRange7dControl  RadarQualityIqiSummaryGetParamsDateRange = "7dControl"
	RadarQualityIqiSummaryGetParamsDateRange14dControl RadarQualityIqiSummaryGetParamsDateRange = "14dControl"
	RadarQualityIqiSummaryGetParamsDateRange28dControl RadarQualityIqiSummaryGetParamsDateRange = "28dControl"
	RadarQualityIqiSummaryGetParamsDateRange12wControl RadarQualityIqiSummaryGetParamsDateRange = "12wControl"
	RadarQualityIqiSummaryGetParamsDateRange24wControl RadarQualityIqiSummaryGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarQualityIqiSummaryGetParamsFormat string

const (
	RadarQualityIqiSummaryGetParamsFormatJson RadarQualityIqiSummaryGetParamsFormat = "JSON"
	RadarQualityIqiSummaryGetParamsFormatCsv  RadarQualityIqiSummaryGetParamsFormat = "CSV"
)
