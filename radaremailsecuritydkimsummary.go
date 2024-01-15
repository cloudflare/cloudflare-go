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

// RadarEmailSecurityDkimSummaryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityDkimSummaryService] method instead.
type RadarEmailSecurityDkimSummaryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityDkimSummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityDkimSummaryService(opts ...option.RequestOption) (r *RadarEmailSecurityDkimSummaryService) {
	r = &RadarEmailSecurityDkimSummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per DKIM validation.
func (r *RadarEmailSecurityDkimSummaryService) List(ctx context.Context, query RadarEmailSecurityDkimSummaryListParams, opts ...option.RequestOption) (res *RadarEmailSecurityDkimSummaryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityDkimSummaryListResponse struct {
	Result  RadarEmailSecurityDkimSummaryListResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailSecurityDkimSummaryListResponseJSON   `json:"-"`
}

// radarEmailSecurityDkimSummaryListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityDkimSummaryListResponse]
type radarEmailSecurityDkimSummaryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDkimSummaryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDkimSummaryListResponseResult struct {
	Meta     RadarEmailSecurityDkimSummaryListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecurityDkimSummaryListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecurityDkimSummaryListResponseResultJSON     `json:"-"`
}

// radarEmailSecurityDkimSummaryListResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecurityDkimSummaryListResponseResult]
type radarEmailSecurityDkimSummaryListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDkimSummaryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDkimSummaryListResponseResultMeta struct {
	DateRange      []RadarEmailSecurityDkimSummaryListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                            `json:"lastUpdated,required"`
	Normalization  string                                                            `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityDkimSummaryListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityDkimSummaryListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecurityDkimSummaryListResponseResultMeta]
type radarEmailSecurityDkimSummaryListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityDkimSummaryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDkimSummaryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityDkimSummaryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityDkimSummaryListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityDkimSummaryListResponseResultMetaDateRange]
type radarEmailSecurityDkimSummaryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDkimSummaryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                         `json:"level"`
	JSON        radarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfo]
type radarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                          `json:"dataSource,required"`
	Description     string                                                                          `json:"description,required"`
	EventType       string                                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                       `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityDkimSummaryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDkimSummaryListResponseResultSummary0 struct {
	Fail string                                                      `json:"FAIL,required"`
	None string                                                      `json:"NONE,required"`
	Pass string                                                      `json:"PASS,required"`
	JSON radarEmailSecurityDkimSummaryListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecurityDkimSummaryListResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityDkimSummaryListResponseResultSummary0]
type radarEmailSecurityDkimSummaryListResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDkimSummaryListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDkimSummaryListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityDkimSummaryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityDkimSummaryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityDkimSummaryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityDkimSummaryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityDkimSummaryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityDkimSummaryListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityDkimSummaryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecurityDkimSummaryListParamsArc string

const (
	RadarEmailSecurityDkimSummaryListParamsArcPass RadarEmailSecurityDkimSummaryListParamsArc = "PASS"
	RadarEmailSecurityDkimSummaryListParamsArcNone RadarEmailSecurityDkimSummaryListParamsArc = "NONE"
	RadarEmailSecurityDkimSummaryListParamsArcFail RadarEmailSecurityDkimSummaryListParamsArc = "FAIL"
)

type RadarEmailSecurityDkimSummaryListParamsDateRange string

const (
	RadarEmailSecurityDkimSummaryListParamsDateRange1d         RadarEmailSecurityDkimSummaryListParamsDateRange = "1d"
	RadarEmailSecurityDkimSummaryListParamsDateRange2d         RadarEmailSecurityDkimSummaryListParamsDateRange = "2d"
	RadarEmailSecurityDkimSummaryListParamsDateRange7d         RadarEmailSecurityDkimSummaryListParamsDateRange = "7d"
	RadarEmailSecurityDkimSummaryListParamsDateRange14d        RadarEmailSecurityDkimSummaryListParamsDateRange = "14d"
	RadarEmailSecurityDkimSummaryListParamsDateRange28d        RadarEmailSecurityDkimSummaryListParamsDateRange = "28d"
	RadarEmailSecurityDkimSummaryListParamsDateRange12w        RadarEmailSecurityDkimSummaryListParamsDateRange = "12w"
	RadarEmailSecurityDkimSummaryListParamsDateRange24w        RadarEmailSecurityDkimSummaryListParamsDateRange = "24w"
	RadarEmailSecurityDkimSummaryListParamsDateRange52w        RadarEmailSecurityDkimSummaryListParamsDateRange = "52w"
	RadarEmailSecurityDkimSummaryListParamsDateRange1dControl  RadarEmailSecurityDkimSummaryListParamsDateRange = "1dControl"
	RadarEmailSecurityDkimSummaryListParamsDateRange2dControl  RadarEmailSecurityDkimSummaryListParamsDateRange = "2dControl"
	RadarEmailSecurityDkimSummaryListParamsDateRange7dControl  RadarEmailSecurityDkimSummaryListParamsDateRange = "7dControl"
	RadarEmailSecurityDkimSummaryListParamsDateRange14dControl RadarEmailSecurityDkimSummaryListParamsDateRange = "14dControl"
	RadarEmailSecurityDkimSummaryListParamsDateRange28dControl RadarEmailSecurityDkimSummaryListParamsDateRange = "28dControl"
	RadarEmailSecurityDkimSummaryListParamsDateRange12wControl RadarEmailSecurityDkimSummaryListParamsDateRange = "12wControl"
	RadarEmailSecurityDkimSummaryListParamsDateRange24wControl RadarEmailSecurityDkimSummaryListParamsDateRange = "24wControl"
)

type RadarEmailSecurityDkimSummaryListParamsDmarc string

const (
	RadarEmailSecurityDkimSummaryListParamsDmarcPass RadarEmailSecurityDkimSummaryListParamsDmarc = "PASS"
	RadarEmailSecurityDkimSummaryListParamsDmarcNone RadarEmailSecurityDkimSummaryListParamsDmarc = "NONE"
	RadarEmailSecurityDkimSummaryListParamsDmarcFail RadarEmailSecurityDkimSummaryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityDkimSummaryListParamsFormat string

const (
	RadarEmailSecurityDkimSummaryListParamsFormatJson RadarEmailSecurityDkimSummaryListParamsFormat = "JSON"
	RadarEmailSecurityDkimSummaryListParamsFormatCsv  RadarEmailSecurityDkimSummaryListParamsFormat = "CSV"
)

type RadarEmailSecurityDkimSummaryListParamsSpf string

const (
	RadarEmailSecurityDkimSummaryListParamsSpfPass RadarEmailSecurityDkimSummaryListParamsSpf = "PASS"
	RadarEmailSecurityDkimSummaryListParamsSpfNone RadarEmailSecurityDkimSummaryListParamsSpf = "NONE"
	RadarEmailSecurityDkimSummaryListParamsSpfFail RadarEmailSecurityDkimSummaryListParamsSpf = "FAIL"
)
