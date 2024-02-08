// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
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
	var env LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/health", accountIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// A list of regions from which to run health checks. Null means every Cloudflare
// data center.
//
// Union satisfied by
// [LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseUnknown]
// or [shared.UnionString].
type LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponse interface {
	ImplementsLoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelope struct {
	Errors   []LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessages `json:"messages,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	Result LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelope]
type loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrors struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrors]
type loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessages struct {
	Code    int64                                                                                       `json:"code,required"`
	Message string                                                                                      `json:"message,required"`
	JSON    loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessages]
type loadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeSuccess bool

const (
	LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeSuccessTrue LoadBalancerPoolHealthAccountLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeSuccess = true
)
