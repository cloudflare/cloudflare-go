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

// AccountLoadBalancerPoolHealthService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountLoadBalancerPoolHealthService] method instead.
type AccountLoadBalancerPoolHealthService struct {
	Options []option.RequestOption
}

// NewAccountLoadBalancerPoolHealthService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerPoolHealthService(opts ...option.RequestOption) (r *AccountLoadBalancerPoolHealthService) {
	r = &AccountLoadBalancerPoolHealthService{}
	r.Options = opts
	return
}

// Fetch the latest pool health status for a single pool.
func (r *AccountLoadBalancerPoolHealthService) AccountLoadBalancerPoolsPoolHealthDetails(ctx context.Context, accountIdentifier string, identifier interface{}, opts ...option.RequestOption) (res *HealthDetail, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%v/health", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HealthDetail struct {
	Errors   []HealthDetailError   `json:"errors"`
	Messages []HealthDetailMessage `json:"messages"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success HealthDetailSuccess `json:"success"`
	JSON    healthDetailJSON    `json:"-"`
}

// healthDetailJSON contains the JSON metadata for the struct [HealthDetail]
type healthDetailJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthDetailError struct {
	Code    int64                 `json:"code,required"`
	Message string                `json:"message,required"`
	JSON    healthDetailErrorJSON `json:"-"`
}

// healthDetailErrorJSON contains the JSON metadata for the struct
// [HealthDetailError]
type healthDetailErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthDetailError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type HealthDetailMessage struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    healthDetailMessageJSON `json:"-"`
}

// healthDetailMessageJSON contains the JSON metadata for the struct
// [HealthDetailMessage]
type healthDetailMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthDetailMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type HealthDetailSuccess bool

const (
	HealthDetailSuccessTrue HealthDetailSuccess = true
)
