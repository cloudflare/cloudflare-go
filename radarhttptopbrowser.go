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

// RadarHTTPTopBrowserService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPTopBrowserService]
// method instead.
type RadarHTTPTopBrowserService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopBrowserService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPTopBrowserService(opts ...option.RequestOption) (r *RadarHTTPTopBrowserService) {
	r = &RadarHTTPTopBrowserService{}
	r.Options = opts
	return
}

// Get the top user agents by HTTP traffic. Values are a percentage out of the
// total traffic.
func (r *RadarHTTPTopBrowserService) List(ctx context.Context, query RadarHTTPTopBrowserListParams, opts ...option.RequestOption) (res *RadarHTTPTopBrowserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/browsers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopBrowserListResponse struct {
	Result  RadarHTTPTopBrowserListResponseResult `json:"result,required"`
	Success bool                                  `json:"success,required"`
	JSON    radarHTTPTopBrowserListResponseJSON   `json:"-"`
}

// radarHTTPTopBrowserListResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTopBrowserListResponse]
type radarHTTPTopBrowserListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserListResponseResult struct {
	Meta RadarHTTPTopBrowserListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopBrowserListResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopBrowserListResponseResultJSON   `json:"-"`
}

// radarHTTPTopBrowserListResponseResultJSON contains the JSON metadata for the
// struct [RadarHTTPTopBrowserListResponseResult]
type radarHTTPTopBrowserListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserListResponseResultMeta struct {
	DateRange      []RadarHTTPTopBrowserListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                  `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopBrowserListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopBrowserListResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopBrowserListResponseResultMetaJSON contains the JSON metadata for the
// struct [RadarHTTPTopBrowserListResponseResultMeta]
type radarHTTPTopBrowserListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                              `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopBrowserListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopBrowserListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPTopBrowserListResponseResultMetaDateRange]
type radarHTTPTopBrowserListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopBrowserListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                               `json:"level"`
	JSON        radarHTTPTopBrowserListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopBrowserListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopBrowserListResponseResultMetaConfidenceInfo]
type radarHTTPTopBrowserListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                `json:"dataSource,required"`
	Description     string                                                                `json:"description,required"`
	EventType       string                                                                `json:"eventType,required"`
	IsInstantaneous interface{}                                                           `json:"isInstantaneous,required"`
	EndTime         time.Time                                                             `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                `json:"linkedUrl"`
	StartTime       time.Time                                                             `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopBrowserListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopBrowserListResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopBrowserListResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopBrowserListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopBrowserListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserListResponseResultTop0 struct {
	Name  string                                        `json:"name,required"`
	Value string                                        `json:"value,required"`
	JSON  radarHTTPTopBrowserListResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopBrowserListResponseResultTop0JSON contains the JSON metadata for the
// struct [RadarHTTPTopBrowserListResponseResultTop0]
type radarHTTPTopBrowserListResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopBrowserListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopBrowserListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopBrowserListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopBrowserListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopBrowserListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopBrowserListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopBrowserListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopBrowserListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopBrowserListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopBrowserListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopBrowserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopBrowserListParamsBotClass string

const (
	RadarHTTPTopBrowserListParamsBotClassLikelyAutomated RadarHTTPTopBrowserListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopBrowserListParamsBotClassLikelyHuman     RadarHTTPTopBrowserListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopBrowserListParamsDateRange string

const (
	RadarHTTPTopBrowserListParamsDateRange1d         RadarHTTPTopBrowserListParamsDateRange = "1d"
	RadarHTTPTopBrowserListParamsDateRange2d         RadarHTTPTopBrowserListParamsDateRange = "2d"
	RadarHTTPTopBrowserListParamsDateRange7d         RadarHTTPTopBrowserListParamsDateRange = "7d"
	RadarHTTPTopBrowserListParamsDateRange14d        RadarHTTPTopBrowserListParamsDateRange = "14d"
	RadarHTTPTopBrowserListParamsDateRange28d        RadarHTTPTopBrowserListParamsDateRange = "28d"
	RadarHTTPTopBrowserListParamsDateRange12w        RadarHTTPTopBrowserListParamsDateRange = "12w"
	RadarHTTPTopBrowserListParamsDateRange24w        RadarHTTPTopBrowserListParamsDateRange = "24w"
	RadarHTTPTopBrowserListParamsDateRange52w        RadarHTTPTopBrowserListParamsDateRange = "52w"
	RadarHTTPTopBrowserListParamsDateRange1dControl  RadarHTTPTopBrowserListParamsDateRange = "1dControl"
	RadarHTTPTopBrowserListParamsDateRange2dControl  RadarHTTPTopBrowserListParamsDateRange = "2dControl"
	RadarHTTPTopBrowserListParamsDateRange7dControl  RadarHTTPTopBrowserListParamsDateRange = "7dControl"
	RadarHTTPTopBrowserListParamsDateRange14dControl RadarHTTPTopBrowserListParamsDateRange = "14dControl"
	RadarHTTPTopBrowserListParamsDateRange28dControl RadarHTTPTopBrowserListParamsDateRange = "28dControl"
	RadarHTTPTopBrowserListParamsDateRange12wControl RadarHTTPTopBrowserListParamsDateRange = "12wControl"
	RadarHTTPTopBrowserListParamsDateRange24wControl RadarHTTPTopBrowserListParamsDateRange = "24wControl"
)

type RadarHTTPTopBrowserListParamsDeviceType string

const (
	RadarHTTPTopBrowserListParamsDeviceTypeDesktop RadarHTTPTopBrowserListParamsDeviceType = "DESKTOP"
	RadarHTTPTopBrowserListParamsDeviceTypeMobile  RadarHTTPTopBrowserListParamsDeviceType = "MOBILE"
	RadarHTTPTopBrowserListParamsDeviceTypeOther   RadarHTTPTopBrowserListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopBrowserListParamsFormat string

const (
	RadarHTTPTopBrowserListParamsFormatJson RadarHTTPTopBrowserListParamsFormat = "JSON"
	RadarHTTPTopBrowserListParamsFormatCsv  RadarHTTPTopBrowserListParamsFormat = "CSV"
)

type RadarHTTPTopBrowserListParamsHTTPProtocol string

const (
	RadarHTTPTopBrowserListParamsHTTPProtocolHTTP  RadarHTTPTopBrowserListParamsHTTPProtocol = "HTTP"
	RadarHTTPTopBrowserListParamsHTTPProtocolHTTPs RadarHTTPTopBrowserListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopBrowserListParamsHTTPVersion string

const (
	RadarHTTPTopBrowserListParamsHTTPVersionHttPv1 RadarHTTPTopBrowserListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopBrowserListParamsHTTPVersionHttPv2 RadarHTTPTopBrowserListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopBrowserListParamsHTTPVersionHttPv3 RadarHTTPTopBrowserListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopBrowserListParamsIPVersion string

const (
	RadarHTTPTopBrowserListParamsIPVersionIPv4 RadarHTTPTopBrowserListParamsIPVersion = "IPv4"
	RadarHTTPTopBrowserListParamsIPVersionIPv6 RadarHTTPTopBrowserListParamsIPVersion = "IPv6"
)

type RadarHTTPTopBrowserListParamsO string

const (
	RadarHTTPTopBrowserListParamsOWindows  RadarHTTPTopBrowserListParamsO = "WINDOWS"
	RadarHTTPTopBrowserListParamsOMacosx   RadarHTTPTopBrowserListParamsO = "MACOSX"
	RadarHTTPTopBrowserListParamsOIos      RadarHTTPTopBrowserListParamsO = "IOS"
	RadarHTTPTopBrowserListParamsOAndroid  RadarHTTPTopBrowserListParamsO = "ANDROID"
	RadarHTTPTopBrowserListParamsOChromeos RadarHTTPTopBrowserListParamsO = "CHROMEOS"
	RadarHTTPTopBrowserListParamsOLinux    RadarHTTPTopBrowserListParamsO = "LINUX"
	RadarHTTPTopBrowserListParamsOSmartTv  RadarHTTPTopBrowserListParamsO = "SMART_TV"
)

type RadarHTTPTopBrowserListParamsTlsVersion string

const (
	RadarHTTPTopBrowserListParamsTlsVersionTlSv1_0  RadarHTTPTopBrowserListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopBrowserListParamsTlsVersionTlSv1_1  RadarHTTPTopBrowserListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopBrowserListParamsTlsVersionTlSv1_2  RadarHTTPTopBrowserListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopBrowserListParamsTlsVersionTlSv1_3  RadarHTTPTopBrowserListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopBrowserListParamsTlsVersionTlSvQuic RadarHTTPTopBrowserListParamsTlsVersion = "TLSvQUIC"
)
