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
func (r *RadarAs112TimeseriesGroupIPVersionService) Get(ctx context.Context, query RadarAs112TimeseriesGroupIPVersionGetParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupIPVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseriesGroupIPVersionGetResponse struct {
	Result  RadarAs112TimeseriesGroupIPVersionGetResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarAs112TimeseriesGroupIPVersionGetResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupIPVersionGetResponseJSON contains the JSON metadata for
// the struct [RadarAs112TimeseriesGroupIPVersionGetResponse]
type radarAs112TimeseriesGroupIPVersionGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupIPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupIPVersionGetResponseResult struct {
	Meta   interface{}                                               `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupIPVersionGetResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupIPVersionGetResponseResultJSON   `json:"-"`
}

// radarAs112TimeseriesGroupIPVersionGetResponseResultJSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupIPVersionGetResponseResult]
type radarAs112TimeseriesGroupIPVersionGetResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupIPVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupIPVersionGetResponseResultSerie0 struct {
	IPv4 []string                                                      `json:"IPv4,required"`
	IPv6 []string                                                      `json:"IPv6,required"`
	JSON radarAs112TimeseriesGroupIPVersionGetResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupIPVersionGetResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarAs112TimeseriesGroupIPVersionGetResponseResultSerie0]
type radarAs112TimeseriesGroupIPVersionGetResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupIPVersionGetResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupIPVersionGetParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupIPVersionGetParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupIPVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupIPVersionGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupIPVersionGetParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupIPVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupIPVersionGetParamsAggInterval string

const (
	RadarAs112TimeseriesGroupIPVersionGetParamsAggInterval15m RadarAs112TimeseriesGroupIPVersionGetParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupIPVersionGetParamsAggInterval1h  RadarAs112TimeseriesGroupIPVersionGetParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupIPVersionGetParamsAggInterval1d  RadarAs112TimeseriesGroupIPVersionGetParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupIPVersionGetParamsAggInterval1w  RadarAs112TimeseriesGroupIPVersionGetParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupIPVersionGetParamsDateRange string

const (
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange1d         RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "1d"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange2d         RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "2d"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange7d         RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "7d"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange14d        RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "14d"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange28d        RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "28d"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange12w        RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "12w"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange24w        RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "24w"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange52w        RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "52w"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange1dControl  RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange2dControl  RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange7dControl  RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange14dControl RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange28dControl RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange12wControl RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupIPVersionGetParamsDateRange24wControl RadarAs112TimeseriesGroupIPVersionGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupIPVersionGetParamsFormat string

const (
	RadarAs112TimeseriesGroupIPVersionGetParamsFormatJson RadarAs112TimeseriesGroupIPVersionGetParamsFormat = "JSON"
	RadarAs112TimeseriesGroupIPVersionGetParamsFormatCsv  RadarAs112TimeseriesGroupIPVersionGetParamsFormat = "CSV"
)
