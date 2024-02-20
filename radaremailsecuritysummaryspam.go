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

// RadarEmailSecuritySummarySpamService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecuritySummarySpamService] method instead.
type RadarEmailSecuritySummarySpamService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecuritySummarySpamService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecuritySummarySpamService(opts ...option.RequestOption) (r *RadarEmailSecuritySummarySpamService) {
	r = &RadarEmailSecuritySummarySpamService{}
	r.Options = opts
	return
}

// Proportion of emails categorized as either spam or legitimate (non-spam).
func (r *RadarEmailSecuritySummarySpamService) List(ctx context.Context, query RadarEmailSecuritySummarySpamListParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummarySpamListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummarySpamListResponseEnvelope
	path := "radar/email/security/summary/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecuritySummarySpamListResponse struct {
	Meta     RadarEmailSecuritySummarySpamListResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummarySpamListResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummarySpamListResponseJSON     `json:"-"`
}

// radarEmailSecuritySummarySpamListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummarySpamListResponse]
type radarEmailSecuritySummarySpamListResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySpamListResponseMeta struct {
	DateRange      []RadarEmailSecuritySummarySpamListResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                      `json:"lastUpdated,required"`
	Normalization  string                                                      `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummarySpamListResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummarySpamListResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummarySpamListResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummarySpamListResponseMeta]
type radarEmailSecuritySummarySpamListResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamListResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySpamListResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummarySpamListResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummarySpamListResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummarySpamListResponseMetaDateRange]
type radarEmailSecuritySummarySpamListResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamListResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySpamListResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummarySpamListResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                   `json:"level"`
	JSON        radarEmailSecuritySummarySpamListResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummarySpamListResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecuritySummarySpamListResponseMetaConfidenceInfo]
type radarEmailSecuritySummarySpamListResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamListResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySpamListResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                    `json:"dataSource,required"`
	Description     string                                                                    `json:"description,required"`
	EventType       string                                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                                 `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummarySpamListResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummarySpamListResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummarySpamListResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummarySpamListResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummarySpamListResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySpamListResponseSummary0 struct {
	NotSpam string                                                `json:"NOT_SPAM,required"`
	Spam    string                                                `json:"SPAM,required"`
	JSON    radarEmailSecuritySummarySpamListResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummarySpamListResponseSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummarySpamListResponseSummary0]
type radarEmailSecuritySummarySpamListResponseSummary0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamListResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySpamListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummarySpamListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummarySpamListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummarySpamListParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySummarySpamListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummarySpamListParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummarySpamListParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummarySpamListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySummarySpamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummarySpamListParamsArc string

const (
	RadarEmailSecuritySummarySpamListParamsArcPass RadarEmailSecuritySummarySpamListParamsArc = "PASS"
	RadarEmailSecuritySummarySpamListParamsArcNone RadarEmailSecuritySummarySpamListParamsArc = "NONE"
	RadarEmailSecuritySummarySpamListParamsArcFail RadarEmailSecuritySummarySpamListParamsArc = "FAIL"
)

type RadarEmailSecuritySummarySpamListParamsDateRange string

const (
	RadarEmailSecuritySummarySpamListParamsDateRange1d         RadarEmailSecuritySummarySpamListParamsDateRange = "1d"
	RadarEmailSecuritySummarySpamListParamsDateRange2d         RadarEmailSecuritySummarySpamListParamsDateRange = "2d"
	RadarEmailSecuritySummarySpamListParamsDateRange7d         RadarEmailSecuritySummarySpamListParamsDateRange = "7d"
	RadarEmailSecuritySummarySpamListParamsDateRange14d        RadarEmailSecuritySummarySpamListParamsDateRange = "14d"
	RadarEmailSecuritySummarySpamListParamsDateRange28d        RadarEmailSecuritySummarySpamListParamsDateRange = "28d"
	RadarEmailSecuritySummarySpamListParamsDateRange12w        RadarEmailSecuritySummarySpamListParamsDateRange = "12w"
	RadarEmailSecuritySummarySpamListParamsDateRange24w        RadarEmailSecuritySummarySpamListParamsDateRange = "24w"
	RadarEmailSecuritySummarySpamListParamsDateRange52w        RadarEmailSecuritySummarySpamListParamsDateRange = "52w"
	RadarEmailSecuritySummarySpamListParamsDateRange1dControl  RadarEmailSecuritySummarySpamListParamsDateRange = "1dControl"
	RadarEmailSecuritySummarySpamListParamsDateRange2dControl  RadarEmailSecuritySummarySpamListParamsDateRange = "2dControl"
	RadarEmailSecuritySummarySpamListParamsDateRange7dControl  RadarEmailSecuritySummarySpamListParamsDateRange = "7dControl"
	RadarEmailSecuritySummarySpamListParamsDateRange14dControl RadarEmailSecuritySummarySpamListParamsDateRange = "14dControl"
	RadarEmailSecuritySummarySpamListParamsDateRange28dControl RadarEmailSecuritySummarySpamListParamsDateRange = "28dControl"
	RadarEmailSecuritySummarySpamListParamsDateRange12wControl RadarEmailSecuritySummarySpamListParamsDateRange = "12wControl"
	RadarEmailSecuritySummarySpamListParamsDateRange24wControl RadarEmailSecuritySummarySpamListParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummarySpamListParamsDKIM string

const (
	RadarEmailSecuritySummarySpamListParamsDKIMPass RadarEmailSecuritySummarySpamListParamsDKIM = "PASS"
	RadarEmailSecuritySummarySpamListParamsDKIMNone RadarEmailSecuritySummarySpamListParamsDKIM = "NONE"
	RadarEmailSecuritySummarySpamListParamsDKIMFail RadarEmailSecuritySummarySpamListParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummarySpamListParamsDmarc string

const (
	RadarEmailSecuritySummarySpamListParamsDmarcPass RadarEmailSecuritySummarySpamListParamsDmarc = "PASS"
	RadarEmailSecuritySummarySpamListParamsDmarcNone RadarEmailSecuritySummarySpamListParamsDmarc = "NONE"
	RadarEmailSecuritySummarySpamListParamsDmarcFail RadarEmailSecuritySummarySpamListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummarySpamListParamsFormat string

const (
	RadarEmailSecuritySummarySpamListParamsFormatJson RadarEmailSecuritySummarySpamListParamsFormat = "JSON"
	RadarEmailSecuritySummarySpamListParamsFormatCsv  RadarEmailSecuritySummarySpamListParamsFormat = "CSV"
)

type RadarEmailSecuritySummarySpamListParamsSPF string

const (
	RadarEmailSecuritySummarySpamListParamsSPFPass RadarEmailSecuritySummarySpamListParamsSPF = "PASS"
	RadarEmailSecuritySummarySpamListParamsSPFNone RadarEmailSecuritySummarySpamListParamsSPF = "NONE"
	RadarEmailSecuritySummarySpamListParamsSPFFail RadarEmailSecuritySummarySpamListParamsSPF = "FAIL"
)

type RadarEmailSecuritySummarySpamListResponseEnvelope struct {
	Result  RadarEmailSecuritySummarySpamListResponse             `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarEmailSecuritySummarySpamListResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummarySpamListResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummarySpamListResponseEnvelope]
type radarEmailSecuritySummarySpamListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
