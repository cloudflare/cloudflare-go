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

// RadarHTTPTopAseTlsVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTopAseTlsVersionService] method instead.
type RadarHTTPTopAseTlsVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopAseTlsVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTopAseTlsVersionService(opts ...option.RequestOption) (r *RadarHTTPTopAseTlsVersionService) {
	r = &RadarHTTPTopAseTlsVersionService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested TLS
// protocol version. Values are a percentage out of the total traffic.
func (r *RadarHTTPTopAseTlsVersionService) Get(ctx context.Context, tlsVersion RadarHTTPTopAseTlsVersionGetParamsTlsVersion, query RadarHTTPTopAseTlsVersionGetParams, opts ...option.RequestOption) (res *RadarHTTPTopAseTlsVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/tls_version/%v", tlsVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopAseTlsVersionGetResponse struct {
	Result  RadarHTTPTopAseTlsVersionGetResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarHTTPTopAseTlsVersionGetResponseJSON   `json:"-"`
}

// radarHTTPTopAseTlsVersionGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseTlsVersionGetResponse]
type radarHTTPTopAseTlsVersionGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseTlsVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseTlsVersionGetResponseResult struct {
	Meta RadarHTTPTopAseTlsVersionGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseTlsVersionGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseTlsVersionGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseTlsVersionGetResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTopAseTlsVersionGetResponseResult]
type radarHTTPTopAseTlsVersionGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseTlsVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseTlsVersionGetResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseTlsVersionGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                       `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseTlsVersionGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseTlsVersionGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseTlsVersionGetResponseResultMeta]
type radarHTTPTopAseTlsVersionGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseTlsVersionGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseTlsVersionGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseTlsVersionGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseTlsVersionGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopAseTlsVersionGetResponseResultMetaDateRange]
type radarHTTPTopAseTlsVersionGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseTlsVersionGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                    `json:"level"`
	JSON        radarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfo]
type radarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                     `json:"dataSource,required"`
	Description     string                                                                     `json:"description,required"`
	EventType       string                                                                     `json:"eventType,required"`
	IsInstantaneous interface{}                                                                `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                                  `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseTlsVersionGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseTlsVersionGetResponseResultTop0 struct {
	ClientASN    int64                                              `json:"clientASN,required"`
	ClientAsName string                                             `json:"clientASName,required"`
	Value        string                                             `json:"value,required"`
	JSON         radarHTTPTopAseTlsVersionGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseTlsVersionGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarHTTPTopAseTlsVersionGetResponseResultTop0]
type radarHTTPTopAseTlsVersionGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseTlsVersionGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseTlsVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseTlsVersionGetParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopAseTlsVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopAseTlsVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopAseTlsVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopAseTlsVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopAseTlsVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopAseTlsVersionGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopAseTlsVersionGetParamsO] `query:"os"`
}

// URLQuery serializes [RadarHTTPTopAseTlsVersionGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopAseTlsVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// TLS version.
type RadarHTTPTopAseTlsVersionGetParamsTlsVersion string

const (
	RadarHTTPTopAseTlsVersionGetParamsTlsVersionTlSv1_0  RadarHTTPTopAseTlsVersionGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseTlsVersionGetParamsTlsVersionTlSv1_1  RadarHTTPTopAseTlsVersionGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseTlsVersionGetParamsTlsVersionTlSv1_2  RadarHTTPTopAseTlsVersionGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseTlsVersionGetParamsTlsVersionTlSv1_3  RadarHTTPTopAseTlsVersionGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseTlsVersionGetParamsTlsVersionTlSvQuic RadarHTTPTopAseTlsVersionGetParamsTlsVersion = "TLSvQUIC"
)

type RadarHTTPTopAseTlsVersionGetParamsBotClass string

const (
	RadarHTTPTopAseTlsVersionGetParamsBotClassLikelyAutomated RadarHTTPTopAseTlsVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseTlsVersionGetParamsBotClassLikelyHuman     RadarHTTPTopAseTlsVersionGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopAseTlsVersionGetParamsDateRange string

const (
	RadarHTTPTopAseTlsVersionGetParamsDateRange1d         RadarHTTPTopAseTlsVersionGetParamsDateRange = "1d"
	RadarHTTPTopAseTlsVersionGetParamsDateRange2d         RadarHTTPTopAseTlsVersionGetParamsDateRange = "2d"
	RadarHTTPTopAseTlsVersionGetParamsDateRange7d         RadarHTTPTopAseTlsVersionGetParamsDateRange = "7d"
	RadarHTTPTopAseTlsVersionGetParamsDateRange14d        RadarHTTPTopAseTlsVersionGetParamsDateRange = "14d"
	RadarHTTPTopAseTlsVersionGetParamsDateRange28d        RadarHTTPTopAseTlsVersionGetParamsDateRange = "28d"
	RadarHTTPTopAseTlsVersionGetParamsDateRange12w        RadarHTTPTopAseTlsVersionGetParamsDateRange = "12w"
	RadarHTTPTopAseTlsVersionGetParamsDateRange24w        RadarHTTPTopAseTlsVersionGetParamsDateRange = "24w"
	RadarHTTPTopAseTlsVersionGetParamsDateRange52w        RadarHTTPTopAseTlsVersionGetParamsDateRange = "52w"
	RadarHTTPTopAseTlsVersionGetParamsDateRange1dControl  RadarHTTPTopAseTlsVersionGetParamsDateRange = "1dControl"
	RadarHTTPTopAseTlsVersionGetParamsDateRange2dControl  RadarHTTPTopAseTlsVersionGetParamsDateRange = "2dControl"
	RadarHTTPTopAseTlsVersionGetParamsDateRange7dControl  RadarHTTPTopAseTlsVersionGetParamsDateRange = "7dControl"
	RadarHTTPTopAseTlsVersionGetParamsDateRange14dControl RadarHTTPTopAseTlsVersionGetParamsDateRange = "14dControl"
	RadarHTTPTopAseTlsVersionGetParamsDateRange28dControl RadarHTTPTopAseTlsVersionGetParamsDateRange = "28dControl"
	RadarHTTPTopAseTlsVersionGetParamsDateRange12wControl RadarHTTPTopAseTlsVersionGetParamsDateRange = "12wControl"
	RadarHTTPTopAseTlsVersionGetParamsDateRange24wControl RadarHTTPTopAseTlsVersionGetParamsDateRange = "24wControl"
)

type RadarHTTPTopAseTlsVersionGetParamsDeviceType string

const (
	RadarHTTPTopAseTlsVersionGetParamsDeviceTypeDesktop RadarHTTPTopAseTlsVersionGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseTlsVersionGetParamsDeviceTypeMobile  RadarHTTPTopAseTlsVersionGetParamsDeviceType = "MOBILE"
	RadarHTTPTopAseTlsVersionGetParamsDeviceTypeOther   RadarHTTPTopAseTlsVersionGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopAseTlsVersionGetParamsFormat string

const (
	RadarHTTPTopAseTlsVersionGetParamsFormatJson RadarHTTPTopAseTlsVersionGetParamsFormat = "JSON"
	RadarHTTPTopAseTlsVersionGetParamsFormatCsv  RadarHTTPTopAseTlsVersionGetParamsFormat = "CSV"
)

type RadarHTTPTopAseTlsVersionGetParamsHTTPProtocol string

const (
	RadarHTTPTopAseTlsVersionGetParamsHTTPProtocolHTTP  RadarHTTPTopAseTlsVersionGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseTlsVersionGetParamsHTTPProtocolHTTPs RadarHTTPTopAseTlsVersionGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopAseTlsVersionGetParamsHTTPVersion string

const (
	RadarHTTPTopAseTlsVersionGetParamsHTTPVersionHttPv1 RadarHTTPTopAseTlsVersionGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseTlsVersionGetParamsHTTPVersionHttPv2 RadarHTTPTopAseTlsVersionGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseTlsVersionGetParamsHTTPVersionHttPv3 RadarHTTPTopAseTlsVersionGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopAseTlsVersionGetParamsIPVersion string

const (
	RadarHTTPTopAseTlsVersionGetParamsIPVersionIPv4 RadarHTTPTopAseTlsVersionGetParamsIPVersion = "IPv4"
	RadarHTTPTopAseTlsVersionGetParamsIPVersionIPv6 RadarHTTPTopAseTlsVersionGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopAseTlsVersionGetParamsO string

const (
	RadarHTTPTopAseTlsVersionGetParamsOWindows  RadarHTTPTopAseTlsVersionGetParamsO = "WINDOWS"
	RadarHTTPTopAseTlsVersionGetParamsOMacosx   RadarHTTPTopAseTlsVersionGetParamsO = "MACOSX"
	RadarHTTPTopAseTlsVersionGetParamsOIos      RadarHTTPTopAseTlsVersionGetParamsO = "IOS"
	RadarHTTPTopAseTlsVersionGetParamsOAndroid  RadarHTTPTopAseTlsVersionGetParamsO = "ANDROID"
	RadarHTTPTopAseTlsVersionGetParamsOChromeos RadarHTTPTopAseTlsVersionGetParamsO = "CHROMEOS"
	RadarHTTPTopAseTlsVersionGetParamsOLinux    RadarHTTPTopAseTlsVersionGetParamsO = "LINUX"
	RadarHTTPTopAseTlsVersionGetParamsOSmartTv  RadarHTTPTopAseTlsVersionGetParamsO = "SMART_TV"
)
