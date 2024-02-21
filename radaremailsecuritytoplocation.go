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
	Options   []option.RequestOption
	Arc       *RadarEmailSecurityTopLocationArcService
	DKIM      *RadarEmailSecurityTopLocationDKIMService
	Dmarc     *RadarEmailSecurityTopLocationDmarcService
	Malicious *RadarEmailSecurityTopLocationMaliciousService
	Spam      *RadarEmailSecurityTopLocationSpamService
	SPF       *RadarEmailSecurityTopLocationSPFService
}

// NewRadarEmailSecurityTopLocationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarEmailSecurityTopLocationService(opts ...option.RequestOption) (r *RadarEmailSecurityTopLocationService) {
	r = &RadarEmailSecurityTopLocationService{}
	r.Options = opts
	r.Arc = NewRadarEmailSecurityTopLocationArcService(opts...)
	r.DKIM = NewRadarEmailSecurityTopLocationDKIMService(opts...)
	r.Dmarc = NewRadarEmailSecurityTopLocationDmarcService(opts...)
	r.Malicious = NewRadarEmailSecurityTopLocationMaliciousService(opts...)
	r.Spam = NewRadarEmailSecurityTopLocationSpamService(opts...)
	r.SPF = NewRadarEmailSecurityTopLocationSPFService(opts...)
	return
}

// Get the top locations by email messages. Values are a percentage out of the
// total emails.
func (r *RadarEmailSecurityTopLocationService) Get(ctx context.Context, query RadarEmailSecurityTopLocationGetParams, opts ...option.RequestOption) (res *RadarEmailSecurityTopLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarEmailSecurityTopLocationGetResponseEnvelope
	path := "radar/email/security/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarEmailSecurityTopLocationGetResponse struct {
	Meta RadarEmailSecurityTopLocationGetResponseMeta   `json:"meta,required"`
	Top0 []RadarEmailSecurityTopLocationGetResponseTop0 `json:"top_0,required"`
	JSON radarEmailSecurityTopLocationGetResponseJSON   `json:"-"`
}

// radarEmailSecurityTopLocationGetResponseJSON contains the JSON metadata for the
// struct [RadarEmailSecurityTopLocationGetResponse]
type radarEmailSecurityTopLocationGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationGetResponseMeta struct {
	DateRange      []RadarEmailSecurityTopLocationGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	ConfidenceInfo RadarEmailSecurityTopLocationGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarEmailSecurityTopLocationGetResponseMetaJSON           `json:"-"`
}

// radarEmailSecurityTopLocationGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopLocationGetResponseMeta]
type radarEmailSecurityTopLocationGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarEmailSecurityTopLocationGetResponseMetaDateRangeJSON `json:"-"`
}

// radarEmailSecurityTopLocationGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarEmailSecurityTopLocationGetResponseMetaDateRange]
type radarEmailSecurityTopLocationGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationGetResponseMetaConfidenceInfo struct {
	Annotations []RadarEmailSecurityTopLocationGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarEmailSecurityTopLocationGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarEmailSecurityTopLocationGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarEmailSecurityTopLocationGetResponseMetaConfidenceInfo]
type radarEmailSecurityTopLocationGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarEmailSecurityTopLocationGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarEmailSecurityTopLocationGetResponseMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarEmailSecurityTopLocationGetResponseMetaConfidenceInfoAnnotation]
type radarEmailSecurityTopLocationGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarEmailSecurityTopLocationGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationGetResponseTop0 struct {
	ClientCountryAlpha2 string                                           `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                           `json:"clientCountryName,required"`
	Value               string                                           `json:"value,required"`
	JSON                radarEmailSecurityTopLocationGetResponseTop0JSON `json:"-"`
}

// radarEmailSecurityTopLocationGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarEmailSecurityTopLocationGetResponseTop0]
type radarEmailSecurityTopLocationGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarEmailSecurityTopLocationGetParams struct {
	// Filter for arc (Authenticated Received Chain).
	Arc param.Field[[]RadarEmailSecurityTopLocationGetParamsArc] `query:"arc"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarEmailSecurityTopLocationGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for dkim.
	DKIM param.Field[[]RadarEmailSecurityTopLocationGetParamsDKIM] `query:"dkim"`
	// Filter for dmarc.
	Dmarc param.Field[[]RadarEmailSecurityTopLocationGetParamsDmarc] `query:"dmarc"`
	// Format results are returned in.
	Format param.Field[RadarEmailSecurityTopLocationGetParamsFormat] `query:"format"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for spf.
	SPF param.Field[[]RadarEmailSecurityTopLocationGetParamsSPF] `query:"spf"`
}

// URLQuery serializes [RadarEmailSecurityTopLocationGetParams]'s query parameters
// as `url.Values`.
func (r RadarEmailSecurityTopLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarEmailSecurityTopLocationGetParamsArc string

const (
	RadarEmailSecurityTopLocationGetParamsArcPass RadarEmailSecurityTopLocationGetParamsArc = "PASS"
	RadarEmailSecurityTopLocationGetParamsArcNone RadarEmailSecurityTopLocationGetParamsArc = "NONE"
	RadarEmailSecurityTopLocationGetParamsArcFail RadarEmailSecurityTopLocationGetParamsArc = "FAIL"
)

type RadarEmailSecurityTopLocationGetParamsDateRange string

const (
	RadarEmailSecurityTopLocationGetParamsDateRange1d         RadarEmailSecurityTopLocationGetParamsDateRange = "1d"
	RadarEmailSecurityTopLocationGetParamsDateRange2d         RadarEmailSecurityTopLocationGetParamsDateRange = "2d"
	RadarEmailSecurityTopLocationGetParamsDateRange7d         RadarEmailSecurityTopLocationGetParamsDateRange = "7d"
	RadarEmailSecurityTopLocationGetParamsDateRange14d        RadarEmailSecurityTopLocationGetParamsDateRange = "14d"
	RadarEmailSecurityTopLocationGetParamsDateRange28d        RadarEmailSecurityTopLocationGetParamsDateRange = "28d"
	RadarEmailSecurityTopLocationGetParamsDateRange12w        RadarEmailSecurityTopLocationGetParamsDateRange = "12w"
	RadarEmailSecurityTopLocationGetParamsDateRange24w        RadarEmailSecurityTopLocationGetParamsDateRange = "24w"
	RadarEmailSecurityTopLocationGetParamsDateRange52w        RadarEmailSecurityTopLocationGetParamsDateRange = "52w"
	RadarEmailSecurityTopLocationGetParamsDateRange1dControl  RadarEmailSecurityTopLocationGetParamsDateRange = "1dControl"
	RadarEmailSecurityTopLocationGetParamsDateRange2dControl  RadarEmailSecurityTopLocationGetParamsDateRange = "2dControl"
	RadarEmailSecurityTopLocationGetParamsDateRange7dControl  RadarEmailSecurityTopLocationGetParamsDateRange = "7dControl"
	RadarEmailSecurityTopLocationGetParamsDateRange14dControl RadarEmailSecurityTopLocationGetParamsDateRange = "14dControl"
	RadarEmailSecurityTopLocationGetParamsDateRange28dControl RadarEmailSecurityTopLocationGetParamsDateRange = "28dControl"
	RadarEmailSecurityTopLocationGetParamsDateRange12wControl RadarEmailSecurityTopLocationGetParamsDateRange = "12wControl"
	RadarEmailSecurityTopLocationGetParamsDateRange24wControl RadarEmailSecurityTopLocationGetParamsDateRange = "24wControl"
)

type RadarEmailSecurityTopLocationGetParamsDKIM string

const (
	RadarEmailSecurityTopLocationGetParamsDKIMPass RadarEmailSecurityTopLocationGetParamsDKIM = "PASS"
	RadarEmailSecurityTopLocationGetParamsDKIMNone RadarEmailSecurityTopLocationGetParamsDKIM = "NONE"
	RadarEmailSecurityTopLocationGetParamsDKIMFail RadarEmailSecurityTopLocationGetParamsDKIM = "FAIL"
)

type RadarEmailSecurityTopLocationGetParamsDmarc string

const (
	RadarEmailSecurityTopLocationGetParamsDmarcPass RadarEmailSecurityTopLocationGetParamsDmarc = "PASS"
	RadarEmailSecurityTopLocationGetParamsDmarcNone RadarEmailSecurityTopLocationGetParamsDmarc = "NONE"
	RadarEmailSecurityTopLocationGetParamsDmarcFail RadarEmailSecurityTopLocationGetParamsDmarc = "FAIL"
)

// Format results are returned in.
type RadarEmailSecurityTopLocationGetParamsFormat string

const (
	RadarEmailSecurityTopLocationGetParamsFormatJson RadarEmailSecurityTopLocationGetParamsFormat = "JSON"
	RadarEmailSecurityTopLocationGetParamsFormatCsv  RadarEmailSecurityTopLocationGetParamsFormat = "CSV"
)

type RadarEmailSecurityTopLocationGetParamsSPF string

const (
	RadarEmailSecurityTopLocationGetParamsSPFPass RadarEmailSecurityTopLocationGetParamsSPF = "PASS"
	RadarEmailSecurityTopLocationGetParamsSPFNone RadarEmailSecurityTopLocationGetParamsSPF = "NONE"
	RadarEmailSecurityTopLocationGetParamsSPFFail RadarEmailSecurityTopLocationGetParamsSPF = "FAIL"
)

type RadarEmailSecurityTopLocationGetResponseEnvelope struct {
	Result  RadarEmailSecurityTopLocationGetResponse             `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarEmailSecurityTopLocationGetResponseEnvelopeJSON `json:"-"`
}

// radarEmailSecurityTopLocationGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarEmailSecurityTopLocationGetResponseEnvelope]
type radarEmailSecurityTopLocationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarEmailSecurityTopLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
