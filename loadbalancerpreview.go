// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// LoadBalancerPreviewService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancerPreviewService]
// method instead.
type LoadBalancerPreviewService struct {
	Options []option.RequestOption
}

// NewLoadBalancerPreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancerPreviewService(opts ...option.RequestOption) (r *LoadBalancerPreviewService) {
	r = &LoadBalancerPreviewService{}
	r.Options = opts
	return
}

// Get the result of a previous preview operation using the provided preview_id.
func (r *LoadBalancerPreviewService) Get(ctx context.Context, previewID interface{}, query LoadBalancerPreviewGetParams, opts ...option.RequestOption) (res *LoadBalancingPreviewResult, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPreviewGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/load_balancers/preview/%v", query.AccountID, previewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancerPreviewGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type LoadBalancerPreviewGetResponseEnvelope struct {
	Errors   []LoadBalancerPreviewGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []LoadBalancerPreviewGetResponseEnvelopeMessages `json:"messages,required"`
	// Resulting health data from a preview operation.
	Result LoadBalancingPreviewResult `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancerPreviewGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancerPreviewGetResponseEnvelopeJSON    `json:"-"`
}

// loadBalancerPreviewGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancerPreviewGetResponseEnvelope]
type loadBalancerPreviewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPreviewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPreviewGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    loadBalancerPreviewGetResponseEnvelopeErrorsJSON `json:"-"`
}

// loadBalancerPreviewGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [LoadBalancerPreviewGetResponseEnvelopeErrors]
type loadBalancerPreviewGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPreviewGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPreviewGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    loadBalancerPreviewGetResponseEnvelopeMessagesJSON `json:"-"`
}

// loadBalancerPreviewGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [LoadBalancerPreviewGetResponseEnvelopeMessages]
type loadBalancerPreviewGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPreviewGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPreviewGetResponseEnvelopeSuccess bool

const (
	LoadBalancerPreviewGetResponseEnvelopeSuccessTrue LoadBalancerPreviewGetResponseEnvelopeSuccess = true
)
