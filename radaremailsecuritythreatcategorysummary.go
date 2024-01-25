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

// RadarEmailSecurityThreatCategorySummaryService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEmailSecurityThreatCategorySummaryService] method instead.
type RadarEmailSecurityThreatCategorySummaryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityThreatCategorySummaryService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityThreatCategorySummaryService(opts ...option.RequestOption) (r *RadarEmailSecurityThreatCategorySummaryService) {
	r = &RadarEmailSecurityThreatCategorySummaryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified in Threat Categories.
func (r *RadarEmailSecurityThreatCategorySummaryService) List(ctx context.Context, query RadarEmailSecurityThreatCategorySummaryListParams, opts ...option.RequestOption) (res *RadarEmailSecurityThreatCategorySummaryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/summary/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityThreatCategorySummaryListResponse struct {
	Result  RadarEmailSecurityThreatCategorySummaryListResponseResult `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarEmailSecurityThreatCategorySummaryListResponseJSON   `json:"-"`
}

// radarEmailSecurityThreatCategorySummaryListResponseJSON contains the JSON
// metadata for the struct [RadarEmailSecurityThreatCategorySummaryListResponse]
type radarEmailSecurityThreatCategorySummaryListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityThreatCategorySummaryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityThreatCategorySummaryListResponseResult struct {
	Meta     RadarEmailSecurityThreatCategorySummaryListResponseResultMeta     `json:"meta,required"`
	Summary0 RadarEmailSecurityThreatCategorySummaryListResponseResultSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecurityThreatCategorySummaryListResponseResultJSON     `json:"-"`
}

// radarEmailSecurityThreatCategorySummaryListResponseResultJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityThreatCategorySummaryListResponseResult]
type radarEmailSecurityThreatCategorySummaryListResponseResultJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityThreatCategorySummaryListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityThreatCategorySummaryListResponseResultMeta struct {
	DateRange      []RadarEmailSecurityThreatCategorySummaryListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                      `json:"lastUpdated,required"`
	Normalization  string                                                                      `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityThreatCategorySummaryListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityThreatCategorySummaryListResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityThreatCategorySummaryListResponseResultMeta]
type radarEmailSecurityThreatCategorySummaryListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityThreatCategorySummaryListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityThreatCategorySummaryListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityThreatCategorySummaryListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityThreatCategorySummaryListResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityThreatCategorySummaryListResponseResultMetaDateRange]
type radarEmailSecurityThreatCategorySummaryListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityThreatCategorySummaryListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                   `json:"level"`
	JSON        radarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfo]
type radarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                    `json:"dataSource,required"`
	Description     string                                                                                    `json:"description,required"`
	EventType       string                                                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                                 `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityThreatCategorySummaryListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityThreatCategorySummaryListResponseResultSummary0 struct {
	BrandImpersonation  string                                                                `json:"BrandImpersonation,required"`
	CredentialHarvester string                                                                `json:"CredentialHarvester,required"`
	IdentityDeception   string                                                                `json:"IdentityDeception,required"`
	Link                string                                                                `json:"Link,required"`
	JSON                radarEmailSecurityThreatCategorySummaryListResponseResultSummary0JSON `json:"-"`
}

// radarEmailSecurityThreatCategorySummaryListResponseResultSummary0JSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityThreatCategorySummaryListResponseResultSummary0]
type radarEmailSecurityThreatCategorySummaryListResponseResultSummary0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityThreatCategorySummaryListResponseResultSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityThreatCategorySummaryListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityThreatCategorySummaryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityThreatCategorySummaryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityThreatCategorySummaryListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityThreatCategorySummaryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityThreatCategorySummaryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityThreatCategorySummaryListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityThreatCategorySummaryListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityThreatCategorySummaryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecurityThreatCategorySummaryListParamsArc string

const (
	RadarEmailSecurityThreatCategorySummaryListParamsArcPass RadarEmailSecurityThreatCategorySummaryListParamsArc = "PASS"
	RadarEmailSecurityThreatCategorySummaryListParamsArcNone RadarEmailSecurityThreatCategorySummaryListParamsArc = "NONE"
	RadarEmailSecurityThreatCategorySummaryListParamsArcFail RadarEmailSecurityThreatCategorySummaryListParamsArc = "FAIL"
)

type RadarEmailSecurityThreatCategorySummaryListParamsDateRange string

const (
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange1d         RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "1d"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange2d         RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "2d"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange7d         RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "7d"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange14d        RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "14d"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange28d        RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "28d"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange12w        RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "12w"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange24w        RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "24w"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange52w        RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "52w"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange1dControl  RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "1dControl"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange2dControl  RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "2dControl"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange7dControl  RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "7dControl"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange14dControl RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "14dControl"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange28dControl RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "28dControl"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange12wControl RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "12wControl"
	RadarEmailSecurityThreatCategorySummaryListParamsDateRange24wControl RadarEmailSecurityThreatCategorySummaryListParamsDateRange = "24wControl"
)

type RadarEmailSecurityThreatCategorySummaryListParamsDkim string

const (
	RadarEmailSecurityThreatCategorySummaryListParamsDkimPass RadarEmailSecurityThreatCategorySummaryListParamsDkim = "PASS"
	RadarEmailSecurityThreatCategorySummaryListParamsDkimNone RadarEmailSecurityThreatCategorySummaryListParamsDkim = "NONE"
	RadarEmailSecurityThreatCategorySummaryListParamsDkimFail RadarEmailSecurityThreatCategorySummaryListParamsDkim = "FAIL"
)

type RadarEmailSecurityThreatCategorySummaryListParamsDmarc string

const (
	RadarEmailSecurityThreatCategorySummaryListParamsDmarcPass RadarEmailSecurityThreatCategorySummaryListParamsDmarc = "PASS"
	RadarEmailSecurityThreatCategorySummaryListParamsDmarcNone RadarEmailSecurityThreatCategorySummaryListParamsDmarc = "NONE"
	RadarEmailSecurityThreatCategorySummaryListParamsDmarcFail RadarEmailSecurityThreatCategorySummaryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityThreatCategorySummaryListParamsFormat string

const (
	RadarEmailSecurityThreatCategorySummaryListParamsFormatJson RadarEmailSecurityThreatCategorySummaryListParamsFormat = "JSON"
	RadarEmailSecurityThreatCategorySummaryListParamsFormatCsv  RadarEmailSecurityThreatCategorySummaryListParamsFormat = "CSV"
)

type RadarEmailSecurityThreatCategorySummaryListParamsSpf string

const (
	RadarEmailSecurityThreatCategorySummaryListParamsSpfPass RadarEmailSecurityThreatCategorySummaryListParamsSpf = "PASS"
	RadarEmailSecurityThreatCategorySummaryListParamsSpfNone RadarEmailSecurityThreatCategorySummaryListParamsSpf = "NONE"
	RadarEmailSecurityThreatCategorySummaryListParamsSpfFail RadarEmailSecurityThreatCategorySummaryListParamsSpf = "FAIL"
)
