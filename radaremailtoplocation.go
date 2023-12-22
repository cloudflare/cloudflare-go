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

// RadarEmailTopLocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarEmailTopLocationService]
// method instead.
type RadarEmailTopLocationService struct {
	Options   []option.RequestOption
	Arcs      *RadarEmailTopLocationArcService
	Dkims     *RadarEmailTopLocationDkimService
	Dmarcs    *RadarEmailTopLocationDmarcService
	Malicious *RadarEmailTopLocationMaliciousService
	Spams     *RadarEmailTopLocationSpamService
	Spfs      *RadarEmailTopLocationSpfService
}

// NewRadarEmailTopLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarEmailTopLocationService(opts ...option.RequestOption) (r *RadarEmailTopLocationService) {
	r = &RadarEmailTopLocationService{}
	r.Options = opts
	r.Arcs = NewRadarEmailTopLocationArcService(opts...)
	r.Dkims = NewRadarEmailTopLocationDkimService(opts...)
	r.Dmarcs = NewRadarEmailTopLocationDmarcService(opts...)
	r.Malicious = NewRadarEmailTopLocationMaliciousService(opts...)
	r.Spams = NewRadarEmailTopLocationSpamService(opts...)
	r.Spfs = NewRadarEmailTopLocationSpfService(opts...)
	return
}

// Get the top locations by HTTP traffic. Values are a percentage out of the total
// traffic.
func (r *RadarEmailTopLocationService) List(ctx context.Context, query RadarEmailTopLocationListParams, opts ...option.RequestOption) (res *RadarEmailTopLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/email/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopLocationListResponse struct {
	Result  RadarEmailTopLocationListResponseResult `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    radarEmailTopLocationListResponseJSON   `json:"-"`
}

// radarEmailTopLocationListResponseJSON contains the JSON metadata for the struct
// [RadarEmailTopLocationListResponse]
type radarEmailTopLocationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationListResponseResult struct {
	Meta RadarEmailTopLocationListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopLocationListResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopLocationListResponseResultJSON   `json:"-"`
}

// radarEmailTopLocationListResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTopLocationListResponseResult]
type radarEmailTopLocationListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationListResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopLocationListResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopLocationListResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopLocationListResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopLocationListResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarEmailTopLocationListResponseResultMeta]
type radarEmailTopLocationListResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopLocationListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopLocationListResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                 `json:"level,required"`
	JSON        radarEmailTopLocationListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopLocationListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailTopLocationListResponseResultMetaConfidenceInfo]
type radarEmailTopLocationListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                  `json:"dataSource,required"`
	Description string                                                                  `json:"description,required"`
	EndTime     time.Time                                                               `json:"endTime,required" format:"date-time"`
	EventType   string                                                                  `json:"eventType,required"`
	StartTime   time.Time                                                               `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarEmailTopLocationListResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationListResponseResultMetaDateRange struct {
	EndTime   time.Time                                                `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopLocationListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopLocationListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailTopLocationListResponseResultMetaDateRange]
type radarEmailTopLocationListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationListResponseResultTop0 struct {
	ClientCountryAlpha2 string                                          `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                          `json:"clientCountryName,required"`
	Value               string                                          `json:"value,required"`
	JSON                radarEmailTopLocationListResponseResultTop0JSON `json:"-"`
}

// radarEmailTopLocationListResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarEmailTopLocationListResponseResultTop0]
type radarEmailTopLocationListResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailTopLocationListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationListParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTopLocationListParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopLocationListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTopLocationListParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTopLocationListParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopLocationListParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTopLocationListParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTopLocationListParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopLocationListParamsArc string

const (
	RadarEmailTopLocationListParamsArcPass RadarEmailTopLocationListParamsArc = "PASS"
	RadarEmailTopLocationListParamsArcNone RadarEmailTopLocationListParamsArc = "NONE"
	RadarEmailTopLocationListParamsArcFail RadarEmailTopLocationListParamsArc = "FAIL"
)

type RadarEmailTopLocationListParamsDateRange string

const (
	RadarEmailTopLocationListParamsDateRange1d         RadarEmailTopLocationListParamsDateRange = "1d"
	RadarEmailTopLocationListParamsDateRange7d         RadarEmailTopLocationListParamsDateRange = "7d"
	RadarEmailTopLocationListParamsDateRange14d        RadarEmailTopLocationListParamsDateRange = "14d"
	RadarEmailTopLocationListParamsDateRange28d        RadarEmailTopLocationListParamsDateRange = "28d"
	RadarEmailTopLocationListParamsDateRange12w        RadarEmailTopLocationListParamsDateRange = "12w"
	RadarEmailTopLocationListParamsDateRange24w        RadarEmailTopLocationListParamsDateRange = "24w"
	RadarEmailTopLocationListParamsDateRange52w        RadarEmailTopLocationListParamsDateRange = "52w"
	RadarEmailTopLocationListParamsDateRange1dControl  RadarEmailTopLocationListParamsDateRange = "1dControl"
	RadarEmailTopLocationListParamsDateRange7dControl  RadarEmailTopLocationListParamsDateRange = "7dControl"
	RadarEmailTopLocationListParamsDateRange14dControl RadarEmailTopLocationListParamsDateRange = "14dControl"
	RadarEmailTopLocationListParamsDateRange28dControl RadarEmailTopLocationListParamsDateRange = "28dControl"
	RadarEmailTopLocationListParamsDateRange12wControl RadarEmailTopLocationListParamsDateRange = "12wControl"
	RadarEmailTopLocationListParamsDateRange24wControl RadarEmailTopLocationListParamsDateRange = "24wControl"
)

type RadarEmailTopLocationListParamsDkim string

const (
	RadarEmailTopLocationListParamsDkimPass RadarEmailTopLocationListParamsDkim = "PASS"
	RadarEmailTopLocationListParamsDkimNone RadarEmailTopLocationListParamsDkim = "NONE"
	RadarEmailTopLocationListParamsDkimFail RadarEmailTopLocationListParamsDkim = "FAIL"
)

type RadarEmailTopLocationListParamsDmarc string

const (
	RadarEmailTopLocationListParamsDmarcPass RadarEmailTopLocationListParamsDmarc = "PASS"
	RadarEmailTopLocationListParamsDmarcNone RadarEmailTopLocationListParamsDmarc = "NONE"
	RadarEmailTopLocationListParamsDmarcFail RadarEmailTopLocationListParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTopLocationListParamsFormat string

const (
	RadarEmailTopLocationListParamsFormatJson RadarEmailTopLocationListParamsFormat = "JSON"
	RadarEmailTopLocationListParamsFormatCsv  RadarEmailTopLocationListParamsFormat = "CSV"
)

type RadarEmailTopLocationListParamsSpf string

const (
	RadarEmailTopLocationListParamsSpfPass RadarEmailTopLocationListParamsSpf = "PASS"
	RadarEmailTopLocationListParamsSpfNone RadarEmailTopLocationListParamsSpf = "NONE"
	RadarEmailTopLocationListParamsSpfFail RadarEmailTopLocationListParamsSpf = "FAIL"
)
