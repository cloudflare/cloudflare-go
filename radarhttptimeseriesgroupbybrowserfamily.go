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

// RadarHTTPTimeseriesGroupByBrowserFamilyService contains methods and other
// services that help with interacting with the cloudflare API. Note, unlike
// clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRadarHTTPTimeseriesGroupByBrowserFamilyService] method instead.
type RadarHTTPTimeseriesGroupByBrowserFamilyService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseriesGroupByBrowserFamilyService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarHTTPTimeseriesGroupByBrowserFamilyService(opts ...option.RequestOption) (r *RadarHTTPTimeseriesGroupByBrowserFamilyService) {
	r = &RadarHTTPTimeseriesGroupByBrowserFamilyService{}
	r.Options = opts
	return
}

// Get a time series of the percentage distribution of traffic of the top user
// agents aggregated in families.
func (r *RadarHTTPTimeseriesGroupByBrowserFamilyService) List(ctx context.Context, query RadarHTTPTimeseriesGroupByBrowserFamilyListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupByBrowserFamilyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/browser_family"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseriesGroupByBrowserFamilyListResponse struct {
	Result  RadarHTTPTimeseriesGroupByBrowserFamilyListResponseResult `json:"result,required"`
	Success bool                                                      `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupByBrowserFamilyListResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByBrowserFamilyListResponseJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupByBrowserFamilyListResponse]
type radarHTTPTimeseriesGroupByBrowserFamilyListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByBrowserFamilyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByBrowserFamilyListResponseResult struct {
	Meta   interface{}                                                     `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupByBrowserFamilyListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupByBrowserFamilyListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByBrowserFamilyListResponseResultJSON contains the JSON
// metadata for the struct
// [RadarHTTPTimeseriesGroupByBrowserFamilyListResponseResult]
type radarHTTPTimeseriesGroupByBrowserFamilyListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByBrowserFamilyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByBrowserFamilyListResponseResultSerie0 struct {
	Timestamps []string                                                            `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupByBrowserFamilyListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupByBrowserFamilyListResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarHTTPTimeseriesGroupByBrowserFamilyListResponseResultSerie0]
type radarHTTPTimeseriesGroupByBrowserFamilyListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByBrowserFamilyListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByBrowserFamilyListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupByBrowserFamilyListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupByBrowserFamilyListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupByBrowserFamilyListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupByBrowserFamilyListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseriesGroupByBrowserFamilyListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupByBrowserFamilyListParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupByBrowserFamilyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupByBrowserFamilyListParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsAggInterval15m RadarHTTPTimeseriesGroupByBrowserFamilyListParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsAggInterval1h  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsAggInterval1d  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsAggInterval1w  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupByBrowserFamilyListParamsBotClass string

const (
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupByBrowserFamilyListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupByBrowserFamilyListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange string

const (
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange1d         RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange2d         RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange7d         RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange14d        RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange28d        RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange12w        RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange24w        RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange52w        RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange1dControl  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange2dControl  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange7dControl  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange14dControl RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange28dControl RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange12wControl RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange24wControl RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDeviceTypeOther   RadarHTTPTimeseriesGroupByBrowserFamilyListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupByBrowserFamilyListParamsFormat string

const (
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsFormatJson RadarHTTPTimeseriesGroupByBrowserFamilyListParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsFormatCsv  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPProtocolHTTPs RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupByBrowserFamilyListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupByBrowserFamilyListParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsIPVersionIPv4 RadarHTTPTimeseriesGroupByBrowserFamilyListParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsIPVersionIPv6 RadarHTTPTimeseriesGroupByBrowserFamilyListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupByBrowserFamilyListParamsO string

const (
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsOWindows  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsOMacosx   RadarHTTPTimeseriesGroupByBrowserFamilyListParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsOIos      RadarHTTPTimeseriesGroupByBrowserFamilyListParamsO = "IOS"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsOAndroid  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsOChromeos RadarHTTPTimeseriesGroupByBrowserFamilyListParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsOLinux    RadarHTTPTimeseriesGroupByBrowserFamilyListParamsO = "LINUX"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsOSmartTv  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsO = "SMART_TV"
)

type RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupByBrowserFamilyListParamsTlsVersion = "TLSvQUIC"
)
