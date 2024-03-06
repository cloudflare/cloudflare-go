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

// RadarHTTPAseOSService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPAseOSService] method
// instead.
type RadarHTTPAseOSService struct {
	Options []option.RequestOption
}

// NewRadarHTTPAseOSService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarHTTPAseOSService(opts ...option.RequestOption) (r *RadarHTTPAseOSService) {
	r = &RadarHTTPAseOSService{}
	r.Options = opts
	return
}

// Get the top autonomous systems, by HTTP traffic, of the requested operating
// systems. Values are a percentage out of the total traffic.
func (r *RadarHTTPAseOSService) Get(ctx context.Context, os RadarHTTPAseOSGetParamsOS, query RadarHTTPAseOSGetParams, opts ...option.RequestOption) (res *RadarHTTPAseOSGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPAseOSGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/ases/os/%v", os)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPAseOSGetResponse struct {
	Meta RadarHTTPAseOSGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPAseOSGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPAseOSGetResponseJSON   `json:"-"`
}

// radarHTTPAseOSGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPAseOSGetResponse]
type radarHTTPAseOSGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseOSGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseOSGetResponseJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseOSGetResponseMeta struct {
	DateRange      []RadarHTTPAseOSGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                      `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPAseOSGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPAseOSGetResponseMetaJSON           `json:"-"`
}

// radarHTTPAseOSGetResponseMetaJSON contains the JSON metadata for the struct
// [RadarHTTPAseOSGetResponseMeta]
type radarHTTPAseOSGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPAseOSGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseOSGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseOSGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                  `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPAseOSGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPAseOSGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarHTTPAseOSGetResponseMetaDateRange]
type radarHTTPAseOSGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseOSGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseOSGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseOSGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPAseOSGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                   `json:"level"`
	JSON        radarHTTPAseOSGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPAseOSGetResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [RadarHTTPAseOSGetResponseMetaConfidenceInfo]
type radarHTTPAseOSGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseOSGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseOSGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseOSGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                    `json:"dataSource,required"`
	Description     string                                                    `json:"description,required"`
	EventType       string                                                    `json:"eventType,required"`
	IsInstantaneous interface{}                                               `json:"isInstantaneous,required"`
	EndTime         time.Time                                                 `json:"endTime" format:"date-time"`
	LinkedURL       string                                                    `json:"linkedUrl"`
	StartTime       time.Time                                                 `json:"startTime" format:"date-time"`
	JSON            radarHTTPAseOSGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPAseOSGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [RadarHTTPAseOSGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPAseOSGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPAseOSGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseOSGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseOSGetResponseTop0 struct {
	ClientASN    int64                             `json:"clientASN,required"`
	ClientAsName string                            `json:"clientASName,required"`
	Value        string                            `json:"value,required"`
	JSON         radarHTTPAseOSGetResponseTop0JSON `json:"-"`
}

// radarHTTPAseOSGetResponseTop0JSON contains the JSON metadata for the struct
// [RadarHTTPAseOSGetResponseTop0]
type radarHTTPAseOSGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPAseOSGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseOSGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type RadarHTTPAseOSGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPAseOSGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPAseOSGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPAseOSGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPAseOSGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPAseOSGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPAseOSGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPAseOSGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPAseOSGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPAseOSGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPAseOSGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP version.
type RadarHTTPAseOSGetParamsOS string

const (
	RadarHTTPAseOSGetParamsOSWindows  RadarHTTPAseOSGetParamsOS = "WINDOWS"
	RadarHTTPAseOSGetParamsOSMacosx   RadarHTTPAseOSGetParamsOS = "MACOSX"
	RadarHTTPAseOSGetParamsOSIos      RadarHTTPAseOSGetParamsOS = "IOS"
	RadarHTTPAseOSGetParamsOSAndroid  RadarHTTPAseOSGetParamsOS = "ANDROID"
	RadarHTTPAseOSGetParamsOSChromeos RadarHTTPAseOSGetParamsOS = "CHROMEOS"
	RadarHTTPAseOSGetParamsOSLinux    RadarHTTPAseOSGetParamsOS = "LINUX"
	RadarHTTPAseOSGetParamsOSSmartTv  RadarHTTPAseOSGetParamsOS = "SMART_TV"
)

type RadarHTTPAseOSGetParamsBotClass string

const (
	RadarHTTPAseOSGetParamsBotClassLikelyAutomated RadarHTTPAseOSGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPAseOSGetParamsBotClassLikelyHuman     RadarHTTPAseOSGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPAseOSGetParamsDateRange string

const (
	RadarHTTPAseOSGetParamsDateRange1d         RadarHTTPAseOSGetParamsDateRange = "1d"
	RadarHTTPAseOSGetParamsDateRange2d         RadarHTTPAseOSGetParamsDateRange = "2d"
	RadarHTTPAseOSGetParamsDateRange7d         RadarHTTPAseOSGetParamsDateRange = "7d"
	RadarHTTPAseOSGetParamsDateRange14d        RadarHTTPAseOSGetParamsDateRange = "14d"
	RadarHTTPAseOSGetParamsDateRange28d        RadarHTTPAseOSGetParamsDateRange = "28d"
	RadarHTTPAseOSGetParamsDateRange12w        RadarHTTPAseOSGetParamsDateRange = "12w"
	RadarHTTPAseOSGetParamsDateRange24w        RadarHTTPAseOSGetParamsDateRange = "24w"
	RadarHTTPAseOSGetParamsDateRange52w        RadarHTTPAseOSGetParamsDateRange = "52w"
	RadarHTTPAseOSGetParamsDateRange1dControl  RadarHTTPAseOSGetParamsDateRange = "1dControl"
	RadarHTTPAseOSGetParamsDateRange2dControl  RadarHTTPAseOSGetParamsDateRange = "2dControl"
	RadarHTTPAseOSGetParamsDateRange7dControl  RadarHTTPAseOSGetParamsDateRange = "7dControl"
	RadarHTTPAseOSGetParamsDateRange14dControl RadarHTTPAseOSGetParamsDateRange = "14dControl"
	RadarHTTPAseOSGetParamsDateRange28dControl RadarHTTPAseOSGetParamsDateRange = "28dControl"
	RadarHTTPAseOSGetParamsDateRange12wControl RadarHTTPAseOSGetParamsDateRange = "12wControl"
	RadarHTTPAseOSGetParamsDateRange24wControl RadarHTTPAseOSGetParamsDateRange = "24wControl"
)

type RadarHTTPAseOSGetParamsDeviceType string

const (
	RadarHTTPAseOSGetParamsDeviceTypeDesktop RadarHTTPAseOSGetParamsDeviceType = "DESKTOP"
	RadarHTTPAseOSGetParamsDeviceTypeMobile  RadarHTTPAseOSGetParamsDeviceType = "MOBILE"
	RadarHTTPAseOSGetParamsDeviceTypeOther   RadarHTTPAseOSGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPAseOSGetParamsFormat string

const (
	RadarHTTPAseOSGetParamsFormatJson RadarHTTPAseOSGetParamsFormat = "JSON"
	RadarHTTPAseOSGetParamsFormatCsv  RadarHTTPAseOSGetParamsFormat = "CSV"
)

type RadarHTTPAseOSGetParamsHTTPProtocol string

const (
	RadarHTTPAseOSGetParamsHTTPProtocolHTTP  RadarHTTPAseOSGetParamsHTTPProtocol = "HTTP"
	RadarHTTPAseOSGetParamsHTTPProtocolHTTPS RadarHTTPAseOSGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPAseOSGetParamsHTTPVersion string

const (
	RadarHTTPAseOSGetParamsHTTPVersionHttPv1 RadarHTTPAseOSGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPAseOSGetParamsHTTPVersionHttPv2 RadarHTTPAseOSGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPAseOSGetParamsHTTPVersionHttPv3 RadarHTTPAseOSGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPAseOSGetParamsIPVersion string

const (
	RadarHTTPAseOSGetParamsIPVersionIPv4 RadarHTTPAseOSGetParamsIPVersion = "IPv4"
	RadarHTTPAseOSGetParamsIPVersionIPv6 RadarHTTPAseOSGetParamsIPVersion = "IPv6"
)

type RadarHTTPAseOSGetParamsTLSVersion string

const (
	RadarHTTPAseOSGetParamsTLSVersionTlSv1_0  RadarHTTPAseOSGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPAseOSGetParamsTLSVersionTlSv1_1  RadarHTTPAseOSGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPAseOSGetParamsTLSVersionTlSv1_2  RadarHTTPAseOSGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPAseOSGetParamsTLSVersionTlSv1_3  RadarHTTPAseOSGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPAseOSGetParamsTLSVersionTlSvQuic RadarHTTPAseOSGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPAseOSGetResponseEnvelope struct {
	Result  RadarHTTPAseOSGetResponse             `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarHTTPAseOSGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPAseOSGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RadarHTTPAseOSGetResponseEnvelope]
type radarHTTPAseOSGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseOSGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r radarHTTPAseOSGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
