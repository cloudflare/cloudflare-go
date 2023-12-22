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

// RadarEmailTopAseDkimService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTopAseDkimService]
// method instead.
type RadarEmailTopAseDkimService struct {
	Options []option.RequestOption
}

// NewRadarEmailTopAseDkimService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTopAseDkimService(opts ...option.RequestOption) (r *RadarEmailTopAseDkimService) {
	r = &RadarEmailTopAseDkimService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by emails classified, of DKIM validations.
func (r *RadarEmailTopAseDkimService) Get(ctx context.Context, dkim RadarEmailTopAseDkimGetParamsDkim, query RadarEmailTopAseDkimGetParams, opts ...option.RequestOption) (res *RadarEmailTopAseDkimGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/top/ases/dkim/%v", dkim)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopAseDkimGetResponse struct {
	Result  RadarEmailTopAseDkimGetResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarEmailTopAseDkimGetResponseJSON   `json:"-"`
}

// radarEmailTopAseDkimGetResponseJSON contains the JSON metadata for the struct
// [RadarEmailTopAseDkimGetResponse]
type radarEmailTopAseDkimGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseDkimGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDkimGetResponseResult struct {
	Meta RadarEmailTopAseDkimGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopAseDkimGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopAseDkimGetResponseResultJSON   `json:"-"`
}

// radarEmailTopAseDkimGetResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTopAseDkimGetResponseResult]
type radarEmailTopAseDkimGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseDkimGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDkimGetResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopAseDkimGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopAseDkimGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopAseDkimGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopAseDkimGetResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarEmailTopAseDkimGetResponseResultMeta]
type radarEmailTopAseDkimGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopAseDkimGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDkimGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopAseDkimGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                               `json:"level,required"`
	JSON        radarEmailTopAseDkimGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopAseDkimGetResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailTopAseDkimGetResponseResultMetaConfidenceInfo]
type radarEmailTopAseDkimGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseDkimGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDkimGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                `json:"dataSource,required"`
	Description string                                                                `json:"description,required"`
	EndTime     time.Time                                                             `json:"endTime,required" format:"date-time"`
	EventType   string                                                                `json:"eventType,required"`
	StartTime   time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopAseDkimGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopAseDkimGetResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailTopAseDkimGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopAseDkimGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseDkimGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDkimGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                              `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopAseDkimGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopAseDkimGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailTopAseDkimGetResponseResultMetaDateRange]
type radarEmailTopAseDkimGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseDkimGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDkimGetResponseResultTop0 struct {
	ClientASN    int64                                         `json:"clientASN,required"`
	ClientAsName string                                        `json:"clientASName,required"`
	Value        string                                        `json:"value,required"`
	JSON         radarEmailTopAseDkimGetResponseResultTop0JSON `json:"-"`
}

// radarEmailTopAseDkimGetResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarEmailTopAseDkimGetResponseResultTop0]
type radarEmailTopAseDkimGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailTopAseDkimGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseDkimGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTopAseDkimGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopAseDkimGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTopAseDkimGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopAseDkimGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTopAseDkimGetParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTopAseDkimGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopAseDkimGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopAseDkimGetParamsDkim string

const (
	RadarEmailTopAseDkimGetParamsDkimPass RadarEmailTopAseDkimGetParamsDkim = "PASS"
	RadarEmailTopAseDkimGetParamsDkimNone RadarEmailTopAseDkimGetParamsDkim = "NONE"
	RadarEmailTopAseDkimGetParamsDkimFail RadarEmailTopAseDkimGetParamsDkim = "FAIL"
)

type RadarEmailTopAseDkimGetParamsArc string

const (
	RadarEmailTopAseDkimGetParamsArcPass RadarEmailTopAseDkimGetParamsArc = "PASS"
	RadarEmailTopAseDkimGetParamsArcNone RadarEmailTopAseDkimGetParamsArc = "NONE"
	RadarEmailTopAseDkimGetParamsArcFail RadarEmailTopAseDkimGetParamsArc = "FAIL"
)

type RadarEmailTopAseDkimGetParamsDateRange string

const (
	RadarEmailTopAseDkimGetParamsDateRange1d         RadarEmailTopAseDkimGetParamsDateRange = "1d"
	RadarEmailTopAseDkimGetParamsDateRange7d         RadarEmailTopAseDkimGetParamsDateRange = "7d"
	RadarEmailTopAseDkimGetParamsDateRange14d        RadarEmailTopAseDkimGetParamsDateRange = "14d"
	RadarEmailTopAseDkimGetParamsDateRange28d        RadarEmailTopAseDkimGetParamsDateRange = "28d"
	RadarEmailTopAseDkimGetParamsDateRange12w        RadarEmailTopAseDkimGetParamsDateRange = "12w"
	RadarEmailTopAseDkimGetParamsDateRange24w        RadarEmailTopAseDkimGetParamsDateRange = "24w"
	RadarEmailTopAseDkimGetParamsDateRange52w        RadarEmailTopAseDkimGetParamsDateRange = "52w"
	RadarEmailTopAseDkimGetParamsDateRange1dControl  RadarEmailTopAseDkimGetParamsDateRange = "1dControl"
	RadarEmailTopAseDkimGetParamsDateRange7dControl  RadarEmailTopAseDkimGetParamsDateRange = "7dControl"
	RadarEmailTopAseDkimGetParamsDateRange14dControl RadarEmailTopAseDkimGetParamsDateRange = "14dControl"
	RadarEmailTopAseDkimGetParamsDateRange28dControl RadarEmailTopAseDkimGetParamsDateRange = "28dControl"
	RadarEmailTopAseDkimGetParamsDateRange12wControl RadarEmailTopAseDkimGetParamsDateRange = "12wControl"
	RadarEmailTopAseDkimGetParamsDateRange24wControl RadarEmailTopAseDkimGetParamsDateRange = "24wControl"
)

type RadarEmailTopAseDkimGetParamsDmarc string

const (
	RadarEmailTopAseDkimGetParamsDmarcPass RadarEmailTopAseDkimGetParamsDmarc = "PASS"
	RadarEmailTopAseDkimGetParamsDmarcNone RadarEmailTopAseDkimGetParamsDmarc = "NONE"
	RadarEmailTopAseDkimGetParamsDmarcFail RadarEmailTopAseDkimGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTopAseDkimGetParamsFormat string

const (
	RadarEmailTopAseDkimGetParamsFormatJson RadarEmailTopAseDkimGetParamsFormat = "JSON"
	RadarEmailTopAseDkimGetParamsFormatCsv  RadarEmailTopAseDkimGetParamsFormat = "CSV"
)

type RadarEmailTopAseDkimGetParamsSpf string

const (
	RadarEmailTopAseDkimGetParamsSpfPass RadarEmailTopAseDkimGetParamsSpf = "PASS"
	RadarEmailTopAseDkimGetParamsSpfNone RadarEmailTopAseDkimGetParamsSpf = "NONE"
	RadarEmailTopAseDkimGetParamsSpfFail RadarEmailTopAseDkimGetParamsSpf = "FAIL"
)
