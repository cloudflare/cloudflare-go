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

// RadarHTTPLocationTLSVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPLocationTLSVersionService] method instead.
type RadarHTTPLocationTLSVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPLocationTLSVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPLocationTLSVersionService(opts ...option.RequestOption) (r *RadarHTTPLocationTLSVersionService) {
	r = &RadarHTTPLocationTLSVersionService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested TLS protocol version.
// Values are a percentage out of the total traffic.
func (r *RadarHTTPLocationTLSVersionService) Get(ctx context.Context, tlsVersion RadarHTTPLocationTLSVersionGetParamsTLSVersion, query RadarHTTPLocationTLSVersionGetParams, opts ...option.RequestOption) (res *RadarHTTPLocationTLSVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPLocationTLSVersionGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/locations/tls_version/%v", tlsVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPLocationTLSVersionGetResponse struct {
	Meta RadarHTTPLocationTLSVersionGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPLocationTLSVersionGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPLocationTLSVersionGetResponseJSON   `json:"-"`
}

// radarHTTPLocationTLSVersionGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPLocationTLSVersionGetResponse]
type radarHTTPLocationTLSVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationTLSVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationTLSVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationTLSVersionGetResponseMeta struct {
	DateRange      []RadarHTTPLocationTLSVersionGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPLocationTLSVersionGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPLocationTLSVersionGetResponseMetaJSON           `json:"-"`
}

// radarHTTPLocationTLSVersionGetResponseMetaJSON contains the JSON metadata for
// the struct [RadarHTTPLocationTLSVersionGetResponseMeta]
type radarHTTPLocationTLSVersionGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPLocationTLSVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationTLSVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationTLSVersionGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPLocationTLSVersionGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPLocationTLSVersionGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPLocationTLSVersionGetResponseMetaDateRange]
type radarHTTPLocationTLSVersionGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationTLSVersionGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationTLSVersionGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationTLSVersionGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarHTTPLocationTLSVersionGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPLocationTLSVersionGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarHTTPLocationTLSVersionGetResponseMetaConfidenceInfo]
type radarHTTPLocationTLSVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationTLSVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationTLSVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarHTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarHTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationTLSVersionGetResponseTop0 struct {
	ClientCountryAlpha2 string                                         `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                         `json:"clientCountryName,required"`
	Value               string                                         `json:"value,required"`
	JSON                radarHTTPLocationTLSVersionGetResponseTop0JSON `json:"-"`
}

// radarHTTPLocationTLSVersionGetResponseTop0JSON contains the JSON metadata for
// the struct [RadarHTTPLocationTLSVersionGetResponseTop0]
type radarHTTPLocationTLSVersionGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPLocationTLSVersionGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationTLSVersionGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationTLSVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPLocationTLSVersionGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPLocationTLSVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPLocationTLSVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPLocationTLSVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPLocationTLSVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPLocationTLSVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPLocationTLSVersionGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPLocationTLSVersionGetParamsOS] `query:"os"`
}

// URLQuery serializes [RadarHTTPLocationTLSVersionGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPLocationTLSVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// TLS version.
type RadarHTTPLocationTLSVersionGetParamsTLSVersion string

const (
	RadarHTTPLocationTLSVersionGetParamsTLSVersionTlSv1_0  RadarHTTPLocationTLSVersionGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPLocationTLSVersionGetParamsTLSVersionTlSv1_1  RadarHTTPLocationTLSVersionGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPLocationTLSVersionGetParamsTLSVersionTlSv1_2  RadarHTTPLocationTLSVersionGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPLocationTLSVersionGetParamsTLSVersionTlSv1_3  RadarHTTPLocationTLSVersionGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPLocationTLSVersionGetParamsTLSVersionTlSvQuic RadarHTTPLocationTLSVersionGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPLocationTLSVersionGetParamsBotClass string

const (
	RadarHTTPLocationTLSVersionGetParamsBotClassLikelyAutomated RadarHTTPLocationTLSVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPLocationTLSVersionGetParamsBotClassLikelyHuman     RadarHTTPLocationTLSVersionGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPLocationTLSVersionGetParamsDateRange string

const (
	RadarHTTPLocationTLSVersionGetParamsDateRange1d         RadarHTTPLocationTLSVersionGetParamsDateRange = "1d"
	RadarHTTPLocationTLSVersionGetParamsDateRange2d         RadarHTTPLocationTLSVersionGetParamsDateRange = "2d"
	RadarHTTPLocationTLSVersionGetParamsDateRange7d         RadarHTTPLocationTLSVersionGetParamsDateRange = "7d"
	RadarHTTPLocationTLSVersionGetParamsDateRange14d        RadarHTTPLocationTLSVersionGetParamsDateRange = "14d"
	RadarHTTPLocationTLSVersionGetParamsDateRange28d        RadarHTTPLocationTLSVersionGetParamsDateRange = "28d"
	RadarHTTPLocationTLSVersionGetParamsDateRange12w        RadarHTTPLocationTLSVersionGetParamsDateRange = "12w"
	RadarHTTPLocationTLSVersionGetParamsDateRange24w        RadarHTTPLocationTLSVersionGetParamsDateRange = "24w"
	RadarHTTPLocationTLSVersionGetParamsDateRange52w        RadarHTTPLocationTLSVersionGetParamsDateRange = "52w"
	RadarHTTPLocationTLSVersionGetParamsDateRange1dControl  RadarHTTPLocationTLSVersionGetParamsDateRange = "1dControl"
	RadarHTTPLocationTLSVersionGetParamsDateRange2dControl  RadarHTTPLocationTLSVersionGetParamsDateRange = "2dControl"
	RadarHTTPLocationTLSVersionGetParamsDateRange7dControl  RadarHTTPLocationTLSVersionGetParamsDateRange = "7dControl"
	RadarHTTPLocationTLSVersionGetParamsDateRange14dControl RadarHTTPLocationTLSVersionGetParamsDateRange = "14dControl"
	RadarHTTPLocationTLSVersionGetParamsDateRange28dControl RadarHTTPLocationTLSVersionGetParamsDateRange = "28dControl"
	RadarHTTPLocationTLSVersionGetParamsDateRange12wControl RadarHTTPLocationTLSVersionGetParamsDateRange = "12wControl"
	RadarHTTPLocationTLSVersionGetParamsDateRange24wControl RadarHTTPLocationTLSVersionGetParamsDateRange = "24wControl"
)

type RadarHTTPLocationTLSVersionGetParamsDeviceType string

const (
	RadarHTTPLocationTLSVersionGetParamsDeviceTypeDesktop RadarHTTPLocationTLSVersionGetParamsDeviceType = "DESKTOP"
	RadarHTTPLocationTLSVersionGetParamsDeviceTypeMobile  RadarHTTPLocationTLSVersionGetParamsDeviceType = "MOBILE"
	RadarHTTPLocationTLSVersionGetParamsDeviceTypeOther   RadarHTTPLocationTLSVersionGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPLocationTLSVersionGetParamsFormat string

const (
	RadarHTTPLocationTLSVersionGetParamsFormatJson RadarHTTPLocationTLSVersionGetParamsFormat = "JSON"
	RadarHTTPLocationTLSVersionGetParamsFormatCsv  RadarHTTPLocationTLSVersionGetParamsFormat = "CSV"
)

type RadarHTTPLocationTLSVersionGetParamsHTTPProtocol string

const (
	RadarHTTPLocationTLSVersionGetParamsHTTPProtocolHTTP  RadarHTTPLocationTLSVersionGetParamsHTTPProtocol = "HTTP"
	RadarHTTPLocationTLSVersionGetParamsHTTPProtocolHTTPS RadarHTTPLocationTLSVersionGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPLocationTLSVersionGetParamsHTTPVersion string

const (
	RadarHTTPLocationTLSVersionGetParamsHTTPVersionHttPv1 RadarHTTPLocationTLSVersionGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPLocationTLSVersionGetParamsHTTPVersionHttPv2 RadarHTTPLocationTLSVersionGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPLocationTLSVersionGetParamsHTTPVersionHttPv3 RadarHTTPLocationTLSVersionGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPLocationTLSVersionGetParamsIPVersion string

const (
	RadarHTTPLocationTLSVersionGetParamsIPVersionIPv4 RadarHTTPLocationTLSVersionGetParamsIPVersion = "IPv4"
	RadarHTTPLocationTLSVersionGetParamsIPVersionIPv6 RadarHTTPLocationTLSVersionGetParamsIPVersion = "IPv6"
)

type RadarHTTPLocationTLSVersionGetParamsOS string

const (
	RadarHTTPLocationTLSVersionGetParamsOSWindows  RadarHTTPLocationTLSVersionGetParamsOS = "WINDOWS"
	RadarHTTPLocationTLSVersionGetParamsOSMacosx   RadarHTTPLocationTLSVersionGetParamsOS = "MACOSX"
	RadarHTTPLocationTLSVersionGetParamsOSIos      RadarHTTPLocationTLSVersionGetParamsOS = "IOS"
	RadarHTTPLocationTLSVersionGetParamsOSAndroid  RadarHTTPLocationTLSVersionGetParamsOS = "ANDROID"
	RadarHTTPLocationTLSVersionGetParamsOSChromeos RadarHTTPLocationTLSVersionGetParamsOS = "CHROMEOS"
	RadarHTTPLocationTLSVersionGetParamsOSLinux    RadarHTTPLocationTLSVersionGetParamsOS = "LINUX"
	RadarHTTPLocationTLSVersionGetParamsOSSmartTv  RadarHTTPLocationTLSVersionGetParamsOS = "SMART_TV"
)

type RadarHTTPLocationTLSVersionGetResponseEnvelope struct {
	Result  RadarHTTPLocationTLSVersionGetResponse             `json:"result,required"`
	Success bool                                               `json:"success,required"`
	JSON    radarHTTPLocationTLSVersionGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPLocationTLSVersionGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [RadarHTTPLocationTLSVersionGetResponseEnvelope]
type radarHTTPLocationTLSVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationTLSVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationTLSVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
