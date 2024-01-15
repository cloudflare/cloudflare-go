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

// ZoneAnalyticsColoService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneAnalyticsColoService] method
// instead.
type ZoneAnalyticsColoService struct {
	Options []option.RequestOption
}

// NewZoneAnalyticsColoService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneAnalyticsColoService(opts ...option.RequestOption) (r *ZoneAnalyticsColoService) {
	r = &ZoneAnalyticsColoService{}
	r.Options = opts
	return
}

// This view provides a breakdown of analytics data by datacenter. Note: This is
// available to Enterprise customers only.
func (r *ZoneAnalyticsColoService) ZoneAnalyticsDeprecatedGetAnalyticsByCoLocations(ctx context.Context, zoneIdentifier string, query ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParams, opts ...option.RequestOption) (res *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/analytics/colos", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponse struct {
	Errors   []ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseError   `json:"errors"`
	Messages []ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseMessage `json:"messages"`
	// The exact parameters/timestamps the analytics service used to return data.
	Query ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQuery `json:"query"`
	// A breakdown of all dashboard analytics data by co-locations. This is limited to
	// Enterprise zones only.
	Result []ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResult `json:"result"`
	// Whether the API call was successful
	Success ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseSuccess `json:"success"`
	JSON    zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseJSON    `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponse]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Query       apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseError struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseErrorJSON `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseError]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseMessage struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseMessageJSON `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseMessage]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The exact parameters/timestamps the analytics service used to return data.
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQuery struct {
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
	Since ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQuerySince `json:"since"`
	// The amount of time (in minutes) that each data point in the timeseries
	// represents. The granularity of the time-series returned (e.g. each bucket in the
	// time series representing 1-minute vs 1-day) is calculated by the API based on
	// the time-range provided to the API.
	TimeDelta int64 `json:"time_delta"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQueryUntil `json:"until"`
	JSON  zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQueryJSON  `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQueryJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQuery]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQueryJSON struct {
	Since       apijson.Field
	TimeDelta   apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQuery) UnmarshalJSON(data []byte) (err error) {
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
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQuerySince interface {
	ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQuerySince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQuerySince)(nil)).Elem(),
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
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQueryUntil interface {
	ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQueryUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseQueryUntil)(nil)).Elem(),
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

type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResult struct {
	// The airport code identifer for the co-location.
	ColoID string `json:"colo_id"`
	// Time deltas containing metadata about each bucket of time. The number of buckets
	// (resolution) is determined by the amount of time between the since and until
	// parameters.
	Timeseries []ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimesery `json:"timeseries"`
	// Breakdown of totals by data type.
	Totals ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotals `json:"totals"`
	JSON   zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultJSON   `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResult]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultJSON struct {
	ColoID      apijson.Field
	Timeseries  apijson.Field
	Totals      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimesery struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesBandwidth `json:"bandwidth"`
	// Breakdown of totals for requests.
	Requests ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesRequests `json:"requests"`
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
	Since ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesSince `json:"since"`
	// Breakdown of totals for threats.
	Threats ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesThreats `json:"threats"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesUntil `json:"until"`
	JSON  zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseryJSON    `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseryJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimesery]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseryJSON struct {
	Bandwidth   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimesery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for bandwidth in the form of bytes.
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesBandwidth struct {
	// The total number of bytes served within the time frame.
	All int64 `json:"all"`
	// The number of bytes that were cached (and served) by Cloudflare.
	Cached int64 `json:"cached"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                                                                                                  `json:"uncached"`
	JSON     zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesBandwidthJSON `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesBandwidthJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesBandwidth]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesBandwidthJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for requests.
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesRequests struct {
	// Total number of requests served.
	All int64 `json:"all"`
	// Total number of cached requests served.
	Cached int64 `json:"cached"`
	// Key/value pairs where the key is a two-digit country code and the value is the
	// number of requests served to that country.
	Country interface{} `json:"country"`
	// A variable list of key/value pairs where the key is a HTTP status code and the
	// value is the number of requests with that code served.
	HTTPStatus interface{} `json:"http_status"`
	// Total number of requests served from the origin.
	Uncached int64                                                                                                 `json:"uncached"`
	JSON     zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesRequestsJSON `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesRequestsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesRequests]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesRequestsJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Country     apijson.Field
	HTTPStatus  apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesRequests) UnmarshalJSON(data []byte) (err error) {
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
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesSince interface {
	ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesSince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesSince)(nil)).Elem(),
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
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}                                                                                          `json:"type"`
	JSON zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesThreatsJSON `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesThreatsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesThreats]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesUntil interface {
	ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTimeseriesUntil)(nil)).Elem(),
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
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotals struct {
	// Breakdown of totals for bandwidth in the form of bytes.
	Bandwidth ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsBandwidth `json:"bandwidth"`
	// Breakdown of totals for requests.
	Requests ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsRequests `json:"requests"`
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
	Since ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsSince `json:"since"`
	// Breakdown of totals for threats.
	Threats ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsThreats `json:"threats"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsUntil `json:"until"`
	JSON  zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsJSON  `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotals]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsJSON struct {
	Bandwidth   apijson.Field
	Requests    apijson.Field
	Since       apijson.Field
	Threats     apijson.Field
	Until       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotals) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for bandwidth in the form of bytes.
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsBandwidth struct {
	// The total number of bytes served within the time frame.
	All int64 `json:"all"`
	// The number of bytes that were cached (and served) by Cloudflare.
	Cached int64 `json:"cached"`
	// The number of bytes that were fetched and served from the origin server.
	Uncached int64                                                                                              `json:"uncached"`
	JSON     zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsBandwidthJSON `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsBandwidthJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsBandwidth]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsBandwidthJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsBandwidth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Breakdown of totals for requests.
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsRequests struct {
	// Total number of requests served.
	All int64 `json:"all"`
	// Total number of cached requests served.
	Cached int64 `json:"cached"`
	// Key/value pairs where the key is a two-digit country code and the value is the
	// number of requests served to that country.
	Country interface{} `json:"country"`
	// A variable list of key/value pairs where the key is a HTTP status code and the
	// value is the number of requests with that code served.
	HTTPStatus interface{} `json:"http_status"`
	// Total number of requests served from the origin.
	Uncached int64                                                                                             `json:"uncached"`
	JSON     zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsRequestsJSON `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsRequestsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsRequests]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsRequestsJSON struct {
	All         apijson.Field
	Cached      apijson.Field
	Country     apijson.Field
	HTTPStatus  apijson.Field
	Uncached    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsRequests) UnmarshalJSON(data []byte) (err error) {
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
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsSince interface {
	ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsSince()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsSince)(nil)).Elem(),
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
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsThreats struct {
	// The total number of identifiable threats received over the time frame.
	All int64 `json:"all"`
	// A list of key/value pairs where the key is a two-digit country code and the
	// value is the number of malicious requests received from that country.
	Country interface{} `json:"country"`
	// The list of key/value pairs where the key is a threat category and the value is
	// the number of requests.
	Type interface{}                                                                                      `json:"type"`
	JSON zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsThreatsJSON `json:"-"`
}

// zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsThreatsJSON
// contains the JSON metadata for the struct
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsThreats]
type zoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsThreatsJSON struct {
	All         apijson.Field
	Country     apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsThreats) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Union satisfied by [shared.UnionString] or [shared.UnionInt].
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsUntil interface {
	ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsUntil()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseResultTotalsUntil)(nil)).Elem(),
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
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseSuccess bool

const (
	ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseSuccessTrue ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsResponseSuccess = true
)

type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParams struct {
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
	Since param.Field[ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsSince] `query:"since"`
	// The (exclusive) end of the requested time frame. This value can be a negative
	// integer representing the number of minutes in the past relative to time the
	// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
	// omitted, the time of the request is used.
	Until param.Field[ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsUntil] `query:"until"`
}

// URLQuery serializes
// [ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParams]'s
// query parameters as `url.Values`.
func (r ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParams) URLQuery() (v url.Values) {
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
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsSince interface {
	ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsSince()
}

// The (exclusive) end of the requested time frame. This value can be a negative
// integer representing the number of minutes in the past relative to time the
// request is made, or can be an absolute timestamp that conforms to RFC 3339. If
// omitted, the time of the request is used.
//
// Satisfied by [shared.UnionString], [shared.UnionInt].
type ZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsUntil interface {
	ImplementsZoneAnalyticsColoZoneAnalyticsDeprecatedGetAnalyticsByCoLocationsParamsUntil()
}
