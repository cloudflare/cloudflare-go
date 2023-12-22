// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAddressAddressMapZoneService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountAddressAddressMapZoneService] method instead.
type AccountAddressAddressMapZoneService struct {
	Options []option.RequestOption
}

// NewAccountAddressAddressMapZoneService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountAddressAddressMapZoneService(opts ...option.RequestOption) (r *AccountAddressAddressMapZoneService) {
	r = &AccountAddressAddressMapZoneService{}
	r.Options = opts
	return
}

// Add a zone as a member of a particular address map.
func (r *AccountAddressAddressMapZoneService) Update(ctx context.Context, accountIdentifier string, addressMapIdentifier string, zoneIdentifier string, opts ...option.RequestOption) (res *APIResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", accountIdentifier, addressMapIdentifier, zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return
}

// Remove a zone as a member of a particular address map.
func (r *AccountAddressAddressMapZoneService) Delete(ctx context.Context, accountIdentifier string, addressMapIdentifier string, zoneIdentifier string, opts ...option.RequestOption) (res *APIResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/addressing/address_maps/%s/zones/%s", accountIdentifier, addressMapIdentifier, zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}
