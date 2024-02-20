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
func (r *UserLoadBalancerPoolHealthService) List(ctx context.Context, poolID string, opts ...option.RequestOption) (res *UserLoadBalancerPoolHealthListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserLoadBalancerPoolHealthListResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/pools/%s/health", poolID)
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
// Union satisfied by [UserLoadBalancerPoolHealthListResponseUnknown] or
// [shared.UnionString].
type UserLoadBalancerPoolHealthListResponse interface {
	ImplementsUserLoadBalancerPoolHealthListResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*UserLoadBalancerPoolHealthListResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type UserLoadBalancerPoolHealthListResponseEnvelope struct {
	Errors   []UserLoadBalancerPoolHealthListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserLoadBalancerPoolHealthListResponseEnvelopeMessages `json:"messages,required"`
	// A list of regions from which to run health checks. Null means every Cloudflare
	// data center.
	Result UserLoadBalancerPoolHealthListResponse `json:"result,required"`
	// Whether the API call was successful
	Success UserLoadBalancerPoolHealthListResponseEnvelopeSuccess `json:"success,required"`
	JSON    userLoadBalancerPoolHealthListResponseEnvelopeJSON    `json:"-"`
}

// userLoadBalancerPoolHealthListResponseEnvelopeJSON contains the JSON metadata
// for the struct [UserLoadBalancerPoolHealthListResponseEnvelope]
type userLoadBalancerPoolHealthListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolHealthListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolHealthListResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    userLoadBalancerPoolHealthListResponseEnvelopeErrorsJSON `json:"-"`
}

// userLoadBalancerPoolHealthListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolHealthListResponseEnvelopeErrors]
type userLoadBalancerPoolHealthListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolHealthListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPoolHealthListResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    userLoadBalancerPoolHealthListResponseEnvelopeMessagesJSON `json:"-"`
}

// userLoadBalancerPoolHealthListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserLoadBalancerPoolHealthListResponseEnvelopeMessages]
type userLoadBalancerPoolHealthListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPoolHealthListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPoolHealthListResponseEnvelopeSuccess bool

const (
	UserLoadBalancerPoolHealthListResponseEnvelopeSuccessTrue UserLoadBalancerPoolHealthListResponseEnvelopeSuccess = true
)
