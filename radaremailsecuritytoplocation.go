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

// RadarEmailSecurityTopLocationService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationService] method instead.
type RadarEmailSecurityTopLocationService struct {
	Options     []option.RequestOption
	Arc         *RadarEmailSecurityTopLocationArcService
	ByDkim      *RadarEmailSecurityTopLocationByDkimService
	ByDmarc     *RadarEmailSecurityTopLocationByDmarcService
	ByMalicious *RadarEmailSecurityTopLocationByMaliciousService
	BySpam      *RadarEmailSecurityTopLocationBySpamService
	BySpf       *RadarEmailSecurityTopLocationBySpfService
}

// NewRadarEmailSecurityTopLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopLocationService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationService) {
	r = &RadarEmailSecurityTopLocationService{}
	r.Options = opts
	r.Arc = NewRadarEmailSecurityTopLocationArcService(opts...)
	r.ByDkim = NewRadarEmailSecurityTopLocationByDkimService(opts...)
	r.ByDmarc = NewRadarEmailSecurityTopLocationByDmarcService(opts...)
	r.ByMalicious = NewRadarEmailSecurityTopLocationByMaliciousService(opts...)
	r.BySpam = NewRadarEmailSecurityTopLocationBySpamService(opts...)
	r.BySpf = NewRadarEmailSecurityTopLocationBySpfService(opts...)
	return
}

// Get the top locations by email messages. Values are a percentage out of the
// total emails.
func (r *RadarEmailSecurityTopLocationService) List(ctx context.Context, query RadarEmailSecurityTopLocationListParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/security/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailSecurityTopLocationListResponse struct {
	Result  RadarEmailSecurityTopLocationListResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailSecurityTopLocationListResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationListResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopLocationListResponse]
type radarEmailSecurityTopLocationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationListResponseResult struct {
	Meta RadarEmailSecurityTopLocationListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationListResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationListResponseResultJSON   `json:"-"`
}

// radarEmailSecurityTopLocationListResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationListResponseResult]
type radarEmailSecurityTopLocationListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationListResponseResultMeta struct {
	DateRange      []RadarEmailSecurityTopLocationListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                            `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationListResponseResultMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationListResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationListResponseResultMeta]
type radarEmailSecurityTopLocationListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationListResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopLocationListResponseResultMetaDateRange]
type radarEmailSecurityTopLocationListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                         `json:"level"`
	JSON        radarEmailSecurityTopLocationListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationListResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailSecurityTopLocationListResponseResultMetaConfidenceInfo]
type radarEmailSecurityTopLocationListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                          `json:"dataSource,required"`
	Description     string                                                                          `json:"description,required"`
	EventType       string                                                                          `json:"eventType,required"`
	IsInstantaneous interface{}                                                                     `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                                       `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationListResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                  `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                  `json:"clientCountryName,required"`
	Value               string                                                  `json:"value,required"`
	JSON                radarEmailSecurityTopLocationListResponseResultTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationListResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationListResponseResultTop0]
type radarEmailSecurityTopLocationListResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopLocationListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailSecurityTopLocationListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopLocationListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailSecurityTopLocationListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationListParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityTopLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecurityTopLocationListParamsArc string

const (
	RadarEmailSecurityTopLocationListParamsArcPass RadarEmailSecurityTopLocationListParamsArc = "PASS"
	RadarEmailSecurityTopLocationListParamsArcNone RadarEmailSecurityTopLocationListParamsArc = "NONE"
	RadarEmailSecurityTopLocationListParamsArcFail RadarEmailSecurityTopLocationListParamsArc = "FAIL"
)

type RadarEmailSecurityTopLocationListParamsDateRange string

const (
	RadarEmailSecurityTopLocationListParamsDateRange1d         RadarEmailSecurityTopLocationListParamsDateRange = "1d"
	RadarEmailSecurityTopLocationListParamsDateRange2d         RadarEmailSecurityTopLocationListParamsDateRange = "2d"
	RadarEmailSecurityTopLocationListParamsDateRange7d         RadarEmailSecurityTopLocationListParamsDateRange = "7d"
	RadarEmailSecurityTopLocationListParamsDateRange14d        RadarEmailSecurityTopLocationListParamsDateRange = "14d"
	RadarEmailSecurityTopLocationListParamsDateRange28d        RadarEmailSecurityTopLocationListParamsDateRange = "28d"
	RadarEmailSecurityTopLocationListParamsDateRange12w        RadarEmailSecurityTopLocationListParamsDateRange = "12w"
	RadarEmailSecurityTopLocationListParamsDateRange24w        RadarEmailSecurityTopLocationListParamsDateRange = "24w"
	RadarEmailSecurityTopLocationListParamsDateRange52w        RadarEmailSecurityTopLocationListParamsDateRange = "52w"
	RadarEmailSecurityTopLocationListParamsDateRange1dControl  RadarEmailSecurityTopLocationListParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationListParamsDateRange2dControl  RadarEmailSecurityTopLocationListParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationListParamsDateRange7dControl  RadarEmailSecurityTopLocationListParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationListParamsDateRange14dControl RadarEmailSecurityTopLocationListParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationListParamsDateRange28dControl RadarEmailSecurityTopLocationListParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationListParamsDateRange12wControl RadarEmailSecurityTopLocationListParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationListParamsDateRange24wControl RadarEmailSecurityTopLocationListParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationListParamsDkim string

const (
	RadarEmailSecurityTopLocationListParamsDkimPass RadarEmailSecurityTopLocationListParamsDkim = "PASS"
	RadarEmailSecurityTopLocationListParamsDkimNone RadarEmailSecurityTopLocationListParamsDkim = "NONE"
	RadarEmailSecurityTopLocationListParamsDkimFail RadarEmailSecurityTopLocationListParamsDkim = "FAIL"
)

type RadarEmailSecurityTopLocationListParamsDmarc string

const (
	RadarEmailSecurityTopLocationListParamsDmarcPass RadarEmailSecurityTopLocationListParamsDmarc = "PASS"
	RadarEmailSecurityTopLocationListParamsDmarcNone RadarEmailSecurityTopLocationListParamsDmarc = "NONE"
	RadarEmailSecurityTopLocationListParamsDmarcFail RadarEmailSecurityTopLocationListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationListParamsFormat string

const (
	RadarEmailSecurityTopLocationListParamsFormatJson RadarEmailSecurityTopLocationListParamsFormat = "JSON"
	RadarEmailSecurityTopLocationListParamsFormatCsv  RadarEmailSecurityTopLocationListParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationListParamsSpf string

const (
	RadarEmailSecurityTopLocationListParamsSpfPass RadarEmailSecurityTopLocationListParamsSpf = "PASS"
	RadarEmailSecurityTopLocationListParamsSpfNone RadarEmailSecurityTopLocationListParamsSpf = "NONE"
	RadarEmailSecurityTopLocationListParamsSpfFail RadarEmailSecurityTopLocationListParamsSpf = "FAIL"
)
