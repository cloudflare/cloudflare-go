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

// HTTPAseBotClassService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPAseBotClassService] method instead.
type HTTPAseBotClassService struct {
	Options []option.RequestOption
}

// NewHTTPAseBotClassService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHTTPAseBotClassService(opts ...option.RequestOption) (r *HTTPAseBotClassService) {
	r = &HTTPAseBotClassService{}
	r.Options = opts
	return
}

// Get the top autonomous systems (AS), by HTTP traffic, of the requested bot
// class. These two categories use Cloudflare's bot score - refer to
// [Bot Scores](https://developers.cloudflare.com/bots/concepts/bot-score) for more
// information. Values are a percentage out of the total traffic.
func (r *HTTPAseBotClassService) Get(ctx context.Context, botClass HTTPAseBotClassGetParamsBotClass, query HTTPAseBotClassGetParams, opts ...option.RequestOption) (res *HTTPAseBotClassGetResponse, err error) {
	var env HTTPAseBotClassGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("radar/http/top/ases/bot_class/%v", botClass)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPAseBotClassGetResponse struct {
	Meta HTTPAseBotClassGetResponseMeta   `json:"meta,required"`
	Top0 []HTTPAseBotClassGetResponseTop0 `json:"top_0,required"`
	JSON httpAseBotClassGetResponseJSON   `json:"-"`
}

// httpAseBotClassGetResponseJSON contains the JSON metadata for the struct
// [HTTPAseBotClassGetResponse]
type httpAseBotClassGetResponseJSON struct {
	Meta        apijson.Field
	Top0        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBotClassGetResponseMeta struct {
	DateRange      []HTTPAseBotClassGetResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    string                                       `json:"lastUpdated,required"`
	ConfidenceInfo HTTPAseBotClassGetResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpAseBotClassGetResponseMetaJSON           `json:"-"`
}

// httpAseBotClassGetResponseMetaJSON contains the JSON metadata for the struct
// [HTTPAseBotClassGetResponseMeta]
type httpAseBotClassGetResponseMetaJSON struct {
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBotClassGetResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                                   `json:"startTime,required" format:"date-time"`
	JSON      httpAseBotClassGetResponseMetaDateRangeJSON `json:"-"`
}

// httpAseBotClassGetResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPAseBotClassGetResponseMetaDateRange]
type httpAseBotClassGetResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBotClassGetResponseMetaConfidenceInfo struct {
	Annotations []HTTPAseBotClassGetResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                    `json:"level"`
	JSON        httpAseBotClassGetResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpAseBotClassGetResponseMetaConfidenceInfoJSON contains the JSON metadata for
// the struct [HTTPAseBotClassGetResponseMetaConfidenceInfo]
type httpAseBotClassGetResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBotClassGetResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                     `json:"dataSource,required"`
	Description     string                                                     `json:"description,required"`
	EventType       string                                                     `json:"eventType,required"`
	IsInstantaneous bool                                                       `json:"isInstantaneous,required"`
	EndTime         time.Time                                                  `json:"endTime" format:"date-time"`
	LinkedURL       string                                                     `json:"linkedUrl"`
	StartTime       time.Time                                                  `json:"startTime" format:"date-time"`
	JSON            httpAseBotClassGetResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpAseBotClassGetResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [HTTPAseBotClassGetResponseMetaConfidenceInfoAnnotation]
type httpAseBotClassGetResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPAseBotClassGetResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPAseBotClassGetResponseTop0 struct {
	ClientASN    int64                              `json:"clientASN,required"`
	ClientAsName string                             `json:"clientASName,required"`
	Value        string                             `json:"value,required"`
	JSON         httpAseBotClassGetResponseTop0JSON `json:"-"`
}

// httpAseBotClassGetResponseTop0JSON contains the JSON metadata for the struct
// [HTTPAseBotClassGetResponseTop0]
type httpAseBotClassGetResponseTop0JSON struct {
	ClientASN    apijson.Field
	ClientAsName apijson.Field
	Value        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponseTop0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseTop0JSON) RawJSON() string {
	return r.raw
}

type HTTPAseBotClassGetParams struct {
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for browser family.
	BrowserFamily param.Field[[]HTTPAseBotClassGetParamsBrowserFamily] `query:"browserFamily"`
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
	DeviceType param.Field[[]HTTPAseBotClassGetParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[HTTPAseBotClassGetParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]HTTPAseBotClassGetParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]HTTPAseBotClassGetParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]HTTPAseBotClassGetParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects in the response.
	Limit param.Field[int64] `query:"limit"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	OS param.Field[[]HTTPAseBotClassGetParamsOS] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]HTTPAseBotClassGetParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPAseBotClassGetParams]'s query parameters as
// `url.Values`.
func (r HTTPAseBotClassGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Bot class.
type HTTPAseBotClassGetParamsBotClass string

const (
	HTTPAseBotClassGetParamsBotClassLikelyAutomated HTTPAseBotClassGetParamsBotClass = "LIKELY_AUTOMATED"
	HTTPAseBotClassGetParamsBotClassLikelyHuman     HTTPAseBotClassGetParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPAseBotClassGetParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsBotClassLikelyAutomated, HTTPAseBotClassGetParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsBrowserFamily string

const (
	HTTPAseBotClassGetParamsBrowserFamilyChrome  HTTPAseBotClassGetParamsBrowserFamily = "CHROME"
	HTTPAseBotClassGetParamsBrowserFamilyEdge    HTTPAseBotClassGetParamsBrowserFamily = "EDGE"
	HTTPAseBotClassGetParamsBrowserFamilyFirefox HTTPAseBotClassGetParamsBrowserFamily = "FIREFOX"
	HTTPAseBotClassGetParamsBrowserFamilySafari  HTTPAseBotClassGetParamsBrowserFamily = "SAFARI"
)

func (r HTTPAseBotClassGetParamsBrowserFamily) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsBrowserFamilyChrome, HTTPAseBotClassGetParamsBrowserFamilyEdge, HTTPAseBotClassGetParamsBrowserFamilyFirefox, HTTPAseBotClassGetParamsBrowserFamilySafari:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsDeviceType string

const (
	HTTPAseBotClassGetParamsDeviceTypeDesktop HTTPAseBotClassGetParamsDeviceType = "DESKTOP"
	HTTPAseBotClassGetParamsDeviceTypeMobile  HTTPAseBotClassGetParamsDeviceType = "MOBILE"
	HTTPAseBotClassGetParamsDeviceTypeOther   HTTPAseBotClassGetParamsDeviceType = "OTHER"
)

func (r HTTPAseBotClassGetParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsDeviceTypeDesktop, HTTPAseBotClassGetParamsDeviceTypeMobile, HTTPAseBotClassGetParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format results are returned in.
type HTTPAseBotClassGetParamsFormat string

const (
	HTTPAseBotClassGetParamsFormatJson HTTPAseBotClassGetParamsFormat = "JSON"
	HTTPAseBotClassGetParamsFormatCsv  HTTPAseBotClassGetParamsFormat = "CSV"
)

func (r HTTPAseBotClassGetParamsFormat) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsFormatJson, HTTPAseBotClassGetParamsFormatCsv:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsHTTPProtocol string

const (
	HTTPAseBotClassGetParamsHTTPProtocolHTTP  HTTPAseBotClassGetParamsHTTPProtocol = "HTTP"
	HTTPAseBotClassGetParamsHTTPProtocolHTTPS HTTPAseBotClassGetParamsHTTPProtocol = "HTTPS"
)

func (r HTTPAseBotClassGetParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsHTTPProtocolHTTP, HTTPAseBotClassGetParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsHTTPVersion string

const (
	HTTPAseBotClassGetParamsHTTPVersionHttPv1 HTTPAseBotClassGetParamsHTTPVersion = "HTTPv1"
	HTTPAseBotClassGetParamsHTTPVersionHttPv2 HTTPAseBotClassGetParamsHTTPVersion = "HTTPv2"
	HTTPAseBotClassGetParamsHTTPVersionHttPv3 HTTPAseBotClassGetParamsHTTPVersion = "HTTPv3"
)

func (r HTTPAseBotClassGetParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsHTTPVersionHttPv1, HTTPAseBotClassGetParamsHTTPVersionHttPv2, HTTPAseBotClassGetParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsIPVersion string

const (
	HTTPAseBotClassGetParamsIPVersionIPv4 HTTPAseBotClassGetParamsIPVersion = "IPv4"
	HTTPAseBotClassGetParamsIPVersionIPv6 HTTPAseBotClassGetParamsIPVersion = "IPv6"
)

func (r HTTPAseBotClassGetParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsIPVersionIPv4, HTTPAseBotClassGetParamsIPVersionIPv6:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsOS string

const (
	HTTPAseBotClassGetParamsOSWindows  HTTPAseBotClassGetParamsOS = "WINDOWS"
	HTTPAseBotClassGetParamsOSMacosx   HTTPAseBotClassGetParamsOS = "MACOSX"
	HTTPAseBotClassGetParamsOSIos      HTTPAseBotClassGetParamsOS = "IOS"
	HTTPAseBotClassGetParamsOSAndroid  HTTPAseBotClassGetParamsOS = "ANDROID"
	HTTPAseBotClassGetParamsOSChromeos HTTPAseBotClassGetParamsOS = "CHROMEOS"
	HTTPAseBotClassGetParamsOSLinux    HTTPAseBotClassGetParamsOS = "LINUX"
	HTTPAseBotClassGetParamsOSSmartTv  HTTPAseBotClassGetParamsOS = "SMART_TV"
)

func (r HTTPAseBotClassGetParamsOS) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsOSWindows, HTTPAseBotClassGetParamsOSMacosx, HTTPAseBotClassGetParamsOSIos, HTTPAseBotClassGetParamsOSAndroid, HTTPAseBotClassGetParamsOSChromeos, HTTPAseBotClassGetParamsOSLinux, HTTPAseBotClassGetParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPAseBotClassGetParamsTLSVersion string

const (
	HTTPAseBotClassGetParamsTLSVersionTlSv1_0  HTTPAseBotClassGetParamsTLSVersion = "TLSv1_0"
	HTTPAseBotClassGetParamsTLSVersionTlSv1_1  HTTPAseBotClassGetParamsTLSVersion = "TLSv1_1"
	HTTPAseBotClassGetParamsTLSVersionTlSv1_2  HTTPAseBotClassGetParamsTLSVersion = "TLSv1_2"
	HTTPAseBotClassGetParamsTLSVersionTlSv1_3  HTTPAseBotClassGetParamsTLSVersion = "TLSv1_3"
	HTTPAseBotClassGetParamsTLSVersionTlSvQuic HTTPAseBotClassGetParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPAseBotClassGetParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPAseBotClassGetParamsTLSVersionTlSv1_0, HTTPAseBotClassGetParamsTLSVersionTlSv1_1, HTTPAseBotClassGetParamsTLSVersionTlSv1_2, HTTPAseBotClassGetParamsTLSVersionTlSv1_3, HTTPAseBotClassGetParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPAseBotClassGetResponseEnvelope struct {
	Result  HTTPAseBotClassGetResponse             `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    httpAseBotClassGetResponseEnvelopeJSON `json:"-"`
}

// httpAseBotClassGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPAseBotClassGetResponseEnvelope]
type httpAseBotClassGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPAseBotClassGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpAseBotClassGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
