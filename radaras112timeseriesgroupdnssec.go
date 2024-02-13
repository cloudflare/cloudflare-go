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

// RadarAs112TimeseriesGroupDNSSECService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAs112TimeseriesGroupDNSSECService] method instead.
type RadarAs112TimeseriesGroupDNSSECService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseriesGroupDNSSECService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseriesGroupDNSSECService(opts ...option.RequestOption) (r *RadarAs112TimeseriesGroupDNSSECService) {
	r = &RadarAs112TimeseriesGroupDNSSECService{}
	r.Options = opts
	return
}

// Percentage distribution of DNS AS112 queries by DNSSEC support over time.
func (r *RadarAs112TimeseriesGroupDNSSECService) List(ctx context.Context, query RadarAs112TimeseriesGroupDNSSECListParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupDNSSECListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TimeseriesGroupDNSSECListResponseEnvelope
	path := "radar/as112/timeseries_groups/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAs112TimeseriesGroupDNSSECListResponse struct {
	Meta   interface{}                                       `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupDNSSECListResponseSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupDNSSECListResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupDNSSECListResponseJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupDNSSECListResponse]
type radarAs112TimeseriesGroupDNSSECListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupDNSSECListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupDNSSECListResponseSerie0 struct {
	NotSupported []string                                              `json:"NOT_SUPPORTED,required"`
	Supported    []string                                              `json:"SUPPORTED,required"`
	JSON         radarAs112TimeseriesGroupDNSSECListResponseSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupDNSSECListResponseSerie0JSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupDNSSECListResponseSerie0]
type radarAs112TimeseriesGroupDNSSECListResponseSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupDNSSECListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupDNSSECListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupDNSSECListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupDNSSECListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupDNSSECListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupDNSSECListParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupDNSSECListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupDNSSECListParamsAggInterval string

const (
	RadarAs112TimeseriesGroupDNSSECListParamsAggInterval15m RadarAs112TimeseriesGroupDNSSECListParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupDNSSECListParamsAggInterval1h  RadarAs112TimeseriesGroupDNSSECListParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupDNSSECListParamsAggInterval1d  RadarAs112TimeseriesGroupDNSSECListParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupDNSSECListParamsAggInterval1w  RadarAs112TimeseriesGroupDNSSECListParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupDNSSECListParamsDateRange string

const (
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange1d         RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "1d"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange2d         RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "2d"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange7d         RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "7d"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange14d        RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "14d"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange28d        RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "28d"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange12w        RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "12w"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange24w        RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "24w"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange52w        RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "52w"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange1dControl  RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange2dControl  RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange7dControl  RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange14dControl RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange28dControl RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange12wControl RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupDNSSECListParamsDateRange24wControl RadarAs112TimeseriesGroupDNSSECListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupDNSSECListParamsFormat string

const (
	RadarAs112TimeseriesGroupDNSSECListParamsFormatJson RadarAs112TimeseriesGroupDNSSECListParamsFormat = "JSON"
	RadarAs112TimeseriesGroupDNSSECListParamsFormatCsv  RadarAs112TimeseriesGroupDNSSECListParamsFormat = "CSV"
)

type RadarAs112TimeseriesGroupDNSSECListResponseEnvelope struct {
	Result  RadarAs112TimeseriesGroupDNSSECListResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarAs112TimeseriesGroupDNSSECListResponseEnvelopeJSON `json:"-"`
}

// radarAs112TimeseriesGroupDNSSECListResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupDNSSECListResponseEnvelope]
type radarAs112TimeseriesGroupDNSSECListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupDNSSECListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
