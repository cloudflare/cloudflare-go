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

// RadarEmailSecuritySummaryMaliciousService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecuritySummaryMaliciousService] method instead.
type RadarEmailSecuritySummaryMaliciousService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySummaryMaliciousService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecuritySummaryMaliciousService(opts ...option.RequestOption) (r *RadarEmailSecuritySummaryMaliciousService) {
	r = &RadarEmailSecuritySummaryMaliciousService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified as MALICIOUS.
func (r *RadarEmailSecuritySummaryMaliciousService) List(ctx context.Context, query RadarEmailSecuritySummaryMaliciousListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryMaliciousListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryMaliciousListResponseEnvelope
	path := "radar/email/security/summary/malicious"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecuritySummaryMaliciousListResponse struct {
	Meta     RadarEmailSecuritySummaryMaliciousListResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryMaliciousListResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryMaliciousListResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryMaliciousListResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryMaliciousListResponse]
type radarEmailSecuritySummaryMaliciousListResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryMaliciousListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryMaliciousListResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryMaliciousListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                           `json:"lastUpdated,required"`
	Normalization  string                                                           `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryMaliciousListResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryMaliciousListResponseMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryMaliciousListResponseMeta]
type radarEmailSecuritySummaryMaliciousListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryMaliciousListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryMaliciousListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryMaliciousListResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryMaliciousListResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryMaliciousListResponseMetaDateRange]
type radarEmailSecuritySummaryMaliciousListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryMaliciousListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                        `json:"level"`
	JSON        radarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                         `json:"dataSource,required"`
	Description     string                                                                         `json:"description,required"`
	EventType       string                                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                                      `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryMaliciousListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryMaliciousListResponseSummary0 struct {
	Malicious    string                                                     `json:"MALICIOUS,required"`
	NotMalicious string                                                     `json:"NOT_MALICIOUS,required"`
	JSON         radarEmailSecuritySummaryMaliciousListResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryMaliciousListResponseSummary0JSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryMaliciousListResponseSummary0]
type radarEmailSecuritySummaryMaliciousListResponseSummary0JSON struct {
	Malicious    apijson.Field
	NotMalicious apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryMaliciousListResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryMaliciousListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummaryMaliciousListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryMaliciousListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummaryMaliciousListParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySummaryMaliciousListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryMaliciousListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryMaliciousListParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummaryMaliciousListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecuritySummaryMaliciousListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryMaliciousListParamsArc string

const (
	RadarEmailSecuritySummaryMaliciousListParamsArcPass RadarEmailSecuritySummaryMaliciousListParamsArc = "PASS"
	RadarEmailSecuritySummaryMaliciousListParamsArcNone RadarEmailSecuritySummaryMaliciousListParamsArc = "NONE"
	RadarEmailSecuritySummaryMaliciousListParamsArcFail RadarEmailSecuritySummaryMaliciousListParamsArc = "FAIL"
)

type RadarEmailSecuritySummaryMaliciousListParamsDateRange string

const (
	RadarEmailSecuritySummaryMaliciousListParamsDateRange1d         RadarEmailSecuritySummaryMaliciousListParamsDateRange = "1d"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange2d         RadarEmailSecuritySummaryMaliciousListParamsDateRange = "2d"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange7d         RadarEmailSecuritySummaryMaliciousListParamsDateRange = "7d"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange14d        RadarEmailSecuritySummaryMaliciousListParamsDateRange = "14d"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange28d        RadarEmailSecuritySummaryMaliciousListParamsDateRange = "28d"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange12w        RadarEmailSecuritySummaryMaliciousListParamsDateRange = "12w"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange24w        RadarEmailSecuritySummaryMaliciousListParamsDateRange = "24w"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange52w        RadarEmailSecuritySummaryMaliciousListParamsDateRange = "52w"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange1dControl  RadarEmailSecuritySummaryMaliciousListParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange2dControl  RadarEmailSecuritySummaryMaliciousListParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange7dControl  RadarEmailSecuritySummaryMaliciousListParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange14dControl RadarEmailSecuritySummaryMaliciousListParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange28dControl RadarEmailSecuritySummaryMaliciousListParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange12wControl RadarEmailSecuritySummaryMaliciousListParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryMaliciousListParamsDateRange24wControl RadarEmailSecuritySummaryMaliciousListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryMaliciousListParamsDKIM string

const (
	RadarEmailSecuritySummaryMaliciousListParamsDKIMPass RadarEmailSecuritySummaryMaliciousListParamsDKIM = "PASS"
	RadarEmailSecuritySummaryMaliciousListParamsDKIMNone RadarEmailSecuritySummaryMaliciousListParamsDKIM = "NONE"
	RadarEmailSecuritySummaryMaliciousListParamsDKIMFail RadarEmailSecuritySummaryMaliciousListParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummaryMaliciousListParamsDmarc string

const (
	RadarEmailSecuritySummaryMaliciousListParamsDmarcPass RadarEmailSecuritySummaryMaliciousListParamsDmarc = "PASS"
	RadarEmailSecuritySummaryMaliciousListParamsDmarcNone RadarEmailSecuritySummaryMaliciousListParamsDmarc = "NONE"
	RadarEmailSecuritySummaryMaliciousListParamsDmarcFail RadarEmailSecuritySummaryMaliciousListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryMaliciousListParamsFormat string

const (
	RadarEmailSecuritySummaryMaliciousListParamsFormatJson RadarEmailSecuritySummaryMaliciousListParamsFormat = "JSON"
	RadarEmailSecuritySummaryMaliciousListParamsFormatCsv  RadarEmailSecuritySummaryMaliciousListParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryMaliciousListParamsSPF string

const (
	RadarEmailSecuritySummaryMaliciousListParamsSPFPass RadarEmailSecuritySummaryMaliciousListParamsSPF = "PASS"
	RadarEmailSecuritySummaryMaliciousListParamsSPFNone RadarEmailSecuritySummaryMaliciousListParamsSPF = "NONE"
	RadarEmailSecuritySummaryMaliciousListParamsSPFFail RadarEmailSecuritySummaryMaliciousListParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryMaliciousListResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryMaliciousListResponse             `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarEmailSecuritySummaryMaliciousListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryMaliciousListResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryMaliciousListResponseEnvelope]
type radarEmailSecuritySummaryMaliciousListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryMaliciousListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
