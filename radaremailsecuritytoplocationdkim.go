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

// RadarEmailSecurityTopLocationDKIMService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarEmailSecurityTopLocationDKIMService] method instead.
type RadarEmailSecurityTopLocationDKIMService struct {
	Options []option.RequestOption
}

// NewRadarEmailSecurityTopLocationDKIMService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopLocationDKIMService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationDKIMService) {
	r = &RadarEmailSecurityTopLocationDKIMService{}
	r.Options = opts
	return
}

// Get the locations, by email DKIM validation.
func (r *RadarEmailSecurityTopLocationDKIMService) Get(ctx context.Context, dkim RadarEmailSecurityTopLocationDKIMGetParamsDKIM, query RadarEmailSecurityTopLocationDKIMGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationDKIMGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopLocationDKIMGetResponseEnvelope
	path := fmt.Sprintf("radar/email/security/top/locations/dkim/%v", dkim)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopLocationDKIMGetResponse struct {
	Meta RadarEmailSecurityTopLocationDKIMGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationDKIMGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationDKIMGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationDKIMGetResponseJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopLocationDKIMGetResponse]
type radarEmailSecurityTopLocationDKIMGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDKIMGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDKIMGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopLocationDKIMGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationDKIMGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationDKIMGetResponseMetaJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationDKIMGetResponseMeta]
type radarEmailSecurityTopLocationDKIMGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDKIMGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDKIMGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationDKIMGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationDKIMGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationDKIMGetResponseMetaDateRange]
type radarEmailSecurityTopLocationDKIMGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDKIMGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationDKIMGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDKIMGetResponseTop0 struct {
	ClientCountryAlpha2 string                                               `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                               `json:"clientCountryName,required"`
	Value               string                                               `json:"value,required"`
	JSON                radarEmailSecurityTopLocationDKIMGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationDKIMGetResponseTop0JSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationDKIMGetResponseTop0]
type radarEmailSecurityTopLocationDKIMGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDKIMGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationDKIMGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopLocationDKIMGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationDKIMGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopLocationDKIMGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationDKIMGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopLocationDKIMGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationDKIMGetParams]'s query
// parameters as `url.Values`.
func (r RadarEmailSecurityTopLocationDKIMGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// DKIM.
type RadarEmailSecurityTopLocationDKIMGetParamsDKIM string

const (
	RadarEmailSecurityTopLocationDKIMGetParamsDKIMPass RadarEmailSecurityTopLocationDKIMGetParamsDKIM = "PASS"
	RadarEmailSecurityTopLocationDKIMGetParamsDKIMNone RadarEmailSecurityTopLocationDKIMGetParamsDKIM = "NONE"
	RadarEmailSecurityTopLocationDKIMGetParamsDKIMFail RadarEmailSecurityTopLocationDKIMGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopLocationDKIMGetParamsArc string

const (
	RadarEmailSecurityTopLocationDKIMGetParamsArcPass RadarEmailSecurityTopLocationDKIMGetParamsArc = "PASS"
	RadarEmailSecurityTopLocationDKIMGetParamsArcNone RadarEmailSecurityTopLocationDKIMGetParamsArc = "NONE"
	RadarEmailSecurityTopLocationDKIMGetParamsArcFail RadarEmailSecurityTopLocationDKIMGetParamsArc = "FAIL"
)

type RadarEmailSecurityTopLocationDKIMGetParamsDateRange string

const (
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange1d         RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "1d"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange2d         RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "2d"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange7d         RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "7d"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange14d        RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "14d"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange28d        RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "28d"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange12w        RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "12w"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange24w        RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "24w"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange52w        RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "52w"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange1dControl  RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange2dControl  RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange7dControl  RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange14dControl RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange28dControl RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange12wControl RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationDKIMGetParamsDateRange24wControl RadarEmailSecurityTopLocationDKIMGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationDKIMGetParamsDmarc string

const (
	RadarEmailSecurityTopLocationDKIMGetParamsDmarcPass RadarEmailSecurityTopLocationDKIMGetParamsDmarc = "PASS"
	RadarEmailSecurityTopLocationDKIMGetParamsDmarcNone RadarEmailSecurityTopLocationDKIMGetParamsDmarc = "NONE"
	RadarEmailSecurityTopLocationDKIMGetParamsDmarcFail RadarEmailSecurityTopLocationDKIMGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationDKIMGetParamsFormat string

const (
	RadarEmailSecurityTopLocationDKIMGetParamsFormatJson RadarEmailSecurityTopLocationDKIMGetParamsFormat = "JSON"
	RadarEmailSecurityTopLocationDKIMGetParamsFormatCsv  RadarEmailSecurityTopLocationDKIMGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationDKIMGetParamsSPF string

const (
	RadarEmailSecurityTopLocationDKIMGetParamsSPFPass RadarEmailSecurityTopLocationDKIMGetParamsSPF = "PASS"
	RadarEmailSecurityTopLocationDKIMGetParamsSPFNone RadarEmailSecurityTopLocationDKIMGetParamsSPF = "NONE"
	RadarEmailSecurityTopLocationDKIMGetParamsSPFFail RadarEmailSecurityTopLocationDKIMGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopLocationDKIMGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopLocationDKIMGetResponse             `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarEmailSecurityTopLocationDKIMGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopLocationDKIMGetResponseEnvelopeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationDKIMGetResponseEnvelope]
type radarEmailSecurityTopLocationDKIMGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationDKIMGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
