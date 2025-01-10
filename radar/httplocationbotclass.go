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

// HTTPLocationBotClassService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPLocationBotClassService] method instead.
type HTTPLocationBotClassService struct {
	Options []option.RequestOption
}

// NewHTTPLocationBotClassService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewHTTPLocationBotClassService(opts ...option.RequestOption) (r *HTTPLocationBotClassService) {
	r = &HTTPLocationBotClassService{}
	r.Options = opts
	return
}

// Get the top locations, by HTTP traffic, of the requested bot class. These two
// categories use Cloudflare's bot score - refer to [Bot
// scores])https://developers.cloudflare.com/bots/concepts/bot-score). Values are a
// percentage out of the total traffic.
func (r *HTTPLocationBotClassService) Get(ctx context.Context, botClass HTTPLocationBotClassGetParamsBotClass, query HTTPLocationBotClassGetParams, opts ...option.RequestOption) (res *HTTPLocationBotClassGetResponse, err error) {
	var env HTTPLocationBotClassGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/locations/bot_class/%v", botClass)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPLocationBotClassGetResponse struct {
	Meta HTTPLocationBotClassGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPLocationBotClassGetResponseTop0 `json:"top_0,required"`
	JSON httpLocationBotClassGetResponseJSON   `json:"-"`
}

// httpLocationBotClassGetResponseJSON contains the JSON metadata for the struct
// [HTTPLocationBotClassGetResponse]
type httpLocationBotClassGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBotClassGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBotClassGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBotClassGetResponseMeta struct {
	DateRange      []HTTPLocationBotClassGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                            `json:"lastUpdated,required"`
	ConfidenceInfo HTTPLocationBotClassGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpLocationBotClassGetResponseMetaJSON           `json:"-"`
}

// httpLocationBotClassGetResponseMetaJSON contains the JSON metadata for the
// struct [HTTPLocationBotClassGetResponseMeta]
type httpLocationBotClassGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPLocationBotClassGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBotClassGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBotClassGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                        `json:"startTime,required" format:"date-time"`
	JSON      httpLocationBotClassGetResponseMetaDateRangeJSON `json:"-"`
}

// httpLocationBotClassGetResponseMetaDateRangeJSON contains the JSON metadata for
// the struct [HTTPLocationBotClassGetResponseMetaDateRange]
type httpLocationBotClassGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBotClassGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBotClassGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBotClassGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPLocationBotClassGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                         `json:"level"`
	JSON        httpLocationBotClassGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpLocationBotClassGetResponseMetaConfidenceInfoJSON contains the JSON metadata
// for the struct [HTTPLocationBotClassGetResponseMetaConfidenceInfo]
type httpLocationBotClassGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBotClassGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBotClassGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBotClassGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                          `json:"dataSource,required"`
	Description     string                                                          `json:"description,required"`
	EventType       string                                                          `json:"eventType,required"`
	IsInstantaneous bool                                                            `json:"isInstantaneous,required"`
	EndTime         time.Time                                                       `json:"endTime" format:"date-time"`
	LinkedURL       string                                                          `json:"linkedUrl"`
	StartTime       time.Time                                                       `json:"startTime" format:"date-time"`
	JSON            httpLocationBotClassGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpLocationBotClassGetResponseMetaConfidenceInfoAnnotationJSON contains the
// JSON metadata for the struct
// [HTTPLocationBotClassGetResponseMetaConfidenceInfoAnnotation]
type httpLocationBotClassGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPLocationBotClassGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBotClassGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBotClassGetResponseTop0 struct {
	ClientCountryAlpha2 string                                  `json:"clientCountryAlpha2,required"`
	ClientCountryName   string                                  `json:"clientCountryName,required"`
	Value               string                                  `json:"value,required"`
	JSON                httpLocationBotClassGetResponseTop0JSON `json:"-"`
}

// httpLocationBotClassGetResponseTop0JSON contains the JSON metadata for the
// struct [HTTPLocationBotClassGetResponseTop0]
type httpLocationBotClassGetResponseTop0JSON struct {
	ClientCountryAlpha2 apijson.Field
	ClientCountryName   apijson.Field
	Value               apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *HTTPLocationBotClassGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBotClassGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPLocationBotClassGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for browser family.
	BrowserFamily param.Field[[]HTTPLocationBotClassGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPLocationBotClassGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPLocationBotClassGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPLocationBotClassGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPLocationBotClassGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPLocationBotClassGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPLocationBotClassGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPLocationBotClassGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPLocationBotClassGetParams]'s query parameters as
// `url.Values`.
func (r HTTPLocationBotClassGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Bot class.
type HTTPLocationBotClassGetParamsBotClass string

const (
	HTTPLocationBotClassGetParamsBotClassLikelyAutomated HTTPLocationBotClassGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPLocationBotClassGetParamsBotClassLikelyHuman     HTTPLocationBotClassGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPLocationBotClassGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPLocationBotClassGetParamsBotClassLikelyAutomated, HTTPLocationBotClassGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPLocationBotClassGetParamsBrowserFamily string

const (
	HTTPLocationBotClassGetParamsBrowserFamilyChrome  HTTPLocationBotClassGetParamsBrowserFamily = "CHROME"
	HTTPLocationBotClassGetParamsBrowserFamilyEdge    HTTPLocationBotClassGetParamsBrowserFamily = "EDGE"
	HTTPLocationBotClassGetParamsBrowserFamilyFirefox HTTPLocationBotClassGetParamsBrowserFamily = "FIREFOX"
	HTTPLocationBotClassGetParamsBrowserFamilySafari  HTTPLocationBotClassGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPLocationBotClassGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPLocationBotClassGetParamsBrowserFamilyChrome, HTTPLocationBotClassGetParamsBrowserFamilyEdge, HTTPLocationBotClassGetParamsBrowserFamilyFirefox, HTTPLocationBotClassGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPLocationBotClassGetParamsDeviceType string

const (
	HTTPLocationBotClassGetParamsDeviceTypeDesktop HTTPLocationBotClassGetParamsDeviceType = "DESKTOP"
	HTTPLocationBotClassGetParamsDeviceTypeMobile  HTTPLocationBotClassGetParamsDeviceType = "MOBILE"
	HTTPLocationBotClassGetParamsDeviceTypeOther   HTTPLocationBotClassGetParamsDeviceType = "OTHER"
)

func (r HTTPLocationBotClassGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPLocationBotClassGetParamsDeviceTypeDesktop, HTTPLocationBotClassGetParamsDeviceTypeMobile, HTTPLocationBotClassGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPLocationBotClassGetParamsFormat string

const (
	HTTPLocationBotClassGetParamsFormatJson HTTPLocationBotClassGetParamsFormat = "JSON"
	HTTPLocationBotClassGetParamsFormatCsv  HTTPLocationBotClassGetParamsFormat = "CSV"
)

func (r HTTPLocationBotClassGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPLocationBotClassGetParamsFormatJson, HTTPLocationBotClassGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPLocationBotClassGetParamsHTTPProtocol string

const (
	HTTPLocationBotClassGetParamsHTTPProtocolHTTP  HTTPLocationBotClassGetParamsHTTPProtocol = "HTTP"
	HTTPLocationBotClassGetParamsHTTPProtocolHTTPS HTTPLocationBotClassGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPLocationBotClassGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPLocationBotClassGetParamsHTTPProtocolHTTP, HTTPLocationBotClassGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPLocationBotClassGetParamsHTTPVersion string

const (
	HTTPLocationBotClassGetParamsHTTPVersionHttPv1 HTTPLocationBotClassGetParamsHTTPVersion = "HTTPv1"
	HTTPLocationBotClassGetParamsHTTPVersionHttPv2 HTTPLocationBotClassGetParamsHTTPVersion = "HTTPv2"
	HTTPLocationBotClassGetParamsHTTPVersionHttPv3 HTTPLocationBotClassGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPLocationBotClassGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationBotClassGetParamsHTTPVersionHttPv1, HTTPLocationBotClassGetParamsHTTPVersionHttPv2, HTTPLocationBotClassGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPLocationBotClassGetParamsIPVersion string

const (
	HTTPLocationBotClassGetParamsIPVersionIPv4 HTTPLocationBotClassGetParamsIPVersion = "IPv4"
	HTTPLocationBotClassGetParamsIPVersionIPv6 HTTPLocationBotClassGetParamsIPVersion = "IPv6"
)

func (r HTTPLocationBotClassGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPLocationBotClassGetParamsIPVersionIPv4, HTTPLocationBotClassGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPLocationBotClassGetParamsOS string

const (
	HTTPLocationBotClassGetParamsOSWindows  HTTPLocationBotClassGetParamsOS = "WINDOWS"
	HTTPLocationBotClassGetParamsOSMacosx   HTTPLocationBotClassGetParamsOS = "MACOSX"
	HTTPLocationBotClassGetParamsOSIos      HTTPLocationBotClassGetParamsOS = "IOS"
	HTTPLocationBotClassGetParamsOSAndroid  HTTPLocationBotClassGetParamsOS = "ANDROID"
	HTTPLocationBotClassGetParamsOSChromeos HTTPLocationBotClassGetParamsOS = "CHROMEOS"
	HTTPLocationBotClassGetParamsOSLinux    HTTPLocationBotClassGetParamsOS = "LINUX"
	HTTPLocationBotClassGetParamsOSSmartTv  HTTPLocationBotClassGetParamsOS = "SMART_TV"
)

func (r HTTPLocationBotClassGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPLocationBotClassGetParamsOSWindows, HTTPLocationBotClassGetParamsOSMacosx, HTTPLocationBotClassGetParamsOSIos, HTTPLocationBotClassGetParamsOSAndroid, HTTPLocationBotClassGetParamsOSChromeos, HTTPLocationBotClassGetParamsOSLinux, HTTPLocationBotClassGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPLocationBotClassGetParamsTLSVersion string

const (
	HTTPLocationBotClassGetParamsTLSVersionTlSv1_0  HTTPLocationBotClassGetParamsTLSVersion = "TLSv1_0"
	HTTPLocationBotClassGetParamsTLSVersionTlSv1_1  HTTPLocationBotClassGetParamsTLSVersion = "TLSv1_1"
	HTTPLocationBotClassGetParamsTLSVersionTlSv1_2  HTTPLocationBotClassGetParamsTLSVersion = "TLSv1_2"
	HTTPLocationBotClassGetParamsTLSVersionTlSv1_3  HTTPLocationBotClassGetParamsTLSVersion = "TLSv1_3"
	HTTPLocationBotClassGetParamsTLSVersionTlSvQuic HTTPLocationBotClassGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPLocationBotClassGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPLocationBotClassGetParamsTLSVersionTlSv1_0, HTTPLocationBotClassGetParamsTLSVersionTlSv1_1, HTTPLocationBotClassGetParamsTLSVersionTlSv1_2, HTTPLocationBotClassGetParamsTLSVersionTlSv1_3, HTTPLocationBotClassGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPLocationBotClassGetResponseEnvelope struct {
	Result  HTTPLocationBotClassGetResponse             `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    httpLocationBotClassGetResponseEnvelopeJSON `json:"-"`
}

// httpLocationBotClassGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [HTTPLocationBotClassGetResponseEnvelope]
type httpLocationBotClassGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPLocationBotClassGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpLocationBotClassGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
