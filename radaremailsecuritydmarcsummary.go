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

// RadarEmailSecurityDmarcSummaryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityDmarcSummaryService] method instead.
type RadarEmailSecurityDmarcSummaryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityDmarcSummaryService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityDmarcSummaryService(opts ...option.RequestOption) (r *RadarEmailSecurityDmarcSummaryService) {
	r = &RadarEmailSecurityDmarcSummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per DMARC validation.
func (r *RadarEmailSecurityDmarcSummaryService) List(ctx context.Context, query RadarEmailSecurityDmarcSummaryListParams, opts ...option.RequestOption) (res *RadarEmailSecurityDmarcSummaryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityDmarcSummaryListResponse struct {
	Result  RadarEmailSecurityDmarcSummaryListResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarEmailSecurityDmarcSummaryListResponseJSON   `json:"-"`
}

// radarEmailSecurityDmarcSummaryListResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityDmarcSummaryListResponse]
type radarEmailSecurityDmarcSummaryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDmarcSummaryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDmarcSummaryListResponseResult struct {
	Meta     RadarEmailSecurityDmarcSummaryListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecurityDmarcSummaryListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecurityDmarcSummaryListResponseResultJSON     `json:"-"`
}

// radarEmailSecurityDmarcSummaryListResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecurityDmarcSummaryListResponseResult]
type radarEmailSecurityDmarcSummaryListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDmarcSummaryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDmarcSummaryListResponseResultMeta struct {
	DateRange      []RadarEmailSecurityDmarcSummaryListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                             `json:"lastUpdated,required"`
	Normalization  string                                                             `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityDmarcSummaryListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityDmarcSummaryListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecurityDmarcSummaryListResponseResultMeta]
type radarEmailSecurityDmarcSummaryListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityDmarcSummaryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDmarcSummaryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                         `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityDmarcSummaryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityDmarcSummaryListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityDmarcSummaryListResponseResultMetaDateRange]
type radarEmailSecurityDmarcSummaryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDmarcSummaryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                          `json:"level"`
	JSON        radarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfo]
type radarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                           `json:"dataSource,required"`
	Description     string                                                                           `json:"description,required"`
	EventType       string                                                                           `json:"eventType,required"`
	IsInstantaneous interface{}                                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                                        `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityDmarcSummaryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDmarcSummaryListResponseResultSummary0 struct {
	Fail string                                                       `json:"FAIL,required"`
	None string                                                       `json:"NONE,required"`
	Pass string                                                       `json:"PASS,required"`
	JSON radarEmailSecurityDmarcSummaryListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecurityDmarcSummaryListResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityDmarcSummaryListResponseResultSummary0]
type radarEmailSecurityDmarcSummaryListResponseResultSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityDmarcSummaryListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityDmarcSummaryListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityDmarcSummaryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityDmarcSummaryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityDmarcSummaryListParamsDkim] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityDmarcSummaryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityDmarcSummaryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityDmarcSummaryListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityDmarcSummaryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecurityDmarcSummaryListParamsArc string

const (
	RadarEmailSecurityDmarcSummaryListParamsArcPass RadarEmailSecurityDmarcSummaryListParamsArc = "PASS"
	RadarEmailSecurityDmarcSummaryListParamsArcNone RadarEmailSecurityDmarcSummaryListParamsArc = "NONE"
	RadarEmailSecurityDmarcSummaryListParamsArcFail RadarEmailSecurityDmarcSummaryListParamsArc = "FAIL"
)

type RadarEmailSecurityDmarcSummaryListParamsDateRange string

const (
	RadarEmailSecurityDmarcSummaryListParamsDateRange1d         RadarEmailSecurityDmarcSummaryListParamsDateRange = "1d"
	RadarEmailSecurityDmarcSummaryListParamsDateRange2d         RadarEmailSecurityDmarcSummaryListParamsDateRange = "2d"
	RadarEmailSecurityDmarcSummaryListParamsDateRange7d         RadarEmailSecurityDmarcSummaryListParamsDateRange = "7d"
	RadarEmailSecurityDmarcSummaryListParamsDateRange14d        RadarEmailSecurityDmarcSummaryListParamsDateRange = "14d"
	RadarEmailSecurityDmarcSummaryListParamsDateRange28d        RadarEmailSecurityDmarcSummaryListParamsDateRange = "28d"
	RadarEmailSecurityDmarcSummaryListParamsDateRange12w        RadarEmailSecurityDmarcSummaryListParamsDateRange = "12w"
	RadarEmailSecurityDmarcSummaryListParamsDateRange24w        RadarEmailSecurityDmarcSummaryListParamsDateRange = "24w"
	RadarEmailSecurityDmarcSummaryListParamsDateRange52w        RadarEmailSecurityDmarcSummaryListParamsDateRange = "52w"
	RadarEmailSecurityDmarcSummaryListParamsDateRange1dControl  RadarEmailSecurityDmarcSummaryListParamsDateRange = "1dControl"
	RadarEmailSecurityDmarcSummaryListParamsDateRange2dControl  RadarEmailSecurityDmarcSummaryListParamsDateRange = "2dControl"
	RadarEmailSecurityDmarcSummaryListParamsDateRange7dControl  RadarEmailSecurityDmarcSummaryListParamsDateRange = "7dControl"
	RadarEmailSecurityDmarcSummaryListParamsDateRange14dControl RadarEmailSecurityDmarcSummaryListParamsDateRange = "14dControl"
	RadarEmailSecurityDmarcSummaryListParamsDateRange28dControl RadarEmailSecurityDmarcSummaryListParamsDateRange = "28dControl"
	RadarEmailSecurityDmarcSummaryListParamsDateRange12wControl RadarEmailSecurityDmarcSummaryListParamsDateRange = "12wControl"
	RadarEmailSecurityDmarcSummaryListParamsDateRange24wControl RadarEmailSecurityDmarcSummaryListParamsDateRange = "24wControl"
)

type RadarEmailSecurityDmarcSummaryListParamsDkim string

const (
	RadarEmailSecurityDmarcSummaryListParamsDkimPass RadarEmailSecurityDmarcSummaryListParamsDkim = "PASS"
	RadarEmailSecurityDmarcSummaryListParamsDkimNone RadarEmailSecurityDmarcSummaryListParamsDkim = "NONE"
	RadarEmailSecurityDmarcSummaryListParamsDkimFail RadarEmailSecurityDmarcSummaryListParamsDkim = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityDmarcSummaryListParamsFormat string

const (
	RadarEmailSecurityDmarcSummaryListParamsFormatJson RadarEmailSecurityDmarcSummaryListParamsFormat = "JSON"
	RadarEmailSecurityDmarcSummaryListParamsFormatCsv  RadarEmailSecurityDmarcSummaryListParamsFormat = "CSV"
)

type RadarEmailSecurityDmarcSummaryListParamsSpf string

const (
	RadarEmailSecurityDmarcSummaryListParamsSpfPass RadarEmailSecurityDmarcSummaryListParamsSpf = "PASS"
	RadarEmailSecurityDmarcSummaryListParamsSpfNone RadarEmailSecurityDmarcSummaryListParamsSpf = "NONE"
	RadarEmailSecurityDmarcSummaryListParamsSpfFail RadarEmailSecurityDmarcSummaryListParamsSpf = "FAIL"
)
