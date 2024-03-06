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

// RadarNetflowService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarNetflowService] method
// instead.
type RadarNetflowService struct {
	Options []option.RequestOption
	Top     *RadarNetflowTopService
}

// NewRadarNetflowService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarNetflowService(opts ...option.RequestOption) (r *RadarNetflowService) {
	r = &RadarNetflowService{}
	r.Options = opts
	r.Top = NewRadarNetflowTopService(opts...)
	return
}

// Get network traffic change over time. Visit
// https://en.wikipedia.org/wiki/NetFlow for more information on NetFlows.
func (r *RadarNetflowService) Timeseries(ctx context.Context, query RadarNetflowTimeseriesParams, opts ...option.RequestOption) (res *RadarNetflowTimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarNetflowTimeseriesResponseEnvelope
	path := "radar/netflows/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarNetflowTimeseriesResponse struct {
	Meta   RadarNetflowTimeseriesResponseMeta   `json:"meta,required"`
	Serie0 RadarNetflowTimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   radarNetflowTimeseriesResponseJSON   `json:"-"`
}

// radarNetflowTimeseriesResponseJSON contains the JSON metadata for the struct
// [RadarNetflowTimeseriesResponse]
type radarNetflowTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTimeseriesResponseMeta struct {
	AggInterval    string                                           `json:"aggInterval,required"`
	DateRange      []RadarNetflowTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                        `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarNetflowTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarNetflowTimeseriesResponseMetaJSON           `json:"-"`
}

// radarNetflowTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [RadarNetflowTimeseriesResponseMeta]
type radarNetflowTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarNetflowTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      radarNetflowTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// radarNetflowTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarNetflowTimeseriesResponseMetaDateRange]
type radarNetflowTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []RadarNetflowTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        radarNetflowTimeseriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarNetflowTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarNetflowTimeseriesResponseMetaConfidenceInfo]
type radarNetflowTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            radarNetflowTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarNetflowTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarNetflowTimeseriesResponseMetaConfidenceInfoAnnotation]
type radarNetflowTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarNetflowTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTimeseriesResponseSerie0 struct {
	Timestamps []time.Time                              `json:"timestamps,required" format:"date-time"`
	Values     []string                                 `json:"values,required"`
	JSON       radarNetflowTimeseriesResponseSerie0JSON `json:"-"`
}

// radarNetflowTimeseriesResponseSerie0JSON contains the JSON metadata for the
// struct [RadarNetflowTimeseriesResponseSerie0]
type radarNetflowTimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarNetflowTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarNetflowTimeseriesParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarNetflowTimeseriesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarNetflowTimeseriesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarNetflowTimeseriesParamsNormalization] `query:"normalization"`
	// Array of network traffic product types.
	Product param.Field[[]RadarNetflowTimeseriesParamsProduct] `query:"product"`
}

// URLQuery serializes [RadarNetflowTimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarNetflowTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarNetflowTimeseriesParamsAggInterval string

const (
	RadarNetflowTimeseriesParamsAggInterval15m RadarNetflowTimeseriesParamsAggInterval = "15m"
	RadarNetflowTimeseriesParamsAggInterval1h  RadarNetflowTimeseriesParamsAggInterval = "1h"
	RadarNetflowTimeseriesParamsAggInterval1d  RadarNetflowTimeseriesParamsAggInterval = "1d"
	RadarNetflowTimeseriesParamsAggInterval1w  RadarNetflowTimeseriesParamsAggInterval = "1w"
)

type RadarNetflowTimeseriesParamsDateRange string

const (
	RadarNetflowTimeseriesParamsDateRange1d         RadarNetflowTimeseriesParamsDateRange = "1d"
	RadarNetflowTimeseriesParamsDateRange2d         RadarNetflowTimeseriesParamsDateRange = "2d"
	RadarNetflowTimeseriesParamsDateRange7d         RadarNetflowTimeseriesParamsDateRange = "7d"
	RadarNetflowTimeseriesParamsDateRange14d        RadarNetflowTimeseriesParamsDateRange = "14d"
	RadarNetflowTimeseriesParamsDateRange28d        RadarNetflowTimeseriesParamsDateRange = "28d"
	RadarNetflowTimeseriesParamsDateRange12w        RadarNetflowTimeseriesParamsDateRange = "12w"
	RadarNetflowTimeseriesParamsDateRange24w        RadarNetflowTimeseriesParamsDateRange = "24w"
	RadarNetflowTimeseriesParamsDateRange52w        RadarNetflowTimeseriesParamsDateRange = "52w"
	RadarNetflowTimeseriesParamsDateRange1dControl  RadarNetflowTimeseriesParamsDateRange = "1dControl"
	RadarNetflowTimeseriesParamsDateRange2dControl  RadarNetflowTimeseriesParamsDateRange = "2dControl"
	RadarNetflowTimeseriesParamsDateRange7dControl  RadarNetflowTimeseriesParamsDateRange = "7dControl"
	RadarNetflowTimeseriesParamsDateRange14dControl RadarNetflowTimeseriesParamsDateRange = "14dControl"
	RadarNetflowTimeseriesParamsDateRange28dControl RadarNetflowTimeseriesParamsDateRange = "28dControl"
	RadarNetflowTimeseriesParamsDateRange12wControl RadarNetflowTimeseriesParamsDateRange = "12wControl"
	RadarNetflowTimeseriesParamsDateRange24wControl RadarNetflowTimeseriesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarNetflowTimeseriesParamsFormat string

const (
	RadarNetflowTimeseriesParamsFormatJson RadarNetflowTimeseriesParamsFormat = "JSON"
	RadarNetflowTimeseriesParamsFormatCsv  RadarNetflowTimeseriesParamsFormat = "CSV"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarNetflowTimeseriesParamsNormalization string

const (
	RadarNetflowTimeseriesParamsNormalizationPercentageChange RadarNetflowTimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	RadarNetflowTimeseriesParamsNormalizationMin0Max          RadarNetflowTimeseriesParamsNormalization = "MIN0_MAX"
)

type RadarNetflowTimeseriesParamsProduct string

const (
	RadarNetflowTimeseriesParamsProductHTTP RadarNetflowTimeseriesParamsProduct = "HTTP"
	RadarNetflowTimeseriesParamsProductAll  RadarNetflowTimeseriesParamsProduct = "ALL"
)

type RadarNetflowTimeseriesResponseEnvelope struct {
	Result  RadarNetflowTimeseriesResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarNetflowTimeseriesResponseEnvelopeJSON `json:"-"`
}

// radarNetflowTimeseriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarNetflowTimeseriesResponseEnvelope]
type radarNetflowTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarNetflowTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
