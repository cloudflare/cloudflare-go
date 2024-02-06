// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// LoadBalancerPoolHealthService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerPoolHealthService]
// method instead.
type LoadBalancerPoolHealthService struct {
	Options []option.RequestOption
}

// NewLoadBalancerPoolHealthService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerPoolHealthService(opts ...option.RequestOption) (r *LoadBalancerPoolHealthService) {
	r = &LoadBalancerPoolHealthService{}
	r.Options = opts
	return
}

// Fetch the latest pool health status for a single pool.
func (r *LoadBalancerPoolHealthService) AccountLoadBalancerPoolsPoolHealthDetails(ctx context.Context, accountIdentifier string, identifier string, opts ...option.RequestOption) (res *LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/health", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponse struct {
	Errors   []LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseError   `json:"errors"`
	Messages []LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseMessage `json:"messages"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseSuccess `json:"success"`
	JSON    loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseJSON    `json:"-"`
}

// loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponse]
type loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseError struct {
	Code    int64                                                                            `json:"code,required"`
	Message string                                                                           `json:"message,required"`
	JSON    loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseErrorJSON `json:"-"`
}

// loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseErrorJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseError]
type loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseMessage struct {
	Code    int64                                                                              `json:"code,required"`
	Message string                                                                             `json:"message,required"`
	JSON    loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseMessageJSON `json:"-"`
}

// loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseMessageJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseMessage]
type loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseSuccess bool

const (
	LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseSuccessTrue LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseSuccess = true
)
