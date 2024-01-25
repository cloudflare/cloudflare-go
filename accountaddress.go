// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAddressService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAddressService] method
// instead.
type AccountAddressService struct {
	Options      []option.RequestOption
	AddressMaps  *AccountAddressAddressMapService
	LoaDocuments *AccountAddressLoaDocumentService
	Prefixes     *AccountAddressPrefixService
}

// NewAccountAddressService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountAddressService(opts ...option.RequestOption) (r *AccountAddressService) {
	r = &AccountAddressService{}
	r.Options = opts
	r.AddressMaps = NewAccountAddressAddressMapService(opts...)
	r.LoaDocuments = NewAccountAddressLoaDocumentService(opts...)
	r.Prefixes = NewAccountAddressPrefixService(opts...)
	return
}
