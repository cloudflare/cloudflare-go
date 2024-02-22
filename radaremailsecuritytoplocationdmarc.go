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

// RadarEmailSecurityTopLocationDMARCService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationDMARCService] method instead.
type RadarEmailSecurityTopLocationDMARCService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationDMARCService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityTopLocationDMARCService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationDMARCService) {
	r = &RadarEmailSecurityTopLocationDMARCService{}
	r.Options = opts
	return
}

// Get the locations by email DMARC validation.
func (r *RadarEmailSecurityTopLocationDMARCService) Get(ctx context.Context, dmarc RadarEmailSecurityTopLocationDMARCGetParamsDMARC, query RadarEmailSecurityTopLocationDMARCGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationDMARCGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopLocationDMARCGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/locations/dmarc/%v", dmarc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopLocationDMARCGetResponse struct {
	Meta RadarEmailSecurityTopLocationDMARCGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationDMARCGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationDMARCGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationDMARCGetResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopLocationDMARCGetResponse]
type radarEmailSecurityTopLocationDMARCGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDMARCGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDMARCGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopLocationDMARCGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                          `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationDMARCGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationDMARCGetResponseMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationDMARCGetResponseMeta]
type radarEmailSecurityTopLocationDMARCGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDMARCGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDMARCGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationDMARCGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationDMARCGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationDMARCGetResponseMetaDateRange]
type radarEmailSecurityTopLocationDMARCGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDMARCGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                       `json:"level"`
	JSON        radarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                        `json:"dataSource,required"`
	Description     string                                                                        `json:"description,required"`
	EventType       string                                                                        `json:"eventType,required"`
	IsInstantaneous interface{}                                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                                     `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationDMARCGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDMARCGetResponseTop0 struct {
	ClientCountryAlpha2 string                                                `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                `json:"clientCountryName,required"`
	Value               string                                                `json:"value,required"`
	JSON                radarEmailSecurityTopLocationDMARCGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationDMARCGetResponseTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationDMARCGetResponseTop0]
type radarEmailSecurityTopLocationDMARCGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDMARCGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDMARCGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	ARC param.Field[[]RadarEmailSecurityTopLocationDMARCGetParamsARC] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationDMARCGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopLocationDMARCGetParamsDKIM] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationDMARCGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopLocationDMARCGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationDMARCGetParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationDMARCGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DMARC.
type RadarEmailSecurityTopLocationDMARCGetParamsDMARC string

const (
	RadarEmailSecurityTopLocationDMARCGetParamsDMARCPass RadarEmailSecurityTopLocationDMARCGetParamsDMARC = "PASS"
	RadarEmailSecurityTopLocationDMARCGetParamsDMARCNone RadarEmailSecurityTopLocationDMARCGetParamsDMARC = "NONE"
	RadarEmailSecurityTopLocationDMARCGetParamsDMARCFail RadarEmailSecurityTopLocationDMARCGetParamsDMARC = "FAIL"
)

type RadarEmailSecurityTopLocationDMARCGetParamsARC string

const (
	RadarEmailSecurityTopLocationDMARCGetParamsARCPass RadarEmailSecurityTopLocationDMARCGetParamsARC = "PASS"
	RadarEmailSecurityTopLocationDMARCGetParamsARCNone RadarEmailSecurityTopLocationDMARCGetParamsARC = "NONE"
	RadarEmailSecurityTopLocationDMARCGetParamsARCFail RadarEmailSecurityTopLocationDMARCGetParamsARC = "FAIL"
)

type RadarEmailSecurityTopLocationDMARCGetParamsDateRange string

const (
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange1d         RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "1d"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange2d         RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "2d"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange7d         RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "7d"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange14d        RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "14d"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange28d        RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "28d"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange12w        RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "12w"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange24w        RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "24w"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange52w        RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "52w"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange1dControl  RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange2dControl  RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange7dControl  RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange14dControl RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange28dControl RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange12wControl RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationDMARCGetParamsDateRange24wControl RadarEmailSecurityTopLocationDMARCGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationDMARCGetParamsDKIM string

const (
	RadarEmailSecurityTopLocationDMARCGetParamsDKIMPass RadarEmailSecurityTopLocationDMARCGetParamsDKIM = "PASS"
	RadarEmailSecurityTopLocationDMARCGetParamsDKIMNone RadarEmailSecurityTopLocationDMARCGetParamsDKIM = "NONE"
	RadarEmailSecurityTopLocationDMARCGetParamsDKIMFail RadarEmailSecurityTopLocationDMARCGetParamsDKIM = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationDMARCGetParamsFormat string

const (
	RadarEmailSecurityTopLocationDMARCGetParamsFormatJson RadarEmailSecurityTopLocationDMARCGetParamsFormat = "JSON"
	RadarEmailSecurityTopLocationDMARCGetParamsFormatCsv  RadarEmailSecurityTopLocationDMARCGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationDMARCGetParamsSPF string

const (
	RadarEmailSecurityTopLocationDMARCGetParamsSPFPass RadarEmailSecurityTopLocationDMARCGetParamsSPF = "PASS"
	RadarEmailSecurityTopLocationDMARCGetParamsSPFNone RadarEmailSecurityTopLocationDMARCGetParamsSPF = "NONE"
	RadarEmailSecurityTopLocationDMARCGetParamsSPFFail RadarEmailSecurityTopLocationDMARCGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopLocationDMARCGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopLocationDMARCGetResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarEmailSecurityTopLocationDMARCGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopLocationDMARCGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationDMARCGetResponseEnvelope]
type radarEmailSecurityTopLocationDMARCGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDMARCGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
