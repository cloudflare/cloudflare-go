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

// RadarHTTPTopLocationHTTPProtocolService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTopLocationHTTPProtocolService] method instead.
type RadarHTTPTopLocationHTTPProtocolService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopLocationHTTPProtocolService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTopLocationHTTPProtocolService(opts ...option.RequestOption) (r *RadarHTTPTopLocationHTTPProtocolService) {
	r = &RadarHTTPTopLocationHTTPProtocolService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested HTTP protocol. Values
// are a percentage out of the total traffic.
func (r *RadarHTTPTopLocationHTTPProtocolService) Get(ctx context.Context, httpProtocol RadarHTTPTopLocationHTTPProtocolGetParamsHTTPProtocol, query RadarHTTPTopLocationHTTPProtocolGetParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationHTTPProtocolGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/http_protocol/%v", httpProtocol)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopLocationHTTPProtocolGetResponse struct {
	Result  RadarHTTPTopLocationHTTPProtocolGetResponseResult `json:"result,required"`
	Success bool                                              `json:"success,required"`
	JSON    radarHTTPTopLocationHTTPProtocolGetResponseJSON   `json:"-"`
}

// radarHTTPTopLocationHTTPProtocolGetResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationHTTPProtocolGetResponse]
type radarHTTPTopLocationHTTPProtocolGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPProtocolGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPProtocolGetResponseResult struct {
	Meta RadarHTTPTopLocationHTTPProtocolGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationHTTPProtocolGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationHTTPProtocolGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationHTTPProtocolGetResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTopLocationHTTPProtocolGetResponseResult]
type radarHTTPTopLocationHTTPProtocolGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPProtocolGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPProtocolGetResponseResultMeta struct {
	ConfidenceInfo RadarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfo `json:"confidenceInfo,required"`
	DateRange      RadarHTTPTopLocationHTTPProtocolGetResponseResultMetaDateRange      `json:"dateRange,required"`
	JSON           radarHTTPTopLocationHTTPProtocolGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopLocationHTTPProtocolGetResponseResultMetaJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationHTTPProtocolGetResponseResultMeta]
type radarHTTPTopLocationHTTPProtocolGetResponseResultMetaJSON struct {
	ConfidenceInfo apijson.Field
	DateRange      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPProtocolGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations,required"`
	Level       int64                                                                           `json:"level,required"`
	JSON        radarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfoJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource  string                                                                            `json:"dataSource,required"`
	Description string                                                                            `json:"description,required"`
	EndTime     time.Time                                                                         `json:"endTime,required" format:"date-time"`
	EventType   string                                                                            `json:"eventType,required"`
	StartTime   time.Time                                                                         `json:"startTime,required" format:"date-time"`
	JSON        radarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
	DataSource  apijson.Field
	Description apijson.Field
	EndTime     apijson.Field
	EventType   apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPProtocolGetResponseResultMetaDateRange struct {
	EndTime   time.Time                                                          `json:"endTime,required" format:"date-time"`
	StartTime time.Time                                                          `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationHTTPProtocolGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationHTTPProtocolGetResponseResultMetaDateRangeJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopLocationHTTPProtocolGetResponseResultMetaDateRange]
type radarHTTPTopLocationHTTPProtocolGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPProtocolGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPProtocolGetResponseResultTop0 struct {
	ClientCountryAlpha2 string                                                    `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                                    `json:"clientCountryName,required"`
	Value               string                                                    `json:"value,required"`
	JSON                radarHTTPTopLocationHTTPProtocolGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationHTTPProtocolGetResponseResultTop0JSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationHTTPProtocolGetResponseResultTop0]
type radarHTTPTopLocationHTTPProtocolGetResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationHTTPProtocolGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationHTTPProtocolGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationHTTPProtocolGetParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopLocationHTTPProtocolGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopLocationHTTPProtocolGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopLocationHTTPProtocolGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopLocationHTTPProtocolGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopLocationHTTPProtocolGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopLocationHTTPProtocolGetParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationHTTPProtocolGetParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTopLocationHTTPProtocolGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopLocationHTTPProtocolGetParamsHTTPProtocol string

const (
	RadarHTTPTopLocationHTTPProtocolGetParamsHTTPProtocolHTTP  RadarHTTPTopLocationHTTPProtocolGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationHTTPProtocolGetParamsHTTPProtocolHTTPs RadarHTTPTopLocationHTTPProtocolGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopLocationHTTPProtocolGetParamsBotClass string

const (
	RadarHTTPTopLocationHTTPProtocolGetParamsBotClassLikelyAutomated RadarHTTPTopLocationHTTPProtocolGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationHTTPProtocolGetParamsBotClassLikelyHuman     RadarHTTPTopLocationHTTPProtocolGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopLocationHTTPProtocolGetParamsDateRange string

const (
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange1d         RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "1d"
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange7d         RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "7d"
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange14d        RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "14d"
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange28d        RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "28d"
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange12w        RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "12w"
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange24w        RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "24w"
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange52w        RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "52w"
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange1dControl  RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "1dControl"
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange7dControl  RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "7dControl"
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange14dControl RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "14dControl"
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange28dControl RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "28dControl"
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange12wControl RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "12wControl"
	RadarHTTPTopLocationHTTPProtocolGetParamsDateRange24wControl RadarHTTPTopLocationHTTPProtocolGetParamsDateRange = "24wControl"
)

type RadarHTTPTopLocationHTTPProtocolGetParamsDeviceType string

const (
	RadarHTTPTopLocationHTTPProtocolGetParamsDeviceTypeDesktop RadarHTTPTopLocationHTTPProtocolGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationHTTPProtocolGetParamsDeviceTypeMobile  RadarHTTPTopLocationHTTPProtocolGetParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationHTTPProtocolGetParamsDeviceTypeOther   RadarHTTPTopLocationHTTPProtocolGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopLocationHTTPProtocolGetParamsFormat string

const (
	RadarHTTPTopLocationHTTPProtocolGetParamsFormatJson RadarHTTPTopLocationHTTPProtocolGetParamsFormat = "JSON"
	RadarHTTPTopLocationHTTPProtocolGetParamsFormatCsv  RadarHTTPTopLocationHTTPProtocolGetParamsFormat = "CSV"
)

type RadarHTTPTopLocationHTTPProtocolGetParamsIPVersion string

const (
	RadarHTTPTopLocationHTTPProtocolGetParamsIPVersionIPv4 RadarHTTPTopLocationHTTPProtocolGetParamsIPVersion = "IPv4"
	RadarHTTPTopLocationHTTPProtocolGetParamsIPVersionIPv6 RadarHTTPTopLocationHTTPProtocolGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopLocationHTTPProtocolGetParamsO string

const (
	RadarHTTPTopLocationHTTPProtocolGetParamsOWindows  RadarHTTPTopLocationHTTPProtocolGetParamsO = "WINDOWS"
	RadarHTTPTopLocationHTTPProtocolGetParamsOMacosx   RadarHTTPTopLocationHTTPProtocolGetParamsO = "MACOSX"
	RadarHTTPTopLocationHTTPProtocolGetParamsOAndroid  RadarHTTPTopLocationHTTPProtocolGetParamsO = "ANDROID"
	RadarHTTPTopLocationHTTPProtocolGetParamsOChromeos RadarHTTPTopLocationHTTPProtocolGetParamsO = "CHROMEOS"
	RadarHTTPTopLocationHTTPProtocolGetParamsOLinux    RadarHTTPTopLocationHTTPProtocolGetParamsO = "LINUX"
	RadarHTTPTopLocationHTTPProtocolGetParamsOSmartTv  RadarHTTPTopLocationHTTPProtocolGetParamsO = "SMART_TV"
)

type RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersion string

const (
	RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersionTlSv1_0  RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersionTlSv1_1  RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersionTlSv1_2  RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersionTlSv1_3  RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersionTlSvQuic RadarHTTPTopLocationHTTPProtocolGetParamsTlsVersion = "TLSvQUIC"
)
