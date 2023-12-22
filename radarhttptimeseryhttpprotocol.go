// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// RadarHTTPTimeseryHTTPProtocolService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseryHTTPProtocolService] method instead.
type RadarHTTPTimeseryHTTPProtocolService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseryHTTPProtocolService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseryHTTPProtocolService(opts ...option.RequestOption) (r *RadarHTTPTimeseryHTTPProtocolService) {
	r = &RadarHTTPTimeseryHTTPProtocolService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic per HTTP protocol over time.
func (r *RadarHTTPTimeseryHTTPProtocolService) List(ctx context.Context, query RadarHTTPTimeseryHTTPProtocolListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseryHTTPProtocolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries/http_protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseryHTTPProtocolListResponse struct {
	Result  RadarHTTPTimeseryHTTPProtocolListResponseResult `json:"result,required"`
	Success bool                                            `json:"success,required"`
	JSON    radarHTTPTimeseryHTTPProtocolListResponseJSON   `json:"-"`
}

// radarHTTPTimeseryHTTPProtocolListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseryHTTPProtocolListResponse]
type radarHTTPTimeseryHTTPProtocolListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryHTTPProtocolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryHTTPProtocolListResponseResult struct {
	Meta   interface{}                                           `json:"meta,required"`
	Serie0 RadarHTTPTimeseryHTTPProtocolListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseryHTTPProtocolListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseryHTTPProtocolListResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseryHTTPProtocolListResponseResult]
type radarHTTPTimeseryHTTPProtocolListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryHTTPProtocolListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryHTTPProtocolListResponseResultSerie0 struct {
	HTTP       []string                                                  `json:"http,required"`
	HTTPs      []string                                                  `json:"https,required"`
	Timestamps []string                                                  `json:"timestamps,required"`
	JSON       radarHTTPTimeseryHTTPProtocolListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseryHTTPProtocolListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarHTTPTimeseryHTTPProtocolListResponseResultSerie0]
type radarHTTPTimeseryHTTPProtocolListResponseResultSerie0JSON struct {
	HTTP        apijson.Field
	HTTPs       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryHTTPProtocolListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryHTTPProtocolListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseryHTTPProtocolListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseryHTTPProtocolListParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseryHTTPProtocolListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseryHTTPProtocolListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseryHTTPProtocolListParamsFormat] `query:"format"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseryHTTPProtocolListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseryHTTPProtocolListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseryHTTPProtocolListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseryHTTPProtocolListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseryHTTPProtocolListParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTimeseryHTTPProtocolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseryHTTPProtocolListParamsAggInterval string

const (
	RadarHTTPTimeseryHTTPProtocolListParamsAggInterval15m RadarHTTPTimeseryHTTPProtocolListParamsAggInterval = "15m"
	RadarHTTPTimeseryHTTPProtocolListParamsAggInterval1h  RadarHTTPTimeseryHTTPProtocolListParamsAggInterval = "1h"
	RadarHTTPTimeseryHTTPProtocolListParamsAggInterval1d  RadarHTTPTimeseryHTTPProtocolListParamsAggInterval = "1d"
	RadarHTTPTimeseryHTTPProtocolListParamsAggInterval1w  RadarHTTPTimeseryHTTPProtocolListParamsAggInterval = "1w"
)

type RadarHTTPTimeseryHTTPProtocolListParamsBotClass string

const (
	RadarHTTPTimeseryHTTPProtocolListParamsBotClassLikelyAutomated RadarHTTPTimeseryHTTPProtocolListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseryHTTPProtocolListParamsBotClassLikelyHuman     RadarHTTPTimeseryHTTPProtocolListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseryHTTPProtocolListParamsDateRange string

const (
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange1d         RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "1d"
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange7d         RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "7d"
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange14d        RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "14d"
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange28d        RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "28d"
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange12w        RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "12w"
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange24w        RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "24w"
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange52w        RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "52w"
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange1dControl  RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "1dControl"
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange7dControl  RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "7dControl"
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange14dControl RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "14dControl"
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange28dControl RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "28dControl"
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange12wControl RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "12wControl"
	RadarHTTPTimeseryHTTPProtocolListParamsDateRange24wControl RadarHTTPTimeseryHTTPProtocolListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseryHTTPProtocolListParamsDeviceType string

const (
	RadarHTTPTimeseryHTTPProtocolListParamsDeviceTypeDesktop RadarHTTPTimeseryHTTPProtocolListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseryHTTPProtocolListParamsDeviceTypeMobile  RadarHTTPTimeseryHTTPProtocolListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseryHTTPProtocolListParamsDeviceTypeOther   RadarHTTPTimeseryHTTPProtocolListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseryHTTPProtocolListParamsFormat string

const (
	RadarHTTPTimeseryHTTPProtocolListParamsFormatJson RadarHTTPTimeseryHTTPProtocolListParamsFormat = "JSON"
	RadarHTTPTimeseryHTTPProtocolListParamsFormatCsv  RadarHTTPTimeseryHTTPProtocolListParamsFormat = "CSV"
)

type RadarHTTPTimeseryHTTPProtocolListParamsHTTPVersion string

const (
	RadarHTTPTimeseryHTTPProtocolListParamsHTTPVersionHttPv1 RadarHTTPTimeseryHTTPProtocolListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseryHTTPProtocolListParamsHTTPVersionHttPv2 RadarHTTPTimeseryHTTPProtocolListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseryHTTPProtocolListParamsHTTPVersionHttPv3 RadarHTTPTimeseryHTTPProtocolListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseryHTTPProtocolListParamsIPVersion string

const (
	RadarHTTPTimeseryHTTPProtocolListParamsIPVersionIPv4 RadarHTTPTimeseryHTTPProtocolListParamsIPVersion = "IPv4"
	RadarHTTPTimeseryHTTPProtocolListParamsIPVersionIPv6 RadarHTTPTimeseryHTTPProtocolListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseryHTTPProtocolListParamsO string

const (
	RadarHTTPTimeseryHTTPProtocolListParamsOWindows  RadarHTTPTimeseryHTTPProtocolListParamsO = "WINDOWS"
	RadarHTTPTimeseryHTTPProtocolListParamsOMacosx   RadarHTTPTimeseryHTTPProtocolListParamsO = "MACOSX"
	RadarHTTPTimeseryHTTPProtocolListParamsOAndroid  RadarHTTPTimeseryHTTPProtocolListParamsO = "ANDROID"
	RadarHTTPTimeseryHTTPProtocolListParamsOChromeos RadarHTTPTimeseryHTTPProtocolListParamsO = "CHROMEOS"
	RadarHTTPTimeseryHTTPProtocolListParamsOLinux    RadarHTTPTimeseryHTTPProtocolListParamsO = "LINUX"
	RadarHTTPTimeseryHTTPProtocolListParamsOSmartTv  RadarHTTPTimeseryHTTPProtocolListParamsO = "SMART_TV"
)

type RadarHTTPTimeseryHTTPProtocolListParamsTlsVersion string

const (
	RadarHTTPTimeseryHTTPProtocolListParamsTlsVersionTlSv1_0  RadarHTTPTimeseryHTTPProtocolListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseryHTTPProtocolListParamsTlsVersionTlSv1_1  RadarHTTPTimeseryHTTPProtocolListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseryHTTPProtocolListParamsTlsVersionTlSv1_2  RadarHTTPTimeseryHTTPProtocolListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseryHTTPProtocolListParamsTlsVersionTlSv1_3  RadarHTTPTimeseryHTTPProtocolListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseryHTTPProtocolListParamsTlsVersionTlSvQuic RadarHTTPTimeseryHTTPProtocolListParamsTlsVersion = "TLSvQUIC"
)
