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

// RadarHTTPTimeseriesGroupByBrowserService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseriesGroupByBrowserService] method instead.
type RadarHTTPTimeseriesGroupByBrowserService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseriesGroupByBrowserService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseriesGroupByBrowserService(opts ...option.RequestOption) (r *RadarHTTPTimeseriesGroupByBrowserService) {
	r = &RadarHTTPTimeseriesGroupByBrowserService{}
	r.Options = opts
	return
}

// Get a time series of the percentage distribution of traffic of the top user
// agents.
func (r *RadarHTTPTimeseriesGroupByBrowserService) List(ctx context.Context, query RadarHTTPTimeseriesGroupByBrowserListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupByBrowserListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/browser"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseriesGroupByBrowserListResponse struct {
	Result  RadarHTTPTimeseriesGroupByBrowserListResponseResult `json:"result,required"`
	Success bool                                                `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupByBrowserListResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByBrowserListResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseriesGroupByBrowserListResponse]
type radarHTTPTimeseriesGroupByBrowserListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByBrowserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByBrowserListResponseResult struct {
	Meta   interface{}                                               `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupByBrowserListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupByBrowserListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByBrowserListResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupByBrowserListResponseResult]
type radarHTTPTimeseriesGroupByBrowserListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByBrowserListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByBrowserListResponseResultSerie0 struct {
	Timestamps []string                                                      `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupByBrowserListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupByBrowserListResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarHTTPTimeseriesGroupByBrowserListResponseResultSerie0]
type radarHTTPTimeseriesGroupByBrowserListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByBrowserListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByBrowserListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupByBrowserListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupByBrowserListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupByBrowserListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupByBrowserListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupByBrowserListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupByBrowserListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupByBrowserListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupByBrowserListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers, verticals, etc) to the top items over
	// the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseriesGroupByBrowserListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupByBrowserListParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupByBrowserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupByBrowserListParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupByBrowserListParamsAggInterval15m RadarHTTPTimeseriesGroupByBrowserListParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupByBrowserListParamsAggInterval1h  RadarHTTPTimeseriesGroupByBrowserListParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupByBrowserListParamsAggInterval1d  RadarHTTPTimeseriesGroupByBrowserListParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupByBrowserListParamsAggInterval1w  RadarHTTPTimeseriesGroupByBrowserListParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupByBrowserListParamsBotClass string

const (
	RadarHTTPTimeseriesGroupByBrowserListParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupByBrowserListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupByBrowserListParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupByBrowserListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupByBrowserListParamsDateRange string

const (
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange1d         RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange2d         RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange7d         RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange14d        RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange28d        RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange12w        RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange24w        RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange52w        RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange1dControl  RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange2dControl  RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange7dControl  RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange14dControl RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange28dControl RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange12wControl RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupByBrowserListParamsDateRange24wControl RadarHTTPTimeseriesGroupByBrowserListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupByBrowserListParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupByBrowserListParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupByBrowserListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupByBrowserListParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupByBrowserListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupByBrowserListParamsDeviceTypeOther   RadarHTTPTimeseriesGroupByBrowserListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupByBrowserListParamsFormat string

const (
	RadarHTTPTimeseriesGroupByBrowserListParamsFormatJson RadarHTTPTimeseriesGroupByBrowserListParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupByBrowserListParamsFormatCsv  RadarHTTPTimeseriesGroupByBrowserListParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupByBrowserListParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupByBrowserListParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupByBrowserListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupByBrowserListParamsHTTPProtocolHTTPs RadarHTTPTimeseriesGroupByBrowserListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupByBrowserListParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupByBrowserListParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupByBrowserListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupByBrowserListParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupByBrowserListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupByBrowserListParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupByBrowserListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupByBrowserListParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupByBrowserListParamsIPVersionIPv4 RadarHTTPTimeseriesGroupByBrowserListParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupByBrowserListParamsIPVersionIPv6 RadarHTTPTimeseriesGroupByBrowserListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupByBrowserListParamsO string

const (
	RadarHTTPTimeseriesGroupByBrowserListParamsOWindows  RadarHTTPTimeseriesGroupByBrowserListParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupByBrowserListParamsOMacosx   RadarHTTPTimeseriesGroupByBrowserListParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupByBrowserListParamsOIos      RadarHTTPTimeseriesGroupByBrowserListParamsO = "IOS"
	RadarHTTPTimeseriesGroupByBrowserListParamsOAndroid  RadarHTTPTimeseriesGroupByBrowserListParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupByBrowserListParamsOChromeos RadarHTTPTimeseriesGroupByBrowserListParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupByBrowserListParamsOLinux    RadarHTTPTimeseriesGroupByBrowserListParamsO = "LINUX"
	RadarHTTPTimeseriesGroupByBrowserListParamsOSmartTv  RadarHTTPTimeseriesGroupByBrowserListParamsO = "SMART_TV"
)

type RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupByBrowserListParamsTlsVersion = "TLSvQUIC"
)
