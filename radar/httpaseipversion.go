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

// HTTPAseIPVersionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPAseIPVersionService] method instead.
type HTTPAseIPVersionService struct {
	Options []option.RequestOption
}

// NewHTTPAseIPVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPAseIPVersionService(opts ...option.RequestOption) (r *HTTPAseIPVersionService) {
	r = &HTTPAseIPVersionService{}
	r.Options = opts
	return
}

// Get the top autonomous systems, by HTTP traffic, of the requested IP version.
// Values are a percentage out of the total traffic.
func (r *HTTPAseIPVersionService) Get(ctx context.Context, ipVersion HTTPAseIPVersionGetParamsIPVersion, query HTTPAseIPVersionGetParams, opts ...option.RequestOption) (res *HTTPAseIPVersionGetResponse, err error) {
	var env HTTPAseIPVersionGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/ip_version/%v", ipVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseIPVersionGetResponse struct {
	Meta HTTPAseIPVersionGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseIPVersionGetResponseTop0 `json:"top_0,required"`
	JSON httpAseIPVersionGetResponseJSON   `json:"-"`
}

// httpAseIPVersionGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseIPVersionGetResponse]
type httpAseIPVersionGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPAseIPVersionGetResponseMeta struct {
	DateRange      []HTTPAseIPVersionGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                        `json:"lastUpdated,required"`
	ConfidenceInfo HTTPAseIPVersionGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpAseIPVersionGetResponseMetaJSON           `json:"-"`
}

// httpAseIPVersionGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseIPVersionGetResponseMeta]
type httpAseIPVersionGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseIPVersionGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                    `json:"startTime,required" format:"date-time"`
	JSON      httpAseIPVersionGetResponseMetaDateRangeJSON `json:"-"`
}

// httpAseIPVersionGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPAseIPVersionGetResponseMetaDateRange]
type httpAseIPVersionGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPAseIPVersionGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                     `json:"level"`
	JSON        httpAseIPVersionGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpAseIPVersionGetResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [HTTPAseIPVersionGetResponseMetaConfidenceInfo]
type httpAseIPVersionGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                      `json:"dataSource,required"`
	Description     string                                                      `json:"description,required"`
	EventType       string                                                      `json:"eventType,required"`
	IsInstantaneous bool                                                        `json:"isInstantaneous,required"`
	EndTime         time.Time                                                   `json:"endTime" format:"date-time"`
	LinkedURL       string                                                      `json:"linkedUrl"`
	StartTime       time.Time                                                   `json:"startTime" format:"date-time"`
	JSON            httpAseIPVersionGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpAseIPVersionGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct
// [HTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotation]
type httpAseIPVersionGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPAseIPVersionGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPAseIPVersionGetResponseTop0 struct {
	ClientASN    int64                               `json:"clientASN,required"`
	ClientAsName string                              `json:"clientASName,required"`
	Value        string                              `json:"value,required"`
	JSON         httpAseIPVersionGetResponseTop0JSON `json:"-"`
}

// httpAseIPVersionGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPAseIPVersionGetResponseTop0]
type httpAseIPVersionGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseIPVersionGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPAseIPVersionGetParamsBotClass] `query:"botClass"`
	// Filter for browser family.
	BrowserFamily param.Field[[]HTTPAseIPVersionGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPAseIPVersionGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPAseIPVersionGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPAseIPVersionGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPAseIPVersionGetParamsHTTPVersion] `query:"httpVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPAseIPVersionGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPAseIPVersionGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPAseIPVersionGetParams]'s query parameters as
// `url.Values`.
func (r HTTPAseIPVersionGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// IP version.
type HTTPAseIPVersionGetParamsIPVersion string

const (
	HTTPAseIPVersionGetParamsIPVersionIPv4 HTTPAseIPVersionGetParamsIPVersion = "IPv4"
	HTTPAseIPVersionGetParamsIPVersionIPv6 HTTPAseIPVersionGetParamsIPVersion = "IPv6"
)

func (r HTTPAseIPVersionGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsIPVersionIPv4, HTTPAseIPVersionGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsBotClass string

const (
	HTTPAseIPVersionGetParamsBotClassLikelyAutomated HTTPAseIPVersionGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseIPVersionGetParamsBotClassLikelyHuman     HTTPAseIPVersionGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseIPVersionGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsBotClassLikelyAutomated, HTTPAseIPVersionGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsBrowserFamily string

const (
	HTTPAseIPVersionGetParamsBrowserFamilyChrome  HTTPAseIPVersionGetParamsBrowserFamily = "CHROME"
	HTTPAseIPVersionGetParamsBrowserFamilyEdge    HTTPAseIPVersionGetParamsBrowserFamily = "EDGE"
	HTTPAseIPVersionGetParamsBrowserFamilyFirefox HTTPAseIPVersionGetParamsBrowserFamily = "FIREFOX"
	HTTPAseIPVersionGetParamsBrowserFamilySafari  HTTPAseIPVersionGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPAseIPVersionGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsBrowserFamilyChrome, HTTPAseIPVersionGetParamsBrowserFamilyEdge, HTTPAseIPVersionGetParamsBrowserFamilyFirefox, HTTPAseIPVersionGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsDeviceType string

const (
	HTTPAseIPVersionGetParamsDeviceTypeDesktop HTTPAseIPVersionGetParamsDeviceType = "DESKTOP"
	HTTPAseIPVersionGetParamsDeviceTypeMobile  HTTPAseIPVersionGetParamsDeviceType = "MOBILE"
	HTTPAseIPVersionGetParamsDeviceTypeOther   HTTPAseIPVersionGetParamsDeviceType = "OTHER"
)

func (r HTTPAseIPVersionGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsDeviceTypeDesktop, HTTPAseIPVersionGetParamsDeviceTypeMobile, HTTPAseIPVersionGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPAseIPVersionGetParamsFormat string

const (
	HTTPAseIPVersionGetParamsFormatJson HTTPAseIPVersionGetParamsFormat = "JSON"
	HTTPAseIPVersionGetParamsFormatCsv  HTTPAseIPVersionGetParamsFormat = "CSV"
)

func (r HTTPAseIPVersionGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsFormatJson, HTTPAseIPVersionGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsHTTPProtocol string

const (
	HTTPAseIPVersionGetParamsHTTPProtocolHTTP  HTTPAseIPVersionGetParamsHTTPProtocol = "HTTP"
	HTTPAseIPVersionGetParamsHTTPProtocolHTTPS HTTPAseIPVersionGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseIPVersionGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsHTTPProtocolHTTP, HTTPAseIPVersionGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsHTTPVersion string

const (
	HTTPAseIPVersionGetParamsHTTPVersionHttPv1 HTTPAseIPVersionGetParamsHTTPVersion = "HTTPv1"
	HTTPAseIPVersionGetParamsHTTPVersionHttPv2 HTTPAseIPVersionGetParamsHTTPVersion = "HTTPv2"
	HTTPAseIPVersionGetParamsHTTPVersionHttPv3 HTTPAseIPVersionGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseIPVersionGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsHTTPVersionHttPv1, HTTPAseIPVersionGetParamsHTTPVersionHttPv2, HTTPAseIPVersionGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsOS string

const (
	HTTPAseIPVersionGetParamsOSWindows  HTTPAseIPVersionGetParamsOS = "WINDOWS"
	HTTPAseIPVersionGetParamsOSMacosx   HTTPAseIPVersionGetParamsOS = "MACOSX"
	HTTPAseIPVersionGetParamsOSIos      HTTPAseIPVersionGetParamsOS = "IOS"
	HTTPAseIPVersionGetParamsOSAndroid  HTTPAseIPVersionGetParamsOS = "ANDROID"
	HTTPAseIPVersionGetParamsOSChromeos HTTPAseIPVersionGetParamsOS = "CHROMEOS"
	HTTPAseIPVersionGetParamsOSLinux    HTTPAseIPVersionGetParamsOS = "LINUX"
	HTTPAseIPVersionGetParamsOSSmartTv  HTTPAseIPVersionGetParamsOS = "SMART_TV"
)

func (r HTTPAseIPVersionGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsOSWindows, HTTPAseIPVersionGetParamsOSMacosx, HTTPAseIPVersionGetParamsOSIos, HTTPAseIPVersionGetParamsOSAndroid, HTTPAseIPVersionGetParamsOSChromeos, HTTPAseIPVersionGetParamsOSLinux, HTTPAseIPVersionGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseIPVersionGetParamsTLSVersion string

const (
	HTTPAseIPVersionGetParamsTLSVersionTlSv1_0  HTTPAseIPVersionGetParamsTLSVersion = "TLSv1_0"
	HTTPAseIPVersionGetParamsTLSVersionTlSv1_1  HTTPAseIPVersionGetParamsTLSVersion = "TLSv1_1"
	HTTPAseIPVersionGetParamsTLSVersionTlSv1_2  HTTPAseIPVersionGetParamsTLSVersion = "TLSv1_2"
	HTTPAseIPVersionGetParamsTLSVersionTlSv1_3  HTTPAseIPVersionGetParamsTLSVersion = "TLSv1_3"
	HTTPAseIPVersionGetParamsTLSVersionTlSvQuic HTTPAseIPVersionGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseIPVersionGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseIPVersionGetParamsTLSVersionTlSv1_0, HTTPAseIPVersionGetParamsTLSVersionTlSv1_1, HTTPAseIPVersionGetParamsTLSVersionTlSv1_2, HTTPAseIPVersionGetParamsTLSVersionTlSv1_3, HTTPAseIPVersionGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseIPVersionGetResponseEnvelope struct {
	Result  HTTPAseIPVersionGetResponse             `json:"result,required"`
	Success bool                                    `json:"success,required"`
	JSON    httpAseIPVersionGetResponseEnvelopeJSON `json:"-"`
}

// httpAseIPVersionGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPAseIPVersionGetResponseEnvelope]
type httpAseIPVersionGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseIPVersionGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseIPVersionGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
