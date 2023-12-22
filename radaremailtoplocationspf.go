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

// RadarEmailTopLocationSpfService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailTopLocationSpfService] method instead.
type RadarEmailTopLocationSpfService struct {
	Options []option.RequestOption
}

// NewRadarEmailTopLocationSpfService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailTopLocationSpfService(opts ...option.RequestOption) (r *RadarEmailTopLocationSpfService) {
	r = &RadarEmailTopLocationSpfService{}
	r.Options = opts
	return
}

// Get the locations, by emails classified, of SPF validations.
func (r *RadarEmailTopLocationSpfService) Get(ctx context.Context, spf RadarEmailTopLocationSpfGetParamsSpf, query RadarEmailTopLocationSpfGetParams, opts ...option.RequestOption) (res *RadarEmailTopLocationSpfGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/top/locations/spf/%v", spf)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopLocationSpfGetResponse struct {
	Result  RadarEmailTopLocationSpfGetResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarEmailTopLocationSpfGetResponseJSON   `json:"-"`
}

// radarEmailTopLocationSpfGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailTopLocationSpfGetResponse]
type radarEmailTopLocationSpfGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpfGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpfGetResponseResult struct {
	Meta RadarEmailTopLocationSpfGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopLocationSpfGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopLocationSpfGetResponseResultJSON   `json:"-"`
}

// radarEmailTopLocationSpfGetResponseResultJSON contains the JSON metadata for the
// struct [RadarEmailTopLocationSpfGetResponseResult]
type radarEmailTopLocationSpfGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpfGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpfGetResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopLocationSpfGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopLocationSpfGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopLocationSpfGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopLocationSpfGetResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarEmailTopLocationSpfGetResponseResultMeta]
type radarEmailTopLocationSpfGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpfGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpfGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopLocationSpfGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                   `json:"level,required"`
	JSON        radarEmailTopLocationSpfGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopLocationSpfGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailTopLocationSpfGetResponseResultMetaConfidenceInfo]
type radarEmailTopLocationSpfGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpfGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpfGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                    `json:"dataSource,required"`
	Description string                                                                    `json:"description,required"`
	EndTime     time.Time                                                                 `json:"endTime,required" format:"date-time"`
	EventType   string                                                                    `json:"eventType,required"`
	StartTime   time.Time                                                                 `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopLocationSpfGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopLocationSpfGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailTopLocationSpfGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopLocationSpfGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpfGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpfGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                  `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopLocationSpfGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopLocationSpfGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailTopLocationSpfGetResponseResultMetaDateRange]
type radarEmailTopLocationSpfGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpfGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpfGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                            `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                            `json:"clientCountryName,required"`
	Value               string                                            `json:"value,required"`
	JSON                radarEmailTopLocationSpfGetResponseResultTop0JSON `json:"-"`
}

// radarEmailTopLocationSpfGetResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarEmailTopLocationSpfGetResponseResultTop0]
type radarEmailTopLocationSpfGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpfGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpfGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTopLocationSpfGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopLocationSpfGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTopLocationSpfGetParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTopLocationSpfGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopLocationSpfGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
}

// URLQuery serializes [RadarEmailTopLocationSpfGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopLocationSpfGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopLocationSpfGetParamsSpf string

const (
	RadarEmailTopLocationSpfGetParamsSpfPass RadarEmailTopLocationSpfGetParamsSpf = "PASS"
	RadarEmailTopLocationSpfGetParamsSpfNone RadarEmailTopLocationSpfGetParamsSpf = "NONE"
	RadarEmailTopLocationSpfGetParamsSpfFail RadarEmailTopLocationSpfGetParamsSpf = "FAIL"
)

type RadarEmailTopLocationSpfGetParamsArc string

const (
	RadarEmailTopLocationSpfGetParamsArcPass RadarEmailTopLocationSpfGetParamsArc = "PASS"
	RadarEmailTopLocationSpfGetParamsArcNone RadarEmailTopLocationSpfGetParamsArc = "NONE"
	RadarEmailTopLocationSpfGetParamsArcFail RadarEmailTopLocationSpfGetParamsArc = "FAIL"
)

type RadarEmailTopLocationSpfGetParamsDateRange string

const (
	RadarEmailTopLocationSpfGetParamsDateRange1d         RadarEmailTopLocationSpfGetParamsDateRange = "1d"
	RadarEmailTopLocationSpfGetParamsDateRange7d         RadarEmailTopLocationSpfGetParamsDateRange = "7d"
	RadarEmailTopLocationSpfGetParamsDateRange14d        RadarEmailTopLocationSpfGetParamsDateRange = "14d"
	RadarEmailTopLocationSpfGetParamsDateRange28d        RadarEmailTopLocationSpfGetParamsDateRange = "28d"
	RadarEmailTopLocationSpfGetParamsDateRange12w        RadarEmailTopLocationSpfGetParamsDateRange = "12w"
	RadarEmailTopLocationSpfGetParamsDateRange24w        RadarEmailTopLocationSpfGetParamsDateRange = "24w"
	RadarEmailTopLocationSpfGetParamsDateRange52w        RadarEmailTopLocationSpfGetParamsDateRange = "52w"
	RadarEmailTopLocationSpfGetParamsDateRange1dControl  RadarEmailTopLocationSpfGetParamsDateRange = "1dControl"
	RadarEmailTopLocationSpfGetParamsDateRange7dControl  RadarEmailTopLocationSpfGetParamsDateRange = "7dControl"
	RadarEmailTopLocationSpfGetParamsDateRange14dControl RadarEmailTopLocationSpfGetParamsDateRange = "14dControl"
	RadarEmailTopLocationSpfGetParamsDateRange28dControl RadarEmailTopLocationSpfGetParamsDateRange = "28dControl"
	RadarEmailTopLocationSpfGetParamsDateRange12wControl RadarEmailTopLocationSpfGetParamsDateRange = "12wControl"
	RadarEmailTopLocationSpfGetParamsDateRange24wControl RadarEmailTopLocationSpfGetParamsDateRange = "24wControl"
)

type RadarEmailTopLocationSpfGetParamsDkim string

const (
	RadarEmailTopLocationSpfGetParamsDkimPass RadarEmailTopLocationSpfGetParamsDkim = "PASS"
	RadarEmailTopLocationSpfGetParamsDkimNone RadarEmailTopLocationSpfGetParamsDkim = "NONE"
	RadarEmailTopLocationSpfGetParamsDkimFail RadarEmailTopLocationSpfGetParamsDkim = "FAIL"
)

type RadarEmailTopLocationSpfGetParamsDmarc string

const (
	RadarEmailTopLocationSpfGetParamsDmarcPass RadarEmailTopLocationSpfGetParamsDmarc = "PASS"
	RadarEmailTopLocationSpfGetParamsDmarcNone RadarEmailTopLocationSpfGetParamsDmarc = "NONE"
	RadarEmailTopLocationSpfGetParamsDmarcFail RadarEmailTopLocationSpfGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTopLocationSpfGetParamsFormat string

const (
	RadarEmailTopLocationSpfGetParamsFormatJson RadarEmailTopLocationSpfGetParamsFormat = "JSON"
	RadarEmailTopLocationSpfGetParamsFormatCsv  RadarEmailTopLocationSpfGetParamsFormat = "CSV"
)
