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

// RadarHTTPTopLocationService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPTopLocationService]
// method instead.
type RadarHTTPTopLocationService struct {
	Options       []option.RequestOption
	BotClasses    *RadarHTTPTopLocationBotClassService
	DeviceTypes   *RadarHTTPTopLocationDeviceTypeService
	HTTPProtocols *RadarHTTPTopLocationHTTPProtocolService
	HTTPVersions  *RadarHTTPTopLocationHTTPVersionService
	IPVersions    *RadarHTTPTopLocationIPVersionService
	Os            *RadarHTTPTopLocationOService
	TlsVersions   *RadarHTTPTopLocationTlsVersionService
}

// NewRadarHTTPTopLocationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPTopLocationService(opts ...option.RequestOption) (r *RadarHTTPTopLocationService) {
	r = &RadarHTTPTopLocationService{}
	r.Options = opts
	r.BotClasses = NewRadarHTTPTopLocationBotClassService(opts...)
	r.DeviceTypes = NewRadarHTTPTopLocationDeviceTypeService(opts...)
	r.HTTPProtocols = NewRadarHTTPTopLocationHTTPProtocolService(opts...)
	r.HTTPVersions = NewRadarHTTPTopLocationHTTPVersionService(opts...)
	r.IPVersions = NewRadarHTTPTopLocationIPVersionService(opts...)
	r.Os = NewRadarHTTPTopLocationOService(opts...)
	r.TlsVersions = NewRadarHTTPTopLocationTlsVersionService(opts...)
	return
}

// Get the top locations by HTTP traffic. Values are a percentage out of the total
// traffic.
func (r *RadarHTTPTopLocationService) List(ctx context.Context, query RadarHTTPTopLocationListParams, opts ...option.RequestOption) (res *RadarHTTPTopLocationListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTopLocationListResponse struct {
	Result  RadarHTTPTopLocationListResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarHTTPTopLocationListResponseJSON   `json:"-"`
}

// radarHTTPTopLocationListResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTopLocationListResponse]
type radarHTTPTopLocationListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationListResponseResult struct {
	Meta RadarHTTPTopLocationListResponseResultMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopLocationListResponseResultTop0 `json:"top_0,required"`
	JSON radarHTTPTopLocationListResponseResultJSON   `json:"-"`
}

// radarHTTPTopLocationListResponseResultJSON contains the JSON metadata for the
// struct [RadarHTTPTopLocationListResponseResult]
type radarHTTPTopLocationListResponseResultJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationListResponseResultMeta struct {
	DateRange      []RadarHTTPTopLocationListResponseResultMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                   `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopLocationListResponseResultMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopLocationListResponseResultMetaJSON           `json:"-"`
}

// radarHTTPTopLocationListResponseResultMetaJSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationListResponseResultMeta]
type radarHTTPTopLocationListResponseResultMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponseResultMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationListResponseResultMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                               `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopLocationListResponseResultMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopLocationListResponseResultMetaDateRangeJSON contains the JSON
// metadata for the struct [RadarHTTPTopLocationListResponseResultMetaDateRange]
type radarHTTPTopLocationListResponseResultMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponseResultMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationListResponseResultMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                                `json:"level"`
	JSON        radarHTTPTopLocationListResponseResultMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopLocationListResponseResultMetaConfidenceInfoJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopLocationListResponseResultMetaConfidenceInfo]
type radarHTTPTopLocationListResponseResultMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponseResultMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                                 `json:"dataSource,required"`
	Description     string                                                                 `json:"description,required"`
	EventType       string                                                                 `json:"eventType,required"`
	IsInstantaneous interface{}                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                                 `json:"linkedUrl"`
	StartTime       time.Time                                                              `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [RadarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotation]
type radarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopLocationListResponseResultMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationListResponseResultTop0 struct {
	ClientCountryAlpha2 string                                         `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                         `json:"clientCountryName,required"`
	Value               string                                         `json:"value,required"`
	JSON                radarHTTPTopLocationListResponseResultTop0JSON `json:"-"`
}

// radarHTTPTopLocationListResponseResultTop0JSON contains the JSON metadata for
// the struct [RadarHTTPTopLocationListResponseResultTop0]
type radarHTTPTopLocationListResponseResultTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *RadarHTTPTopLocationListResponseResultTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopLocationListParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopLocationListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopLocationListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopLocationListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopLocationListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopLocationListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopLocationListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopLocationListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTopLocationListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTopLocationListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopLocationListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopLocationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopLocationListParamsBotClass string

const (
	RadarHTTPTopLocationListParamsBotClassLikelyAutomated RadarHTTPTopLocationListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopLocationListParamsBotClassLikelyHuman     RadarHTTPTopLocationListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopLocationListParamsDateRange string

const (
	RadarHTTPTopLocationListParamsDateRange1d         RadarHTTPTopLocationListParamsDateRange = "1d"
	RadarHTTPTopLocationListParamsDateRange2d         RadarHTTPTopLocationListParamsDateRange = "2d"
	RadarHTTPTopLocationListParamsDateRange7d         RadarHTTPTopLocationListParamsDateRange = "7d"
	RadarHTTPTopLocationListParamsDateRange14d        RadarHTTPTopLocationListParamsDateRange = "14d"
	RadarHTTPTopLocationListParamsDateRange28d        RadarHTTPTopLocationListParamsDateRange = "28d"
	RadarHTTPTopLocationListParamsDateRange12w        RadarHTTPTopLocationListParamsDateRange = "12w"
	RadarHTTPTopLocationListParamsDateRange24w        RadarHTTPTopLocationListParamsDateRange = "24w"
	RadarHTTPTopLocationListParamsDateRange52w        RadarHTTPTopLocationListParamsDateRange = "52w"
	RadarHTTPTopLocationListParamsDateRange1dControl  RadarHTTPTopLocationListParamsDateRange = "1dControl"
	RadarHTTPTopLocationListParamsDateRange2dControl  RadarHTTPTopLocationListParamsDateRange = "2dControl"
	RadarHTTPTopLocationListParamsDateRange7dControl  RadarHTTPTopLocationListParamsDateRange = "7dControl"
	RadarHTTPTopLocationListParamsDateRange14dControl RadarHTTPTopLocationListParamsDateRange = "14dControl"
	RadarHTTPTopLocationListParamsDateRange28dControl RadarHTTPTopLocationListParamsDateRange = "28dControl"
	RadarHTTPTopLocationListParamsDateRange12wControl RadarHTTPTopLocationListParamsDateRange = "12wControl"
	RadarHTTPTopLocationListParamsDateRange24wControl RadarHTTPTopLocationListParamsDateRange = "24wControl"
)

type RadarHTTPTopLocationListParamsDeviceType string

const (
	RadarHTTPTopLocationListParamsDeviceTypeDesktop RadarHTTPTopLocationListParamsDeviceType = "DESKTOP"
	RadarHTTPTopLocationListParamsDeviceTypeMobile  RadarHTTPTopLocationListParamsDeviceType = "MOBILE"
	RadarHTTPTopLocationListParamsDeviceTypeOther   RadarHTTPTopLocationListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopLocationListParamsFormat string

const (
	RadarHTTPTopLocationListParamsFormatJson RadarHTTPTopLocationListParamsFormat = "JSON"
	RadarHTTPTopLocationListParamsFormatCsv  RadarHTTPTopLocationListParamsFormat = "CSV"
)

type RadarHTTPTopLocationListParamsHTTPProtocol string

const (
	RadarHTTPTopLocationListParamsHTTPProtocolHTTP  RadarHTTPTopLocationListParamsHTTPProtocol = "HTTP"
	RadarHTTPTopLocationListParamsHTTPProtocolHTTPs RadarHTTPTopLocationListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopLocationListParamsHTTPVersion string

const (
	RadarHTTPTopLocationListParamsHTTPVersionHttPv1 RadarHTTPTopLocationListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopLocationListParamsHTTPVersionHttPv2 RadarHTTPTopLocationListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopLocationListParamsHTTPVersionHttPv3 RadarHTTPTopLocationListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopLocationListParamsIPVersion string

const (
	RadarHTTPTopLocationListParamsIPVersionIPv4 RadarHTTPTopLocationListParamsIPVersion = "IPv4"
	RadarHTTPTopLocationListParamsIPVersionIPv6 RadarHTTPTopLocationListParamsIPVersion = "IPv6"
)

type RadarHTTPTopLocationListParamsO string

const (
	RadarHTTPTopLocationListParamsOWindows  RadarHTTPTopLocationListParamsO = "WINDOWS"
	RadarHTTPTopLocationListParamsOMacosx   RadarHTTPTopLocationListParamsO = "MACOSX"
	RadarHTTPTopLocationListParamsOIos      RadarHTTPTopLocationListParamsO = "IOS"
	RadarHTTPTopLocationListParamsOAndroid  RadarHTTPTopLocationListParamsO = "ANDROID"
	RadarHTTPTopLocationListParamsOChromeos RadarHTTPTopLocationListParamsO = "CHROMEOS"
	RadarHTTPTopLocationListParamsOLinux    RadarHTTPTopLocationListParamsO = "LINUX"
	RadarHTTPTopLocationListParamsOSmartTv  RadarHTTPTopLocationListParamsO = "SMART_TV"
)

type RadarHTTPTopLocationListParamsTlsVersion string

const (
	RadarHTTPTopLocationListParamsTlsVersionTlSv1_0  RadarHTTPTopLocationListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTopLocationListParamsTlsVersionTlSv1_1  RadarHTTPTopLocationListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTopLocationListParamsTlsVersionTlSv1_2  RadarHTTPTopLocationListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTopLocationListParamsTlsVersionTlSv1_3  RadarHTTPTopLocationListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTopLocationListParamsTlsVersionTlSvQuic RadarHTTPTopLocationListParamsTlsVersion = "TLSvQUIC"
)
