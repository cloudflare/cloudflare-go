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

// RadarAttackLayer3TimeseriesGroupVerticalService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer3TimeseriesGroupVerticalService] method instead.
type RadarAttackLayer3TimeseriesGroupVerticalService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TimeseriesGroupVerticalService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer3TimeseriesGroupVerticalService(opts ...option.RequestOption) (r *RadarAttackLayer3TimeseriesGroupVerticalService) {
	r = &RadarAttackLayer3TimeseriesGroupVerticalService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by vertical used over time.
func (r *RadarAttackLayer3TimeseriesGroupVerticalService) List(ctx context.Context, query RadarAttackLayer3TimeseriesGroupVerticalListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupVerticalListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3TimeseriesGroupVerticalListResponse struct {
	Result  RadarAttackLayer3TimeseriesGroupVerticalListResponseResult `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupVerticalListResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupVerticalListResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupVerticalListResponse]
type radarAttackLayer3TimeseriesGroupVerticalListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupVerticalListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupVerticalListResponseResult struct {
	Meta   interface{}                                                      `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupVerticalListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupVerticalListResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupVerticalListResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupVerticalListResponseResult]
type radarAttackLayer3TimeseriesGroupVerticalListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupVerticalListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupVerticalListResponseResultSerie0 struct {
	Timestamps []string                                                             `json:"timestamps,required"`
	JSON       radarAttackLayer3TimeseriesGroupVerticalListResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupVerticalListResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupVerticalListResponseResultSerie0]
type radarAttackLayer3TimeseriesGroupVerticalListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupVerticalListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupVerticalListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupVerticalListParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupVerticalListParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupVerticalListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupVerticalListParamsIPVersion] `query:"ipVersion"`
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
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupVerticalListParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupVerticalListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupVerticalListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupVerticalListParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupVerticalListParamsAggInterval15m RadarAttackLayer3TimeseriesGroupVerticalListParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupVerticalListParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupVerticalListParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupVerticalListParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange1d         RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange2d         RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange7d         RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange14d        RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange28d        RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange12w        RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange24w        RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange52w        RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupVerticalListParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupVerticalListParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupVerticalListParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsDirectionTarget RadarAttackLayer3TimeseriesGroupVerticalListParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupVerticalListParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupVerticalListParamsFormatJson RadarAttackLayer3TimeseriesGroupVerticalListParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsFormatCsv  RadarAttackLayer3TimeseriesGroupVerticalListParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupVerticalListParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupVerticalListParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupVerticalListParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupVerticalListParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupVerticalListParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupVerticalListParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupVerticalListParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupVerticalListParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupVerticalListParamsNormalization = "MIN0_MAX"
)
