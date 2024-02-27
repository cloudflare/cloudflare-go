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

// RadarEmailSecurityTopAseSPFService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseSPFService] method instead.
type RadarEmailSecurityTopAseSPFService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopAseSPFService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseSPFService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseSPFService) {
	r = &RadarEmailSecurityTopAseSPFService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS) by email SPF validation.
func (r *RadarEmailSecurityTopAseSPFService) Get(ctx context.Context, spf RadarEmailSecurityTopAseSPFGetParamsSPF, query RadarEmailSecurityTopAseSPFGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseSPFGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopAseSPFGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/ases/spf/%v", spf)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopAseSPFGetResponse struct {
	Meta RadarEmailSecurityTopAseSPFGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseSPFGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseSPFGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseSPFGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseSPFGetResponse]
type radarEmailSecurityTopAseSPFGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSPFGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSPFGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopAseSPFGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseSPFGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseSPFGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseSPFGetResponseMeta]
type radarEmailSecurityTopAseSPFGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSPFGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSPFGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseSPFGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseSPFGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseSPFGetResponseMetaDateRange]
type radarEmailSecurityTopAseSPFGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSPFGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseSPFGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSPFGetResponseTop0 struct {
	ClientASN    int64                                          `json:"clientASN,required"`
	ClientAsName string                                         `json:"clientASName,required"`
	Value        string                                         `json:"value,required"`
	JSON         radarEmailSecurityTopAseSPFGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseSPFGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseSPFGetResponseTop0]
type radarEmailSecurityTopAseSPFGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSPFGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSPFGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTopAseSPFGetParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseSPFGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopAseSPFGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTopAseSPFGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseSPFGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailSecurityTopAseSPFGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecurityTopAseSPFGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// SPF.
type RadarEmailSecurityTopAseSPFGetParamsSPF string

const (
	RadarEmailSecurityTopAseSPFGetParamsSPFPass RadarEmailSecurityTopAseSPFGetParamsSPF = "PASS"
	RadarEmailSecurityTopAseSPFGetParamsSPFNone RadarEmailSecurityTopAseSPFGetParamsSPF = "NONE"
	RadarEmailSecurityTopAseSPFGetParamsSPFFail RadarEmailSecurityTopAseSPFGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopAseSPFGetParamsARC string

const (
	RadarEmailSecurityTopAseSPFGetParamsARCPass RadarEmailSecurityTopAseSPFGetParamsARC = "PASS"
	RadarEmailSecurityTopAseSPFGetParamsARCNone RadarEmailSecurityTopAseSPFGetParamsARC = "NONE"
	RadarEmailSecurityTopAseSPFGetParamsARCFail RadarEmailSecurityTopAseSPFGetParamsARC = "FAIL"
)

type RadarEmailSecurityTopAseSPFGetParamsDateRange string

const (
	RadarEmailSecurityTopAseSPFGetParamsDateRange1d         RadarEmailSecurityTopAseSPFGetParamsDateRange = "1d"
	RadarEmailSecurityTopAseSPFGetParamsDateRange2d         RadarEmailSecurityTopAseSPFGetParamsDateRange = "2d"
	RadarEmailSecurityTopAseSPFGetParamsDateRange7d         RadarEmailSecurityTopAseSPFGetParamsDateRange = "7d"
	RadarEmailSecurityTopAseSPFGetParamsDateRange14d        RadarEmailSecurityTopAseSPFGetParamsDateRange = "14d"
	RadarEmailSecurityTopAseSPFGetParamsDateRange28d        RadarEmailSecurityTopAseSPFGetParamsDateRange = "28d"
	RadarEmailSecurityTopAseSPFGetParamsDateRange12w        RadarEmailSecurityTopAseSPFGetParamsDateRange = "12w"
	RadarEmailSecurityTopAseSPFGetParamsDateRange24w        RadarEmailSecurityTopAseSPFGetParamsDateRange = "24w"
	RadarEmailSecurityTopAseSPFGetParamsDateRange52w        RadarEmailSecurityTopAseSPFGetParamsDateRange = "52w"
	RadarEmailSecurityTopAseSPFGetParamsDateRange1dControl  RadarEmailSecurityTopAseSPFGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseSPFGetParamsDateRange2dControl  RadarEmailSecurityTopAseSPFGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseSPFGetParamsDateRange7dControl  RadarEmailSecurityTopAseSPFGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseSPFGetParamsDateRange14dControl RadarEmailSecurityTopAseSPFGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseSPFGetParamsDateRange28dControl RadarEmailSecurityTopAseSPFGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseSPFGetParamsDateRange12wControl RadarEmailSecurityTopAseSPFGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseSPFGetParamsDateRange24wControl RadarEmailSecurityTopAseSPFGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseSPFGetParamsDKIM string

const (
	RadarEmailSecurityTopAseSPFGetParamsDKIMPass RadarEmailSecurityTopAseSPFGetParamsDKIM = "PASS"
	RadarEmailSecurityTopAseSPFGetParamsDKIMNone RadarEmailSecurityTopAseSPFGetParamsDKIM = "NONE"
	RadarEmailSecurityTopAseSPFGetParamsDKIMFail RadarEmailSecurityTopAseSPFGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopAseSPFGetParamsDMARC string

const (
	RadarEmailSecurityTopAseSPFGetParamsDMARCPass RadarEmailSecurityTopAseSPFGetParamsDMARC = "PASS"
	RadarEmailSecurityTopAseSPFGetParamsDMARCNone RadarEmailSecurityTopAseSPFGetParamsDMARC = "NONE"
	RadarEmailSecurityTopAseSPFGetParamsDMARCFail RadarEmailSecurityTopAseSPFGetParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseSPFGetParamsFormat string

const (
	RadarEmailSecurityTopAseSPFGetParamsFormatJson RadarEmailSecurityTopAseSPFGetParamsFormat = "JSON"
	RadarEmailSecurityTopAseSPFGetParamsFormatCsv  RadarEmailSecurityTopAseSPFGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopAseSPFGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopAseSPFGetResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarEmailSecurityTopAseSPFGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopAseSPFGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseSPFGetResponseEnvelope]
type radarEmailSecurityTopAseSPFGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSPFGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
