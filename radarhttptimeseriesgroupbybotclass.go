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

// RadarHTTPTimeseriesGroupByBotClassService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseriesGroupByBotClassService] method instead.
type RadarHTTPTimeseriesGroupByBotClassService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseriesGroupByBotClassService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewRadarHTTPTimeseriesGroupByBotClassService(opts ...option.RequestOption) (r *RadarHTTPTimeseriesGroupByBotClassService) {
	r = &RadarHTTPTimeseriesGroupByBotClassService{}
	r.Options = opts
	return
}

// Get a time series of the percentage distribution of traffic classified as
// automated or human. Visit
// https://developers.cloudflare.com/radar/concepts/bot-classes/ for more
// information.
func (r *RadarHTTPTimeseriesGroupByBotClassService) List(ctx context.Context, query RadarHTTPTimeseriesGroupByBotClassListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupByBotClassListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/bot_class"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseriesGroupByBotClassListResponse struct {
	Result  RadarHTTPTimeseriesGroupByBotClassListResponseResult `json:"result,required"`
	Success bool                                                 `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupByBotClassListResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByBotClassListResponseJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseriesGroupByBotClassListResponse]
type radarHTTPTimeseriesGroupByBotClassListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByBotClassListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByBotClassListResponseResult struct {
	Meta   interface{}                                                `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupByBotClassListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupByBotClassListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByBotClassListResponseResultJSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupByBotClassListResponseResult]
type radarHTTPTimeseriesGroupByBotClassListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByBotClassListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByBotClassListResponseResultSerie0 struct {
	Bot        []string                                                       `json:"bot,required"`
	Human      []string                                                       `json:"human,required"`
	Timestamps []string                                                       `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupByBotClassListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupByBotClassListResponseResultSerie0JSON contains the JSON
// metadata for the struct
// [RadarHTTPTimeseriesGroupByBotClassListResponseResultSerie0]
type radarHTTPTimeseriesGroupByBotClassListResponseResultSerie0JSON struct {
	Bot         apijson.Field
	Human       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByBotClassListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByBotClassListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupByBotClassListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupByBotClassListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupByBotClassListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupByBotClassListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupByBotClassListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupByBotClassListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupByBotClassListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseriesGroupByBotClassListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupByBotClassListParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseriesGroupByBotClassListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupByBotClassListParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupByBotClassListParamsAggInterval15m RadarHTTPTimeseriesGroupByBotClassListParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupByBotClassListParamsAggInterval1h  RadarHTTPTimeseriesGroupByBotClassListParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupByBotClassListParamsAggInterval1d  RadarHTTPTimeseriesGroupByBotClassListParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupByBotClassListParamsAggInterval1w  RadarHTTPTimeseriesGroupByBotClassListParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupByBotClassListParamsDateRange string

const (
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange1d         RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange2d         RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange7d         RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange14d        RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange28d        RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange12w        RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange24w        RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange52w        RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange1dControl  RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange2dControl  RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange7dControl  RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange14dControl RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange28dControl RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange12wControl RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupByBotClassListParamsDateRange24wControl RadarHTTPTimeseriesGroupByBotClassListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupByBotClassListParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupByBotClassListParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupByBotClassListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupByBotClassListParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupByBotClassListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupByBotClassListParamsDeviceTypeOther   RadarHTTPTimeseriesGroupByBotClassListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupByBotClassListParamsFormat string

const (
	RadarHTTPTimeseriesGroupByBotClassListParamsFormatJson RadarHTTPTimeseriesGroupByBotClassListParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupByBotClassListParamsFormatCsv  RadarHTTPTimeseriesGroupByBotClassListParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupByBotClassListParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupByBotClassListParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupByBotClassListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupByBotClassListParamsHTTPProtocolHTTPs RadarHTTPTimeseriesGroupByBotClassListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupByBotClassListParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupByBotClassListParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupByBotClassListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupByBotClassListParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupByBotClassListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupByBotClassListParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupByBotClassListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupByBotClassListParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupByBotClassListParamsIPVersionIPv4 RadarHTTPTimeseriesGroupByBotClassListParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupByBotClassListParamsIPVersionIPv6 RadarHTTPTimeseriesGroupByBotClassListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupByBotClassListParamsO string

const (
	RadarHTTPTimeseriesGroupByBotClassListParamsOWindows  RadarHTTPTimeseriesGroupByBotClassListParamsO = "WINDOWS"
	RadarHTTPTimeseriesGroupByBotClassListParamsOMacosx   RadarHTTPTimeseriesGroupByBotClassListParamsO = "MACOSX"
	RadarHTTPTimeseriesGroupByBotClassListParamsOIos      RadarHTTPTimeseriesGroupByBotClassListParamsO = "IOS"
	RadarHTTPTimeseriesGroupByBotClassListParamsOAndroid  RadarHTTPTimeseriesGroupByBotClassListParamsO = "ANDROID"
	RadarHTTPTimeseriesGroupByBotClassListParamsOChromeos RadarHTTPTimeseriesGroupByBotClassListParamsO = "CHROMEOS"
	RadarHTTPTimeseriesGroupByBotClassListParamsOLinux    RadarHTTPTimeseriesGroupByBotClassListParamsO = "LINUX"
	RadarHTTPTimeseriesGroupByBotClassListParamsOSmartTv  RadarHTTPTimeseriesGroupByBotClassListParamsO = "SMART_TV"
)

type RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupByBotClassListParamsTlsVersion = "TLSvQUIC"
)
