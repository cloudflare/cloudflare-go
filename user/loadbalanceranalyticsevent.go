// File generated from our OpenAPI spec by Stainless.

package user

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/option"
)

// LoadBalancerAnalyticsEventService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewLoadBalancerAnalyticsEventService] method instead.
type LoadBalancerAnalyticsEventService struct {
	Options []option.RequestOption
}

// NewLoadBalancerAnalyticsEventService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLoadBalancerAnalyticsEventService(opts ...option.RequestOption) (r *LoadBalancerAnalyticsEventService) {
	r = &LoadBalancerAnalyticsEventService{}
	r.Options = opts
	return
}

// List origin health changes.
func (r *LoadBalancerAnalyticsEventService) List(ctx context.Context, query LoadBalancerAnalyticsEventListParams, opts ...option.RequestOption) (res *[]LoadBalancingAnalytics, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerAnalyticsEventListResponseEnvelope
	path := "user/load_balancing_analytics/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancingAnalytics struct {
	ID        int64                      `json:"id"`
	Origins   []interface{}              `json:"origins"`
	Pool      interface{}                `json:"pool"`
	Timestamp time.Time                  `json:"timestamp" format:"date-time"`
	JSON      loadBalancingAnalyticsJSON `json:"-"`
}

// loadBalancingAnalyticsJSON contains the JSON metadata for the struct
// [LoadBalancingAnalytics]
type loadBalancingAnalyticsJSON struct {
	ID          apijson.Field
	Origins     apijson.Field
	Pool        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingAnalytics) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingAnalyticsJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerAnalyticsEventListParams struct {
	// If true, filter events where the origin status is healthy. If false, filter
	// events where the origin status is unhealthy.
	OriginHealthy param.Field[bool] `query:"origin_healthy"`
	// The name for the origin to filter.
	OriginName param.Field[string] `query:"origin_name"`
	// If true, filter events where the pool status is healthy. If false, filter events
	// where the pool status is unhealthy.
	PoolHealthy param.Field[bool]   `query:"pool_healthy"`
	PoolID      param.Field[string] `query:"pool_id"`
	// The name for the pool to filter.
	PoolName param.Field[string] `query:"pool_name"`
	// Start date and time of requesting data period in the ISO8601 format.
	Since param.Field[time.Time] `query:"since" format:"date-time"`
	// End date and time of requesting data period in the ISO8601 format.
	Until param.Field[time.Time] `query:"until" format:"date-time"`
}

// URLQuery serializes [LoadBalancerAnalyticsEventListParams]'s query parameters as
// `url.Values`.
func (r LoadBalancerAnalyticsEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type LoadBalancerAnalyticsEventListResponseEnvelope struct {
	Errors   []LoadBalancerAnalyticsEventListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerAnalyticsEventListResponseEnvelopeMessages `json:"messages,required"`
	Result   []LoadBalancingAnalytics                                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    LoadBalancerAnalyticsEventListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo LoadBalancerAnalyticsEventListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       loadBalancerAnalyticsEventListResponseEnvelopeJSON       `json:"-"`
}

// loadBalancerAnalyticsEventListResponseEnvelopeJSON contains the JSON metadata
// for the struct [LoadBalancerAnalyticsEventListResponseEnvelope]
type loadBalancerAnalyticsEventListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerAnalyticsEventListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerAnalyticsEventListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerAnalyticsEventListResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    loadBalancerAnalyticsEventListResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerAnalyticsEventListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [LoadBalancerAnalyticsEventListResponseEnvelopeErrors]
type loadBalancerAnalyticsEventListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerAnalyticsEventListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerAnalyticsEventListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type LoadBalancerAnalyticsEventListResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    loadBalancerAnalyticsEventListResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerAnalyticsEventListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [LoadBalancerAnalyticsEventListResponseEnvelopeMessages]
type loadBalancerAnalyticsEventListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerAnalyticsEventListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerAnalyticsEventListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerAnalyticsEventListResponseEnvelopeSuccess bool

const (
	LoadBalancerAnalyticsEventListResponseEnvelopeSuccessTrue LoadBalancerAnalyticsEventListResponseEnvelopeSuccess = true
)

type LoadBalancerAnalyticsEventListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       loadBalancerAnalyticsEventListResponseEnvelopeResultInfoJSON `json:"-"`
}

// loadBalancerAnalyticsEventListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [LoadBalancerAnalyticsEventListResponseEnvelopeResultInfo]
type loadBalancerAnalyticsEventListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerAnalyticsEventListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancerAnalyticsEventListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}
