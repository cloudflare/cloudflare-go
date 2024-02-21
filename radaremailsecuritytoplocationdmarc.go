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

// RadarEmailSecurityTopLocationDmarcService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationDmarcService] method instead.
type RadarEmailSecurityTopLocationDmarcService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationDmarcService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarEmailSecurityTopLocationDmarcService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationDmarcService) {
	r = &RadarEmailSecurityTopLocationDmarcService{}
	r.Options = opts
	return
}

// Get the locations by email DMARC validation.
func (r *RadarEmailSecurityTopLocationDmarcService) Get(ctx context.Context, dmarc RadarEmailSecurityTopLocationDmarcGetParamsDmarc, query RadarEmailSecurityTopLocationDmarcGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationDmarcGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopLocationDmarcGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/locations/dmarc/%v", dmarc)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopLocationDmarcGetResponse struct {
	Meta RadarEmailSecurityTopLocationDmarcGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationDmarcGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationDmarcGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationDmarcGetResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopLocationDmarcGetResponse]
type radarEmailSecurityTopLocationDmarcGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDmarcGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDmarcGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopLocationDmarcGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                          `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationDmarcGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationDmarcGetResponseMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationDmarcGetResponseMeta]
type radarEmailSecurityTopLocationDmarcGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDmarcGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDmarcGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationDmarcGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationDmarcGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationDmarcGetResponseMetaDateRange]
type radarEmailSecurityTopLocationDmarcGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDmarcGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                       `json:"level"`
	JSON        radarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                        `json:"dataSource,required"`
	Description     string                                                                        `json:"description,required"`
	EventType       string                                                                        `json:"eventType,required"`
	IsInstantaneous interface{}                                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                     `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                        `json:"linkedUrl"`
	StartTime       time.Time                                                                     `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationDmarcGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDmarcGetResponseTop0 struct {
	ClientCountryAlpha2 string                                                `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                `json:"clientCountryName,required"`
	Value               string                                                `json:"value,required"`
	JSON                radarEmailSecurityTopLocationDmarcGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationDmarcGetResponseTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationDmarcGetResponseTop0]
type radarEmailSecurityTopLocationDmarcGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDmarcGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDmarcGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopLocationDmarcGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationDmarcGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopLocationDmarcGetParamsDKIM] `query:"dkim"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationDmarcGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopLocationDmarcGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationDmarcGetParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationDmarcGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DMARC.
type RadarEmailSecurityTopLocationDmarcGetParamsDmarc string

const (
	RadarEmailSecurityTopLocationDmarcGetParamsDmarcPass RadarEmailSecurityTopLocationDmarcGetParamsDmarc = "PASS"
	RadarEmailSecurityTopLocationDmarcGetParamsDmarcNone RadarEmailSecurityTopLocationDmarcGetParamsDmarc = "NONE"
	RadarEmailSecurityTopLocationDmarcGetParamsDmarcFail RadarEmailSecurityTopLocationDmarcGetParamsDmarc = "FAIL"
)

type RadarEmailSecurityTopLocationDmarcGetParamsArc string

const (
	RadarEmailSecurityTopLocationDmarcGetParamsArcPass RadarEmailSecurityTopLocationDmarcGetParamsArc = "PASS"
	RadarEmailSecurityTopLocationDmarcGetParamsArcNone RadarEmailSecurityTopLocationDmarcGetParamsArc = "NONE"
	RadarEmailSecurityTopLocationDmarcGetParamsArcFail RadarEmailSecurityTopLocationDmarcGetParamsArc = "FAIL"
)

type RadarEmailSecurityTopLocationDmarcGetParamsDateRange string

const (
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange1d         RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "1d"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange2d         RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "2d"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange7d         RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "7d"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange14d        RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "14d"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange28d        RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "28d"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange12w        RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "12w"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange24w        RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "24w"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange52w        RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "52w"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange1dControl  RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange2dControl  RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange7dControl  RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange14dControl RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange28dControl RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange12wControl RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationDmarcGetParamsDateRange24wControl RadarEmailSecurityTopLocationDmarcGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationDmarcGetParamsDKIM string

const (
	RadarEmailSecurityTopLocationDmarcGetParamsDKIMPass RadarEmailSecurityTopLocationDmarcGetParamsDKIM = "PASS"
	RadarEmailSecurityTopLocationDmarcGetParamsDKIMNone RadarEmailSecurityTopLocationDmarcGetParamsDKIM = "NONE"
	RadarEmailSecurityTopLocationDmarcGetParamsDKIMFail RadarEmailSecurityTopLocationDmarcGetParamsDKIM = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationDmarcGetParamsFormat string

const (
	RadarEmailSecurityTopLocationDmarcGetParamsFormatJson RadarEmailSecurityTopLocationDmarcGetParamsFormat = "JSON"
	RadarEmailSecurityTopLocationDmarcGetParamsFormatCsv  RadarEmailSecurityTopLocationDmarcGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationDmarcGetParamsSPF string

const (
	RadarEmailSecurityTopLocationDmarcGetParamsSPFPass RadarEmailSecurityTopLocationDmarcGetParamsSPF = "PASS"
	RadarEmailSecurityTopLocationDmarcGetParamsSPFNone RadarEmailSecurityTopLocationDmarcGetParamsSPF = "NONE"
	RadarEmailSecurityTopLocationDmarcGetParamsSPFFail RadarEmailSecurityTopLocationDmarcGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopLocationDmarcGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopLocationDmarcGetResponse             `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarEmailSecurityTopLocationDmarcGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopLocationDmarcGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationDmarcGetResponseEnvelope]
type radarEmailSecurityTopLocationDmarcGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDmarcGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
