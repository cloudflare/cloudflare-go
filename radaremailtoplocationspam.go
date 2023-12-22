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

// RadarEmailTopLocationSpamService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarEmailTopLocationSpamService] method instead.
type RadarEmailTopLocationSpamService struct {
	Options []option.RequestOption
}

// NewRadarEmailTopLocationSpamService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailTopLocationSpamService(opts ...option.RequestOption) (r *RadarEmailTopLocationSpamService) {
	r = &RadarEmailTopLocationSpamService{}
	r.Options = opts
	return
}

// Get the locations, by emails classified, of Spam validations.
func (r *RadarEmailTopLocationSpamService) Get(ctx context.Context, spam RadarEmailTopLocationSpamGetParamsSpam, query RadarEmailTopLocationSpamGetParams, opts ...option.RequestOption) (res *RadarEmailTopLocationSpamGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/email/top/locations/spam/%v", spam)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarEmailTopLocationSpamGetResponse struct {
	Result  RadarEmailTopLocationSpamGetResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarEmailTopLocationSpamGetResponseJSON   `json:"-"`
}

// radarEmailTopLocationSpamGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailTopLocationSpamGetResponse]
type radarEmailTopLocationSpamGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpamGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpamGetResponseResult struct {
	Meta RadarEmailTopLocationSpamGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarEmailTopLocationSpamGetResponseResultTop0 `json:"top_0,required"`
	JSON radarEmailTopLocationSpamGetResponseResultJSON   `json:"-"`
}

// radarEmailTopLocationSpamGetResponseResultJSON contains the JSON metadata for
// the struct [RadarEmailTopLocationSpamGetResponseResult]
type radarEmailTopLocationSpamGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpamGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpamGetResponseResultMeta struct {
	ConfidenceInfo RadarEmailTopLocationSpamGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarEmailTopLocationSpamGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarEmailTopLocationSpamGetResponseResultMetaJSON           `json:"-"`
}

// radarEmailTopLocationSpamGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarEmailTopLocationSpamGetResponseResultMeta]
type radarEmailTopLocationSpamGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpamGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpamGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarEmailTopLocationSpamGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                    `json:"level,required"`
	JSON        radarEmailTopLocationSpamGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailTopLocationSpamGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailTopLocationSpamGetResponseResultMetaConfidenceInfo]
type radarEmailTopLocationSpamGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpamGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpamGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                     `json:"dataSource,required"`
	Description string                                                                     `json:"description,required"`
	EndTime     time.Time                                                                  `json:"endTime,required" format:"date-time"`
	EventType   string                                                                     `json:"eventType,required"`
	StartTime   time.Time                                                                  `json:"startTime,required" format:"date-time"`
	JSON        radarEmailTopLocationSpamGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailTopLocationSpamGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailTopLocationSpamGetResponseResultMetaConfidenceInfoAnnotation]
type radarEmailTopLocationSpamGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpamGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpamGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                   `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarEmailTopLocationSpamGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarEmailTopLocationSpamGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailTopLocationSpamGetResponseResultMetaDateRange]
type radarEmailTopLocationSpamGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpamGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpamGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                             `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                             `json:"clientCountryName,required"`
	Value               string                                             `json:"value,required"`
	JSON                radarEmailTopLocationSpamGetResponseResultTop0JSON `json:"-"`
}

// radarEmailTopLocationSpamGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarEmailTopLocationSpamGetResponseResultTop0]
type radarEmailTopLocationSpamGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailTopLocationSpamGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailTopLocationSpamGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailTopLocationSpamGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailTopLocationSpamGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	Dkim param.Field[[]RadarEmailTopLocationSpamGetParamsDkim] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailTopLocationSpamGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailTopLocationSpamGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	Spf param.Field[[]RadarEmailTopLocationSpamGetParamsSpf] `query:"spf"`
}

// URLQuery serializes [RadarEmailTopLocationSpamGetParams]'s query parameters as
// `url.Values`.
func (r RadarEmailTopLocationSpamGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailTopLocationSpamGetParamsSpam string

const (
	RadarEmailTopLocationSpamGetParamsSpamSpam    RadarEmailTopLocationSpamGetParamsSpam = "SPAM"
	RadarEmailTopLocationSpamGetParamsSpamNotSpam RadarEmailTopLocationSpamGetParamsSpam = "NOT_SPAM"
)

type RadarEmailTopLocationSpamGetParamsArc string

const (
	RadarEmailTopLocationSpamGetParamsArcPass RadarEmailTopLocationSpamGetParamsArc = "PASS"
	RadarEmailTopLocationSpamGetParamsArcNone RadarEmailTopLocationSpamGetParamsArc = "NONE"
	RadarEmailTopLocationSpamGetParamsArcFail RadarEmailTopLocationSpamGetParamsArc = "FAIL"
)

type RadarEmailTopLocationSpamGetParamsDateRange string

const (
	RadarEmailTopLocationSpamGetParamsDateRange1d         RadarEmailTopLocationSpamGetParamsDateRange = "1d"
	RadarEmailTopLocationSpamGetParamsDateRange7d         RadarEmailTopLocationSpamGetParamsDateRange = "7d"
	RadarEmailTopLocationSpamGetParamsDateRange14d        RadarEmailTopLocationSpamGetParamsDateRange = "14d"
	RadarEmailTopLocationSpamGetParamsDateRange28d        RadarEmailTopLocationSpamGetParamsDateRange = "28d"
	RadarEmailTopLocationSpamGetParamsDateRange12w        RadarEmailTopLocationSpamGetParamsDateRange = "12w"
	RadarEmailTopLocationSpamGetParamsDateRange24w        RadarEmailTopLocationSpamGetParamsDateRange = "24w"
	RadarEmailTopLocationSpamGetParamsDateRange52w        RadarEmailTopLocationSpamGetParamsDateRange = "52w"
	RadarEmailTopLocationSpamGetParamsDateRange1dControl  RadarEmailTopLocationSpamGetParamsDateRange = "1dControl"
	RadarEmailTopLocationSpamGetParamsDateRange7dControl  RadarEmailTopLocationSpamGetParamsDateRange = "7dControl"
	RadarEmailTopLocationSpamGetParamsDateRange14dControl RadarEmailTopLocationSpamGetParamsDateRange = "14dControl"
	RadarEmailTopLocationSpamGetParamsDateRange28dControl RadarEmailTopLocationSpamGetParamsDateRange = "28dControl"
	RadarEmailTopLocationSpamGetParamsDateRange12wControl RadarEmailTopLocationSpamGetParamsDateRange = "12wControl"
	RadarEmailTopLocationSpamGetParamsDateRange24wControl RadarEmailTopLocationSpamGetParamsDateRange = "24wControl"
)

type RadarEmailTopLocationSpamGetParamsDkim string

const (
	RadarEmailTopLocationSpamGetParamsDkimPass RadarEmailTopLocationSpamGetParamsDkim = "PASS"
	RadarEmailTopLocationSpamGetParamsDkimNone RadarEmailTopLocationSpamGetParamsDkim = "NONE"
	RadarEmailTopLocationSpamGetParamsDkimFail RadarEmailTopLocationSpamGetParamsDkim = "FAIL"
)

type RadarEmailTopLocationSpamGetParamsDmarc string

const (
	RadarEmailTopLocationSpamGetParamsDmarcPass RadarEmailTopLocationSpamGetParamsDmarc = "PASS"
	RadarEmailTopLocationSpamGetParamsDmarcNone RadarEmailTopLocationSpamGetParamsDmarc = "NONE"
	RadarEmailTopLocationSpamGetParamsDmarcFail RadarEmailTopLocationSpamGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailTopLocationSpamGetParamsFormat string

const (
	RadarEmailTopLocationSpamGetParamsFormatJson RadarEmailTopLocationSpamGetParamsFormat = "JSON"
	RadarEmailTopLocationSpamGetParamsFormatCsv  RadarEmailTopLocationSpamGetParamsFormat = "CSV"
)

type RadarEmailTopLocationSpamGetParamsSpf string

const (
	RadarEmailTopLocationSpamGetParamsSpfPass RadarEmailTopLocationSpamGetParamsSpf = "PASS"
	RadarEmailTopLocationSpamGetParamsSpfNone RadarEmailTopLocationSpamGetParamsSpf = "NONE"
	RadarEmailTopLocationSpamGetParamsSpfFail RadarEmailTopLocationSpamGetParamsSpf = "FAIL"
)
