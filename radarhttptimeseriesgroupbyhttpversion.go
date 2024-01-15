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

// RadarHTTPTimeseriesGroupByHTTPVersionService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseriesGroupByHTTPVersionService] method instead.
type RadarHTTPTimeseriesGroupByHTTPVersionService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseriesGroupByHTTPVersionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarHTTPTimeseriesGroupByHTTPVersionService(opts ...option.RequestOption) (r *RadarHTTPTimeseriesGroupByHTTPVersionService) {
	r = &RadarHTTPTimeseriesGroupByHTTPVersionService{}
	r.Options = opts
	return
}

// Get a time series of the percentage distribution of traffic per HTTP protocol
// version.
func (r *RadarHTTPTimeseriesGroupByHTTPVersionService) List(ctx context.Context, query RadarHTTPTimeseriesGroupByHTTPVersionListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupByHTTPVersionListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseriesGroupByHTTPVersionListResponse struct {
	Result  RadarHTTPTimeseriesGroupByHTTPVersionListResponseResult `json:"result,required"`
	Success bool                                                    `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupByHTTPVersionListResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByHTTPVersionListResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupByHTTPVersionListResponse]
type radarHTTPTimeseriesGroupByHTTPVersionListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByHTTPVersionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByHTTPVersionListResponseResult struct {
	Meta   interface{}                                                   `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupByHTTPVersionListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupByHTTPVersionListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByHTTPVersionListResponseResultJSON contains the JSON
// metadata for the struct
// [RadarHTTPTimeseriesGroupByHTTPVersionListResponseResult]
type radarHTTPTimeseriesGroupByHTTPVersionListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByHTTPVersionListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByHTTPVersionListResponseResultSerie0 struct {
	HTTP1X     []string                                                          `json:"HTTP/1.x,required"`
	HTTP2      []string                                                          `json:"HTTP/2,required"`
	HTTP3      []string                                                          `json:"HTTP/3,required"`
	Timestamps []string                                                          `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupByHTTPVersionListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupByHTTPVersionListResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarHTTPTimeseriesGroupByHTTPVersionListResponseResultSerie0]
type radarHTTPTimeseriesGroupByHTTPVersionListResponseResultSerie0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByHTTPVersionListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByHTTPVersionListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupByHTTPVersionListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupByHTTPVersionListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupByHTTPVersionListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupByHTTPVersionListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupByHTTPVersionListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupByHTTPVersionListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseriesGroupByHTTPVersionListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupByHTTPVersionListParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupByHTTPVersionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupByHTTPVersionListParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsAggInterval15m RadarHTTPTimeseriesGroupByHTTPVersionListParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsAggInterval1h  RadarHTTPTimeseriesGroupByHTTPVersionListParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsAggInterval1d  RadarHTTPTimeseriesGroupByHTTPVersionListParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsAggInterval1w  RadarHTTPTimeseriesGroupByHTTPVersionListParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupByHTTPVersionListParamsBotClass string

const (
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupByHTTPVersionListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupByHTTPVersionListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange string

const (
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange1d         RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange2d         RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange7d         RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange14d        RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange28d        RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange12w        RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange24w        RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange52w        RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange1dControl  RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange2dControl  RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange7dControl  RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange14dControl RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange28dControl RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange12wControl RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange24wControl RadarHTTPTimeseriesGroupByHTTPVersionListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupByHTTPVersionListParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupByHTTPVersionListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupByHTTPVersionListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsDeviceTypeOther   RadarHTTPTimeseriesGroupByHTTPVersionListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupByHTTPVersionListParamsFormat string

const (
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsFormatJson RadarHTTPTimeseriesGroupByHTTPVersionListParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsFormatCsv  RadarHTTPTimeseriesGroupByHTTPVersionListParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupByHTTPVersionListParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupByHTTPVersionListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsHTTPProtocolHTTPs RadarHTTPTimeseriesGroupByHTTPVersionListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupByHTTPVersionListParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsIPVersionIPv4 RadarHTTPTimeseriesGroupByHTTPVersionListParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsIPVersionIPv6 RadarHTTPTimeseriesGroupByHTTPVersionListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupByHTTPVersionListParamsO string

const (
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsOWindows  RadarHTTPTimeseriesGroupByHTTPVersionListParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsOMacosx   RadarHTTPTimeseriesGroupByHTTPVersionListParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsOIos      RadarHTTPTimeseriesGroupByHTTPVersionListParamsO = "IOS"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsOAndroid  RadarHTTPTimeseriesGroupByHTTPVersionListParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsOChromeos RadarHTTPTimeseriesGroupByHTTPVersionListParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsOLinux    RadarHTTPTimeseriesGroupByHTTPVersionListParamsO = "LINUX"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsOSmartTv  RadarHTTPTimeseriesGroupByHTTPVersionListParamsO = "SMART_TV"
)

type RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupByHTTPVersionListParamsTlsVersion = "TLSvQUIC"
)
