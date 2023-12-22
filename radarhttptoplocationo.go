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

// RadarHTTPTopLocationOService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPTopLocationOService]
// method instead.
type RadarHTTPTopLocationOService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopLocationOService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPTopLocationOService(opts ...option.RequestOption) (r *RadarHTTPTopLocationOService) {
	r = &RadarHTTPTopLocationOService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested operating systems.
// Values are a percentage out of the total traffic.
func (r *RadarHTTPTopLocationOService) Get(ctx context.Context, os RadarHTTPTopLocationOGetParamsOs, query RadarHTTPTopLocationOGetParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationOGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/os/%v", os)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopLocationOGetResponse struct {
	Result  RadarHTTPTopLocationOGetResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarHTTPTopLocationOGetResponseJSON   `json:"-"`
}

// radarHTTPTopLocationOGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTopLocationOGetResponse]
type radarHTTPTopLocationOGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationOGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationOGetResponseResult struct {
	Meta RadarHTTPTopLocationOGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationOGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationOGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationOGetResponseResultJSON contains the JSON metadata for the
// struct [RadarHTTPTopLocationOGetResponseResult]
type radarHTTPTopLocationOGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationOGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationOGetResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationOGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarHTTPTopLocationOGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarHTTPTopLocationOGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopLocationOGetResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationOGetResponseResultMeta]
type radarHTTPTopLocationOGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationOGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationOGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationOGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                `json:"level,required"`
	JSON        radarHTTPTopLocationOGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopLocationOGetResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationOGetResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationOGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationOGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationOGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                 `json:"dataSource,required"`
	Description string                                                                 `json:"description,required"`
	EndTime     time.Time                                                              `json:"endTime,required" format:"date-time"`
	EventType   string                                                                 `json:"eventType,required"`
	StartTime   time.Time                                                              `json:"startTime,required" format:"date-time"`
	JSON        radarHTTPTopLocationOGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationOGetResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationOGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationOGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationOGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationOGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                               `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationOGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationOGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationOGetResponseResultMetaDateRange]
type radarHTTPTopLocationOGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationOGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationOGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                         `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                         `json:"clientCountryName,required"`
	Value               string                                         `json:"value,required"`
	JSON                radarHTTPTopLocationOGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationOGetResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationOGetResponseResultTop0]
type radarHTTPTopLocationOGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationOGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationOGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationOGetParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopLocationOGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopLocationOGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopLocationOGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopLocationOGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationOGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopLocationOGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopLocationOGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationOGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopLocationOGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopLocationOGetParamsOs string

const (
	RadarHTTPTopLocationOGetParamsOsWindows  RadarHTTPTopLocationOGetParamsOs = "WINDOWS"
	RadarHTTPTopLocationOGetParamsOsMacosx   RadarHTTPTopLocationOGetParamsOs = "MACOSX"
	RadarHTTPTopLocationOGetParamsOsAndroid  RadarHTTPTopLocationOGetParamsOs = "ANDROID"
	RadarHTTPTopLocationOGetParamsOsChromeos RadarHTTPTopLocationOGetParamsOs = "CHROMEOS"
	RadarHTTPTopLocationOGetParamsOsLinux    RadarHTTPTopLocationOGetParamsOs = "LINUX"
	RadarHTTPTopLocationOGetParamsOsSmartTv  RadarHTTPTopLocationOGetParamsOs = "SMART_TV"
)

type RadarHTTPTopLocationOGetParamsBotClass string

const (
	RadarHTTPTopLocationOGetParamsBotClassLikelyAutomated RadarHTTPTopLocationOGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationOGetParamsBotClassLikelyHuman     RadarHTTPTopLocationOGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopLocationOGetParamsDateRange string

const (
	RadarHTTPTopLocationOGetParamsDateRange1d         RadarHTTPTopLocationOGetParamsDateRange = "1d"
	RadarHTTPTopLocationOGetParamsDateRange7d         RadarHTTPTopLocationOGetParamsDateRange = "7d"
	RadarHTTPTopLocationOGetParamsDateRange14d        RadarHTTPTopLocationOGetParamsDateRange = "14d"
	RadarHTTPTopLocationOGetParamsDateRange28d        RadarHTTPTopLocationOGetParamsDateRange = "28d"
	RadarHTTPTopLocationOGetParamsDateRange12w        RadarHTTPTopLocationOGetParamsDateRange = "12w"
	RadarHTTPTopLocationOGetParamsDateRange24w        RadarHTTPTopLocationOGetParamsDateRange = "24w"
	RadarHTTPTopLocationOGetParamsDateRange52w        RadarHTTPTopLocationOGetParamsDateRange = "52w"
	RadarHTTPTopLocationOGetParamsDateRange1dControl  RadarHTTPTopLocationOGetParamsDateRange = "1dControl"
	RadarHTTPTopLocationOGetParamsDateRange7dControl  RadarHTTPTopLocationOGetParamsDateRange = "7dControl"
	RadarHTTPTopLocationOGetParamsDateRange14dControl RadarHTTPTopLocationOGetParamsDateRange = "14dControl"
	RadarHTTPTopLocationOGetParamsDateRange28dControl RadarHTTPTopLocationOGetParamsDateRange = "28dControl"
	RadarHTTPTopLocationOGetParamsDateRange12wControl RadarHTTPTopLocationOGetParamsDateRange = "12wControl"
	RadarHTTPTopLocationOGetParamsDateRange24wControl RadarHTTPTopLocationOGetParamsDateRange = "24wControl"
)

type RadarHTTPTopLocationOGetParamsDeviceType string

const (
	RadarHTTPTopLocationOGetParamsDeviceTypeDesktop RadarHTTPTopLocationOGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationOGetParamsDeviceTypeMobile  RadarHTTPTopLocationOGetParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationOGetParamsDeviceTypeOther   RadarHTTPTopLocationOGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopLocationOGetParamsFormat string

const (
	RadarHTTPTopLocationOGetParamsFormatJson RadarHTTPTopLocationOGetParamsFormat = "JSON"
	RadarHTTPTopLocationOGetParamsFormatCsv  RadarHTTPTopLocationOGetParamsFormat = "CSV"
)

type RadarHTTPTopLocationOGetParamsHTTPProtocol string

const (
	RadarHTTPTopLocationOGetParamsHTTPProtocolHTTP  RadarHTTPTopLocationOGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationOGetParamsHTTPProtocolHTTPs RadarHTTPTopLocationOGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopLocationOGetParamsHTTPVersion string

const (
	RadarHTTPTopLocationOGetParamsHTTPVersionHttPv1 RadarHTTPTopLocationOGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationOGetParamsHTTPVersionHttPv2 RadarHTTPTopLocationOGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationOGetParamsHTTPVersionHttPv3 RadarHTTPTopLocationOGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopLocationOGetParamsIPVersion string

const (
	RadarHTTPTopLocationOGetParamsIPVersionIPv4 RadarHTTPTopLocationOGetParamsIPVersion = "IPv4"
	RadarHTTPTopLocationOGetParamsIPVersionIPv6 RadarHTTPTopLocationOGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopLocationOGetParamsTlsVersion string

const (
	RadarHTTPTopLocationOGetParamsTlsVersionTlSv1_0  RadarHTTPTopLocationOGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationOGetParamsTlsVersionTlSv1_1  RadarHTTPTopLocationOGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationOGetParamsTlsVersionTlSv1_2  RadarHTTPTopLocationOGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationOGetParamsTlsVersionTlSv1_3  RadarHTTPTopLocationOGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationOGetParamsTlsVersionTlSvQuic RadarHTTPTopLocationOGetParamsTlsVersion = "TLSvQUIC"
)
