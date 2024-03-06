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

// RadarAttackLayer3Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarAttackLayer3Service] method
// instead.
type RadarAttackLayer3Service struct {
	Options          []option.RequestOption
	Summary          *RadarAttackLayer3SummaryService
	TimeseriesGroups *RadarAttackLayer3TimeseriesGroupService
	Top              *RadarAttackLayer3TopService
}

// NewRadarAttackLayer3Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarAttackLayer3Service(opts ...option.RequestOption) (r *RadarAttackLayer3Service) {
	r = &RadarAttackLayer3Service{}
	r.Options = opts
	r.Summary = NewRadarAttackLayer3SummaryService(opts...)
	r.TimeseriesGroups = NewRadarAttackLayer3TimeseriesGroupService(opts...)
	r.Top = NewRadarAttackLayer3TopService(opts...)
	return
}

// Get attacks change over time by bytes.
func (r *RadarAttackLayer3Service) Timeseries(ctx context.Context, query RadarAttackLayer3TimeseriesParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TimeseriesResponseEnvelope
	path := "radar/attacks/layer3/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer3TimeseriesResponse struct {
	Meta   interface{}                               `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer3TimeseriesResponse]
type radarAttackLayer3TimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesResponseSerie0 struct {
	Timestamps []string                                      `json:"timestamps,required"`
	Values     []string                                      `json:"values,required"`
	JSON       radarAttackLayer3TimeseriesResponseSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesResponseSerie0JSON contains the JSON metadata for the
// struct [RadarAttackLayer3TimeseriesResponseSerie0]
type radarAttackLayer3TimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type RadarAttackLayer3TimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Measurement units, eg. bytes.
	Metric param.Field[RadarAttackLayer3TimeseriesParamsMetric] `query:"metric"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer3TimeseriesParamsNormalization] `query:"normalization"`
	// Array of L3/4 attack types.
	Protocol param.Field[[]RadarAttackLayer3TimeseriesParamsProtocol] `query:"protocol"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer3TimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesParamsAggInterval15m RadarAttackLayer3TimeseriesParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesParamsAggInterval1h  RadarAttackLayer3TimeseriesParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesParamsAggInterval1d  RadarAttackLayer3TimeseriesParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesParamsAggInterval1w  RadarAttackLayer3TimeseriesParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesParamsDateRange string

const (
	RadarAttackLayer3TimeseriesParamsDateRange1d         RadarAttackLayer3TimeseriesParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesParamsDateRange2d         RadarAttackLayer3TimeseriesParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesParamsDateRange7d         RadarAttackLayer3TimeseriesParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesParamsDateRange14d        RadarAttackLayer3TimeseriesParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesParamsDateRange28d        RadarAttackLayer3TimeseriesParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesParamsDateRange12w        RadarAttackLayer3TimeseriesParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesParamsDateRange24w        RadarAttackLayer3TimeseriesParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesParamsDateRange52w        RadarAttackLayer3TimeseriesParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesParamsDateRange1dControl  RadarAttackLayer3TimeseriesParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesParamsDateRange2dControl  RadarAttackLayer3TimeseriesParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesParamsDateRange7dControl  RadarAttackLayer3TimeseriesParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesParamsDateRange14dControl RadarAttackLayer3TimeseriesParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesParamsDateRange28dControl RadarAttackLayer3TimeseriesParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesParamsDateRange12wControl RadarAttackLayer3TimeseriesParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesParamsDateRange24wControl RadarAttackLayer3TimeseriesParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesParamsDirection string

const (
	RadarAttackLayer3TimeseriesParamsDirectionOrigin RadarAttackLayer3TimeseriesParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesParamsDirectionTarget RadarAttackLayer3TimeseriesParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesParamsFormat string

const (
	RadarAttackLayer3TimeseriesParamsFormatJson RadarAttackLayer3TimeseriesParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesParamsFormatCsv  RadarAttackLayer3TimeseriesParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesParamsIPVersionIPv4 RadarAttackLayer3TimeseriesParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesParamsIPVersionIPv6 RadarAttackLayer3TimeseriesParamsIPVersion = "IPv6"
)

// Measurement units, eg. bytes.
type RadarAttackLayer3TimeseriesParamsMetric string

const (
	RadarAttackLayer3TimeseriesParamsMetricBytes    RadarAttackLayer3TimeseriesParamsMetric = "BYTES"
	RadarAttackLayer3TimeseriesParamsMetricBytesOld RadarAttackLayer3TimeseriesParamsMetric = "BYTES_OLD"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesParamsNormalization string

const (
	RadarAttackLayer3TimeseriesParamsNormalizationPercentageChange RadarAttackLayer3TimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	RadarAttackLayer3TimeseriesParamsNormalizationMin0Max          RadarAttackLayer3TimeseriesParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesParamsProtocol string

const (
	RadarAttackLayer3TimeseriesParamsProtocolUdp  RadarAttackLayer3TimeseriesParamsProtocol = "UDP"
	RadarAttackLayer3TimeseriesParamsProtocolTcp  RadarAttackLayer3TimeseriesParamsProtocol = "TCP"
	RadarAttackLayer3TimeseriesParamsProtocolIcmp RadarAttackLayer3TimeseriesParamsProtocol = "ICMP"
	RadarAttackLayer3TimeseriesParamsProtocolGRE  RadarAttackLayer3TimeseriesParamsProtocol = "GRE"
)

type RadarAttackLayer3TimeseriesResponseEnvelope struct {
	Result  RadarAttackLayer3TimeseriesResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarAttackLayer3TimeseriesResponseEnvelope]
type radarAttackLayer3TimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarAttackLayer3TimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
