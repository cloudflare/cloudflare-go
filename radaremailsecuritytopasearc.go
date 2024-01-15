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
	path := fmt.Sprintf("radar/email/security/top/ases/arc/%v", arc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTopAseArcGetResponse struct {
	Result  RadarEmailSecurityTopAseArcGetResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarEmailSecurityTopAseArcGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseArcGetResponse]
type radarEmailSecurityTopAseArcGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseArcGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetResponseResult struct {
	Meta RadarEmailSecurityTopAseArcGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseArcGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseArcGetResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseArcGetResponseResult]
type radarEmailSecurityTopAseArcGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseArcGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopAseArcGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseArcGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseArcGetResponseResultMeta]
type radarEmailSecurityTopAseArcGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseArcGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseArcGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopAseArcGetResponseResultMetaDateRange]
type radarEmailSecurityTopAseArcGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseArcGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseArcGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetResponseResultTop0 struct {
	ClientASN    int64                                                `json:"clientASN,required"`
	ClientAsName string                                               `json:"clientASName,required"`
	Value        string                                               `json:"value,required"`
	JSON         radarEmailSecurityTopAseArcGetResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseArcGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseArcGetResponseResultTop0]
type radarEmailSecurityTopAseArcGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseArcGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseArcGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseArcGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityTopAseArcGetParamsDkim] `query:"dkim"`
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
	Spf param.Field[[]RadarEmailSecurityTopAseArcGetParamsSpf] `query:"spf"`
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

type RadarEmailSecurityTopAseArcGetParamsDkim string

const (
	RadarEmailSecurityTopAseArcGetParamsDkimPass RadarEmailSecurityTopAseArcGetParamsDkim = "PASS"
	RadarEmailSecurityTopAseArcGetParamsDkimNone RadarEmailSecurityTopAseArcGetParamsDkim = "NONE"
	RadarEmailSecurityTopAseArcGetParamsDkimFail RadarEmailSecurityTopAseArcGetParamsDkim = "FAIL"
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

type RadarEmailSecurityTopAseArcGetParamsSpf string

const (
	RadarEmailSecurityTopAseArcGetParamsSpfPass RadarEmailSecurityTopAseArcGetParamsSpf = "PASS"
	RadarEmailSecurityTopAseArcGetParamsSpfNone RadarEmailSecurityTopAseArcGetParamsSpf = "NONE"
	RadarEmailSecurityTopAseArcGetParamsSpfFail RadarEmailSecurityTopAseArcGetParamsSpf = "FAIL"
)
