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

// RadarAs112TimeseryProtocolService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAs112TimeseryProtocolService] method instead.
type RadarAs112TimeseryProtocolService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseryProtocolService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseryProtocolService(opts ...option.RequestOption) (r *RadarAs112TimeseryProtocolService) {
	r = &RadarAs112TimeseryProtocolService{}
	r.Options = opts
	return
}

// Percentage distribution of dns requests classified per Protocol over time.
func (r *RadarAs112TimeseryProtocolService) List(ctx context.Context, query RadarAs112TimeseryProtocolListParams, opts ...option.RequestOption) (res *RadarAs112TimeseryProtocolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseryProtocolListResponse struct {
	Result  RadarAs112TimeseryProtocolListResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarAs112TimeseryProtocolListResponseJSON   `json:"-"`
}

// radarAs112TimeseryProtocolListResponseJSON contains the JSON metadata for the
// struct [RadarAs112TimeseryProtocolListResponse]
type radarAs112TimeseryProtocolListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryProtocolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryProtocolListResponseResult struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 RadarAs112TimeseryProtocolListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseryProtocolListResponseResultJSON   `json:"-"`
}

// radarAs112TimeseryProtocolListResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112TimeseryProtocolListResponseResult]
type radarAs112TimeseryProtocolListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryProtocolListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryProtocolListResponseResultSerie0 struct {
	Tcp  []string                                               `json:"tcp,required"`
	Udp  []string                                               `json:"udp,required"`
	JSON radarAs112TimeseryProtocolListResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseryProtocolListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarAs112TimeseryProtocolListResponseResultSerie0]
type radarAs112TimeseryProtocolListResponseResultSerie0JSON struct {
	Tcp         apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryProtocolListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryProtocolListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseryProtocolListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseryProtocolListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseryProtocolListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseryProtocolListParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TimeseryProtocolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseryProtocolListParamsAggInterval string

const (
	RadarAs112TimeseryProtocolListParamsAggInterval15m RadarAs112TimeseryProtocolListParamsAggInterval = "15m"
	RadarAs112TimeseryProtocolListParamsAggInterval1h  RadarAs112TimeseryProtocolListParamsAggInterval = "1h"
	RadarAs112TimeseryProtocolListParamsAggInterval1d  RadarAs112TimeseryProtocolListParamsAggInterval = "1d"
	RadarAs112TimeseryProtocolListParamsAggInterval1w  RadarAs112TimeseryProtocolListParamsAggInterval = "1w"
)

type RadarAs112TimeseryProtocolListParamsDateRange string

const (
	RadarAs112TimeseryProtocolListParamsDateRange1d         RadarAs112TimeseryProtocolListParamsDateRange = "1d"
	RadarAs112TimeseryProtocolListParamsDateRange7d         RadarAs112TimeseryProtocolListParamsDateRange = "7d"
	RadarAs112TimeseryProtocolListParamsDateRange14d        RadarAs112TimeseryProtocolListParamsDateRange = "14d"
	RadarAs112TimeseryProtocolListParamsDateRange28d        RadarAs112TimeseryProtocolListParamsDateRange = "28d"
	RadarAs112TimeseryProtocolListParamsDateRange12w        RadarAs112TimeseryProtocolListParamsDateRange = "12w"
	RadarAs112TimeseryProtocolListParamsDateRange24w        RadarAs112TimeseryProtocolListParamsDateRange = "24w"
	RadarAs112TimeseryProtocolListParamsDateRange52w        RadarAs112TimeseryProtocolListParamsDateRange = "52w"
	RadarAs112TimeseryProtocolListParamsDateRange1dControl  RadarAs112TimeseryProtocolListParamsDateRange = "1dControl"
	RadarAs112TimeseryProtocolListParamsDateRange7dControl  RadarAs112TimeseryProtocolListParamsDateRange = "7dControl"
	RadarAs112TimeseryProtocolListParamsDateRange14dControl RadarAs112TimeseryProtocolListParamsDateRange = "14dControl"
	RadarAs112TimeseryProtocolListParamsDateRange28dControl RadarAs112TimeseryProtocolListParamsDateRange = "28dControl"
	RadarAs112TimeseryProtocolListParamsDateRange12wControl RadarAs112TimeseryProtocolListParamsDateRange = "12wControl"
	RadarAs112TimeseryProtocolListParamsDateRange24wControl RadarAs112TimeseryProtocolListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseryProtocolListParamsFormat string

const (
	RadarAs112TimeseryProtocolListParamsFormatJson RadarAs112TimeseryProtocolListParamsFormat = "JSON"
	RadarAs112TimeseryProtocolListParamsFormatCsv  RadarAs112TimeseryProtocolListParamsFormat = "CSV"
)
