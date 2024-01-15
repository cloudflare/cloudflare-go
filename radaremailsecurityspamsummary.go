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

// RadarEmailSecuritySpamSummaryService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecuritySpamSummaryService] method instead.
type RadarEmailSecuritySpamSummaryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySpamSummaryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySpamSummaryService(opts ...option.RequestOption) (r *RadarEmailSecuritySpamSummaryService) {
	r = &RadarEmailSecuritySpamSummaryService{}
	r.Options = opts
	return
}

// Proportion of emails categorized as either spam or legitimate (non-spam).
func (r *RadarEmailSecuritySpamSummaryService) List(ctx context.Context, query RadarEmailSecuritySpamSummaryListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySpamSummaryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecuritySpamSummaryListResponse struct {
	Result  RadarEmailSecuritySpamSummaryListResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailSecuritySpamSummaryListResponseJSON   `json:"-"`
}

// radarEmailSecuritySpamSummaryListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySpamSummaryListResponse]
type radarEmailSecuritySpamSummaryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpamSummaryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamSummaryListResponseResult struct {
	Meta     RadarEmailSecuritySpamSummaryListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySpamSummaryListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySpamSummaryListResponseResultJSON     `json:"-"`
}

// radarEmailSecuritySpamSummaryListResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySpamSummaryListResponseResult]
type radarEmailSecuritySpamSummaryListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpamSummaryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamSummaryListResponseResultMeta struct {
	DateRange      []RadarEmailSecuritySpamSummaryListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                            `json:"lastUpdated,required"`
	Normalization  string                                                            `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySpamSummaryListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecuritySpamSummaryListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySpamSummaryListResponseResultMeta]
type radarEmailSecuritySpamSummaryListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySpamSummaryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamSummaryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySpamSummaryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySpamSummaryListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySpamSummaryListResponseResultMetaDateRange]
type radarEmailSecuritySpamSummaryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpamSummaryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                         `json:"level"`
	JSON        radarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfo]
type radarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                          `json:"dataSource,required"`
	Description     string                                                                          `json:"description,required"`
	EventType       string                                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                       `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySpamSummaryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamSummaryListResponseResultSummary0 struct {
	NotSpam string                                                      `json:"NOT_SPAM,required"`
	Spam    string                                                      `json:"SPAM,required"`
	JSON    radarEmailSecuritySpamSummaryListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecuritySpamSummaryListResponseResultSummary0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySpamSummaryListResponseResultSummary0]
type radarEmailSecuritySpamSummaryListResponseResultSummary0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySpamSummaryListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySpamSummaryListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySpamSummaryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySpamSummaryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecuritySpamSummaryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySpamSummaryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySpamSummaryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecuritySpamSummaryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySpamSummaryListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySpamSummaryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySpamSummaryListParamsArc string

const (
	RadarEmailSecuritySpamSummaryListParamsArcPass RadarEmailSecuritySpamSummaryListParamsArc = "PASS"
	RadarEmailSecuritySpamSummaryListParamsArcNone RadarEmailSecuritySpamSummaryListParamsArc = "NONE"
	RadarEmailSecuritySpamSummaryListParamsArcFail RadarEmailSecuritySpamSummaryListParamsArc = "FAIL"
)

type RadarEmailSecuritySpamSummaryListParamsDateRange string

const (
	RadarEmailSecuritySpamSummaryListParamsDateRange1d         RadarEmailSecuritySpamSummaryListParamsDateRange = "1d"
	RadarEmailSecuritySpamSummaryListParamsDateRange2d         RadarEmailSecuritySpamSummaryListParamsDateRange = "2d"
	RadarEmailSecuritySpamSummaryListParamsDateRange7d         RadarEmailSecuritySpamSummaryListParamsDateRange = "7d"
	RadarEmailSecuritySpamSummaryListParamsDateRange14d        RadarEmailSecuritySpamSummaryListParamsDateRange = "14d"
	RadarEmailSecuritySpamSummaryListParamsDateRange28d        RadarEmailSecuritySpamSummaryListParamsDateRange = "28d"
	RadarEmailSecuritySpamSummaryListParamsDateRange12w        RadarEmailSecuritySpamSummaryListParamsDateRange = "12w"
	RadarEmailSecuritySpamSummaryListParamsDateRange24w        RadarEmailSecuritySpamSummaryListParamsDateRange = "24w"
	RadarEmailSecuritySpamSummaryListParamsDateRange52w        RadarEmailSecuritySpamSummaryListParamsDateRange = "52w"
	RadarEmailSecuritySpamSummaryListParamsDateRange1dControl  RadarEmailSecuritySpamSummaryListParamsDateRange = "1dControl"
	RadarEmailSecuritySpamSummaryListParamsDateRange2dControl  RadarEmailSecuritySpamSummaryListParamsDateRange = "2dControl"
	RadarEmailSecuritySpamSummaryListParamsDateRange7dControl  RadarEmailSecuritySpamSummaryListParamsDateRange = "7dControl"
	RadarEmailSecuritySpamSummaryListParamsDateRange14dControl RadarEmailSecuritySpamSummaryListParamsDateRange = "14dControl"
	RadarEmailSecuritySpamSummaryListParamsDateRange28dControl RadarEmailSecuritySpamSummaryListParamsDateRange = "28dControl"
	RadarEmailSecuritySpamSummaryListParamsDateRange12wControl RadarEmailSecuritySpamSummaryListParamsDateRange = "12wControl"
	RadarEmailSecuritySpamSummaryListParamsDateRange24wControl RadarEmailSecuritySpamSummaryListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySpamSummaryListParamsDkim string

const (
	RadarEmailSecuritySpamSummaryListParamsDkimPass RadarEmailSecuritySpamSummaryListParamsDkim = "PASS"
	RadarEmailSecuritySpamSummaryListParamsDkimNone RadarEmailSecuritySpamSummaryListParamsDkim = "NONE"
	RadarEmailSecuritySpamSummaryListParamsDkimFail RadarEmailSecuritySpamSummaryListParamsDkim = "FAIL"
)

type RadarEmailSecuritySpamSummaryListParamsDmarc string

const (
	RadarEmailSecuritySpamSummaryListParamsDmarcPass RadarEmailSecuritySpamSummaryListParamsDmarc = "PASS"
	RadarEmailSecuritySpamSummaryListParamsDmarcNone RadarEmailSecuritySpamSummaryListParamsDmarc = "NONE"
	RadarEmailSecuritySpamSummaryListParamsDmarcFail RadarEmailSecuritySpamSummaryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySpamSummaryListParamsFormat string

const (
	RadarEmailSecuritySpamSummaryListParamsFormatJson RadarEmailSecuritySpamSummaryListParamsFormat = "JSON"
	RadarEmailSecuritySpamSummaryListParamsFormatCsv  RadarEmailSecuritySpamSummaryListParamsFormat = "CSV"
)

type RadarEmailSecuritySpamSummaryListParamsSpf string

const (
	RadarEmailSecuritySpamSummaryListParamsSpfPass RadarEmailSecuritySpamSummaryListParamsSpf = "PASS"
	RadarEmailSecuritySpamSummaryListParamsSpfNone RadarEmailSecuritySpamSummaryListParamsSpf = "NONE"
	RadarEmailSecuritySpamSummaryListParamsSpfFail RadarEmailSecuritySpamSummaryListParamsSpf = "FAIL"
)
