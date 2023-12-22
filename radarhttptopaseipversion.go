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

// RadarHTTPTopAseIPVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTopAseIPVersionService] method instead.
type RadarHTTPTopAseIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopAseIPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTopAseIPVersionService(opts ...option.RequestOption) (r *RadarHTTPTopAseIPVersionService) {
	r = &RadarHTTPTopAseIPVersionService{}
	r.Options = opts
	return
}

// Get the top autonomous systems, by HTTP traffic, of the requested IP protocol
// version. Values are a percentage out of the total traffic.
func (r *RadarHTTPTopAseIPVersionService) Get(ctx context.Context, ipVersion RadarHTTPTopAseIPVersionGetParamsIPVersion, query RadarHTTPTopAseIPVersionGetParams, opts ...option.RequestOption) (res *RadarHTTPTopAseIPVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopAseIPVersionGetResponse struct {
	Result  RadarHTTPTopAseIPVersionGetResponseResult `json:"result,required"`
	Success bool                                      `json:"success,required"`
	JSON    radarHTTPTopAseIPVersionGetResponseJSON   `json:"-"`
}

// radarHTTPTopAseIPVersionGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseIPVersionGetResponse]
type radarHTTPTopAseIPVersionGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseIPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseIPVersionGetResponseResult struct {
	Meta RadarHTTPTopAseIPVersionGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseIPVersionGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseIPVersionGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseIPVersionGetResponseResultJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseIPVersionGetResponseResult]
type radarHTTPTopAseIPVersionGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseIPVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseIPVersionGetResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarHTTPTopAseIPVersionGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarHTTPTopAseIPVersionGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseIPVersionGetResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarHTTPTopAseIPVersionGetResponseResultMeta]
type radarHTTPTopAseIPVersionGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseIPVersionGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                   `json:"level,required"`
	JSON        radarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfo]
type radarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                    `json:"dataSource,required"`
	Description string                                                                    `json:"description,required"`
	EndTime     time.Time                                                                 `json:"endTime,required" format:"date-time"`
	EventType   string                                                                    `json:"eventType,required"`
	StartTime   time.Time                                                                 `json:"startTime,required" format:"date-time"`
	JSON        radarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseIPVersionGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseIPVersionGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                  `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                  `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseIPVersionGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseIPVersionGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPTopAseIPVersionGetResponseResultMetaDateRange]
type radarHTTPTopAseIPVersionGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseIPVersionGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseIPVersionGetResponseResultTop0 struct {
	ClientASN    int64                                             `json:"clientASN,required"`
	ClientAsName string                                            `json:"clientASName,required"`
	Value        string                                            `json:"value,required"`
	JSON         radarHTTPTopAseIPVersionGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseIPVersionGetResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarHTTPTopAseIPVersionGetResponseResultTop0]
type radarHTTPTopAseIPVersionGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseIPVersionGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseIPVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseIPVersionGetParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopAseIPVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopAseIPVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopAseIPVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopAseIPVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopAseIPVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopAseIPVersionGetParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopAseIPVersionGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseIPVersionGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopAseIPVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopAseIPVersionGetParamsIPVersion string

const (
	RadarHTTPTopAseIPVersionGetParamsIPVersionIPv4 RadarHTTPTopAseIPVersionGetParamsIPVersion = "IPv4"
	RadarHTTPTopAseIPVersionGetParamsIPVersionIPv6 RadarHTTPTopAseIPVersionGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopAseIPVersionGetParamsBotClass string

const (
	RadarHTTPTopAseIPVersionGetParamsBotClassLikelyAutomated RadarHTTPTopAseIPVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseIPVersionGetParamsBotClassLikelyHuman     RadarHTTPTopAseIPVersionGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopAseIPVersionGetParamsDateRange string

const (
	RadarHTTPTopAseIPVersionGetParamsDateRange1d         RadarHTTPTopAseIPVersionGetParamsDateRange = "1d"
	RadarHTTPTopAseIPVersionGetParamsDateRange7d         RadarHTTPTopAseIPVersionGetParamsDateRange = "7d"
	RadarHTTPTopAseIPVersionGetParamsDateRange14d        RadarHTTPTopAseIPVersionGetParamsDateRange = "14d"
	RadarHTTPTopAseIPVersionGetParamsDateRange28d        RadarHTTPTopAseIPVersionGetParamsDateRange = "28d"
	RadarHTTPTopAseIPVersionGetParamsDateRange12w        RadarHTTPTopAseIPVersionGetParamsDateRange = "12w"
	RadarHTTPTopAseIPVersionGetParamsDateRange24w        RadarHTTPTopAseIPVersionGetParamsDateRange = "24w"
	RadarHTTPTopAseIPVersionGetParamsDateRange52w        RadarHTTPTopAseIPVersionGetParamsDateRange = "52w"
	RadarHTTPTopAseIPVersionGetParamsDateRange1dControl  RadarHTTPTopAseIPVersionGetParamsDateRange = "1dControl"
	RadarHTTPTopAseIPVersionGetParamsDateRange7dControl  RadarHTTPTopAseIPVersionGetParamsDateRange = "7dControl"
	RadarHTTPTopAseIPVersionGetParamsDateRange14dControl RadarHTTPTopAseIPVersionGetParamsDateRange = "14dControl"
	RadarHTTPTopAseIPVersionGetParamsDateRange28dControl RadarHTTPTopAseIPVersionGetParamsDateRange = "28dControl"
	RadarHTTPTopAseIPVersionGetParamsDateRange12wControl RadarHTTPTopAseIPVersionGetParamsDateRange = "12wControl"
	RadarHTTPTopAseIPVersionGetParamsDateRange24wControl RadarHTTPTopAseIPVersionGetParamsDateRange = "24wControl"
)

type RadarHTTPTopAseIPVersionGetParamsDeviceType string

const (
	RadarHTTPTopAseIPVersionGetParamsDeviceTypeDesktop RadarHTTPTopAseIPVersionGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseIPVersionGetParamsDeviceTypeMobile  RadarHTTPTopAseIPVersionGetParamsDeviceType = "MOBILE"
	RadarHTTPTopAseIPVersionGetParamsDeviceTypeOther   RadarHTTPTopAseIPVersionGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopAseIPVersionGetParamsFormat string

const (
	RadarHTTPTopAseIPVersionGetParamsFormatJson RadarHTTPTopAseIPVersionGetParamsFormat = "JSON"
	RadarHTTPTopAseIPVersionGetParamsFormatCsv  RadarHTTPTopAseIPVersionGetParamsFormat = "CSV"
)

type RadarHTTPTopAseIPVersionGetParamsHTTPProtocol string

const (
	RadarHTTPTopAseIPVersionGetParamsHTTPProtocolHTTP  RadarHTTPTopAseIPVersionGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseIPVersionGetParamsHTTPProtocolHTTPs RadarHTTPTopAseIPVersionGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopAseIPVersionGetParamsHTTPVersion string

const (
	RadarHTTPTopAseIPVersionGetParamsHTTPVersionHttPv1 RadarHTTPTopAseIPVersionGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseIPVersionGetParamsHTTPVersionHttPv2 RadarHTTPTopAseIPVersionGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseIPVersionGetParamsHTTPVersionHttPv3 RadarHTTPTopAseIPVersionGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopAseIPVersionGetParamsO string

const (
	RadarHTTPTopAseIPVersionGetParamsOWindows  RadarHTTPTopAseIPVersionGetParamsO = "WINDOWS"
	RadarHTTPTopAseIPVersionGetParamsOMacosx   RadarHTTPTopAseIPVersionGetParamsO = "MACOSX"
	RadarHTTPTopAseIPVersionGetParamsOAndroid  RadarHTTPTopAseIPVersionGetParamsO = "ANDROID"
	RadarHTTPTopAseIPVersionGetParamsOChromeos RadarHTTPTopAseIPVersionGetParamsO = "CHROMEOS"
	RadarHTTPTopAseIPVersionGetParamsOLinux    RadarHTTPTopAseIPVersionGetParamsO = "LINUX"
	RadarHTTPTopAseIPVersionGetParamsOSmartTv  RadarHTTPTopAseIPVersionGetParamsO = "SMART_TV"
)

type RadarHTTPTopAseIPVersionGetParamsTlsVersion string

const (
	RadarHTTPTopAseIPVersionGetParamsTlsVersionTlSv1_0  RadarHTTPTopAseIPVersionGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseIPVersionGetParamsTlsVersionTlSv1_1  RadarHTTPTopAseIPVersionGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseIPVersionGetParamsTlsVersionTlSv1_2  RadarHTTPTopAseIPVersionGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseIPVersionGetParamsTlsVersionTlSv1_3  RadarHTTPTopAseIPVersionGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseIPVersionGetParamsTlsVersionTlSvQuic RadarHTTPTopAseIPVersionGetParamsTlsVersion = "TLSvQUIC"
)
