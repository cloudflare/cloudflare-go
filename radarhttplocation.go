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

// RadarHTTPLocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPLocationService] method
// instead.
type RadarHTTPLocationService struct {
	Options      []option.RequestOption
	BotClass     *RadarHTTPLocationBotClassService
	DeviceType   *RadarHTTPLocationDeviceTypeService
	HTTPProtocol *RadarHTTPLocationHTTPProtocolService
	HTTPMethod   *RadarHTTPLocationHTTPMethodService
	IPVersion    *RadarHTTPLocationIPVersionService
	OS           *RadarHTTPLocationOSService
	TLSVersion   *RadarHTTPLocationTLSVersionService
}

// NewRadarHTTPLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPLocationService(opts ...option.RequestOption) (r *RadarHTTPLocationService) {
	r = &RadarHTTPLocationService{}
	r.Options = opts
	r.BotClass = NewRadarHTTPLocationBotClassService(opts...)
	r.DeviceType = NewRadarHTTPLocationDeviceTypeService(opts...)
	r.HTTPProtocol = NewRadarHTTPLocationHTTPProtocolService(opts...)
	r.HTTPMethod = NewRadarHTTPLocationHTTPMethodService(opts...)
	r.IPVersion = NewRadarHTTPLocationIPVersionService(opts...)
	r.OS = NewRadarHTTPLocationOSService(opts...)
	r.TLSVersion = NewRadarHTTPLocationTLSVersionService(opts...)
	return
}

// Get the top locations by HTTP traffic. Values are a percentage out of the total
// traffic.
func (r *RadarHTTPLocationService) Get(ctx context.Context, query RadarHTTPLocationGetParams, opts ...option.RequestOption) (res *RadarHTTPLocationGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPLocationGetResponseEnvelope
	path := "radar/http/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPLocationGetResponse struct {
	Meta RadarHTTPLocationGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPLocationGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPLocationGetResponseJSON   `json:"-"`
}

// radarHTTPLocationGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPLocationGetResponse]
type radarHTTPLocationGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationGetResponseMeta struct {
	DateRange      []RadarHTTPLocationGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                         `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPLocationGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPLocationGetResponseMetaJSON           `json:"-"`
}

// radarHTTPLocationGetResponseMetaJSON contains the JSON metadata for the struct
// [RadarHTTPLocationGetResponseMeta]
type radarHTTPLocationGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPLocationGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPLocationGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPLocationGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarHTTPLocationGetResponseMetaDateRange]
type radarHTTPLocationGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPLocationGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                      `json:"level"`
	JSON        radarHTTPLocationGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPLocationGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarHTTPLocationGetResponseMetaConfidenceInfo]
type radarHTTPLocationGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                       `json:"dataSource,required"`
	Description     string                                                       `json:"description,required"`
	EventType       string                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                    `json:"startTime" format:"date-time"`
	JSON            radarHTTPLocationGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPLocationGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarHTTPLocationGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPLocationGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPLocationGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationGetResponseTop0 struct {
	ClientCountryAlpha2 string                               `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                               `json:"clientCountryName,required"`
	Value               string                               `json:"value,required"`
	JSON                radarHTTPLocationGetResponseTop0JSON `json:"-"`
}

// radarHTTPLocationGetResponseTop0JSON contains the JSON metadata for the struct
// [RadarHTTPLocationGetResponseTop0]
type radarHTTPLocationGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPLocationGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPLocationGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPLocationGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPLocationGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPLocationGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPLocationGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPLocationGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPLocationGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPLocationGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPLocationGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPLocationGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPLocationGetParamsBotClass string

const (
	RadarHTTPLocationGetParamsBotClassLikelyAutomated RadarHTTPLocationGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPLocationGetParamsBotClassLikelyHuman     RadarHTTPLocationGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPLocationGetParamsDateRange string

const (
	RadarHTTPLocationGetParamsDateRange1d         RadarHTTPLocationGetParamsDateRange = "1d"
	RadarHTTPLocationGetParamsDateRange2d         RadarHTTPLocationGetParamsDateRange = "2d"
	RadarHTTPLocationGetParamsDateRange7d         RadarHTTPLocationGetParamsDateRange = "7d"
	RadarHTTPLocationGetParamsDateRange14d        RadarHTTPLocationGetParamsDateRange = "14d"
	RadarHTTPLocationGetParamsDateRange28d        RadarHTTPLocationGetParamsDateRange = "28d"
	RadarHTTPLocationGetParamsDateRange12w        RadarHTTPLocationGetParamsDateRange = "12w"
	RadarHTTPLocationGetParamsDateRange24w        RadarHTTPLocationGetParamsDateRange = "24w"
	RadarHTTPLocationGetParamsDateRange52w        RadarHTTPLocationGetParamsDateRange = "52w"
	RadarHTTPLocationGetParamsDateRange1dControl  RadarHTTPLocationGetParamsDateRange = "1dControl"
	RadarHTTPLocationGetParamsDateRange2dControl  RadarHTTPLocationGetParamsDateRange = "2dControl"
	RadarHTTPLocationGetParamsDateRange7dControl  RadarHTTPLocationGetParamsDateRange = "7dControl"
	RadarHTTPLocationGetParamsDateRange14dControl RadarHTTPLocationGetParamsDateRange = "14dControl"
	RadarHTTPLocationGetParamsDateRange28dControl RadarHTTPLocationGetParamsDateRange = "28dControl"
	RadarHTTPLocationGetParamsDateRange12wControl RadarHTTPLocationGetParamsDateRange = "12wControl"
	RadarHTTPLocationGetParamsDateRange24wControl RadarHTTPLocationGetParamsDateRange = "24wControl"
)

type RadarHTTPLocationGetParamsDeviceType string

const (
	RadarHTTPLocationGetParamsDeviceTypeDesktop RadarHTTPLocationGetParamsDeviceType = "DESKTOP"
	RadarHTTPLocationGetParamsDeviceTypeMobile  RadarHTTPLocationGetParamsDeviceType = "MOBILE"
	RadarHTTPLocationGetParamsDeviceTypeOther   RadarHTTPLocationGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPLocationGetParamsFormat string

const (
	RadarHTTPLocationGetParamsFormatJson RadarHTTPLocationGetParamsFormat = "JSON"
	RadarHTTPLocationGetParamsFormatCsv  RadarHTTPLocationGetParamsFormat = "CSV"
)

type RadarHTTPLocationGetParamsHTTPProtocol string

const (
	RadarHTTPLocationGetParamsHTTPProtocolHTTP  RadarHTTPLocationGetParamsHTTPProtocol = "HTTP"
	RadarHTTPLocationGetParamsHTTPProtocolHTTPS RadarHTTPLocationGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPLocationGetParamsHTTPVersion string

const (
	RadarHTTPLocationGetParamsHTTPVersionHttPv1 RadarHTTPLocationGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPLocationGetParamsHTTPVersionHttPv2 RadarHTTPLocationGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPLocationGetParamsHTTPVersionHttPv3 RadarHTTPLocationGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPLocationGetParamsIPVersion string

const (
	RadarHTTPLocationGetParamsIPVersionIPv4 RadarHTTPLocationGetParamsIPVersion = "IPv4"
	RadarHTTPLocationGetParamsIPVersionIPv6 RadarHTTPLocationGetParamsIPVersion = "IPv6"
)

type RadarHTTPLocationGetParamsOS string

const (
	RadarHTTPLocationGetParamsOSWindows  RadarHTTPLocationGetParamsOS = "WINDOWS"
	RadarHTTPLocationGetParamsOSMacosx   RadarHTTPLocationGetParamsOS = "MACOSX"
	RadarHTTPLocationGetParamsOSIos      RadarHTTPLocationGetParamsOS = "IOS"
	RadarHTTPLocationGetParamsOSAndroid  RadarHTTPLocationGetParamsOS = "ANDROID"
	RadarHTTPLocationGetParamsOSChromeos RadarHTTPLocationGetParamsOS = "CHROMEOS"
	RadarHTTPLocationGetParamsOSLinux    RadarHTTPLocationGetParamsOS = "LINUX"
	RadarHTTPLocationGetParamsOSSmartTv  RadarHTTPLocationGetParamsOS = "SMART_TV"
)

type RadarHTTPLocationGetParamsTLSVersion string

const (
	RadarHTTPLocationGetParamsTLSVersionTlSv1_0  RadarHTTPLocationGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPLocationGetParamsTLSVersionTlSv1_1  RadarHTTPLocationGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPLocationGetParamsTLSVersionTlSv1_2  RadarHTTPLocationGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPLocationGetParamsTLSVersionTlSv1_3  RadarHTTPLocationGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPLocationGetParamsTLSVersionTlSvQuic RadarHTTPLocationGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPLocationGetResponseEnvelope struct {
	Result  RadarHTTPLocationGetResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarHTTPLocationGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPLocationGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarHTTPLocationGetResponseEnvelope]
type radarHTTPLocationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
