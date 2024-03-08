// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// RadarHTTPLocationIPVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPLocationIPVersionService] method instead.
type RadarHTTPLocationIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPLocationIPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPLocationIPVersionService(opts ...option.RequestOption) (r *RadarHTTPLocationIPVersionService) {
	r = &RadarHTTPLocationIPVersionService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested IP protocol version.
// Values are a percentage out of the total traffic.
func (r *RadarHTTPLocationIPVersionService) Get(ctx context.Context, ipVersion RadarHTTPLocationIPVersionGetParamsIPVersion, query RadarHTTPLocationIPVersionGetParams, opts ...option.RequestOption) (res *RadarHTTPLocationIPVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPLocationIPVersionGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/locations/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPLocationIPVersionGetResponse struct {
	Meta RadarHTTPLocationIPVersionGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPLocationIPVersionGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPLocationIPVersionGetResponseJSON   `json:"-"`
}

// radarHTTPLocationIPVersionGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPLocationIPVersionGetResponse]
type radarHTTPLocationIPVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationIPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationIPVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationIPVersionGetResponseMeta struct {
	DateRange      []RadarHTTPLocationIPVersionGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPLocationIPVersionGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPLocationIPVersionGetResponseMetaJSON           `json:"-"`
}

// radarHTTPLocationIPVersionGetResponseMetaJSON contains the JSON metadata for the
// struct [RadarHTTPLocationIPVersionGetResponseMeta]
type radarHTTPLocationIPVersionGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPLocationIPVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationIPVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationIPVersionGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPLocationIPVersionGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPLocationIPVersionGetResponseMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPLocationIPVersionGetResponseMetaDateRange]
type radarHTTPLocationIPVersionGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationIPVersionGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationIPVersionGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationIPVersionGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        radarHTTPLocationIPVersionGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPLocationIPVersionGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarHTTPLocationIPVersionGetResponseMetaConfidenceInfo]
type radarHTTPLocationIPVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationIPVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationIPVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            radarHTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarHTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationIPVersionGetResponseTop0 struct {
	ClientCountryAlpha2 string                                        `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                        `json:"clientCountryName,required"`
	Value               string                                        `json:"value,required"`
	JSON                radarHTTPLocationIPVersionGetResponseTop0JSON `json:"-"`
}

// radarHTTPLocationIPVersionGetResponseTop0JSON contains the JSON metadata for the
// struct [RadarHTTPLocationIPVersionGetResponseTop0]
type radarHTTPLocationIPVersionGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPLocationIPVersionGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationIPVersionGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPLocationIPVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPLocationIPVersionGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPLocationIPVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPLocationIPVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPLocationIPVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPLocationIPVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPLocationIPVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPLocationIPVersionGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPLocationIPVersionGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPLocationIPVersionGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPLocationIPVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP version.
type RadarHTTPLocationIPVersionGetParamsIPVersion string

const (
	RadarHTTPLocationIPVersionGetParamsIPVersionIPv4 RadarHTTPLocationIPVersionGetParamsIPVersion = "IPv4"
	RadarHTTPLocationIPVersionGetParamsIPVersionIPv6 RadarHTTPLocationIPVersionGetParamsIPVersion = "IPv6"
)

type RadarHTTPLocationIPVersionGetParamsBotClass string

const (
	RadarHTTPLocationIPVersionGetParamsBotClassLikelyAutomated RadarHTTPLocationIPVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPLocationIPVersionGetParamsBotClassLikelyHuman     RadarHTTPLocationIPVersionGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPLocationIPVersionGetParamsDateRange string

const (
	RadarHTTPLocationIPVersionGetParamsDateRange1d         RadarHTTPLocationIPVersionGetParamsDateRange = "1d"
	RadarHTTPLocationIPVersionGetParamsDateRange2d         RadarHTTPLocationIPVersionGetParamsDateRange = "2d"
	RadarHTTPLocationIPVersionGetParamsDateRange7d         RadarHTTPLocationIPVersionGetParamsDateRange = "7d"
	RadarHTTPLocationIPVersionGetParamsDateRange14d        RadarHTTPLocationIPVersionGetParamsDateRange = "14d"
	RadarHTTPLocationIPVersionGetParamsDateRange28d        RadarHTTPLocationIPVersionGetParamsDateRange = "28d"
	RadarHTTPLocationIPVersionGetParamsDateRange12w        RadarHTTPLocationIPVersionGetParamsDateRange = "12w"
	RadarHTTPLocationIPVersionGetParamsDateRange24w        RadarHTTPLocationIPVersionGetParamsDateRange = "24w"
	RadarHTTPLocationIPVersionGetParamsDateRange52w        RadarHTTPLocationIPVersionGetParamsDateRange = "52w"
	RadarHTTPLocationIPVersionGetParamsDateRange1dControl  RadarHTTPLocationIPVersionGetParamsDateRange = "1dControl"
	RadarHTTPLocationIPVersionGetParamsDateRange2dControl  RadarHTTPLocationIPVersionGetParamsDateRange = "2dControl"
	RadarHTTPLocationIPVersionGetParamsDateRange7dControl  RadarHTTPLocationIPVersionGetParamsDateRange = "7dControl"
	RadarHTTPLocationIPVersionGetParamsDateRange14dControl RadarHTTPLocationIPVersionGetParamsDateRange = "14dControl"
	RadarHTTPLocationIPVersionGetParamsDateRange28dControl RadarHTTPLocationIPVersionGetParamsDateRange = "28dControl"
	RadarHTTPLocationIPVersionGetParamsDateRange12wControl RadarHTTPLocationIPVersionGetParamsDateRange = "12wControl"
	RadarHTTPLocationIPVersionGetParamsDateRange24wControl RadarHTTPLocationIPVersionGetParamsDateRange = "24wControl"
)

type RadarHTTPLocationIPVersionGetParamsDeviceType string

const (
	RadarHTTPLocationIPVersionGetParamsDeviceTypeDesktop RadarHTTPLocationIPVersionGetParamsDeviceType = "DESKTOP"
	RadarHTTPLocationIPVersionGetParamsDeviceTypeMobile  RadarHTTPLocationIPVersionGetParamsDeviceType = "MOBILE"
	RadarHTTPLocationIPVersionGetParamsDeviceTypeOther   RadarHTTPLocationIPVersionGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPLocationIPVersionGetParamsFormat string

const (
	RadarHTTPLocationIPVersionGetParamsFormatJson RadarHTTPLocationIPVersionGetParamsFormat = "JSON"
	RadarHTTPLocationIPVersionGetParamsFormatCsv  RadarHTTPLocationIPVersionGetParamsFormat = "CSV"
)

type RadarHTTPLocationIPVersionGetParamsHTTPProtocol string

const (
	RadarHTTPLocationIPVersionGetParamsHTTPProtocolHTTP  RadarHTTPLocationIPVersionGetParamsHTTPProtocol = "HTTP"
	RadarHTTPLocationIPVersionGetParamsHTTPProtocolHTTPS RadarHTTPLocationIPVersionGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPLocationIPVersionGetParamsHTTPVersion string

const (
	RadarHTTPLocationIPVersionGetParamsHTTPVersionHttPv1 RadarHTTPLocationIPVersionGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPLocationIPVersionGetParamsHTTPVersionHttPv2 RadarHTTPLocationIPVersionGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPLocationIPVersionGetParamsHTTPVersionHttPv3 RadarHTTPLocationIPVersionGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPLocationIPVersionGetParamsOS string

const (
	RadarHTTPLocationIPVersionGetParamsOSWindows  RadarHTTPLocationIPVersionGetParamsOS = "WINDOWS"
	RadarHTTPLocationIPVersionGetParamsOSMacosx   RadarHTTPLocationIPVersionGetParamsOS = "MACOSX"
	RadarHTTPLocationIPVersionGetParamsOSIos      RadarHTTPLocationIPVersionGetParamsOS = "IOS"
	RadarHTTPLocationIPVersionGetParamsOSAndroid  RadarHTTPLocationIPVersionGetParamsOS = "ANDROID"
	RadarHTTPLocationIPVersionGetParamsOSChromeos RadarHTTPLocationIPVersionGetParamsOS = "CHROMEOS"
	RadarHTTPLocationIPVersionGetParamsOSLinux    RadarHTTPLocationIPVersionGetParamsOS = "LINUX"
	RadarHTTPLocationIPVersionGetParamsOSSmartTv  RadarHTTPLocationIPVersionGetParamsOS = "SMART_TV"
)

type RadarHTTPLocationIPVersionGetParamsTLSVersion string

const (
	RadarHTTPLocationIPVersionGetParamsTLSVersionTlSv1_0  RadarHTTPLocationIPVersionGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPLocationIPVersionGetParamsTLSVersionTlSv1_1  RadarHTTPLocationIPVersionGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPLocationIPVersionGetParamsTLSVersionTlSv1_2  RadarHTTPLocationIPVersionGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPLocationIPVersionGetParamsTLSVersionTlSv1_3  RadarHTTPLocationIPVersionGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPLocationIPVersionGetParamsTLSVersionTlSvQuic RadarHTTPLocationIPVersionGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPLocationIPVersionGetResponseEnvelope struct {
	Result  RadarHTTPLocationIPVersionGetResponse             `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarHTTPLocationIPVersionGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPLocationIPVersionGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarHTTPLocationIPVersionGetResponseEnvelope]
type radarHTTPLocationIPVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPLocationIPVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPLocationIPVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
