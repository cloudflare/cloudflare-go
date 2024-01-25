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
func (r *UserLoadBalancerPoolHealthService) LoadBalancerPoolsPoolHealthDetails(ctx context.Context, identifier string, opts ...option.RequestOption) (res *UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/pools/%s/health", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponse struct {
	Errors   []UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseError   `json:"errors"`
	Messages []UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseMessage `json:"messages"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseSuccess `json:"success"`
	JSON    userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseJSON    `json:"-"`
}

// userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponse]
type userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseError struct {
	Code    int64                                                                         `json:"code,required"`
	Message string                                                                        `json:"message,required"`
	JSON    userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseErrorJSON `json:"-"`
}

// userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseErrorJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseError]
type userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseMessage struct {
	Code    int64                                                                           `json:"code,required"`
	Message string                                                                          `json:"message,required"`
	JSON    userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseMessageJSON `json:"-"`
}

// userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseMessageJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseMessage]
type userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseSuccess bool

const (
	UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseSuccessTrue UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseSuccess = true
)
