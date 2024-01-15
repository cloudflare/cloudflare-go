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

// RadarHTTPTimeseriesGroupByDeviceTypeService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseriesGroupByDeviceTypeService] method instead.
type RadarHTTPTimeseriesGroupByDeviceTypeService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseriesGroupByDeviceTypeService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarHTTPTimeseriesGroupByDeviceTypeService(opts ...option.RequestOption) (r *RadarHTTPTimeseriesGroupByDeviceTypeService) {
	r = &RadarHTTPTimeseriesGroupByDeviceTypeService{}
	r.Options = opts
	return
}

// Get a time series of the percentage distribution of traffic per device type.
func (r *RadarHTTPTimeseriesGroupByDeviceTypeService) List(ctx context.Context, query RadarHTTPTimeseriesGroupByDeviceTypeListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupByDeviceTypeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/device_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseriesGroupByDeviceTypeListResponse struct {
	Result  RadarHTTPTimeseriesGroupByDeviceTypeListResponseResult `json:"result,required"`
	Success bool                                                   `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupByDeviceTypeListResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByDeviceTypeListResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupByDeviceTypeListResponse]
type radarHTTPTimeseriesGroupByDeviceTypeListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByDeviceTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByDeviceTypeListResponseResult struct {
	Meta   interface{}                                                  `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupByDeviceTypeListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupByDeviceTypeListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByDeviceTypeListResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupByDeviceTypeListResponseResult]
type radarHTTPTimeseriesGroupByDeviceTypeListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByDeviceTypeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByDeviceTypeListResponseResultSerie0 struct {
	Desktop    []string                                                         `json:"desktop,required"`
	Mobile     []string                                                         `json:"mobile,required"`
	Other      []string                                                         `json:"other,required"`
	Timestamps []string                                                         `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupByDeviceTypeListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupByDeviceTypeListResponseResultSerie0JSON contains the
// JSON metadata for the struct
// [RadarHTTPTimeseriesGroupByDeviceTypeListResponseResultSerie0]
type radarHTTPTimeseriesGroupByDeviceTypeListResponseResultSerie0JSON struct {
	Desktop     apijson.Field
	Mobile      apijson.Field
	Other       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByDeviceTypeListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByDeviceTypeListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupByDeviceTypeListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupByDeviceTypeListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupByDeviceTypeListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupByDeviceTypeListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseriesGroupByDeviceTypeListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupByDeviceTypeListParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupByDeviceTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupByDeviceTypeListParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsAggInterval15m RadarHTTPTimeseriesGroupByDeviceTypeListParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsAggInterval1h  RadarHTTPTimeseriesGroupByDeviceTypeListParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsAggInterval1d  RadarHTTPTimeseriesGroupByDeviceTypeListParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsAggInterval1w  RadarHTTPTimeseriesGroupByDeviceTypeListParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupByDeviceTypeListParamsBotClass string

const (
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupByDeviceTypeListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupByDeviceTypeListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange string

const (
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange1d         RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange2d         RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange7d         RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange14d        RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange28d        RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange12w        RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange24w        RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange52w        RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange1dControl  RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange2dControl  RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange7dControl  RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange14dControl RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange28dControl RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange12wControl RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange24wControl RadarHTTPTimeseriesGroupByDeviceTypeListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupByDeviceTypeListParamsFormat string

const (
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsFormatJson RadarHTTPTimeseriesGroupByDeviceTypeListParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsFormatCsv  RadarHTTPTimeseriesGroupByDeviceTypeListParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPProtocolHTTPs RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupByDeviceTypeListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupByDeviceTypeListParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsIPVersionIPv4 RadarHTTPTimeseriesGroupByDeviceTypeListParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsIPVersionIPv6 RadarHTTPTimeseriesGroupByDeviceTypeListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupByDeviceTypeListParamsO string

const (
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsOWindows  RadarHTTPTimeseriesGroupByDeviceTypeListParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsOMacosx   RadarHTTPTimeseriesGroupByDeviceTypeListParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsOIos      RadarHTTPTimeseriesGroupByDeviceTypeListParamsO = "IOS"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsOAndroid  RadarHTTPTimeseriesGroupByDeviceTypeListParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsOChromeos RadarHTTPTimeseriesGroupByDeviceTypeListParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsOLinux    RadarHTTPTimeseriesGroupByDeviceTypeListParamsO = "LINUX"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsOSmartTv  RadarHTTPTimeseriesGroupByDeviceTypeListParamsO = "SMART_TV"
)

type RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupByDeviceTypeListParamsTlsVersion = "TLSvQUIC"
)
