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

// HTTPLocationIPVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationIPVersionService] method instead.
type HTTPLocationIPVersionService struct {
	Options []option.RequestOption
}

// NewHTTPLocationIPVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPLocationIPVersionService(opts ...option.RequestOption) (r *HTTPLocationIPVersionService) {
	r = &HTTPLocationIPVersionService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested IP version. Values are
// a percentage out of the total traffic.
func (r *HTTPLocationIPVersionService) Get(ctx context.Context, ipVersion HTTPLocationIPVersionGetParamsIPVersion, query HTTPLocationIPVersionGetParams, opts ...option.RequestOption) (res *HTTPLocationIPVersionGetResponse, err error) {
	var env HTTPLocationIPVersionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationIPVersionGetResponse struct {
	Meta HTTPLocationIPVersionGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationIPVersionGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationIPVersionGetResponseJSON   `json:"-"`
}

// httpLocationIPVersionGetResponseJSON contains the JSON metadata for the struct
// [HTTPLocationIPVersionGetResponse]
type httpLocationIPVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationIPVersionGetResponseMeta struct {
	DateRange      []HTTPLocationIPVersionGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                             `json:"lastUpdated,required"`
	ConfidenceInfo HTTPLocationIPVersionGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpLocationIPVersionGetResponseMetaJSON           `json:"-"`
}

// httpLocationIPVersionGetResponseMetaJSON contains the JSON metadata for the
// struct [HTTPLocationIPVersionGetResponseMeta]
type httpLocationIPVersionGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationIPVersionGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                         `json:"startTime,required" format:"date-time"`
	JSON      httpLocationIPVersionGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationIPVersionGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [HTTPLocationIPVersionGetResponseMetaDateRange]
type httpLocationIPVersionGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationIPVersionGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                          `json:"level"`
	JSON        httpLocationIPVersionGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpLocationIPVersionGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [HTTPLocationIPVersionGetResponseMetaConfidenceInfo]
type httpLocationIPVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                           `json:"dataSource,required"`
	Description     string                                                           `json:"description,required"`
	EventType       string                                                           `json:"eventType,required"`
	IsInstantaneous bool                                                             `json:"isInstantaneous,required"`
	EndTime         time.Time                                                        `json:"endTime" format:"date-time"`
	LinkedURL       string                                                           `json:"linkedUrl"`
	StartTime       time.Time                                                        `json:"startTime" format:"date-time"`
	JSON            httpLocationIPVersionGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationIPVersionGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [HTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotation]
type httpLocationIPVersionGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationIPVersionGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationIPVersionGetResponseTop0 struct {
	ClientCountryAlpha2 string                                   `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                   `json:"clientCountryName,required"`
	Value               string                                   `json:"value,required"`
	JSON                httpLocationIPVersionGetResponseTop0JSON `json:"-"`
}

// httpLocationIPVersionGetResponseTop0JSON contains the JSON metadata for the
// struct [HTTPLocationIPVersionGetResponseTop0]
type httpLocationIPVersionGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationIPVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPLocationIPVersionGetParamsBotClass] `query:"botClass"`
	// Filter for browser family.
	BrowserFamily param.Field[[]HTTPLocationIPVersionGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPLocationIPVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPLocationIPVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPLocationIPVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPLocationIPVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPLocationIPVersionGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPLocationIPVersionGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPLocationIPVersionGetParams]'s query parameters as
// `url.Values`.
func (r HTTPLocationIPVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// IP version.
type HTTPLocationIPVersionGetParamsIPVersion string

const (
	HTTPLocationIPVersionGetParamsIPVersionIPv4 HTTPLocationIPVersionGetParamsIPVersion = "IPv4"
	HTTPLocationIPVersionGetParamsIPVersionIPv6 HTTPLocationIPVersionGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationIPVersionGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsIPVersionIPv4, HTTPLocationIPVersionGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsBotClass string

const (
	HTTPLocationIPVersionGetParamsBotClassLikelyAutomated HTTPLocationIPVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationIPVersionGetParamsBotClassLikelyHuman     HTTPLocationIPVersionGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationIPVersionGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsBotClassLikelyAutomated, HTTPLocationIPVersionGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsBrowserFamily string

const (
	HTTPLocationIPVersionGetParamsBrowserFamilyChrome  HTTPLocationIPVersionGetParamsBrowserFamily = "CHROME"
	HTTPLocationIPVersionGetParamsBrowserFamilyEdge    HTTPLocationIPVersionGetParamsBrowserFamily = "EDGE"
	HTTPLocationIPVersionGetParamsBrowserFamilyFirefox HTTPLocationIPVersionGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationIPVersionGetParamsBrowserFamilySafari  HTTPLocationIPVersionGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationIPVersionGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsBrowserFamilyChrome, HTTPLocationIPVersionGetParamsBrowserFamilyEdge, HTTPLocationIPVersionGetParamsBrowserFamilyFirefox, HTTPLocationIPVersionGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsDeviceType string

const (
	HTTPLocationIPVersionGetParamsDeviceTypeDesktop HTTPLocationIPVersionGetParamsDeviceType = "DESKTOP"
	HTTPLocationIPVersionGetParamsDeviceTypeMobile  HTTPLocationIPVersionGetParamsDeviceType = "MOBILE"
	HTTPLocationIPVersionGetParamsDeviceTypeOther   HTTPLocationIPVersionGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationIPVersionGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsDeviceTypeDesktop, HTTPLocationIPVersionGetParamsDeviceTypeMobile, HTTPLocationIPVersionGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPLocationIPVersionGetParamsFormat string

const (
	HTTPLocationIPVersionGetParamsFormatJson HTTPLocationIPVersionGetParamsFormat = "JSON"
	HTTPLocationIPVersionGetParamsFormatCsv  HTTPLocationIPVersionGetParamsFormat = "CSV"
)

func (r HTTPLocationIPVersionGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsFormatJson, HTTPLocationIPVersionGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsHTTPProtocol string

const (
	HTTPLocationIPVersionGetParamsHTTPProtocolHTTP  HTTPLocationIPVersionGetParamsHTTPProtocol = "HTTP"
	HTTPLocationIPVersionGetParamsHTTPProtocolHTTPS HTTPLocationIPVersionGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationIPVersionGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsHTTPProtocolHTTP, HTTPLocationIPVersionGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsHTTPVersion string

const (
	HTTPLocationIPVersionGetParamsHTTPVersionHttPv1 HTTPLocationIPVersionGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationIPVersionGetParamsHTTPVersionHttPv2 HTTPLocationIPVersionGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationIPVersionGetParamsHTTPVersionHttPv3 HTTPLocationIPVersionGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationIPVersionGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsHTTPVersionHttPv1, HTTPLocationIPVersionGetParamsHTTPVersionHttPv2, HTTPLocationIPVersionGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsOS string

const (
	HTTPLocationIPVersionGetParamsOSWindows  HTTPLocationIPVersionGetParamsOS = "WINDOWS"
	HTTPLocationIPVersionGetParamsOSMacosx   HTTPLocationIPVersionGetParamsOS = "MACOSX"
	HTTPLocationIPVersionGetParamsOSIos      HTTPLocationIPVersionGetParamsOS = "IOS"
	HTTPLocationIPVersionGetParamsOSAndroid  HTTPLocationIPVersionGetParamsOS = "ANDROID"
	HTTPLocationIPVersionGetParamsOSChromeos HTTPLocationIPVersionGetParamsOS = "CHROMEOS"
	HTTPLocationIPVersionGetParamsOSLinux    HTTPLocationIPVersionGetParamsOS = "LINUX"
	HTTPLocationIPVersionGetParamsOSSmartTv  HTTPLocationIPVersionGetParamsOS = "SMART_TV"
)

func (r HTTPLocationIPVersionGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsOSWindows, HTTPLocationIPVersionGetParamsOSMacosx, HTTPLocationIPVersionGetParamsOSIos, HTTPLocationIPVersionGetParamsOSAndroid, HTTPLocationIPVersionGetParamsOSChromeos, HTTPLocationIPVersionGetParamsOSLinux, HTTPLocationIPVersionGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetParamsTLSVersion string

const (
	HTTPLocationIPVersionGetParamsTLSVersionTlSv1_0  HTTPLocationIPVersionGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationIPVersionGetParamsTLSVersionTlSv1_1  HTTPLocationIPVersionGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationIPVersionGetParamsTLSVersionTlSv1_2  HTTPLocationIPVersionGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationIPVersionGetParamsTLSVersionTlSv1_3  HTTPLocationIPVersionGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationIPVersionGetParamsTLSVersionTlSvQuic HTTPLocationIPVersionGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationIPVersionGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationIPVersionGetParamsTLSVersionTlSv1_0, HTTPLocationIPVersionGetParamsTLSVersionTlSv1_1, HTTPLocationIPVersionGetParamsTLSVersionTlSv1_2, HTTPLocationIPVersionGetParamsTLSVersionTlSv1_3, HTTPLocationIPVersionGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationIPVersionGetResponseEnvelope struct {
	Result  HTTPLocationIPVersionGetResponse             `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    httpLocationIPVersionGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationIPVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPLocationIPVersionGetResponseEnvelope]
type httpLocationIPVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationIPVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationIPVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
