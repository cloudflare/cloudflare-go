// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarHTTPTopAseService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPTopAseService] method
// instead.
type RadarHTTPTopAseService struct {
	Options       []option.RequestOption
	BotClasses    *RadarHTTPTopAseBotClassService
	DeviceTypes   *RadarHTTPTopAseDeviceTypeService
	HTTPProtocols *RadarHTTPTopAseHTTPProtocolService
	HTTPVersions  *RadarHTTPTopAseHTTPVersionService
	IPVersions    *RadarHTTPTopAseIPVersionService
	Os            *RadarHTTPTopAseOService
	TlsVersions   *RadarHTTPTopAseTlsVersionService
}

// NewRadarHTTPTopAseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarHTTPTopAseService(opts ...option.RequestOption) (r *RadarHTTPTopAseService) {
	r = &RadarHTTPTopAseService{}
	r.Options = opts
	r.BotClasses = NewRadarHTTPTopAseBotClassService(opts...)
	r.DeviceTypes = NewRadarHTTPTopAseDeviceTypeService(opts...)
	r.HTTPProtocols = NewRadarHTTPTopAseHTTPProtocolService(opts...)
	r.HTTPVersions = NewRadarHTTPTopAseHTTPVersionService(opts...)
	r.IPVersions = NewRadarHTTPTopAseIPVersionService(opts...)
	r.Os = NewRadarHTTPTopAseOService(opts...)
	r.TlsVersions = NewRadarHTTPTopAseTlsVersionService(opts...)
	return
}

// Get the top autonomous systems by HTTP traffic. Values are a percentage out of
// the total traffic.
func (r *RadarHTTPTopAseService) List(ctx context.Context, query RadarHTTPTopAseListParams, opts ...option.RequestOption) (res *RadarHTTPTopAseListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/ases"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopAseListResponse struct {
	Result  RadarHTTPTopAseListResponseResult `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    radarHTTPTopAseListResponseJSON   `json:"-"`
}

// radarHTTPTopAseListResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTopAseListResponse]
type radarHTTPTopAseListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseListResponseResult struct {
	Meta RadarHTTPTopAseListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopAseListResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopAseListResponseResultJSON   `json:"-"`
}

// radarHTTPTopAseListResponseResultJSON contains the JSON metadata for the struct
// [RadarHTTPTopAseListResponseResult]
type radarHTTPTopAseListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseListResponseResultMeta struct {
	DateRange      []RadarHTTPTopAseListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopAseListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopAseListResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopAseListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListResponseResultMeta]
type radarHTTPTopAseListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopAseListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopAseListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopAseListResponseResultMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarHTTPTopAseListResponseResultMetaDateRange]
type radarHTTPTopAseListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        radarHTTPTopAseListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopAseListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarHTTPTopAseListResponseResultMetaConfidenceInfo]
type radarHTTPTopAseListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopAseListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous interface{}                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopAseListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseListResponseResultTop0 struct {
	ClientASN    int64                                     `json:"clientASN,required"`
	ClientAsName string                                    `json:"clientASName,required"`
	Value        string                                    `json:"value,required"`
	JSON         radarHTTPTopAseListResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopAseListResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarHTTPTopAseListResponseResultTop0]
type radarHTTPTopAseListResponseResultTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *RadarHTTPTopAseListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopAseListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopAseListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopAseListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopAseListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopAseListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopAseListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopAseListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopAseListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopAseListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopAseListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopAseListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopAseListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopAseListParamsBotClass string

const (
	RadarHTTPTopAseListParamsBotClassLikelyAutomated RadarHTTPTopAseListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopAseListParamsBotClassLikelyHuman     RadarHTTPTopAseListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopAseListParamsDateRange string

const (
	RadarHTTPTopAseListParamsDateRange1d         RadarHTTPTopAseListParamsDateRange = "1d"
	RadarHTTPTopAseListParamsDateRange2d         RadarHTTPTopAseListParamsDateRange = "2d"
	RadarHTTPTopAseListParamsDateRange7d         RadarHTTPTopAseListParamsDateRange = "7d"
	RadarHTTPTopAseListParamsDateRange14d        RadarHTTPTopAseListParamsDateRange = "14d"
	RadarHTTPTopAseListParamsDateRange28d        RadarHTTPTopAseListParamsDateRange = "28d"
	RadarHTTPTopAseListParamsDateRange12w        RadarHTTPTopAseListParamsDateRange = "12w"
	RadarHTTPTopAseListParamsDateRange24w        RadarHTTPTopAseListParamsDateRange = "24w"
	RadarHTTPTopAseListParamsDateRange52w        RadarHTTPTopAseListParamsDateRange = "52w"
	RadarHTTPTopAseListParamsDateRange1dControl  RadarHTTPTopAseListParamsDateRange = "1dControl"
	RadarHTTPTopAseListParamsDateRange2dControl  RadarHTTPTopAseListParamsDateRange = "2dControl"
	RadarHTTPTopAseListParamsDateRange7dControl  RadarHTTPTopAseListParamsDateRange = "7dControl"
	RadarHTTPTopAseListParamsDateRange14dControl RadarHTTPTopAseListParamsDateRange = "14dControl"
	RadarHTTPTopAseListParamsDateRange28dControl RadarHTTPTopAseListParamsDateRange = "28dControl"
	RadarHTTPTopAseListParamsDateRange12wControl RadarHTTPTopAseListParamsDateRange = "12wControl"
	RadarHTTPTopAseListParamsDateRange24wControl RadarHTTPTopAseListParamsDateRange = "24wControl"
)

type RadarHTTPTopAseListParamsDeviceType string

const (
	RadarHTTPTopAseListParamsDeviceTypeDesktop RadarHTTPTopAseListParamsDeviceType = "DESKTOP"
	RadarHTTPTopAseListParamsDeviceTypeMobile  RadarHTTPTopAseListParamsDeviceType = "MOBILE"
	RadarHTTPTopAseListParamsDeviceTypeOther   RadarHTTPTopAseListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopAseListParamsFormat string

const (
	RadarHTTPTopAseListParamsFormatJson RadarHTTPTopAseListParamsFormat = "JSON"
	RadarHTTPTopAseListParamsFormatCsv  RadarHTTPTopAseListParamsFormat = "CSV"
)

type RadarHTTPTopAseListParamsHTTPProtocol string

const (
	RadarHTTPTopAseListParamsHTTPProtocolHTTP  RadarHTTPTopAseListParamsHTTPProtocol = "HTTP"
	RadarHTTPTopAseListParamsHTTPProtocolHTTPs RadarHTTPTopAseListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopAseListParamsHTTPVersion string

const (
	RadarHTTPTopAseListParamsHTTPVersionHttPv1 RadarHTTPTopAseListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopAseListParamsHTTPVersionHttPv2 RadarHTTPTopAseListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopAseListParamsHTTPVersionHttPv3 RadarHTTPTopAseListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopAseListParamsIPVersion string

const (
	RadarHTTPTopAseListParamsIPVersionIPv4 RadarHTTPTopAseListParamsIPVersion = "IPv4"
	RadarHTTPTopAseListParamsIPVersionIPv6 RadarHTTPTopAseListParamsIPVersion = "IPv6"
)

type RadarHTTPTopAseListParamsO string

const (
	RadarHTTPTopAseListParamsOWindows  RadarHTTPTopAseListParamsO = "WINDOWS"
	RadarHTTPTopAseListParamsOMacosx   RadarHTTPTopAseListParamsO = "MACOSX"
	RadarHTTPTopAseListParamsOIos      RadarHTTPTopAseListParamsO = "IOS"
	RadarHTTPTopAseListParamsOAndroid  RadarHTTPTopAseListParamsO = "ANDROID"
	RadarHTTPTopAseListParamsOChromeos RadarHTTPTopAseListParamsO = "CHROMEOS"
	RadarHTTPTopAseListParamsOLinux    RadarHTTPTopAseListParamsO = "LINUX"
	RadarHTTPTopAseListParamsOSmartTv  RadarHTTPTopAseListParamsO = "SMART_TV"
)

type RadarHTTPTopAseListParamsTlsVersion string

const (
	RadarHTTPTopAseListParamsTlsVersionTlSv1_0  RadarHTTPTopAseListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopAseListParamsTlsVersionTlSv1_1  RadarHTTPTopAseListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopAseListParamsTlsVersionTlSv1_2  RadarHTTPTopAseListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopAseListParamsTlsVersionTlSv1_3  RadarHTTPTopAseListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopAseListParamsTlsVersionTlSvQuic RadarHTTPTopAseListParamsTlsVersion = "TLSvQUIC"
)
