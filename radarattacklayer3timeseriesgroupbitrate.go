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

// RadarAttackLayer3TimeseriesGroupBitrateService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer3TimeseriesGroupBitrateService] method instead.
type RadarAttackLayer3TimeseriesGroupBitrateService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TimeseriesGroupBitrateService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer3TimeseriesGroupBitrateService(opts ...option.RequestOption) (r *RadarAttackLayer3TimeseriesGroupBitrateService) {
	r = &RadarAttackLayer3TimeseriesGroupBitrateService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by bitrate over time.
func (r *RadarAttackLayer3TimeseriesGroupBitrateService) List(ctx context.Context, query RadarAttackLayer3TimeseriesGroupBitrateListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupBitrateListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/bitrate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3TimeseriesGroupBitrateListResponse struct {
	Result  RadarAttackLayer3TimeseriesGroupBitrateListResponseResult `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupBitrateListResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupBitrateListResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupBitrateListResponse]
type radarAttackLayer3TimeseriesGroupBitrateListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupBitrateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupBitrateListResponseResult struct {
	Meta   interface{}                                                     `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupBitrateListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupBitrateListResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupBitrateListResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupBitrateListResponseResult]
type radarAttackLayer3TimeseriesGroupBitrateListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupBitrateListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupBitrateListResponseResultSerie0 struct {
	_1GbpsTo10Gbps   []string                                                            `json:"_1_GBPS_TO_10_GBPS,required"`
	_10GbpsTo100Gbps []string                                                            `json:"_10_GBPS_TO_100_GBPS,required"`
	_500MbpsTo1Gbps  []string                                                            `json:"_500_MBPS_TO_1_GBPS,required"`
	Over100Gbps      []string                                                            `json:"OVER_100_GBPS,required"`
	Timestamps       []string                                                            `json:"timestamps,required"`
	Under500Mbps     []string                                                            `json:"UNDER_500_MBPS,required"`
	JSON             radarAttackLayer3TimeseriesGroupBitrateListResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupBitrateListResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupBitrateListResponseResultSerie0]
type radarAttackLayer3TimeseriesGroupBitrateListResponseResultSerie0JSON struct {
	_1GbpsTo10Gbps   apijson.Field
	_10GbpsTo100Gbps apijson.Field
	_500MbpsTo1Gbps  apijson.Field
	Over100Gbps      apijson.Field
	Timestamps       apijson.Field
	Under500Mbps     apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupBitrateListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupBitrateListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupBitrateListParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupBitrateListParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupBitrateListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupBitrateListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupBitrateListParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupBitrateListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupBitrateListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupBitrateListParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupBitrateListParamsAggInterval15m RadarAttackLayer3TimeseriesGroupBitrateListParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupBitrateListParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupBitrateListParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupBitrateListParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange1d         RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange2d         RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange7d         RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange14d        RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange28d        RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange12w        RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange24w        RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange52w        RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupBitrateListParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupBitrateListParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupBitrateListParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsDirectionTarget RadarAttackLayer3TimeseriesGroupBitrateListParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupBitrateListParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupBitrateListParamsFormatJson RadarAttackLayer3TimeseriesGroupBitrateListParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsFormatCsv  RadarAttackLayer3TimeseriesGroupBitrateListParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupBitrateListParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupBitrateListParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupBitrateListParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupBitrateListParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupBitrateListParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupBitrateListParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupBitrateListParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupBitrateListParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocolGre  RadarAttackLayer3TimeseriesGroupBitrateListParamsProtocol = "GRE"
)
