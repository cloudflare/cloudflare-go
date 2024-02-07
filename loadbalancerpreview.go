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
func (r *LoadBalancerPreviewService) Get(ctx context.Context, accountIdentifier string, previewID interface{}, opts ...option.RequestOption) (res *LoadBalancerPreviewGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/preview/%v", accountIdentifier, previewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LoadBalancerPreviewGetResponse struct {
	Errors   []LoadBalancerPreviewGetResponseError   `json:"errors"`
	Messages []LoadBalancerPreviewGetResponseMessage `json:"messages"`
	// Resulting health data from a preview operation.
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success LoadBalancerPreviewGetResponseSuccess `json:"success"`
	JSON    loadBalancerPreviewGetResponseJSON    `json:"-"`
}

// loadBalancerPreviewGetResponseJSON contains the JSON metadata for the struct
// [LoadBalancerPreviewGetResponse]
type loadBalancerPreviewGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPreviewGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPreviewGetResponseError struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    loadBalancerPreviewGetResponseErrorJSON `json:"-"`
}

// loadBalancerPreviewGetResponseErrorJSON contains the JSON metadata for the
// struct [LoadBalancerPreviewGetResponseError]
type loadBalancerPreviewGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPreviewGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type LoadBalancerPreviewGetResponseMessage struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    loadBalancerPreviewGetResponseMessageJSON `json:"-"`
}

// loadBalancerPreviewGetResponseMessageJSON contains the JSON metadata for the
// struct [LoadBalancerPreviewGetResponseMessage]
type loadBalancerPreviewGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LoadBalancerPreviewGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type LoadBalancerPreviewGetResponseSuccess bool

const (
	LoadBalancerPreviewGetResponseSuccessTrue LoadBalancerPreviewGetResponseSuccess = true
)
