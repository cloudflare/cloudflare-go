// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserLoadBalancerPoolHealthService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewUserLoadBalancerPoolHealthService] method instead.
type UserLoadBalancerPoolHealthService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerPoolHealthService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserLoadBalancerPoolHealthService(opts ...option.RequestOption) (r *UserLoadBalancerPoolHealthService) {
	r = &UserLoadBalancerPoolHealthService{}
	r.Options = opts
	return
}

// Fetch the latest pool health status for a single pool.
func (r *UserLoadBalancerPoolHealthService) LoadBalancerPoolsPoolHealthDetails(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *HealthDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/pools/%v/health", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
