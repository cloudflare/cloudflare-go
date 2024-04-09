// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// LoadBalancingPreviewService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewLoadBalancingPreviewService]
// method instead.
type LoadBalancingPreviewService struct {
	Options []option.RequestOption
}

// NewLoadBalancingPreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewLoadBalancingPreviewService(opts ...option.RequestOption) (r *LoadBalancingPreviewService) {
	r = &LoadBalancingPreviewService{}
	r.Options = opts
	return
}

// Get the result of a previous preview operation using the provided preview_id.
func (r *LoadBalancingPreviewService) Get(ctx context.Context, previewID string, opts ...option.RequestOption) (res *LoadBalancingPreviewGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancingPreviewGetResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/preview/%s", previewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancingPreviewGetResponse map[string]LoadBalancingPreviewGetResponse

type LoadBalancingPreviewGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Resulting health data from a preview operation.
	Result LoadBalancingPreviewGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success LoadBalancingPreviewGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    loadBalancingPreviewGetResponseEnvelopeJSON    `json:"-"`
}

// loadBalancingPreviewGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [LoadBalancingPreviewGetResponseEnvelope]
type loadBalancingPreviewGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPreviewGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPreviewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancingPreviewGetResponseEnvelopeSuccess bool

const (
	LoadBalancingPreviewGetResponseEnvelopeSuccessTrue LoadBalancingPreviewGetResponseEnvelopeSuccess = true
)

func (r LoadBalancingPreviewGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LoadBalancingPreviewGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
