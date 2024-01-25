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

// RadarEmailSecuritySpfSummaryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecuritySpfSummaryService] method instead.
type RadarEmailSecuritySpfSummaryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySpfSummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySpfSummaryService(opts ...option.RequestOption) (r *RadarEmailSecuritySpfSummaryService) {
	r = &RadarEmailSecuritySpfSummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per SPF validation.
func (r *RadarEmailSecuritySpfSummaryService) List(ctx context.Context, query RadarEmailSecuritySpfSummaryListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySpfSummaryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecuritySpfSummaryListResponse struct {
	Result  RadarEmailSecuritySpfSummaryListResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarEmailSecuritySpfSummaryListResponseJSON   `json:"-"`
}

// radarEmailSecuritySpfSummaryListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySpfSummaryListResponse]
type radarEmailSecuritySpfSummaryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpfSummaryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfSummaryListResponseResult struct {
	Meta     RadarEmailSecuritySpfSummaryListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySpfSummaryListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySpfSummaryListResponseResultJSON     `json:"-"`
}

// radarEmailSecuritySpfSummaryListResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySpfSummaryListResponseResult]
type radarEmailSecuritySpfSummaryListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpfSummaryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfSummaryListResponseResultMeta struct {
	DateRange      []RadarEmailSecuritySpfSummaryListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                           `json:"lastUpdated,required"`
	Normalization  string                                                           `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySpfSummaryListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecuritySpfSummaryListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySpfSummaryListResponseResultMeta]
type radarEmailSecuritySpfSummaryListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySpfSummaryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfSummaryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySpfSummaryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySpfSummaryListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySpfSummaryListResponseResultMetaDateRange]
type radarEmailSecuritySpfSummaryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpfSummaryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                        `json:"level"`
	JSON        radarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfo]
type radarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                         `json:"dataSource,required"`
	Description     string                                                                         `json:"description,required"`
	EventType       string                                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                                      `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySpfSummaryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfSummaryListResponseResultSummary0 struct {
	Fail string                                                     `json:"FAIL,required"`
	None string                                                     `json:"NONE,required"`
	Pass string                                                     `json:"PASS,required"`
	JSON radarEmailSecuritySpfSummaryListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecuritySpfSummaryListResponseResultSummary0JSON contains the JSON
// metadata for the struct [RadarEmailSecuritySpfSummaryListResponseResultSummary0]
type radarEmailSecuritySpfSummaryListResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpfSummaryListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpfSummaryListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySpfSummaryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySpfSummaryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecuritySpfSummaryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySpfSummaryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySpfSummaryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailSecuritySpfSummaryListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySpfSummaryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySpfSummaryListParamsArc string

const (
	RadarEmailSecuritySpfSummaryListParamsArcPass RadarEmailSecuritySpfSummaryListParamsArc = "PASS"
	RadarEmailSecuritySpfSummaryListParamsArcNone RadarEmailSecuritySpfSummaryListParamsArc = "NONE"
	RadarEmailSecuritySpfSummaryListParamsArcFail RadarEmailSecuritySpfSummaryListParamsArc = "FAIL"
)

type RadarEmailSecuritySpfSummaryListParamsDateRange string

const (
	RadarEmailSecuritySpfSummaryListParamsDateRange1d         RadarEmailSecuritySpfSummaryListParamsDateRange = "1d"
	RadarEmailSecuritySpfSummaryListParamsDateRange2d         RadarEmailSecuritySpfSummaryListParamsDateRange = "2d"
	RadarEmailSecuritySpfSummaryListParamsDateRange7d         RadarEmailSecuritySpfSummaryListParamsDateRange = "7d"
	RadarEmailSecuritySpfSummaryListParamsDateRange14d        RadarEmailSecuritySpfSummaryListParamsDateRange = "14d"
	RadarEmailSecuritySpfSummaryListParamsDateRange28d        RadarEmailSecuritySpfSummaryListParamsDateRange = "28d"
	RadarEmailSecuritySpfSummaryListParamsDateRange12w        RadarEmailSecuritySpfSummaryListParamsDateRange = "12w"
	RadarEmailSecuritySpfSummaryListParamsDateRange24w        RadarEmailSecuritySpfSummaryListParamsDateRange = "24w"
	RadarEmailSecuritySpfSummaryListParamsDateRange52w        RadarEmailSecuritySpfSummaryListParamsDateRange = "52w"
	RadarEmailSecuritySpfSummaryListParamsDateRange1dControl  RadarEmailSecuritySpfSummaryListParamsDateRange = "1dControl"
	RadarEmailSecuritySpfSummaryListParamsDateRange2dControl  RadarEmailSecuritySpfSummaryListParamsDateRange = "2dControl"
	RadarEmailSecuritySpfSummaryListParamsDateRange7dControl  RadarEmailSecuritySpfSummaryListParamsDateRange = "7dControl"
	RadarEmailSecuritySpfSummaryListParamsDateRange14dControl RadarEmailSecuritySpfSummaryListParamsDateRange = "14dControl"
	RadarEmailSecuritySpfSummaryListParamsDateRange28dControl RadarEmailSecuritySpfSummaryListParamsDateRange = "28dControl"
	RadarEmailSecuritySpfSummaryListParamsDateRange12wControl RadarEmailSecuritySpfSummaryListParamsDateRange = "12wControl"
	RadarEmailSecuritySpfSummaryListParamsDateRange24wControl RadarEmailSecuritySpfSummaryListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySpfSummaryListParamsDkim string

const (
	RadarEmailSecuritySpfSummaryListParamsDkimPass RadarEmailSecuritySpfSummaryListParamsDkim = "PASS"
	RadarEmailSecuritySpfSummaryListParamsDkimNone RadarEmailSecuritySpfSummaryListParamsDkim = "NONE"
	RadarEmailSecuritySpfSummaryListParamsDkimFail RadarEmailSecuritySpfSummaryListParamsDkim = "FAIL"
)

type RadarEmailSecuritySpfSummaryListParamsDmarc string

const (
	RadarEmailSecuritySpfSummaryListParamsDmarcPass RadarEmailSecuritySpfSummaryListParamsDmarc = "PASS"
	RadarEmailSecuritySpfSummaryListParamsDmarcNone RadarEmailSecuritySpfSummaryListParamsDmarc = "NONE"
	RadarEmailSecuritySpfSummaryListParamsDmarcFail RadarEmailSecuritySpfSummaryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySpfSummaryListParamsFormat string

const (
	RadarEmailSecuritySpfSummaryListParamsFormatJson RadarEmailSecuritySpfSummaryListParamsFormat = "JSON"
	RadarEmailSecuritySpfSummaryListParamsFormatCsv  RadarEmailSecuritySpfSummaryListParamsFormat = "CSV"
)
