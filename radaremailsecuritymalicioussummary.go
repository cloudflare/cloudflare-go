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

// RadarEmailSecurityMaliciousSummaryService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityMaliciousSummaryService] method instead.
type RadarEmailSecurityMaliciousSummaryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityMaliciousSummaryService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityMaliciousSummaryService(opts ...option.RequestOption) (r *RadarEmailSecurityMaliciousSummaryService) {
	r = &RadarEmailSecurityMaliciousSummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified as MALICIOUS.
func (r *RadarEmailSecurityMaliciousSummaryService) List(ctx context.Context, query RadarEmailSecurityMaliciousSummaryListParams, opts ...option.RequestOption) (res *RadarEmailSecurityMaliciousSummaryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityMaliciousSummaryListResponse struct {
	Result  RadarEmailSecurityMaliciousSummaryListResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailSecurityMaliciousSummaryListResponseJSON   `json:"-"`
}

// radarEmailSecurityMaliciousSummaryListResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityMaliciousSummaryListResponse]
type radarEmailSecurityMaliciousSummaryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityMaliciousSummaryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityMaliciousSummaryListResponseResult struct {
	Meta     RadarEmailSecurityMaliciousSummaryListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecurityMaliciousSummaryListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecurityMaliciousSummaryListResponseResultJSON     `json:"-"`
}

// radarEmailSecurityMaliciousSummaryListResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecurityMaliciousSummaryListResponseResult]
type radarEmailSecurityMaliciousSummaryListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityMaliciousSummaryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityMaliciousSummaryListResponseResultMeta struct {
	DateRange      []RadarEmailSecurityMaliciousSummaryListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                 `json:"lastUpdated,required"`
	Normalization  string                                                                 `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityMaliciousSummaryListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityMaliciousSummaryListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityMaliciousSummaryListResponseResultMeta]
type radarEmailSecurityMaliciousSummaryListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityMaliciousSummaryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityMaliciousSummaryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityMaliciousSummaryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityMaliciousSummaryListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityMaliciousSummaryListResponseResultMetaDateRange]
type radarEmailSecurityMaliciousSummaryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityMaliciousSummaryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                              `json:"level"`
	JSON        radarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfo]
type radarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                               `json:"dataSource,required"`
	Description     string                                                                               `json:"description,required"`
	EventType       string                                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                                            `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityMaliciousSummaryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityMaliciousSummaryListResponseResultSummary0 struct {
	Malicious    string                                                           `json:"MALICIOUS,required"`
	NotMalicious string                                                           `json:"NOT_MALICIOUS,required"`
	JSON         radarEmailSecurityMaliciousSummaryListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecurityMaliciousSummaryListResponseResultSummary0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityMaliciousSummaryListResponseResultSummary0]
type radarEmailSecurityMaliciousSummaryListResponseResultSummary0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityMaliciousSummaryListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityMaliciousSummaryListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityMaliciousSummaryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityMaliciousSummaryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityMaliciousSummaryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityMaliciousSummaryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityMaliciousSummaryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityMaliciousSummaryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityMaliciousSummaryListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityMaliciousSummaryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecurityMaliciousSummaryListParamsArc string

const (
	RadarEmailSecurityMaliciousSummaryListParamsArcPass RadarEmailSecurityMaliciousSummaryListParamsArc = "PASS"
	RadarEmailSecurityMaliciousSummaryListParamsArcNone RadarEmailSecurityMaliciousSummaryListParamsArc = "NONE"
	RadarEmailSecurityMaliciousSummaryListParamsArcFail RadarEmailSecurityMaliciousSummaryListParamsArc = "FAIL"
)

type RadarEmailSecurityMaliciousSummaryListParamsDateRange string

const (
	RadarEmailSecurityMaliciousSummaryListParamsDateRange1d         RadarEmailSecurityMaliciousSummaryListParamsDateRange = "1d"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange2d         RadarEmailSecurityMaliciousSummaryListParamsDateRange = "2d"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange7d         RadarEmailSecurityMaliciousSummaryListParamsDateRange = "7d"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange14d        RadarEmailSecurityMaliciousSummaryListParamsDateRange = "14d"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange28d        RadarEmailSecurityMaliciousSummaryListParamsDateRange = "28d"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange12w        RadarEmailSecurityMaliciousSummaryListParamsDateRange = "12w"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange24w        RadarEmailSecurityMaliciousSummaryListParamsDateRange = "24w"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange52w        RadarEmailSecurityMaliciousSummaryListParamsDateRange = "52w"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange1dControl  RadarEmailSecurityMaliciousSummaryListParamsDateRange = "1dControl"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange2dControl  RadarEmailSecurityMaliciousSummaryListParamsDateRange = "2dControl"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange7dControl  RadarEmailSecurityMaliciousSummaryListParamsDateRange = "7dControl"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange14dControl RadarEmailSecurityMaliciousSummaryListParamsDateRange = "14dControl"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange28dControl RadarEmailSecurityMaliciousSummaryListParamsDateRange = "28dControl"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange12wControl RadarEmailSecurityMaliciousSummaryListParamsDateRange = "12wControl"
	RadarEmailSecurityMaliciousSummaryListParamsDateRange24wControl RadarEmailSecurityMaliciousSummaryListParamsDateRange = "24wControl"
)

type RadarEmailSecurityMaliciousSummaryListParamsDkim string

const (
	RadarEmailSecurityMaliciousSummaryListParamsDkimPass RadarEmailSecurityMaliciousSummaryListParamsDkim = "PASS"
	RadarEmailSecurityMaliciousSummaryListParamsDkimNone RadarEmailSecurityMaliciousSummaryListParamsDkim = "NONE"
	RadarEmailSecurityMaliciousSummaryListParamsDkimFail RadarEmailSecurityMaliciousSummaryListParamsDkim = "FAIL"
)

type RadarEmailSecurityMaliciousSummaryListParamsDmarc string

const (
	RadarEmailSecurityMaliciousSummaryListParamsDmarcPass RadarEmailSecurityMaliciousSummaryListParamsDmarc = "PASS"
	RadarEmailSecurityMaliciousSummaryListParamsDmarcNone RadarEmailSecurityMaliciousSummaryListParamsDmarc = "NONE"
	RadarEmailSecurityMaliciousSummaryListParamsDmarcFail RadarEmailSecurityMaliciousSummaryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityMaliciousSummaryListParamsFormat string

const (
	RadarEmailSecurityMaliciousSummaryListParamsFormatJson RadarEmailSecurityMaliciousSummaryListParamsFormat = "JSON"
	RadarEmailSecurityMaliciousSummaryListParamsFormatCsv  RadarEmailSecurityMaliciousSummaryListParamsFormat = "CSV"
)

type RadarEmailSecurityMaliciousSummaryListParamsSpf string

const (
	RadarEmailSecurityMaliciousSummaryListParamsSpfPass RadarEmailSecurityMaliciousSummaryListParamsSpf = "PASS"
	RadarEmailSecurityMaliciousSummaryListParamsSpfNone RadarEmailSecurityMaliciousSummaryListParamsSpf = "NONE"
	RadarEmailSecurityMaliciousSummaryListParamsSpfFail RadarEmailSecurityMaliciousSummaryListParamsSpf = "FAIL"
)
