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
	path := fmt.Sprintf("radar/email/security/top/ases/malicious/%v", malicious)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTopAseMaliciousGetResponse struct {
	Result  RadarEmailSecurityTopAseMaliciousGetResponseResult `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarEmailSecurityTopAseMaliciousGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseMaliciousGetResponse]
type radarEmailSecurityTopAseMaliciousGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseMaliciousGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetResponseResult struct {
	Meta RadarEmailSecurityTopAseMaliciousGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseMaliciousGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseMaliciousGetResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseMaliciousGetResponseResult]
type radarEmailSecurityTopAseMaliciousGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseMaliciousGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopAseMaliciousGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                               `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseMaliciousGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseMaliciousGetResponseResultMeta]
type radarEmailSecurityTopAseMaliciousGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseMaliciousGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                           `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseMaliciousGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopAseMaliciousGetResponseResultMetaDateRange]
type radarEmailSecurityTopAseMaliciousGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseMaliciousGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                            `json:"level"`
	JSON        radarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                             `json:"dataSource,required"`
	Description     string                                                                             `json:"description,required"`
	EventType       string                                                                             `json:"eventType,required"`
	IsInstantaneous interface{}                                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                          `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                             `json:"linkedUrl"`
	StartTime       time.Time                                                                          `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseMaliciousGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetResponseResultTop0 struct {
	ClientASN    int64                                                      `json:"clientASN,required"`
	ClientAsName string                                                     `json:"clientASName,required"`
	Value        string                                                     `json:"value,required"`
	JSON         radarEmailSecurityTopAseMaliciousGetResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseMaliciousGetResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopAseMaliciousGetResponseResultTop0]
type radarEmailSecurityTopAseMaliciousGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseMaliciousGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseMaliciousGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopAseMaliciousGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseMaliciousGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityTopAseMaliciousGetParamsDkim] `query:"dkim"`
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
	Spf param.Field[[]RadarEmailSecurityTopAseMaliciousGetParamsSpf] `query:"spf"`
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

type RadarEmailSecurityTopAseMaliciousGetParamsDkim string

const (
	RadarEmailSecurityTopAseMaliciousGetParamsDkimPass RadarEmailSecurityTopAseMaliciousGetParamsDkim = "PASS"
	RadarEmailSecurityTopAseMaliciousGetParamsDkimNone RadarEmailSecurityTopAseMaliciousGetParamsDkim = "NONE"
	RadarEmailSecurityTopAseMaliciousGetParamsDkimFail RadarEmailSecurityTopAseMaliciousGetParamsDkim = "FAIL"
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

type RadarEmailSecurityTopAseMaliciousGetParamsSpf string

const (
	RadarEmailSecurityTopAseMaliciousGetParamsSpfPass RadarEmailSecurityTopAseMaliciousGetParamsSpf = "PASS"
	RadarEmailSecurityTopAseMaliciousGetParamsSpfNone RadarEmailSecurityTopAseMaliciousGetParamsSpf = "NONE"
	RadarEmailSecurityTopAseMaliciousGetParamsSpfFail RadarEmailSecurityTopAseMaliciousGetParamsSpf = "FAIL"
)
