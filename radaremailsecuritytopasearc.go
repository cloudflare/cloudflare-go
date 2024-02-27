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

// RadarEmailSecurityTopAseARCService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseARCService] method instead.
type RadarEmailSecurityTopAseARCService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopAseARCService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseARCService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseARCService) {
	r = &RadarEmailSecurityTopAseARCService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS) by emails ARC validation.
func (r *RadarEmailSecurityTopAseARCService) Get(ctx context.Context, arc RadarEmailSecurityTopAseARCGetParamsARC, query RadarEmailSecurityTopAseARCGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseARCGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopAseARCGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/ases/arc/%v", arc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopAseARCGetResponse struct {
	Meta RadarEmailSecurityTopAseARCGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseARCGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseARCGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseARCGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseARCGetResponse]
type radarEmailSecurityTopAseARCGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseARCGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseARCGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopAseARCGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseARCGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseARCGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseARCGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseARCGetResponseMeta]
type radarEmailSecurityTopAseARCGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseARCGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseARCGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseARCGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseARCGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseARCGetResponseMetaDateRange]
type radarEmailSecurityTopAseARCGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseARCGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseARCGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseARCGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarEmailSecurityTopAseARCGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseARCGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopAseARCGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopAseARCGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseARCGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseARCGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseARCGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseARCGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopAseARCGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseARCGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseARCGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseARCGetResponseTop0 struct {
	ClientASN    int64                                          `json:"clientASN,required"`
	ClientAsName string                                         `json:"clientASName,required"`
	Value        string                                         `json:"value,required"`
	JSON         radarEmailSecurityTopAseARCGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseARCGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseARCGetResponseTop0]
type radarEmailSecurityTopAseARCGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseARCGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseARCGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseARCGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopAseARCGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTopAseARCGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseARCGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopAseARCGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopAseARCGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecurityTopAseARCGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ARC.
type RadarEmailSecurityTopAseARCGetParamsARC string

const (
	RadarEmailSecurityTopAseARCGetParamsARCPass RadarEmailSecurityTopAseARCGetParamsARC = "PASS"
	RadarEmailSecurityTopAseARCGetParamsARCNone RadarEmailSecurityTopAseARCGetParamsARC = "NONE"
	RadarEmailSecurityTopAseARCGetParamsARCFail RadarEmailSecurityTopAseARCGetParamsARC = "FAIL"
)

type RadarEmailSecurityTopAseARCGetParamsDateRange string

const (
	RadarEmailSecurityTopAseARCGetParamsDateRange1d         RadarEmailSecurityTopAseARCGetParamsDateRange = "1d"
	RadarEmailSecurityTopAseARCGetParamsDateRange2d         RadarEmailSecurityTopAseARCGetParamsDateRange = "2d"
	RadarEmailSecurityTopAseARCGetParamsDateRange7d         RadarEmailSecurityTopAseARCGetParamsDateRange = "7d"
	RadarEmailSecurityTopAseARCGetParamsDateRange14d        RadarEmailSecurityTopAseARCGetParamsDateRange = "14d"
	RadarEmailSecurityTopAseARCGetParamsDateRange28d        RadarEmailSecurityTopAseARCGetParamsDateRange = "28d"
	RadarEmailSecurityTopAseARCGetParamsDateRange12w        RadarEmailSecurityTopAseARCGetParamsDateRange = "12w"
	RadarEmailSecurityTopAseARCGetParamsDateRange24w        RadarEmailSecurityTopAseARCGetParamsDateRange = "24w"
	RadarEmailSecurityTopAseARCGetParamsDateRange52w        RadarEmailSecurityTopAseARCGetParamsDateRange = "52w"
	RadarEmailSecurityTopAseARCGetParamsDateRange1dControl  RadarEmailSecurityTopAseARCGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseARCGetParamsDateRange2dControl  RadarEmailSecurityTopAseARCGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseARCGetParamsDateRange7dControl  RadarEmailSecurityTopAseARCGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseARCGetParamsDateRange14dControl RadarEmailSecurityTopAseARCGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseARCGetParamsDateRange28dControl RadarEmailSecurityTopAseARCGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseARCGetParamsDateRange12wControl RadarEmailSecurityTopAseARCGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseARCGetParamsDateRange24wControl RadarEmailSecurityTopAseARCGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseARCGetParamsDKIM string

const (
	RadarEmailSecurityTopAseARCGetParamsDKIMPass RadarEmailSecurityTopAseARCGetParamsDKIM = "PASS"
	RadarEmailSecurityTopAseARCGetParamsDKIMNone RadarEmailSecurityTopAseARCGetParamsDKIM = "NONE"
	RadarEmailSecurityTopAseARCGetParamsDKIMFail RadarEmailSecurityTopAseARCGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopAseARCGetParamsDMARC string

const (
	RadarEmailSecurityTopAseARCGetParamsDMARCPass RadarEmailSecurityTopAseARCGetParamsDMARC = "PASS"
	RadarEmailSecurityTopAseARCGetParamsDMARCNone RadarEmailSecurityTopAseARCGetParamsDMARC = "NONE"
	RadarEmailSecurityTopAseARCGetParamsDMARCFail RadarEmailSecurityTopAseARCGetParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseARCGetParamsFormat string

const (
	RadarEmailSecurityTopAseARCGetParamsFormatJson RadarEmailSecurityTopAseARCGetParamsFormat = "JSON"
	RadarEmailSecurityTopAseARCGetParamsFormatCsv  RadarEmailSecurityTopAseARCGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopAseARCGetParamsSPF string

const (
	RadarEmailSecurityTopAseARCGetParamsSPFPass RadarEmailSecurityTopAseARCGetParamsSPF = "PASS"
	RadarEmailSecurityTopAseARCGetParamsSPFNone RadarEmailSecurityTopAseARCGetParamsSPF = "NONE"
	RadarEmailSecurityTopAseARCGetParamsSPFFail RadarEmailSecurityTopAseARCGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopAseARCGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopAseARCGetResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarEmailSecurityTopAseARCGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopAseARCGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseARCGetResponseEnvelope]
type radarEmailSecurityTopAseARCGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseARCGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
