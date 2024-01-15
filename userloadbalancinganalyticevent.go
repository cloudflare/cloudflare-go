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
func (r *UserLoadBalancingAnalyticEventService) LoadBalancerHealthcheckEventsListHealthcheckEvents(ctx context.Context, query UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsParams, opts ...option.RequestOption) (res *UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "user/load_balancing_analytics/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponse struct {
	Errors     []UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseError    `json:"errors"`
	Messages   []UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseMessage  `json:"messages"`
	Result     []UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResult   `json:"result"`
	ResultInfo UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseSuccess `json:"success"`
	JSON    userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseJSON    `json:"-"`
}

// userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseJSON
// contains the JSON metadata for the struct
// [UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponse]
type userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseError struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseErrorJSON `json:"-"`
}

// userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseErrorJSON
// contains the JSON metadata for the struct
// [UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseError]
type userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseMessage struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseMessageJSON `json:"-"`
}

// userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseMessageJSON
// contains the JSON metadata for the struct
// [UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseMessage]
type userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResult struct {
	ID        int64                                                                                              `json:"id"`
	Origins   []interface{}                                                                                      `json:"origins"`
	Pool      interface{}                                                                                        `json:"pool"`
	Timestamp time.Time                                                                                          `json:"timestamp" format:"date-time"`
	JSON      userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResultJSON `json:"-"`
}

// userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResultJSON
// contains the JSON metadata for the struct
// [UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResult]
type userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResultJSON struct {
	ID          apijson.Field
	Origins     apijson.Field
	Pool        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                `json:"total_count"`
	JSON       userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResultInfoJSON `json:"-"`
}

// userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResultInfo]
type userLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseSuccess bool

const (
	UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseSuccessTrue UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsResponseSuccess = true
)

type UserLoadBalancingAnalyticEventLoadBalancerHealthcheckEventsListHealthcheckEventsParams struct {
	Identifier param.Field[string] `query:"identifier"`
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
