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

// RadarEmailSecuritySummaryThreatCategoryService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEmailSecuritySummaryThreatCategoryService] method instead.
type RadarEmailSecuritySummaryThreatCategoryService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySummaryThreatCategoryService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecuritySummaryThreatCategoryService(opts ...option.RequestOption) (r *RadarEmailSecuritySummaryThreatCategoryService) {
	r = &RadarEmailSecuritySummaryThreatCategoryService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified in Threat Categories.
func (r *RadarEmailSecuritySummaryThreatCategoryService) List(ctx context.Context, query RadarEmailSecuritySummaryThreatCategoryListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryThreatCategoryListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryThreatCategoryListResponseEnvelope
	path := "radar/email/security/summary/threat_category"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecuritySummaryThreatCategoryListResponse struct {
	Meta     RadarEmailSecuritySummaryThreatCategoryListResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryThreatCategoryListResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryThreatCategoryListResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryListResponseJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryThreatCategoryListResponse]
type radarEmailSecuritySummaryThreatCategoryListResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryThreatCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryThreatCategoryListResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryThreatCategoryListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                `json:"lastUpdated,required"`
	Normalization  string                                                                `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryThreatCategoryListResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryListResponseMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryThreatCategoryListResponseMeta]
type radarEmailSecuritySummaryThreatCategoryListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryThreatCategoryListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryThreatCategoryListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                            `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryThreatCategoryListResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryListResponseMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryThreatCategoryListResponseMetaDateRange]
type radarEmailSecuritySummaryThreatCategoryListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryThreatCategoryListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                             `json:"level"`
	JSON        radarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                              `json:"dataSource,required"`
	Description     string                                                                              `json:"description,required"`
	EventType       string                                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                                           `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryThreatCategoryListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryThreatCategoryListResponseSummary0 struct {
	BrandImpersonation  string                                                          `json:"BrandImpersonation,required"`
	CredentialHarvester string                                                          `json:"CredentialHarvester,required"`
	IdentityDeception   string                                                          `json:"IdentityDeception,required"`
	Link                string                                                          `json:"Link,required"`
	JSON                radarEmailSecuritySummaryThreatCategoryListResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryListResponseSummary0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryThreatCategoryListResponseSummary0]
type radarEmailSecuritySummaryThreatCategoryListResponseSummary0JSON struct {
	BrandImpersonation  apijson.Field
	CredentialHarvester apijson.Field
	IdentityDeception   apijson.Field
	Link                apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryThreatCategoryListResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryThreatCategoryListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummaryThreatCategoryListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryThreatCategoryListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummaryThreatCategoryListParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySummaryThreatCategoryListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryThreatCategoryListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryThreatCategoryListParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummaryThreatCategoryListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecuritySummaryThreatCategoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryThreatCategoryListParamsArc string

const (
	RadarEmailSecuritySummaryThreatCategoryListParamsArcPass RadarEmailSecuritySummaryThreatCategoryListParamsArc = "PASS"
	RadarEmailSecuritySummaryThreatCategoryListParamsArcNone RadarEmailSecuritySummaryThreatCategoryListParamsArc = "NONE"
	RadarEmailSecuritySummaryThreatCategoryListParamsArcFail RadarEmailSecuritySummaryThreatCategoryListParamsArc = "FAIL"
)

type RadarEmailSecuritySummaryThreatCategoryListParamsDateRange string

const (
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange1d         RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "1d"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange2d         RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "2d"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange7d         RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "7d"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange14d        RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "14d"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange28d        RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "28d"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange12w        RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "12w"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange24w        RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "24w"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange52w        RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "52w"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange1dControl  RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange2dControl  RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange7dControl  RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange14dControl RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange28dControl RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange12wControl RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryThreatCategoryListParamsDateRange24wControl RadarEmailSecuritySummaryThreatCategoryListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryThreatCategoryListParamsDKIM string

const (
	RadarEmailSecuritySummaryThreatCategoryListParamsDKIMPass RadarEmailSecuritySummaryThreatCategoryListParamsDKIM = "PASS"
	RadarEmailSecuritySummaryThreatCategoryListParamsDKIMNone RadarEmailSecuritySummaryThreatCategoryListParamsDKIM = "NONE"
	RadarEmailSecuritySummaryThreatCategoryListParamsDKIMFail RadarEmailSecuritySummaryThreatCategoryListParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummaryThreatCategoryListParamsDmarc string

const (
	RadarEmailSecuritySummaryThreatCategoryListParamsDmarcPass RadarEmailSecuritySummaryThreatCategoryListParamsDmarc = "PASS"
	RadarEmailSecuritySummaryThreatCategoryListParamsDmarcNone RadarEmailSecuritySummaryThreatCategoryListParamsDmarc = "NONE"
	RadarEmailSecuritySummaryThreatCategoryListParamsDmarcFail RadarEmailSecuritySummaryThreatCategoryListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryThreatCategoryListParamsFormat string

const (
	RadarEmailSecuritySummaryThreatCategoryListParamsFormatJson RadarEmailSecuritySummaryThreatCategoryListParamsFormat = "JSON"
	RadarEmailSecuritySummaryThreatCategoryListParamsFormatCsv  RadarEmailSecuritySummaryThreatCategoryListParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryThreatCategoryListParamsSPF string

const (
	RadarEmailSecuritySummaryThreatCategoryListParamsSPFPass RadarEmailSecuritySummaryThreatCategoryListParamsSPF = "PASS"
	RadarEmailSecuritySummaryThreatCategoryListParamsSPFNone RadarEmailSecuritySummaryThreatCategoryListParamsSPF = "NONE"
	RadarEmailSecuritySummaryThreatCategoryListParamsSPFFail RadarEmailSecuritySummaryThreatCategoryListParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryThreatCategoryListResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryThreatCategoryListResponse             `json:"result,required"`
	Success bool                                                            `json:"success,required"`
	JSON    radarEmailSecuritySummaryThreatCategoryListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryThreatCategoryListResponseEnvelopeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryThreatCategoryListResponseEnvelope]
type radarEmailSecuritySummaryThreatCategoryListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryThreatCategoryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
