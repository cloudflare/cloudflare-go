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

// RadarAttackLayer7MitigationProductService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7MitigationProductService] method instead.
type RadarAttackLayer7MitigationProductService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7MitigationProductService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer7MitigationProductService(opts ...option.RequestOption) (r *RadarAttackLayer7MitigationProductService) {
	r = &RadarAttackLayer7MitigationProductService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by mitigation product used.
func (r *RadarAttackLayer7MitigationProductService) List(ctx context.Context, query RadarAttackLayer7MitigationProductListParams, opts ...option.RequestOption) (res *RadarAttackLayer7MitigationProductListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/mitigation_product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Percentage distribution of attacks by mitigation product used over time.
func (r *RadarAttackLayer7MitigationProductService) ListTimeseriesGroups(ctx context.Context, query RadarAttackLayer7MitigationProductListTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarAttackLayer7MitigationProductListTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/mitigation_product"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7MitigationProductListResponse struct {
	Result  RadarAttackLayer7MitigationProductListResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAttackLayer7MitigationProductListResponseJSON   `json:"-"`
}

// radarAttackLayer7MitigationProductListResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7MitigationProductListResponse]
type radarAttackLayer7MitigationProductListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7MitigationProductListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7MitigationProductListResponseResult struct {
	Meta     RadarAttackLayer7MitigationProductListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7MitigationProductListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7MitigationProductListResponseResultJSON     `json:"-"`
}

// radarAttackLayer7MitigationProductListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer7MitigationProductListResponseResult]
type radarAttackLayer7MitigationProductListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7MitigationProductListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7MitigationProductListResponseResultMeta struct {
	DateRange      []RadarAttackLayer7MitigationProductListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                 `json:"lastUpdated,required"`
	Normalization  string                                                                 `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7MitigationProductListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7MitigationProductListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7MitigationProductListResponseResultMeta]
type radarAttackLayer7MitigationProductListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7MitigationProductListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7MitigationProductListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7MitigationProductListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7MitigationProductListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7MitigationProductListResponseResultMetaDateRange]
type radarAttackLayer7MitigationProductListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7MitigationProductListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                              `json:"level"`
	JSON        radarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfo]
type radarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                               `json:"dataSource,required"`
	Description     string                                                                               `json:"description,required"`
	EventType       string                                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                                            `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource      apijson.Field
	Description     apijson.Field
	EventType       apijson.Field
	IsInstantaneous apijson.Field
	EndTime         apijson.Field
	LinkedURL       apijson.Field
	StartTime       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RadarAttackLayer7MitigationProductListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7MitigationProductListResponseResultSummary0 struct {
	Ddos string                                                           `json:"DDOS,required"`
	Waf  string                                                           `json:"WAF,required"`
	JSON radarAttackLayer7MitigationProductListResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer7MitigationProductListResponseResultSummary0JSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7MitigationProductListResponseResultSummary0]
type radarAttackLayer7MitigationProductListResponseResultSummary0JSON struct {
	Ddos        apijson.Field
	Waf         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7MitigationProductListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7MitigationProductListTimeseriesGroupsResponse struct {
	Result  RadarAttackLayer7MitigationProductListTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                                                 `json:"success,required"`
	JSON    radarAttackLayer7MitigationProductListTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarAttackLayer7MitigationProductListTimeseriesGroupsResponseJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7MitigationProductListTimeseriesGroupsResponse]
type radarAttackLayer7MitigationProductListTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7MitigationProductListTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7MitigationProductListTimeseriesGroupsResponseResult struct {
	Meta   interface{}                                                                `json:"meta,required"`
	Serie0 RadarAttackLayer7MitigationProductListTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7MitigationProductListTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7MitigationProductListTimeseriesGroupsResponseResultJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7MitigationProductListTimeseriesGroupsResponseResult]
type radarAttackLayer7MitigationProductListTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7MitigationProductListTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7MitigationProductListTimeseriesGroupsResponseResultSerie0 struct {
	Ddos       []string                                                                       `json:"DDOS,required"`
	Timestamps []string                                                                       `json:"timestamps,required"`
	JSON       radarAttackLayer7MitigationProductListTimeseriesGroupsResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7MitigationProductListTimeseriesGroupsResponseResultSerie0JSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7MitigationProductListTimeseriesGroupsResponseResultSerie0]
type radarAttackLayer7MitigationProductListTimeseriesGroupsResponseResultSerie0JSON struct {
	Ddos        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7MitigationProductListTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7MitigationProductListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7MitigationProductListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7MitigationProductListParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7MitigationProductListParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7MitigationProductListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7MitigationProductListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7MitigationProductListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7MitigationProductListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7MitigationProductListParamsDateRange string

const (
	RadarAttackLayer7MitigationProductListParamsDateRange1d         RadarAttackLayer7MitigationProductListParamsDateRange = "1d"
	RadarAttackLayer7MitigationProductListParamsDateRange2d         RadarAttackLayer7MitigationProductListParamsDateRange = "2d"
	RadarAttackLayer7MitigationProductListParamsDateRange7d         RadarAttackLayer7MitigationProductListParamsDateRange = "7d"
	RadarAttackLayer7MitigationProductListParamsDateRange14d        RadarAttackLayer7MitigationProductListParamsDateRange = "14d"
	RadarAttackLayer7MitigationProductListParamsDateRange28d        RadarAttackLayer7MitigationProductListParamsDateRange = "28d"
	RadarAttackLayer7MitigationProductListParamsDateRange12w        RadarAttackLayer7MitigationProductListParamsDateRange = "12w"
	RadarAttackLayer7MitigationProductListParamsDateRange24w        RadarAttackLayer7MitigationProductListParamsDateRange = "24w"
	RadarAttackLayer7MitigationProductListParamsDateRange52w        RadarAttackLayer7MitigationProductListParamsDateRange = "52w"
	RadarAttackLayer7MitigationProductListParamsDateRange1dControl  RadarAttackLayer7MitigationProductListParamsDateRange = "1dControl"
	RadarAttackLayer7MitigationProductListParamsDateRange2dControl  RadarAttackLayer7MitigationProductListParamsDateRange = "2dControl"
	RadarAttackLayer7MitigationProductListParamsDateRange7dControl  RadarAttackLayer7MitigationProductListParamsDateRange = "7dControl"
	RadarAttackLayer7MitigationProductListParamsDateRange14dControl RadarAttackLayer7MitigationProductListParamsDateRange = "14dControl"
	RadarAttackLayer7MitigationProductListParamsDateRange28dControl RadarAttackLayer7MitigationProductListParamsDateRange = "28dControl"
	RadarAttackLayer7MitigationProductListParamsDateRange12wControl RadarAttackLayer7MitigationProductListParamsDateRange = "12wControl"
	RadarAttackLayer7MitigationProductListParamsDateRange24wControl RadarAttackLayer7MitigationProductListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7MitigationProductListParamsFormat string

const (
	RadarAttackLayer7MitigationProductListParamsFormatJson RadarAttackLayer7MitigationProductListParamsFormat = "JSON"
	RadarAttackLayer7MitigationProductListParamsFormatCsv  RadarAttackLayer7MitigationProductListParamsFormat = "CSV"
)

type RadarAttackLayer7MitigationProductListParamsHTTPMethod string

const (
	RadarAttackLayer7MitigationProductListParamsHTTPMethodGet             RadarAttackLayer7MitigationProductListParamsHTTPMethod = "GET"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodPost            RadarAttackLayer7MitigationProductListParamsHTTPMethod = "POST"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodDelete          RadarAttackLayer7MitigationProductListParamsHTTPMethod = "DELETE"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodPut             RadarAttackLayer7MitigationProductListParamsHTTPMethod = "PUT"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodHead            RadarAttackLayer7MitigationProductListParamsHTTPMethod = "HEAD"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodPurge           RadarAttackLayer7MitigationProductListParamsHTTPMethod = "PURGE"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodOptions         RadarAttackLayer7MitigationProductListParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodPropfind        RadarAttackLayer7MitigationProductListParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodMkcol           RadarAttackLayer7MitigationProductListParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodPatch           RadarAttackLayer7MitigationProductListParamsHTTPMethod = "PATCH"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodACL             RadarAttackLayer7MitigationProductListParamsHTTPMethod = "ACL"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodBcopy           RadarAttackLayer7MitigationProductListParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodBdelete         RadarAttackLayer7MitigationProductListParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodBmove           RadarAttackLayer7MitigationProductListParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodBpropfind       RadarAttackLayer7MitigationProductListParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodBproppatch      RadarAttackLayer7MitigationProductListParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodCheckin         RadarAttackLayer7MitigationProductListParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodCheckout        RadarAttackLayer7MitigationProductListParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodConnect         RadarAttackLayer7MitigationProductListParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodCopy            RadarAttackLayer7MitigationProductListParamsHTTPMethod = "COPY"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodLabel           RadarAttackLayer7MitigationProductListParamsHTTPMethod = "LABEL"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodLock            RadarAttackLayer7MitigationProductListParamsHTTPMethod = "LOCK"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodMerge           RadarAttackLayer7MitigationProductListParamsHTTPMethod = "MERGE"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodMkactivity      RadarAttackLayer7MitigationProductListParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodMkworkspace     RadarAttackLayer7MitigationProductListParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodMove            RadarAttackLayer7MitigationProductListParamsHTTPMethod = "MOVE"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodNotify          RadarAttackLayer7MitigationProductListParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodOrderpatch      RadarAttackLayer7MitigationProductListParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodPoll            RadarAttackLayer7MitigationProductListParamsHTTPMethod = "POLL"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodProppatch       RadarAttackLayer7MitigationProductListParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodReport          RadarAttackLayer7MitigationProductListParamsHTTPMethod = "REPORT"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodSearch          RadarAttackLayer7MitigationProductListParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodSubscribe       RadarAttackLayer7MitigationProductListParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodTrace           RadarAttackLayer7MitigationProductListParamsHTTPMethod = "TRACE"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodUncheckout      RadarAttackLayer7MitigationProductListParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodUnlock          RadarAttackLayer7MitigationProductListParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodUnsubscribe     RadarAttackLayer7MitigationProductListParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodUpdate          RadarAttackLayer7MitigationProductListParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodVersioncontrol  RadarAttackLayer7MitigationProductListParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodBaselinecontrol RadarAttackLayer7MitigationProductListParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodXmsenumatts     RadarAttackLayer7MitigationProductListParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodRpcOutData      RadarAttackLayer7MitigationProductListParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodRpcInData       RadarAttackLayer7MitigationProductListParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodJson            RadarAttackLayer7MitigationProductListParamsHTTPMethod = "JSON"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodCook            RadarAttackLayer7MitigationProductListParamsHTTPMethod = "COOK"
	RadarAttackLayer7MitigationProductListParamsHTTPMethodTrack           RadarAttackLayer7MitigationProductListParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7MitigationProductListParamsHTTPVersion string

const (
	RadarAttackLayer7MitigationProductListParamsHTTPVersionHttPv1 RadarAttackLayer7MitigationProductListParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7MitigationProductListParamsHTTPVersionHttPv2 RadarAttackLayer7MitigationProductListParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7MitigationProductListParamsHTTPVersionHttPv3 RadarAttackLayer7MitigationProductListParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7MitigationProductListParamsIPVersion string

const (
	RadarAttackLayer7MitigationProductListParamsIPVersionIPv4 RadarAttackLayer7MitigationProductListParamsIPVersion = "IPv4"
	RadarAttackLayer7MitigationProductListParamsIPVersionIPv6 RadarAttackLayer7MitigationProductListParamsIPVersion = "IPv6"
)

type RadarAttackLayer7MitigationProductListTimeseriesGroupsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsNormalization] `query:"normalization"`
}

// URLQuery serializes
// [RadarAttackLayer7MitigationProductListTimeseriesGroupsParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7MitigationProductListTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsAggInterval string

const (
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsAggInterval15m RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsAggInterval = "15m"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsAggInterval1h  RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsAggInterval = "1h"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsAggInterval1d  RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsAggInterval = "1d"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsAggInterval1w  RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsAggInterval = "1w"
)

type RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange string

const (
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange1d         RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "1d"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange2d         RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "2d"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange7d         RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "7d"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange14d        RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "14d"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange28d        RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "28d"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange12w        RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "12w"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange24w        RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "24w"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange52w        RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "52w"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange1dControl  RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "1dControl"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange2dControl  RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "2dControl"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange7dControl  RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "7dControl"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange14dControl RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "14dControl"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange28dControl RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "28dControl"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange12wControl RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "12wControl"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange24wControl RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsFormat string

const (
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsFormatJson RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsFormat = "JSON"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsFormatCsv  RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsFormat = "CSV"
)

type RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod string

const (
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodGet             RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "GET"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodPost            RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "POST"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodDelete          RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "DELETE"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodPut             RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "PUT"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodHead            RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "HEAD"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodPurge           RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "PURGE"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodOptions         RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodPropfind        RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodMkcol           RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodPatch           RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "PATCH"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodACL             RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "ACL"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodBcopy           RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodBdelete         RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodBmove           RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodBpropfind       RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodBproppatch      RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodCheckin         RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodCheckout        RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodConnect         RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodCopy            RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "COPY"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodLabel           RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "LABEL"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodLock            RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "LOCK"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodMerge           RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "MERGE"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodMkactivity      RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodMkworkspace     RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodMove            RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "MOVE"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodNotify          RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodOrderpatch      RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodPoll            RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "POLL"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodProppatch       RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodReport          RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "REPORT"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodSearch          RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodSubscribe       RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodTrace           RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "TRACE"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodUncheckout      RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodUnlock          RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodUnsubscribe     RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodUpdate          RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodVersioncontrol  RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodBaselinecontrol RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodXmsenumatts     RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodRpcOutData      RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodRpcInData       RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodJson            RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "JSON"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodCook            RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "COOK"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethodTrack           RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPVersion string

const (
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPVersionHttPv1 RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPVersionHttPv2 RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPVersionHttPv3 RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsIPVersion string

const (
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsIPVersionIPv4 RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsIPVersion = "IPv4"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsIPVersionIPv6 RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsIPVersion = "IPv6"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsNormalization string

const (
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsNormalizationPercentage RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsNormalizationMin0Max    RadarAttackLayer7MitigationProductListTimeseriesGroupsParamsNormalization = "MIN0_MAX"
)
