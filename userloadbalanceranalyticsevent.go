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

// UserLoadBalancerAnalyticsEventService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewUserLoadBalancerAnalyticsEventService] method instead.
type UserLoadBalancerAnalyticsEventService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerAnalyticsEventService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserLoadBalancerAnalyticsEventService(opts ...option.RequestOption) (r *UserLoadBalancerAnalyticsEventService) {
	r = &UserLoadBalancerAnalyticsEventService{}
	r.Options = opts
	return
}

// List origin health changes.
func (r *UserLoadBalancerAnalyticsEventService) List(ctx context.Context, query UserLoadBalancerAnalyticsEventListParams, opts ...option.RequestOption) (res *[]UserLoadBalancerAnalyticsEventListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerAnalyticsEventListResponseEnvelope
	path := "user/load_balancing_analytics/events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserLoadBalancerAnalyticsEventListResponse struct {
	ID        int64                                          `json:"id"`
	Origins   []interface{}                                  `json:"origins"`
	Pool      interface{}                                    `json:"pool"`
	Timestamp time.Time                                      `json:"timestamp" format:"date-time"`
	JSON      userLoadBalancerAnalyticsEventListResponseJSON `json:"-"`
}

// userLoadBalancerAnalyticsEventListResponseJSON contains the JSON metadata for
// the struct [UserLoadBalancerAnalyticsEventListResponse]
type userLoadBalancerAnalyticsEventListResponseJSON struct {
	ID          apijson.Field
	Origins     apijson.Field
	Pool        apijson.Field
	Timestamp   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerAnalyticsEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerAnalyticsEventListParams struct {
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

// URLQuery serializes [UserLoadBalancerAnalyticsEventListParams]'s query
// parameters as `url.Values`.
func (r UserLoadBalancerAnalyticsEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserLoadBalancerAnalyticsEventListResponseEnvelope struct {
	Errors   []UserLoadBalancerAnalyticsEventListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerAnalyticsEventListResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserLoadBalancerAnalyticsEventListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserLoadBalancerAnalyticsEventListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserLoadBalancerAnalyticsEventListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userLoadBalancerAnalyticsEventListResponseEnvelopeJSON       `json:"-"`
}

// userLoadBalancerAnalyticsEventListResponseEnvelopeJSON contains the JSON
// metadata for the struct [UserLoadBalancerAnalyticsEventListResponseEnvelope]
type userLoadBalancerAnalyticsEventListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerAnalyticsEventListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerAnalyticsEventListResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    userLoadBalancerAnalyticsEventListResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerAnalyticsEventListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerAnalyticsEventListResponseEnvelopeErrors]
type userLoadBalancerAnalyticsEventListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerAnalyticsEventListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerAnalyticsEventListResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    userLoadBalancerAnalyticsEventListResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerAnalyticsEventListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [UserLoadBalancerAnalyticsEventListResponseEnvelopeMessages]
type userLoadBalancerAnalyticsEventListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerAnalyticsEventListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerAnalyticsEventListResponseEnvelopeSuccess bool

const (
	UserLoadBalancerAnalyticsEventListResponseEnvelopeSuccessTrue UserLoadBalancerAnalyticsEventListResponseEnvelopeSuccess = true
)

type UserLoadBalancerAnalyticsEventListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                          `json:"total_count"`
	JSON       userLoadBalancerAnalyticsEventListResponseEnvelopeResultInfoJSON `json:"-"`
}

// userLoadBalancerAnalyticsEventListResponseEnvelopeResultInfoJSON contains the
// JSON metadata for the struct
// [UserLoadBalancerAnalyticsEventListResponseEnvelopeResultInfo]
type userLoadBalancerAnalyticsEventListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerAnalyticsEventListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
