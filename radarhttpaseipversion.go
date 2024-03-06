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

// RadarHTTPAseIPVersionService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPAseIPVersionService]
// method instead.
type RadarHTTPAseIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPAseIPVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPAseIPVersionService(opts ...option.RequestOption) (r *RadarHTTPAseIPVersionService) {
	r = &RadarHTTPAseIPVersionService{}
	r.Options = opts
	return
}

// Get the top autonomous systems, by HTTP traffic, of the requested IP protocol
// version. Values are a percentage out of the total traffic.
func (r *RadarHTTPAseIPVersionService) Get(ctx context.Context, ipVersion RadarHTTPAseIPVersionGetParamsIPVersion, query RadarHTTPAseIPVersionGetParams, opts ...option.RequestOption) (res *RadarHTTPAseIPVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPAseIPVersionGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/ases/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPAseIPVersionGetResponse struct {
	Meta RadarHTTPAseIPVersionGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPAseIPVersionGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPAseIPVersionGetResponseJSON   `json:"-"`
}

// radarHTTPAseIPVersionGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPAseIPVersionGetResponse]
type radarHTTPAseIPVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseIPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseIPVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseIPVersionGetResponseMeta struct {
	DateRange      []RadarHTTPAseIPVersionGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                             `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPAseIPVersionGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPAseIPVersionGetResponseMetaJSON           `json:"-"`
}

// radarHTTPAseIPVersionGetResponseMetaJSON contains the JSON metadata for the
// struct [RadarHTTPAseIPVersionGetResponseMeta]
type radarHTTPAseIPVersionGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPAseIPVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseIPVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseIPVersionGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPAseIPVersionGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPAseIPVersionGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [RadarHTTPAseIPVersionGetResponseMetaDateRange]
type radarHTTPAseIPVersionGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseIPVersionGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseIPVersionGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseIPVersionGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                          `json:"level"`
	JSON        radarHTTPAseIPVersionGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPAseIPVersionGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarHTTPAseIPVersionGetResponseMetaConfidenceInfo]
type radarHTTPAseIPVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseIPVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseIPVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                           `json:"dataSource,required"`
	Description     string                                                           `json:"description,required"`
	EventType       string                                                           `json:"eventType,required"`
	IsInstantaneous interface{}                                                      `json:"isInstantaneous,required"`
	EndTime         time.Time                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                        `json:"startTime" format:"date-time"`
	JSON            radarHTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarHTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseIPVersionGetResponseTop0 struct {
	ClientASN    int64                                    `json:"clientASN,required"`
	ClientAsName string                                   `json:"clientASName,required"`
	Value        string                                   `json:"value,required"`
	JSON         radarHTTPAseIPVersionGetResponseTop0JSON `json:"-"`
}

// radarHTTPAseIPVersionGetResponseTop0JSON contains the JSON metadata for the
// struct [RadarHTTPAseIPVersionGetResponseTop0]
type radarHTTPAseIPVersionGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPAseIPVersionGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseIPVersionGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseIPVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPAseIPVersionGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPAseIPVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPAseIPVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPAseIPVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPAseIPVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPAseIPVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPAseIPVersionGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPAseIPVersionGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPAseIPVersionGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPAseIPVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP version.
type RadarHTTPAseIPVersionGetParamsIPVersion string

const (
	RadarHTTPAseIPVersionGetParamsIPVersionIPv4 RadarHTTPAseIPVersionGetParamsIPVersion = "IPv4"
	RadarHTTPAseIPVersionGetParamsIPVersionIPv6 RadarHTTPAseIPVersionGetParamsIPVersion = "IPv6"
)

type RadarHTTPAseIPVersionGetParamsBotClass string

const (
	RadarHTTPAseIPVersionGetParamsBotClassLikelyAutomated RadarHTTPAseIPVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPAseIPVersionGetParamsBotClassLikelyHuman     RadarHTTPAseIPVersionGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPAseIPVersionGetParamsDateRange string

const (
	RadarHTTPAseIPVersionGetParamsDateRange1d         RadarHTTPAseIPVersionGetParamsDateRange = "1d"
	RadarHTTPAseIPVersionGetParamsDateRange2d         RadarHTTPAseIPVersionGetParamsDateRange = "2d"
	RadarHTTPAseIPVersionGetParamsDateRange7d         RadarHTTPAseIPVersionGetParamsDateRange = "7d"
	RadarHTTPAseIPVersionGetParamsDateRange14d        RadarHTTPAseIPVersionGetParamsDateRange = "14d"
	RadarHTTPAseIPVersionGetParamsDateRange28d        RadarHTTPAseIPVersionGetParamsDateRange = "28d"
	RadarHTTPAseIPVersionGetParamsDateRange12w        RadarHTTPAseIPVersionGetParamsDateRange = "12w"
	RadarHTTPAseIPVersionGetParamsDateRange24w        RadarHTTPAseIPVersionGetParamsDateRange = "24w"
	RadarHTTPAseIPVersionGetParamsDateRange52w        RadarHTTPAseIPVersionGetParamsDateRange = "52w"
	RadarHTTPAseIPVersionGetParamsDateRange1dControl  RadarHTTPAseIPVersionGetParamsDateRange = "1dControl"
	RadarHTTPAseIPVersionGetParamsDateRange2dControl  RadarHTTPAseIPVersionGetParamsDateRange = "2dControl"
	RadarHTTPAseIPVersionGetParamsDateRange7dControl  RadarHTTPAseIPVersionGetParamsDateRange = "7dControl"
	RadarHTTPAseIPVersionGetParamsDateRange14dControl RadarHTTPAseIPVersionGetParamsDateRange = "14dControl"
	RadarHTTPAseIPVersionGetParamsDateRange28dControl RadarHTTPAseIPVersionGetParamsDateRange = "28dControl"
	RadarHTTPAseIPVersionGetParamsDateRange12wControl RadarHTTPAseIPVersionGetParamsDateRange = "12wControl"
	RadarHTTPAseIPVersionGetParamsDateRange24wControl RadarHTTPAseIPVersionGetParamsDateRange = "24wControl"
)

type RadarHTTPAseIPVersionGetParamsDeviceType string

const (
	RadarHTTPAseIPVersionGetParamsDeviceTypeDesktop RadarHTTPAseIPVersionGetParamsDeviceType = "DESKTOP"
	RadarHTTPAseIPVersionGetParamsDeviceTypeMobile  RadarHTTPAseIPVersionGetParamsDeviceType = "MOBILE"
	RadarHTTPAseIPVersionGetParamsDeviceTypeOther   RadarHTTPAseIPVersionGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPAseIPVersionGetParamsFormat string

const (
	RadarHTTPAseIPVersionGetParamsFormatJson RadarHTTPAseIPVersionGetParamsFormat = "JSON"
	RadarHTTPAseIPVersionGetParamsFormatCsv  RadarHTTPAseIPVersionGetParamsFormat = "CSV"
)

type RadarHTTPAseIPVersionGetParamsHTTPProtocol string

const (
	RadarHTTPAseIPVersionGetParamsHTTPProtocolHTTP  RadarHTTPAseIPVersionGetParamsHTTPProtocol = "HTTP"
	RadarHTTPAseIPVersionGetParamsHTTPProtocolHTTPS RadarHTTPAseIPVersionGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPAseIPVersionGetParamsHTTPVersion string

const (
	RadarHTTPAseIPVersionGetParamsHTTPVersionHttPv1 RadarHTTPAseIPVersionGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPAseIPVersionGetParamsHTTPVersionHttPv2 RadarHTTPAseIPVersionGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPAseIPVersionGetParamsHTTPVersionHttPv3 RadarHTTPAseIPVersionGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPAseIPVersionGetParamsOS string

const (
	RadarHTTPAseIPVersionGetParamsOSWindows  RadarHTTPAseIPVersionGetParamsOS = "WINDOWS"
	RadarHTTPAseIPVersionGetParamsOSMacosx   RadarHTTPAseIPVersionGetParamsOS = "MACOSX"
	RadarHTTPAseIPVersionGetParamsOSIos      RadarHTTPAseIPVersionGetParamsOS = "IOS"
	RadarHTTPAseIPVersionGetParamsOSAndroid  RadarHTTPAseIPVersionGetParamsOS = "ANDROID"
	RadarHTTPAseIPVersionGetParamsOSChromeos RadarHTTPAseIPVersionGetParamsOS = "CHROMEOS"
	RadarHTTPAseIPVersionGetParamsOSLinux    RadarHTTPAseIPVersionGetParamsOS = "LINUX"
	RadarHTTPAseIPVersionGetParamsOSSmartTv  RadarHTTPAseIPVersionGetParamsOS = "SMART_TV"
)

type RadarHTTPAseIPVersionGetParamsTLSVersion string

const (
	RadarHTTPAseIPVersionGetParamsTLSVersionTlSv1_0  RadarHTTPAseIPVersionGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPAseIPVersionGetParamsTLSVersionTlSv1_1  RadarHTTPAseIPVersionGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPAseIPVersionGetParamsTLSVersionTlSv1_2  RadarHTTPAseIPVersionGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPAseIPVersionGetParamsTLSVersionTlSv1_3  RadarHTTPAseIPVersionGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPAseIPVersionGetParamsTLSVersionTlSvQuic RadarHTTPAseIPVersionGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPAseIPVersionGetResponseEnvelope struct {
	Result  RadarHTTPAseIPVersionGetResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarHTTPAseIPVersionGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPAseIPVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarHTTPAseIPVersionGetResponseEnvelope]
type radarHTTPAseIPVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseIPVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseIPVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
