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

// RadarAs112Service contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRadarAs112Service] method instead.
type RadarAs112Service struct {
	Options          []option.RequestOption
	Summary          *RadarAs112SummaryService
	TimeseriesGroups *RadarAs112TimeseriesGroupService
	Top              *RadarAs112TopService
}

// NewRadarAs112Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarAs112Service(opts ...option.RequestOption) (r *RadarAs112Service) {
	r = &RadarAs112Service{}
	r.Options = opts
	r.Summary = NewRadarAs112SummaryService(opts...)
	r.TimeseriesGroups = NewRadarAs112TimeseriesGroupService(opts...)
	r.Top = NewRadarAs112TopService(opts...)
	return
}

// Get AS112 queries change over time.
func (r *RadarAs112Service) Timeseries(ctx context.Context, query RadarAs112TimeseriesParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TimeseriesResponseEnvelope
	path := "radar/as112/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAs112TimeseriesResponse struct {
	Meta   RadarAs112TimeseriesResponseMeta   `json:"meta,required"`
	Serie0 RadarAs112TimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesResponseJSON   `json:"-"`
}

// radarAs112TimeseriesResponseJSON contains the JSON metadata for the struct
// [RadarAs112TimeseriesResponse]
type radarAs112TimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesResponseMeta struct {
	AggInterval    string                                         `json:"aggInterval,required"`
	DateRange      []RadarAs112TimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                      `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarAs112TimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112TimeseriesResponseMetaJSON           `json:"-"`
}

// radarAs112TimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [RadarAs112TimeseriesResponseMeta]
type radarAs112TimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// radarAs112TimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarAs112TimeseriesResponseMetaDateRange]
type radarAs112TimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesResponseMetaConfidenceInfo struct {
	Annotations []RadarAs112TimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                      `json:"level"`
	JSON        radarAs112TimeseriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesResponseMetaConfidenceInfo]
type radarAs112TimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                       `json:"dataSource,required"`
	Description     string                                                       `json:"description,required"`
	EventType       string                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                    `json:"startTime" format:"date-time"`
	JSON            radarAs112TimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarAs112TimeseriesResponseMetaConfidenceInfoAnnotation]
type radarAs112TimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112TimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesResponseSerie0 struct {
	Timestamps []time.Time                            `json:"timestamps,required" format:"date-time"`
	Values     []string                               `json:"values,required"`
	JSON       radarAs112TimeseriesResponseSerie0JSON `json:"-"`
}

// radarAs112TimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [RadarAs112TimeseriesResponseSerie0]
type radarAs112TimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarAs112TimeseriesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesParamsAggInterval string

const (
	RadarAs112TimeseriesParamsAggInterval15m RadarAs112TimeseriesParamsAggInterval = "15m"
	RadarAs112TimeseriesParamsAggInterval1h  RadarAs112TimeseriesParamsAggInterval = "1h"
	RadarAs112TimeseriesParamsAggInterval1d  RadarAs112TimeseriesParamsAggInterval = "1d"
	RadarAs112TimeseriesParamsAggInterval1w  RadarAs112TimeseriesParamsAggInterval = "1w"
)

type RadarAs112TimeseriesParamsDateRange string

const (
	RadarAs112TimeseriesParamsDateRange1d         RadarAs112TimeseriesParamsDateRange = "1d"
	RadarAs112TimeseriesParamsDateRange2d         RadarAs112TimeseriesParamsDateRange = "2d"
	RadarAs112TimeseriesParamsDateRange7d         RadarAs112TimeseriesParamsDateRange = "7d"
	RadarAs112TimeseriesParamsDateRange14d        RadarAs112TimeseriesParamsDateRange = "14d"
	RadarAs112TimeseriesParamsDateRange28d        RadarAs112TimeseriesParamsDateRange = "28d"
	RadarAs112TimeseriesParamsDateRange12w        RadarAs112TimeseriesParamsDateRange = "12w"
	RadarAs112TimeseriesParamsDateRange24w        RadarAs112TimeseriesParamsDateRange = "24w"
	RadarAs112TimeseriesParamsDateRange52w        RadarAs112TimeseriesParamsDateRange = "52w"
	RadarAs112TimeseriesParamsDateRange1dControl  RadarAs112TimeseriesParamsDateRange = "1dControl"
	RadarAs112TimeseriesParamsDateRange2dControl  RadarAs112TimeseriesParamsDateRange = "2dControl"
	RadarAs112TimeseriesParamsDateRange7dControl  RadarAs112TimeseriesParamsDateRange = "7dControl"
	RadarAs112TimeseriesParamsDateRange14dControl RadarAs112TimeseriesParamsDateRange = "14dControl"
	RadarAs112TimeseriesParamsDateRange28dControl RadarAs112TimeseriesParamsDateRange = "28dControl"
	RadarAs112TimeseriesParamsDateRange12wControl RadarAs112TimeseriesParamsDateRange = "12wControl"
	RadarAs112TimeseriesParamsDateRange24wControl RadarAs112TimeseriesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesParamsFormat string

const (
	RadarAs112TimeseriesParamsFormatJson RadarAs112TimeseriesParamsFormat = "JSON"
	RadarAs112TimeseriesParamsFormatCsv  RadarAs112TimeseriesParamsFormat = "CSV"
)

type RadarAs112TimeseriesResponseEnvelope struct {
	Result  RadarAs112TimeseriesResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarAs112TimeseriesResponseEnvelopeJSON `json:"-"`
}

// radarAs112TimeseriesResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarAs112TimeseriesResponseEnvelope]
type radarAs112TimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
