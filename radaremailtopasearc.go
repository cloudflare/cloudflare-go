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

// RadarEmailTopAseArcService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTopAseArcService]
// method instead.
type RadarEmailTopAseArcService struct {
	Options []option.RequestOption
}

// NewRadarEmailTopAseArcService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTopAseArcService(opts ...option.RequestOption) (r *RadarEmailTopAseArcService) {
	r = &RadarEmailTopAseArcService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by emails classified, of ARC validations.
func (r *RadarEmailTopAseArcService) Get(ctx context.Context, arc RadarEmailTopAseArcGetParamsArc, query RadarEmailTopAseArcGetParams, opts ...option.RequestOption) (res *RadarEmailTopAseArcGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/top/ases/arc/%v", arc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopAseArcGetResponse struct {
	Result  RadarEmailTopAseArcGetResponseResult `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarEmailTopAseArcGetResponseJSON   `json:"-"`
}

// radarEmailTopAseArcGetResponseJSON contains the JSON metadata for the struct
// [RadarEmailTopAseArcGetResponse]
type radarEmailTopAseArcGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseArcGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseArcGetResponseResult struct {
	Meta RadarEmailTopAseArcGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopAseArcGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopAseArcGetResponseResultJSON   `json:"-"`
}

// radarEmailTopAseArcGetResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTopAseArcGetResponseResult]
type radarEmailTopAseArcGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseArcGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseArcGetResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopAseArcGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopAseArcGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopAseArcGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopAseArcGetResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarEmailTopAseArcGetResponseResultMeta]
type radarEmailTopAseArcGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopAseArcGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseArcGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopAseArcGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                              `json:"level,required"`
	JSON        radarEmailTopAseArcGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopAseArcGetResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarEmailTopAseArcGetResponseResultMetaConfidenceInfo]
type radarEmailTopAseArcGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseArcGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseArcGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                               `json:"dataSource,required"`
	Description string                                                               `json:"description,required"`
	EndTime     time.Time                                                            `json:"endTime,required" format:"date-time"`
	EventType   string                                                               `json:"eventType,required"`
	StartTime   time.Time                                                            `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopAseArcGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopAseArcGetResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailTopAseArcGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopAseArcGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseArcGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseArcGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                             `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopAseArcGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopAseArcGetResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarEmailTopAseArcGetResponseResultMetaDateRange]
type radarEmailTopAseArcGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseArcGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseArcGetResponseResultTop0 struct {
	ClientASN    int64                                        `json:"clientASN,required"`
	ClientAsName string                                       `json:"clientASName,required"`
	Value        string                                       `json:"value,required"`
	JSON         radarEmailTopAseArcGetResponseResultTop0JSON `json:"-"`
}

// radarEmailTopAseArcGetResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarEmailTopAseArcGetResponseResultTop0]
type radarEmailTopAseArcGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailTopAseArcGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseArcGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopAseArcGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTopAseArcGetParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTopAseArcGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopAseArcGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTopAseArcGetParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTopAseArcGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopAseArcGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopAseArcGetParamsArc string

const (
	RadarEmailTopAseArcGetParamsArcPass RadarEmailTopAseArcGetParamsArc = "PASS"
	RadarEmailTopAseArcGetParamsArcNone RadarEmailTopAseArcGetParamsArc = "NONE"
	RadarEmailTopAseArcGetParamsArcFail RadarEmailTopAseArcGetParamsArc = "FAIL"
)

type RadarEmailTopAseArcGetParamsDateRange string

const (
	RadarEmailTopAseArcGetParamsDateRange1d         RadarEmailTopAseArcGetParamsDateRange = "1d"
	RadarEmailTopAseArcGetParamsDateRange7d         RadarEmailTopAseArcGetParamsDateRange = "7d"
	RadarEmailTopAseArcGetParamsDateRange14d        RadarEmailTopAseArcGetParamsDateRange = "14d"
	RadarEmailTopAseArcGetParamsDateRange28d        RadarEmailTopAseArcGetParamsDateRange = "28d"
	RadarEmailTopAseArcGetParamsDateRange12w        RadarEmailTopAseArcGetParamsDateRange = "12w"
	RadarEmailTopAseArcGetParamsDateRange24w        RadarEmailTopAseArcGetParamsDateRange = "24w"
	RadarEmailTopAseArcGetParamsDateRange52w        RadarEmailTopAseArcGetParamsDateRange = "52w"
	RadarEmailTopAseArcGetParamsDateRange1dControl  RadarEmailTopAseArcGetParamsDateRange = "1dControl"
	RadarEmailTopAseArcGetParamsDateRange7dControl  RadarEmailTopAseArcGetParamsDateRange = "7dControl"
	RadarEmailTopAseArcGetParamsDateRange14dControl RadarEmailTopAseArcGetParamsDateRange = "14dControl"
	RadarEmailTopAseArcGetParamsDateRange28dControl RadarEmailTopAseArcGetParamsDateRange = "28dControl"
	RadarEmailTopAseArcGetParamsDateRange12wControl RadarEmailTopAseArcGetParamsDateRange = "12wControl"
	RadarEmailTopAseArcGetParamsDateRange24wControl RadarEmailTopAseArcGetParamsDateRange = "24wControl"
)

type RadarEmailTopAseArcGetParamsDkim string

const (
	RadarEmailTopAseArcGetParamsDkimPass RadarEmailTopAseArcGetParamsDkim = "PASS"
	RadarEmailTopAseArcGetParamsDkimNone RadarEmailTopAseArcGetParamsDkim = "NONE"
	RadarEmailTopAseArcGetParamsDkimFail RadarEmailTopAseArcGetParamsDkim = "FAIL"
)

type RadarEmailTopAseArcGetParamsDmarc string

const (
	RadarEmailTopAseArcGetParamsDmarcPass RadarEmailTopAseArcGetParamsDmarc = "PASS"
	RadarEmailTopAseArcGetParamsDmarcNone RadarEmailTopAseArcGetParamsDmarc = "NONE"
	RadarEmailTopAseArcGetParamsDmarcFail RadarEmailTopAseArcGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTopAseArcGetParamsFormat string

const (
	RadarEmailTopAseArcGetParamsFormatJson RadarEmailTopAseArcGetParamsFormat = "JSON"
	RadarEmailTopAseArcGetParamsFormatCsv  RadarEmailTopAseArcGetParamsFormat = "CSV"
)

type RadarEmailTopAseArcGetParamsSpf string

const (
	RadarEmailTopAseArcGetParamsSpfPass RadarEmailTopAseArcGetParamsSpf = "PASS"
	RadarEmailTopAseArcGetParamsSpfNone RadarEmailTopAseArcGetParamsSpf = "NONE"
	RadarEmailTopAseArcGetParamsSpfFail RadarEmailTopAseArcGetParamsSpf = "FAIL"
)
