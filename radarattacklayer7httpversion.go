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

// RadarAttackLayer7HTTPVersionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7HTTPVersionService] method instead.
type RadarAttackLayer7HTTPVersionService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7HTTPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7HTTPVersionService(opts ...option.RequestOption) (r *RadarAttackLayer7HTTPVersionService) {
	r = &RadarAttackLayer7HTTPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by http version used over time.
func (r *RadarAttackLayer7HTTPVersionService) ListTimeseriesGroups(ctx context.Context, query RadarAttackLayer7HTTPVersionListTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarAttackLayer7HTTPVersionListTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7HTTPVersionListTimeseriesGroupsResponse struct {
	Result  RadarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                                           `json:"success,required"`
	JSON    radarAttackLayer7HTTPVersionListTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarAttackLayer7HTTPVersionListTimeseriesGroupsResponseJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7HTTPVersionListTimeseriesGroupsResponse]
type radarAttackLayer7HTTPVersionListTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7HTTPVersionListTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResult struct {
	Meta   interface{}                                                          `json:"meta,required"`
	Serie0 RadarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResultJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResult]
type radarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResultSerie0 struct {
	HTTP1X     []string                                                                 `json:"HTTP/1.x,required"`
	Timestamps []string                                                                 `json:"timestamps,required"`
	JSON       radarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResultSerie0]
type radarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResultSerie0JSON struct {
	HTTP1X      apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7HTTPVersionListTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7HTTPVersionListTimeseriesGroupsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod] `query:"httpMethod"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7HTTPVersionListTimeseriesGroupsParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer7HTTPVersionListTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsAggInterval string

const (
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsAggInterval15m RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsAggInterval = "15m"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsAggInterval1h  RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsAggInterval = "1h"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsAggInterval1d  RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsAggInterval = "1d"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsAggInterval1w  RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsAggInterval = "1w"
)

type RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange string

const (
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange1d         RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "1d"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange2d         RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "2d"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange7d         RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "7d"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange14d        RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "14d"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange28d        RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "28d"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange12w        RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "12w"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange24w        RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "24w"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange52w        RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "52w"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange1dControl  RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "1dControl"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange2dControl  RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "2dControl"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange7dControl  RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "7dControl"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange14dControl RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "14dControl"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange28dControl RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "28dControl"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange12wControl RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "12wControl"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange24wControl RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsFormat string

const (
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsFormatJson RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsFormat = "JSON"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsFormatCsv  RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsFormat = "CSV"
)

type RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod string

const (
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodGet             RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "GET"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodPost            RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "POST"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodDelete          RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "DELETE"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodPut             RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "PUT"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodHead            RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "HEAD"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodPurge           RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "PURGE"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodOptions         RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodPropfind        RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodMkcol           RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodPatch           RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "PATCH"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodACL             RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "ACL"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodBcopy           RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodBdelete         RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodBmove           RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodBpropfind       RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodBproppatch      RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodCheckin         RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodCheckout        RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodConnect         RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodCopy            RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "COPY"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodLabel           RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "LABEL"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodLock            RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "LOCK"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodMerge           RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "MERGE"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodMkactivity      RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodMkworkspace     RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodMove            RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "MOVE"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodNotify          RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodOrderpatch      RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodPoll            RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "POLL"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodProppatch       RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodReport          RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "REPORT"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodSearch          RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodSubscribe       RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodTrace           RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "TRACE"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodUncheckout      RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodUnlock          RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodUnsubscribe     RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodUpdate          RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodVersioncontrol  RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodBaselinecontrol RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodXmsenumatts     RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodRpcOutData      RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodRpcInData       RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodJson            RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "JSON"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodCook            RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "COOK"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethodTrack           RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsIPVersion string

const (
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsIPVersionIPv4 RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsIPVersion = "IPv4"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsIPVersionIPv6 RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsIPVersion = "IPv6"
)

type RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProduct string

const (
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProductDdos               RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProduct = "DDOS"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProductWaf                RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProduct = "WAF"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProductBotManagement      RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProductAccessRules        RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProductIPReputation       RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProductAPIShield          RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProductDataLossPrevention RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsNormalization string

const (
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsNormalizationPercentage RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsNormalizationMin0Max    RadarAttackLayer7HTTPVersionListTimeseriesGroupsParamsNormalization = "MIN0_MAX"
)
