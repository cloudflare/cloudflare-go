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

// RadarAttackLayer7IndustryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer7IndustryService] method instead.
type RadarAttackLayer7IndustryService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7IndustryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7IndustryService(opts ...option.RequestOption) (r *RadarAttackLayer7IndustryService) {
	r = &RadarAttackLayer7IndustryService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by industry used over time.
func (r *RadarAttackLayer7IndustryService) ListTimeseriesGroups(ctx context.Context, query RadarAttackLayer7IndustryListTimeseriesGroupsParams, opts ...option.RequestOption) (res *RadarAttackLayer7IndustryListTimeseriesGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/timeseries_groups/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get the Industry of attacks.
func (r *RadarAttackLayer7IndustryService) ListTops(ctx context.Context, query RadarAttackLayer7IndustryListTopsParams, opts ...option.RequestOption) (res *RadarAttackLayer7IndustryListTopsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/top/industry"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7IndustryListTimeseriesGroupsResponse struct {
	Result  RadarAttackLayer7IndustryListTimeseriesGroupsResponseResult `json:"result,required"`
	Success bool                                                        `json:"success,required"`
	JSON    radarAttackLayer7IndustryListTimeseriesGroupsResponseJSON   `json:"-"`
}

// radarAttackLayer7IndustryListTimeseriesGroupsResponseJSON contains the JSON
// metadata for the struct [RadarAttackLayer7IndustryListTimeseriesGroupsResponse]
type radarAttackLayer7IndustryListTimeseriesGroupsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7IndustryListTimeseriesGroupsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IndustryListTimeseriesGroupsResponseResult struct {
	Meta   interface{}                                                       `json:"meta,required"`
	Serie0 RadarAttackLayer7IndustryListTimeseriesGroupsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarAttackLayer7IndustryListTimeseriesGroupsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7IndustryListTimeseriesGroupsResponseResultJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7IndustryListTimeseriesGroupsResponseResult]
type radarAttackLayer7IndustryListTimeseriesGroupsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7IndustryListTimeseriesGroupsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IndustryListTimeseriesGroupsResponseResultSerie0 struct {
	Timestamps []string                                                              `json:"timestamps,required"`
	JSON       radarAttackLayer7IndustryListTimeseriesGroupsResponseResultSerie0JSON `json:"-"`
}

// radarAttackLayer7IndustryListTimeseriesGroupsResponseResultSerie0JSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7IndustryListTimeseriesGroupsResponseResultSerie0]
type radarAttackLayer7IndustryListTimeseriesGroupsResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7IndustryListTimeseriesGroupsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IndustryListTopsResponse struct {
	Result  RadarAttackLayer7IndustryListTopsResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarAttackLayer7IndustryListTopsResponseJSON   `json:"-"`
}

// radarAttackLayer7IndustryListTopsResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7IndustryListTopsResponse]
type radarAttackLayer7IndustryListTopsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7IndustryListTopsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IndustryListTopsResponseResult struct {
	Meta RadarAttackLayer7IndustryListTopsResponseResultMeta   `json:"meta,required"`
	Top0 []RadarAttackLayer7IndustryListTopsResponseResultTop0 `json:"top_0,required"`
	JSON radarAttackLayer7IndustryListTopsResponseResultJSON   `json:"-"`
}

// radarAttackLayer7IndustryListTopsResponseResultJSON contains the JSON metadata
// for the struct [RadarAttackLayer7IndustryListTopsResponseResult]
type radarAttackLayer7IndustryListTopsResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7IndustryListTopsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IndustryListTopsResponseResultMeta struct {
	DateRange      []RadarAttackLayer7IndustryListTopsResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7IndustryListTopsResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7IndustryListTopsResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarAttackLayer7IndustryListTopsResponseResultMeta]
type radarAttackLayer7IndustryListTopsResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7IndustryListTopsResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IndustryListTopsResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7IndustryListTopsResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7IndustryListTopsResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7IndustryListTopsResponseResultMetaDateRange]
type radarAttackLayer7IndustryListTopsResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7IndustryListTopsResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                         `json:"level"`
	JSON        radarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfo]
type radarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                          `json:"dataSource,required"`
	Description     string                                                                          `json:"description,required"`
	EventType       string                                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                       `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7IndustryListTopsResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IndustryListTopsResponseResultTop0 struct {
	Name  string                                                  `json:"name,required"`
	Value string                                                  `json:"value,required"`
	JSON  radarAttackLayer7IndustryListTopsResponseResultTop0JSON `json:"-"`
}

// radarAttackLayer7IndustryListTopsResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7IndustryListTopsResponseResultTop0]
type radarAttackLayer7IndustryListTopsResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7IndustryListTopsResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7IndustryListTimeseriesGroupsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarAttackLayer7IndustryListTimeseriesGroupsParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7IndustryListTimeseriesGroupsParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7IndustryListTimeseriesGroupsParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[RadarAttackLayer7IndustryListTimeseriesGroupsParamsNormalization] `query:"normalization"`
}

// URLQuery serializes [RadarAttackLayer7IndustryListTimeseriesGroupsParams]'s
// query parameters as `url.Values`.
func (r RadarAttackLayer7IndustryListTimeseriesGroupsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarAttackLayer7IndustryListTimeseriesGroupsParamsAggInterval string

const (
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsAggInterval15m RadarAttackLayer7IndustryListTimeseriesGroupsParamsAggInterval = "15m"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsAggInterval1h  RadarAttackLayer7IndustryListTimeseriesGroupsParamsAggInterval = "1h"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsAggInterval1d  RadarAttackLayer7IndustryListTimeseriesGroupsParamsAggInterval = "1d"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsAggInterval1w  RadarAttackLayer7IndustryListTimeseriesGroupsParamsAggInterval = "1w"
)

type RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange string

const (
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange1d         RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "1d"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange2d         RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "2d"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange7d         RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "7d"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange14d        RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "14d"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange28d        RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "28d"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange12w        RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "12w"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange24w        RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "24w"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange52w        RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "52w"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange1dControl  RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "1dControl"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange2dControl  RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "2dControl"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange7dControl  RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "7dControl"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange14dControl RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "14dControl"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange28dControl RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "28dControl"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange12wControl RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "12wControl"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange24wControl RadarAttackLayer7IndustryListTimeseriesGroupsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7IndustryListTimeseriesGroupsParamsFormat string

const (
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsFormatJson RadarAttackLayer7IndustryListTimeseriesGroupsParamsFormat = "JSON"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsFormatCsv  RadarAttackLayer7IndustryListTimeseriesGroupsParamsFormat = "CSV"
)

type RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod string

const (
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodGet             RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "GET"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodPost            RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "POST"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodDelete          RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "DELETE"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodPut             RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "PUT"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodHead            RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "HEAD"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodPurge           RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "PURGE"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodOptions         RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodPropfind        RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodMkcol           RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodPatch           RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "PATCH"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodACL             RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "ACL"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodBcopy           RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodBdelete         RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodBmove           RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodBpropfind       RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodBproppatch      RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodCheckin         RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodCheckout        RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodConnect         RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodCopy            RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "COPY"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodLabel           RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "LABEL"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodLock            RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "LOCK"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodMerge           RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "MERGE"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodMkactivity      RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodMkworkspace     RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodMove            RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "MOVE"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodNotify          RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodOrderpatch      RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodPoll            RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "POLL"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodProppatch       RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodReport          RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "REPORT"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodSearch          RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodSubscribe       RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodTrace           RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "TRACE"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodUncheckout      RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodUnlock          RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodUnsubscribe     RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodUpdate          RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodVersioncontrol  RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodBaselinecontrol RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodXmsenumatts     RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodRpcOutData      RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodRpcInData       RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodJson            RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "JSON"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodCook            RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "COOK"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethodTrack           RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPVersion string

const (
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPVersionHttPv1 RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPVersionHttPv2 RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPVersionHttPv3 RadarAttackLayer7IndustryListTimeseriesGroupsParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7IndustryListTimeseriesGroupsParamsIPVersion string

const (
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsIPVersionIPv4 RadarAttackLayer7IndustryListTimeseriesGroupsParamsIPVersion = "IPv4"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsIPVersionIPv6 RadarAttackLayer7IndustryListTimeseriesGroupsParamsIPVersion = "IPv6"
)

type RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProduct string

const (
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProductDdos               RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProduct = "DDOS"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProductWaf                RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProduct = "WAF"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProductBotManagement      RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProductAccessRules        RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProductIPReputation       RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProductAPIShield          RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProductDataLossPrevention RadarAttackLayer7IndustryListTimeseriesGroupsParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type RadarAttackLayer7IndustryListTimeseriesGroupsParamsNormalization string

const (
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsNormalizationPercentage RadarAttackLayer7IndustryListTimeseriesGroupsParamsNormalization = "PERCENTAGE"
	RadarAttackLayer7IndustryListTimeseriesGroupsParamsNormalizationMin0Max    RadarAttackLayer7IndustryListTimeseriesGroupsParamsNormalization = "MIN0_MAX"
)

type RadarAttackLayer7IndustryListTopsParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7IndustryListTopsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7IndustryListTopsParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7IndustryListTopsParams]'s query parameters
// as `url.Values`.
func (r RadarAttackLayer7IndustryListTopsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7IndustryListTopsParamsDateRange string

const (
	RadarAttackLayer7IndustryListTopsParamsDateRange1d         RadarAttackLayer7IndustryListTopsParamsDateRange = "1d"
	RadarAttackLayer7IndustryListTopsParamsDateRange2d         RadarAttackLayer7IndustryListTopsParamsDateRange = "2d"
	RadarAttackLayer7IndustryListTopsParamsDateRange7d         RadarAttackLayer7IndustryListTopsParamsDateRange = "7d"
	RadarAttackLayer7IndustryListTopsParamsDateRange14d        RadarAttackLayer7IndustryListTopsParamsDateRange = "14d"
	RadarAttackLayer7IndustryListTopsParamsDateRange28d        RadarAttackLayer7IndustryListTopsParamsDateRange = "28d"
	RadarAttackLayer7IndustryListTopsParamsDateRange12w        RadarAttackLayer7IndustryListTopsParamsDateRange = "12w"
	RadarAttackLayer7IndustryListTopsParamsDateRange24w        RadarAttackLayer7IndustryListTopsParamsDateRange = "24w"
	RadarAttackLayer7IndustryListTopsParamsDateRange52w        RadarAttackLayer7IndustryListTopsParamsDateRange = "52w"
	RadarAttackLayer7IndustryListTopsParamsDateRange1dControl  RadarAttackLayer7IndustryListTopsParamsDateRange = "1dControl"
	RadarAttackLayer7IndustryListTopsParamsDateRange2dControl  RadarAttackLayer7IndustryListTopsParamsDateRange = "2dControl"
	RadarAttackLayer7IndustryListTopsParamsDateRange7dControl  RadarAttackLayer7IndustryListTopsParamsDateRange = "7dControl"
	RadarAttackLayer7IndustryListTopsParamsDateRange14dControl RadarAttackLayer7IndustryListTopsParamsDateRange = "14dControl"
	RadarAttackLayer7IndustryListTopsParamsDateRange28dControl RadarAttackLayer7IndustryListTopsParamsDateRange = "28dControl"
	RadarAttackLayer7IndustryListTopsParamsDateRange12wControl RadarAttackLayer7IndustryListTopsParamsDateRange = "12wControl"
	RadarAttackLayer7IndustryListTopsParamsDateRange24wControl RadarAttackLayer7IndustryListTopsParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7IndustryListTopsParamsFormat string

const (
	RadarAttackLayer7IndustryListTopsParamsFormatJson RadarAttackLayer7IndustryListTopsParamsFormat = "JSON"
	RadarAttackLayer7IndustryListTopsParamsFormatCsv  RadarAttackLayer7IndustryListTopsParamsFormat = "CSV"
)
