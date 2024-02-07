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

// RadarAttackLayer3TimeseriesGroupVectorService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer3TimeseriesGroupVectorService] method instead.
type RadarAttackLayer3TimeseriesGroupVectorService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TimeseriesGroupVectorService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer3TimeseriesGroupVectorService(opts ...option.RequestOption) (r *RadarAttackLayer3TimeseriesGroupVectorService) {
	r = &RadarAttackLayer3TimeseriesGroupVectorService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by vector used over time.
func (r *RadarAttackLayer3TimeseriesGroupVectorService) List(ctx context.Context, query RadarAttackLayer3TimeseriesGroupVectorListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupVectorListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/vector"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3TimeseriesGroupVectorListResponse struct {
	Result  RadarAttackLayer3TimeseriesGroupVectorListResponseResult `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupVectorListResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupVectorListResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupVectorListResponse]
type radarAttackLayer3TimeseriesGroupVectorListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupVectorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupVectorListResponseResult struct {
	Meta   interface{}                                                    `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupVectorListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupVectorListResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupVectorListResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupVectorListResponseResult]
type radarAttackLayer3TimeseriesGroupVectorListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupVectorListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupVectorListResponseResultSerie0 struct {
	Timestamps []string                                                           `json:"timestamps,required"`
	JSON       radarAttackLayer3TimeseriesGroupVectorListResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupVectorListResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupVectorListResponseResultSerie0]
type radarAttackLayer3TimeseriesGroupVectorListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupVectorListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupVectorListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupVectorListParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupVectorListParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupVectorListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupVectorListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupVectorListParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupVectorListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupVectorListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupVectorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupVectorListParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupVectorListParamsAggInterval15m RadarAttackLayer3TimeseriesGroupVectorListParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupVectorListParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupVectorListParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupVectorListParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupVectorListParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupVectorListParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupVectorListParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange1d         RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange2d         RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange7d         RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange14d        RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange28d        RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange12w        RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange24w        RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange52w        RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupVectorListParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupVectorListParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupVectorListParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupVectorListParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupVectorListParamsDirectionTarget RadarAttackLayer3TimeseriesGroupVectorListParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupVectorListParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupVectorListParamsFormatJson RadarAttackLayer3TimeseriesGroupVectorListParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupVectorListParamsFormatCsv  RadarAttackLayer3TimeseriesGroupVectorListParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupVectorListParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupVectorListParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupVectorListParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupVectorListParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupVectorListParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupVectorListParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupVectorListParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupVectorListParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupVectorListParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupVectorListParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupVectorListParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupVectorListParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupVectorListParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupVectorListParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupVectorListParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupVectorListParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupVectorListParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupVectorListParamsProtocolGre  RadarAttackLayer3TimeseriesGroupVectorListParamsProtocol = "GRE"
)
