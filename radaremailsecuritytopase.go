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

// RadarEmailSecurityTopAseService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseService] method instead.
type RadarEmailSecurityTopAseService struct {
	Options   []option.RequestOption
	Arc       *RadarEmailSecurityTopAseArcService
	DKIM      *RadarEmailSecurityTopAseDKIMService
	Dmarc     *RadarEmailSecurityTopAseDmarcService
	Malicious *RadarEmailSecurityTopAseMaliciousService
	Spam      *RadarEmailSecurityTopAseSpamService
	SPF       *RadarEmailSecurityTopAseSPFService
}

// NewRadarEmailSecurityTopAseService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseService) {
	r = &RadarEmailSecurityTopAseService{}
	r.Options = opts
	r.Arc = NewRadarEmailSecurityTopAseArcService(opts...)
	r.DKIM = NewRadarEmailSecurityTopAseDKIMService(opts...)
	r.Dmarc = NewRadarEmailSecurityTopAseDmarcService(opts...)
	r.Malicious = NewRadarEmailSecurityTopAseMaliciousService(opts...)
	r.Spam = NewRadarEmailSecurityTopAseSpamService(opts...)
	r.SPF = NewRadarEmailSecurityTopAseSPFService(opts...)
	return
}

// Get the top autonomous systems (AS) by email messages. Values are a percentage
// out of the total emails.
func (r *RadarEmailSecurityTopAseService) Get(ctx context.Context, query RadarEmailSecurityTopAseGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopAseGetResponseEnvelope
	path := "radar/email/security/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopAseGetResponse struct {
	Meta RadarEmailSecurityTopAseGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseGetResponse]
type radarEmailSecurityTopAseGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopAseGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseGetResponseMetaJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseGetResponseMeta]
type radarEmailSecurityTopAseGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseGetResponseMetaDateRange]
type radarEmailSecurityTopAseGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarEmailSecurityTopAseGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopAseGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopAseGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseGetResponseTop0 struct {
	ClientAsn    int64                                       `json:"clientASN,required"`
	ClientAsName string                                      `json:"clientASName,required"`
	Value        string                                      `json:"value,required"`
	JSON         radarEmailSecurityTopAseGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseGetResponseTop0JSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseGetResponseTop0]
type radarEmailSecurityTopAseGetResponseTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopAseGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopAseGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopAseGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopAseGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopAseGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecurityTopAseGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecurityTopAseGetParamsArc string

const (
	RadarEmailSecurityTopAseGetParamsArcPass RadarEmailSecurityTopAseGetParamsArc = "PASS"
	RadarEmailSecurityTopAseGetParamsArcNone RadarEmailSecurityTopAseGetParamsArc = "NONE"
	RadarEmailSecurityTopAseGetParamsArcFail RadarEmailSecurityTopAseGetParamsArc = "FAIL"
)

type RadarEmailSecurityTopAseGetParamsDateRange string

const (
	RadarEmailSecurityTopAseGetParamsDateRange1d         RadarEmailSecurityTopAseGetParamsDateRange = "1d"
	RadarEmailSecurityTopAseGetParamsDateRange2d         RadarEmailSecurityTopAseGetParamsDateRange = "2d"
	RadarEmailSecurityTopAseGetParamsDateRange7d         RadarEmailSecurityTopAseGetParamsDateRange = "7d"
	RadarEmailSecurityTopAseGetParamsDateRange14d        RadarEmailSecurityTopAseGetParamsDateRange = "14d"
	RadarEmailSecurityTopAseGetParamsDateRange28d        RadarEmailSecurityTopAseGetParamsDateRange = "28d"
	RadarEmailSecurityTopAseGetParamsDateRange12w        RadarEmailSecurityTopAseGetParamsDateRange = "12w"
	RadarEmailSecurityTopAseGetParamsDateRange24w        RadarEmailSecurityTopAseGetParamsDateRange = "24w"
	RadarEmailSecurityTopAseGetParamsDateRange52w        RadarEmailSecurityTopAseGetParamsDateRange = "52w"
	RadarEmailSecurityTopAseGetParamsDateRange1dControl  RadarEmailSecurityTopAseGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseGetParamsDateRange2dControl  RadarEmailSecurityTopAseGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseGetParamsDateRange7dControl  RadarEmailSecurityTopAseGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseGetParamsDateRange14dControl RadarEmailSecurityTopAseGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseGetParamsDateRange28dControl RadarEmailSecurityTopAseGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseGetParamsDateRange12wControl RadarEmailSecurityTopAseGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseGetParamsDateRange24wControl RadarEmailSecurityTopAseGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseGetParamsDKIM string

const (
	RadarEmailSecurityTopAseGetParamsDKIMPass RadarEmailSecurityTopAseGetParamsDKIM = "PASS"
	RadarEmailSecurityTopAseGetParamsDKIMNone RadarEmailSecurityTopAseGetParamsDKIM = "NONE"
	RadarEmailSecurityTopAseGetParamsDKIMFail RadarEmailSecurityTopAseGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopAseGetParamsDmarc string

const (
	RadarEmailSecurityTopAseGetParamsDmarcPass RadarEmailSecurityTopAseGetParamsDmarc = "PASS"
	RadarEmailSecurityTopAseGetParamsDmarcNone RadarEmailSecurityTopAseGetParamsDmarc = "NONE"
	RadarEmailSecurityTopAseGetParamsDmarcFail RadarEmailSecurityTopAseGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseGetParamsFormat string

const (
	RadarEmailSecurityTopAseGetParamsFormatJson RadarEmailSecurityTopAseGetParamsFormat = "JSON"
	RadarEmailSecurityTopAseGetParamsFormatCsv  RadarEmailSecurityTopAseGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopAseGetParamsSPF string

const (
	RadarEmailSecurityTopAseGetParamsSPFPass RadarEmailSecurityTopAseGetParamsSPF = "PASS"
	RadarEmailSecurityTopAseGetParamsSPFNone RadarEmailSecurityTopAseGetParamsSPF = "NONE"
	RadarEmailSecurityTopAseGetParamsSPFFail RadarEmailSecurityTopAseGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopAseGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopAseGetResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailSecurityTopAseGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopAseGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseGetResponseEnvelope]
type radarEmailSecurityTopAseGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
