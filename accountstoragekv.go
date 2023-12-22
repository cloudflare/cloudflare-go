// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountStorageKvService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountStorageKvService] method
// instead.
type AccountStorageKvService struct {
	Options    []option.RequestOption
	Namespaces *AccountStorageKvNamespaceService
}

// NewAccountStorageKvService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountStorageKvService(opts ...option.RequestOption) (r *AccountStorageKvService) {
	r = &AccountStorageKvService{}
	r.Options = opts
	r.Namespaces = NewAccountStorageKvNamespaceService(opts...)
	return
}
