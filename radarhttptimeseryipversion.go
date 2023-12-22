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

// RadarHTTPTimeseryIPVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseryIPVersionService] method instead.
type RadarHTTPTimeseryIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseryIPVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseryIPVersionService(opts ...option.RequestOption) (r *RadarHTTPTimeseryIPVersionService) {
	r = &RadarHTTPTimeseryIPVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic per IP protocol version over time.
func (r *RadarHTTPTimeseryIPVersionService) List(ctx context.Context, query RadarHTTPTimeseryIPVersionListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseryIPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseryIPVersionListResponse struct {
	Result  RadarHTTPTimeseryIPVersionListResponseResult `json:"result,required"`
	Success bool                                         `json:"success,required"`
	JSON    radarHTTPTimeseryIPVersionListResponseJSON   `json:"-"`
}

// radarHTTPTimeseryIPVersionListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseryIPVersionListResponse]
type radarHTTPTimeseryIPVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryIPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryIPVersionListResponseResult struct {
	Meta   interface{}                                        `json:"meta,required"`
	Serie0 RadarHTTPTimeseryIPVersionListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseryIPVersionListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseryIPVersionListResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseryIPVersionListResponseResult]
type radarHTTPTimeseryIPVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryIPVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryIPVersionListResponseResultSerie0 struct {
	IPv4       []string                                               `json:"IPv4,required"`
	IPv6       []string                                               `json:"IPv6,required"`
	Timestamps []string                                               `json:"timestamps,required"`
	JSON       radarHTTPTimeseryIPVersionListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseryIPVersionListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarHTTPTimeseryIPVersionListResponseResultSerie0]
type radarHTTPTimeseryIPVersionListResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryIPVersionListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryIPVersionListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseryIPVersionListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseryIPVersionListParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseryIPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseryIPVersionListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseryIPVersionListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseryIPVersionListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseryIPVersionListParamsHTTPVersion] `query:"httpVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseryIPVersionListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseryIPVersionListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseryIPVersionListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTimeseryIPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseryIPVersionListParamsAggInterval string

const (
	RadarHTTPTimeseryIPVersionListParamsAggInterval15m RadarHTTPTimeseryIPVersionListParamsAggInterval = "15m"
	RadarHTTPTimeseryIPVersionListParamsAggInterval1h  RadarHTTPTimeseryIPVersionListParamsAggInterval = "1h"
	RadarHTTPTimeseryIPVersionListParamsAggInterval1d  RadarHTTPTimeseryIPVersionListParamsAggInterval = "1d"
	RadarHTTPTimeseryIPVersionListParamsAggInterval1w  RadarHTTPTimeseryIPVersionListParamsAggInterval = "1w"
)

type RadarHTTPTimeseryIPVersionListParamsBotClass string

const (
	RadarHTTPTimeseryIPVersionListParamsBotClassLikelyAutomated RadarHTTPTimeseryIPVersionListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseryIPVersionListParamsBotClassLikelyHuman     RadarHTTPTimeseryIPVersionListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseryIPVersionListParamsDateRange string

const (
	RadarHTTPTimeseryIPVersionListParamsDateRange1d         RadarHTTPTimeseryIPVersionListParamsDateRange = "1d"
	RadarHTTPTimeseryIPVersionListParamsDateRange7d         RadarHTTPTimeseryIPVersionListParamsDateRange = "7d"
	RadarHTTPTimeseryIPVersionListParamsDateRange14d        RadarHTTPTimeseryIPVersionListParamsDateRange = "14d"
	RadarHTTPTimeseryIPVersionListParamsDateRange28d        RadarHTTPTimeseryIPVersionListParamsDateRange = "28d"
	RadarHTTPTimeseryIPVersionListParamsDateRange12w        RadarHTTPTimeseryIPVersionListParamsDateRange = "12w"
	RadarHTTPTimeseryIPVersionListParamsDateRange24w        RadarHTTPTimeseryIPVersionListParamsDateRange = "24w"
	RadarHTTPTimeseryIPVersionListParamsDateRange52w        RadarHTTPTimeseryIPVersionListParamsDateRange = "52w"
	RadarHTTPTimeseryIPVersionListParamsDateRange1dControl  RadarHTTPTimeseryIPVersionListParamsDateRange = "1dControl"
	RadarHTTPTimeseryIPVersionListParamsDateRange7dControl  RadarHTTPTimeseryIPVersionListParamsDateRange = "7dControl"
	RadarHTTPTimeseryIPVersionListParamsDateRange14dControl RadarHTTPTimeseryIPVersionListParamsDateRange = "14dControl"
	RadarHTTPTimeseryIPVersionListParamsDateRange28dControl RadarHTTPTimeseryIPVersionListParamsDateRange = "28dControl"
	RadarHTTPTimeseryIPVersionListParamsDateRange12wControl RadarHTTPTimeseryIPVersionListParamsDateRange = "12wControl"
	RadarHTTPTimeseryIPVersionListParamsDateRange24wControl RadarHTTPTimeseryIPVersionListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseryIPVersionListParamsDeviceType string

const (
	RadarHTTPTimeseryIPVersionListParamsDeviceTypeDesktop RadarHTTPTimeseryIPVersionListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseryIPVersionListParamsDeviceTypeMobile  RadarHTTPTimeseryIPVersionListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseryIPVersionListParamsDeviceTypeOther   RadarHTTPTimeseryIPVersionListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseryIPVersionListParamsFormat string

const (
	RadarHTTPTimeseryIPVersionListParamsFormatJson RadarHTTPTimeseryIPVersionListParamsFormat = "JSON"
	RadarHTTPTimeseryIPVersionListParamsFormatCsv  RadarHTTPTimeseryIPVersionListParamsFormat = "CSV"
)

type RadarHTTPTimeseryIPVersionListParamsHTTPProtocol string

const (
	RadarHTTPTimeseryIPVersionListParamsHTTPProtocolHTTP  RadarHTTPTimeseryIPVersionListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseryIPVersionListParamsHTTPProtocolHTTPs RadarHTTPTimeseryIPVersionListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseryIPVersionListParamsHTTPVersion string

const (
	RadarHTTPTimeseryIPVersionListParamsHTTPVersionHttPv1 RadarHTTPTimeseryIPVersionListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseryIPVersionListParamsHTTPVersionHttPv2 RadarHTTPTimeseryIPVersionListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseryIPVersionListParamsHTTPVersionHttPv3 RadarHTTPTimeseryIPVersionListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseryIPVersionListParamsO string

const (
	RadarHTTPTimeseryIPVersionListParamsOWindows  RadarHTTPTimeseryIPVersionListParamsO = "WINDOWS"
	RadarHTTPTimeseryIPVersionListParamsOMacosx   RadarHTTPTimeseryIPVersionListParamsO = "MACOSX"
	RadarHTTPTimeseryIPVersionListParamsOAndroid  RadarHTTPTimeseryIPVersionListParamsO = "ANDROID"
	RadarHTTPTimeseryIPVersionListParamsOChromeos RadarHTTPTimeseryIPVersionListParamsO = "CHROMEOS"
	RadarHTTPTimeseryIPVersionListParamsOLinux    RadarHTTPTimeseryIPVersionListParamsO = "LINUX"
	RadarHTTPTimeseryIPVersionListParamsOSmartTv  RadarHTTPTimeseryIPVersionListParamsO = "SMART_TV"
)

type RadarHTTPTimeseryIPVersionListParamsTlsVersion string

const (
	RadarHTTPTimeseryIPVersionListParamsTlsVersionTlSv1_0  RadarHTTPTimeseryIPVersionListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseryIPVersionListParamsTlsVersionTlSv1_1  RadarHTTPTimeseryIPVersionListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseryIPVersionListParamsTlsVersionTlSv1_2  RadarHTTPTimeseryIPVersionListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseryIPVersionListParamsTlsVersionTlSv1_3  RadarHTTPTimeseryIPVersionListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseryIPVersionListParamsTlsVersionTlSvQuic RadarHTTPTimeseryIPVersionListParamsTlsVersion = "TLSvQUIC"
)
