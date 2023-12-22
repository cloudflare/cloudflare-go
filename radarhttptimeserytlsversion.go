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

// RadarHTTPTimeseryTlsVersionService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseryTlsVersionService] method instead.
type RadarHTTPTimeseryTlsVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseryTlsVersionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseryTlsVersionService(opts ...option.RequestOption) (r *RadarHTTPTimeseryTlsVersionService) {
	r = &RadarHTTPTimeseryTlsVersionService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic per TLS protocol version over time.
func (r *RadarHTTPTimeseryTlsVersionService) List(ctx context.Context, query RadarHTTPTimeseryTlsVersionListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseryTlsVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseryTlsVersionListResponse struct {
	Result  RadarHTTPTimeseryTlsVersionListResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarHTTPTimeseryTlsVersionListResponseJSON   `json:"-"`
}

// radarHTTPTimeseryTlsVersionListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseryTlsVersionListResponse]
type radarHTTPTimeseryTlsVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryTlsVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryTlsVersionListResponseResult struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarHTTPTimeseryTlsVersionListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseryTlsVersionListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseryTlsVersionListResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseryTlsVersionListResponseResult]
type radarHTTPTimeseryTlsVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryTlsVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryTlsVersionListResponseResultSerie0 struct {
	Timestamps []string                                                `json:"timestamps,required"`
	Tls1_0     []string                                                `json:"TLS 1.0,required"`
	Tls1_1     []string                                                `json:"TLS 1.1,required"`
	Tls1_2     []string                                                `json:"TLS 1.2,required"`
	Tls1_3     []string                                                `json:"TLS 1.3,required"`
	TlsQuic    []string                                                `json:"TLS QUIC,required"`
	JSON       radarHTTPTimeseryTlsVersionListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseryTlsVersionListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarHTTPTimeseryTlsVersionListResponseResultSerie0]
type radarHTTPTimeseryTlsVersionListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Tls1_0      apijson.Field
	Tls1_1      apijson.Field
	Tls1_2      apijson.Field
	Tls1_3      apijson.Field
	TlsQuic     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryTlsVersionListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryTlsVersionListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseryTlsVersionListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseryTlsVersionListParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseryTlsVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseryTlsVersionListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseryTlsVersionListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseryTlsVersionListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseryTlsVersionListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseryTlsVersionListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseryTlsVersionListParamsO] `query:"os"`
}

// URLQuery serializes [RadarHTTPTimeseryTlsVersionListParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTimeseryTlsVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseryTlsVersionListParamsAggInterval string

const (
	RadarHTTPTimeseryTlsVersionListParamsAggInterval15m RadarHTTPTimeseryTlsVersionListParamsAggInterval = "15m"
	RadarHTTPTimeseryTlsVersionListParamsAggInterval1h  RadarHTTPTimeseryTlsVersionListParamsAggInterval = "1h"
	RadarHTTPTimeseryTlsVersionListParamsAggInterval1d  RadarHTTPTimeseryTlsVersionListParamsAggInterval = "1d"
	RadarHTTPTimeseryTlsVersionListParamsAggInterval1w  RadarHTTPTimeseryTlsVersionListParamsAggInterval = "1w"
)

type RadarHTTPTimeseryTlsVersionListParamsBotClass string

const (
	RadarHTTPTimeseryTlsVersionListParamsBotClassLikelyAutomated RadarHTTPTimeseryTlsVersionListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseryTlsVersionListParamsBotClassLikelyHuman     RadarHTTPTimeseryTlsVersionListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseryTlsVersionListParamsDateRange string

const (
	RadarHTTPTimeseryTlsVersionListParamsDateRange1d         RadarHTTPTimeseryTlsVersionListParamsDateRange = "1d"
	RadarHTTPTimeseryTlsVersionListParamsDateRange7d         RadarHTTPTimeseryTlsVersionListParamsDateRange = "7d"
	RadarHTTPTimeseryTlsVersionListParamsDateRange14d        RadarHTTPTimeseryTlsVersionListParamsDateRange = "14d"
	RadarHTTPTimeseryTlsVersionListParamsDateRange28d        RadarHTTPTimeseryTlsVersionListParamsDateRange = "28d"
	RadarHTTPTimeseryTlsVersionListParamsDateRange12w        RadarHTTPTimeseryTlsVersionListParamsDateRange = "12w"
	RadarHTTPTimeseryTlsVersionListParamsDateRange24w        RadarHTTPTimeseryTlsVersionListParamsDateRange = "24w"
	RadarHTTPTimeseryTlsVersionListParamsDateRange52w        RadarHTTPTimeseryTlsVersionListParamsDateRange = "52w"
	RadarHTTPTimeseryTlsVersionListParamsDateRange1dControl  RadarHTTPTimeseryTlsVersionListParamsDateRange = "1dControl"
	RadarHTTPTimeseryTlsVersionListParamsDateRange7dControl  RadarHTTPTimeseryTlsVersionListParamsDateRange = "7dControl"
	RadarHTTPTimeseryTlsVersionListParamsDateRange14dControl RadarHTTPTimeseryTlsVersionListParamsDateRange = "14dControl"
	RadarHTTPTimeseryTlsVersionListParamsDateRange28dControl RadarHTTPTimeseryTlsVersionListParamsDateRange = "28dControl"
	RadarHTTPTimeseryTlsVersionListParamsDateRange12wControl RadarHTTPTimeseryTlsVersionListParamsDateRange = "12wControl"
	RadarHTTPTimeseryTlsVersionListParamsDateRange24wControl RadarHTTPTimeseryTlsVersionListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseryTlsVersionListParamsDeviceType string

const (
	RadarHTTPTimeseryTlsVersionListParamsDeviceTypeDesktop RadarHTTPTimeseryTlsVersionListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseryTlsVersionListParamsDeviceTypeMobile  RadarHTTPTimeseryTlsVersionListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseryTlsVersionListParamsDeviceTypeOther   RadarHTTPTimeseryTlsVersionListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseryTlsVersionListParamsFormat string

const (
	RadarHTTPTimeseryTlsVersionListParamsFormatJson RadarHTTPTimeseryTlsVersionListParamsFormat = "JSON"
	RadarHTTPTimeseryTlsVersionListParamsFormatCsv  RadarHTTPTimeseryTlsVersionListParamsFormat = "CSV"
)

type RadarHTTPTimeseryTlsVersionListParamsHTTPProtocol string

const (
	RadarHTTPTimeseryTlsVersionListParamsHTTPProtocolHTTP  RadarHTTPTimeseryTlsVersionListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseryTlsVersionListParamsHTTPProtocolHTTPs RadarHTTPTimeseryTlsVersionListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseryTlsVersionListParamsHTTPVersion string

const (
	RadarHTTPTimeseryTlsVersionListParamsHTTPVersionHttPv1 RadarHTTPTimeseryTlsVersionListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseryTlsVersionListParamsHTTPVersionHttPv2 RadarHTTPTimeseryTlsVersionListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseryTlsVersionListParamsHTTPVersionHttPv3 RadarHTTPTimeseryTlsVersionListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseryTlsVersionListParamsIPVersion string

const (
	RadarHTTPTimeseryTlsVersionListParamsIPVersionIPv4 RadarHTTPTimeseryTlsVersionListParamsIPVersion = "IPv4"
	RadarHTTPTimeseryTlsVersionListParamsIPVersionIPv6 RadarHTTPTimeseryTlsVersionListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseryTlsVersionListParamsO string

const (
	RadarHTTPTimeseryTlsVersionListParamsOWindows  RadarHTTPTimeseryTlsVersionListParamsO = "WINDOWS"
	RadarHTTPTimeseryTlsVersionListParamsOMacosx   RadarHTTPTimeseryTlsVersionListParamsO = "MACOSX"
	RadarHTTPTimeseryTlsVersionListParamsOAndroid  RadarHTTPTimeseryTlsVersionListParamsO = "ANDROID"
	RadarHTTPTimeseryTlsVersionListParamsOChromeos RadarHTTPTimeseryTlsVersionListParamsO = "CHROMEOS"
	RadarHTTPTimeseryTlsVersionListParamsOLinux    RadarHTTPTimeseryTlsVersionListParamsO = "LINUX"
	RadarHTTPTimeseryTlsVersionListParamsOSmartTv  RadarHTTPTimeseryTlsVersionListParamsO = "SMART_TV"
)
