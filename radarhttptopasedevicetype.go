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

// RadarHTTPTopAseDeviceTypeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTopAseDeviceTypeService] method instead.
type RadarHTTPTopAseDeviceTypeService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopAseDeviceTypeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTopAseDeviceTypeService(opts ...option.RequestOption) (r *RadarHTTPTopAseDeviceTypeService) {
	r = &RadarHTTPTopAseDeviceTypeService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested device
// type. Values are a percentage out of the total traffic.
func (r *RadarHTTPTopAseDeviceTypeService) Get(ctx context.Context, deviceType RadarHTTPTopAseDeviceTypeGetParamsDeviceType, query RadarHTTPTopAseDeviceTypeGetParams, opts ...option.RequestOption) (res *RadarHTTPTopAseDeviceTypeGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/device_type/%v", deviceType)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopAseDeviceTypeGetResponse struct {
	Result  RadarHTTPTopAseDeviceTypeGetResponseResult `json:"result,required"`
	Success bool                                       `json:"success,required"`
	JSON    radarHTTPTopAseDeviceTypeGetResponseJSON   `json:"-"`
}

// radarHTTPTopAseDeviceTypeGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseDeviceTypeGetResponse]
type radarHTTPTopAseDeviceTypeGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseDeviceTypeGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseDeviceTypeGetResponseResult struct {
	Meta RadarHTTPTopAseDeviceTypeGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseDeviceTypeGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseDeviceTypeGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseDeviceTypeGetResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTopAseDeviceTypeGetResponseResult]
type radarHTTPTopAseDeviceTypeGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseDeviceTypeGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseDeviceTypeGetResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarHTTPTopAseDeviceTypeGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarHTTPTopAseDeviceTypeGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseDeviceTypeGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseDeviceTypeGetResponseResultMeta]
type radarHTTPTopAseDeviceTypeGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseDeviceTypeGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                    `json:"level,required"`
	JSON        radarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfo]
type radarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                     `json:"dataSource,required"`
	Description string                                                                     `json:"description,required"`
	EndTime     time.Time                                                                  `json:"endTime,required" format:"date-time"`
	EventType   string                                                                     `json:"eventType,required"`
	StartTime   time.Time                                                                  `json:"startTime,required" format:"date-time"`
	JSON        radarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseDeviceTypeGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseDeviceTypeGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                   `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                   `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseDeviceTypeGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseDeviceTypeGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopAseDeviceTypeGetResponseResultMetaDateRange]
type radarHTTPTopAseDeviceTypeGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseDeviceTypeGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseDeviceTypeGetResponseResultTop0 struct {
	ClientASN    int64                                              `json:"clientASN,required"`
	ClientAsName string                                             `json:"clientASName,required"`
	Value        string                                             `json:"value,required"`
	JSON         radarHTTPTopAseDeviceTypeGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseDeviceTypeGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarHTTPTopAseDeviceTypeGetResponseResultTop0]
type radarHTTPTopAseDeviceTypeGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseDeviceTypeGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseDeviceTypeGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseDeviceTypeGetParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopAseDeviceTypeGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopAseDeviceTypeGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopAseDeviceTypeGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopAseDeviceTypeGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopAseDeviceTypeGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopAseDeviceTypeGetParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopAseDeviceTypeGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseDeviceTypeGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopAseDeviceTypeGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopAseDeviceTypeGetParamsDeviceType string

const (
	RadarHTTPTopAseDeviceTypeGetParamsDeviceTypeDesktop RadarHTTPTopAseDeviceTypeGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseDeviceTypeGetParamsDeviceTypeMobile  RadarHTTPTopAseDeviceTypeGetParamsDeviceType = "MOBILE"
	RadarHTTPTopAseDeviceTypeGetParamsDeviceTypeOther   RadarHTTPTopAseDeviceTypeGetParamsDeviceType = "OTHER"
)

type RadarHTTPTopAseDeviceTypeGetParamsBotClass string

const (
	RadarHTTPTopAseDeviceTypeGetParamsBotClassLikelyAutomated RadarHTTPTopAseDeviceTypeGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseDeviceTypeGetParamsBotClassLikelyHuman     RadarHTTPTopAseDeviceTypeGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopAseDeviceTypeGetParamsDateRange string

const (
	RadarHTTPTopAseDeviceTypeGetParamsDateRange1d         RadarHTTPTopAseDeviceTypeGetParamsDateRange = "1d"
	RadarHTTPTopAseDeviceTypeGetParamsDateRange7d         RadarHTTPTopAseDeviceTypeGetParamsDateRange = "7d"
	RadarHTTPTopAseDeviceTypeGetParamsDateRange14d        RadarHTTPTopAseDeviceTypeGetParamsDateRange = "14d"
	RadarHTTPTopAseDeviceTypeGetParamsDateRange28d        RadarHTTPTopAseDeviceTypeGetParamsDateRange = "28d"
	RadarHTTPTopAseDeviceTypeGetParamsDateRange12w        RadarHTTPTopAseDeviceTypeGetParamsDateRange = "12w"
	RadarHTTPTopAseDeviceTypeGetParamsDateRange24w        RadarHTTPTopAseDeviceTypeGetParamsDateRange = "24w"
	RadarHTTPTopAseDeviceTypeGetParamsDateRange52w        RadarHTTPTopAseDeviceTypeGetParamsDateRange = "52w"
	RadarHTTPTopAseDeviceTypeGetParamsDateRange1dControl  RadarHTTPTopAseDeviceTypeGetParamsDateRange = "1dControl"
	RadarHTTPTopAseDeviceTypeGetParamsDateRange7dControl  RadarHTTPTopAseDeviceTypeGetParamsDateRange = "7dControl"
	RadarHTTPTopAseDeviceTypeGetParamsDateRange14dControl RadarHTTPTopAseDeviceTypeGetParamsDateRange = "14dControl"
	RadarHTTPTopAseDeviceTypeGetParamsDateRange28dControl RadarHTTPTopAseDeviceTypeGetParamsDateRange = "28dControl"
	RadarHTTPTopAseDeviceTypeGetParamsDateRange12wControl RadarHTTPTopAseDeviceTypeGetParamsDateRange = "12wControl"
	RadarHTTPTopAseDeviceTypeGetParamsDateRange24wControl RadarHTTPTopAseDeviceTypeGetParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarHTTPTopAseDeviceTypeGetParamsFormat string

const (
	RadarHTTPTopAseDeviceTypeGetParamsFormatJson RadarHTTPTopAseDeviceTypeGetParamsFormat = "JSON"
	RadarHTTPTopAseDeviceTypeGetParamsFormatCsv  RadarHTTPTopAseDeviceTypeGetParamsFormat = "CSV"
)

type RadarHTTPTopAseDeviceTypeGetParamsHTTPProtocol string

const (
	RadarHTTPTopAseDeviceTypeGetParamsHTTPProtocolHTTP  RadarHTTPTopAseDeviceTypeGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseDeviceTypeGetParamsHTTPProtocolHTTPs RadarHTTPTopAseDeviceTypeGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopAseDeviceTypeGetParamsHTTPVersion string

const (
	RadarHTTPTopAseDeviceTypeGetParamsHTTPVersionHttPv1 RadarHTTPTopAseDeviceTypeGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseDeviceTypeGetParamsHTTPVersionHttPv2 RadarHTTPTopAseDeviceTypeGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseDeviceTypeGetParamsHTTPVersionHttPv3 RadarHTTPTopAseDeviceTypeGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopAseDeviceTypeGetParamsIPVersion string

const (
	RadarHTTPTopAseDeviceTypeGetParamsIPVersionIPv4 RadarHTTPTopAseDeviceTypeGetParamsIPVersion = "IPv4"
	RadarHTTPTopAseDeviceTypeGetParamsIPVersionIPv6 RadarHTTPTopAseDeviceTypeGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopAseDeviceTypeGetParamsO string

const (
	RadarHTTPTopAseDeviceTypeGetParamsOWindows  RadarHTTPTopAseDeviceTypeGetParamsO = "WINDOWS"
	RadarHTTPTopAseDeviceTypeGetParamsOMacosx   RadarHTTPTopAseDeviceTypeGetParamsO = "MACOSX"
	RadarHTTPTopAseDeviceTypeGetParamsOAndroid  RadarHTTPTopAseDeviceTypeGetParamsO = "ANDROID"
	RadarHTTPTopAseDeviceTypeGetParamsOChromeos RadarHTTPTopAseDeviceTypeGetParamsO = "CHROMEOS"
	RadarHTTPTopAseDeviceTypeGetParamsOLinux    RadarHTTPTopAseDeviceTypeGetParamsO = "LINUX"
	RadarHTTPTopAseDeviceTypeGetParamsOSmartTv  RadarHTTPTopAseDeviceTypeGetParamsO = "SMART_TV"
)

type RadarHTTPTopAseDeviceTypeGetParamsTlsVersion string

const (
	RadarHTTPTopAseDeviceTypeGetParamsTlsVersionTlSv1_0  RadarHTTPTopAseDeviceTypeGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseDeviceTypeGetParamsTlsVersionTlSv1_1  RadarHTTPTopAseDeviceTypeGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseDeviceTypeGetParamsTlsVersionTlSv1_2  RadarHTTPTopAseDeviceTypeGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseDeviceTypeGetParamsTlsVersionTlSv1_3  RadarHTTPTopAseDeviceTypeGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseDeviceTypeGetParamsTlsVersionTlSvQuic RadarHTTPTopAseDeviceTypeGetParamsTlsVersion = "TLSvQUIC"
)
