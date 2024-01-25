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

// RadarAttackLayer7SummaryHTTPVersionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7SummaryHTTPVersionService] method instead.
type RadarAttackLayer7SummaryHTTPVersionService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7SummaryHTTPVersionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer7SummaryHTTPVersionService(opts ...option.RequestOption) (r *RadarAttackLayer7SummaryHTTPVersionService) {
	r = &RadarAttackLayer7SummaryHTTPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by http version used.
func (r *RadarAttackLayer7SummaryHTTPVersionService) List(ctx context.Context, query RadarAttackLayer7SummaryHTTPVersionListParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryHTTPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7SummaryHTTPVersionListResponse struct {
	Result  RadarAttackLayer7SummaryHTTPVersionListResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarAttackLayer7SummaryHTTPVersionListResponseJSON   `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionListResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryHTTPVersionListResponse]
type radarAttackLayer7SummaryHTTPVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionListResponseResult struct {
	Meta     RadarAttackLayer7SummaryHTTPVersionListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryHTTPVersionListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryHTTPVersionListResponseResultJSON     `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryHTTPVersionListResponseResult]
type radarAttackLayer7SummaryHTTPVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionListResponseResultMeta struct {
	DateRange      []RadarAttackLayer7SummaryHTTPVersionListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                  `json:"lastUpdated,required"`
	Normalization  string                                                                  `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryHTTPVersionListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryHTTPVersionListResponseResultMeta]
type radarAttackLayer7SummaryHTTPVersionListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPVersionListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryHTTPVersionListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryHTTPVersionListResponseResultMetaDateRange]
type radarAttackLayer7SummaryHTTPVersionListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPVersionListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                               `json:"level"`
	JSON        radarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfo]
type radarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                `json:"dataSource,required"`
	Description     string                                                                                `json:"description,required"`
	EventType       string                                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                                             `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryHTTPVersionListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionListResponseResultSummary0 struct {
	HTTP1X string                                                            `json:"HTTP/1.x,required"`
	HTTP2  string                                                            `json:"HTTP/2,required"`
	HTTP3  string                                                            `json:"HTTP/3,required"`
	JSON   radarAttackLayer7SummaryHTTPVersionListResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPVersionListResponseResultSummary0JSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryHTTPVersionListResponseResultSummary0]
type radarAttackLayer7SummaryHTTPVersionListResponseResultSummary0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPVersionListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPVersionListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryHTTPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryHTTPVersionListParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod] `query:"httpMethod"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7SummaryHTTPVersionListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryHTTPVersionListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7SummaryHTTPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryHTTPVersionListParamsDateRange string

const (
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange1d         RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "1d"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange2d         RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "2d"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange7d         RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "7d"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange14d        RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "14d"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange28d        RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "28d"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange12w        RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "12w"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange24w        RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "24w"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange52w        RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "52w"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange1dControl  RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange2dControl  RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange7dControl  RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange14dControl RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange28dControl RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange12wControl RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryHTTPVersionListParamsDateRange24wControl RadarAttackLayer7SummaryHTTPVersionListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7SummaryHTTPVersionListParamsFormat string

const (
	RadarAttackLayer7SummaryHTTPVersionListParamsFormatJson RadarAttackLayer7SummaryHTTPVersionListParamsFormat = "JSON"
	RadarAttackLayer7SummaryHTTPVersionListParamsFormatCsv  RadarAttackLayer7SummaryHTTPVersionListParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodGet             RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodPost            RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodDelete          RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodPut             RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodHead            RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodPurge           RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodOptions         RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodPropfind        RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodMkcol           RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodPatch           RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodACL             RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodBcopy           RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodBdelete         RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodBmove           RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodCheckin         RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodCheckout        RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodConnect         RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodCopy            RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodLabel           RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodLock            RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodMerge           RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodMove            RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodNotify          RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodPoll            RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodProppatch       RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodReport          RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodSearch          RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodTrace           RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodUnlock          RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodUpdate          RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodJson            RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodCook            RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethodTrack           RadarAttackLayer7SummaryHTTPVersionListParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7SummaryHTTPVersionListParamsIPVersion string

const (
	RadarAttackLayer7SummaryHTTPVersionListParamsIPVersionIPv4 RadarAttackLayer7SummaryHTTPVersionListParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryHTTPVersionListParamsIPVersionIPv6 RadarAttackLayer7SummaryHTTPVersionListParamsIPVersion = "IPv6"
)

type RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProductDdos               RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProductWaf                RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProductBotManagement      RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProductAccessRules        RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProductIPReputation       RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProductAPIShield          RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryHTTPVersionListParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)
