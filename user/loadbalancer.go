// File generated from our OpenAPI spec by Stainless.

package user

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// LoadBalancerService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerService] method
// instead.
type LoadBalancerService struct {
	Options   []option.RequestOption
	Monitors  *LoadBalancerMonitorService
	Pools     *LoadBalancerPoolService
	Preview   *LoadBalancerPreviewService
	Analytics *LoadBalancerAnalyticsService
}

// NewLoadBalancerService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLoadBalancerService(opts ...option.RequestOption) (r *LoadBalancerService) {
	r = &LoadBalancerService{}
	r.Options = opts
	r.Monitors = NewLoadBalancerMonitorService(opts...)
	r.Pools = NewLoadBalancerPoolService(opts...)
	r.Preview = NewLoadBalancerPreviewService(opts...)
	r.Analytics = NewLoadBalancerAnalyticsService(opts...)
	return
}
