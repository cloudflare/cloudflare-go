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

// RadarAttackLayer7VerticalService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer7VerticalService] method instead.
type RadarAttackLayer7VerticalService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7VerticalService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7VerticalService(opts ...option.RequestOption) (r *RadarAttackLayer7VerticalService) {
	r = &RadarAttackLayer7VerticalService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by vertical used over time.
func (r *RadarAttackLayer7VerticalService) ListTimeseriesGroups(ctx context.Context, query RadarAttackLayer7VerticalListTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarAttackLayer7VerticalListTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the Verticals of attacks.
func (r *RadarAttackLayer7VerticalService) ListTops(ctx context.Context, query RadarAttackLayer7VerticalListTopsParams, opts ...option.RequestOption) (res *RadarAttackLayer7VerticalListTopsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/vertical"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7VerticalListTimeseriesGroupsResponse struct {
	Result  RadarAttackLayer7VerticalListTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                                        `json:"success,required"`
	JSON    radarAttackLayer7VerticalListTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarAttackLayer7VerticalListTimeseriesGroupsResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer7VerticalListTimeseriesGroupsResponse]
type radarAttackLayer7VerticalListTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7VerticalListTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7VerticalListTimeseriesGroupsResponseResult struct {
	Meta   interface{}                                                       `json:"meta,required"`
	Serie0 RadarAttackLayer7VerticalListTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7VerticalListTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7VerticalListTimeseriesGroupsResponseResultJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7VerticalListTimeseriesGroupsResponseResult]
type radarAttackLayer7VerticalListTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7VerticalListTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7VerticalListTimeseriesGroupsResponseResultSerie0 struct {
	Timestamps []string                                                              `json:"timestamps,required"`
	JSON       radarAttackLayer7VerticalListTimeseriesGroupsResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7VerticalListTimeseriesGroupsResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7VerticalListTimeseriesGroupsResponseResultSerie0]
type radarAttackLayer7VerticalListTimeseriesGroupsResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7VerticalListTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7VerticalListTopsResponse struct {
	Result  RadarAttackLayer7VerticalListTopsResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer7VerticalListTopsResponseJSON   `json:"-"`
}

// radarAttackLayer7VerticalListTopsResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7VerticalListTopsResponse]
type radarAttackLayer7VerticalListTopsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7VerticalListTopsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7VerticalListTopsResponseResult struct {
	Meta RadarAttackLayer7VerticalListTopsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7VerticalListTopsResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7VerticalListTopsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7VerticalListTopsResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer7VerticalListTopsResponseResult]
type radarAttackLayer7VerticalListTopsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7VerticalListTopsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7VerticalListTopsResponseResultMeta struct {
	DateRange      []RadarAttackLayer7VerticalListTopsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7VerticalListTopsResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7VerticalListTopsResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer7VerticalListTopsResponseResultMeta]
type radarAttackLayer7VerticalListTopsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7VerticalListTopsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7VerticalListTopsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7VerticalListTopsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7VerticalListTopsResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7VerticalListTopsResponseResultMetaDateRange]
type radarAttackLayer7VerticalListTopsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7VerticalListTopsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                         `json:"level"`
	JSON        radarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfo]
type radarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                          `json:"dataSource,required"`
	Description     string                                                                          `json:"description,required"`
	EventType       string                                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                       `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7VerticalListTopsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7VerticalListTopsResponseResultTop0 struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  radarAttackLayer7VerticalListTopsResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7VerticalListTopsResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7VerticalListTopsResponseResultTop0]
type radarAttackLayer7VerticalListTopsResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7VerticalListTopsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7VerticalListTimeseriesGroupsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7VerticalListTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7VerticalListTimeseriesGroupsParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7VerticalListTimeseriesGroupsParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7VerticalListTimeseriesGroupsParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7VerticalListTimeseriesGroupsParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer7VerticalListTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7VerticalListTimeseriesGroupsParamsAggInterval string

const (
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsAggInterval15m RadarAttackLayer7VerticalListTimeseriesGroupsParamsAggInterval = "15m"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsAggInterval1h  RadarAttackLayer7VerticalListTimeseriesGroupsParamsAggInterval = "1h"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsAggInterval1d  RadarAttackLayer7VerticalListTimeseriesGroupsParamsAggInterval = "1d"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsAggInterval1w  RadarAttackLayer7VerticalListTimeseriesGroupsParamsAggInterval = "1w"
)

type RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange string

const (
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange1d         RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "1d"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange2d         RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "2d"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange7d         RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "7d"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange14d        RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "14d"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange28d        RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "28d"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange12w        RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "12w"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange24w        RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "24w"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange52w        RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "52w"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange1dControl  RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "1dControl"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange2dControl  RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "2dControl"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange7dControl  RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "7dControl"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange14dControl RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "14dControl"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange28dControl RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "28dControl"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange12wControl RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "12wControl"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange24wControl RadarAttackLayer7VerticalListTimeseriesGroupsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7VerticalListTimeseriesGroupsParamsFormat string

const (
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsFormatJson RadarAttackLayer7VerticalListTimeseriesGroupsParamsFormat = "JSON"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsFormatCsv  RadarAttackLayer7VerticalListTimeseriesGroupsParamsFormat = "CSV"
)

type RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod string

const (
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodGet             RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "GET"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodPost            RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "POST"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodDelete          RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "DELETE"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodPut             RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "PUT"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodHead            RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "HEAD"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodPurge           RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "PURGE"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodOptions         RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodPropfind        RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodMkcol           RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodPatch           RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "PATCH"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodACL             RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "ACL"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodBcopy           RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodBdelete         RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodBmove           RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodBpropfind       RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodBproppatch      RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodCheckin         RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodCheckout        RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodConnect         RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodCopy            RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "COPY"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodLabel           RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "LABEL"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodLock            RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "LOCK"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodMerge           RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "MERGE"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodMkactivity      RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodMkworkspace     RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodMove            RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "MOVE"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodNotify          RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodOrderpatch      RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodPoll            RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "POLL"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodProppatch       RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodReport          RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "REPORT"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodSearch          RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodSubscribe       RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodTrace           RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "TRACE"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodUncheckout      RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodUnlock          RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodUnsubscribe     RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodUpdate          RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodVersioncontrol  RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodBaselinecontrol RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodXmsenumatts     RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodRpcOutData      RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodRpcInData       RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodJson            RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "JSON"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodCook            RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "COOK"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethodTrack           RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPVersion string

const (
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPVersionHttPv1 RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPVersionHttPv2 RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPVersionHttPv3 RadarAttackLayer7VerticalListTimeseriesGroupsParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7VerticalListTimeseriesGroupsParamsIPVersion string

const (
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsIPVersionIPv4 RadarAttackLayer7VerticalListTimeseriesGroupsParamsIPVersion = "IPv4"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsIPVersionIPv6 RadarAttackLayer7VerticalListTimeseriesGroupsParamsIPVersion = "IPv6"
)

type RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProduct string

const (
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProductDdos               RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProduct = "DDOS"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProductWaf                RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProduct = "WAF"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProductBotManagement      RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProductAccessRules        RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProductIPReputation       RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProductAPIShield          RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProductDataLossPrevention RadarAttackLayer7VerticalListTimeseriesGroupsParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7VerticalListTimeseriesGroupsParamsNormalization string

const (
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsNormalizationPercentage RadarAttackLayer7VerticalListTimeseriesGroupsParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7VerticalListTimeseriesGroupsParamsNormalizationMin0Max    RadarAttackLayer7VerticalListTimeseriesGroupsParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7VerticalListTopsParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7VerticalListTopsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7VerticalListTopsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7VerticalListTopsParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer7VerticalListTopsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7VerticalListTopsParamsDateRange string

const (
	RadarAttackLayer7VerticalListTopsParamsDateRange1d         RadarAttackLayer7VerticalListTopsParamsDateRange = "1d"
	RadarAttackLayer7VerticalListTopsParamsDateRange2d         RadarAttackLayer7VerticalListTopsParamsDateRange = "2d"
	RadarAttackLayer7VerticalListTopsParamsDateRange7d         RadarAttackLayer7VerticalListTopsParamsDateRange = "7d"
	RadarAttackLayer7VerticalListTopsParamsDateRange14d        RadarAttackLayer7VerticalListTopsParamsDateRange = "14d"
	RadarAttackLayer7VerticalListTopsParamsDateRange28d        RadarAttackLayer7VerticalListTopsParamsDateRange = "28d"
	RadarAttackLayer7VerticalListTopsParamsDateRange12w        RadarAttackLayer7VerticalListTopsParamsDateRange = "12w"
	RadarAttackLayer7VerticalListTopsParamsDateRange24w        RadarAttackLayer7VerticalListTopsParamsDateRange = "24w"
	RadarAttackLayer7VerticalListTopsParamsDateRange52w        RadarAttackLayer7VerticalListTopsParamsDateRange = "52w"
	RadarAttackLayer7VerticalListTopsParamsDateRange1dControl  RadarAttackLayer7VerticalListTopsParamsDateRange = "1dControl"
	RadarAttackLayer7VerticalListTopsParamsDateRange2dControl  RadarAttackLayer7VerticalListTopsParamsDateRange = "2dControl"
	RadarAttackLayer7VerticalListTopsParamsDateRange7dControl  RadarAttackLayer7VerticalListTopsParamsDateRange = "7dControl"
	RadarAttackLayer7VerticalListTopsParamsDateRange14dControl RadarAttackLayer7VerticalListTopsParamsDateRange = "14dControl"
	RadarAttackLayer7VerticalListTopsParamsDateRange28dControl RadarAttackLayer7VerticalListTopsParamsDateRange = "28dControl"
	RadarAttackLayer7VerticalListTopsParamsDateRange12wControl RadarAttackLayer7VerticalListTopsParamsDateRange = "12wControl"
	RadarAttackLayer7VerticalListTopsParamsDateRange24wControl RadarAttackLayer7VerticalListTopsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7VerticalListTopsParamsFormat string

const (
	RadarAttackLayer7VerticalListTopsParamsFormatJson RadarAttackLayer7VerticalListTopsParamsFormat = "JSON"
	RadarAttackLayer7VerticalListTopsParamsFormatCsv  RadarAttackLayer7VerticalListTopsParamsFormat = "CSV"
)
