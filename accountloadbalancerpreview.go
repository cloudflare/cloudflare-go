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

// AccountLoadBalancerPreviewService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountLoadBalancerPreviewService] method instead.
type AccountLoadBalancerPreviewService struct {
	Options []option.RequestOption
}

// NewAccountLoadBalancerPreviewService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountLoadBalancerPreviewService(opts ...option.RequestOption) (r *AccountLoadBalancerPreviewService) {
	r = &AccountLoadBalancerPreviewService{}
	r.Options = opts
	return
}

// Get the result of a previous preview operation using the provided preview_id.
func (r *AccountLoadBalancerPreviewService) Get(ctx context.Context, accountIdentifier string, previewID interface{}, opts ...option.RequestOption) (res *AccountLoadBalancerPreviewGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/load_balancers/preview/%v", accountIdentifier, previewID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountLoadBalancerPreviewGetResponse struct {
	Errors   []AccountLoadBalancerPreviewGetResponseError   `json:"errors"`
	Messages []AccountLoadBalancerPreviewGetResponseMessage `json:"messages"`
	// Resulting health data from a preview operation.
	Result interface{} `json:"result"`
	// Whether the API call was successful
	Success AccountLoadBalancerPreviewGetResponseSuccess `json:"success"`
	JSON    accountLoadBalancerPreviewGetResponseJSON    `json:"-"`
}

// accountLoadBalancerPreviewGetResponseJSON contains the JSON metadata for the
// struct [AccountLoadBalancerPreviewGetResponse]
type accountLoadBalancerPreviewGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPreviewGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPreviewGetResponseError struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accountLoadBalancerPreviewGetResponseErrorJSON `json:"-"`
}

// accountLoadBalancerPreviewGetResponseErrorJSON contains the JSON metadata for
// the struct [AccountLoadBalancerPreviewGetResponseError]
type accountLoadBalancerPreviewGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPreviewGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountLoadBalancerPreviewGetResponseMessage struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accountLoadBalancerPreviewGetResponseMessageJSON `json:"-"`
}

// accountLoadBalancerPreviewGetResponseMessageJSON contains the JSON metadata for
// the struct [AccountLoadBalancerPreviewGetResponseMessage]
type accountLoadBalancerPreviewGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountLoadBalancerPreviewGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountLoadBalancerPreviewGetResponseSuccess bool

const (
	AccountLoadBalancerPreviewGetResponseSuccessTrue AccountLoadBalancerPreviewGetResponseSuccess = true
)
