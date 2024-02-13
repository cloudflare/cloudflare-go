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
	var env UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s/health", identifier)
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
// [UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseUnknown] or
// [shared.UnionString].
type UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponse interface {
	ImplementsUserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessages `json:"messages,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	Result UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponse `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelope]
type userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrors struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrors]
type userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessages struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessages]
type userLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeSuccessTrue UserLoadBalancerPoolHealthLoadBalancerPoolsPoolHealthDetailsResponseEnvelopeSuccess = true
)
