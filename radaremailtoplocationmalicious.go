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

// RadarEmailTopLocationMaliciousService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailTopLocationMaliciousService] method instead.
type RadarEmailTopLocationMaliciousService struct {
	Options []option.RequestOption
}

// NewRadarEmailTopLocationMaliciousService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailTopLocationMaliciousService(opts ...option.RequestOption) (r *RadarEmailTopLocationMaliciousService) {
	r = &RadarEmailTopLocationMaliciousService{}
	r.Options = opts
	return
}

// Get the locations, by emails classified, of Malicious validations.
func (r *RadarEmailTopLocationMaliciousService) Get(ctx context.Context, malicious RadarEmailTopLocationMaliciousGetParamsMalicious, query RadarEmailTopLocationMaliciousGetParams, opts ...option.RequestOption) (res *RadarEmailTopLocationMaliciousGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/top/locations/malicious/%v", malicious)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopLocationMaliciousGetResponse struct {
	Result  RadarEmailTopLocationMaliciousGetResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarEmailTopLocationMaliciousGetResponseJSON   `json:"-"`
}

// radarEmailTopLocationMaliciousGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailTopLocationMaliciousGetResponse]
type radarEmailTopLocationMaliciousGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationMaliciousGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationMaliciousGetResponseResult struct {
	Meta RadarEmailTopLocationMaliciousGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopLocationMaliciousGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopLocationMaliciousGetResponseResultJSON   `json:"-"`
}

// radarEmailTopLocationMaliciousGetResponseResultJSON contains the JSON metadata
// for the struct [RadarEmailTopLocationMaliciousGetResponseResult]
type radarEmailTopLocationMaliciousGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationMaliciousGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationMaliciousGetResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopLocationMaliciousGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopLocationMaliciousGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopLocationMaliciousGetResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarEmailTopLocationMaliciousGetResponseResultMeta]
type radarEmailTopLocationMaliciousGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopLocationMaliciousGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                         `json:"level,required"`
	JSON        radarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfo]
type radarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                          `json:"dataSource,required"`
	Description string                                                                          `json:"description,required"`
	EndTime     time.Time                                                                       `json:"endTime,required" format:"date-time"`
	EventType   string                                                                          `json:"eventType,required"`
	StartTime   time.Time                                                                       `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationMaliciousGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationMaliciousGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                        `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopLocationMaliciousGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopLocationMaliciousGetResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarEmailTopLocationMaliciousGetResponseResultMetaDateRange]
type radarEmailTopLocationMaliciousGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationMaliciousGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationMaliciousGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                  `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                  `json:"clientCountryName,required"`
	Value               string                                                  `json:"value,required"`
	JSON                radarEmailTopLocationMaliciousGetResponseResultTop0JSON `json:"-"`
}

// radarEmailTopLocationMaliciousGetResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarEmailTopLocationMaliciousGetResponseResultTop0]
type radarEmailTopLocationMaliciousGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailTopLocationMaliciousGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationMaliciousGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTopLocationMaliciousGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopLocationMaliciousGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTopLocationMaliciousGetParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTopLocationMaliciousGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopLocationMaliciousGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTopLocationMaliciousGetParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTopLocationMaliciousGetParams]'s query parameters
// as `url.Values`.
func (r RadarEmailTopLocationMaliciousGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopLocationMaliciousGetParamsMalicious string

const (
	RadarEmailTopLocationMaliciousGetParamsMaliciousMalicious    RadarEmailTopLocationMaliciousGetParamsMalicious = "MALICIOUS"
	RadarEmailTopLocationMaliciousGetParamsMaliciousNotMalicious RadarEmailTopLocationMaliciousGetParamsMalicious = "NOT_MALICIOUS"
)

type RadarEmailTopLocationMaliciousGetParamsArc string

const (
	RadarEmailTopLocationMaliciousGetParamsArcPass RadarEmailTopLocationMaliciousGetParamsArc = "PASS"
	RadarEmailTopLocationMaliciousGetParamsArcNone RadarEmailTopLocationMaliciousGetParamsArc = "NONE"
	RadarEmailTopLocationMaliciousGetParamsArcFail RadarEmailTopLocationMaliciousGetParamsArc = "FAIL"
)

type RadarEmailTopLocationMaliciousGetParamsDateRange string

const (
	RadarEmailTopLocationMaliciousGetParamsDateRange1d         RadarEmailTopLocationMaliciousGetParamsDateRange = "1d"
	RadarEmailTopLocationMaliciousGetParamsDateRange7d         RadarEmailTopLocationMaliciousGetParamsDateRange = "7d"
	RadarEmailTopLocationMaliciousGetParamsDateRange14d        RadarEmailTopLocationMaliciousGetParamsDateRange = "14d"
	RadarEmailTopLocationMaliciousGetParamsDateRange28d        RadarEmailTopLocationMaliciousGetParamsDateRange = "28d"
	RadarEmailTopLocationMaliciousGetParamsDateRange12w        RadarEmailTopLocationMaliciousGetParamsDateRange = "12w"
	RadarEmailTopLocationMaliciousGetParamsDateRange24w        RadarEmailTopLocationMaliciousGetParamsDateRange = "24w"
	RadarEmailTopLocationMaliciousGetParamsDateRange52w        RadarEmailTopLocationMaliciousGetParamsDateRange = "52w"
	RadarEmailTopLocationMaliciousGetParamsDateRange1dControl  RadarEmailTopLocationMaliciousGetParamsDateRange = "1dControl"
	RadarEmailTopLocationMaliciousGetParamsDateRange7dControl  RadarEmailTopLocationMaliciousGetParamsDateRange = "7dControl"
	RadarEmailTopLocationMaliciousGetParamsDateRange14dControl RadarEmailTopLocationMaliciousGetParamsDateRange = "14dControl"
	RadarEmailTopLocationMaliciousGetParamsDateRange28dControl RadarEmailTopLocationMaliciousGetParamsDateRange = "28dControl"
	RadarEmailTopLocationMaliciousGetParamsDateRange12wControl RadarEmailTopLocationMaliciousGetParamsDateRange = "12wControl"
	RadarEmailTopLocationMaliciousGetParamsDateRange24wControl RadarEmailTopLocationMaliciousGetParamsDateRange = "24wControl"
)

type RadarEmailTopLocationMaliciousGetParamsDkim string

const (
	RadarEmailTopLocationMaliciousGetParamsDkimPass RadarEmailTopLocationMaliciousGetParamsDkim = "PASS"
	RadarEmailTopLocationMaliciousGetParamsDkimNone RadarEmailTopLocationMaliciousGetParamsDkim = "NONE"
	RadarEmailTopLocationMaliciousGetParamsDkimFail RadarEmailTopLocationMaliciousGetParamsDkim = "FAIL"
)

type RadarEmailTopLocationMaliciousGetParamsDmarc string

const (
	RadarEmailTopLocationMaliciousGetParamsDmarcPass RadarEmailTopLocationMaliciousGetParamsDmarc = "PASS"
	RadarEmailTopLocationMaliciousGetParamsDmarcNone RadarEmailTopLocationMaliciousGetParamsDmarc = "NONE"
	RadarEmailTopLocationMaliciousGetParamsDmarcFail RadarEmailTopLocationMaliciousGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTopLocationMaliciousGetParamsFormat string

const (
	RadarEmailTopLocationMaliciousGetParamsFormatJson RadarEmailTopLocationMaliciousGetParamsFormat = "JSON"
	RadarEmailTopLocationMaliciousGetParamsFormatCsv  RadarEmailTopLocationMaliciousGetParamsFormat = "CSV"
)

type RadarEmailTopLocationMaliciousGetParamsSpf string

const (
	RadarEmailTopLocationMaliciousGetParamsSpfPass RadarEmailTopLocationMaliciousGetParamsSpf = "PASS"
	RadarEmailTopLocationMaliciousGetParamsSpfNone RadarEmailTopLocationMaliciousGetParamsSpf = "NONE"
	RadarEmailTopLocationMaliciousGetParamsSpfFail RadarEmailTopLocationMaliciousGetParamsSpf = "FAIL"
)
