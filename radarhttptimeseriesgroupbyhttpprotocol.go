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

// RadarHTTPTimeseriesGroupByHTTPProtocolService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarHTTPTimeseriesGroupByHTTPProtocolService] method instead.
type RadarHTTPTimeseriesGroupByHTTPProtocolService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseriesGroupByHTTPProtocolService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarHTTPTimeseriesGroupByHTTPProtocolService(opts ...option.RequestOption) (r *RadarHTTPTimeseriesGroupByHTTPProtocolService) {
	r = &RadarHTTPTimeseriesGroupByHTTPProtocolService{}
	r.Options = opts
	return
}

// Get a time series of the percentage distribution of traffic per HTTP protocol.
func (r *RadarHTTPTimeseriesGroupByHTTPProtocolService) List(ctx context.Context, query RadarHTTPTimeseriesGroupByHTTPProtocolListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupByHTTPProtocolListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/http_protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseriesGroupByHTTPProtocolListResponse struct {
	Result  RadarHTTPTimeseriesGroupByHTTPProtocolListResponseResult `json:"result,required"`
	Success bool                                                     `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupByHTTPProtocolListResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByHTTPProtocolListResponseJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupByHTTPProtocolListResponse]
type radarHTTPTimeseriesGroupByHTTPProtocolListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByHTTPProtocolListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByHTTPProtocolListResponseResult struct {
	Meta   interface{}                                                    `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupByHTTPProtocolListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupByHTTPProtocolListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByHTTPProtocolListResponseResultJSON contains the JSON
// metadata for the struct
// [RadarHTTPTimeseriesGroupByHTTPProtocolListResponseResult]
type radarHTTPTimeseriesGroupByHTTPProtocolListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByHTTPProtocolListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByHTTPProtocolListResponseResultSerie0 struct {
	HTTP       []string                                                           `json:"http,required"`
	HTTPs      []string                                                           `json:"https,required"`
	Timestamps []string                                                           `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupByHTTPProtocolListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupByHTTPProtocolListResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarHTTPTimeseriesGroupByHTTPProtocolListResponseResultSerie0]
type radarHTTPTimeseriesGroupByHTTPProtocolListResponseResultSerie0JSON struct {
	HTTP        apijson.Field
	HTTPs       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByHTTPProtocolListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByHTTPProtocolListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupByHTTPProtocolListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupByHTTPProtocolListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupByHTTPProtocolListParamsFormat] `query:"format"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupByHTTPProtocolListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupByHTTPProtocolListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseriesGroupByHTTPProtocolListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupByHTTPProtocolListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupByHTTPProtocolListParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupByHTTPProtocolListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupByHTTPProtocolListParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsAggInterval15m RadarHTTPTimeseriesGroupByHTTPProtocolListParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsAggInterval1h  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsAggInterval1d  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsAggInterval1w  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupByHTTPProtocolListParamsBotClass string

const (
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupByHTTPProtocolListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupByHTTPProtocolListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange string

const (
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange1d         RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange2d         RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange7d         RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange14d        RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange28d        RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange12w        RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange24w        RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange52w        RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange1dControl  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange2dControl  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange7dControl  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange14dControl RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange28dControl RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange12wControl RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange24wControl RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDeviceTypeOther   RadarHTTPTimeseriesGroupByHTTPProtocolListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupByHTTPProtocolListParamsFormat string

const (
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsFormatJson RadarHTTPTimeseriesGroupByHTTPProtocolListParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsFormatCsv  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupByHTTPProtocolListParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupByHTTPProtocolListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupByHTTPProtocolListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupByHTTPProtocolListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupByHTTPProtocolListParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsIPVersionIPv4 RadarHTTPTimeseriesGroupByHTTPProtocolListParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsIPVersionIPv6 RadarHTTPTimeseriesGroupByHTTPProtocolListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupByHTTPProtocolListParamsO string

const (
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsOWindows  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsOMacosx   RadarHTTPTimeseriesGroupByHTTPProtocolListParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsOIos      RadarHTTPTimeseriesGroupByHTTPProtocolListParamsO = "IOS"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsOAndroid  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsOChromeos RadarHTTPTimeseriesGroupByHTTPProtocolListParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsOLinux    RadarHTTPTimeseriesGroupByHTTPProtocolListParamsO = "LINUX"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsOSmartTv  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsO = "SMART_TV"
)

type RadarHTTPTimeseriesGroupByHTTPProtocolListParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupByHTTPProtocolListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupByHTTPProtocolListParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupByHTTPProtocolListParamsTlsVersion = "TLSvQUIC"
)
