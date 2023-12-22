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

// RadarAttackLayer3TimeseriesGroupService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer3TimeseriesGroupService] method instead.
type RadarAttackLayer3TimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TimeseriesGroupService(opts ...option.RequestOption) (r *RadarAttackLayer3TimeseriesGroupService) {
	r = &RadarAttackLayer3TimeseriesGroupService{}
	r.Options = opts
	return
}

// Get percentage of what type of network protocols are used in layer 3/4 attacks,
// over time.
func (r *RadarAttackLayer3TimeseriesGroupService) List(ctx context.Context, query RadarAttackLayer3TimeseriesGroupListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3TimeseriesGroupListResponse struct {
	Result  RadarAttackLayer3TimeseriesGroupListResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupListResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupListResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TimeseriesGroupListResponse]
type radarAttackLayer3TimeseriesGroupListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupListResponseResult struct {
	Meta   RadarAttackLayer3TimeseriesGroupListResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupListResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupListResponseResult]
type radarAttackLayer3TimeseriesGroupListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupListResponseResultMeta struct {
	AggInterval    string                                                               `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarAttackLayer3TimeseriesGroupListResponseResultMetaDateRange      `json:"dateRange,required"`
	LastUpdated    time.Time                                                            `json:"lastUpdated,required" format:"date-time"`
	JSON           radarAttackLayer3TimeseriesGroupListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3TimeseriesGroupListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupListResponseResultMeta]
type radarAttackLayer3TimeseriesGroupListResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                            `json:"level,required"`
	JSON        radarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfo]
type radarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                             `json:"dataSource,required"`
	Description string                                                                             `json:"description,required"`
	EndTime     time.Time                                                                          `json:"endTime,required" format:"date-time"`
	EventType   string                                                                             `json:"eventType,required"`
	LinkedURL   string                                                                             `json:"linkedUrl,required"`
	StartTime   time.Time                                                                          `json:"startTime,required" format:"date-time"`
	JSON        radarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	LinkedURL   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                           `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TimeseriesGroupListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupListResponseResultMetaDateRange]
type radarAttackLayer3TimeseriesGroupListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupListResponseResultSerie0 struct {
	Gre        []string                                                     `json:"gre,required"`
	Icmp       []string                                                     `json:"icmp,required"`
	Tcp        []string                                                     `json:"tcp,required"`
	Timestamps []string                                                     `json:"timestamps,required"`
	Udp        []string                                                     `json:"udp,required"`
	JSON       radarAttackLayer3TimeseriesGroupListResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupListResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupListResponseResultSerie0]
type radarAttackLayer3TimeseriesGroupListResponseResultSerie0JSON struct {
	Gre         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Timestamps  apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupListParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupListParamsAggInterval15m RadarAttackLayer3TimeseriesGroupListParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupListParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupListParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupListParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupListParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupListParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupListParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupListParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupListParamsDateRange1d         RadarAttackLayer3TimeseriesGroupListParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupListParamsDateRange7d         RadarAttackLayer3TimeseriesGroupListParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupListParamsDateRange14d        RadarAttackLayer3TimeseriesGroupListParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupListParamsDateRange28d        RadarAttackLayer3TimeseriesGroupListParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupListParamsDateRange12w        RadarAttackLayer3TimeseriesGroupListParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupListParamsDateRange24w        RadarAttackLayer3TimeseriesGroupListParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupListParamsDateRange52w        RadarAttackLayer3TimeseriesGroupListParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupListParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupListParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupListParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupListParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupListParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupListParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupListParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupListParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupListParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupListParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupListParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupListParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupListParamsFormatJson RadarAttackLayer3TimeseriesGroupListParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupListParamsFormatCsv  RadarAttackLayer3TimeseriesGroupListParamsFormat = "CSV"
)
