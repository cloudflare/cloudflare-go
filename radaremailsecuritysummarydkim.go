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

// RadarEmailSecuritySummaryDKIMService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecuritySummaryDKIMService] method instead.
type RadarEmailSecuritySummaryDKIMService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySummaryDKIMService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySummaryDKIMService(opts ...option.RequestOption) (r *RadarEmailSecuritySummaryDKIMService) {
	r = &RadarEmailSecuritySummaryDKIMService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per DKIM validation.
func (r *RadarEmailSecuritySummaryDKIMService) List(ctx context.Context, query RadarEmailSecuritySummaryDKIMListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryDKIMListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryDKIMListResponseEnvelope
	path := "radar/email/security/summary/dkim"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecuritySummaryDKIMListResponse struct {
	Meta     RadarEmailSecuritySummaryDKIMListResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryDKIMListResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryDKIMListResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryDKIMListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummaryDKIMListResponse]
type radarEmailSecuritySummaryDKIMListResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDKIMListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDKIMListResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryDKIMListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	Normalization  string                                                      `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryDKIMListResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryDKIMListResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryDKIMListResponseMeta]
type radarEmailSecuritySummaryDKIMListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDKIMListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDKIMListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryDKIMListResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryDKIMListResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryDKIMListResponseMetaDateRange]
type radarEmailSecuritySummaryDKIMListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDKIMListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryDKIMListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDKIMListResponseSummary0 struct {
	Fail string                                                `json:"FAIL,required"`
	None string                                                `json:"NONE,required"`
	Pass string                                                `json:"PASS,required"`
	JSON radarEmailSecuritySummaryDKIMListResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryDKIMListResponseSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryDKIMListResponseSummary0]
type radarEmailSecuritySummaryDKIMListResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDKIMListResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDKIMListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummaryDKIMListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryDKIMListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySummaryDKIMListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryDKIMListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryDKIMListParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummaryDKIMListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySummaryDKIMListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryDKIMListParamsArc string

const (
	RadarEmailSecuritySummaryDKIMListParamsArcPass RadarEmailSecuritySummaryDKIMListParamsArc = "PASS"
	RadarEmailSecuritySummaryDKIMListParamsArcNone RadarEmailSecuritySummaryDKIMListParamsArc = "NONE"
	RadarEmailSecuritySummaryDKIMListParamsArcFail RadarEmailSecuritySummaryDKIMListParamsArc = "FAIL"
)

type RadarEmailSecuritySummaryDKIMListParamsDateRange string

const (
	RadarEmailSecuritySummaryDKIMListParamsDateRange1d         RadarEmailSecuritySummaryDKIMListParamsDateRange = "1d"
	RadarEmailSecuritySummaryDKIMListParamsDateRange2d         RadarEmailSecuritySummaryDKIMListParamsDateRange = "2d"
	RadarEmailSecuritySummaryDKIMListParamsDateRange7d         RadarEmailSecuritySummaryDKIMListParamsDateRange = "7d"
	RadarEmailSecuritySummaryDKIMListParamsDateRange14d        RadarEmailSecuritySummaryDKIMListParamsDateRange = "14d"
	RadarEmailSecuritySummaryDKIMListParamsDateRange28d        RadarEmailSecuritySummaryDKIMListParamsDateRange = "28d"
	RadarEmailSecuritySummaryDKIMListParamsDateRange12w        RadarEmailSecuritySummaryDKIMListParamsDateRange = "12w"
	RadarEmailSecuritySummaryDKIMListParamsDateRange24w        RadarEmailSecuritySummaryDKIMListParamsDateRange = "24w"
	RadarEmailSecuritySummaryDKIMListParamsDateRange52w        RadarEmailSecuritySummaryDKIMListParamsDateRange = "52w"
	RadarEmailSecuritySummaryDKIMListParamsDateRange1dControl  RadarEmailSecuritySummaryDKIMListParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryDKIMListParamsDateRange2dControl  RadarEmailSecuritySummaryDKIMListParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryDKIMListParamsDateRange7dControl  RadarEmailSecuritySummaryDKIMListParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryDKIMListParamsDateRange14dControl RadarEmailSecuritySummaryDKIMListParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryDKIMListParamsDateRange28dControl RadarEmailSecuritySummaryDKIMListParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryDKIMListParamsDateRange12wControl RadarEmailSecuritySummaryDKIMListParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryDKIMListParamsDateRange24wControl RadarEmailSecuritySummaryDKIMListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryDKIMListParamsDmarc string

const (
	RadarEmailSecuritySummaryDKIMListParamsDmarcPass RadarEmailSecuritySummaryDKIMListParamsDmarc = "PASS"
	RadarEmailSecuritySummaryDKIMListParamsDmarcNone RadarEmailSecuritySummaryDKIMListParamsDmarc = "NONE"
	RadarEmailSecuritySummaryDKIMListParamsDmarcFail RadarEmailSecuritySummaryDKIMListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryDKIMListParamsFormat string

const (
	RadarEmailSecuritySummaryDKIMListParamsFormatJson RadarEmailSecuritySummaryDKIMListParamsFormat = "JSON"
	RadarEmailSecuritySummaryDKIMListParamsFormatCsv  RadarEmailSecuritySummaryDKIMListParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryDKIMListParamsSPF string

const (
	RadarEmailSecuritySummaryDKIMListParamsSPFPass RadarEmailSecuritySummaryDKIMListParamsSPF = "PASS"
	RadarEmailSecuritySummaryDKIMListParamsSPFNone RadarEmailSecuritySummaryDKIMListParamsSPF = "NONE"
	RadarEmailSecuritySummaryDKIMListParamsSPFFail RadarEmailSecuritySummaryDKIMListParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryDKIMListResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryDKIMListResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarEmailSecuritySummaryDKIMListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryDKIMListResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryDKIMListResponseEnvelope]
type radarEmailSecuritySummaryDKIMListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDKIMListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
