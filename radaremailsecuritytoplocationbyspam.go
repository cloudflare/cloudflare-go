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

// RadarEmailSecurityTopLocationBySpamService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationBySpamService] method instead.
type RadarEmailSecurityTopLocationBySpamService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationBySpamService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityTopLocationBySpamService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationBySpamService) {
	r = &RadarEmailSecurityTopLocationBySpamService{}
	r.Options = opts
	return
}

// Get the top locations by emails classified as Spam or not.
func (r *RadarEmailSecurityTopLocationBySpamService) List(ctx context.Context, spam RadarEmailSecurityTopLocationBySpamListParamsSpam, query RadarEmailSecurityTopLocationBySpamListParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationBySpamListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/security/top/locations/spam/%v", spam)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTopLocationBySpamListResponse struct {
	Result  RadarEmailSecurityTopLocationBySpamListResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarEmailSecurityTopLocationBySpamListResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationBySpamListResponseJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationBySpamListResponse]
type radarEmailSecurityTopLocationBySpamListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationBySpamListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpamListResponseResult struct {
	Meta RadarEmailSecurityTopLocationBySpamListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationBySpamListResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationBySpamListResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopLocationBySpamListResponseResultJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationBySpamListResponseResult]
type radarEmailSecurityTopLocationBySpamListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationBySpamListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpamListResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopLocationBySpamListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                                  `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationBySpamListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationBySpamListResponseResultMetaJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationBySpamListResponseResultMeta]
type radarEmailSecurityTopLocationBySpamListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationBySpamListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpamListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationBySpamListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationBySpamListResponseResultMetaDateRangeJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopLocationBySpamListResponseResultMetaDateRange]
type radarEmailSecurityTopLocationBySpamListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationBySpamListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                               `json:"level"`
	JSON        radarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfoJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                                `json:"dataSource,required"`
	Description     string                                                                                `json:"description,required"`
	EventType       string                                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                                             `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationBySpamListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpamListResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                        `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                        `json:"clientCountryName,required"`
	Value               string                                                        `json:"value,required"`
	JSON                radarEmailSecurityTopLocationBySpamListResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationBySpamListResponseResultTop0JSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationBySpamListResponseResultTop0]
type radarEmailSecurityTopLocationBySpamListResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationBySpamListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationBySpamListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopLocationBySpamListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationBySpamListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityTopLocationBySpamListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopLocationBySpamListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationBySpamListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityTopLocationBySpamListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationBySpamListParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationBySpamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Spam.
type RadarEmailSecurityTopLocationBySpamListParamsSpam string

const (
	RadarEmailSecurityTopLocationBySpamListParamsSpamSpam    RadarEmailSecurityTopLocationBySpamListParamsSpam = "SPAM"
	RadarEmailSecurityTopLocationBySpamListParamsSpamNotSpam RadarEmailSecurityTopLocationBySpamListParamsSpam = "NOT_SPAM"
)

type RadarEmailSecurityTopLocationBySpamListParamsArc string

const (
	RadarEmailSecurityTopLocationBySpamListParamsArcPass RadarEmailSecurityTopLocationBySpamListParamsArc = "PASS"
	RadarEmailSecurityTopLocationBySpamListParamsArcNone RadarEmailSecurityTopLocationBySpamListParamsArc = "NONE"
	RadarEmailSecurityTopLocationBySpamListParamsArcFail RadarEmailSecurityTopLocationBySpamListParamsArc = "FAIL"
)

type RadarEmailSecurityTopLocationBySpamListParamsDateRange string

const (
	RadarEmailSecurityTopLocationBySpamListParamsDateRange1d         RadarEmailSecurityTopLocationBySpamListParamsDateRange = "1d"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange2d         RadarEmailSecurityTopLocationBySpamListParamsDateRange = "2d"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange7d         RadarEmailSecurityTopLocationBySpamListParamsDateRange = "7d"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange14d        RadarEmailSecurityTopLocationBySpamListParamsDateRange = "14d"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange28d        RadarEmailSecurityTopLocationBySpamListParamsDateRange = "28d"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange12w        RadarEmailSecurityTopLocationBySpamListParamsDateRange = "12w"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange24w        RadarEmailSecurityTopLocationBySpamListParamsDateRange = "24w"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange52w        RadarEmailSecurityTopLocationBySpamListParamsDateRange = "52w"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange1dControl  RadarEmailSecurityTopLocationBySpamListParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange2dControl  RadarEmailSecurityTopLocationBySpamListParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange7dControl  RadarEmailSecurityTopLocationBySpamListParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange14dControl RadarEmailSecurityTopLocationBySpamListParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange28dControl RadarEmailSecurityTopLocationBySpamListParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange12wControl RadarEmailSecurityTopLocationBySpamListParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationBySpamListParamsDateRange24wControl RadarEmailSecurityTopLocationBySpamListParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationBySpamListParamsDkim string

const (
	RadarEmailSecurityTopLocationBySpamListParamsDkimPass RadarEmailSecurityTopLocationBySpamListParamsDkim = "PASS"
	RadarEmailSecurityTopLocationBySpamListParamsDkimNone RadarEmailSecurityTopLocationBySpamListParamsDkim = "NONE"
	RadarEmailSecurityTopLocationBySpamListParamsDkimFail RadarEmailSecurityTopLocationBySpamListParamsDkim = "FAIL"
)

type RadarEmailSecurityTopLocationBySpamListParamsDmarc string

const (
	RadarEmailSecurityTopLocationBySpamListParamsDmarcPass RadarEmailSecurityTopLocationBySpamListParamsDmarc = "PASS"
	RadarEmailSecurityTopLocationBySpamListParamsDmarcNone RadarEmailSecurityTopLocationBySpamListParamsDmarc = "NONE"
	RadarEmailSecurityTopLocationBySpamListParamsDmarcFail RadarEmailSecurityTopLocationBySpamListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationBySpamListParamsFormat string

const (
	RadarEmailSecurityTopLocationBySpamListParamsFormatJson RadarEmailSecurityTopLocationBySpamListParamsFormat = "JSON"
	RadarEmailSecurityTopLocationBySpamListParamsFormatCsv  RadarEmailSecurityTopLocationBySpamListParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationBySpamListParamsSpf string

const (
	RadarEmailSecurityTopLocationBySpamListParamsSpfPass RadarEmailSecurityTopLocationBySpamListParamsSpf = "PASS"
	RadarEmailSecurityTopLocationBySpamListParamsSpfNone RadarEmailSecurityTopLocationBySpamListParamsSpf = "NONE"
	RadarEmailSecurityTopLocationBySpamListParamsSpfFail RadarEmailSecurityTopLocationBySpamListParamsSpf = "FAIL"
)
