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

// RadarHTTPTopAseHTTPProtocolService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTopAseHTTPProtocolService] method instead.
type RadarHTTPTopAseHTTPProtocolService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopAseHTTPProtocolService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTopAseHTTPProtocolService(opts ...option.RequestOption) (r *RadarHTTPTopAseHTTPProtocolService) {
	r = &RadarHTTPTopAseHTTPProtocolService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested HTTP
// protocol. Values are a percentage out of the total traffic.
func (r *RadarHTTPTopAseHTTPProtocolService) Get(ctx context.Context, httpProtocol RadarHTTPTopAseHTTPProtocolGetParamsHTTPProtocol, query RadarHTTPTopAseHTTPProtocolGetParams, opts ...option.RequestOption) (res *RadarHTTPTopAseHTTPProtocolGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/http_protocol/%v", httpProtocol)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopAseHTTPProtocolGetResponse struct {
	Result  RadarHTTPTopAseHTTPProtocolGetResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarHTTPTopAseHTTPProtocolGetResponseJSON   `json:"-"`
}

// radarHTTPTopAseHTTPProtocolGetResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseHTTPProtocolGetResponse]
type radarHTTPTopAseHTTPProtocolGetResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseHTTPProtocolGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPProtocolGetResponseResult struct {
	Meta RadarHTTPTopAseHTTPProtocolGetResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseHTTPProtocolGetResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseHTTPProtocolGetResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseHTTPProtocolGetResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTopAseHTTPProtocolGetResponseResult]
type radarHTTPTopAseHTTPProtocolGetResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseHTTPProtocolGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPProtocolGetResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseHTTPProtocolGetResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                         `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseHTTPProtocolGetResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseHTTPProtocolGetResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseHTTPProtocolGetResponseResultMeta]
type radarHTTPTopAseHTTPProtocolGetResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseHTTPProtocolGetResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPProtocolGetResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                     `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseHTTPProtocolGetResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseHTTPProtocolGetResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopAseHTTPProtocolGetResponseResultMetaDateRange]
type radarHTTPTopAseHTTPProtocolGetResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseHTTPProtocolGetResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                      `json:"level"`
	JSON        radarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfo]
type radarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                       `json:"dataSource,required"`
	Description     string                                                                       `json:"description,required"`
	EventType       string                                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                                    `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseHTTPProtocolGetResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPProtocolGetResponseResultTop0 struct {
	ClientASN    int64                                                `json:"clientASN,required"`
	ClientAsName string                                               `json:"clientASName,required"`
	Value        string                                               `json:"value,required"`
	JSON         radarHTTPTopAseHTTPProtocolGetResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseHTTPProtocolGetResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarHTTPTopAseHTTPProtocolGetResponseResultTop0]
type radarHTTPTopAseHTTPProtocolGetResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseHTTPProtocolGetResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseHTTPProtocolGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseHTTPProtocolGetParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopAseHTTPProtocolGetParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopAseHTTPProtocolGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopAseHTTPProtocolGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopAseHTTPProtocolGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopAseHTTPProtocolGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopAseHTTPProtocolGetParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopAseHTTPProtocolGetParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseHTTPProtocolGetParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopAseHTTPProtocolGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// HTTP Protocol.
type RadarHTTPTopAseHTTPProtocolGetParamsHTTPProtocol string

const (
	RadarHTTPTopAseHTTPProtocolGetParamsHTTPProtocolHTTP  RadarHTTPTopAseHTTPProtocolGetParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseHTTPProtocolGetParamsHTTPProtocolHTTPs RadarHTTPTopAseHTTPProtocolGetParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopAseHTTPProtocolGetParamsBotClass string

const (
	RadarHTTPTopAseHTTPProtocolGetParamsBotClassLikelyAutomated RadarHTTPTopAseHTTPProtocolGetParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseHTTPProtocolGetParamsBotClassLikelyHuman     RadarHTTPTopAseHTTPProtocolGetParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopAseHTTPProtocolGetParamsDateRange string

const (
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange1d         RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "1d"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange2d         RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "2d"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange7d         RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "7d"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange14d        RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "14d"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange28d        RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "28d"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange12w        RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "12w"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange24w        RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "24w"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange52w        RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "52w"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange1dControl  RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "1dControl"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange2dControl  RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "2dControl"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange7dControl  RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "7dControl"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange14dControl RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "14dControl"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange28dControl RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "28dControl"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange12wControl RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "12wControl"
	RadarHTTPTopAseHTTPProtocolGetParamsDateRange24wControl RadarHTTPTopAseHTTPProtocolGetParamsDateRange = "24wControl"
)

type RadarHTTPTopAseHTTPProtocolGetParamsDeviceType string

const (
	RadarHTTPTopAseHTTPProtocolGetParamsDeviceTypeDesktop RadarHTTPTopAseHTTPProtocolGetParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseHTTPProtocolGetParamsDeviceTypeMobile  RadarHTTPTopAseHTTPProtocolGetParamsDeviceType = "MOBILE"
	RadarHTTPTopAseHTTPProtocolGetParamsDeviceTypeOther   RadarHTTPTopAseHTTPProtocolGetParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopAseHTTPProtocolGetParamsFormat string

const (
	RadarHTTPTopAseHTTPProtocolGetParamsFormatJson RadarHTTPTopAseHTTPProtocolGetParamsFormat = "JSON"
	RadarHTTPTopAseHTTPProtocolGetParamsFormatCsv  RadarHTTPTopAseHTTPProtocolGetParamsFormat = "CSV"
)

type RadarHTTPTopAseHTTPProtocolGetParamsIPVersion string

const (
	RadarHTTPTopAseHTTPProtocolGetParamsIPVersionIPv4 RadarHTTPTopAseHTTPProtocolGetParamsIPVersion = "IPv4"
	RadarHTTPTopAseHTTPProtocolGetParamsIPVersionIPv6 RadarHTTPTopAseHTTPProtocolGetParamsIPVersion = "IPv6"
)

type RadarHTTPTopAseHTTPProtocolGetParamsO string

const (
	RadarHTTPTopAseHTTPProtocolGetParamsOWindows  RadarHTTPTopAseHTTPProtocolGetParamsO = "WINDOWS"
	RadarHTTPTopAseHTTPProtocolGetParamsOMacosx   RadarHTTPTopAseHTTPProtocolGetParamsO = "MACOSX"
	RadarHTTPTopAseHTTPProtocolGetParamsOIos      RadarHTTPTopAseHTTPProtocolGetParamsO = "IOS"
	RadarHTTPTopAseHTTPProtocolGetParamsOAndroid  RadarHTTPTopAseHTTPProtocolGetParamsO = "ANDROID"
	RadarHTTPTopAseHTTPProtocolGetParamsOChromeos RadarHTTPTopAseHTTPProtocolGetParamsO = "CHROMEOS"
	RadarHTTPTopAseHTTPProtocolGetParamsOLinux    RadarHTTPTopAseHTTPProtocolGetParamsO = "LINUX"
	RadarHTTPTopAseHTTPProtocolGetParamsOSmartTv  RadarHTTPTopAseHTTPProtocolGetParamsO = "SMART_TV"
)

type RadarHTTPTopAseHTTPProtocolGetParamsTlsVersion string

const (
	RadarHTTPTopAseHTTPProtocolGetParamsTlsVersionTlSv1_0  RadarHTTPTopAseHTTPProtocolGetParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseHTTPProtocolGetParamsTlsVersionTlSv1_1  RadarHTTPTopAseHTTPProtocolGetParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseHTTPProtocolGetParamsTlsVersionTlSv1_2  RadarHTTPTopAseHTTPProtocolGetParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseHTTPProtocolGetParamsTlsVersionTlSv1_3  RadarHTTPTopAseHTTPProtocolGetParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseHTTPProtocolGetParamsTlsVersionTlSvQuic RadarHTTPTopAseHTTPProtocolGetParamsTlsVersion = "TLSvQUIC"
)
