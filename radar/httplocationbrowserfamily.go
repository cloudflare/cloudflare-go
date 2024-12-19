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

// HTTPLocationBrowserFamilyService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationBrowserFamilyService] method instead.
type HTTPLocationBrowserFamilyService struct {
	Options []option.RequestOption
}

// NewHTTPLocationBrowserFamilyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewHTTPLocationBrowserFamilyService(opts ...option.RequestOption) (r *HTTPLocationBrowserFamilyService) {
	r = &HTTPLocationBrowserFamilyService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested browser family. Values
// are a percentage out of the total traffic.
func (r *HTTPLocationBrowserFamilyService) Get(ctx context.Context, browserFamily HTTPLocationBrowserFamilyGetParamsBrowserFamily, query HTTPLocationBrowserFamilyGetParams, opts ...option.RequestOption) (res *HTTPLocationBrowserFamilyGetResponse, err error) {
	var env HTTPLocationBrowserFamilyGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/browser_family/%v", browserFamily)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationBrowserFamilyGetResponse struct {
	Meta HTTPLocationBrowserFamilyGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationBrowserFamilyGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationBrowserFamilyGetResponseJSON   `json:"-"`
}

// httpLocationBrowserFamilyGetResponseJSON contains the JSON metadata for the
// struct [HTTPLocationBrowserFamilyGetResponse]
type httpLocationBrowserFamilyGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBrowserFamilyGetResponseMeta struct {
	DateRange      []HTTPLocationBrowserFamilyGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                                 `json:"lastUpdated,required"`
	ConfidenceInfo HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpLocationBrowserFamilyGetResponseMetaJSON           `json:"-"`
}

// httpLocationBrowserFamilyGetResponseMetaJSON contains the JSON metadata for the
// struct [HTTPLocationBrowserFamilyGetResponseMeta]
type httpLocationBrowserFamilyGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBrowserFamilyGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                             `json:"startTime,required" format:"date-time"`
	JSON      httpLocationBrowserFamilyGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationBrowserFamilyGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [HTTPLocationBrowserFamilyGetResponseMetaDateRange]
type httpLocationBrowserFamilyGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                              `json:"level"`
	JSON        httpLocationBrowserFamilyGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpLocationBrowserFamilyGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfo]
type httpLocationBrowserFamilyGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                               `json:"dataSource,required"`
	Description     string                                                               `json:"description,required"`
	EventType       string                                                               `json:"eventType,required"`
	IsInstantaneous bool                                                                 `json:"isInstantaneous,required"`
	EndTime         time.Time                                                            `json:"endTime" format:"date-time"`
	LinkedURL       string                                                               `json:"linkedUrl"`
	StartTime       time.Time                                                            `json:"startTime" format:"date-time"`
	JSON            httpLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotationJSON contains
// the JSON metadata for the struct
// [HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotation]
type httpLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBrowserFamilyGetResponseTop0 struct {
	ClientCountryAlpha2 string                                       `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                       `json:"clientCountryName,required"`
	Value               string                                       `json:"value,required"`
	JSON                httpLocationBrowserFamilyGetResponseTop0JSON `json:"-"`
}

// httpLocationBrowserFamilyGetResponseTop0JSON contains the JSON metadata for the
// struct [HTTPLocationBrowserFamilyGetResponseTop0]
type httpLocationBrowserFamilyGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBrowserFamilyGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPLocationBrowserFamilyGetParamsBotClass] `query:"botClass"`
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
	DeviceType param.Field[[]HTTPLocationBrowserFamilyGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPLocationBrowserFamilyGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPLocationBrowserFamilyGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPLocationBrowserFamilyGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPLocationBrowserFamilyGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPLocationBrowserFamilyGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPLocationBrowserFamilyGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPLocationBrowserFamilyGetParams]'s query parameters as
// `url.Values`.
func (r HTTPLocationBrowserFamilyGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Browser family.
type HTTPLocationBrowserFamilyGetParamsBrowserFamily string

const (
	HTTPLocationBrowserFamilyGetParamsBrowserFamilyChrome  HTTPLocationBrowserFamilyGetParamsBrowserFamily = "CHROME"
	HTTPLocationBrowserFamilyGetParamsBrowserFamilyEdge    HTTPLocationBrowserFamilyGetParamsBrowserFamily = "EDGE"
	HTTPLocationBrowserFamilyGetParamsBrowserFamilyFirefox HTTPLocationBrowserFamilyGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationBrowserFamilyGetParamsBrowserFamilySafari  HTTPLocationBrowserFamilyGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationBrowserFamilyGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsBrowserFamilyChrome, HTTPLocationBrowserFamilyGetParamsBrowserFamilyEdge, HTTPLocationBrowserFamilyGetParamsBrowserFamilyFirefox, HTTPLocationBrowserFamilyGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsBotClass string

const (
	HTTPLocationBrowserFamilyGetParamsBotClassLikelyAutomated HTTPLocationBrowserFamilyGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationBrowserFamilyGetParamsBotClassLikelyHuman     HTTPLocationBrowserFamilyGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationBrowserFamilyGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsBotClassLikelyAutomated, HTTPLocationBrowserFamilyGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsDeviceType string

const (
	HTTPLocationBrowserFamilyGetParamsDeviceTypeDesktop HTTPLocationBrowserFamilyGetParamsDeviceType = "DESKTOP"
	HTTPLocationBrowserFamilyGetParamsDeviceTypeMobile  HTTPLocationBrowserFamilyGetParamsDeviceType = "MOBILE"
	HTTPLocationBrowserFamilyGetParamsDeviceTypeOther   HTTPLocationBrowserFamilyGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationBrowserFamilyGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsDeviceTypeDesktop, HTTPLocationBrowserFamilyGetParamsDeviceTypeMobile, HTTPLocationBrowserFamilyGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPLocationBrowserFamilyGetParamsFormat string

const (
	HTTPLocationBrowserFamilyGetParamsFormatJson HTTPLocationBrowserFamilyGetParamsFormat = "JSON"
	HTTPLocationBrowserFamilyGetParamsFormatCsv  HTTPLocationBrowserFamilyGetParamsFormat = "CSV"
)

func (r HTTPLocationBrowserFamilyGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsFormatJson, HTTPLocationBrowserFamilyGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsHTTPProtocol string

const (
	HTTPLocationBrowserFamilyGetParamsHTTPProtocolHTTP  HTTPLocationBrowserFamilyGetParamsHTTPProtocol = "HTTP"
	HTTPLocationBrowserFamilyGetParamsHTTPProtocolHTTPS HTTPLocationBrowserFamilyGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationBrowserFamilyGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsHTTPProtocolHTTP, HTTPLocationBrowserFamilyGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsHTTPVersion string

const (
	HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv1 HTTPLocationBrowserFamilyGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv2 HTTPLocationBrowserFamilyGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv3 HTTPLocationBrowserFamilyGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationBrowserFamilyGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv1, HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv2, HTTPLocationBrowserFamilyGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsIPVersion string

const (
	HTTPLocationBrowserFamilyGetParamsIPVersionIPv4 HTTPLocationBrowserFamilyGetParamsIPVersion = "IPv4"
	HTTPLocationBrowserFamilyGetParamsIPVersionIPv6 HTTPLocationBrowserFamilyGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationBrowserFamilyGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsIPVersionIPv4, HTTPLocationBrowserFamilyGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsOS string

const (
	HTTPLocationBrowserFamilyGetParamsOSWindows  HTTPLocationBrowserFamilyGetParamsOS = "WINDOWS"
	HTTPLocationBrowserFamilyGetParamsOSMacosx   HTTPLocationBrowserFamilyGetParamsOS = "MACOSX"
	HTTPLocationBrowserFamilyGetParamsOSIos      HTTPLocationBrowserFamilyGetParamsOS = "IOS"
	HTTPLocationBrowserFamilyGetParamsOSAndroid  HTTPLocationBrowserFamilyGetParamsOS = "ANDROID"
	HTTPLocationBrowserFamilyGetParamsOSChromeos HTTPLocationBrowserFamilyGetParamsOS = "CHROMEOS"
	HTTPLocationBrowserFamilyGetParamsOSLinux    HTTPLocationBrowserFamilyGetParamsOS = "LINUX"
	HTTPLocationBrowserFamilyGetParamsOSSmartTv  HTTPLocationBrowserFamilyGetParamsOS = "SMART_TV"
)

func (r HTTPLocationBrowserFamilyGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsOSWindows, HTTPLocationBrowserFamilyGetParamsOSMacosx, HTTPLocationBrowserFamilyGetParamsOSIos, HTTPLocationBrowserFamilyGetParamsOSAndroid, HTTPLocationBrowserFamilyGetParamsOSChromeos, HTTPLocationBrowserFamilyGetParamsOSLinux, HTTPLocationBrowserFamilyGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetParamsTLSVersion string

const (
	HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_0  HTTPLocationBrowserFamilyGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_1  HTTPLocationBrowserFamilyGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_2  HTTPLocationBrowserFamilyGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_3  HTTPLocationBrowserFamilyGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationBrowserFamilyGetParamsTLSVersionTlSvQuic HTTPLocationBrowserFamilyGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationBrowserFamilyGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_0, HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_1, HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_2, HTTPLocationBrowserFamilyGetParamsTLSVersionTlSv1_3, HTTPLocationBrowserFamilyGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationBrowserFamilyGetResponseEnvelope struct {
	Result  HTTPLocationBrowserFamilyGetResponse             `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    httpLocationBrowserFamilyGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationBrowserFamilyGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [HTTPLocationBrowserFamilyGetResponseEnvelope]
type httpLocationBrowserFamilyGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBrowserFamilyGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBrowserFamilyGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
