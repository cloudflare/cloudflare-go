// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserLoadBalancerMonitorReferenceService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewUserLoadBalancerMonitorReferenceService] method instead.
type UserLoadBalancerMonitorReferenceService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerMonitorReferenceService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserLoadBalancerMonitorReferenceService(opts ...option.RequestOption) (r *UserLoadBalancerMonitorReferenceService) {
	r = &UserLoadBalancerMonitorReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided monitor.
func (r *UserLoadBalancerMonitorReferenceService) LoadBalancerMonitorsListMonitorReferences(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *ReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/monitors/%v/references", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
