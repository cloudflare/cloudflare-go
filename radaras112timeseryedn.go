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

// RadarAs112TimeseryEdnService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAs112TimeseryEdnService]
// method instead.
type RadarAs112TimeseryEdnService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseryEdnService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseryEdnService(opts ...option.RequestOption) (r *RadarAs112TimeseryEdnService) {
	r = &RadarAs112TimeseryEdnService{}
	r.Options = opts
	return
}

// Percentage distribution of dns requests classified per EDNS over time.
func (r *RadarAs112TimeseryEdnService) List(ctx context.Context, query RadarAs112TimeseryEdnListParams, opts ...option.RequestOption) (res *RadarAs112TimeseryEdnListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries/edns"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseryEdnListResponse struct {
	Result  RadarAs112TimeseryEdnListResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarAs112TimeseryEdnListResponseJSON   `json:"-"`
}

// radarAs112TimeseryEdnListResponseJSON contains the JSON metadata for the struct
// [RadarAs112TimeseryEdnListResponse]
type radarAs112TimeseryEdnListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryEdnListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryEdnListResponseResult struct {
	Meta   interface{}                                   `json:"meta,required"`
	Serie0 RadarAs112TimeseryEdnListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseryEdnListResponseResultJSON   `json:"-"`
}

// radarAs112TimeseryEdnListResponseResultJSON contains the JSON metadata for the
// struct [RadarAs112TimeseryEdnListResponseResult]
type radarAs112TimeseryEdnListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryEdnListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryEdnListResponseResultSerie0 struct {
	NotSupported []string                                          `json:"NOT_SUPPORTED,required"`
	Supported    []string                                          `json:"SUPPORTED,required"`
	JSON         radarAs112TimeseryEdnListResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseryEdnListResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarAs112TimeseryEdnListResponseResultSerie0]
type radarAs112TimeseryEdnListResponseResultSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112TimeseryEdnListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryEdnListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseryEdnListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseryEdnListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseryEdnListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseryEdnListParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TimeseryEdnListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseryEdnListParamsAggInterval string

const (
	RadarAs112TimeseryEdnListParamsAggInterval15m RadarAs112TimeseryEdnListParamsAggInterval = "15m"
	RadarAs112TimeseryEdnListParamsAggInterval1h  RadarAs112TimeseryEdnListParamsAggInterval = "1h"
	RadarAs112TimeseryEdnListParamsAggInterval1d  RadarAs112TimeseryEdnListParamsAggInterval = "1d"
	RadarAs112TimeseryEdnListParamsAggInterval1w  RadarAs112TimeseryEdnListParamsAggInterval = "1w"
)

type RadarAs112TimeseryEdnListParamsDateRange string

const (
	RadarAs112TimeseryEdnListParamsDateRange1d         RadarAs112TimeseryEdnListParamsDateRange = "1d"
	RadarAs112TimeseryEdnListParamsDateRange7d         RadarAs112TimeseryEdnListParamsDateRange = "7d"
	RadarAs112TimeseryEdnListParamsDateRange14d        RadarAs112TimeseryEdnListParamsDateRange = "14d"
	RadarAs112TimeseryEdnListParamsDateRange28d        RadarAs112TimeseryEdnListParamsDateRange = "28d"
	RadarAs112TimeseryEdnListParamsDateRange12w        RadarAs112TimeseryEdnListParamsDateRange = "12w"
	RadarAs112TimeseryEdnListParamsDateRange24w        RadarAs112TimeseryEdnListParamsDateRange = "24w"
	RadarAs112TimeseryEdnListParamsDateRange52w        RadarAs112TimeseryEdnListParamsDateRange = "52w"
	RadarAs112TimeseryEdnListParamsDateRange1dControl  RadarAs112TimeseryEdnListParamsDateRange = "1dControl"
	RadarAs112TimeseryEdnListParamsDateRange7dControl  RadarAs112TimeseryEdnListParamsDateRange = "7dControl"
	RadarAs112TimeseryEdnListParamsDateRange14dControl RadarAs112TimeseryEdnListParamsDateRange = "14dControl"
	RadarAs112TimeseryEdnListParamsDateRange28dControl RadarAs112TimeseryEdnListParamsDateRange = "28dControl"
	RadarAs112TimeseryEdnListParamsDateRange12wControl RadarAs112TimeseryEdnListParamsDateRange = "12wControl"
	RadarAs112TimeseryEdnListParamsDateRange24wControl RadarAs112TimeseryEdnListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseryEdnListParamsFormat string

const (
	RadarAs112TimeseryEdnListParamsFormatJson RadarAs112TimeseryEdnListParamsFormat = "JSON"
	RadarAs112TimeseryEdnListParamsFormatCsv  RadarAs112TimeseryEdnListParamsFormat = "CSV"
)
