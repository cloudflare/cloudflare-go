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

// RadarEmailSecurityTopAseService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopAseService] method instead.
type RadarEmailSecurityTopAseService struct {
	Options   []option.RequestOption
	Arc       *RadarEmailSecurityTopAseArcService
	Dkim      *RadarEmailSecurityTopAseDkimService
	Dmarc     *RadarEmailSecurityTopAseDmarcService
	Malicious *RadarEmailSecurityTopAseMaliciousService
	Spam      *RadarEmailSecurityTopAseSpamService
	Spf       *RadarEmailSecurityTopAseSpfService
}

// NewRadarEmailSecurityTopAseService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopAseService(opts ...option.RequestOption) (r *RadarEmailSecurityTopAseService) {
	r = &RadarEmailSecurityTopAseService{}
	r.Options = opts
	r.Arc = NewRadarEmailSecurityTopAseArcService(opts...)
	r.Dkim = NewRadarEmailSecurityTopAseDkimService(opts...)
	r.Dmarc = NewRadarEmailSecurityTopAseDmarcService(opts...)
	r.Malicious = NewRadarEmailSecurityTopAseMaliciousService(opts...)
	r.Spam = NewRadarEmailSecurityTopAseSpamService(opts...)
	r.Spf = NewRadarEmailSecurityTopAseSpfService(opts...)
	return
}

// Get the top autonomous systems (AS) by email messages. Values are a percentage
// out of the total emails.
func (r *RadarEmailSecurityTopAseService) List(ctx context.Context, query RadarEmailSecurityTopAseListParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopAseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTopAseListResponse struct {
	Result  RadarEmailSecurityTopAseListResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarEmailSecurityTopAseListResponseJSON   `json:"-"`
}

// radarEmailSecurityTopAseListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopAseListResponse]
type radarEmailSecurityTopAseListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListResponseResult struct {
	Meta RadarEmailSecurityTopAseListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopAseListResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopAseListResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopAseListResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopAseListResponseResult]
type radarEmailSecurityTopAseListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopAseListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopAseListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopAseListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopAseListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseListResponseResultMeta]
type radarEmailSecurityTopAseListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopAseListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopAseListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopAseListResponseResultMetaDateRange]
type radarEmailSecurityTopAseListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopAseListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarEmailSecurityTopAseListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopAseListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopAseListResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopAseListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopAseListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopAseListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopAseListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopAseListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopAseListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListResponseResultTop0 struct {
	ClientASN    int64                                              `json:"clientASN,required"`
	ClientAsName string                                             `json:"clientASName,required"`
	Value        string                                             `json:"value,required"`
	JSON         radarEmailSecurityTopAseListResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopAseListResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopAseListResponseResultTop0]
type radarEmailSecurityTopAseListResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarEmailSecurityTopAseListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopAseListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopAseListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopAseListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityTopAseListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopAseListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopAseListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityTopAseListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopAseListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailSecurityTopAseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecurityTopAseListParamsArc string

const (
	RadarEmailSecurityTopAseListParamsArcPass RadarEmailSecurityTopAseListParamsArc = "PASS"
	RadarEmailSecurityTopAseListParamsArcNone RadarEmailSecurityTopAseListParamsArc = "NONE"
	RadarEmailSecurityTopAseListParamsArcFail RadarEmailSecurityTopAseListParamsArc = "FAIL"
)

type RadarEmailSecurityTopAseListParamsDateRange string

const (
	RadarEmailSecurityTopAseListParamsDateRange1d         RadarEmailSecurityTopAseListParamsDateRange = "1d"
	RadarEmailSecurityTopAseListParamsDateRange2d         RadarEmailSecurityTopAseListParamsDateRange = "2d"
	RadarEmailSecurityTopAseListParamsDateRange7d         RadarEmailSecurityTopAseListParamsDateRange = "7d"
	RadarEmailSecurityTopAseListParamsDateRange14d        RadarEmailSecurityTopAseListParamsDateRange = "14d"
	RadarEmailSecurityTopAseListParamsDateRange28d        RadarEmailSecurityTopAseListParamsDateRange = "28d"
	RadarEmailSecurityTopAseListParamsDateRange12w        RadarEmailSecurityTopAseListParamsDateRange = "12w"
	RadarEmailSecurityTopAseListParamsDateRange24w        RadarEmailSecurityTopAseListParamsDateRange = "24w"
	RadarEmailSecurityTopAseListParamsDateRange52w        RadarEmailSecurityTopAseListParamsDateRange = "52w"
	RadarEmailSecurityTopAseListParamsDateRange1dControl  RadarEmailSecurityTopAseListParamsDateRange = "1dControl"
	RadarEmailSecurityTopAseListParamsDateRange2dControl  RadarEmailSecurityTopAseListParamsDateRange = "2dControl"
	RadarEmailSecurityTopAseListParamsDateRange7dControl  RadarEmailSecurityTopAseListParamsDateRange = "7dControl"
	RadarEmailSecurityTopAseListParamsDateRange14dControl RadarEmailSecurityTopAseListParamsDateRange = "14dControl"
	RadarEmailSecurityTopAseListParamsDateRange28dControl RadarEmailSecurityTopAseListParamsDateRange = "28dControl"
	RadarEmailSecurityTopAseListParamsDateRange12wControl RadarEmailSecurityTopAseListParamsDateRange = "12wControl"
	RadarEmailSecurityTopAseListParamsDateRange24wControl RadarEmailSecurityTopAseListParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopAseListParamsDkim string

const (
	RadarEmailSecurityTopAseListParamsDkimPass RadarEmailSecurityTopAseListParamsDkim = "PASS"
	RadarEmailSecurityTopAseListParamsDkimNone RadarEmailSecurityTopAseListParamsDkim = "NONE"
	RadarEmailSecurityTopAseListParamsDkimFail RadarEmailSecurityTopAseListParamsDkim = "FAIL"
)

type RadarEmailSecurityTopAseListParamsDmarc string

const (
	RadarEmailSecurityTopAseListParamsDmarcPass RadarEmailSecurityTopAseListParamsDmarc = "PASS"
	RadarEmailSecurityTopAseListParamsDmarcNone RadarEmailSecurityTopAseListParamsDmarc = "NONE"
	RadarEmailSecurityTopAseListParamsDmarcFail RadarEmailSecurityTopAseListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopAseListParamsFormat string

const (
	RadarEmailSecurityTopAseListParamsFormatJson RadarEmailSecurityTopAseListParamsFormat = "JSON"
	RadarEmailSecurityTopAseListParamsFormatCsv  RadarEmailSecurityTopAseListParamsFormat = "CSV"
)

type RadarEmailSecurityTopAseListParamsSpf string

const (
	RadarEmailSecurityTopAseListParamsSpfPass RadarEmailSecurityTopAseListParamsSpf = "PASS"
	RadarEmailSecurityTopAseListParamsSpfNone RadarEmailSecurityTopAseListParamsSpf = "NONE"
	RadarEmailSecurityTopAseListParamsSpfFail RadarEmailSecurityTopAseListParamsSpf = "FAIL"
)
