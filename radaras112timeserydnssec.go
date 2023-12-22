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

// RadarAs112TimeseryDnssecService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAs112TimeseryDnssecService] method instead.
type RadarAs112TimeseryDnssecService struct {
	Options []option.RequestOption
}

// NewRadarAs112TimeseryDnssecService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAs112TimeseryDnssecService(opts ...option.RequestOption) (r *RadarAs112TimeseryDnssecService) {
	r = &RadarAs112TimeseryDnssecService{}
	r.Options = opts
	return
}

// Percentage distribution of dns requests classified per DNSSEC over time.
func (r *RadarAs112TimeseryDnssecService) List(ctx context.Context, query RadarAs112TimeseryDnssecListParams, opts ...option.RequestOption) (res *RadarAs112TimeseryDnssecListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/as112/timeseries/dnssec"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAs112TimeseryDnssecListResponse struct {
	Result  RadarAs112TimeseryDnssecListResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarAs112TimeseryDnssecListResponseJSON   `json:"-"`
}

// radarAs112TimeseryDnssecListResponseJSON contains the JSON metadata for the
// struct [RadarAs112TimeseryDnssecListResponse]
type radarAs112TimeseryDnssecListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryDnssecListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryDnssecListResponseResult struct {
	Meta   interface{}                                      `json:"meta,required"`
	Serie0 RadarAs112TimeseryDnssecListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAs112TimeseryDnssecListResponseResultJSON   `json:"-"`
}

// radarAs112TimeseryDnssecListResponseResultJSON contains the JSON metadata for
// the struct [RadarAs112TimeseryDnssecListResponseResult]
type radarAs112TimeseryDnssecListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAs112TimeseryDnssecListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryDnssecListResponseResultSerie0 struct {
	NotSupported []string                                             `json:"NOT_SUPPORTED,required"`
	Supported    []string                                             `json:"SUPPORTED,required"`
	JSON         radarAs112TimeseryDnssecListResponseResultSerie0JSON `json:"-"`
}

// radarAs112TimeseryDnssecListResponseResultSerie0JSON contains the JSON metadata
// for the struct [RadarAs112TimeseryDnssecListResponseResultSerie0]
type radarAs112TimeseryDnssecListResponseResultSerie0JSON struct {
	NotSupported apijson.Field
	Supported    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarAs112TimeseryDnssecListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAs112TimeseryDnssecListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAs112TimeseryDnssecListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAs112TimeseryDnssecListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAs112TimeseryDnssecListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAs112TimeseryDnssecListParams]'s query parameters as
// `url.Values`.
func (r RadarAs112TimeseryDnssecListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAs112TimeseryDnssecListParamsAggInterval string

const (
	RadarAs112TimeseryDnssecListParamsAggInterval15m RadarAs112TimeseryDnssecListParamsAggInterval = "15m"
	RadarAs112TimeseryDnssecListParamsAggInterval1h  RadarAs112TimeseryDnssecListParamsAggInterval = "1h"
	RadarAs112TimeseryDnssecListParamsAggInterval1d  RadarAs112TimeseryDnssecListParamsAggInterval = "1d"
	RadarAs112TimeseryDnssecListParamsAggInterval1w  RadarAs112TimeseryDnssecListParamsAggInterval = "1w"
)

type RadarAs112TimeseryDnssecListParamsDateRange string

const (
	RadarAs112TimeseryDnssecListParamsDateRange1d         RadarAs112TimeseryDnssecListParamsDateRange = "1d"
	RadarAs112TimeseryDnssecListParamsDateRange7d         RadarAs112TimeseryDnssecListParamsDateRange = "7d"
	RadarAs112TimeseryDnssecListParamsDateRange14d        RadarAs112TimeseryDnssecListParamsDateRange = "14d"
	RadarAs112TimeseryDnssecListParamsDateRange28d        RadarAs112TimeseryDnssecListParamsDateRange = "28d"
	RadarAs112TimeseryDnssecListParamsDateRange12w        RadarAs112TimeseryDnssecListParamsDateRange = "12w"
	RadarAs112TimeseryDnssecListParamsDateRange24w        RadarAs112TimeseryDnssecListParamsDateRange = "24w"
	RadarAs112TimeseryDnssecListParamsDateRange52w        RadarAs112TimeseryDnssecListParamsDateRange = "52w"
	RadarAs112TimeseryDnssecListParamsDateRange1dControl  RadarAs112TimeseryDnssecListParamsDateRange = "1dControl"
	RadarAs112TimeseryDnssecListParamsDateRange7dControl  RadarAs112TimeseryDnssecListParamsDateRange = "7dControl"
	RadarAs112TimeseryDnssecListParamsDateRange14dControl RadarAs112TimeseryDnssecListParamsDateRange = "14dControl"
	RadarAs112TimeseryDnssecListParamsDateRange28dControl RadarAs112TimeseryDnssecListParamsDateRange = "28dControl"
	RadarAs112TimeseryDnssecListParamsDateRange12wControl RadarAs112TimeseryDnssecListParamsDateRange = "12wControl"
	RadarAs112TimeseryDnssecListParamsDateRange24wControl RadarAs112TimeseryDnssecListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAs112TimeseryDnssecListParamsFormat string

const (
	RadarAs112TimeseryDnssecListParamsFormatJson RadarAs112TimeseryDnssecListParamsFormat = "JSON"
	RadarAs112TimeseryDnssecListParamsFormatCsv  RadarAs112TimeseryDnssecListParamsFormat = "CSV"
)
