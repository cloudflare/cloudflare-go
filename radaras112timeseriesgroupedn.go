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

// RadarAs112TimeseriesGroupEdnService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAs112TimeseriesGroupEdnService] method instead.
type RadarAs112TimeseriesGroupEdnService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseriesGroupEdnService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseriesGroupEdnService(opts ...option.RequestOption) (r *RadarAs112TimeseriesGroupEdnService) {
	r = &RadarAs112TimeseriesGroupEdnService{}
	r.Options = opts
	return
}

// Percentage distribution of AS112 DNS queries by EDNS support over time.
func (r *RadarAs112TimeseriesGroupEdnService) Get(ctx context.Context, query RadarAs112TimeseriesGroupEdnGetParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupEdnGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseriesGroupEdnGetResponse struct {
	Result  RadarAs112TimeseriesGroupEdnGetResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarAs112TimeseriesGroupEdnGetResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupEdnGetResponseJSON contains the JSON metadata for the
// struct [RadarAs112TimeseriesGroupEdnGetResponse]
type radarAs112TimeseriesGroupEdnGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupEdnGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupEdnGetResponseResult struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupEdnGetResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupEdnGetResponseResultJSON   `json:"-"`
}

// radarAs112TimeseriesGroupEdnGetResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupEdnGetResponseResult]
type radarAs112TimeseriesGroupEdnGetResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupEdnGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupEdnGetResponseResultSerie0 struct {
	NotSupported []string                                                `json:"NOT_SUPPORTED,required"`
	Supported    []string                                                `json:"SUPPORTED,required"`
	JSON         radarAs112TimeseriesGroupEdnGetResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupEdnGetResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupEdnGetResponseResultSerie0]
type radarAs112TimeseriesGroupEdnGetResponseResultSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupEdnGetResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupEdnGetParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupEdnGetParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupEdnGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupEdnGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupEdnGetParams]'s query parameters
// as `url.Values`.
func (r RadarAs112TimeseriesGroupEdnGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupEdnGetParamsAggInterval string

const (
	RadarAs112TimeseriesGroupEdnGetParamsAggInterval15m RadarAs112TimeseriesGroupEdnGetParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupEdnGetParamsAggInterval1h  RadarAs112TimeseriesGroupEdnGetParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupEdnGetParamsAggInterval1d  RadarAs112TimeseriesGroupEdnGetParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupEdnGetParamsAggInterval1w  RadarAs112TimeseriesGroupEdnGetParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupEdnGetParamsDateRange string

const (
	RadarAs112TimeseriesGroupEdnGetParamsDateRange1d         RadarAs112TimeseriesGroupEdnGetParamsDateRange = "1d"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange2d         RadarAs112TimeseriesGroupEdnGetParamsDateRange = "2d"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange7d         RadarAs112TimeseriesGroupEdnGetParamsDateRange = "7d"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange14d        RadarAs112TimeseriesGroupEdnGetParamsDateRange = "14d"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange28d        RadarAs112TimeseriesGroupEdnGetParamsDateRange = "28d"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange12w        RadarAs112TimeseriesGroupEdnGetParamsDateRange = "12w"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange24w        RadarAs112TimeseriesGroupEdnGetParamsDateRange = "24w"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange52w        RadarAs112TimeseriesGroupEdnGetParamsDateRange = "52w"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange1dControl  RadarAs112TimeseriesGroupEdnGetParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange2dControl  RadarAs112TimeseriesGroupEdnGetParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange7dControl  RadarAs112TimeseriesGroupEdnGetParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange14dControl RadarAs112TimeseriesGroupEdnGetParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange28dControl RadarAs112TimeseriesGroupEdnGetParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange12wControl RadarAs112TimeseriesGroupEdnGetParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupEdnGetParamsDateRange24wControl RadarAs112TimeseriesGroupEdnGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupEdnGetParamsFormat string

const (
	RadarAs112TimeseriesGroupEdnGetParamsFormatJson RadarAs112TimeseriesGroupEdnGetParamsFormat = "JSON"
	RadarAs112TimeseriesGroupEdnGetParamsFormatCsv  RadarAs112TimeseriesGroupEdnGetParamsFormat = "CSV"
)
