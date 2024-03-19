// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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
func (r *LoadBalancerPreviewService) Get(ctx context.Context, previewID interface{}, opts ...option.RequestOption) (res *LoadBalancingPreviewResult, err error) {
	opts = append(r.Options[:], opts...)
	var env LoadBalancerPreviewGetResponseEnvelope
	path := fmt.Sprintf("user/load_balancers/preview/%v", previewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LoadBalancingPreviewResult map[string]LoadBalancingPreviewResultItem

type LoadBalancingPreviewResultItem struct {
	Healthy bool                                          `json:"healthy"`
	Origins []map[string]LoadBalancingPreviewResultOrigin `json:"origins"`
	JSON    loadBalancingPreviewResultItemJSON            `json:"-"`
}

// loadBalancingPreviewResultItemJSON contains the JSON metadata for the struct
// [LoadBalancingPreviewResultItem]
type loadBalancingPreviewResultItemJSON struct {
	Healthy     apijson.Field
	Origins     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancingPreviewResultItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPreviewResultItemJSON) RawJSON() string {
	return r.raw
}

// The origin ipv4/ipv6 address or domain name mapped to it's health data.
type LoadBalancingPreviewResultOrigin struct {
	FailureReason string                               `json:"failure_reason"`
	Healthy       bool                                 `json:"healthy"`
	ResponseCode  float64                              `json:"response_code"`
	RTT           string                               `json:"rtt"`
	JSON          loadBalancingPreviewResultOriginJSON `json:"-"`
}

// loadBalancingPreviewResultOriginJSON contains the JSON metadata for the struct
// [LoadBalancingPreviewResultOrigin]
type loadBalancingPreviewResultOriginJSON struct {
	FailureReason apijson.Field
	Healthy       apijson.Field
	ResponseCode  apijson.Field
	RTT           apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *LoadBalancingPreviewResultOrigin) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r loadBalancingPreviewResultOriginJSON) RawJSON() string {
	return r.raw
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

func (r loadBalancerPreviewGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r loadBalancerPreviewGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r loadBalancerPreviewGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type LoadBalancerPreviewGetResponseEnvelopeSuccess bool

const (
	LoadBalancerPreviewGetResponseEnvelopeSuccessTrue LoadBalancerPreviewGetResponseEnvelopeSuccess = true
)
