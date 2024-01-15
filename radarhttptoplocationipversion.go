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

// RadarHTTPTopLocationIPVersionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTopLocationIPVersionService] method instead.
type RadarHTTPTopLocationIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopLocationIPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTopLocationIPVersionService(opts ...option.RequestOption) (r *RadarHTTPTopLocationIPVersionService) {
	r = &RadarHTTPTopLocationIPVersionService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested IP protocol version.
// Values are a percentage out of the total traffic.
func (r *RadarHTTPTopLocationIPVersionService) Get(ctx context.Context, ipVersion RadarHTTPTopLocationIPVersionGetParamsIPVersion, query RadarHTTPTopLocationIPVersionGetParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationIPVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopLocationIPVersionGetResponse struct {
	Result  RadarHTTPTopLocationIPVersionGetResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarHTTPTopLocationIPVersionGetResponseJSON   `json:"-"`
}

// radarHTTPTopLocationIPVersionGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopLocationIPVersionGetResponse]
type radarHTTPTopLocationIPVersionGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationIPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationIPVersionGetResponseResult struct {
	Meta RadarHTTPTopLocationIPVersionGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationIPVersionGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationIPVersionGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationIPVersionGetResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationIPVersionGetResponseResult]
type radarHTTPTopLocationIPVersionGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationIPVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationIPVersionGetResponseResultMeta struct {
	DateRange      []RadarHTTPTopLocationIPVersionGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                           `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopLocationIPVersionGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopLocationIPVersionGetResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationIPVersionGetResponseResultMeta]
type radarHTTPTopLocationIPVersionGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationIPVersionGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationIPVersionGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                       `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationIPVersionGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationIPVersionGetResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationIPVersionGetResponseResultMetaDateRange]
type radarHTTPTopLocationIPVersionGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationIPVersionGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                        `json:"level"`
	JSON        radarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                         `json:"dataSource,required"`
	Description     string                                                                         `json:"description,required"`
	EventType       string                                                                         `json:"eventType,required"`
	IsInstantaneous interface{}                                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                      `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                         `json:"linkedUrl"`
	StartTime       time.Time                                                                      `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopLocationIPVersionGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationIPVersionGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                 `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                 `json:"clientCountryName,required"`
	Value               string                                                 `json:"value,required"`
	JSON                radarHTTPTopLocationIPVersionGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationIPVersionGetResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationIPVersionGetResponseResultTop0]
type radarHTTPTopLocationIPVersionGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationIPVersionGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationIPVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationIPVersionGetParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopLocationIPVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopLocationIPVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopLocationIPVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopLocationIPVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationIPVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopLocationIPVersionGetParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopLocationIPVersionGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationIPVersionGetParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTopLocationIPVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// IP version.
type RadarHTTPTopLocationIPVersionGetParamsIPVersion string

const (
	RadarHTTPTopLocationIPVersionGetParamsIPVersionIPv4 RadarHTTPTopLocationIPVersionGetParamsIPVersion = "IPv4"
	RadarHTTPTopLocationIPVersionGetParamsIPVersionIPv6 RadarHTTPTopLocationIPVersionGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopLocationIPVersionGetParamsBotClass string

const (
	RadarHTTPTopLocationIPVersionGetParamsBotClassLikelyAutomated RadarHTTPTopLocationIPVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationIPVersionGetParamsBotClassLikelyHuman     RadarHTTPTopLocationIPVersionGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopLocationIPVersionGetParamsDateRange string

const (
	RadarHTTPTopLocationIPVersionGetParamsDateRange1d         RadarHTTPTopLocationIPVersionGetParamsDateRange = "1d"
	RadarHTTPTopLocationIPVersionGetParamsDateRange2d         RadarHTTPTopLocationIPVersionGetParamsDateRange = "2d"
	RadarHTTPTopLocationIPVersionGetParamsDateRange7d         RadarHTTPTopLocationIPVersionGetParamsDateRange = "7d"
	RadarHTTPTopLocationIPVersionGetParamsDateRange14d        RadarHTTPTopLocationIPVersionGetParamsDateRange = "14d"
	RadarHTTPTopLocationIPVersionGetParamsDateRange28d        RadarHTTPTopLocationIPVersionGetParamsDateRange = "28d"
	RadarHTTPTopLocationIPVersionGetParamsDateRange12w        RadarHTTPTopLocationIPVersionGetParamsDateRange = "12w"
	RadarHTTPTopLocationIPVersionGetParamsDateRange24w        RadarHTTPTopLocationIPVersionGetParamsDateRange = "24w"
	RadarHTTPTopLocationIPVersionGetParamsDateRange52w        RadarHTTPTopLocationIPVersionGetParamsDateRange = "52w"
	RadarHTTPTopLocationIPVersionGetParamsDateRange1dControl  RadarHTTPTopLocationIPVersionGetParamsDateRange = "1dControl"
	RadarHTTPTopLocationIPVersionGetParamsDateRange2dControl  RadarHTTPTopLocationIPVersionGetParamsDateRange = "2dControl"
	RadarHTTPTopLocationIPVersionGetParamsDateRange7dControl  RadarHTTPTopLocationIPVersionGetParamsDateRange = "7dControl"
	RadarHTTPTopLocationIPVersionGetParamsDateRange14dControl RadarHTTPTopLocationIPVersionGetParamsDateRange = "14dControl"
	RadarHTTPTopLocationIPVersionGetParamsDateRange28dControl RadarHTTPTopLocationIPVersionGetParamsDateRange = "28dControl"
	RadarHTTPTopLocationIPVersionGetParamsDateRange12wControl RadarHTTPTopLocationIPVersionGetParamsDateRange = "12wControl"
	RadarHTTPTopLocationIPVersionGetParamsDateRange24wControl RadarHTTPTopLocationIPVersionGetParamsDateRange = "24wControl"
)

type RadarHTTPTopLocationIPVersionGetParamsDeviceType string

const (
	RadarHTTPTopLocationIPVersionGetParamsDeviceTypeDesktop RadarHTTPTopLocationIPVersionGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationIPVersionGetParamsDeviceTypeMobile  RadarHTTPTopLocationIPVersionGetParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationIPVersionGetParamsDeviceTypeOther   RadarHTTPTopLocationIPVersionGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopLocationIPVersionGetParamsFormat string

const (
	RadarHTTPTopLocationIPVersionGetParamsFormatJson RadarHTTPTopLocationIPVersionGetParamsFormat = "JSON"
	RadarHTTPTopLocationIPVersionGetParamsFormatCsv  RadarHTTPTopLocationIPVersionGetParamsFormat = "CSV"
)

type RadarHTTPTopLocationIPVersionGetParamsHTTPProtocol string

const (
	RadarHTTPTopLocationIPVersionGetParamsHTTPProtocolHTTP  RadarHTTPTopLocationIPVersionGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationIPVersionGetParamsHTTPProtocolHTTPs RadarHTTPTopLocationIPVersionGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopLocationIPVersionGetParamsHTTPVersion string

const (
	RadarHTTPTopLocationIPVersionGetParamsHTTPVersionHttPv1 RadarHTTPTopLocationIPVersionGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationIPVersionGetParamsHTTPVersionHttPv2 RadarHTTPTopLocationIPVersionGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationIPVersionGetParamsHTTPVersionHttPv3 RadarHTTPTopLocationIPVersionGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopLocationIPVersionGetParamsO string

const (
	RadarHTTPTopLocationIPVersionGetParamsOWindows  RadarHTTPTopLocationIPVersionGetParamsO = "WINDOWS"
	RadarHTTPTopLocationIPVersionGetParamsOMacosx   RadarHTTPTopLocationIPVersionGetParamsO = "MACOSX"
	RadarHTTPTopLocationIPVersionGetParamsOIos      RadarHTTPTopLocationIPVersionGetParamsO = "IOS"
	RadarHTTPTopLocationIPVersionGetParamsOAndroid  RadarHTTPTopLocationIPVersionGetParamsO = "ANDROID"
	RadarHTTPTopLocationIPVersionGetParamsOChromeos RadarHTTPTopLocationIPVersionGetParamsO = "CHROMEOS"
	RadarHTTPTopLocationIPVersionGetParamsOLinux    RadarHTTPTopLocationIPVersionGetParamsO = "LINUX"
	RadarHTTPTopLocationIPVersionGetParamsOSmartTv  RadarHTTPTopLocationIPVersionGetParamsO = "SMART_TV"
)

type RadarHTTPTopLocationIPVersionGetParamsTlsVersion string

const (
	RadarHTTPTopLocationIPVersionGetParamsTlsVersionTlSv1_0  RadarHTTPTopLocationIPVersionGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationIPVersionGetParamsTlsVersionTlSv1_1  RadarHTTPTopLocationIPVersionGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationIPVersionGetParamsTlsVersionTlSv1_2  RadarHTTPTopLocationIPVersionGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationIPVersionGetParamsTlsVersionTlSv1_3  RadarHTTPTopLocationIPVersionGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationIPVersionGetParamsTlsVersionTlSvQuic RadarHTTPTopLocationIPVersionGetParamsTlsVersion = "TLSvQUIC"
)
