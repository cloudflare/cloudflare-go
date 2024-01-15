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

// RadarAttackLayer7HTTPMethodService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer7HTTPMethodService] method instead.
type RadarAttackLayer7HTTPMethodService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7HTTPMethodService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7HTTPMethodService(opts ...option.RequestOption) (r *RadarAttackLayer7HTTPMethodService) {
	r = &RadarAttackLayer7HTTPMethodService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by http method used over time.
func (r *RadarAttackLayer7HTTPMethodService) ListTimeseriesGroups(ctx context.Context, query RadarAttackLayer7HTTPMethodListTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarAttackLayer7HTTPMethodListTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/http_method"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7HTTPMethodListTimeseriesGroupsResponse struct {
	Result  RadarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                                          `json:"success,required"`
	JSON    radarAttackLayer7HTTPMethodListTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarAttackLayer7HTTPMethodListTimeseriesGroupsResponseJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7HTTPMethodListTimeseriesGroupsResponse]
type radarAttackLayer7HTTPMethodListTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7HTTPMethodListTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResult struct {
	Meta   interface{}                                                         `json:"meta,required"`
	Serie0 RadarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResultJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResult]
type radarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResultSerie0 struct {
	Get        []string                                                                `json:"GET,required"`
	Timestamps []string                                                                `json:"timestamps,required"`
	JSON       radarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResultSerie0]
type radarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResultSerie0JSON struct {
	Get         apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7HTTPMethodListTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7HTTPMethodListTimeseriesGroupsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsFormat] `query:"format"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7HTTPMethodListTimeseriesGroupsParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer7HTTPMethodListTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsAggInterval string

const (
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsAggInterval15m RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsAggInterval = "15m"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsAggInterval1h  RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsAggInterval = "1h"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsAggInterval1d  RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsAggInterval = "1d"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsAggInterval1w  RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsAggInterval = "1w"
)

type RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange string

const (
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange1d         RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "1d"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange2d         RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "2d"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange7d         RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "7d"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange14d        RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "14d"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange28d        RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "28d"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange12w        RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "12w"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange24w        RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "24w"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange52w        RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "52w"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange1dControl  RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "1dControl"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange2dControl  RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "2dControl"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange7dControl  RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "7dControl"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange14dControl RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "14dControl"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange28dControl RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "28dControl"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange12wControl RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "12wControl"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange24wControl RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsFormat string

const (
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsFormatJson RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsFormat = "JSON"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsFormatCsv  RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsFormat = "CSV"
)

type RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsHTTPVersion string

const (
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsHTTPVersionHttPv1 RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsHTTPVersionHttPv2 RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsHTTPVersionHttPv3 RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsIPVersion string

const (
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsIPVersionIPv4 RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsIPVersion = "IPv4"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsIPVersionIPv6 RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsIPVersion = "IPv6"
)

type RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProduct string

const (
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProductDdos               RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProduct = "DDOS"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProductWaf                RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProduct = "WAF"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProductBotManagement      RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProductAccessRules        RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProductIPReputation       RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProductAPIShield          RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProductDataLossPrevention RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsNormalization string

const (
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsNormalizationPercentage RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsNormalizationMin0Max    RadarAttackLayer7HTTPMethodListTimeseriesGroupsParamsNormalization = "MIN0_MAX"
)
