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

// RadarAttackLayer7ManagedRuleService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7ManagedRuleService] method instead.
type RadarAttackLayer7ManagedRuleService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7ManagedRuleService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7ManagedRuleService(opts ...option.RequestOption) (r *RadarAttackLayer7ManagedRuleService) {
	r = &RadarAttackLayer7ManagedRuleService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by managed rules used over time.
func (r *RadarAttackLayer7ManagedRuleService) ListTimeseriesGroups(ctx context.Context, query RadarAttackLayer7ManagedRuleListTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarAttackLayer7ManagedRuleListTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/managed_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7ManagedRuleListTimeseriesGroupsResponse struct {
	Result  RadarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                                           `json:"success,required"`
	JSON    radarAttackLayer7ManagedRuleListTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarAttackLayer7ManagedRuleListTimeseriesGroupsResponseJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7ManagedRuleListTimeseriesGroupsResponse]
type radarAttackLayer7ManagedRuleListTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7ManagedRuleListTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResult struct {
	Meta   interface{}                                                          `json:"meta,required"`
	Serie0 RadarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResultJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResult]
type radarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResultSerie0 struct {
	Bot        []string                                                                 `json:"Bot,required"`
	Timestamps []string                                                                 `json:"timestamps,required"`
	JSON       radarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResultSerie0]
type radarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResultSerie0JSON struct {
	Bot         apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7ManagedRuleListTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7ManagedRuleListTimeseriesGroupsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7ManagedRuleListTimeseriesGroupsParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer7ManagedRuleListTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsAggInterval string

const (
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsAggInterval15m RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsAggInterval = "15m"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsAggInterval1h  RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsAggInterval = "1h"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsAggInterval1d  RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsAggInterval = "1d"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsAggInterval1w  RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsAggInterval = "1w"
)

type RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange string

const (
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange1d         RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "1d"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange2d         RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "2d"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange7d         RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "7d"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange14d        RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "14d"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange28d        RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "28d"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange12w        RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "12w"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange24w        RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "24w"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange52w        RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "52w"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange1dControl  RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "1dControl"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange2dControl  RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "2dControl"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange7dControl  RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "7dControl"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange14dControl RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "14dControl"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange28dControl RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "28dControl"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange12wControl RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "12wControl"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange24wControl RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsFormat string

const (
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsFormatJson RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsFormat = "JSON"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsFormatCsv  RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsFormat = "CSV"
)

type RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod string

const (
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodGet             RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "GET"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodPost            RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "POST"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodDelete          RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "DELETE"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodPut             RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "PUT"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodHead            RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "HEAD"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodPurge           RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "PURGE"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodOptions         RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodPropfind        RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodMkcol           RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodPatch           RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "PATCH"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodACL             RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "ACL"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodBcopy           RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodBdelete         RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodBmove           RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodBpropfind       RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodBproppatch      RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodCheckin         RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodCheckout        RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodConnect         RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodCopy            RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "COPY"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodLabel           RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "LABEL"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodLock            RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "LOCK"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodMerge           RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "MERGE"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodMkactivity      RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodMkworkspace     RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodMove            RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "MOVE"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodNotify          RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodOrderpatch      RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodPoll            RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "POLL"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodProppatch       RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodReport          RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "REPORT"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodSearch          RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodSubscribe       RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodTrace           RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "TRACE"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodUncheckout      RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodUnlock          RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodUnsubscribe     RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodUpdate          RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodVersioncontrol  RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodBaselinecontrol RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodXmsenumatts     RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodRpcOutData      RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodRpcInData       RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodJson            RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "JSON"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodCook            RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "COOK"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethodTrack           RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPVersion string

const (
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPVersionHttPv1 RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPVersionHttPv2 RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPVersionHttPv3 RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsIPVersion string

const (
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsIPVersionIPv4 RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsIPVersion = "IPv4"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsIPVersionIPv6 RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsIPVersion = "IPv6"
)

type RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProduct string

const (
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProductDdos               RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProduct = "DDOS"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProductWaf                RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProduct = "WAF"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProductBotManagement      RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProductAccessRules        RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProductIPReputation       RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProductAPIShield          RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProductDataLossPrevention RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsNormalization string

const (
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsNormalizationPercentage RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsNormalizationMin0Max    RadarAttackLayer7ManagedRuleListTimeseriesGroupsParamsNormalization = "MIN0_MAX"
)
