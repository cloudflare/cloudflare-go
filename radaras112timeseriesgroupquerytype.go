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

// RadarAs112TimeseriesGroupQueryTypeService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAs112TimeseriesGroupQueryTypeService] method instead.
type RadarAs112TimeseriesGroupQueryTypeService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseriesGroupQueryTypeService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAs112TimeseriesGroupQueryTypeService(opts ...option.RequestOption) (r *RadarAs112TimeseriesGroupQueryTypeService) {
	r = &RadarAs112TimeseriesGroupQueryTypeService{}
	r.Options = opts
	return
}

// Percentage distribution of AS112 DNS queries by Query Type over time.
func (r *RadarAs112TimeseriesGroupQueryTypeService) Get(ctx context.Context, query RadarAs112TimeseriesGroupQueryTypeGetParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupQueryTypeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseriesGroupQueryTypeGetResponse struct {
	Result  RadarAs112TimeseriesGroupQueryTypeGetResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarAs112TimeseriesGroupQueryTypeGetResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupQueryTypeGetResponseJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupQueryTypeGetResponse]
type radarAs112TimeseriesGroupQueryTypeGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupQueryTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupQueryTypeGetResponseResult struct {
	Meta   interface{}                                               `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupQueryTypeGetResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupQueryTypeGetResponseResultJSON   `json:"-"`
}

// radarAs112TimeseriesGroupQueryTypeGetResponseResultJSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupQueryTypeGetResponseResult]
type radarAs112TimeseriesGroupQueryTypeGetResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupQueryTypeGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupQueryTypeGetResponseResultSerie0 struct {
	A    []string                                                      `json:"A,required"`
	Aaaa []string                                                      `json:"AAAA,required"`
	Ptr  []string                                                      `json:"PTR,required"`
	Soa  []string                                                      `json:"SOA,required"`
	Srv  []string                                                      `json:"SRV,required"`
	JSON radarAs112TimeseriesGroupQueryTypeGetResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupQueryTypeGetResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarAs112TimeseriesGroupQueryTypeGetResponseResultSerie0]
type radarAs112TimeseriesGroupQueryTypeGetResponseResultSerie0JSON struct {
	A           apijson.Field
	Aaaa        apijson.Field
	Ptr         apijson.Field
	Soa         apijson.Field
	Srv         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupQueryTypeGetResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupQueryTypeGetParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupQueryTypeGetParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupQueryTypeGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupQueryTypeGetParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupQueryTypeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupQueryTypeGetParamsAggInterval string

const (
	RadarAs112TimeseriesGroupQueryTypeGetParamsAggInterval15m RadarAs112TimeseriesGroupQueryTypeGetParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupQueryTypeGetParamsAggInterval1h  RadarAs112TimeseriesGroupQueryTypeGetParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupQueryTypeGetParamsAggInterval1d  RadarAs112TimeseriesGroupQueryTypeGetParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupQueryTypeGetParamsAggInterval1w  RadarAs112TimeseriesGroupQueryTypeGetParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange string

const (
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange1d         RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "1d"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange2d         RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "2d"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange7d         RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "7d"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange14d        RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "14d"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange28d        RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "28d"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange12w        RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "12w"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange24w        RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "24w"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange52w        RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "52w"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange1dControl  RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange2dControl  RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange7dControl  RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange14dControl RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange28dControl RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange12wControl RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange24wControl RadarAs112TimeseriesGroupQueryTypeGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupQueryTypeGetParamsFormat string

const (
	RadarAs112TimeseriesGroupQueryTypeGetParamsFormatJson RadarAs112TimeseriesGroupQueryTypeGetParamsFormat = "JSON"
	RadarAs112TimeseriesGroupQueryTypeGetParamsFormatCsv  RadarAs112TimeseriesGroupQueryTypeGetParamsFormat = "CSV"
)
