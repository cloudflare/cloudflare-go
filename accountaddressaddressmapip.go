// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAddressAddressMapIPService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountAddressAddressMapIPService] method instead.
type AccountAddressAddressMapIPService struct {
	Options []option.RequestOption
}

// NewAccountAddressAddressMapIPService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressAddressMapIPService(opts ...option.RequestOption) (r *AccountAddressAddressMapIPService) {
	r = &AccountAddressAddressMapIPService{}
	r.Options = opts
	return
}

// Add an IP from a prefix owned by the account to a particular address map.
func (r *AccountAddressAddressMapIPService) Update(ctx context.Context, accountIdentifier string, addressMapIdentifier string, ipAddress string, opts ...option.RequestOption) (res *APIResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", accountIdentifier, addressMapIdentifier, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Remove an IP from a particular address map.
func (r *AccountAddressAddressMapIPService) Delete(ctx context.Context, accountIdentifier string, addressMapIdentifier string, ipAddress string, opts ...option.RequestOption) (res *APIResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/ips/%s", accountIdentifier, addressMapIdentifier, ipAddress)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}
