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

// RadarAs112TimeseryResponseCodeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAs112TimeseryResponseCodeService] method instead.
type RadarAs112TimeseryResponseCodeService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseryResponseCodeService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseryResponseCodeService(opts ...option.RequestOption) (r *RadarAs112TimeseryResponseCodeService) {
	r = &RadarAs112TimeseryResponseCodeService{}
	r.Options = opts
	return
}

// Percentage distribution of dns requests classified per Response Codes over time.
func (r *RadarAs112TimeseryResponseCodeService) List(ctx context.Context, query RadarAs112TimeseryResponseCodeListParams, opts ...option.RequestOption) (res *RadarAs112TimeseryResponseCodeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries/response_codes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseryResponseCodeListResponse struct {
	Result  RadarAs112TimeseryResponseCodeListResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarAs112TimeseryResponseCodeListResponseJSON   `json:"-"`
}

// radarAs112TimeseryResponseCodeListResponseJSON contains the JSON metadata for
// the struct [RadarAs112TimeseryResponseCodeListResponse]
type radarAs112TimeseryResponseCodeListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryResponseCodeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryResponseCodeListResponseResult struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarAs112TimeseryResponseCodeListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseryResponseCodeListResponseResultJSON   `json:"-"`
}

// radarAs112TimeseryResponseCodeListResponseResultJSON contains the JSON metadata
// for the struct [RadarAs112TimeseryResponseCodeListResponseResult]
type radarAs112TimeseryResponseCodeListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryResponseCodeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryResponseCodeListResponseResultSerie0 struct {
	Noerror  []string                                                   `json:"NOERROR,required"`
	Nxdomain []string                                                   `json:"NXDOMAIN,required"`
	JSON     radarAs112TimeseryResponseCodeListResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseryResponseCodeListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarAs112TimeseryResponseCodeListResponseResultSerie0]
type radarAs112TimeseryResponseCodeListResponseResultSerie0JSON struct {
	Noerror     apijson.Field
	Nxdomain    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryResponseCodeListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryResponseCodeListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseryResponseCodeListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseryResponseCodeListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseryResponseCodeListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseryResponseCodeListParams]'s query
// parameters as `url.Values`.
func (r RadarAs112TimeseryResponseCodeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseryResponseCodeListParamsAggInterval string

const (
	RadarAs112TimeseryResponseCodeListParamsAggInterval15m RadarAs112TimeseryResponseCodeListParamsAggInterval = "15m"
	RadarAs112TimeseryResponseCodeListParamsAggInterval1h  RadarAs112TimeseryResponseCodeListParamsAggInterval = "1h"
	RadarAs112TimeseryResponseCodeListParamsAggInterval1d  RadarAs112TimeseryResponseCodeListParamsAggInterval = "1d"
	RadarAs112TimeseryResponseCodeListParamsAggInterval1w  RadarAs112TimeseryResponseCodeListParamsAggInterval = "1w"
)

type RadarAs112TimeseryResponseCodeListParamsDateRange string

const (
	RadarAs112TimeseryResponseCodeListParamsDateRange1d         RadarAs112TimeseryResponseCodeListParamsDateRange = "1d"
	RadarAs112TimeseryResponseCodeListParamsDateRange7d         RadarAs112TimeseryResponseCodeListParamsDateRange = "7d"
	RadarAs112TimeseryResponseCodeListParamsDateRange14d        RadarAs112TimeseryResponseCodeListParamsDateRange = "14d"
	RadarAs112TimeseryResponseCodeListParamsDateRange28d        RadarAs112TimeseryResponseCodeListParamsDateRange = "28d"
	RadarAs112TimeseryResponseCodeListParamsDateRange12w        RadarAs112TimeseryResponseCodeListParamsDateRange = "12w"
	RadarAs112TimeseryResponseCodeListParamsDateRange24w        RadarAs112TimeseryResponseCodeListParamsDateRange = "24w"
	RadarAs112TimeseryResponseCodeListParamsDateRange52w        RadarAs112TimeseryResponseCodeListParamsDateRange = "52w"
	RadarAs112TimeseryResponseCodeListParamsDateRange1dControl  RadarAs112TimeseryResponseCodeListParamsDateRange = "1dControl"
	RadarAs112TimeseryResponseCodeListParamsDateRange7dControl  RadarAs112TimeseryResponseCodeListParamsDateRange = "7dControl"
	RadarAs112TimeseryResponseCodeListParamsDateRange14dControl RadarAs112TimeseryResponseCodeListParamsDateRange = "14dControl"
	RadarAs112TimeseryResponseCodeListParamsDateRange28dControl RadarAs112TimeseryResponseCodeListParamsDateRange = "28dControl"
	RadarAs112TimeseryResponseCodeListParamsDateRange12wControl RadarAs112TimeseryResponseCodeListParamsDateRange = "12wControl"
	RadarAs112TimeseryResponseCodeListParamsDateRange24wControl RadarAs112TimeseryResponseCodeListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseryResponseCodeListParamsFormat string

const (
	RadarAs112TimeseryResponseCodeListParamsFormatJson RadarAs112TimeseryResponseCodeListParamsFormat = "JSON"
	RadarAs112TimeseryResponseCodeListParamsFormatCsv  RadarAs112TimeseryResponseCodeListParamsFormat = "CSV"
)
