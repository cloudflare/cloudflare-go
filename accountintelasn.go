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

// AccountIntelASNService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountIntelASNService] method
// instead.
type AccountIntelASNService struct {
	Options []option.RequestOption
	Subnets *AccountIntelASNSubnetService
}

// NewAccountIntelASNService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountIntelASNService(opts ...option.RequestOption) (r *AccountIntelASNService) {
	r = &AccountIntelASNService{}
	r.Options = opts
	r.Subnets = NewAccountIntelASNSubnetService(opts...)
	return
}

// Get ASN Overview
func (r *AccountIntelASNService) Get(ctx context.Context, accountIdentifier string, asn interface{}, opts ...option.RequestOption) (res *AccountIntelASNGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/asn/%v", accountIdentifier, asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type AccountIntelASNGetResponse struct {
	Errors   []AccountIntelASNGetResponseError   `json:"errors"`
	Messages []AccountIntelASNGetResponseMessage `json:"messages"`
	Result   int64                               `json:"result"`
	// Whether the API call was successful
	Success AccountIntelASNGetResponseSuccess `json:"success"`
	JSON    accountIntelASNGetResponseJSON    `json:"-"`
}

// accountIntelASNGetResponseJSON contains the JSON metadata for the struct
// [AccountIntelASNGetResponse]
type accountIntelASNGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelASNGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelASNGetResponseError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    accountIntelASNGetResponseErrorJSON `json:"-"`
}

// accountIntelASNGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountIntelASNGetResponseError]
type accountIntelASNGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelASNGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountIntelASNGetResponseMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    accountIntelASNGetResponseMessageJSON `json:"-"`
}

// accountIntelASNGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountIntelASNGetResponseMessage]
type accountIntelASNGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountIntelASNGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountIntelASNGetResponseSuccess bool

const (
	AccountIntelASNGetResponseSuccessTrue AccountIntelASNGetResponseSuccess = true
)
