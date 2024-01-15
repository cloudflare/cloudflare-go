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

// RadarEmailSecurityArcSummaryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityArcSummaryService] method instead.
type RadarEmailSecurityArcSummaryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityArcSummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityArcSummaryService(opts ...option.RequestOption) (r *RadarEmailSecurityArcSummaryService) {
	r = &RadarEmailSecurityArcSummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per ARC validation.
func (r *RadarEmailSecurityArcSummaryService) List(ctx context.Context, query RadarEmailSecurityArcSummaryListParams, opts ...option.RequestOption) (res *RadarEmailSecurityArcSummaryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityArcSummaryListResponse struct {
	Result  RadarEmailSecurityArcSummaryListResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarEmailSecurityArcSummaryListResponseJSON   `json:"-"`
}

// radarEmailSecurityArcSummaryListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityArcSummaryListResponse]
type radarEmailSecurityArcSummaryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityArcSummaryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityArcSummaryListResponseResult struct {
	Meta     RadarEmailSecurityArcSummaryListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecurityArcSummaryListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecurityArcSummaryListResponseResultJSON     `json:"-"`
}

// radarEmailSecurityArcSummaryListResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecurityArcSummaryListResponseResult]
type radarEmailSecurityArcSummaryListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityArcSummaryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityArcSummaryListResponseResultMeta struct {
	DateRange      []RadarEmailSecurityArcSummaryListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                           `json:"lastUpdated,required"`
	Normalization  string                                                           `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityArcSummaryListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityArcSummaryListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecurityArcSummaryListResponseResultMeta]
type radarEmailSecurityArcSummaryListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityArcSummaryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityArcSummaryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityArcSummaryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityArcSummaryListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityArcSummaryListResponseResultMetaDateRange]
type radarEmailSecurityArcSummaryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityArcSummaryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                        `json:"level"`
	JSON        radarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfo]
type radarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                         `json:"dataSource,required"`
	Description     string                                                                         `json:"description,required"`
	EventType       string                                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                                      `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityArcSummaryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityArcSummaryListResponseResultSummary0 struct {
	Fail string                                                     `json:"FAIL,required"`
	None string                                                     `json:"NONE,required"`
	Pass string                                                     `json:"PASS,required"`
	JSON radarEmailSecurityArcSummaryListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecurityArcSummaryListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityArcSummaryListResponseResultSummary0]
type radarEmailSecurityArcSummaryListResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityArcSummaryListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityArcSummaryListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityArcSummaryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityArcSummaryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityArcSummaryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityArcSummaryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityArcSummaryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityArcSummaryListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityArcSummaryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecurityArcSummaryListParamsDateRange string

const (
	RadarEmailSecurityArcSummaryListParamsDateRange1d         RadarEmailSecurityArcSummaryListParamsDateRange = "1d"
	RadarEmailSecurityArcSummaryListParamsDateRange2d         RadarEmailSecurityArcSummaryListParamsDateRange = "2d"
	RadarEmailSecurityArcSummaryListParamsDateRange7d         RadarEmailSecurityArcSummaryListParamsDateRange = "7d"
	RadarEmailSecurityArcSummaryListParamsDateRange14d        RadarEmailSecurityArcSummaryListParamsDateRange = "14d"
	RadarEmailSecurityArcSummaryListParamsDateRange28d        RadarEmailSecurityArcSummaryListParamsDateRange = "28d"
	RadarEmailSecurityArcSummaryListParamsDateRange12w        RadarEmailSecurityArcSummaryListParamsDateRange = "12w"
	RadarEmailSecurityArcSummaryListParamsDateRange24w        RadarEmailSecurityArcSummaryListParamsDateRange = "24w"
	RadarEmailSecurityArcSummaryListParamsDateRange52w        RadarEmailSecurityArcSummaryListParamsDateRange = "52w"
	RadarEmailSecurityArcSummaryListParamsDateRange1dControl  RadarEmailSecurityArcSummaryListParamsDateRange = "1dControl"
	RadarEmailSecurityArcSummaryListParamsDateRange2dControl  RadarEmailSecurityArcSummaryListParamsDateRange = "2dControl"
	RadarEmailSecurityArcSummaryListParamsDateRange7dControl  RadarEmailSecurityArcSummaryListParamsDateRange = "7dControl"
	RadarEmailSecurityArcSummaryListParamsDateRange14dControl RadarEmailSecurityArcSummaryListParamsDateRange = "14dControl"
	RadarEmailSecurityArcSummaryListParamsDateRange28dControl RadarEmailSecurityArcSummaryListParamsDateRange = "28dControl"
	RadarEmailSecurityArcSummaryListParamsDateRange12wControl RadarEmailSecurityArcSummaryListParamsDateRange = "12wControl"
	RadarEmailSecurityArcSummaryListParamsDateRange24wControl RadarEmailSecurityArcSummaryListParamsDateRange = "24wControl"
)

type RadarEmailSecurityArcSummaryListParamsDkim string

const (
	RadarEmailSecurityArcSummaryListParamsDkimPass RadarEmailSecurityArcSummaryListParamsDkim = "PASS"
	RadarEmailSecurityArcSummaryListParamsDkimNone RadarEmailSecurityArcSummaryListParamsDkim = "NONE"
	RadarEmailSecurityArcSummaryListParamsDkimFail RadarEmailSecurityArcSummaryListParamsDkim = "FAIL"
)

type RadarEmailSecurityArcSummaryListParamsDmarc string

const (
	RadarEmailSecurityArcSummaryListParamsDmarcPass RadarEmailSecurityArcSummaryListParamsDmarc = "PASS"
	RadarEmailSecurityArcSummaryListParamsDmarcNone RadarEmailSecurityArcSummaryListParamsDmarc = "NONE"
	RadarEmailSecurityArcSummaryListParamsDmarcFail RadarEmailSecurityArcSummaryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityArcSummaryListParamsFormat string

const (
	RadarEmailSecurityArcSummaryListParamsFormatJson RadarEmailSecurityArcSummaryListParamsFormat = "JSON"
	RadarEmailSecurityArcSummaryListParamsFormatCsv  RadarEmailSecurityArcSummaryListParamsFormat = "CSV"
)

type RadarEmailSecurityArcSummaryListParamsSpf string

const (
	RadarEmailSecurityArcSummaryListParamsSpfPass RadarEmailSecurityArcSummaryListParamsSpf = "PASS"
	RadarEmailSecurityArcSummaryListParamsSpfNone RadarEmailSecurityArcSummaryListParamsSpf = "NONE"
	RadarEmailSecurityArcSummaryListParamsSpfFail RadarEmailSecurityArcSummaryListParamsSpf = "FAIL"
)
