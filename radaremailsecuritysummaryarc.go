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

// RadarEmailSecuritySummaryArcService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecuritySummaryArcService] method instead.
type RadarEmailSecuritySummaryArcService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySummaryArcService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySummaryArcService(opts ...option.RequestOption) (r *RadarEmailSecuritySummaryArcService) {
	r = &RadarEmailSecuritySummaryArcService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per ARC validation.
func (r *RadarEmailSecuritySummaryArcService) List(ctx context.Context, query RadarEmailSecuritySummaryArcListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryArcListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryArcListResponseEnvelope
	path := "radar/email/security/summary/arc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecuritySummaryArcListResponse struct {
	Meta     RadarEmailSecuritySummaryArcListResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryArcListResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryArcListResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryArcListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryArcListResponse]
type radarEmailSecuritySummaryArcListResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryArcListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryArcListResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryArcListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	Normalization  string                                                     `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryArcListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryArcListResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryArcListResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryArcListResponseMeta]
type radarEmailSecuritySummaryArcListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryArcListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryArcListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryArcListResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryArcListResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryArcListResponseMetaDateRange]
type radarEmailSecuritySummaryArcListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryArcListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryArcListResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryArcListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarEmailSecuritySummaryArcListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryArcListResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryArcListResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryArcListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryArcListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryArcListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryArcListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryArcListResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryArcListResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryArcListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryArcListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryArcListResponseSummary0 struct {
	Fail string                                               `json:"FAIL,required"`
	None string                                               `json:"NONE,required"`
	Pass string                                               `json:"PASS,required"`
	JSON radarEmailSecuritySummaryArcListResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryArcListResponseSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryArcListResponseSummary0]
type radarEmailSecuritySummaryArcListResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryArcListResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryArcListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryArcListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummaryArcListParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySummaryArcListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryArcListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryArcListParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummaryArcListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySummaryArcListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryArcListParamsDateRange string

const (
	RadarEmailSecuritySummaryArcListParamsDateRange1d         RadarEmailSecuritySummaryArcListParamsDateRange = "1d"
	RadarEmailSecuritySummaryArcListParamsDateRange2d         RadarEmailSecuritySummaryArcListParamsDateRange = "2d"
	RadarEmailSecuritySummaryArcListParamsDateRange7d         RadarEmailSecuritySummaryArcListParamsDateRange = "7d"
	RadarEmailSecuritySummaryArcListParamsDateRange14d        RadarEmailSecuritySummaryArcListParamsDateRange = "14d"
	RadarEmailSecuritySummaryArcListParamsDateRange28d        RadarEmailSecuritySummaryArcListParamsDateRange = "28d"
	RadarEmailSecuritySummaryArcListParamsDateRange12w        RadarEmailSecuritySummaryArcListParamsDateRange = "12w"
	RadarEmailSecuritySummaryArcListParamsDateRange24w        RadarEmailSecuritySummaryArcListParamsDateRange = "24w"
	RadarEmailSecuritySummaryArcListParamsDateRange52w        RadarEmailSecuritySummaryArcListParamsDateRange = "52w"
	RadarEmailSecuritySummaryArcListParamsDateRange1dControl  RadarEmailSecuritySummaryArcListParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryArcListParamsDateRange2dControl  RadarEmailSecuritySummaryArcListParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryArcListParamsDateRange7dControl  RadarEmailSecuritySummaryArcListParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryArcListParamsDateRange14dControl RadarEmailSecuritySummaryArcListParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryArcListParamsDateRange28dControl RadarEmailSecuritySummaryArcListParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryArcListParamsDateRange12wControl RadarEmailSecuritySummaryArcListParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryArcListParamsDateRange24wControl RadarEmailSecuritySummaryArcListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryArcListParamsDKIM string

const (
	RadarEmailSecuritySummaryArcListParamsDKIMPass RadarEmailSecuritySummaryArcListParamsDKIM = "PASS"
	RadarEmailSecuritySummaryArcListParamsDKIMNone RadarEmailSecuritySummaryArcListParamsDKIM = "NONE"
	RadarEmailSecuritySummaryArcListParamsDKIMFail RadarEmailSecuritySummaryArcListParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummaryArcListParamsDmarc string

const (
	RadarEmailSecuritySummaryArcListParamsDmarcPass RadarEmailSecuritySummaryArcListParamsDmarc = "PASS"
	RadarEmailSecuritySummaryArcListParamsDmarcNone RadarEmailSecuritySummaryArcListParamsDmarc = "NONE"
	RadarEmailSecuritySummaryArcListParamsDmarcFail RadarEmailSecuritySummaryArcListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryArcListParamsFormat string

const (
	RadarEmailSecuritySummaryArcListParamsFormatJson RadarEmailSecuritySummaryArcListParamsFormat = "JSON"
	RadarEmailSecuritySummaryArcListParamsFormatCsv  RadarEmailSecuritySummaryArcListParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryArcListParamsSPF string

const (
	RadarEmailSecuritySummaryArcListParamsSPFPass RadarEmailSecuritySummaryArcListParamsSPF = "PASS"
	RadarEmailSecuritySummaryArcListParamsSPFNone RadarEmailSecuritySummaryArcListParamsSPF = "NONE"
	RadarEmailSecuritySummaryArcListParamsSPFFail RadarEmailSecuritySummaryArcListParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryArcListResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryArcListResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailSecuritySummaryArcListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryArcListResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryArcListResponseEnvelope]
type radarEmailSecuritySummaryArcListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryArcListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
