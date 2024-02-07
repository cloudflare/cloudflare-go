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

// RadarHTTPService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewRadarHTTPService] method instead.
type RadarHTTPService struct {
	Options    []option.RequestOption
	TLSVersion *RadarHTTPTLSVersionService
}

// NewRadarHTTPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRadarHTTPService(opts ...option.RequestOption) (r *RadarHTTPService) {
	r = &RadarHTTPService{}
	r.Options = opts
	r.TLSVersion = NewRadarHTTPTLSVersionService(opts...)
	return
}

// Get a time series of the percentage distribution of traffic classified as
// automated or human. Visit
// https://developers.cloudflare.com/radar/concepts/bot-classes/ for more
// information.
func (r *RadarHTTPService) BotClasses(ctx context.Context, query RadarHTTPBotClassesParams, opts ...option.RequestOption) (res *RadarHTTPBotClassesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/bot_class"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a time series of the percentage distribution of traffic of the top user
// agents aggregated in families.
func (r *RadarHTTPService) BrowserFamilies(ctx context.Context, query RadarHTTPBrowserFamiliesParams, opts ...option.RequestOption) (res *RadarHTTPBrowserFamiliesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/browser_family"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a time series of the percentage distribution of traffic of the top user
// agents.
func (r *RadarHTTPService) Browsers(ctx context.Context, query RadarHTTPBrowsersParams, opts ...option.RequestOption) (res *RadarHTTPBrowsersResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/browser"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a time series of the percentage distribution of traffic per device type.
func (r *RadarHTTPService) DeviceTypes(ctx context.Context, query RadarHTTPDeviceTypesParams, opts ...option.RequestOption) (res *RadarHTTPDeviceTypesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/device_type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a time series of the percentage distribution of traffic per HTTP protocol.
func (r *RadarHTTPService) HTTPProtocols(ctx context.Context, query RadarHTTPHTTPProtocolsParams, opts ...option.RequestOption) (res *RadarHttphttpProtocolsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/http_protocol"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a time series of the percentage distribution of traffic per HTTP protocol
// version.
func (r *RadarHTTPService) HTTPVersions(ctx context.Context, query RadarHTTPHTTPVersionsParams, opts ...option.RequestOption) (res *RadarHttphttpVersionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/http_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a time series of the percentage distribution of traffic per IP protocol
// version.
func (r *RadarHTTPService) IPVersions(ctx context.Context, query RadarHTTPIPVersionsParams, opts ...option.RequestOption) (res *RadarHttpipVersionsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/ip_version"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a time series of the percentage distribution of traffic of the top operating
// systems.
func (r *RadarHTTPService) Oss(ctx context.Context, query RadarHTTPOssParams, opts ...option.RequestOption) (res *RadarHTTPOssResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "radar/http/timeseries_groups/os"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type RadarHTTPBotClassesResponse struct {
	Result  RadarHTTPBotClassesResponseResult `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    radarHTTPBotClassesResponseJSON   `json:"-"`
}

// radarHTTPBotClassesResponseJSON contains the JSON metadata for the struct
// [RadarHTTPBotClassesResponse]
type radarHTTPBotClassesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPBotClassesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPBotClassesResponseResult struct {
	Meta   interface{}                             `json:"meta,required"`
	Serie0 RadarHTTPBotClassesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPBotClassesResponseResultJSON   `json:"-"`
}

// radarHTTPBotClassesResponseResultJSON contains the JSON metadata for the struct
// [RadarHTTPBotClassesResponseResult]
type radarHTTPBotClassesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPBotClassesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPBotClassesResponseResultSerie0 struct {
	Bot        []string                                    `json:"bot,required"`
	Human      []string                                    `json:"human,required"`
	Timestamps []string                                    `json:"timestamps,required"`
	JSON       radarHTTPBotClassesResponseResultSerie0JSON `json:"-"`
}

// radarHTTPBotClassesResponseResultSerie0JSON contains the JSON metadata for the
// struct [RadarHTTPBotClassesResponseResultSerie0]
type radarHTTPBotClassesResponseResultSerie0JSON struct {
	Bot         apijson.Field
	Human       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPBotClassesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPBrowserFamiliesResponse struct {
	Result  RadarHTTPBrowserFamiliesResponseResult `json:"result,required"`
	Success bool                                   `json:"success,required"`
	JSON    radarHTTPBrowserFamiliesResponseJSON   `json:"-"`
}

// radarHTTPBrowserFamiliesResponseJSON contains the JSON metadata for the struct
// [RadarHTTPBrowserFamiliesResponse]
type radarHTTPBrowserFamiliesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPBrowserFamiliesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPBrowserFamiliesResponseResult struct {
	Meta   interface{}                                  `json:"meta,required"`
	Serie0 RadarHTTPBrowserFamiliesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPBrowserFamiliesResponseResultJSON   `json:"-"`
}

// radarHTTPBrowserFamiliesResponseResultJSON contains the JSON metadata for the
// struct [RadarHTTPBrowserFamiliesResponseResult]
type radarHTTPBrowserFamiliesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPBrowserFamiliesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPBrowserFamiliesResponseResultSerie0 struct {
	Timestamps []string                                         `json:"timestamps,required"`
	JSON       radarHTTPBrowserFamiliesResponseResultSerie0JSON `json:"-"`
}

// radarHTTPBrowserFamiliesResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarHTTPBrowserFamiliesResponseResultSerie0]
type radarHTTPBrowserFamiliesResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPBrowserFamiliesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPBrowsersResponse struct {
	Result  RadarHTTPBrowsersResponseResult `json:"result,required"`
	Success bool                            `json:"success,required"`
	JSON    radarHTTPBrowsersResponseJSON   `json:"-"`
}

// radarHTTPBrowsersResponseJSON contains the JSON metadata for the struct
// [RadarHTTPBrowsersResponse]
type radarHTTPBrowsersResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPBrowsersResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPBrowsersResponseResult struct {
	Meta   interface{}                           `json:"meta,required"`
	Serie0 RadarHTTPBrowsersResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPBrowsersResponseResultJSON   `json:"-"`
}

// radarHTTPBrowsersResponseResultJSON contains the JSON metadata for the struct
// [RadarHTTPBrowsersResponseResult]
type radarHTTPBrowsersResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPBrowsersResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPBrowsersResponseResultSerie0 struct {
	Timestamps []string                                  `json:"timestamps,required"`
	JSON       radarHTTPBrowsersResponseResultSerie0JSON `json:"-"`
}

// radarHTTPBrowsersResponseResultSerie0JSON contains the JSON metadata for the
// struct [RadarHTTPBrowsersResponseResultSerie0]
type radarHTTPBrowsersResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPBrowsersResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPDeviceTypesResponse struct {
	Result  RadarHTTPDeviceTypesResponseResult `json:"result,required"`
	Success bool                               `json:"success,required"`
	JSON    radarHTTPDeviceTypesResponseJSON   `json:"-"`
}

// radarHTTPDeviceTypesResponseJSON contains the JSON metadata for the struct
// [RadarHTTPDeviceTypesResponse]
type radarHTTPDeviceTypesResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPDeviceTypesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPDeviceTypesResponseResult struct {
	Meta   interface{}                              `json:"meta,required"`
	Serie0 RadarHTTPDeviceTypesResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPDeviceTypesResponseResultJSON   `json:"-"`
}

// radarHTTPDeviceTypesResponseResultJSON contains the JSON metadata for the struct
// [RadarHTTPDeviceTypesResponseResult]
type radarHTTPDeviceTypesResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPDeviceTypesResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPDeviceTypesResponseResultSerie0 struct {
	Desktop    []string                                     `json:"desktop,required"`
	Mobile     []string                                     `json:"mobile,required"`
	Other      []string                                     `json:"other,required"`
	Timestamps []string                                     `json:"timestamps,required"`
	JSON       radarHTTPDeviceTypesResponseResultSerie0JSON `json:"-"`
}

// radarHTTPDeviceTypesResponseResultSerie0JSON contains the JSON metadata for the
// struct [RadarHTTPDeviceTypesResponseResultSerie0]
type radarHTTPDeviceTypesResponseResultSerie0JSON struct {
	Desktop     apijson.Field
	Mobile      apijson.Field
	Other       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPDeviceTypesResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHttphttpProtocolsResponse struct {
	Result  RadarHttphttpProtocolsResponseResult `json:"result,required"`
	Success bool                                 `json:"success,required"`
	JSON    radarHttphttpProtocolsResponseJSON   `json:"-"`
}

// radarHttphttpProtocolsResponseJSON contains the JSON metadata for the struct
// [RadarHttphttpProtocolsResponse]
type radarHttphttpProtocolsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHttphttpProtocolsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHttphttpProtocolsResponseResult struct {
	Meta   interface{}                                `json:"meta,required"`
	Serie0 RadarHttphttpProtocolsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHttphttpProtocolsResponseResultJSON   `json:"-"`
}

// radarHttphttpProtocolsResponseResultJSON contains the JSON metadata for the
// struct [RadarHttphttpProtocolsResponseResult]
type radarHttphttpProtocolsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHttphttpProtocolsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHttphttpProtocolsResponseResultSerie0 struct {
	HTTP       []string                                       `json:"http,required"`
	HTTPs      []string                                       `json:"https,required"`
	Timestamps []string                                       `json:"timestamps,required"`
	JSON       radarHttphttpProtocolsResponseResultSerie0JSON `json:"-"`
}

// radarHttphttpProtocolsResponseResultSerie0JSON contains the JSON metadata for
// the struct [RadarHttphttpProtocolsResponseResultSerie0]
type radarHttphttpProtocolsResponseResultSerie0JSON struct {
	HTTP        apijson.Field
	HTTPs       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHttphttpProtocolsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHttphttpVersionsResponse struct {
	Result  RadarHttphttpVersionsResponseResult `json:"result,required"`
	Success bool                                `json:"success,required"`
	JSON    radarHttphttpVersionsResponseJSON   `json:"-"`
}

// radarHttphttpVersionsResponseJSON contains the JSON metadata for the struct
// [RadarHttphttpVersionsResponse]
type radarHttphttpVersionsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHttphttpVersionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHttphttpVersionsResponseResult struct {
	Meta   interface{}                               `json:"meta,required"`
	Serie0 RadarHttphttpVersionsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHttphttpVersionsResponseResultJSON   `json:"-"`
}

// radarHttphttpVersionsResponseResultJSON contains the JSON metadata for the
// struct [RadarHttphttpVersionsResponseResult]
type radarHttphttpVersionsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHttphttpVersionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHttphttpVersionsResponseResultSerie0 struct {
	HTTP1X     []string                                      `json:"HTTP/1.x,required"`
	HTTP2      []string                                      `json:"HTTP/2,required"`
	HTTP3      []string                                      `json:"HTTP/3,required"`
	Timestamps []string                                      `json:"timestamps,required"`
	JSON       radarHttphttpVersionsResponseResultSerie0JSON `json:"-"`
}

// radarHttphttpVersionsResponseResultSerie0JSON contains the JSON metadata for the
// struct [RadarHttphttpVersionsResponseResultSerie0]
type radarHttphttpVersionsResponseResultSerie0JSON struct {
	HTTP1X      apijson.Field
	HTTP2       apijson.Field
	HTTP3       apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHttphttpVersionsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHttpipVersionsResponse struct {
	Result  RadarHttpipVersionsResponseResult `json:"result,required"`
	Success bool                              `json:"success,required"`
	JSON    radarHttpipVersionsResponseJSON   `json:"-"`
}

// radarHttpipVersionsResponseJSON contains the JSON metadata for the struct
// [RadarHttpipVersionsResponse]
type radarHttpipVersionsResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHttpipVersionsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHttpipVersionsResponseResult struct {
	Meta   interface{}                             `json:"meta,required"`
	Serie0 RadarHttpipVersionsResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHttpipVersionsResponseResultJSON   `json:"-"`
}

// radarHttpipVersionsResponseResultJSON contains the JSON metadata for the struct
// [RadarHttpipVersionsResponseResult]
type radarHttpipVersionsResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHttpipVersionsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHttpipVersionsResponseResultSerie0 struct {
	IPv4       []string                                    `json:"IPv4,required"`
	IPv6       []string                                    `json:"IPv6,required"`
	Timestamps []string                                    `json:"timestamps,required"`
	JSON       radarHttpipVersionsResponseResultSerie0JSON `json:"-"`
}

// radarHttpipVersionsResponseResultSerie0JSON contains the JSON metadata for the
// struct [RadarHttpipVersionsResponseResultSerie0]
type radarHttpipVersionsResponseResultSerie0JSON struct {
	IPv4        apijson.Field
	IPv6        apijson.Field
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHttpipVersionsResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPOssResponse struct {
	Result  RadarHTTPOssResponseResult `json:"result,required"`
	Success bool                       `json:"success,required"`
	JSON    radarHTTPOssResponseJSON   `json:"-"`
}

// radarHTTPOssResponseJSON contains the JSON metadata for the struct
// [RadarHTTPOssResponse]
type radarHTTPOssResponseJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPOssResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPOssResponseResult struct {
	Meta   interface{}                      `json:"meta,required"`
	Serie0 RadarHTTPOssResponseResultSerie0 `json:"serie_0,required"`
	JSON   radarHTTPOssResponseResultJSON   `json:"-"`
}

// radarHTTPOssResponseResultJSON contains the JSON metadata for the struct
// [RadarHTTPOssResponseResult]
type radarHTTPOssResponseResultJSON struct {
	Meta        apijson.Field
	Serie0      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPOssResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPOssResponseResultSerie0 struct {
	Timestamps []string                             `json:"timestamps,required"`
	JSON       radarHTTPOssResponseResultSerie0JSON `json:"-"`
}

// radarHTTPOssResponseResultSerie0JSON contains the JSON metadata for the struct
// [RadarHTTPOssResponseResultSerie0]
type radarHTTPOssResponseResultSerie0JSON struct {
	Timestamps  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RadarHTTPOssResponseResultSerie0) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type RadarHTTPBotClassesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPBotClassesParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPBotClassesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPBotClassesParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPBotClassesParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPBotClassesParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPBotClassesParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPBotClassesParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPBotClassesParamsO] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPBotClassesParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPBotClassesParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPBotClassesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPBotClassesParamsAggInterval string

const (
	RadarHTTPBotClassesParamsAggInterval15m RadarHTTPBotClassesParamsAggInterval = "15m"
	RadarHTTPBotClassesParamsAggInterval1h  RadarHTTPBotClassesParamsAggInterval = "1h"
	RadarHTTPBotClassesParamsAggInterval1d  RadarHTTPBotClassesParamsAggInterval = "1d"
	RadarHTTPBotClassesParamsAggInterval1w  RadarHTTPBotClassesParamsAggInterval = "1w"
)

type RadarHTTPBotClassesParamsDateRange string

const (
	RadarHTTPBotClassesParamsDateRange1d         RadarHTTPBotClassesParamsDateRange = "1d"
	RadarHTTPBotClassesParamsDateRange2d         RadarHTTPBotClassesParamsDateRange = "2d"
	RadarHTTPBotClassesParamsDateRange7d         RadarHTTPBotClassesParamsDateRange = "7d"
	RadarHTTPBotClassesParamsDateRange14d        RadarHTTPBotClassesParamsDateRange = "14d"
	RadarHTTPBotClassesParamsDateRange28d        RadarHTTPBotClassesParamsDateRange = "28d"
	RadarHTTPBotClassesParamsDateRange12w        RadarHTTPBotClassesParamsDateRange = "12w"
	RadarHTTPBotClassesParamsDateRange24w        RadarHTTPBotClassesParamsDateRange = "24w"
	RadarHTTPBotClassesParamsDateRange52w        RadarHTTPBotClassesParamsDateRange = "52w"
	RadarHTTPBotClassesParamsDateRange1dControl  RadarHTTPBotClassesParamsDateRange = "1dControl"
	RadarHTTPBotClassesParamsDateRange2dControl  RadarHTTPBotClassesParamsDateRange = "2dControl"
	RadarHTTPBotClassesParamsDateRange7dControl  RadarHTTPBotClassesParamsDateRange = "7dControl"
	RadarHTTPBotClassesParamsDateRange14dControl RadarHTTPBotClassesParamsDateRange = "14dControl"
	RadarHTTPBotClassesParamsDateRange28dControl RadarHTTPBotClassesParamsDateRange = "28dControl"
	RadarHTTPBotClassesParamsDateRange12wControl RadarHTTPBotClassesParamsDateRange = "12wControl"
	RadarHTTPBotClassesParamsDateRange24wControl RadarHTTPBotClassesParamsDateRange = "24wControl"
)

type RadarHTTPBotClassesParamsDeviceType string

const (
	RadarHTTPBotClassesParamsDeviceTypeDesktop RadarHTTPBotClassesParamsDeviceType = "DESKTOP"
	RadarHTTPBotClassesParamsDeviceTypeMobile  RadarHTTPBotClassesParamsDeviceType = "MOBILE"
	RadarHTTPBotClassesParamsDeviceTypeOther   RadarHTTPBotClassesParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPBotClassesParamsFormat string

const (
	RadarHTTPBotClassesParamsFormatJson RadarHTTPBotClassesParamsFormat = "JSON"
	RadarHTTPBotClassesParamsFormatCsv  RadarHTTPBotClassesParamsFormat = "CSV"
)

type RadarHTTPBotClassesParamsHTTPProtocol string

const (
	RadarHTTPBotClassesParamsHTTPProtocolHTTP  RadarHTTPBotClassesParamsHTTPProtocol = "HTTP"
	RadarHTTPBotClassesParamsHTTPProtocolHTTPs RadarHTTPBotClassesParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPBotClassesParamsHTTPVersion string

const (
	RadarHTTPBotClassesParamsHTTPVersionHttPv1 RadarHTTPBotClassesParamsHTTPVersion = "HTTPv1"
	RadarHTTPBotClassesParamsHTTPVersionHttPv2 RadarHTTPBotClassesParamsHTTPVersion = "HTTPv2"
	RadarHTTPBotClassesParamsHTTPVersionHttPv3 RadarHTTPBotClassesParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPBotClassesParamsIPVersion string

const (
	RadarHTTPBotClassesParamsIPVersionIPv4 RadarHTTPBotClassesParamsIPVersion = "IPv4"
	RadarHTTPBotClassesParamsIPVersionIPv6 RadarHTTPBotClassesParamsIPVersion = "IPv6"
)

type RadarHTTPBotClassesParamsO string

const (
	RadarHTTPBotClassesParamsOWindows  RadarHTTPBotClassesParamsO = "WINDOWS"
	RadarHTTPBotClassesParamsOMacosx   RadarHTTPBotClassesParamsO = "MACOSX"
	RadarHTTPBotClassesParamsOIos      RadarHTTPBotClassesParamsO = "IOS"
	RadarHTTPBotClassesParamsOAndroid  RadarHTTPBotClassesParamsO = "ANDROID"
	RadarHTTPBotClassesParamsOChromeos RadarHTTPBotClassesParamsO = "CHROMEOS"
	RadarHTTPBotClassesParamsOLinux    RadarHTTPBotClassesParamsO = "LINUX"
	RadarHTTPBotClassesParamsOSmartTv  RadarHTTPBotClassesParamsO = "SMART_TV"
)

type RadarHTTPBotClassesParamsTLSVersion string

const (
	RadarHTTPBotClassesParamsTLSVersionTlSv1_0  RadarHTTPBotClassesParamsTLSVersion = "TLSv1_0"
	RadarHTTPBotClassesParamsTLSVersionTlSv1_1  RadarHTTPBotClassesParamsTLSVersion = "TLSv1_1"
	RadarHTTPBotClassesParamsTLSVersionTlSv1_2  RadarHTTPBotClassesParamsTLSVersion = "TLSv1_2"
	RadarHTTPBotClassesParamsTLSVersionTlSv1_3  RadarHTTPBotClassesParamsTLSVersion = "TLSv1_3"
	RadarHTTPBotClassesParamsTLSVersionTlSvQuic RadarHTTPBotClassesParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPBrowserFamiliesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPBrowserFamiliesParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPBrowserFamiliesParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPBrowserFamiliesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPBrowserFamiliesParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPBrowserFamiliesParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPBrowserFamiliesParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPBrowserFamiliesParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPBrowserFamiliesParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPBrowserFamiliesParamsO] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPBrowserFamiliesParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPBrowserFamiliesParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPBrowserFamiliesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPBrowserFamiliesParamsAggInterval string

const (
	RadarHTTPBrowserFamiliesParamsAggInterval15m RadarHTTPBrowserFamiliesParamsAggInterval = "15m"
	RadarHTTPBrowserFamiliesParamsAggInterval1h  RadarHTTPBrowserFamiliesParamsAggInterval = "1h"
	RadarHTTPBrowserFamiliesParamsAggInterval1d  RadarHTTPBrowserFamiliesParamsAggInterval = "1d"
	RadarHTTPBrowserFamiliesParamsAggInterval1w  RadarHTTPBrowserFamiliesParamsAggInterval = "1w"
)

type RadarHTTPBrowserFamiliesParamsBotClass string

const (
	RadarHTTPBrowserFamiliesParamsBotClassLikelyAutomated RadarHTTPBrowserFamiliesParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPBrowserFamiliesParamsBotClassLikelyHuman     RadarHTTPBrowserFamiliesParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPBrowserFamiliesParamsDateRange string

const (
	RadarHTTPBrowserFamiliesParamsDateRange1d         RadarHTTPBrowserFamiliesParamsDateRange = "1d"
	RadarHTTPBrowserFamiliesParamsDateRange2d         RadarHTTPBrowserFamiliesParamsDateRange = "2d"
	RadarHTTPBrowserFamiliesParamsDateRange7d         RadarHTTPBrowserFamiliesParamsDateRange = "7d"
	RadarHTTPBrowserFamiliesParamsDateRange14d        RadarHTTPBrowserFamiliesParamsDateRange = "14d"
	RadarHTTPBrowserFamiliesParamsDateRange28d        RadarHTTPBrowserFamiliesParamsDateRange = "28d"
	RadarHTTPBrowserFamiliesParamsDateRange12w        RadarHTTPBrowserFamiliesParamsDateRange = "12w"
	RadarHTTPBrowserFamiliesParamsDateRange24w        RadarHTTPBrowserFamiliesParamsDateRange = "24w"
	RadarHTTPBrowserFamiliesParamsDateRange52w        RadarHTTPBrowserFamiliesParamsDateRange = "52w"
	RadarHTTPBrowserFamiliesParamsDateRange1dControl  RadarHTTPBrowserFamiliesParamsDateRange = "1dControl"
	RadarHTTPBrowserFamiliesParamsDateRange2dControl  RadarHTTPBrowserFamiliesParamsDateRange = "2dControl"
	RadarHTTPBrowserFamiliesParamsDateRange7dControl  RadarHTTPBrowserFamiliesParamsDateRange = "7dControl"
	RadarHTTPBrowserFamiliesParamsDateRange14dControl RadarHTTPBrowserFamiliesParamsDateRange = "14dControl"
	RadarHTTPBrowserFamiliesParamsDateRange28dControl RadarHTTPBrowserFamiliesParamsDateRange = "28dControl"
	RadarHTTPBrowserFamiliesParamsDateRange12wControl RadarHTTPBrowserFamiliesParamsDateRange = "12wControl"
	RadarHTTPBrowserFamiliesParamsDateRange24wControl RadarHTTPBrowserFamiliesParamsDateRange = "24wControl"
)

type RadarHTTPBrowserFamiliesParamsDeviceType string

const (
	RadarHTTPBrowserFamiliesParamsDeviceTypeDesktop RadarHTTPBrowserFamiliesParamsDeviceType = "DESKTOP"
	RadarHTTPBrowserFamiliesParamsDeviceTypeMobile  RadarHTTPBrowserFamiliesParamsDeviceType = "MOBILE"
	RadarHTTPBrowserFamiliesParamsDeviceTypeOther   RadarHTTPBrowserFamiliesParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPBrowserFamiliesParamsFormat string

const (
	RadarHTTPBrowserFamiliesParamsFormatJson RadarHTTPBrowserFamiliesParamsFormat = "JSON"
	RadarHTTPBrowserFamiliesParamsFormatCsv  RadarHTTPBrowserFamiliesParamsFormat = "CSV"
)

type RadarHTTPBrowserFamiliesParamsHTTPProtocol string

const (
	RadarHTTPBrowserFamiliesParamsHTTPProtocolHTTP  RadarHTTPBrowserFamiliesParamsHTTPProtocol = "HTTP"
	RadarHTTPBrowserFamiliesParamsHTTPProtocolHTTPs RadarHTTPBrowserFamiliesParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPBrowserFamiliesParamsHTTPVersion string

const (
	RadarHTTPBrowserFamiliesParamsHTTPVersionHttPv1 RadarHTTPBrowserFamiliesParamsHTTPVersion = "HTTPv1"
	RadarHTTPBrowserFamiliesParamsHTTPVersionHttPv2 RadarHTTPBrowserFamiliesParamsHTTPVersion = "HTTPv2"
	RadarHTTPBrowserFamiliesParamsHTTPVersionHttPv3 RadarHTTPBrowserFamiliesParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPBrowserFamiliesParamsIPVersion string

const (
	RadarHTTPBrowserFamiliesParamsIPVersionIPv4 RadarHTTPBrowserFamiliesParamsIPVersion = "IPv4"
	RadarHTTPBrowserFamiliesParamsIPVersionIPv6 RadarHTTPBrowserFamiliesParamsIPVersion = "IPv6"
)

type RadarHTTPBrowserFamiliesParamsO string

const (
	RadarHTTPBrowserFamiliesParamsOWindows  RadarHTTPBrowserFamiliesParamsO = "WINDOWS"
	RadarHTTPBrowserFamiliesParamsOMacosx   RadarHTTPBrowserFamiliesParamsO = "MACOSX"
	RadarHTTPBrowserFamiliesParamsOIos      RadarHTTPBrowserFamiliesParamsO = "IOS"
	RadarHTTPBrowserFamiliesParamsOAndroid  RadarHTTPBrowserFamiliesParamsO = "ANDROID"
	RadarHTTPBrowserFamiliesParamsOChromeos RadarHTTPBrowserFamiliesParamsO = "CHROMEOS"
	RadarHTTPBrowserFamiliesParamsOLinux    RadarHTTPBrowserFamiliesParamsO = "LINUX"
	RadarHTTPBrowserFamiliesParamsOSmartTv  RadarHTTPBrowserFamiliesParamsO = "SMART_TV"
)

type RadarHTTPBrowserFamiliesParamsTLSVersion string

const (
	RadarHTTPBrowserFamiliesParamsTLSVersionTlSv1_0  RadarHTTPBrowserFamiliesParamsTLSVersion = "TLSv1_0"
	RadarHTTPBrowserFamiliesParamsTLSVersionTlSv1_1  RadarHTTPBrowserFamiliesParamsTLSVersion = "TLSv1_1"
	RadarHTTPBrowserFamiliesParamsTLSVersionTlSv1_2  RadarHTTPBrowserFamiliesParamsTLSVersion = "TLSv1_2"
	RadarHTTPBrowserFamiliesParamsTLSVersionTlSv1_3  RadarHTTPBrowserFamiliesParamsTLSVersion = "TLSv1_3"
	RadarHTTPBrowserFamiliesParamsTLSVersionTlSvQuic RadarHTTPBrowserFamiliesParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPBrowsersParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPBrowsersParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPBrowsersParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPBrowsersParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPBrowsersParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPBrowsersParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPBrowsersParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPBrowsersParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPBrowsersParamsIPVersion] `query:"ipVersion"`
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
	Os param.Field[[]RadarHTTPBrowsersParamsO] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPBrowsersParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPBrowsersParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPBrowsersParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPBrowsersParamsAggInterval string

const (
	RadarHTTPBrowsersParamsAggInterval15m RadarHTTPBrowsersParamsAggInterval = "15m"
	RadarHTTPBrowsersParamsAggInterval1h  RadarHTTPBrowsersParamsAggInterval = "1h"
	RadarHTTPBrowsersParamsAggInterval1d  RadarHTTPBrowsersParamsAggInterval = "1d"
	RadarHTTPBrowsersParamsAggInterval1w  RadarHTTPBrowsersParamsAggInterval = "1w"
)

type RadarHTTPBrowsersParamsBotClass string

const (
	RadarHTTPBrowsersParamsBotClassLikelyAutomated RadarHTTPBrowsersParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPBrowsersParamsBotClassLikelyHuman     RadarHTTPBrowsersParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPBrowsersParamsDateRange string

const (
	RadarHTTPBrowsersParamsDateRange1d         RadarHTTPBrowsersParamsDateRange = "1d"
	RadarHTTPBrowsersParamsDateRange2d         RadarHTTPBrowsersParamsDateRange = "2d"
	RadarHTTPBrowsersParamsDateRange7d         RadarHTTPBrowsersParamsDateRange = "7d"
	RadarHTTPBrowsersParamsDateRange14d        RadarHTTPBrowsersParamsDateRange = "14d"
	RadarHTTPBrowsersParamsDateRange28d        RadarHTTPBrowsersParamsDateRange = "28d"
	RadarHTTPBrowsersParamsDateRange12w        RadarHTTPBrowsersParamsDateRange = "12w"
	RadarHTTPBrowsersParamsDateRange24w        RadarHTTPBrowsersParamsDateRange = "24w"
	RadarHTTPBrowsersParamsDateRange52w        RadarHTTPBrowsersParamsDateRange = "52w"
	RadarHTTPBrowsersParamsDateRange1dControl  RadarHTTPBrowsersParamsDateRange = "1dControl"
	RadarHTTPBrowsersParamsDateRange2dControl  RadarHTTPBrowsersParamsDateRange = "2dControl"
	RadarHTTPBrowsersParamsDateRange7dControl  RadarHTTPBrowsersParamsDateRange = "7dControl"
	RadarHTTPBrowsersParamsDateRange14dControl RadarHTTPBrowsersParamsDateRange = "14dControl"
	RadarHTTPBrowsersParamsDateRange28dControl RadarHTTPBrowsersParamsDateRange = "28dControl"
	RadarHTTPBrowsersParamsDateRange12wControl RadarHTTPBrowsersParamsDateRange = "12wControl"
	RadarHTTPBrowsersParamsDateRange24wControl RadarHTTPBrowsersParamsDateRange = "24wControl"
)

type RadarHTTPBrowsersParamsDeviceType string

const (
	RadarHTTPBrowsersParamsDeviceTypeDesktop RadarHTTPBrowsersParamsDeviceType = "DESKTOP"
	RadarHTTPBrowsersParamsDeviceTypeMobile  RadarHTTPBrowsersParamsDeviceType = "MOBILE"
	RadarHTTPBrowsersParamsDeviceTypeOther   RadarHTTPBrowsersParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPBrowsersParamsFormat string

const (
	RadarHTTPBrowsersParamsFormatJson RadarHTTPBrowsersParamsFormat = "JSON"
	RadarHTTPBrowsersParamsFormatCsv  RadarHTTPBrowsersParamsFormat = "CSV"
)

type RadarHTTPBrowsersParamsHTTPProtocol string

const (
	RadarHTTPBrowsersParamsHTTPProtocolHTTP  RadarHTTPBrowsersParamsHTTPProtocol = "HTTP"
	RadarHTTPBrowsersParamsHTTPProtocolHTTPs RadarHTTPBrowsersParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPBrowsersParamsHTTPVersion string

const (
	RadarHTTPBrowsersParamsHTTPVersionHttPv1 RadarHTTPBrowsersParamsHTTPVersion = "HTTPv1"
	RadarHTTPBrowsersParamsHTTPVersionHttPv2 RadarHTTPBrowsersParamsHTTPVersion = "HTTPv2"
	RadarHTTPBrowsersParamsHTTPVersionHttPv3 RadarHTTPBrowsersParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPBrowsersParamsIPVersion string

const (
	RadarHTTPBrowsersParamsIPVersionIPv4 RadarHTTPBrowsersParamsIPVersion = "IPv4"
	RadarHTTPBrowsersParamsIPVersionIPv6 RadarHTTPBrowsersParamsIPVersion = "IPv6"
)

type RadarHTTPBrowsersParamsO string

const (
	RadarHTTPBrowsersParamsOWindows  RadarHTTPBrowsersParamsO = "WINDOWS"
	RadarHTTPBrowsersParamsOMacosx   RadarHTTPBrowsersParamsO = "MACOSX"
	RadarHTTPBrowsersParamsOIos      RadarHTTPBrowsersParamsO = "IOS"
	RadarHTTPBrowsersParamsOAndroid  RadarHTTPBrowsersParamsO = "ANDROID"
	RadarHTTPBrowsersParamsOChromeos RadarHTTPBrowsersParamsO = "CHROMEOS"
	RadarHTTPBrowsersParamsOLinux    RadarHTTPBrowsersParamsO = "LINUX"
	RadarHTTPBrowsersParamsOSmartTv  RadarHTTPBrowsersParamsO = "SMART_TV"
)

type RadarHTTPBrowsersParamsTLSVersion string

const (
	RadarHTTPBrowsersParamsTLSVersionTlSv1_0  RadarHTTPBrowsersParamsTLSVersion = "TLSv1_0"
	RadarHTTPBrowsersParamsTLSVersionTlSv1_1  RadarHTTPBrowsersParamsTLSVersion = "TLSv1_1"
	RadarHTTPBrowsersParamsTLSVersionTlSv1_2  RadarHTTPBrowsersParamsTLSVersion = "TLSv1_2"
	RadarHTTPBrowsersParamsTLSVersionTlSv1_3  RadarHTTPBrowsersParamsTLSVersion = "TLSv1_3"
	RadarHTTPBrowsersParamsTLSVersionTlSvQuic RadarHTTPBrowsersParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPDeviceTypesParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPDeviceTypesParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPDeviceTypesParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPDeviceTypesParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Format results are returned in.
	Format param.Field[RadarHTTPDeviceTypesParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPDeviceTypesParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPDeviceTypesParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPDeviceTypesParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHTTPDeviceTypesParamsO] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPDeviceTypesParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPDeviceTypesParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPDeviceTypesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPDeviceTypesParamsAggInterval string

const (
	RadarHTTPDeviceTypesParamsAggInterval15m RadarHTTPDeviceTypesParamsAggInterval = "15m"
	RadarHTTPDeviceTypesParamsAggInterval1h  RadarHTTPDeviceTypesParamsAggInterval = "1h"
	RadarHTTPDeviceTypesParamsAggInterval1d  RadarHTTPDeviceTypesParamsAggInterval = "1d"
	RadarHTTPDeviceTypesParamsAggInterval1w  RadarHTTPDeviceTypesParamsAggInterval = "1w"
)

type RadarHTTPDeviceTypesParamsBotClass string

const (
	RadarHTTPDeviceTypesParamsBotClassLikelyAutomated RadarHTTPDeviceTypesParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPDeviceTypesParamsBotClassLikelyHuman     RadarHTTPDeviceTypesParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPDeviceTypesParamsDateRange string

const (
	RadarHTTPDeviceTypesParamsDateRange1d         RadarHTTPDeviceTypesParamsDateRange = "1d"
	RadarHTTPDeviceTypesParamsDateRange2d         RadarHTTPDeviceTypesParamsDateRange = "2d"
	RadarHTTPDeviceTypesParamsDateRange7d         RadarHTTPDeviceTypesParamsDateRange = "7d"
	RadarHTTPDeviceTypesParamsDateRange14d        RadarHTTPDeviceTypesParamsDateRange = "14d"
	RadarHTTPDeviceTypesParamsDateRange28d        RadarHTTPDeviceTypesParamsDateRange = "28d"
	RadarHTTPDeviceTypesParamsDateRange12w        RadarHTTPDeviceTypesParamsDateRange = "12w"
	RadarHTTPDeviceTypesParamsDateRange24w        RadarHTTPDeviceTypesParamsDateRange = "24w"
	RadarHTTPDeviceTypesParamsDateRange52w        RadarHTTPDeviceTypesParamsDateRange = "52w"
	RadarHTTPDeviceTypesParamsDateRange1dControl  RadarHTTPDeviceTypesParamsDateRange = "1dControl"
	RadarHTTPDeviceTypesParamsDateRange2dControl  RadarHTTPDeviceTypesParamsDateRange = "2dControl"
	RadarHTTPDeviceTypesParamsDateRange7dControl  RadarHTTPDeviceTypesParamsDateRange = "7dControl"
	RadarHTTPDeviceTypesParamsDateRange14dControl RadarHTTPDeviceTypesParamsDateRange = "14dControl"
	RadarHTTPDeviceTypesParamsDateRange28dControl RadarHTTPDeviceTypesParamsDateRange = "28dControl"
	RadarHTTPDeviceTypesParamsDateRange12wControl RadarHTTPDeviceTypesParamsDateRange = "12wControl"
	RadarHTTPDeviceTypesParamsDateRange24wControl RadarHTTPDeviceTypesParamsDateRange = "24wControl"
)

// Format results are returned in.
type RadarHTTPDeviceTypesParamsFormat string

const (
	RadarHTTPDeviceTypesParamsFormatJson RadarHTTPDeviceTypesParamsFormat = "JSON"
	RadarHTTPDeviceTypesParamsFormatCsv  RadarHTTPDeviceTypesParamsFormat = "CSV"
)

type RadarHTTPDeviceTypesParamsHTTPProtocol string

const (
	RadarHTTPDeviceTypesParamsHTTPProtocolHTTP  RadarHTTPDeviceTypesParamsHTTPProtocol = "HTTP"
	RadarHTTPDeviceTypesParamsHTTPProtocolHTTPs RadarHTTPDeviceTypesParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPDeviceTypesParamsHTTPVersion string

const (
	RadarHTTPDeviceTypesParamsHTTPVersionHttPv1 RadarHTTPDeviceTypesParamsHTTPVersion = "HTTPv1"
	RadarHTTPDeviceTypesParamsHTTPVersionHttPv2 RadarHTTPDeviceTypesParamsHTTPVersion = "HTTPv2"
	RadarHTTPDeviceTypesParamsHTTPVersionHttPv3 RadarHTTPDeviceTypesParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPDeviceTypesParamsIPVersion string

const (
	RadarHTTPDeviceTypesParamsIPVersionIPv4 RadarHTTPDeviceTypesParamsIPVersion = "IPv4"
	RadarHTTPDeviceTypesParamsIPVersionIPv6 RadarHTTPDeviceTypesParamsIPVersion = "IPv6"
)

type RadarHTTPDeviceTypesParamsO string

const (
	RadarHTTPDeviceTypesParamsOWindows  RadarHTTPDeviceTypesParamsO = "WINDOWS"
	RadarHTTPDeviceTypesParamsOMacosx   RadarHTTPDeviceTypesParamsO = "MACOSX"
	RadarHTTPDeviceTypesParamsOIos      RadarHTTPDeviceTypesParamsO = "IOS"
	RadarHTTPDeviceTypesParamsOAndroid  RadarHTTPDeviceTypesParamsO = "ANDROID"
	RadarHTTPDeviceTypesParamsOChromeos RadarHTTPDeviceTypesParamsO = "CHROMEOS"
	RadarHTTPDeviceTypesParamsOLinux    RadarHTTPDeviceTypesParamsO = "LINUX"
	RadarHTTPDeviceTypesParamsOSmartTv  RadarHTTPDeviceTypesParamsO = "SMART_TV"
)

type RadarHTTPDeviceTypesParamsTLSVersion string

const (
	RadarHTTPDeviceTypesParamsTLSVersionTlSv1_0  RadarHTTPDeviceTypesParamsTLSVersion = "TLSv1_0"
	RadarHTTPDeviceTypesParamsTLSVersionTlSv1_1  RadarHTTPDeviceTypesParamsTLSVersion = "TLSv1_1"
	RadarHTTPDeviceTypesParamsTLSVersionTlSv1_2  RadarHTTPDeviceTypesParamsTLSVersion = "TLSv1_2"
	RadarHTTPDeviceTypesParamsTLSVersionTlSv1_3  RadarHTTPDeviceTypesParamsTLSVersion = "TLSv1_3"
	RadarHTTPDeviceTypesParamsTLSVersionTlSvQuic RadarHTTPDeviceTypesParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPHTTPProtocolsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHttphttpProtocolsParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHttphttpProtocolsParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHttphttpProtocolsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHttphttpProtocolsParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHttphttpProtocolsParamsFormat] `query:"format"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHttphttpProtocolsParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHttphttpProtocolsParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHttphttpProtocolsParamsO] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHttphttpProtocolsParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPHTTPProtocolsParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPHTTPProtocolsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHttphttpProtocolsParamsAggInterval string

const (
	RadarHttphttpProtocolsParamsAggInterval15m RadarHttphttpProtocolsParamsAggInterval = "15m"
	RadarHttphttpProtocolsParamsAggInterval1h  RadarHttphttpProtocolsParamsAggInterval = "1h"
	RadarHttphttpProtocolsParamsAggInterval1d  RadarHttphttpProtocolsParamsAggInterval = "1d"
	RadarHttphttpProtocolsParamsAggInterval1w  RadarHttphttpProtocolsParamsAggInterval = "1w"
)

type RadarHttphttpProtocolsParamsBotClass string

const (
	RadarHttphttpProtocolsParamsBotClassLikelyAutomated RadarHttphttpProtocolsParamsBotClass = "LIKELY_AUTOMATED"
	RadarHttphttpProtocolsParamsBotClassLikelyHuman     RadarHttphttpProtocolsParamsBotClass = "LIKELY_HUMAN"
)

type RadarHttphttpProtocolsParamsDateRange string

const (
	RadarHttphttpProtocolsParamsDateRange1d         RadarHttphttpProtocolsParamsDateRange = "1d"
	RadarHttphttpProtocolsParamsDateRange2d         RadarHttphttpProtocolsParamsDateRange = "2d"
	RadarHttphttpProtocolsParamsDateRange7d         RadarHttphttpProtocolsParamsDateRange = "7d"
	RadarHttphttpProtocolsParamsDateRange14d        RadarHttphttpProtocolsParamsDateRange = "14d"
	RadarHttphttpProtocolsParamsDateRange28d        RadarHttphttpProtocolsParamsDateRange = "28d"
	RadarHttphttpProtocolsParamsDateRange12w        RadarHttphttpProtocolsParamsDateRange = "12w"
	RadarHttphttpProtocolsParamsDateRange24w        RadarHttphttpProtocolsParamsDateRange = "24w"
	RadarHttphttpProtocolsParamsDateRange52w        RadarHttphttpProtocolsParamsDateRange = "52w"
	RadarHttphttpProtocolsParamsDateRange1dControl  RadarHttphttpProtocolsParamsDateRange = "1dControl"
	RadarHttphttpProtocolsParamsDateRange2dControl  RadarHttphttpProtocolsParamsDateRange = "2dControl"
	RadarHttphttpProtocolsParamsDateRange7dControl  RadarHttphttpProtocolsParamsDateRange = "7dControl"
	RadarHttphttpProtocolsParamsDateRange14dControl RadarHttphttpProtocolsParamsDateRange = "14dControl"
	RadarHttphttpProtocolsParamsDateRange28dControl RadarHttphttpProtocolsParamsDateRange = "28dControl"
	RadarHttphttpProtocolsParamsDateRange12wControl RadarHttphttpProtocolsParamsDateRange = "12wControl"
	RadarHttphttpProtocolsParamsDateRange24wControl RadarHttphttpProtocolsParamsDateRange = "24wControl"
)

type RadarHttphttpProtocolsParamsDeviceType string

const (
	RadarHttphttpProtocolsParamsDeviceTypeDesktop RadarHttphttpProtocolsParamsDeviceType = "DESKTOP"
	RadarHttphttpProtocolsParamsDeviceTypeMobile  RadarHttphttpProtocolsParamsDeviceType = "MOBILE"
	RadarHttphttpProtocolsParamsDeviceTypeOther   RadarHttphttpProtocolsParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHttphttpProtocolsParamsFormat string

const (
	RadarHttphttpProtocolsParamsFormatJson RadarHttphttpProtocolsParamsFormat = "JSON"
	RadarHttphttpProtocolsParamsFormatCsv  RadarHttphttpProtocolsParamsFormat = "CSV"
)

type RadarHttphttpProtocolsParamsHTTPVersion string

const (
	RadarHttphttpProtocolsParamsHTTPVersionHttPv1 RadarHttphttpProtocolsParamsHTTPVersion = "HTTPv1"
	RadarHttphttpProtocolsParamsHTTPVersionHttPv2 RadarHttphttpProtocolsParamsHTTPVersion = "HTTPv2"
	RadarHttphttpProtocolsParamsHTTPVersionHttPv3 RadarHttphttpProtocolsParamsHTTPVersion = "HTTPv3"
)

type RadarHttphttpProtocolsParamsIPVersion string

const (
	RadarHttphttpProtocolsParamsIPVersionIPv4 RadarHttphttpProtocolsParamsIPVersion = "IPv4"
	RadarHttphttpProtocolsParamsIPVersionIPv6 RadarHttphttpProtocolsParamsIPVersion = "IPv6"
)

type RadarHttphttpProtocolsParamsO string

const (
	RadarHttphttpProtocolsParamsOWindows  RadarHttphttpProtocolsParamsO = "WINDOWS"
	RadarHttphttpProtocolsParamsOMacosx   RadarHttphttpProtocolsParamsO = "MACOSX"
	RadarHttphttpProtocolsParamsOIos      RadarHttphttpProtocolsParamsO = "IOS"
	RadarHttphttpProtocolsParamsOAndroid  RadarHttphttpProtocolsParamsO = "ANDROID"
	RadarHttphttpProtocolsParamsOChromeos RadarHttphttpProtocolsParamsO = "CHROMEOS"
	RadarHttphttpProtocolsParamsOLinux    RadarHttphttpProtocolsParamsO = "LINUX"
	RadarHttphttpProtocolsParamsOSmartTv  RadarHttphttpProtocolsParamsO = "SMART_TV"
)

type RadarHttphttpProtocolsParamsTLSVersion string

const (
	RadarHttphttpProtocolsParamsTLSVersionTlSv1_0  RadarHttphttpProtocolsParamsTLSVersion = "TLSv1_0"
	RadarHttphttpProtocolsParamsTLSVersionTlSv1_1  RadarHttphttpProtocolsParamsTLSVersion = "TLSv1_1"
	RadarHttphttpProtocolsParamsTLSVersionTlSv1_2  RadarHttphttpProtocolsParamsTLSVersion = "TLSv1_2"
	RadarHttphttpProtocolsParamsTLSVersionTlSv1_3  RadarHttphttpProtocolsParamsTLSVersion = "TLSv1_3"
	RadarHttphttpProtocolsParamsTLSVersionTlSvQuic RadarHttphttpProtocolsParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPHTTPVersionsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHttphttpVersionsParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHttphttpVersionsParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHttphttpVersionsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHttphttpVersionsParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHttphttpVersionsParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHttphttpVersionsParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHttphttpVersionsParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHttphttpVersionsParamsO] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHttphttpVersionsParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPHTTPVersionsParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPHTTPVersionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHttphttpVersionsParamsAggInterval string

const (
	RadarHttphttpVersionsParamsAggInterval15m RadarHttphttpVersionsParamsAggInterval = "15m"
	RadarHttphttpVersionsParamsAggInterval1h  RadarHttphttpVersionsParamsAggInterval = "1h"
	RadarHttphttpVersionsParamsAggInterval1d  RadarHttphttpVersionsParamsAggInterval = "1d"
	RadarHttphttpVersionsParamsAggInterval1w  RadarHttphttpVersionsParamsAggInterval = "1w"
)

type RadarHttphttpVersionsParamsBotClass string

const (
	RadarHttphttpVersionsParamsBotClassLikelyAutomated RadarHttphttpVersionsParamsBotClass = "LIKELY_AUTOMATED"
	RadarHttphttpVersionsParamsBotClassLikelyHuman     RadarHttphttpVersionsParamsBotClass = "LIKELY_HUMAN"
)

type RadarHttphttpVersionsParamsDateRange string

const (
	RadarHttphttpVersionsParamsDateRange1d         RadarHttphttpVersionsParamsDateRange = "1d"
	RadarHttphttpVersionsParamsDateRange2d         RadarHttphttpVersionsParamsDateRange = "2d"
	RadarHttphttpVersionsParamsDateRange7d         RadarHttphttpVersionsParamsDateRange = "7d"
	RadarHttphttpVersionsParamsDateRange14d        RadarHttphttpVersionsParamsDateRange = "14d"
	RadarHttphttpVersionsParamsDateRange28d        RadarHttphttpVersionsParamsDateRange = "28d"
	RadarHttphttpVersionsParamsDateRange12w        RadarHttphttpVersionsParamsDateRange = "12w"
	RadarHttphttpVersionsParamsDateRange24w        RadarHttphttpVersionsParamsDateRange = "24w"
	RadarHttphttpVersionsParamsDateRange52w        RadarHttphttpVersionsParamsDateRange = "52w"
	RadarHttphttpVersionsParamsDateRange1dControl  RadarHttphttpVersionsParamsDateRange = "1dControl"
	RadarHttphttpVersionsParamsDateRange2dControl  RadarHttphttpVersionsParamsDateRange = "2dControl"
	RadarHttphttpVersionsParamsDateRange7dControl  RadarHttphttpVersionsParamsDateRange = "7dControl"
	RadarHttphttpVersionsParamsDateRange14dControl RadarHttphttpVersionsParamsDateRange = "14dControl"
	RadarHttphttpVersionsParamsDateRange28dControl RadarHttphttpVersionsParamsDateRange = "28dControl"
	RadarHttphttpVersionsParamsDateRange12wControl RadarHttphttpVersionsParamsDateRange = "12wControl"
	RadarHttphttpVersionsParamsDateRange24wControl RadarHttphttpVersionsParamsDateRange = "24wControl"
)

type RadarHttphttpVersionsParamsDeviceType string

const (
	RadarHttphttpVersionsParamsDeviceTypeDesktop RadarHttphttpVersionsParamsDeviceType = "DESKTOP"
	RadarHttphttpVersionsParamsDeviceTypeMobile  RadarHttphttpVersionsParamsDeviceType = "MOBILE"
	RadarHttphttpVersionsParamsDeviceTypeOther   RadarHttphttpVersionsParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHttphttpVersionsParamsFormat string

const (
	RadarHttphttpVersionsParamsFormatJson RadarHttphttpVersionsParamsFormat = "JSON"
	RadarHttphttpVersionsParamsFormatCsv  RadarHttphttpVersionsParamsFormat = "CSV"
)

type RadarHttphttpVersionsParamsHTTPProtocol string

const (
	RadarHttphttpVersionsParamsHTTPProtocolHTTP  RadarHttphttpVersionsParamsHTTPProtocol = "HTTP"
	RadarHttphttpVersionsParamsHTTPProtocolHTTPs RadarHttphttpVersionsParamsHTTPProtocol = "HTTPS"
)

type RadarHttphttpVersionsParamsIPVersion string

const (
	RadarHttphttpVersionsParamsIPVersionIPv4 RadarHttphttpVersionsParamsIPVersion = "IPv4"
	RadarHttphttpVersionsParamsIPVersionIPv6 RadarHttphttpVersionsParamsIPVersion = "IPv6"
)

type RadarHttphttpVersionsParamsO string

const (
	RadarHttphttpVersionsParamsOWindows  RadarHttphttpVersionsParamsO = "WINDOWS"
	RadarHttphttpVersionsParamsOMacosx   RadarHttphttpVersionsParamsO = "MACOSX"
	RadarHttphttpVersionsParamsOIos      RadarHttphttpVersionsParamsO = "IOS"
	RadarHttphttpVersionsParamsOAndroid  RadarHttphttpVersionsParamsO = "ANDROID"
	RadarHttphttpVersionsParamsOChromeos RadarHttphttpVersionsParamsO = "CHROMEOS"
	RadarHttphttpVersionsParamsOLinux    RadarHttphttpVersionsParamsO = "LINUX"
	RadarHttphttpVersionsParamsOSmartTv  RadarHttphttpVersionsParamsO = "SMART_TV"
)

type RadarHttphttpVersionsParamsTLSVersion string

const (
	RadarHttphttpVersionsParamsTLSVersionTlSv1_0  RadarHttphttpVersionsParamsTLSVersion = "TLSv1_0"
	RadarHttphttpVersionsParamsTLSVersionTlSv1_1  RadarHttphttpVersionsParamsTLSVersion = "TLSv1_1"
	RadarHttphttpVersionsParamsTLSVersionTlSv1_2  RadarHttphttpVersionsParamsTLSVersion = "TLSv1_2"
	RadarHttphttpVersionsParamsTLSVersionTlSv1_3  RadarHttphttpVersionsParamsTLSVersion = "TLSv1_3"
	RadarHttphttpVersionsParamsTLSVersionTlSvQuic RadarHttphttpVersionsParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPIPVersionsParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHttpipVersionsParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHttpipVersionsParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHttpipVersionsParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHttpipVersionsParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHttpipVersionsParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHttpipVersionsParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHttpipVersionsParamsHTTPVersion] `query:"httpVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for os name.
	Os param.Field[[]RadarHttpipVersionsParamsO] `query:"os"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHttpipVersionsParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPIPVersionsParams]'s query parameters as
// `url.Values`.
func (r RadarHTTPIPVersionsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHttpipVersionsParamsAggInterval string

const (
	RadarHttpipVersionsParamsAggInterval15m RadarHttpipVersionsParamsAggInterval = "15m"
	RadarHttpipVersionsParamsAggInterval1h  RadarHttpipVersionsParamsAggInterval = "1h"
	RadarHttpipVersionsParamsAggInterval1d  RadarHttpipVersionsParamsAggInterval = "1d"
	RadarHttpipVersionsParamsAggInterval1w  RadarHttpipVersionsParamsAggInterval = "1w"
)

type RadarHttpipVersionsParamsBotClass string

const (
	RadarHttpipVersionsParamsBotClassLikelyAutomated RadarHttpipVersionsParamsBotClass = "LIKELY_AUTOMATED"
	RadarHttpipVersionsParamsBotClassLikelyHuman     RadarHttpipVersionsParamsBotClass = "LIKELY_HUMAN"
)

type RadarHttpipVersionsParamsDateRange string

const (
	RadarHttpipVersionsParamsDateRange1d         RadarHttpipVersionsParamsDateRange = "1d"
	RadarHttpipVersionsParamsDateRange2d         RadarHttpipVersionsParamsDateRange = "2d"
	RadarHttpipVersionsParamsDateRange7d         RadarHttpipVersionsParamsDateRange = "7d"
	RadarHttpipVersionsParamsDateRange14d        RadarHttpipVersionsParamsDateRange = "14d"
	RadarHttpipVersionsParamsDateRange28d        RadarHttpipVersionsParamsDateRange = "28d"
	RadarHttpipVersionsParamsDateRange12w        RadarHttpipVersionsParamsDateRange = "12w"
	RadarHttpipVersionsParamsDateRange24w        RadarHttpipVersionsParamsDateRange = "24w"
	RadarHttpipVersionsParamsDateRange52w        RadarHttpipVersionsParamsDateRange = "52w"
	RadarHttpipVersionsParamsDateRange1dControl  RadarHttpipVersionsParamsDateRange = "1dControl"
	RadarHttpipVersionsParamsDateRange2dControl  RadarHttpipVersionsParamsDateRange = "2dControl"
	RadarHttpipVersionsParamsDateRange7dControl  RadarHttpipVersionsParamsDateRange = "7dControl"
	RadarHttpipVersionsParamsDateRange14dControl RadarHttpipVersionsParamsDateRange = "14dControl"
	RadarHttpipVersionsParamsDateRange28dControl RadarHttpipVersionsParamsDateRange = "28dControl"
	RadarHttpipVersionsParamsDateRange12wControl RadarHttpipVersionsParamsDateRange = "12wControl"
	RadarHttpipVersionsParamsDateRange24wControl RadarHttpipVersionsParamsDateRange = "24wControl"
)

type RadarHttpipVersionsParamsDeviceType string

const (
	RadarHttpipVersionsParamsDeviceTypeDesktop RadarHttpipVersionsParamsDeviceType = "DESKTOP"
	RadarHttpipVersionsParamsDeviceTypeMobile  RadarHttpipVersionsParamsDeviceType = "MOBILE"
	RadarHttpipVersionsParamsDeviceTypeOther   RadarHttpipVersionsParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHttpipVersionsParamsFormat string

const (
	RadarHttpipVersionsParamsFormatJson RadarHttpipVersionsParamsFormat = "JSON"
	RadarHttpipVersionsParamsFormatCsv  RadarHttpipVersionsParamsFormat = "CSV"
)

type RadarHttpipVersionsParamsHTTPProtocol string

const (
	RadarHttpipVersionsParamsHTTPProtocolHTTP  RadarHttpipVersionsParamsHTTPProtocol = "HTTP"
	RadarHttpipVersionsParamsHTTPProtocolHTTPs RadarHttpipVersionsParamsHTTPProtocol = "HTTPS"
)

type RadarHttpipVersionsParamsHTTPVersion string

const (
	RadarHttpipVersionsParamsHTTPVersionHttPv1 RadarHttpipVersionsParamsHTTPVersion = "HTTPv1"
	RadarHttpipVersionsParamsHTTPVersionHttPv2 RadarHttpipVersionsParamsHTTPVersion = "HTTPv2"
	RadarHttpipVersionsParamsHTTPVersionHttPv3 RadarHttpipVersionsParamsHTTPVersion = "HTTPv3"
)

type RadarHttpipVersionsParamsO string

const (
	RadarHttpipVersionsParamsOWindows  RadarHttpipVersionsParamsO = "WINDOWS"
	RadarHttpipVersionsParamsOMacosx   RadarHttpipVersionsParamsO = "MACOSX"
	RadarHttpipVersionsParamsOIos      RadarHttpipVersionsParamsO = "IOS"
	RadarHttpipVersionsParamsOAndroid  RadarHttpipVersionsParamsO = "ANDROID"
	RadarHttpipVersionsParamsOChromeos RadarHttpipVersionsParamsO = "CHROMEOS"
	RadarHttpipVersionsParamsOLinux    RadarHttpipVersionsParamsO = "LINUX"
	RadarHttpipVersionsParamsOSmartTv  RadarHttpipVersionsParamsO = "SMART_TV"
)

type RadarHttpipVersionsParamsTLSVersion string

const (
	RadarHttpipVersionsParamsTLSVersionTlSv1_0  RadarHttpipVersionsParamsTLSVersion = "TLSv1_0"
	RadarHttpipVersionsParamsTLSVersionTlSv1_1  RadarHttpipVersionsParamsTLSVersion = "TLSv1_1"
	RadarHttpipVersionsParamsTLSVersionTlSv1_2  RadarHttpipVersionsParamsTLSVersion = "TLSv1_2"
	RadarHttpipVersionsParamsTLSVersionTlSv1_3  RadarHttpipVersionsParamsTLSVersion = "TLSv1_3"
	RadarHttpipVersionsParamsTLSVersionTlSvQuic RadarHttpipVersionsParamsTLSVersion = "TLSvQUIC"
)

type RadarHTTPOssParams struct {
	// Aggregation interval results should be returned in (for example, in 15 minutes
	// or 1 hour intervals). Refer to
	// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
	AggInterval param.Field[RadarHTTPOssParamsAggInterval] `query:"aggInterval"`
	// Array of comma separated list of ASNs, start with `-` to exclude from results.
	// For example, `-174, 3356` excludes results from AS174, but includes results from
	// AS3356.
	Asn param.Field[[]string] `query:"asn"`
	// Filter for bot class. Refer to
	// [Bot classes](https://developers.cloudflare.com/radar/concepts/bot-classes/).
	BotClass param.Field[[]RadarHTTPOssParamsBotClass] `query:"botClass"`
	// End of the date range (inclusive).
	DateEnd param.Field[[]time.Time] `query:"dateEnd" format:"date-time"`
	// For example, use `7d` and `7dControl` to compare this week with the previous
	// week. Use this parameter or set specific start and end dates (`dateStart` and
	// `dateEnd` parameters).
	DateRange param.Field[[]RadarHTTPOssParamsDateRange] `query:"dateRange"`
	// Array of datetimes to filter the start of a series.
	DateStart param.Field[[]time.Time] `query:"dateStart" format:"date-time"`
	// Filter for device type.
	DeviceType param.Field[[]RadarHTTPOssParamsDeviceType] `query:"deviceType"`
	// Format results are returned in.
	Format param.Field[RadarHTTPOssParamsFormat] `query:"format"`
	// Filter for http protocol.
	HTTPProtocol param.Field[[]RadarHTTPOssParamsHTTPProtocol] `query:"httpProtocol"`
	// Filter for http version.
	HTTPVersion param.Field[[]RadarHTTPOssParamsHTTPVersion] `query:"httpVersion"`
	// Filter for ip version.
	IPVersion param.Field[[]RadarHTTPOssParamsIPVersion] `query:"ipVersion"`
	// Array of comma separated list of locations (alpha-2 country codes). Start with
	// `-` to exclude from results. For example, `-US,PT` excludes results from the US,
	// but includes results from PT.
	Location param.Field[[]string] `query:"location"`
	// Array of names that will be used to name the series in responses.
	Name param.Field[[]string] `query:"name"`
	// Filter for tls version.
	TLSVersion param.Field[[]RadarHTTPOssParamsTLSVersion] `query:"tlsVersion"`
}

// URLQuery serializes [RadarHTTPOssParams]'s query parameters as `url.Values`.
func (r RadarHTTPOssParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Aggregation interval results should be returned in (for example, in 15 minutes
// or 1 hour intervals). Refer to
// [Aggregation intervals](https://developers.cloudflare.com/radar/concepts/aggregation-intervals/).
type RadarHTTPOssParamsAggInterval string

const (
	RadarHTTPOssParamsAggInterval15m RadarHTTPOssParamsAggInterval = "15m"
	RadarHTTPOssParamsAggInterval1h  RadarHTTPOssParamsAggInterval = "1h"
	RadarHTTPOssParamsAggInterval1d  RadarHTTPOssParamsAggInterval = "1d"
	RadarHTTPOssParamsAggInterval1w  RadarHTTPOssParamsAggInterval = "1w"
)

type RadarHTTPOssParamsBotClass string

const (
	RadarHTTPOssParamsBotClassLikelyAutomated RadarHTTPOssParamsBotClass = "LIKELY_AUTOMATED"
	RadarHTTPOssParamsBotClassLikelyHuman     RadarHTTPOssParamsBotClass = "LIKELY_HUMAN"
)

type RadarHTTPOssParamsDateRange string

const (
	RadarHTTPOssParamsDateRange1d         RadarHTTPOssParamsDateRange = "1d"
	RadarHTTPOssParamsDateRange2d         RadarHTTPOssParamsDateRange = "2d"
	RadarHTTPOssParamsDateRange7d         RadarHTTPOssParamsDateRange = "7d"
	RadarHTTPOssParamsDateRange14d        RadarHTTPOssParamsDateRange = "14d"
	RadarHTTPOssParamsDateRange28d        RadarHTTPOssParamsDateRange = "28d"
	RadarHTTPOssParamsDateRange12w        RadarHTTPOssParamsDateRange = "12w"
	RadarHTTPOssParamsDateRange24w        RadarHTTPOssParamsDateRange = "24w"
	RadarHTTPOssParamsDateRange52w        RadarHTTPOssParamsDateRange = "52w"
	RadarHTTPOssParamsDateRange1dControl  RadarHTTPOssParamsDateRange = "1dControl"
	RadarHTTPOssParamsDateRange2dControl  RadarHTTPOssParamsDateRange = "2dControl"
	RadarHTTPOssParamsDateRange7dControl  RadarHTTPOssParamsDateRange = "7dControl"
	RadarHTTPOssParamsDateRange14dControl RadarHTTPOssParamsDateRange = "14dControl"
	RadarHTTPOssParamsDateRange28dControl RadarHTTPOssParamsDateRange = "28dControl"
	RadarHTTPOssParamsDateRange12wControl RadarHTTPOssParamsDateRange = "12wControl"
	RadarHTTPOssParamsDateRange24wControl RadarHTTPOssParamsDateRange = "24wControl"
)

type RadarHTTPOssParamsDeviceType string

const (
	RadarHTTPOssParamsDeviceTypeDesktop RadarHTTPOssParamsDeviceType = "DESKTOP"
	RadarHTTPOssParamsDeviceTypeMobile  RadarHTTPOssParamsDeviceType = "MOBILE"
	RadarHTTPOssParamsDeviceTypeOther   RadarHTTPOssParamsDeviceType = "OTHER"
)

// Format results are returned in.
type RadarHTTPOssParamsFormat string

const (
	RadarHTTPOssParamsFormatJson RadarHTTPOssParamsFormat = "JSON"
	RadarHTTPOssParamsFormatCsv  RadarHTTPOssParamsFormat = "CSV"
)

type RadarHTTPOssParamsHTTPProtocol string

const (
	RadarHTTPOssParamsHTTPProtocolHTTP  RadarHTTPOssParamsHTTPProtocol = "HTTP"
	RadarHTTPOssParamsHTTPProtocolHTTPs RadarHTTPOssParamsHTTPProtocol = "HTTPS"
)

type RadarHTTPOssParamsHTTPVersion string

const (
	RadarHTTPOssParamsHTTPVersionHttPv1 RadarHTTPOssParamsHTTPVersion = "HTTPv1"
	RadarHTTPOssParamsHTTPVersionHttPv2 RadarHTTPOssParamsHTTPVersion = "HTTPv2"
	RadarHTTPOssParamsHTTPVersionHttPv3 RadarHTTPOssParamsHTTPVersion = "HTTPv3"
)

type RadarHTTPOssParamsIPVersion string

const (
	RadarHTTPOssParamsIPVersionIPv4 RadarHTTPOssParamsIPVersion = "IPv4"
	RadarHTTPOssParamsIPVersionIPv6 RadarHTTPOssParamsIPVersion = "IPv6"
)

type RadarHTTPOssParamsTLSVersion string

const (
	RadarHTTPOssParamsTLSVersionTlSv1_0  RadarHTTPOssParamsTLSVersion = "TLSv1_0"
	RadarHTTPOssParamsTLSVersionTlSv1_1  RadarHTTPOssParamsTLSVersion = "TLSv1_1"
	RadarHTTPOssParamsTLSVersionTlSv1_2  RadarHTTPOssParamsTLSVersion = "TLSv1_2"
	RadarHTTPOssParamsTLSVersionTlSv1_3  RadarHTTPOssParamsTLSVersion = "TLSv1_3"
	RadarHTTPOssParamsTLSVersionTlSvQuic RadarHTTPOssParamsTLSVersion = "TLSvQUIC"
)
