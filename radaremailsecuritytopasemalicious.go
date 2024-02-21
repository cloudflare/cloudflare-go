// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEmailSecurityTopAseMaliciousService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseMaliciousService] method instead.
type RadarEmailSecurityTopAseMaliciousService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopAseMaliciousService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseMaliciousService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseMaliciousService) {
	r = &RadarEmailSecurityTopAseMaliciousService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by emails classified as Malicious or not.
func (r *RadarEmailSecurityTopAseMaliciousService) Get(ctx context.Context, malicious RadarEmailSecurityTopAseMaliciousGetParamsMalicious, query RadarEmailSecurityTopAseMaliciousGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseMaliciousGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopAseMaliciousGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/ases/malicious/%v", malicious)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopAseMaliciousGetResponse struct {
	Meta RadarEmailSecurityTopAseMaliciousGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseMaliciousGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseMaliciousGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseMaliciousGetResponse]
type radarEmailSecurityTopAseMaliciousGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseMaliciousGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopAseMaliciousGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseMaliciousGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseMaliciousGetResponseMeta]
type radarEmailSecurityTopAseMaliciousGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseMaliciousGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseMaliciousGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopAseMaliciousGetResponseMetaDateRange]
type radarEmailSecurityTopAseMaliciousGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseMaliciousGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseMaliciousGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetResponseTop0 struct {
	ClientAsn    int64                                                `json:"clientASN,required"`
	ClientAsName string                                               `json:"clientASName,required"`
	Value        string                                               `json:"value,required"`
	JSON         radarEmailSecurityTopAseMaliciousGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseMaliciousGetResponseTop0]
type radarEmailSecurityTopAseMaliciousGetResponseTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseMaliciousGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopAseMaliciousGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseMaliciousGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopAseMaliciousGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopAseMaliciousGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseMaliciousGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopAseMaliciousGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopAseMaliciousGetParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopAseMaliciousGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Malicious.
type RadarEmailSecurityTopAseMaliciousGetParamsMalicious string

const (
	RadarEmailSecurityTopAseMaliciousGetParamsMaliciousMalicious    RadarEmailSecurityTopAseMaliciousGetParamsMalicious = "MALICIOUS"
	RadarEmailSecurityTopAseMaliciousGetParamsMaliciousNotMalicious RadarEmailSecurityTopAseMaliciousGetParamsMalicious = "NOT_MALICIOUS"
)

type RadarEmailSecurityTopAseMaliciousGetParamsArc string

const (
	RadarEmailSecurityTopAseMaliciousGetParamsArcPass RadarEmailSecurityTopAseMaliciousGetParamsArc = "PASS"
	RadarEmailSecurityTopAseMaliciousGetParamsArcNone RadarEmailSecurityTopAseMaliciousGetParamsArc = "NONE"
	RadarEmailSecurityTopAseMaliciousGetParamsArcFail RadarEmailSecurityTopAseMaliciousGetParamsArc = "FAIL"
)

type RadarEmailSecurityTopAseMaliciousGetParamsDateRange string

const (
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange1d         RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "1d"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange2d         RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "2d"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange7d         RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "7d"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange14d        RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "14d"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange28d        RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "28d"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange12w        RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "12w"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange24w        RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "24w"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange52w        RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "52w"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange1dControl  RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange2dControl  RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange7dControl  RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange14dControl RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange28dControl RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange12wControl RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseMaliciousGetParamsDateRange24wControl RadarEmailSecurityTopAseMaliciousGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseMaliciousGetParamsDKIM string

const (
	RadarEmailSecurityTopAseMaliciousGetParamsDKIMPass RadarEmailSecurityTopAseMaliciousGetParamsDKIM = "PASS"
	RadarEmailSecurityTopAseMaliciousGetParamsDKIMNone RadarEmailSecurityTopAseMaliciousGetParamsDKIM = "NONE"
	RadarEmailSecurityTopAseMaliciousGetParamsDKIMFail RadarEmailSecurityTopAseMaliciousGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopAseMaliciousGetParamsDmarc string

const (
	RadarEmailSecurityTopAseMaliciousGetParamsDmarcPass RadarEmailSecurityTopAseMaliciousGetParamsDmarc = "PASS"
	RadarEmailSecurityTopAseMaliciousGetParamsDmarcNone RadarEmailSecurityTopAseMaliciousGetParamsDmarc = "NONE"
	RadarEmailSecurityTopAseMaliciousGetParamsDmarcFail RadarEmailSecurityTopAseMaliciousGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseMaliciousGetParamsFormat string

const (
	RadarEmailSecurityTopAseMaliciousGetParamsFormatJson RadarEmailSecurityTopAseMaliciousGetParamsFormat = "JSON"
	RadarEmailSecurityTopAseMaliciousGetParamsFormatCsv  RadarEmailSecurityTopAseMaliciousGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopAseMaliciousGetParamsSPF string

const (
	RadarEmailSecurityTopAseMaliciousGetParamsSPFPass RadarEmailSecurityTopAseMaliciousGetParamsSPF = "PASS"
	RadarEmailSecurityTopAseMaliciousGetParamsSPFNone RadarEmailSecurityTopAseMaliciousGetParamsSPF = "NONE"
	RadarEmailSecurityTopAseMaliciousGetParamsSPFFail RadarEmailSecurityTopAseMaliciousGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopAseMaliciousGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopAseMaliciousGetResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarEmailSecurityTopAseMaliciousGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseMaliciousGetResponseEnvelope]
type radarEmailSecurityTopAseMaliciousGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseMaliciousGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
