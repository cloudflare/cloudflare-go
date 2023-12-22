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

// RadarEmailTopLocationDmarcService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailTopLocationDmarcService] method instead.
type RadarEmailTopLocationDmarcService struct {
	Options []option.RequestOption
}

// NewRadarEmailTopLocationDmarcService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailTopLocationDmarcService(opts ...option.RequestOption) (r *RadarEmailTopLocationDmarcService) {
	r = &RadarEmailTopLocationDmarcService{}
	r.Options = opts
	return
}

// Get the locations, by emails classified, of DMARC validations.
func (r *RadarEmailTopLocationDmarcService) Get(ctx context.Context, dmarc RadarEmailTopLocationDmarcGetParamsDmarc, query RadarEmailTopLocationDmarcGetParams, opts ...option.RequestOption) (res *RadarEmailTopLocationDmarcGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/top/locations/dmarc/%v", dmarc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopLocationDmarcGetResponse struct {
	Result  RadarEmailTopLocationDmarcGetResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarEmailTopLocationDmarcGetResponseJSON   `json:"-"`
}

// radarEmailTopLocationDmarcGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailTopLocationDmarcGetResponse]
type radarEmailTopLocationDmarcGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationDmarcGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDmarcGetResponseResult struct {
	Meta RadarEmailTopLocationDmarcGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopLocationDmarcGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopLocationDmarcGetResponseResultJSON   `json:"-"`
}

// radarEmailTopLocationDmarcGetResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailTopLocationDmarcGetResponseResult]
type radarEmailTopLocationDmarcGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationDmarcGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDmarcGetResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopLocationDmarcGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopLocationDmarcGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopLocationDmarcGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailTopLocationDmarcGetResponseResultMeta]
type radarEmailTopLocationDmarcGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopLocationDmarcGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                     `json:"level,required"`
	JSON        radarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfo]
type radarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                      `json:"dataSource,required"`
	Description string                                                                      `json:"description,required"`
	EndTime     time.Time                                                                   `json:"endTime,required" format:"date-time"`
	EventType   string                                                                      `json:"eventType,required"`
	StartTime   time.Time                                                                   `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationDmarcGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDmarcGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                    `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopLocationDmarcGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopLocationDmarcGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailTopLocationDmarcGetResponseResultMetaDateRange]
type radarEmailTopLocationDmarcGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationDmarcGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDmarcGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                              `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                              `json:"clientCountryName,required"`
	Value               string                                              `json:"value,required"`
	JSON                radarEmailTopLocationDmarcGetResponseResultTop0JSON `json:"-"`
}

// radarEmailTopLocationDmarcGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarEmailTopLocationDmarcGetResponseResultTop0]
type radarEmailTopLocationDmarcGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailTopLocationDmarcGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDmarcGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTopLocationDmarcGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopLocationDmarcGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTopLocationDmarcGetParamsDkim] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopLocationDmarcGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTopLocationDmarcGetParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTopLocationDmarcGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopLocationDmarcGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopLocationDmarcGetParamsDmarc string

const (
	RadarEmailTopLocationDmarcGetParamsDmarcPass RadarEmailTopLocationDmarcGetParamsDmarc = "PASS"
	RadarEmailTopLocationDmarcGetParamsDmarcNone RadarEmailTopLocationDmarcGetParamsDmarc = "NONE"
	RadarEmailTopLocationDmarcGetParamsDmarcFail RadarEmailTopLocationDmarcGetParamsDmarc = "FAIL"
)

type RadarEmailTopLocationDmarcGetParamsArc string

const (
	RadarEmailTopLocationDmarcGetParamsArcPass RadarEmailTopLocationDmarcGetParamsArc = "PASS"
	RadarEmailTopLocationDmarcGetParamsArcNone RadarEmailTopLocationDmarcGetParamsArc = "NONE"
	RadarEmailTopLocationDmarcGetParamsArcFail RadarEmailTopLocationDmarcGetParamsArc = "FAIL"
)

type RadarEmailTopLocationDmarcGetParamsDateRange string

const (
	RadarEmailTopLocationDmarcGetParamsDateRange1d         RadarEmailTopLocationDmarcGetParamsDateRange = "1d"
	RadarEmailTopLocationDmarcGetParamsDateRange7d         RadarEmailTopLocationDmarcGetParamsDateRange = "7d"
	RadarEmailTopLocationDmarcGetParamsDateRange14d        RadarEmailTopLocationDmarcGetParamsDateRange = "14d"
	RadarEmailTopLocationDmarcGetParamsDateRange28d        RadarEmailTopLocationDmarcGetParamsDateRange = "28d"
	RadarEmailTopLocationDmarcGetParamsDateRange12w        RadarEmailTopLocationDmarcGetParamsDateRange = "12w"
	RadarEmailTopLocationDmarcGetParamsDateRange24w        RadarEmailTopLocationDmarcGetParamsDateRange = "24w"
	RadarEmailTopLocationDmarcGetParamsDateRange52w        RadarEmailTopLocationDmarcGetParamsDateRange = "52w"
	RadarEmailTopLocationDmarcGetParamsDateRange1dControl  RadarEmailTopLocationDmarcGetParamsDateRange = "1dControl"
	RadarEmailTopLocationDmarcGetParamsDateRange7dControl  RadarEmailTopLocationDmarcGetParamsDateRange = "7dControl"
	RadarEmailTopLocationDmarcGetParamsDateRange14dControl RadarEmailTopLocationDmarcGetParamsDateRange = "14dControl"
	RadarEmailTopLocationDmarcGetParamsDateRange28dControl RadarEmailTopLocationDmarcGetParamsDateRange = "28dControl"
	RadarEmailTopLocationDmarcGetParamsDateRange12wControl RadarEmailTopLocationDmarcGetParamsDateRange = "12wControl"
	RadarEmailTopLocationDmarcGetParamsDateRange24wControl RadarEmailTopLocationDmarcGetParamsDateRange = "24wControl"
)

type RadarEmailTopLocationDmarcGetParamsDkim string

const (
	RadarEmailTopLocationDmarcGetParamsDkimPass RadarEmailTopLocationDmarcGetParamsDkim = "PASS"
	RadarEmailTopLocationDmarcGetParamsDkimNone RadarEmailTopLocationDmarcGetParamsDkim = "NONE"
	RadarEmailTopLocationDmarcGetParamsDkimFail RadarEmailTopLocationDmarcGetParamsDkim = "FAIL"
)

// Format results are returned in.
type RadarEmailTopLocationDmarcGetParamsFormat string

const (
	RadarEmailTopLocationDmarcGetParamsFormatJson RadarEmailTopLocationDmarcGetParamsFormat = "JSON"
	RadarEmailTopLocationDmarcGetParamsFormatCsv  RadarEmailTopLocationDmarcGetParamsFormat = "CSV"
)

type RadarEmailTopLocationDmarcGetParamsSpf string

const (
	RadarEmailTopLocationDmarcGetParamsSpfPass RadarEmailTopLocationDmarcGetParamsSpf = "PASS"
	RadarEmailTopLocationDmarcGetParamsSpfNone RadarEmailTopLocationDmarcGetParamsSpf = "NONE"
	RadarEmailTopLocationDmarcGetParamsSpfFail RadarEmailTopLocationDmarcGetParamsSpf = "FAIL"
)
