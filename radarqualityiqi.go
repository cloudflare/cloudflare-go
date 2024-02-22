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

// RadarQualityIQIService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarQualityIQIService] method
// instead.
type RadarQualityIQIService struct {
	Options []option.RequestOption
}

// NewRadarQualityIQIService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarQualityIQIService(opts ...option.RequestOption) (r *RadarQualityIQIService) {
	r = &RadarQualityIQIService{}
	r.Options = opts
	return
}

// Get a summary (percentiles) of bandwidth, latency or DNS response time from the
// Radar Internet Quality Index (IQI).
func (r *RadarQualityIQIService) Summary(ctx context.Context, query RadarQualityIQISummaryParams, opts ...option.RequestOption) (res *RadarQualityIQISummaryResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarQualityIQISummaryResponseEnvelope
	path := "radar/quality/iqi/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get a time series (percentiles) of bandwidth, latency or DNS response time from
// the Radar Internet Quality Index (IQI).
func (r *RadarQualityIQIService) TimeseriesGroups(ctx context.Context, query RadarQualityIQITimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarQualityIQITimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarQualityIQITimeseriesGroupsResponseEnvelope
	path := "radar/quality/iqi/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarQualityIQISummaryResponse struct {
	Meta     RadarQualityIQISummaryResponseMeta     `json:"meta,required"`
	Summary0 RadarQualityIQISummaryResponseSummary0 `json:"summary_0,required"`
	JSON     radarQualityIQISummaryResponseJSON     `json:"-"`
}

// radarQualityIQISummaryResponseJSON contains the JSON metadata for the struct
// [RadarQualityIQISummaryResponse]
type radarQualityIQISummaryResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIQISummaryResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIQISummaryResponseMeta struct {
	DateRange      []RadarQualityIQISummaryResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	Normalization  string                                           `json:"normalization,required"`
	ConfidenceInfo RadarQualityIQISummaryResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualityIQISummaryResponseMetaJSON           `json:"-"`
}

// radarQualityIQISummaryResponseMetaJSON contains the JSON metadata for the struct
// [RadarQualityIQISummaryResponseMeta]
type radarQualityIQISummaryResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualityIQISummaryResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIQISummaryResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      radarQualityIQISummaryResponseMetaDateRangeJSON `json:"-"`
}

// radarQualityIQISummaryResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarQualityIQISummaryResponseMetaDateRange]
type radarQualityIQISummaryResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIQISummaryResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIQISummaryResponseMetaConfidenceInfo struct {
	Annotations []RadarQualityIQISummaryResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        radarQualityIQISummaryResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualityIQISummaryResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarQualityIQISummaryResponseMetaConfidenceInfo]
type radarQualityIQISummaryResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIQISummaryResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIQISummaryResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            radarQualityIQISummaryResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualityIQISummaryResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarQualityIQISummaryResponseMetaConfidenceInfoAnnotation]
type radarQualityIQISummaryResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualityIQISummaryResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIQISummaryResponseSummary0 struct {
	P25  string                                     `json:"p25,required"`
	P50  string                                     `json:"p50,required"`
	P75  string                                     `json:"p75,required"`
	JSON radarQualityIQISummaryResponseSummary0JSON `json:"-"`
}

// radarQualityIQISummaryResponseSummary0JSON contains the JSON metadata for the
// struct [RadarQualityIQISummaryResponseSummary0]
type radarQualityIQISummaryResponseSummary0JSON struct {
	P25         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIQISummaryResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIQITimeseriesGroupsResponse struct {
	Meta   interface{}                                   `json:"meta,required"`
	Serie0 RadarQualityIQITimeseriesGroupsResponseSerie0 `json:"serie_0,required"`
	JSON   radarQualityIQITimeseriesGroupsResponseJSON   `json:"-"`
}

// radarQualityIQITimeseriesGroupsResponseJSON contains the JSON metadata for the
// struct [RadarQualityIQITimeseriesGroupsResponse]
type radarQualityIQITimeseriesGroupsResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIQITimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIQITimeseriesGroupsResponseSerie0 struct {
	P25        []string                                          `json:"p25,required"`
	P50        []string                                          `json:"p50,required"`
	P75        []string                                          `json:"p75,required"`
	Timestamps []string                                          `json:"timestamps,required"`
	JSON       radarQualityIQITimeseriesGroupsResponseSerie0JSON `json:"-"`
}

// radarQualityIQITimeseriesGroupsResponseSerie0JSON contains the JSON metadata for
// the struct [RadarQualityIQITimeseriesGroupsResponseSerie0]
type radarQualityIQITimeseriesGroupsResponseSerie0JSON struct {
	P25         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIQITimeseriesGroupsResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIQISummaryParams struct {
	// Which metric to return: bandwidth, latency or DNS response time.
	Metric param.Field[RadarQualityIQISummaryParamsMetric] `query:"metric,required"`
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
	DateRange param.Field[[]RadarQualityIQISummaryParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualityIQISummaryParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualityIQISummaryParams]'s query parameters as
// `url.Values`.
func (r RadarQualityIQISummaryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which metric to return: bandwidth, latency or DNS response time.
type RadarQualityIQISummaryParamsMetric string

const (
	RadarQualityIQISummaryParamsMetricBandwidth RadarQualityIQISummaryParamsMetric = "BANDWIDTH"
	RadarQualityIQISummaryParamsMetricDNS       RadarQualityIQISummaryParamsMetric = "DNS"
	RadarQualityIQISummaryParamsMetricLatency   RadarQualityIQISummaryParamsMetric = "LATENCY"
)

type RadarQualityIQISummaryParamsDateRange string

const (
	RadarQualityIQISummaryParamsDateRange1d         RadarQualityIQISummaryParamsDateRange = "1d"
	RadarQualityIQISummaryParamsDateRange2d         RadarQualityIQISummaryParamsDateRange = "2d"
	RadarQualityIQISummaryParamsDateRange7d         RadarQualityIQISummaryParamsDateRange = "7d"
	RadarQualityIQISummaryParamsDateRange14d        RadarQualityIQISummaryParamsDateRange = "14d"
	RadarQualityIQISummaryParamsDateRange28d        RadarQualityIQISummaryParamsDateRange = "28d"
	RadarQualityIQISummaryParamsDateRange12w        RadarQualityIQISummaryParamsDateRange = "12w"
	RadarQualityIQISummaryParamsDateRange24w        RadarQualityIQISummaryParamsDateRange = "24w"
	RadarQualityIQISummaryParamsDateRange52w        RadarQualityIQISummaryParamsDateRange = "52w"
	RadarQualityIQISummaryParamsDateRange1dControl  RadarQualityIQISummaryParamsDateRange = "1dControl"
	RadarQualityIQISummaryParamsDateRange2dControl  RadarQualityIQISummaryParamsDateRange = "2dControl"
	RadarQualityIQISummaryParamsDateRange7dControl  RadarQualityIQISummaryParamsDateRange = "7dControl"
	RadarQualityIQISummaryParamsDateRange14dControl RadarQualityIQISummaryParamsDateRange = "14dControl"
	RadarQualityIQISummaryParamsDateRange28dControl RadarQualityIQISummaryParamsDateRange = "28dControl"
	RadarQualityIQISummaryParamsDateRange12wControl RadarQualityIQISummaryParamsDateRange = "12wControl"
	RadarQualityIQISummaryParamsDateRange24wControl RadarQualityIQISummaryParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarQualityIQISummaryParamsFormat string

const (
	RadarQualityIQISummaryParamsFormatJson RadarQualityIQISummaryParamsFormat = "JSON"
	RadarQualityIQISummaryParamsFormatCsv  RadarQualityIQISummaryParamsFormat = "CSV"
)

type RadarQualityIQISummaryResponseEnvelope struct {
	Result  RadarQualityIQISummaryResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarQualityIQISummaryResponseEnvelopeJSON `json:"-"`
}

// radarQualityIQISummaryResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarQualityIQISummaryResponseEnvelope]
type radarQualityIQISummaryResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIQISummaryResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIQITimeseriesGroupsParams struct {
	// Which metric to return: bandwidth, latency or DNS response time.
	Metric param.Field[RadarQualityIQITimeseriesGroupsParamsMetric] `query:"metric,required"`
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarQualityIQITimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarQualityIQITimeseriesGroupsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualityIQITimeseriesGroupsParamsFormat] `query:"format"`
	// Enable interpolation for all series (using the average).
	Interpolation param.Field[bool] `query:"interpolation"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualityIQITimeseriesGroupsParams]'s query parameters
// as `url.Values`.
func (r RadarQualityIQITimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which metric to return: bandwidth, latency or DNS response time.
type RadarQualityIQITimeseriesGroupsParamsMetric string

const (
	RadarQualityIQITimeseriesGroupsParamsMetricBandwidth RadarQualityIQITimeseriesGroupsParamsMetric = "BANDWIDTH"
	RadarQualityIQITimeseriesGroupsParamsMetricDNS       RadarQualityIQITimeseriesGroupsParamsMetric = "DNS"
	RadarQualityIQITimeseriesGroupsParamsMetricLatency   RadarQualityIQITimeseriesGroupsParamsMetric = "LATENCY"
)

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarQualityIQITimeseriesGroupsParamsAggInterval string

const (
	RadarQualityIQITimeseriesGroupsParamsAggInterval15m RadarQualityIQITimeseriesGroupsParamsAggInterval = "15m"
	RadarQualityIQITimeseriesGroupsParamsAggInterval1h  RadarQualityIQITimeseriesGroupsParamsAggInterval = "1h"
	RadarQualityIQITimeseriesGroupsParamsAggInterval1d  RadarQualityIQITimeseriesGroupsParamsAggInterval = "1d"
	RadarQualityIQITimeseriesGroupsParamsAggInterval1w  RadarQualityIQITimeseriesGroupsParamsAggInterval = "1w"
)

type RadarQualityIQITimeseriesGroupsParamsDateRange string

const (
	RadarQualityIQITimeseriesGroupsParamsDateRange1d         RadarQualityIQITimeseriesGroupsParamsDateRange = "1d"
	RadarQualityIQITimeseriesGroupsParamsDateRange2d         RadarQualityIQITimeseriesGroupsParamsDateRange = "2d"
	RadarQualityIQITimeseriesGroupsParamsDateRange7d         RadarQualityIQITimeseriesGroupsParamsDateRange = "7d"
	RadarQualityIQITimeseriesGroupsParamsDateRange14d        RadarQualityIQITimeseriesGroupsParamsDateRange = "14d"
	RadarQualityIQITimeseriesGroupsParamsDateRange28d        RadarQualityIQITimeseriesGroupsParamsDateRange = "28d"
	RadarQualityIQITimeseriesGroupsParamsDateRange12w        RadarQualityIQITimeseriesGroupsParamsDateRange = "12w"
	RadarQualityIQITimeseriesGroupsParamsDateRange24w        RadarQualityIQITimeseriesGroupsParamsDateRange = "24w"
	RadarQualityIQITimeseriesGroupsParamsDateRange52w        RadarQualityIQITimeseriesGroupsParamsDateRange = "52w"
	RadarQualityIQITimeseriesGroupsParamsDateRange1dControl  RadarQualityIQITimeseriesGroupsParamsDateRange = "1dControl"
	RadarQualityIQITimeseriesGroupsParamsDateRange2dControl  RadarQualityIQITimeseriesGroupsParamsDateRange = "2dControl"
	RadarQualityIQITimeseriesGroupsParamsDateRange7dControl  RadarQualityIQITimeseriesGroupsParamsDateRange = "7dControl"
	RadarQualityIQITimeseriesGroupsParamsDateRange14dControl RadarQualityIQITimeseriesGroupsParamsDateRange = "14dControl"
	RadarQualityIQITimeseriesGroupsParamsDateRange28dControl RadarQualityIQITimeseriesGroupsParamsDateRange = "28dControl"
	RadarQualityIQITimeseriesGroupsParamsDateRange12wControl RadarQualityIQITimeseriesGroupsParamsDateRange = "12wControl"
	RadarQualityIQITimeseriesGroupsParamsDateRange24wControl RadarQualityIQITimeseriesGroupsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarQualityIQITimeseriesGroupsParamsFormat string

const (
	RadarQualityIQITimeseriesGroupsParamsFormatJson RadarQualityIQITimeseriesGroupsParamsFormat = "JSON"
	RadarQualityIQITimeseriesGroupsParamsFormatCsv  RadarQualityIQITimeseriesGroupsParamsFormat = "CSV"
)

type RadarQualityIQITimeseriesGroupsResponseEnvelope struct {
	Result  RadarQualityIQITimeseriesGroupsResponse             `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarQualityIQITimeseriesGroupsResponseEnvelopeJSON `json:"-"`
}

// radarQualityIQITimeseriesGroupsResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarQualityIQITimeseriesGroupsResponseEnvelope]
type radarQualityIQITimeseriesGroupsResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIQITimeseriesGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
