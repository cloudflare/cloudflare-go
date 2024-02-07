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

// RadarAttackLayer3TimeseriesGroupProtocolService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer3TimeseriesGroupProtocolService] method instead.
type RadarAttackLayer3TimeseriesGroupProtocolService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TimeseriesGroupProtocolService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer3TimeseriesGroupProtocolService(opts ...option.RequestOption) (r *RadarAttackLayer3TimeseriesGroupProtocolService) {
	r = &RadarAttackLayer3TimeseriesGroupProtocolService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by protocol used over time.
func (r *RadarAttackLayer3TimeseriesGroupProtocolService) List(ctx context.Context, query RadarAttackLayer3TimeseriesGroupProtocolListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupProtocolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TimeseriesGroupProtocolListResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer3TimeseriesGroupProtocolListResponse struct {
	Meta   interface{}                                                `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupProtocolListResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupProtocolListResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupProtocolListResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupProtocolListResponse]
type radarAttackLayer3TimeseriesGroupProtocolListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupProtocolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupProtocolListResponseSerie0 struct {
	Gre        []string                                                       `json:"GRE,required"`
	Icmp       []string                                                       `json:"ICMP,required"`
	Tcp        []string                                                       `json:"TCP,required"`
	Timestamps []string                                                       `json:"timestamps,required"`
	Udp        []string                                                       `json:"UDP,required"`
	JSON       radarAttackLayer3TimeseriesGroupProtocolListResponseSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupProtocolListResponseSerie0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupProtocolListResponseSerie0]
type radarAttackLayer3TimeseriesGroupProtocolListResponseSerie0JSON struct {
	Gre         apijson.Field
	Icmp        apijson.Field
	Tcp         apijson.Field
	Timestamps  apijson.Field
	Udp         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupProtocolListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupProtocolListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupProtocolListParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupProtocolListParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupProtocolListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupProtocolListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupProtocolListParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupProtocolListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupProtocolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupProtocolListParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupProtocolListParamsAggInterval15m RadarAttackLayer3TimeseriesGroupProtocolListParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupProtocolListParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupProtocolListParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupProtocolListParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange1d         RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange2d         RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange7d         RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange14d        RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange28d        RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange12w        RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange24w        RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange52w        RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupProtocolListParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupProtocolListParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupProtocolListParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsDirectionTarget RadarAttackLayer3TimeseriesGroupProtocolListParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupProtocolListParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupProtocolListParamsFormatJson RadarAttackLayer3TimeseriesGroupProtocolListParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsFormatCsv  RadarAttackLayer3TimeseriesGroupProtocolListParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupProtocolListParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupProtocolListParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupProtocolListParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupProtocolListParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupProtocolListParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupProtocolListParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupProtocolListParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupProtocolListParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupProtocolListParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupProtocolListResponseEnvelope struct {
	Result  RadarAttackLayer3TimeseriesGroupProtocolListResponse             `json:"result,required"`
	Success bool                                                             `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupProtocolListResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupProtocolListResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupProtocolListResponseEnvelope]
type radarAttackLayer3TimeseriesGroupProtocolListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupProtocolListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
