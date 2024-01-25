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

// RadarHTTPTopBrowserFamilyService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTopBrowserFamilyService] method instead.
type RadarHTTPTopBrowserFamilyService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopBrowserFamilyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTopBrowserFamilyService(opts ...option.RequestOption) (r *RadarHTTPTopBrowserFamilyService) {
	r = &RadarHTTPTopBrowserFamilyService{}
	r.Options = opts
	return
}

// Get the top user agents aggregated in families by HTTP traffic. Values are a
// percentage out of the total traffic.
func (r *RadarHTTPTopBrowserFamilyService) List(ctx context.Context, query RadarHTTPTopBrowserFamilyListParams, opts ...option.RequestOption) (res *RadarHTTPTopBrowserFamilyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/browser_families"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopBrowserFamilyListResponse struct {
	Result  RadarHTTPTopBrowserFamilyListResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarHTTPTopBrowserFamilyListResponseJSON   `json:"-"`
}

// radarHTTPTopBrowserFamilyListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopBrowserFamilyListResponse]
type radarHTTPTopBrowserFamilyListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserFamilyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamilyListResponseResult struct {
	Meta RadarHTTPTopBrowserFamilyListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopBrowserFamilyListResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopBrowserFamilyListResponseResultJSON   `json:"-"`
}

// radarHTTPTopBrowserFamilyListResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTopBrowserFamilyListResponseResult]
type radarHTTPTopBrowserFamilyListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserFamilyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamilyListResponseResultMeta struct {
	DateRange      []RadarHTTPTopBrowserFamilyListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                        `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopBrowserFamilyListResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopBrowserFamilyListResponseResultMetaJSON contains the JSON metadata
// for the struct [RadarHTTPTopBrowserFamilyListResponseResultMeta]
type radarHTTPTopBrowserFamilyListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserFamilyListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamilyListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                                    `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopBrowserFamilyListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopBrowserFamilyListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopBrowserFamilyListResponseResultMetaDateRange]
type radarHTTPTopBrowserFamilyListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserFamilyListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                     `json:"level"`
	JSON        radarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfoJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfo]
type radarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                      `json:"dataSource,required"`
	Description     string                                                                      `json:"description,required"`
	EventType       string                                                                      `json:"eventType,required"`
	IsInstantaneous interface{}                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                                   `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfoAnnotationJSON
// contains the JSON metadata for the struct
// [RadarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopBrowserFamilyListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamilyListResponseResultTop0 struct {
	Name  string                                              `json:"name,required"`
	Value string                                              `json:"value,required"`
	JSON  radarHTTPTopBrowserFamilyListResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopBrowserFamilyListResponseResultTop0JSON contains the JSON metadata
// for the struct [RadarHTTPTopBrowserFamilyListResponseResultTop0]
type radarHTTPTopBrowserFamilyListResponseResultTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserFamilyListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamilyListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopBrowserFamilyListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopBrowserFamilyListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopBrowserFamilyListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopBrowserFamilyListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopBrowserFamilyListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopBrowserFamilyListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopBrowserFamilyListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopBrowserFamilyListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopBrowserFamilyListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopBrowserFamilyListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopBrowserFamilyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopBrowserFamilyListParamsBotClass string

const (
	RadarHTTPTopBrowserFamilyListParamsBotClassLikelyAutomated RadarHTTPTopBrowserFamilyListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopBrowserFamilyListParamsBotClassLikelyHuman     RadarHTTPTopBrowserFamilyListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopBrowserFamilyListParamsDateRange string

const (
	RadarHTTPTopBrowserFamilyListParamsDateRange1d         RadarHTTPTopBrowserFamilyListParamsDateRange = "1d"
	RadarHTTPTopBrowserFamilyListParamsDateRange2d         RadarHTTPTopBrowserFamilyListParamsDateRange = "2d"
	RadarHTTPTopBrowserFamilyListParamsDateRange7d         RadarHTTPTopBrowserFamilyListParamsDateRange = "7d"
	RadarHTTPTopBrowserFamilyListParamsDateRange14d        RadarHTTPTopBrowserFamilyListParamsDateRange = "14d"
	RadarHTTPTopBrowserFamilyListParamsDateRange28d        RadarHTTPTopBrowserFamilyListParamsDateRange = "28d"
	RadarHTTPTopBrowserFamilyListParamsDateRange12w        RadarHTTPTopBrowserFamilyListParamsDateRange = "12w"
	RadarHTTPTopBrowserFamilyListParamsDateRange24w        RadarHTTPTopBrowserFamilyListParamsDateRange = "24w"
	RadarHTTPTopBrowserFamilyListParamsDateRange52w        RadarHTTPTopBrowserFamilyListParamsDateRange = "52w"
	RadarHTTPTopBrowserFamilyListParamsDateRange1dControl  RadarHTTPTopBrowserFamilyListParamsDateRange = "1dControl"
	RadarHTTPTopBrowserFamilyListParamsDateRange2dControl  RadarHTTPTopBrowserFamilyListParamsDateRange = "2dControl"
	RadarHTTPTopBrowserFamilyListParamsDateRange7dControl  RadarHTTPTopBrowserFamilyListParamsDateRange = "7dControl"
	RadarHTTPTopBrowserFamilyListParamsDateRange14dControl RadarHTTPTopBrowserFamilyListParamsDateRange = "14dControl"
	RadarHTTPTopBrowserFamilyListParamsDateRange28dControl RadarHTTPTopBrowserFamilyListParamsDateRange = "28dControl"
	RadarHTTPTopBrowserFamilyListParamsDateRange12wControl RadarHTTPTopBrowserFamilyListParamsDateRange = "12wControl"
	RadarHTTPTopBrowserFamilyListParamsDateRange24wControl RadarHTTPTopBrowserFamilyListParamsDateRange = "24wControl"
)

type RadarHTTPTopBrowserFamilyListParamsDeviceType string

const (
	RadarHTTPTopBrowserFamilyListParamsDeviceTypeDesktop RadarHTTPTopBrowserFamilyListParamsDeviceType = "DESKTOP"
	RadarHTTPTopBrowserFamilyListParamsDeviceTypeMobile  RadarHTTPTopBrowserFamilyListParamsDeviceType = "MOBILE"
	RadarHTTPTopBrowserFamilyListParamsDeviceTypeOther   RadarHTTPTopBrowserFamilyListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopBrowserFamilyListParamsFormat string

const (
	RadarHTTPTopBrowserFamilyListParamsFormatJson RadarHTTPTopBrowserFamilyListParamsFormat = "JSON"
	RadarHTTPTopBrowserFamilyListParamsFormatCsv  RadarHTTPTopBrowserFamilyListParamsFormat = "CSV"
)

type RadarHTTPTopBrowserFamilyListParamsHTTPProtocol string

const (
	RadarHTTPTopBrowserFamilyListParamsHTTPProtocolHTTP  RadarHTTPTopBrowserFamilyListParamsHTTPProtocol = "HTTP"
	RadarHTTPTopBrowserFamilyListParamsHTTPProtocolHTTPs RadarHTTPTopBrowserFamilyListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopBrowserFamilyListParamsHTTPVersion string

const (
	RadarHTTPTopBrowserFamilyListParamsHTTPVersionHttPv1 RadarHTTPTopBrowserFamilyListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopBrowserFamilyListParamsHTTPVersionHttPv2 RadarHTTPTopBrowserFamilyListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopBrowserFamilyListParamsHTTPVersionHttPv3 RadarHTTPTopBrowserFamilyListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopBrowserFamilyListParamsIPVersion string

const (
	RadarHTTPTopBrowserFamilyListParamsIPVersionIPv4 RadarHTTPTopBrowserFamilyListParamsIPVersion = "IPv4"
	RadarHTTPTopBrowserFamilyListParamsIPVersionIPv6 RadarHTTPTopBrowserFamilyListParamsIPVersion = "IPv6"
)

type RadarHTTPTopBrowserFamilyListParamsO string

const (
	RadarHTTPTopBrowserFamilyListParamsOWindows  RadarHTTPTopBrowserFamilyListParamsO = "WINDOWS"
	RadarHTTPTopBrowserFamilyListParamsOMacosx   RadarHTTPTopBrowserFamilyListParamsO = "MACOSX"
	RadarHTTPTopBrowserFamilyListParamsOIos      RadarHTTPTopBrowserFamilyListParamsO = "IOS"
	RadarHTTPTopBrowserFamilyListParamsOAndroid  RadarHTTPTopBrowserFamilyListParamsO = "ANDROID"
	RadarHTTPTopBrowserFamilyListParamsOChromeos RadarHTTPTopBrowserFamilyListParamsO = "CHROMEOS"
	RadarHTTPTopBrowserFamilyListParamsOLinux    RadarHTTPTopBrowserFamilyListParamsO = "LINUX"
	RadarHTTPTopBrowserFamilyListParamsOSmartTv  RadarHTTPTopBrowserFamilyListParamsO = "SMART_TV"
)

type RadarHTTPTopBrowserFamilyListParamsTlsVersion string

const (
	RadarHTTPTopBrowserFamilyListParamsTlsVersionTlSv1_0  RadarHTTPTopBrowserFamilyListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopBrowserFamilyListParamsTlsVersionTlSv1_1  RadarHTTPTopBrowserFamilyListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopBrowserFamilyListParamsTlsVersionTlSv1_2  RadarHTTPTopBrowserFamilyListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopBrowserFamilyListParamsTlsVersionTlSv1_3  RadarHTTPTopBrowserFamilyListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopBrowserFamilyListParamsTlsVersionTlSvQuic RadarHTTPTopBrowserFamilyListParamsTlsVersion = "TLSvQUIC"
)
