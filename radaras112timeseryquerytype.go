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

// RadarAs112TimeseryQueryTypeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAs112TimeseryQueryTypeService] method instead.
type RadarAs112TimeseryQueryTypeService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseryQueryTypeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseryQueryTypeService(opts ...option.RequestOption) (r *RadarAs112TimeseryQueryTypeService) {
	r = &RadarAs112TimeseryQueryTypeService{}
	r.Options = opts
	return
}

// Percentage distribution of dns requests classified per Query Type over time.
func (r *RadarAs112TimeseryQueryTypeService) List(ctx context.Context, query RadarAs112TimeseryQueryTypeListParams, opts ...option.RequestOption) (res *RadarAs112TimeseryQueryTypeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries/query_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseryQueryTypeListResponse struct {
	Result  RadarAs112TimeseryQueryTypeListResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarAs112TimeseryQueryTypeListResponseJSON   `json:"-"`
}

// radarAs112TimeseryQueryTypeListResponseJSON contains the JSON metadata for the
// struct [RadarAs112TimeseryQueryTypeListResponse]
type radarAs112TimeseryQueryTypeListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryQueryTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryQueryTypeListResponseResult struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarAs112TimeseryQueryTypeListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseryQueryTypeListResponseResultJSON   `json:"-"`
}

// radarAs112TimeseryQueryTypeListResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112TimeseryQueryTypeListResponseResult]
type radarAs112TimeseryQueryTypeListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryQueryTypeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryQueryTypeListResponseResultSerie0 struct {
	A    []string                                                `json:"A,required"`
	Aaaa []string                                                `json:"AAAA,required"`
	Ptr  []string                                                `json:"PTR,required"`
	Soa  []string                                                `json:"SOA,required"`
	Srv  []string                                                `json:"SRV,required"`
	JSON radarAs112TimeseryQueryTypeListResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseryQueryTypeListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarAs112TimeseryQueryTypeListResponseResultSerie0]
type radarAs112TimeseryQueryTypeListResponseResultSerie0JSON struct {
	A           apijson.Field
	Aaaa        apijson.Field
	Ptr         apijson.Field
	Soa         apijson.Field
	Srv         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryQueryTypeListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryQueryTypeListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseryQueryTypeListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseryQueryTypeListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseryQueryTypeListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseryQueryTypeListParams]'s query parameters
// as `url.Values`.
func (r RadarAs112TimeseryQueryTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseryQueryTypeListParamsAggInterval string

const (
	RadarAs112TimeseryQueryTypeListParamsAggInterval15m RadarAs112TimeseryQueryTypeListParamsAggInterval = "15m"
	RadarAs112TimeseryQueryTypeListParamsAggInterval1h  RadarAs112TimeseryQueryTypeListParamsAggInterval = "1h"
	RadarAs112TimeseryQueryTypeListParamsAggInterval1d  RadarAs112TimeseryQueryTypeListParamsAggInterval = "1d"
	RadarAs112TimeseryQueryTypeListParamsAggInterval1w  RadarAs112TimeseryQueryTypeListParamsAggInterval = "1w"
)

type RadarAs112TimeseryQueryTypeListParamsDateRange string

const (
	RadarAs112TimeseryQueryTypeListParamsDateRange1d         RadarAs112TimeseryQueryTypeListParamsDateRange = "1d"
	RadarAs112TimeseryQueryTypeListParamsDateRange7d         RadarAs112TimeseryQueryTypeListParamsDateRange = "7d"
	RadarAs112TimeseryQueryTypeListParamsDateRange14d        RadarAs112TimeseryQueryTypeListParamsDateRange = "14d"
	RadarAs112TimeseryQueryTypeListParamsDateRange28d        RadarAs112TimeseryQueryTypeListParamsDateRange = "28d"
	RadarAs112TimeseryQueryTypeListParamsDateRange12w        RadarAs112TimeseryQueryTypeListParamsDateRange = "12w"
	RadarAs112TimeseryQueryTypeListParamsDateRange24w        RadarAs112TimeseryQueryTypeListParamsDateRange = "24w"
	RadarAs112TimeseryQueryTypeListParamsDateRange52w        RadarAs112TimeseryQueryTypeListParamsDateRange = "52w"
	RadarAs112TimeseryQueryTypeListParamsDateRange1dControl  RadarAs112TimeseryQueryTypeListParamsDateRange = "1dControl"
	RadarAs112TimeseryQueryTypeListParamsDateRange7dControl  RadarAs112TimeseryQueryTypeListParamsDateRange = "7dControl"
	RadarAs112TimeseryQueryTypeListParamsDateRange14dControl RadarAs112TimeseryQueryTypeListParamsDateRange = "14dControl"
	RadarAs112TimeseryQueryTypeListParamsDateRange28dControl RadarAs112TimeseryQueryTypeListParamsDateRange = "28dControl"
	RadarAs112TimeseryQueryTypeListParamsDateRange12wControl RadarAs112TimeseryQueryTypeListParamsDateRange = "12wControl"
	RadarAs112TimeseryQueryTypeListParamsDateRange24wControl RadarAs112TimeseryQueryTypeListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseryQueryTypeListParamsFormat string

const (
	RadarAs112TimeseryQueryTypeListParamsFormatJson RadarAs112TimeseryQueryTypeListParamsFormat = "JSON"
	RadarAs112TimeseryQueryTypeListParamsFormatCsv  RadarAs112TimeseryQueryTypeListParamsFormat = "CSV"
)
