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

// RadarHTTPTimeseryHTTPVersionService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseryHTTPVersionService] method instead.
type RadarHTTPTimeseryHTTPVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseryHTTPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseryHTTPVersionService(opts ...option.RequestOption) (r *RadarHTTPTimeseryHTTPVersionService) {
	r = &RadarHTTPTimeseryHTTPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic per HTTP protocol version over time.
func (r *RadarHTTPTimeseryHTTPVersionService) List(ctx context.Context, query RadarHTTPTimeseryHTTPVersionListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseryHTTPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseryHTTPVersionListResponse struct {
	Result  RadarHTTPTimeseryHTTPVersionListResponseResult `json:"result,required"`
	Success bool                                           `json:"success,required"`
	JSON    radarHTTPTimeseryHTTPVersionListResponseJSON   `json:"-"`
}

// radarHTTPTimeseryHTTPVersionListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseryHTTPVersionListResponse]
type radarHTTPTimeseryHTTPVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryHTTPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryHTTPVersionListResponseResult struct {
	Meta   interface{}                                          `json:"meta,required"`
	Serie0 RadarHTTPTimeseryHTTPVersionListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseryHTTPVersionListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseryHTTPVersionListResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseryHTTPVersionListResponseResult]
type radarHTTPTimeseryHTTPVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryHTTPVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryHTTPVersionListResponseResultSerie0 struct {
	HTTP1X     []string                                                 `json:"HTTP/1.x,required"`
	HTTP2      []string                                                 `json:"HTTP/2,required"`
	HTTP3      []string                                                 `json:"HTTP/3,required"`
	Timestamps []string                                                 `json:"timestamps,required"`
	JSON       radarHTTPTimeseryHTTPVersionListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseryHTTPVersionListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarHTTPTimeseryHTTPVersionListResponseResultSerie0]
type radarHTTPTimeseryHTTPVersionListResponseResultSerie0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryHTTPVersionListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryHTTPVersionListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseryHTTPVersionListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseryHTTPVersionListParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseryHTTPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseryHTTPVersionListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseryHTTPVersionListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseryHTTPVersionListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseryHTTPVersionListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseryHTTPVersionListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseryHTTPVersionListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseryHTTPVersionListParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTimeseryHTTPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseryHTTPVersionListParamsAggInterval string

const (
	RadarHTTPTimeseryHTTPVersionListParamsAggInterval15m RadarHTTPTimeseryHTTPVersionListParamsAggInterval = "15m"
	RadarHTTPTimeseryHTTPVersionListParamsAggInterval1h  RadarHTTPTimeseryHTTPVersionListParamsAggInterval = "1h"
	RadarHTTPTimeseryHTTPVersionListParamsAggInterval1d  RadarHTTPTimeseryHTTPVersionListParamsAggInterval = "1d"
	RadarHTTPTimeseryHTTPVersionListParamsAggInterval1w  RadarHTTPTimeseryHTTPVersionListParamsAggInterval = "1w"
)

type RadarHTTPTimeseryHTTPVersionListParamsBotClass string

const (
	RadarHTTPTimeseryHTTPVersionListParamsBotClassLikelyAutomated RadarHTTPTimeseryHTTPVersionListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseryHTTPVersionListParamsBotClassLikelyHuman     RadarHTTPTimeseryHTTPVersionListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseryHTTPVersionListParamsDateRange string

const (
	RadarHTTPTimeseryHTTPVersionListParamsDateRange1d         RadarHTTPTimeseryHTTPVersionListParamsDateRange = "1d"
	RadarHTTPTimeseryHTTPVersionListParamsDateRange7d         RadarHTTPTimeseryHTTPVersionListParamsDateRange = "7d"
	RadarHTTPTimeseryHTTPVersionListParamsDateRange14d        RadarHTTPTimeseryHTTPVersionListParamsDateRange = "14d"
	RadarHTTPTimeseryHTTPVersionListParamsDateRange28d        RadarHTTPTimeseryHTTPVersionListParamsDateRange = "28d"
	RadarHTTPTimeseryHTTPVersionListParamsDateRange12w        RadarHTTPTimeseryHTTPVersionListParamsDateRange = "12w"
	RadarHTTPTimeseryHTTPVersionListParamsDateRange24w        RadarHTTPTimeseryHTTPVersionListParamsDateRange = "24w"
	RadarHTTPTimeseryHTTPVersionListParamsDateRange52w        RadarHTTPTimeseryHTTPVersionListParamsDateRange = "52w"
	RadarHTTPTimeseryHTTPVersionListParamsDateRange1dControl  RadarHTTPTimeseryHTTPVersionListParamsDateRange = "1dControl"
	RadarHTTPTimeseryHTTPVersionListParamsDateRange7dControl  RadarHTTPTimeseryHTTPVersionListParamsDateRange = "7dControl"
	RadarHTTPTimeseryHTTPVersionListParamsDateRange14dControl RadarHTTPTimeseryHTTPVersionListParamsDateRange = "14dControl"
	RadarHTTPTimeseryHTTPVersionListParamsDateRange28dControl RadarHTTPTimeseryHTTPVersionListParamsDateRange = "28dControl"
	RadarHTTPTimeseryHTTPVersionListParamsDateRange12wControl RadarHTTPTimeseryHTTPVersionListParamsDateRange = "12wControl"
	RadarHTTPTimeseryHTTPVersionListParamsDateRange24wControl RadarHTTPTimeseryHTTPVersionListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseryHTTPVersionListParamsDeviceType string

const (
	RadarHTTPTimeseryHTTPVersionListParamsDeviceTypeDesktop RadarHTTPTimeseryHTTPVersionListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseryHTTPVersionListParamsDeviceTypeMobile  RadarHTTPTimeseryHTTPVersionListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseryHTTPVersionListParamsDeviceTypeOther   RadarHTTPTimeseryHTTPVersionListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseryHTTPVersionListParamsFormat string

const (
	RadarHTTPTimeseryHTTPVersionListParamsFormatJson RadarHTTPTimeseryHTTPVersionListParamsFormat = "JSON"
	RadarHTTPTimeseryHTTPVersionListParamsFormatCsv  RadarHTTPTimeseryHTTPVersionListParamsFormat = "CSV"
)

type RadarHTTPTimeseryHTTPVersionListParamsHTTPProtocol string

const (
	RadarHTTPTimeseryHTTPVersionListParamsHTTPProtocolHTTP  RadarHTTPTimeseryHTTPVersionListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseryHTTPVersionListParamsHTTPProtocolHTTPs RadarHTTPTimeseryHTTPVersionListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseryHTTPVersionListParamsIPVersion string

const (
	RadarHTTPTimeseryHTTPVersionListParamsIPVersionIPv4 RadarHTTPTimeseryHTTPVersionListParamsIPVersion = "IPv4"
	RadarHTTPTimeseryHTTPVersionListParamsIPVersionIPv6 RadarHTTPTimeseryHTTPVersionListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseryHTTPVersionListParamsO string

const (
	RadarHTTPTimeseryHTTPVersionListParamsOWindows  RadarHTTPTimeseryHTTPVersionListParamsO = "WINDOWS"
	RadarHTTPTimeseryHTTPVersionListParamsOMacosx   RadarHTTPTimeseryHTTPVersionListParamsO = "MACOSX"
	RadarHTTPTimeseryHTTPVersionListParamsOAndroid  RadarHTTPTimeseryHTTPVersionListParamsO = "ANDROID"
	RadarHTTPTimeseryHTTPVersionListParamsOChromeos RadarHTTPTimeseryHTTPVersionListParamsO = "CHROMEOS"
	RadarHTTPTimeseryHTTPVersionListParamsOLinux    RadarHTTPTimeseryHTTPVersionListParamsO = "LINUX"
	RadarHTTPTimeseryHTTPVersionListParamsOSmartTv  RadarHTTPTimeseryHTTPVersionListParamsO = "SMART_TV"
)

type RadarHTTPTimeseryHTTPVersionListParamsTlsVersion string

const (
	RadarHTTPTimeseryHTTPVersionListParamsTlsVersionTlSv1_0  RadarHTTPTimeseryHTTPVersionListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseryHTTPVersionListParamsTlsVersionTlSv1_1  RadarHTTPTimeseryHTTPVersionListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseryHTTPVersionListParamsTlsVersionTlSv1_2  RadarHTTPTimeseryHTTPVersionListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseryHTTPVersionListParamsTlsVersionTlSv1_3  RadarHTTPTimeseryHTTPVersionListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseryHTTPVersionListParamsTlsVersionTlSvQuic RadarHTTPTimeseryHTTPVersionListParamsTlsVersion = "TLSvQUIC"
)
