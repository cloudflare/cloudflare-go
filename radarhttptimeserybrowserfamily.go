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

// RadarHTTPTimeseryBrowserFamilyService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewRadarHTTPTimeseryBrowserFamilyService] method instead.
type RadarHTTPTimeseryBrowserFamilyService struct {
	Options []option.RequestOption
}

// NewRadarHTTPTimeseryBrowserFamilyService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewRadarHTTPTimeseryBrowserFamilyService(opts ...option.RequestOption) (r *RadarHTTPTimeseryBrowserFamilyService) {
	r = &RadarHTTPTimeseryBrowserFamilyService{}
	r.Options = opts
	return
}

// Percentage distribution of traffic of the top user agents aggregated in families
// in the selected time range, over time.
func (r *RadarHTTPTimeseryBrowserFamilyService) List(ctx context.Context, query RadarHTTPTimeseryBrowserFamilyListParams, opts ...option.RequestOption) (res *RadarHTTPTimeseryBrowserFamilyListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries/browser_family"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPTimeseryBrowserFamilyListResponse struct {
	Result  RadarHTTPTimeseryBrowserFamilyListResponseResult `json:"result,required"`
	Success bool                                             `json:"success,required"`
	JSON    radarHTTPTimeseryBrowserFamilyListResponseJSON   `json:"-"`
}

// radarHTTPTimeseryBrowserFamilyListResponseJSON contains the JSON metadata for
// the struct [RadarHTTPTimeseryBrowserFamilyListResponse]
type radarHTTPTimeseryBrowserFamilyListResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryBrowserFamilyListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryBrowserFamilyListResponseResult struct {
	Meta   interface{}                                            `json:"meta,required"`
	Serie0 RadarHTTPTimeseryBrowserFamilyListResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPTimeseryBrowserFamilyListResponseResultJSON   `json:"-"`
}

// radarHTTPTimeseryBrowserFamilyListResponseResultJSON contains the JSON metadata
// for the struct [RadarHTTPTimeseryBrowserFamilyListResponseResult]
type radarHTTPTimeseryBrowserFamilyListResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryBrowserFamilyListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryBrowserFamilyListResponseResultSerie0 struct {
	BrowserName []string                                                   `json:"<browser name>,required"`
	Timestamps  []string                                                   `json:"timestamps,required"`
	JSON        radarHTTPTimeseryBrowserFamilyListResponseResultSerie0JSON `json:"-"`
}

// radarHTTPTimeseryBrowserFamilyListResponseResultSerie0JSON contains the JSON
// metadata for the struct [RadarHTTPTimeseryBrowserFamilyListResponseResultSerie0]
type radarHTTPTimeseryBrowserFamilyListResponseResultSerie0JSON struct {
	BrowserName apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPTimeseryBrowserFamilyListResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPTimeseryBrowserFamilyListParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPTimeseryBrowserFamilyListParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	ASN param.Field[[]string] `query:"asn"`
	// Array of datetimes to filter the end of a series.
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPTimeseryBrowserFamilyListParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPTimeseryBrowserFamilyListParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPTimeseryBrowserFamilyListParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPTimeseryBrowserFamilyListParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPTimeseryBrowserFamilyListParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPTimeseryBrowserFamilyListParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPTimeseryBrowserFamilyListParamsO] `query:"os"`
	// Filter for tls version.
	TlsVersion param.Field[[]RadarHTTPTimeseryBrowserFamilyListParamsTlsVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPTimeseryBrowserFamilyListParams]'s query
// parameters as `url.Values`.
func (r RadarHTTPTimeseryBrowserFamilyListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPTimeseryBrowserFamilyListParamsAggInterval string

const (
	RadarHTTPTimeseryBrowserFamilyListParamsAggInterval15m RadarHTTPTimeseryBrowserFamilyListParamsAggInterval = "15m"
	RadarHTTPTimeseryBrowserFamilyListParamsAggInterval1h  RadarHTTPTimeseryBrowserFamilyListParamsAggInterval = "1h"
	RadarHTTPTimeseryBrowserFamilyListParamsAggInterval1d  RadarHTTPTimeseryBrowserFamilyListParamsAggInterval = "1d"
	RadarHTTPTimeseryBrowserFamilyListParamsAggInterval1w  RadarHTTPTimeseryBrowserFamilyListParamsAggInterval = "1w"
)

type RadarHTTPTimeseryBrowserFamilyListParamsDateRange string

const (
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange1d         RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "1d"
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange7d         RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "7d"
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange14d        RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "14d"
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange28d        RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "28d"
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange12w        RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "12w"
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange24w        RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "24w"
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange52w        RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "52w"
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange1dControl  RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "1dControl"
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange7dControl  RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "7dControl"
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange14dControl RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "14dControl"
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange28dControl RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "28dControl"
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange12wControl RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "12wControl"
	RadarHTTPTimeseryBrowserFamilyListParamsDateRange24wControl RadarHTTPTimeseryBrowserFamilyListParamsDateRange = "24wControl"
)

type RadarHTTPTimeseryBrowserFamilyListParamsDeviceType string

const (
	RadarHTTPTimeseryBrowserFamilyListParamsDeviceTypeDesktop RadarHTTPTimeseryBrowserFamilyListParamsDeviceType = "DESKTOP"
	RadarHTTPTimeseryBrowserFamilyListParamsDeviceTypeMobile  RadarHTTPTimeseryBrowserFamilyListParamsDeviceType = "MOBILE"
	RadarHTTPTimeseryBrowserFamilyListParamsDeviceTypeOther   RadarHTTPTimeseryBrowserFamilyListParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPTimeseryBrowserFamilyListParamsFormat string

const (
	RadarHTTPTimeseryBrowserFamilyListParamsFormatJson RadarHTTPTimeseryBrowserFamilyListParamsFormat = "JSON"
	RadarHTTPTimeseryBrowserFamilyListParamsFormatCsv  RadarHTTPTimeseryBrowserFamilyListParamsFormat = "CSV"
)

type RadarHTTPTimeseryBrowserFamilyListParamsHTTPProtocol string

const (
	RadarHTTPTimeseryBrowserFamilyListParamsHTTPProtocolHTTP  RadarHTTPTimeseryBrowserFamilyListParamsHTTPProtocol = "HTTP"
	RadarHTTPTimeseryBrowserFamilyListParamsHTTPProtocolHTTPs RadarHTTPTimeseryBrowserFamilyListParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPTimeseryBrowserFamilyListParamsHTTPVersion string

const (
	RadarHTTPTimeseryBrowserFamilyListParamsHTTPVersionHttPv1 RadarHTTPTimeseryBrowserFamilyListParamsHTTPVersion = "HTTPv1"
	RadarHTTPTimeseryBrowserFamilyListParamsHTTPVersionHttPv2 RadarHTTPTimeseryBrowserFamilyListParamsHTTPVersion = "HTTPv2"
	RadarHTTPTimeseryBrowserFamilyListParamsHTTPVersionHttPv3 RadarHTTPTimeseryBrowserFamilyListParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPTimeseryBrowserFamilyListParamsIPVersion string

const (
	RadarHTTPTimeseryBrowserFamilyListParamsIPVersionIPv4 RadarHTTPTimeseryBrowserFamilyListParamsIPVersion = "IPv4"
	RadarHTTPTimeseryBrowserFamilyListParamsIPVersionIPv6 RadarHTTPTimeseryBrowserFamilyListParamsIPVersion = "IPv6"
)

type RadarHTTPTimeseryBrowserFamilyListParamsO string

const (
	RadarHTTPTimeseryBrowserFamilyListParamsOWindows  RadarHTTPTimeseryBrowserFamilyListParamsO = "WINDOWS"
	RadarHTTPTimeseryBrowserFamilyListParamsOMacosx   RadarHTTPTimeseryBrowserFamilyListParamsO = "MACOSX"
	RadarHTTPTimeseryBrowserFamilyListParamsOAndroid  RadarHTTPTimeseryBrowserFamilyListParamsO = "ANDROID"
	RadarHTTPTimeseryBrowserFamilyListParamsOChromeos RadarHTTPTimeseryBrowserFamilyListParamsO = "CHROMEOS"
	RadarHTTPTimeseryBrowserFamilyListParamsOLinux    RadarHTTPTimeseryBrowserFamilyListParamsO = "LINUX"
	RadarHTTPTimeseryBrowserFamilyListParamsOSmartTv  RadarHTTPTimeseryBrowserFamilyListParamsO = "SMART_TV"
)

type RadarHTTPTimeseryBrowserFamilyListParamsTlsVersion string

const (
	RadarHTTPTimeseryBrowserFamilyListParamsTlsVersionTlSv1_0  RadarHTTPTimeseryBrowserFamilyListParamsTlsVersion = "TLSv1_0"
	RadarHTTPTimeseryBrowserFamilyListParamsTlsVersionTlSv1_1  RadarHTTPTimeseryBrowserFamilyListParamsTlsVersion = "TLSv1_1"
	RadarHTTPTimeseryBrowserFamilyListParamsTlsVersionTlSv1_2  RadarHTTPTimeseryBrowserFamilyListParamsTlsVersion = "TLSv1_2"
	RadarHTTPTimeseryBrowserFamilyListParamsTlsVersionTlSv1_3  RadarHTTPTimeseryBrowserFamilyListParamsTlsVersion = "TLSv1_3"
	RadarHTTPTimeseryBrowserFamilyListParamsTlsVersionTlSvQuic RadarHTTPTimeseryBrowserFamilyListParamsTlsVersion = "TLSvQUIC"
)
