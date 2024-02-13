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

// RadarEmailSecuritySummaryDmarcService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecuritySummaryDmarcService] method instead.
type RadarEmailSecuritySummaryDmarcService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySummaryDmarcService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySummaryDmarcService(opts ...option.RequestOption) (r *RadarEmailSecuritySummaryDmarcService) {
	r = &RadarEmailSecuritySummaryDmarcService{}
	r.Options = opts
	return
}

// Percentage distribution of emails classified per DMARC validation.
func (r *RadarEmailSecuritySummaryDmarcService) List(ctx context.Context, query RadarEmailSecuritySummaryDmarcListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummaryDmarcListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummaryDmarcListResponseEnvelope
	path := "radar/email/security/summary/dmarc"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecuritySummaryDmarcListResponse struct {
	Meta     RadarEmailSecuritySummaryDmarcListResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummaryDmarcListResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummaryDmarcListResponseJSON     `json:"-"`
}

// radarEmailSecuritySummaryDmarcListResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummaryDmarcListResponse]
type radarEmailSecuritySummaryDmarcListResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDmarcListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDmarcListResponseMeta struct {
	DateRange      []RadarEmailSecuritySummaryDmarcListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	Normalization  string                                                       `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummaryDmarcListResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummaryDmarcListResponseMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummaryDmarcListResponseMeta]
type radarEmailSecuritySummaryDmarcListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDmarcListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDmarcListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummaryDmarcListResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummaryDmarcListResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummaryDmarcListResponseMetaDateRange]
type radarEmailSecuritySummaryDmarcListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDmarcListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfo]
type radarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummaryDmarcListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDmarcListResponseSummary0 struct {
	Fail string                                                 `json:"FAIL,required"`
	None string                                                 `json:"NONE,required"`
	Pass string                                                 `json:"PASS,required"`
	JSON radarEmailSecuritySummaryDmarcListResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummaryDmarcListResponseSummary0JSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryDmarcListResponseSummary0]
type radarEmailSecuritySummaryDmarcListResponseSummary0JSON struct {
	Fail        apijson.Field
	None        apijson.Field
	Pass        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDmarcListResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummaryDmarcListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummaryDmarcListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummaryDmarcListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummaryDmarcListParamsDKIM] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummaryDmarcListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummaryDmarcListParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummaryDmarcListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecuritySummaryDmarcListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummaryDmarcListParamsArc string

const (
	RadarEmailSecuritySummaryDmarcListParamsArcPass RadarEmailSecuritySummaryDmarcListParamsArc = "PASS"
	RadarEmailSecuritySummaryDmarcListParamsArcNone RadarEmailSecuritySummaryDmarcListParamsArc = "NONE"
	RadarEmailSecuritySummaryDmarcListParamsArcFail RadarEmailSecuritySummaryDmarcListParamsArc = "FAIL"
)

type RadarEmailSecuritySummaryDmarcListParamsDateRange string

const (
	RadarEmailSecuritySummaryDmarcListParamsDateRange1d         RadarEmailSecuritySummaryDmarcListParamsDateRange = "1d"
	RadarEmailSecuritySummaryDmarcListParamsDateRange2d         RadarEmailSecuritySummaryDmarcListParamsDateRange = "2d"
	RadarEmailSecuritySummaryDmarcListParamsDateRange7d         RadarEmailSecuritySummaryDmarcListParamsDateRange = "7d"
	RadarEmailSecuritySummaryDmarcListParamsDateRange14d        RadarEmailSecuritySummaryDmarcListParamsDateRange = "14d"
	RadarEmailSecuritySummaryDmarcListParamsDateRange28d        RadarEmailSecuritySummaryDmarcListParamsDateRange = "28d"
	RadarEmailSecuritySummaryDmarcListParamsDateRange12w        RadarEmailSecuritySummaryDmarcListParamsDateRange = "12w"
	RadarEmailSecuritySummaryDmarcListParamsDateRange24w        RadarEmailSecuritySummaryDmarcListParamsDateRange = "24w"
	RadarEmailSecuritySummaryDmarcListParamsDateRange52w        RadarEmailSecuritySummaryDmarcListParamsDateRange = "52w"
	RadarEmailSecuritySummaryDmarcListParamsDateRange1dControl  RadarEmailSecuritySummaryDmarcListParamsDateRange = "1dControl"
	RadarEmailSecuritySummaryDmarcListParamsDateRange2dControl  RadarEmailSecuritySummaryDmarcListParamsDateRange = "2dControl"
	RadarEmailSecuritySummaryDmarcListParamsDateRange7dControl  RadarEmailSecuritySummaryDmarcListParamsDateRange = "7dControl"
	RadarEmailSecuritySummaryDmarcListParamsDateRange14dControl RadarEmailSecuritySummaryDmarcListParamsDateRange = "14dControl"
	RadarEmailSecuritySummaryDmarcListParamsDateRange28dControl RadarEmailSecuritySummaryDmarcListParamsDateRange = "28dControl"
	RadarEmailSecuritySummaryDmarcListParamsDateRange12wControl RadarEmailSecuritySummaryDmarcListParamsDateRange = "12wControl"
	RadarEmailSecuritySummaryDmarcListParamsDateRange24wControl RadarEmailSecuritySummaryDmarcListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummaryDmarcListParamsDKIM string

const (
	RadarEmailSecuritySummaryDmarcListParamsDKIMPass RadarEmailSecuritySummaryDmarcListParamsDKIM = "PASS"
	RadarEmailSecuritySummaryDmarcListParamsDKIMNone RadarEmailSecuritySummaryDmarcListParamsDKIM = "NONE"
	RadarEmailSecuritySummaryDmarcListParamsDKIMFail RadarEmailSecuritySummaryDmarcListParamsDKIM = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummaryDmarcListParamsFormat string

const (
	RadarEmailSecuritySummaryDmarcListParamsFormatJson RadarEmailSecuritySummaryDmarcListParamsFormat = "JSON"
	RadarEmailSecuritySummaryDmarcListParamsFormatCsv  RadarEmailSecuritySummaryDmarcListParamsFormat = "CSV"
)

type RadarEmailSecuritySummaryDmarcListParamsSPF string

const (
	RadarEmailSecuritySummaryDmarcListParamsSPFPass RadarEmailSecuritySummaryDmarcListParamsSPF = "PASS"
	RadarEmailSecuritySummaryDmarcListParamsSPFNone RadarEmailSecuritySummaryDmarcListParamsSPF = "NONE"
	RadarEmailSecuritySummaryDmarcListParamsSPFFail RadarEmailSecuritySummaryDmarcListParamsSPF = "FAIL"
)

type RadarEmailSecuritySummaryDmarcListResponseEnvelope struct {
	Result  RadarEmailSecuritySummaryDmarcListResponse             `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarEmailSecuritySummaryDmarcListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummaryDmarcListResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummaryDmarcListResponseEnvelope]
type radarEmailSecuritySummaryDmarcListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummaryDmarcListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
