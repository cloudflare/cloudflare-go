// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountLoadBalancerService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountLoadBalancerService]
// method instead.
type AccountLoadBalancerService struct {
	Options  []option.RequestOption
	Monitors *AccountLoadBalancerMonitorService
	Pools    *AccountLoadBalancerPoolService
	Previews *AccountLoadBalancerPreviewService
	Regions  *AccountLoadBalancerRegionService
	Searches *AccountLoadBalancerSearchService
}

// NewAccountLoadBalancerService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerService(opts ...option.RequestOption) (r *AccountLoadBalancerService) {
	r = &AccountLoadBalancerService{}
	r.Options = opts
	r.Monitors = NewAccountLoadBalancerMonitorService(opts...)
	r.Pools = NewAccountLoadBalancerPoolService(opts...)
	r.Previews = NewAccountLoadBalancerPreviewService(opts...)
	r.Regions = NewAccountLoadBalancerRegionService(opts...)
	r.Searches = NewAccountLoadBalancerSearchService(opts...)
	return
}
