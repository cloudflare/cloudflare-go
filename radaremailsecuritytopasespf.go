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

// RadarEmailSecurityTopAseSpfService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseSpfService] method instead.
type RadarEmailSecurityTopAseSpfService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopAseSpfService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseSpfService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseSpfService) {
	r = &RadarEmailSecurityTopAseSpfService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS) by email SPF validation.
func (r *RadarEmailSecurityTopAseSpfService) Get(ctx context.Context, spf RadarEmailSecurityTopAseSpfGetParamsSpf, query RadarEmailSecurityTopAseSpfGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseSpfGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/ases/spf/%v", spf)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTopAseSpfGetResponse struct {
	Result  RadarEmailSecurityTopAseSpfGetResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarEmailSecurityTopAseSpfGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseSpfGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseSpfGetResponse]
type radarEmailSecurityTopAseSpfGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSpfGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpfGetResponseResult struct {
	Meta RadarEmailSecurityTopAseSpfGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseSpfGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseSpfGetResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopAseSpfGetResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseSpfGetResponseResult]
type radarEmailSecurityTopAseSpfGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSpfGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpfGetResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopAseSpfGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseSpfGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseSpfGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseSpfGetResponseResultMeta]
type radarEmailSecurityTopAseSpfGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSpfGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpfGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseSpfGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseSpfGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopAseSpfGetResponseResultMetaDateRange]
type radarEmailSecurityTopAseSpfGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSpfGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseSpfGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpfGetResponseResultTop0 struct {
	ClientASN    int64                                                `json:"clientASN,required"`
	ClientAsName string                                               `json:"clientASName,required"`
	Value        string                                               `json:"value,required"`
	JSON         radarEmailSecurityTopAseSpfGetResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseSpfGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseSpfGetResponseResultTop0]
type radarEmailSecurityTopAseSpfGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseSpfGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseSpfGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopAseSpfGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseSpfGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityTopAseSpfGetParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopAseSpfGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseSpfGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailSecurityTopAseSpfGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecurityTopAseSpfGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// SPF.
type RadarEmailSecurityTopAseSpfGetParamsSpf string

const (
	RadarEmailSecurityTopAseSpfGetParamsSpfPass RadarEmailSecurityTopAseSpfGetParamsSpf = "PASS"
	RadarEmailSecurityTopAseSpfGetParamsSpfNone RadarEmailSecurityTopAseSpfGetParamsSpf = "NONE"
	RadarEmailSecurityTopAseSpfGetParamsSpfFail RadarEmailSecurityTopAseSpfGetParamsSpf = "FAIL"
)

type RadarEmailSecurityTopAseSpfGetParamsArc string

const (
	RadarEmailSecurityTopAseSpfGetParamsArcPass RadarEmailSecurityTopAseSpfGetParamsArc = "PASS"
	RadarEmailSecurityTopAseSpfGetParamsArcNone RadarEmailSecurityTopAseSpfGetParamsArc = "NONE"
	RadarEmailSecurityTopAseSpfGetParamsArcFail RadarEmailSecurityTopAseSpfGetParamsArc = "FAIL"
)

type RadarEmailSecurityTopAseSpfGetParamsDateRange string

const (
	RadarEmailSecurityTopAseSpfGetParamsDateRange1d         RadarEmailSecurityTopAseSpfGetParamsDateRange = "1d"
	RadarEmailSecurityTopAseSpfGetParamsDateRange2d         RadarEmailSecurityTopAseSpfGetParamsDateRange = "2d"
	RadarEmailSecurityTopAseSpfGetParamsDateRange7d         RadarEmailSecurityTopAseSpfGetParamsDateRange = "7d"
	RadarEmailSecurityTopAseSpfGetParamsDateRange14d        RadarEmailSecurityTopAseSpfGetParamsDateRange = "14d"
	RadarEmailSecurityTopAseSpfGetParamsDateRange28d        RadarEmailSecurityTopAseSpfGetParamsDateRange = "28d"
	RadarEmailSecurityTopAseSpfGetParamsDateRange12w        RadarEmailSecurityTopAseSpfGetParamsDateRange = "12w"
	RadarEmailSecurityTopAseSpfGetParamsDateRange24w        RadarEmailSecurityTopAseSpfGetParamsDateRange = "24w"
	RadarEmailSecurityTopAseSpfGetParamsDateRange52w        RadarEmailSecurityTopAseSpfGetParamsDateRange = "52w"
	RadarEmailSecurityTopAseSpfGetParamsDateRange1dControl  RadarEmailSecurityTopAseSpfGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseSpfGetParamsDateRange2dControl  RadarEmailSecurityTopAseSpfGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseSpfGetParamsDateRange7dControl  RadarEmailSecurityTopAseSpfGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseSpfGetParamsDateRange14dControl RadarEmailSecurityTopAseSpfGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseSpfGetParamsDateRange28dControl RadarEmailSecurityTopAseSpfGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseSpfGetParamsDateRange12wControl RadarEmailSecurityTopAseSpfGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseSpfGetParamsDateRange24wControl RadarEmailSecurityTopAseSpfGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseSpfGetParamsDkim string

const (
	RadarEmailSecurityTopAseSpfGetParamsDkimPass RadarEmailSecurityTopAseSpfGetParamsDkim = "PASS"
	RadarEmailSecurityTopAseSpfGetParamsDkimNone RadarEmailSecurityTopAseSpfGetParamsDkim = "NONE"
	RadarEmailSecurityTopAseSpfGetParamsDkimFail RadarEmailSecurityTopAseSpfGetParamsDkim = "FAIL"
)

type RadarEmailSecurityTopAseSpfGetParamsDmarc string

const (
	RadarEmailSecurityTopAseSpfGetParamsDmarcPass RadarEmailSecurityTopAseSpfGetParamsDmarc = "PASS"
	RadarEmailSecurityTopAseSpfGetParamsDmarcNone RadarEmailSecurityTopAseSpfGetParamsDmarc = "NONE"
	RadarEmailSecurityTopAseSpfGetParamsDmarcFail RadarEmailSecurityTopAseSpfGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseSpfGetParamsFormat string

const (
	RadarEmailSecurityTopAseSpfGetParamsFormatJson RadarEmailSecurityTopAseSpfGetParamsFormat = "JSON"
	RadarEmailSecurityTopAseSpfGetParamsFormatCsv  RadarEmailSecurityTopAseSpfGetParamsFormat = "CSV"
)
