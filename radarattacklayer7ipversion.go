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

// RadarAttackLayer7IPVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer7IPVersionService] method instead.
type RadarAttackLayer7IPVersionService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7IPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7IPVersionService(opts ...option.RequestOption) (r *RadarAttackLayer7IPVersionService) {
	r = &RadarAttackLayer7IPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by ip version used over time.
func (r *RadarAttackLayer7IPVersionService) ListTimeseriesGroups(ctx context.Context, query RadarAttackLayer7IPVersionListTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarAttackLayer7IPVersionListTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7IPVersionListTimeseriesGroupsResponse struct {
	Result  RadarAttackLayer7IPVersionListTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                                         `json:"success,required"`
	JSON    radarAttackLayer7IPVersionListTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarAttackLayer7IPVersionListTimeseriesGroupsResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer7IPVersionListTimeseriesGroupsResponse]
type radarAttackLayer7IPVersionListTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7IPVersionListTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IPVersionListTimeseriesGroupsResponseResult struct {
	Meta   interface{}                                                        `json:"meta,required"`
	Serie0 RadarAttackLayer7IPVersionListTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7IPVersionListTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7IPVersionListTimeseriesGroupsResponseResultJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7IPVersionListTimeseriesGroupsResponseResult]
type radarAttackLayer7IPVersionListTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7IPVersionListTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IPVersionListTimeseriesGroupsResponseResultSerie0 struct {
	IPv4       []string                                                               `json:"IPv4,required"`
	IPv6       []string                                                               `json:"IPv6,required"`
	Timestamps []string                                                               `json:"timestamps,required"`
	JSON       radarAttackLayer7IPVersionListTimeseriesGroupsResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7IPVersionListTimeseriesGroupsResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7IPVersionListTimeseriesGroupsResponseResultSerie0]
type radarAttackLayer7IPVersionListTimeseriesGroupsResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7IPVersionListTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IPVersionListTimeseriesGroupsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7IPVersionListTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7IPVersionListTimeseriesGroupsParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPVersion] `query:"httpVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7IPVersionListTimeseriesGroupsParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7IPVersionListTimeseriesGroupsParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer7IPVersionListTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7IPVersionListTimeseriesGroupsParamsAggInterval string

const (
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsAggInterval15m RadarAttackLayer7IPVersionListTimeseriesGroupsParamsAggInterval = "15m"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsAggInterval1h  RadarAttackLayer7IPVersionListTimeseriesGroupsParamsAggInterval = "1h"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsAggInterval1d  RadarAttackLayer7IPVersionListTimeseriesGroupsParamsAggInterval = "1d"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsAggInterval1w  RadarAttackLayer7IPVersionListTimeseriesGroupsParamsAggInterval = "1w"
)

type RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange string

const (
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange1d         RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "1d"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange2d         RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "2d"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange7d         RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "7d"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange14d        RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "14d"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange28d        RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "28d"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange12w        RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "12w"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange24w        RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "24w"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange52w        RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "52w"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange1dControl  RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "1dControl"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange2dControl  RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "2dControl"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange7dControl  RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "7dControl"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange14dControl RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "14dControl"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange28dControl RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "28dControl"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange12wControl RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "12wControl"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange24wControl RadarAttackLayer7IPVersionListTimeseriesGroupsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7IPVersionListTimeseriesGroupsParamsFormat string

const (
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsFormatJson RadarAttackLayer7IPVersionListTimeseriesGroupsParamsFormat = "JSON"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsFormatCsv  RadarAttackLayer7IPVersionListTimeseriesGroupsParamsFormat = "CSV"
)

type RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod string

const (
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodGet             RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "GET"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodPost            RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "POST"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodDelete          RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "DELETE"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodPut             RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "PUT"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodHead            RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "HEAD"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodPurge           RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "PURGE"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodOptions         RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodPropfind        RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodMkcol           RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodPatch           RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "PATCH"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodACL             RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "ACL"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodBcopy           RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodBdelete         RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodBmove           RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodBpropfind       RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodBproppatch      RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodCheckin         RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodCheckout        RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodConnect         RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodCopy            RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "COPY"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodLabel           RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "LABEL"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodLock            RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "LOCK"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodMerge           RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "MERGE"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodMkactivity      RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodMkworkspace     RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodMove            RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "MOVE"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodNotify          RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodOrderpatch      RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodPoll            RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "POLL"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodProppatch       RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodReport          RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "REPORT"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodSearch          RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodSubscribe       RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodTrace           RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "TRACE"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodUncheckout      RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodUnlock          RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodUnsubscribe     RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodUpdate          RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodVersioncontrol  RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodBaselinecontrol RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodXmsenumatts     RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodRpcOutData      RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodRpcInData       RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodJson            RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "JSON"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodCook            RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "COOK"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethodTrack           RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPVersion string

const (
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPVersionHttPv1 RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPVersionHttPv2 RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPVersionHttPv3 RadarAttackLayer7IPVersionListTimeseriesGroupsParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProduct string

const (
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProductDdos               RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProduct = "DDOS"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProductWaf                RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProduct = "WAF"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProductBotManagement      RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProductAccessRules        RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProductIPReputation       RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProductAPIShield          RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProductDataLossPrevention RadarAttackLayer7IPVersionListTimeseriesGroupsParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7IPVersionListTimeseriesGroupsParamsNormalization string

const (
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsNormalizationPercentage RadarAttackLayer7IPVersionListTimeseriesGroupsParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7IPVersionListTimeseriesGroupsParamsNormalizationMin0Max    RadarAttackLayer7IPVersionListTimeseriesGroupsParamsNormalization = "MIN0_MAX"
)
