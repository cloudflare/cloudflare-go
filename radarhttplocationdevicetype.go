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

// RadarHTTPLocationDeviceTypeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPLocationDeviceTypeService] method instead.
type RadarHTTPLocationDeviceTypeService struct {
	Options []option.RequestOption
}

// NewRadarHTTPLocationDeviceTypeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPLocationDeviceTypeService(opts ...option.RequestOption) (r *RadarHTTPLocationDeviceTypeService) {
	r = &RadarHTTPLocationDeviceTypeService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested device type. Values are
// a percentage out of the total traffic.
func (r *RadarHTTPLocationDeviceTypeService) Get(ctx context.Context, deviceType RadarHTTPLocationDeviceTypeGetParamsDeviceType, query RadarHTTPLocationDeviceTypeGetParams, opts ...option.RequestOption) (res *RadarHTTPLocationDeviceTypeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPLocationDeviceTypeGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/locations/device_type/%v", deviceType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPLocationDeviceTypeGetResponse struct {
	Meta RadarHTTPLocationDeviceTypeGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPLocationDeviceTypeGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPLocationDeviceTypeGetResponseJSON   `json:"-"`
}

// radarHTTPLocationDeviceTypeGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPLocationDeviceTypeGetResponse]
type radarHTTPLocationDeviceTypeGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationDeviceTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationDeviceTypeGetResponseMeta struct {
	DateRange      []RadarHTTPLocationDeviceTypeGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPLocationDeviceTypeGetResponseMetaJSON           `json:"-"`
}

// radarHTTPLocationDeviceTypeGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarHTTPLocationDeviceTypeGetResponseMeta]
type radarHTTPLocationDeviceTypeGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPLocationDeviceTypeGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationDeviceTypeGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPLocationDeviceTypeGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPLocationDeviceTypeGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPLocationDeviceTypeGetResponseMetaDateRange]
type radarHTTPLocationDeviceTypeGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationDeviceTypeGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfo]
type radarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPLocationDeviceTypeGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationDeviceTypeGetResponseTop0 struct {
	ClientCountryAlpha2 string                                         `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                         `json:"clientCountryName,required"`
	Value               string                                         `json:"value,required"`
	JSON                radarHTTPLocationDeviceTypeGetResponseTop0JSON `json:"-"`
}

// radarHTTPLocationDeviceTypeGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarHTTPLocationDeviceTypeGetResponseTop0]
type radarHTTPLocationDeviceTypeGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPLocationDeviceTypeGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationDeviceTypeGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPLocationDeviceTypeGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPLocationDeviceTypeGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarHTTPLocationDeviceTypeGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPLocationDeviceTypeGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPLocationDeviceTypeGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPLocationDeviceTypeGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPLocationDeviceTypeGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPLocationDeviceTypeGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPLocationDeviceTypeGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPLocationDeviceTypeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Device type.
type RadarHTTPLocationDeviceTypeGetParamsDeviceType string

const (
	RadarHTTPLocationDeviceTypeGetParamsDeviceTypeDesktop RadarHTTPLocationDeviceTypeGetParamsDeviceType = "DESKTOP"
	RadarHTTPLocationDeviceTypeGetParamsDeviceTypeMobile  RadarHTTPLocationDeviceTypeGetParamsDeviceType = "MOBILE"
	RadarHTTPLocationDeviceTypeGetParamsDeviceTypeOther   RadarHTTPLocationDeviceTypeGetParamsDeviceType = "OTHER"
)

type RadarHTTPLocationDeviceTypeGetParamsBotClass string

const (
	RadarHTTPLocationDeviceTypeGetParamsBotClassLikelyAutomated RadarHTTPLocationDeviceTypeGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPLocationDeviceTypeGetParamsBotClassLikelyHuman     RadarHTTPLocationDeviceTypeGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPLocationDeviceTypeGetParamsDateRange string

const (
	RadarHTTPLocationDeviceTypeGetParamsDateRange1d         RadarHTTPLocationDeviceTypeGetParamsDateRange = "1d"
	RadarHTTPLocationDeviceTypeGetParamsDateRange2d         RadarHTTPLocationDeviceTypeGetParamsDateRange = "2d"
	RadarHTTPLocationDeviceTypeGetParamsDateRange7d         RadarHTTPLocationDeviceTypeGetParamsDateRange = "7d"
	RadarHTTPLocationDeviceTypeGetParamsDateRange14d        RadarHTTPLocationDeviceTypeGetParamsDateRange = "14d"
	RadarHTTPLocationDeviceTypeGetParamsDateRange28d        RadarHTTPLocationDeviceTypeGetParamsDateRange = "28d"
	RadarHTTPLocationDeviceTypeGetParamsDateRange12w        RadarHTTPLocationDeviceTypeGetParamsDateRange = "12w"
	RadarHTTPLocationDeviceTypeGetParamsDateRange24w        RadarHTTPLocationDeviceTypeGetParamsDateRange = "24w"
	RadarHTTPLocationDeviceTypeGetParamsDateRange52w        RadarHTTPLocationDeviceTypeGetParamsDateRange = "52w"
	RadarHTTPLocationDeviceTypeGetParamsDateRange1dControl  RadarHTTPLocationDeviceTypeGetParamsDateRange = "1dControl"
	RadarHTTPLocationDeviceTypeGetParamsDateRange2dControl  RadarHTTPLocationDeviceTypeGetParamsDateRange = "2dControl"
	RadarHTTPLocationDeviceTypeGetParamsDateRange7dControl  RadarHTTPLocationDeviceTypeGetParamsDateRange = "7dControl"
	RadarHTTPLocationDeviceTypeGetParamsDateRange14dControl RadarHTTPLocationDeviceTypeGetParamsDateRange = "14dControl"
	RadarHTTPLocationDeviceTypeGetParamsDateRange28dControl RadarHTTPLocationDeviceTypeGetParamsDateRange = "28dControl"
	RadarHTTPLocationDeviceTypeGetParamsDateRange12wControl RadarHTTPLocationDeviceTypeGetParamsDateRange = "12wControl"
	RadarHTTPLocationDeviceTypeGetParamsDateRange24wControl RadarHTTPLocationDeviceTypeGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarHTTPLocationDeviceTypeGetParamsFormat string

const (
	RadarHTTPLocationDeviceTypeGetParamsFormatJson RadarHTTPLocationDeviceTypeGetParamsFormat = "JSON"
	RadarHTTPLocationDeviceTypeGetParamsFormatCsv  RadarHTTPLocationDeviceTypeGetParamsFormat = "CSV"
)

type RadarHTTPLocationDeviceTypeGetParamsHTTPProtocol string

const (
	RadarHTTPLocationDeviceTypeGetParamsHTTPProtocolHTTP  RadarHTTPLocationDeviceTypeGetParamsHTTPProtocol = "HTTP"
	RadarHTTPLocationDeviceTypeGetParamsHTTPProtocolHTTPS RadarHTTPLocationDeviceTypeGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPLocationDeviceTypeGetParamsHTTPVersion string

const (
	RadarHTTPLocationDeviceTypeGetParamsHTTPVersionHttPv1 RadarHTTPLocationDeviceTypeGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPLocationDeviceTypeGetParamsHTTPVersionHttPv2 RadarHTTPLocationDeviceTypeGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPLocationDeviceTypeGetParamsHTTPVersionHttPv3 RadarHTTPLocationDeviceTypeGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPLocationDeviceTypeGetParamsIPVersion string

const (
	RadarHTTPLocationDeviceTypeGetParamsIPVersionIPv4 RadarHTTPLocationDeviceTypeGetParamsIPVersion = "IPv4"
	RadarHTTPLocationDeviceTypeGetParamsIPVersionIPv6 RadarHTTPLocationDeviceTypeGetParamsIPVersion = "IPv6"
)

type RadarHTTPLocationDeviceTypeGetParamsOS string

const (
	RadarHTTPLocationDeviceTypeGetParamsOSWindows  RadarHTTPLocationDeviceTypeGetParamsOS = "WINDOWS"
	RadarHTTPLocationDeviceTypeGetParamsOSMacosx   RadarHTTPLocationDeviceTypeGetParamsOS = "MACOSX"
	RadarHTTPLocationDeviceTypeGetParamsOSIos      RadarHTTPLocationDeviceTypeGetParamsOS = "IOS"
	RadarHTTPLocationDeviceTypeGetParamsOSAndroid  RadarHTTPLocationDeviceTypeGetParamsOS = "ANDROID"
	RadarHTTPLocationDeviceTypeGetParamsOSChromeos RadarHTTPLocationDeviceTypeGetParamsOS = "CHROMEOS"
	RadarHTTPLocationDeviceTypeGetParamsOSLinux    RadarHTTPLocationDeviceTypeGetParamsOS = "LINUX"
	RadarHTTPLocationDeviceTypeGetParamsOSSmartTv  RadarHTTPLocationDeviceTypeGetParamsOS = "SMART_TV"
)

type RadarHTTPLocationDeviceTypeGetParamsTLSVersion string

const (
	RadarHTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_0  RadarHTTPLocationDeviceTypeGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_1  RadarHTTPLocationDeviceTypeGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_2  RadarHTTPLocationDeviceTypeGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPLocationDeviceTypeGetParamsTLSVersionTlSv1_3  RadarHTTPLocationDeviceTypeGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPLocationDeviceTypeGetParamsTLSVersionTlSvQuic RadarHTTPLocationDeviceTypeGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPLocationDeviceTypeGetResponseEnvelope struct {
	Result  RadarHTTPLocationDeviceTypeGetResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarHTTPLocationDeviceTypeGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPLocationDeviceTypeGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarHTTPLocationDeviceTypeGetResponseEnvelope]
type radarHTTPLocationDeviceTypeGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationDeviceTypeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
