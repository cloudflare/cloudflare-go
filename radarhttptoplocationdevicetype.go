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

// RadarHTTPTopLocationDeviceTypeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTopLocationDeviceTypeService] method instead.
type RadarHTTPTopLocationDeviceTypeService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopLocationDeviceTypeService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTopLocationDeviceTypeService(opts ...option.RequestOption) (r *RadarHTTPTopLocationDeviceTypeService) {
	r = &RadarHTTPTopLocationDeviceTypeService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested device type. Values are
// a percentage out of the total traffic.
func (r *RadarHTTPTopLocationDeviceTypeService) Get(ctx context.Context, deviceType RadarHTTPTopLocationDeviceTypeGetParamsDeviceType, query RadarHTTPTopLocationDeviceTypeGetParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationDeviceTypeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/device_type/%v", deviceType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopLocationDeviceTypeGetResponse struct {
	Result  RadarHTTPTopLocationDeviceTypeGetResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarHTTPTopLocationDeviceTypeGetResponseJSON   `json:"-"`
}

// radarHTTPTopLocationDeviceTypeGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopLocationDeviceTypeGetResponse]
type radarHTTPTopLocationDeviceTypeGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationDeviceTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationDeviceTypeGetResponseResult struct {
	Meta RadarHTTPTopLocationDeviceTypeGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationDeviceTypeGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationDeviceTypeGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationDeviceTypeGetResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationDeviceTypeGetResponseResult]
type radarHTTPTopLocationDeviceTypeGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationDeviceTypeGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationDeviceTypeGetResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarHTTPTopLocationDeviceTypeGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarHTTPTopLocationDeviceTypeGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopLocationDeviceTypeGetResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationDeviceTypeGetResponseResultMeta]
type radarHTTPTopLocationDeviceTypeGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationDeviceTypeGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                         `json:"level,required"`
	JSON        radarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                          `json:"dataSource,required"`
	Description string                                                                          `json:"description,required"`
	EndTime     time.Time                                                                       `json:"endTime,required" format:"date-time"`
	EventType   string                                                                          `json:"eventType,required"`
	StartTime   time.Time                                                                       `json:"startTime,required" format:"date-time"`
	JSON        radarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationDeviceTypeGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationDeviceTypeGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                        `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                        `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationDeviceTypeGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationDeviceTypeGetResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationDeviceTypeGetResponseResultMetaDateRange]
type radarHTTPTopLocationDeviceTypeGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationDeviceTypeGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationDeviceTypeGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                  `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                  `json:"clientCountryName,required"`
	Value               string                                                  `json:"value,required"`
	JSON                radarHTTPTopLocationDeviceTypeGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationDeviceTypeGetResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationDeviceTypeGetResponseResultTop0]
type radarHTTPTopLocationDeviceTypeGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationDeviceTypeGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationDeviceTypeGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationDeviceTypeGetParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopLocationDeviceTypeGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopLocationDeviceTypeGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopLocationDeviceTypeGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationDeviceTypeGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopLocationDeviceTypeGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopLocationDeviceTypeGetParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopLocationDeviceTypeGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationDeviceTypeGetParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTopLocationDeviceTypeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopLocationDeviceTypeGetParamsDeviceType string

const (
	RadarHTTPTopLocationDeviceTypeGetParamsDeviceTypeDesktop RadarHTTPTopLocationDeviceTypeGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationDeviceTypeGetParamsDeviceTypeMobile  RadarHTTPTopLocationDeviceTypeGetParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationDeviceTypeGetParamsDeviceTypeOther   RadarHTTPTopLocationDeviceTypeGetParamsDeviceType = "OTHER"
)

type RadarHTTPTopLocationDeviceTypeGetParamsBotClass string

const (
	RadarHTTPTopLocationDeviceTypeGetParamsBotClassLikelyAutomated RadarHTTPTopLocationDeviceTypeGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationDeviceTypeGetParamsBotClassLikelyHuman     RadarHTTPTopLocationDeviceTypeGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopLocationDeviceTypeGetParamsDateRange string

const (
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange1d         RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "1d"
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange7d         RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "7d"
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange14d        RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "14d"
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange28d        RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "28d"
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange12w        RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "12w"
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange24w        RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "24w"
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange52w        RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "52w"
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange1dControl  RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "1dControl"
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange7dControl  RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "7dControl"
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange14dControl RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "14dControl"
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange28dControl RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "28dControl"
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange12wControl RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "12wControl"
	RadarHTTPTopLocationDeviceTypeGetParamsDateRange24wControl RadarHTTPTopLocationDeviceTypeGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarHTTPTopLocationDeviceTypeGetParamsFormat string

const (
	RadarHTTPTopLocationDeviceTypeGetParamsFormatJson RadarHTTPTopLocationDeviceTypeGetParamsFormat = "JSON"
	RadarHTTPTopLocationDeviceTypeGetParamsFormatCsv  RadarHTTPTopLocationDeviceTypeGetParamsFormat = "CSV"
)

type RadarHTTPTopLocationDeviceTypeGetParamsHTTPProtocol string

const (
	RadarHTTPTopLocationDeviceTypeGetParamsHTTPProtocolHTTP  RadarHTTPTopLocationDeviceTypeGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationDeviceTypeGetParamsHTTPProtocolHTTPs RadarHTTPTopLocationDeviceTypeGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopLocationDeviceTypeGetParamsHTTPVersion string

const (
	RadarHTTPTopLocationDeviceTypeGetParamsHTTPVersionHttPv1 RadarHTTPTopLocationDeviceTypeGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationDeviceTypeGetParamsHTTPVersionHttPv2 RadarHTTPTopLocationDeviceTypeGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationDeviceTypeGetParamsHTTPVersionHttPv3 RadarHTTPTopLocationDeviceTypeGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopLocationDeviceTypeGetParamsIPVersion string

const (
	RadarHTTPTopLocationDeviceTypeGetParamsIPVersionIPv4 RadarHTTPTopLocationDeviceTypeGetParamsIPVersion = "IPv4"
	RadarHTTPTopLocationDeviceTypeGetParamsIPVersionIPv6 RadarHTTPTopLocationDeviceTypeGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopLocationDeviceTypeGetParamsO string

const (
	RadarHTTPTopLocationDeviceTypeGetParamsOWindows  RadarHTTPTopLocationDeviceTypeGetParamsO = "WINDOWS"
	RadarHTTPTopLocationDeviceTypeGetParamsOMacosx   RadarHTTPTopLocationDeviceTypeGetParamsO = "MACOSX"
	RadarHTTPTopLocationDeviceTypeGetParamsOAndroid  RadarHTTPTopLocationDeviceTypeGetParamsO = "ANDROID"
	RadarHTTPTopLocationDeviceTypeGetParamsOChromeos RadarHTTPTopLocationDeviceTypeGetParamsO = "CHROMEOS"
	RadarHTTPTopLocationDeviceTypeGetParamsOLinux    RadarHTTPTopLocationDeviceTypeGetParamsO = "LINUX"
	RadarHTTPTopLocationDeviceTypeGetParamsOSmartTv  RadarHTTPTopLocationDeviceTypeGetParamsO = "SMART_TV"
)

type RadarHTTPTopLocationDeviceTypeGetParamsTlsVersion string

const (
	RadarHTTPTopLocationDeviceTypeGetParamsTlsVersionTlSv1_0  RadarHTTPTopLocationDeviceTypeGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationDeviceTypeGetParamsTlsVersionTlSv1_1  RadarHTTPTopLocationDeviceTypeGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationDeviceTypeGetParamsTlsVersionTlSv1_2  RadarHTTPTopLocationDeviceTypeGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationDeviceTypeGetParamsTlsVersionTlSv1_3  RadarHTTPTopLocationDeviceTypeGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationDeviceTypeGetParamsTlsVersionTlSvQuic RadarHTTPTopLocationDeviceTypeGetParamsTlsVersion = "TLSvQUIC"
)
