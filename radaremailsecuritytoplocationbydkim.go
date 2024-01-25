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

// RadarEmailSecurityTopLocationByDkimService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationByDkimService] method instead.
type RadarEmailSecurityTopLocationByDkimService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationByDkimService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityTopLocationByDkimService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationByDkimService) {
	r = &RadarEmailSecurityTopLocationByDkimService{}
	r.Options = opts
	return
}

// Get the locations, by email DKIM validation.
func (r *RadarEmailSecurityTopLocationByDkimService) List(ctx context.Context, dkim RadarEmailSecurityTopLocationByDkimListParamsDkim, query RadarEmailSecurityTopLocationByDkimListParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationByDkimListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/locations/dkim/%v", dkim)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTopLocationByDkimListResponse struct {
	Result  RadarEmailSecurityTopLocationByDkimListResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarEmailSecurityTopLocationByDkimListResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationByDkimListResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationByDkimListResponse]
type radarEmailSecurityTopLocationByDkimListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByDkimListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDkimListResponseResult struct {
	Meta RadarEmailSecurityTopLocationByDkimListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationByDkimListResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationByDkimListResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopLocationByDkimListResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationByDkimListResponseResult]
type radarEmailSecurityTopLocationByDkimListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByDkimListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDkimListResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopLocationByDkimListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                  `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationByDkimListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationByDkimListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationByDkimListResponseResultMeta]
type radarEmailSecurityTopLocationByDkimListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByDkimListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDkimListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationByDkimListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationByDkimListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopLocationByDkimListResponseResultMetaDateRange]
type radarEmailSecurityTopLocationByDkimListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByDkimListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                               `json:"level"`
	JSON        radarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                `json:"dataSource,required"`
	Description     string                                                                                `json:"description,required"`
	EventType       string                                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                                             `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationByDkimListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDkimListResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                        `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                        `json:"clientCountryName,required"`
	Value               string                                                        `json:"value,required"`
	JSON                radarEmailSecurityTopLocationByDkimListResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationByDkimListResponseResultTop0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationByDkimListResponseResultTop0]
type radarEmailSecurityTopLocationByDkimListResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByDkimListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDkimListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopLocationByDkimListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationByDkimListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopLocationByDkimListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationByDkimListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityTopLocationByDkimListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationByDkimListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationByDkimListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DKIM.
type RadarEmailSecurityTopLocationByDkimListParamsDkim string

const (
	RadarEmailSecurityTopLocationByDkimListParamsDkimPass RadarEmailSecurityTopLocationByDkimListParamsDkim = "PASS"
	RadarEmailSecurityTopLocationByDkimListParamsDkimNone RadarEmailSecurityTopLocationByDkimListParamsDkim = "NONE"
	RadarEmailSecurityTopLocationByDkimListParamsDkimFail RadarEmailSecurityTopLocationByDkimListParamsDkim = "FAIL"
)

type RadarEmailSecurityTopLocationByDkimListParamsArc string

const (
	RadarEmailSecurityTopLocationByDkimListParamsArcPass RadarEmailSecurityTopLocationByDkimListParamsArc = "PASS"
	RadarEmailSecurityTopLocationByDkimListParamsArcNone RadarEmailSecurityTopLocationByDkimListParamsArc = "NONE"
	RadarEmailSecurityTopLocationByDkimListParamsArcFail RadarEmailSecurityTopLocationByDkimListParamsArc = "FAIL"
)

type RadarEmailSecurityTopLocationByDkimListParamsDateRange string

const (
	RadarEmailSecurityTopLocationByDkimListParamsDateRange1d         RadarEmailSecurityTopLocationByDkimListParamsDateRange = "1d"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange2d         RadarEmailSecurityTopLocationByDkimListParamsDateRange = "2d"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange7d         RadarEmailSecurityTopLocationByDkimListParamsDateRange = "7d"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange14d        RadarEmailSecurityTopLocationByDkimListParamsDateRange = "14d"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange28d        RadarEmailSecurityTopLocationByDkimListParamsDateRange = "28d"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange12w        RadarEmailSecurityTopLocationByDkimListParamsDateRange = "12w"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange24w        RadarEmailSecurityTopLocationByDkimListParamsDateRange = "24w"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange52w        RadarEmailSecurityTopLocationByDkimListParamsDateRange = "52w"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange1dControl  RadarEmailSecurityTopLocationByDkimListParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange2dControl  RadarEmailSecurityTopLocationByDkimListParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange7dControl  RadarEmailSecurityTopLocationByDkimListParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange14dControl RadarEmailSecurityTopLocationByDkimListParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange28dControl RadarEmailSecurityTopLocationByDkimListParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange12wControl RadarEmailSecurityTopLocationByDkimListParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationByDkimListParamsDateRange24wControl RadarEmailSecurityTopLocationByDkimListParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationByDkimListParamsDmarc string

const (
	RadarEmailSecurityTopLocationByDkimListParamsDmarcPass RadarEmailSecurityTopLocationByDkimListParamsDmarc = "PASS"
	RadarEmailSecurityTopLocationByDkimListParamsDmarcNone RadarEmailSecurityTopLocationByDkimListParamsDmarc = "NONE"
	RadarEmailSecurityTopLocationByDkimListParamsDmarcFail RadarEmailSecurityTopLocationByDkimListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationByDkimListParamsFormat string

const (
	RadarEmailSecurityTopLocationByDkimListParamsFormatJson RadarEmailSecurityTopLocationByDkimListParamsFormat = "JSON"
	RadarEmailSecurityTopLocationByDkimListParamsFormatCsv  RadarEmailSecurityTopLocationByDkimListParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationByDkimListParamsSpf string

const (
	RadarEmailSecurityTopLocationByDkimListParamsSpfPass RadarEmailSecurityTopLocationByDkimListParamsSpf = "PASS"
	RadarEmailSecurityTopLocationByDkimListParamsSpfNone RadarEmailSecurityTopLocationByDkimListParamsSpf = "NONE"
	RadarEmailSecurityTopLocationByDkimListParamsSpfFail RadarEmailSecurityTopLocationByDkimListParamsSpf = "FAIL"
)
