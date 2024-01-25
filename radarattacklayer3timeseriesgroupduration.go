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

// RadarAttackLayer3TimeseriesGroupDurationService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer3TimeseriesGroupDurationService] method instead.
type RadarAttackLayer3TimeseriesGroupDurationService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TimeseriesGroupDurationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer3TimeseriesGroupDurationService(opts ...option.RequestOption) (r *RadarAttackLayer3TimeseriesGroupDurationService) {
	r = &RadarAttackLayer3TimeseriesGroupDurationService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by duration over time.
func (r *RadarAttackLayer3TimeseriesGroupDurationService) List(ctx context.Context, query RadarAttackLayer3TimeseriesGroupDurationListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupDurationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer3/timeseries_groups/duration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer3TimeseriesGroupDurationListResponse struct {
	Result  RadarAttackLayer3TimeseriesGroupDurationListResponseResult `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupDurationListResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupDurationListResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupDurationListResponse]
type radarAttackLayer3TimeseriesGroupDurationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupDurationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupDurationListResponseResult struct {
	Meta   interface{}                                                      `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupDurationListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupDurationListResponseResultJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupDurationListResponseResultJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupDurationListResponseResult]
type radarAttackLayer3TimeseriesGroupDurationListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupDurationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupDurationListResponseResultSerie0 struct {
	Number1HourTo3Hours  []string                                                             `json:"_1_HOUR_TO_3_HOURS,required"`
	Number10MinsTo20Mins []string                                                             `json:"_10_MINS_TO_20_MINS,required"`
	Number20MinsTo40Mins []string                                                             `json:"_20_MINS_TO_40_MINS,required"`
	Number40MinsTo1Hour  []string                                                             `json:"_40_MINS_TO_1_HOUR,required"`
	Over3Hours           []string                                                             `json:"OVER_3_HOURS,required"`
	Timestamps           []string                                                             `json:"timestamps,required"`
	Under10Mins          []string                                                             `json:"UNDER_10_MINS,required"`
	JSON                 radarAttackLayer3TimeseriesGroupDurationListResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupDurationListResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupDurationListResponseResultSerie0]
type radarAttackLayer3TimeseriesGroupDurationListResponseResultSerie0JSON struct {
	Number1HourTo3Hours  apijson.Field
	Number10MinsTo20Mins apijson.Field
	Number20MinsTo40Mins apijson.Field
	Number40MinsTo1Hour  apijson.Field
	Over3Hours           apijson.Field
	Timestamps           apijson.Field
	Under10Mins          apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupDurationListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupDurationListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupDurationListParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupDurationListParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupDurationListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupDurationListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupDurationListParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupDurationListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupDurationListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupDurationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupDurationListParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupDurationListParamsAggInterval15m RadarAttackLayer3TimeseriesGroupDurationListParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupDurationListParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupDurationListParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupDurationListParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupDurationListParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupDurationListParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupDurationListParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange1d         RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange2d         RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange7d         RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange14d        RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange28d        RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange12w        RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange24w        RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange52w        RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupDurationListParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupDurationListParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupDurationListParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupDurationListParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupDurationListParamsDirectionTarget RadarAttackLayer3TimeseriesGroupDurationListParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupDurationListParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupDurationListParamsFormatJson RadarAttackLayer3TimeseriesGroupDurationListParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupDurationListParamsFormatCsv  RadarAttackLayer3TimeseriesGroupDurationListParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupDurationListParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupDurationListParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupDurationListParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupDurationListParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupDurationListParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupDurationListParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupDurationListParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupDurationListParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupDurationListParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupDurationListParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupDurationListParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupDurationListParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupDurationListParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupDurationListParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupDurationListParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupDurationListParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupDurationListParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupDurationListParamsProtocolGre  RadarAttackLayer3TimeseriesGroupDurationListParamsProtocol = "GRE"
)
