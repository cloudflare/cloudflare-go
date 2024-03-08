// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// UserLoadBalancerService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewUserLoadBalancerService] method
// instead.
type UserLoadBalancerService struct {
	Options   []option.RequestOption
	Monitors  *UserLoadBalancerMonitorService
	Pools     *UserLoadBalancerPoolService
	Preview   *UserLoadBalancerPreviewService
	Analytics *UserLoadBalancerAnalyticsService
}

// NewUserLoadBalancerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserLoadBalancerService(opts ...option.RequestOption) (r *UserLoadBalancerService) {
	r = &UserLoadBalancerService{}
	r.Options = opts
	r.Monitors = NewUserLoadBalancerMonitorService(opts...)
	r.Pools = NewUserLoadBalancerPoolService(opts...)
	r.Preview = NewUserLoadBalancerPreviewService(opts...)
	r.Analytics = NewUserLoadBalancerAnalyticsService(opts...)
	return
}
