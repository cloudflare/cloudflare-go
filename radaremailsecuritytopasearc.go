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

// RadarEmailSecurityTopAseArcService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseArcService] method instead.
type RadarEmailSecurityTopAseArcService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopAseArcService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseArcService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseArcService) {
	r = &RadarEmailSecurityTopAseArcService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS) by emails ARC validation.
func (r *RadarEmailSecurityTopAseArcService) Get(ctx context.Context, arc RadarEmailSecurityTopAseArcGetParamsArc, query RadarEmailSecurityTopAseArcGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseArcGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopAseArcGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/ases/arc/%v", arc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopAseArcGetResponse struct {
	Meta RadarEmailSecurityTopAseArcGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseArcGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseArcGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseArcGetResponse]
type radarEmailSecurityTopAseArcGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseArcGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopAseArcGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseArcGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseArcGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseArcGetResponseMeta]
type radarEmailSecurityTopAseArcGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseArcGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseArcGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseArcGetResponseMetaDateRange]
type radarEmailSecurityTopAseArcGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseArcGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseArcGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarEmailSecurityTopAseArcGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopAseArcGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopAseArcGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseArcGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseArcGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopAseArcGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseArcGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseArcGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetResponseTop0 struct {
	ClientAsn    int64                                          `json:"clientASN,required"`
	ClientAsName string                                         `json:"clientASName,required"`
	Value        string                                         `json:"value,required"`
	JSON         radarEmailSecurityTopAseArcGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseArcGetResponseTop0]
type radarEmailSecurityTopAseArcGetResponseTop0JSON struct {
	ClientAsn    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseArcGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseArcGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopAseArcGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopAseArcGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseArcGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopAseArcGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopAseArcGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecurityTopAseArcGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ARC.
type RadarEmailSecurityTopAseArcGetParamsArc string

const (
	RadarEmailSecurityTopAseArcGetParamsArcPass RadarEmailSecurityTopAseArcGetParamsArc = "PASS"
	RadarEmailSecurityTopAseArcGetParamsArcNone RadarEmailSecurityTopAseArcGetParamsArc = "NONE"
	RadarEmailSecurityTopAseArcGetParamsArcFail RadarEmailSecurityTopAseArcGetParamsArc = "FAIL"
)

type RadarEmailSecurityTopAseArcGetParamsDateRange string

const (
	RadarEmailSecurityTopAseArcGetParamsDateRange1d         RadarEmailSecurityTopAseArcGetParamsDateRange = "1d"
	RadarEmailSecurityTopAseArcGetParamsDateRange2d         RadarEmailSecurityTopAseArcGetParamsDateRange = "2d"
	RadarEmailSecurityTopAseArcGetParamsDateRange7d         RadarEmailSecurityTopAseArcGetParamsDateRange = "7d"
	RadarEmailSecurityTopAseArcGetParamsDateRange14d        RadarEmailSecurityTopAseArcGetParamsDateRange = "14d"
	RadarEmailSecurityTopAseArcGetParamsDateRange28d        RadarEmailSecurityTopAseArcGetParamsDateRange = "28d"
	RadarEmailSecurityTopAseArcGetParamsDateRange12w        RadarEmailSecurityTopAseArcGetParamsDateRange = "12w"
	RadarEmailSecurityTopAseArcGetParamsDateRange24w        RadarEmailSecurityTopAseArcGetParamsDateRange = "24w"
	RadarEmailSecurityTopAseArcGetParamsDateRange52w        RadarEmailSecurityTopAseArcGetParamsDateRange = "52w"
	RadarEmailSecurityTopAseArcGetParamsDateRange1dControl  RadarEmailSecurityTopAseArcGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseArcGetParamsDateRange2dControl  RadarEmailSecurityTopAseArcGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseArcGetParamsDateRange7dControl  RadarEmailSecurityTopAseArcGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseArcGetParamsDateRange14dControl RadarEmailSecurityTopAseArcGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseArcGetParamsDateRange28dControl RadarEmailSecurityTopAseArcGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseArcGetParamsDateRange12wControl RadarEmailSecurityTopAseArcGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseArcGetParamsDateRange24wControl RadarEmailSecurityTopAseArcGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseArcGetParamsDKIM string

const (
	RadarEmailSecurityTopAseArcGetParamsDKIMPass RadarEmailSecurityTopAseArcGetParamsDKIM = "PASS"
	RadarEmailSecurityTopAseArcGetParamsDKIMNone RadarEmailSecurityTopAseArcGetParamsDKIM = "NONE"
	RadarEmailSecurityTopAseArcGetParamsDKIMFail RadarEmailSecurityTopAseArcGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopAseArcGetParamsDmarc string

const (
	RadarEmailSecurityTopAseArcGetParamsDmarcPass RadarEmailSecurityTopAseArcGetParamsDmarc = "PASS"
	RadarEmailSecurityTopAseArcGetParamsDmarcNone RadarEmailSecurityTopAseArcGetParamsDmarc = "NONE"
	RadarEmailSecurityTopAseArcGetParamsDmarcFail RadarEmailSecurityTopAseArcGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseArcGetParamsFormat string

const (
	RadarEmailSecurityTopAseArcGetParamsFormatJson RadarEmailSecurityTopAseArcGetParamsFormat = "JSON"
	RadarEmailSecurityTopAseArcGetParamsFormatCsv  RadarEmailSecurityTopAseArcGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopAseArcGetParamsSPF string

const (
	RadarEmailSecurityTopAseArcGetParamsSPFPass RadarEmailSecurityTopAseArcGetParamsSPF = "PASS"
	RadarEmailSecurityTopAseArcGetParamsSPFNone RadarEmailSecurityTopAseArcGetParamsSPF = "NONE"
	RadarEmailSecurityTopAseArcGetParamsSPFFail RadarEmailSecurityTopAseArcGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopAseArcGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopAseArcGetResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarEmailSecurityTopAseArcGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseArcGetResponseEnvelope]
type radarEmailSecurityTopAseArcGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseArcGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
