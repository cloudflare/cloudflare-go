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

// RadarAttackLayer3TimeseryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer3TimeseryService] method instead.
type RadarAttackLayer3TimeseryService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TimeseryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3TimeseryService(opts ...option.RequestOption) (r *RadarAttackLayer3TimeseryService) {
	r = &RadarAttackLayer3TimeseryService{}
	r.Options = opts
	return
}

// Get layer 3/4 attacks change over time. Values are normalized using min-max by
// default, with the minimum set to 0. When asking for multiple time series, you
// can also get the percentual relative change of the 1st/main series, with respect
// to the 2nd/control series - for example, to get the relative change of this week
// from the previous week, the 1st series would have a date range of 7d, the 2nd, a
// date range of 7dControl, and the normalization would be set to
// PERCENTAGE_CHANGE.
func (r *RadarAttackLayer3TimeseryService) List(ctx context.Context, query RadarAttackLayer3TimeseryListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3TimeseryListResponse struct {
	Result  RadarAttackLayer3TimeseryListResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarAttackLayer3TimeseryListResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseryListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3TimeseryListResponse]
type radarAttackLayer3TimeseryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseryListResponseResult struct {
	Meta   RadarAttackLayer3TimeseryListResponseResultMeta   `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseryListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseryListResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseryListResponseResultJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TimeseryListResponseResult]
type radarAttackLayer3TimeseryListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseryListResponseResultMeta struct {
	AggInterval    string                                                        `json:"aggInterval,required"`
	ConfidenceInfo RadarAttackLayer3TimeseryListResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarAttackLayer3TimeseryListResponseResultMetaDateRange      `json:"dateRange,required"`
	LastUpdated    time.Time                                                     `json:"lastUpdated,required" format:"date-time"`
	JSON           radarAttackLayer3TimeseryListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer3TimeseryListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer3TimeseryListResponseResultMeta]
type radarAttackLayer3TimeseryListResponseResultMetaJSON struct {
	AggInterval    apijson.Field
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseryListResponseResultMetaConfidenceInfo struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                         `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TimeseryListResponseResultMetaConfidenceInfoJSON `json:"-"`
}

// radarAttackLayer3TimeseryListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseryListResponseResultMetaConfidenceInfo]
type radarAttackLayer3TimeseryListResponseResultMetaConfidenceInfoJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer3TimeseryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer3TimeseryListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseryListResponseResultMetaDateRange]
type radarAttackLayer3TimeseryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseryListResponseResultSerie0 struct {
	Timestamps []time.Time                                           `json:"timestamps,required" format:"date-time"`
	Values     []string                                              `json:"values,required"`
	JSON       radarAttackLayer3TimeseryListResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseryListResponseResultSerie0JSON contains the JSON metadata
// for the struct [RadarAttackLayer3TimeseryListResponseResultSerie0]
type radarAttackLayer3TimeseryListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseryListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseryListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of L3/4 attack types.
	Attack param.Field[[]RadarAttackLayer3TimeseryListParamsAttack] `query:"attack"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseryListParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer3TimeseryListParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3TimeseryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseryListParamsAggInterval string

const (
	RadarAttackLayer3TimeseryListParamsAggInterval15m RadarAttackLayer3TimeseryListParamsAggInterval = "15m"
	RadarAttackLayer3TimeseryListParamsAggInterval1h  RadarAttackLayer3TimeseryListParamsAggInterval = "1h"
	RadarAttackLayer3TimeseryListParamsAggInterval1d  RadarAttackLayer3TimeseryListParamsAggInterval = "1d"
	RadarAttackLayer3TimeseryListParamsAggInterval1w  RadarAttackLayer3TimeseryListParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseryListParamsAttack string

const (
	RadarAttackLayer3TimeseryListParamsAttackUdp  RadarAttackLayer3TimeseryListParamsAttack = "UDP"
	RadarAttackLayer3TimeseryListParamsAttackTcp  RadarAttackLayer3TimeseryListParamsAttack = "TCP"
	RadarAttackLayer3TimeseryListParamsAttackIcmp RadarAttackLayer3TimeseryListParamsAttack = "ICMP"
	RadarAttackLayer3TimeseryListParamsAttackGre  RadarAttackLayer3TimeseryListParamsAttack = "GRE"
)

type RadarAttackLayer3TimeseryListParamsDateRange string

const (
	RadarAttackLayer3TimeseryListParamsDateRange1d         RadarAttackLayer3TimeseryListParamsDateRange = "1d"
	RadarAttackLayer3TimeseryListParamsDateRange7d         RadarAttackLayer3TimeseryListParamsDateRange = "7d"
	RadarAttackLayer3TimeseryListParamsDateRange14d        RadarAttackLayer3TimeseryListParamsDateRange = "14d"
	RadarAttackLayer3TimeseryListParamsDateRange28d        RadarAttackLayer3TimeseryListParamsDateRange = "28d"
	RadarAttackLayer3TimeseryListParamsDateRange12w        RadarAttackLayer3TimeseryListParamsDateRange = "12w"
	RadarAttackLayer3TimeseryListParamsDateRange24w        RadarAttackLayer3TimeseryListParamsDateRange = "24w"
	RadarAttackLayer3TimeseryListParamsDateRange52w        RadarAttackLayer3TimeseryListParamsDateRange = "52w"
	RadarAttackLayer3TimeseryListParamsDateRange1dControl  RadarAttackLayer3TimeseryListParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseryListParamsDateRange7dControl  RadarAttackLayer3TimeseryListParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseryListParamsDateRange14dControl RadarAttackLayer3TimeseryListParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseryListParamsDateRange28dControl RadarAttackLayer3TimeseryListParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseryListParamsDateRange12wControl RadarAttackLayer3TimeseryListParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseryListParamsDateRange24wControl RadarAttackLayer3TimeseryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer3TimeseryListParamsFormat string

const (
	RadarAttackLayer3TimeseryListParamsFormatJson RadarAttackLayer3TimeseryListParamsFormat = "JSON"
	RadarAttackLayer3TimeseryListParamsFormatCsv  RadarAttackLayer3TimeseryListParamsFormat = "CSV"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseryListParamsNormalization string

const (
	RadarAttackLayer3TimeseryListParamsNormalizationPercentageChange RadarAttackLayer3TimeseryListParamsNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3TimeseryListParamsNormalizationMin0Max          RadarAttackLayer3TimeseryListParamsNormalization = "MIN0_MAX"
)
