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

// RadarAttackLayer7SummaryService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarAttackLayer7SummaryService] method instead.
type RadarAttackLayer7SummaryService struct {
	Options      []option.RequestOption
	HTTPMethod   *RadarAttackLayer7SummaryHTTPMethodService
	HTTPVersion  *RadarAttackLayer7SummaryHTTPVersionService
	IPVersion    *RadarAttackLayer7SummaryIPVersionService
	ManagedRules *RadarAttackLayer7SummaryManagedRuleService
}

// NewRadarAttackLayer7SummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarAttackLayer7SummaryService(opts ...option.RequestOption) (r *RadarAttackLayer7SummaryService) {
	r = &RadarAttackLayer7SummaryService{}
	r.Options = opts
	r.HTTPMethod = NewRadarAttackLayer7SummaryHTTPMethodService(opts...)
	r.HTTPVersion = NewRadarAttackLayer7SummaryHTTPVersionService(opts...)
	r.IPVersion = NewRadarAttackLayer7SummaryIPVersionService(opts...)
	r.ManagedRules = NewRadarAttackLayer7SummaryManagedRuleService(opts...)
	return
}

// Percentage distribution of mitigation techniques in Layer 7 attacks.
func (r *RadarAttackLayer7SummaryService) List(ctx context.Context, query RadarAttackLayer7SummaryListParams, opts ...option.RequestOption) (res *RadarAttackLayer7SummaryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/attacks/layer7/summary"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarAttackLayer7SummaryListResponse struct {
	Result  RadarAttackLayer7SummaryListResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarAttackLayer7SummaryListResponseJSON   `json:"-"`
}

// radarAttackLayer7SummaryListResponseJSON contains the JSON metadata for the
// struct [RadarAttackLayer7SummaryListResponse]
type radarAttackLayer7SummaryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryListResponseResult struct {
	Meta     RadarAttackLayer7SummaryListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarAttackLayer7SummaryListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarAttackLayer7SummaryListResponseResultJSON     `json:"-"`
}

// radarAttackLayer7SummaryListResponseResultJSON contains the JSON metadata for
// the struct [RadarAttackLayer7SummaryListResponseResult]
type radarAttackLayer7SummaryListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryListResponseResultMeta struct {
	DateRange      []RadarAttackLayer7SummaryListResponseResultMetaDateRange    `json:"dateRange,required"`
	ConfidenceInfo RadarAttackLayer7SummaryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarAttackLayer7SummaryListResponseResultMetaJSON           `json:"-"`
}

// radarAttackLayer7SummaryListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarAttackLayer7SummaryListResponseResultMeta]
type radarAttackLayer7SummaryListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarAttackLayer7SummaryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarAttackLayer7SummaryListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarAttackLayer7SummaryListResponseResultMetaDateRange]
type radarAttackLayer7SummaryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarAttackLayer7SummaryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarAttackLayer7SummaryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarAttackLayer7SummaryListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarAttackLayer7SummaryListResponseResultMetaConfidenceInfo]
type radarAttackLayer7SummaryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarAttackLayer7SummaryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarAttackLayer7SummaryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarAttackLayer7SummaryListResponseResultMetaConfidenceInfoAnnotation]
type radarAttackLayer7SummaryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarAttackLayer7SummaryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryListResponseResultSummary0 struct {
	AccessRules        string                                                 `json:"ACCESS_RULES,required"`
	APIShield          string                                                 `json:"API_SHIELD,required"`
	BotManagement      string                                                 `json:"BOT_MANAGEMENT,required"`
	DataLossPrevention string                                                 `json:"DATA_LOSS_PREVENTION,required"`
	Ddos               string                                                 `json:"DDOS,required"`
	IPReputation       string                                                 `json:"IP_REPUTATION,required"`
	Waf                string                                                 `json:"WAF,required"`
	JSON               radarAttackLayer7SummaryListResponseResultSummary0JSON `json:"-"`
}

// radarAttackLayer7SummaryListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarAttackLayer7SummaryListResponseResultSummary0]
type radarAttackLayer7SummaryListResponseResultSummary0JSON struct {
	AccessRules        apijson.Field
	APIShield          apijson.Field
	BotManagement      apijson.Field
	DataLossPrevention apijson.Field
	Ddos               apijson.Field
	IPReputation       apijson.Field
	Waf                apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *RadarAttackLayer7SummaryListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarAttackLayer7SummaryListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarAttackLayer7SummaryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarAttackLayer7SummaryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarAttackLayer7SummaryListParams]'s query parameters as
// `url.Values`.
func (r RadarAttackLayer7SummaryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarAttackLayer7SummaryListParamsDateRange string

const (
	RadarAttackLayer7SummaryListParamsDateRange1d         RadarAttackLayer7SummaryListParamsDateRange = "1d"
	RadarAttackLayer7SummaryListParamsDateRange2d         RadarAttackLayer7SummaryListParamsDateRange = "2d"
	RadarAttackLayer7SummaryListParamsDateRange7d         RadarAttackLayer7SummaryListParamsDateRange = "7d"
	RadarAttackLayer7SummaryListParamsDateRange14d        RadarAttackLayer7SummaryListParamsDateRange = "14d"
	RadarAttackLayer7SummaryListParamsDateRange28d        RadarAttackLayer7SummaryListParamsDateRange = "28d"
	RadarAttackLayer7SummaryListParamsDateRange12w        RadarAttackLayer7SummaryListParamsDateRange = "12w"
	RadarAttackLayer7SummaryListParamsDateRange24w        RadarAttackLayer7SummaryListParamsDateRange = "24w"
	RadarAttackLayer7SummaryListParamsDateRange52w        RadarAttackLayer7SummaryListParamsDateRange = "52w"
	RadarAttackLayer7SummaryListParamsDateRange1dControl  RadarAttackLayer7SummaryListParamsDateRange = "1dControl"
	RadarAttackLayer7SummaryListParamsDateRange2dControl  RadarAttackLayer7SummaryListParamsDateRange = "2dControl"
	RadarAttackLayer7SummaryListParamsDateRange7dControl  RadarAttackLayer7SummaryListParamsDateRange = "7dControl"
	RadarAttackLayer7SummaryListParamsDateRange14dControl RadarAttackLayer7SummaryListParamsDateRange = "14dControl"
	RadarAttackLayer7SummaryListParamsDateRange28dControl RadarAttackLayer7SummaryListParamsDateRange = "28dControl"
	RadarAttackLayer7SummaryListParamsDateRange12wControl RadarAttackLayer7SummaryListParamsDateRange = "12wControl"
	RadarAttackLayer7SummaryListParamsDateRange24wControl RadarAttackLayer7SummaryListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarAttackLayer7SummaryListParamsFormat string

const (
	RadarAttackLayer7SummaryListParamsFormatJson RadarAttackLayer7SummaryListParamsFormat = "JSON"
	RadarAttackLayer7SummaryListParamsFormatCsv  RadarAttackLayer7SummaryListParamsFormat = "CSV"
)
