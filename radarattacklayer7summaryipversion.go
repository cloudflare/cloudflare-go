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

// RadarAttackLayer7SummaryIPVersionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7SummaryIPVersionService] method instead.
type RadarAttackLayer7SummaryIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7SummaryIPVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7SummaryIPVersionService(opts ...option.RequestOption) (r *RadarAttackLayer7SummaryIPVersionService) {
	r = &RadarAttackLayer7SummaryIPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by ip version used.
func (r *RadarAttackLayer7SummaryIPVersionService) List(ctx context.Context, query RadarAttackLayer7SummaryIPVersionListParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryIPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7SummaryIPVersionListResponse struct {
	Result  RadarAttackLayer7SummaryIPVersionListResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarAttackLayer7SummaryIPVersionListResponseJSON   `json:"-"`
}

// radarAttackLayer7SummaryIPVersionListResponseJSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryIPVersionListResponse]
type radarAttackLayer7SummaryIPVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryIPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionListResponseResult struct {
	Meta     RadarAttackLayer7SummaryIPVersionListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryIPVersionListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryIPVersionListResponseResultJSON     `json:"-"`
}

// radarAttackLayer7SummaryIPVersionListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryIPVersionListResponseResult]
type radarAttackLayer7SummaryIPVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryIPVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionListResponseResultMeta struct {
	DateRange      []RadarAttackLayer7SummaryIPVersionListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                `json:"lastUpdated,required"`
	Normalization  string                                                                `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryIPVersionListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryIPVersionListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryIPVersionListResponseResultMeta]
type radarAttackLayer7SummaryIPVersionListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryIPVersionListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryIPVersionListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryIPVersionListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryIPVersionListResponseResultMetaDateRange]
type radarAttackLayer7SummaryIPVersionListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryIPVersionListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                             `json:"level"`
	JSON        radarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfo]
type radarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                              `json:"dataSource,required"`
	Description     string                                                                              `json:"description,required"`
	EventType       string                                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                                           `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryIPVersionListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionListResponseResultSummary0 struct {
	IPv4 string                                                          `json:"IPv4,required"`
	IPv6 string                                                          `json:"IPv6,required"`
	JSON radarAttackLayer7SummaryIPVersionListResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryIPVersionListResponseResultSummary0JSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryIPVersionListResponseResultSummary0]
type radarAttackLayer7SummaryIPVersionListResponseResultSummary0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryIPVersionListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryIPVersionListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryIPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryIPVersionListParamsFormat] `query:"format"`
	// Filter for http method.
	HTTPMethod param.Field[[]RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod] `query:"httpMethod"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryIPVersionListParamsHTTPVersion] `query:"httpVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryIPVersionListParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryIPVersionListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7SummaryIPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryIPVersionListParamsDateRange string

const (
	RadarAttackLayer7SummaryIPVersionListParamsDateRange1d         RadarAttackLayer7SummaryIPVersionListParamsDateRange = "1d"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange2d         RadarAttackLayer7SummaryIPVersionListParamsDateRange = "2d"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange7d         RadarAttackLayer7SummaryIPVersionListParamsDateRange = "7d"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange14d        RadarAttackLayer7SummaryIPVersionListParamsDateRange = "14d"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange28d        RadarAttackLayer7SummaryIPVersionListParamsDateRange = "28d"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange12w        RadarAttackLayer7SummaryIPVersionListParamsDateRange = "12w"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange24w        RadarAttackLayer7SummaryIPVersionListParamsDateRange = "24w"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange52w        RadarAttackLayer7SummaryIPVersionListParamsDateRange = "52w"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange1dControl  RadarAttackLayer7SummaryIPVersionListParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange2dControl  RadarAttackLayer7SummaryIPVersionListParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange7dControl  RadarAttackLayer7SummaryIPVersionListParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange14dControl RadarAttackLayer7SummaryIPVersionListParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange28dControl RadarAttackLayer7SummaryIPVersionListParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange12wControl RadarAttackLayer7SummaryIPVersionListParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryIPVersionListParamsDateRange24wControl RadarAttackLayer7SummaryIPVersionListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7SummaryIPVersionListParamsFormat string

const (
	RadarAttackLayer7SummaryIPVersionListParamsFormatJson RadarAttackLayer7SummaryIPVersionListParamsFormat = "JSON"
	RadarAttackLayer7SummaryIPVersionListParamsFormatCsv  RadarAttackLayer7SummaryIPVersionListParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod string

const (
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodGet             RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "GET"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodPost            RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "POST"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodDelete          RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "DELETE"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodPut             RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "PUT"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodHead            RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "HEAD"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodPurge           RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "PURGE"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodOptions         RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "OPTIONS"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodPropfind        RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "PROPFIND"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodMkcol           RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "MKCOL"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodPatch           RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "PATCH"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodACL             RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "ACL"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodBcopy           RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "BCOPY"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodBdelete         RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "BDELETE"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodBmove           RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "BMOVE"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodBpropfind       RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "BPROPFIND"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodBproppatch      RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "BPROPPATCH"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodCheckin         RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "CHECKIN"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodCheckout        RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "CHECKOUT"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodConnect         RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "CONNECT"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodCopy            RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "COPY"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodLabel           RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "LABEL"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodLock            RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "LOCK"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodMerge           RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "MERGE"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodMkactivity      RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "MKACTIVITY"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodMkworkspace     RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "MKWORKSPACE"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodMove            RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "MOVE"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodNotify          RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "NOTIFY"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodOrderpatch      RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "ORDERPATCH"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodPoll            RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "POLL"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodProppatch       RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "PROPPATCH"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodReport          RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "REPORT"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodSearch          RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "SEARCH"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodSubscribe       RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "SUBSCRIBE"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodTrace           RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "TRACE"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodUncheckout      RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "UNCHECKOUT"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodUnlock          RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "UNLOCK"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodUnsubscribe     RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "UNSUBSCRIBE"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodUpdate          RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "UPDATE"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodVersioncontrol  RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "VERSIONCONTROL"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodBaselinecontrol RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "BASELINECONTROL"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodXmsenumatts     RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "XMSENUMATTS"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodRpcOutData      RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "RPC_OUT_DATA"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodRpcInData       RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "RPC_IN_DATA"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodJson            RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "JSON"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodCook            RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "COOK"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPMethodTrack           RadarAttackLayer7SummaryIPVersionListParamsHTTPMethod = "TRACK"
)

type RadarAttackLayer7SummaryIPVersionListParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryIPVersionListParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryIPVersionListParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryIPVersionListParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryIPVersionListParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryIPVersionListParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7SummaryIPVersionListParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryIPVersionListParamsMitigationProductDdos               RadarAttackLayer7SummaryIPVersionListParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryIPVersionListParamsMitigationProductWaf                RadarAttackLayer7SummaryIPVersionListParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryIPVersionListParamsMitigationProductBotManagement      RadarAttackLayer7SummaryIPVersionListParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryIPVersionListParamsMitigationProductAccessRules        RadarAttackLayer7SummaryIPVersionListParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryIPVersionListParamsMitigationProductIPReputation       RadarAttackLayer7SummaryIPVersionListParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryIPVersionListParamsMitigationProductAPIShield          RadarAttackLayer7SummaryIPVersionListParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryIPVersionListParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryIPVersionListParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)
