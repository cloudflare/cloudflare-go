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
	Options          []option.RequestOption
	TimeseriesGroups *RadarQualityIqiTimeseriesGroupService
}

// NewRadarQualityIqiService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarQualityIqiService(opts ...option.RequestOption) (r *RadarQualityIqiService) {
	r = &RadarQualityIqiService{}
	r.Options = opts
	r.TimeseriesGroups = NewRadarQualityIqiTimeseriesGroupService(opts...)
	return
}

// Get a summary (percentiles) of bandwidth, latency or DNS response time from the
// Radar Internet Quality Index (IQI).
func (r *RadarQualityIqiService) List(ctx context.Context, query RadarQualityIqiListParams, opts ...option.RequestOption) (res *RadarQualityIqiListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarQualityIqiListResponseEnvelope
	path := "radar/quality/iqi/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarQualityIqiListResponse struct {
	Meta     RadarQualityIqiListResponseMeta     `json:"meta,required"`
	Summary0 RadarQualityIqiListResponseSummary0 `json:"summary_0,required"`
	JSON     radarQualityIqiListResponseJSON     `json:"-"`
}

// radarQualityIqiListResponseJSON contains the JSON metadata for the struct
// [RadarQualityIqiListResponse]
type radarQualityIqiListResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiListResponseMeta struct {
	DateRange      []RadarQualityIqiListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                        `json:"lastUpdated,required"`
	Normalization  string                                        `json:"normalization,required"`
	ConfidenceInfo RadarQualityIqiListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarQualityIqiListResponseMetaJSON           `json:"-"`
}

// radarQualityIqiListResponseMetaJSON contains the JSON metadata for the struct
// [RadarQualityIqiListResponseMeta]
type radarQualityIqiListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarQualityIqiListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      radarQualityIqiListResponseMetaDateRangeJSON `json:"-"`
}

// radarQualityIqiListResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarQualityIqiListResponseMetaDateRange]
type radarQualityIqiListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiListResponseMetaConfidenceInfo struct {
	Annotations []RadarQualityIqiListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                     `json:"level"`
	JSON        radarQualityIqiListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarQualityIqiListResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [RadarQualityIqiListResponseMetaConfidenceInfo]
type radarQualityIqiListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                      `json:"dataSource,required"`
	Description     string                                                      `json:"description,required"`
	EventType       string                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                   `json:"startTime" format:"date-time"`
	JSON            radarQualityIqiListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarQualityIqiListResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarQualityIqiListResponseMetaConfidenceInfoAnnotation]
type radarQualityIqiListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarQualityIqiListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiListResponseSummary0 struct {
	P25  string                                  `json:"p25,required"`
	P50  string                                  `json:"p50,required"`
	P75  string                                  `json:"p75,required"`
	JSON radarQualityIqiListResponseSummary0JSON `json:"-"`
}

// radarQualityIqiListResponseSummary0JSON contains the JSON metadata for the
// struct [RadarQualityIqiListResponseSummary0]
type radarQualityIqiListResponseSummary0JSON struct {
	P25         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiListResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiListParams struct {
	// Which metric to return: bandwidth, latency or DNS response time.
	Metric param.Field[RadarQualityIqiListParamsMetric] `query:"metric,required"`
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
	DateRange param.Field[[]RadarQualityIqiListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualityIqiListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualityIqiListParams]'s query parameters as
// `url.Values`.
func (r RadarQualityIqiListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which metric to return: bandwidth, latency or DNS response time.
type RadarQualityIqiListParamsMetric string

const (
	RadarQualityIqiListParamsMetricBandwidth RadarQualityIqiListParamsMetric = "BANDWIDTH"
	RadarQualityIqiListParamsMetricDNS       RadarQualityIqiListParamsMetric = "DNS"
	RadarQualityIqiListParamsMetricLatency   RadarQualityIqiListParamsMetric = "LATENCY"
)

type RadarQualityIqiListParamsDateRange string

const (
	RadarQualityIqiListParamsDateRange1d         RadarQualityIqiListParamsDateRange = "1d"
	RadarQualityIqiListParamsDateRange2d         RadarQualityIqiListParamsDateRange = "2d"
	RadarQualityIqiListParamsDateRange7d         RadarQualityIqiListParamsDateRange = "7d"
	RadarQualityIqiListParamsDateRange14d        RadarQualityIqiListParamsDateRange = "14d"
	RadarQualityIqiListParamsDateRange28d        RadarQualityIqiListParamsDateRange = "28d"
	RadarQualityIqiListParamsDateRange12w        RadarQualityIqiListParamsDateRange = "12w"
	RadarQualityIqiListParamsDateRange24w        RadarQualityIqiListParamsDateRange = "24w"
	RadarQualityIqiListParamsDateRange52w        RadarQualityIqiListParamsDateRange = "52w"
	RadarQualityIqiListParamsDateRange1dControl  RadarQualityIqiListParamsDateRange = "1dControl"
	RadarQualityIqiListParamsDateRange2dControl  RadarQualityIqiListParamsDateRange = "2dControl"
	RadarQualityIqiListParamsDateRange7dControl  RadarQualityIqiListParamsDateRange = "7dControl"
	RadarQualityIqiListParamsDateRange14dControl RadarQualityIqiListParamsDateRange = "14dControl"
	RadarQualityIqiListParamsDateRange28dControl RadarQualityIqiListParamsDateRange = "28dControl"
	RadarQualityIqiListParamsDateRange12wControl RadarQualityIqiListParamsDateRange = "12wControl"
	RadarQualityIqiListParamsDateRange24wControl RadarQualityIqiListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarQualityIqiListParamsFormat string

const (
	RadarQualityIqiListParamsFormatJson RadarQualityIqiListParamsFormat = "JSON"
	RadarQualityIqiListParamsFormatCsv  RadarQualityIqiListParamsFormat = "CSV"
)

type RadarQualityIqiListResponseEnvelope struct {
	Result  RadarQualityIqiListResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarQualityIqiListResponseEnvelopeJSON `json:"-"`
}

// radarQualityIqiListResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarQualityIqiListResponseEnvelope]
type radarQualityIqiListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
