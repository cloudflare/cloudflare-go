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

// RadarAs112TimeseriesGroupIPVersionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAs112TimeseriesGroupIPVersionService] method instead.
type RadarAs112TimeseriesGroupIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseriesGroupIPVersionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAs112TimeseriesGroupIPVersionService(opts ...option.RequestOption) (r *RadarAs112TimeseriesGroupIPVersionService) {
	r = &RadarAs112TimeseriesGroupIPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of AS112 DNS queries by IP Version over time.
func (r *RadarAs112TimeseriesGroupIPVersionService) List(ctx context.Context, query RadarAs112TimeseriesGroupIPVersionListParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupIPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAs112TimeseriesGroupIPVersionListResponseEnvelope
	path := "radar/as112/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAs112TimeseriesGroupIPVersionListResponse struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupIPVersionListResponseSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupIPVersionListResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupIPVersionListResponseJSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupIPVersionListResponse]
type radarAs112TimeseriesGroupIPVersionListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupIPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupIPVersionListResponseSerie0 struct {
	IPv4 []string                                                 `json:"IPv4,required"`
	IPv6 []string                                                 `json:"IPv6,required"`
	JSON radarAs112TimeseriesGroupIPVersionListResponseSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupIPVersionListResponseSerie0JSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupIPVersionListResponseSerie0]
type radarAs112TimeseriesGroupIPVersionListResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupIPVersionListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupIPVersionListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupIPVersionListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupIPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupIPVersionListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupIPVersionListParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupIPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupIPVersionListParamsAggInterval string

const (
	RadarAs112TimeseriesGroupIPVersionListParamsAggInterval15m RadarAs112TimeseriesGroupIPVersionListParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupIPVersionListParamsAggInterval1h  RadarAs112TimeseriesGroupIPVersionListParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupIPVersionListParamsAggInterval1d  RadarAs112TimeseriesGroupIPVersionListParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupIPVersionListParamsAggInterval1w  RadarAs112TimeseriesGroupIPVersionListParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupIPVersionListParamsDateRange string

const (
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange1d         RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "1d"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange2d         RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "2d"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange7d         RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "7d"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange14d        RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "14d"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange28d        RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "28d"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange12w        RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "12w"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange24w        RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "24w"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange52w        RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "52w"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange1dControl  RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange2dControl  RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange7dControl  RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange14dControl RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange28dControl RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange12wControl RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupIPVersionListParamsDateRange24wControl RadarAs112TimeseriesGroupIPVersionListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupIPVersionListParamsFormat string

const (
	RadarAs112TimeseriesGroupIPVersionListParamsFormatJson RadarAs112TimeseriesGroupIPVersionListParamsFormat = "JSON"
	RadarAs112TimeseriesGroupIPVersionListParamsFormatCsv  RadarAs112TimeseriesGroupIPVersionListParamsFormat = "CSV"
)

type RadarAs112TimeseriesGroupIPVersionListResponseEnvelope struct {
	Result  RadarAs112TimeseriesGroupIPVersionListResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarAs112TimeseriesGroupIPVersionListResponseEnvelopeJSON `json:"-"`
}

// radarAs112TimeseriesGroupIPVersionListResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupIPVersionListResponseEnvelope]
type radarAs112TimeseriesGroupIPVersionListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupIPVersionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
