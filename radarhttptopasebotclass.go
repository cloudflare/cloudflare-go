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

// RadarHTTPTopAseBotClassService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTopAseBotClassService] method instead.
type RadarHTTPTopAseBotClassService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopAseBotClassService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPTopAseBotClassService(opts ...option.RequestOption) (r *RadarHTTPTopAseBotClassService) {
	r = &RadarHTTPTopAseBotClassService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested bot
// class. These two categories use Cloudflare's bot score - refer to
// [Bot Scores](https://developers.cloudflare.com/bots/concepts/bot-score) for more
// information. Values are a percentage out of the total traffic.
func (r *RadarHTTPTopAseBotClassService) Get(ctx context.Context, botClass RadarHTTPTopAseBotClassGetParamsBotClass, query RadarHTTPTopAseBotClassGetParams, opts ...option.RequestOption) (res *RadarHTTPTopAseBotClassGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/bot_class/%v", botClass)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopAseBotClassGetResponse struct {
	Result  RadarHTTPTopAseBotClassGetResponseResult `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarHTTPTopAseBotClassGetResponseJSON   `json:"-"`
}

// radarHTTPTopAseBotClassGetResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTopAseBotClassGetResponse]
type radarHTTPTopAseBotClassGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseBotClassGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseBotClassGetResponseResult struct {
	Meta RadarHTTPTopAseBotClassGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseBotClassGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseBotClassGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseBotClassGetResponseResultJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseBotClassGetResponseResult]
type radarHTTPTopAseBotClassGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseBotClassGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseBotClassGetResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseBotClassGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                     `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseBotClassGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseBotClassGetResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarHTTPTopAseBotClassGetResponseResultMeta]
type radarHTTPTopAseBotClassGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseBotClassGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseBotClassGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                 `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseBotClassGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseBotClassGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPTopAseBotClassGetResponseResultMetaDateRange]
type radarHTTPTopAseBotClassGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseBotClassGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                  `json:"level"`
	JSON        radarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfo]
type radarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                   `json:"dataSource,required"`
	Description     string                                                                   `json:"description,required"`
	EventType       string                                                                   `json:"eventType,required"`
	IsInstantaneous interface{}                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                   `json:"linkedUrl"`
	StartTime       time.Time                                                                `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseBotClassGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseBotClassGetResponseResultTop0 struct {
	ClientASN    int64                                            `json:"clientASN,required"`
	ClientAsName string                                           `json:"clientASName,required"`
	Value        string                                           `json:"value,required"`
	JSON         radarHTTPTopAseBotClassGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseBotClassGetResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarHTTPTopAseBotClassGetResponseResultTop0]
type radarHTTPTopAseBotClassGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseBotClassGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseBotClassGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopAseBotClassGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopAseBotClassGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopAseBotClassGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopAseBotClassGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopAseBotClassGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopAseBotClassGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopAseBotClassGetParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopAseBotClassGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseBotClassGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopAseBotClassGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Bot class.
type RadarHTTPTopAseBotClassGetParamsBotClass string

const (
	RadarHTTPTopAseBotClassGetParamsBotClassLikelyAutomated RadarHTTPTopAseBotClassGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseBotClassGetParamsBotClassLikelyHuman     RadarHTTPTopAseBotClassGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopAseBotClassGetParamsDateRange string

const (
	RadarHTTPTopAseBotClassGetParamsDateRange1d         RadarHTTPTopAseBotClassGetParamsDateRange = "1d"
	RadarHTTPTopAseBotClassGetParamsDateRange2d         RadarHTTPTopAseBotClassGetParamsDateRange = "2d"
	RadarHTTPTopAseBotClassGetParamsDateRange7d         RadarHTTPTopAseBotClassGetParamsDateRange = "7d"
	RadarHTTPTopAseBotClassGetParamsDateRange14d        RadarHTTPTopAseBotClassGetParamsDateRange = "14d"
	RadarHTTPTopAseBotClassGetParamsDateRange28d        RadarHTTPTopAseBotClassGetParamsDateRange = "28d"
	RadarHTTPTopAseBotClassGetParamsDateRange12w        RadarHTTPTopAseBotClassGetParamsDateRange = "12w"
	RadarHTTPTopAseBotClassGetParamsDateRange24w        RadarHTTPTopAseBotClassGetParamsDateRange = "24w"
	RadarHTTPTopAseBotClassGetParamsDateRange52w        RadarHTTPTopAseBotClassGetParamsDateRange = "52w"
	RadarHTTPTopAseBotClassGetParamsDateRange1dControl  RadarHTTPTopAseBotClassGetParamsDateRange = "1dControl"
	RadarHTTPTopAseBotClassGetParamsDateRange2dControl  RadarHTTPTopAseBotClassGetParamsDateRange = "2dControl"
	RadarHTTPTopAseBotClassGetParamsDateRange7dControl  RadarHTTPTopAseBotClassGetParamsDateRange = "7dControl"
	RadarHTTPTopAseBotClassGetParamsDateRange14dControl RadarHTTPTopAseBotClassGetParamsDateRange = "14dControl"
	RadarHTTPTopAseBotClassGetParamsDateRange28dControl RadarHTTPTopAseBotClassGetParamsDateRange = "28dControl"
	RadarHTTPTopAseBotClassGetParamsDateRange12wControl RadarHTTPTopAseBotClassGetParamsDateRange = "12wControl"
	RadarHTTPTopAseBotClassGetParamsDateRange24wControl RadarHTTPTopAseBotClassGetParamsDateRange = "24wControl"
)

type RadarHTTPTopAseBotClassGetParamsDeviceType string

const (
	RadarHTTPTopAseBotClassGetParamsDeviceTypeDesktop RadarHTTPTopAseBotClassGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseBotClassGetParamsDeviceTypeMobile  RadarHTTPTopAseBotClassGetParamsDeviceType = "MOBILE"
	RadarHTTPTopAseBotClassGetParamsDeviceTypeOther   RadarHTTPTopAseBotClassGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopAseBotClassGetParamsFormat string

const (
	RadarHTTPTopAseBotClassGetParamsFormatJson RadarHTTPTopAseBotClassGetParamsFormat = "JSON"
	RadarHTTPTopAseBotClassGetParamsFormatCsv  RadarHTTPTopAseBotClassGetParamsFormat = "CSV"
)

type RadarHTTPTopAseBotClassGetParamsHTTPProtocol string

const (
	RadarHTTPTopAseBotClassGetParamsHTTPProtocolHTTP  RadarHTTPTopAseBotClassGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseBotClassGetParamsHTTPProtocolHTTPs RadarHTTPTopAseBotClassGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopAseBotClassGetParamsHTTPVersion string

const (
	RadarHTTPTopAseBotClassGetParamsHTTPVersionHttPv1 RadarHTTPTopAseBotClassGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseBotClassGetParamsHTTPVersionHttPv2 RadarHTTPTopAseBotClassGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseBotClassGetParamsHTTPVersionHttPv3 RadarHTTPTopAseBotClassGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopAseBotClassGetParamsIPVersion string

const (
	RadarHTTPTopAseBotClassGetParamsIPVersionIPv4 RadarHTTPTopAseBotClassGetParamsIPVersion = "IPv4"
	RadarHTTPTopAseBotClassGetParamsIPVersionIPv6 RadarHTTPTopAseBotClassGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopAseBotClassGetParamsO string

const (
	RadarHTTPTopAseBotClassGetParamsOWindows  RadarHTTPTopAseBotClassGetParamsO = "WINDOWS"
	RadarHTTPTopAseBotClassGetParamsOMacosx   RadarHTTPTopAseBotClassGetParamsO = "MACOSX"
	RadarHTTPTopAseBotClassGetParamsOIos      RadarHTTPTopAseBotClassGetParamsO = "IOS"
	RadarHTTPTopAseBotClassGetParamsOAndroid  RadarHTTPTopAseBotClassGetParamsO = "ANDROID"
	RadarHTTPTopAseBotClassGetParamsOChromeos RadarHTTPTopAseBotClassGetParamsO = "CHROMEOS"
	RadarHTTPTopAseBotClassGetParamsOLinux    RadarHTTPTopAseBotClassGetParamsO = "LINUX"
	RadarHTTPTopAseBotClassGetParamsOSmartTv  RadarHTTPTopAseBotClassGetParamsO = "SMART_TV"
)

type RadarHTTPTopAseBotClassGetParamsTlsVersion string

const (
	RadarHTTPTopAseBotClassGetParamsTlsVersionTlSv1_0  RadarHTTPTopAseBotClassGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseBotClassGetParamsTlsVersionTlSv1_1  RadarHTTPTopAseBotClassGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseBotClassGetParamsTlsVersionTlSv1_2  RadarHTTPTopAseBotClassGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseBotClassGetParamsTlsVersionTlSv1_3  RadarHTTPTopAseBotClassGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseBotClassGetParamsTlsVersionTlSvQuic RadarHTTPTopAseBotClassGetParamsTlsVersion = "TLSvQUIC"
)
