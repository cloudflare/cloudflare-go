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

// RadarHTTPTopLocationTlsVersionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTopLocationTlsVersionService] method instead.
type RadarHTTPTopLocationTlsVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopLocationTlsVersionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTopLocationTlsVersionService(opts ...option.RequestOption) (r *RadarHTTPTopLocationTlsVersionService) {
	r = &RadarHTTPTopLocationTlsVersionService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested TLS protocol version.
// Values are a percentage out of the total traffic.
func (r *RadarHTTPTopLocationTlsVersionService) Get(ctx context.Context, tlsVersion RadarHTTPTopLocationTlsVersionGetParamsTlsVersion, query RadarHTTPTopLocationTlsVersionGetParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationTlsVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/tls_version/%v", tlsVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopLocationTlsVersionGetResponse struct {
	Result  RadarHTTPTopLocationTlsVersionGetResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarHTTPTopLocationTlsVersionGetResponseJSON   `json:"-"`
}

// radarHTTPTopLocationTlsVersionGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopLocationTlsVersionGetResponse]
type radarHTTPTopLocationTlsVersionGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationTlsVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationTlsVersionGetResponseResult struct {
	Meta RadarHTTPTopLocationTlsVersionGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationTlsVersionGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationTlsVersionGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationTlsVersionGetResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationTlsVersionGetResponseResult]
type radarHTTPTopLocationTlsVersionGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationTlsVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationTlsVersionGetResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarHTTPTopLocationTlsVersionGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarHTTPTopLocationTlsVersionGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopLocationTlsVersionGetResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationTlsVersionGetResponseResultMeta]
type radarHTTPTopLocationTlsVersionGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationTlsVersionGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                         `json:"level,required"`
	JSON        radarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                          `json:"dataSource,required"`
	Description string                                                                          `json:"description,required"`
	EndTime     time.Time                                                                       `json:"endTime,required" format:"date-time"`
	EventType   string                                                                          `json:"eventType,required"`
	StartTime   time.Time                                                                       `json:"startTime,required" format:"date-time"`
	JSON        radarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationTlsVersionGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationTlsVersionGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                        `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationTlsVersionGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationTlsVersionGetResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationTlsVersionGetResponseResultMetaDateRange]
type radarHTTPTopLocationTlsVersionGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationTlsVersionGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationTlsVersionGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                  `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                  `json:"clientCountryName,required"`
	Value               string                                                  `json:"value,required"`
	JSON                radarHTTPTopLocationTlsVersionGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationTlsVersionGetResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationTlsVersionGetResponseResultTop0]
type radarHTTPTopLocationTlsVersionGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationTlsVersionGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationTlsVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationTlsVersionGetParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopLocationTlsVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopLocationTlsVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopLocationTlsVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopLocationTlsVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationTlsVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopLocationTlsVersionGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopLocationTlsVersionGetParamsO] `query:"os"`
}

// URLQuery serializes [RadarHTTPTopLocationTlsVersionGetParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTopLocationTlsVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopLocationTlsVersionGetParamsTlsVersion string

const (
	RadarHTTPTopLocationTlsVersionGetParamsTlsVersionTlSv1_0  RadarHTTPTopLocationTlsVersionGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationTlsVersionGetParamsTlsVersionTlSv1_1  RadarHTTPTopLocationTlsVersionGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationTlsVersionGetParamsTlsVersionTlSv1_2  RadarHTTPTopLocationTlsVersionGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationTlsVersionGetParamsTlsVersionTlSv1_3  RadarHTTPTopLocationTlsVersionGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationTlsVersionGetParamsTlsVersionTlSvQuic RadarHTTPTopLocationTlsVersionGetParamsTlsVersion = "TLSvQUIC"
)

type RadarHTTPTopLocationTlsVersionGetParamsBotClass string

const (
	RadarHTTPTopLocationTlsVersionGetParamsBotClassLikelyAutomated RadarHTTPTopLocationTlsVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationTlsVersionGetParamsBotClassLikelyHuman     RadarHTTPTopLocationTlsVersionGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopLocationTlsVersionGetParamsDateRange string

const (
	RadarHTTPTopLocationTlsVersionGetParamsDateRange1d         RadarHTTPTopLocationTlsVersionGetParamsDateRange = "1d"
	RadarHTTPTopLocationTlsVersionGetParamsDateRange7d         RadarHTTPTopLocationTlsVersionGetParamsDateRange = "7d"
	RadarHTTPTopLocationTlsVersionGetParamsDateRange14d        RadarHTTPTopLocationTlsVersionGetParamsDateRange = "14d"
	RadarHTTPTopLocationTlsVersionGetParamsDateRange28d        RadarHTTPTopLocationTlsVersionGetParamsDateRange = "28d"
	RadarHTTPTopLocationTlsVersionGetParamsDateRange12w        RadarHTTPTopLocationTlsVersionGetParamsDateRange = "12w"
	RadarHTTPTopLocationTlsVersionGetParamsDateRange24w        RadarHTTPTopLocationTlsVersionGetParamsDateRange = "24w"
	RadarHTTPTopLocationTlsVersionGetParamsDateRange52w        RadarHTTPTopLocationTlsVersionGetParamsDateRange = "52w"
	RadarHTTPTopLocationTlsVersionGetParamsDateRange1dControl  RadarHTTPTopLocationTlsVersionGetParamsDateRange = "1dControl"
	RadarHTTPTopLocationTlsVersionGetParamsDateRange7dControl  RadarHTTPTopLocationTlsVersionGetParamsDateRange = "7dControl"
	RadarHTTPTopLocationTlsVersionGetParamsDateRange14dControl RadarHTTPTopLocationTlsVersionGetParamsDateRange = "14dControl"
	RadarHTTPTopLocationTlsVersionGetParamsDateRange28dControl RadarHTTPTopLocationTlsVersionGetParamsDateRange = "28dControl"
	RadarHTTPTopLocationTlsVersionGetParamsDateRange12wControl RadarHTTPTopLocationTlsVersionGetParamsDateRange = "12wControl"
	RadarHTTPTopLocationTlsVersionGetParamsDateRange24wControl RadarHTTPTopLocationTlsVersionGetParamsDateRange = "24wControl"
)

type RadarHTTPTopLocationTlsVersionGetParamsDeviceType string

const (
	RadarHTTPTopLocationTlsVersionGetParamsDeviceTypeDesktop RadarHTTPTopLocationTlsVersionGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationTlsVersionGetParamsDeviceTypeMobile  RadarHTTPTopLocationTlsVersionGetParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationTlsVersionGetParamsDeviceTypeOther   RadarHTTPTopLocationTlsVersionGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopLocationTlsVersionGetParamsFormat string

const (
	RadarHTTPTopLocationTlsVersionGetParamsFormatJson RadarHTTPTopLocationTlsVersionGetParamsFormat = "JSON"
	RadarHTTPTopLocationTlsVersionGetParamsFormatCsv  RadarHTTPTopLocationTlsVersionGetParamsFormat = "CSV"
)

type RadarHTTPTopLocationTlsVersionGetParamsHTTPProtocol string

const (
	RadarHTTPTopLocationTlsVersionGetParamsHTTPProtocolHTTP  RadarHTTPTopLocationTlsVersionGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationTlsVersionGetParamsHTTPProtocolHTTPs RadarHTTPTopLocationTlsVersionGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopLocationTlsVersionGetParamsHTTPVersion string

const (
	RadarHTTPTopLocationTlsVersionGetParamsHTTPVersionHttPv1 RadarHTTPTopLocationTlsVersionGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationTlsVersionGetParamsHTTPVersionHttPv2 RadarHTTPTopLocationTlsVersionGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationTlsVersionGetParamsHTTPVersionHttPv3 RadarHTTPTopLocationTlsVersionGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopLocationTlsVersionGetParamsIPVersion string

const (
	RadarHTTPTopLocationTlsVersionGetParamsIPVersionIPv4 RadarHTTPTopLocationTlsVersionGetParamsIPVersion = "IPv4"
	RadarHTTPTopLocationTlsVersionGetParamsIPVersionIPv6 RadarHTTPTopLocationTlsVersionGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopLocationTlsVersionGetParamsO string

const (
	RadarHTTPTopLocationTlsVersionGetParamsOWindows  RadarHTTPTopLocationTlsVersionGetParamsO = "WINDOWS"
	RadarHTTPTopLocationTlsVersionGetParamsOMacosx   RadarHTTPTopLocationTlsVersionGetParamsO = "MACOSX"
	RadarHTTPTopLocationTlsVersionGetParamsOAndroid  RadarHTTPTopLocationTlsVersionGetParamsO = "ANDROID"
	RadarHTTPTopLocationTlsVersionGetParamsOChromeos RadarHTTPTopLocationTlsVersionGetParamsO = "CHROMEOS"
	RadarHTTPTopLocationTlsVersionGetParamsOLinux    RadarHTTPTopLocationTlsVersionGetParamsO = "LINUX"
	RadarHTTPTopLocationTlsVersionGetParamsOSmartTv  RadarHTTPTopLocationTlsVersionGetParamsO = "SMART_TV"
)
