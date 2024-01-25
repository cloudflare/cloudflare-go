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

// RadarAttackLayer7SummaryHTTPMethodService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarAttackLayer7SummaryHTTPMethodService] method instead.
type RadarAttackLayer7SummaryHTTPMethodService struct {
	Options []option.RequestOption
}

// NewRadarAttackLayer7SummaryHTTPMethodService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarAttackLayer7SummaryHTTPMethodService(opts ...option.RequestOption) (r *RadarAttackLayer7SummaryHTTPMethodService) {
	r = &RadarAttackLayer7SummaryHTTPMethodService{}
	r.Options = opts
	return
}

// Percentage distribution of attacks by http method used.
func (r *RadarAttackLayer7SummaryHTTPMethodService) List(ctx context.Context, query RadarAttackLayer7SummaryHTTPMethodListParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryHTTPMethodListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary/http_method"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7SummaryHTTPMethodListResponse struct {
	Result  RadarAttackLayer7SummaryHTTPMethodListResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarAttackLayer7SummaryHTTPMethodListResponseJSON   `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodListResponseJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryHTTPMethodListResponse]
type radarAttackLayer7SummaryHTTPMethodListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPMethodListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodListResponseResult struct {
	Meta     RadarAttackLayer7SummaryHTTPMethodListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryHTTPMethodListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryHTTPMethodListResponseResultJSON     `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodListResponseResultJSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryHTTPMethodListResponseResult]
type radarAttackLayer7SummaryHTTPMethodListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPMethodListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodListResponseResultMeta struct {
	DateRange      []RadarAttackLayer7SummaryHTTPMethodListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                 `json:"lastUpdated,required"`
	Normalization  string                                                                 `json:"normalization,required"`
	ConfidenceInfo RadarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryHTTPMethodListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryHTTPMethodListResponseResultMeta]
type radarAttackLayer7SummaryHTTPMethodListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPMethodListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryHTTPMethodListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarAttackLayer7SummaryHTTPMethodListResponseResultMetaDateRange]
type radarAttackLayer7SummaryHTTPMethodListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPMethodListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                              `json:"level"`
	JSON        radarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfo]
type radarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                               `json:"dataSource,required"`
	Description     string                                                                               `json:"description,required"`
	EventType       string                                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                                            `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryHTTPMethodListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodListResponseResultSummary0 struct {
	Get  string                                                           `json:"GET,required"`
	Post string                                                           `json:"POST,required"`
	JSON radarAttackLayer7SummaryHTTPMethodListResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryHTTPMethodListResponseResultSummary0JSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryHTTPMethodListResponseResultSummary0]
type radarAttackLayer7SummaryHTTPMethodListResponseResultSummary0JSON struct {
	Get         apijson.Field
	Post        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryHTTPMethodListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryHTTPMethodListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryHTTPMethodListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryHTTPMethodListParamsFormat] `query:"format"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarAttackLayer7SummaryHTTPMethodListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarAttackLayer7SummaryHTTPMethodListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of L7 mitigation products.
	MitigationProduct param.Field[[]RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProduct] `query:"mitigationProduct"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryHTTPMethodListParams]'s query
// parameters as `url.Values`.
func (r RadarAttackLayer7SummaryHTTPMethodListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryHTTPMethodListParamsDateRange string

const (
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange1d         RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "1d"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange2d         RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "2d"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange7d         RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "7d"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange14d        RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "14d"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange28d        RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "28d"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange12w        RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "12w"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange24w        RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "24w"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange52w        RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "52w"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange1dControl  RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange2dControl  RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange7dControl  RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange14dControl RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange28dControl RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange12wControl RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryHTTPMethodListParamsDateRange24wControl RadarAttackLayer7SummaryHTTPMethodListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7SummaryHTTPMethodListParamsFormat string

const (
	RadarAttackLayer7SummaryHTTPMethodListParamsFormatJson RadarAttackLayer7SummaryHTTPMethodListParamsFormat = "JSON"
	RadarAttackLayer7SummaryHTTPMethodListParamsFormatCsv  RadarAttackLayer7SummaryHTTPMethodListParamsFormat = "CSV"
)

type RadarAttackLayer7SummaryHTTPMethodListParamsHTTPVersion string

const (
	RadarAttackLayer7SummaryHTTPMethodListParamsHTTPVersionHttPv1 RadarAttackLayer7SummaryHTTPMethodListParamsHTTPVersion = "HTTPv1"
	RadarAttackLayer7SummaryHTTPMethodListParamsHTTPVersionHttPv2 RadarAttackLayer7SummaryHTTPMethodListParamsHTTPVersion = "HTTPv2"
	RadarAttackLayer7SummaryHTTPMethodListParamsHTTPVersionHttPv3 RadarAttackLayer7SummaryHTTPMethodListParamsHTTPVersion = "HTTPv3"
)

type RadarAttackLayer7SummaryHTTPMethodListParamsIPVersion string

const (
	RadarAttackLayer7SummaryHTTPMethodListParamsIPVersionIPv4 RadarAttackLayer7SummaryHTTPMethodListParamsIPVersion = "IPv4"
	RadarAttackLayer7SummaryHTTPMethodListParamsIPVersionIPv6 RadarAttackLayer7SummaryHTTPMethodListParamsIPVersion = "IPv6"
)

type RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProduct string

const (
	RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProductDdos               RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProduct = "DDOS"
	RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProductWaf                RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProduct = "WAF"
	RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProductBotManagement      RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProduct = "BOT_MANAGEMENT"
	RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProductAccessRules        RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProduct = "ACCESS_RULES"
	RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProductIPReputation       RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProduct = "IP_REPUTATION"
	RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProductAPIShield          RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProduct = "API_SHIELD"
	RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProductDataLossPrevention RadarAttackLayer7SummaryHTTPMethodListParamsMitigationProduct = "DATA_LOSS_PREVENTION"
)
