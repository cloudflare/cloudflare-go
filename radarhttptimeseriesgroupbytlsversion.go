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

// RadarHTTPTimeseriesGroupByTlsVersionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseriesGroupByTlsVersionService] method instead.
type RadarHTTPTimeseriesGroupByTlsVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseriesGroupByTlsVersionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarHTTPTimeseriesGroupByTlsVersionService(opts ...option.RequestOption) (r *RadarHTTPTimeseriesGroupByTlsVersionService) {
	r = &RadarHTTPTimeseriesGroupByTlsVersionService{}
	r.Options = opts
	return
}

// Get a time series of the percentage distribution of traffic per TLS protocol
// version.
func (r *RadarHTTPTimeseriesGroupByTlsVersionService) List(ctx context.Context, query RadarHTTPTimeseriesGroupByTlsVersionListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupByTlsVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/tls_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseriesGroupByTlsVersionListResponse struct {
	Result  RadarHTTPTimeseriesGroupByTlsVersionListResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupByTlsVersionListResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByTlsVersionListResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupByTlsVersionListResponse]
type radarHTTPTimeseriesGroupByTlsVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByTlsVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByTlsVersionListResponseResult struct {
	Meta   interface{}                                                  `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupByTlsVersionListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupByTlsVersionListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByTlsVersionListResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupByTlsVersionListResponseResult]
type radarHTTPTimeseriesGroupByTlsVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByTlsVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByTlsVersionListResponseResultSerie0 struct {
	Timestamps []string                                                         `json:"timestamps,required"`
	Tls1_0     []string                                                         `json:"TLS 1.0,required"`
	Tls1_1     []string                                                         `json:"TLS 1.1,required"`
	Tls1_2     []string                                                         `json:"TLS 1.2,required"`
	Tls1_3     []string                                                         `json:"TLS 1.3,required"`
	TlsQuic    []string                                                         `json:"TLS QUIC,required"`
	JSON       radarHTTPTimeseriesGroupByTlsVersionListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupByTlsVersionListResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarHTTPTimeseriesGroupByTlsVersionListResponseResultSerie0]
type radarHTTPTimeseriesGroupByTlsVersionListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	Tls1_0      apijson.Field
	Tls1_1      apijson.Field
	Tls1_2      apijson.Field
	Tls1_3      apijson.Field
	TlsQuic     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByTlsVersionListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByTlsVersionListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupByTlsVersionListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupByTlsVersionListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupByTlsVersionListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupByTlsVersionListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupByTlsVersionListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseriesGroupByTlsVersionListParamsO] `query:"os"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupByTlsVersionListParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupByTlsVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupByTlsVersionListParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupByTlsVersionListParamsAggInterval15m RadarHTTPTimeseriesGroupByTlsVersionListParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsAggInterval1h  RadarHTTPTimeseriesGroupByTlsVersionListParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsAggInterval1d  RadarHTTPTimeseriesGroupByTlsVersionListParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsAggInterval1w  RadarHTTPTimeseriesGroupByTlsVersionListParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupByTlsVersionListParamsBotClass string

const (
	RadarHTTPTimeseriesGroupByTlsVersionListParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupByTlsVersionListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupByTlsVersionListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange string

const (
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange1d         RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange2d         RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange7d         RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange14d        RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange28d        RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange12w        RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange24w        RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange52w        RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange1dControl  RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange2dControl  RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange7dControl  RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange14dControl RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange28dControl RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange12wControl RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange24wControl RadarHTTPTimeseriesGroupByTlsVersionListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupByTlsVersionListParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupByTlsVersionListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupByTlsVersionListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsDeviceTypeOther   RadarHTTPTimeseriesGroupByTlsVersionListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupByTlsVersionListParamsFormat string

const (
	RadarHTTPTimeseriesGroupByTlsVersionListParamsFormatJson RadarHTTPTimeseriesGroupByTlsVersionListParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsFormatCsv  RadarHTTPTimeseriesGroupByTlsVersionListParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPProtocolHTTPs RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupByTlsVersionListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupByTlsVersionListParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupByTlsVersionListParamsIPVersionIPv4 RadarHTTPTimeseriesGroupByTlsVersionListParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsIPVersionIPv6 RadarHTTPTimeseriesGroupByTlsVersionListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupByTlsVersionListParamsO string

const (
	RadarHTTPTimeseriesGroupByTlsVersionListParamsOWindows  RadarHTTPTimeseriesGroupByTlsVersionListParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsOMacosx   RadarHTTPTimeseriesGroupByTlsVersionListParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsOIos      RadarHTTPTimeseriesGroupByTlsVersionListParamsO = "IOS"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsOAndroid  RadarHTTPTimeseriesGroupByTlsVersionListParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsOChromeos RadarHTTPTimeseriesGroupByTlsVersionListParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsOLinux    RadarHTTPTimeseriesGroupByTlsVersionListParamsO = "LINUX"
	RadarHTTPTimeseriesGroupByTlsVersionListParamsOSmartTv  RadarHTTPTimeseriesGroupByTlsVersionListParamsO = "SMART_TV"
)
