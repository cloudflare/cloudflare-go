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

// RadarEmailSecurityTopLocationBySpfService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationBySpfService] method instead.
type RadarEmailSecurityTopLocationBySpfService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationBySpfService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityTopLocationBySpfService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationBySpfService) {
	r = &RadarEmailSecurityTopLocationBySpfService{}
	r.Options = opts
	return
}

// Get the top locations by email SPF validation.
func (r *RadarEmailSecurityTopLocationBySpfService) List(ctx context.Context, spf RadarEmailSecurityTopLocationBySpfListParamsSpf, query RadarEmailSecurityTopLocationBySpfListParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationBySpfListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/locations/spf/%v", spf)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTopLocationBySpfListResponse struct {
	Result  RadarEmailSecurityTopLocationBySpfListResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailSecurityTopLocationBySpfListResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationBySpfListResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationBySpfListResponse]
type radarEmailSecurityTopLocationBySpfListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationBySpfListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpfListResponseResult struct {
	Meta RadarEmailSecurityTopLocationBySpfListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationBySpfListResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationBySpfListResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopLocationBySpfListResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationBySpfListResponseResult]
type radarEmailSecurityTopLocationBySpfListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationBySpfListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpfListResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopLocationBySpfListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                 `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationBySpfListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationBySpfListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationBySpfListResponseResultMeta]
type radarEmailSecurityTopLocationBySpfListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationBySpfListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpfListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                             `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationBySpfListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationBySpfListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopLocationBySpfListResponseResultMetaDateRange]
type radarEmailSecurityTopLocationBySpfListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationBySpfListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                              `json:"level"`
	JSON        radarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                               `json:"dataSource,required"`
	Description     string                                                                               `json:"description,required"`
	EventType       string                                                                               `json:"eventType,required"`
	IsInstantaneous interface{}                                                                          `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                                            `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationBySpfListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpfListResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                       `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                       `json:"clientCountryName,required"`
	Value               string                                                       `json:"value,required"`
	JSON                radarEmailSecurityTopLocationBySpfListResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationBySpfListResponseResultTop0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationBySpfListResponseResultTop0]
type radarEmailSecurityTopLocationBySpfListResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationBySpfListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpfListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopLocationBySpfListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationBySpfListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityTopLocationBySpfListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopLocationBySpfListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationBySpfListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationBySpfListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationBySpfListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// SPF.
type RadarEmailSecurityTopLocationBySpfListParamsSpf string

const (
	RadarEmailSecurityTopLocationBySpfListParamsSpfPass RadarEmailSecurityTopLocationBySpfListParamsSpf = "PASS"
	RadarEmailSecurityTopLocationBySpfListParamsSpfNone RadarEmailSecurityTopLocationBySpfListParamsSpf = "NONE"
	RadarEmailSecurityTopLocationBySpfListParamsSpfFail RadarEmailSecurityTopLocationBySpfListParamsSpf = "FAIL"
)

type RadarEmailSecurityTopLocationBySpfListParamsArc string

const (
	RadarEmailSecurityTopLocationBySpfListParamsArcPass RadarEmailSecurityTopLocationBySpfListParamsArc = "PASS"
	RadarEmailSecurityTopLocationBySpfListParamsArcNone RadarEmailSecurityTopLocationBySpfListParamsArc = "NONE"
	RadarEmailSecurityTopLocationBySpfListParamsArcFail RadarEmailSecurityTopLocationBySpfListParamsArc = "FAIL"
)

type RadarEmailSecurityTopLocationBySpfListParamsDateRange string

const (
	RadarEmailSecurityTopLocationBySpfListParamsDateRange1d         RadarEmailSecurityTopLocationBySpfListParamsDateRange = "1d"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange2d         RadarEmailSecurityTopLocationBySpfListParamsDateRange = "2d"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange7d         RadarEmailSecurityTopLocationBySpfListParamsDateRange = "7d"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange14d        RadarEmailSecurityTopLocationBySpfListParamsDateRange = "14d"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange28d        RadarEmailSecurityTopLocationBySpfListParamsDateRange = "28d"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange12w        RadarEmailSecurityTopLocationBySpfListParamsDateRange = "12w"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange24w        RadarEmailSecurityTopLocationBySpfListParamsDateRange = "24w"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange52w        RadarEmailSecurityTopLocationBySpfListParamsDateRange = "52w"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange1dControl  RadarEmailSecurityTopLocationBySpfListParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange2dControl  RadarEmailSecurityTopLocationBySpfListParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange7dControl  RadarEmailSecurityTopLocationBySpfListParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange14dControl RadarEmailSecurityTopLocationBySpfListParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange28dControl RadarEmailSecurityTopLocationBySpfListParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange12wControl RadarEmailSecurityTopLocationBySpfListParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationBySpfListParamsDateRange24wControl RadarEmailSecurityTopLocationBySpfListParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationBySpfListParamsDkim string

const (
	RadarEmailSecurityTopLocationBySpfListParamsDkimPass RadarEmailSecurityTopLocationBySpfListParamsDkim = "PASS"
	RadarEmailSecurityTopLocationBySpfListParamsDkimNone RadarEmailSecurityTopLocationBySpfListParamsDkim = "NONE"
	RadarEmailSecurityTopLocationBySpfListParamsDkimFail RadarEmailSecurityTopLocationBySpfListParamsDkim = "FAIL"
)

type RadarEmailSecurityTopLocationBySpfListParamsDmarc string

const (
	RadarEmailSecurityTopLocationBySpfListParamsDmarcPass RadarEmailSecurityTopLocationBySpfListParamsDmarc = "PASS"
	RadarEmailSecurityTopLocationBySpfListParamsDmarcNone RadarEmailSecurityTopLocationBySpfListParamsDmarc = "NONE"
	RadarEmailSecurityTopLocationBySpfListParamsDmarcFail RadarEmailSecurityTopLocationBySpfListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationBySpfListParamsFormat string

const (
	RadarEmailSecurityTopLocationBySpfListParamsFormatJson RadarEmailSecurityTopLocationBySpfListParamsFormat = "JSON"
	RadarEmailSecurityTopLocationBySpfListParamsFormatCsv  RadarEmailSecurityTopLocationBySpfListParamsFormat = "CSV"
)
