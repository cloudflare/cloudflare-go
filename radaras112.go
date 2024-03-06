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

// RadarAS112Service contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRadarAS112Service] method instead.
type RadarAS112Service struct {
	Options          []option.RequestOption
	Summary          *RadarAS112SummaryService
	TimeseriesGroups *RadarAS112TimeseriesGroupService
	Top              *RadarAS112TopService
}

// NewRadarAS112Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarAS112Service(opts ...option.RequestOption) (r *RadarAS112Service) {
	r = &RadarAS112Service{}
	r.Options = opts
	r.Summary = NewRadarAS112SummaryService(opts...)
	r.TimeseriesGroups = NewRadarAS112TimeseriesGroupService(opts...)
	r.Top = NewRadarAS112TopService(opts...)
	return
}

// Get AS112 queries change over time.
func (r *RadarAS112Service) Timeseries(ctx context.Context, query RadarAS112TimeseriesParams, opts ...option.RequestOption) (res *RadarAS112TimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAS112TimeseriesResponseEnvelope
	path := "radar/as112/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAS112TimeseriesResponse struct {
	Meta   RadarAS112TimeseriesResponseMeta   `json:"meta,required"`
	Serie0 RadarAS112TimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   radarAS112TimeseriesResponseJSON   `json:"-"`
}

// radarAS112TimeseriesResponseJSON contains the JSON metadata for the struct
// [RadarAS112TimeseriesResponse]
type radarAS112TimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TimeseriesResponseMeta struct {
	AggInterval    string                                         `json:"aggInterval,required"`
	DateRange      []RadarAS112TimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                      `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarAS112TimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAS112TimeseriesResponseMetaJSON           `json:"-"`
}

// radarAS112TimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [RadarAS112TimeseriesResponseMeta]
type radarAS112TimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAS112TimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAS112TimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// radarAS112TimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarAS112TimeseriesResponseMetaDateRange]
type radarAS112TimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TimeseriesResponseMetaConfidenceInfo struct {
	Annotations []RadarAS112TimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                      `json:"level"`
	JSON        radarAS112TimeseriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAS112TimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarAS112TimeseriesResponseMetaConfidenceInfo]
type radarAS112TimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                       `json:"dataSource,required"`
	Description     string                                                       `json:"description,required"`
	EventType       string                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                    `json:"startTime" format:"date-time"`
	JSON            radarAS112TimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAS112TimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarAS112TimeseriesResponseMetaConfidenceInfoAnnotation]
type radarAS112TimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAS112TimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarAS112TimeseriesResponseSerie0 struct {
	Timestamps []time.Time                            `json:"timestamps,required" format:"date-time"`
	Values     []string                               `json:"values,required"`
	JSON       radarAS112TimeseriesResponseSerie0JSON `json:"-"`
}

// radarAS112TimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [RadarAS112TimeseriesResponseSerie0]
type radarAS112TimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAS112TimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAS112TimeseriesParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAS112TimeseriesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAS112TimeseriesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAS112TimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarAS112TimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAS112TimeseriesParamsAggInterval string

const (
	RadarAS112TimeseriesParamsAggInterval15m RadarAS112TimeseriesParamsAggInterval = "15m"
	RadarAS112TimeseriesParamsAggInterval1h  RadarAS112TimeseriesParamsAggInterval = "1h"
	RadarAS112TimeseriesParamsAggInterval1d  RadarAS112TimeseriesParamsAggInterval = "1d"
	RadarAS112TimeseriesParamsAggInterval1w  RadarAS112TimeseriesParamsAggInterval = "1w"
)

type RadarAS112TimeseriesParamsDateRange string

const (
	RadarAS112TimeseriesParamsDateRange1d         RadarAS112TimeseriesParamsDateRange = "1d"
	RadarAS112TimeseriesParamsDateRange2d         RadarAS112TimeseriesParamsDateRange = "2d"
	RadarAS112TimeseriesParamsDateRange7d         RadarAS112TimeseriesParamsDateRange = "7d"
	RadarAS112TimeseriesParamsDateRange14d        RadarAS112TimeseriesParamsDateRange = "14d"
	RadarAS112TimeseriesParamsDateRange28d        RadarAS112TimeseriesParamsDateRange = "28d"
	RadarAS112TimeseriesParamsDateRange12w        RadarAS112TimeseriesParamsDateRange = "12w"
	RadarAS112TimeseriesParamsDateRange24w        RadarAS112TimeseriesParamsDateRange = "24w"
	RadarAS112TimeseriesParamsDateRange52w        RadarAS112TimeseriesParamsDateRange = "52w"
	RadarAS112TimeseriesParamsDateRange1dControl  RadarAS112TimeseriesParamsDateRange = "1dControl"
	RadarAS112TimeseriesParamsDateRange2dControl  RadarAS112TimeseriesParamsDateRange = "2dControl"
	RadarAS112TimeseriesParamsDateRange7dControl  RadarAS112TimeseriesParamsDateRange = "7dControl"
	RadarAS112TimeseriesParamsDateRange14dControl RadarAS112TimeseriesParamsDateRange = "14dControl"
	RadarAS112TimeseriesParamsDateRange28dControl RadarAS112TimeseriesParamsDateRange = "28dControl"
	RadarAS112TimeseriesParamsDateRange12wControl RadarAS112TimeseriesParamsDateRange = "12wControl"
	RadarAS112TimeseriesParamsDateRange24wControl RadarAS112TimeseriesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAS112TimeseriesParamsFormat string

const (
	RadarAS112TimeseriesParamsFormatJson RadarAS112TimeseriesParamsFormat = "JSON"
	RadarAS112TimeseriesParamsFormatCsv  RadarAS112TimeseriesParamsFormat = "CSV"
)

type RadarAS112TimeseriesResponseEnvelope struct {
	Result  RadarAS112TimeseriesResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarAS112TimeseriesResponseEnvelopeJSON `json:"-"`
}

// radarAS112TimeseriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAS112TimeseriesResponseEnvelope]
type radarAS112TimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAS112TimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAS112TimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
