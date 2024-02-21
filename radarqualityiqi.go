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

// RadarQualityIqiService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarQualityIqiService] method
// instead.
type RadarQualityIqiService struct {
	Options []option.RequestOption
}

// NewRadarQualityIqiService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarQualityIqiService(opts ...option.RequestOption) (r *RadarQualityIqiService) {
	r = &RadarQualityIqiService{}
	r.Options = opts
	return
}

// Get a summary (percentiles) of bandwidth, latency or DNS response time from the
// Radar Internet Quality Index (IQI).
func (r *RadarQualityIqiService) Summary(ctx context.Context, query RadarQualityIqiSummaryParams, opts ...option.RequestOption) (res *RadarQualityIqiSummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarQualityIqiSummaryResponseEnvelope
	path := "radar/quality/iqi/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarQualityIqiSummaryResponse struct {
	Meta     RadarQualityIqiSummaryResponseMeta     `json:"meta,required"`
	Summary0 RadarQualityIqiSummaryResponseSummary0 `json:"summary_0,required"`
	JSON     radarQualityIqiSummaryResponseJSON     `json:"-"`
}

// radarQualityIqiSummaryResponseJSON contains the JSON metadata for the struct
// [RadarQualityIqiSummaryResponse]
type radarQualityIqiSummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryResponseMeta struct {
	DateRange      []RadarQualityIqiSummaryResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	Normalization  string                                           `json:"normalization,required"`
	ConfidenceInfo RadarQualityIqiSummaryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualityIqiSummaryResponseMetaJSON           `json:"-"`
}

// radarQualityIqiSummaryResponseMetaJSON contains the JSON metadata for the struct
// [RadarQualityIqiSummaryResponseMeta]
type radarQualityIqiSummaryResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      radarQualityIqiSummaryResponseMetaDateRangeJSON `json:"-"`
}

// radarQualityIqiSummaryResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarQualityIqiSummaryResponseMetaDateRange]
type radarQualityIqiSummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryResponseMetaConfidenceInfo struct {
	Annotations []RadarQualityIqiSummaryResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        radarQualityIqiSummaryResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualityIqiSummaryResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarQualityIqiSummaryResponseMetaConfidenceInfo]
type radarQualityIqiSummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            radarQualityIqiSummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualityIqiSummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarQualityIqiSummaryResponseMetaConfidenceInfoAnnotation]
type radarQualityIqiSummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualityIqiSummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryResponseSummary0 struct {
	P25  string                                     `json:"p25,required"`
	P50  string                                     `json:"p50,required"`
	P75  string                                     `json:"p75,required"`
	JSON radarQualityIqiSummaryResponseSummary0JSON `json:"-"`
}

// radarQualityIqiSummaryResponseSummary0JSON contains the JSON metadata for the
// struct [RadarQualityIqiSummaryResponseSummary0]
type radarQualityIqiSummaryResponseSummary0JSON struct {
	P25         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiSummaryParams struct {
	// Which metric to return: bandwidth, latency or DNS response time.
	Metric param.Field[RadarQualityIqiSummaryParamsMetric] `query:"metric,required"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarQualityIqiSummaryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualityIqiSummaryParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualityIqiSummaryParams]'s query parameters as
// `url.Values`.
func (r RadarQualityIqiSummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which metric to return: bandwidth, latency or DNS response time.
type RadarQualityIqiSummaryParamsMetric string

const (
	RadarQualityIqiSummaryParamsMetricBandwidth RadarQualityIqiSummaryParamsMetric = "BANDWIDTH"
	RadarQualityIqiSummaryParamsMetricDNS       RadarQualityIqiSummaryParamsMetric = "DNS"
	RadarQualityIqiSummaryParamsMetricLatency   RadarQualityIqiSummaryParamsMetric = "LATENCY"
)

type RadarQualityIqiSummaryParamsDateRange string

const (
	RadarQualityIqiSummaryParamsDateRange1d         RadarQualityIqiSummaryParamsDateRange = "1d"
	RadarQualityIqiSummaryParamsDateRange2d         RadarQualityIqiSummaryParamsDateRange = "2d"
	RadarQualityIqiSummaryParamsDateRange7d         RadarQualityIqiSummaryParamsDateRange = "7d"
	RadarQualityIqiSummaryParamsDateRange14d        RadarQualityIqiSummaryParamsDateRange = "14d"
	RadarQualityIqiSummaryParamsDateRange28d        RadarQualityIqiSummaryParamsDateRange = "28d"
	RadarQualityIqiSummaryParamsDateRange12w        RadarQualityIqiSummaryParamsDateRange = "12w"
	RadarQualityIqiSummaryParamsDateRange24w        RadarQualityIqiSummaryParamsDateRange = "24w"
	RadarQualityIqiSummaryParamsDateRange52w        RadarQualityIqiSummaryParamsDateRange = "52w"
	RadarQualityIqiSummaryParamsDateRange1dControl  RadarQualityIqiSummaryParamsDateRange = "1dControl"
	RadarQualityIqiSummaryParamsDateRange2dControl  RadarQualityIqiSummaryParamsDateRange = "2dControl"
	RadarQualityIqiSummaryParamsDateRange7dControl  RadarQualityIqiSummaryParamsDateRange = "7dControl"
	RadarQualityIqiSummaryParamsDateRange14dControl RadarQualityIqiSummaryParamsDateRange = "14dControl"
	RadarQualityIqiSummaryParamsDateRange28dControl RadarQualityIqiSummaryParamsDateRange = "28dControl"
	RadarQualityIqiSummaryParamsDateRange12wControl RadarQualityIqiSummaryParamsDateRange = "12wControl"
	RadarQualityIqiSummaryParamsDateRange24wControl RadarQualityIqiSummaryParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarQualityIqiSummaryParamsFormat string

const (
	RadarQualityIqiSummaryParamsFormatJson RadarQualityIqiSummaryParamsFormat = "JSON"
	RadarQualityIqiSummaryParamsFormatCsv  RadarQualityIqiSummaryParamsFormat = "CSV"
)

type RadarQualityIqiSummaryResponseEnvelope struct {
	Result  RadarQualityIqiSummaryResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarQualityIqiSummaryResponseEnvelopeJSON `json:"-"`
}

// radarQualityIqiSummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarQualityIqiSummaryResponseEnvelope]
type radarQualityIqiSummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiSummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
