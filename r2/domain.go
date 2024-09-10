// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package r2

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// DomainService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainService] method instead.
type DomainService struct {
	Options []option.RequestOption
	Custom  *DomainCustomService
	Managed *DomainManagedService
}

// NewDomainService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDomainService(opts ...option.RequestOption) (r *DomainService) {
	r = &DomainService{}
	r.Options = opts
	r.Custom = NewDomainCustomService(opts...)
	r.Managed = NewDomainManagedService(opts...)
	return
}
