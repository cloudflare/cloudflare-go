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

// RadarHTTPTimeseriesGroupByOService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseriesGroupByOService] method instead.
type RadarHTTPTimeseriesGroupByOService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseriesGroupByOService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseriesGroupByOService(opts ...option.RequestOption) (r *RadarHTTPTimeseriesGroupByOService) {
	r = &RadarHTTPTimeseriesGroupByOService{}
	r.Options = opts
	return
}

// Get a time series of the percentage distribution of traffic of the top operating
// systems.
func (r *RadarHTTPTimeseriesGroupByOService) List(ctx context.Context, query RadarHTTPTimeseriesGroupByOListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseriesGroupByOListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/os"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseriesGroupByOListResponse struct {
	Result  RadarHTTPTimeseriesGroupByOListResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarHTTPTimeseriesGroupByOListResponseJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByOListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseriesGroupByOListResponse]
type radarHTTPTimeseriesGroupByOListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByOListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByOListResponseResult struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarHTTPTimeseriesGroupByOListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseriesGroupByOListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseriesGroupByOListResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseriesGroupByOListResponseResult]
type radarHTTPTimeseriesGroupByOListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByOListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByOListResponseResultSerie0 struct {
	Timestamps []string                                                `json:"timestamps,required"`
	JSON       radarHTTPTimeseriesGroupByOListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseriesGroupByOListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarHTTPTimeseriesGroupByOListResponseResultSerie0]
type radarHTTPTimeseriesGroupByOListResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseriesGroupByOListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseriesGroupByOListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseriesGroupByOListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseriesGroupByOListParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseriesGroupByOListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseriesGroupByOListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseriesGroupByOListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseriesGroupByOListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseriesGroupByOListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseriesGroupByOListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseriesGroupByOListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseriesGroupByOListParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTimeseriesGroupByOListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseriesGroupByOListParamsAggInterval string

const (
	RadarHTTPTimeseriesGroupByOListParamsAggInterval15m RadarHTTPTimeseriesGroupByOListParamsAggInterval = "15m"
	RadarHTTPTimeseriesGroupByOListParamsAggInterval1h  RadarHTTPTimeseriesGroupByOListParamsAggInterval = "1h"
	RadarHTTPTimeseriesGroupByOListParamsAggInterval1d  RadarHTTPTimeseriesGroupByOListParamsAggInterval = "1d"
	RadarHTTPTimeseriesGroupByOListParamsAggInterval1w  RadarHTTPTimeseriesGroupByOListParamsAggInterval = "1w"
)

type RadarHTTPTimeseriesGroupByOListParamsBotClass string

const (
	RadarHTTPTimeseriesGroupByOListParamsBotClassLikelyAutomated RadarHTTPTimeseriesGroupByOListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseriesGroupByOListParamsBotClassLikelyHuman     RadarHTTPTimeseriesGroupByOListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseriesGroupByOListParamsDateRange string

const (
	RadarHTTPTimeseriesGroupByOListParamsDateRange1d         RadarHTTPTimeseriesGroupByOListParamsDateRange = "1d"
	RadarHTTPTimeseriesGroupByOListParamsDateRange2d         RadarHTTPTimeseriesGroupByOListParamsDateRange = "2d"
	RadarHTTPTimeseriesGroupByOListParamsDateRange7d         RadarHTTPTimeseriesGroupByOListParamsDateRange = "7d"
	RadarHTTPTimeseriesGroupByOListParamsDateRange14d        RadarHTTPTimeseriesGroupByOListParamsDateRange = "14d"
	RadarHTTPTimeseriesGroupByOListParamsDateRange28d        RadarHTTPTimeseriesGroupByOListParamsDateRange = "28d"
	RadarHTTPTimeseriesGroupByOListParamsDateRange12w        RadarHTTPTimeseriesGroupByOListParamsDateRange = "12w"
	RadarHTTPTimeseriesGroupByOListParamsDateRange24w        RadarHTTPTimeseriesGroupByOListParamsDateRange = "24w"
	RadarHTTPTimeseriesGroupByOListParamsDateRange52w        RadarHTTPTimeseriesGroupByOListParamsDateRange = "52w"
	RadarHTTPTimeseriesGroupByOListParamsDateRange1dControl  RadarHTTPTimeseriesGroupByOListParamsDateRange = "1dControl"
	RadarHTTPTimeseriesGroupByOListParamsDateRange2dControl  RadarHTTPTimeseriesGroupByOListParamsDateRange = "2dControl"
	RadarHTTPTimeseriesGroupByOListParamsDateRange7dControl  RadarHTTPTimeseriesGroupByOListParamsDateRange = "7dControl"
	RadarHTTPTimeseriesGroupByOListParamsDateRange14dControl RadarHTTPTimeseriesGroupByOListParamsDateRange = "14dControl"
	RadarHTTPTimeseriesGroupByOListParamsDateRange28dControl RadarHTTPTimeseriesGroupByOListParamsDateRange = "28dControl"
	RadarHTTPTimeseriesGroupByOListParamsDateRange12wControl RadarHTTPTimeseriesGroupByOListParamsDateRange = "12wControl"
	RadarHTTPTimeseriesGroupByOListParamsDateRange24wControl RadarHTTPTimeseriesGroupByOListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseriesGroupByOListParamsDeviceType string

const (
	RadarHTTPTimeseriesGroupByOListParamsDeviceTypeDesktop RadarHTTPTimeseriesGroupByOListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseriesGroupByOListParamsDeviceTypeMobile  RadarHTTPTimeseriesGroupByOListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseriesGroupByOListParamsDeviceTypeOther   RadarHTTPTimeseriesGroupByOListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseriesGroupByOListParamsFormat string

const (
	RadarHTTPTimeseriesGroupByOListParamsFormatJson RadarHTTPTimeseriesGroupByOListParamsFormat = "JSON"
	RadarHTTPTimeseriesGroupByOListParamsFormatCsv  RadarHTTPTimeseriesGroupByOListParamsFormat = "CSV"
)

type RadarHTTPTimeseriesGroupByOListParamsHTTPProtocol string

const (
	RadarHTTPTimeseriesGroupByOListParamsHTTPProtocolHTTP  RadarHTTPTimeseriesGroupByOListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseriesGroupByOListParamsHTTPProtocolHTTPs RadarHTTPTimeseriesGroupByOListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseriesGroupByOListParamsHTTPVersion string

const (
	RadarHTTPTimeseriesGroupByOListParamsHTTPVersionHttPv1 RadarHTTPTimeseriesGroupByOListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseriesGroupByOListParamsHTTPVersionHttPv2 RadarHTTPTimeseriesGroupByOListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseriesGroupByOListParamsHTTPVersionHttPv3 RadarHTTPTimeseriesGroupByOListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseriesGroupByOListParamsIPVersion string

const (
	RadarHTTPTimeseriesGroupByOListParamsIPVersionIPv4 RadarHTTPTimeseriesGroupByOListParamsIPVersion = "IPv4"
	RadarHTTPTimeseriesGroupByOListParamsIPVersionIPv6 RadarHTTPTimeseriesGroupByOListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseriesGroupByOListParamsTlsVersion string

const (
	RadarHTTPTimeseriesGroupByOListParamsTlsVersionTlSv1_0  RadarHTTPTimeseriesGroupByOListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseriesGroupByOListParamsTlsVersionTlSv1_1  RadarHTTPTimeseriesGroupByOListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseriesGroupByOListParamsTlsVersionTlSv1_2  RadarHTTPTimeseriesGroupByOListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseriesGroupByOListParamsTlsVersionTlSv1_3  RadarHTTPTimeseriesGroupByOListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseriesGroupByOListParamsTlsVersionTlSvQuic RadarHTTPTimeseriesGroupByOListParamsTlsVersion = "TLSvQUIC"
)
