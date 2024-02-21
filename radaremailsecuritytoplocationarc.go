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

// RadarEmailSecurityTopLocationArcService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationArcService] method instead.
type RadarEmailSecurityTopLocationArcService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationArcService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopLocationArcService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationArcService) {
	r = &RadarEmailSecurityTopLocationArcService{}
	r.Options = opts
	return
}

// Get the locations, by emails ARC validation.
func (r *RadarEmailSecurityTopLocationArcService) Get(ctx context.Context, arc RadarEmailSecurityTopLocationArcGetParamsArc, query RadarEmailSecurityTopLocationArcGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationArcGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopLocationArcGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/locations/arc/%v", arc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopLocationArcGetResponse struct {
	Meta RadarEmailSecurityTopLocationArcGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationArcGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationArcGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationArcGetResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopLocationArcGetResponse]
type radarEmailSecurityTopLocationArcGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationArcGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationArcGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopLocationArcGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationArcGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationArcGetResponseMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationArcGetResponseMeta]
type radarEmailSecurityTopLocationArcGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationArcGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationArcGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationArcGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationArcGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationArcGetResponseMetaDateRange]
type radarEmailSecurityTopLocationArcGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationArcGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationArcGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationArcGetResponseTop0 struct {
	ClientCountryAlpha2 string                                              `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                              `json:"clientCountryName,required"`
	Value               string                                              `json:"value,required"`
	JSON                radarEmailSecurityTopLocationArcGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationArcGetResponseTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationArcGetResponseTop0]
type radarEmailSecurityTopLocationArcGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationArcGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationArcGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationArcGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopLocationArcGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopLocationArcGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationArcGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopLocationArcGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationArcGetParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationArcGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// ARC.
type RadarEmailSecurityTopLocationArcGetParamsArc string

const (
	RadarEmailSecurityTopLocationArcGetParamsArcPass RadarEmailSecurityTopLocationArcGetParamsArc = "PASS"
	RadarEmailSecurityTopLocationArcGetParamsArcNone RadarEmailSecurityTopLocationArcGetParamsArc = "NONE"
	RadarEmailSecurityTopLocationArcGetParamsArcFail RadarEmailSecurityTopLocationArcGetParamsArc = "FAIL"
)

type RadarEmailSecurityTopLocationArcGetParamsDateRange string

const (
	RadarEmailSecurityTopLocationArcGetParamsDateRange1d         RadarEmailSecurityTopLocationArcGetParamsDateRange = "1d"
	RadarEmailSecurityTopLocationArcGetParamsDateRange2d         RadarEmailSecurityTopLocationArcGetParamsDateRange = "2d"
	RadarEmailSecurityTopLocationArcGetParamsDateRange7d         RadarEmailSecurityTopLocationArcGetParamsDateRange = "7d"
	RadarEmailSecurityTopLocationArcGetParamsDateRange14d        RadarEmailSecurityTopLocationArcGetParamsDateRange = "14d"
	RadarEmailSecurityTopLocationArcGetParamsDateRange28d        RadarEmailSecurityTopLocationArcGetParamsDateRange = "28d"
	RadarEmailSecurityTopLocationArcGetParamsDateRange12w        RadarEmailSecurityTopLocationArcGetParamsDateRange = "12w"
	RadarEmailSecurityTopLocationArcGetParamsDateRange24w        RadarEmailSecurityTopLocationArcGetParamsDateRange = "24w"
	RadarEmailSecurityTopLocationArcGetParamsDateRange52w        RadarEmailSecurityTopLocationArcGetParamsDateRange = "52w"
	RadarEmailSecurityTopLocationArcGetParamsDateRange1dControl  RadarEmailSecurityTopLocationArcGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationArcGetParamsDateRange2dControl  RadarEmailSecurityTopLocationArcGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationArcGetParamsDateRange7dControl  RadarEmailSecurityTopLocationArcGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationArcGetParamsDateRange14dControl RadarEmailSecurityTopLocationArcGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationArcGetParamsDateRange28dControl RadarEmailSecurityTopLocationArcGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationArcGetParamsDateRange12wControl RadarEmailSecurityTopLocationArcGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationArcGetParamsDateRange24wControl RadarEmailSecurityTopLocationArcGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationArcGetParamsDKIM string

const (
	RadarEmailSecurityTopLocationArcGetParamsDKIMPass RadarEmailSecurityTopLocationArcGetParamsDKIM = "PASS"
	RadarEmailSecurityTopLocationArcGetParamsDKIMNone RadarEmailSecurityTopLocationArcGetParamsDKIM = "NONE"
	RadarEmailSecurityTopLocationArcGetParamsDKIMFail RadarEmailSecurityTopLocationArcGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopLocationArcGetParamsDmarc string

const (
	RadarEmailSecurityTopLocationArcGetParamsDmarcPass RadarEmailSecurityTopLocationArcGetParamsDmarc = "PASS"
	RadarEmailSecurityTopLocationArcGetParamsDmarcNone RadarEmailSecurityTopLocationArcGetParamsDmarc = "NONE"
	RadarEmailSecurityTopLocationArcGetParamsDmarcFail RadarEmailSecurityTopLocationArcGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationArcGetParamsFormat string

const (
	RadarEmailSecurityTopLocationArcGetParamsFormatJson RadarEmailSecurityTopLocationArcGetParamsFormat = "JSON"
	RadarEmailSecurityTopLocationArcGetParamsFormatCsv  RadarEmailSecurityTopLocationArcGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationArcGetParamsSPF string

const (
	RadarEmailSecurityTopLocationArcGetParamsSPFPass RadarEmailSecurityTopLocationArcGetParamsSPF = "PASS"
	RadarEmailSecurityTopLocationArcGetParamsSPFNone RadarEmailSecurityTopLocationArcGetParamsSPF = "NONE"
	RadarEmailSecurityTopLocationArcGetParamsSPFFail RadarEmailSecurityTopLocationArcGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopLocationArcGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopLocationArcGetResponse             `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarEmailSecurityTopLocationArcGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopLocationArcGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationArcGetResponseEnvelope]
type radarEmailSecurityTopLocationArcGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationArcGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
