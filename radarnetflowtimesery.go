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

// RadarNetflowTimeseryService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarNetflowTimeseryService]
// method instead.
type RadarNetflowTimeseryService struct {
	Options []option.RequestOption
}

// NewRadarNetflowTimeseryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarNetflowTimeseryService(opts ...option.RequestOption) (r *RadarNetflowTimeseryService) {
	r = &RadarNetflowTimeseryService{}
	r.Options = opts
	return
}

// Get network traffic change over time. Values are normalized using min-max by
// default, with the minimum set to 0. When asking for multiple time series, you
// can also get the percentage relative change of the first/main series, with
// respect to the second/control series - for example, to get the relative change
// of this week from the previous week, the first series would have a date range of
// `7d`, the second, a date range of `7dControl`, and the normalization would be
// set to `PERCENTAGE_CHANGE`.
func (r *RadarNetflowTimeseryService) List(ctx context.Context, query RadarNetflowTimeseryListParams, opts ...option.RequestOption) (res *RadarNetflowTimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/netflows/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarNetflowTimeseryListResponse struct {
	Result  RadarNetflowTimeseryListResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarNetflowTimeseryListResponseJSON   `json:"-"`
}

// radarNetflowTimeseryListResponseJSON contains the JSON metadata for the struct
// [RadarNetflowTimeseryListResponse]
type radarNetflowTimeseryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarNetflowTimeseryListResponseResult struct {
	Meta   RadarNetflowTimeseryListResponseResultMeta   `json:"meta,required"`
	Serie0 RadarNetflowTimeseryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarNetflowTimeseryListResponseResultJSON   `json:"-"`
}

// radarNetflowTimeseryListResponseResultJSON contains the JSON metadata for the
// struct [RadarNetflowTimeseryListResponseResult]
type radarNetflowTimeseryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTimeseryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarNetflowTimeseryListResponseResultMeta struct {
	AggInterval    string                                                   `json:"aggInterval,required"`
	ConfidenceInfo RadarNetflowTimeseryListResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarNetflowTimeseryListResponseResultMetaDateRange      `json:"dateRange,required"`
	LastUpdated    time.Time                                                `json:"lastUpdated,required" format:"date-time"`
	JSON           radarNetflowTimeseryListResponseResultMetaJSON           `json:"-"`
}

// radarNetflowTimeseryListResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarNetflowTimeseryListResponseResultMeta]
type radarNetflowTimeseryListResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarNetflowTimeseryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarNetflowTimeseryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarNetflowTimeseryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                `json:"level,required"`
	JSON        radarNetflowTimeseryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarNetflowTimeseryListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarNetflowTimeseryListResponseResultMetaConfidenceInfo]
type radarNetflowTimeseryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTimeseryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarNetflowTimeseryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                 `json:"dataSource,required"`
	Description string                                                                 `json:"description,required"`
	EndTime     time.Time                                                              `json:"endTime,required" format:"date-time"`
	EventType   string                                                                 `json:"eventType,required"`
	LinkedURL   string                                                                 `json:"linkedUrl,required"`
	StartTime   time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON        radarNetflowTimeseryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarNetflowTimeseryListResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarNetflowTimeseryListResponseResultMetaConfidenceInfoAnnotation]
type radarNetflowTimeseryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	LinkedURL   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTimeseryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarNetflowTimeseryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarNetflowTimeseryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarNetflowTimeseryListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarNetflowTimeseryListResponseResultMetaDateRange]
type radarNetflowTimeseryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTimeseryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarNetflowTimeseryListResponseResultSerie0 struct {
	Timestamps []time.Time                                      `json:"timestamps,required" format:"date-time"`
	Values     []string                                         `json:"values,required"`
	JSON       radarNetflowTimeseryListResponseResultSerie0JSON `json:"-"`
}

// radarNetflowTimeseryListResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarNetflowTimeseryListResponseResultSerie0]
type radarNetflowTimeseryListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarNetflowTimeseryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarNetflowTimeseryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarNetflowTimeseryListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarNetflowTimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarNetflowTimeseryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarNetflowTimeseryListParamsNormalization] `query:"normalization"`
	// Array of network traffic product types.
	Product param.Field[[]RadarNetflowTimeseryListParamsProduct] `query:"product"`
}

// URLQuery serializes [RadarNetflowTimeseryListParams]'s query parameters as
// `url.Values`.
func (r RadarNetflowTimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarNetflowTimeseryListParamsAggInterval string

const (
	RadarNetflowTimeseryListParamsAggInterval15m RadarNetflowTimeseryListParamsAggInterval = "15m"
	RadarNetflowTimeseryListParamsAggInterval1h  RadarNetflowTimeseryListParamsAggInterval = "1h"
	RadarNetflowTimeseryListParamsAggInterval1d  RadarNetflowTimeseryListParamsAggInterval = "1d"
	RadarNetflowTimeseryListParamsAggInterval1w  RadarNetflowTimeseryListParamsAggInterval = "1w"
)

type RadarNetflowTimeseryListParamsDateRange string

const (
	RadarNetflowTimeseryListParamsDateRange1d         RadarNetflowTimeseryListParamsDateRange = "1d"
	RadarNetflowTimeseryListParamsDateRange7d         RadarNetflowTimeseryListParamsDateRange = "7d"
	RadarNetflowTimeseryListParamsDateRange14d        RadarNetflowTimeseryListParamsDateRange = "14d"
	RadarNetflowTimeseryListParamsDateRange28d        RadarNetflowTimeseryListParamsDateRange = "28d"
	RadarNetflowTimeseryListParamsDateRange12w        RadarNetflowTimeseryListParamsDateRange = "12w"
	RadarNetflowTimeseryListParamsDateRange24w        RadarNetflowTimeseryListParamsDateRange = "24w"
	RadarNetflowTimeseryListParamsDateRange52w        RadarNetflowTimeseryListParamsDateRange = "52w"
	RadarNetflowTimeseryListParamsDateRange1dControl  RadarNetflowTimeseryListParamsDateRange = "1dControl"
	RadarNetflowTimeseryListParamsDateRange7dControl  RadarNetflowTimeseryListParamsDateRange = "7dControl"
	RadarNetflowTimeseryListParamsDateRange14dControl RadarNetflowTimeseryListParamsDateRange = "14dControl"
	RadarNetflowTimeseryListParamsDateRange28dControl RadarNetflowTimeseryListParamsDateRange = "28dControl"
	RadarNetflowTimeseryListParamsDateRange12wControl RadarNetflowTimeseryListParamsDateRange = "12wControl"
	RadarNetflowTimeseryListParamsDateRange24wControl RadarNetflowTimeseryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarNetflowTimeseryListParamsFormat string

const (
	RadarNetflowTimeseryListParamsFormatJson RadarNetflowTimeseryListParamsFormat = "JSON"
	RadarNetflowTimeseryListParamsFormatCsv  RadarNetflowTimeseryListParamsFormat = "CSV"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarNetflowTimeseryListParamsNormalization string

const (
	RadarNetflowTimeseryListParamsNormalizationPercentageChange RadarNetflowTimeseryListParamsNormalization = "PERCENTAGE_CHANGE"
	RadarNetflowTimeseryListParamsNormalizationMin0Max          RadarNetflowTimeseryListParamsNormalization = "MIN0_MAX"
)

type RadarNetflowTimeseryListParamsProduct string

const (
	RadarNetflowTimeseryListParamsProductHTTP RadarNetflowTimeseryListParamsProduct = "HTTP"
	RadarNetflowTimeseryListParamsProductAll  RadarNetflowTimeseryListParamsProduct = "ALL"
)
