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

// HTTPAseTLSVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPAseTLSVersionService] method instead.
type HTTPAseTLSVersionService struct {
	Options []option.RequestOption
}

// NewHTTPAseTLSVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPAseTLSVersionService(opts ...option.RequestOption) (r *HTTPAseTLSVersionService) {
	r = &HTTPAseTLSVersionService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested TLS
// protocol version. Values are a percentage out of the total traffic.
func (r *HTTPAseTLSVersionService) Get(ctx context.Context, tlsVersion HTTPAseTLSVersionGetParamsTLSVersion, query HTTPAseTLSVersionGetParams, opts ...option.RequestOption) (res *HTTPAseTLSVersionGetResponse, err error) {
	var env HTTPAseTLSVersionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/tls_version/%v", tlsVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseTLSVersionGetResponse struct {
	Meta HTTPAseTLSVersionGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseTLSVersionGetResponseTop0 `json:"top_0,required"`
	JSON httpAseTLSVersionGetResponseJSON   `json:"-"`
}

// httpAseTLSVersionGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseTLSVersionGetResponse]
type httpAseTLSVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetResponseMeta struct {
	DateRange      []HTTPAseTLSVersionGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                         `json:"lastUpdated,required"`
	ConfidenceInfo HTTPAseTLSVersionGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpAseTLSVersionGetResponseMetaJSON           `json:"-"`
}

// httpAseTLSVersionGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseTLSVersionGetResponseMeta]
type httpAseTLSVersionGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                     `json:"startTime,required" format:"date-time"`
	JSON      httpAseTLSVersionGetResponseMetaDateRangeJSON `json:"-"`
}

// httpAseTLSVersionGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPAseTLSVersionGetResponseMetaDateRange]
type httpAseTLSVersionGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                      `json:"level"`
	JSON        httpAseTLSVersionGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpAseTLSVersionGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPAseTLSVersionGetResponseMetaConfidenceInfo]
type httpAseTLSVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                       `json:"dataSource,required"`
	Description     string                                                       `json:"description,required"`
	EventType       string                                                       `json:"eventType,required"`
	IsInstantaneous bool                                                         `json:"isInstantaneous,required"`
	EndTime         time.Time                                                    `json:"endTime" format:"date-time"`
	LinkedURL       string                                                       `json:"linkedUrl"`
	StartTime       time.Time                                                    `json:"startTime" format:"date-time"`
	JSON            httpAseTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpAseTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [HTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotation]
type httpAseTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPAseTLSVersionGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetResponseTop0 struct {
	ClientASN    int64                                `json:"clientASN,required"`
	ClientAsName string                               `json:"clientASName,required"`
	Value        string                               `json:"value,required"`
	JSON         httpAseTLSVersionGetResponseTop0JSON `json:"-"`
}

// httpAseTLSVersionGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPAseTLSVersionGetResponseTop0]
type httpAseTLSVersionGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseTLSVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseTLSVersionGetParamsBotClass] `query:"botClass"`
	// Filter for browser family.
	BrowserFamily param.Field[[]HTTPAseTLSVersionGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPAseTLSVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPAseTLSVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPAseTLSVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPAseTLSVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPAseTLSVersionGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPAseTLSVersionGetParamsOS] `query:"os"`
}

// URLQuery serializes [HTTPAseTLSVersionGetParams]'s query parameters as
// `url.Values`.
func (r HTTPAseTLSVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// TLS version.
type HTTPAseTLSVersionGetParamsTLSVersion string

const (
	HTTPAseTLSVersionGetParamsTLSVersionTlSv1_0  HTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_0"
	HTTPAseTLSVersionGetParamsTLSVersionTlSv1_1  HTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_1"
	HTTPAseTLSVersionGetParamsTLSVersionTlSv1_2  HTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_2"
	HTTPAseTLSVersionGetParamsTLSVersionTlSv1_3  HTTPAseTLSVersionGetParamsTLSVersion = "TLSv1_3"
	HTTPAseTLSVersionGetParamsTLSVersionTlSvQuic HTTPAseTLSVersionGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseTLSVersionGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsTLSVersionTlSv1_0, HTTPAseTLSVersionGetParamsTLSVersionTlSv1_1, HTTPAseTLSVersionGetParamsTLSVersionTlSv1_2, HTTPAseTLSVersionGetParamsTLSVersionTlSv1_3, HTTPAseTLSVersionGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsBotClass string

const (
	HTTPAseTLSVersionGetParamsBotClassLikelyAutomated HTTPAseTLSVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseTLSVersionGetParamsBotClassLikelyHuman     HTTPAseTLSVersionGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseTLSVersionGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsBotClassLikelyAutomated, HTTPAseTLSVersionGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsBrowserFamily string

const (
	HTTPAseTLSVersionGetParamsBrowserFamilyChrome  HTTPAseTLSVersionGetParamsBrowserFamily = "CHROME"
	HTTPAseTLSVersionGetParamsBrowserFamilyEdge    HTTPAseTLSVersionGetParamsBrowserFamily = "EDGE"
	HTTPAseTLSVersionGetParamsBrowserFamilyFirefox HTTPAseTLSVersionGetParamsBrowserFamily = "FIREFOX"
	HTTPAseTLSVersionGetParamsBrowserFamilySafari  HTTPAseTLSVersionGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPAseTLSVersionGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsBrowserFamilyChrome, HTTPAseTLSVersionGetParamsBrowserFamilyEdge, HTTPAseTLSVersionGetParamsBrowserFamilyFirefox, HTTPAseTLSVersionGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsDeviceType string

const (
	HTTPAseTLSVersionGetParamsDeviceTypeDesktop HTTPAseTLSVersionGetParamsDeviceType = "DESKTOP"
	HTTPAseTLSVersionGetParamsDeviceTypeMobile  HTTPAseTLSVersionGetParamsDeviceType = "MOBILE"
	HTTPAseTLSVersionGetParamsDeviceTypeOther   HTTPAseTLSVersionGetParamsDeviceType = "OTHER"
)

func (r HTTPAseTLSVersionGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsDeviceTypeDesktop, HTTPAseTLSVersionGetParamsDeviceTypeMobile, HTTPAseTLSVersionGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPAseTLSVersionGetParamsFormat string

const (
	HTTPAseTLSVersionGetParamsFormatJson HTTPAseTLSVersionGetParamsFormat = "JSON"
	HTTPAseTLSVersionGetParamsFormatCsv  HTTPAseTLSVersionGetParamsFormat = "CSV"
)

func (r HTTPAseTLSVersionGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsFormatJson, HTTPAseTLSVersionGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsHTTPProtocol string

const (
	HTTPAseTLSVersionGetParamsHTTPProtocolHTTP  HTTPAseTLSVersionGetParamsHTTPProtocol = "HTTP"
	HTTPAseTLSVersionGetParamsHTTPProtocolHTTPS HTTPAseTLSVersionGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseTLSVersionGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsHTTPProtocolHTTP, HTTPAseTLSVersionGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsHTTPVersion string

const (
	HTTPAseTLSVersionGetParamsHTTPVersionHttPv1 HTTPAseTLSVersionGetParamsHTTPVersion = "HTTPv1"
	HTTPAseTLSVersionGetParamsHTTPVersionHttPv2 HTTPAseTLSVersionGetParamsHTTPVersion = "HTTPv2"
	HTTPAseTLSVersionGetParamsHTTPVersionHttPv3 HTTPAseTLSVersionGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseTLSVersionGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsHTTPVersionHttPv1, HTTPAseTLSVersionGetParamsHTTPVersionHttPv2, HTTPAseTLSVersionGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsIPVersion string

const (
	HTTPAseTLSVersionGetParamsIPVersionIPv4 HTTPAseTLSVersionGetParamsIPVersion = "IPv4"
	HTTPAseTLSVersionGetParamsIPVersionIPv6 HTTPAseTLSVersionGetParamsIPVersion = "IPv6"
)

func (r HTTPAseTLSVersionGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsIPVersionIPv4, HTTPAseTLSVersionGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetParamsOS string

const (
	HTTPAseTLSVersionGetParamsOSWindows  HTTPAseTLSVersionGetParamsOS = "WINDOWS"
	HTTPAseTLSVersionGetParamsOSMacosx   HTTPAseTLSVersionGetParamsOS = "MACOSX"
	HTTPAseTLSVersionGetParamsOSIos      HTTPAseTLSVersionGetParamsOS = "IOS"
	HTTPAseTLSVersionGetParamsOSAndroid  HTTPAseTLSVersionGetParamsOS = "ANDROID"
	HTTPAseTLSVersionGetParamsOSChromeos HTTPAseTLSVersionGetParamsOS = "CHROMEOS"
	HTTPAseTLSVersionGetParamsOSLinux    HTTPAseTLSVersionGetParamsOS = "LINUX"
	HTTPAseTLSVersionGetParamsOSSmartTv  HTTPAseTLSVersionGetParamsOS = "SMART_TV"
)

func (r HTTPAseTLSVersionGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseTLSVersionGetParamsOSWindows, HTTPAseTLSVersionGetParamsOSMacosx, HTTPAseTLSVersionGetParamsOSIos, HTTPAseTLSVersionGetParamsOSAndroid, HTTPAseTLSVersionGetParamsOSChromeos, HTTPAseTLSVersionGetParamsOSLinux, HTTPAseTLSVersionGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseTLSVersionGetResponseEnvelope struct {
	Result  HTTPAseTLSVersionGetResponse             `json:"result,required"`
	Success bool                                     `json:"success,required"`
	JSON    httpAseTLSVersionGetResponseEnvelopeJSON `json:"-"`
}

// httpAseTLSVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPAseTLSVersionGetResponseEnvelope]
type httpAseTLSVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseTLSVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseTLSVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
