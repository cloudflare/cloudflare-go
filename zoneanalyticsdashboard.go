// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// ZoneAnalyticsDashboardService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAnalyticsDashboardService]
// method instead.
type ZoneAnalyticsDashboardService struct {
	Options []option.RequestOption
}

// NewZoneAnalyticsDashboardService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAnalyticsDashboardService(opts ...option.RequestOption) (r *ZoneAnalyticsDashboardService) {
	r = &ZoneAnalyticsDashboardService{}
	r.Options = opts
	return
}

// The dashboard view provides both totals and timeseries data for the given zone
// and time period across the entire Cloudflare network.
func (r *ZoneAnalyticsDashboardService) ZoneAnalyticsDeprecatedGetDashboard(ctx context.Context, zoneIdentifier string, query ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParams, opts ...option.RequestOption) (res *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/analytics/dashboard", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponse struct {
	Errors   []ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseError   `json:"errors"`
	Messages []ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseMessage `json:"messages"`
	// The exact parameters/timestamps the analytics service used to return data.
	Query ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQuery `json:"query"`
	// Totals and timeseries data.
	Result ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResult `json:"result"`
	// Whether the API call was successful
	Success ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseSuccess `json:"success"`
	JSON    zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseJSON    `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseJSON contains
// the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponse]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Query       apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseError struct {
	Code    int64                                                                      `json:"code,required"`
	Message string                                                                     `json:"message,required"`
	JSON    zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseErrorJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseError]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseMessage struct {
	Code    int64                                                                        `json:"code,required"`
	Message string                                                                       `json:"message,required"`
	JSON    zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseMessageJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseMessage]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The exact parameters/timestamps the analytics service used to return data.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQuery struct {
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQuerySince `json:"since"`
	// The amount of time (in minutes) that each data point in the timeseries
	// represents. The granularity of the time-series returned (e.g. each bucket in the
	// time series representing 1-minute vs 1-day) is calculated by the API based on
	// the time-range provided to the API.
	TimeDelta int64 `json:"time_delta"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQueryUntil `json:"until"`
	JSON  zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQueryJSON  `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQueryJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQuery]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQueryJSON struct {
	Since       apijson.Field
	TimeDelta   apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQuery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (inclusive) beginning of the requested time frame. This value can be a
// negative integer representing the number of minutes in the past relative to time
// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
// At this point in time, it cannot exceed a time in the past greater than one
// year.
//
// Ranges that the Cloudflare web application provides will provide the following
// period length for each point:
//
// - Last 60 minutes (from -59 to -1): 1 minute resolution
// - Last 7 hours (from -419 to -60): 15 minutes resolution
// - Last 15 hours (from -899 to -420): 30 minutes resolution
// - Last 72 hours (from -4320 to -900): 1 hour resolution
// - Older than 3 days (-525600 to -4320): 1 day resolution.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQuerySince interface {
	ImplementsZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQuerySince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQuerySince)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQueryUntil interface {
	ImplementsZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQueryUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseQueryUntil)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Totals and timeseries data.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResult struct {
	// Time deltas containing metadata about each bucket of time. The number of buckets
	// (resolution) is determined by the amount of time between the since and until
	// parameters.
	Timeseries []ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimesery `json:"timeseries"`
	// Breakdown of totals by data type.
	Totals ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotals `json:"totals"`
	JSON   zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultJSON   `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResult]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultJSON struct {
	Timeseries  apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimesery struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidth `json:"bandwidth"`
	// Breakdown of totals for pageviews.
	Pageviews ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesPageviews `json:"pageviews"`
	// Breakdown of totals for requests.
	Requests ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequests `json:"requests"`
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesSince `json:"since"`
	// Breakdown of totals for threats.
	Threats ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesThreats `json:"threats"`
	Uniques ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesUniques `json:"uniques"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesUntil `json:"until"`
	JSON  zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseryJSON    `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseryJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimesery]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseryJSON struct {
	Bandwidth   apijson.Field
	Pageviews   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Uniques     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimesery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for bandwidth in the form of bytes.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidth struct {
	// The total number of bytes served within the time frame.
	All int64 `json:"all"`
	// The number of bytes that were cached (and served) by Cloudflare.
	Cached int64 `json:"cached"`
	// A variable list of key/value pairs where the key represents the type of content
	// served, and the value is the number in bytes served.
	ContentType interface{} `json:"content_type"`
	// A variable list of key/value pairs where the key is a two-digit country code and
	// the value is the number of bytes served to that country.
	Country interface{} `json:"country"`
	// A break down of bytes served over HTTPS.
	Ssl ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSsl `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SslProtocols ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSslProtocols `json:"ssl_protocols"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                                                                                          `json:"uncached"`
	JSON     zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidth]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthJSON struct {
	All          apijson.Field
	Cached       apijson.Field
	ContentType  apijson.Field
	Country      apijson.Field
	Ssl          apijson.Field
	SslProtocols apijson.Field
	Uncached     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A break down of bytes served over HTTPS.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSsl struct {
	// The number of bytes served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of bytes served over HTTP.
	Unencrypted int64                                                                                             `json:"unencrypted"`
	JSON        zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSslJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSslJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSsl]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSslJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A breakdown of requests by their SSL protocol.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSslProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                                                                                      `json:"TLSv1.3"`
	JSON    zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSslProtocolsJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSslProtocolsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSslProtocols]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSslProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesBandwidthSslProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for pageviews.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesPageviews struct {
	// The total number of pageviews served within the time range.
	All int64 `json:"all"`
	// A variable list of key/value pairs representing the search engine and number of
	// hits.
	SearchEngine interface{}                                                                                    `json:"search_engine"`
	JSON         zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesPageviewsJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesPageviewsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesPageviews]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesPageviewsJSON struct {
	All          apijson.Field
	SearchEngine apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesPageviews) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for requests.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequests struct {
	// Total number of requests served.
	All int64 `json:"all"`
	// Total number of cached requests served.
	Cached int64 `json:"cached"`
	// A variable list of key/value pairs where the key represents the type of content
	// served, and the value is the number of requests.
	ContentType interface{} `json:"content_type"`
	// A variable list of key/value pairs where the key is a two-digit country code and
	// the value is the number of requests served to that country.
	Country interface{} `json:"country"`
	// Key/value pairs where the key is a HTTP status code and the value is the number
	// of requests served with that code.
	HTTPStatus interface{} `json:"http_status"`
	// A break down of requests served over HTTPS.
	Ssl ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSsl `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SslProtocols ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSslProtocols `json:"ssl_protocols"`
	// Total number of requests served from the origin.
	Uncached int64                                                                                         `json:"uncached"`
	JSON     zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequests]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsJSON struct {
	All          apijson.Field
	Cached       apijson.Field
	ContentType  apijson.Field
	Country      apijson.Field
	HTTPStatus   apijson.Field
	Ssl          apijson.Field
	SslProtocols apijson.Field
	Uncached     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequests) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A break down of requests served over HTTPS.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSsl struct {
	// The number of requests served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of requests served over HTTP.
	Unencrypted int64                                                                                            `json:"unencrypted"`
	JSON        zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSslJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSslJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSsl]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSslJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A breakdown of requests by their SSL protocol.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSslProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                                                                                     `json:"TLSv1.3"`
	JSON    zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSslProtocolsJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSslProtocolsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSslProtocols]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSslProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesRequestsSslProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (inclusive) beginning of the requested time frame. This value can be a
// negative integer representing the number of minutes in the past relative to time
// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
// At this point in time, it cannot exceed a time in the past greater than one
// year.
//
// Ranges that the Cloudflare web application provides will provide the following
// period length for each point:
//
// - Last 60 minutes (from -59 to -1): 1 minute resolution
// - Last 7 hours (from -419 to -60): 15 minutes resolution
// - Last 15 hours (from -899 to -420): 30 minutes resolution
// - Last 72 hours (from -4320 to -900): 1 hour resolution
// - Older than 3 days (-525600 to -4320): 1 day resolution.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesSince interface {
	ImplementsZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesSince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesSince)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Breakdown of totals for threats.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}                                                                                  `json:"type"`
	JSON zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesThreatsJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesThreatsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesThreats]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesUniques struct {
	// Total number of unique IP addresses within the time range.
	All  int64                                                                                        `json:"all"`
	JSON zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesUniquesJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesUniquesJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesUniques]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesUniquesJSON struct {
	All         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesUniques) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesUntil interface {
	ImplementsZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTimeseriesUntil)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Breakdown of totals by data type.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotals struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidth `json:"bandwidth"`
	// Breakdown of totals for pageviews.
	Pageviews ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsPageviews `json:"pageviews"`
	// Breakdown of totals for requests.
	Requests ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequests `json:"requests"`
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsSince `json:"since"`
	// Breakdown of totals for threats.
	Threats ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsThreats `json:"threats"`
	Uniques ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsUniques `json:"uniques"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsUntil `json:"until"`
	JSON  zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsJSON  `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotals]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsJSON struct {
	Bandwidth   apijson.Field
	Pageviews   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Uniques     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for bandwidth in the form of bytes.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidth struct {
	// The total number of bytes served within the time frame.
	All int64 `json:"all"`
	// The number of bytes that were cached (and served) by Cloudflare.
	Cached int64 `json:"cached"`
	// A variable list of key/value pairs where the key represents the type of content
	// served, and the value is the number in bytes served.
	ContentType interface{} `json:"content_type"`
	// A variable list of key/value pairs where the key is a two-digit country code and
	// the value is the number of bytes served to that country.
	Country interface{} `json:"country"`
	// A break down of bytes served over HTTPS.
	Ssl ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSsl `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SslProtocols ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSslProtocols `json:"ssl_protocols"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                                                                                      `json:"uncached"`
	JSON     zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidth]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthJSON struct {
	All          apijson.Field
	Cached       apijson.Field
	ContentType  apijson.Field
	Country      apijson.Field
	Ssl          apijson.Field
	SslProtocols apijson.Field
	Uncached     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A break down of bytes served over HTTPS.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSsl struct {
	// The number of bytes served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of bytes served over HTTP.
	Unencrypted int64                                                                                         `json:"unencrypted"`
	JSON        zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSslJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSslJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSsl]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSslJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A breakdown of requests by their SSL protocol.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSslProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                                                                                  `json:"TLSv1.3"`
	JSON    zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSslProtocolsJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSslProtocolsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSslProtocols]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSslProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsBandwidthSslProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for pageviews.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsPageviews struct {
	// The total number of pageviews served within the time range.
	All int64 `json:"all"`
	// A variable list of key/value pairs representing the search engine and number of
	// hits.
	SearchEngine interface{}                                                                                `json:"search_engine"`
	JSON         zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsPageviewsJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsPageviewsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsPageviews]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsPageviewsJSON struct {
	All          apijson.Field
	SearchEngine apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsPageviews) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for requests.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequests struct {
	// Total number of requests served.
	All int64 `json:"all"`
	// Total number of cached requests served.
	Cached int64 `json:"cached"`
	// A variable list of key/value pairs where the key represents the type of content
	// served, and the value is the number of requests.
	ContentType interface{} `json:"content_type"`
	// A variable list of key/value pairs where the key is a two-digit country code and
	// the value is the number of requests served to that country.
	Country interface{} `json:"country"`
	// Key/value pairs where the key is a HTTP status code and the value is the number
	// of requests served with that code.
	HTTPStatus interface{} `json:"http_status"`
	// A break down of requests served over HTTPS.
	Ssl ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSsl `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SslProtocols ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSslProtocols `json:"ssl_protocols"`
	// Total number of requests served from the origin.
	Uncached int64                                                                                     `json:"uncached"`
	JSON     zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequests]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsJSON struct {
	All          apijson.Field
	Cached       apijson.Field
	ContentType  apijson.Field
	Country      apijson.Field
	HTTPStatus   apijson.Field
	Ssl          apijson.Field
	SslProtocols apijson.Field
	Uncached     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequests) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A break down of requests served over HTTPS.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSsl struct {
	// The number of requests served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of requests served over HTTP.
	Unencrypted int64                                                                                        `json:"unencrypted"`
	JSON        zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSslJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSslJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSsl]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSslJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A breakdown of requests by their SSL protocol.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSslProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                                                                                 `json:"TLSv1.3"`
	JSON    zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSslProtocolsJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSslProtocolsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSslProtocols]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSslProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsRequestsSslProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (inclusive) beginning of the requested time frame. This value can be a
// negative integer representing the number of minutes in the past relative to time
// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
// At this point in time, it cannot exceed a time in the past greater than one
// year.
//
// Ranges that the Cloudflare web application provides will provide the following
// period length for each point:
//
// - Last 60 minutes (from -59 to -1): 1 minute resolution
// - Last 7 hours (from -419 to -60): 15 minutes resolution
// - Last 15 hours (from -899 to -420): 30 minutes resolution
// - Last 72 hours (from -4320 to -900): 1 hour resolution
// - Older than 3 days (-525600 to -4320): 1 day resolution.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsSince interface {
	ImplementsZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsSince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsSince)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Breakdown of totals for threats.
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}                                                                              `json:"type"`
	JSON zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsThreatsJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsThreatsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsThreats]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsUniques struct {
	// Total number of unique IP addresses within the time range.
	All  int64                                                                                    `json:"all"`
	JSON zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsUniquesJSON `json:"-"`
}

// zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsUniquesJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsUniques]
type zoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsUniquesJSON struct {
	All         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsUniques) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsUntil interface {
	ImplementsZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseResultTotalsUntil)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter:         gjson.String,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.Number,
			DiscriminatorValue: "",
			Type:               reflect.TypeOf(shared.UnionInt(0)),
		},
	)
}

// Whether the API call was successful
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseSuccess bool

const (
	ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseSuccessTrue ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardResponseSuccess = true
)

type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParams struct {
	// When set to true, the API will move the requested time window backward, until it
	// finds a region with completely aggregated data.
	//
	// The API response _may not represent the requested time window_.
	Continuous param.Field[bool] `query:"continuous"`
	// The (inclusive) beginning of the requested time frame. This value can be a
	// negative integer representing the number of minutes in the past relative to time
	// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
	// At this point in time, it cannot exceed a time in the past greater than one
	// year.
	//
	// Ranges that the Cloudflare web application provides will provide the following
	// period length for each point:
	//
	// - Last 60 minutes (from -59 to -1): 1 minute resolution
	// - Last 7 hours (from -419 to -60): 15 minutes resolution
	// - Last 15 hours (from -899 to -420): 30 minutes resolution
	// - Last 72 hours (from -4320 to -900): 1 hour resolution
	// - Older than 3 days (-525600 to -4320): 1 day resolution.
	Since param.Field[ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsSince] `query:"since"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until param.Field[ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsUntil] `query:"until"`
}

// URLQuery serializes
// [ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParams]'s query
// parameters as `url.Values`.
func (r ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// The (inclusive) beginning of the requested time frame. This value can be a
// negative integer representing the number of minutes in the past relative to time
// the request is made, or can be an absolute timestamp that conforms to RFC 3339.
// At this point in time, it cannot exceed a time in the past greater than one
// year.
//
// Ranges that the Cloudflare web application provides will provide the following
// period length for each point:
//
// - Last 60 minutes (from -59 to -1): 1 minute resolution
// - Last 7 hours (from -419 to -60): 15 minutes resolution
// - Last 15 hours (from -899 to -420): 30 minutes resolution
// - Last 72 hours (from -4320 to -900): 1 hour resolution
// - Older than 3 days (-525600 to -4320): 1 day resolution.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsSince interface {
	ImplementsZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsSince()
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsUntil interface {
	ImplementsZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParamsUntil()
}
