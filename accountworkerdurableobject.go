// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountWorkerDurableObjectService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountWorkerDurableObjectService] method instead.
type AccountWorkerDurableObjectService struct {
	Options    []option.RequestOption
	Namespaces *AccountWorkerDurableObjectNamespaceService
}

// NewAccountWorkerDurableObjectService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerDurableObjectService(opts ...option.RequestOption) (r *AccountWorkerDurableObjectService) {
	r = &AccountWorkerDurableObjectService{}
	r.Options = opts
	r.Namespaces = NewAccountWorkerDurableObjectNamespaceService(opts...)
	return
}
