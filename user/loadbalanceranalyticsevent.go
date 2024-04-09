// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *LoadBalancerAnalyticsEventService) List(ctx context.Context, query LoadBalancerAnalyticsEventListParams, opts ...option.RequestOption) (res *pagination.SinglePage[Analytics], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "user/load_balancing_analytics/events"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List origin health changes.
func (r *LoadBalancerAnalyticsEventService) ListAutoPaging(ctx context.Context, query LoadBalancerAnalyticsEventListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Analytics] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
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
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
