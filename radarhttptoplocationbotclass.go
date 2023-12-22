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

// RadarHTTPTopLocationBotClassService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTopLocationBotClassService] method instead.
type RadarHTTPTopLocationBotClassService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopLocationBotClassService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTopLocationBotClassService(opts ...option.RequestOption) (r *RadarHTTPTopLocationBotClassService) {
	r = &RadarHTTPTopLocationBotClassService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested bot class. These two
// categories use Cloudflare's bot score - refer to [Bot
// scores])https://developers.cloudflare.com/bots/concepts/bot-score). Values are a
// percentage out of the total traffic.
func (r *RadarHTTPTopLocationBotClassService) Get(ctx context.Context, botClass RadarHTTPTopLocationBotClassGetParamsBotClass, query RadarHTTPTopLocationBotClassGetParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationBotClassGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/bot_class/%v", botClass)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopLocationBotClassGetResponse struct {
	Result  RadarHTTPTopLocationBotClassGetResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarHTTPTopLocationBotClassGetResponseJSON   `json:"-"`
}

// radarHTTPTopLocationBotClassGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopLocationBotClassGetResponse]
type radarHTTPTopLocationBotClassGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationBotClassGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationBotClassGetResponseResult struct {
	Meta RadarHTTPTopLocationBotClassGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationBotClassGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationBotClassGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationBotClassGetResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationBotClassGetResponseResult]
type radarHTTPTopLocationBotClassGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationBotClassGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationBotClassGetResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarHTTPTopLocationBotClassGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarHTTPTopLocationBotClassGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopLocationBotClassGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationBotClassGetResponseResultMeta]
type radarHTTPTopLocationBotClassGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationBotClassGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                       `json:"level,required"`
	JSON        radarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                        `json:"dataSource,required"`
	Description string                                                                        `json:"description,required"`
	EndTime     time.Time                                                                     `json:"endTime,required" format:"date-time"`
	EventType   string                                                                        `json:"eventType,required"`
	StartTime   time.Time                                                                     `json:"startTime,required" format:"date-time"`
	JSON        radarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationBotClassGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationBotClassGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                      `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                      `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationBotClassGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationBotClassGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationBotClassGetResponseResultMetaDateRange]
type radarHTTPTopLocationBotClassGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationBotClassGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationBotClassGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                `json:"clientCountryName,required"`
	Value               string                                                `json:"value,required"`
	JSON                radarHTTPTopLocationBotClassGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationBotClassGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationBotClassGetResponseResultTop0]
type radarHTTPTopLocationBotClassGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationBotClassGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationBotClassGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopLocationBotClassGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopLocationBotClassGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopLocationBotClassGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopLocationBotClassGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationBotClassGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopLocationBotClassGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopLocationBotClassGetParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopLocationBotClassGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationBotClassGetParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTopLocationBotClassGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopLocationBotClassGetParamsBotClass string

const (
	RadarHTTPTopLocationBotClassGetParamsBotClassLikelyAutomated RadarHTTPTopLocationBotClassGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationBotClassGetParamsBotClassLikelyHuman     RadarHTTPTopLocationBotClassGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopLocationBotClassGetParamsDateRange string

const (
	RadarHTTPTopLocationBotClassGetParamsDateRange1d         RadarHTTPTopLocationBotClassGetParamsDateRange = "1d"
	RadarHTTPTopLocationBotClassGetParamsDateRange7d         RadarHTTPTopLocationBotClassGetParamsDateRange = "7d"
	RadarHTTPTopLocationBotClassGetParamsDateRange14d        RadarHTTPTopLocationBotClassGetParamsDateRange = "14d"
	RadarHTTPTopLocationBotClassGetParamsDateRange28d        RadarHTTPTopLocationBotClassGetParamsDateRange = "28d"
	RadarHTTPTopLocationBotClassGetParamsDateRange12w        RadarHTTPTopLocationBotClassGetParamsDateRange = "12w"
	RadarHTTPTopLocationBotClassGetParamsDateRange24w        RadarHTTPTopLocationBotClassGetParamsDateRange = "24w"
	RadarHTTPTopLocationBotClassGetParamsDateRange52w        RadarHTTPTopLocationBotClassGetParamsDateRange = "52w"
	RadarHTTPTopLocationBotClassGetParamsDateRange1dControl  RadarHTTPTopLocationBotClassGetParamsDateRange = "1dControl"
	RadarHTTPTopLocationBotClassGetParamsDateRange7dControl  RadarHTTPTopLocationBotClassGetParamsDateRange = "7dControl"
	RadarHTTPTopLocationBotClassGetParamsDateRange14dControl RadarHTTPTopLocationBotClassGetParamsDateRange = "14dControl"
	RadarHTTPTopLocationBotClassGetParamsDateRange28dControl RadarHTTPTopLocationBotClassGetParamsDateRange = "28dControl"
	RadarHTTPTopLocationBotClassGetParamsDateRange12wControl RadarHTTPTopLocationBotClassGetParamsDateRange = "12wControl"
	RadarHTTPTopLocationBotClassGetParamsDateRange24wControl RadarHTTPTopLocationBotClassGetParamsDateRange = "24wControl"
)

type RadarHTTPTopLocationBotClassGetParamsDeviceType string

const (
	RadarHTTPTopLocationBotClassGetParamsDeviceTypeDesktop RadarHTTPTopLocationBotClassGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationBotClassGetParamsDeviceTypeMobile  RadarHTTPTopLocationBotClassGetParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationBotClassGetParamsDeviceTypeOther   RadarHTTPTopLocationBotClassGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopLocationBotClassGetParamsFormat string

const (
	RadarHTTPTopLocationBotClassGetParamsFormatJson RadarHTTPTopLocationBotClassGetParamsFormat = "JSON"
	RadarHTTPTopLocationBotClassGetParamsFormatCsv  RadarHTTPTopLocationBotClassGetParamsFormat = "CSV"
)

type RadarHTTPTopLocationBotClassGetParamsHTTPProtocol string

const (
	RadarHTTPTopLocationBotClassGetParamsHTTPProtocolHTTP  RadarHTTPTopLocationBotClassGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationBotClassGetParamsHTTPProtocolHTTPs RadarHTTPTopLocationBotClassGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopLocationBotClassGetParamsHTTPVersion string

const (
	RadarHTTPTopLocationBotClassGetParamsHTTPVersionHttPv1 RadarHTTPTopLocationBotClassGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationBotClassGetParamsHTTPVersionHttPv2 RadarHTTPTopLocationBotClassGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationBotClassGetParamsHTTPVersionHttPv3 RadarHTTPTopLocationBotClassGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopLocationBotClassGetParamsIPVersion string

const (
	RadarHTTPTopLocationBotClassGetParamsIPVersionIPv4 RadarHTTPTopLocationBotClassGetParamsIPVersion = "IPv4"
	RadarHTTPTopLocationBotClassGetParamsIPVersionIPv6 RadarHTTPTopLocationBotClassGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopLocationBotClassGetParamsO string

const (
	RadarHTTPTopLocationBotClassGetParamsOWindows  RadarHTTPTopLocationBotClassGetParamsO = "WINDOWS"
	RadarHTTPTopLocationBotClassGetParamsOMacosx   RadarHTTPTopLocationBotClassGetParamsO = "MACOSX"
	RadarHTTPTopLocationBotClassGetParamsOAndroid  RadarHTTPTopLocationBotClassGetParamsO = "ANDROID"
	RadarHTTPTopLocationBotClassGetParamsOChromeos RadarHTTPTopLocationBotClassGetParamsO = "CHROMEOS"
	RadarHTTPTopLocationBotClassGetParamsOLinux    RadarHTTPTopLocationBotClassGetParamsO = "LINUX"
	RadarHTTPTopLocationBotClassGetParamsOSmartTv  RadarHTTPTopLocationBotClassGetParamsO = "SMART_TV"
)

type RadarHTTPTopLocationBotClassGetParamsTlsVersion string

const (
	RadarHTTPTopLocationBotClassGetParamsTlsVersionTlSv1_0  RadarHTTPTopLocationBotClassGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationBotClassGetParamsTlsVersionTlSv1_1  RadarHTTPTopLocationBotClassGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationBotClassGetParamsTlsVersionTlSv1_2  RadarHTTPTopLocationBotClassGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationBotClassGetParamsTlsVersionTlSv1_3  RadarHTTPTopLocationBotClassGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationBotClassGetParamsTlsVersionTlSvQuic RadarHTTPTopLocationBotClassGetParamsTlsVersion = "TLSvQUIC"
)
