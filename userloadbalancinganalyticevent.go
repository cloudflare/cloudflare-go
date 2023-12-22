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

// UserLoadBalancingAnalyticEventService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewUserLoadBalancingAnalyticEventService] method instead.
type UserLoadBalancingAnalyticEventService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancingAnalyticEventService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserLoadBalancingAnalyticEventService(opts ...option.RequestOption) (r *UserLoadBalancingAnalyticEventService) {
	r = &UserLoadBalancingAnalyticEventService{}
	r.Options = opts
	return
}

// List origin health changes.
func (r *UserLoadBalancingAnalyticEventService) LoadBalancerHealthcheckEventsListHealthcheckEvents(ctx context.Context, query UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsParams, opts ...option.RequestOption) (res *AnalyticsResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancing_analytics/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AnalyticsResponseCollection struct {
	Errors     []AnalyticsResponseCollectionError    `json:"errors"`
	Messages   []AnalyticsResponseCollectionMessage  `json:"messages"`
	Result     []AnalyticsResponseCollectionResult   `json:"result"`
	ResultInfo AnalyticsResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AnalyticsResponseCollectionSuccess `json:"success"`
	JSON    analyticsResponseCollectionJSON    `json:"-"`
}

// analyticsResponseCollectionJSON contains the JSON metadata for the struct
// [AnalyticsResponseCollection]
type analyticsResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsResponseCollectionError struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    analyticsResponseCollectionErrorJSON `json:"-"`
}

// analyticsResponseCollectionErrorJSON contains the JSON metadata for the struct
// [AnalyticsResponseCollectionError]
type analyticsResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsResponseCollectionMessage struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    analyticsResponseCollectionMessageJSON `json:"-"`
}

// analyticsResponseCollectionMessageJSON contains the JSON metadata for the struct
// [AnalyticsResponseCollectionMessage]
type analyticsResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsResponseCollectionResult struct {
	ID        int64                                 `json:"id"`
	Origins   []interface{}                         `json:"origins"`
	Pool      interface{}                           `json:"pool"`
	Timestamp time.Time                             `json:"timestamp" format:"date-time"`
	JSON      analyticsResponseCollectionResultJSON `json:"-"`
}

// analyticsResponseCollectionResultJSON contains the JSON metadata for the struct
// [AnalyticsResponseCollectionResult]
type analyticsResponseCollectionResultJSON struct {
	ID          apijson.Field
	Origins     apijson.Field
	Pool        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsResponseCollectionResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AnalyticsResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       analyticsResponseCollectionResultInfoJSON `json:"-"`
}

// analyticsResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [AnalyticsResponseCollectionResultInfo]
type analyticsResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AnalyticsResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AnalyticsResponseCollectionSuccess bool

const (
	AnalyticsResponseCollectionSuccessTrue AnalyticsResponseCollectionSuccess = true
)

type UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsParams struct {
	Identifier param.Field[interface{}] `query:"identifier"`
	// If true, filter events where the origin status is healthy. If false, filter
	// events where the origin status is unhealthy.
	OriginHealthy param.Field[bool] `query:"origin_healthy"`
	// The name for the origin to filter.
	OriginName param.Field[string] `query:"origin_name"`
	// If true, filter events where the pool status is healthy. If false, filter events
	// where the pool status is unhealthy.
	PoolHealthy param.Field[bool] `query:"pool_healthy"`
	// The name for the pool to filter.
	PoolName param.Field[string] `query:"pool_name"`
	// Start date and time of requesting data period in the ISO8601 format.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// End date and time of requesting data period in the ISO8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes
// [UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsParams]'s
// query parameters as `url.Values`.
func (r UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
