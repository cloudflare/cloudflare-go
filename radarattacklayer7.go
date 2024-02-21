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

// RadarAttackLayer7Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAttackLayer7Service] method
// instead.
type RadarAttackLayer7Service struct {
	Options          []option.RequestOption
	Summary          *RadarAttackLayer7SummaryService
	TimeseriesGroups *RadarAttackLayer7TimeseriesGroupService
	Top              *RadarAttackLayer7TopService
}

// NewRadarAttackLayer7Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7Service(opts ...option.RequestOption) (r *RadarAttackLayer7Service) {
	r = &RadarAttackLayer7Service{}
	r.Options = opts
	r.Summary = NewRadarAttackLayer7SummaryService(opts...)
	r.TimeseriesGroups = NewRadarAttackLayer7TimeseriesGroupService(opts...)
	r.Top = NewRadarAttackLayer7TopService(opts...)
	return
}

// Get attacks change over time by bytes.
func (r *RadarAttackLayer7Service) Timeseries(ctx context.Context, query RadarAttackLayer7TimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer7TimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer7TimeseriesResponseEnvelope
	path := "radar/attacks/layer3/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer7TimeseriesResponse struct {
	Meta   interface{}                               `json:"meta,required"`
	Serie0 RadarAttackLayer7TimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7TimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer7TimeseriesResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7TimeseriesResponse]
type radarAttackLayer7TimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesResponseSerie0 struct {
	Timestamps []string                                      `json:"timestamps,required"`
	Values     []string                                      `json:"values,required"`
	JSON       radarAttackLayer7TimeseriesResponseSerie0JSON `json:"-"`
}

// radarAttackLayer7TimeseriesResponseSerie0JSON contains the JSON metadata for the
// struct [RadarAttackLayer7TimeseriesResponseSerie0]
type radarAttackLayer7TimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7TimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7TimeseriesParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7TimeseriesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer7TimeseriesParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7TimeseriesParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7TimeseriesParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Measurement units, eg. bytes.
	Metric param.Field[RadarAttackLayer7TimeseriesParamsMetric] `query:"metric"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7TimeseriesParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer7TimeseriesParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer7TimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7TimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7TimeseriesParamsAggInterval string

const (
	RadarAttackLayer7TimeseriesParamsAggInterval15m RadarAttackLayer7TimeseriesParamsAggInterval = "15m"
	RadarAttackLayer7TimeseriesParamsAggInterval1h  RadarAttackLayer7TimeseriesParamsAggInterval = "1h"
	RadarAttackLayer7TimeseriesParamsAggInterval1d  RadarAttackLayer7TimeseriesParamsAggInterval = "1d"
	RadarAttackLayer7TimeseriesParamsAggInterval1w  RadarAttackLayer7TimeseriesParamsAggInterval = "1w"
)

type RadarAttackLayer7TimeseriesParamsDateRange string

const (
	RadarAttackLayer7TimeseriesParamsDateRange1d         RadarAttackLayer7TimeseriesParamsDateRange = "1d"
	RadarAttackLayer7TimeseriesParamsDateRange2d         RadarAttackLayer7TimeseriesParamsDateRange = "2d"
	RadarAttackLayer7TimeseriesParamsDateRange7d         RadarAttackLayer7TimeseriesParamsDateRange = "7d"
	RadarAttackLayer7TimeseriesParamsDateRange14d        RadarAttackLayer7TimeseriesParamsDateRange = "14d"
	RadarAttackLayer7TimeseriesParamsDateRange28d        RadarAttackLayer7TimeseriesParamsDateRange = "28d"
	RadarAttackLayer7TimeseriesParamsDateRange12w        RadarAttackLayer7TimeseriesParamsDateRange = "12w"
	RadarAttackLayer7TimeseriesParamsDateRange24w        RadarAttackLayer7TimeseriesParamsDateRange = "24w"
	RadarAttackLayer7TimeseriesParamsDateRange52w        RadarAttackLayer7TimeseriesParamsDateRange = "52w"
	RadarAttackLayer7TimeseriesParamsDateRange1dControl  RadarAttackLayer7TimeseriesParamsDateRange = "1dControl"
	RadarAttackLayer7TimeseriesParamsDateRange2dControl  RadarAttackLayer7TimeseriesParamsDateRange = "2dControl"
	RadarAttackLayer7TimeseriesParamsDateRange7dControl  RadarAttackLayer7TimeseriesParamsDateRange = "7dControl"
	RadarAttackLayer7TimeseriesParamsDateRange14dControl RadarAttackLayer7TimeseriesParamsDateRange = "14dControl"
	RadarAttackLayer7TimeseriesParamsDateRange28dControl RadarAttackLayer7TimeseriesParamsDateRange = "28dControl"
	RadarAttackLayer7TimeseriesParamsDateRange12wControl RadarAttackLayer7TimeseriesParamsDateRange = "12wControl"
	RadarAttackLayer7TimeseriesParamsDateRange24wControl RadarAttackLayer7TimeseriesParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer7TimeseriesParamsDirection string

const (
	RadarAttackLayer7TimeseriesParamsDirectionOrigin RadarAttackLayer7TimeseriesParamsDirection = "ORIGIN"
	RadarAttackLayer7TimeseriesParamsDirectionTarget RadarAttackLayer7TimeseriesParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer7TimeseriesParamsFormat string

const (
	RadarAttackLayer7TimeseriesParamsFormatJson RadarAttackLayer7TimeseriesParamsFormat = "JSON"
	RadarAttackLayer7TimeseriesParamsFormatCsv  RadarAttackLayer7TimeseriesParamsFormat = "CSV"
)

type RadarAttackLayer7TimeseriesParamsIPVersion string

const (
	RadarAttackLayer7TimeseriesParamsIPVersionIPv4 RadarAttackLayer7TimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer7TimeseriesParamsIPVersionIPv6 RadarAttackLayer7TimeseriesParamsIPVersion = "IPv6"
)

// Measurement units, eg. bytes.
type RadarAttackLayer7TimeseriesParamsMetric string

const (
	RadarAttackLayer7TimeseriesParamsMetricBytes    RadarAttackLayer7TimeseriesParamsMetric = "BYTES"
	RadarAttackLayer7TimeseriesParamsMetricBytesOld RadarAttackLayer7TimeseriesParamsMetric = "BYTES_OLD"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7TimeseriesParamsNormalization string

const (
	RadarAttackLayer7TimeseriesParamsNormalizationPercentageChange RadarAttackLayer7TimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer7TimeseriesParamsNormalizationMin0Max          RadarAttackLayer7TimeseriesParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7TimeseriesParamsProtocol string

const (
	RadarAttackLayer7TimeseriesParamsProtocolUdp  RadarAttackLayer7TimeseriesParamsProtocol = "UDP"
	RadarAttackLayer7TimeseriesParamsProtocolTcp  RadarAttackLayer7TimeseriesParamsProtocol = "TCP"
	RadarAttackLayer7TimeseriesParamsProtocolIcmp RadarAttackLayer7TimeseriesParamsProtocol = "ICMP"
	RadarAttackLayer7TimeseriesParamsProtocolGre  RadarAttackLayer7TimeseriesParamsProtocol = "GRE"
)

type RadarAttackLayer7TimeseriesResponseEnvelope struct {
	Result  RadarAttackLayer7TimeseriesResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer7TimeseriesResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer7TimeseriesResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAttackLayer7TimeseriesResponseEnvelope]
type radarAttackLayer7TimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7TimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
