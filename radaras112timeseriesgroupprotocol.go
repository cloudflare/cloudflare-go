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

// RadarAs112TimeseriesGroupProtocolService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAs112TimeseriesGroupProtocolService] method instead.
type RadarAs112TimeseriesGroupProtocolService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseriesGroupProtocolService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseriesGroupProtocolService(opts ...option.RequestOption) (r *RadarAs112TimeseriesGroupProtocolService) {
	r = &RadarAs112TimeseriesGroupProtocolService{}
	r.Options = opts
	return
}

// Percentage distribution of AS112 dns requests classified per Protocol over time.
func (r *RadarAs112TimeseriesGroupProtocolService) Get(ctx context.Context, query RadarAs112TimeseriesGroupProtocolGetParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupProtocolGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseriesGroupProtocolGetResponse struct {
	Result  RadarAs112TimeseriesGroupProtocolGetResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarAs112TimeseriesGroupProtocolGetResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupProtocolGetResponseJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupProtocolGetResponse]
type radarAs112TimeseriesGroupProtocolGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupProtocolGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupProtocolGetResponseResult struct {
	Meta   interface{}                                              `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupProtocolGetResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupProtocolGetResponseResultJSON   `json:"-"`
}

// radarAs112TimeseriesGroupProtocolGetResponseResultJSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupProtocolGetResponseResult]
type radarAs112TimeseriesGroupProtocolGetResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupProtocolGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupProtocolGetResponseResultSerie0 struct {
	Tcp  []string                                                     `json:"tcp,required"`
	Udp  []string                                                     `json:"udp,required"`
	JSON radarAs112TimeseriesGroupProtocolGetResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupProtocolGetResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarAs112TimeseriesGroupProtocolGetResponseResultSerie0]
type radarAs112TimeseriesGroupProtocolGetResponseResultSerie0JSON struct {
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupProtocolGetResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupProtocolGetParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupProtocolGetParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupProtocolGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupProtocolGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupProtocolGetParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupProtocolGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupProtocolGetParamsAggInterval string

const (
	RadarAs112TimeseriesGroupProtocolGetParamsAggInterval15m RadarAs112TimeseriesGroupProtocolGetParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupProtocolGetParamsAggInterval1h  RadarAs112TimeseriesGroupProtocolGetParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupProtocolGetParamsAggInterval1d  RadarAs112TimeseriesGroupProtocolGetParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupProtocolGetParamsAggInterval1w  RadarAs112TimeseriesGroupProtocolGetParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupProtocolGetParamsDateRange string

const (
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange1d         RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "1d"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange2d         RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "2d"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange7d         RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "7d"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange14d        RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "14d"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange28d        RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "28d"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange12w        RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "12w"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange24w        RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "24w"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange52w        RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "52w"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange1dControl  RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange2dControl  RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange7dControl  RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange14dControl RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange28dControl RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange12wControl RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupProtocolGetParamsDateRange24wControl RadarAs112TimeseriesGroupProtocolGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupProtocolGetParamsFormat string

const (
	RadarAs112TimeseriesGroupProtocolGetParamsFormatJson RadarAs112TimeseriesGroupProtocolGetParamsFormat = "JSON"
	RadarAs112TimeseriesGroupProtocolGetParamsFormatCsv  RadarAs112TimeseriesGroupProtocolGetParamsFormat = "CSV"
)
