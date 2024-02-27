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

// RadarHTTPAseHTTPMethodService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPAseHTTPMethodService]
// method instead.
type RadarHTTPAseHTTPMethodService struct {
	Options []option.RequestOption
}

// NewRadarHTTPAseHTTPMethodService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPAseHTTPMethodService(opts ...option.RequestOption) (r *RadarHTTPAseHTTPMethodService) {
	r = &RadarHTTPAseHTTPMethodService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested HTTP
// protocol version. Values are a percentage out of the total traffic.
func (r *RadarHTTPAseHTTPMethodService) Get(ctx context.Context, httpVersion RadarHTTPAseHTTPMethodGetParamsHTTPVersion, query RadarHTTPAseHTTPMethodGetParams, opts ...option.RequestOption) (res *RadarHTTPAseHTTPMethodGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPAseHTTPMethodGetResponseEnvelope
	path := fmt.Sprintf("radar/http/top/ases/http_version/%v", httpVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPAseHTTPMethodGetResponse struct {
	Meta RadarHTTPAseHTTPMethodGetResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPAseHTTPMethodGetResponseTop0 `json:"top_0,required"`
	JSON radarHTTPAseHTTPMethodGetResponseJSON   `json:"-"`
}

// radarHTTPAseHTTPMethodGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPAseHTTPMethodGetResponse]
type radarHTTPAseHTTPMethodGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseHTTPMethodGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseHTTPMethodGetResponseMeta struct {
	DateRange      []RadarHTTPAseHTTPMethodGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPAseHTTPMethodGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPAseHTTPMethodGetResponseMetaJSON           `json:"-"`
}

// radarHTTPAseHTTPMethodGetResponseMetaJSON contains the JSON metadata for the
// struct [RadarHTTPAseHTTPMethodGetResponseMeta]
type radarHTTPAseHTTPMethodGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPAseHTTPMethodGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseHTTPMethodGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPAseHTTPMethodGetResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPAseHTTPMethodGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarHTTPAseHTTPMethodGetResponseMetaDateRange]
type radarHTTPAseHTTPMethodGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseHTTPMethodGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseHTTPMethodGetResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPAseHTTPMethodGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        radarHTTPAseHTTPMethodGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPAseHTTPMethodGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarHTTPAseHTTPMethodGetResponseMetaConfidenceInfo]
type radarHTTPAseHTTPMethodGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseHTTPMethodGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseHTTPMethodGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            radarHTTPAseHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPAseHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarHTTPAseHTTPMethodGetResponseMetaConfidenceInfoAnnotation]
type radarHTTPAseHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPAseHTTPMethodGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseHTTPMethodGetResponseTop0 struct {
	ClientASN    int64                                     `json:"clientASN,required"`
	ClientAsName string                                    `json:"clientASName,required"`
	Value        string                                    `json:"value,required"`
	JSON         radarHTTPAseHTTPMethodGetResponseTop0JSON `json:"-"`
}

// radarHTTPAseHTTPMethodGetResponseTop0JSON contains the JSON metadata for the
// struct [RadarHTTPAseHTTPMethodGetResponseTop0]
type radarHTTPAseHTTPMethodGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPAseHTTPMethodGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPAseHTTPMethodGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPAseHTTPMethodGetParamsBotClass] `query:"botClass"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPAseHTTPMethodGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPAseHTTPMethodGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPAseHTTPMethodGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPAseHTTPMethodGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPAseHTTPMethodGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPAseHTTPMethodGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPAseHTTPMethodGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPAseHTTPMethodGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPAseHTTPMethodGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// HTTP version.
type RadarHTTPAseHTTPMethodGetParamsHTTPVersion string

const (
	RadarHTTPAseHTTPMethodGetParamsHTTPVersionHttPv1 RadarHTTPAseHTTPMethodGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPAseHTTPMethodGetParamsHTTPVersionHttPv2 RadarHTTPAseHTTPMethodGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPAseHTTPMethodGetParamsHTTPVersionHttPv3 RadarHTTPAseHTTPMethodGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPAseHTTPMethodGetParamsBotClass string

const (
	RadarHTTPAseHTTPMethodGetParamsBotClassLikelyAutomated RadarHTTPAseHTTPMethodGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPAseHTTPMethodGetParamsBotClassLikelyHuman     RadarHTTPAseHTTPMethodGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPAseHTTPMethodGetParamsDateRange string

const (
	RadarHTTPAseHTTPMethodGetParamsDateRange1d         RadarHTTPAseHTTPMethodGetParamsDateRange = "1d"
	RadarHTTPAseHTTPMethodGetParamsDateRange2d         RadarHTTPAseHTTPMethodGetParamsDateRange = "2d"
	RadarHTTPAseHTTPMethodGetParamsDateRange7d         RadarHTTPAseHTTPMethodGetParamsDateRange = "7d"
	RadarHTTPAseHTTPMethodGetParamsDateRange14d        RadarHTTPAseHTTPMethodGetParamsDateRange = "14d"
	RadarHTTPAseHTTPMethodGetParamsDateRange28d        RadarHTTPAseHTTPMethodGetParamsDateRange = "28d"
	RadarHTTPAseHTTPMethodGetParamsDateRange12w        RadarHTTPAseHTTPMethodGetParamsDateRange = "12w"
	RadarHTTPAseHTTPMethodGetParamsDateRange24w        RadarHTTPAseHTTPMethodGetParamsDateRange = "24w"
	RadarHTTPAseHTTPMethodGetParamsDateRange52w        RadarHTTPAseHTTPMethodGetParamsDateRange = "52w"
	RadarHTTPAseHTTPMethodGetParamsDateRange1dControl  RadarHTTPAseHTTPMethodGetParamsDateRange = "1dControl"
	RadarHTTPAseHTTPMethodGetParamsDateRange2dControl  RadarHTTPAseHTTPMethodGetParamsDateRange = "2dControl"
	RadarHTTPAseHTTPMethodGetParamsDateRange7dControl  RadarHTTPAseHTTPMethodGetParamsDateRange = "7dControl"
	RadarHTTPAseHTTPMethodGetParamsDateRange14dControl RadarHTTPAseHTTPMethodGetParamsDateRange = "14dControl"
	RadarHTTPAseHTTPMethodGetParamsDateRange28dControl RadarHTTPAseHTTPMethodGetParamsDateRange = "28dControl"
	RadarHTTPAseHTTPMethodGetParamsDateRange12wControl RadarHTTPAseHTTPMethodGetParamsDateRange = "12wControl"
	RadarHTTPAseHTTPMethodGetParamsDateRange24wControl RadarHTTPAseHTTPMethodGetParamsDateRange = "24wControl"
)

type RadarHTTPAseHTTPMethodGetParamsDeviceType string

const (
	RadarHTTPAseHTTPMethodGetParamsDeviceTypeDesktop RadarHTTPAseHTTPMethodGetParamsDeviceType = "DESKTOP"
	RadarHTTPAseHTTPMethodGetParamsDeviceTypeMobile  RadarHTTPAseHTTPMethodGetParamsDeviceType = "MOBILE"
	RadarHTTPAseHTTPMethodGetParamsDeviceTypeOther   RadarHTTPAseHTTPMethodGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPAseHTTPMethodGetParamsFormat string

const (
	RadarHTTPAseHTTPMethodGetParamsFormatJson RadarHTTPAseHTTPMethodGetParamsFormat = "JSON"
	RadarHTTPAseHTTPMethodGetParamsFormatCsv  RadarHTTPAseHTTPMethodGetParamsFormat = "CSV"
)

type RadarHTTPAseHTTPMethodGetParamsHTTPProtocol string

const (
	RadarHTTPAseHTTPMethodGetParamsHTTPProtocolHTTP  RadarHTTPAseHTTPMethodGetParamsHTTPProtocol = "HTTP"
	RadarHTTPAseHTTPMethodGetParamsHTTPProtocolHTTPS RadarHTTPAseHTTPMethodGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPAseHTTPMethodGetParamsIPVersion string

const (
	RadarHTTPAseHTTPMethodGetParamsIPVersionIPv4 RadarHTTPAseHTTPMethodGetParamsIPVersion = "IPv4"
	RadarHTTPAseHTTPMethodGetParamsIPVersionIPv6 RadarHTTPAseHTTPMethodGetParamsIPVersion = "IPv6"
)

type RadarHTTPAseHTTPMethodGetParamsOS string

const (
	RadarHTTPAseHTTPMethodGetParamsOSWindows  RadarHTTPAseHTTPMethodGetParamsOS = "WINDOWS"
	RadarHTTPAseHTTPMethodGetParamsOSMacosx   RadarHTTPAseHTTPMethodGetParamsOS = "MACOSX"
	RadarHTTPAseHTTPMethodGetParamsOSIos      RadarHTTPAseHTTPMethodGetParamsOS = "IOS"
	RadarHTTPAseHTTPMethodGetParamsOSAndroid  RadarHTTPAseHTTPMethodGetParamsOS = "ANDROID"
	RadarHTTPAseHTTPMethodGetParamsOSChromeos RadarHTTPAseHTTPMethodGetParamsOS = "CHROMEOS"
	RadarHTTPAseHTTPMethodGetParamsOSLinux    RadarHTTPAseHTTPMethodGetParamsOS = "LINUX"
	RadarHTTPAseHTTPMethodGetParamsOSSmartTv  RadarHTTPAseHTTPMethodGetParamsOS = "SMART_TV"
)

type RadarHTTPAseHTTPMethodGetParamsTLSVersion string

const (
	RadarHTTPAseHTTPMethodGetParamsTLSVersionTlSv1_0  RadarHTTPAseHTTPMethodGetParamsTLSVersion = "TLSv1_0"
	RadarHTTPAseHTTPMethodGetParamsTLSVersionTlSv1_1  RadarHTTPAseHTTPMethodGetParamsTLSVersion = "TLSv1_1"
	RadarHTTPAseHTTPMethodGetParamsTLSVersionTlSv1_2  RadarHTTPAseHTTPMethodGetParamsTLSVersion = "TLSv1_2"
	RadarHTTPAseHTTPMethodGetParamsTLSVersionTlSv1_3  RadarHTTPAseHTTPMethodGetParamsTLSVersion = "TLSv1_3"
	RadarHTTPAseHTTPMethodGetParamsTLSVersionTlSvQuic RadarHTTPAseHTTPMethodGetParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPAseHTTPMethodGetResponseEnvelope struct {
	Result  RadarHTTPAseHTTPMethodGetResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarHTTPAseHTTPMethodGetResponseEnvelopeJSON `json:"-"`
}

// radarHTTPAseHTTPMethodGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarHTTPAseHTTPMethodGetResponseEnvelope]
type radarHTTPAseHTTPMethodGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPAseHTTPMethodGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
