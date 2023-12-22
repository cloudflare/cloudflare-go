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

// RadarEmailTopLocationDkimService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailTopLocationDkimService] method instead.
type RadarEmailTopLocationDkimService struct {
	Options []option.RequestOption
}

// NewRadarEmailTopLocationDkimService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailTopLocationDkimService(opts ...option.RequestOption) (r *RadarEmailTopLocationDkimService) {
	r = &RadarEmailTopLocationDkimService{}
	r.Options = opts
	return
}

// Get the locations, by emails classified, of DKIM validations.
func (r *RadarEmailTopLocationDkimService) Get(ctx context.Context, dkim RadarEmailTopLocationDkimGetParamsDkim, query RadarEmailTopLocationDkimGetParams, opts ...option.RequestOption) (res *RadarEmailTopLocationDkimGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/top/locations/dkim/%v", dkim)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopLocationDkimGetResponse struct {
	Result  RadarEmailTopLocationDkimGetResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarEmailTopLocationDkimGetResponseJSON   `json:"-"`
}

// radarEmailTopLocationDkimGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailTopLocationDkimGetResponse]
type radarEmailTopLocationDkimGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationDkimGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDkimGetResponseResult struct {
	Meta RadarEmailTopLocationDkimGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopLocationDkimGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopLocationDkimGetResponseResultJSON   `json:"-"`
}

// radarEmailTopLocationDkimGetResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailTopLocationDkimGetResponseResult]
type radarEmailTopLocationDkimGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationDkimGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDkimGetResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopLocationDkimGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopLocationDkimGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopLocationDkimGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopLocationDkimGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailTopLocationDkimGetResponseResultMeta]
type radarEmailTopLocationDkimGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopLocationDkimGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDkimGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopLocationDkimGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                    `json:"level,required"`
	JSON        radarEmailTopLocationDkimGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopLocationDkimGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailTopLocationDkimGetResponseResultMetaConfidenceInfo]
type radarEmailTopLocationDkimGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationDkimGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDkimGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                     `json:"dataSource,required"`
	Description string                                                                     `json:"description,required"`
	EndTime     time.Time                                                                  `json:"endTime,required" format:"date-time"`
	EventType   string                                                                     `json:"eventType,required"`
	StartTime   time.Time                                                                  `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopLocationDkimGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopLocationDkimGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailTopLocationDkimGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopLocationDkimGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationDkimGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDkimGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                   `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopLocationDkimGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopLocationDkimGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailTopLocationDkimGetResponseResultMetaDateRange]
type radarEmailTopLocationDkimGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationDkimGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDkimGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                             `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                             `json:"clientCountryName,required"`
	Value               string                                             `json:"value,required"`
	JSON                radarEmailTopLocationDkimGetResponseResultTop0JSON `json:"-"`
}

// radarEmailTopLocationDkimGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarEmailTopLocationDkimGetResponseResultTop0]
type radarEmailTopLocationDkimGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailTopLocationDkimGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationDkimGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTopLocationDkimGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopLocationDkimGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTopLocationDkimGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopLocationDkimGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTopLocationDkimGetParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTopLocationDkimGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopLocationDkimGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopLocationDkimGetParamsDkim string

const (
	RadarEmailTopLocationDkimGetParamsDkimPass RadarEmailTopLocationDkimGetParamsDkim = "PASS"
	RadarEmailTopLocationDkimGetParamsDkimNone RadarEmailTopLocationDkimGetParamsDkim = "NONE"
	RadarEmailTopLocationDkimGetParamsDkimFail RadarEmailTopLocationDkimGetParamsDkim = "FAIL"
)

type RadarEmailTopLocationDkimGetParamsArc string

const (
	RadarEmailTopLocationDkimGetParamsArcPass RadarEmailTopLocationDkimGetParamsArc = "PASS"
	RadarEmailTopLocationDkimGetParamsArcNone RadarEmailTopLocationDkimGetParamsArc = "NONE"
	RadarEmailTopLocationDkimGetParamsArcFail RadarEmailTopLocationDkimGetParamsArc = "FAIL"
)

type RadarEmailTopLocationDkimGetParamsDateRange string

const (
	RadarEmailTopLocationDkimGetParamsDateRange1d         RadarEmailTopLocationDkimGetParamsDateRange = "1d"
	RadarEmailTopLocationDkimGetParamsDateRange7d         RadarEmailTopLocationDkimGetParamsDateRange = "7d"
	RadarEmailTopLocationDkimGetParamsDateRange14d        RadarEmailTopLocationDkimGetParamsDateRange = "14d"
	RadarEmailTopLocationDkimGetParamsDateRange28d        RadarEmailTopLocationDkimGetParamsDateRange = "28d"
	RadarEmailTopLocationDkimGetParamsDateRange12w        RadarEmailTopLocationDkimGetParamsDateRange = "12w"
	RadarEmailTopLocationDkimGetParamsDateRange24w        RadarEmailTopLocationDkimGetParamsDateRange = "24w"
	RadarEmailTopLocationDkimGetParamsDateRange52w        RadarEmailTopLocationDkimGetParamsDateRange = "52w"
	RadarEmailTopLocationDkimGetParamsDateRange1dControl  RadarEmailTopLocationDkimGetParamsDateRange = "1dControl"
	RadarEmailTopLocationDkimGetParamsDateRange7dControl  RadarEmailTopLocationDkimGetParamsDateRange = "7dControl"
	RadarEmailTopLocationDkimGetParamsDateRange14dControl RadarEmailTopLocationDkimGetParamsDateRange = "14dControl"
	RadarEmailTopLocationDkimGetParamsDateRange28dControl RadarEmailTopLocationDkimGetParamsDateRange = "28dControl"
	RadarEmailTopLocationDkimGetParamsDateRange12wControl RadarEmailTopLocationDkimGetParamsDateRange = "12wControl"
	RadarEmailTopLocationDkimGetParamsDateRange24wControl RadarEmailTopLocationDkimGetParamsDateRange = "24wControl"
)

type RadarEmailTopLocationDkimGetParamsDmarc string

const (
	RadarEmailTopLocationDkimGetParamsDmarcPass RadarEmailTopLocationDkimGetParamsDmarc = "PASS"
	RadarEmailTopLocationDkimGetParamsDmarcNone RadarEmailTopLocationDkimGetParamsDmarc = "NONE"
	RadarEmailTopLocationDkimGetParamsDmarcFail RadarEmailTopLocationDkimGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTopLocationDkimGetParamsFormat string

const (
	RadarEmailTopLocationDkimGetParamsFormatJson RadarEmailTopLocationDkimGetParamsFormat = "JSON"
	RadarEmailTopLocationDkimGetParamsFormatCsv  RadarEmailTopLocationDkimGetParamsFormat = "CSV"
)

type RadarEmailTopLocationDkimGetParamsSpf string

const (
	RadarEmailTopLocationDkimGetParamsSpfPass RadarEmailTopLocationDkimGetParamsSpf = "PASS"
	RadarEmailTopLocationDkimGetParamsSpfNone RadarEmailTopLocationDkimGetParamsSpf = "NONE"
	RadarEmailTopLocationDkimGetParamsSpfFail RadarEmailTopLocationDkimGetParamsSpf = "FAIL"
)
