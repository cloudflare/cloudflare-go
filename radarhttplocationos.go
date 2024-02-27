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

// RadarHTTPLocationOSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPLocationOSService]
// method instead.
type RadarHTTPLocationOSService struct {
	Options []option.RequestOption
}

// NewRadarHTTPLocationOSService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPLocationOSService(opts ...option.RequestOption) (r *RadarHTTPLocationOSService) {
	r = &RadarHTTPLocationOSService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested operating systems.
// Values are a percentage out of the total traffic.
func (r *RadarHTTPLocationOSService) Get(ctx context.Context, os RadarHTTPLocationOSGetParamsOS, query RadarHTTPLocationOSGetParams, opts ...option.RequestOption) (res *RadarHTTPLocationOSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPLocationOSGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/locations/os/%v", os)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPLocationOSGetResponse struct {
	Meta RadarHTTPLocationOSGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPLocationOSGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPLocationOSGetResponseJSON   `json:"-"`
}

// radarHTTPLocationOSGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPLocationOSGetResponse]
type radarHTTPLocationOSGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationOSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationOSGetResponseMeta struct {
	DateRange      []RadarHTTPLocationOSGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                           `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPLocationOSGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPLocationOSGetResponseMetaJSON           `json:"-"`
}

// radarHTTPLocationOSGetResponseMetaJSON contains the JSON metadata for the struct
// [RadarHTTPLocationOSGetResponseMeta]
type radarHTTPLocationOSGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPLocationOSGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationOSGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                       `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPLocationOSGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPLocationOSGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarHTTPLocationOSGetResponseMetaDateRange]
type radarHTTPLocationOSGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationOSGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationOSGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPLocationOSGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                        `json:"level"`
	JSON        radarHTTPLocationOSGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPLocationOSGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarHTTPLocationOSGetResponseMetaConfidenceInfo]
type radarHTTPLocationOSGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationOSGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationOSGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                         `json:"dataSource,required"`
	Description     string                                                         `json:"description,required"`
	EventType       string                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                      `json:"startTime" format:"date-time"`
	JSON            radarHTTPLocationOSGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPLocationOSGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarHTTPLocationOSGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPLocationOSGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPLocationOSGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationOSGetResponseTop0 struct {
	ClientCountryAlpha2 string                                 `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                 `json:"clientCountryName,required"`
	Value               string                                 `json:"value,required"`
	JSON                radarHTTPLocationOSGetResponseTop0JSON `json:"-"`
}

// radarHTTPLocationOSGetResponseTop0JSON contains the JSON metadata for the struct
// [RadarHTTPLocationOSGetResponseTop0]
type radarHTTPLocationOSGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPLocationOSGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPLocationOSGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPLocationOSGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPLocationOSGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPLocationOSGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPLocationOSGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPLocationOSGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPLocationOSGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPLocationOSGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPLocationOSGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPLocationOSGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPLocationOSGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP version.
type RadarHTTPLocationOSGetParamsOS string

const (
	RadarHTTPLocationOSGetParamsOSWindows  RadarHTTPLocationOSGetParamsOS = "WINDOWS"
	RadarHTTPLocationOSGetParamsOSMacosx   RadarHTTPLocationOSGetParamsOS = "MACOSX"
	RadarHTTPLocationOSGetParamsOSIos      RadarHTTPLocationOSGetParamsOS = "IOS"
	RadarHTTPLocationOSGetParamsOSAndroid  RadarHTTPLocationOSGetParamsOS = "ANDROID"
	RadarHTTPLocationOSGetParamsOSChromeos RadarHTTPLocationOSGetParamsOS = "CHROMEOS"
	RadarHTTPLocationOSGetParamsOSLinux    RadarHTTPLocationOSGetParamsOS = "LINUX"
	RadarHTTPLocationOSGetParamsOSSmartTv  RadarHTTPLocationOSGetParamsOS = "SMART_TV"
)

type RadarHTTPLocationOSGetParamsBotClass string

const (
	RadarHTTPLocationOSGetParamsBotClassLikelyAutomated RadarHTTPLocationOSGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPLocationOSGetParamsBotClassLikelyHuman     RadarHTTPLocationOSGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPLocationOSGetParamsDateRange string

const (
	RadarHTTPLocationOSGetParamsDateRange1d         RadarHTTPLocationOSGetParamsDateRange = "1d"
	RadarHTTPLocationOSGetParamsDateRange2d         RadarHTTPLocationOSGetParamsDateRange = "2d"
	RadarHTTPLocationOSGetParamsDateRange7d         RadarHTTPLocationOSGetParamsDateRange = "7d"
	RadarHTTPLocationOSGetParamsDateRange14d        RadarHTTPLocationOSGetParamsDateRange = "14d"
	RadarHTTPLocationOSGetParamsDateRange28d        RadarHTTPLocationOSGetParamsDateRange = "28d"
	RadarHTTPLocationOSGetParamsDateRange12w        RadarHTTPLocationOSGetParamsDateRange = "12w"
	RadarHTTPLocationOSGetParamsDateRange24w        RadarHTTPLocationOSGetParamsDateRange = "24w"
	RadarHTTPLocationOSGetParamsDateRange52w        RadarHTTPLocationOSGetParamsDateRange = "52w"
	RadarHTTPLocationOSGetParamsDateRange1dControl  RadarHTTPLocationOSGetParamsDateRange = "1dControl"
	RadarHTTPLocationOSGetParamsDateRange2dControl  RadarHTTPLocationOSGetParamsDateRange = "2dControl"
	RadarHTTPLocationOSGetParamsDateRange7dControl  RadarHTTPLocationOSGetParamsDateRange = "7dControl"
	RadarHTTPLocationOSGetParamsDateRange14dControl RadarHTTPLocationOSGetParamsDateRange = "14dControl"
	RadarHTTPLocationOSGetParamsDateRange28dControl RadarHTTPLocationOSGetParamsDateRange = "28dControl"
	RadarHTTPLocationOSGetParamsDateRange12wControl RadarHTTPLocationOSGetParamsDateRange = "12wControl"
	RadarHTTPLocationOSGetParamsDateRange24wControl RadarHTTPLocationOSGetParamsDateRange = "24wControl"
)

type RadarHTTPLocationOSGetParamsDeviceType string

const (
	RadarHTTPLocationOSGetParamsDeviceTypeDesktop RadarHTTPLocationOSGetParamsDeviceType = "DESKTOP"
	RadarHTTPLocationOSGetParamsDeviceTypeMobile  RadarHTTPLocationOSGetParamsDeviceType = "MOBILE"
	RadarHTTPLocationOSGetParamsDeviceTypeOther   RadarHTTPLocationOSGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPLocationOSGetParamsFormat string

const (
	RadarHTTPLocationOSGetParamsFormatJson RadarHTTPLocationOSGetParamsFormat = "JSON"
	RadarHTTPLocationOSGetParamsFormatCsv  RadarHTTPLocationOSGetParamsFormat = "CSV"
)

type RadarHTTPLocationOSGetParamsHTTPProtocol string

const (
	RadarHTTPLocationOSGetParamsHTTPProtocolHTTP  RadarHTTPLocationOSGetParamsHTTPProtocol = "HTTP"
	RadarHTTPLocationOSGetParamsHTTPProtocolHTTPS RadarHTTPLocationOSGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPLocationOSGetParamsHTTPVersion string

const (
	RadarHTTPLocationOSGetParamsHTTPVersionHttPv1 RadarHTTPLocationOSGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPLocationOSGetParamsHTTPVersionHttPv2 RadarHTTPLocationOSGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPLocationOSGetParamsHTTPVersionHttPv3 RadarHTTPLocationOSGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPLocationOSGetParamsIPVersion string

const (
	RadarHTTPLocationOSGetParamsIPVersionIPv4 RadarHTTPLocationOSGetParamsIPVersion = "IPv4"
	RadarHTTPLocationOSGetParamsIPVersionIPv6 RadarHTTPLocationOSGetParamsIPVersion = "IPv6"
)

type RadarHTTPLocationOSGetParamsTLSVersion string

const (
	RadarHTTPLocationOSGetParamsTLSVersionTlSv1_0  RadarHTTPLocationOSGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPLocationOSGetParamsTLSVersionTlSv1_1  RadarHTTPLocationOSGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPLocationOSGetParamsTLSVersionTlSv1_2  RadarHTTPLocationOSGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPLocationOSGetParamsTLSVersionTlSv1_3  RadarHTTPLocationOSGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPLocationOSGetParamsTLSVersionTlSvQuic RadarHTTPLocationOSGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPLocationOSGetResponseEnvelope struct {
	Result  RadarHTTPLocationOSGetResponse             `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarHTTPLocationOSGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPLocationOSGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarHTTPLocationOSGetResponseEnvelope]
type radarHTTPLocationOSGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationOSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
