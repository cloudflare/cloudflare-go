// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarEmailTopAseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTopAseService] method
// instead.
type RadarEmailTopAseService struct {
	Options   []option.RequestOption
	Arcs      *RadarEmailTopAseArcService
	Dkims     *RadarEmailTopAseDkimService
	Dmarcs    *RadarEmailTopAseDmarcService
	Malicious *RadarEmailTopAseMaliciousService
	Spams     *RadarEmailTopAseSpamService
	Spfs      *RadarEmailTopAseSpfService
}

// NewRadarEmailTopAseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTopAseService(opts ...option.RequestOption) (r *RadarEmailTopAseService) {
	r = &RadarEmailTopAseService{}
	r.Options = opts
	r.Arcs = NewRadarEmailTopAseArcService(opts...)
	r.Dkims = NewRadarEmailTopAseDkimService(opts...)
	r.Dmarcs = NewRadarEmailTopAseDmarcService(opts...)
	r.Malicious = NewRadarEmailTopAseMaliciousService(opts...)
	r.Spams = NewRadarEmailTopAseSpamService(opts...)
	r.Spfs = NewRadarEmailTopAseSpfService(opts...)
	return
}

// Get the top autonomous systems (AS) by HTTP traffic. Values are a percentage out
// of the total traffic.
func (r *RadarEmailTopAseService) List(ctx context.Context, query RadarEmailTopAseListParams, opts ...option.RequestOption) (res *RadarEmailTopAseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopAseListResponse struct {
	Result  RadarEmailTopAseListResponseResult `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    radarEmailTopAseListResponseJSON   `json:"-"`
}

// radarEmailTopAseListResponseJSON contains the JSON metadata for the struct
// [RadarEmailTopAseListResponse]
type radarEmailTopAseListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseListResponseResult struct {
	Meta RadarEmailTopAseListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopAseListResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopAseListResponseResultJSON   `json:"-"`
}

// radarEmailTopAseListResponseResultJSON contains the JSON metadata for the struct
// [RadarEmailTopAseListResponseResult]
type radarEmailTopAseListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseListResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopAseListResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopAseListResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopAseListResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopAseListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarEmailTopAseListResponseResultMeta]
type radarEmailTopAseListResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopAseListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopAseListResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                            `json:"level,required"`
	JSON        radarEmailTopAseListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopAseListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarEmailTopAseListResponseResultMetaConfidenceInfo]
type radarEmailTopAseListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                             `json:"dataSource,required"`
	Description string                                                             `json:"description,required"`
	EndTime     time.Time                                                          `json:"endTime,required" format:"date-time"`
	EventType   string                                                             `json:"eventType,required"`
	StartTime   time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopAseListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopAseListResponseResultMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarEmailTopAseListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopAseListResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseListResponseResultMetaDateRange struct {
	EndTime   time.Time                                           `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                           `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopAseListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopAseListResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarEmailTopAseListResponseResultMetaDateRange]
type radarEmailTopAseListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopAseListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseListResponseResultTop0 struct {
	ClientASN    int64                                      `json:"clientASN,required"`
	ClientAsName string                                     `json:"clientASName,required"`
	Value        string                                     `json:"value,required"`
	JSON         radarEmailTopAseListResponseResultTop0JSON `json:"-"`
}

// radarEmailTopAseListResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarEmailTopAseListResponseResultTop0]
type radarEmailTopAseListResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailTopAseListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopAseListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTopAseListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopAseListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTopAseListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTopAseListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopAseListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTopAseListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTopAseListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopAseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopAseListParamsArc string

const (
	RadarEmailTopAseListParamsArcPass RadarEmailTopAseListParamsArc = "PASS"
	RadarEmailTopAseListParamsArcNone RadarEmailTopAseListParamsArc = "NONE"
	RadarEmailTopAseListParamsArcFail RadarEmailTopAseListParamsArc = "FAIL"
)

type RadarEmailTopAseListParamsDateRange string

const (
	RadarEmailTopAseListParamsDateRange1d         RadarEmailTopAseListParamsDateRange = "1d"
	RadarEmailTopAseListParamsDateRange7d         RadarEmailTopAseListParamsDateRange = "7d"
	RadarEmailTopAseListParamsDateRange14d        RadarEmailTopAseListParamsDateRange = "14d"
	RadarEmailTopAseListParamsDateRange28d        RadarEmailTopAseListParamsDateRange = "28d"
	RadarEmailTopAseListParamsDateRange12w        RadarEmailTopAseListParamsDateRange = "12w"
	RadarEmailTopAseListParamsDateRange24w        RadarEmailTopAseListParamsDateRange = "24w"
	RadarEmailTopAseListParamsDateRange52w        RadarEmailTopAseListParamsDateRange = "52w"
	RadarEmailTopAseListParamsDateRange1dControl  RadarEmailTopAseListParamsDateRange = "1dControl"
	RadarEmailTopAseListParamsDateRange7dControl  RadarEmailTopAseListParamsDateRange = "7dControl"
	RadarEmailTopAseListParamsDateRange14dControl RadarEmailTopAseListParamsDateRange = "14dControl"
	RadarEmailTopAseListParamsDateRange28dControl RadarEmailTopAseListParamsDateRange = "28dControl"
	RadarEmailTopAseListParamsDateRange12wControl RadarEmailTopAseListParamsDateRange = "12wControl"
	RadarEmailTopAseListParamsDateRange24wControl RadarEmailTopAseListParamsDateRange = "24wControl"
)

type RadarEmailTopAseListParamsDkim string

const (
	RadarEmailTopAseListParamsDkimPass RadarEmailTopAseListParamsDkim = "PASS"
	RadarEmailTopAseListParamsDkimNone RadarEmailTopAseListParamsDkim = "NONE"
	RadarEmailTopAseListParamsDkimFail RadarEmailTopAseListParamsDkim = "FAIL"
)

type RadarEmailTopAseListParamsDmarc string

const (
	RadarEmailTopAseListParamsDmarcPass RadarEmailTopAseListParamsDmarc = "PASS"
	RadarEmailTopAseListParamsDmarcNone RadarEmailTopAseListParamsDmarc = "NONE"
	RadarEmailTopAseListParamsDmarcFail RadarEmailTopAseListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTopAseListParamsFormat string

const (
	RadarEmailTopAseListParamsFormatJson RadarEmailTopAseListParamsFormat = "JSON"
	RadarEmailTopAseListParamsFormatCsv  RadarEmailTopAseListParamsFormat = "CSV"
)

type RadarEmailTopAseListParamsSpf string

const (
	RadarEmailTopAseListParamsSpfPass RadarEmailTopAseListParamsSpf = "PASS"
	RadarEmailTopAseListParamsSpfNone RadarEmailTopAseListParamsSpf = "NONE"
	RadarEmailTopAseListParamsSpfFail RadarEmailTopAseListParamsSpf = "FAIL"
)
