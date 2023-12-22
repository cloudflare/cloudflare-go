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
func (r *ZoneAnalyticsDashboardService) ZoneAnalyticsDeprecatedGetDashboard(ctx context.Context, zoneIdentifier string, query ZoneAnalyticsDashboardZoneAnalyticsDeprecatedGetDashboardParams, opts ...option.RequestOption) (res *DashboardResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/analytics/dashboard", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type DashboardResponse struct {
	Errors   []DashboardResponseError   `json:"errors"`
	Messages []DashboardResponseMessage `json:"messages"`
	// The exact parameters/timestamps the analytics service used to return data.
	Query DashboardResponseQuery `json:"query"`
	// Totals and timeseries data.
	Result DashboardResponseResult `json:"result"`
	// Whether the API call was successful
	Success DashboardResponseSuccess `json:"success"`
	JSON    dashboardResponseJSON    `json:"-"`
}

// dashboardResponseJSON contains the JSON metadata for the struct
// [DashboardResponse]
type dashboardResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Query       apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DashboardResponseError struct {
	Code    int64                      `json:"code,required"`
	Message string                     `json:"message,required"`
	JSON    dashboardResponseErrorJSON `json:"-"`
}

// dashboardResponseErrorJSON contains the JSON metadata for the struct
// [DashboardResponseError]
type dashboardResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DashboardResponseMessage struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    dashboardResponseMessageJSON `json:"-"`
}

// dashboardResponseMessageJSON contains the JSON metadata for the struct
// [DashboardResponseMessage]
type dashboardResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The exact parameters/timestamps the analytics service used to return data.
type DashboardResponseQuery struct {
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
	Since DashboardResponseQuerySince `json:"since"`
	// The amount of time (in minutes) that each data point in the timeseries
	// represents. The granularity of the time-series returned (e.g. each bucket in the
	// time series representing 1-minute vs 1-day) is calculated by the API based on
	// the time-range provided to the API.
	TimeDelta int64 `json:"time_delta"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until DashboardResponseQueryUntil `json:"until"`
	JSON  dashboardResponseQueryJSON  `json:"-"`
}

// dashboardResponseQueryJSON contains the JSON metadata for the struct
// [DashboardResponseQuery]
type dashboardResponseQueryJSON struct {
	Since       apijson.Field
	TimeDelta   apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseQuery) UnmarshalJSON(data []byte) (err error) {
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
type DashboardResponseQuerySince interface {
	ImplementsDashboardResponseQuerySince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DashboardResponseQuerySince)(nil)).Elem(),
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
type DashboardResponseQueryUntil interface {
	ImplementsDashboardResponseQueryUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DashboardResponseQueryUntil)(nil)).Elem(),
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
type DashboardResponseResult struct {
	// Time deltas containing metadata about each bucket of time. The number of buckets
	// (resolution) is determined by the amount of time between the since and until
	// parameters.
	Timeseries []DashboardResponseResultTimesery `json:"timeseries"`
	// Breakdown of totals by data type.
	Totals DashboardResponseResultTotals `json:"totals"`
	JSON   dashboardResponseResultJSON   `json:"-"`
}

// dashboardResponseResultJSON contains the JSON metadata for the struct
// [DashboardResponseResult]
type dashboardResponseResultJSON struct {
	Timeseries  apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DashboardResponseResultTimesery struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth DashboardResponseResultTimeseriesBandwidth `json:"bandwidth"`
	// Breakdown of totals for pageviews.
	Pageviews DashboardResponseResultTimeseriesPageviews `json:"pageviews"`
	// Breakdown of totals for requests.
	Requests DashboardResponseResultTimeseriesRequests `json:"requests"`
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
	Since DashboardResponseResultTimeseriesSince `json:"since"`
	// Breakdown of totals for threats.
	Threats DashboardResponseResultTimeseriesThreats `json:"threats"`
	Uniques DashboardResponseResultTimeseriesUniques `json:"uniques"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until DashboardResponseResultTimeseriesUntil `json:"until"`
	JSON  dashboardResponseResultTimeseryJSON    `json:"-"`
}

// dashboardResponseResultTimeseryJSON contains the JSON metadata for the struct
// [DashboardResponseResultTimesery]
type dashboardResponseResultTimeseryJSON struct {
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

func (r *DashboardResponseResultTimesery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for bandwidth in the form of bytes.
type DashboardResponseResultTimeseriesBandwidth struct {
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
	Ssl DashboardResponseResultTimeseriesBandwidthSsl `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SslProtocols DashboardResponseResultTimeseriesBandwidthSslProtocols `json:"ssl_protocols"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                                          `json:"uncached"`
	JSON     dashboardResponseResultTimeseriesBandwidthJSON `json:"-"`
}

// dashboardResponseResultTimeseriesBandwidthJSON contains the JSON metadata for
// the struct [DashboardResponseResultTimeseriesBandwidth]
type dashboardResponseResultTimeseriesBandwidthJSON struct {
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

func (r *DashboardResponseResultTimeseriesBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A break down of bytes served over HTTPS.
type DashboardResponseResultTimeseriesBandwidthSsl struct {
	// The number of bytes served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of bytes served over HTTP.
	Unencrypted int64                                             `json:"unencrypted"`
	JSON        dashboardResponseResultTimeseriesBandwidthSslJSON `json:"-"`
}

// dashboardResponseResultTimeseriesBandwidthSslJSON contains the JSON metadata for
// the struct [DashboardResponseResultTimeseriesBandwidthSsl]
type dashboardResponseResultTimeseriesBandwidthSslJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResultTimeseriesBandwidthSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A breakdown of requests by their SSL protocol.
type DashboardResponseResultTimeseriesBandwidthSslProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                                      `json:"TLSv1.3"`
	JSON    dashboardResponseResultTimeseriesBandwidthSslProtocolsJSON `json:"-"`
}

// dashboardResponseResultTimeseriesBandwidthSslProtocolsJSON contains the JSON
// metadata for the struct [DashboardResponseResultTimeseriesBandwidthSslProtocols]
type dashboardResponseResultTimeseriesBandwidthSslProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResultTimeseriesBandwidthSslProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for pageviews.
type DashboardResponseResultTimeseriesPageviews struct {
	// The total number of pageviews served within the time range.
	All int64 `json:"all"`
	// A variable list of key/value pairs representing the search engine and number of
	// hits.
	SearchEngine interface{}                                    `json:"search_engine"`
	JSON         dashboardResponseResultTimeseriesPageviewsJSON `json:"-"`
}

// dashboardResponseResultTimeseriesPageviewsJSON contains the JSON metadata for
// the struct [DashboardResponseResultTimeseriesPageviews]
type dashboardResponseResultTimeseriesPageviewsJSON struct {
	All          apijson.Field
	SearchEngine apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DashboardResponseResultTimeseriesPageviews) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for requests.
type DashboardResponseResultTimeseriesRequests struct {
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
	Ssl DashboardResponseResultTimeseriesRequestsSsl `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SslProtocols DashboardResponseResultTimeseriesRequestsSslProtocols `json:"ssl_protocols"`
	// Total number of requests served from the origin.
	Uncached int64                                         `json:"uncached"`
	JSON     dashboardResponseResultTimeseriesRequestsJSON `json:"-"`
}

// dashboardResponseResultTimeseriesRequestsJSON contains the JSON metadata for the
// struct [DashboardResponseResultTimeseriesRequests]
type dashboardResponseResultTimeseriesRequestsJSON struct {
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

func (r *DashboardResponseResultTimeseriesRequests) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A break down of requests served over HTTPS.
type DashboardResponseResultTimeseriesRequestsSsl struct {
	// The number of requests served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of requests served over HTTP.
	Unencrypted int64                                            `json:"unencrypted"`
	JSON        dashboardResponseResultTimeseriesRequestsSslJSON `json:"-"`
}

// dashboardResponseResultTimeseriesRequestsSslJSON contains the JSON metadata for
// the struct [DashboardResponseResultTimeseriesRequestsSsl]
type dashboardResponseResultTimeseriesRequestsSslJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResultTimeseriesRequestsSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A breakdown of requests by their SSL protocol.
type DashboardResponseResultTimeseriesRequestsSslProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                                     `json:"TLSv1.3"`
	JSON    dashboardResponseResultTimeseriesRequestsSslProtocolsJSON `json:"-"`
}

// dashboardResponseResultTimeseriesRequestsSslProtocolsJSON contains the JSON
// metadata for the struct [DashboardResponseResultTimeseriesRequestsSslProtocols]
type dashboardResponseResultTimeseriesRequestsSslProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResultTimeseriesRequestsSslProtocols) UnmarshalJSON(data []byte) (err error) {
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
type DashboardResponseResultTimeseriesSince interface {
	ImplementsDashboardResponseResultTimeseriesSince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DashboardResponseResultTimeseriesSince)(nil)).Elem(),
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
type DashboardResponseResultTimeseriesThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}                                  `json:"type"`
	JSON dashboardResponseResultTimeseriesThreatsJSON `json:"-"`
}

// dashboardResponseResultTimeseriesThreatsJSON contains the JSON metadata for the
// struct [DashboardResponseResultTimeseriesThreats]
type dashboardResponseResultTimeseriesThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResultTimeseriesThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DashboardResponseResultTimeseriesUniques struct {
	// Total number of unique IP addresses within the time range.
	All  int64                                        `json:"all"`
	JSON dashboardResponseResultTimeseriesUniquesJSON `json:"-"`
}

// dashboardResponseResultTimeseriesUniquesJSON contains the JSON metadata for the
// struct [DashboardResponseResultTimeseriesUniques]
type dashboardResponseResultTimeseriesUniquesJSON struct {
	All         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResultTimeseriesUniques) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type DashboardResponseResultTimeseriesUntil interface {
	ImplementsDashboardResponseResultTimeseriesUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DashboardResponseResultTimeseriesUntil)(nil)).Elem(),
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
type DashboardResponseResultTotals struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth DashboardResponseResultTotalsBandwidth `json:"bandwidth"`
	// Breakdown of totals for pageviews.
	Pageviews DashboardResponseResultTotalsPageviews `json:"pageviews"`
	// Breakdown of totals for requests.
	Requests DashboardResponseResultTotalsRequests `json:"requests"`
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
	Since DashboardResponseResultTotalsSince `json:"since"`
	// Breakdown of totals for threats.
	Threats DashboardResponseResultTotalsThreats `json:"threats"`
	Uniques DashboardResponseResultTotalsUniques `json:"uniques"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until DashboardResponseResultTotalsUntil `json:"until"`
	JSON  dashboardResponseResultTotalsJSON  `json:"-"`
}

// dashboardResponseResultTotalsJSON contains the JSON metadata for the struct
// [DashboardResponseResultTotals]
type dashboardResponseResultTotalsJSON struct {
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

func (r *DashboardResponseResultTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for bandwidth in the form of bytes.
type DashboardResponseResultTotalsBandwidth struct {
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
	Ssl DashboardResponseResultTotalsBandwidthSsl `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SslProtocols DashboardResponseResultTotalsBandwidthSslProtocols `json:"ssl_protocols"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                                      `json:"uncached"`
	JSON     dashboardResponseResultTotalsBandwidthJSON `json:"-"`
}

// dashboardResponseResultTotalsBandwidthJSON contains the JSON metadata for the
// struct [DashboardResponseResultTotalsBandwidth]
type dashboardResponseResultTotalsBandwidthJSON struct {
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

func (r *DashboardResponseResultTotalsBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A break down of bytes served over HTTPS.
type DashboardResponseResultTotalsBandwidthSsl struct {
	// The number of bytes served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of bytes served over HTTP.
	Unencrypted int64                                         `json:"unencrypted"`
	JSON        dashboardResponseResultTotalsBandwidthSslJSON `json:"-"`
}

// dashboardResponseResultTotalsBandwidthSslJSON contains the JSON metadata for the
// struct [DashboardResponseResultTotalsBandwidthSsl]
type dashboardResponseResultTotalsBandwidthSslJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResultTotalsBandwidthSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A breakdown of requests by their SSL protocol.
type DashboardResponseResultTotalsBandwidthSslProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                                  `json:"TLSv1.3"`
	JSON    dashboardResponseResultTotalsBandwidthSslProtocolsJSON `json:"-"`
}

// dashboardResponseResultTotalsBandwidthSslProtocolsJSON contains the JSON
// metadata for the struct [DashboardResponseResultTotalsBandwidthSslProtocols]
type dashboardResponseResultTotalsBandwidthSslProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResultTotalsBandwidthSslProtocols) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for pageviews.
type DashboardResponseResultTotalsPageviews struct {
	// The total number of pageviews served within the time range.
	All int64 `json:"all"`
	// A variable list of key/value pairs representing the search engine and number of
	// hits.
	SearchEngine interface{}                                `json:"search_engine"`
	JSON         dashboardResponseResultTotalsPageviewsJSON `json:"-"`
}

// dashboardResponseResultTotalsPageviewsJSON contains the JSON metadata for the
// struct [DashboardResponseResultTotalsPageviews]
type dashboardResponseResultTotalsPageviewsJSON struct {
	All          apijson.Field
	SearchEngine apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DashboardResponseResultTotalsPageviews) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for requests.
type DashboardResponseResultTotalsRequests struct {
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
	Ssl DashboardResponseResultTotalsRequestsSsl `json:"ssl"`
	// A breakdown of requests by their SSL protocol.
	SslProtocols DashboardResponseResultTotalsRequestsSslProtocols `json:"ssl_protocols"`
	// Total number of requests served from the origin.
	Uncached int64                                     `json:"uncached"`
	JSON     dashboardResponseResultTotalsRequestsJSON `json:"-"`
}

// dashboardResponseResultTotalsRequestsJSON contains the JSON metadata for the
// struct [DashboardResponseResultTotalsRequests]
type dashboardResponseResultTotalsRequestsJSON struct {
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

func (r *DashboardResponseResultTotalsRequests) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A break down of requests served over HTTPS.
type DashboardResponseResultTotalsRequestsSsl struct {
	// The number of requests served over HTTPS.
	Encrypted int64 `json:"encrypted"`
	// The number of requests served over HTTP.
	Unencrypted int64                                        `json:"unencrypted"`
	JSON        dashboardResponseResultTotalsRequestsSslJSON `json:"-"`
}

// dashboardResponseResultTotalsRequestsSslJSON contains the JSON metadata for the
// struct [DashboardResponseResultTotalsRequestsSsl]
type dashboardResponseResultTotalsRequestsSslJSON struct {
	Encrypted   apijson.Field
	Unencrypted apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResultTotalsRequestsSsl) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// A breakdown of requests by their SSL protocol.
type DashboardResponseResultTotalsRequestsSslProtocols struct {
	// The number of requests served over HTTP.
	None int64 `json:"none"`
	// The number of requests served over TLS v1.0.
	TlSv1 int64 `json:"TLSv1"`
	// The number of requests served over TLS v1.1.
	TlSv1_1 int64 `json:"TLSv1.1"`
	// The number of requests served over TLS v1.2.
	TlSv1_2 int64 `json:"TLSv1.2"`
	// The number of requests served over TLS v1.3.
	TlSv1_3 int64                                                 `json:"TLSv1.3"`
	JSON    dashboardResponseResultTotalsRequestsSslProtocolsJSON `json:"-"`
}

// dashboardResponseResultTotalsRequestsSslProtocolsJSON contains the JSON metadata
// for the struct [DashboardResponseResultTotalsRequestsSslProtocols]
type dashboardResponseResultTotalsRequestsSslProtocolsJSON struct {
	None        apijson.Field
	TlSv1       apijson.Field
	TlSv1_1     apijson.Field
	TlSv1_2     apijson.Field
	TlSv1_3     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResultTotalsRequestsSslProtocols) UnmarshalJSON(data []byte) (err error) {
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
type DashboardResponseResultTotalsSince interface {
	ImplementsDashboardResponseResultTotalsSince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DashboardResponseResultTotalsSince)(nil)).Elem(),
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
type DashboardResponseResultTotalsThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}                              `json:"type"`
	JSON dashboardResponseResultTotalsThreatsJSON `json:"-"`
}

// dashboardResponseResultTotalsThreatsJSON contains the JSON metadata for the
// struct [DashboardResponseResultTotalsThreats]
type dashboardResponseResultTotalsThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResultTotalsThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type DashboardResponseResultTotalsUniques struct {
	// Total number of unique IP addresses within the time range.
	All  int64                                    `json:"all"`
	JSON dashboardResponseResultTotalsUniquesJSON `json:"-"`
}

// dashboardResponseResultTotalsUniquesJSON contains the JSON metadata for the
// struct [DashboardResponseResultTotalsUniques]
type dashboardResponseResultTotalsUniquesJSON struct {
	All         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DashboardResponseResultTotalsUniques) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type DashboardResponseResultTotalsUntil interface {
	ImplementsDashboardResponseResultTotalsUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*DashboardResponseResultTotalsUntil)(nil)).Elem(),
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
type DashboardResponseSuccess bool

const (
	DashboardResponseSuccessTrue DashboardResponseSuccess = true
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
