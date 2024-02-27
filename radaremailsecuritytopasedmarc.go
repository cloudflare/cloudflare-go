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

// RadarEmailSecurityTopAseDMARCService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseDMARCService] method instead.
type RadarEmailSecurityTopAseDMARCService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopAseDMARCService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseDMARCService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseDMARCService) {
	r = &RadarEmailSecurityTopAseDMARCService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS) by emails DMARC validation.
func (r *RadarEmailSecurityTopAseDMARCService) Get(ctx context.Context, dmarc RadarEmailSecurityTopAseDMARCGetParamsDMARC, query RadarEmailSecurityTopAseDMARCGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseDMARCGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopAseDMARCGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/ases/dmarc/%v", dmarc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopAseDMARCGetResponse struct {
	Meta RadarEmailSecurityTopAseDMARCGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseDMARCGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseDMARCGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseDMARCGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseDMARCGetResponse]
type radarEmailSecurityTopAseDMARCGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDMARCGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDMARCGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopAseDMARCGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseDMARCGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseDMARCGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseDMARCGetResponseMeta]
type radarEmailSecurityTopAseDMARCGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDMARCGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDMARCGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseDMARCGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseDMARCGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseDMARCGetResponseMetaDateRange]
type radarEmailSecurityTopAseDMARCGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDMARCGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseDMARCGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDMARCGetResponseTop0 struct {
	ClientASN    int64                                            `json:"clientASN,required"`
	ClientAsName string                                           `json:"clientASName,required"`
	Value        string                                           `json:"value,required"`
	JSON         radarEmailSecurityTopAseDMARCGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseDMARCGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseDMARCGetResponseTop0]
type radarEmailSecurityTopAseDMARCGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDMARCGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseDMARCGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTopAseDMARCGetParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseDMARCGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopAseDMARCGetParamsDKIM] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseDMARCGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopAseDMARCGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopAseDMARCGetParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityTopAseDMARCGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DMARC.
type RadarEmailSecurityTopAseDMARCGetParamsDMARC string

const (
	RadarEmailSecurityTopAseDMARCGetParamsDMARCPass RadarEmailSecurityTopAseDMARCGetParamsDMARC = "PASS"
	RadarEmailSecurityTopAseDMARCGetParamsDMARCNone RadarEmailSecurityTopAseDMARCGetParamsDMARC = "NONE"
	RadarEmailSecurityTopAseDMARCGetParamsDMARCFail RadarEmailSecurityTopAseDMARCGetParamsDMARC = "FAIL"
)

type RadarEmailSecurityTopAseDMARCGetParamsARC string

const (
	RadarEmailSecurityTopAseDMARCGetParamsARCPass RadarEmailSecurityTopAseDMARCGetParamsARC = "PASS"
	RadarEmailSecurityTopAseDMARCGetParamsARCNone RadarEmailSecurityTopAseDMARCGetParamsARC = "NONE"
	RadarEmailSecurityTopAseDMARCGetParamsARCFail RadarEmailSecurityTopAseDMARCGetParamsARC = "FAIL"
)

type RadarEmailSecurityTopAseDMARCGetParamsDateRange string

const (
	RadarEmailSecurityTopAseDMARCGetParamsDateRange1d         RadarEmailSecurityTopAseDMARCGetParamsDateRange = "1d"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange2d         RadarEmailSecurityTopAseDMARCGetParamsDateRange = "2d"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange7d         RadarEmailSecurityTopAseDMARCGetParamsDateRange = "7d"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange14d        RadarEmailSecurityTopAseDMARCGetParamsDateRange = "14d"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange28d        RadarEmailSecurityTopAseDMARCGetParamsDateRange = "28d"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange12w        RadarEmailSecurityTopAseDMARCGetParamsDateRange = "12w"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange24w        RadarEmailSecurityTopAseDMARCGetParamsDateRange = "24w"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange52w        RadarEmailSecurityTopAseDMARCGetParamsDateRange = "52w"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange1dControl  RadarEmailSecurityTopAseDMARCGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange2dControl  RadarEmailSecurityTopAseDMARCGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange7dControl  RadarEmailSecurityTopAseDMARCGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange14dControl RadarEmailSecurityTopAseDMARCGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange28dControl RadarEmailSecurityTopAseDMARCGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange12wControl RadarEmailSecurityTopAseDMARCGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseDMARCGetParamsDateRange24wControl RadarEmailSecurityTopAseDMARCGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseDMARCGetParamsDKIM string

const (
	RadarEmailSecurityTopAseDMARCGetParamsDKIMPass RadarEmailSecurityTopAseDMARCGetParamsDKIM = "PASS"
	RadarEmailSecurityTopAseDMARCGetParamsDKIMNone RadarEmailSecurityTopAseDMARCGetParamsDKIM = "NONE"
	RadarEmailSecurityTopAseDMARCGetParamsDKIMFail RadarEmailSecurityTopAseDMARCGetParamsDKIM = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseDMARCGetParamsFormat string

const (
	RadarEmailSecurityTopAseDMARCGetParamsFormatJson RadarEmailSecurityTopAseDMARCGetParamsFormat = "JSON"
	RadarEmailSecurityTopAseDMARCGetParamsFormatCsv  RadarEmailSecurityTopAseDMARCGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopAseDMARCGetParamsSPF string

const (
	RadarEmailSecurityTopAseDMARCGetParamsSPFPass RadarEmailSecurityTopAseDMARCGetParamsSPF = "PASS"
	RadarEmailSecurityTopAseDMARCGetParamsSPFNone RadarEmailSecurityTopAseDMARCGetParamsSPF = "NONE"
	RadarEmailSecurityTopAseDMARCGetParamsSPFFail RadarEmailSecurityTopAseDMARCGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopAseDMARCGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopAseDMARCGetResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailSecurityTopAseDMARCGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopAseDMARCGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseDMARCGetResponseEnvelope]
type radarEmailSecurityTopAseDMARCGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseDMARCGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
