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

// RadarHTTPTopAseOService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPTopAseOService] method
// instead.
type RadarHTTPTopAseOService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopAseOService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPTopAseOService(opts ...option.RequestOption) (r *RadarHTTPTopAseOService) {
	r = &RadarHTTPTopAseOService{}
	r.Options = opts
	return
}

// Get the top autonomous systems, by HTTP traffic, of the requested operating
// systems. Values are a percentage out of the total traffic.
func (r *RadarHTTPTopAseOService) Get(ctx context.Context, os RadarHTTPTopAseOGetParamsOs, query RadarHTTPTopAseOGetParams, opts ...option.RequestOption) (res *RadarHTTPTopAseOGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/os/%v", os)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopAseOGetResponse struct {
	Result  RadarHTTPTopAseOGetResponseResult `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    radarHTTPTopAseOGetResponseJSON   `json:"-"`
}

// radarHTTPTopAseOGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTopAseOGetResponse]
type radarHTTPTopAseOGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseOGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseOGetResponseResult struct {
	Meta RadarHTTPTopAseOGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseOGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseOGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseOGetResponseResultJSON contains the JSON metadata for the struct
// [RadarHTTPTopAseOGetResponseResult]
type radarHTTPTopAseOGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseOGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseOGetResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseOGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseOGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseOGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseOGetResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseOGetResponseResultMeta]
type radarHTTPTopAseOGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseOGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseOGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseOGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseOGetResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseOGetResponseResultMetaDateRange]
type radarHTTPTopAseOGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseOGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseOGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseOGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        radarHTTPTopAseOGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseOGetResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarHTTPTopAseOGetResponseResultMetaConfidenceInfo]
type radarHTTPTopAseOGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseOGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseOGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseOGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseOGetResponseResultMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseOGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseOGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseOGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseOGetResponseResultTop0 struct {
	ClientASN    int64                                     `json:"clientASN,required"`
	ClientAsName string                                    `json:"clientASName,required"`
	Value        string                                    `json:"value,required"`
	JSON         radarHTTPTopAseOGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseOGetResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarHTTPTopAseOGetResponseResultTop0]
type radarHTTPTopAseOGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseOGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseOGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseOGetParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopAseOGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopAseOGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopAseOGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopAseOGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopAseOGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopAseOGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopAseOGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseOGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopAseOGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP version.
type RadarHTTPTopAseOGetParamsOs string

const (
	RadarHTTPTopAseOGetParamsOsWindows  RadarHTTPTopAseOGetParamsOs = "WINDOWS"
	RadarHTTPTopAseOGetParamsOsMacosx   RadarHTTPTopAseOGetParamsOs = "MACOSX"
	RadarHTTPTopAseOGetParamsOsIos      RadarHTTPTopAseOGetParamsOs = "IOS"
	RadarHTTPTopAseOGetParamsOsAndroid  RadarHTTPTopAseOGetParamsOs = "ANDROID"
	RadarHTTPTopAseOGetParamsOsChromeos RadarHTTPTopAseOGetParamsOs = "CHROMEOS"
	RadarHTTPTopAseOGetParamsOsLinux    RadarHTTPTopAseOGetParamsOs = "LINUX"
	RadarHTTPTopAseOGetParamsOsSmartTv  RadarHTTPTopAseOGetParamsOs = "SMART_TV"
)

type RadarHTTPTopAseOGetParamsBotClass string

const (
	RadarHTTPTopAseOGetParamsBotClassLikelyAutomated RadarHTTPTopAseOGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseOGetParamsBotClassLikelyHuman     RadarHTTPTopAseOGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopAseOGetParamsDateRange string

const (
	RadarHTTPTopAseOGetParamsDateRange1d         RadarHTTPTopAseOGetParamsDateRange = "1d"
	RadarHTTPTopAseOGetParamsDateRange2d         RadarHTTPTopAseOGetParamsDateRange = "2d"
	RadarHTTPTopAseOGetParamsDateRange7d         RadarHTTPTopAseOGetParamsDateRange = "7d"
	RadarHTTPTopAseOGetParamsDateRange14d        RadarHTTPTopAseOGetParamsDateRange = "14d"
	RadarHTTPTopAseOGetParamsDateRange28d        RadarHTTPTopAseOGetParamsDateRange = "28d"
	RadarHTTPTopAseOGetParamsDateRange12w        RadarHTTPTopAseOGetParamsDateRange = "12w"
	RadarHTTPTopAseOGetParamsDateRange24w        RadarHTTPTopAseOGetParamsDateRange = "24w"
	RadarHTTPTopAseOGetParamsDateRange52w        RadarHTTPTopAseOGetParamsDateRange = "52w"
	RadarHTTPTopAseOGetParamsDateRange1dControl  RadarHTTPTopAseOGetParamsDateRange = "1dControl"
	RadarHTTPTopAseOGetParamsDateRange2dControl  RadarHTTPTopAseOGetParamsDateRange = "2dControl"
	RadarHTTPTopAseOGetParamsDateRange7dControl  RadarHTTPTopAseOGetParamsDateRange = "7dControl"
	RadarHTTPTopAseOGetParamsDateRange14dControl RadarHTTPTopAseOGetParamsDateRange = "14dControl"
	RadarHTTPTopAseOGetParamsDateRange28dControl RadarHTTPTopAseOGetParamsDateRange = "28dControl"
	RadarHTTPTopAseOGetParamsDateRange12wControl RadarHTTPTopAseOGetParamsDateRange = "12wControl"
	RadarHTTPTopAseOGetParamsDateRange24wControl RadarHTTPTopAseOGetParamsDateRange = "24wControl"
)

type RadarHTTPTopAseOGetParamsDeviceType string

const (
	RadarHTTPTopAseOGetParamsDeviceTypeDesktop RadarHTTPTopAseOGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseOGetParamsDeviceTypeMobile  RadarHTTPTopAseOGetParamsDeviceType = "MOBILE"
	RadarHTTPTopAseOGetParamsDeviceTypeOther   RadarHTTPTopAseOGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopAseOGetParamsFormat string

const (
	RadarHTTPTopAseOGetParamsFormatJson RadarHTTPTopAseOGetParamsFormat = "JSON"
	RadarHTTPTopAseOGetParamsFormatCsv  RadarHTTPTopAseOGetParamsFormat = "CSV"
)

type RadarHTTPTopAseOGetParamsHTTPProtocol string

const (
	RadarHTTPTopAseOGetParamsHTTPProtocolHTTP  RadarHTTPTopAseOGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseOGetParamsHTTPProtocolHTTPs RadarHTTPTopAseOGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopAseOGetParamsHTTPVersion string

const (
	RadarHTTPTopAseOGetParamsHTTPVersionHttPv1 RadarHTTPTopAseOGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseOGetParamsHTTPVersionHttPv2 RadarHTTPTopAseOGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseOGetParamsHTTPVersionHttPv3 RadarHTTPTopAseOGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopAseOGetParamsIPVersion string

const (
	RadarHTTPTopAseOGetParamsIPVersionIPv4 RadarHTTPTopAseOGetParamsIPVersion = "IPv4"
	RadarHTTPTopAseOGetParamsIPVersionIPv6 RadarHTTPTopAseOGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopAseOGetParamsTlsVersion string

const (
	RadarHTTPTopAseOGetParamsTlsVersionTlSv1_0  RadarHTTPTopAseOGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseOGetParamsTlsVersionTlSv1_1  RadarHTTPTopAseOGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseOGetParamsTlsVersionTlSv1_2  RadarHTTPTopAseOGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseOGetParamsTlsVersionTlSv1_3  RadarHTTPTopAseOGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseOGetParamsTlsVersionTlSvQuic RadarHTTPTopAseOGetParamsTlsVersion = "TLSvQUIC"
)
