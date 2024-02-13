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
func (r *RadarAs112TimeseriesGroupEdnService) List(ctx context.Context, query RadarAs112TimeseriesGroupEdnListParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupEdnListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TimeseriesGroupEdnListResponseEnvelope
	path := "radar/as112/timeseries_groups/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAs112TimeseriesGroupEdnListResponse struct {
	Meta   interface{}                                    `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupEdnListResponseSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupEdnListResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupEdnListResponseJSON contains the JSON metadata for the
// struct [RadarAs112TimeseriesGroupEdnListResponse]
type radarAs112TimeseriesGroupEdnListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupEdnListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupEdnListResponseSerie0 struct {
	NotSupported []string                                           `json:"NOT_SUPPORTED,required"`
	Supported    []string                                           `json:"SUPPORTED,required"`
	JSON         radarAs112TimeseriesGroupEdnListResponseSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupEdnListResponseSerie0JSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupEdnListResponseSerie0]
type radarAs112TimeseriesGroupEdnListResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupEdnListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupEdnListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupEdnListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupEdnListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupEdnListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupEdnListParams]'s query parameters
// as `url.Values`.
func (r RadarAs112TimeseriesGroupEdnListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupEdnListParamsAggInterval string

const (
	RadarAs112TimeseriesGroupEdnListParamsAggInterval15m RadarAs112TimeseriesGroupEdnListParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupEdnListParamsAggInterval1h  RadarAs112TimeseriesGroupEdnListParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupEdnListParamsAggInterval1d  RadarAs112TimeseriesGroupEdnListParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupEdnListParamsAggInterval1w  RadarAs112TimeseriesGroupEdnListParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupEdnListParamsDateRange string

const (
	RadarAs112TimeseriesGroupEdnListParamsDateRange1d         RadarAs112TimeseriesGroupEdnListParamsDateRange = "1d"
	RadarAs112TimeseriesGroupEdnListParamsDateRange2d         RadarAs112TimeseriesGroupEdnListParamsDateRange = "2d"
	RadarAs112TimeseriesGroupEdnListParamsDateRange7d         RadarAs112TimeseriesGroupEdnListParamsDateRange = "7d"
	RadarAs112TimeseriesGroupEdnListParamsDateRange14d        RadarAs112TimeseriesGroupEdnListParamsDateRange = "14d"
	RadarAs112TimeseriesGroupEdnListParamsDateRange28d        RadarAs112TimeseriesGroupEdnListParamsDateRange = "28d"
	RadarAs112TimeseriesGroupEdnListParamsDateRange12w        RadarAs112TimeseriesGroupEdnListParamsDateRange = "12w"
	RadarAs112TimeseriesGroupEdnListParamsDateRange24w        RadarAs112TimeseriesGroupEdnListParamsDateRange = "24w"
	RadarAs112TimeseriesGroupEdnListParamsDateRange52w        RadarAs112TimeseriesGroupEdnListParamsDateRange = "52w"
	RadarAs112TimeseriesGroupEdnListParamsDateRange1dControl  RadarAs112TimeseriesGroupEdnListParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupEdnListParamsDateRange2dControl  RadarAs112TimeseriesGroupEdnListParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupEdnListParamsDateRange7dControl  RadarAs112TimeseriesGroupEdnListParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupEdnListParamsDateRange14dControl RadarAs112TimeseriesGroupEdnListParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupEdnListParamsDateRange28dControl RadarAs112TimeseriesGroupEdnListParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupEdnListParamsDateRange12wControl RadarAs112TimeseriesGroupEdnListParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupEdnListParamsDateRange24wControl RadarAs112TimeseriesGroupEdnListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupEdnListParamsFormat string

const (
	RadarAs112TimeseriesGroupEdnListParamsFormatJson RadarAs112TimeseriesGroupEdnListParamsFormat = "JSON"
	RadarAs112TimeseriesGroupEdnListParamsFormatCsv  RadarAs112TimeseriesGroupEdnListParamsFormat = "CSV"
)

type RadarAs112TimeseriesGroupEdnListResponseEnvelope struct {
	Result  RadarAs112TimeseriesGroupEdnListResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAs112TimeseriesGroupEdnListResponseEnvelopeJSON `json:"-"`
}

// radarAs112TimeseriesGroupEdnListResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupEdnListResponseEnvelope]
type radarAs112TimeseriesGroupEdnListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupEdnListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
