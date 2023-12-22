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

// RadarEmailTopLocationArcService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailTopLocationArcService] method instead.
type RadarEmailTopLocationArcService struct {
	Options []option.RequestOption
}

// NewRadarEmailTopLocationArcService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailTopLocationArcService(opts ...option.RequestOption) (r *RadarEmailTopLocationArcService) {
	r = &RadarEmailTopLocationArcService{}
	r.Options = opts
	return
}

// Get the locations, by emails classified, of ARC validations.
func (r *RadarEmailTopLocationArcService) Get(ctx context.Context, arc RadarEmailTopLocationArcGetParamsArc, query RadarEmailTopLocationArcGetParams, opts ...option.RequestOption) (res *RadarEmailTopLocationArcGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/top/locations/arc/%v", arc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopLocationArcGetResponse struct {
	Result  RadarEmailTopLocationArcGetResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarEmailTopLocationArcGetResponseJSON   `json:"-"`
}

// radarEmailTopLocationArcGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailTopLocationArcGetResponse]
type radarEmailTopLocationArcGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationArcGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationArcGetResponseResult struct {
	Meta RadarEmailTopLocationArcGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopLocationArcGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopLocationArcGetResponseResultJSON   `json:"-"`
}

// radarEmailTopLocationArcGetResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTopLocationArcGetResponseResult]
type radarEmailTopLocationArcGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationArcGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationArcGetResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopLocationArcGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopLocationArcGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopLocationArcGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopLocationArcGetResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarEmailTopLocationArcGetResponseResultMeta]
type radarEmailTopLocationArcGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopLocationArcGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationArcGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopLocationArcGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                   `json:"level,required"`
	JSON        radarEmailTopLocationArcGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopLocationArcGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailTopLocationArcGetResponseResultMetaConfidenceInfo]
type radarEmailTopLocationArcGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationArcGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationArcGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                    `json:"dataSource,required"`
	Description string                                                                    `json:"description,required"`
	EndTime     time.Time                                                                 `json:"endTime,required" format:"date-time"`
	EventType   string                                                                    `json:"eventType,required"`
	StartTime   time.Time                                                                 `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopLocationArcGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopLocationArcGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailTopLocationArcGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopLocationArcGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationArcGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationArcGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                  `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopLocationArcGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopLocationArcGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailTopLocationArcGetResponseResultMetaDateRange]
type radarEmailTopLocationArcGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationArcGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationArcGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                            `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                            `json:"clientCountryName,required"`
	Value               string                                            `json:"value,required"`
	JSON                radarEmailTopLocationArcGetResponseResultTop0JSON `json:"-"`
}

// radarEmailTopLocationArcGetResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarEmailTopLocationArcGetResponseResultTop0]
type radarEmailTopLocationArcGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailTopLocationArcGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationArcGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopLocationArcGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTopLocationArcGetParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTopLocationArcGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopLocationArcGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTopLocationArcGetParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTopLocationArcGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopLocationArcGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopLocationArcGetParamsArc string

const (
	RadarEmailTopLocationArcGetParamsArcPass RadarEmailTopLocationArcGetParamsArc = "PASS"
	RadarEmailTopLocationArcGetParamsArcNone RadarEmailTopLocationArcGetParamsArc = "NONE"
	RadarEmailTopLocationArcGetParamsArcFail RadarEmailTopLocationArcGetParamsArc = "FAIL"
)

type RadarEmailTopLocationArcGetParamsDateRange string

const (
	RadarEmailTopLocationArcGetParamsDateRange1d         RadarEmailTopLocationArcGetParamsDateRange = "1d"
	RadarEmailTopLocationArcGetParamsDateRange7d         RadarEmailTopLocationArcGetParamsDateRange = "7d"
	RadarEmailTopLocationArcGetParamsDateRange14d        RadarEmailTopLocationArcGetParamsDateRange = "14d"
	RadarEmailTopLocationArcGetParamsDateRange28d        RadarEmailTopLocationArcGetParamsDateRange = "28d"
	RadarEmailTopLocationArcGetParamsDateRange12w        RadarEmailTopLocationArcGetParamsDateRange = "12w"
	RadarEmailTopLocationArcGetParamsDateRange24w        RadarEmailTopLocationArcGetParamsDateRange = "24w"
	RadarEmailTopLocationArcGetParamsDateRange52w        RadarEmailTopLocationArcGetParamsDateRange = "52w"
	RadarEmailTopLocationArcGetParamsDateRange1dControl  RadarEmailTopLocationArcGetParamsDateRange = "1dControl"
	RadarEmailTopLocationArcGetParamsDateRange7dControl  RadarEmailTopLocationArcGetParamsDateRange = "7dControl"
	RadarEmailTopLocationArcGetParamsDateRange14dControl RadarEmailTopLocationArcGetParamsDateRange = "14dControl"
	RadarEmailTopLocationArcGetParamsDateRange28dControl RadarEmailTopLocationArcGetParamsDateRange = "28dControl"
	RadarEmailTopLocationArcGetParamsDateRange12wControl RadarEmailTopLocationArcGetParamsDateRange = "12wControl"
	RadarEmailTopLocationArcGetParamsDateRange24wControl RadarEmailTopLocationArcGetParamsDateRange = "24wControl"
)

type RadarEmailTopLocationArcGetParamsDkim string

const (
	RadarEmailTopLocationArcGetParamsDkimPass RadarEmailTopLocationArcGetParamsDkim = "PASS"
	RadarEmailTopLocationArcGetParamsDkimNone RadarEmailTopLocationArcGetParamsDkim = "NONE"
	RadarEmailTopLocationArcGetParamsDkimFail RadarEmailTopLocationArcGetParamsDkim = "FAIL"
)

type RadarEmailTopLocationArcGetParamsDmarc string

const (
	RadarEmailTopLocationArcGetParamsDmarcPass RadarEmailTopLocationArcGetParamsDmarc = "PASS"
	RadarEmailTopLocationArcGetParamsDmarcNone RadarEmailTopLocationArcGetParamsDmarc = "NONE"
	RadarEmailTopLocationArcGetParamsDmarcFail RadarEmailTopLocationArcGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTopLocationArcGetParamsFormat string

const (
	RadarEmailTopLocationArcGetParamsFormatJson RadarEmailTopLocationArcGetParamsFormat = "JSON"
	RadarEmailTopLocationArcGetParamsFormatCsv  RadarEmailTopLocationArcGetParamsFormat = "CSV"
)

type RadarEmailTopLocationArcGetParamsSpf string

const (
	RadarEmailTopLocationArcGetParamsSpfPass RadarEmailTopLocationArcGetParamsSpf = "PASS"
	RadarEmailTopLocationArcGetParamsSpfNone RadarEmailTopLocationArcGetParamsSpf = "NONE"
	RadarEmailTopLocationArcGetParamsSpfFail RadarEmailTopLocationArcGetParamsSpf = "FAIL"
)
