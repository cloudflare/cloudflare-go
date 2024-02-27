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

// RadarHTTPTopService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPTopService] method
// instead.
type RadarHTTPTopService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTopService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarHTTPTopService(opts ...option.RequestOption) (r *RadarHTTPTopService) {
	r = &RadarHTTPTopService{}
	r.Options = opts
	return
}

// Get the top user agents aggregated in families by HTTP traffic. Values are a
// percentage out of the total traffic.
func (r *RadarHTTPTopService) BrowserFamilies(ctx context.Context, query RadarHTTPTopBrowserFamiliesParams, opts ...option.RequestOption) (res *RadarHTTPTopBrowserFamiliesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPTopBrowserFamiliesResponseEnvelope
	path := "radar/http/top/browser_families"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the top user agents by HTTP traffic. Values are a percentage out of the
// total traffic.
func (r *RadarHTTPTopService) Browsers(ctx context.Context, query RadarHTTPTopBrowsersParams, opts ...option.RequestOption) (res *RadarHTTPTopBrowsersResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RadarHTTPTopBrowsersResponseEnvelope
	path := "radar/http/top/browsers"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RadarHTTPTopBrowserFamiliesResponse struct {
	Meta RadarHTTPTopBrowserFamiliesResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopBrowserFamiliesResponseTop0 `json:"top_0,required"`
	JSON radarHTTPTopBrowserFamiliesResponseJSON   `json:"-"`
}

// radarHTTPTopBrowserFamiliesResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTopBrowserFamiliesResponse]
type radarHTTPTopBrowserFamiliesResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserFamiliesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamiliesResponseMeta struct {
	DateRange      []RadarHTTPTopBrowserFamiliesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopBrowserFamiliesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopBrowserFamiliesResponseMetaJSON           `json:"-"`
}

// radarHTTPTopBrowserFamiliesResponseMetaJSON contains the JSON metadata for the
// struct [RadarHTTPTopBrowserFamiliesResponseMeta]
type radarHTTPTopBrowserFamiliesResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserFamiliesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamiliesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                            `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopBrowserFamiliesResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopBrowserFamiliesResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [RadarHTTPTopBrowserFamiliesResponseMetaDateRange]
type radarHTTPTopBrowserFamiliesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserFamiliesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamiliesResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopBrowserFamiliesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                             `json:"level"`
	JSON        radarHTTPTopBrowserFamiliesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopBrowserFamiliesResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [RadarHTTPTopBrowserFamiliesResponseMetaConfidenceInfo]
type radarHTTPTopBrowserFamiliesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserFamiliesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamiliesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                              `json:"dataSource,required"`
	Description     string                                                              `json:"description,required"`
	EventType       string                                                              `json:"eventType,required"`
	IsInstantaneous interface{}                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                           `json:"endTime" format:"date-time"`
	LinkedURL       string                                                              `json:"linkedUrl"`
	StartTime       time.Time                                                           `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopBrowserFamiliesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopBrowserFamiliesResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [RadarHTTPTopBrowserFamiliesResponseMetaConfidenceInfoAnnotation]
type radarHTTPTopBrowserFamiliesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopBrowserFamiliesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamiliesResponseTop0 struct {
	Name  string                                      `json:"name,required"`
	Value string                                      `json:"value,required"`
	JSON  radarHTTPTopBrowserFamiliesResponseTop0JSON `json:"-"`
}

// radarHTTPTopBrowserFamiliesResponseTop0JSON contains the JSON metadata for the
// struct [RadarHTTPTopBrowserFamiliesResponseTop0]
type radarHTTPTopBrowserFamiliesResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserFamiliesResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowsersResponse struct {
	Meta RadarHTTPTopBrowsersResponseMeta   `json:"meta,required"`
	Top0 []RadarHTTPTopBrowsersResponseTop0 `json:"top_0,required"`
	JSON radarHTTPTopBrowsersResponseJSON   `json:"-"`
}

// radarHTTPTopBrowsersResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTopBrowsersResponse]
type radarHTTPTopBrowsersResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowsersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowsersResponseMeta struct {
	DateRange      []RadarHTTPTopBrowsersResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                         `json:"lastUpdated,required"`
	ConfidenceInfo RadarHTTPTopBrowsersResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           radarHTTPTopBrowsersResponseMetaJSON           `json:"-"`
}

// radarHTTPTopBrowsersResponseMetaJSON contains the JSON metadata for the struct
// [RadarHTTPTopBrowsersResponseMeta]
type radarHTTPTopBrowsersResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *RadarHTTPTopBrowsersResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowsersResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      radarHTTPTopBrowsersResponseMetaDateRangeJSON `json:"-"`
}

// radarHTTPTopBrowsersResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [RadarHTTPTopBrowsersResponseMetaDateRange]
type radarHTTPTopBrowsersResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowsersResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowsersResponseMetaConfidenceInfo struct {
	Annotations []RadarHTTPTopBrowsersResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                      `json:"level"`
	JSON        radarHTTPTopBrowsersResponseMetaConfidenceInfoJSON         `json:"-"`
}

// radarHTTPTopBrowsersResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [RadarHTTPTopBrowsersResponseMetaConfidenceInfo]
type radarHTTPTopBrowsersResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowsersResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowsersResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                       `json:"dataSource,required"`
	Description     string                                                       `json:"description,required"`
	EventType       string                                                       `json:"eventType,required"`
	IsInstantaneous interface{}                                                  `json:"isInstantaneous,required"`
	EndTime         time.Time                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                    `json:"startTime" format:"date-time"`
	JSON            radarHTTPTopBrowsersResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// radarHTTPTopBrowsersResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [RadarHTTPTopBrowsersResponseMetaConfidenceInfoAnnotation]
type radarHTTPTopBrowsersResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *RadarHTTPTopBrowsersResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowsersResponseTop0 struct {
	Name  string                               `json:"name,required"`
	Value string                               `json:"value,required"`
	JSON  radarHTTPTopBrowsersResponseTop0JSON `json:"-"`
}

// radarHTTPTopBrowsersResponseTop0JSON contains the JSON metadata for the struct
// [RadarHTTPTopBrowsersResponseTop0]
type radarHTTPTopBrowsersResponseTop0JSON struct {
	Name        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowsersResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowserFamiliesParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopBrowserFamiliesParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopBrowserFamiliesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopBrowserFamiliesParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopBrowserFamiliesParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopBrowserFamiliesParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopBrowserFamiliesParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopBrowserFamiliesParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPTopBrowserFamiliesParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPTopBrowserFamiliesParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopBrowserFamiliesParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopBrowserFamiliesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopBrowserFamiliesParamsBotClass string

const (
	RadarHTTPTopBrowserFamiliesParamsBotClassLikelyAutomated RadarHTTPTopBrowserFamiliesParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopBrowserFamiliesParamsBotClassLikelyHuman     RadarHTTPTopBrowserFamiliesParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopBrowserFamiliesParamsDateRange string

const (
	RadarHTTPTopBrowserFamiliesParamsDateRange1d         RadarHTTPTopBrowserFamiliesParamsDateRange = "1d"
	RadarHTTPTopBrowserFamiliesParamsDateRange2d         RadarHTTPTopBrowserFamiliesParamsDateRange = "2d"
	RadarHTTPTopBrowserFamiliesParamsDateRange7d         RadarHTTPTopBrowserFamiliesParamsDateRange = "7d"
	RadarHTTPTopBrowserFamiliesParamsDateRange14d        RadarHTTPTopBrowserFamiliesParamsDateRange = "14d"
	RadarHTTPTopBrowserFamiliesParamsDateRange28d        RadarHTTPTopBrowserFamiliesParamsDateRange = "28d"
	RadarHTTPTopBrowserFamiliesParamsDateRange12w        RadarHTTPTopBrowserFamiliesParamsDateRange = "12w"
	RadarHTTPTopBrowserFamiliesParamsDateRange24w        RadarHTTPTopBrowserFamiliesParamsDateRange = "24w"
	RadarHTTPTopBrowserFamiliesParamsDateRange52w        RadarHTTPTopBrowserFamiliesParamsDateRange = "52w"
	RadarHTTPTopBrowserFamiliesParamsDateRange1dControl  RadarHTTPTopBrowserFamiliesParamsDateRange = "1dControl"
	RadarHTTPTopBrowserFamiliesParamsDateRange2dControl  RadarHTTPTopBrowserFamiliesParamsDateRange = "2dControl"
	RadarHTTPTopBrowserFamiliesParamsDateRange7dControl  RadarHTTPTopBrowserFamiliesParamsDateRange = "7dControl"
	RadarHTTPTopBrowserFamiliesParamsDateRange14dControl RadarHTTPTopBrowserFamiliesParamsDateRange = "14dControl"
	RadarHTTPTopBrowserFamiliesParamsDateRange28dControl RadarHTTPTopBrowserFamiliesParamsDateRange = "28dControl"
	RadarHTTPTopBrowserFamiliesParamsDateRange12wControl RadarHTTPTopBrowserFamiliesParamsDateRange = "12wControl"
	RadarHTTPTopBrowserFamiliesParamsDateRange24wControl RadarHTTPTopBrowserFamiliesParamsDateRange = "24wControl"
)

type RadarHTTPTopBrowserFamiliesParamsDeviceType string

const (
	RadarHTTPTopBrowserFamiliesParamsDeviceTypeDesktop RadarHTTPTopBrowserFamiliesParamsDeviceType = "DESKTOP"
	RadarHTTPTopBrowserFamiliesParamsDeviceTypeMobile  RadarHTTPTopBrowserFamiliesParamsDeviceType = "MOBILE"
	RadarHTTPTopBrowserFamiliesParamsDeviceTypeOther   RadarHTTPTopBrowserFamiliesParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopBrowserFamiliesParamsFormat string

const (
	RadarHTTPTopBrowserFamiliesParamsFormatJson RadarHTTPTopBrowserFamiliesParamsFormat = "JSON"
	RadarHTTPTopBrowserFamiliesParamsFormatCsv  RadarHTTPTopBrowserFamiliesParamsFormat = "CSV"
)

type RadarHTTPTopBrowserFamiliesParamsHTTPProtocol string

const (
	RadarHTTPTopBrowserFamiliesParamsHTTPProtocolHTTP  RadarHTTPTopBrowserFamiliesParamsHTTPProtocol = "HTTP"
	RadarHTTPTopBrowserFamiliesParamsHTTPProtocolHTTPS RadarHTTPTopBrowserFamiliesParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopBrowserFamiliesParamsHTTPVersion string

const (
	RadarHTTPTopBrowserFamiliesParamsHTTPVersionHttPv1 RadarHTTPTopBrowserFamiliesParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopBrowserFamiliesParamsHTTPVersionHttPv2 RadarHTTPTopBrowserFamiliesParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopBrowserFamiliesParamsHTTPVersionHttPv3 RadarHTTPTopBrowserFamiliesParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopBrowserFamiliesParamsIPVersion string

const (
	RadarHTTPTopBrowserFamiliesParamsIPVersionIPv4 RadarHTTPTopBrowserFamiliesParamsIPVersion = "IPv4"
	RadarHTTPTopBrowserFamiliesParamsIPVersionIPv6 RadarHTTPTopBrowserFamiliesParamsIPVersion = "IPv6"
)

type RadarHTTPTopBrowserFamiliesParamsOS string

const (
	RadarHTTPTopBrowserFamiliesParamsOSWindows  RadarHTTPTopBrowserFamiliesParamsOS = "WINDOWS"
	RadarHTTPTopBrowserFamiliesParamsOSMacosx   RadarHTTPTopBrowserFamiliesParamsOS = "MACOSX"
	RadarHTTPTopBrowserFamiliesParamsOSIos      RadarHTTPTopBrowserFamiliesParamsOS = "IOS"
	RadarHTTPTopBrowserFamiliesParamsOSAndroid  RadarHTTPTopBrowserFamiliesParamsOS = "ANDROID"
	RadarHTTPTopBrowserFamiliesParamsOSChromeos RadarHTTPTopBrowserFamiliesParamsOS = "CHROMEOS"
	RadarHTTPTopBrowserFamiliesParamsOSLinux    RadarHTTPTopBrowserFamiliesParamsOS = "LINUX"
	RadarHTTPTopBrowserFamiliesParamsOSSmartTv  RadarHTTPTopBrowserFamiliesParamsOS = "SMART_TV"
)

type RadarHTTPTopBrowserFamiliesParamsTLSVersion string

const (
	RadarHTTPTopBrowserFamiliesParamsTLSVersionTlSv1_0  RadarHTTPTopBrowserFamiliesParamsTLSVersion = "TLSv1_0"
	RadarHTTPTopBrowserFamiliesParamsTLSVersionTlSv1_1  RadarHTTPTopBrowserFamiliesParamsTLSVersion = "TLSv1_1"
	RadarHTTPTopBrowserFamiliesParamsTLSVersionTlSv1_2  RadarHTTPTopBrowserFamiliesParamsTLSVersion = "TLSv1_2"
	RadarHTTPTopBrowserFamiliesParamsTLSVersionTlSv1_3  RadarHTTPTopBrowserFamiliesParamsTLSVersion = "TLSv1_3"
	RadarHTTPTopBrowserFamiliesParamsTLSVersionTlSvQuic RadarHTTPTopBrowserFamiliesParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPTopBrowserFamiliesResponseEnvelope struct {
	Result  RadarHTTPTopBrowserFamiliesResponse             `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarHTTPTopBrowserFamiliesResponseEnvelopeJSON `json:"-"`
}

// radarHTTPTopBrowserFamiliesResponseEnvelopeJSON contains the JSON metadata for
// the struct [RadarHTTPTopBrowserFamiliesResponseEnvelope]
type radarHTTPTopBrowserFamiliesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowserFamiliesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTopBrowsersParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTopBrowsersParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTopBrowsersParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTopBrowsersParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTopBrowsersParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTopBrowsersParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTopBrowsersParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTopBrowsersParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]RadarHTTPTopBrowsersParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPTopBrowsersParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTopBrowsersParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTopBrowsersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RadarHTTPTopBrowsersParamsBotClass string

const (
	RadarHTTPTopBrowsersParamsBotClassLikelyAutomated RadarHTTPTopBrowsersParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTopBrowsersParamsBotClassLikelyHuman     RadarHTTPTopBrowsersParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTopBrowsersParamsDateRange string

const (
	RadarHTTPTopBrowsersParamsDateRange1d         RadarHTTPTopBrowsersParamsDateRange = "1d"
	RadarHTTPTopBrowsersParamsDateRange2d         RadarHTTPTopBrowsersParamsDateRange = "2d"
	RadarHTTPTopBrowsersParamsDateRange7d         RadarHTTPTopBrowsersParamsDateRange = "7d"
	RadarHTTPTopBrowsersParamsDateRange14d        RadarHTTPTopBrowsersParamsDateRange = "14d"
	RadarHTTPTopBrowsersParamsDateRange28d        RadarHTTPTopBrowsersParamsDateRange = "28d"
	RadarHTTPTopBrowsersParamsDateRange12w        RadarHTTPTopBrowsersParamsDateRange = "12w"
	RadarHTTPTopBrowsersParamsDateRange24w        RadarHTTPTopBrowsersParamsDateRange = "24w"
	RadarHTTPTopBrowsersParamsDateRange52w        RadarHTTPTopBrowsersParamsDateRange = "52w"
	RadarHTTPTopBrowsersParamsDateRange1dControl  RadarHTTPTopBrowsersParamsDateRange = "1dControl"
	RadarHTTPTopBrowsersParamsDateRange2dControl  RadarHTTPTopBrowsersParamsDateRange = "2dControl"
	RadarHTTPTopBrowsersParamsDateRange7dControl  RadarHTTPTopBrowsersParamsDateRange = "7dControl"
	RadarHTTPTopBrowsersParamsDateRange14dControl RadarHTTPTopBrowsersParamsDateRange = "14dControl"
	RadarHTTPTopBrowsersParamsDateRange28dControl RadarHTTPTopBrowsersParamsDateRange = "28dControl"
	RadarHTTPTopBrowsersParamsDateRange12wControl RadarHTTPTopBrowsersParamsDateRange = "12wControl"
	RadarHTTPTopBrowsersParamsDateRange24wControl RadarHTTPTopBrowsersParamsDateRange = "24wControl"
)

type RadarHTTPTopBrowsersParamsDeviceType string

const (
	RadarHTTPTopBrowsersParamsDeviceTypeDesktop RadarHTTPTopBrowsersParamsDeviceType = "DESKTOP"
	RadarHTTPTopBrowsersParamsDeviceTypeMobile  RadarHTTPTopBrowsersParamsDeviceType = "MOBILE"
	RadarHTTPTopBrowsersParamsDeviceTypeOther   RadarHTTPTopBrowsersParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTopBrowsersParamsFormat string

const (
	RadarHTTPTopBrowsersParamsFormatJson RadarHTTPTopBrowsersParamsFormat = "JSON"
	RadarHTTPTopBrowsersParamsFormatCsv  RadarHTTPTopBrowsersParamsFormat = "CSV"
)

type RadarHTTPTopBrowsersParamsHTTPProtocol string

const (
	RadarHTTPTopBrowsersParamsHTTPProtocolHTTP  RadarHTTPTopBrowsersParamsHTTPProtocol = "HTTP"
	RadarHTTPTopBrowsersParamsHTTPProtocolHTTPS RadarHTTPTopBrowsersParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTopBrowsersParamsHTTPVersion string

const (
	RadarHTTPTopBrowsersParamsHTTPVersionHttPv1 RadarHTTPTopBrowsersParamsHTTPVersion = "HTTPv1"
	RadarHTTPTopBrowsersParamsHTTPVersionHttPv2 RadarHTTPTopBrowsersParamsHTTPVersion = "HTTPv2"
	RadarHTTPTopBrowsersParamsHTTPVersionHttPv3 RadarHTTPTopBrowsersParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTopBrowsersParamsIPVersion string

const (
	RadarHTTPTopBrowsersParamsIPVersionIPv4 RadarHTTPTopBrowsersParamsIPVersion = "IPv4"
	RadarHTTPTopBrowsersParamsIPVersionIPv6 RadarHTTPTopBrowsersParamsIPVersion = "IPv6"
)

type RadarHTTPTopBrowsersParamsOS string

const (
	RadarHTTPTopBrowsersParamsOSWindows  RadarHTTPTopBrowsersParamsOS = "WINDOWS"
	RadarHTTPTopBrowsersParamsOSMacosx   RadarHTTPTopBrowsersParamsOS = "MACOSX"
	RadarHTTPTopBrowsersParamsOSIos      RadarHTTPTopBrowsersParamsOS = "IOS"
	RadarHTTPTopBrowsersParamsOSAndroid  RadarHTTPTopBrowsersParamsOS = "ANDROID"
	RadarHTTPTopBrowsersParamsOSChromeos RadarHTTPTopBrowsersParamsOS = "CHROMEOS"
	RadarHTTPTopBrowsersParamsOSLinux    RadarHTTPTopBrowsersParamsOS = "LINUX"
	RadarHTTPTopBrowsersParamsOSSmartTv  RadarHTTPTopBrowsersParamsOS = "SMART_TV"
)

type RadarHTTPTopBrowsersParamsTLSVersion string

const (
	RadarHTTPTopBrowsersParamsTLSVersionTlSv1_0  RadarHTTPTopBrowsersParamsTLSVersion = "TLSv1_0"
	RadarHTTPTopBrowsersParamsTLSVersionTlSv1_1  RadarHTTPTopBrowsersParamsTLSVersion = "TLSv1_1"
	RadarHTTPTopBrowsersParamsTLSVersionTlSv1_2  RadarHTTPTopBrowsersParamsTLSVersion = "TLSv1_2"
	RadarHTTPTopBrowsersParamsTLSVersionTlSv1_3  RadarHTTPTopBrowsersParamsTLSVersion = "TLSv1_3"
	RadarHTTPTopBrowsersParamsTLSVersionTlSvQuic RadarHTTPTopBrowsersParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPTopBrowsersResponseEnvelope struct {
	Result  RadarHTTPTopBrowsersResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    radarHTTPTopBrowsersResponseEnvelopeJSON `json:"-"`
}

// radarHTTPTopBrowsersResponseEnvelopeJSON contains the JSON metadata for the
// struct [RadarHTTPTopBrowsersResponseEnvelope]
type radarHTTPTopBrowsersResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTopBrowsersResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
