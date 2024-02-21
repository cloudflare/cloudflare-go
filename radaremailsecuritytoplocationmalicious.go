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

// RadarEmailSecurityTopLocationMaliciousService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEmailSecurityTopLocationMaliciousService] method instead.
type RadarEmailSecurityTopLocationMaliciousService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationMaliciousService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityTopLocationMaliciousService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationMaliciousService) {
	r = &RadarEmailSecurityTopLocationMaliciousService{}
	r.Options = opts
	return
}

// Get the locations by emails classified as malicious or not.
func (r *RadarEmailSecurityTopLocationMaliciousService) Get(ctx context.Context, malicious RadarEmailSecurityTopLocationMaliciousGetParamsMalicious, query RadarEmailSecurityTopLocationMaliciousGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationMaliciousGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopLocationMaliciousGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/locations/malicious/%v", malicious)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopLocationMaliciousGetResponse struct {
	Meta RadarEmailSecurityTopLocationMaliciousGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationMaliciousGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationMaliciousGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationMaliciousGetResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationMaliciousGetResponse]
type radarEmailSecurityTopLocationMaliciousGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationMaliciousGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationMaliciousGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopLocationMaliciousGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationMaliciousGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationMaliciousGetResponseMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationMaliciousGetResponseMeta]
type radarEmailSecurityTopLocationMaliciousGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationMaliciousGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationMaliciousGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationMaliciousGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationMaliciousGetResponseMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopLocationMaliciousGetResponseMetaDateRange]
type radarEmailSecurityTopLocationMaliciousGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationMaliciousGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                           `json:"level"`
	JSON        radarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                            `json:"dataSource,required"`
	Description     string                                                                            `json:"description,required"`
	EventType       string                                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                                         `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationMaliciousGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationMaliciousGetResponseTop0 struct {
	ClientCountryAlpha2 string                                                    `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                    `json:"clientCountryName,required"`
	Value               string                                                    `json:"value,required"`
	JSON                radarEmailSecurityTopLocationMaliciousGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationMaliciousGetResponseTop0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationMaliciousGetResponseTop0]
type radarEmailSecurityTopLocationMaliciousGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationMaliciousGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationMaliciousGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopLocationMaliciousGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationMaliciousGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopLocationMaliciousGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopLocationMaliciousGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationMaliciousGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopLocationMaliciousGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationMaliciousGetParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationMaliciousGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Malicious.
type RadarEmailSecurityTopLocationMaliciousGetParamsMalicious string

const (
	RadarEmailSecurityTopLocationMaliciousGetParamsMaliciousMalicious    RadarEmailSecurityTopLocationMaliciousGetParamsMalicious = "MALICIOUS"
	RadarEmailSecurityTopLocationMaliciousGetParamsMaliciousNotMalicious RadarEmailSecurityTopLocationMaliciousGetParamsMalicious = "NOT_MALICIOUS"
)

type RadarEmailSecurityTopLocationMaliciousGetParamsArc string

const (
	RadarEmailSecurityTopLocationMaliciousGetParamsArcPass RadarEmailSecurityTopLocationMaliciousGetParamsArc = "PASS"
	RadarEmailSecurityTopLocationMaliciousGetParamsArcNone RadarEmailSecurityTopLocationMaliciousGetParamsArc = "NONE"
	RadarEmailSecurityTopLocationMaliciousGetParamsArcFail RadarEmailSecurityTopLocationMaliciousGetParamsArc = "FAIL"
)

type RadarEmailSecurityTopLocationMaliciousGetParamsDateRange string

const (
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange1d         RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "1d"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange2d         RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "2d"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange7d         RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "7d"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange14d        RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "14d"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange28d        RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "28d"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange12w        RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "12w"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange24w        RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "24w"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange52w        RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "52w"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange1dControl  RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange2dControl  RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange7dControl  RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange14dControl RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange28dControl RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange12wControl RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationMaliciousGetParamsDateRange24wControl RadarEmailSecurityTopLocationMaliciousGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationMaliciousGetParamsDKIM string

const (
	RadarEmailSecurityTopLocationMaliciousGetParamsDKIMPass RadarEmailSecurityTopLocationMaliciousGetParamsDKIM = "PASS"
	RadarEmailSecurityTopLocationMaliciousGetParamsDKIMNone RadarEmailSecurityTopLocationMaliciousGetParamsDKIM = "NONE"
	RadarEmailSecurityTopLocationMaliciousGetParamsDKIMFail RadarEmailSecurityTopLocationMaliciousGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopLocationMaliciousGetParamsDmarc string

const (
	RadarEmailSecurityTopLocationMaliciousGetParamsDmarcPass RadarEmailSecurityTopLocationMaliciousGetParamsDmarc = "PASS"
	RadarEmailSecurityTopLocationMaliciousGetParamsDmarcNone RadarEmailSecurityTopLocationMaliciousGetParamsDmarc = "NONE"
	RadarEmailSecurityTopLocationMaliciousGetParamsDmarcFail RadarEmailSecurityTopLocationMaliciousGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationMaliciousGetParamsFormat string

const (
	RadarEmailSecurityTopLocationMaliciousGetParamsFormatJson RadarEmailSecurityTopLocationMaliciousGetParamsFormat = "JSON"
	RadarEmailSecurityTopLocationMaliciousGetParamsFormatCsv  RadarEmailSecurityTopLocationMaliciousGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationMaliciousGetParamsSPF string

const (
	RadarEmailSecurityTopLocationMaliciousGetParamsSPFPass RadarEmailSecurityTopLocationMaliciousGetParamsSPF = "PASS"
	RadarEmailSecurityTopLocationMaliciousGetParamsSPFNone RadarEmailSecurityTopLocationMaliciousGetParamsSPF = "NONE"
	RadarEmailSecurityTopLocationMaliciousGetParamsSPFFail RadarEmailSecurityTopLocationMaliciousGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopLocationMaliciousGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopLocationMaliciousGetResponse             `json:"result,required"`
	Success bool                                                          `json:"success,required"`
	JSON    radarEmailSecurityTopLocationMaliciousGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopLocationMaliciousGetResponseEnvelopeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationMaliciousGetResponseEnvelope]
type radarEmailSecurityTopLocationMaliciousGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationMaliciousGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
