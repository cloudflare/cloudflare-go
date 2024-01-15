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

// RadarAttackLayer7SummaryManagedRuleService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7SummaryManagedRuleService] method instead.
type RadarAttackLayer7SummaryManagedRuleService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7SummaryManagedRuleService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer7SummaryManagedRuleService(opts ...option.RequestOption) (r *RadarAttackLayer7SummaryManagedRuleService) {
	r = &RadarAttackLayer7SummaryManagedRuleService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by managed rules used.
func (r *RadarAttackLayer7SummaryManagedRuleService) List(ctx context.Context, query RadarAttackLayer7SummaryManagedRuleListParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryManagedRuleListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/managed_rules"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7SummaryManagedRuleListResponse struct {
	Result  RadarAttackLayer7SummaryManagedRuleListResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarAttackLayer7SummaryManagedRuleListResponseJSON   `json:"-"`
}

// radarAttackLayer7SummaryManagedRuleListResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryManagedRuleListResponse]
type radarAttackLayer7SummaryManagedRuleListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryManagedRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryManagedRuleListResponseResult struct {
	Meta     RadarAttackLayer7SummaryManagedRuleListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryManagedRuleListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryManagedRuleListResponseResultJSON     `json:"-"`
}

// radarAttackLayer7SummaryManagedRuleListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryManagedRuleListResponseResult]
type radarAttackLayer7SummaryManagedRuleListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryManagedRuleListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryManagedRuleListResponseResultMeta struct {
	DateRange      []RadarAttackLayer7SummaryManagedRuleListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                  `json:"lastUpdated,required"`
	Normalization  string                                                                  `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryManagedRuleListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryManagedRuleListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryManagedRuleListResponseResultMeta]
type radarAttackLayer7SummaryManagedRuleListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryManagedRuleListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryManagedRuleListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryManagedRuleListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryManagedRuleListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryManagedRuleListResponseResultMetaDateRange]
type radarAttackLayer7SummaryManagedRuleListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryManagedRuleListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                               `json:"level"`
	JSON        radarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfo]
type radarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                `json:"dataSource,required"`
	Description     string                                                                                `json:"description,required"`
	EventType       string                                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                                             `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryManagedRuleListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryManagedRuleListResponseResultSummary0 struct {
	Bot         string                                                            `json:"Bot,required"`
	HTTPAnomaly string                                                            `json:"HTTP Anomaly,required"`
	JSON        radarAttackLayer7SummaryManagedRuleListResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryManagedRuleListResponseResultSummary0JSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryManagedRuleListResponseResultSummary0]
type radarAttackLayer7SummaryManagedRuleListResponseResultSummary0JSON struct {
	Bot         apijson.Field
	HTTPAnomaly apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryManagedRuleListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryManagedRuleListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryManagedRuleListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryManagedRuleListParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryManagedRuleListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7SummaryManagedRuleListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryManagedRuleListParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryManagedRuleListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7SummaryManagedRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryManagedRuleListParamsDateRange string

const (
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange1d         RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "1d"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange2d         RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "2d"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange7d         RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "7d"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange14d        RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "14d"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange28d        RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "28d"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange12w        RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "12w"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange24w        RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "24w"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange52w        RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "52w"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange1dControl  RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange2dControl  RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange7dControl  RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange14dControl RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange28dControl RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange12wControl RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryManagedRuleListParamsDateRange24wControl RadarAttackLayer7SummaryManagedRuleListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7SummaryManagedRuleListParamsFormat string

const (
	RadarAttackLayer7SummaryManagedRuleListParamsFormatJson RadarAttackLayer7SummaryManagedRuleListParamsFormat = "JSON"
	RadarAttackLayer7SummaryManagedRuleListParamsFormatCsv  RadarAttackLayer7SummaryManagedRuleListParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodGet             RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodPost            RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodDelete          RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodPut             RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodHead            RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodPurge           RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodOptions         RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodPropfind        RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodMkcol           RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodPatch           RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodACL             RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodBcopy           RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodBdelete         RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodBmove           RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodCheckin         RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodCheckout        RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodConnect         RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodCopy            RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodLabel           RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodLock            RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodMerge           RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodMove            RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodNotify          RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodPoll            RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodProppatch       RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodReport          RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodSearch          RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodTrace           RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodUnlock          RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodUpdate          RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodJson            RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodCook            RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethodTrack           RadarAttackLayer7SummaryManagedRuleListParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7SummaryManagedRuleListParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryManagedRuleListParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryManagedRuleListParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryManagedRuleListParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryManagedRuleListParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7SummaryManagedRuleListParamsIPVersion string

const (
	RadarAttackLayer7SummaryManagedRuleListParamsIPVersionIPv4 RadarAttackLayer7SummaryManagedRuleListParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryManagedRuleListParamsIPVersionIPv6 RadarAttackLayer7SummaryManagedRuleListParamsIPVersion = "IPv6"
)

type RadarAttackLayer7SummaryManagedRuleListParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryManagedRuleListParamsMitigationProductDdos               RadarAttackLayer7SummaryManagedRuleListParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryManagedRuleListParamsMitigationProductWaf                RadarAttackLayer7SummaryManagedRuleListParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryManagedRuleListParamsMitigationProductBotManagement      RadarAttackLayer7SummaryManagedRuleListParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryManagedRuleListParamsMitigationProductAccessRules        RadarAttackLayer7SummaryManagedRuleListParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryManagedRuleListParamsMitigationProductIPReputation       RadarAttackLayer7SummaryManagedRuleListParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryManagedRuleListParamsMitigationProductAPIShield          RadarAttackLayer7SummaryManagedRuleListParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryManagedRuleListParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryManagedRuleListParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)
