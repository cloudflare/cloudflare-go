// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// HTTPAseBrowserFamilyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPAseBrowserFamilyService] method instead.
type HTTPAseBrowserFamilyService struct {
	Options []option.RequestOption
}

// NewHTTPAseBrowserFamilyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPAseBrowserFamilyService(opts ...option.RequestOption) (r *HTTPAseBrowserFamilyService) {
	r = &HTTPAseBrowserFamilyService{}
	r.Options = opts
	return
}

// Retrieves the top autonomous systems, by HTTP requests, of the requested browser
// family.
func (r *HTTPAseBrowserFamilyService) Get(ctx context.Context, browserFamily HTTPAseBrowserFamilyGetParamsBrowserFamily, query HTTPAseBrowserFamilyGetParams, opts ...option.RequestOption) (res *HTTPAseBrowserFamilyGetResponse, err error) {
	var env HTTPAseBrowserFamilyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/browser_family/%v", browserFamily)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseBrowserFamilyGetResponse struct {
	Meta HTTPAseBrowserFamilyGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseBrowserFamilyGetResponseTop0 `json:"top_0,required"`
	JSON httpAseBrowserFamilyGetResponseJSON   `json:"-"`
}

// httpAseBrowserFamilyGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseBrowserFamilyGetResponse]
type httpAseBrowserFamilyGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBrowserFamilyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBrowserFamilyGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBrowserFamilyGetResponseMeta struct {
	DateRange      []HTTPAseBrowserFamilyGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	ConfidenceInfo HTTPAseBrowserFamilyGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpAseBrowserFamilyGetResponseMetaJSON           `json:"-"`
}

// httpAseBrowserFamilyGetResponseMetaJSON contains the JSON metadata for the
// struct [HTTPAseBrowserFamilyGetResponseMeta]
type httpAseBrowserFamilyGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseBrowserFamilyGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBrowserFamilyGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBrowserFamilyGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      httpAseBrowserFamilyGetResponseMetaDateRangeJSON `json:"-"`
}

// httpAseBrowserFamilyGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [HTTPAseBrowserFamilyGetResponseMetaDateRange]
type httpAseBrowserFamilyGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBrowserFamilyGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBrowserFamilyGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBrowserFamilyGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPAseBrowserFamilyGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        httpAseBrowserFamilyGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpAseBrowserFamilyGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPAseBrowserFamilyGetResponseMetaConfidenceInfo]
type httpAseBrowserFamilyGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBrowserFamilyGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBrowserFamilyGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBrowserFamilyGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            httpAseBrowserFamilyGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpAseBrowserFamilyGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [HTTPAseBrowserFamilyGetResponseMetaConfidenceInfoAnnotation]
type httpAseBrowserFamilyGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPAseBrowserFamilyGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBrowserFamilyGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBrowserFamilyGetResponseTop0 struct {
	ClientASN    int64                                   `json:"clientASN,required"`
	ClientAsName string                                  `json:"clientASName,required"`
	Value        string                                  `json:"value,required"`
	JSON         httpAseBrowserFamilyGetResponseTop0JSON `json:"-"`
}

// httpAseBrowserFamilyGetResponseTop0JSON contains the JSON metadata for the
// struct [HTTPAseBrowserFamilyGetResponseTop0]
type httpAseBrowserFamilyGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseBrowserFamilyGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBrowserFamilyGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseBrowserFamilyGetParams struct {
	// Filters results by Autonomous System. Specify one or more Autonomous System
	// Numbers (ASNs) as a comma-separated list. Prefix with `-` to exclude ASNs from
	// results. For example, `-174, 3356` excludes results from AS174, but includes
	// results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseBrowserFamilyGetParamsBotClass] `query:"botClass"`
	// Filters results by continent. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude continents from results. For example, `-EU,NA`
	// excludes results from EU, but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by date range. For example, use `7d` and `7dcontrol` to compare
	// this week with the previous week. Use this parameter or set specific start and
	// end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by device type.
	DeviceType param.Field[[]HTTPAseBrowserFamilyGetParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPAseBrowserFamilyGetParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPAseBrowserFamilyGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPAseBrowserFamilyGetParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPAseBrowserFamilyGetParamsIPVersion] `query:"ipVersion"`
	// Limits the number of objects returned in the response.
	Limit param.Field[int64] `query:"limit"`
	// Filters results by location. Specify a comma-separated list of alpha-2 codes.
	// Prefix with `-` to exclude locations from results. For example, `-US,PT`
	// excludes results from the US, but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Filters results by operating system.
	OS param.Field[[]HTTPAseBrowserFamilyGetParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPAseBrowserFamilyGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPAseBrowserFamilyGetParams]'s query parameters as
// `url.Values`.
func (r HTTPAseBrowserFamilyGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Browser family.
type HTTPAseBrowserFamilyGetParamsBrowserFamily string

const (
	HTTPAseBrowserFamilyGetParamsBrowserFamilyChrome  HTTPAseBrowserFamilyGetParamsBrowserFamily = "CHROME"
	HTTPAseBrowserFamilyGetParamsBrowserFamilyEdge    HTTPAseBrowserFamilyGetParamsBrowserFamily = "EDGE"
	HTTPAseBrowserFamilyGetParamsBrowserFamilyFirefox HTTPAseBrowserFamilyGetParamsBrowserFamily = "FIREFOX"
	HTTPAseBrowserFamilyGetParamsBrowserFamilySafari  HTTPAseBrowserFamilyGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPAseBrowserFamilyGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPAseBrowserFamilyGetParamsBrowserFamilyChrome, HTTPAseBrowserFamilyGetParamsBrowserFamilyEdge, HTTPAseBrowserFamilyGetParamsBrowserFamilyFirefox, HTTPAseBrowserFamilyGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPAseBrowserFamilyGetParamsBotClass string

const (
	HTTPAseBrowserFamilyGetParamsBotClassLikelyAutomated HTTPAseBrowserFamilyGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseBrowserFamilyGetParamsBotClassLikelyHuman     HTTPAseBrowserFamilyGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseBrowserFamilyGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseBrowserFamilyGetParamsBotClassLikelyAutomated, HTTPAseBrowserFamilyGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseBrowserFamilyGetParamsDeviceType string

const (
	HTTPAseBrowserFamilyGetParamsDeviceTypeDesktop HTTPAseBrowserFamilyGetParamsDeviceType = "DESKTOP"
	HTTPAseBrowserFamilyGetParamsDeviceTypeMobile  HTTPAseBrowserFamilyGetParamsDeviceType = "MOBILE"
	HTTPAseBrowserFamilyGetParamsDeviceTypeOther   HTTPAseBrowserFamilyGetParamsDeviceType = "OTHER"
)

func (r HTTPAseBrowserFamilyGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseBrowserFamilyGetParamsDeviceTypeDesktop, HTTPAseBrowserFamilyGetParamsDeviceTypeMobile, HTTPAseBrowserFamilyGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPAseBrowserFamilyGetParamsFormat string

const (
	HTTPAseBrowserFamilyGetParamsFormatJson HTTPAseBrowserFamilyGetParamsFormat = "JSON"
	HTTPAseBrowserFamilyGetParamsFormatCsv  HTTPAseBrowserFamilyGetParamsFormat = "CSV"
)

func (r HTTPAseBrowserFamilyGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseBrowserFamilyGetParamsFormatJson, HTTPAseBrowserFamilyGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseBrowserFamilyGetParamsHTTPProtocol string

const (
	HTTPAseBrowserFamilyGetParamsHTTPProtocolHTTP  HTTPAseBrowserFamilyGetParamsHTTPProtocol = "HTTP"
	HTTPAseBrowserFamilyGetParamsHTTPProtocolHTTPS HTTPAseBrowserFamilyGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseBrowserFamilyGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseBrowserFamilyGetParamsHTTPProtocolHTTP, HTTPAseBrowserFamilyGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseBrowserFamilyGetParamsHTTPVersion string

const (
	HTTPAseBrowserFamilyGetParamsHTTPVersionHttPv1 HTTPAseBrowserFamilyGetParamsHTTPVersion = "HTTPv1"
	HTTPAseBrowserFamilyGetParamsHTTPVersionHttPv2 HTTPAseBrowserFamilyGetParamsHTTPVersion = "HTTPv2"
	HTTPAseBrowserFamilyGetParamsHTTPVersionHttPv3 HTTPAseBrowserFamilyGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseBrowserFamilyGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseBrowserFamilyGetParamsHTTPVersionHttPv1, HTTPAseBrowserFamilyGetParamsHTTPVersionHttPv2, HTTPAseBrowserFamilyGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseBrowserFamilyGetParamsIPVersion string

const (
	HTTPAseBrowserFamilyGetParamsIPVersionIPv4 HTTPAseBrowserFamilyGetParamsIPVersion = "IPv4"
	HTTPAseBrowserFamilyGetParamsIPVersionIPv6 HTTPAseBrowserFamilyGetParamsIPVersion = "IPv6"
)

func (r HTTPAseBrowserFamilyGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseBrowserFamilyGetParamsIPVersionIPv4, HTTPAseBrowserFamilyGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseBrowserFamilyGetParamsOS string

const (
	HTTPAseBrowserFamilyGetParamsOSWindows  HTTPAseBrowserFamilyGetParamsOS = "WINDOWS"
	HTTPAseBrowserFamilyGetParamsOSMacosx   HTTPAseBrowserFamilyGetParamsOS = "MACOSX"
	HTTPAseBrowserFamilyGetParamsOSIos      HTTPAseBrowserFamilyGetParamsOS = "IOS"
	HTTPAseBrowserFamilyGetParamsOSAndroid  HTTPAseBrowserFamilyGetParamsOS = "ANDROID"
	HTTPAseBrowserFamilyGetParamsOSChromeos HTTPAseBrowserFamilyGetParamsOS = "CHROMEOS"
	HTTPAseBrowserFamilyGetParamsOSLinux    HTTPAseBrowserFamilyGetParamsOS = "LINUX"
	HTTPAseBrowserFamilyGetParamsOSSmartTv  HTTPAseBrowserFamilyGetParamsOS = "SMART_TV"
)

func (r HTTPAseBrowserFamilyGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseBrowserFamilyGetParamsOSWindows, HTTPAseBrowserFamilyGetParamsOSMacosx, HTTPAseBrowserFamilyGetParamsOSIos, HTTPAseBrowserFamilyGetParamsOSAndroid, HTTPAseBrowserFamilyGetParamsOSChromeos, HTTPAseBrowserFamilyGetParamsOSLinux, HTTPAseBrowserFamilyGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseBrowserFamilyGetParamsTLSVersion string

const (
	HTTPAseBrowserFamilyGetParamsTLSVersionTlSv1_0  HTTPAseBrowserFamilyGetParamsTLSVersion = "TLSv1_0"
	HTTPAseBrowserFamilyGetParamsTLSVersionTlSv1_1  HTTPAseBrowserFamilyGetParamsTLSVersion = "TLSv1_1"
	HTTPAseBrowserFamilyGetParamsTLSVersionTlSv1_2  HTTPAseBrowserFamilyGetParamsTLSVersion = "TLSv1_2"
	HTTPAseBrowserFamilyGetParamsTLSVersionTlSv1_3  HTTPAseBrowserFamilyGetParamsTLSVersion = "TLSv1_3"
	HTTPAseBrowserFamilyGetParamsTLSVersionTlSvQuic HTTPAseBrowserFamilyGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseBrowserFamilyGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseBrowserFamilyGetParamsTLSVersionTlSv1_0, HTTPAseBrowserFamilyGetParamsTLSVersionTlSv1_1, HTTPAseBrowserFamilyGetParamsTLSVersionTlSv1_2, HTTPAseBrowserFamilyGetParamsTLSVersionTlSv1_3, HTTPAseBrowserFamilyGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseBrowserFamilyGetResponseEnvelope struct {
	Result  HTTPAseBrowserFamilyGetResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    httpAseBrowserFamilyGetResponseEnvelopeJSON `json:"-"`
}

// httpAseBrowserFamilyGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPAseBrowserFamilyGetResponseEnvelope]
type httpAseBrowserFamilyGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBrowserFamilyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBrowserFamilyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
