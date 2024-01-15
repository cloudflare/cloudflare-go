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

// RadarAs112TimeseriesGroupDnssecService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAs112TimeseriesGroupDnssecService] method instead.
type RadarAs112TimeseriesGroupDnssecService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseriesGroupDnssecService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseriesGroupDnssecService(opts ...option.RequestOption) (r *RadarAs112TimeseriesGroupDnssecService) {
	r = &RadarAs112TimeseriesGroupDnssecService{}
	r.Options = opts
	return
}

// Percentage distribution of DNS AS112 queries by DNSSEC support over time.
func (r *RadarAs112TimeseriesGroupDnssecService) Get(ctx context.Context, query RadarAs112TimeseriesGroupDnssecGetParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupDnssecGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseriesGroupDnssecGetResponse struct {
	Result  RadarAs112TimeseriesGroupDnssecGetResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarAs112TimeseriesGroupDnssecGetResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupDnssecGetResponseJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupDnssecGetResponse]
type radarAs112TimeseriesGroupDnssecGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupDnssecGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupDnssecGetResponseResult struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupDnssecGetResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupDnssecGetResponseResultJSON   `json:"-"`
}

// radarAs112TimeseriesGroupDnssecGetResponseResultJSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupDnssecGetResponseResult]
type radarAs112TimeseriesGroupDnssecGetResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupDnssecGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupDnssecGetResponseResultSerie0 struct {
	NotSupported []string                                                   `json:"NOT_SUPPORTED,required"`
	Supported    []string                                                   `json:"SUPPORTED,required"`
	JSON         radarAs112TimeseriesGroupDnssecGetResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupDnssecGetResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupDnssecGetResponseResultSerie0]
type radarAs112TimeseriesGroupDnssecGetResponseResultSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupDnssecGetResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupDnssecGetParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupDnssecGetParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupDnssecGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupDnssecGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupDnssecGetParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupDnssecGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupDnssecGetParamsAggInterval string

const (
	RadarAs112TimeseriesGroupDnssecGetParamsAggInterval15m RadarAs112TimeseriesGroupDnssecGetParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupDnssecGetParamsAggInterval1h  RadarAs112TimeseriesGroupDnssecGetParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupDnssecGetParamsAggInterval1d  RadarAs112TimeseriesGroupDnssecGetParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupDnssecGetParamsAggInterval1w  RadarAs112TimeseriesGroupDnssecGetParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupDnssecGetParamsDateRange string

const (
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange1d         RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "1d"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange2d         RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "2d"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange7d         RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "7d"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange14d        RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "14d"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange28d        RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "28d"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange12w        RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "12w"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange24w        RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "24w"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange52w        RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "52w"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange1dControl  RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange2dControl  RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange7dControl  RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange14dControl RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange28dControl RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange12wControl RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupDnssecGetParamsDateRange24wControl RadarAs112TimeseriesGroupDnssecGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupDnssecGetParamsFormat string

const (
	RadarAs112TimeseriesGroupDnssecGetParamsFormatJson RadarAs112TimeseriesGroupDnssecGetParamsFormat = "JSON"
	RadarAs112TimeseriesGroupDnssecGetParamsFormatCsv  RadarAs112TimeseriesGroupDnssecGetParamsFormat = "CSV"
)
