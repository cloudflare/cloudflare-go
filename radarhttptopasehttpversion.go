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

// RadarHTTPTopAseHTTPVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTopAseHTTPVersionService] method instead.
type RadarHTTPTopAseHTTPVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopAseHTTPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTopAseHTTPVersionService(opts ...option.RequestOption) (r *RadarHTTPTopAseHTTPVersionService) {
	r = &RadarHTTPTopAseHTTPVersionService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested HTTP
// protocol version. Values are a percentage out of the total traffic.
func (r *RadarHTTPTopAseHTTPVersionService) Get(ctx context.Context, httpVersion RadarHTTPTopAseHTTPVersionGetParamsHTTPVersion, query RadarHTTPTopAseHTTPVersionGetParams, opts ...option.RequestOption) (res *RadarHTTPTopAseHTTPVersionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/http_version/%v", httpVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopAseHTTPVersionGetResponse struct {
	Result  RadarHTTPTopAseHTTPVersionGetResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarHTTPTopAseHTTPVersionGetResponseJSON   `json:"-"`
}

// radarHTTPTopAseHTTPVersionGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseHTTPVersionGetResponse]
type radarHTTPTopAseHTTPVersionGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseHTTPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPVersionGetResponseResult struct {
	Meta RadarHTTPTopAseHTTPVersionGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseHTTPVersionGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseHTTPVersionGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseHTTPVersionGetResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTopAseHTTPVersionGetResponseResult]
type radarHTTPTopAseHTTPVersionGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseHTTPVersionGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPVersionGetResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseHTTPVersionGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseHTTPVersionGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseHTTPVersionGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseHTTPVersionGetResponseResultMeta]
type radarHTTPTopAseHTTPVersionGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseHTTPVersionGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPVersionGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseHTTPVersionGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseHTTPVersionGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopAseHTTPVersionGetResponseResultMetaDateRange]
type radarHTTPTopAseHTTPVersionGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseHTTPVersionGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfo]
type radarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseHTTPVersionGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPVersionGetResponseResultTop0 struct {
	ClientASN    int64                                               `json:"clientASN,required"`
	ClientAsName string                                              `json:"clientASName,required"`
	Value        string                                              `json:"value,required"`
	JSON         radarHTTPTopAseHTTPVersionGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseHTTPVersionGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarHTTPTopAseHTTPVersionGetResponseResultTop0]
type radarHTTPTopAseHTTPVersionGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseHTTPVersionGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseHTTPVersionGetParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopAseHTTPVersionGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopAseHTTPVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopAseHTTPVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopAseHTTPVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopAseHTTPVersionGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopAseHTTPVersionGetParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopAseHTTPVersionGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseHTTPVersionGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopAseHTTPVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// HTTP version.
type RadarHTTPTopAseHTTPVersionGetParamsHTTPVersion string

const (
	RadarHTTPTopAseHTTPVersionGetParamsHTTPVersionHttPv1 RadarHTTPTopAseHTTPVersionGetParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseHTTPVersionGetParamsHTTPVersionHttPv2 RadarHTTPTopAseHTTPVersionGetParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseHTTPVersionGetParamsHTTPVersionHttPv3 RadarHTTPTopAseHTTPVersionGetParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopAseHTTPVersionGetParamsBotClass string

const (
	RadarHTTPTopAseHTTPVersionGetParamsBotClassLikelyAutomated RadarHTTPTopAseHTTPVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseHTTPVersionGetParamsBotClassLikelyHuman     RadarHTTPTopAseHTTPVersionGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopAseHTTPVersionGetParamsDateRange string

const (
	RadarHTTPTopAseHTTPVersionGetParamsDateRange1d         RadarHTTPTopAseHTTPVersionGetParamsDateRange = "1d"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange2d         RadarHTTPTopAseHTTPVersionGetParamsDateRange = "2d"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange7d         RadarHTTPTopAseHTTPVersionGetParamsDateRange = "7d"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange14d        RadarHTTPTopAseHTTPVersionGetParamsDateRange = "14d"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange28d        RadarHTTPTopAseHTTPVersionGetParamsDateRange = "28d"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange12w        RadarHTTPTopAseHTTPVersionGetParamsDateRange = "12w"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange24w        RadarHTTPTopAseHTTPVersionGetParamsDateRange = "24w"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange52w        RadarHTTPTopAseHTTPVersionGetParamsDateRange = "52w"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange1dControl  RadarHTTPTopAseHTTPVersionGetParamsDateRange = "1dControl"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange2dControl  RadarHTTPTopAseHTTPVersionGetParamsDateRange = "2dControl"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange7dControl  RadarHTTPTopAseHTTPVersionGetParamsDateRange = "7dControl"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange14dControl RadarHTTPTopAseHTTPVersionGetParamsDateRange = "14dControl"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange28dControl RadarHTTPTopAseHTTPVersionGetParamsDateRange = "28dControl"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange12wControl RadarHTTPTopAseHTTPVersionGetParamsDateRange = "12wControl"
	RadarHTTPTopAseHTTPVersionGetParamsDateRange24wControl RadarHTTPTopAseHTTPVersionGetParamsDateRange = "24wControl"
)

type RadarHTTPTopAseHTTPVersionGetParamsDeviceType string

const (
	RadarHTTPTopAseHTTPVersionGetParamsDeviceTypeDesktop RadarHTTPTopAseHTTPVersionGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseHTTPVersionGetParamsDeviceTypeMobile  RadarHTTPTopAseHTTPVersionGetParamsDeviceType = "MOBILE"
	RadarHTTPTopAseHTTPVersionGetParamsDeviceTypeOther   RadarHTTPTopAseHTTPVersionGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopAseHTTPVersionGetParamsFormat string

const (
	RadarHTTPTopAseHTTPVersionGetParamsFormatJson RadarHTTPTopAseHTTPVersionGetParamsFormat = "JSON"
	RadarHTTPTopAseHTTPVersionGetParamsFormatCsv  RadarHTTPTopAseHTTPVersionGetParamsFormat = "CSV"
)

type RadarHTTPTopAseHTTPVersionGetParamsHTTPProtocol string

const (
	RadarHTTPTopAseHTTPVersionGetParamsHTTPProtocolHTTP  RadarHTTPTopAseHTTPVersionGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseHTTPVersionGetParamsHTTPProtocolHTTPs RadarHTTPTopAseHTTPVersionGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopAseHTTPVersionGetParamsIPVersion string

const (
	RadarHTTPTopAseHTTPVersionGetParamsIPVersionIPv4 RadarHTTPTopAseHTTPVersionGetParamsIPVersion = "IPv4"
	RadarHTTPTopAseHTTPVersionGetParamsIPVersionIPv6 RadarHTTPTopAseHTTPVersionGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopAseHTTPVersionGetParamsO string

const (
	RadarHTTPTopAseHTTPVersionGetParamsOWindows  RadarHTTPTopAseHTTPVersionGetParamsO = "WINDOWS"
	RadarHTTPTopAseHTTPVersionGetParamsOMacosx   RadarHTTPTopAseHTTPVersionGetParamsO = "MACOSX"
	RadarHTTPTopAseHTTPVersionGetParamsOIos      RadarHTTPTopAseHTTPVersionGetParamsO = "IOS"
	RadarHTTPTopAseHTTPVersionGetParamsOAndroid  RadarHTTPTopAseHTTPVersionGetParamsO = "ANDROID"
	RadarHTTPTopAseHTTPVersionGetParamsOChromeos RadarHTTPTopAseHTTPVersionGetParamsO = "CHROMEOS"
	RadarHTTPTopAseHTTPVersionGetParamsOLinux    RadarHTTPTopAseHTTPVersionGetParamsO = "LINUX"
	RadarHTTPTopAseHTTPVersionGetParamsOSmartTv  RadarHTTPTopAseHTTPVersionGetParamsO = "SMART_TV"
)

type RadarHTTPTopAseHTTPVersionGetParamsTlsVersion string

const (
	RadarHTTPTopAseHTTPVersionGetParamsTlsVersionTlSv1_0  RadarHTTPTopAseHTTPVersionGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseHTTPVersionGetParamsTlsVersionTlSv1_1  RadarHTTPTopAseHTTPVersionGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseHTTPVersionGetParamsTlsVersionTlSv1_2  RadarHTTPTopAseHTTPVersionGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseHTTPVersionGetParamsTlsVersionTlSv1_3  RadarHTTPTopAseHTTPVersionGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseHTTPVersionGetParamsTlsVersionTlSvQuic RadarHTTPTopAseHTTPVersionGetParamsTlsVersion = "TLSvQUIC"
)
