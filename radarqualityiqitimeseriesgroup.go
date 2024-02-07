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

// RadarQualityIqiTimeseriesGroupService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarQualityIqiTimeseriesGroupService] method instead.
type RadarQualityIqiTimeseriesGroupService struct {
	Options []option.RequestOption
}

// NewRadarQualityIqiTimeseriesGroupService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarQualityIqiTimeseriesGroupService(opts ...option.RequestOption) (r *RadarQualityIqiTimeseriesGroupService) {
	r = &RadarQualityIqiTimeseriesGroupService{}
	r.Options = opts
	return
}

// Get a time series (percentiles) of bandwidth, latency or DNS response time from
// the Radar Internet Quality Index (IQI).
func (r *RadarQualityIqiTimeseriesGroupService) List(ctx context.Context, query RadarQualityIqiTimeseriesGroupListParams, opts ...option.RequestOption) (res *RadarQualityIqiTimeseriesGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/quality/iqi/timeseries_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarQualityIqiTimeseriesGroupListResponse struct {
	Result  RadarQualityIqiTimeseriesGroupListResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarQualityIqiTimeseriesGroupListResponseJSON   `json:"-"`
}

// radarQualityIqiTimeseriesGroupListResponseJSON contains the JSON metadata for
// the struct [RadarQualityIqiTimeseriesGroupListResponse]
type radarQualityIqiTimeseriesGroupListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiTimeseriesGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiTimeseriesGroupListResponseResult struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarQualityIqiTimeseriesGroupListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarQualityIqiTimeseriesGroupListResponseResultJSON   `json:"-"`
}

// radarQualityIqiTimeseriesGroupListResponseResultJSON contains the JSON metadata
// for the struct [RadarQualityIqiTimeseriesGroupListResponseResult]
type radarQualityIqiTimeseriesGroupListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiTimeseriesGroupListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiTimeseriesGroupListResponseResultSerie0 struct {
	P25        []string                                                   `json:"p25,required"`
	P50        []string                                                   `json:"p50,required"`
	P75        []string                                                   `json:"p75,required"`
	Timestamps []string                                                   `json:"timestamps,required"`
	JSON       radarQualityIqiTimeseriesGroupListResponseResultSerie0JSON `json:"-"`
}

// radarQualityIqiTimeseriesGroupListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarQualityIqiTimeseriesGroupListResponseResultSerie0]
type radarQualityIqiTimeseriesGroupListResponseResultSerie0JSON struct {
	P25         apijson.Field
	P50         apijson.Field
	P75         apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarQualityIqiTimeseriesGroupListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarQualityIqiTimeseriesGroupListParams struct {
	// Which metric to return: bandwidth, latency or DNS response time.
	Metric param.Field[RadarQualityIqiTimeseriesGroupListParamsMetric] `query:"metric,required"`
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarQualityIqiTimeseriesGroupListParamsAggInterval] `query:"aggInterval"`
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
	DateRange param.Field[[]RadarQualityIqiTimeseriesGroupListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarQualityIqiTimeseriesGroupListParamsFormat] `query:"format"`
	// Enable interpolation for all series (using the average).
	Interpolation param.Field[bool] `query:"interpolation"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarQualityIqiTimeseriesGroupListParams]'s query
// parameters as `url.Values`.
func (r RadarQualityIqiTimeseriesGroupListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Which metric to return: bandwidth, latency or DNS response time.
type RadarQualityIqiTimeseriesGroupListParamsMetric string

const (
	RadarQualityIqiTimeseriesGroupListParamsMetricBandwidth RadarQualityIqiTimeseriesGroupListParamsMetric = "BANDWIDTH"
	RadarQualityIqiTimeseriesGroupListParamsMetricDNS       RadarQualityIqiTimeseriesGroupListParamsMetric = "DNS"
	RadarQualityIqiTimeseriesGroupListParamsMetricLatency   RadarQualityIqiTimeseriesGroupListParamsMetric = "LATENCY"
)

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarQualityIqiTimeseriesGroupListParamsAggInterval string

const (
	RadarQualityIqiTimeseriesGroupListParamsAggInterval15m RadarQualityIqiTimeseriesGroupListParamsAggInterval = "15m"
	RadarQualityIqiTimeseriesGroupListParamsAggInterval1h  RadarQualityIqiTimeseriesGroupListParamsAggInterval = "1h"
	RadarQualityIqiTimeseriesGroupListParamsAggInterval1d  RadarQualityIqiTimeseriesGroupListParamsAggInterval = "1d"
	RadarQualityIqiTimeseriesGroupListParamsAggInterval1w  RadarQualityIqiTimeseriesGroupListParamsAggInterval = "1w"
)

type RadarQualityIqiTimeseriesGroupListParamsDateRange string

const (
	RadarQualityIqiTimeseriesGroupListParamsDateRange1d         RadarQualityIqiTimeseriesGroupListParamsDateRange = "1d"
	RadarQualityIqiTimeseriesGroupListParamsDateRange2d         RadarQualityIqiTimeseriesGroupListParamsDateRange = "2d"
	RadarQualityIqiTimeseriesGroupListParamsDateRange7d         RadarQualityIqiTimeseriesGroupListParamsDateRange = "7d"
	RadarQualityIqiTimeseriesGroupListParamsDateRange14d        RadarQualityIqiTimeseriesGroupListParamsDateRange = "14d"
	RadarQualityIqiTimeseriesGroupListParamsDateRange28d        RadarQualityIqiTimeseriesGroupListParamsDateRange = "28d"
	RadarQualityIqiTimeseriesGroupListParamsDateRange12w        RadarQualityIqiTimeseriesGroupListParamsDateRange = "12w"
	RadarQualityIqiTimeseriesGroupListParamsDateRange24w        RadarQualityIqiTimeseriesGroupListParamsDateRange = "24w"
	RadarQualityIqiTimeseriesGroupListParamsDateRange52w        RadarQualityIqiTimeseriesGroupListParamsDateRange = "52w"
	RadarQualityIqiTimeseriesGroupListParamsDateRange1dControl  RadarQualityIqiTimeseriesGroupListParamsDateRange = "1dControl"
	RadarQualityIqiTimeseriesGroupListParamsDateRange2dControl  RadarQualityIqiTimeseriesGroupListParamsDateRange = "2dControl"
	RadarQualityIqiTimeseriesGroupListParamsDateRange7dControl  RadarQualityIqiTimeseriesGroupListParamsDateRange = "7dControl"
	RadarQualityIqiTimeseriesGroupListParamsDateRange14dControl RadarQualityIqiTimeseriesGroupListParamsDateRange = "14dControl"
	RadarQualityIqiTimeseriesGroupListParamsDateRange28dControl RadarQualityIqiTimeseriesGroupListParamsDateRange = "28dControl"
	RadarQualityIqiTimeseriesGroupListParamsDateRange12wControl RadarQualityIqiTimeseriesGroupListParamsDateRange = "12wControl"
	RadarQualityIqiTimeseriesGroupListParamsDateRange24wControl RadarQualityIqiTimeseriesGroupListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarQualityIqiTimeseriesGroupListParamsFormat string

const (
	RadarQualityIqiTimeseriesGroupListParamsFormatJson RadarQualityIqiTimeseriesGroupListParamsFormat = "JSON"
	RadarQualityIqiTimeseriesGroupListParamsFormatCsv  RadarQualityIqiTimeseriesGroupListParamsFormat = "CSV"
)
