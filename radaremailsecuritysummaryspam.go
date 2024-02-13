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
func (r *RadarEmailSecuritySummarySpamService) Get(ctx context.Context, query RadarEmailSecuritySummarySpamGetParams, opts ...option.RequestOption) (res *RadarEmailSecuritySummarySpamGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecuritySummarySpamGetResponseEnvelope
	path := "radar/email/security/summary/spam"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecuritySummarySpamGetResponse struct {
	Meta     RadarEmailSecuritySummarySpamGetResponseMeta     `json:"meta,required"`
	Summary0 RadarEmailSecuritySummarySpamGetResponseSummary0 `json:"summary_0,required"`
	JSON     radarEmailSecuritySummarySpamGetResponseJSON     `json:"-"`
}

// radarEmailSecuritySummarySpamGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecuritySummarySpamGetResponse]
type radarEmailSecuritySummarySpamGetResponseJSON struct {
	Meta        apijson.Field
	Summary0    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySpamGetResponseMeta struct {
	DateRange      []RadarEmailSecuritySummarySpamGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	Normalization  string                                                     `json:"normalization,required"`
	ConfidenceInfo RadarEmailSecuritySummarySpamGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecuritySummarySpamGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecuritySummarySpamGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecuritySummarySpamGetResponseMeta]
type radarEmailSecuritySummarySpamGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	Normalization  apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySpamGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecuritySummarySpamGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecuritySummarySpamGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecuritySummarySpamGetResponseMetaDateRange]
type radarEmailSecuritySummarySpamGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySpamGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecuritySummarySpamGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarEmailSecuritySummarySpamGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecuritySummarySpamGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecuritySummarySpamGetResponseMetaConfidenceInfo]
type radarEmailSecuritySummarySpamGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySpamGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarEmailSecuritySummarySpamGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecuritySummarySpamGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecuritySummarySpamGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecuritySummarySpamGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecuritySummarySpamGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySpamGetResponseSummary0 struct {
	NotSpam string                                               `json:"NOT_SPAM,required"`
	Spam    string                                               `json:"SPAM,required"`
	JSON    radarEmailSecuritySummarySpamGetResponseSummary0JSON `json:"-"`
}

// radarEmailSecuritySummarySpamGetResponseSummary0JSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummarySpamGetResponseSummary0]
type radarEmailSecuritySummarySpamGetResponseSummary0JSON struct {
	NotSpam     apijson.Field
	Spam        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamGetResponseSummary0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecuritySummarySpamGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecuritySummarySpamGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecuritySummarySpamGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecuritySummarySpamGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecuritySummarySpamGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecuritySummarySpamGetParamsFormat] `query:"format"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecuritySummarySpamGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecuritySummarySpamGetParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecuritySummarySpamGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecuritySummarySpamGetParamsArc string

const (
	RadarEmailSecuritySummarySpamGetParamsArcPass RadarEmailSecuritySummarySpamGetParamsArc = "PASS"
	RadarEmailSecuritySummarySpamGetParamsArcNone RadarEmailSecuritySummarySpamGetParamsArc = "NONE"
	RadarEmailSecuritySummarySpamGetParamsArcFail RadarEmailSecuritySummarySpamGetParamsArc = "FAIL"
)

type RadarEmailSecuritySummarySpamGetParamsDateRange string

const (
	RadarEmailSecuritySummarySpamGetParamsDateRange1d         RadarEmailSecuritySummarySpamGetParamsDateRange = "1d"
	RadarEmailSecuritySummarySpamGetParamsDateRange2d         RadarEmailSecuritySummarySpamGetParamsDateRange = "2d"
	RadarEmailSecuritySummarySpamGetParamsDateRange7d         RadarEmailSecuritySummarySpamGetParamsDateRange = "7d"
	RadarEmailSecuritySummarySpamGetParamsDateRange14d        RadarEmailSecuritySummarySpamGetParamsDateRange = "14d"
	RadarEmailSecuritySummarySpamGetParamsDateRange28d        RadarEmailSecuritySummarySpamGetParamsDateRange = "28d"
	RadarEmailSecuritySummarySpamGetParamsDateRange12w        RadarEmailSecuritySummarySpamGetParamsDateRange = "12w"
	RadarEmailSecuritySummarySpamGetParamsDateRange24w        RadarEmailSecuritySummarySpamGetParamsDateRange = "24w"
	RadarEmailSecuritySummarySpamGetParamsDateRange52w        RadarEmailSecuritySummarySpamGetParamsDateRange = "52w"
	RadarEmailSecuritySummarySpamGetParamsDateRange1dControl  RadarEmailSecuritySummarySpamGetParamsDateRange = "1dControl"
	RadarEmailSecuritySummarySpamGetParamsDateRange2dControl  RadarEmailSecuritySummarySpamGetParamsDateRange = "2dControl"
	RadarEmailSecuritySummarySpamGetParamsDateRange7dControl  RadarEmailSecuritySummarySpamGetParamsDateRange = "7dControl"
	RadarEmailSecuritySummarySpamGetParamsDateRange14dControl RadarEmailSecuritySummarySpamGetParamsDateRange = "14dControl"
	RadarEmailSecuritySummarySpamGetParamsDateRange28dControl RadarEmailSecuritySummarySpamGetParamsDateRange = "28dControl"
	RadarEmailSecuritySummarySpamGetParamsDateRange12wControl RadarEmailSecuritySummarySpamGetParamsDateRange = "12wControl"
	RadarEmailSecuritySummarySpamGetParamsDateRange24wControl RadarEmailSecuritySummarySpamGetParamsDateRange = "24wControl"
)

type RadarEmailSecuritySummarySpamGetParamsDKIM string

const (
	RadarEmailSecuritySummarySpamGetParamsDKIMPass RadarEmailSecuritySummarySpamGetParamsDKIM = "PASS"
	RadarEmailSecuritySummarySpamGetParamsDKIMNone RadarEmailSecuritySummarySpamGetParamsDKIM = "NONE"
	RadarEmailSecuritySummarySpamGetParamsDKIMFail RadarEmailSecuritySummarySpamGetParamsDKIM = "FAIL"
)

type RadarEmailSecuritySummarySpamGetParamsDmarc string

const (
	RadarEmailSecuritySummarySpamGetParamsDmarcPass RadarEmailSecuritySummarySpamGetParamsDmarc = "PASS"
	RadarEmailSecuritySummarySpamGetParamsDmarcNone RadarEmailSecuritySummarySpamGetParamsDmarc = "NONE"
	RadarEmailSecuritySummarySpamGetParamsDmarcFail RadarEmailSecuritySummarySpamGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecuritySummarySpamGetParamsFormat string

const (
	RadarEmailSecuritySummarySpamGetParamsFormatJson RadarEmailSecuritySummarySpamGetParamsFormat = "JSON"
	RadarEmailSecuritySummarySpamGetParamsFormatCsv  RadarEmailSecuritySummarySpamGetParamsFormat = "CSV"
)

type RadarEmailSecuritySummarySpamGetParamsSPF string

const (
	RadarEmailSecuritySummarySpamGetParamsSPFPass RadarEmailSecuritySummarySpamGetParamsSPF = "PASS"
	RadarEmailSecuritySummarySpamGetParamsSPFNone RadarEmailSecuritySummarySpamGetParamsSPF = "NONE"
	RadarEmailSecuritySummarySpamGetParamsSPFFail RadarEmailSecuritySummarySpamGetParamsSPF = "FAIL"
)

type RadarEmailSecuritySummarySpamGetResponseEnvelope struct {
	Result  RadarEmailSecuritySummarySpamGetResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailSecuritySummarySpamGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecuritySummarySpamGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecuritySummarySpamGetResponseEnvelope]
type radarEmailSecuritySummarySpamGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecuritySummarySpamGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
