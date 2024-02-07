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

// RadarAttackLayer3TimeseriesGroupIPVersionService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer3TimeseriesGroupIPVersionService] method instead.
type RadarAttackLayer3TimeseriesGroupIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TimeseriesGroupIPVersionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer3TimeseriesGroupIPVersionService(opts ...option.RequestOption) (r *RadarAttackLayer3TimeseriesGroupIPVersionService) {
	r = &RadarAttackLayer3TimeseriesGroupIPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by ip version used over time.
func (r *RadarAttackLayer3TimeseriesGroupIPVersionService) List(ctx context.Context, query RadarAttackLayer3TimeseriesGroupIPVersionListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupIPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TimeseriesGroupIPVersionListResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer3TimeseriesGroupIPVersionListResponse struct {
	Meta   interface{}                                                 `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupIPVersionListResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupIPVersionListResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupIPVersionListResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupIPVersionListResponse]
type radarAttackLayer3TimeseriesGroupIPVersionListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupIPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupIPVersionListResponseSerie0 struct {
	IPv4       []string                                                        `json:"IPv4,required"`
	IPv6       []string                                                        `json:"IPv6,required"`
	Timestamps []string                                                        `json:"timestamps,required"`
	JSON       radarAttackLayer3TimeseriesGroupIPVersionListResponseSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupIPVersionListResponseSerie0JSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupIPVersionListResponseSerie0]
type radarAttackLayer3TimeseriesGroupIPVersionListResponseSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupIPVersionListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupIPVersionListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupIPVersionListParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupIPVersionListParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupIPVersionListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupIPVersionListParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupIPVersionListParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupIPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupIPVersionListParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsAggInterval15m RadarAttackLayer3TimeseriesGroupIPVersionListParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupIPVersionListParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupIPVersionListParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupIPVersionListParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange1d         RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange2d         RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange7d         RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange14d        RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange28d        RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange12w        RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange24w        RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange52w        RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupIPVersionListParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupIPVersionListParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupIPVersionListParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsDirectionTarget RadarAttackLayer3TimeseriesGroupIPVersionListParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupIPVersionListParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsFormatJson RadarAttackLayer3TimeseriesGroupIPVersionListParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsFormatCsv  RadarAttackLayer3TimeseriesGroupIPVersionListParamsFormat = "CSV"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupIPVersionListParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupIPVersionListParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupIPVersionListParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocol string

const (
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocolUdp  RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocolTcp  RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocolIcmp RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocolGre  RadarAttackLayer3TimeseriesGroupIPVersionListParamsProtocol = "GRE"
)

type RadarAttackLayer3TimeseriesGroupIPVersionListResponseEnvelope struct {
	Result  RadarAttackLayer3TimeseriesGroupIPVersionListResponse             `json:"result,required"`
	Success bool                                                              `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupIPVersionListResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupIPVersionListResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupIPVersionListResponseEnvelope]
type radarAttackLayer3TimeseriesGroupIPVersionListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupIPVersionListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
