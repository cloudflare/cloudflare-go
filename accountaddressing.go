// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountAddressingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountAddressingService] method
// instead.
type AccountAddressingService struct {
	Options  []option.RequestOption
	Prefixes *AccountAddressingPrefixService
	Services *AccountAddressingServiceService
}

// NewAccountAddressingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAddressingService(opts ...option.RequestOption) (r *AccountAddressingService) {
	r = &AccountAddressingService{}
	r.Options = opts
	r.Prefixes = NewAccountAddressingPrefixService(opts...)
	r.Services = NewAccountAddressingServiceService(opts...)
	return
}
