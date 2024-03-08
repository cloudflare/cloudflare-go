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

// RadarHTTPAseTLSVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPAseTLSVersionService]
// method instead.
type RadarHTTPAseTLSVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPAseTLSVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPAseTLSVersionService(opts ...option.RequestOption) (r *RadarHTTPAseTLSVersionService) {
	r = &RadarHTTPAseTLSVersionService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested TLS
// protocol version. Values are a percentage out of the total traffic.
func (r *RadarHTTPAseTLSVersionService) Get(ctx context.Context, tlsVersion RadarHTTPAseTLSVersionGetParamsTLSVersion, query RadarHTTPAseTLSVersionGetParams, opts ...option.RequestOption) (res *RadarHTTPAseTLSVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPAseTLSVersionGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/ases/tls_version/%v", tlsVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPAseTLSVersionGetResponse struct {
	Meta RadarHTTPAseTLSVersionGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPAseTLSVersionGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPAseTLSVersionGetResponseJSON   `json:"-"`
}

// radarHTTPAseTLSVersionGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPAseTLSVersionGetResponse]
type radarHTTPAseTLSVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseTLSVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseTLSVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseTLSVersionGetResponseMeta struct {
	DateRange      []RadarHTTPAseTLSVersionGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPAseTLSVersionGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPAseTLSVersionGetResponseMetaJSON           `json:"-"`
}

// radarHTTPAseTLSVersionGetResponseMetaJSON contains the JSON metadata for the
// struct [RadarHTTPAseTLSVersionGetResponseMeta]
type radarHTTPAseTLSVersionGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPAseTLSVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseTLSVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseTLSVersionGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPAseTLSVersionGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPAseTLSVersionGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarHTTPAseTLSVersionGetResponseMetaDateRange]
type radarHTTPAseTLSVersionGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseTLSVersionGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseTLSVersionGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseTLSVersionGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        radarHTTPAseTLSVersionGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPAseTLSVersionGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarHTTPAseTLSVersionGetResponseMetaConfidenceInfo]
type radarHTTPAseTLSVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseTLSVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseTLSVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            radarHTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarHTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseTLSVersionGetResponseTop0 struct {
	ClientASN    int64                                     `json:"clientASN,required"`
	ClientAsName string                                    `json:"clientASName,required"`
	Value        string                                    `json:"value,required"`
	JSON         radarHTTPAseTLSVersionGetResponseTop0JSON `json:"-"`
}

// radarHTTPAseTLSVersionGetResponseTop0JSON contains the JSON metadata for the
// struct [RadarHTTPAseTLSVersionGetResponseTop0]
type radarHTTPAseTLSVersionGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPAseTLSVersionGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseTLSVersionGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseTLSVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPAseTLSVersionGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPAseTLSVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPAseTLSVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPAseTLSVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPAseTLSVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPAseTLSVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPAseTLSVersionGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPAseTLSVersionGetParamsOS] `query:"os"`
}

// URLQuery serializes [RadarHTTPAseTLSVersionGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPAseTLSVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// TLS version.
type RadarHTTPAseTLSVersionGetParamsTLSVersion string

const (
	RadarHTTPAseTLSVersionGetParamsTLSVersionTlSv1_0  RadarHTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPAseTLSVersionGetParamsTLSVersionTlSv1_1  RadarHTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPAseTLSVersionGetParamsTLSVersionTlSv1_2  RadarHTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPAseTLSVersionGetParamsTLSVersionTlSv1_3  RadarHTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPAseTLSVersionGetParamsTLSVersionTlSvQuic RadarHTTPAseTLSVersionGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPAseTLSVersionGetParamsBotClass string

const (
	RadarHTTPAseTLSVersionGetParamsBotClassLikelyAutomated RadarHTTPAseTLSVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPAseTLSVersionGetParamsBotClassLikelyHuman     RadarHTTPAseTLSVersionGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPAseTLSVersionGetParamsDateRange string

const (
	RadarHTTPAseTLSVersionGetParamsDateRange1d         RadarHTTPAseTLSVersionGetParamsDateRange = "1d"
	RadarHTTPAseTLSVersionGetParamsDateRange2d         RadarHTTPAseTLSVersionGetParamsDateRange = "2d"
	RadarHTTPAseTLSVersionGetParamsDateRange7d         RadarHTTPAseTLSVersionGetParamsDateRange = "7d"
	RadarHTTPAseTLSVersionGetParamsDateRange14d        RadarHTTPAseTLSVersionGetParamsDateRange = "14d"
	RadarHTTPAseTLSVersionGetParamsDateRange28d        RadarHTTPAseTLSVersionGetParamsDateRange = "28d"
	RadarHTTPAseTLSVersionGetParamsDateRange12w        RadarHTTPAseTLSVersionGetParamsDateRange = "12w"
	RadarHTTPAseTLSVersionGetParamsDateRange24w        RadarHTTPAseTLSVersionGetParamsDateRange = "24w"
	RadarHTTPAseTLSVersionGetParamsDateRange52w        RadarHTTPAseTLSVersionGetParamsDateRange = "52w"
	RadarHTTPAseTLSVersionGetParamsDateRange1dControl  RadarHTTPAseTLSVersionGetParamsDateRange = "1dControl"
	RadarHTTPAseTLSVersionGetParamsDateRange2dControl  RadarHTTPAseTLSVersionGetParamsDateRange = "2dControl"
	RadarHTTPAseTLSVersionGetParamsDateRange7dControl  RadarHTTPAseTLSVersionGetParamsDateRange = "7dControl"
	RadarHTTPAseTLSVersionGetParamsDateRange14dControl RadarHTTPAseTLSVersionGetParamsDateRange = "14dControl"
	RadarHTTPAseTLSVersionGetParamsDateRange28dControl RadarHTTPAseTLSVersionGetParamsDateRange = "28dControl"
	RadarHTTPAseTLSVersionGetParamsDateRange12wControl RadarHTTPAseTLSVersionGetParamsDateRange = "12wControl"
	RadarHTTPAseTLSVersionGetParamsDateRange24wControl RadarHTTPAseTLSVersionGetParamsDateRange = "24wControl"
)

type RadarHTTPAseTLSVersionGetParamsDeviceType string

const (
	RadarHTTPAseTLSVersionGetParamsDeviceTypeDesktop RadarHTTPAseTLSVersionGetParamsDeviceType = "DESKTOP"
	RadarHTTPAseTLSVersionGetParamsDeviceTypeMobile  RadarHTTPAseTLSVersionGetParamsDeviceType = "MOBILE"
	RadarHTTPAseTLSVersionGetParamsDeviceTypeOther   RadarHTTPAseTLSVersionGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPAseTLSVersionGetParamsFormat string

const (
	RadarHTTPAseTLSVersionGetParamsFormatJson RadarHTTPAseTLSVersionGetParamsFormat = "JSON"
	RadarHTTPAseTLSVersionGetParamsFormatCsv  RadarHTTPAseTLSVersionGetParamsFormat = "CSV"
)

type RadarHTTPAseTLSVersionGetParamsHTTPProtocol string

const (
	RadarHTTPAseTLSVersionGetParamsHTTPProtocolHTTP  RadarHTTPAseTLSVersionGetParamsHTTPProtocol = "HTTP"
	RadarHTTPAseTLSVersionGetParamsHTTPProtocolHTTPS RadarHTTPAseTLSVersionGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPAseTLSVersionGetParamsHTTPVersion string

const (
	RadarHTTPAseTLSVersionGetParamsHTTPVersionHttPv1 RadarHTTPAseTLSVersionGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPAseTLSVersionGetParamsHTTPVersionHttPv2 RadarHTTPAseTLSVersionGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPAseTLSVersionGetParamsHTTPVersionHttPv3 RadarHTTPAseTLSVersionGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPAseTLSVersionGetParamsIPVersion string

const (
	RadarHTTPAseTLSVersionGetParamsIPVersionIPv4 RadarHTTPAseTLSVersionGetParamsIPVersion = "IPv4"
	RadarHTTPAseTLSVersionGetParamsIPVersionIPv6 RadarHTTPAseTLSVersionGetParamsIPVersion = "IPv6"
)

type RadarHTTPAseTLSVersionGetParamsOS string

const (
	RadarHTTPAseTLSVersionGetParamsOSWindows  RadarHTTPAseTLSVersionGetParamsOS = "WINDOWS"
	RadarHTTPAseTLSVersionGetParamsOSMacosx   RadarHTTPAseTLSVersionGetParamsOS = "MACOSX"
	RadarHTTPAseTLSVersionGetParamsOSIos      RadarHTTPAseTLSVersionGetParamsOS = "IOS"
	RadarHTTPAseTLSVersionGetParamsOSAndroid  RadarHTTPAseTLSVersionGetParamsOS = "ANDROID"
	RadarHTTPAseTLSVersionGetParamsOSChromeos RadarHTTPAseTLSVersionGetParamsOS = "CHROMEOS"
	RadarHTTPAseTLSVersionGetParamsOSLinux    RadarHTTPAseTLSVersionGetParamsOS = "LINUX"
	RadarHTTPAseTLSVersionGetParamsOSSmartTv  RadarHTTPAseTLSVersionGetParamsOS = "SMART_TV"
)

type RadarHTTPAseTLSVersionGetResponseEnvelope struct {
	Result  RadarHTTPAseTLSVersionGetResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarHTTPAseTLSVersionGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPAseTLSVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarHTTPAseTLSVersionGetResponseEnvelope]
type radarHTTPAseTLSVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseTLSVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseTLSVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
