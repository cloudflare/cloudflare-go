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

// RadarEmailSecurityTopLocationByMaliciousService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarEmailSecurityTopLocationByMaliciousService] method instead.
type RadarEmailSecurityTopLocationByMaliciousService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationByMaliciousService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityTopLocationByMaliciousService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationByMaliciousService) {
	r = &RadarEmailSecurityTopLocationByMaliciousService{}
	r.Options = opts
	return
}

// Get the locations by emails classified as malicious or not.
func (r *RadarEmailSecurityTopLocationByMaliciousService) List(ctx context.Context, malicious RadarEmailSecurityTopLocationByMaliciousListParamsMalicious, query RadarEmailSecurityTopLocationByMaliciousListParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationByMaliciousListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/locations/malicious/%v", malicious)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTopLocationByMaliciousListResponse struct {
	Result  RadarEmailSecurityTopLocationByMaliciousListResponseResult `json:"result,required"`
	Success bool                                                       `json:"success,required"`
	JSON    radarEmailSecurityTopLocationByMaliciousListResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationByMaliciousListResponseJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationByMaliciousListResponse]
type radarEmailSecurityTopLocationByMaliciousListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByMaliciousListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByMaliciousListResponseResult struct {
	Meta RadarEmailSecurityTopLocationByMaliciousListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationByMaliciousListResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationByMaliciousListResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopLocationByMaliciousListResponseResultJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationByMaliciousListResponseResult]
type radarEmailSecurityTopLocationByMaliciousListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByMaliciousListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByMaliciousListResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopLocationByMaliciousListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                       `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationByMaliciousListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationByMaliciousListResponseResultMetaJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopLocationByMaliciousListResponseResultMeta]
type radarEmailSecurityTopLocationByMaliciousListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByMaliciousListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByMaliciousListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationByMaliciousListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationByMaliciousListResponseResultMetaDateRangeJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationByMaliciousListResponseResultMetaDateRange]
type radarEmailSecurityTopLocationByMaliciousListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByMaliciousListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                                    `json:"level"`
	JSON        radarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                     `json:"dataSource,required"`
	Description     string                                                                                     `json:"description,required"`
	EventType       string                                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                                  `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationByMaliciousListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByMaliciousListResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                             `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                             `json:"clientCountryName,required"`
	Value               string                                                             `json:"value,required"`
	JSON                radarEmailSecurityTopLocationByMaliciousListResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationByMaliciousListResponseResultTop0JSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopLocationByMaliciousListResponseResultTop0]
type radarEmailSecurityTopLocationByMaliciousListResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationByMaliciousListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationByMaliciousListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopLocationByMaliciousListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationByMaliciousListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityTopLocationByMaliciousListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopLocationByMaliciousListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationByMaliciousListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityTopLocationByMaliciousListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationByMaliciousListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationByMaliciousListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Malicious.
type RadarEmailSecurityTopLocationByMaliciousListParamsMalicious string

const (
	RadarEmailSecurityTopLocationByMaliciousListParamsMaliciousMalicious    RadarEmailSecurityTopLocationByMaliciousListParamsMalicious = "MALICIOUS"
	RadarEmailSecurityTopLocationByMaliciousListParamsMaliciousNotMalicious RadarEmailSecurityTopLocationByMaliciousListParamsMalicious = "NOT_MALICIOUS"
)

type RadarEmailSecurityTopLocationByMaliciousListParamsArc string

const (
	RadarEmailSecurityTopLocationByMaliciousListParamsArcPass RadarEmailSecurityTopLocationByMaliciousListParamsArc = "PASS"
	RadarEmailSecurityTopLocationByMaliciousListParamsArcNone RadarEmailSecurityTopLocationByMaliciousListParamsArc = "NONE"
	RadarEmailSecurityTopLocationByMaliciousListParamsArcFail RadarEmailSecurityTopLocationByMaliciousListParamsArc = "FAIL"
)

type RadarEmailSecurityTopLocationByMaliciousListParamsDateRange string

const (
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange1d         RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "1d"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange2d         RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "2d"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange7d         RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "7d"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange14d        RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "14d"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange28d        RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "28d"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange12w        RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "12w"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange24w        RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "24w"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange52w        RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "52w"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange1dControl  RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange2dControl  RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange7dControl  RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange14dControl RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange28dControl RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange12wControl RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationByMaliciousListParamsDateRange24wControl RadarEmailSecurityTopLocationByMaliciousListParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationByMaliciousListParamsDkim string

const (
	RadarEmailSecurityTopLocationByMaliciousListParamsDkimPass RadarEmailSecurityTopLocationByMaliciousListParamsDkim = "PASS"
	RadarEmailSecurityTopLocationByMaliciousListParamsDkimNone RadarEmailSecurityTopLocationByMaliciousListParamsDkim = "NONE"
	RadarEmailSecurityTopLocationByMaliciousListParamsDkimFail RadarEmailSecurityTopLocationByMaliciousListParamsDkim = "FAIL"
)

type RadarEmailSecurityTopLocationByMaliciousListParamsDmarc string

const (
	RadarEmailSecurityTopLocationByMaliciousListParamsDmarcPass RadarEmailSecurityTopLocationByMaliciousListParamsDmarc = "PASS"
	RadarEmailSecurityTopLocationByMaliciousListParamsDmarcNone RadarEmailSecurityTopLocationByMaliciousListParamsDmarc = "NONE"
	RadarEmailSecurityTopLocationByMaliciousListParamsDmarcFail RadarEmailSecurityTopLocationByMaliciousListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationByMaliciousListParamsFormat string

const (
	RadarEmailSecurityTopLocationByMaliciousListParamsFormatJson RadarEmailSecurityTopLocationByMaliciousListParamsFormat = "JSON"
	RadarEmailSecurityTopLocationByMaliciousListParamsFormatCsv  RadarEmailSecurityTopLocationByMaliciousListParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationByMaliciousListParamsSpf string

const (
	RadarEmailSecurityTopLocationByMaliciousListParamsSpfPass RadarEmailSecurityTopLocationByMaliciousListParamsSpf = "PASS"
	RadarEmailSecurityTopLocationByMaliciousListParamsSpfNone RadarEmailSecurityTopLocationByMaliciousListParamsSpf = "NONE"
	RadarEmailSecurityTopLocationByMaliciousListParamsSpfFail RadarEmailSecurityTopLocationByMaliciousListParamsSpf = "FAIL"
)
