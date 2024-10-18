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

// HTTPLocationHTTPMethodService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationHTTPMethodService] method instead.
type HTTPLocationHTTPMethodService struct {
	Options []option.RequestOption
}

// NewHTTPLocationHTTPMethodService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPLocationHTTPMethodService(opts ...option.RequestOption) (r *HTTPLocationHTTPMethodService) {
	r = &HTTPLocationHTTPMethodService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested HTTP version. Values
// are a percentage out of the total traffic.
func (r *HTTPLocationHTTPMethodService) Get(ctx context.Context, httpVersion HTTPLocationHTTPMethodGetParamsHTTPVersion, query HTTPLocationHTTPMethodGetParams, opts ...option.RequestOption) (res *HTTPLocationHTTPMethodGetResponse, err error) {
	var env HTTPLocationHTTPMethodGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/http_version/%v", httpVersion)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationHTTPMethodGetResponse struct {
	Meta HTTPLocationHTTPMethodGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationHTTPMethodGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationHTTPMethodGetResponseJSON   `json:"-"`
}

// httpLocationHTTPMethodGetResponseJSON contains the JSON metadata for the struct
// [HTTPLocationHTTPMethodGetResponse]
type httpLocationHTTPMethodGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationHTTPMethodGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPMethodGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationHTTPMethodGetResponseMeta struct {
	DateRange      []HTTPLocationHTTPMethodGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                              `json:"lastUpdated,required"`
	ConfidenceInfo HTTPLocationHTTPMethodGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpLocationHTTPMethodGetResponseMetaJSON           `json:"-"`
}

// httpLocationHTTPMethodGetResponseMetaJSON contains the JSON metadata for the
// struct [HTTPLocationHTTPMethodGetResponseMeta]
type httpLocationHTTPMethodGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationHTTPMethodGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPMethodGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationHTTPMethodGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                          `json:"startTime,required" format:"date-time"`
	JSON      httpLocationHTTPMethodGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationHTTPMethodGetResponseMetaDateRangeJSON contains the JSON metadata
// for the struct [HTTPLocationHTTPMethodGetResponseMetaDateRange]
type httpLocationHTTPMethodGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationHTTPMethodGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPMethodGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationHTTPMethodGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                           `json:"level"`
	JSON        httpLocationHTTPMethodGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpLocationHTTPMethodGetResponseMetaConfidenceInfoJSON contains the JSON
// metadata for the struct [HTTPLocationHTTPMethodGetResponseMetaConfidenceInfo]
type httpLocationHTTPMethodGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationHTTPMethodGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPMethodGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                            `json:"dataSource,required"`
	Description     string                                                            `json:"description,required"`
	EventType       string                                                            `json:"eventType,required"`
	IsInstantaneous bool                                                              `json:"isInstantaneous,required"`
	EndTime         time.Time                                                         `json:"endTime" format:"date-time"`
	LinkedURL       string                                                            `json:"linkedUrl"`
	StartTime       time.Time                                                         `json:"startTime" format:"date-time"`
	JSON            httpLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [HTTPLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotation]
type httpLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPMethodGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationHTTPMethodGetResponseTop0 struct {
	ClientCountryAlpha2 string                                    `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                    `json:"clientCountryName,required"`
	Value               string                                    `json:"value,required"`
	JSON                httpLocationHTTPMethodGetResponseTop0JSON `json:"-"`
}

// httpLocationHTTPMethodGetResponseTop0JSON contains the JSON metadata for the
// struct [HTTPLocationHTTPMethodGetResponseTop0]
type httpLocationHTTPMethodGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationHTTPMethodGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPMethodGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationHTTPMethodGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPLocationHTTPMethodGetParamsBotClass] `query:"botClass"`
	// Filter for browser family.
	BrowserFamily param.Field[[]HTTPLocationHTTPMethodGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPLocationHTTPMethodGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPLocationHTTPMethodGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPLocationHTTPMethodGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPLocationHTTPMethodGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPLocationHTTPMethodGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPLocationHTTPMethodGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPLocationHTTPMethodGetParams]'s query parameters as
// `url.Values`.
func (r HTTPLocationHTTPMethodGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// HTTP version.
type HTTPLocationHTTPMethodGetParamsHTTPVersion string

const (
	HTTPLocationHTTPMethodGetParamsHTTPVersionHttPv1 HTTPLocationHTTPMethodGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationHTTPMethodGetParamsHTTPVersionHttPv2 HTTPLocationHTTPMethodGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationHTTPMethodGetParamsHTTPVersionHttPv3 HTTPLocationHTTPMethodGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationHTTPMethodGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPMethodGetParamsHTTPVersionHttPv1, HTTPLocationHTTPMethodGetParamsHTTPVersionHttPv2, HTTPLocationHTTPMethodGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationHTTPMethodGetParamsBotClass string

const (
	HTTPLocationHTTPMethodGetParamsBotClassLikelyAutomated HTTPLocationHTTPMethodGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationHTTPMethodGetParamsBotClassLikelyHuman     HTTPLocationHTTPMethodGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationHTTPMethodGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPMethodGetParamsBotClassLikelyAutomated, HTTPLocationHTTPMethodGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationHTTPMethodGetParamsBrowserFamily string

const (
	HTTPLocationHTTPMethodGetParamsBrowserFamilyChrome  HTTPLocationHTTPMethodGetParamsBrowserFamily = "CHROME"
	HTTPLocationHTTPMethodGetParamsBrowserFamilyEdge    HTTPLocationHTTPMethodGetParamsBrowserFamily = "EDGE"
	HTTPLocationHTTPMethodGetParamsBrowserFamilyFirefox HTTPLocationHTTPMethodGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationHTTPMethodGetParamsBrowserFamilySafari  HTTPLocationHTTPMethodGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationHTTPMethodGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPMethodGetParamsBrowserFamilyChrome, HTTPLocationHTTPMethodGetParamsBrowserFamilyEdge, HTTPLocationHTTPMethodGetParamsBrowserFamilyFirefox, HTTPLocationHTTPMethodGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPLocationHTTPMethodGetParamsDeviceType string

const (
	HTTPLocationHTTPMethodGetParamsDeviceTypeDesktop HTTPLocationHTTPMethodGetParamsDeviceType = "DESKTOP"
	HTTPLocationHTTPMethodGetParamsDeviceTypeMobile  HTTPLocationHTTPMethodGetParamsDeviceType = "MOBILE"
	HTTPLocationHTTPMethodGetParamsDeviceTypeOther   HTTPLocationHTTPMethodGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationHTTPMethodGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPMethodGetParamsDeviceTypeDesktop, HTTPLocationHTTPMethodGetParamsDeviceTypeMobile, HTTPLocationHTTPMethodGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPLocationHTTPMethodGetParamsFormat string

const (
	HTTPLocationHTTPMethodGetParamsFormatJson HTTPLocationHTTPMethodGetParamsFormat = "JSON"
	HTTPLocationHTTPMethodGetParamsFormatCsv  HTTPLocationHTTPMethodGetParamsFormat = "CSV"
)

func (r HTTPLocationHTTPMethodGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPMethodGetParamsFormatJson, HTTPLocationHTTPMethodGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationHTTPMethodGetParamsHTTPProtocol string

const (
	HTTPLocationHTTPMethodGetParamsHTTPProtocolHTTP  HTTPLocationHTTPMethodGetParamsHTTPProtocol = "HTTP"
	HTTPLocationHTTPMethodGetParamsHTTPProtocolHTTPS HTTPLocationHTTPMethodGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationHTTPMethodGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPMethodGetParamsHTTPProtocolHTTP, HTTPLocationHTTPMethodGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationHTTPMethodGetParamsIPVersion string

const (
	HTTPLocationHTTPMethodGetParamsIPVersionIPv4 HTTPLocationHTTPMethodGetParamsIPVersion = "IPv4"
	HTTPLocationHTTPMethodGetParamsIPVersionIPv6 HTTPLocationHTTPMethodGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationHTTPMethodGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPMethodGetParamsIPVersionIPv4, HTTPLocationHTTPMethodGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationHTTPMethodGetParamsOS string

const (
	HTTPLocationHTTPMethodGetParamsOSWindows  HTTPLocationHTTPMethodGetParamsOS = "WINDOWS"
	HTTPLocationHTTPMethodGetParamsOSMacosx   HTTPLocationHTTPMethodGetParamsOS = "MACOSX"
	HTTPLocationHTTPMethodGetParamsOSIos      HTTPLocationHTTPMethodGetParamsOS = "IOS"
	HTTPLocationHTTPMethodGetParamsOSAndroid  HTTPLocationHTTPMethodGetParamsOS = "ANDROID"
	HTTPLocationHTTPMethodGetParamsOSChromeos HTTPLocationHTTPMethodGetParamsOS = "CHROMEOS"
	HTTPLocationHTTPMethodGetParamsOSLinux    HTTPLocationHTTPMethodGetParamsOS = "LINUX"
	HTTPLocationHTTPMethodGetParamsOSSmartTv  HTTPLocationHTTPMethodGetParamsOS = "SMART_TV"
)

func (r HTTPLocationHTTPMethodGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPMethodGetParamsOSWindows, HTTPLocationHTTPMethodGetParamsOSMacosx, HTTPLocationHTTPMethodGetParamsOSIos, HTTPLocationHTTPMethodGetParamsOSAndroid, HTTPLocationHTTPMethodGetParamsOSChromeos, HTTPLocationHTTPMethodGetParamsOSLinux, HTTPLocationHTTPMethodGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationHTTPMethodGetParamsTLSVersion string

const (
	HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_0  HTTPLocationHTTPMethodGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_1  HTTPLocationHTTPMethodGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_2  HTTPLocationHTTPMethodGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_3  HTTPLocationHTTPMethodGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationHTTPMethodGetParamsTLSVersionTlSvQuic HTTPLocationHTTPMethodGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationHTTPMethodGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_0, HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_1, HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_2, HTTPLocationHTTPMethodGetParamsTLSVersionTlSv1_3, HTTPLocationHTTPMethodGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationHTTPMethodGetResponseEnvelope struct {
	Result  HTTPLocationHTTPMethodGetResponse             `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    httpLocationHTTPMethodGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationHTTPMethodGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPLocationHTTPMethodGetResponseEnvelope]
type httpLocationHTTPMethodGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationHTTPMethodGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationHTTPMethodGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
