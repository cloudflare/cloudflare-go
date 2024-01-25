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

// RadarHTTPTimeseriesGroupByIPVersionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseriesGroupByIPVersionService] method instead.
type RadarHTTPTimeseriesGroupByIPVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseriesGroupByIPVersionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarHTTPTimeseriesGroupByIPVersionService(opts ...option.RequestOption) (r *RadarHTTPTimeseriesGroupByIPVersionService) {
	r = &RadarHTTPTimeseriesGroupByIPVersionService{}
	r.Options = opts
	return
}

// Get a time series of the percentage distribution of traffic per IP protocol
// version.
func (r *RadarHTTPTimeseriesGroupByIPVersionService) List(ctx context.Context, query RadarHTTPTimeseriesGroupByIPVersionListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupByIPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseriesGroupByIPVersionListResponse struct {
	Result  RadarHTTPTimeseriesGroupByIPVersionListResponseResult `json:"result,required"`
	Success bool                                                  `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupByIPVersionListResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByIPVersionListResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupByIPVersionListResponse]
type radarHTTPTimeseriesGroupByIPVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByIPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByIPVersionListResponseResult struct {
	Meta   interface{}                                                 `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupByIPVersionListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupByIPVersionListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByIPVersionListResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupByIPVersionListResponseResult]
type radarHTTPTimeseriesGroupByIPVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByIPVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByIPVersionListResponseResultSerie0 struct {
	IPv4       []string                                                        `json:"IPv4,required"`
	IPv6       []string                                                        `json:"IPv6,required"`
	Timestamps []string                                                        `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupByIPVersionListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupByIPVersionListResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarHTTPTimeseriesGroupByIPVersionListResponseResultSerie0]
type radarHTTPTimeseriesGroupByIPVersionListResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByIPVersionListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByIPVersionListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupByIPVersionListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupByIPVersionListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupByIPVersionListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupByIPVersionListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPVersion] `query:"httpVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseriesGroupByIPVersionListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupByIPVersionListParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupByIPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupByIPVersionListParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupByIPVersionListParamsAggInterval15m RadarHTTPTimeseriesGroupByIPVersionListParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupByIPVersionListParamsAggInterval1h  RadarHTTPTimeseriesGroupByIPVersionListParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupByIPVersionListParamsAggInterval1d  RadarHTTPTimeseriesGroupByIPVersionListParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupByIPVersionListParamsAggInterval1w  RadarHTTPTimeseriesGroupByIPVersionListParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupByIPVersionListParamsBotClass string

const (
	RadarHTTPTimeseriesGroupByIPVersionListParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupByIPVersionListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupByIPVersionListParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupByIPVersionListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange string

const (
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange1d         RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange2d         RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange7d         RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange14d        RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange28d        RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange12w        RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange24w        RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange52w        RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange1dControl  RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange2dControl  RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange7dControl  RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange14dControl RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange28dControl RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange12wControl RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange24wControl RadarHTTPTimeseriesGroupByIPVersionListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupByIPVersionListParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupByIPVersionListParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupByIPVersionListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupByIPVersionListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupByIPVersionListParamsDeviceTypeOther   RadarHTTPTimeseriesGroupByIPVersionListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupByIPVersionListParamsFormat string

const (
	RadarHTTPTimeseriesGroupByIPVersionListParamsFormatJson RadarHTTPTimeseriesGroupByIPVersionListParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupByIPVersionListParamsFormatCsv  RadarHTTPTimeseriesGroupByIPVersionListParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPProtocolHTTPs RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupByIPVersionListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupByIPVersionListParamsO string

const (
	RadarHTTPTimeseriesGroupByIPVersionListParamsOWindows  RadarHTTPTimeseriesGroupByIPVersionListParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupByIPVersionListParamsOMacosx   RadarHTTPTimeseriesGroupByIPVersionListParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupByIPVersionListParamsOIos      RadarHTTPTimeseriesGroupByIPVersionListParamsO = "IOS"
	RadarHTTPTimeseriesGroupByIPVersionListParamsOAndroid  RadarHTTPTimeseriesGroupByIPVersionListParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupByIPVersionListParamsOChromeos RadarHTTPTimeseriesGroupByIPVersionListParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupByIPVersionListParamsOLinux    RadarHTTPTimeseriesGroupByIPVersionListParamsO = "LINUX"
	RadarHTTPTimeseriesGroupByIPVersionListParamsOSmartTv  RadarHTTPTimeseriesGroupByIPVersionListParamsO = "SMART_TV"
)

type RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupByIPVersionListParamsTlsVersion = "TLSvQUIC"
)
