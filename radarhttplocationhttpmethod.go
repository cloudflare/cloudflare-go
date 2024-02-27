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

// RadarHTTPLocationHTTPMethodService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPLocationHTTPMethodService] method instead.
type RadarHTTPLocationHTTPMethodService struct {
	Options []option.RequestOption
}

// NewRadarHTTPLocationHTTPMethodService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPLocationHTTPMethodService(opts ...option.RequestOption) (r *RadarHTTPLocationHTTPMethodService) {
	r = &RadarHTTPLocationHTTPMethodService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested HTTP protocol. Values
// are a percentage out of the total traffic.
func (r *RadarHTTPLocationHTTPMethodService) Get(ctx context.Context, httpVersion RadarHTTPLocationHTTPMethodGetParamsHTTPVersion, query RadarHTTPLocationHTTPMethodGetParams, opts ...option.RequestOption) (res *RadarHTTPLocationHTTPMethodGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPLocationHTTPMethodGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/locations/http_version/%v", httpVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPLocationHTTPMethodGetResponse struct {
	Meta RadarHTTPLocationHTTPMethodGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPLocationHTTPMethodGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPLocationHTTPMethodGetResponseJSON   `json:"-"`
}

// radarHTTPLocationHTTPMethodGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPLocationHTTPMethodGetResponse]
type radarHTTPLocationHTTPMethodGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationHTTPMethodGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationHTTPMethodGetResponseMeta struct {
	DateRange      []RadarHTTPLocationHTTPMethodGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPLocationHTTPMethodGetResponseMetaJSON           `json:"-"`
}

// radarHTTPLocationHTTPMethodGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarHTTPLocationHTTPMethodGetResponseMeta]
type radarHTTPLocationHTTPMethodGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPLocationHTTPMethodGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationHTTPMethodGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPLocationHTTPMethodGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPLocationHTTPMethodGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPLocationHTTPMethodGetResponseMetaDateRange]
type radarHTTPLocationHTTPMethodGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationHTTPMethodGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfo]
type radarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationHTTPMethodGetResponseTop0 struct {
	ClientCountryAlpha2 string                                         `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                         `json:"clientCountryName,required"`
	Value               string                                         `json:"value,required"`
	JSON                radarHTTPLocationHTTPMethodGetResponseTop0JSON `json:"-"`
}

// radarHTTPLocationHTTPMethodGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarHTTPLocationHTTPMethodGetResponseTop0]
type radarHTTPLocationHTTPMethodGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPLocationHTTPMethodGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationHTTPMethodGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPLocationHTTPMethodGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPLocationHTTPMethodGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPLocationHTTPMethodGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPLocationHTTPMethodGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPLocationHTTPMethodGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPLocationHTTPMethodGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPLocationHTTPMethodGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPLocationHTTPMethodGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPLocationHTTPMethodGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPLocationHTTPMethodGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// HTTP version.
type RadarHTTPLocationHTTPMethodGetParamsHTTPVersion string

const (
	RadarHTTPLocationHTTPMethodGetParamsHTTPVersionHttPv1 RadarHTTPLocationHTTPMethodGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPLocationHTTPMethodGetParamsHTTPVersionHttPv2 RadarHTTPLocationHTTPMethodGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPLocationHTTPMethodGetParamsHTTPVersionHttPv3 RadarHTTPLocationHTTPMethodGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPLocationHTTPMethodGetParamsBotClass string

const (
	RadarHTTPLocationHTTPMethodGetParamsBotClassLikelyAutomated RadarHTTPLocationHTTPMethodGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPLocationHTTPMethodGetParamsBotClassLikelyHuman     RadarHTTPLocationHTTPMethodGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPLocationHTTPMethodGetParamsDateRange string

const (
	RadarHTTPLocationHTTPMethodGetParamsDateRange1d         RadarHTTPLocationHTTPMethodGetParamsDateRange = "1d"
	RadarHTTPLocationHTTPMethodGetParamsDateRange2d         RadarHTTPLocationHTTPMethodGetParamsDateRange = "2d"
	RadarHTTPLocationHTTPMethodGetParamsDateRange7d         RadarHTTPLocationHTTPMethodGetParamsDateRange = "7d"
	RadarHTTPLocationHTTPMethodGetParamsDateRange14d        RadarHTTPLocationHTTPMethodGetParamsDateRange = "14d"
	RadarHTTPLocationHTTPMethodGetParamsDateRange28d        RadarHTTPLocationHTTPMethodGetParamsDateRange = "28d"
	RadarHTTPLocationHTTPMethodGetParamsDateRange12w        RadarHTTPLocationHTTPMethodGetParamsDateRange = "12w"
	RadarHTTPLocationHTTPMethodGetParamsDateRange24w        RadarHTTPLocationHTTPMethodGetParamsDateRange = "24w"
	RadarHTTPLocationHTTPMethodGetParamsDateRange52w        RadarHTTPLocationHTTPMethodGetParamsDateRange = "52w"
	RadarHTTPLocationHTTPMethodGetParamsDateRange1dControl  RadarHTTPLocationHTTPMethodGetParamsDateRange = "1dControl"
	RadarHTTPLocationHTTPMethodGetParamsDateRange2dControl  RadarHTTPLocationHTTPMethodGetParamsDateRange = "2dControl"
	RadarHTTPLocationHTTPMethodGetParamsDateRange7dControl  RadarHTTPLocationHTTPMethodGetParamsDateRange = "7dControl"
	RadarHTTPLocationHTTPMethodGetParamsDateRange14dControl RadarHTTPLocationHTTPMethodGetParamsDateRange = "14dControl"
	RadarHTTPLocationHTTPMethodGetParamsDateRange28dControl RadarHTTPLocationHTTPMethodGetParamsDateRange = "28dControl"
	RadarHTTPLocationHTTPMethodGetParamsDateRange12wControl RadarHTTPLocationHTTPMethodGetParamsDateRange = "12wControl"
	RadarHTTPLocationHTTPMethodGetParamsDateRange24wControl RadarHTTPLocationHTTPMethodGetParamsDateRange = "24wControl"
)

type RadarHTTPLocationHTTPMethodGetParamsDeviceType string

const (
	RadarHTTPLocationHTTPMethodGetParamsDeviceTypeDesktop RadarHTTPLocationHTTPMethodGetParamsDeviceType = "DESKTOP"
	RadarHTTPLocationHTTPMethodGetParamsDeviceTypeMobile  RadarHTTPLocationHTTPMethodGetParamsDeviceType = "MOBILE"
	RadarHTTPLocationHTTPMethodGetParamsDeviceTypeOther   RadarHTTPLocationHTTPMethodGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPLocationHTTPMethodGetParamsFormat string

const (
	RadarHTTPLocationHTTPMethodGetParamsFormatJson RadarHTTPLocationHTTPMethodGetParamsFormat = "JSON"
	RadarHTTPLocationHTTPMethodGetParamsFormatCsv  RadarHTTPLocationHTTPMethodGetParamsFormat = "CSV"
)

type RadarHTTPLocationHTTPMethodGetParamsHTTPProtocol string

const (
	RadarHTTPLocationHTTPMethodGetParamsHTTPProtocolHTTP  RadarHTTPLocationHTTPMethodGetParamsHTTPProtocol = "HTTP"
	RadarHTTPLocationHTTPMethodGetParamsHTTPProtocolHTTPS RadarHTTPLocationHTTPMethodGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPLocationHTTPMethodGetParamsIPVersion string

const (
	RadarHTTPLocationHTTPMethodGetParamsIPVersionIPv4 RadarHTTPLocationHTTPMethodGetParamsIPVersion = "IPv4"
	RadarHTTPLocationHTTPMethodGetParamsIPVersionIPv6 RadarHTTPLocationHTTPMethodGetParamsIPVersion = "IPv6"
)

type RadarHTTPLocationHTTPMethodGetParamsOS string

const (
	RadarHTTPLocationHTTPMethodGetParamsOSWindows  RadarHTTPLocationHTTPMethodGetParamsOS = "WINDOWS"
	RadarHTTPLocationHTTPMethodGetParamsOSMacosx   RadarHTTPLocationHTTPMethodGetParamsOS = "MACOSX"
	RadarHTTPLocationHTTPMethodGetParamsOSIos      RadarHTTPLocationHTTPMethodGetParamsOS = "IOS"
	RadarHTTPLocationHTTPMethodGetParamsOSAndroid  RadarHTTPLocationHTTPMethodGetParamsOS = "ANDROID"
	RadarHTTPLocationHTTPMethodGetParamsOSChromeos RadarHTTPLocationHTTPMethodGetParamsOS = "CHROMEOS"
	RadarHTTPLocationHTTPMethodGetParamsOSLinux    RadarHTTPLocationHTTPMethodGetParamsOS = "LINUX"
	RadarHTTPLocationHTTPMethodGetParamsOSSmartTv  RadarHTTPLocationHTTPMethodGetParamsOS = "SMART_TV"
)

type RadarHTTPLocationHTTPMethodGetParamsTLSVersion string

const (
	RadarHTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_0  RadarHTTPLocationHTTPMethodGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_1  RadarHTTPLocationHTTPMethodGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_2  RadarHTTPLocationHTTPMethodGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_3  RadarHTTPLocationHTTPMethodGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPLocationHTTPMethodGetParamsTLSVersionTlSvQuic RadarHTTPLocationHTTPMethodGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPLocationHTTPMethodGetResponseEnvelope struct {
	Result  RadarHTTPLocationHTTPMethodGetResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarHTTPLocationHTTPMethodGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPLocationHTTPMethodGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarHTTPLocationHTTPMethodGetResponseEnvelope]
type radarHTTPLocationHTTPMethodGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationHTTPMethodGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
