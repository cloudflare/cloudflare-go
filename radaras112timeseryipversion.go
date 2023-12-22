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

// RadarAs112TimeseryIPVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAs112TimeseryIPVersionService] method instead.
type RadarAs112TimeseryIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseryIPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseryIPVersionService(opts ...option.RequestOption) (r *RadarAs112TimeseryIPVersionService) {
	r = &RadarAs112TimeseryIPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of dns requests classified per IP Version over time.
func (r *RadarAs112TimeseryIPVersionService) List(ctx context.Context, query RadarAs112TimeseryIPVersionListParams, opts ...option.RequestOption) (res *RadarAs112TimeseryIPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseryIPVersionListResponse struct {
	Result  RadarAs112TimeseryIPVersionListResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarAs112TimeseryIPVersionListResponseJSON   `json:"-"`
}

// radarAs112TimeseryIPVersionListResponseJSON contains the JSON metadata for the
// struct [RadarAs112TimeseryIPVersionListResponse]
type radarAs112TimeseryIPVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryIPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryIPVersionListResponseResult struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarAs112TimeseryIPVersionListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseryIPVersionListResponseResultJSON   `json:"-"`
}

// radarAs112TimeseryIPVersionListResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112TimeseryIPVersionListResponseResult]
type radarAs112TimeseryIPVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryIPVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryIPVersionListResponseResultSerie0 struct {
	IPv4 []string                                                `json:"IPv4,required"`
	IPv6 []string                                                `json:"IPv6,required"`
	JSON radarAs112TimeseryIPVersionListResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseryIPVersionListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarAs112TimeseryIPVersionListResponseResultSerie0]
type radarAs112TimeseryIPVersionListResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryIPVersionListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryIPVersionListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseryIPVersionListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseryIPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseryIPVersionListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseryIPVersionListParams]'s query parameters
// as `url.Values`.
func (r RadarAs112TimeseryIPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseryIPVersionListParamsAggInterval string

const (
	RadarAs112TimeseryIPVersionListParamsAggInterval15m RadarAs112TimeseryIPVersionListParamsAggInterval = "15m"
	RadarAs112TimeseryIPVersionListParamsAggInterval1h  RadarAs112TimeseryIPVersionListParamsAggInterval = "1h"
	RadarAs112TimeseryIPVersionListParamsAggInterval1d  RadarAs112TimeseryIPVersionListParamsAggInterval = "1d"
	RadarAs112TimeseryIPVersionListParamsAggInterval1w  RadarAs112TimeseryIPVersionListParamsAggInterval = "1w"
)

type RadarAs112TimeseryIPVersionListParamsDateRange string

const (
	RadarAs112TimeseryIPVersionListParamsDateRange1d         RadarAs112TimeseryIPVersionListParamsDateRange = "1d"
	RadarAs112TimeseryIPVersionListParamsDateRange7d         RadarAs112TimeseryIPVersionListParamsDateRange = "7d"
	RadarAs112TimeseryIPVersionListParamsDateRange14d        RadarAs112TimeseryIPVersionListParamsDateRange = "14d"
	RadarAs112TimeseryIPVersionListParamsDateRange28d        RadarAs112TimeseryIPVersionListParamsDateRange = "28d"
	RadarAs112TimeseryIPVersionListParamsDateRange12w        RadarAs112TimeseryIPVersionListParamsDateRange = "12w"
	RadarAs112TimeseryIPVersionListParamsDateRange24w        RadarAs112TimeseryIPVersionListParamsDateRange = "24w"
	RadarAs112TimeseryIPVersionListParamsDateRange52w        RadarAs112TimeseryIPVersionListParamsDateRange = "52w"
	RadarAs112TimeseryIPVersionListParamsDateRange1dControl  RadarAs112TimeseryIPVersionListParamsDateRange = "1dControl"
	RadarAs112TimeseryIPVersionListParamsDateRange7dControl  RadarAs112TimeseryIPVersionListParamsDateRange = "7dControl"
	RadarAs112TimeseryIPVersionListParamsDateRange14dControl RadarAs112TimeseryIPVersionListParamsDateRange = "14dControl"
	RadarAs112TimeseryIPVersionListParamsDateRange28dControl RadarAs112TimeseryIPVersionListParamsDateRange = "28dControl"
	RadarAs112TimeseryIPVersionListParamsDateRange12wControl RadarAs112TimeseryIPVersionListParamsDateRange = "12wControl"
	RadarAs112TimeseryIPVersionListParamsDateRange24wControl RadarAs112TimeseryIPVersionListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseryIPVersionListParamsFormat string

const (
	RadarAs112TimeseryIPVersionListParamsFormatJson RadarAs112TimeseryIPVersionListParamsFormat = "JSON"
	RadarAs112TimeseryIPVersionListParamsFormatCsv  RadarAs112TimeseryIPVersionListParamsFormat = "CSV"
)
