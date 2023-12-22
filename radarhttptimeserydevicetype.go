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

// RadarHTTPTimeseryDeviceTypeService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseryDeviceTypeService] method instead.
type RadarHTTPTimeseryDeviceTypeService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseryDeviceTypeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseryDeviceTypeService(opts ...option.RequestOption) (r *RadarHTTPTimeseryDeviceTypeService) {
	r = &RadarHTTPTimeseryDeviceTypeService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic per device type over time.
func (r *RadarHTTPTimeseryDeviceTypeService) List(ctx context.Context, query RadarHTTPTimeseryDeviceTypeListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseryDeviceTypeListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries/device_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseryDeviceTypeListResponse struct {
	Result  RadarHTTPTimeseryDeviceTypeListResponseResult `json:"result,required"`
	Success bool                                          `json:"success,required"`
	JSON    radarHTTPTimeseryDeviceTypeListResponseJSON   `json:"-"`
}

// radarHTTPTimeseryDeviceTypeListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseryDeviceTypeListResponse]
type radarHTTPTimeseryDeviceTypeListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryDeviceTypeListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryDeviceTypeListResponseResult struct {
	Meta   interface{}                                         `json:"meta,required"`
	Serie0 RadarHTTPTimeseryDeviceTypeListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseryDeviceTypeListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseryDeviceTypeListResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseryDeviceTypeListResponseResult]
type radarHTTPTimeseryDeviceTypeListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryDeviceTypeListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryDeviceTypeListResponseResultSerie0 struct {
	Desktop    []string                                                `json:"desktop,required"`
	Mobile     []string                                                `json:"mobile,required"`
	Other      []string                                                `json:"other,required"`
	Timestamps []string                                                `json:"timestamps,required"`
	JSON       radarHTTPTimeseryDeviceTypeListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseryDeviceTypeListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarHTTPTimeseryDeviceTypeListResponseResultSerie0]
type radarHTTPTimeseryDeviceTypeListResponseResultSerie0JSON struct {
	Desktop     apijson.Field
	Mobile      apijson.Field
	Other       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryDeviceTypeListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryDeviceTypeListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseryDeviceTypeListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPTimeseryDeviceTypeListParamsBotClass] `query:"botClass"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseryDeviceTypeListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseryDeviceTypeListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseryDeviceTypeListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseryDeviceTypeListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseryDeviceTypeListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseryDeviceTypeListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseryDeviceTypeListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseryDeviceTypeListParams]'s query parameters
// as `url.Values`.
func (r RadarHTTPTimeseryDeviceTypeListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseryDeviceTypeListParamsAggInterval string

const (
	RadarHTTPTimeseryDeviceTypeListParamsAggInterval15m RadarHTTPTimeseryDeviceTypeListParamsAggInterval = "15m"
	RadarHTTPTimeseryDeviceTypeListParamsAggInterval1h  RadarHTTPTimeseryDeviceTypeListParamsAggInterval = "1h"
	RadarHTTPTimeseryDeviceTypeListParamsAggInterval1d  RadarHTTPTimeseryDeviceTypeListParamsAggInterval = "1d"
	RadarHTTPTimeseryDeviceTypeListParamsAggInterval1w  RadarHTTPTimeseryDeviceTypeListParamsAggInterval = "1w"
)

type RadarHTTPTimeseryDeviceTypeListParamsBotClass string

const (
	RadarHTTPTimeseryDeviceTypeListParamsBotClassLikelyAutomated RadarHTTPTimeseryDeviceTypeListParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPTimeseryDeviceTypeListParamsBotClassLikelyHuman     RadarHTTPTimeseryDeviceTypeListParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPTimeseryDeviceTypeListParamsDateRange string

const (
	RadarHTTPTimeseryDeviceTypeListParamsDateRange1d         RadarHTTPTimeseryDeviceTypeListParamsDateRange = "1d"
	RadarHTTPTimeseryDeviceTypeListParamsDateRange7d         RadarHTTPTimeseryDeviceTypeListParamsDateRange = "7d"
	RadarHTTPTimeseryDeviceTypeListParamsDateRange14d        RadarHTTPTimeseryDeviceTypeListParamsDateRange = "14d"
	RadarHTTPTimeseryDeviceTypeListParamsDateRange28d        RadarHTTPTimeseryDeviceTypeListParamsDateRange = "28d"
	RadarHTTPTimeseryDeviceTypeListParamsDateRange12w        RadarHTTPTimeseryDeviceTypeListParamsDateRange = "12w"
	RadarHTTPTimeseryDeviceTypeListParamsDateRange24w        RadarHTTPTimeseryDeviceTypeListParamsDateRange = "24w"
	RadarHTTPTimeseryDeviceTypeListParamsDateRange52w        RadarHTTPTimeseryDeviceTypeListParamsDateRange = "52w"
	RadarHTTPTimeseryDeviceTypeListParamsDateRange1dControl  RadarHTTPTimeseryDeviceTypeListParamsDateRange = "1dControl"
	RadarHTTPTimeseryDeviceTypeListParamsDateRange7dControl  RadarHTTPTimeseryDeviceTypeListParamsDateRange = "7dControl"
	RadarHTTPTimeseryDeviceTypeListParamsDateRange14dControl RadarHTTPTimeseryDeviceTypeListParamsDateRange = "14dControl"
	RadarHTTPTimeseryDeviceTypeListParamsDateRange28dControl RadarHTTPTimeseryDeviceTypeListParamsDateRange = "28dControl"
	RadarHTTPTimeseryDeviceTypeListParamsDateRange12wControl RadarHTTPTimeseryDeviceTypeListParamsDateRange = "12wControl"
	RadarHTTPTimeseryDeviceTypeListParamsDateRange24wControl RadarHTTPTimeseryDeviceTypeListParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarHTTPTimeseryDeviceTypeListParamsFormat string

const (
	RadarHTTPTimeseryDeviceTypeListParamsFormatJson RadarHTTPTimeseryDeviceTypeListParamsFormat = "JSON"
	RadarHTTPTimeseryDeviceTypeListParamsFormatCsv  RadarHTTPTimeseryDeviceTypeListParamsFormat = "CSV"
)

type RadarHTTPTimeseryDeviceTypeListParamsHTTPProtocol string

const (
	RadarHTTPTimeseryDeviceTypeListParamsHTTPProtocolHTTP  RadarHTTPTimeseryDeviceTypeListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseryDeviceTypeListParamsHTTPProtocolHTTPs RadarHTTPTimeseryDeviceTypeListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseryDeviceTypeListParamsHTTPVersion string

const (
	RadarHTTPTimeseryDeviceTypeListParamsHTTPVersionHttPv1 RadarHTTPTimeseryDeviceTypeListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseryDeviceTypeListParamsHTTPVersionHttPv2 RadarHTTPTimeseryDeviceTypeListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseryDeviceTypeListParamsHTTPVersionHttPv3 RadarHTTPTimeseryDeviceTypeListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseryDeviceTypeListParamsIPVersion string

const (
	RadarHTTPTimeseryDeviceTypeListParamsIPVersionIPv4 RadarHTTPTimeseryDeviceTypeListParamsIPVersion = "IPv4"
	RadarHTTPTimeseryDeviceTypeListParamsIPVersionIPv6 RadarHTTPTimeseryDeviceTypeListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseryDeviceTypeListParamsO string

const (
	RadarHTTPTimeseryDeviceTypeListParamsOWindows  RadarHTTPTimeseryDeviceTypeListParamsO = "WINDOWS"
	RadarHTTPTimeseryDeviceTypeListParamsOMacosx   RadarHTTPTimeseryDeviceTypeListParamsO = "MACOSX"
	RadarHTTPTimeseryDeviceTypeListParamsOAndroid  RadarHTTPTimeseryDeviceTypeListParamsO = "ANDROID"
	RadarHTTPTimeseryDeviceTypeListParamsOChromeos RadarHTTPTimeseryDeviceTypeListParamsO = "CHROMEOS"
	RadarHTTPTimeseryDeviceTypeListParamsOLinux    RadarHTTPTimeseryDeviceTypeListParamsO = "LINUX"
	RadarHTTPTimeseryDeviceTypeListParamsOSmartTv  RadarHTTPTimeseryDeviceTypeListParamsO = "SMART_TV"
)

type RadarHTTPTimeseryDeviceTypeListParamsTlsVersion string

const (
	RadarHTTPTimeseryDeviceTypeListParamsTlsVersionTlSv1_0  RadarHTTPTimeseryDeviceTypeListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseryDeviceTypeListParamsTlsVersionTlSv1_1  RadarHTTPTimeseryDeviceTypeListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseryDeviceTypeListParamsTlsVersionTlSv1_2  RadarHTTPTimeseryDeviceTypeListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseryDeviceTypeListParamsTlsVersionTlSv1_3  RadarHTTPTimeseryDeviceTypeListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseryDeviceTypeListParamsTlsVersionTlSvQuic RadarHTTPTimeseryDeviceTypeListParamsTlsVersion = "TLSvQUIC"
)
