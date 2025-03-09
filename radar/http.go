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

// HTTPService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHTTPService] method instead.
type HTTPService struct {
	Options          []option.RequestOption
	Locations        *HTTPLocationService
	Ases             *HTTPAseService
	Summary          *HTTPSummaryService
	TimeseriesGroups *HTTPTimeseriesGroupService
	Top              *HTTPTopService
}

// NewHTTPService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHTTPService(opts ...option.RequestOption) (r *HTTPService) {
	r = &HTTPService{}
	r.Options = opts
	r.Locations = NewHTTPLocationService(opts...)
	r.Ases = NewHTTPAseService(opts...)
	r.Summary = NewHTTPSummaryService(opts...)
	r.TimeseriesGroups = NewHTTPTimeseriesGroupService(opts...)
	r.Top = NewHTTPTopService(opts...)
	return
}

// Retrieves the HTTP requests over time.
func (r *HTTPService) Timeseries(ctx context.Context, query HTTPTimeseriesParams, opts ...option.RequestOption) (res *HTTPTimeseriesResponse, err error) {
	var env HTTPTimeseriesResponseEnvelope
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type HTTPTimeseriesResponse struct {
	Meta   HTTPTimeseriesResponseMeta   `json:"meta,required"`
	Serie0 HTTPTimeseriesResponseSerie0 `json:"serie_0,required"`
	JSON   httpTimeseriesResponseJSON   `json:"-"`
}

// httpTimeseriesResponseJSON contains the JSON metadata for the struct
// [HTTPTimeseriesResponse]
type httpTimeseriesResponseJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesResponseMeta struct {
	AggInterval    string                                   `json:"aggInterval,required"`
	DateRange      []HTTPTimeseriesResponseMetaDateRange    `json:"dateRange,required"`
	LastUpdated    time.Time                                `json:"lastUpdated,required" format:"date-time"`
	ConfidenceInfo HTTPTimeseriesResponseMetaConfidenceInfo `json:"confidenceInfo"`
	JSON           httpTimeseriesResponseMetaJSON           `json:"-"`
}

// httpTimeseriesResponseMetaJSON contains the JSON metadata for the struct
// [HTTPTimeseriesResponseMeta]
type httpTimeseriesResponseMetaJSON struct {
	AggInterval    apijson.Field
	DateRange      apijson.Field
	LastUpdated    apijson.Field
	ConfidenceInfo apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *HTTPTimeseriesResponseMeta) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseMetaJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesResponseMetaDateRange struct {
	// Adjusted end of date range.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Adjusted start of date range.
	StartTime time.Time                               `json:"startTime,required" format:"date-time"`
	JSON      httpTimeseriesResponseMetaDateRangeJSON `json:"-"`
}

// httpTimeseriesResponseMetaDateRangeJSON contains the JSON metadata for the
// struct [HTTPTimeseriesResponseMetaDateRange]
type httpTimeseriesResponseMetaDateRangeJSON struct {
	EndTime     apijson.Field
	StartTime   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesResponseMetaDateRange) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseMetaDateRangeJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesResponseMetaConfidenceInfo struct {
	Annotations []HTTPTimeseriesResponseMetaConfidenceInfoAnnotation `json:"annotations"`
	Level       int64                                                `json:"level"`
	JSON        httpTimeseriesResponseMetaConfidenceInfoJSON         `json:"-"`
}

// httpTimeseriesResponseMetaConfidenceInfoJSON contains the JSON metadata for the
// struct [HTTPTimeseriesResponseMetaConfidenceInfo]
type httpTimeseriesResponseMetaConfidenceInfoJSON struct {
	Annotations apijson.Field
	Level       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesResponseMetaConfidenceInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseMetaConfidenceInfoJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesResponseMetaConfidenceInfoAnnotation struct {
	DataSource      string                                                 `json:"dataSource,required"`
	Description     string                                                 `json:"description,required"`
	EventType       string                                                 `json:"eventType,required"`
	IsInstantaneous bool                                                   `json:"isInstantaneous,required"`
	EndTime         time.Time                                              `json:"endTime" format:"date-time"`
	LinkedURL       string                                                 `json:"linkedUrl"`
	StartTime       time.Time                                              `json:"startTime" format:"date-time"`
	JSON            httpTimeseriesResponseMetaConfidenceInfoAnnotationJSON `json:"-"`
}

// httpTimeseriesResponseMetaConfidenceInfoAnnotationJSON contains the JSON
// metadata for the struct [HTTPTimeseriesResponseMetaConfidenceInfoAnnotation]
type httpTimeseriesResponseMetaConfidenceInfoAnnotationJSON struct {
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

func (r *HTTPTimeseriesResponseMetaConfidenceInfoAnnotation) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseMetaConfidenceInfoAnnotationJSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesResponseSerie0 struct {
	Timestamps []time.Time                      `json:"timestamps,required" format:"date-time"`
	Values     []string                         `json:"values,required"`
	JSON       httpTimeseriesResponseSerie0JSON `json:"-"`
}

// httpTimeseriesResponseSerie0JSON contains the JSON metadata for the struct
// [HTTPTimeseriesResponseSerie0]
type httpTimeseriesResponseSerie0JSON struct {
	Timestamps  apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesResponseSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseSerie0JSON) RawJSON() string {
	return r.raw
}

type HTTPTimeseriesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[HTTPTimeseriesParamsAggInterval] `query:"aggInterval"`
	// Comma-separated list of Autonomous System Numbers (ASNs). Prefix with `-` to
	// exclude ASNs from results. For example, `-174, 3356` excludes results from
	// AS174, but includes results from AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filters results by bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]HTTPTimeseriesParamsBotClass] `query:"botClass"`
	// Comma-separated list of continents (alpha-2 continent codes). Prefix with `-` to
	// exclude continents from results. For example, `-EU,NA` excludes results from EU,
	// but includes results from NA.
	Continent param.Field[[]string] `query:"continent"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// Filters results by the specified date range. For example, use `7d` and
	// `7dcontrol` to compare this week with the previous week. Use this parameter or
	// set specific start and end dates (`dateStart` and `dateEnd` parameters).
	DateRange param.Field[[]string] `query:"dateRange"`
	// Start of the date range.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filters results by device type.
	DeviceType param.Field[[]HTTPTimeseriesParamsDeviceType] `query:"deviceType"`
	// Format in which results will be returned.
	Format param.Field[HTTPTimeseriesParamsFormat] `query:"format"`
	// Filters results by HTTP protocol (HTTP vs. HTTPS).
	HTTPProtocol param.Field[[]HTTPTimeseriesParamsHTTPProtocol] `query:"httpProtocol"`
	// Filters results by HTTP version.
	HTTPVersion param.Field[[]HTTPTimeseriesParamsHTTPVersion] `query:"httpVersion"`
	// Filters results by IP version (Ipv4 vs. IPv6).
	IPVersion param.Field[[]HTTPTimeseriesParamsIPVersion] `query:"ipVersion"`
	// Comma-separated list of locations (alpha-2 codes). Prefix with `-` to exclude
	// locations from results. For example, `-US,PT` excludes results from the US, but
	// includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names used to label the series in the response.
	Name param.Field[[]string] `query:"name"`
	// Normalization method applied. Refer to
	// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
	Normalization param.Field[HTTPTimeseriesParamsNormalization] `query:"normalization"`
	// Filters results by operating system.
	OS param.Field[[]HTTPTimeseriesParamsOS] `query:"os"`
	// Filters results by TLS version.
	TLSVersion param.Field[[]HTTPTimeseriesParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [HTTPTimeseriesParams]'s query parameters as `url.Values`.
func (r HTTPTimeseriesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type HTTPTimeseriesParamsAggInterval string

const (
	HTTPTimeseriesParamsAggInterval15m HTTPTimeseriesParamsAggInterval = "15m"
	HTTPTimeseriesParamsAggInterval1h  HTTPTimeseriesParamsAggInterval = "1h"
	HTTPTimeseriesParamsAggInterval1d  HTTPTimeseriesParamsAggInterval = "1d"
	HTTPTimeseriesParamsAggInterval1w  HTTPTimeseriesParamsAggInterval = "1w"
)

func (r HTTPTimeseriesParamsAggInterval) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsAggInterval15m, HTTPTimeseriesParamsAggInterval1h, HTTPTimeseriesParamsAggInterval1d, HTTPTimeseriesParamsAggInterval1w:
		return true
	}
	return false
}

type HTTPTimeseriesParamsBotClass string

const (
	HTTPTimeseriesParamsBotClassLikelyAutomated HTTPTimeseriesParamsBotClass = "LIKELY_AUTOMATED"
	HTTPTimeseriesParamsBotClassLikelyHuman     HTTPTimeseriesParamsBotClass = "LIKELY_HUMAN"
)

func (r HTTPTimeseriesParamsBotClass) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsBotClassLikelyAutomated, HTTPTimeseriesParamsBotClassLikelyHuman:
		return true
	}
	return false
}

type HTTPTimeseriesParamsDeviceType string

const (
	HTTPTimeseriesParamsDeviceTypeDesktop HTTPTimeseriesParamsDeviceType = "DESKTOP"
	HTTPTimeseriesParamsDeviceTypeMobile  HTTPTimeseriesParamsDeviceType = "MOBILE"
	HTTPTimeseriesParamsDeviceTypeOther   HTTPTimeseriesParamsDeviceType = "OTHER"
)

func (r HTTPTimeseriesParamsDeviceType) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsDeviceTypeDesktop, HTTPTimeseriesParamsDeviceTypeMobile, HTTPTimeseriesParamsDeviceTypeOther:
		return true
	}
	return false
}

// Format in which results will be returned.
type HTTPTimeseriesParamsFormat string

const (
	HTTPTimeseriesParamsFormatJson HTTPTimeseriesParamsFormat = "JSON"
	HTTPTimeseriesParamsFormatCsv  HTTPTimeseriesParamsFormat = "CSV"
)

func (r HTTPTimeseriesParamsFormat) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsFormatJson, HTTPTimeseriesParamsFormatCsv:
		return true
	}
	return false
}

type HTTPTimeseriesParamsHTTPProtocol string

const (
	HTTPTimeseriesParamsHTTPProtocolHTTP  HTTPTimeseriesParamsHTTPProtocol = "HTTP"
	HTTPTimeseriesParamsHTTPProtocolHTTPS HTTPTimeseriesParamsHTTPProtocol = "HTTPS"
)

func (r HTTPTimeseriesParamsHTTPProtocol) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsHTTPProtocolHTTP, HTTPTimeseriesParamsHTTPProtocolHTTPS:
		return true
	}
	return false
}

type HTTPTimeseriesParamsHTTPVersion string

const (
	HTTPTimeseriesParamsHTTPVersionHttPv1 HTTPTimeseriesParamsHTTPVersion = "HTTPv1"
	HTTPTimeseriesParamsHTTPVersionHttPv2 HTTPTimeseriesParamsHTTPVersion = "HTTPv2"
	HTTPTimeseriesParamsHTTPVersionHttPv3 HTTPTimeseriesParamsHTTPVersion = "HTTPv3"
)

func (r HTTPTimeseriesParamsHTTPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsHTTPVersionHttPv1, HTTPTimeseriesParamsHTTPVersionHttPv2, HTTPTimeseriesParamsHTTPVersionHttPv3:
		return true
	}
	return false
}

type HTTPTimeseriesParamsIPVersion string

const (
	HTTPTimeseriesParamsIPVersionIPv4 HTTPTimeseriesParamsIPVersion = "IPv4"
	HTTPTimeseriesParamsIPVersionIPv6 HTTPTimeseriesParamsIPVersion = "IPv6"
)

func (r HTTPTimeseriesParamsIPVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsIPVersionIPv4, HTTPTimeseriesParamsIPVersionIPv6:
		return true
	}
	return false
}

// Normalization method applied. Refer to
// [Normalization methods](https://developers.cloudflare.com/radar/concepts/normalization/).
type HTTPTimeseriesParamsNormalization string

const (
	HTTPTimeseriesParamsNormalizationPercentageChange HTTPTimeseriesParamsNormalization = "PERCENTAGE_CHANGE"
	HTTPTimeseriesParamsNormalizationMin0Max          HTTPTimeseriesParamsNormalization = "MIN0_MAX"
)

func (r HTTPTimeseriesParamsNormalization) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsNormalizationPercentageChange, HTTPTimeseriesParamsNormalizationMin0Max:
		return true
	}
	return false
}

type HTTPTimeseriesParamsOS string

const (
	HTTPTimeseriesParamsOSWindows  HTTPTimeseriesParamsOS = "WINDOWS"
	HTTPTimeseriesParamsOSMacosx   HTTPTimeseriesParamsOS = "MACOSX"
	HTTPTimeseriesParamsOSIos      HTTPTimeseriesParamsOS = "IOS"
	HTTPTimeseriesParamsOSAndroid  HTTPTimeseriesParamsOS = "ANDROID"
	HTTPTimeseriesParamsOSChromeos HTTPTimeseriesParamsOS = "CHROMEOS"
	HTTPTimeseriesParamsOSLinux    HTTPTimeseriesParamsOS = "LINUX"
	HTTPTimeseriesParamsOSSmartTv  HTTPTimeseriesParamsOS = "SMART_TV"
)

func (r HTTPTimeseriesParamsOS) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsOSWindows, HTTPTimeseriesParamsOSMacosx, HTTPTimeseriesParamsOSIos, HTTPTimeseriesParamsOSAndroid, HTTPTimeseriesParamsOSChromeos, HTTPTimeseriesParamsOSLinux, HTTPTimeseriesParamsOSSmartTv:
		return true
	}
	return false
}

type HTTPTimeseriesParamsTLSVersion string

const (
	HTTPTimeseriesParamsTLSVersionTlSv1_0  HTTPTimeseriesParamsTLSVersion = "TLSv1_0"
	HTTPTimeseriesParamsTLSVersionTlSv1_1  HTTPTimeseriesParamsTLSVersion = "TLSv1_1"
	HTTPTimeseriesParamsTLSVersionTlSv1_2  HTTPTimeseriesParamsTLSVersion = "TLSv1_2"
	HTTPTimeseriesParamsTLSVersionTlSv1_3  HTTPTimeseriesParamsTLSVersion = "TLSv1_3"
	HTTPTimeseriesParamsTLSVersionTlSvQuic HTTPTimeseriesParamsTLSVersion = "TLSvQUIC"
)

func (r HTTPTimeseriesParamsTLSVersion) IsKnown() bool {
	switch r {
	case HTTPTimeseriesParamsTLSVersionTlSv1_0, HTTPTimeseriesParamsTLSVersionTlSv1_1, HTTPTimeseriesParamsTLSVersionTlSv1_2, HTTPTimeseriesParamsTLSVersionTlSv1_3, HTTPTimeseriesParamsTLSVersionTlSvQuic:
		return true
	}
	return false
}

type HTTPTimeseriesResponseEnvelope struct {
	Result  HTTPTimeseriesResponse             `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    httpTimeseriesResponseEnvelopeJSON `json:"-"`
}

// httpTimeseriesResponseEnvelopeJSON contains the JSON metadata for the struct
// [HTTPTimeseriesResponseEnvelope]
type httpTimeseriesResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HTTPTimeseriesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r httpTimeseriesResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}
