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

// RadarHTTPTimeseryOService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRadarHTTPTimeseryOService] method
// instead.
type RadarHTTPTimeseryOService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseryOService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseryOService(opts ...option.RequestOption) (r *RadarHTTPTimeseryOService) {
	r = &RadarHTTPTimeseryOService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic of the top operating systems in the selected
// time range, over time.
func (r *RadarHTTPTimeseryOService) List(ctx context.Context, query RadarHTTPTimeseryOListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseryOListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries/os"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseryOListResponse struct {
	Result  RadarHTTPTimeseryOListResponseResult `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarHTTPTimeseryOListResponseJSON   `json:"-"`
}

// radarHTTPTimeseryOListResponseJSON contains the JSON metadata for the struct
// [RadarHTTPTimeseryOListResponse]
type radarHTTPTimeseryOListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryOListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryOListResponseResult struct {
	Meta   interface{}                                `json:"meta,required"`
	Serie0 RadarHTTPTimeseryOListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseryOListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseryOListResponseResultJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseryOListResponseResult]
type radarHTTPTimeseryOListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryOListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryOListResponseResultSerie0 struct {
	OsName     []string                                       `json:"<OS name>,required"`
	Timestamps []string                                       `json:"timestamps,required"`
	JSON       radarHTTPTimeseryOListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseryOListResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarHTTPTimeseryOListResponseResultSerie0]
type radarHTTPTimeseryOListResponseResultSerie0JSON struct {
	OsName      apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryOListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryOListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseryOListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseryOListParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseryOListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseryOListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseryOListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseryOListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseryOListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseryOListParamsIPVersion] `query:"ipVersion"`
	// Limit the number of objects (eg browsers) to the top items over the time range.
	LimitPerGroup param.Field[int64] `query:"limitPerGroup"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseryOListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseryOListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTimeseryOListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseryOListParamsAggInterval string

const (
	RadarHTTPTimeseryOListParamsAggInterval15m RadarHTTPTimeseryOListParamsAggInterval = "15m"
	RadarHTTPTimeseryOListParamsAggInterval1h  RadarHTTPTimeseryOListParamsAggInterval = "1h"
	RadarHTTPTimeseryOListParamsAggInterval1d  RadarHTTPTimeseryOListParamsAggInterval = "1d"
	RadarHTTPTimeseryOListParamsAggInterval1w  RadarHTTPTimeseryOListParamsAggInterval = "1w"
)

type RadarHTTPTimeseryOListParamsBotClass string

const (
	RadarHTTPTimeseryOListParamsBotClassLikelyAutomated RadarHTTPTimeseryOListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseryOListParamsBotClassLikelyHuman     RadarHTTPTimeseryOListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseryOListParamsDateRange string

const (
	RadarHTTPTimeseryOListParamsDateRange1d         RadarHTTPTimeseryOListParamsDateRange = "1d"
	RadarHTTPTimeseryOListParamsDateRange7d         RadarHTTPTimeseryOListParamsDateRange = "7d"
	RadarHTTPTimeseryOListParamsDateRange14d        RadarHTTPTimeseryOListParamsDateRange = "14d"
	RadarHTTPTimeseryOListParamsDateRange28d        RadarHTTPTimeseryOListParamsDateRange = "28d"
	RadarHTTPTimeseryOListParamsDateRange12w        RadarHTTPTimeseryOListParamsDateRange = "12w"
	RadarHTTPTimeseryOListParamsDateRange24w        RadarHTTPTimeseryOListParamsDateRange = "24w"
	RadarHTTPTimeseryOListParamsDateRange52w        RadarHTTPTimeseryOListParamsDateRange = "52w"
	RadarHTTPTimeseryOListParamsDateRange1dControl  RadarHTTPTimeseryOListParamsDateRange = "1dControl"
	RadarHTTPTimeseryOListParamsDateRange7dControl  RadarHTTPTimeseryOListParamsDateRange = "7dControl"
	RadarHTTPTimeseryOListParamsDateRange14dControl RadarHTTPTimeseryOListParamsDateRange = "14dControl"
	RadarHTTPTimeseryOListParamsDateRange28dControl RadarHTTPTimeseryOListParamsDateRange = "28dControl"
	RadarHTTPTimeseryOListParamsDateRange12wControl RadarHTTPTimeseryOListParamsDateRange = "12wControl"
	RadarHTTPTimeseryOListParamsDateRange24wControl RadarHTTPTimeseryOListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseryOListParamsDeviceType string

const (
	RadarHTTPTimeseryOListParamsDeviceTypeDesktop RadarHTTPTimeseryOListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseryOListParamsDeviceTypeMobile  RadarHTTPTimeseryOListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseryOListParamsDeviceTypeOther   RadarHTTPTimeseryOListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseryOListParamsFormat string

const (
	RadarHTTPTimeseryOListParamsFormatJson RadarHTTPTimeseryOListParamsFormat = "JSON"
	RadarHTTPTimeseryOListParamsFormatCsv  RadarHTTPTimeseryOListParamsFormat = "CSV"
)

type RadarHTTPTimeseryOListParamsHTTPProtocol string

const (
	RadarHTTPTimeseryOListParamsHTTPProtocolHTTP  RadarHTTPTimeseryOListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseryOListParamsHTTPProtocolHTTPs RadarHTTPTimeseryOListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseryOListParamsHTTPVersion string

const (
	RadarHTTPTimeseryOListParamsHTTPVersionHttPv1 RadarHTTPTimeseryOListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseryOListParamsHTTPVersionHttPv2 RadarHTTPTimeseryOListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseryOListParamsHTTPVersionHttPv3 RadarHTTPTimeseryOListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseryOListParamsIPVersion string

const (
	RadarHTTPTimeseryOListParamsIPVersionIPv4 RadarHTTPTimeseryOListParamsIPVersion = "IPv4"
	RadarHTTPTimeseryOListParamsIPVersionIPv6 RadarHTTPTimeseryOListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseryOListParamsTlsVersion string

const (
	RadarHTTPTimeseryOListParamsTlsVersionTlSv1_0  RadarHTTPTimeseryOListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseryOListParamsTlsVersionTlSv1_1  RadarHTTPTimeseryOListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseryOListParamsTlsVersionTlSv1_2  RadarHTTPTimeseryOListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseryOListParamsTlsVersionTlSv1_3  RadarHTTPTimeseryOListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseryOListParamsTlsVersionTlSvQuic RadarHTTPTimeseryOListParamsTlsVersion = "TLSvQUIC"
)
