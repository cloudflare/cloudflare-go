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

// Get attacks change over time by bytes.
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
	Meta   interface{}                                       `json:"meta,required"`
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

type RadarAttackLayer3TimeseryListResponseResultSerie0 struct {
	Timestamps []string                                              `json:"timestamps,required"`
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
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseryListParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseryListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseryListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Measurement units, eg. bytes.
	Metric param.Field[RadarAttackLayer3TimeseryListParamsMetric] `query:"metric"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseryListParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TimeseryListParamsProtocol] `query:"protocol"`
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

type RadarAttackLayer3TimeseryListParamsDateRange string

const (
	RadarAttackLayer3TimeseryListParamsDateRange1d         RadarAttackLayer3TimeseryListParamsDateRange = "1d"
	RadarAttackLayer3TimeseryListParamsDateRange2d         RadarAttackLayer3TimeseryListParamsDateRange = "2d"
	RadarAttackLayer3TimeseryListParamsDateRange7d         RadarAttackLayer3TimeseryListParamsDateRange = "7d"
	RadarAttackLayer3TimeseryListParamsDateRange14d        RadarAttackLayer3TimeseryListParamsDateRange = "14d"
	RadarAttackLayer3TimeseryListParamsDateRange28d        RadarAttackLayer3TimeseryListParamsDateRange = "28d"
	RadarAttackLayer3TimeseryListParamsDateRange12w        RadarAttackLayer3TimeseryListParamsDateRange = "12w"
	RadarAttackLayer3TimeseryListParamsDateRange24w        RadarAttackLayer3TimeseryListParamsDateRange = "24w"
	RadarAttackLayer3TimeseryListParamsDateRange52w        RadarAttackLayer3TimeseryListParamsDateRange = "52w"
	RadarAttackLayer3TimeseryListParamsDateRange1dControl  RadarAttackLayer3TimeseryListParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseryListParamsDateRange2dControl  RadarAttackLayer3TimeseryListParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseryListParamsDateRange7dControl  RadarAttackLayer3TimeseryListParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseryListParamsDateRange14dControl RadarAttackLayer3TimeseryListParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseryListParamsDateRange28dControl RadarAttackLayer3TimeseryListParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseryListParamsDateRange12wControl RadarAttackLayer3TimeseryListParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseryListParamsDateRange24wControl RadarAttackLayer3TimeseryListParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseryListParamsDirection string

const (
	RadarAttackLayer3TimeseryListParamsDirectionOrigin RadarAttackLayer3TimeseryListParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseryListParamsDirectionTarget RadarAttackLayer3TimeseryListParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseryListParamsFormat string

const (
	RadarAttackLayer3TimeseryListParamsFormatJson RadarAttackLayer3TimeseryListParamsFormat = "JSON"
	RadarAttackLayer3TimeseryListParamsFormatCsv  RadarAttackLayer3TimeseryListParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseryListParamsIPVersion string

const (
	RadarAttackLayer3TimeseryListParamsIPVersionIPv4 RadarAttackLayer3TimeseryListParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseryListParamsIPVersionIPv6 RadarAttackLayer3TimeseryListParamsIPVersion = "IPv6"
)

// Measurement units, eg. bytes.
type RadarAttackLayer3TimeseryListParamsMetric string

const (
	RadarAttackLayer3TimeseryListParamsMetricBytes    RadarAttackLayer3TimeseryListParamsMetric = "BYTES"
	RadarAttackLayer3TimeseryListParamsMetricBytesOld RadarAttackLayer3TimeseryListParamsMetric = "BYTES_OLD"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseryListParamsNormalization string

const (
	RadarAttackLayer3TimeseryListParamsNormalizationPercentageChange RadarAttackLayer3TimeseryListParamsNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3TimeseryListParamsNormalizationMin0Max          RadarAttackLayer3TimeseryListParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseryListParamsProtocol string

const (
	RadarAttackLayer3TimeseryListParamsProtocolUdp  RadarAttackLayer3TimeseryListParamsProtocol = "UDP"
	RadarAttackLayer3TimeseryListParamsProtocolTcp  RadarAttackLayer3TimeseryListParamsProtocol = "TCP"
	RadarAttackLayer3TimeseryListParamsProtocolIcmp RadarAttackLayer3TimeseryListParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseryListParamsProtocolGre  RadarAttackLayer3TimeseryListParamsProtocol = "GRE"
)
