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

// RadarAttackLayer3TimeseriesGroupIndustryService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarAttackLayer3TimeseriesGroupIndustryService] method instead.
type RadarAttackLayer3TimeseriesGroupIndustryService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer3TimeseriesGroupIndustryService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer3TimeseriesGroupIndustryService(opts ...option.RequestOption) (r *RadarAttackLayer3TimeseriesGroupIndustryService) {
	r = &RadarAttackLayer3TimeseriesGroupIndustryService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by industry used over time.
func (r *RadarAttackLayer3TimeseriesGroupIndustryService) List(ctx context.Context, query RadarAttackLayer3TimeseriesGroupIndustryListParams, opts ...option.RequestOption) (res *RadarAttackLayer3TimeseriesGroupIndustryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarAttackLayer3TimeseriesGroupIndustryListResponseEnvelope
	path := "radar/attacks/layer3/timeseries_groups/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarAttackLayer3TimeseriesGroupIndustryListResponse struct {
	Meta   interface{}                                                `json:"meta,required"`
	Serie0 RadarAttackLayer3TimeseriesGroupIndustryListResponseSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer3TimeseriesGroupIndustryListResponseJSON   `json:"-"`
}

// radarAttackLayer3TimeseriesGroupIndustryListResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer3TimeseriesGroupIndustryListResponse]
type radarAttackLayer3TimeseriesGroupIndustryListResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupIndustryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupIndustryListResponseSerie0 struct {
	Timestamps  []string                                                       `json:"timestamps,required"`
	ExtraFields map[string][]string                                            `json:"-,extras"`
	JSON        radarAttackLayer3TimeseriesGroupIndustryListResponseSerie0JSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupIndustryListResponseSerie0JSON contains the JSON
// metadata for the struct
// [RadarAttackLayer3TimeseriesGroupIndustryListResponseSerie0]
type radarAttackLayer3TimeseriesGroupIndustryListResponseSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupIndustryListResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer3TimeseriesGroupIndustryListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer3TimeseriesGroupIndustryListParamsAggInterval] `query:"aggInterval"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Together with the `location` parameter, will apply the filter to origin or
	// target location.
	Direction param.Field[RadarAttackLayer3TimeseriesGroupIndustryListParamsDirection] `query:"direction"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer3TimeseriesGroupIndustryListParamsFormat] `query:"format"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer3TimeseriesGroupIndustryListParamsIPVersion] `query:"ipVersion"`
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
	Normalization param.Field[RadarAttackLayer3TimeseriesGroupIndustryListParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer3TimeseriesGroupIndustryListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer3TimeseriesGroupIndustryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer3TimeseriesGroupIndustryListParamsAggInterval string

const (
	RadarAttackLayer3TimeseriesGroupIndustryListParamsAggInterval15m RadarAttackLayer3TimeseriesGroupIndustryListParamsAggInterval = "15m"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsAggInterval1h  RadarAttackLayer3TimeseriesGroupIndustryListParamsAggInterval = "1h"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsAggInterval1d  RadarAttackLayer3TimeseriesGroupIndustryListParamsAggInterval = "1d"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsAggInterval1w  RadarAttackLayer3TimeseriesGroupIndustryListParamsAggInterval = "1w"
)

type RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange string

const (
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange1d         RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "1d"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange2d         RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "2d"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange7d         RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "7d"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange14d        RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "14d"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange28d        RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "28d"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange12w        RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "12w"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange24w        RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "24w"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange52w        RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "52w"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange1dControl  RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "1dControl"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange2dControl  RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "2dControl"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange7dControl  RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "7dControl"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange14dControl RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "14dControl"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange28dControl RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "28dControl"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange12wControl RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "12wControl"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange24wControl RadarAttackLayer3TimeseriesGroupIndustryListParamsDateRange = "24wControl"
)

// Together with the `location` parameter, will apply the filter to origin or
// target location.
type RadarAttackLayer3TimeseriesGroupIndustryListParamsDirection string

const (
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDirectionOrigin RadarAttackLayer3TimeseriesGroupIndustryListParamsDirection = "ORIGIN"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsDirectionTarget RadarAttackLayer3TimeseriesGroupIndustryListParamsDirection = "TARGET"
)

// Format results are returned in.
type RadarAttackLayer3TimeseriesGroupIndustryListParamsFormat string

const (
	RadarAttackLayer3TimeseriesGroupIndustryListParamsFormatJson RadarAttackLayer3TimeseriesGroupIndustryListParamsFormat = "JSON"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsFormatCsv  RadarAttackLayer3TimeseriesGroupIndustryListParamsFormat = "CSV"
)

type RadarAttackLayer3TimeseriesGroupIndustryListParamsIPVersion string

const (
	RadarAttackLayer3TimeseriesGroupIndustryListParamsIPVersionIPv4 RadarAttackLayer3TimeseriesGroupIndustryListParamsIPVersion = "IPv4"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsIPVersionIPv6 RadarAttackLayer3TimeseriesGroupIndustryListParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer3TimeseriesGroupIndustryListParamsNormalization string

const (
	RadarAttackLayer3TimeseriesGroupIndustryListParamsNormalizationPercentage RadarAttackLayer3TimeseriesGroupIndustryListParamsNormalization = "PERCENTAGE"
	RadarAttackLayer3TimeseriesGroupIndustryListParamsNormalizationMin0Max    RadarAttackLayer3TimeseriesGroupIndustryListParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer3TimeseriesGroupIndustryListResponseEnvelope struct {
	Result  RadarAttackLayer3TimeseriesGroupIndustryListResponse             `json:"result,required"`
	Success bool                                                             `json:"success,required"`
	JSON    radarAttackLayer3TimeseriesGroupIndustryListResponseEnvelopeJSON `json:"-"`
}

// radarAttackLayer3TimeseriesGroupIndustryListResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer3TimeseriesGroupIndustryListResponseEnvelope]
type radarAttackLayer3TimeseriesGroupIndustryListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer3TimeseriesGroupIndustryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
