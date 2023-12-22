// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserLoadBalancerPoolReferenceService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewUserLoadBalancerPoolReferenceService] method instead.
type UserLoadBalancerPoolReferenceService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerPoolReferenceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserLoadBalancerPoolReferenceService(opts ...option.RequestOption) (r *UserLoadBalancerPoolReferenceService) {
	r = &UserLoadBalancerPoolReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided pool.
func (r *UserLoadBalancerPoolReferenceService) LoadBalancerPoolsListPoolReferences(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *SchemasReferencesResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/pools/%v/references", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
