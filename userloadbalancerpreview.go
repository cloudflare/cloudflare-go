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

// UserLoadBalancerPreviewService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewUserLoadBalancerPreviewService] method instead.
type UserLoadBalancerPreviewService struct {
	Options []option.RequestOption
}

// NewUserLoadBalancerPreviewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserLoadBalancerPreviewService(opts ...option.RequestOption) (r *UserLoadBalancerPreviewService) {
	r = &UserLoadBalancerPreviewService{}
	r.Options = opts
	return
}

// Get the result of a previous preview operation using the provided preview_id.
func (r *UserLoadBalancerPreviewService) Get(ctx context.Context, previewID interface{}, opts ...option.RequestOption) (res *UserLoadBalancerPreviewGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("user/load_balancers/preview/%v", previewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type UserLoadBalancerPreviewGetResponse struct {
	Errors   []UserLoadBalancerPreviewGetResponseError   `json:"errors"`
	Messages []UserLoadBalancerPreviewGetResponseMessage `json:"messages"`
	// Resulting health data from a preview operation.
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success UserLoadBalancerPreviewGetResponseSuccess `json:"success"`
	JSON    userLoadBalancerPreviewGetResponseJSON    `json:"-"`
}

// userLoadBalancerPreviewGetResponseJSON contains the JSON metadata for the struct
// [UserLoadBalancerPreviewGetResponse]
type userLoadBalancerPreviewGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPreviewGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPreviewGetResponseError struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    userLoadBalancerPreviewGetResponseErrorJSON `json:"-"`
}

// userLoadBalancerPreviewGetResponseErrorJSON contains the JSON metadata for the
// struct [UserLoadBalancerPreviewGetResponseError]
type userLoadBalancerPreviewGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPreviewGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserLoadBalancerPreviewGetResponseMessage struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    userLoadBalancerPreviewGetResponseMessageJSON `json:"-"`
}

// userLoadBalancerPreviewGetResponseMessageJSON contains the JSON metadata for the
// struct [UserLoadBalancerPreviewGetResponseMessage]
type userLoadBalancerPreviewGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserLoadBalancerPreviewGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserLoadBalancerPreviewGetResponseSuccess bool

const (
	UserLoadBalancerPreviewGetResponseSuccessTrue UserLoadBalancerPreviewGetResponseSuccess = true
)
