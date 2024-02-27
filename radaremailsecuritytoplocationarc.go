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

// RadarEmailSecurityTopLocationARCService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationARCService] method instead.
type RadarEmailSecurityTopLocationARCService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationARCService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopLocationARCService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationARCService) {
	r = &RadarEmailSecurityTopLocationARCService{}
	r.Options = opts
	return
}

// Get the locations, by emails ARC validation.
func (r *RadarEmailSecurityTopLocationARCService) Get(ctx context.Context, arc RadarEmailSecurityTopLocationARCGetParamsARC, query RadarEmailSecurityTopLocationARCGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationARCGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopLocationARCGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/locations/arc/%v", arc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopLocationARCGetResponse struct {
	Meta RadarEmailSecurityTopLocationARCGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationARCGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationARCGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationARCGetResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopLocationARCGetResponse]
type radarEmailSecurityTopLocationARCGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationARCGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationARCGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopLocationARCGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationARCGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationARCGetResponseMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationARCGetResponseMeta]
type radarEmailSecurityTopLocationARCGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationARCGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationARCGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationARCGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationARCGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationARCGetResponseMetaDateRange]
type radarEmailSecurityTopLocationARCGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationARCGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationARCGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationARCGetResponseTop0 struct {
	ClientCountryAlpha2 string                                              `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                              `json:"clientCountryName,required"`
	Value               string                                              `json:"value,required"`
	JSON                radarEmailSecurityTopLocationARCGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationARCGetResponseTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationARCGetResponseTop0]
type radarEmailSecurityTopLocationARCGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationARCGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationARCGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationARCGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopLocationARCGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	DMARC param.Field[[]RadarEmailSecurityTopLocationARCGetParamsDMARC] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationARCGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopLocationARCGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationARCGetParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationARCGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ARC.
type RadarEmailSecurityTopLocationARCGetParamsARC string

const (
	RadarEmailSecurityTopLocationARCGetParamsARCPass RadarEmailSecurityTopLocationARCGetParamsARC = "PASS"
	RadarEmailSecurityTopLocationARCGetParamsARCNone RadarEmailSecurityTopLocationARCGetParamsARC = "NONE"
	RadarEmailSecurityTopLocationARCGetParamsARCFail RadarEmailSecurityTopLocationARCGetParamsARC = "FAIL"
)

type RadarEmailSecurityTopLocationARCGetParamsDateRange string

const (
	RadarEmailSecurityTopLocationARCGetParamsDateRange1d         RadarEmailSecurityTopLocationARCGetParamsDateRange = "1d"
	RadarEmailSecurityTopLocationARCGetParamsDateRange2d         RadarEmailSecurityTopLocationARCGetParamsDateRange = "2d"
	RadarEmailSecurityTopLocationARCGetParamsDateRange7d         RadarEmailSecurityTopLocationARCGetParamsDateRange = "7d"
	RadarEmailSecurityTopLocationARCGetParamsDateRange14d        RadarEmailSecurityTopLocationARCGetParamsDateRange = "14d"
	RadarEmailSecurityTopLocationARCGetParamsDateRange28d        RadarEmailSecurityTopLocationARCGetParamsDateRange = "28d"
	RadarEmailSecurityTopLocationARCGetParamsDateRange12w        RadarEmailSecurityTopLocationARCGetParamsDateRange = "12w"
	RadarEmailSecurityTopLocationARCGetParamsDateRange24w        RadarEmailSecurityTopLocationARCGetParamsDateRange = "24w"
	RadarEmailSecurityTopLocationARCGetParamsDateRange52w        RadarEmailSecurityTopLocationARCGetParamsDateRange = "52w"
	RadarEmailSecurityTopLocationARCGetParamsDateRange1dControl  RadarEmailSecurityTopLocationARCGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationARCGetParamsDateRange2dControl  RadarEmailSecurityTopLocationARCGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationARCGetParamsDateRange7dControl  RadarEmailSecurityTopLocationARCGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationARCGetParamsDateRange14dControl RadarEmailSecurityTopLocationARCGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationARCGetParamsDateRange28dControl RadarEmailSecurityTopLocationARCGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationARCGetParamsDateRange12wControl RadarEmailSecurityTopLocationARCGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationARCGetParamsDateRange24wControl RadarEmailSecurityTopLocationARCGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationARCGetParamsDKIM string

const (
	RadarEmailSecurityTopLocationARCGetParamsDKIMPass RadarEmailSecurityTopLocationARCGetParamsDKIM = "PASS"
	RadarEmailSecurityTopLocationARCGetParamsDKIMNone RadarEmailSecurityTopLocationARCGetParamsDKIM = "NONE"
	RadarEmailSecurityTopLocationARCGetParamsDKIMFail RadarEmailSecurityTopLocationARCGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopLocationARCGetParamsDMARC string

const (
	RadarEmailSecurityTopLocationARCGetParamsDMARCPass RadarEmailSecurityTopLocationARCGetParamsDMARC = "PASS"
	RadarEmailSecurityTopLocationARCGetParamsDMARCNone RadarEmailSecurityTopLocationARCGetParamsDMARC = "NONE"
	RadarEmailSecurityTopLocationARCGetParamsDMARCFail RadarEmailSecurityTopLocationARCGetParamsDMARC = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationARCGetParamsFormat string

const (
	RadarEmailSecurityTopLocationARCGetParamsFormatJson RadarEmailSecurityTopLocationARCGetParamsFormat = "JSON"
	RadarEmailSecurityTopLocationARCGetParamsFormatCsv  RadarEmailSecurityTopLocationARCGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationARCGetParamsSPF string

const (
	RadarEmailSecurityTopLocationARCGetParamsSPFPass RadarEmailSecurityTopLocationARCGetParamsSPF = "PASS"
	RadarEmailSecurityTopLocationARCGetParamsSPFNone RadarEmailSecurityTopLocationARCGetParamsSPF = "NONE"
	RadarEmailSecurityTopLocationARCGetParamsSPFFail RadarEmailSecurityTopLocationARCGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopLocationARCGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopLocationARCGetResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarEmailSecurityTopLocationARCGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopLocationARCGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationARCGetResponseEnvelope]
type radarEmailSecurityTopLocationARCGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationARCGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
