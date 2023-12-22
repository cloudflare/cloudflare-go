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

// RadarEmailTopAseSpfService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTopAseSpfService]
// method instead.
type RadarEmailTopAseSpfService struct {
	Options []option.RequestOption
}

// NewRadarEmailTopAseSpfService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTopAseSpfService(opts ...option.RequestOption) (r *RadarEmailTopAseSpfService) {
	r = &RadarEmailTopAseSpfService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by emails classified, of SPF validations.
func (r *RadarEmailTopAseSpfService) Get(ctx context.Context, spf RadarEmailTopAseSpfGetParamsSpf, query RadarEmailTopAseSpfGetParams, opts ...option.RequestOption) (res *RadarEmailTopAseSpfGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/top/ases/spf/%v", spf)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopAseSpfGetResponse struct {
	Result  RadarEmailTopAseSpfGetResponseResult `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarEmailTopAseSpfGetResponseJSON   `json:"-"`
}

// radarEmailTopAseSpfGetResponseJSON contains the JSON metadata for the struct
// [RadarEmailTopAseSpfGetResponse]
type radarEmailTopAseSpfGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseSpfGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpfGetResponseResult struct {
	Meta RadarEmailTopAseSpfGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopAseSpfGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopAseSpfGetResponseResultJSON   `json:"-"`
}

// radarEmailTopAseSpfGetResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTopAseSpfGetResponseResult]
type radarEmailTopAseSpfGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseSpfGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpfGetResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopAseSpfGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopAseSpfGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopAseSpfGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopAseSpfGetResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarEmailTopAseSpfGetResponseResultMeta]
type radarEmailTopAseSpfGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopAseSpfGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpfGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopAseSpfGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                              `json:"level,required"`
	JSON        radarEmailTopAseSpfGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopAseSpfGetResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarEmailTopAseSpfGetResponseResultMetaConfidenceInfo]
type radarEmailTopAseSpfGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseSpfGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpfGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                               `json:"dataSource,required"`
	Description string                                                               `json:"description,required"`
	EndTime     time.Time                                                            `json:"endTime,required" format:"date-time"`
	EventType   string                                                               `json:"eventType,required"`
	StartTime   time.Time                                                            `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopAseSpfGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopAseSpfGetResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailTopAseSpfGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopAseSpfGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseSpfGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpfGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                             `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopAseSpfGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopAseSpfGetResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarEmailTopAseSpfGetResponseResultMetaDateRange]
type radarEmailTopAseSpfGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseSpfGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpfGetResponseResultTop0 struct {
	ClientASN    int64                                        `json:"clientASN,required"`
	ClientAsName string                                       `json:"clientASName,required"`
	Value        string                                       `json:"value,required"`
	JSON         radarEmailTopAseSpfGetResponseResultTop0JSON `json:"-"`
}

// radarEmailTopAseSpfGetResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarEmailTopAseSpfGetResponseResultTop0]
type radarEmailTopAseSpfGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailTopAseSpfGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseSpfGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTopAseSpfGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopAseSpfGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTopAseSpfGetParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTopAseSpfGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopAseSpfGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailTopAseSpfGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopAseSpfGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopAseSpfGetParamsSpf string

const (
	RadarEmailTopAseSpfGetParamsSpfPass RadarEmailTopAseSpfGetParamsSpf = "PASS"
	RadarEmailTopAseSpfGetParamsSpfNone RadarEmailTopAseSpfGetParamsSpf = "NONE"
	RadarEmailTopAseSpfGetParamsSpfFail RadarEmailTopAseSpfGetParamsSpf = "FAIL"
)

type RadarEmailTopAseSpfGetParamsArc string

const (
	RadarEmailTopAseSpfGetParamsArcPass RadarEmailTopAseSpfGetParamsArc = "PASS"
	RadarEmailTopAseSpfGetParamsArcNone RadarEmailTopAseSpfGetParamsArc = "NONE"
	RadarEmailTopAseSpfGetParamsArcFail RadarEmailTopAseSpfGetParamsArc = "FAIL"
)

type RadarEmailTopAseSpfGetParamsDateRange string

const (
	RadarEmailTopAseSpfGetParamsDateRange1d         RadarEmailTopAseSpfGetParamsDateRange = "1d"
	RadarEmailTopAseSpfGetParamsDateRange7d         RadarEmailTopAseSpfGetParamsDateRange = "7d"
	RadarEmailTopAseSpfGetParamsDateRange14d        RadarEmailTopAseSpfGetParamsDateRange = "14d"
	RadarEmailTopAseSpfGetParamsDateRange28d        RadarEmailTopAseSpfGetParamsDateRange = "28d"
	RadarEmailTopAseSpfGetParamsDateRange12w        RadarEmailTopAseSpfGetParamsDateRange = "12w"
	RadarEmailTopAseSpfGetParamsDateRange24w        RadarEmailTopAseSpfGetParamsDateRange = "24w"
	RadarEmailTopAseSpfGetParamsDateRange52w        RadarEmailTopAseSpfGetParamsDateRange = "52w"
	RadarEmailTopAseSpfGetParamsDateRange1dControl  RadarEmailTopAseSpfGetParamsDateRange = "1dControl"
	RadarEmailTopAseSpfGetParamsDateRange7dControl  RadarEmailTopAseSpfGetParamsDateRange = "7dControl"
	RadarEmailTopAseSpfGetParamsDateRange14dControl RadarEmailTopAseSpfGetParamsDateRange = "14dControl"
	RadarEmailTopAseSpfGetParamsDateRange28dControl RadarEmailTopAseSpfGetParamsDateRange = "28dControl"
	RadarEmailTopAseSpfGetParamsDateRange12wControl RadarEmailTopAseSpfGetParamsDateRange = "12wControl"
	RadarEmailTopAseSpfGetParamsDateRange24wControl RadarEmailTopAseSpfGetParamsDateRange = "24wControl"
)

type RadarEmailTopAseSpfGetParamsDkim string

const (
	RadarEmailTopAseSpfGetParamsDkimPass RadarEmailTopAseSpfGetParamsDkim = "PASS"
	RadarEmailTopAseSpfGetParamsDkimNone RadarEmailTopAseSpfGetParamsDkim = "NONE"
	RadarEmailTopAseSpfGetParamsDkimFail RadarEmailTopAseSpfGetParamsDkim = "FAIL"
)

type RadarEmailTopAseSpfGetParamsDmarc string

const (
	RadarEmailTopAseSpfGetParamsDmarcPass RadarEmailTopAseSpfGetParamsDmarc = "PASS"
	RadarEmailTopAseSpfGetParamsDmarcNone RadarEmailTopAseSpfGetParamsDmarc = "NONE"
	RadarEmailTopAseSpfGetParamsDmarcFail RadarEmailTopAseSpfGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTopAseSpfGetParamsFormat string

const (
	RadarEmailTopAseSpfGetParamsFormatJson RadarEmailTopAseSpfGetParamsFormat = "JSON"
	RadarEmailTopAseSpfGetParamsFormatCsv  RadarEmailTopAseSpfGetParamsFormat = "CSV"
)
