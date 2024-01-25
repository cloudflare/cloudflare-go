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

// RadarAs112TimeseryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAs112TimeseryService] method
// instead.
type RadarAs112TimeseryService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseryService(opts ...option.RequestOption) (r *RadarAs112TimeseryService) {
	r = &RadarAs112TimeseryService{}
	r.Options = opts
	return
}

// Get AS112 queries change over time.
func (r *RadarAs112TimeseryService) List(ctx context.Context, query RadarAs112TimeseryListParams, opts ...option.RequestOption) (res *RadarAs112TimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseryListResponse struct {
	Result  RadarAs112TimeseryListResponseResult `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarAs112TimeseryListResponseJSON   `json:"-"`
}

// radarAs112TimeseryListResponseJSON contains the JSON metadata for the struct
// [RadarAs112TimeseryListResponse]
type radarAs112TimeseryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryListResponseResult struct {
	Meta   RadarAs112TimeseryListResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAs112TimeseryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseryListResponseResultJSON   `json:"-"`
}

// radarAs112TimeseryListResponseResultJSON contains the JSON metadata for the
// struct [RadarAs112TimeseryListResponseResult]
type radarAs112TimeseryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryListResponseResultMeta struct {
	AggInterval    string                                                 `json:"aggInterval,required"`
	DateRange      []RadarAs112TimeseryListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                              `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo RadarAs112TimeseryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAs112TimeseryListResponseResultMetaJSON           `json:"-"`
}

// radarAs112TimeseryListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarAs112TimeseryListResponseResultMeta]
type radarAs112TimeseryListResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAs112TimeseryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarAs112TimeseryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAs112TimeseryListResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarAs112TimeseryListResponseResultMetaDateRange]
type radarAs112TimeseryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAs112TimeseryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        radarAs112TimeseryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAs112TimeseryListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarAs112TimeseryListResponseResultMetaConfidenceInfo]
type radarAs112TimeseryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            radarAs112TimeseryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAs112TimeseryListResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarAs112TimeseryListResponseResultMetaConfidenceInfoAnnotation]
type radarAs112TimeseryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAs112TimeseryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryListResponseResultSerie0 struct {
	Timestamps []time.Time                                    `json:"timestamps,required" format:"date-time"`
	Values     []string                                       `json:"values,required"`
	JSON       radarAs112TimeseryListResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseryListResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarAs112TimeseryListResponseResultSerie0]
type radarAs112TimeseryListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseryListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseryListParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseryListParamsAggInterval string

const (
	RadarAs112TimeseryListParamsAggInterval15m RadarAs112TimeseryListParamsAggInterval = "15m"
	RadarAs112TimeseryListParamsAggInterval1h  RadarAs112TimeseryListParamsAggInterval = "1h"
	RadarAs112TimeseryListParamsAggInterval1d  RadarAs112TimeseryListParamsAggInterval = "1d"
	RadarAs112TimeseryListParamsAggInterval1w  RadarAs112TimeseryListParamsAggInterval = "1w"
)

type RadarAs112TimeseryListParamsDateRange string

const (
	RadarAs112TimeseryListParamsDateRange1d         RadarAs112TimeseryListParamsDateRange = "1d"
	RadarAs112TimeseryListParamsDateRange2d         RadarAs112TimeseryListParamsDateRange = "2d"
	RadarAs112TimeseryListParamsDateRange7d         RadarAs112TimeseryListParamsDateRange = "7d"
	RadarAs112TimeseryListParamsDateRange14d        RadarAs112TimeseryListParamsDateRange = "14d"
	RadarAs112TimeseryListParamsDateRange28d        RadarAs112TimeseryListParamsDateRange = "28d"
	RadarAs112TimeseryListParamsDateRange12w        RadarAs112TimeseryListParamsDateRange = "12w"
	RadarAs112TimeseryListParamsDateRange24w        RadarAs112TimeseryListParamsDateRange = "24w"
	RadarAs112TimeseryListParamsDateRange52w        RadarAs112TimeseryListParamsDateRange = "52w"
	RadarAs112TimeseryListParamsDateRange1dControl  RadarAs112TimeseryListParamsDateRange = "1dControl"
	RadarAs112TimeseryListParamsDateRange2dControl  RadarAs112TimeseryListParamsDateRange = "2dControl"
	RadarAs112TimeseryListParamsDateRange7dControl  RadarAs112TimeseryListParamsDateRange = "7dControl"
	RadarAs112TimeseryListParamsDateRange14dControl RadarAs112TimeseryListParamsDateRange = "14dControl"
	RadarAs112TimeseryListParamsDateRange28dControl RadarAs112TimeseryListParamsDateRange = "28dControl"
	RadarAs112TimeseryListParamsDateRange12wControl RadarAs112TimeseryListParamsDateRange = "12wControl"
	RadarAs112TimeseryListParamsDateRange24wControl RadarAs112TimeseryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseryListParamsFormat string

const (
	RadarAs112TimeseryListParamsFormatJson RadarAs112TimeseryListParamsFormat = "JSON"
	RadarAs112TimeseryListParamsFormatCsv  RadarAs112TimeseryListParamsFormat = "CSV"
)
