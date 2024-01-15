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

// RadarEmailSecurityTopLocationByDmarcService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationByDmarcService] method instead.
type RadarEmailSecurityTopLocationByDmarcService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationByDmarcService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityTopLocationByDmarcService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationByDmarcService) {
	r = &RadarEmailSecurityTopLocationByDmarcService{}
	r.Options = opts
	return
}

// Get the locations by email DMARC validation.
func (r *RadarEmailSecurityTopLocationByDmarcService) List(ctx context.Context, dmarc RadarEmailSecurityTopLocationByDmarcListParamsDmarc, query RadarEmailSecurityTopLocationByDmarcListParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationByDmarcListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/locations/dmarc/%v", dmarc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTopLocationByDmarcListResponse struct {
	Result  RadarEmailSecurityTopLocationByDmarcListResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarEmailSecurityTopLocationByDmarcListResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationByDmarcListResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationByDmarcListResponse]
type radarEmailSecurityTopLocationByDmarcListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByDmarcListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDmarcListResponseResult struct {
	Meta RadarEmailSecurityTopLocationByDmarcListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationByDmarcListResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationByDmarcListResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopLocationByDmarcListResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationByDmarcListResponseResult]
type radarEmailSecurityTopLocationByDmarcListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByDmarcListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDmarcListResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopLocationByDmarcListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                   `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationByDmarcListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationByDmarcListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationByDmarcListResponseResultMeta]
type radarEmailSecurityTopLocationByDmarcListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByDmarcListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDmarcListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                               `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationByDmarcListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationByDmarcListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopLocationByDmarcListResponseResultMetaDateRange]
type radarEmailSecurityTopLocationByDmarcListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByDmarcListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                `json:"level"`
	JSON        radarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                 `json:"dataSource,required"`
	Description     string                                                                                 `json:"description,required"`
	EventType       string                                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                                              `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationByDmarcListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDmarcListResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                         `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                         `json:"clientCountryName,required"`
	Value               string                                                         `json:"value,required"`
	JSON                radarEmailSecurityTopLocationByDmarcListResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationByDmarcListResponseResultTop0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationByDmarcListResponseResultTop0]
type radarEmailSecurityTopLocationByDmarcListResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByDmarcListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByDmarcListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopLocationByDmarcListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationByDmarcListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityTopLocationByDmarcListParamsDkim] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationByDmarcListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityTopLocationByDmarcListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationByDmarcListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationByDmarcListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DMARC.
type RadarEmailSecurityTopLocationByDmarcListParamsDmarc string

const (
	RadarEmailSecurityTopLocationByDmarcListParamsDmarcPass RadarEmailSecurityTopLocationByDmarcListParamsDmarc = "PASS"
	RadarEmailSecurityTopLocationByDmarcListParamsDmarcNone RadarEmailSecurityTopLocationByDmarcListParamsDmarc = "NONE"
	RadarEmailSecurityTopLocationByDmarcListParamsDmarcFail RadarEmailSecurityTopLocationByDmarcListParamsDmarc = "FAIL"
)

type RadarEmailSecurityTopLocationByDmarcListParamsArc string

const (
	RadarEmailSecurityTopLocationByDmarcListParamsArcPass RadarEmailSecurityTopLocationByDmarcListParamsArc = "PASS"
	RadarEmailSecurityTopLocationByDmarcListParamsArcNone RadarEmailSecurityTopLocationByDmarcListParamsArc = "NONE"
	RadarEmailSecurityTopLocationByDmarcListParamsArcFail RadarEmailSecurityTopLocationByDmarcListParamsArc = "FAIL"
)

type RadarEmailSecurityTopLocationByDmarcListParamsDateRange string

const (
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange1d         RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "1d"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange2d         RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "2d"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange7d         RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "7d"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange14d        RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "14d"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange28d        RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "28d"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange12w        RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "12w"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange24w        RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "24w"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange52w        RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "52w"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange1dControl  RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange2dControl  RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange7dControl  RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange14dControl RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange28dControl RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange12wControl RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationByDmarcListParamsDateRange24wControl RadarEmailSecurityTopLocationByDmarcListParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationByDmarcListParamsDkim string

const (
	RadarEmailSecurityTopLocationByDmarcListParamsDkimPass RadarEmailSecurityTopLocationByDmarcListParamsDkim = "PASS"
	RadarEmailSecurityTopLocationByDmarcListParamsDkimNone RadarEmailSecurityTopLocationByDmarcListParamsDkim = "NONE"
	RadarEmailSecurityTopLocationByDmarcListParamsDkimFail RadarEmailSecurityTopLocationByDmarcListParamsDkim = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationByDmarcListParamsFormat string

const (
	RadarEmailSecurityTopLocationByDmarcListParamsFormatJson RadarEmailSecurityTopLocationByDmarcListParamsFormat = "JSON"
	RadarEmailSecurityTopLocationByDmarcListParamsFormatCsv  RadarEmailSecurityTopLocationByDmarcListParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationByDmarcListParamsSpf string

const (
	RadarEmailSecurityTopLocationByDmarcListParamsSpfPass RadarEmailSecurityTopLocationByDmarcListParamsSpf = "PASS"
	RadarEmailSecurityTopLocationByDmarcListParamsSpfNone RadarEmailSecurityTopLocationByDmarcListParamsSpf = "NONE"
	RadarEmailSecurityTopLocationByDmarcListParamsSpfFail RadarEmailSecurityTopLocationByDmarcListParamsSpf = "FAIL"
)
