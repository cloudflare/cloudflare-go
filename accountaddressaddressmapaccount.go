// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAddressAddressMapAccountService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAddressAddressMapAccountService] method instead.
type AccountAddressAddressMapAccountService struct {
	Options []option.RequestOption
}

// NewAccountAddressAddressMapAccountService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressAddressMapAccountService(opts ...option.RequestOption) (r *AccountAddressAddressMapAccountService) {
	r = &AccountAddressAddressMapAccountService{}
	r.Options = opts
	return
}

// Add an account as a member of a particular address map.
func (r *AccountAddressAddressMapAccountService) Update(ctx context.Context, accountIdentifier1 string, addressMapIdentifier string, accountIdentifier string, opts ...option.RequestOption) (res *APIResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", accountIdentifier1, addressMapIdentifier, accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Remove an account as a member of a particular address map.
func (r *AccountAddressAddressMapAccountService) Delete(ctx context.Context, accountIdentifier1 string, addressMapIdentifier string, accountIdentifier string, opts ...option.RequestOption) (res *APIResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/accounts/%s", accountIdentifier1, addressMapIdentifier, accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}
