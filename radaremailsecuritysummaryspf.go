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

// RadarEmailSecuritySummarySPFService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecuritySummarySPFService] method instead.
type RadarEmailSecuritySummarySPFService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySummarySPFService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySummarySPFService(opts ...option.RequestOption) (r *RadarEmailSecuritySummarySPFService) {
	r = &RadarEmailSecuritySummarySPFService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per SPF validation.
func (r *RadarEmailSecuritySummarySPFService) List(ctx context.Context, query RadarEmailSecuritySummarySPFListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummarySPFListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummarySPFListResponseEnvelope
	path := "radar/email/security/summary/spf"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecuritySummarySPFListResponse struct {
	Meta     RadarEmailSecuritySummarySPFListResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummarySPFListResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummarySPFListResponseJSON     `json:"-"`
}

// radarEmailSecuritySummarySPFListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummarySPFListResponse]
type radarEmailSecuritySummarySPFListResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySPFListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySPFListResponseMeta struct {
	DateRange      []RadarEmailSecuritySummarySPFListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	Normalization  string                                                     `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummarySPFListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummarySPFListResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummarySPFListResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummarySPFListResponseMeta]
type radarEmailSecuritySummarySPFListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySPFListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySPFListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummarySPFListResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummarySPFListResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummarySPFListResponseMetaDateRange]
type radarEmailSecuritySummarySPFListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySPFListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySPFListResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummarySPFListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarEmailSecuritySummarySPFListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummarySPFListResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummarySPFListResponseMetaConfidenceInfo]
type radarEmailSecuritySummarySPFListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySPFListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySPFListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummarySPFListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummarySPFListResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummarySPFListResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummarySPFListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummarySPFListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySPFListResponseSummary0 struct {
	Fail string                                               `json:"FAIL,required"`
	None string                                               `json:"NONE,required"`
	Pass string                                               `json:"PASS,required"`
	JSON radarEmailSecuritySummarySPFListResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummarySPFListResponseSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummarySPFListResponseSummary0]
type radarEmailSecuritySummarySPFListResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySPFListResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySPFListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummarySPFListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummarySPFListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummarySPFListParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySummarySPFListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummarySPFListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailSecuritySummarySPFListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySummarySPFListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummarySPFListParamsArc string

const (
	RadarEmailSecuritySummarySPFListParamsArcPass RadarEmailSecuritySummarySPFListParamsArc = "PASS"
	RadarEmailSecuritySummarySPFListParamsArcNone RadarEmailSecuritySummarySPFListParamsArc = "NONE"
	RadarEmailSecuritySummarySPFListParamsArcFail RadarEmailSecuritySummarySPFListParamsArc = "FAIL"
)

type RadarEmailSecuritySummarySPFListParamsDateRange string

const (
	RadarEmailSecuritySummarySPFListParamsDateRange1d         RadarEmailSecuritySummarySPFListParamsDateRange = "1d"
	RadarEmailSecuritySummarySPFListParamsDateRange2d         RadarEmailSecuritySummarySPFListParamsDateRange = "2d"
	RadarEmailSecuritySummarySPFListParamsDateRange7d         RadarEmailSecuritySummarySPFListParamsDateRange = "7d"
	RadarEmailSecuritySummarySPFListParamsDateRange14d        RadarEmailSecuritySummarySPFListParamsDateRange = "14d"
	RadarEmailSecuritySummarySPFListParamsDateRange28d        RadarEmailSecuritySummarySPFListParamsDateRange = "28d"
	RadarEmailSecuritySummarySPFListParamsDateRange12w        RadarEmailSecuritySummarySPFListParamsDateRange = "12w"
	RadarEmailSecuritySummarySPFListParamsDateRange24w        RadarEmailSecuritySummarySPFListParamsDateRange = "24w"
	RadarEmailSecuritySummarySPFListParamsDateRange52w        RadarEmailSecuritySummarySPFListParamsDateRange = "52w"
	RadarEmailSecuritySummarySPFListParamsDateRange1dControl  RadarEmailSecuritySummarySPFListParamsDateRange = "1dControl"
	RadarEmailSecuritySummarySPFListParamsDateRange2dControl  RadarEmailSecuritySummarySPFListParamsDateRange = "2dControl"
	RadarEmailSecuritySummarySPFListParamsDateRange7dControl  RadarEmailSecuritySummarySPFListParamsDateRange = "7dControl"
	RadarEmailSecuritySummarySPFListParamsDateRange14dControl RadarEmailSecuritySummarySPFListParamsDateRange = "14dControl"
	RadarEmailSecuritySummarySPFListParamsDateRange28dControl RadarEmailSecuritySummarySPFListParamsDateRange = "28dControl"
	RadarEmailSecuritySummarySPFListParamsDateRange12wControl RadarEmailSecuritySummarySPFListParamsDateRange = "12wControl"
	RadarEmailSecuritySummarySPFListParamsDateRange24wControl RadarEmailSecuritySummarySPFListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummarySPFListParamsDKIM string

const (
	RadarEmailSecuritySummarySPFListParamsDKIMPass RadarEmailSecuritySummarySPFListParamsDKIM = "PASS"
	RadarEmailSecuritySummarySPFListParamsDKIMNone RadarEmailSecuritySummarySPFListParamsDKIM = "NONE"
	RadarEmailSecuritySummarySPFListParamsDKIMFail RadarEmailSecuritySummarySPFListParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummarySPFListParamsDmarc string

const (
	RadarEmailSecuritySummarySPFListParamsDmarcPass RadarEmailSecuritySummarySPFListParamsDmarc = "PASS"
	RadarEmailSecuritySummarySPFListParamsDmarcNone RadarEmailSecuritySummarySPFListParamsDmarc = "NONE"
	RadarEmailSecuritySummarySPFListParamsDmarcFail RadarEmailSecuritySummarySPFListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummarySPFListParamsFormat string

const (
	RadarEmailSecuritySummarySPFListParamsFormatJson RadarEmailSecuritySummarySPFListParamsFormat = "JSON"
	RadarEmailSecuritySummarySPFListParamsFormatCsv  RadarEmailSecuritySummarySPFListParamsFormat = "CSV"
)

type RadarEmailSecuritySummarySPFListResponseEnvelope struct {
	Result  RadarEmailSecuritySummarySPFListResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailSecuritySummarySPFListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummarySPFListResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummarySPFListResponseEnvelope]
type radarEmailSecuritySummarySPFListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySPFListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
