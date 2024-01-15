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

// AccountAddressingServiceService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountAddressingServiceService] method instead.
type AccountAddressingServiceService struct {
	Options []option.RequestOption
}

// NewAccountAddressingServiceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressingServiceService(opts ...option.RequestOption) (r *AccountAddressingServiceService) {
	r = &AccountAddressingServiceService{}
	r.Options = opts
	return
}

// Bring-Your-Own IP (BYOIP) prefixes onboarded to Cloudflare must be bound to a
// service running on the Cloudflare network to enable a Cloudflare product on the
// IP addresses. This endpoint can be used as a reference of available services on
// the Cloudflare network, and their service IDs.
func (r *AccountAddressingServiceService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *AccountAddressingServiceListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/services", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountAddressingServiceListResponse struct {
	Errors   []AccountAddressingServiceListResponseError   `json:"errors"`
	Messages []AccountAddressingServiceListResponseMessage `json:"messages"`
	Result   []AccountAddressingServiceListResponseResult  `json:"result"`
	// Whether the API call was successful
	Success AccountAddressingServiceListResponseSuccess `json:"success"`
	JSON    accountAddressingServiceListResponseJSON    `json:"-"`
}

// accountAddressingServiceListResponseJSON contains the JSON metadata for the
// struct [AccountAddressingServiceListResponse]
type accountAddressingServiceListResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingServiceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingServiceListResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accountAddressingServiceListResponseErrorJSON `json:"-"`
}

// accountAddressingServiceListResponseErrorJSON contains the JSON metadata for the
// struct [AccountAddressingServiceListResponseError]
type accountAddressingServiceListResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingServiceListResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingServiceListResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    accountAddressingServiceListResponseMessageJSON `json:"-"`
}

// accountAddressingServiceListResponseMessageJSON contains the JSON metadata for
// the struct [AccountAddressingServiceListResponseMessage]
type accountAddressingServiceListResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingServiceListResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountAddressingServiceListResponseResult struct {
	// Identifier
	ID string `json:"id"`
	// Name of a service running on the Cloudflare network
	Name string                                         `json:"name"`
	JSON accountAddressingServiceListResponseResultJSON `json:"-"`
}

// accountAddressingServiceListResponseResultJSON contains the JSON metadata for
// the struct [AccountAddressingServiceListResponseResult]
type accountAddressingServiceListResponseResultJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountAddressingServiceListResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountAddressingServiceListResponseSuccess bool

const (
	AccountAddressingServiceListResponseSuccessTrue AccountAddressingServiceListResponseSuccess = true
)
