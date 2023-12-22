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

// RadarHTTPTimeseryBotClassService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseryBotClassService] method instead.
type RadarHTTPTimeseryBotClassService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseryBotClassService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseryBotClassService(opts ...option.RequestOption) (r *RadarHTTPTimeseryBotClassService) {
	r = &RadarHTTPTimeseryBotClassService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic classified as automated or human over time.
func (r *RadarHTTPTimeseryBotClassService) List(ctx context.Context, query RadarHTTPTimeseryBotClassListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseryBotClassListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries/bot_class"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseryBotClassListResponse struct {
	Result  RadarHTTPTimeseryBotClassListResponseResult `json:"result,required"`
	Success bool                                        `json:"success,required"`
	JSON    radarHTTPTimeseryBotClassListResponseJSON   `json:"-"`
}

// radarHTTPTimeseryBotClassListResponseJSON contains the JSON metadata for the
// struct [RadarHTTPTimeseryBotClassListResponse]
type radarHTTPTimeseryBotClassListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryBotClassListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryBotClassListResponseResult struct {
	Meta   interface{}                                       `json:"meta,required"`
	Serie0 RadarHTTPTimeseryBotClassListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseryBotClassListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseryBotClassListResponseResultJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseryBotClassListResponseResult]
type radarHTTPTimeseryBotClassListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryBotClassListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryBotClassListResponseResultSerie0 struct {
	Bot        []string                                              `json:"bot,required"`
	Human      []string                                              `json:"human,required"`
	Timestamps []string                                              `json:"timestamps,required"`
	JSON       radarHTTPTimeseryBotClassListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseryBotClassListResponseResultSerie0JSON contains the JSON metadata
// for the struct [RadarHTTPTimeseryBotClassListResponseResultSerie0]
type radarHTTPTimeseryBotClassListResponseResultSerie0JSON struct {
	Bot         apijson.Field
	Human       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryBotClassListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryBotClassListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseryBotClassListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseryBotClassListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseryBotClassListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseryBotClassListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseryBotClassListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseryBotClassListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseryBotClassListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseryBotClassListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseryBotClassListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseryBotClassListParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPTimeseryBotClassListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseryBotClassListParamsAggInterval string

const (
	RadarHTTPTimeseryBotClassListParamsAggInterval15m RadarHTTPTimeseryBotClassListParamsAggInterval = "15m"
	RadarHTTPTimeseryBotClassListParamsAggInterval1h  RadarHTTPTimeseryBotClassListParamsAggInterval = "1h"
	RadarHTTPTimeseryBotClassListParamsAggInterval1d  RadarHTTPTimeseryBotClassListParamsAggInterval = "1d"
	RadarHTTPTimeseryBotClassListParamsAggInterval1w  RadarHTTPTimeseryBotClassListParamsAggInterval = "1w"
)

type RadarHTTPTimeseryBotClassListParamsDateRange string

const (
	RadarHTTPTimeseryBotClassListParamsDateRange1d         RadarHTTPTimeseryBotClassListParamsDateRange = "1d"
	RadarHTTPTimeseryBotClassListParamsDateRange7d         RadarHTTPTimeseryBotClassListParamsDateRange = "7d"
	RadarHTTPTimeseryBotClassListParamsDateRange14d        RadarHTTPTimeseryBotClassListParamsDateRange = "14d"
	RadarHTTPTimeseryBotClassListParamsDateRange28d        RadarHTTPTimeseryBotClassListParamsDateRange = "28d"
	RadarHTTPTimeseryBotClassListParamsDateRange12w        RadarHTTPTimeseryBotClassListParamsDateRange = "12w"
	RadarHTTPTimeseryBotClassListParamsDateRange24w        RadarHTTPTimeseryBotClassListParamsDateRange = "24w"
	RadarHTTPTimeseryBotClassListParamsDateRange52w        RadarHTTPTimeseryBotClassListParamsDateRange = "52w"
	RadarHTTPTimeseryBotClassListParamsDateRange1dControl  RadarHTTPTimeseryBotClassListParamsDateRange = "1dControl"
	RadarHTTPTimeseryBotClassListParamsDateRange7dControl  RadarHTTPTimeseryBotClassListParamsDateRange = "7dControl"
	RadarHTTPTimeseryBotClassListParamsDateRange14dControl RadarHTTPTimeseryBotClassListParamsDateRange = "14dControl"
	RadarHTTPTimeseryBotClassListParamsDateRange28dControl RadarHTTPTimeseryBotClassListParamsDateRange = "28dControl"
	RadarHTTPTimeseryBotClassListParamsDateRange12wControl RadarHTTPTimeseryBotClassListParamsDateRange = "12wControl"
	RadarHTTPTimeseryBotClassListParamsDateRange24wControl RadarHTTPTimeseryBotClassListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseryBotClassListParamsDeviceType string

const (
	RadarHTTPTimeseryBotClassListParamsDeviceTypeDesktop RadarHTTPTimeseryBotClassListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseryBotClassListParamsDeviceTypeMobile  RadarHTTPTimeseryBotClassListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseryBotClassListParamsDeviceTypeOther   RadarHTTPTimeseryBotClassListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseryBotClassListParamsFormat string

const (
	RadarHTTPTimeseryBotClassListParamsFormatJson RadarHTTPTimeseryBotClassListParamsFormat = "JSON"
	RadarHTTPTimeseryBotClassListParamsFormatCsv  RadarHTTPTimeseryBotClassListParamsFormat = "CSV"
)

type RadarHTTPTimeseryBotClassListParamsHTTPProtocol string

const (
	RadarHTTPTimeseryBotClassListParamsHTTPProtocolHTTP  RadarHTTPTimeseryBotClassListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseryBotClassListParamsHTTPProtocolHTTPs RadarHTTPTimeseryBotClassListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseryBotClassListParamsHTTPVersion string

const (
	RadarHTTPTimeseryBotClassListParamsHTTPVersionHttPv1 RadarHTTPTimeseryBotClassListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseryBotClassListParamsHTTPVersionHttPv2 RadarHTTPTimeseryBotClassListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseryBotClassListParamsHTTPVersionHttPv3 RadarHTTPTimeseryBotClassListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseryBotClassListParamsIPVersion string

const (
	RadarHTTPTimeseryBotClassListParamsIPVersionIPv4 RadarHTTPTimeseryBotClassListParamsIPVersion = "IPv4"
	RadarHTTPTimeseryBotClassListParamsIPVersionIPv6 RadarHTTPTimeseryBotClassListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseryBotClassListParamsO string

const (
	RadarHTTPTimeseryBotClassListParamsOWindows  RadarHTTPTimeseryBotClassListParamsO = "WINDOWS"
	RadarHTTPTimeseryBotClassListParamsOMacosx   RadarHTTPTimeseryBotClassListParamsO = "MACOSX"
	RadarHTTPTimeseryBotClassListParamsOAndroid  RadarHTTPTimeseryBotClassListParamsO = "ANDROID"
	RadarHTTPTimeseryBotClassListParamsOChromeos RadarHTTPTimeseryBotClassListParamsO = "CHROMEOS"
	RadarHTTPTimeseryBotClassListParamsOLinux    RadarHTTPTimeseryBotClassListParamsO = "LINUX"
	RadarHTTPTimeseryBotClassListParamsOSmartTv  RadarHTTPTimeseryBotClassListParamsO = "SMART_TV"
)

type RadarHTTPTimeseryBotClassListParamsTlsVersion string

const (
	RadarHTTPTimeseryBotClassListParamsTlsVersionTlSv1_0  RadarHTTPTimeseryBotClassListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseryBotClassListParamsTlsVersionTlSv1_1  RadarHTTPTimeseryBotClassListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseryBotClassListParamsTlsVersionTlSv1_2  RadarHTTPTimeseryBotClassListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseryBotClassListParamsTlsVersionTlSv1_3  RadarHTTPTimeseryBotClassListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseryBotClassListParamsTlsVersionTlSvQuic RadarHTTPTimeseryBotClassListParamsTlsVersion = "TLSvQUIC"
)
