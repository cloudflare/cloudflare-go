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

// RadarAs112TimeseriesGroupResponseCodeService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAs112TimeseriesGroupResponseCodeService] method instead.
type RadarAs112TimeseriesGroupResponseCodeService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseriesGroupResponseCodeService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAs112TimeseriesGroupResponseCodeService(opts ...option.RequestOption) (r *RadarAs112TimeseriesGroupResponseCodeService) {
	r = &RadarAs112TimeseriesGroupResponseCodeService{}
	r.Options = opts
	return
}

// Percentage distribution of AS112 dns requests classified per Response Codes over
// time.
func (r *RadarAs112TimeseriesGroupResponseCodeService) Get(ctx context.Context, query RadarAs112TimeseriesGroupResponseCodeGetParams, opts ...option.RequestOption) (res *RadarAs112TimeseriesGroupResponseCodeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries_groups/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseriesGroupResponseCodeGetResponse struct {
	Result  RadarAs112TimeseriesGroupResponseCodeGetResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarAs112TimeseriesGroupResponseCodeGetResponseJSON   `json:"-"`
}

// radarAs112TimeseriesGroupResponseCodeGetResponseJSON contains the JSON metadata
// for the struct [RadarAs112TimeseriesGroupResponseCodeGetResponse]
type radarAs112TimeseriesGroupResponseCodeGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupResponseCodeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupResponseCodeGetResponseResult struct {
	Meta   interface{}                                                  `json:"meta,required"`
	Serie0 RadarAs112TimeseriesGroupResponseCodeGetResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseriesGroupResponseCodeGetResponseResultJSON   `json:"-"`
}

// radarAs112TimeseriesGroupResponseCodeGetResponseResultJSON contains the JSON
// metadata for the struct [RadarAs112TimeseriesGroupResponseCodeGetResponseResult]
type radarAs112TimeseriesGroupResponseCodeGetResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupResponseCodeGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupResponseCodeGetResponseResultSerie0 struct {
	Noerror  []string                                                         `json:"NOERROR,required"`
	Nxdomain []string                                                         `json:"NXDOMAIN,required"`
	JSON     radarAs112TimeseriesGroupResponseCodeGetResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseriesGroupResponseCodeGetResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarAs112TimeseriesGroupResponseCodeGetResponseResultSerie0]
type radarAs112TimeseriesGroupResponseCodeGetResponseResultSerie0JSON struct {
	Noerror     apijson.Field
	Nxdomain    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseriesGroupResponseCodeGetResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseriesGroupResponseCodeGetParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseriesGroupResponseCodeGetParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseriesGroupResponseCodeGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseriesGroupResponseCodeGetParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseriesGroupResponseCodeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseriesGroupResponseCodeGetParamsAggInterval string

const (
	RadarAs112TimeseriesGroupResponseCodeGetParamsAggInterval15m RadarAs112TimeseriesGroupResponseCodeGetParamsAggInterval = "15m"
	RadarAs112TimeseriesGroupResponseCodeGetParamsAggInterval1h  RadarAs112TimeseriesGroupResponseCodeGetParamsAggInterval = "1h"
	RadarAs112TimeseriesGroupResponseCodeGetParamsAggInterval1d  RadarAs112TimeseriesGroupResponseCodeGetParamsAggInterval = "1d"
	RadarAs112TimeseriesGroupResponseCodeGetParamsAggInterval1w  RadarAs112TimeseriesGroupResponseCodeGetParamsAggInterval = "1w"
)

type RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange string

const (
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange1d         RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "1d"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange2d         RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "2d"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange7d         RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "7d"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange14d        RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "14d"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange28d        RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "28d"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange12w        RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "12w"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange24w        RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "24w"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange52w        RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "52w"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange1dControl  RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "1dControl"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange2dControl  RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "2dControl"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange7dControl  RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "7dControl"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange14dControl RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "14dControl"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange28dControl RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "28dControl"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange12wControl RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "12wControl"
	RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange24wControl RadarAs112TimeseriesGroupResponseCodeGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseriesGroupResponseCodeGetParamsFormat string

const (
	RadarAs112TimeseriesGroupResponseCodeGetParamsFormatJson RadarAs112TimeseriesGroupResponseCodeGetParamsFormat = "JSON"
	RadarAs112TimeseriesGroupResponseCodeGetParamsFormatCsv  RadarAs112TimeseriesGroupResponseCodeGetParamsFormat = "CSV"
)
