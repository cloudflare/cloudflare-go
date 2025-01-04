// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// HTTPLocationService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationService] method instead.
type HTTPLocationService struct {
	Options       []option.RequestOption
	BotClass      *HTTPLocationBotClassService
	DeviceType    *HTTPLocationDeviceTypeService
	HTTPProtocol  *HTTPLocationHTTPProtocolService
	HTTPMethod    *HTTPLocationHTTPMethodService
	IPVersion     *HTTPLocationIPVersionService
	OS            *HTTPLocationOSService
	TLSVersion    *HTTPLocationTLSVersionService
	BrowserFamily *HTTPLocationBrowserFamilyService
}

// NewHTTPLocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHTTPLocationService(opts ...option.RequestOption) (r *HTTPLocationService) {
	r = &HTTPLocationService{}
	r.Options = opts
	r.BotClass = NewHTTPLocationBotClassService(opts...)
	r.DeviceType = NewHTTPLocationDeviceTypeService(opts...)
	r.HTTPProtocol = NewHTTPLocationHTTPProtocolService(opts...)
	r.HTTPMethod = NewHTTPLocationHTTPMethodService(opts...)
	r.IPVersion = NewHTTPLocationIPVersionService(opts...)
	r.OS = NewHTTPLocationOSService(opts...)
	r.TLSVersion = NewHTTPLocationTLSVersionService(opts...)
	r.BrowserFamily = NewHTTPLocationBrowserFamilyService(opts...)
	return
}

// Get the top locations by HTTP traffic. Values are a percentage out of the total
// traffic.
func (r *HTTPLocationService) Get(ctx context.Context, query HTTPLocationGetParams, opts ...option.RequestOption) (res *HTTPLocationGetResponse, err error) {
	var env HTTPLocationGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/top/locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationGetResponse struct {
	Meta HTTPLocationGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationGetResponseJSON   `json:"-"`
}

// httpLocationGetResponseJSON contains the JSON metadata for the struct
// [HTTPLocationGetResponse]
type httpLocationGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationGetResponseMeta struct {
	DateRange      []HTTPLocationGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                    `json:"lastUpdated,required"`
	ConfidenceInfo HTTPLocationGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpLocationGetResponseMetaJSON           `json:"-"`
}

// httpLocationGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPLocationGetResponseMeta]
type httpLocationGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                `json:"startTime,required" format:"date-time"`
	JSON      httpLocationGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPLocationGetResponseMetaDateRange]
type httpLocationGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                 `json:"level"`
	JSON        httpLocationGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpLocationGetResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [HTTPLocationGetResponseMetaConfidenceInfo]
type httpLocationGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                  `json:"dataSource,required"`
	Description     string                                                  `json:"description,required"`
	EventType       string                                                  `json:"eventType,required"`
	IsInstantaneous bool                                                    `json:"isInstantaneous,required"`
	EndTime         time.Time                                               `json:"endTime" format:"date-time"`
	LinkedURL       string                                                  `json:"linkedUrl"`
	StartTime       time.Time                                               `json:"startTime" format:"date-time"`
	JSON            httpLocationGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [HTTPLocationGetResponseMetaConfidenceInfoAnnotation]
type httpLocationGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationGetResponseTop0 struct {
	ClientCountryAlpha2 string                          `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                          `json:"clientCountryName,required"`
	Value               string                          `json:"value,required"`
	JSON                httpLocationGetResponseTop0JSON `json:"-"`
}

// httpLocationGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPLocationGetResponseTop0]
type httpLocationGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPLocationGetParamsBotClass] `query:"botClass"`
	// Filter for browser family.
	BrowserFamily param.Field[[]HTTPLocationGetParamsBrowserFamily] `query:"browserFamily"`
	// Array of comma separated list of continents (alpha-2 continent codes). Start
	// with `-` to exclude from results. For example, `-EU,NA` excludes results from
	// Europe, but includes results from North America.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]HTTPLocationGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPLocationGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPLocationGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPLocationGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPLocationGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPLocationGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPLocationGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPLocationGetParams]'s query parameters as `url.Values`.
func (r HTTPLocationGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type HTTPLocationGetParamsBotClass string

const (
	HTTPLocationGetParamsBotClassLikelyAutomated HTTPLocationGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationGetParamsBotClassLikelyHuman     HTTPLocationGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsBotClassLikelyAutomated, HTTPLocationGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationGetParamsBrowserFamily string

const (
	HTTPLocationGetParamsBrowserFamilyChrome  HTTPLocationGetParamsBrowserFamily = "CHROME"
	HTTPLocationGetParamsBrowserFamilyEdge    HTTPLocationGetParamsBrowserFamily = "EDGE"
	HTTPLocationGetParamsBrowserFamilyFirefox HTTPLocationGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationGetParamsBrowserFamilySafari  HTTPLocationGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsBrowserFamilyChrome, HTTPLocationGetParamsBrowserFamilyEdge, HTTPLocationGetParamsBrowserFamilyFirefox, HTTPLocationGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPLocationGetParamsDeviceType string

const (
	HTTPLocationGetParamsDeviceTypeDesktop HTTPLocationGetParamsDeviceType = "DESKTOP"
	HTTPLocationGetParamsDeviceTypeMobile  HTTPLocationGetParamsDeviceType = "MOBILE"
	HTTPLocationGetParamsDeviceTypeOther   HTTPLocationGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsDeviceTypeDesktop, HTTPLocationGetParamsDeviceTypeMobile, HTTPLocationGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPLocationGetParamsFormat string

const (
	HTTPLocationGetParamsFormatJson HTTPLocationGetParamsFormat = "JSON"
	HTTPLocationGetParamsFormatCsv  HTTPLocationGetParamsFormat = "CSV"
)

func (r HTTPLocationGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsFormatJson, HTTPLocationGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationGetParamsHTTPProtocol string

const (
	HTTPLocationGetParamsHTTPProtocolHTTP  HTTPLocationGetParamsHTTPProtocol = "HTTP"
	HTTPLocationGetParamsHTTPProtocolHTTPS HTTPLocationGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsHTTPProtocolHTTP, HTTPLocationGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationGetParamsHTTPVersion string

const (
	HTTPLocationGetParamsHTTPVersionHttPv1 HTTPLocationGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationGetParamsHTTPVersionHttPv2 HTTPLocationGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationGetParamsHTTPVersionHttPv3 HTTPLocationGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsHTTPVersionHttPv1, HTTPLocationGetParamsHTTPVersionHttPv2, HTTPLocationGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationGetParamsIPVersion string

const (
	HTTPLocationGetParamsIPVersionIPv4 HTTPLocationGetParamsIPVersion = "IPv4"
	HTTPLocationGetParamsIPVersionIPv6 HTTPLocationGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsIPVersionIPv4, HTTPLocationGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationGetParamsOS string

const (
	HTTPLocationGetParamsOSWindows  HTTPLocationGetParamsOS = "WINDOWS"
	HTTPLocationGetParamsOSMacosx   HTTPLocationGetParamsOS = "MACOSX"
	HTTPLocationGetParamsOSIos      HTTPLocationGetParamsOS = "IOS"
	HTTPLocationGetParamsOSAndroid  HTTPLocationGetParamsOS = "ANDROID"
	HTTPLocationGetParamsOSChromeos HTTPLocationGetParamsOS = "CHROMEOS"
	HTTPLocationGetParamsOSLinux    HTTPLocationGetParamsOS = "LINUX"
	HTTPLocationGetParamsOSSmartTv  HTTPLocationGetParamsOS = "SMART_TV"
)

func (r HTTPLocationGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsOSWindows, HTTPLocationGetParamsOSMacosx, HTTPLocationGetParamsOSIos, HTTPLocationGetParamsOSAndroid, HTTPLocationGetParamsOSChromeos, HTTPLocationGetParamsOSLinux, HTTPLocationGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationGetParamsTLSVersion string

const (
	HTTPLocationGetParamsTLSVersionTlSv1_0  HTTPLocationGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationGetParamsTLSVersionTlSv1_1  HTTPLocationGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationGetParamsTLSVersionTlSv1_2  HTTPLocationGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationGetParamsTLSVersionTlSv1_3  HTTPLocationGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationGetParamsTLSVersionTlSvQuic HTTPLocationGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationGetParamsTLSVersionTlSv1_0, HTTPLocationGetParamsTLSVersionTlSv1_1, HTTPLocationGetParamsTLSVersionTlSv1_2, HTTPLocationGetParamsTLSVersionTlSv1_3, HTTPLocationGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationGetResponseEnvelope struct {
	Result  HTTPLocationGetResponse             `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    httpLocationGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPLocationGetResponseEnvelope]
type httpLocationGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
