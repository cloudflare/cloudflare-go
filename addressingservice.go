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

// AddressingServiceService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAddressingServiceService] method
// instead.
type AddressingServiceService struct {
	Options []option.RequestOption
}

// NewAddressingServiceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAddressingServiceService(opts ...option.RequestOption) (r *AddressingServiceService) {
	r = &AddressingServiceService{}
	r.Options = opts
	return
}

// Bring-Your-Own IP (BYOIP) prefixes onboarded to Cloudflare must be bound to a
// service running on the Cloudflare network to enable a Cloudflare product on the
// IP addresses. This endpoint can be used as a reference of available services on
// the Cloudflare network, and their service IDs.
func (r *AddressingServiceService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AddressingServiceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/services", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AddressingServiceListResponse struct {
	Errors   []AddressingServiceListResponseError   `json:"errors"`
	Messages []AddressingServiceListResponseMessage `json:"messages"`
	Result   []AddressingServiceListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AddressingServiceListResponseSuccess `json:"success"`
	JSON    addressingServiceListResponseJSON    `json:"-"`
}

// addressingServiceListResponseJSON contains the JSON metadata for the struct
// [AddressingServiceListResponse]
type addressingServiceListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingServiceListResponseError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    addressingServiceListResponseErrorJSON `json:"-"`
}

// addressingServiceListResponseErrorJSON contains the JSON metadata for the struct
// [AddressingServiceListResponseError]
type addressingServiceListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingServiceListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingServiceListResponseMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    addressingServiceListResponseMessageJSON `json:"-"`
}

// addressingServiceListResponseMessageJSON contains the JSON metadata for the
// struct [AddressingServiceListResponseMessage]
type addressingServiceListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingServiceListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AddressingServiceListResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Name of a service running on the Cloudflare network
	Name string                                  `json:"name"`
	JSON addressingServiceListResponseResultJSON `json:"-"`
}

// addressingServiceListResponseResultJSON contains the JSON metadata for the
// struct [AddressingServiceListResponseResult]
type addressingServiceListResponseResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressingServiceListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AddressingServiceListResponseSuccess bool

const (
	AddressingServiceListResponseSuccessTrue AddressingServiceListResponseSuccess = true
)
